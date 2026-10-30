// generated code - do not edit
package gml

// insertion point
func (inst *A_Row) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "A_rows":
		switch reverseField.Fieldname {
		case "Row":
			if _a_rows, ok := stage.A_rows_Row_reverseMap[inst]; ok {
				res = _a_rows.Name
			}
		}
	}
	return
}

func (inst *A_controlPoint) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *A_refLocation) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *A_rows) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractContinuousCoverageType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractCoordinateOperationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractCoordinateSystemType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractCoverageType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractCurveSegmentType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDQ_Completeness_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDQ_Element_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDQ_LogicalConsistency_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDQ_PositionalAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDQ_Result_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDQ_TemporalAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDQ_ThematicAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractDatumType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractEX_GeographicExtent_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractFeatureCollectionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractFeatureMemberType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractFeatureType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGMLType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeneralConversionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeneralDerivedCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeneralOperationParameterPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "OperationMethodType":
		switch reverseField.Fieldname {
		case "Parameter":
			if _operationmethodtype, ok := stage.OperationMethodType_Parameter_reverseMap[inst]; ok {
				res = _operationmethodtype.Name
			}
		case "IncludesParameter":
			if _operationmethodtype, ok := stage.OperationMethodType_IncludesParameter_reverseMap[inst]; ok {
				res = _operationmethodtype.Name
			}
		case "GeneralOperationParameter":
			if _operationmethodtype, ok := stage.OperationMethodType_GeneralOperationParameter_reverseMap[inst]; ok {
				res = _operationmethodtype.Name
			}
		case "UsesParameter":
			if _operationmethodtype, ok := stage.OperationMethodType_UsesParameter_reverseMap[inst]; ok {
				res = _operationmethodtype.Name
			}
		}
	case "OperationParameterGroupType":
		switch reverseField.Fieldname {
		case "Parameter":
			if _operationparametergrouptype, ok := stage.OperationParameterGroupType_Parameter_reverseMap[inst]; ok {
				res = _operationparametergrouptype.Name
			}
		case "IncludesParameter":
			if _operationparametergrouptype, ok := stage.OperationParameterGroupType_IncludesParameter_reverseMap[inst]; ok {
				res = _operationparametergrouptype.Name
			}
		case "GeneralOperationParameter":
			if _operationparametergrouptype, ok := stage.OperationParameterGroupType_GeneralOperationParameter_reverseMap[inst]; ok {
				res = _operationparametergrouptype.Name
			}
		case "UsesParameter":
			if _operationparametergrouptype, ok := stage.OperationParameterGroupType_UsesParameter_reverseMap[inst]; ok {
				res = _operationparametergrouptype.Name
			}
		}
	}
	return
}

func (inst *AbstractGeneralOperationParameterType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeneralParameterValuePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ConversionType":
		switch reverseField.Fieldname {
		case "ParameterValue":
			if _conversiontype, ok := stage.ConversionType_ParameterValue_reverseMap[inst]; ok {
				res = _conversiontype.Name
			}
		case "UsesValue":
			if _conversiontype, ok := stage.ConversionType_UsesValue_reverseMap[inst]; ok {
				res = _conversiontype.Name
			}
		case "IncludesValue":
			if _conversiontype, ok := stage.ConversionType_IncludesValue_reverseMap[inst]; ok {
				res = _conversiontype.Name
			}
		}
	case "ParameterValueGroupType":
		switch reverseField.Fieldname {
		case "ParameterValue":
			if _parametervaluegrouptype, ok := stage.ParameterValueGroupType_ParameterValue_reverseMap[inst]; ok {
				res = _parametervaluegrouptype.Name
			}
		case "UsesValue":
			if _parametervaluegrouptype, ok := stage.ParameterValueGroupType_UsesValue_reverseMap[inst]; ok {
				res = _parametervaluegrouptype.Name
			}
		case "IncludesValue":
			if _parametervaluegrouptype, ok := stage.ParameterValueGroupType_IncludesValue_reverseMap[inst]; ok {
				res = _parametervaluegrouptype.Name
			}
		}
	case "TransformationType":
		switch reverseField.Fieldname {
		case "ParameterValue":
			if _transformationtype, ok := stage.TransformationType_ParameterValue_reverseMap[inst]; ok {
				res = _transformationtype.Name
			}
		case "UsesValue":
			if _transformationtype, ok := stage.TransformationType_UsesValue_reverseMap[inst]; ok {
				res = _transformationtype.Name
			}
		case "IncludesValue":
			if _transformationtype, ok := stage.TransformationType_IncludesValue_reverseMap[inst]; ok {
				res = _transformationtype.Name
			}
		}
	}
	return
}

func (inst *AbstractGeneralParameterValueType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeneralTransformationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeometricAggregateType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeometricPrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGeometryType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractGriddedSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractMD_Identification_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractMemberType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractMetaDataType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractMetadataPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractParametricCurveSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractRS_ReferenceSystem_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractRingPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "PolygonPatchType":
		switch reverseField.Fieldname {
		case "Interior":
			if _polygonpatchtype, ok := stage.PolygonPatchType_Interior_reverseMap[inst]; ok {
				res = _polygonpatchtype.Name
			}
		}
	case "PolygonType":
		switch reverseField.Fieldname {
		case "Interior":
			if _polygontype, ok := stage.PolygonType_Interior_reverseMap[inst]; ok {
				res = _polygontype.Name
			}
		}
	}
	return
}

func (inst *AbstractRingType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractSolidType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractSurfacePatchType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeComplexType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeGeometricPrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeObjectType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimePrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeSliceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTimeTopologyPrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTopoPrimitiveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AbstractTopologyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AffineCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AffineCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "AffineCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_AffineCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *AffinePlacementType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "AffinePlacement":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_AffinePlacement_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *AngleChoiceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Angle_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ArcByBulgeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ArcByBulge":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ArcByBulge_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "ArcByBulge":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_ArcByBulge_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ArcByCenterPointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ArcByCenterPoint":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ArcByCenterPoint_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "ArcByCenterPoint":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_ArcByCenterPoint_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ArcStringByBulgeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ArcStringByBulge":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ArcStringByBulge_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "ArcStringByBulge":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_ArcStringByBulge_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ArcStringType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ArcString":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ArcString_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "ArcString":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_ArcString_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ArcType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Arc":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Arc_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Arc":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Arc_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ArrayAssociationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ArrayType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Array":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Array_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *AssociationRoleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "BagType":
		switch reverseField.Fieldname {
		case "Member":
			if _bagtype, ok := stage.BagType_Member_reverseMap[inst]; ok {
				res = _bagtype.Name
			}
		}
	}
	return
}

func (inst *BSplineType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "BSpline":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_BSpline_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "BSpline":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_BSpline_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *BagType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Bag":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Bag_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *BaseUnitType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "BaseUnit":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_BaseUnit_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *BezierType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Bezier":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Bezier_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Bezier":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Bezier_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *Binary_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Binary_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Boolean) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Boolean":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Boolean_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *BooleanPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Boolean_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *BoundedFeatureType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *BoundingShapeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Address_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Address_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Citation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Citation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Contact_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Contact_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_DateTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Date_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CI_Citation_Type":
		switch reverseField.Fieldname {
		case "Date":
			if _ci_citation_type, ok := stage.CI_Citation_Type_Date_reverseMap[inst]; ok {
				res = _ci_citation_type.Name
			}
		}
	case "PT_LocaleContainer_Type":
		switch reverseField.Fieldname {
		case "Date":
			if _pt_localecontainer_type, ok := stage.PT_LocaleContainer_Type_Date_reverseMap[inst]; ok {
				res = _pt_localecontainer_type.Name
			}
		}
	}
	return
}

func (inst *CI_Date_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_OnLineFunctionCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_OnlineResource_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_DigitalTransferOptions_Type":
		switch reverseField.Fieldname {
		case "OnLine":
			if _md_digitaltransferoptions_type, ok := stage.MD_DigitalTransferOptions_Type_OnLine_reverseMap[inst]; ok {
				res = _md_digitaltransferoptions_type.Name
			}
		}
	}
	return
}

func (inst *CI_OnlineResource_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_PresentationFormCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CI_Citation_Type":
		switch reverseField.Fieldname {
		case "PresentationForm":
			if _ci_citation_type, ok := stage.CI_Citation_Type_PresentationForm_reverseMap[inst]; ok {
				res = _ci_citation_type.Name
			}
		}
	}
	return
}

func (inst *CI_ResponsibleParty_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "PointOfContact":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_PointOfContact_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "CI_Citation_Type":
		switch reverseField.Fieldname {
		case "CitedResponsibleParty":
			if _ci_citation_type, ok := stage.CI_Citation_Type_CitedResponsibleParty_reverseMap[inst]; ok {
				res = _ci_citation_type.Name
			}
		}
	case "LI_ProcessStep_Type":
		switch reverseField.Fieldname {
		case "Processor":
			if _li_processstep_type, ok := stage.LI_ProcessStep_Type_Processor_reverseMap[inst]; ok {
				res = _li_processstep_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "PointOfContact":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_PointOfContact_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_MaintenanceInformation_Type":
		switch reverseField.Fieldname {
		case "Contact":
			if _md_maintenanceinformation_type, ok := stage.MD_MaintenanceInformation_Type_Contact_reverseMap[inst]; ok {
				res = _md_maintenanceinformation_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "PointOfContact":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_PointOfContact_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	case "MD_Usage_Type":
		switch reverseField.Fieldname {
		case "UserContactInfo":
			if _md_usage_type, ok := stage.MD_Usage_Type_UserContactInfo_reverseMap[inst]; ok {
				res = _md_usage_type.Name
			}
		}
	case "PT_LocaleContainer_Type":
		switch reverseField.Fieldname {
		case "ResponsibleParty":
			if _pt_localecontainer_type, ok := stage.PT_LocaleContainer_Type_ResponsibleParty_reverseMap[inst]; ok {
				res = _pt_localecontainer_type.Name
			}
		}
	}
	return
}

func (inst *CI_ResponsibleParty_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_RoleCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Series_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Series_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Telephone_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CI_Telephone_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CartesianCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CartesianCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CartesianCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CartesianCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *Category) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Category":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Category_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *CategoryExtentType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CategoryExtent":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CategoryExtent_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *CategoryPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CharacterString_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractDQ_Completeness_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _abstractdq_completeness_type, ok := stage.AbstractDQ_Completeness_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _abstractdq_completeness_type.Name
			}
		}
	case "AbstractDQ_Element_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _abstractdq_element_type, ok := stage.AbstractDQ_Element_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _abstractdq_element_type.Name
			}
		}
	case "AbstractDQ_LogicalConsistency_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _abstractdq_logicalconsistency_type, ok := stage.AbstractDQ_LogicalConsistency_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _abstractdq_logicalconsistency_type.Name
			}
		}
	case "AbstractDQ_PositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _abstractdq_positionalaccuracy_type, ok := stage.AbstractDQ_PositionalAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _abstractdq_positionalaccuracy_type.Name
			}
		}
	case "AbstractDQ_TemporalAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _abstractdq_temporalaccuracy_type, ok := stage.AbstractDQ_TemporalAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _abstractdq_temporalaccuracy_type.Name
			}
		}
	case "AbstractDQ_ThematicAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _abstractdq_thematicaccuracy_type, ok := stage.AbstractDQ_ThematicAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _abstractdq_thematicaccuracy_type.Name
			}
		}
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "Credit":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_Credit_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "CI_Address_Type":
		switch reverseField.Fieldname {
		case "DeliveryPoint":
			if _ci_address_type, ok := stage.CI_Address_Type_DeliveryPoint_reverseMap[inst]; ok {
				res = _ci_address_type.Name
			}
		case "ElectronicMailAddress":
			if _ci_address_type, ok := stage.CI_Address_Type_ElectronicMailAddress_reverseMap[inst]; ok {
				res = _ci_address_type.Name
			}
		}
	case "CI_Citation_Type":
		switch reverseField.Fieldname {
		case "AlternateTitle":
			if _ci_citation_type, ok := stage.CI_Citation_Type_AlternateTitle_reverseMap[inst]; ok {
				res = _ci_citation_type.Name
			}
		}
	case "CI_Telephone_Type":
		switch reverseField.Fieldname {
		case "Voice":
			if _ci_telephone_type, ok := stage.CI_Telephone_Type_Voice_reverseMap[inst]; ok {
				res = _ci_telephone_type.Name
			}
		case "Facsimile":
			if _ci_telephone_type, ok := stage.CI_Telephone_Type_Facsimile_reverseMap[inst]; ok {
				res = _ci_telephone_type.Name
			}
		}
	case "DQ_AbsoluteExternalPositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_absoluteexternalpositionalaccuracy_type, ok := stage.DQ_AbsoluteExternalPositionalAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_absoluteexternalpositionalaccuracy_type.Name
			}
		}
	case "DQ_AccuracyOfATimeMeasurement_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_accuracyofatimemeasurement_type, ok := stage.DQ_AccuracyOfATimeMeasurement_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_accuracyofatimemeasurement_type.Name
			}
		}
	case "DQ_CompletenessCommission_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_completenesscommission_type, ok := stage.DQ_CompletenessCommission_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_completenesscommission_type.Name
			}
		}
	case "DQ_CompletenessOmission_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_completenessomission_type, ok := stage.DQ_CompletenessOmission_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_completenessomission_type.Name
			}
		}
	case "DQ_ConceptualConsistency_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_conceptualconsistency_type, ok := stage.DQ_ConceptualConsistency_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_conceptualconsistency_type.Name
			}
		}
	case "DQ_DomainConsistency_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_domainconsistency_type, ok := stage.DQ_DomainConsistency_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_domainconsistency_type.Name
			}
		}
	case "DQ_FormatConsistency_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_formatconsistency_type, ok := stage.DQ_FormatConsistency_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_formatconsistency_type.Name
			}
		}
	case "DQ_GriddedDataPositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_griddeddatapositionalaccuracy_type, ok := stage.DQ_GriddedDataPositionalAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_griddeddatapositionalaccuracy_type.Name
			}
		}
	case "DQ_NonQuantitativeAttributeAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_nonquantitativeattributeaccuracy_type, ok := stage.DQ_NonQuantitativeAttributeAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_nonquantitativeattributeaccuracy_type.Name
			}
		}
	case "DQ_QuantitativeAttributeAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_quantitativeattributeaccuracy_type, ok := stage.DQ_QuantitativeAttributeAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_quantitativeattributeaccuracy_type.Name
			}
		}
	case "DQ_RelativeInternalPositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_relativeinternalpositionalaccuracy_type, ok := stage.DQ_RelativeInternalPositionalAccuracy_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_relativeinternalpositionalaccuracy_type.Name
			}
		}
	case "DQ_TemporalConsistency_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_temporalconsistency_type, ok := stage.DQ_TemporalConsistency_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_temporalconsistency_type.Name
			}
		}
	case "DQ_TemporalValidity_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_temporalvalidity_type, ok := stage.DQ_TemporalValidity_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_temporalvalidity_type.Name
			}
		}
	case "DQ_ThematicClassificationCorrectness_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_thematicclassificationcorrectness_type, ok := stage.DQ_ThematicClassificationCorrectness_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_thematicclassificationcorrectness_type.Name
			}
		}
	case "DQ_TopologicalConsistency_Type":
		switch reverseField.Fieldname {
		case "NameOfMeasure":
			if _dq_topologicalconsistency_type, ok := stage.DQ_TopologicalConsistency_Type_NameOfMeasure_reverseMap[inst]; ok {
				res = _dq_topologicalconsistency_type.Name
			}
		}
	case "MD_Constraints_Type":
		switch reverseField.Fieldname {
		case "UseLimitation":
			if _md_constraints_type, ok := stage.MD_Constraints_Type_UseLimitation_reverseMap[inst]; ok {
				res = _md_constraints_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "Credit":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_Credit_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		case "Language":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_Language_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_Keywords_Type":
		switch reverseField.Fieldname {
		case "Keyword":
			if _md_keywords_type, ok := stage.MD_Keywords_Type_Keyword_reverseMap[inst]; ok {
				res = _md_keywords_type.Name
			}
		}
	case "MD_LegalConstraints_Type":
		switch reverseField.Fieldname {
		case "UseLimitation":
			if _md_legalconstraints_type, ok := stage.MD_LegalConstraints_Type_UseLimitation_reverseMap[inst]; ok {
				res = _md_legalconstraints_type.Name
			}
		case "OtherConstraints":
			if _md_legalconstraints_type, ok := stage.MD_LegalConstraints_Type_OtherConstraints_reverseMap[inst]; ok {
				res = _md_legalconstraints_type.Name
			}
		}
	case "MD_MaintenanceInformation_Type":
		switch reverseField.Fieldname {
		case "MaintenanceNote":
			if _md_maintenanceinformation_type, ok := stage.MD_MaintenanceInformation_Type_MaintenanceNote_reverseMap[inst]; ok {
				res = _md_maintenanceinformation_type.Name
			}
		}
	case "MD_SecurityConstraints_Type":
		switch reverseField.Fieldname {
		case "UseLimitation":
			if _md_securityconstraints_type, ok := stage.MD_SecurityConstraints_Type_UseLimitation_reverseMap[inst]; ok {
				res = _md_securityconstraints_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "Credit":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_Credit_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *CircleByCenterPointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CircleByCenterPoint":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CircleByCenterPoint_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "CircleByCenterPoint":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_CircleByCenterPoint_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CircleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Circle":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Circle_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Circle":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Circle_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ClothoidType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Clothoid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Clothoid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Clothoid":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Clothoid_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CompositeCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CompositeCurve":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CompositeCurve_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveArrayPropertyType":
		switch reverseField.Fieldname {
		case "CompositeCurve":
			if _curvearraypropertytype, ok := stage.CurveArrayPropertyType_CompositeCurve_reverseMap[inst]; ok {
				res = _curvearraypropertytype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "CompositeCurve":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_CompositeCurve_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CompositeSolidType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CompositeSolid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CompositeSolid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "CompositeSolid":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_CompositeSolid_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SolidArrayPropertyType":
		switch reverseField.Fieldname {
		case "CompositeSolid":
			if _solidarraypropertytype, ok := stage.SolidArrayPropertyType_CompositeSolid_reverseMap[inst]; ok {
				res = _solidarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CompositeSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CompositeSurface":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CompositeSurface_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "CompositeSurface":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_CompositeSurface_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SurfaceArrayPropertyType":
		switch reverseField.Fieldname {
		case "CompositeSurface":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_CompositeSurface_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CompositeValueType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CompositeValue":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CompositeValue_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *CompoundCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CompoundCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CompoundCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CompoundCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *ConcatenatedOperationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ConcatenatedOperationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ConcatenatedOperation":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ConcatenatedOperation_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *ConeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SurfacePatchArrayPropertyType":
		switch reverseField.Fieldname {
		case "Cone":
			if _surfacepatcharraypropertytype, ok := stage.SurfacePatchArrayPropertyType_Cone_reverseMap[inst]; ok {
				res = _surfacepatcharraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ConventionalUnitType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ConventionalUnit":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ConventionalUnit_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *ConversionPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ConversionToPreferredUnitType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ConversionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Conversion":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Conversion_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *CoordinateOperationAccuracy) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractCoordinateOperationType":
		switch reverseField.Fieldname {
		case "CoordinateOperationAccuracy":
			if _abstractcoordinateoperationtype, ok := stage.AbstractCoordinateOperationType_CoordinateOperationAccuracy_reverseMap[inst]; ok {
				res = _abstractcoordinateoperationtype.Name
			}
		}
	case "AbstractGeneralConversionType":
		switch reverseField.Fieldname {
		case "CoordinateOperationAccuracy":
			if _abstractgeneralconversiontype, ok := stage.AbstractGeneralConversionType_CoordinateOperationAccuracy_reverseMap[inst]; ok {
				res = _abstractgeneralconversiontype.Name
			}
		}
	case "AbstractGeneralTransformationType":
		switch reverseField.Fieldname {
		case "CoordinateOperationAccuracy":
			if _abstractgeneraltransformationtype, ok := stage.AbstractGeneralTransformationType_CoordinateOperationAccuracy_reverseMap[inst]; ok {
				res = _abstractgeneraltransformationtype.Name
			}
		}
	case "ConcatenatedOperationType":
		switch reverseField.Fieldname {
		case "CoordinateOperationAccuracy":
			if _concatenatedoperationtype, ok := stage.ConcatenatedOperationType_CoordinateOperationAccuracy_reverseMap[inst]; ok {
				res = _concatenatedoperationtype.Name
			}
		}
	case "ConversionType":
		switch reverseField.Fieldname {
		case "CoordinateOperationAccuracy":
			if _conversiontype, ok := stage.ConversionType_CoordinateOperationAccuracy_reverseMap[inst]; ok {
				res = _conversiontype.Name
			}
		}
	case "PassThroughOperationType":
		switch reverseField.Fieldname {
		case "CoordinateOperationAccuracy":
			if _passthroughoperationtype, ok := stage.PassThroughOperationType_CoordinateOperationAccuracy_reverseMap[inst]; ok {
				res = _passthroughoperationtype.Name
			}
		}
	case "TransformationType":
		switch reverseField.Fieldname {
		case "CoordinateOperationAccuracy":
			if _transformationtype, ok := stage.TransformationType_CoordinateOperationAccuracy_reverseMap[inst]; ok {
				res = _transformationtype.Name
			}
		}
	}
	return
}

func (inst *CoordinateOperationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ConcatenatedOperationType":
		switch reverseField.Fieldname {
		case "CoordOperation":
			if _concatenatedoperationtype, ok := stage.ConcatenatedOperationType_CoordOperation_reverseMap[inst]; ok {
				res = _concatenatedoperationtype.Name
			}
		case "UsesSingleOperation":
			if _concatenatedoperationtype, ok := stage.ConcatenatedOperationType_UsesSingleOperation_reverseMap[inst]; ok {
				res = _concatenatedoperationtype.Name
			}
		case "UsesOperation":
			if _concatenatedoperationtype, ok := stage.ConcatenatedOperationType_UsesOperation_reverseMap[inst]; ok {
				res = _concatenatedoperationtype.Name
			}
		}
	}
	return
}

func (inst *CoordinateSystemAxisPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractCoordinateSystemType":
		switch reverseField.Fieldname {
		case "Axis":
			if _abstractcoordinatesystemtype, ok := stage.AbstractCoordinateSystemType_Axis_reverseMap[inst]; ok {
				res = _abstractcoordinatesystemtype.Name
			}
		case "UsesAxis":
			if _abstractcoordinatesystemtype, ok := stage.AbstractCoordinateSystemType_UsesAxis_reverseMap[inst]; ok {
				res = _abstractcoordinatesystemtype.Name
			}
		}
	case "AffineCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _affinecstype, ok := stage.AffineCSType_Axis_reverseMap[inst]; ok {
				res = _affinecstype.Name
			}
		case "UsesAxis":
			if _affinecstype, ok := stage.AffineCSType_UsesAxis_reverseMap[inst]; ok {
				res = _affinecstype.Name
			}
		}
	case "CartesianCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _cartesiancstype, ok := stage.CartesianCSType_Axis_reverseMap[inst]; ok {
				res = _cartesiancstype.Name
			}
		case "UsesAxis":
			if _cartesiancstype, ok := stage.CartesianCSType_UsesAxis_reverseMap[inst]; ok {
				res = _cartesiancstype.Name
			}
		}
	case "CylindricalCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _cylindricalcstype, ok := stage.CylindricalCSType_Axis_reverseMap[inst]; ok {
				res = _cylindricalcstype.Name
			}
		case "UsesAxis":
			if _cylindricalcstype, ok := stage.CylindricalCSType_UsesAxis_reverseMap[inst]; ok {
				res = _cylindricalcstype.Name
			}
		}
	case "EllipsoidalCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _ellipsoidalcstype, ok := stage.EllipsoidalCSType_Axis_reverseMap[inst]; ok {
				res = _ellipsoidalcstype.Name
			}
		case "UsesAxis":
			if _ellipsoidalcstype, ok := stage.EllipsoidalCSType_UsesAxis_reverseMap[inst]; ok {
				res = _ellipsoidalcstype.Name
			}
		}
	case "LinearCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _linearcstype, ok := stage.LinearCSType_Axis_reverseMap[inst]; ok {
				res = _linearcstype.Name
			}
		case "UsesAxis":
			if _linearcstype, ok := stage.LinearCSType_UsesAxis_reverseMap[inst]; ok {
				res = _linearcstype.Name
			}
		}
	case "ObliqueCartesianCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _obliquecartesiancstype, ok := stage.ObliqueCartesianCSType_Axis_reverseMap[inst]; ok {
				res = _obliquecartesiancstype.Name
			}
		case "UsesAxis":
			if _obliquecartesiancstype, ok := stage.ObliqueCartesianCSType_UsesAxis_reverseMap[inst]; ok {
				res = _obliquecartesiancstype.Name
			}
		}
	case "PolarCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _polarcstype, ok := stage.PolarCSType_Axis_reverseMap[inst]; ok {
				res = _polarcstype.Name
			}
		case "UsesAxis":
			if _polarcstype, ok := stage.PolarCSType_UsesAxis_reverseMap[inst]; ok {
				res = _polarcstype.Name
			}
		}
	case "SphericalCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _sphericalcstype, ok := stage.SphericalCSType_Axis_reverseMap[inst]; ok {
				res = _sphericalcstype.Name
			}
		case "UsesAxis":
			if _sphericalcstype, ok := stage.SphericalCSType_UsesAxis_reverseMap[inst]; ok {
				res = _sphericalcstype.Name
			}
		}
	case "TemporalCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _temporalcstype, ok := stage.TemporalCSType_Axis_reverseMap[inst]; ok {
				res = _temporalcstype.Name
			}
		case "UsesAxis":
			if _temporalcstype, ok := stage.TemporalCSType_UsesAxis_reverseMap[inst]; ok {
				res = _temporalcstype.Name
			}
		}
	case "TimeCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _timecstype, ok := stage.TimeCSType_Axis_reverseMap[inst]; ok {
				res = _timecstype.Name
			}
		case "UsesAxis":
			if _timecstype, ok := stage.TimeCSType_UsesAxis_reverseMap[inst]; ok {
				res = _timecstype.Name
			}
		}
	case "UserDefinedCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _userdefinedcstype, ok := stage.UserDefinedCSType_Axis_reverseMap[inst]; ok {
				res = _userdefinedcstype.Name
			}
		case "UsesAxis":
			if _userdefinedcstype, ok := stage.UserDefinedCSType_UsesAxis_reverseMap[inst]; ok {
				res = _userdefinedcstype.Name
			}
		}
	case "VerticalCSType":
		switch reverseField.Fieldname {
		case "Axis":
			if _verticalcstype, ok := stage.VerticalCSType_Axis_reverseMap[inst]; ok {
				res = _verticalcstype.Name
			}
		case "UsesAxis":
			if _verticalcstype, ok := stage.VerticalCSType_UsesAxis_reverseMap[inst]; ok {
				res = _verticalcstype.Name
			}
		}
	}
	return
}

func (inst *CoordinateSystemAxisType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CoordinateSystemAxis":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CoordinateSystemAxis_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *CoordinateSystemPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Count) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Count":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Count_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *CountPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Country_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CoverageFunctionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CoverageFunction":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CoverageFunction_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *CubicSplineType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CubicSpline":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CubicSpline_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "CubicSpline":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_CubicSpline_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CurveArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CurvePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CompositeCurveType":
		switch reverseField.Fieldname {
		case "CurveMember":
			if _compositecurvetype, ok := stage.CompositeCurveType_CurveMember_reverseMap[inst]; ok {
				res = _compositecurvetype.Name
			}
		}
	case "MultiCurveType":
		switch reverseField.Fieldname {
		case "CurveMember":
			if _multicurvetype, ok := stage.MultiCurveType_CurveMember_reverseMap[inst]; ok {
				res = _multicurvetype.Name
			}
		}
	case "RingType":
		switch reverseField.Fieldname {
		case "CurveMember":
			if _ringtype, ok := stage.RingType_CurveMember_reverseMap[inst]; ok {
				res = _ringtype.Name
			}
		}
	}
	return
}

func (inst *CurveSegmentArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Curve":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Curve_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveArrayPropertyType":
		switch reverseField.Fieldname {
		case "Curve":
			if _curvearraypropertytype, ok := stage.CurveArrayPropertyType_Curve_reverseMap[inst]; ok {
				res = _curvearraypropertytype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Curve":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Curve_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CylinderType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SurfacePatchArrayPropertyType":
		switch reverseField.Fieldname {
		case "Cylinder":
			if _surfacepatcharraypropertytype, ok := stage.SurfacePatchArrayPropertyType_Cylinder_reverseMap[inst]; ok {
				res = _surfacepatcharraypropertytype.Name
			}
		}
	}
	return
}

func (inst *CylindricalCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *CylindricalCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CylindricalCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CylindricalCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *DMSAngleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_AbsoluteExternalPositionalAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_AbsoluteExternalPositionalAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_AccuracyOfATimeMeasurement_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_AccuracyOfATimeMeasurement_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_CompletenessCommission_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_CompletenessCommission_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_CompletenessOmission_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_CompletenessOmission_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_Completeness_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_ConceptualConsistency_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_ConceptualConsistency_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_ConformanceResult_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_ConformanceResult_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_DataQuality_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_DataQuality_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_DomainConsistency_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_DomainConsistency_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_Element_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DQ_DataQuality_Type":
		switch reverseField.Fieldname {
		case "Report":
			if _dq_dataquality_type, ok := stage.DQ_DataQuality_Type_Report_reverseMap[inst]; ok {
				res = _dq_dataquality_type.Name
			}
		}
	}
	return
}

func (inst *DQ_EvaluationMethodTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_FormatConsistency_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_FormatConsistency_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_GriddedDataPositionalAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_GriddedDataPositionalAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_LogicalConsistency_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_NonQuantitativeAttributeAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_NonQuantitativeAttributeAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_PositionalAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_QuantitativeAttributeAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_QuantitativeAttributeAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_QuantitativeResult_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_QuantitativeResult_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_RelativeInternalPositionalAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_RelativeInternalPositionalAccuracy_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_Result_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_Scope_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_Scope_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_TemporalAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_TemporalConsistency_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_TemporalConsistency_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_TemporalValidity_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_TemporalValidity_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_ThematicAccuracy_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_ThematicClassificationCorrectness_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_ThematicClassificationCorrectness_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_TopologicalConsistency_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DQ_TopologicalConsistency_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_AssociationTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Association_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_Association_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DS_InitiativeTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DataBlockType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DataBlock":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DataBlock_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *DateTime_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractDQ_Completeness_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _abstractdq_completeness_type, ok := stage.AbstractDQ_Completeness_Type_DateTime_reverseMap[inst]; ok {
				res = _abstractdq_completeness_type.Name
			}
		}
	case "AbstractDQ_Element_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _abstractdq_element_type, ok := stage.AbstractDQ_Element_Type_DateTime_reverseMap[inst]; ok {
				res = _abstractdq_element_type.Name
			}
		}
	case "AbstractDQ_LogicalConsistency_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _abstractdq_logicalconsistency_type, ok := stage.AbstractDQ_LogicalConsistency_Type_DateTime_reverseMap[inst]; ok {
				res = _abstractdq_logicalconsistency_type.Name
			}
		}
	case "AbstractDQ_PositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _abstractdq_positionalaccuracy_type, ok := stage.AbstractDQ_PositionalAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _abstractdq_positionalaccuracy_type.Name
			}
		}
	case "AbstractDQ_TemporalAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _abstractdq_temporalaccuracy_type, ok := stage.AbstractDQ_TemporalAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _abstractdq_temporalaccuracy_type.Name
			}
		}
	case "AbstractDQ_ThematicAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _abstractdq_thematicaccuracy_type, ok := stage.AbstractDQ_ThematicAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _abstractdq_thematicaccuracy_type.Name
			}
		}
	case "DQ_AbsoluteExternalPositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_absoluteexternalpositionalaccuracy_type, ok := stage.DQ_AbsoluteExternalPositionalAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_absoluteexternalpositionalaccuracy_type.Name
			}
		}
	case "DQ_AccuracyOfATimeMeasurement_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_accuracyofatimemeasurement_type, ok := stage.DQ_AccuracyOfATimeMeasurement_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_accuracyofatimemeasurement_type.Name
			}
		}
	case "DQ_CompletenessCommission_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_completenesscommission_type, ok := stage.DQ_CompletenessCommission_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_completenesscommission_type.Name
			}
		}
	case "DQ_CompletenessOmission_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_completenessomission_type, ok := stage.DQ_CompletenessOmission_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_completenessomission_type.Name
			}
		}
	case "DQ_ConceptualConsistency_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_conceptualconsistency_type, ok := stage.DQ_ConceptualConsistency_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_conceptualconsistency_type.Name
			}
		}
	case "DQ_DomainConsistency_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_domainconsistency_type, ok := stage.DQ_DomainConsistency_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_domainconsistency_type.Name
			}
		}
	case "DQ_FormatConsistency_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_formatconsistency_type, ok := stage.DQ_FormatConsistency_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_formatconsistency_type.Name
			}
		}
	case "DQ_GriddedDataPositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_griddeddatapositionalaccuracy_type, ok := stage.DQ_GriddedDataPositionalAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_griddeddatapositionalaccuracy_type.Name
			}
		}
	case "DQ_NonQuantitativeAttributeAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_nonquantitativeattributeaccuracy_type, ok := stage.DQ_NonQuantitativeAttributeAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_nonquantitativeattributeaccuracy_type.Name
			}
		}
	case "DQ_QuantitativeAttributeAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_quantitativeattributeaccuracy_type, ok := stage.DQ_QuantitativeAttributeAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_quantitativeattributeaccuracy_type.Name
			}
		}
	case "DQ_RelativeInternalPositionalAccuracy_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_relativeinternalpositionalaccuracy_type, ok := stage.DQ_RelativeInternalPositionalAccuracy_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_relativeinternalpositionalaccuracy_type.Name
			}
		}
	case "DQ_TemporalConsistency_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_temporalconsistency_type, ok := stage.DQ_TemporalConsistency_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_temporalconsistency_type.Name
			}
		}
	case "DQ_TemporalValidity_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_temporalvalidity_type, ok := stage.DQ_TemporalValidity_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_temporalvalidity_type.Name
			}
		}
	case "DQ_ThematicClassificationCorrectness_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_thematicclassificationcorrectness_type, ok := stage.DQ_ThematicClassificationCorrectness_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_thematicclassificationcorrectness_type.Name
			}
		}
	case "DQ_TopologicalConsistency_Type":
		switch reverseField.Fieldname {
		case "DateTime":
			if _dq_topologicalconsistency_type, ok := stage.DQ_TopologicalConsistency_Type_DateTime_reverseMap[inst]; ok {
				res = _dq_topologicalconsistency_type.Name
			}
		}
	}
	return
}

func (inst *Date_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DatumPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Decimal_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DefinitionBaseType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DefinitionProxyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DefinitionProxy":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DefinitionProxy_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *DefinitionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Definition":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Definition_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *DegreesType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DerivationUnitTermType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ConventionalUnitType":
		switch reverseField.Fieldname {
		case "DerivationUnitTerm":
			if _conventionalunittype, ok := stage.ConventionalUnitType_DerivationUnitTerm_reverseMap[inst]; ok {
				res = _conventionalunittype.Name
			}
		}
	case "DerivedUnitType":
		switch reverseField.Fieldname {
		case "DerivationUnitTerm":
			if _derivedunittype, ok := stage.DerivedUnitType_DerivationUnitTerm_reverseMap[inst]; ok {
				res = _derivedunittype.Name
			}
		}
	}
	return
}

func (inst *DerivedCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DerivedCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DerivedCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DerivedCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *DerivedUnitType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DerivedUnit":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DerivedUnit_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *DictionaryEntryType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DictionaryType":
		switch reverseField.Fieldname {
		case "DictionaryEntry":
			if _dictionarytype, ok := stage.DictionaryType_DictionaryEntry_reverseMap[inst]; ok {
				res = _dictionarytype.Name
			}
		case "DefinitionMember":
			if _dictionarytype, ok := stage.DictionaryType_DefinitionMember_reverseMap[inst]; ok {
				res = _dictionarytype.Name
			}
		}
	}
	return
}

func (inst *DictionaryType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Dictionary":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Dictionary_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "DefinitionCollection":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DefinitionCollection_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *DirectPositionListType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectPositionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArcStringByBulgeType":
		switch reverseField.Fieldname {
		case "Pos":
			if _arcstringbybulgetype, ok := stage.ArcStringByBulgeType_Pos_reverseMap[inst]; ok {
				res = _arcstringbybulgetype.Name
			}
		}
	case "ArcStringType":
		switch reverseField.Fieldname {
		case "Pos":
			if _arcstringtype, ok := stage.ArcStringType_Pos_reverseMap[inst]; ok {
				res = _arcstringtype.Name
			}
		}
	case "BSplineType":
		switch reverseField.Fieldname {
		case "Pos":
			if _bsplinetype, ok := stage.BSplineType_Pos_reverseMap[inst]; ok {
				res = _bsplinetype.Name
			}
		}
	case "BezierType":
		switch reverseField.Fieldname {
		case "Pos":
			if _beziertype, ok := stage.BezierType_Pos_reverseMap[inst]; ok {
				res = _beziertype.Name
			}
		}
	case "CircleType":
		switch reverseField.Fieldname {
		case "Pos":
			if _circletype, ok := stage.CircleType_Pos_reverseMap[inst]; ok {
				res = _circletype.Name
			}
		}
	case "CubicSplineType":
		switch reverseField.Fieldname {
		case "Pos":
			if _cubicsplinetype, ok := stage.CubicSplineType_Pos_reverseMap[inst]; ok {
				res = _cubicsplinetype.Name
			}
		}
	case "LineStringSegmentType":
		switch reverseField.Fieldname {
		case "Pos":
			if _linestringsegmenttype, ok := stage.LineStringSegmentType_Pos_reverseMap[inst]; ok {
				res = _linestringsegmenttype.Name
			}
		}
	case "LineStringType":
		switch reverseField.Fieldname {
		case "Pos":
			if _linestringtype, ok := stage.LineStringType_Pos_reverseMap[inst]; ok {
				res = _linestringtype.Name
			}
		}
	case "LinearRingType":
		switch reverseField.Fieldname {
		case "Pos":
			if _linearringtype, ok := stage.LinearRingType_Pos_reverseMap[inst]; ok {
				res = _linearringtype.Name
			}
		}
	}
	return
}

func (inst *DirectedEdgePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "FaceType":
		switch reverseField.Fieldname {
		case "DirectedEdge":
			if _facetype, ok := stage.FaceType_DirectedEdge_reverseMap[inst]; ok {
				res = _facetype.Name
			}
		}
	case "NodeType":
		switch reverseField.Fieldname {
		case "DirectedEdge":
			if _nodetype, ok := stage.NodeType_DirectedEdge_reverseMap[inst]; ok {
				res = _nodetype.Name
			}
		}
	case "TopoCurveType":
		switch reverseField.Fieldname {
		case "DirectedEdge":
			if _topocurvetype, ok := stage.TopoCurveType_DirectedEdge_reverseMap[inst]; ok {
				res = _topocurvetype.Name
			}
		}
	}
	return
}

func (inst *DirectedFacePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "EdgeType":
		switch reverseField.Fieldname {
		case "DirectedFace":
			if _edgetype, ok := stage.EdgeType_DirectedFace_reverseMap[inst]; ok {
				res = _edgetype.Name
			}
		}
	case "TopoSolidType":
		switch reverseField.Fieldname {
		case "DirectedFace":
			if _toposolidtype, ok := stage.TopoSolidType_DirectedFace_reverseMap[inst]; ok {
				res = _toposolidtype.Name
			}
		}
	case "TopoSurfaceType":
		switch reverseField.Fieldname {
		case "DirectedFace":
			if _toposurfacetype, ok := stage.TopoSurfaceType_DirectedFace_reverseMap[inst]; ok {
				res = _toposurfacetype.Name
			}
		}
	}
	return
}

func (inst *DirectedNodePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectedObservationAtDistanceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DirectedObservationAtDistance":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DirectedObservationAtDistance_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "FeatureArrayPropertyType":
		switch reverseField.Fieldname {
		case "DirectedObservationAtDistance":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_DirectedObservationAtDistance_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *DirectedObservationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DirectedObservation":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DirectedObservation_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "FeatureArrayPropertyType":
		switch reverseField.Fieldname {
		case "DirectedObservation":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_DirectedObservation_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *DirectedTopoSolidPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TopoVolumeType":
		switch reverseField.Fieldname {
		case "DirectedTopoSolid":
			if _topovolumetype, ok := stage.TopoVolumeType_DirectedTopoSolid_reverseMap[inst]; ok {
				res = _topovolumetype.Name
			}
		}
	}
	return
}

func (inst *DirectionDescriptionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectionPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DirectionVectorType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DiscreteCoverageType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "MultiPointCoverage":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiPointCoverage_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "MultiCurveCoverage":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiCurveCoverage_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "MultiSurfaceCoverage":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiSurfaceCoverage_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "MultiSolidCoverage":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiSolidCoverage_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "GridCoverage":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GridCoverage_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "RectifiedGridCoverage":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_RectifiedGridCoverage_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "FeatureArrayPropertyType":
		switch reverseField.Fieldname {
		case "MultiPointCoverage":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_MultiPointCoverage_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		case "MultiCurveCoverage":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_MultiCurveCoverage_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		case "MultiSurfaceCoverage":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_MultiSurfaceCoverage_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		case "MultiSolidCoverage":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_MultiSolidCoverage_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		case "GridCoverage":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_GridCoverage_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		case "RectifiedGridCoverage":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_RectifiedGridCoverage_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *Distance_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DomainOfValidity) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _abstractcrstype, ok := stage.AbstractCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _abstractcrstype.Name
			}
		}
	case "AbstractGeneralDerivedCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _abstractgeneralderivedcrstype, ok := stage.AbstractGeneralDerivedCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _abstractgeneralderivedcrstype.Name
			}
		}
	case "CompoundCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _compoundcrstype, ok := stage.CompoundCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _compoundcrstype.Name
			}
		}
	case "DerivedCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _derivedcrstype, ok := stage.DerivedCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _derivedcrstype.Name
			}
		}
	case "EngineeringCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _engineeringcrstype, ok := stage.EngineeringCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _engineeringcrstype.Name
			}
		}
	case "GeocentricCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _geocentriccrstype, ok := stage.GeocentricCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _geocentriccrstype.Name
			}
		}
	case "GeodeticCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _geodeticcrstype, ok := stage.GeodeticCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _geodeticcrstype.Name
			}
		}
	case "GeographicCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _geographiccrstype, ok := stage.GeographicCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _geographiccrstype.Name
			}
		}
	case "ImageCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _imagecrstype, ok := stage.ImageCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _imagecrstype.Name
			}
		}
	case "ProjectedCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _projectedcrstype, ok := stage.ProjectedCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _projectedcrstype.Name
			}
		}
	case "TemporalCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _temporalcrstype, ok := stage.TemporalCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _temporalcrstype.Name
			}
		}
	case "VerticalCRSType":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _verticalcrstype, ok := stage.VerticalCRSType_DomainOfValidity_reverseMap[inst]; ok {
				res = _verticalcrstype.Name
			}
		}
	}
	return
}

func (inst *DomainSetType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DynamicFeatureCollectionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DynamicFeatureCollection":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DynamicFeatureCollection_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "DynamicFeatureMemberType":
		switch reverseField.Fieldname {
		case "DynamicFeatureCollection":
			if _dynamicfeaturemembertype, ok := stage.DynamicFeatureMemberType_DynamicFeatureCollection_reverseMap[inst]; ok {
				res = _dynamicfeaturemembertype.Name
			}
		}
	case "FeatureArrayPropertyType":
		switch reverseField.Fieldname {
		case "DynamicFeatureCollection":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_DynamicFeatureCollection_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *DynamicFeatureMemberType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *DynamicFeatureType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "DynamicFeature":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_DynamicFeature_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "DynamicFeatureMemberType":
		switch reverseField.Fieldname {
		case "DynamicFeature":
			if _dynamicfeaturemembertype, ok := stage.DynamicFeatureMemberType_DynamicFeature_reverseMap[inst]; ok {
				res = _dynamicfeaturemembertype.Name
			}
		}
	case "FeatureArrayPropertyType":
		switch reverseField.Fieldname {
		case "DynamicFeature":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_DynamicFeature_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *EX_BoundingPolygon_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_BoundingPolygon_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_Extent_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractRS_ReferenceSystem_Type":
		switch reverseField.Fieldname {
		case "DomainOfValidity":
			if _abstractrs_referencesystem_type, ok := stage.AbstractRS_ReferenceSystem_Type_DomainOfValidity_reverseMap[inst]; ok {
				res = _abstractrs_referencesystem_type.Name
			}
		}
	case "LI_Source_Type":
		switch reverseField.Fieldname {
		case "SourceExtent":
			if _li_source_type, ok := stage.LI_Source_Type_SourceExtent_reverseMap[inst]; ok {
				res = _li_source_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "Extent":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_Extent_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	}
	return
}

func (inst *EX_Extent_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_GeographicBoundingBox_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_GeographicBoundingBox_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_GeographicDescription_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_GeographicDescription_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_GeographicExtent_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "EX_Extent_Type":
		switch reverseField.Fieldname {
		case "GeographicElement":
			if _ex_extent_type, ok := stage.EX_Extent_Type_GeographicElement_reverseMap[inst]; ok {
				res = _ex_extent_type.Name
			}
		}
	case "EX_SpatialTemporalExtent_Type":
		switch reverseField.Fieldname {
		case "SpatialExtent":
			if _ex_spatialtemporalextent_type, ok := stage.EX_SpatialTemporalExtent_Type_SpatialExtent_reverseMap[inst]; ok {
				res = _ex_spatialtemporalextent_type.Name
			}
		}
	}
	return
}

func (inst *EX_SpatialTemporalExtent_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_SpatialTemporalExtent_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_TemporalExtent_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "EX_Extent_Type":
		switch reverseField.Fieldname {
		case "TemporalElement":
			if _ex_extent_type, ok := stage.EX_Extent_Type_TemporalElement_reverseMap[inst]; ok {
				res = _ex_extent_type.Name
			}
		}
	}
	return
}

func (inst *EX_TemporalExtent_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EX_VerticalExtent_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "EX_Extent_Type":
		switch reverseField.Fieldname {
		case "VerticalElement":
			if _ex_extent_type, ok := stage.EX_Extent_Type_VerticalElement_reverseMap[inst]; ok {
				res = _ex_extent_type.Name
			}
		}
	}
	return
}

func (inst *EX_VerticalExtent_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EdgeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Edge":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Edge_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "TopoPrimitiveArrayAssociationType":
		switch reverseField.Fieldname {
		case "Edge":
			if _topoprimitivearrayassociationtype, ok := stage.TopoPrimitiveArrayAssociationType_Edge_reverseMap[inst]; ok {
				res = _topoprimitivearrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *EllipsoidPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EllipsoidType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Ellipsoid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Ellipsoid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *EllipsoidalCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EllipsoidalCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "EllipsoidalCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_EllipsoidalCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *EngineeringCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EngineeringCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "EngineeringCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_EngineeringCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *EngineeringDatumPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *EngineeringDatumType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "EngineeringDatum":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_EngineeringDatum_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *EnvelopeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Envelope":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Envelope_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *EnvelopeWithTimePeriodType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "EnvelopeWithTimePeriod":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_EnvelopeWithTimePeriod_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *FaceOrTopoSolidPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Face":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Face_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "TopoPrimitiveArrayAssociationType":
		switch reverseField.Fieldname {
		case "Face":
			if _topoprimitivearrayassociationtype, ok := stage.TopoPrimitiveArrayAssociationType_Face_reverseMap[inst]; ok {
				res = _topoprimitivearrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *FeatureArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FeatureCollectionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "FeatureCollection":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_FeatureCollection_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "FeatureArrayPropertyType":
		switch reverseField.Fieldname {
		case "FeatureCollection":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_FeatureCollection_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *FeaturePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractFeatureCollectionType":
		switch reverseField.Fieldname {
		case "FeatureMember":
			if _abstractfeaturecollectiontype, ok := stage.AbstractFeatureCollectionType_FeatureMember_reverseMap[inst]; ok {
				res = _abstractfeaturecollectiontype.Name
			}
		}
	case "FeatureCollectionType":
		switch reverseField.Fieldname {
		case "FeatureMember":
			if _featurecollectiontype, ok := stage.FeatureCollectionType_FeatureMember_reverseMap[inst]; ok {
				res = _featurecollectiontype.Name
			}
		}
	}
	return
}

func (inst *FileType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "File":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_File_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *FormulaCitation) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *FormulaType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GM_Object_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "EX_BoundingPolygon_Type":
		switch reverseField.Fieldname {
		case "Polygon":
			if _ex_boundingpolygon_type, ok := stage.EX_BoundingPolygon_Type_Polygon_reverseMap[inst]; ok {
				res = _ex_boundingpolygon_type.Name
			}
		}
	}
	return
}

func (inst *GM_Point_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeneralConversionPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeneralTransformationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GenericMetaDataType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GenericMetaData":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GenericMetaData_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *GenericName_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeocentricCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeocentricCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GeocentricCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GeocentricCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *GeodesicStringType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GeodesicString":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GeodesicString_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "GeodesicString":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_GeodesicString_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *GeodesicType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Geodesic":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Geodesic_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "Geodesic":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_Geodesic_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *GeodeticCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeodeticCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GeodeticCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GeodeticCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *GeodeticDatumPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeodeticDatumType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GeodeticDatum":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GeodeticDatum_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *GeographicCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeographicCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GeographicCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GeographicCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *GeometricComplexPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeometricComplexType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GeometricComplex":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GeometricComplex_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "GeometricComplex":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_GeometricComplex_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *GeometricPrimitivePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GeometricComplexType":
		switch reverseField.Fieldname {
		case "Element":
			if _geometriccomplextype, ok := stage.GeometricComplexType_Element_reverseMap[inst]; ok {
				res = _geometriccomplextype.Name
			}
		}
	}
	return
}

func (inst *GeometryArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GeometryPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MultiGeometryType":
		switch reverseField.Fieldname {
		case "GeometryMember":
			if _multigeometrytype, ok := stage.MultiGeometryType_GeometryMember_reverseMap[inst]; ok {
				res = _multigeometrytype.Name
			}
		}
	}
	return
}

func (inst *GridEnvelopeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GridFunctionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "GridFunction":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_GridFunction_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *GridLimitsType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *GridType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Grid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Grid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Grid":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Grid_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *HistoryPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *IdentifiedObjectType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ImageCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ImageCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ImageCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ImageCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *ImageDatumPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ImageDatumType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ImageDatum":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ImageDatum_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *IndirectEntryType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DictionaryType":
		switch reverseField.Fieldname {
		case "IndirectEntry":
			if _dictionarytype, ok := stage.DictionaryType_IndirectEntry_reverseMap[inst]; ok {
				res = _dictionarytype.Name
			}
		}
	}
	return
}

func (inst *InlinePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Integer_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *KnotPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "BSplineType":
		switch reverseField.Fieldname {
		case "Knot":
			if _bsplinetype, ok := stage.BSplineType_Knot_reverseMap[inst]; ok {
				res = _bsplinetype.Name
			}
		}
	}
	return
}

func (inst *KnotType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LI_Lineage_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LI_Lineage_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LI_ProcessStep_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "LI_Lineage_Type":
		switch reverseField.Fieldname {
		case "ProcessStep":
			if _li_lineage_type, ok := stage.LI_Lineage_Type_ProcessStep_reverseMap[inst]; ok {
				res = _li_lineage_type.Name
			}
		}
	case "LI_Source_Type":
		switch reverseField.Fieldname {
		case "SourceStep":
			if _li_source_type, ok := stage.LI_Source_Type_SourceStep_reverseMap[inst]; ok {
				res = _li_source_type.Name
			}
		}
	}
	return
}

func (inst *LI_ProcessStep_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LI_Source_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "LI_Lineage_Type":
		switch reverseField.Fieldname {
		case "Source":
			if _li_lineage_type, ok := stage.LI_Lineage_Type_Source_reverseMap[inst]; ok {
				res = _li_lineage_type.Name
			}
		}
	case "LI_ProcessStep_Type":
		switch reverseField.Fieldname {
		case "Source":
			if _li_processstep_type, ok := stage.LI_ProcessStep_Type_Source_reverseMap[inst]; ok {
				res = _li_processstep_type.Name
			}
		}
	}
	return
}

func (inst *LI_Source_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LanguageCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Length_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LineStringSegmentArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TinType":
		switch reverseField.Fieldname {
		case "StopLines":
			if _tintype, ok := stage.TinType_StopLines_reverseMap[inst]; ok {
				res = _tintype.Name
			}
		case "BreakLines":
			if _tintype, ok := stage.TinType_BreakLines_reverseMap[inst]; ok {
				res = _tintype.Name
			}
		}
	}
	return
}

func (inst *LineStringSegmentType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "LineStringSegment":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_LineStringSegment_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "LineStringSegment":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_LineStringSegment_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	case "LineStringSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "LineStringSegment":
			if _linestringsegmentarraypropertytype, ok := stage.LineStringSegmentArrayPropertyType_LineStringSegment_reverseMap[inst]; ok {
				res = _linestringsegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *LineStringType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "LineString":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_LineString_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveArrayPropertyType":
		switch reverseField.Fieldname {
		case "LineString":
			if _curvearraypropertytype, ok := stage.CurveArrayPropertyType_LineString_reverseMap[inst]; ok {
				res = _curvearraypropertytype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "LineString":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_LineString_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *LinearCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LinearCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "LinearCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_LinearCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *LinearRingPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LinearRingType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "LinearRing":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_LinearRing_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveArrayPropertyType":
		switch reverseField.Fieldname {
		case "LinearRing":
			if _curvearraypropertytype, ok := stage.CurveArrayPropertyType_LinearRing_reverseMap[inst]; ok {
				res = _curvearraypropertytype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "LinearRing":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_LinearRing_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *LocalName_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LocalisedCharacterString_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "PT_FreeText_Type":
		switch reverseField.Fieldname {
		case "TextGroup":
			if _pt_freetext_type, ok := stage.PT_FreeText_Type_TextGroup_reverseMap[inst]; ok {
				res = _pt_freetext_type.Name
			}
		}
	case "PT_LocaleContainer_Type":
		switch reverseField.Fieldname {
		case "LocalisedString":
			if _pt_localecontainer_type, ok := stage.PT_LocaleContainer_Type_LocalisedString_reverseMap[inst]; ok {
				res = _pt_localecontainer_type.Name
			}
		}
	}
	return
}

func (inst *LocalisedCharacterString_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LocationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_AggregateInformation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "AggregationInfo":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_AggregationInfo_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "AggregationInfo":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_AggregationInfo_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "AggregationInfo":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_AggregationInfo_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_AggregateInformation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_BrowseGraphic_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "GraphicOverview":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_GraphicOverview_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "GraphicOverview":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_GraphicOverview_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "GraphicOverview":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_GraphicOverview_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_BrowseGraphic_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_CharacterSetCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "CharacterSet":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_CharacterSet_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_ClassificationCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Constraints_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "ResourceConstraints":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_ResourceConstraints_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceConstraints":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_ResourceConstraints_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceConstraints":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_ResourceConstraints_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_Constraints_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_DataIdentification_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_DataIdentification_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_DigitalTransferOptions_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_Distribution_Type":
		switch reverseField.Fieldname {
		case "TransferOptions":
			if _md_distribution_type, ok := stage.MD_Distribution_Type_TransferOptions_reverseMap[inst]; ok {
				res = _md_distribution_type.Name
			}
		}
	case "MD_Distributor_Type":
		switch reverseField.Fieldname {
		case "DistributorTransferOptions":
			if _md_distributor_type, ok := stage.MD_Distributor_Type_DistributorTransferOptions_reverseMap[inst]; ok {
				res = _md_distributor_type.Name
			}
		}
	}
	return
}

func (inst *MD_DigitalTransferOptions_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_DistributionUnits_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Distribution_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Distribution_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Distributor_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_Distribution_Type":
		switch reverseField.Fieldname {
		case "Distributor":
			if _md_distribution_type, ok := stage.MD_Distribution_Type_Distributor_reverseMap[inst]; ok {
				res = _md_distribution_type.Name
			}
		}
	case "MD_Format_Type":
		switch reverseField.Fieldname {
		case "FormatDistributor":
			if _md_format_type, ok := stage.MD_Format_Type_FormatDistributor_reverseMap[inst]; ok {
				res = _md_format_type.Name
			}
		}
	}
	return
}

func (inst *MD_Distributor_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Format_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "ResourceFormat":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_ResourceFormat_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceFormat":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_ResourceFormat_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_Distribution_Type":
		switch reverseField.Fieldname {
		case "DistributionFormat":
			if _md_distribution_type, ok := stage.MD_Distribution_Type_DistributionFormat_reverseMap[inst]; ok {
				res = _md_distribution_type.Name
			}
		}
	case "MD_Distributor_Type":
		switch reverseField.Fieldname {
		case "DistributorFormat":
			if _md_distributor_type, ok := stage.MD_Distributor_Type_DistributorFormat_reverseMap[inst]; ok {
				res = _md_distributor_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceFormat":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_ResourceFormat_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_Format_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Identification_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Identifier_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CI_Citation_Type":
		switch reverseField.Fieldname {
		case "Identifier":
			if _ci_citation_type, ok := stage.CI_Citation_Type_Identifier_reverseMap[inst]; ok {
				res = _ci_citation_type.Name
			}
		}
	}
	return
}

func (inst *MD_Identifier_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_KeywordTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Keywords_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "DescriptiveKeywords":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_DescriptiveKeywords_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "DescriptiveKeywords":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_DescriptiveKeywords_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "DescriptiveKeywords":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_DescriptiveKeywords_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_Keywords_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_LegalConstraints_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_LegalConstraints_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_MaintenanceFrequencyCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_MaintenanceInformation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "ResourceMaintenance":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_ResourceMaintenance_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceMaintenance":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_ResourceMaintenance_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceMaintenance":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_ResourceMaintenance_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_MaintenanceInformation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_MediumFormatCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_Medium_Type":
		switch reverseField.Fieldname {
		case "MediumFormat":
			if _md_medium_type, ok := stage.MD_Medium_Type_MediumFormat_reverseMap[inst]; ok {
				res = _md_medium_type.Name
			}
		}
	}
	return
}

func (inst *MD_MediumNameCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Medium_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Medium_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ProgressCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "Status":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_Status_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "Status":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_Status_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "Status":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_Status_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_ReferenceSystem_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ReferenceSystem_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_RepresentativeFraction_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_RepresentativeFraction_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Resolution_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "SpatialResolution":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_SpatialResolution_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_Resolution_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_RestrictionCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_LegalConstraints_Type":
		switch reverseField.Fieldname {
		case "AccessConstraints":
			if _md_legalconstraints_type, ok := stage.MD_LegalConstraints_Type_AccessConstraints_reverseMap[inst]; ok {
				res = _md_legalconstraints_type.Name
			}
		case "UseConstraints":
			if _md_legalconstraints_type, ok := stage.MD_LegalConstraints_Type_UseConstraints_reverseMap[inst]; ok {
				res = _md_legalconstraints_type.Name
			}
		}
	}
	return
}

func (inst *MD_ScopeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_MaintenanceInformation_Type":
		switch reverseField.Fieldname {
		case "UpdateScope":
			if _md_maintenanceinformation_type, ok := stage.MD_MaintenanceInformation_Type_UpdateScope_reverseMap[inst]; ok {
				res = _md_maintenanceinformation_type.Name
			}
		}
	}
	return
}

func (inst *MD_ScopeDescription_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DQ_Scope_Type":
		switch reverseField.Fieldname {
		case "LevelDescription":
			if _dq_scope_type, ok := stage.DQ_Scope_Type_LevelDescription_reverseMap[inst]; ok {
				res = _dq_scope_type.Name
			}
		}
	case "MD_MaintenanceInformation_Type":
		switch reverseField.Fieldname {
		case "UpdateScopeDescription":
			if _md_maintenanceinformation_type, ok := stage.MD_MaintenanceInformation_Type_UpdateScopeDescription_reverseMap[inst]; ok {
				res = _md_maintenanceinformation_type.Name
			}
		}
	}
	return
}

func (inst *MD_ScopeDescription_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_SecurityConstraints_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_SecurityConstraints_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ServiceIdentification_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ServiceIdentification_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_SpatialRepresentationTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "SpatialRepresentationType":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_SpatialRepresentationType_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_StandardOrderProcess_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_Distributor_Type":
		switch reverseField.Fieldname {
		case "DistributionOrderProcess":
			if _md_distributor_type, ok := stage.MD_Distributor_Type_DistributionOrderProcess_reverseMap[inst]; ok {
				res = _md_distributor_type.Name
			}
		}
	}
	return
}

func (inst *MD_StandardOrderProcess_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_TopicCategoryCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "TopicCategory":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_TopicCategory_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_Usage_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractMD_Identification_Type":
		switch reverseField.Fieldname {
		case "ResourceSpecificUsage":
			if _abstractmd_identification_type, ok := stage.AbstractMD_Identification_Type_ResourceSpecificUsage_reverseMap[inst]; ok {
				res = _abstractmd_identification_type.Name
			}
		}
	case "MD_DataIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceSpecificUsage":
			if _md_dataidentification_type, ok := stage.MD_DataIdentification_Type_ResourceSpecificUsage_reverseMap[inst]; ok {
				res = _md_dataidentification_type.Name
			}
		}
	case "MD_ServiceIdentification_Type":
		switch reverseField.Fieldname {
		case "ResourceSpecificUsage":
			if _md_serviceidentification_type, ok := stage.MD_ServiceIdentification_Type_ResourceSpecificUsage_reverseMap[inst]; ok {
				res = _md_serviceidentification_type.Name
			}
		}
	}
	return
}

func (inst *MD_Usage_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MappingRuleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "CoverageMappingRule":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_CoverageMappingRule_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *Measure_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MemberName_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MemberName_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MetaDataPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractcrstype, ok := stage.AbstractCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractcrstype.Name
			}
		}
	case "AbstractContinuousCoverageType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractcontinuouscoveragetype, ok := stage.AbstractContinuousCoverageType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractcontinuouscoveragetype.Name
			}
		}
	case "AbstractCoordinateOperationType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractcoordinateoperationtype, ok := stage.AbstractCoordinateOperationType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractcoordinateoperationtype.Name
			}
		}
	case "AbstractCoordinateSystemType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractcoordinatesystemtype, ok := stage.AbstractCoordinateSystemType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractcoordinatesystemtype.Name
			}
		}
	case "AbstractCoverageType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractcoveragetype, ok := stage.AbstractCoverageType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractcoveragetype.Name
			}
		}
	case "AbstractCurveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractcurvetype, ok := stage.AbstractCurveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractcurvetype.Name
			}
		}
	case "AbstractDatumType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractdatumtype, ok := stage.AbstractDatumType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractdatumtype.Name
			}
		}
	case "AbstractFeatureCollectionType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractfeaturecollectiontype, ok := stage.AbstractFeatureCollectionType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractfeaturecollectiontype.Name
			}
		}
	case "AbstractFeatureType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractfeaturetype, ok := stage.AbstractFeatureType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractfeaturetype.Name
			}
		}
	case "AbstractGMLType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgmltype, ok := stage.AbstractGMLType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgmltype.Name
			}
		}
	case "AbstractGeneralConversionType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgeneralconversiontype, ok := stage.AbstractGeneralConversionType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgeneralconversiontype.Name
			}
		}
	case "AbstractGeneralDerivedCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgeneralderivedcrstype, ok := stage.AbstractGeneralDerivedCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgeneralderivedcrstype.Name
			}
		}
	case "AbstractGeneralOperationParameterType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgeneraloperationparametertype, ok := stage.AbstractGeneralOperationParameterType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgeneraloperationparametertype.Name
			}
		}
	case "AbstractGeneralTransformationType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgeneraltransformationtype, ok := stage.AbstractGeneralTransformationType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgeneraltransformationtype.Name
			}
		}
	case "AbstractGeometricAggregateType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgeometricaggregatetype, ok := stage.AbstractGeometricAggregateType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgeometricaggregatetype.Name
			}
		}
	case "AbstractGeometricPrimitiveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgeometricprimitivetype, ok := stage.AbstractGeometricPrimitiveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgeometricprimitivetype.Name
			}
		}
	case "AbstractGeometryType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractgeometrytype, ok := stage.AbstractGeometryType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractgeometrytype.Name
			}
		}
	case "AbstractRingType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractringtype, ok := stage.AbstractRingType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractringtype.Name
			}
		}
	case "AbstractSolidType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractsolidtype, ok := stage.AbstractSolidType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractsolidtype.Name
			}
		}
	case "AbstractSurfaceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstractsurfacetype, ok := stage.AbstractSurfaceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstractsurfacetype.Name
			}
		}
	case "AbstractTimeComplexType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttimecomplextype, ok := stage.AbstractTimeComplexType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttimecomplextype.Name
			}
		}
	case "AbstractTimeGeometricPrimitiveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttimegeometricprimitivetype, ok := stage.AbstractTimeGeometricPrimitiveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttimegeometricprimitivetype.Name
			}
		}
	case "AbstractTimeObjectType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttimeobjecttype, ok := stage.AbstractTimeObjectType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttimeobjecttype.Name
			}
		}
	case "AbstractTimePrimitiveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttimeprimitivetype, ok := stage.AbstractTimePrimitiveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttimeprimitivetype.Name
			}
		}
	case "AbstractTimeSliceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttimeslicetype, ok := stage.AbstractTimeSliceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttimeslicetype.Name
			}
		}
	case "AbstractTimeTopologyPrimitiveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttimetopologyprimitivetype, ok := stage.AbstractTimeTopologyPrimitiveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttimetopologyprimitivetype.Name
			}
		}
	case "AbstractTopoPrimitiveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttopoprimitivetype, ok := stage.AbstractTopoPrimitiveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttopoprimitivetype.Name
			}
		}
	case "AbstractTopologyType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _abstracttopologytype, ok := stage.AbstractTopologyType_MetaDataProperty_reverseMap[inst]; ok {
				res = _abstracttopologytype.Name
			}
		}
	case "AffineCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _affinecstype, ok := stage.AffineCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _affinecstype.Name
			}
		}
	case "ArrayType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _arraytype, ok := stage.ArrayType_MetaDataProperty_reverseMap[inst]; ok {
				res = _arraytype.Name
			}
		}
	case "BagType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _bagtype, ok := stage.BagType_MetaDataProperty_reverseMap[inst]; ok {
				res = _bagtype.Name
			}
		}
	case "BaseUnitType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _baseunittype, ok := stage.BaseUnitType_MetaDataProperty_reverseMap[inst]; ok {
				res = _baseunittype.Name
			}
		}
	case "BoundedFeatureType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _boundedfeaturetype, ok := stage.BoundedFeatureType_MetaDataProperty_reverseMap[inst]; ok {
				res = _boundedfeaturetype.Name
			}
		}
	case "CartesianCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _cartesiancstype, ok := stage.CartesianCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _cartesiancstype.Name
			}
		}
	case "CompositeCurveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _compositecurvetype, ok := stage.CompositeCurveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _compositecurvetype.Name
			}
		}
	case "CompositeSolidType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _compositesolidtype, ok := stage.CompositeSolidType_MetaDataProperty_reverseMap[inst]; ok {
				res = _compositesolidtype.Name
			}
		}
	case "CompositeSurfaceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _compositesurfacetype, ok := stage.CompositeSurfaceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _compositesurfacetype.Name
			}
		}
	case "CompositeValueType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _compositevaluetype, ok := stage.CompositeValueType_MetaDataProperty_reverseMap[inst]; ok {
				res = _compositevaluetype.Name
			}
		}
	case "CompoundCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _compoundcrstype, ok := stage.CompoundCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _compoundcrstype.Name
			}
		}
	case "ConcatenatedOperationType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _concatenatedoperationtype, ok := stage.ConcatenatedOperationType_MetaDataProperty_reverseMap[inst]; ok {
				res = _concatenatedoperationtype.Name
			}
		}
	case "ConventionalUnitType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _conventionalunittype, ok := stage.ConventionalUnitType_MetaDataProperty_reverseMap[inst]; ok {
				res = _conventionalunittype.Name
			}
		}
	case "ConversionType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _conversiontype, ok := stage.ConversionType_MetaDataProperty_reverseMap[inst]; ok {
				res = _conversiontype.Name
			}
		}
	case "CoordinateSystemAxisType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _coordinatesystemaxistype, ok := stage.CoordinateSystemAxisType_MetaDataProperty_reverseMap[inst]; ok {
				res = _coordinatesystemaxistype.Name
			}
		}
	case "CurveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _curvetype, ok := stage.CurveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _curvetype.Name
			}
		}
	case "CylindricalCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _cylindricalcstype, ok := stage.CylindricalCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _cylindricalcstype.Name
			}
		}
	case "DefinitionBaseType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _definitionbasetype, ok := stage.DefinitionBaseType_MetaDataProperty_reverseMap[inst]; ok {
				res = _definitionbasetype.Name
			}
		}
	case "DefinitionProxyType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _definitionproxytype, ok := stage.DefinitionProxyType_MetaDataProperty_reverseMap[inst]; ok {
				res = _definitionproxytype.Name
			}
		}
	case "DefinitionType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _definitiontype, ok := stage.DefinitionType_MetaDataProperty_reverseMap[inst]; ok {
				res = _definitiontype.Name
			}
		}
	case "DerivedCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _derivedcrstype, ok := stage.DerivedCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _derivedcrstype.Name
			}
		}
	case "DerivedUnitType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _derivedunittype, ok := stage.DerivedUnitType_MetaDataProperty_reverseMap[inst]; ok {
				res = _derivedunittype.Name
			}
		}
	case "DictionaryType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _dictionarytype, ok := stage.DictionaryType_MetaDataProperty_reverseMap[inst]; ok {
				res = _dictionarytype.Name
			}
		}
	case "DirectedObservationAtDistanceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _directedobservationatdistancetype, ok := stage.DirectedObservationAtDistanceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _directedobservationatdistancetype.Name
			}
		}
	case "DirectedObservationType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _directedobservationtype, ok := stage.DirectedObservationType_MetaDataProperty_reverseMap[inst]; ok {
				res = _directedobservationtype.Name
			}
		}
	case "DiscreteCoverageType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _discretecoveragetype, ok := stage.DiscreteCoverageType_MetaDataProperty_reverseMap[inst]; ok {
				res = _discretecoveragetype.Name
			}
		}
	case "DynamicFeatureCollectionType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _dynamicfeaturecollectiontype, ok := stage.DynamicFeatureCollectionType_MetaDataProperty_reverseMap[inst]; ok {
				res = _dynamicfeaturecollectiontype.Name
			}
		}
	case "DynamicFeatureType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _dynamicfeaturetype, ok := stage.DynamicFeatureType_MetaDataProperty_reverseMap[inst]; ok {
				res = _dynamicfeaturetype.Name
			}
		}
	case "EdgeType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _edgetype, ok := stage.EdgeType_MetaDataProperty_reverseMap[inst]; ok {
				res = _edgetype.Name
			}
		}
	case "EllipsoidType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _ellipsoidtype, ok := stage.EllipsoidType_MetaDataProperty_reverseMap[inst]; ok {
				res = _ellipsoidtype.Name
			}
		}
	case "EllipsoidalCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _ellipsoidalcstype, ok := stage.EllipsoidalCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _ellipsoidalcstype.Name
			}
		}
	case "EngineeringCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _engineeringcrstype, ok := stage.EngineeringCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _engineeringcrstype.Name
			}
		}
	case "EngineeringDatumType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _engineeringdatumtype, ok := stage.EngineeringDatumType_MetaDataProperty_reverseMap[inst]; ok {
				res = _engineeringdatumtype.Name
			}
		}
	case "FaceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _facetype, ok := stage.FaceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _facetype.Name
			}
		}
	case "FeatureCollectionType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _featurecollectiontype, ok := stage.FeatureCollectionType_MetaDataProperty_reverseMap[inst]; ok {
				res = _featurecollectiontype.Name
			}
		}
	case "GeocentricCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _geocentriccrstype, ok := stage.GeocentricCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _geocentriccrstype.Name
			}
		}
	case "GeodeticCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _geodeticcrstype, ok := stage.GeodeticCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _geodeticcrstype.Name
			}
		}
	case "GeodeticDatumType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _geodeticdatumtype, ok := stage.GeodeticDatumType_MetaDataProperty_reverseMap[inst]; ok {
				res = _geodeticdatumtype.Name
			}
		}
	case "GeographicCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _geographiccrstype, ok := stage.GeographicCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _geographiccrstype.Name
			}
		}
	case "GeometricComplexType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _geometriccomplextype, ok := stage.GeometricComplexType_MetaDataProperty_reverseMap[inst]; ok {
				res = _geometriccomplextype.Name
			}
		}
	case "GridType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _gridtype, ok := stage.GridType_MetaDataProperty_reverseMap[inst]; ok {
				res = _gridtype.Name
			}
		}
	case "IdentifiedObjectType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _identifiedobjecttype, ok := stage.IdentifiedObjectType_MetaDataProperty_reverseMap[inst]; ok {
				res = _identifiedobjecttype.Name
			}
		}
	case "ImageCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _imagecrstype, ok := stage.ImageCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _imagecrstype.Name
			}
		}
	case "ImageDatumType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _imagedatumtype, ok := stage.ImageDatumType_MetaDataProperty_reverseMap[inst]; ok {
				res = _imagedatumtype.Name
			}
		}
	case "LineStringType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _linestringtype, ok := stage.LineStringType_MetaDataProperty_reverseMap[inst]; ok {
				res = _linestringtype.Name
			}
		}
	case "LinearCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _linearcstype, ok := stage.LinearCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _linearcstype.Name
			}
		}
	case "LinearRingType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _linearringtype, ok := stage.LinearRingType_MetaDataProperty_reverseMap[inst]; ok {
				res = _linearringtype.Name
			}
		}
	case "MovingObjectStatusType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _movingobjectstatustype, ok := stage.MovingObjectStatusType_MetaDataProperty_reverseMap[inst]; ok {
				res = _movingobjectstatustype.Name
			}
		}
	case "MultiCurveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _multicurvetype, ok := stage.MultiCurveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _multicurvetype.Name
			}
		}
	case "MultiGeometryType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _multigeometrytype, ok := stage.MultiGeometryType_MetaDataProperty_reverseMap[inst]; ok {
				res = _multigeometrytype.Name
			}
		}
	case "MultiPointType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _multipointtype, ok := stage.MultiPointType_MetaDataProperty_reverseMap[inst]; ok {
				res = _multipointtype.Name
			}
		}
	case "MultiSolidType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _multisolidtype, ok := stage.MultiSolidType_MetaDataProperty_reverseMap[inst]; ok {
				res = _multisolidtype.Name
			}
		}
	case "MultiSurfaceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _multisurfacetype, ok := stage.MultiSurfaceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _multisurfacetype.Name
			}
		}
	case "NodeType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _nodetype, ok := stage.NodeType_MetaDataProperty_reverseMap[inst]; ok {
				res = _nodetype.Name
			}
		}
	case "ObliqueCartesianCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _obliquecartesiancstype, ok := stage.ObliqueCartesianCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _obliquecartesiancstype.Name
			}
		}
	case "ObservationType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _observationtype, ok := stage.ObservationType_MetaDataProperty_reverseMap[inst]; ok {
				res = _observationtype.Name
			}
		}
	case "OperationMethodType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _operationmethodtype, ok := stage.OperationMethodType_MetaDataProperty_reverseMap[inst]; ok {
				res = _operationmethodtype.Name
			}
		}
	case "OperationParameterGroupType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _operationparametergrouptype, ok := stage.OperationParameterGroupType_MetaDataProperty_reverseMap[inst]; ok {
				res = _operationparametergrouptype.Name
			}
		}
	case "OperationParameterType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _operationparametertype, ok := stage.OperationParameterType_MetaDataProperty_reverseMap[inst]; ok {
				res = _operationparametertype.Name
			}
		}
	case "OrientableCurveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _orientablecurvetype, ok := stage.OrientableCurveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _orientablecurvetype.Name
			}
		}
	case "OrientableSurfaceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _orientablesurfacetype, ok := stage.OrientableSurfaceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _orientablesurfacetype.Name
			}
		}
	case "PassThroughOperationType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _passthroughoperationtype, ok := stage.PassThroughOperationType_MetaDataProperty_reverseMap[inst]; ok {
				res = _passthroughoperationtype.Name
			}
		}
	case "PointType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _pointtype, ok := stage.PointType_MetaDataProperty_reverseMap[inst]; ok {
				res = _pointtype.Name
			}
		}
	case "PolarCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _polarcstype, ok := stage.PolarCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _polarcstype.Name
			}
		}
	case "PolygonType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _polygontype, ok := stage.PolygonType_MetaDataProperty_reverseMap[inst]; ok {
				res = _polygontype.Name
			}
		}
	case "PrimeMeridianType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _primemeridiantype, ok := stage.PrimeMeridianType_MetaDataProperty_reverseMap[inst]; ok {
				res = _primemeridiantype.Name
			}
		}
	case "ProjectedCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _projectedcrstype, ok := stage.ProjectedCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _projectedcrstype.Name
			}
		}
	case "RectifiedGridType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _rectifiedgridtype, ok := stage.RectifiedGridType_MetaDataProperty_reverseMap[inst]; ok {
				res = _rectifiedgridtype.Name
			}
		}
	case "RingType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _ringtype, ok := stage.RingType_MetaDataProperty_reverseMap[inst]; ok {
				res = _ringtype.Name
			}
		}
	case "ShellType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _shelltype, ok := stage.ShellType_MetaDataProperty_reverseMap[inst]; ok {
				res = _shelltype.Name
			}
		}
	case "SolidType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _solidtype, ok := stage.SolidType_MetaDataProperty_reverseMap[inst]; ok {
				res = _solidtype.Name
			}
		}
	case "SphericalCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _sphericalcstype, ok := stage.SphericalCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _sphericalcstype.Name
			}
		}
	case "SurfaceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _surfacetype, ok := stage.SurfaceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _surfacetype.Name
			}
		}
	case "TemporalCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _temporalcrstype, ok := stage.TemporalCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _temporalcrstype.Name
			}
		}
	case "TemporalCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _temporalcstype, ok := stage.TemporalCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _temporalcstype.Name
			}
		}
	case "TemporalDatumBaseType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _temporaldatumbasetype, ok := stage.TemporalDatumBaseType_MetaDataProperty_reverseMap[inst]; ok {
				res = _temporaldatumbasetype.Name
			}
		}
	case "TemporalDatumType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _temporaldatumtype, ok := stage.TemporalDatumType_MetaDataProperty_reverseMap[inst]; ok {
				res = _temporaldatumtype.Name
			}
		}
	case "TimeCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timecstype, ok := stage.TimeCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timecstype.Name
			}
		}
	case "TimeCalendarEraType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timecalendareratype, ok := stage.TimeCalendarEraType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timecalendareratype.Name
			}
		}
	case "TimeCalendarType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timecalendartype, ok := stage.TimeCalendarType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timecalendartype.Name
			}
		}
	case "TimeClockType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timeclocktype, ok := stage.TimeClockType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timeclocktype.Name
			}
		}
	case "TimeCoordinateSystemType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timecoordinatesystemtype, ok := stage.TimeCoordinateSystemType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timecoordinatesystemtype.Name
			}
		}
	case "TimeEdgeType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timeedgetype, ok := stage.TimeEdgeType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timeedgetype.Name
			}
		}
	case "TimeInstantType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timeinstanttype, ok := stage.TimeInstantType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timeinstanttype.Name
			}
		}
	case "TimeNodeType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timenodetype, ok := stage.TimeNodeType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timenodetype.Name
			}
		}
	case "TimeOrdinalEraType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timeordinaleratype, ok := stage.TimeOrdinalEraType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timeordinaleratype.Name
			}
		}
	case "TimeOrdinalReferenceSystemType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timeordinalreferencesystemtype, ok := stage.TimeOrdinalReferenceSystemType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timeordinalreferencesystemtype.Name
			}
		}
	case "TimePeriodType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timeperiodtype, ok := stage.TimePeriodType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timeperiodtype.Name
			}
		}
	case "TimeReferenceSystemType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timereferencesystemtype, ok := stage.TimeReferenceSystemType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timereferencesystemtype.Name
			}
		}
	case "TimeTopologyComplexType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _timetopologycomplextype, ok := stage.TimeTopologyComplexType_MetaDataProperty_reverseMap[inst]; ok {
				res = _timetopologycomplextype.Name
			}
		}
	case "TinType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _tintype, ok := stage.TinType_MetaDataProperty_reverseMap[inst]; ok {
				res = _tintype.Name
			}
		}
	case "TopoComplexType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _topocomplextype, ok := stage.TopoComplexType_MetaDataProperty_reverseMap[inst]; ok {
				res = _topocomplextype.Name
			}
		}
	case "TopoCurveType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _topocurvetype, ok := stage.TopoCurveType_MetaDataProperty_reverseMap[inst]; ok {
				res = _topocurvetype.Name
			}
		}
	case "TopoPointType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _topopointtype, ok := stage.TopoPointType_MetaDataProperty_reverseMap[inst]; ok {
				res = _topopointtype.Name
			}
		}
	case "TopoSolidType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _toposolidtype, ok := stage.TopoSolidType_MetaDataProperty_reverseMap[inst]; ok {
				res = _toposolidtype.Name
			}
		}
	case "TopoSurfaceType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _toposurfacetype, ok := stage.TopoSurfaceType_MetaDataProperty_reverseMap[inst]; ok {
				res = _toposurfacetype.Name
			}
		}
	case "TopoVolumeType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _topovolumetype, ok := stage.TopoVolumeType_MetaDataProperty_reverseMap[inst]; ok {
				res = _topovolumetype.Name
			}
		}
	case "TransformationType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _transformationtype, ok := stage.TransformationType_MetaDataProperty_reverseMap[inst]; ok {
				res = _transformationtype.Name
			}
		}
	case "UnitDefinitionType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _unitdefinitiontype, ok := stage.UnitDefinitionType_MetaDataProperty_reverseMap[inst]; ok {
				res = _unitdefinitiontype.Name
			}
		}
	case "UserDefinedCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _userdefinedcstype, ok := stage.UserDefinedCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _userdefinedcstype.Name
			}
		}
	case "ValueArrayType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _valuearraytype, ok := stage.ValueArrayType_MetaDataProperty_reverseMap[inst]; ok {
				res = _valuearraytype.Name
			}
		}
	case "VerticalCRSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _verticalcrstype, ok := stage.VerticalCRSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _verticalcrstype.Name
			}
		}
	case "VerticalCSType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _verticalcstype, ok := stage.VerticalCSType_MetaDataProperty_reverseMap[inst]; ok {
				res = _verticalcstype.Name
			}
		}
	case "VerticalDatumType":
		switch reverseField.Fieldname {
		case "MetaDataProperty":
			if _verticaldatumtype, ok := stage.VerticalDatumType_MetaDataProperty_reverseMap[inst]; ok {
				res = _verticaldatumtype.Name
			}
		}
	}
	return
}

func (inst *MovingObjectStatusType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "MovingObjectStatus":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MovingObjectStatus_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "HistoryPropertyType":
		switch reverseField.Fieldname {
		case "MovingObjectStatus":
			if _historypropertytype, ok := stage.HistoryPropertyType_MovingObjectStatus_reverseMap[inst]; ok {
				res = _historypropertytype.Name
			}
		}
	}
	return
}

func (inst *MultiCurvePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MultiCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "MultiCurve":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiCurve_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "MultiCurve":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_MultiCurve_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *MultiGeometryPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MultiGeometryType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "MultiGeometry":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiGeometry_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "MultiGeometry":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_MultiGeometry_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *MultiPointPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MultiPointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "MultiPoint":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiPoint_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "MultiPoint":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_MultiPoint_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *MultiSolidPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MultiSolidType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "MultiSolid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiSolid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "MultiSolid":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_MultiSolid_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *MultiSurfacePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MultiSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "MultiSurface":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_MultiSurface_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "MultiSurface":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_MultiSurface_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *MultiplicityRange_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Multiplicity_Type":
		switch reverseField.Fieldname {
		case "Range_XSD":
			if _multiplicity_type, ok := stage.Multiplicity_Type_Range_XSD_reverseMap[inst]; ok {
				res = _multiplicity_type.Name
			}
		}
	}
	return
}

func (inst *MultiplicityRange_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Multiplicity_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Multiplicity_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *NodeOrEdgePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TopoSolidType":
		switch reverseField.Fieldname {
		case "Isolated":
			if _toposolidtype, ok := stage.TopoSolidType_Isolated_reverseMap[inst]; ok {
				res = _toposolidtype.Name
			}
		}
	}
	return
}

func (inst *NodePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "FaceType":
		switch reverseField.Fieldname {
		case "Isolated":
			if _facetype, ok := stage.FaceType_Isolated_reverseMap[inst]; ok {
				res = _facetype.Name
			}
		}
	}
	return
}

func (inst *NodeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Node":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Node_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "TopoPrimitiveArrayAssociationType":
		switch reverseField.Fieldname {
		case "Node":
			if _topoprimitivearrayassociationtype, ok := stage.TopoPrimitiveArrayAssociationType_Node_reverseMap[inst]; ok {
				res = _topoprimitivearrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *Number_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ObliqueCartesianCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ObliqueCartesianCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ObliqueCartesianCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ObliqueCartesianCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *ObservationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Observation":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Observation_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "FeatureArrayPropertyType":
		switch reverseField.Fieldname {
		case "Observation":
			if _featurearraypropertytype, ok := stage.FeatureArrayPropertyType_Observation_reverseMap[inst]; ok {
				res = _featurearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *OffsetCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "OffsetCurve":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_OffsetCurve_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveSegmentArrayPropertyType":
		switch reverseField.Fieldname {
		case "OffsetCurve":
			if _curvesegmentarraypropertytype, ok := stage.CurveSegmentArrayPropertyType_OffsetCurve_reverseMap[inst]; ok {
				res = _curvesegmentarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *OperationMethodPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OperationMethodType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "OperationMethod":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_OperationMethod_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *OperationParameterGroupPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OperationParameterGroupType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "OperationParameterGroup":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_OperationParameterGroup_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *OperationParameterPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OperationParameterType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "OperationParameter":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_OperationParameter_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *OperationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *OrientableCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "OrientableCurve":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_OrientableCurve_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveArrayPropertyType":
		switch reverseField.Fieldname {
		case "OrientableCurve":
			if _curvearraypropertytype, ok := stage.CurveArrayPropertyType_OrientableCurve_reverseMap[inst]; ok {
				res = _curvearraypropertytype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "OrientableCurve":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_OrientableCurve_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *OrientableSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "OrientableSurface":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_OrientableSurface_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "OrientableSurface":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_OrientableSurface_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SurfaceArrayPropertyType":
		switch reverseField.Fieldname {
		case "OrientableSurface":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_OrientableSurface_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *PT_FreeText_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PT_FreeText_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PT_LocaleContainer_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PT_LocaleContainer_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PT_Locale_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PT_Locale_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ParameterValueGroupType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ParameterValueGroup":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ParameterValueGroup_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *ParameterValueType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ParameterValue":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ParameterValue_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *PassThroughOperationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PassThroughOperationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "PassThroughOperation":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_PassThroughOperation_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *PointArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PointPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArcStringByBulgeType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _arcstringbybulgetype, ok := stage.ArcStringByBulgeType_PointProperty_reverseMap[inst]; ok {
				res = _arcstringbybulgetype.Name
			}
		case "PointRep":
			if _arcstringbybulgetype, ok := stage.ArcStringByBulgeType_PointRep_reverseMap[inst]; ok {
				res = _arcstringbybulgetype.Name
			}
		}
	case "ArcStringType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _arcstringtype, ok := stage.ArcStringType_PointProperty_reverseMap[inst]; ok {
				res = _arcstringtype.Name
			}
		case "PointRep":
			if _arcstringtype, ok := stage.ArcStringType_PointRep_reverseMap[inst]; ok {
				res = _arcstringtype.Name
			}
		}
	case "BSplineType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _bsplinetype, ok := stage.BSplineType_PointProperty_reverseMap[inst]; ok {
				res = _bsplinetype.Name
			}
		case "PointRep":
			if _bsplinetype, ok := stage.BSplineType_PointRep_reverseMap[inst]; ok {
				res = _bsplinetype.Name
			}
		}
	case "BezierType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _beziertype, ok := stage.BezierType_PointProperty_reverseMap[inst]; ok {
				res = _beziertype.Name
			}
		case "PointRep":
			if _beziertype, ok := stage.BezierType_PointRep_reverseMap[inst]; ok {
				res = _beziertype.Name
			}
		}
	case "CircleType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _circletype, ok := stage.CircleType_PointProperty_reverseMap[inst]; ok {
				res = _circletype.Name
			}
		case "PointRep":
			if _circletype, ok := stage.CircleType_PointRep_reverseMap[inst]; ok {
				res = _circletype.Name
			}
		}
	case "CubicSplineType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _cubicsplinetype, ok := stage.CubicSplineType_PointProperty_reverseMap[inst]; ok {
				res = _cubicsplinetype.Name
			}
		case "PointRep":
			if _cubicsplinetype, ok := stage.CubicSplineType_PointRep_reverseMap[inst]; ok {
				res = _cubicsplinetype.Name
			}
		}
	case "LineStringSegmentType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _linestringsegmenttype, ok := stage.LineStringSegmentType_PointProperty_reverseMap[inst]; ok {
				res = _linestringsegmenttype.Name
			}
		case "PointRep":
			if _linestringsegmenttype, ok := stage.LineStringSegmentType_PointRep_reverseMap[inst]; ok {
				res = _linestringsegmenttype.Name
			}
		}
	case "LineStringType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _linestringtype, ok := stage.LineStringType_PointProperty_reverseMap[inst]; ok {
				res = _linestringtype.Name
			}
		case "PointRep":
			if _linestringtype, ok := stage.LineStringType_PointRep_reverseMap[inst]; ok {
				res = _linestringtype.Name
			}
		}
	case "LinearRingType":
		switch reverseField.Fieldname {
		case "PointProperty":
			if _linearringtype, ok := stage.LinearRingType_PointProperty_reverseMap[inst]; ok {
				res = _linearringtype.Name
			}
		case "PointRep":
			if _linearringtype, ok := stage.LinearRingType_PointRep_reverseMap[inst]; ok {
				res = _linearringtype.Name
			}
		}
	case "MultiPointType":
		switch reverseField.Fieldname {
		case "PointMember":
			if _multipointtype, ok := stage.MultiPointType_PointMember_reverseMap[inst]; ok {
				res = _multipointtype.Name
			}
		}
	}
	return
}

func (inst *PointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Point":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Point_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Point":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Point_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "PointArrayPropertyType":
		switch reverseField.Fieldname {
		case "Point":
			if _pointarraypropertytype, ok := stage.PointArrayPropertyType_Point_reverseMap[inst]; ok {
				res = _pointarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *PolarCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PolarCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "PolarCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_PolarCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *PolygonPatchType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SurfacePatchArrayPropertyType":
		switch reverseField.Fieldname {
		case "PolygonPatch":
			if _surfacepatcharraypropertytype, ok := stage.SurfacePatchArrayPropertyType_PolygonPatch_reverseMap[inst]; ok {
				res = _surfacepatcharraypropertytype.Name
			}
		}
	}
	return
}

func (inst *PolygonType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Polygon":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Polygon_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Polygon":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Polygon_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SurfaceArrayPropertyType":
		switch reverseField.Fieldname {
		case "Polygon":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_Polygon_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *PrimeMeridianPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *PrimeMeridianType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "PrimeMeridian":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_PrimeMeridian_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *PriorityLocationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ProcedurePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ProjectedCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ProjectedCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ProjectedCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ProjectedCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *Quantity) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Quantity":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Quantity_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *QuantityExtentType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "QuantityExtent":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_QuantityExtent_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *QuantityPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RS_Identifier_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RS_Identifier_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RS_ReferenceSystem_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RangeSetType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Real_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_Medium_Type":
		switch reverseField.Fieldname {
		case "Density":
			if _md_medium_type, ok := stage.MD_Medium_Type_Density_reverseMap[inst]; ok {
				res = _md_medium_type.Name
			}
		}
	}
	return
}

func (inst *RecordType_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RecordType_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Record_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "DQ_QuantitativeResult_Type":
		switch reverseField.Fieldname {
		case "Value":
			if _dq_quantitativeresult_type, ok := stage.DQ_QuantitativeResult_Type_Value_reverseMap[inst]; ok {
				res = _dq_quantitativeresult_type.Name
			}
		}
	}
	return
}

func (inst *RectangleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SurfacePatchArrayPropertyType":
		switch reverseField.Fieldname {
		case "Rectangle":
			if _surfacepatcharraypropertytype, ok := stage.SurfacePatchArrayPropertyType_Rectangle_reverseMap[inst]; ok {
				res = _surfacepatcharraypropertytype.Name
			}
		}
	}
	return
}

func (inst *RectifiedGridType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "RectifiedGrid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_RectifiedGrid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "RectifiedGrid":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_RectifiedGrid_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *ReferenceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RelatedTimeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AbstractTimeGeometricPrimitiveType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _abstracttimegeometricprimitivetype, ok := stage.AbstractTimeGeometricPrimitiveType_RelatedTime_reverseMap[inst]; ok {
				res = _abstracttimegeometricprimitivetype.Name
			}
		}
	case "AbstractTimePrimitiveType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _abstracttimeprimitivetype, ok := stage.AbstractTimePrimitiveType_RelatedTime_reverseMap[inst]; ok {
				res = _abstracttimeprimitivetype.Name
			}
		}
	case "AbstractTimeTopologyPrimitiveType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _abstracttimetopologyprimitivetype, ok := stage.AbstractTimeTopologyPrimitiveType_RelatedTime_reverseMap[inst]; ok {
				res = _abstracttimetopologyprimitivetype.Name
			}
		}
	case "TimeEdgeType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _timeedgetype, ok := stage.TimeEdgeType_RelatedTime_reverseMap[inst]; ok {
				res = _timeedgetype.Name
			}
		}
	case "TimeInstantType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _timeinstanttype, ok := stage.TimeInstantType_RelatedTime_reverseMap[inst]; ok {
				res = _timeinstanttype.Name
			}
		}
	case "TimeNodeType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _timenodetype, ok := stage.TimeNodeType_RelatedTime_reverseMap[inst]; ok {
				res = _timenodetype.Name
			}
		}
	case "TimeOrdinalEraType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _timeordinaleratype, ok := stage.TimeOrdinalEraType_RelatedTime_reverseMap[inst]; ok {
				res = _timeordinaleratype.Name
			}
		}
	case "TimePeriodType":
		switch reverseField.Fieldname {
		case "RelatedTime":
			if _timeperiodtype, ok := stage.TimePeriodType_RelatedTime_reverseMap[inst]; ok {
				res = _timeperiodtype.Name
			}
		}
	}
	return
}

func (inst *ResultType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RingPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *RingType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Ring":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Ring_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "CurveArrayPropertyType":
		switch reverseField.Fieldname {
		case "Ring":
			if _curvearraypropertytype, ok := stage.CurveArrayPropertyType_Ring_reverseMap[inst]; ok {
				res = _curvearraypropertytype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Ring":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Ring_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *SC_CRS_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Scale_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ScopedName_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SecondDefiningParameter) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SecondDefiningParameter_1) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SequenceRuleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ShellPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SolidType":
		switch reverseField.Fieldname {
		case "Interior":
			if _solidtype, ok := stage.SolidType_Interior_reverseMap[inst]; ok {
				res = _solidtype.Name
			}
		}
	}
	return
}

func (inst *ShellType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Shell":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Shell_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Shell":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Shell_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SurfaceArrayPropertyType":
		switch reverseField.Fieldname {
		case "Shell":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_Shell_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *SingleCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CompoundCRSType":
		switch reverseField.Fieldname {
		case "ComponentReferenceSystem":
			if _compoundcrstype, ok := stage.CompoundCRSType_ComponentReferenceSystem_reverseMap[inst]; ok {
				res = _compoundcrstype.Name
			}
		case "IncludesSingleCRS":
			if _compoundcrstype, ok := stage.CompoundCRSType_IncludesSingleCRS_reverseMap[inst]; ok {
				res = _compoundcrstype.Name
			}
		}
	}
	return
}

func (inst *SingleOperationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SolidArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SolidPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CompositeSolidType":
		switch reverseField.Fieldname {
		case "SolidMember":
			if _compositesolidtype, ok := stage.CompositeSolidType_SolidMember_reverseMap[inst]; ok {
				res = _compositesolidtype.Name
			}
		}
	case "MultiSolidType":
		switch reverseField.Fieldname {
		case "SolidMember":
			if _multisolidtype, ok := stage.MultiSolidType_SolidMember_reverseMap[inst]; ok {
				res = _multisolidtype.Name
			}
		}
	}
	return
}

func (inst *SolidType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Solid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Solid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Solid":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Solid_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SolidArrayPropertyType":
		switch reverseField.Fieldname {
		case "Solid":
			if _solidarraypropertytype, ok := stage.SolidArrayPropertyType_Solid_reverseMap[inst]; ok {
				res = _solidarraypropertytype.Name
			}
		}
	}
	return
}

func (inst *SphereType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SurfacePatchArrayPropertyType":
		switch reverseField.Fieldname {
		case "Sphere":
			if _surfacepatcharraypropertytype, ok := stage.SurfacePatchArrayPropertyType_Sphere_reverseMap[inst]; ok {
				res = _surfacepatcharraypropertytype.Name
			}
		}
	}
	return
}

func (inst *SphericalCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SphericalCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "SphericalCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_SphericalCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *StringOrRefType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SurfaceArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SurfacePatchArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *SurfacePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CompositeSurfaceType":
		switch reverseField.Fieldname {
		case "SurfaceMember":
			if _compositesurfacetype, ok := stage.CompositeSurfaceType_SurfaceMember_reverseMap[inst]; ok {
				res = _compositesurfacetype.Name
			}
		}
	case "MultiSurfaceType":
		switch reverseField.Fieldname {
		case "SurfaceMember":
			if _multisurfacetype, ok := stage.MultiSurfaceType_SurfaceMember_reverseMap[inst]; ok {
				res = _multisurfacetype.Name
			}
		}
	case "ShellType":
		switch reverseField.Fieldname {
		case "SurfaceMember":
			if _shelltype, ok := stage.ShellType_SurfaceMember_reverseMap[inst]; ok {
				res = _shelltype.Name
			}
		}
	}
	return
}

func (inst *SurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Surface":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Surface_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "PolyhedralSurface":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_PolyhedralSurface_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		case "TriangulatedSurface":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TriangulatedSurface_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Surface":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Surface_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		case "PolyhedralSurface":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_PolyhedralSurface_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		case "TriangulatedSurface":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_TriangulatedSurface_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SurfaceArrayPropertyType":
		switch reverseField.Fieldname {
		case "Surface":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_Surface_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		case "PolyhedralSurface":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_PolyhedralSurface_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		case "TriangulatedSurface":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_TriangulatedSurface_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *TM_PeriodDuration_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TM_Primitive_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TargetPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TemporalCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TemporalCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TemporalCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TemporalCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TemporalCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TemporalCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TemporalCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TemporalCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TemporalDatumBaseType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TemporalDatumPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TemporalDatumType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TemporalDatum":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TemporalDatum_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeCalendarEraPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TimeCalendarType":
		switch reverseField.Fieldname {
		case "ReferenceFrame":
			if _timecalendartype, ok := stage.TimeCalendarType_ReferenceFrame_reverseMap[inst]; ok {
				res = _timecalendartype.Name
			}
		}
	}
	return
}

func (inst *TimeCalendarEraType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeCalendarPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TimeClockType":
		switch reverseField.Fieldname {
		case "DateBasis":
			if _timeclocktype, ok := stage.TimeClockType_DateBasis_reverseMap[inst]; ok {
				res = _timeclocktype.Name
			}
		}
	}
	return
}

func (inst *TimeCalendarType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeCalendar":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeCalendar_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeClockPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeClockType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeClock":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeClock_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeCoordinateSystemType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeCoordinateSystem":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeCoordinateSystem_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeEdgePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TimeNodeType":
		switch reverseField.Fieldname {
		case "PreviousEdge":
			if _timenodetype, ok := stage.TimeNodeType_PreviousEdge_reverseMap[inst]; ok {
				res = _timenodetype.Name
			}
		case "NextEdge":
			if _timenodetype, ok := stage.TimeNodeType_NextEdge_reverseMap[inst]; ok {
				res = _timenodetype.Name
			}
		}
	}
	return
}

func (inst *TimeEdgeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeEdge":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeEdge_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeInstantPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeInstantType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeInstant":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeInstant_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeIntervalLengthType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeNodePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeNodeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeNode":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeNode_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeOrdinalEraPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TimeOrdinalEraType":
		switch reverseField.Fieldname {
		case "Member":
			if _timeordinaleratype, ok := stage.TimeOrdinalEraType_Member_reverseMap[inst]; ok {
				res = _timeordinaleratype.Name
			}
		}
	case "TimeOrdinalReferenceSystemType":
		switch reverseField.Fieldname {
		case "Component":
			if _timeordinalreferencesystemtype, ok := stage.TimeOrdinalReferenceSystemType_Component_reverseMap[inst]; ok {
				res = _timeordinalreferencesystemtype.Name
			}
		}
	}
	return
}

func (inst *TimeOrdinalEraType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeOrdinalReferenceSystemType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeOrdinalReferenceSystem":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeOrdinalReferenceSystem_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimePeriodPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimePeriodType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimePeriod":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimePeriod_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimePositionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimePrimitivePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeReferenceSystemType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeReferenceSystem":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeReferenceSystem_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeTopologyComplexPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TimeTopologyComplexType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TimeTopologyComplex":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TimeTopologyComplex_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TimeTopologyPrimitivePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TimeTopologyComplexType":
		switch reverseField.Fieldname {
		case "Primitive":
			if _timetopologycomplextype, ok := stage.TimeTopologyComplexType_Primitive_reverseMap[inst]; ok {
				res = _timetopologycomplextype.Name
			}
		}
	}
	return
}

func (inst *TinType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Tin":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Tin_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "GeometryArrayPropertyType":
		switch reverseField.Fieldname {
		case "Tin":
			if _geometryarraypropertytype, ok := stage.GeometryArrayPropertyType_Tin_reverseMap[inst]; ok {
				res = _geometryarraypropertytype.Name
			}
		}
	case "SurfaceArrayPropertyType":
		switch reverseField.Fieldname {
		case "Tin":
			if _surfacearraypropertytype, ok := stage.SurfaceArrayPropertyType_Tin_reverseMap[inst]; ok {
				res = _surfacearraypropertytype.Name
			}
		}
	}
	return
}

func (inst *TopoComplexPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TopoComplexType":
		switch reverseField.Fieldname {
		case "SuperComplex":
			if _topocomplextype, ok := stage.TopoComplexType_SuperComplex_reverseMap[inst]; ok {
				res = _topocomplextype.Name
			}
		case "SubComplex":
			if _topocomplextype, ok := stage.TopoComplexType_SubComplex_reverseMap[inst]; ok {
				res = _topocomplextype.Name
			}
		}
	}
	return
}

func (inst *TopoComplexType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TopoComplex":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TopoComplex_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TopoCurvePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoCurveType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoPointPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoPointType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoPrimitiveArrayAssociationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoPrimitiveMemberType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "TopoComplexType":
		switch reverseField.Fieldname {
		case "TopoPrimitiveMember":
			if _topocomplextype, ok := stage.TopoComplexType_TopoPrimitiveMember_reverseMap[inst]; ok {
				res = _topocomplextype.Name
			}
		}
	}
	return
}

func (inst *TopoSolidPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoSolidType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "TopoSolid":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_TopoSolid_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "TopoPrimitiveArrayAssociationType":
		switch reverseField.Fieldname {
		case "TopoSolid":
			if _topoprimitivearrayassociationtype, ok := stage.TopoPrimitiveArrayAssociationType_TopoSolid_reverseMap[inst]; ok {
				res = _topoprimitivearrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TopoSurfacePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoSurfaceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoVolumePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TopoVolumeType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TransformationPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TransformationType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "Transformation":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_Transformation_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *TriangleType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "SurfacePatchArrayPropertyType":
		switch reverseField.Fieldname {
		case "Triangle":
			if _surfacepatcharraypropertytype, ok := stage.SurfacePatchArrayPropertyType_Triangle_reverseMap[inst]; ok {
				res = _surfacepatcharraypropertytype.Name
			}
		}
	}
	return
}

func (inst *TypeName_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TypeName_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *URL_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UnitDefinitionType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "UnitDefinition":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_UnitDefinition_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *UnitOfMeasureType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UnitOfMeasure_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UnlimitedInteger_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UnlimitedInteger_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UomAngle_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UomArea_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UomLength_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UomScale_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UomTime_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UomVelocity_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UomVolume_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UserDefinedCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *UserDefinedCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "UserDefinedCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_UserDefinedCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *ValueArrayPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ValueArrayType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "ValueArray":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_ValueArray_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	case "RangeSetType":
		switch reverseField.Fieldname {
		case "ValueArray":
			if _rangesettype, ok := stage.RangeSetType_ValueArray_reverseMap[inst]; ok {
				res = _rangesettype.Name
			}
		}
	}
	return
}

func (inst *ValuePropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "CompositeValueType":
		switch reverseField.Fieldname {
		case "ValueComponent":
			if _compositevaluetype, ok := stage.CompositeValueType_ValueComponent_reverseMap[inst]; ok {
				res = _compositevaluetype.Name
			}
		}
	case "ValueArrayType":
		switch reverseField.Fieldname {
		case "ValueComponent":
			if _valuearraytype, ok := stage.ValueArrayType_ValueComponent_reverseMap[inst]; ok {
				res = _valuearraytype.Name
			}
		}
	}
	return
}

func (inst *VectorType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AffinePlacementType":
		switch reverseField.Fieldname {
		case "RefDirection":
			if _affineplacementtype, ok := stage.AffinePlacementType_RefDirection_reverseMap[inst]; ok {
				res = _affineplacementtype.Name
			}
		}
	case "ArcStringByBulgeType":
		switch reverseField.Fieldname {
		case "Normal":
			if _arcstringbybulgetype, ok := stage.ArcStringByBulgeType_Normal_reverseMap[inst]; ok {
				res = _arcstringbybulgetype.Name
			}
		}
	case "RectifiedGridType":
		switch reverseField.Fieldname {
		case "OffsetVector":
			if _rectifiedgridtype, ok := stage.RectifiedGridType_OffsetVector_reverseMap[inst]; ok {
				res = _rectifiedgridtype.Name
			}
		}
	}
	return
}

func (inst *VerticalCRSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *VerticalCRSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "VerticalCRS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_VerticalCRS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *VerticalCSPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *VerticalCSType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "VerticalCS":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_VerticalCS_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}

func (inst *VerticalDatumPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *VerticalDatumType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArrayAssociationType":
		switch reverseField.Fieldname {
		case "VerticalDatum":
			if _arrayassociationtype, ok := stage.ArrayAssociationType_VerticalDatum_reverseMap[inst]; ok {
				res = _arrayassociationtype.Name
			}
		}
	}
	return
}
