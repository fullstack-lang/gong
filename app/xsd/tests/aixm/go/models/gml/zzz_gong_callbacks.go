// generated code - do not edit
package gml

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
func (a_row *A_Row) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_RowCreateCallback != nil {
		stage.OnAfterA_RowCreateCallback.OnAfterCreate(stage, a_row)
	}
}

func (a_row *A_Row) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_RowUpdateCallback != nil {
		var frontA_Row *A_Row
		if front != nil {
			frontA_Row, _ = front.(*A_Row)
		}
		stage.OnAfterA_RowUpdateCallback.OnAfterUpdate(stage, a_row, frontA_Row)
	}
}

func (a_row *A_Row) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_RowDeleteCallback != nil {
		var frontA_Row *A_Row
		if front != nil {
			frontA_Row, _ = front.(*A_Row)
		}
		stage.OnAfterA_RowDeleteCallback.OnAfterDelete(stage, a_row, frontA_Row)
	}
}

func (a_controlpoint *A_controlPoint) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_controlPointCreateCallback != nil {
		stage.OnAfterA_controlPointCreateCallback.OnAfterCreate(stage, a_controlpoint)
	}
}

func (a_controlpoint *A_controlPoint) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_controlPointUpdateCallback != nil {
		var frontA_controlPoint *A_controlPoint
		if front != nil {
			frontA_controlPoint, _ = front.(*A_controlPoint)
		}
		stage.OnAfterA_controlPointUpdateCallback.OnAfterUpdate(stage, a_controlpoint, frontA_controlPoint)
	}
}

func (a_controlpoint *A_controlPoint) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_controlPointDeleteCallback != nil {
		var frontA_controlPoint *A_controlPoint
		if front != nil {
			frontA_controlPoint, _ = front.(*A_controlPoint)
		}
		stage.OnAfterA_controlPointDeleteCallback.OnAfterDelete(stage, a_controlpoint, frontA_controlPoint)
	}
}

func (a_reflocation *A_refLocation) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_refLocationCreateCallback != nil {
		stage.OnAfterA_refLocationCreateCallback.OnAfterCreate(stage, a_reflocation)
	}
}

func (a_reflocation *A_refLocation) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_refLocationUpdateCallback != nil {
		var frontA_refLocation *A_refLocation
		if front != nil {
			frontA_refLocation, _ = front.(*A_refLocation)
		}
		stage.OnAfterA_refLocationUpdateCallback.OnAfterUpdate(stage, a_reflocation, frontA_refLocation)
	}
}

func (a_reflocation *A_refLocation) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_refLocationDeleteCallback != nil {
		var frontA_refLocation *A_refLocation
		if front != nil {
			frontA_refLocation, _ = front.(*A_refLocation)
		}
		stage.OnAfterA_refLocationDeleteCallback.OnAfterDelete(stage, a_reflocation, frontA_refLocation)
	}
}

func (a_rows *A_rows) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterA_rowsCreateCallback != nil {
		stage.OnAfterA_rowsCreateCallback.OnAfterCreate(stage, a_rows)
	}
}

func (a_rows *A_rows) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_rowsUpdateCallback != nil {
		var frontA_rows *A_rows
		if front != nil {
			frontA_rows, _ = front.(*A_rows)
		}
		stage.OnAfterA_rowsUpdateCallback.OnAfterUpdate(stage, a_rows, frontA_rows)
	}
}

func (a_rows *A_rows) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterA_rowsDeleteCallback != nil {
		var frontA_rows *A_rows
		if front != nil {
			frontA_rows, _ = front.(*A_rows)
		}
		stage.OnAfterA_rowsDeleteCallback.OnAfterDelete(stage, a_rows, frontA_rows)
	}
}

func (abstractcrstype *AbstractCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCRSTypeCreateCallback != nil {
		stage.OnAfterAbstractCRSTypeCreateCallback.OnAfterCreate(stage, abstractcrstype)
	}
}

func (abstractcrstype *AbstractCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCRSTypeUpdateCallback != nil {
		var frontAbstractCRSType *AbstractCRSType
		if front != nil {
			frontAbstractCRSType, _ = front.(*AbstractCRSType)
		}
		stage.OnAfterAbstractCRSTypeUpdateCallback.OnAfterUpdate(stage, abstractcrstype, frontAbstractCRSType)
	}
}

func (abstractcrstype *AbstractCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCRSTypeDeleteCallback != nil {
		var frontAbstractCRSType *AbstractCRSType
		if front != nil {
			frontAbstractCRSType, _ = front.(*AbstractCRSType)
		}
		stage.OnAfterAbstractCRSTypeDeleteCallback.OnAfterDelete(stage, abstractcrstype, frontAbstractCRSType)
	}
}

func (abstractcontinuouscoveragetype *AbstractContinuousCoverageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractContinuousCoverageTypeCreateCallback != nil {
		stage.OnAfterAbstractContinuousCoverageTypeCreateCallback.OnAfterCreate(stage, abstractcontinuouscoveragetype)
	}
}

func (abstractcontinuouscoveragetype *AbstractContinuousCoverageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractContinuousCoverageTypeUpdateCallback != nil {
		var frontAbstractContinuousCoverageType *AbstractContinuousCoverageType
		if front != nil {
			frontAbstractContinuousCoverageType, _ = front.(*AbstractContinuousCoverageType)
		}
		stage.OnAfterAbstractContinuousCoverageTypeUpdateCallback.OnAfterUpdate(stage, abstractcontinuouscoveragetype, frontAbstractContinuousCoverageType)
	}
}

func (abstractcontinuouscoveragetype *AbstractContinuousCoverageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractContinuousCoverageTypeDeleteCallback != nil {
		var frontAbstractContinuousCoverageType *AbstractContinuousCoverageType
		if front != nil {
			frontAbstractContinuousCoverageType, _ = front.(*AbstractContinuousCoverageType)
		}
		stage.OnAfterAbstractContinuousCoverageTypeDeleteCallback.OnAfterDelete(stage, abstractcontinuouscoveragetype, frontAbstractContinuousCoverageType)
	}
}

func (abstractcoordinateoperationtype *AbstractCoordinateOperationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCoordinateOperationTypeCreateCallback != nil {
		stage.OnAfterAbstractCoordinateOperationTypeCreateCallback.OnAfterCreate(stage, abstractcoordinateoperationtype)
	}
}

func (abstractcoordinateoperationtype *AbstractCoordinateOperationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCoordinateOperationTypeUpdateCallback != nil {
		var frontAbstractCoordinateOperationType *AbstractCoordinateOperationType
		if front != nil {
			frontAbstractCoordinateOperationType, _ = front.(*AbstractCoordinateOperationType)
		}
		stage.OnAfterAbstractCoordinateOperationTypeUpdateCallback.OnAfterUpdate(stage, abstractcoordinateoperationtype, frontAbstractCoordinateOperationType)
	}
}

func (abstractcoordinateoperationtype *AbstractCoordinateOperationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCoordinateOperationTypeDeleteCallback != nil {
		var frontAbstractCoordinateOperationType *AbstractCoordinateOperationType
		if front != nil {
			frontAbstractCoordinateOperationType, _ = front.(*AbstractCoordinateOperationType)
		}
		stage.OnAfterAbstractCoordinateOperationTypeDeleteCallback.OnAfterDelete(stage, abstractcoordinateoperationtype, frontAbstractCoordinateOperationType)
	}
}

func (abstractcoordinatesystemtype *AbstractCoordinateSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCoordinateSystemTypeCreateCallback != nil {
		stage.OnAfterAbstractCoordinateSystemTypeCreateCallback.OnAfterCreate(stage, abstractcoordinatesystemtype)
	}
}

func (abstractcoordinatesystemtype *AbstractCoordinateSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCoordinateSystemTypeUpdateCallback != nil {
		var frontAbstractCoordinateSystemType *AbstractCoordinateSystemType
		if front != nil {
			frontAbstractCoordinateSystemType, _ = front.(*AbstractCoordinateSystemType)
		}
		stage.OnAfterAbstractCoordinateSystemTypeUpdateCallback.OnAfterUpdate(stage, abstractcoordinatesystemtype, frontAbstractCoordinateSystemType)
	}
}

func (abstractcoordinatesystemtype *AbstractCoordinateSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCoordinateSystemTypeDeleteCallback != nil {
		var frontAbstractCoordinateSystemType *AbstractCoordinateSystemType
		if front != nil {
			frontAbstractCoordinateSystemType, _ = front.(*AbstractCoordinateSystemType)
		}
		stage.OnAfterAbstractCoordinateSystemTypeDeleteCallback.OnAfterDelete(stage, abstractcoordinatesystemtype, frontAbstractCoordinateSystemType)
	}
}

func (abstractcoveragetype *AbstractCoverageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCoverageTypeCreateCallback != nil {
		stage.OnAfterAbstractCoverageTypeCreateCallback.OnAfterCreate(stage, abstractcoveragetype)
	}
}

func (abstractcoveragetype *AbstractCoverageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCoverageTypeUpdateCallback != nil {
		var frontAbstractCoverageType *AbstractCoverageType
		if front != nil {
			frontAbstractCoverageType, _ = front.(*AbstractCoverageType)
		}
		stage.OnAfterAbstractCoverageTypeUpdateCallback.OnAfterUpdate(stage, abstractcoveragetype, frontAbstractCoverageType)
	}
}

func (abstractcoveragetype *AbstractCoverageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCoverageTypeDeleteCallback != nil {
		var frontAbstractCoverageType *AbstractCoverageType
		if front != nil {
			frontAbstractCoverageType, _ = front.(*AbstractCoverageType)
		}
		stage.OnAfterAbstractCoverageTypeDeleteCallback.OnAfterDelete(stage, abstractcoveragetype, frontAbstractCoverageType)
	}
}

func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCurveSegmentTypeCreateCallback != nil {
		stage.OnAfterAbstractCurveSegmentTypeCreateCallback.OnAfterCreate(stage, abstractcurvesegmenttype)
	}
}

func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveSegmentTypeUpdateCallback != nil {
		var frontAbstractCurveSegmentType *AbstractCurveSegmentType
		if front != nil {
			frontAbstractCurveSegmentType, _ = front.(*AbstractCurveSegmentType)
		}
		stage.OnAfterAbstractCurveSegmentTypeUpdateCallback.OnAfterUpdate(stage, abstractcurvesegmenttype, frontAbstractCurveSegmentType)
	}
}

func (abstractcurvesegmenttype *AbstractCurveSegmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveSegmentTypeDeleteCallback != nil {
		var frontAbstractCurveSegmentType *AbstractCurveSegmentType
		if front != nil {
			frontAbstractCurveSegmentType, _ = front.(*AbstractCurveSegmentType)
		}
		stage.OnAfterAbstractCurveSegmentTypeDeleteCallback.OnAfterDelete(stage, abstractcurvesegmenttype, frontAbstractCurveSegmentType)
	}
}

func (abstractcurvetype *AbstractCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractCurveTypeCreateCallback != nil {
		stage.OnAfterAbstractCurveTypeCreateCallback.OnAfterCreate(stage, abstractcurvetype)
	}
}

func (abstractcurvetype *AbstractCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveTypeUpdateCallback != nil {
		var frontAbstractCurveType *AbstractCurveType
		if front != nil {
			frontAbstractCurveType, _ = front.(*AbstractCurveType)
		}
		stage.OnAfterAbstractCurveTypeUpdateCallback.OnAfterUpdate(stage, abstractcurvetype, frontAbstractCurveType)
	}
}

func (abstractcurvetype *AbstractCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractCurveTypeDeleteCallback != nil {
		var frontAbstractCurveType *AbstractCurveType
		if front != nil {
			frontAbstractCurveType, _ = front.(*AbstractCurveType)
		}
		stage.OnAfterAbstractCurveTypeDeleteCallback.OnAfterDelete(stage, abstractcurvetype, frontAbstractCurveType)
	}
}

func (abstractdq_completeness_type *AbstractDQ_Completeness_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDQ_Completeness_TypeCreateCallback != nil {
		stage.OnAfterAbstractDQ_Completeness_TypeCreateCallback.OnAfterCreate(stage, abstractdq_completeness_type)
	}
}

func (abstractdq_completeness_type *AbstractDQ_Completeness_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_Completeness_TypeUpdateCallback != nil {
		var frontAbstractDQ_Completeness_Type *AbstractDQ_Completeness_Type
		if front != nil {
			frontAbstractDQ_Completeness_Type, _ = front.(*AbstractDQ_Completeness_Type)
		}
		stage.OnAfterAbstractDQ_Completeness_TypeUpdateCallback.OnAfterUpdate(stage, abstractdq_completeness_type, frontAbstractDQ_Completeness_Type)
	}
}

func (abstractdq_completeness_type *AbstractDQ_Completeness_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_Completeness_TypeDeleteCallback != nil {
		var frontAbstractDQ_Completeness_Type *AbstractDQ_Completeness_Type
		if front != nil {
			frontAbstractDQ_Completeness_Type, _ = front.(*AbstractDQ_Completeness_Type)
		}
		stage.OnAfterAbstractDQ_Completeness_TypeDeleteCallback.OnAfterDelete(stage, abstractdq_completeness_type, frontAbstractDQ_Completeness_Type)
	}
}

func (abstractdq_element_type *AbstractDQ_Element_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDQ_Element_TypeCreateCallback != nil {
		stage.OnAfterAbstractDQ_Element_TypeCreateCallback.OnAfterCreate(stage, abstractdq_element_type)
	}
}

func (abstractdq_element_type *AbstractDQ_Element_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_Element_TypeUpdateCallback != nil {
		var frontAbstractDQ_Element_Type *AbstractDQ_Element_Type
		if front != nil {
			frontAbstractDQ_Element_Type, _ = front.(*AbstractDQ_Element_Type)
		}
		stage.OnAfterAbstractDQ_Element_TypeUpdateCallback.OnAfterUpdate(stage, abstractdq_element_type, frontAbstractDQ_Element_Type)
	}
}

func (abstractdq_element_type *AbstractDQ_Element_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_Element_TypeDeleteCallback != nil {
		var frontAbstractDQ_Element_Type *AbstractDQ_Element_Type
		if front != nil {
			frontAbstractDQ_Element_Type, _ = front.(*AbstractDQ_Element_Type)
		}
		stage.OnAfterAbstractDQ_Element_TypeDeleteCallback.OnAfterDelete(stage, abstractdq_element_type, frontAbstractDQ_Element_Type)
	}
}

func (abstractdq_logicalconsistency_type *AbstractDQ_LogicalConsistency_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDQ_LogicalConsistency_TypeCreateCallback != nil {
		stage.OnAfterAbstractDQ_LogicalConsistency_TypeCreateCallback.OnAfterCreate(stage, abstractdq_logicalconsistency_type)
	}
}

func (abstractdq_logicalconsistency_type *AbstractDQ_LogicalConsistency_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_LogicalConsistency_TypeUpdateCallback != nil {
		var frontAbstractDQ_LogicalConsistency_Type *AbstractDQ_LogicalConsistency_Type
		if front != nil {
			frontAbstractDQ_LogicalConsistency_Type, _ = front.(*AbstractDQ_LogicalConsistency_Type)
		}
		stage.OnAfterAbstractDQ_LogicalConsistency_TypeUpdateCallback.OnAfterUpdate(stage, abstractdq_logicalconsistency_type, frontAbstractDQ_LogicalConsistency_Type)
	}
}

func (abstractdq_logicalconsistency_type *AbstractDQ_LogicalConsistency_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_LogicalConsistency_TypeDeleteCallback != nil {
		var frontAbstractDQ_LogicalConsistency_Type *AbstractDQ_LogicalConsistency_Type
		if front != nil {
			frontAbstractDQ_LogicalConsistency_Type, _ = front.(*AbstractDQ_LogicalConsistency_Type)
		}
		stage.OnAfterAbstractDQ_LogicalConsistency_TypeDeleteCallback.OnAfterDelete(stage, abstractdq_logicalconsistency_type, frontAbstractDQ_LogicalConsistency_Type)
	}
}

func (abstractdq_positionalaccuracy_type *AbstractDQ_PositionalAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDQ_PositionalAccuracy_TypeCreateCallback != nil {
		stage.OnAfterAbstractDQ_PositionalAccuracy_TypeCreateCallback.OnAfterCreate(stage, abstractdq_positionalaccuracy_type)
	}
}

func (abstractdq_positionalaccuracy_type *AbstractDQ_PositionalAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_PositionalAccuracy_TypeUpdateCallback != nil {
		var frontAbstractDQ_PositionalAccuracy_Type *AbstractDQ_PositionalAccuracy_Type
		if front != nil {
			frontAbstractDQ_PositionalAccuracy_Type, _ = front.(*AbstractDQ_PositionalAccuracy_Type)
		}
		stage.OnAfterAbstractDQ_PositionalAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, abstractdq_positionalaccuracy_type, frontAbstractDQ_PositionalAccuracy_Type)
	}
}

func (abstractdq_positionalaccuracy_type *AbstractDQ_PositionalAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_PositionalAccuracy_TypeDeleteCallback != nil {
		var frontAbstractDQ_PositionalAccuracy_Type *AbstractDQ_PositionalAccuracy_Type
		if front != nil {
			frontAbstractDQ_PositionalAccuracy_Type, _ = front.(*AbstractDQ_PositionalAccuracy_Type)
		}
		stage.OnAfterAbstractDQ_PositionalAccuracy_TypeDeleteCallback.OnAfterDelete(stage, abstractdq_positionalaccuracy_type, frontAbstractDQ_PositionalAccuracy_Type)
	}
}

func (abstractdq_result_type *AbstractDQ_Result_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDQ_Result_TypeCreateCallback != nil {
		stage.OnAfterAbstractDQ_Result_TypeCreateCallback.OnAfterCreate(stage, abstractdq_result_type)
	}
}

func (abstractdq_result_type *AbstractDQ_Result_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_Result_TypeUpdateCallback != nil {
		var frontAbstractDQ_Result_Type *AbstractDQ_Result_Type
		if front != nil {
			frontAbstractDQ_Result_Type, _ = front.(*AbstractDQ_Result_Type)
		}
		stage.OnAfterAbstractDQ_Result_TypeUpdateCallback.OnAfterUpdate(stage, abstractdq_result_type, frontAbstractDQ_Result_Type)
	}
}

func (abstractdq_result_type *AbstractDQ_Result_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_Result_TypeDeleteCallback != nil {
		var frontAbstractDQ_Result_Type *AbstractDQ_Result_Type
		if front != nil {
			frontAbstractDQ_Result_Type, _ = front.(*AbstractDQ_Result_Type)
		}
		stage.OnAfterAbstractDQ_Result_TypeDeleteCallback.OnAfterDelete(stage, abstractdq_result_type, frontAbstractDQ_Result_Type)
	}
}

func (abstractdq_temporalaccuracy_type *AbstractDQ_TemporalAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDQ_TemporalAccuracy_TypeCreateCallback != nil {
		stage.OnAfterAbstractDQ_TemporalAccuracy_TypeCreateCallback.OnAfterCreate(stage, abstractdq_temporalaccuracy_type)
	}
}

func (abstractdq_temporalaccuracy_type *AbstractDQ_TemporalAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_TemporalAccuracy_TypeUpdateCallback != nil {
		var frontAbstractDQ_TemporalAccuracy_Type *AbstractDQ_TemporalAccuracy_Type
		if front != nil {
			frontAbstractDQ_TemporalAccuracy_Type, _ = front.(*AbstractDQ_TemporalAccuracy_Type)
		}
		stage.OnAfterAbstractDQ_TemporalAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, abstractdq_temporalaccuracy_type, frontAbstractDQ_TemporalAccuracy_Type)
	}
}

func (abstractdq_temporalaccuracy_type *AbstractDQ_TemporalAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_TemporalAccuracy_TypeDeleteCallback != nil {
		var frontAbstractDQ_TemporalAccuracy_Type *AbstractDQ_TemporalAccuracy_Type
		if front != nil {
			frontAbstractDQ_TemporalAccuracy_Type, _ = front.(*AbstractDQ_TemporalAccuracy_Type)
		}
		stage.OnAfterAbstractDQ_TemporalAccuracy_TypeDeleteCallback.OnAfterDelete(stage, abstractdq_temporalaccuracy_type, frontAbstractDQ_TemporalAccuracy_Type)
	}
}

func (abstractdq_thematicaccuracy_type *AbstractDQ_ThematicAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDQ_ThematicAccuracy_TypeCreateCallback != nil {
		stage.OnAfterAbstractDQ_ThematicAccuracy_TypeCreateCallback.OnAfterCreate(stage, abstractdq_thematicaccuracy_type)
	}
}

func (abstractdq_thematicaccuracy_type *AbstractDQ_ThematicAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_ThematicAccuracy_TypeUpdateCallback != nil {
		var frontAbstractDQ_ThematicAccuracy_Type *AbstractDQ_ThematicAccuracy_Type
		if front != nil {
			frontAbstractDQ_ThematicAccuracy_Type, _ = front.(*AbstractDQ_ThematicAccuracy_Type)
		}
		stage.OnAfterAbstractDQ_ThematicAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, abstractdq_thematicaccuracy_type, frontAbstractDQ_ThematicAccuracy_Type)
	}
}

func (abstractdq_thematicaccuracy_type *AbstractDQ_ThematicAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDQ_ThematicAccuracy_TypeDeleteCallback != nil {
		var frontAbstractDQ_ThematicAccuracy_Type *AbstractDQ_ThematicAccuracy_Type
		if front != nil {
			frontAbstractDQ_ThematicAccuracy_Type, _ = front.(*AbstractDQ_ThematicAccuracy_Type)
		}
		stage.OnAfterAbstractDQ_ThematicAccuracy_TypeDeleteCallback.OnAfterDelete(stage, abstractdq_thematicaccuracy_type, frontAbstractDQ_ThematicAccuracy_Type)
	}
}

func (abstractdatumtype *AbstractDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractDatumTypeCreateCallback != nil {
		stage.OnAfterAbstractDatumTypeCreateCallback.OnAfterCreate(stage, abstractdatumtype)
	}
}

func (abstractdatumtype *AbstractDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDatumTypeUpdateCallback != nil {
		var frontAbstractDatumType *AbstractDatumType
		if front != nil {
			frontAbstractDatumType, _ = front.(*AbstractDatumType)
		}
		stage.OnAfterAbstractDatumTypeUpdateCallback.OnAfterUpdate(stage, abstractdatumtype, frontAbstractDatumType)
	}
}

func (abstractdatumtype *AbstractDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractDatumTypeDeleteCallback != nil {
		var frontAbstractDatumType *AbstractDatumType
		if front != nil {
			frontAbstractDatumType, _ = front.(*AbstractDatumType)
		}
		stage.OnAfterAbstractDatumTypeDeleteCallback.OnAfterDelete(stage, abstractdatumtype, frontAbstractDatumType)
	}
}

func (abstractex_geographicextent_type *AbstractEX_GeographicExtent_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractEX_GeographicExtent_TypeCreateCallback != nil {
		stage.OnAfterAbstractEX_GeographicExtent_TypeCreateCallback.OnAfterCreate(stage, abstractex_geographicextent_type)
	}
}

func (abstractex_geographicextent_type *AbstractEX_GeographicExtent_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractEX_GeographicExtent_TypeUpdateCallback != nil {
		var frontAbstractEX_GeographicExtent_Type *AbstractEX_GeographicExtent_Type
		if front != nil {
			frontAbstractEX_GeographicExtent_Type, _ = front.(*AbstractEX_GeographicExtent_Type)
		}
		stage.OnAfterAbstractEX_GeographicExtent_TypeUpdateCallback.OnAfterUpdate(stage, abstractex_geographicextent_type, frontAbstractEX_GeographicExtent_Type)
	}
}

func (abstractex_geographicextent_type *AbstractEX_GeographicExtent_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractEX_GeographicExtent_TypeDeleteCallback != nil {
		var frontAbstractEX_GeographicExtent_Type *AbstractEX_GeographicExtent_Type
		if front != nil {
			frontAbstractEX_GeographicExtent_Type, _ = front.(*AbstractEX_GeographicExtent_Type)
		}
		stage.OnAfterAbstractEX_GeographicExtent_TypeDeleteCallback.OnAfterDelete(stage, abstractex_geographicextent_type, frontAbstractEX_GeographicExtent_Type)
	}
}

func (abstractfeaturecollectiontype *AbstractFeatureCollectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractFeatureCollectionTypeCreateCallback != nil {
		stage.OnAfterAbstractFeatureCollectionTypeCreateCallback.OnAfterCreate(stage, abstractfeaturecollectiontype)
	}
}

func (abstractfeaturecollectiontype *AbstractFeatureCollectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureCollectionTypeUpdateCallback != nil {
		var frontAbstractFeatureCollectionType *AbstractFeatureCollectionType
		if front != nil {
			frontAbstractFeatureCollectionType, _ = front.(*AbstractFeatureCollectionType)
		}
		stage.OnAfterAbstractFeatureCollectionTypeUpdateCallback.OnAfterUpdate(stage, abstractfeaturecollectiontype, frontAbstractFeatureCollectionType)
	}
}

func (abstractfeaturecollectiontype *AbstractFeatureCollectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureCollectionTypeDeleteCallback != nil {
		var frontAbstractFeatureCollectionType *AbstractFeatureCollectionType
		if front != nil {
			frontAbstractFeatureCollectionType, _ = front.(*AbstractFeatureCollectionType)
		}
		stage.OnAfterAbstractFeatureCollectionTypeDeleteCallback.OnAfterDelete(stage, abstractfeaturecollectiontype, frontAbstractFeatureCollectionType)
	}
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractFeatureMemberTypeCreateCallback != nil {
		stage.OnAfterAbstractFeatureMemberTypeCreateCallback.OnAfterCreate(stage, abstractfeaturemembertype)
	}
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureMemberTypeUpdateCallback != nil {
		var frontAbstractFeatureMemberType *AbstractFeatureMemberType
		if front != nil {
			frontAbstractFeatureMemberType, _ = front.(*AbstractFeatureMemberType)
		}
		stage.OnAfterAbstractFeatureMemberTypeUpdateCallback.OnAfterUpdate(stage, abstractfeaturemembertype, frontAbstractFeatureMemberType)
	}
}

func (abstractfeaturemembertype *AbstractFeatureMemberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureMemberTypeDeleteCallback != nil {
		var frontAbstractFeatureMemberType *AbstractFeatureMemberType
		if front != nil {
			frontAbstractFeatureMemberType, _ = front.(*AbstractFeatureMemberType)
		}
		stage.OnAfterAbstractFeatureMemberTypeDeleteCallback.OnAfterDelete(stage, abstractfeaturemembertype, frontAbstractFeatureMemberType)
	}
}

func (abstractfeaturetype *AbstractFeatureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractFeatureTypeCreateCallback != nil {
		stage.OnAfterAbstractFeatureTypeCreateCallback.OnAfterCreate(stage, abstractfeaturetype)
	}
}

func (abstractfeaturetype *AbstractFeatureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureTypeUpdateCallback != nil {
		var frontAbstractFeatureType *AbstractFeatureType
		if front != nil {
			frontAbstractFeatureType, _ = front.(*AbstractFeatureType)
		}
		stage.OnAfterAbstractFeatureTypeUpdateCallback.OnAfterUpdate(stage, abstractfeaturetype, frontAbstractFeatureType)
	}
}

func (abstractfeaturetype *AbstractFeatureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractFeatureTypeDeleteCallback != nil {
		var frontAbstractFeatureType *AbstractFeatureType
		if front != nil {
			frontAbstractFeatureType, _ = front.(*AbstractFeatureType)
		}
		stage.OnAfterAbstractFeatureTypeDeleteCallback.OnAfterDelete(stage, abstractfeaturetype, frontAbstractFeatureType)
	}
}

func (abstractgmltype *AbstractGMLType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGMLTypeCreateCallback != nil {
		stage.OnAfterAbstractGMLTypeCreateCallback.OnAfterCreate(stage, abstractgmltype)
	}
}

func (abstractgmltype *AbstractGMLType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGMLTypeUpdateCallback != nil {
		var frontAbstractGMLType *AbstractGMLType
		if front != nil {
			frontAbstractGMLType, _ = front.(*AbstractGMLType)
		}
		stage.OnAfterAbstractGMLTypeUpdateCallback.OnAfterUpdate(stage, abstractgmltype, frontAbstractGMLType)
	}
}

func (abstractgmltype *AbstractGMLType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGMLTypeDeleteCallback != nil {
		var frontAbstractGMLType *AbstractGMLType
		if front != nil {
			frontAbstractGMLType, _ = front.(*AbstractGMLType)
		}
		stage.OnAfterAbstractGMLTypeDeleteCallback.OnAfterDelete(stage, abstractgmltype, frontAbstractGMLType)
	}
}

func (abstractgeneralconversiontype *AbstractGeneralConversionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeneralConversionTypeCreateCallback != nil {
		stage.OnAfterAbstractGeneralConversionTypeCreateCallback.OnAfterCreate(stage, abstractgeneralconversiontype)
	}
}

func (abstractgeneralconversiontype *AbstractGeneralConversionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralConversionTypeUpdateCallback != nil {
		var frontAbstractGeneralConversionType *AbstractGeneralConversionType
		if front != nil {
			frontAbstractGeneralConversionType, _ = front.(*AbstractGeneralConversionType)
		}
		stage.OnAfterAbstractGeneralConversionTypeUpdateCallback.OnAfterUpdate(stage, abstractgeneralconversiontype, frontAbstractGeneralConversionType)
	}
}

func (abstractgeneralconversiontype *AbstractGeneralConversionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralConversionTypeDeleteCallback != nil {
		var frontAbstractGeneralConversionType *AbstractGeneralConversionType
		if front != nil {
			frontAbstractGeneralConversionType, _ = front.(*AbstractGeneralConversionType)
		}
		stage.OnAfterAbstractGeneralConversionTypeDeleteCallback.OnAfterDelete(stage, abstractgeneralconversiontype, frontAbstractGeneralConversionType)
	}
}

func (abstractgeneralderivedcrstype *AbstractGeneralDerivedCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeneralDerivedCRSTypeCreateCallback != nil {
		stage.OnAfterAbstractGeneralDerivedCRSTypeCreateCallback.OnAfterCreate(stage, abstractgeneralderivedcrstype)
	}
}

func (abstractgeneralderivedcrstype *AbstractGeneralDerivedCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralDerivedCRSTypeUpdateCallback != nil {
		var frontAbstractGeneralDerivedCRSType *AbstractGeneralDerivedCRSType
		if front != nil {
			frontAbstractGeneralDerivedCRSType, _ = front.(*AbstractGeneralDerivedCRSType)
		}
		stage.OnAfterAbstractGeneralDerivedCRSTypeUpdateCallback.OnAfterUpdate(stage, abstractgeneralderivedcrstype, frontAbstractGeneralDerivedCRSType)
	}
}

func (abstractgeneralderivedcrstype *AbstractGeneralDerivedCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralDerivedCRSTypeDeleteCallback != nil {
		var frontAbstractGeneralDerivedCRSType *AbstractGeneralDerivedCRSType
		if front != nil {
			frontAbstractGeneralDerivedCRSType, _ = front.(*AbstractGeneralDerivedCRSType)
		}
		stage.OnAfterAbstractGeneralDerivedCRSTypeDeleteCallback.OnAfterDelete(stage, abstractgeneralderivedcrstype, frontAbstractGeneralDerivedCRSType)
	}
}

func (abstractgeneraloperationparameterpropertytype *AbstractGeneralOperationParameterPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeneralOperationParameterPropertyTypeCreateCallback != nil {
		stage.OnAfterAbstractGeneralOperationParameterPropertyTypeCreateCallback.OnAfterCreate(stage, abstractgeneraloperationparameterpropertytype)
	}
}

func (abstractgeneraloperationparameterpropertytype *AbstractGeneralOperationParameterPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralOperationParameterPropertyTypeUpdateCallback != nil {
		var frontAbstractGeneralOperationParameterPropertyType *AbstractGeneralOperationParameterPropertyType
		if front != nil {
			frontAbstractGeneralOperationParameterPropertyType, _ = front.(*AbstractGeneralOperationParameterPropertyType)
		}
		stage.OnAfterAbstractGeneralOperationParameterPropertyTypeUpdateCallback.OnAfterUpdate(stage, abstractgeneraloperationparameterpropertytype, frontAbstractGeneralOperationParameterPropertyType)
	}
}

func (abstractgeneraloperationparameterpropertytype *AbstractGeneralOperationParameterPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralOperationParameterPropertyTypeDeleteCallback != nil {
		var frontAbstractGeneralOperationParameterPropertyType *AbstractGeneralOperationParameterPropertyType
		if front != nil {
			frontAbstractGeneralOperationParameterPropertyType, _ = front.(*AbstractGeneralOperationParameterPropertyType)
		}
		stage.OnAfterAbstractGeneralOperationParameterPropertyTypeDeleteCallback.OnAfterDelete(stage, abstractgeneraloperationparameterpropertytype, frontAbstractGeneralOperationParameterPropertyType)
	}
}

func (abstractgeneraloperationparametertype *AbstractGeneralOperationParameterType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeneralOperationParameterTypeCreateCallback != nil {
		stage.OnAfterAbstractGeneralOperationParameterTypeCreateCallback.OnAfterCreate(stage, abstractgeneraloperationparametertype)
	}
}

func (abstractgeneraloperationparametertype *AbstractGeneralOperationParameterType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralOperationParameterTypeUpdateCallback != nil {
		var frontAbstractGeneralOperationParameterType *AbstractGeneralOperationParameterType
		if front != nil {
			frontAbstractGeneralOperationParameterType, _ = front.(*AbstractGeneralOperationParameterType)
		}
		stage.OnAfterAbstractGeneralOperationParameterTypeUpdateCallback.OnAfterUpdate(stage, abstractgeneraloperationparametertype, frontAbstractGeneralOperationParameterType)
	}
}

func (abstractgeneraloperationparametertype *AbstractGeneralOperationParameterType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralOperationParameterTypeDeleteCallback != nil {
		var frontAbstractGeneralOperationParameterType *AbstractGeneralOperationParameterType
		if front != nil {
			frontAbstractGeneralOperationParameterType, _ = front.(*AbstractGeneralOperationParameterType)
		}
		stage.OnAfterAbstractGeneralOperationParameterTypeDeleteCallback.OnAfterDelete(stage, abstractgeneraloperationparametertype, frontAbstractGeneralOperationParameterType)
	}
}

func (abstractgeneralparametervaluepropertytype *AbstractGeneralParameterValuePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeneralParameterValuePropertyTypeCreateCallback != nil {
		stage.OnAfterAbstractGeneralParameterValuePropertyTypeCreateCallback.OnAfterCreate(stage, abstractgeneralparametervaluepropertytype)
	}
}

func (abstractgeneralparametervaluepropertytype *AbstractGeneralParameterValuePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralParameterValuePropertyTypeUpdateCallback != nil {
		var frontAbstractGeneralParameterValuePropertyType *AbstractGeneralParameterValuePropertyType
		if front != nil {
			frontAbstractGeneralParameterValuePropertyType, _ = front.(*AbstractGeneralParameterValuePropertyType)
		}
		stage.OnAfterAbstractGeneralParameterValuePropertyTypeUpdateCallback.OnAfterUpdate(stage, abstractgeneralparametervaluepropertytype, frontAbstractGeneralParameterValuePropertyType)
	}
}

func (abstractgeneralparametervaluepropertytype *AbstractGeneralParameterValuePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralParameterValuePropertyTypeDeleteCallback != nil {
		var frontAbstractGeneralParameterValuePropertyType *AbstractGeneralParameterValuePropertyType
		if front != nil {
			frontAbstractGeneralParameterValuePropertyType, _ = front.(*AbstractGeneralParameterValuePropertyType)
		}
		stage.OnAfterAbstractGeneralParameterValuePropertyTypeDeleteCallback.OnAfterDelete(stage, abstractgeneralparametervaluepropertytype, frontAbstractGeneralParameterValuePropertyType)
	}
}

func (abstractgeneralparametervaluetype *AbstractGeneralParameterValueType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeneralParameterValueTypeCreateCallback != nil {
		stage.OnAfterAbstractGeneralParameterValueTypeCreateCallback.OnAfterCreate(stage, abstractgeneralparametervaluetype)
	}
}

func (abstractgeneralparametervaluetype *AbstractGeneralParameterValueType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralParameterValueTypeUpdateCallback != nil {
		var frontAbstractGeneralParameterValueType *AbstractGeneralParameterValueType
		if front != nil {
			frontAbstractGeneralParameterValueType, _ = front.(*AbstractGeneralParameterValueType)
		}
		stage.OnAfterAbstractGeneralParameterValueTypeUpdateCallback.OnAfterUpdate(stage, abstractgeneralparametervaluetype, frontAbstractGeneralParameterValueType)
	}
}

func (abstractgeneralparametervaluetype *AbstractGeneralParameterValueType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralParameterValueTypeDeleteCallback != nil {
		var frontAbstractGeneralParameterValueType *AbstractGeneralParameterValueType
		if front != nil {
			frontAbstractGeneralParameterValueType, _ = front.(*AbstractGeneralParameterValueType)
		}
		stage.OnAfterAbstractGeneralParameterValueTypeDeleteCallback.OnAfterDelete(stage, abstractgeneralparametervaluetype, frontAbstractGeneralParameterValueType)
	}
}

func (abstractgeneraltransformationtype *AbstractGeneralTransformationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeneralTransformationTypeCreateCallback != nil {
		stage.OnAfterAbstractGeneralTransformationTypeCreateCallback.OnAfterCreate(stage, abstractgeneraltransformationtype)
	}
}

func (abstractgeneraltransformationtype *AbstractGeneralTransformationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralTransformationTypeUpdateCallback != nil {
		var frontAbstractGeneralTransformationType *AbstractGeneralTransformationType
		if front != nil {
			frontAbstractGeneralTransformationType, _ = front.(*AbstractGeneralTransformationType)
		}
		stage.OnAfterAbstractGeneralTransformationTypeUpdateCallback.OnAfterUpdate(stage, abstractgeneraltransformationtype, frontAbstractGeneralTransformationType)
	}
}

func (abstractgeneraltransformationtype *AbstractGeneralTransformationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeneralTransformationTypeDeleteCallback != nil {
		var frontAbstractGeneralTransformationType *AbstractGeneralTransformationType
		if front != nil {
			frontAbstractGeneralTransformationType, _ = front.(*AbstractGeneralTransformationType)
		}
		stage.OnAfterAbstractGeneralTransformationTypeDeleteCallback.OnAfterDelete(stage, abstractgeneraltransformationtype, frontAbstractGeneralTransformationType)
	}
}

func (abstractgeometricaggregatetype *AbstractGeometricAggregateType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeometricAggregateTypeCreateCallback != nil {
		stage.OnAfterAbstractGeometricAggregateTypeCreateCallback.OnAfterCreate(stage, abstractgeometricaggregatetype)
	}
}

func (abstractgeometricaggregatetype *AbstractGeometricAggregateType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometricAggregateTypeUpdateCallback != nil {
		var frontAbstractGeometricAggregateType *AbstractGeometricAggregateType
		if front != nil {
			frontAbstractGeometricAggregateType, _ = front.(*AbstractGeometricAggregateType)
		}
		stage.OnAfterAbstractGeometricAggregateTypeUpdateCallback.OnAfterUpdate(stage, abstractgeometricaggregatetype, frontAbstractGeometricAggregateType)
	}
}

func (abstractgeometricaggregatetype *AbstractGeometricAggregateType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometricAggregateTypeDeleteCallback != nil {
		var frontAbstractGeometricAggregateType *AbstractGeometricAggregateType
		if front != nil {
			frontAbstractGeometricAggregateType, _ = front.(*AbstractGeometricAggregateType)
		}
		stage.OnAfterAbstractGeometricAggregateTypeDeleteCallback.OnAfterDelete(stage, abstractgeometricaggregatetype, frontAbstractGeometricAggregateType)
	}
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeometricPrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractGeometricPrimitiveTypeCreateCallback.OnAfterCreate(stage, abstractgeometricprimitivetype)
	}
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometricPrimitiveTypeUpdateCallback != nil {
		var frontAbstractGeometricPrimitiveType *AbstractGeometricPrimitiveType
		if front != nil {
			frontAbstractGeometricPrimitiveType, _ = front.(*AbstractGeometricPrimitiveType)
		}
		stage.OnAfterAbstractGeometricPrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstractgeometricprimitivetype, frontAbstractGeometricPrimitiveType)
	}
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometricPrimitiveTypeDeleteCallback != nil {
		var frontAbstractGeometricPrimitiveType *AbstractGeometricPrimitiveType
		if front != nil {
			frontAbstractGeometricPrimitiveType, _ = front.(*AbstractGeometricPrimitiveType)
		}
		stage.OnAfterAbstractGeometricPrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstractgeometricprimitivetype, frontAbstractGeometricPrimitiveType)
	}
}

func (abstractgeometrytype *AbstractGeometryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGeometryTypeCreateCallback != nil {
		stage.OnAfterAbstractGeometryTypeCreateCallback.OnAfterCreate(stage, abstractgeometrytype)
	}
}

func (abstractgeometrytype *AbstractGeometryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometryTypeUpdateCallback != nil {
		var frontAbstractGeometryType *AbstractGeometryType
		if front != nil {
			frontAbstractGeometryType, _ = front.(*AbstractGeometryType)
		}
		stage.OnAfterAbstractGeometryTypeUpdateCallback.OnAfterUpdate(stage, abstractgeometrytype, frontAbstractGeometryType)
	}
}

func (abstractgeometrytype *AbstractGeometryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGeometryTypeDeleteCallback != nil {
		var frontAbstractGeometryType *AbstractGeometryType
		if front != nil {
			frontAbstractGeometryType, _ = front.(*AbstractGeometryType)
		}
		stage.OnAfterAbstractGeometryTypeDeleteCallback.OnAfterDelete(stage, abstractgeometrytype, frontAbstractGeometryType)
	}
}

func (abstractgriddedsurfacetype *AbstractGriddedSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractGriddedSurfaceTypeCreateCallback != nil {
		stage.OnAfterAbstractGriddedSurfaceTypeCreateCallback.OnAfterCreate(stage, abstractgriddedsurfacetype)
	}
}

func (abstractgriddedsurfacetype *AbstractGriddedSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGriddedSurfaceTypeUpdateCallback != nil {
		var frontAbstractGriddedSurfaceType *AbstractGriddedSurfaceType
		if front != nil {
			frontAbstractGriddedSurfaceType, _ = front.(*AbstractGriddedSurfaceType)
		}
		stage.OnAfterAbstractGriddedSurfaceTypeUpdateCallback.OnAfterUpdate(stage, abstractgriddedsurfacetype, frontAbstractGriddedSurfaceType)
	}
}

func (abstractgriddedsurfacetype *AbstractGriddedSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractGriddedSurfaceTypeDeleteCallback != nil {
		var frontAbstractGriddedSurfaceType *AbstractGriddedSurfaceType
		if front != nil {
			frontAbstractGriddedSurfaceType, _ = front.(*AbstractGriddedSurfaceType)
		}
		stage.OnAfterAbstractGriddedSurfaceTypeDeleteCallback.OnAfterDelete(stage, abstractgriddedsurfacetype, frontAbstractGriddedSurfaceType)
	}
}

func (abstractmd_identification_type *AbstractMD_Identification_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMD_Identification_TypeCreateCallback != nil {
		stage.OnAfterAbstractMD_Identification_TypeCreateCallback.OnAfterCreate(stage, abstractmd_identification_type)
	}
}

func (abstractmd_identification_type *AbstractMD_Identification_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMD_Identification_TypeUpdateCallback != nil {
		var frontAbstractMD_Identification_Type *AbstractMD_Identification_Type
		if front != nil {
			frontAbstractMD_Identification_Type, _ = front.(*AbstractMD_Identification_Type)
		}
		stage.OnAfterAbstractMD_Identification_TypeUpdateCallback.OnAfterUpdate(stage, abstractmd_identification_type, frontAbstractMD_Identification_Type)
	}
}

func (abstractmd_identification_type *AbstractMD_Identification_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMD_Identification_TypeDeleteCallback != nil {
		var frontAbstractMD_Identification_Type *AbstractMD_Identification_Type
		if front != nil {
			frontAbstractMD_Identification_Type, _ = front.(*AbstractMD_Identification_Type)
		}
		stage.OnAfterAbstractMD_Identification_TypeDeleteCallback.OnAfterDelete(stage, abstractmd_identification_type, frontAbstractMD_Identification_Type)
	}
}

func (abstractmembertype *AbstractMemberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMemberTypeCreateCallback != nil {
		stage.OnAfterAbstractMemberTypeCreateCallback.OnAfterCreate(stage, abstractmembertype)
	}
}

func (abstractmembertype *AbstractMemberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMemberTypeUpdateCallback != nil {
		var frontAbstractMemberType *AbstractMemberType
		if front != nil {
			frontAbstractMemberType, _ = front.(*AbstractMemberType)
		}
		stage.OnAfterAbstractMemberTypeUpdateCallback.OnAfterUpdate(stage, abstractmembertype, frontAbstractMemberType)
	}
}

func (abstractmembertype *AbstractMemberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMemberTypeDeleteCallback != nil {
		var frontAbstractMemberType *AbstractMemberType
		if front != nil {
			frontAbstractMemberType, _ = front.(*AbstractMemberType)
		}
		stage.OnAfterAbstractMemberTypeDeleteCallback.OnAfterDelete(stage, abstractmembertype, frontAbstractMemberType)
	}
}

func (abstractmetadatatype *AbstractMetaDataType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMetaDataTypeCreateCallback != nil {
		stage.OnAfterAbstractMetaDataTypeCreateCallback.OnAfterCreate(stage, abstractmetadatatype)
	}
}

func (abstractmetadatatype *AbstractMetaDataType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMetaDataTypeUpdateCallback != nil {
		var frontAbstractMetaDataType *AbstractMetaDataType
		if front != nil {
			frontAbstractMetaDataType, _ = front.(*AbstractMetaDataType)
		}
		stage.OnAfterAbstractMetaDataTypeUpdateCallback.OnAfterUpdate(stage, abstractmetadatatype, frontAbstractMetaDataType)
	}
}

func (abstractmetadatatype *AbstractMetaDataType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMetaDataTypeDeleteCallback != nil {
		var frontAbstractMetaDataType *AbstractMetaDataType
		if front != nil {
			frontAbstractMetaDataType, _ = front.(*AbstractMetaDataType)
		}
		stage.OnAfterAbstractMetaDataTypeDeleteCallback.OnAfterDelete(stage, abstractmetadatatype, frontAbstractMetaDataType)
	}
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractMetadataPropertyTypeCreateCallback != nil {
		stage.OnAfterAbstractMetadataPropertyTypeCreateCallback.OnAfterCreate(stage, abstractmetadatapropertytype)
	}
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMetadataPropertyTypeUpdateCallback != nil {
		var frontAbstractMetadataPropertyType *AbstractMetadataPropertyType
		if front != nil {
			frontAbstractMetadataPropertyType, _ = front.(*AbstractMetadataPropertyType)
		}
		stage.OnAfterAbstractMetadataPropertyTypeUpdateCallback.OnAfterUpdate(stage, abstractmetadatapropertytype, frontAbstractMetadataPropertyType)
	}
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractMetadataPropertyTypeDeleteCallback != nil {
		var frontAbstractMetadataPropertyType *AbstractMetadataPropertyType
		if front != nil {
			frontAbstractMetadataPropertyType, _ = front.(*AbstractMetadataPropertyType)
		}
		stage.OnAfterAbstractMetadataPropertyTypeDeleteCallback.OnAfterDelete(stage, abstractmetadatapropertytype, frontAbstractMetadataPropertyType)
	}
}

func (abstractparametriccurvesurfacetype *AbstractParametricCurveSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractParametricCurveSurfaceTypeCreateCallback != nil {
		stage.OnAfterAbstractParametricCurveSurfaceTypeCreateCallback.OnAfterCreate(stage, abstractparametriccurvesurfacetype)
	}
}

func (abstractparametriccurvesurfacetype *AbstractParametricCurveSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractParametricCurveSurfaceTypeUpdateCallback != nil {
		var frontAbstractParametricCurveSurfaceType *AbstractParametricCurveSurfaceType
		if front != nil {
			frontAbstractParametricCurveSurfaceType, _ = front.(*AbstractParametricCurveSurfaceType)
		}
		stage.OnAfterAbstractParametricCurveSurfaceTypeUpdateCallback.OnAfterUpdate(stage, abstractparametriccurvesurfacetype, frontAbstractParametricCurveSurfaceType)
	}
}

func (abstractparametriccurvesurfacetype *AbstractParametricCurveSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractParametricCurveSurfaceTypeDeleteCallback != nil {
		var frontAbstractParametricCurveSurfaceType *AbstractParametricCurveSurfaceType
		if front != nil {
			frontAbstractParametricCurveSurfaceType, _ = front.(*AbstractParametricCurveSurfaceType)
		}
		stage.OnAfterAbstractParametricCurveSurfaceTypeDeleteCallback.OnAfterDelete(stage, abstractparametriccurvesurfacetype, frontAbstractParametricCurveSurfaceType)
	}
}

func (abstractrs_referencesystem_type *AbstractRS_ReferenceSystem_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractRS_ReferenceSystem_TypeCreateCallback != nil {
		stage.OnAfterAbstractRS_ReferenceSystem_TypeCreateCallback.OnAfterCreate(stage, abstractrs_referencesystem_type)
	}
}

func (abstractrs_referencesystem_type *AbstractRS_ReferenceSystem_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRS_ReferenceSystem_TypeUpdateCallback != nil {
		var frontAbstractRS_ReferenceSystem_Type *AbstractRS_ReferenceSystem_Type
		if front != nil {
			frontAbstractRS_ReferenceSystem_Type, _ = front.(*AbstractRS_ReferenceSystem_Type)
		}
		stage.OnAfterAbstractRS_ReferenceSystem_TypeUpdateCallback.OnAfterUpdate(stage, abstractrs_referencesystem_type, frontAbstractRS_ReferenceSystem_Type)
	}
}

func (abstractrs_referencesystem_type *AbstractRS_ReferenceSystem_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRS_ReferenceSystem_TypeDeleteCallback != nil {
		var frontAbstractRS_ReferenceSystem_Type *AbstractRS_ReferenceSystem_Type
		if front != nil {
			frontAbstractRS_ReferenceSystem_Type, _ = front.(*AbstractRS_ReferenceSystem_Type)
		}
		stage.OnAfterAbstractRS_ReferenceSystem_TypeDeleteCallback.OnAfterDelete(stage, abstractrs_referencesystem_type, frontAbstractRS_ReferenceSystem_Type)
	}
}

func (abstractringpropertytype *AbstractRingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractRingPropertyTypeCreateCallback != nil {
		stage.OnAfterAbstractRingPropertyTypeCreateCallback.OnAfterCreate(stage, abstractringpropertytype)
	}
}

func (abstractringpropertytype *AbstractRingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingPropertyTypeUpdateCallback != nil {
		var frontAbstractRingPropertyType *AbstractRingPropertyType
		if front != nil {
			frontAbstractRingPropertyType, _ = front.(*AbstractRingPropertyType)
		}
		stage.OnAfterAbstractRingPropertyTypeUpdateCallback.OnAfterUpdate(stage, abstractringpropertytype, frontAbstractRingPropertyType)
	}
}

func (abstractringpropertytype *AbstractRingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingPropertyTypeDeleteCallback != nil {
		var frontAbstractRingPropertyType *AbstractRingPropertyType
		if front != nil {
			frontAbstractRingPropertyType, _ = front.(*AbstractRingPropertyType)
		}
		stage.OnAfterAbstractRingPropertyTypeDeleteCallback.OnAfterDelete(stage, abstractringpropertytype, frontAbstractRingPropertyType)
	}
}

func (abstractringtype *AbstractRingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractRingTypeCreateCallback != nil {
		stage.OnAfterAbstractRingTypeCreateCallback.OnAfterCreate(stage, abstractringtype)
	}
}

func (abstractringtype *AbstractRingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingTypeUpdateCallback != nil {
		var frontAbstractRingType *AbstractRingType
		if front != nil {
			frontAbstractRingType, _ = front.(*AbstractRingType)
		}
		stage.OnAfterAbstractRingTypeUpdateCallback.OnAfterUpdate(stage, abstractringtype, frontAbstractRingType)
	}
}

func (abstractringtype *AbstractRingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractRingTypeDeleteCallback != nil {
		var frontAbstractRingType *AbstractRingType
		if front != nil {
			frontAbstractRingType, _ = front.(*AbstractRingType)
		}
		stage.OnAfterAbstractRingTypeDeleteCallback.OnAfterDelete(stage, abstractringtype, frontAbstractRingType)
	}
}

func (abstractsolidtype *AbstractSolidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSolidTypeCreateCallback != nil {
		stage.OnAfterAbstractSolidTypeCreateCallback.OnAfterCreate(stage, abstractsolidtype)
	}
}

func (abstractsolidtype *AbstractSolidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSolidTypeUpdateCallback != nil {
		var frontAbstractSolidType *AbstractSolidType
		if front != nil {
			frontAbstractSolidType, _ = front.(*AbstractSolidType)
		}
		stage.OnAfterAbstractSolidTypeUpdateCallback.OnAfterUpdate(stage, abstractsolidtype, frontAbstractSolidType)
	}
}

func (abstractsolidtype *AbstractSolidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSolidTypeDeleteCallback != nil {
		var frontAbstractSolidType *AbstractSolidType
		if front != nil {
			frontAbstractSolidType, _ = front.(*AbstractSolidType)
		}
		stage.OnAfterAbstractSolidTypeDeleteCallback.OnAfterDelete(stage, abstractsolidtype, frontAbstractSolidType)
	}
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSurfacePatchTypeCreateCallback != nil {
		stage.OnAfterAbstractSurfacePatchTypeCreateCallback.OnAfterCreate(stage, abstractsurfacepatchtype)
	}
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfacePatchTypeUpdateCallback != nil {
		var frontAbstractSurfacePatchType *AbstractSurfacePatchType
		if front != nil {
			frontAbstractSurfacePatchType, _ = front.(*AbstractSurfacePatchType)
		}
		stage.OnAfterAbstractSurfacePatchTypeUpdateCallback.OnAfterUpdate(stage, abstractsurfacepatchtype, frontAbstractSurfacePatchType)
	}
}

func (abstractsurfacepatchtype *AbstractSurfacePatchType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfacePatchTypeDeleteCallback != nil {
		var frontAbstractSurfacePatchType *AbstractSurfacePatchType
		if front != nil {
			frontAbstractSurfacePatchType, _ = front.(*AbstractSurfacePatchType)
		}
		stage.OnAfterAbstractSurfacePatchTypeDeleteCallback.OnAfterDelete(stage, abstractsurfacepatchtype, frontAbstractSurfacePatchType)
	}
}

func (abstractsurfacetype *AbstractSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractSurfaceTypeCreateCallback != nil {
		stage.OnAfterAbstractSurfaceTypeCreateCallback.OnAfterCreate(stage, abstractsurfacetype)
	}
}

func (abstractsurfacetype *AbstractSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfaceTypeUpdateCallback != nil {
		var frontAbstractSurfaceType *AbstractSurfaceType
		if front != nil {
			frontAbstractSurfaceType, _ = front.(*AbstractSurfaceType)
		}
		stage.OnAfterAbstractSurfaceTypeUpdateCallback.OnAfterUpdate(stage, abstractsurfacetype, frontAbstractSurfaceType)
	}
}

func (abstractsurfacetype *AbstractSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractSurfaceTypeDeleteCallback != nil {
		var frontAbstractSurfaceType *AbstractSurfaceType
		if front != nil {
			frontAbstractSurfaceType, _ = front.(*AbstractSurfaceType)
		}
		stage.OnAfterAbstractSurfaceTypeDeleteCallback.OnAfterDelete(stage, abstractsurfacetype, frontAbstractSurfaceType)
	}
}

func (abstracttimecomplextype *AbstractTimeComplexType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeComplexTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeComplexTypeCreateCallback.OnAfterCreate(stage, abstracttimecomplextype)
	}
}

func (abstracttimecomplextype *AbstractTimeComplexType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeComplexTypeUpdateCallback != nil {
		var frontAbstractTimeComplexType *AbstractTimeComplexType
		if front != nil {
			frontAbstractTimeComplexType, _ = front.(*AbstractTimeComplexType)
		}
		stage.OnAfterAbstractTimeComplexTypeUpdateCallback.OnAfterUpdate(stage, abstracttimecomplextype, frontAbstractTimeComplexType)
	}
}

func (abstracttimecomplextype *AbstractTimeComplexType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeComplexTypeDeleteCallback != nil {
		var frontAbstractTimeComplexType *AbstractTimeComplexType
		if front != nil {
			frontAbstractTimeComplexType, _ = front.(*AbstractTimeComplexType)
		}
		stage.OnAfterAbstractTimeComplexTypeDeleteCallback.OnAfterDelete(stage, abstracttimecomplextype, frontAbstractTimeComplexType)
	}
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeGeometricPrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeGeometricPrimitiveTypeCreateCallback.OnAfterCreate(stage, abstracttimegeometricprimitivetype)
	}
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeGeometricPrimitiveTypeUpdateCallback != nil {
		var frontAbstractTimeGeometricPrimitiveType *AbstractTimeGeometricPrimitiveType
		if front != nil {
			frontAbstractTimeGeometricPrimitiveType, _ = front.(*AbstractTimeGeometricPrimitiveType)
		}
		stage.OnAfterAbstractTimeGeometricPrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstracttimegeometricprimitivetype, frontAbstractTimeGeometricPrimitiveType)
	}
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeGeometricPrimitiveTypeDeleteCallback != nil {
		var frontAbstractTimeGeometricPrimitiveType *AbstractTimeGeometricPrimitiveType
		if front != nil {
			frontAbstractTimeGeometricPrimitiveType, _ = front.(*AbstractTimeGeometricPrimitiveType)
		}
		stage.OnAfterAbstractTimeGeometricPrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstracttimegeometricprimitivetype, frontAbstractTimeGeometricPrimitiveType)
	}
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeObjectTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeObjectTypeCreateCallback.OnAfterCreate(stage, abstracttimeobjecttype)
	}
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeObjectTypeUpdateCallback != nil {
		var frontAbstractTimeObjectType *AbstractTimeObjectType
		if front != nil {
			frontAbstractTimeObjectType, _ = front.(*AbstractTimeObjectType)
		}
		stage.OnAfterAbstractTimeObjectTypeUpdateCallback.OnAfterUpdate(stage, abstracttimeobjecttype, frontAbstractTimeObjectType)
	}
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeObjectTypeDeleteCallback != nil {
		var frontAbstractTimeObjectType *AbstractTimeObjectType
		if front != nil {
			frontAbstractTimeObjectType, _ = front.(*AbstractTimeObjectType)
		}
		stage.OnAfterAbstractTimeObjectTypeDeleteCallback.OnAfterDelete(stage, abstracttimeobjecttype, frontAbstractTimeObjectType)
	}
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimePrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractTimePrimitiveTypeCreateCallback.OnAfterCreate(stage, abstracttimeprimitivetype)
	}
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimePrimitiveTypeUpdateCallback != nil {
		var frontAbstractTimePrimitiveType *AbstractTimePrimitiveType
		if front != nil {
			frontAbstractTimePrimitiveType, _ = front.(*AbstractTimePrimitiveType)
		}
		stage.OnAfterAbstractTimePrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstracttimeprimitivetype, frontAbstractTimePrimitiveType)
	}
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimePrimitiveTypeDeleteCallback != nil {
		var frontAbstractTimePrimitiveType *AbstractTimePrimitiveType
		if front != nil {
			frontAbstractTimePrimitiveType, _ = front.(*AbstractTimePrimitiveType)
		}
		stage.OnAfterAbstractTimePrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstracttimeprimitivetype, frontAbstractTimePrimitiveType)
	}
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeSliceTypeCreateCallback.OnAfterCreate(stage, abstracttimeslicetype)
	}
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeSliceTypeUpdateCallback != nil {
		var frontAbstractTimeSliceType *AbstractTimeSliceType
		if front != nil {
			frontAbstractTimeSliceType, _ = front.(*AbstractTimeSliceType)
		}
		stage.OnAfterAbstractTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, abstracttimeslicetype, frontAbstractTimeSliceType)
	}
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeSliceTypeDeleteCallback != nil {
		var frontAbstractTimeSliceType *AbstractTimeSliceType
		if front != nil {
			frontAbstractTimeSliceType, _ = front.(*AbstractTimeSliceType)
		}
		stage.OnAfterAbstractTimeSliceTypeDeleteCallback.OnAfterDelete(stage, abstracttimeslicetype, frontAbstractTimeSliceType)
	}
}

func (abstracttimetopologyprimitivetype *AbstractTimeTopologyPrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTimeTopologyPrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractTimeTopologyPrimitiveTypeCreateCallback.OnAfterCreate(stage, abstracttimetopologyprimitivetype)
	}
}

func (abstracttimetopologyprimitivetype *AbstractTimeTopologyPrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeTopologyPrimitiveTypeUpdateCallback != nil {
		var frontAbstractTimeTopologyPrimitiveType *AbstractTimeTopologyPrimitiveType
		if front != nil {
			frontAbstractTimeTopologyPrimitiveType, _ = front.(*AbstractTimeTopologyPrimitiveType)
		}
		stage.OnAfterAbstractTimeTopologyPrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstracttimetopologyprimitivetype, frontAbstractTimeTopologyPrimitiveType)
	}
}

func (abstracttimetopologyprimitivetype *AbstractTimeTopologyPrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTimeTopologyPrimitiveTypeDeleteCallback != nil {
		var frontAbstractTimeTopologyPrimitiveType *AbstractTimeTopologyPrimitiveType
		if front != nil {
			frontAbstractTimeTopologyPrimitiveType, _ = front.(*AbstractTimeTopologyPrimitiveType)
		}
		stage.OnAfterAbstractTimeTopologyPrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstracttimetopologyprimitivetype, frontAbstractTimeTopologyPrimitiveType)
	}
}

func (abstracttopoprimitivetype *AbstractTopoPrimitiveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTopoPrimitiveTypeCreateCallback != nil {
		stage.OnAfterAbstractTopoPrimitiveTypeCreateCallback.OnAfterCreate(stage, abstracttopoprimitivetype)
	}
}

func (abstracttopoprimitivetype *AbstractTopoPrimitiveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTopoPrimitiveTypeUpdateCallback != nil {
		var frontAbstractTopoPrimitiveType *AbstractTopoPrimitiveType
		if front != nil {
			frontAbstractTopoPrimitiveType, _ = front.(*AbstractTopoPrimitiveType)
		}
		stage.OnAfterAbstractTopoPrimitiveTypeUpdateCallback.OnAfterUpdate(stage, abstracttopoprimitivetype, frontAbstractTopoPrimitiveType)
	}
}

func (abstracttopoprimitivetype *AbstractTopoPrimitiveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTopoPrimitiveTypeDeleteCallback != nil {
		var frontAbstractTopoPrimitiveType *AbstractTopoPrimitiveType
		if front != nil {
			frontAbstractTopoPrimitiveType, _ = front.(*AbstractTopoPrimitiveType)
		}
		stage.OnAfterAbstractTopoPrimitiveTypeDeleteCallback.OnAfterDelete(stage, abstracttopoprimitivetype, frontAbstractTopoPrimitiveType)
	}
}

func (abstracttopologytype *AbstractTopologyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractTopologyTypeCreateCallback != nil {
		stage.OnAfterAbstractTopologyTypeCreateCallback.OnAfterCreate(stage, abstracttopologytype)
	}
}

func (abstracttopologytype *AbstractTopologyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTopologyTypeUpdateCallback != nil {
		var frontAbstractTopologyType *AbstractTopologyType
		if front != nil {
			frontAbstractTopologyType, _ = front.(*AbstractTopologyType)
		}
		stage.OnAfterAbstractTopologyTypeUpdateCallback.OnAfterUpdate(stage, abstracttopologytype, frontAbstractTopologyType)
	}
}

func (abstracttopologytype *AbstractTopologyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractTopologyTypeDeleteCallback != nil {
		var frontAbstractTopologyType *AbstractTopologyType
		if front != nil {
			frontAbstractTopologyType, _ = front.(*AbstractTopologyType)
		}
		stage.OnAfterAbstractTopologyTypeDeleteCallback.OnAfterDelete(stage, abstracttopologytype, frontAbstractTopologyType)
	}
}

func (affinecspropertytype *AffineCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAffineCSPropertyTypeCreateCallback != nil {
		stage.OnAfterAffineCSPropertyTypeCreateCallback.OnAfterCreate(stage, affinecspropertytype)
	}
}

func (affinecspropertytype *AffineCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAffineCSPropertyTypeUpdateCallback != nil {
		var frontAffineCSPropertyType *AffineCSPropertyType
		if front != nil {
			frontAffineCSPropertyType, _ = front.(*AffineCSPropertyType)
		}
		stage.OnAfterAffineCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, affinecspropertytype, frontAffineCSPropertyType)
	}
}

func (affinecspropertytype *AffineCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAffineCSPropertyTypeDeleteCallback != nil {
		var frontAffineCSPropertyType *AffineCSPropertyType
		if front != nil {
			frontAffineCSPropertyType, _ = front.(*AffineCSPropertyType)
		}
		stage.OnAfterAffineCSPropertyTypeDeleteCallback.OnAfterDelete(stage, affinecspropertytype, frontAffineCSPropertyType)
	}
}

func (affinecstype *AffineCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAffineCSTypeCreateCallback != nil {
		stage.OnAfterAffineCSTypeCreateCallback.OnAfterCreate(stage, affinecstype)
	}
}

func (affinecstype *AffineCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAffineCSTypeUpdateCallback != nil {
		var frontAffineCSType *AffineCSType
		if front != nil {
			frontAffineCSType, _ = front.(*AffineCSType)
		}
		stage.OnAfterAffineCSTypeUpdateCallback.OnAfterUpdate(stage, affinecstype, frontAffineCSType)
	}
}

func (affinecstype *AffineCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAffineCSTypeDeleteCallback != nil {
		var frontAffineCSType *AffineCSType
		if front != nil {
			frontAffineCSType, _ = front.(*AffineCSType)
		}
		stage.OnAfterAffineCSTypeDeleteCallback.OnAfterDelete(stage, affinecstype, frontAffineCSType)
	}
}

func (affineplacementtype *AffinePlacementType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAffinePlacementTypeCreateCallback != nil {
		stage.OnAfterAffinePlacementTypeCreateCallback.OnAfterCreate(stage, affineplacementtype)
	}
}

func (affineplacementtype *AffinePlacementType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAffinePlacementTypeUpdateCallback != nil {
		var frontAffinePlacementType *AffinePlacementType
		if front != nil {
			frontAffinePlacementType, _ = front.(*AffinePlacementType)
		}
		stage.OnAfterAffinePlacementTypeUpdateCallback.OnAfterUpdate(stage, affineplacementtype, frontAffinePlacementType)
	}
}

func (affineplacementtype *AffinePlacementType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAffinePlacementTypeDeleteCallback != nil {
		var frontAffinePlacementType *AffinePlacementType
		if front != nil {
			frontAffinePlacementType, _ = front.(*AffinePlacementType)
		}
		stage.OnAfterAffinePlacementTypeDeleteCallback.OnAfterDelete(stage, affineplacementtype, frontAffinePlacementType)
	}
}

func (anglechoicetype *AngleChoiceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngleChoiceTypeCreateCallback != nil {
		stage.OnAfterAngleChoiceTypeCreateCallback.OnAfterCreate(stage, anglechoicetype)
	}
}

func (anglechoicetype *AngleChoiceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleChoiceTypeUpdateCallback != nil {
		var frontAngleChoiceType *AngleChoiceType
		if front != nil {
			frontAngleChoiceType, _ = front.(*AngleChoiceType)
		}
		stage.OnAfterAngleChoiceTypeUpdateCallback.OnAfterUpdate(stage, anglechoicetype, frontAngleChoiceType)
	}
}

func (anglechoicetype *AngleChoiceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngleChoiceTypeDeleteCallback != nil {
		var frontAngleChoiceType *AngleChoiceType
		if front != nil {
			frontAngleChoiceType, _ = front.(*AngleChoiceType)
		}
		stage.OnAfterAngleChoiceTypeDeleteCallback.OnAfterDelete(stage, anglechoicetype, frontAngleChoiceType)
	}
}

func (angle_propertytype *Angle_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngle_PropertyTypeCreateCallback != nil {
		stage.OnAfterAngle_PropertyTypeCreateCallback.OnAfterCreate(stage, angle_propertytype)
	}
}

func (angle_propertytype *Angle_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngle_PropertyTypeUpdateCallback != nil {
		var frontAngle_PropertyType *Angle_PropertyType
		if front != nil {
			frontAngle_PropertyType, _ = front.(*Angle_PropertyType)
		}
		stage.OnAfterAngle_PropertyTypeUpdateCallback.OnAfterUpdate(stage, angle_propertytype, frontAngle_PropertyType)
	}
}

func (angle_propertytype *Angle_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngle_PropertyTypeDeleteCallback != nil {
		var frontAngle_PropertyType *Angle_PropertyType
		if front != nil {
			frontAngle_PropertyType, _ = front.(*Angle_PropertyType)
		}
		stage.OnAfterAngle_PropertyTypeDeleteCallback.OnAfterDelete(stage, angle_propertytype, frontAngle_PropertyType)
	}
}

func (arcbybulgetype *ArcByBulgeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcByBulgeTypeCreateCallback != nil {
		stage.OnAfterArcByBulgeTypeCreateCallback.OnAfterCreate(stage, arcbybulgetype)
	}
}

func (arcbybulgetype *ArcByBulgeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcByBulgeTypeUpdateCallback != nil {
		var frontArcByBulgeType *ArcByBulgeType
		if front != nil {
			frontArcByBulgeType, _ = front.(*ArcByBulgeType)
		}
		stage.OnAfterArcByBulgeTypeUpdateCallback.OnAfterUpdate(stage, arcbybulgetype, frontArcByBulgeType)
	}
}

func (arcbybulgetype *ArcByBulgeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcByBulgeTypeDeleteCallback != nil {
		var frontArcByBulgeType *ArcByBulgeType
		if front != nil {
			frontArcByBulgeType, _ = front.(*ArcByBulgeType)
		}
		stage.OnAfterArcByBulgeTypeDeleteCallback.OnAfterDelete(stage, arcbybulgetype, frontArcByBulgeType)
	}
}

func (arcbycenterpointtype *ArcByCenterPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcByCenterPointTypeCreateCallback != nil {
		stage.OnAfterArcByCenterPointTypeCreateCallback.OnAfterCreate(stage, arcbycenterpointtype)
	}
}

func (arcbycenterpointtype *ArcByCenterPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcByCenterPointTypeUpdateCallback != nil {
		var frontArcByCenterPointType *ArcByCenterPointType
		if front != nil {
			frontArcByCenterPointType, _ = front.(*ArcByCenterPointType)
		}
		stage.OnAfterArcByCenterPointTypeUpdateCallback.OnAfterUpdate(stage, arcbycenterpointtype, frontArcByCenterPointType)
	}
}

func (arcbycenterpointtype *ArcByCenterPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcByCenterPointTypeDeleteCallback != nil {
		var frontArcByCenterPointType *ArcByCenterPointType
		if front != nil {
			frontArcByCenterPointType, _ = front.(*ArcByCenterPointType)
		}
		stage.OnAfterArcByCenterPointTypeDeleteCallback.OnAfterDelete(stage, arcbycenterpointtype, frontArcByCenterPointType)
	}
}

func (arcstringbybulgetype *ArcStringByBulgeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcStringByBulgeTypeCreateCallback != nil {
		stage.OnAfterArcStringByBulgeTypeCreateCallback.OnAfterCreate(stage, arcstringbybulgetype)
	}
}

func (arcstringbybulgetype *ArcStringByBulgeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcStringByBulgeTypeUpdateCallback != nil {
		var frontArcStringByBulgeType *ArcStringByBulgeType
		if front != nil {
			frontArcStringByBulgeType, _ = front.(*ArcStringByBulgeType)
		}
		stage.OnAfterArcStringByBulgeTypeUpdateCallback.OnAfterUpdate(stage, arcstringbybulgetype, frontArcStringByBulgeType)
	}
}

func (arcstringbybulgetype *ArcStringByBulgeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcStringByBulgeTypeDeleteCallback != nil {
		var frontArcStringByBulgeType *ArcStringByBulgeType
		if front != nil {
			frontArcStringByBulgeType, _ = front.(*ArcStringByBulgeType)
		}
		stage.OnAfterArcStringByBulgeTypeDeleteCallback.OnAfterDelete(stage, arcstringbybulgetype, frontArcStringByBulgeType)
	}
}

func (arcstringtype *ArcStringType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcStringTypeCreateCallback != nil {
		stage.OnAfterArcStringTypeCreateCallback.OnAfterCreate(stage, arcstringtype)
	}
}

func (arcstringtype *ArcStringType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcStringTypeUpdateCallback != nil {
		var frontArcStringType *ArcStringType
		if front != nil {
			frontArcStringType, _ = front.(*ArcStringType)
		}
		stage.OnAfterArcStringTypeUpdateCallback.OnAfterUpdate(stage, arcstringtype, frontArcStringType)
	}
}

func (arcstringtype *ArcStringType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcStringTypeDeleteCallback != nil {
		var frontArcStringType *ArcStringType
		if front != nil {
			frontArcStringType, _ = front.(*ArcStringType)
		}
		stage.OnAfterArcStringTypeDeleteCallback.OnAfterDelete(stage, arcstringtype, frontArcStringType)
	}
}

func (arctype *ArcType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcTypeCreateCallback != nil {
		stage.OnAfterArcTypeCreateCallback.OnAfterCreate(stage, arctype)
	}
}

func (arctype *ArcType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcTypeUpdateCallback != nil {
		var frontArcType *ArcType
		if front != nil {
			frontArcType, _ = front.(*ArcType)
		}
		stage.OnAfterArcTypeUpdateCallback.OnAfterUpdate(stage, arctype, frontArcType)
	}
}

func (arctype *ArcType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcTypeDeleteCallback != nil {
		var frontArcType *ArcType
		if front != nil {
			frontArcType, _ = front.(*ArcType)
		}
		stage.OnAfterArcTypeDeleteCallback.OnAfterDelete(stage, arctype, frontArcType)
	}
}

func (arrayassociationtype *ArrayAssociationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrayAssociationTypeCreateCallback != nil {
		stage.OnAfterArrayAssociationTypeCreateCallback.OnAfterCreate(stage, arrayassociationtype)
	}
}

func (arrayassociationtype *ArrayAssociationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrayAssociationTypeUpdateCallback != nil {
		var frontArrayAssociationType *ArrayAssociationType
		if front != nil {
			frontArrayAssociationType, _ = front.(*ArrayAssociationType)
		}
		stage.OnAfterArrayAssociationTypeUpdateCallback.OnAfterUpdate(stage, arrayassociationtype, frontArrayAssociationType)
	}
}

func (arrayassociationtype *ArrayAssociationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrayAssociationTypeDeleteCallback != nil {
		var frontArrayAssociationType *ArrayAssociationType
		if front != nil {
			frontArrayAssociationType, _ = front.(*ArrayAssociationType)
		}
		stage.OnAfterArrayAssociationTypeDeleteCallback.OnAfterDelete(stage, arrayassociationtype, frontArrayAssociationType)
	}
}

func (arraytype *ArrayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrayTypeCreateCallback != nil {
		stage.OnAfterArrayTypeCreateCallback.OnAfterCreate(stage, arraytype)
	}
}

func (arraytype *ArrayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrayTypeUpdateCallback != nil {
		var frontArrayType *ArrayType
		if front != nil {
			frontArrayType, _ = front.(*ArrayType)
		}
		stage.OnAfterArrayTypeUpdateCallback.OnAfterUpdate(stage, arraytype, frontArrayType)
	}
}

func (arraytype *ArrayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrayTypeDeleteCallback != nil {
		var frontArrayType *ArrayType
		if front != nil {
			frontArrayType, _ = front.(*ArrayType)
		}
		stage.OnAfterArrayTypeDeleteCallback.OnAfterDelete(stage, arraytype, frontArrayType)
	}
}

func (associationroletype *AssociationRoleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAssociationRoleTypeCreateCallback != nil {
		stage.OnAfterAssociationRoleTypeCreateCallback.OnAfterCreate(stage, associationroletype)
	}
}

func (associationroletype *AssociationRoleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAssociationRoleTypeUpdateCallback != nil {
		var frontAssociationRoleType *AssociationRoleType
		if front != nil {
			frontAssociationRoleType, _ = front.(*AssociationRoleType)
		}
		stage.OnAfterAssociationRoleTypeUpdateCallback.OnAfterUpdate(stage, associationroletype, frontAssociationRoleType)
	}
}

func (associationroletype *AssociationRoleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAssociationRoleTypeDeleteCallback != nil {
		var frontAssociationRoleType *AssociationRoleType
		if front != nil {
			frontAssociationRoleType, _ = front.(*AssociationRoleType)
		}
		stage.OnAfterAssociationRoleTypeDeleteCallback.OnAfterDelete(stage, associationroletype, frontAssociationRoleType)
	}
}

func (bsplinetype *BSplineType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBSplineTypeCreateCallback != nil {
		stage.OnAfterBSplineTypeCreateCallback.OnAfterCreate(stage, bsplinetype)
	}
}

func (bsplinetype *BSplineType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBSplineTypeUpdateCallback != nil {
		var frontBSplineType *BSplineType
		if front != nil {
			frontBSplineType, _ = front.(*BSplineType)
		}
		stage.OnAfterBSplineTypeUpdateCallback.OnAfterUpdate(stage, bsplinetype, frontBSplineType)
	}
}

func (bsplinetype *BSplineType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBSplineTypeDeleteCallback != nil {
		var frontBSplineType *BSplineType
		if front != nil {
			frontBSplineType, _ = front.(*BSplineType)
		}
		stage.OnAfterBSplineTypeDeleteCallback.OnAfterDelete(stage, bsplinetype, frontBSplineType)
	}
}

func (bagtype *BagType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBagTypeCreateCallback != nil {
		stage.OnAfterBagTypeCreateCallback.OnAfterCreate(stage, bagtype)
	}
}

func (bagtype *BagType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBagTypeUpdateCallback != nil {
		var frontBagType *BagType
		if front != nil {
			frontBagType, _ = front.(*BagType)
		}
		stage.OnAfterBagTypeUpdateCallback.OnAfterUpdate(stage, bagtype, frontBagType)
	}
}

func (bagtype *BagType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBagTypeDeleteCallback != nil {
		var frontBagType *BagType
		if front != nil {
			frontBagType, _ = front.(*BagType)
		}
		stage.OnAfterBagTypeDeleteCallback.OnAfterDelete(stage, bagtype, frontBagType)
	}
}

func (baseunittype *BaseUnitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBaseUnitTypeCreateCallback != nil {
		stage.OnAfterBaseUnitTypeCreateCallback.OnAfterCreate(stage, baseunittype)
	}
}

func (baseunittype *BaseUnitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBaseUnitTypeUpdateCallback != nil {
		var frontBaseUnitType *BaseUnitType
		if front != nil {
			frontBaseUnitType, _ = front.(*BaseUnitType)
		}
		stage.OnAfterBaseUnitTypeUpdateCallback.OnAfterUpdate(stage, baseunittype, frontBaseUnitType)
	}
}

func (baseunittype *BaseUnitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBaseUnitTypeDeleteCallback != nil {
		var frontBaseUnitType *BaseUnitType
		if front != nil {
			frontBaseUnitType, _ = front.(*BaseUnitType)
		}
		stage.OnAfterBaseUnitTypeDeleteCallback.OnAfterDelete(stage, baseunittype, frontBaseUnitType)
	}
}

func (beziertype *BezierType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBezierTypeCreateCallback != nil {
		stage.OnAfterBezierTypeCreateCallback.OnAfterCreate(stage, beziertype)
	}
}

func (beziertype *BezierType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBezierTypeUpdateCallback != nil {
		var frontBezierType *BezierType
		if front != nil {
			frontBezierType, _ = front.(*BezierType)
		}
		stage.OnAfterBezierTypeUpdateCallback.OnAfterUpdate(stage, beziertype, frontBezierType)
	}
}

func (beziertype *BezierType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBezierTypeDeleteCallback != nil {
		var frontBezierType *BezierType
		if front != nil {
			frontBezierType, _ = front.(*BezierType)
		}
		stage.OnAfterBezierTypeDeleteCallback.OnAfterDelete(stage, beziertype, frontBezierType)
	}
}

func (binary_propertytype *Binary_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBinary_PropertyTypeCreateCallback != nil {
		stage.OnAfterBinary_PropertyTypeCreateCallback.OnAfterCreate(stage, binary_propertytype)
	}
}

func (binary_propertytype *Binary_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBinary_PropertyTypeUpdateCallback != nil {
		var frontBinary_PropertyType *Binary_PropertyType
		if front != nil {
			frontBinary_PropertyType, _ = front.(*Binary_PropertyType)
		}
		stage.OnAfterBinary_PropertyTypeUpdateCallback.OnAfterUpdate(stage, binary_propertytype, frontBinary_PropertyType)
	}
}

func (binary_propertytype *Binary_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBinary_PropertyTypeDeleteCallback != nil {
		var frontBinary_PropertyType *Binary_PropertyType
		if front != nil {
			frontBinary_PropertyType, _ = front.(*Binary_PropertyType)
		}
		stage.OnAfterBinary_PropertyTypeDeleteCallback.OnAfterDelete(stage, binary_propertytype, frontBinary_PropertyType)
	}
}

func (binary_type *Binary_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBinary_TypeCreateCallback != nil {
		stage.OnAfterBinary_TypeCreateCallback.OnAfterCreate(stage, binary_type)
	}
}

func (binary_type *Binary_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBinary_TypeUpdateCallback != nil {
		var frontBinary_Type *Binary_Type
		if front != nil {
			frontBinary_Type, _ = front.(*Binary_Type)
		}
		stage.OnAfterBinary_TypeUpdateCallback.OnAfterUpdate(stage, binary_type, frontBinary_Type)
	}
}

func (binary_type *Binary_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBinary_TypeDeleteCallback != nil {
		var frontBinary_Type *Binary_Type
		if front != nil {
			frontBinary_Type, _ = front.(*Binary_Type)
		}
		stage.OnAfterBinary_TypeDeleteCallback.OnAfterDelete(stage, binary_type, frontBinary_Type)
	}
}

func (boolean *Boolean) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBooleanCreateCallback != nil {
		stage.OnAfterBooleanCreateCallback.OnAfterCreate(stage, boolean)
	}
}

func (boolean *Boolean) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBooleanUpdateCallback != nil {
		var frontBoolean *Boolean
		if front != nil {
			frontBoolean, _ = front.(*Boolean)
		}
		stage.OnAfterBooleanUpdateCallback.OnAfterUpdate(stage, boolean, frontBoolean)
	}
}

func (boolean *Boolean) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBooleanDeleteCallback != nil {
		var frontBoolean *Boolean
		if front != nil {
			frontBoolean, _ = front.(*Boolean)
		}
		stage.OnAfterBooleanDeleteCallback.OnAfterDelete(stage, boolean, frontBoolean)
	}
}

func (booleanpropertytype *BooleanPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBooleanPropertyTypeCreateCallback != nil {
		stage.OnAfterBooleanPropertyTypeCreateCallback.OnAfterCreate(stage, booleanpropertytype)
	}
}

func (booleanpropertytype *BooleanPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBooleanPropertyTypeUpdateCallback != nil {
		var frontBooleanPropertyType *BooleanPropertyType
		if front != nil {
			frontBooleanPropertyType, _ = front.(*BooleanPropertyType)
		}
		stage.OnAfterBooleanPropertyTypeUpdateCallback.OnAfterUpdate(stage, booleanpropertytype, frontBooleanPropertyType)
	}
}

func (booleanpropertytype *BooleanPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBooleanPropertyTypeDeleteCallback != nil {
		var frontBooleanPropertyType *BooleanPropertyType
		if front != nil {
			frontBooleanPropertyType, _ = front.(*BooleanPropertyType)
		}
		stage.OnAfterBooleanPropertyTypeDeleteCallback.OnAfterDelete(stage, booleanpropertytype, frontBooleanPropertyType)
	}
}

func (boolean_propertytype *Boolean_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBoolean_PropertyTypeCreateCallback != nil {
		stage.OnAfterBoolean_PropertyTypeCreateCallback.OnAfterCreate(stage, boolean_propertytype)
	}
}

func (boolean_propertytype *Boolean_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoolean_PropertyTypeUpdateCallback != nil {
		var frontBoolean_PropertyType *Boolean_PropertyType
		if front != nil {
			frontBoolean_PropertyType, _ = front.(*Boolean_PropertyType)
		}
		stage.OnAfterBoolean_PropertyTypeUpdateCallback.OnAfterUpdate(stage, boolean_propertytype, frontBoolean_PropertyType)
	}
}

func (boolean_propertytype *Boolean_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoolean_PropertyTypeDeleteCallback != nil {
		var frontBoolean_PropertyType *Boolean_PropertyType
		if front != nil {
			frontBoolean_PropertyType, _ = front.(*Boolean_PropertyType)
		}
		stage.OnAfterBoolean_PropertyTypeDeleteCallback.OnAfterDelete(stage, boolean_propertytype, frontBoolean_PropertyType)
	}
}

func (boundedfeaturetype *BoundedFeatureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBoundedFeatureTypeCreateCallback != nil {
		stage.OnAfterBoundedFeatureTypeCreateCallback.OnAfterCreate(stage, boundedfeaturetype)
	}
}

func (boundedfeaturetype *BoundedFeatureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoundedFeatureTypeUpdateCallback != nil {
		var frontBoundedFeatureType *BoundedFeatureType
		if front != nil {
			frontBoundedFeatureType, _ = front.(*BoundedFeatureType)
		}
		stage.OnAfterBoundedFeatureTypeUpdateCallback.OnAfterUpdate(stage, boundedfeaturetype, frontBoundedFeatureType)
	}
}

func (boundedfeaturetype *BoundedFeatureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoundedFeatureTypeDeleteCallback != nil {
		var frontBoundedFeatureType *BoundedFeatureType
		if front != nil {
			frontBoundedFeatureType, _ = front.(*BoundedFeatureType)
		}
		stage.OnAfterBoundedFeatureTypeDeleteCallback.OnAfterDelete(stage, boundedfeaturetype, frontBoundedFeatureType)
	}
}

func (boundingshapetype *BoundingShapeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBoundingShapeTypeCreateCallback != nil {
		stage.OnAfterBoundingShapeTypeCreateCallback.OnAfterCreate(stage, boundingshapetype)
	}
}

func (boundingshapetype *BoundingShapeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoundingShapeTypeUpdateCallback != nil {
		var frontBoundingShapeType *BoundingShapeType
		if front != nil {
			frontBoundingShapeType, _ = front.(*BoundingShapeType)
		}
		stage.OnAfterBoundingShapeTypeUpdateCallback.OnAfterUpdate(stage, boundingshapetype, frontBoundingShapeType)
	}
}

func (boundingshapetype *BoundingShapeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBoundingShapeTypeDeleteCallback != nil {
		var frontBoundingShapeType *BoundingShapeType
		if front != nil {
			frontBoundingShapeType, _ = front.(*BoundingShapeType)
		}
		stage.OnAfterBoundingShapeTypeDeleteCallback.OnAfterDelete(stage, boundingshapetype, frontBoundingShapeType)
	}
}

func (ci_address_propertytype *CI_Address_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Address_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_Address_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_address_propertytype)
	}
}

func (ci_address_propertytype *CI_Address_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Address_PropertyTypeUpdateCallback != nil {
		var frontCI_Address_PropertyType *CI_Address_PropertyType
		if front != nil {
			frontCI_Address_PropertyType, _ = front.(*CI_Address_PropertyType)
		}
		stage.OnAfterCI_Address_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_address_propertytype, frontCI_Address_PropertyType)
	}
}

func (ci_address_propertytype *CI_Address_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Address_PropertyTypeDeleteCallback != nil {
		var frontCI_Address_PropertyType *CI_Address_PropertyType
		if front != nil {
			frontCI_Address_PropertyType, _ = front.(*CI_Address_PropertyType)
		}
		stage.OnAfterCI_Address_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_address_propertytype, frontCI_Address_PropertyType)
	}
}

func (ci_address_type *CI_Address_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Address_TypeCreateCallback != nil {
		stage.OnAfterCI_Address_TypeCreateCallback.OnAfterCreate(stage, ci_address_type)
	}
}

func (ci_address_type *CI_Address_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Address_TypeUpdateCallback != nil {
		var frontCI_Address_Type *CI_Address_Type
		if front != nil {
			frontCI_Address_Type, _ = front.(*CI_Address_Type)
		}
		stage.OnAfterCI_Address_TypeUpdateCallback.OnAfterUpdate(stage, ci_address_type, frontCI_Address_Type)
	}
}

func (ci_address_type *CI_Address_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Address_TypeDeleteCallback != nil {
		var frontCI_Address_Type *CI_Address_Type
		if front != nil {
			frontCI_Address_Type, _ = front.(*CI_Address_Type)
		}
		stage.OnAfterCI_Address_TypeDeleteCallback.OnAfterDelete(stage, ci_address_type, frontCI_Address_Type)
	}
}

func (ci_citation_propertytype *CI_Citation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Citation_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_Citation_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_citation_propertytype)
	}
}

func (ci_citation_propertytype *CI_Citation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Citation_PropertyTypeUpdateCallback != nil {
		var frontCI_Citation_PropertyType *CI_Citation_PropertyType
		if front != nil {
			frontCI_Citation_PropertyType, _ = front.(*CI_Citation_PropertyType)
		}
		stage.OnAfterCI_Citation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_citation_propertytype, frontCI_Citation_PropertyType)
	}
}

func (ci_citation_propertytype *CI_Citation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Citation_PropertyTypeDeleteCallback != nil {
		var frontCI_Citation_PropertyType *CI_Citation_PropertyType
		if front != nil {
			frontCI_Citation_PropertyType, _ = front.(*CI_Citation_PropertyType)
		}
		stage.OnAfterCI_Citation_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_citation_propertytype, frontCI_Citation_PropertyType)
	}
}

func (ci_citation_type *CI_Citation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Citation_TypeCreateCallback != nil {
		stage.OnAfterCI_Citation_TypeCreateCallback.OnAfterCreate(stage, ci_citation_type)
	}
}

func (ci_citation_type *CI_Citation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Citation_TypeUpdateCallback != nil {
		var frontCI_Citation_Type *CI_Citation_Type
		if front != nil {
			frontCI_Citation_Type, _ = front.(*CI_Citation_Type)
		}
		stage.OnAfterCI_Citation_TypeUpdateCallback.OnAfterUpdate(stage, ci_citation_type, frontCI_Citation_Type)
	}
}

func (ci_citation_type *CI_Citation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Citation_TypeDeleteCallback != nil {
		var frontCI_Citation_Type *CI_Citation_Type
		if front != nil {
			frontCI_Citation_Type, _ = front.(*CI_Citation_Type)
		}
		stage.OnAfterCI_Citation_TypeDeleteCallback.OnAfterDelete(stage, ci_citation_type, frontCI_Citation_Type)
	}
}

func (ci_contact_propertytype *CI_Contact_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Contact_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_Contact_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_contact_propertytype)
	}
}

func (ci_contact_propertytype *CI_Contact_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Contact_PropertyTypeUpdateCallback != nil {
		var frontCI_Contact_PropertyType *CI_Contact_PropertyType
		if front != nil {
			frontCI_Contact_PropertyType, _ = front.(*CI_Contact_PropertyType)
		}
		stage.OnAfterCI_Contact_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_contact_propertytype, frontCI_Contact_PropertyType)
	}
}

func (ci_contact_propertytype *CI_Contact_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Contact_PropertyTypeDeleteCallback != nil {
		var frontCI_Contact_PropertyType *CI_Contact_PropertyType
		if front != nil {
			frontCI_Contact_PropertyType, _ = front.(*CI_Contact_PropertyType)
		}
		stage.OnAfterCI_Contact_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_contact_propertytype, frontCI_Contact_PropertyType)
	}
}

func (ci_contact_type *CI_Contact_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Contact_TypeCreateCallback != nil {
		stage.OnAfterCI_Contact_TypeCreateCallback.OnAfterCreate(stage, ci_contact_type)
	}
}

func (ci_contact_type *CI_Contact_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Contact_TypeUpdateCallback != nil {
		var frontCI_Contact_Type *CI_Contact_Type
		if front != nil {
			frontCI_Contact_Type, _ = front.(*CI_Contact_Type)
		}
		stage.OnAfterCI_Contact_TypeUpdateCallback.OnAfterUpdate(stage, ci_contact_type, frontCI_Contact_Type)
	}
}

func (ci_contact_type *CI_Contact_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Contact_TypeDeleteCallback != nil {
		var frontCI_Contact_Type *CI_Contact_Type
		if front != nil {
			frontCI_Contact_Type, _ = front.(*CI_Contact_Type)
		}
		stage.OnAfterCI_Contact_TypeDeleteCallback.OnAfterDelete(stage, ci_contact_type, frontCI_Contact_Type)
	}
}

func (ci_datetypecode_propertytype *CI_DateTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_DateTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_DateTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_datetypecode_propertytype)
	}
}

func (ci_datetypecode_propertytype *CI_DateTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_DateTypeCode_PropertyTypeUpdateCallback != nil {
		var frontCI_DateTypeCode_PropertyType *CI_DateTypeCode_PropertyType
		if front != nil {
			frontCI_DateTypeCode_PropertyType, _ = front.(*CI_DateTypeCode_PropertyType)
		}
		stage.OnAfterCI_DateTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_datetypecode_propertytype, frontCI_DateTypeCode_PropertyType)
	}
}

func (ci_datetypecode_propertytype *CI_DateTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_DateTypeCode_PropertyTypeDeleteCallback != nil {
		var frontCI_DateTypeCode_PropertyType *CI_DateTypeCode_PropertyType
		if front != nil {
			frontCI_DateTypeCode_PropertyType, _ = front.(*CI_DateTypeCode_PropertyType)
		}
		stage.OnAfterCI_DateTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_datetypecode_propertytype, frontCI_DateTypeCode_PropertyType)
	}
}

func (ci_date_propertytype *CI_Date_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Date_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_Date_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_date_propertytype)
	}
}

func (ci_date_propertytype *CI_Date_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Date_PropertyTypeUpdateCallback != nil {
		var frontCI_Date_PropertyType *CI_Date_PropertyType
		if front != nil {
			frontCI_Date_PropertyType, _ = front.(*CI_Date_PropertyType)
		}
		stage.OnAfterCI_Date_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_date_propertytype, frontCI_Date_PropertyType)
	}
}

func (ci_date_propertytype *CI_Date_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Date_PropertyTypeDeleteCallback != nil {
		var frontCI_Date_PropertyType *CI_Date_PropertyType
		if front != nil {
			frontCI_Date_PropertyType, _ = front.(*CI_Date_PropertyType)
		}
		stage.OnAfterCI_Date_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_date_propertytype, frontCI_Date_PropertyType)
	}
}

func (ci_date_type *CI_Date_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Date_TypeCreateCallback != nil {
		stage.OnAfterCI_Date_TypeCreateCallback.OnAfterCreate(stage, ci_date_type)
	}
}

func (ci_date_type *CI_Date_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Date_TypeUpdateCallback != nil {
		var frontCI_Date_Type *CI_Date_Type
		if front != nil {
			frontCI_Date_Type, _ = front.(*CI_Date_Type)
		}
		stage.OnAfterCI_Date_TypeUpdateCallback.OnAfterUpdate(stage, ci_date_type, frontCI_Date_Type)
	}
}

func (ci_date_type *CI_Date_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Date_TypeDeleteCallback != nil {
		var frontCI_Date_Type *CI_Date_Type
		if front != nil {
			frontCI_Date_Type, _ = front.(*CI_Date_Type)
		}
		stage.OnAfterCI_Date_TypeDeleteCallback.OnAfterDelete(stage, ci_date_type, frontCI_Date_Type)
	}
}

func (ci_onlinefunctioncode_propertytype *CI_OnLineFunctionCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_OnLineFunctionCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_OnLineFunctionCode_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_onlinefunctioncode_propertytype)
	}
}

func (ci_onlinefunctioncode_propertytype *CI_OnLineFunctionCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_OnLineFunctionCode_PropertyTypeUpdateCallback != nil {
		var frontCI_OnLineFunctionCode_PropertyType *CI_OnLineFunctionCode_PropertyType
		if front != nil {
			frontCI_OnLineFunctionCode_PropertyType, _ = front.(*CI_OnLineFunctionCode_PropertyType)
		}
		stage.OnAfterCI_OnLineFunctionCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_onlinefunctioncode_propertytype, frontCI_OnLineFunctionCode_PropertyType)
	}
}

func (ci_onlinefunctioncode_propertytype *CI_OnLineFunctionCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_OnLineFunctionCode_PropertyTypeDeleteCallback != nil {
		var frontCI_OnLineFunctionCode_PropertyType *CI_OnLineFunctionCode_PropertyType
		if front != nil {
			frontCI_OnLineFunctionCode_PropertyType, _ = front.(*CI_OnLineFunctionCode_PropertyType)
		}
		stage.OnAfterCI_OnLineFunctionCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_onlinefunctioncode_propertytype, frontCI_OnLineFunctionCode_PropertyType)
	}
}

func (ci_onlineresource_propertytype *CI_OnlineResource_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_OnlineResource_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_OnlineResource_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_onlineresource_propertytype)
	}
}

func (ci_onlineresource_propertytype *CI_OnlineResource_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_OnlineResource_PropertyTypeUpdateCallback != nil {
		var frontCI_OnlineResource_PropertyType *CI_OnlineResource_PropertyType
		if front != nil {
			frontCI_OnlineResource_PropertyType, _ = front.(*CI_OnlineResource_PropertyType)
		}
		stage.OnAfterCI_OnlineResource_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_onlineresource_propertytype, frontCI_OnlineResource_PropertyType)
	}
}

func (ci_onlineresource_propertytype *CI_OnlineResource_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_OnlineResource_PropertyTypeDeleteCallback != nil {
		var frontCI_OnlineResource_PropertyType *CI_OnlineResource_PropertyType
		if front != nil {
			frontCI_OnlineResource_PropertyType, _ = front.(*CI_OnlineResource_PropertyType)
		}
		stage.OnAfterCI_OnlineResource_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_onlineresource_propertytype, frontCI_OnlineResource_PropertyType)
	}
}

func (ci_onlineresource_type *CI_OnlineResource_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_OnlineResource_TypeCreateCallback != nil {
		stage.OnAfterCI_OnlineResource_TypeCreateCallback.OnAfterCreate(stage, ci_onlineresource_type)
	}
}

func (ci_onlineresource_type *CI_OnlineResource_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_OnlineResource_TypeUpdateCallback != nil {
		var frontCI_OnlineResource_Type *CI_OnlineResource_Type
		if front != nil {
			frontCI_OnlineResource_Type, _ = front.(*CI_OnlineResource_Type)
		}
		stage.OnAfterCI_OnlineResource_TypeUpdateCallback.OnAfterUpdate(stage, ci_onlineresource_type, frontCI_OnlineResource_Type)
	}
}

func (ci_onlineresource_type *CI_OnlineResource_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_OnlineResource_TypeDeleteCallback != nil {
		var frontCI_OnlineResource_Type *CI_OnlineResource_Type
		if front != nil {
			frontCI_OnlineResource_Type, _ = front.(*CI_OnlineResource_Type)
		}
		stage.OnAfterCI_OnlineResource_TypeDeleteCallback.OnAfterDelete(stage, ci_onlineresource_type, frontCI_OnlineResource_Type)
	}
}

func (ci_presentationformcode_propertytype *CI_PresentationFormCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_PresentationFormCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_PresentationFormCode_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_presentationformcode_propertytype)
	}
}

func (ci_presentationformcode_propertytype *CI_PresentationFormCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_PresentationFormCode_PropertyTypeUpdateCallback != nil {
		var frontCI_PresentationFormCode_PropertyType *CI_PresentationFormCode_PropertyType
		if front != nil {
			frontCI_PresentationFormCode_PropertyType, _ = front.(*CI_PresentationFormCode_PropertyType)
		}
		stage.OnAfterCI_PresentationFormCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_presentationformcode_propertytype, frontCI_PresentationFormCode_PropertyType)
	}
}

func (ci_presentationformcode_propertytype *CI_PresentationFormCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_PresentationFormCode_PropertyTypeDeleteCallback != nil {
		var frontCI_PresentationFormCode_PropertyType *CI_PresentationFormCode_PropertyType
		if front != nil {
			frontCI_PresentationFormCode_PropertyType, _ = front.(*CI_PresentationFormCode_PropertyType)
		}
		stage.OnAfterCI_PresentationFormCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_presentationformcode_propertytype, frontCI_PresentationFormCode_PropertyType)
	}
}

func (ci_responsibleparty_propertytype *CI_ResponsibleParty_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_ResponsibleParty_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_ResponsibleParty_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_responsibleparty_propertytype)
	}
}

func (ci_responsibleparty_propertytype *CI_ResponsibleParty_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_ResponsibleParty_PropertyTypeUpdateCallback != nil {
		var frontCI_ResponsibleParty_PropertyType *CI_ResponsibleParty_PropertyType
		if front != nil {
			frontCI_ResponsibleParty_PropertyType, _ = front.(*CI_ResponsibleParty_PropertyType)
		}
		stage.OnAfterCI_ResponsibleParty_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_responsibleparty_propertytype, frontCI_ResponsibleParty_PropertyType)
	}
}

func (ci_responsibleparty_propertytype *CI_ResponsibleParty_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_ResponsibleParty_PropertyTypeDeleteCallback != nil {
		var frontCI_ResponsibleParty_PropertyType *CI_ResponsibleParty_PropertyType
		if front != nil {
			frontCI_ResponsibleParty_PropertyType, _ = front.(*CI_ResponsibleParty_PropertyType)
		}
		stage.OnAfterCI_ResponsibleParty_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_responsibleparty_propertytype, frontCI_ResponsibleParty_PropertyType)
	}
}

func (ci_responsibleparty_type *CI_ResponsibleParty_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_ResponsibleParty_TypeCreateCallback != nil {
		stage.OnAfterCI_ResponsibleParty_TypeCreateCallback.OnAfterCreate(stage, ci_responsibleparty_type)
	}
}

func (ci_responsibleparty_type *CI_ResponsibleParty_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_ResponsibleParty_TypeUpdateCallback != nil {
		var frontCI_ResponsibleParty_Type *CI_ResponsibleParty_Type
		if front != nil {
			frontCI_ResponsibleParty_Type, _ = front.(*CI_ResponsibleParty_Type)
		}
		stage.OnAfterCI_ResponsibleParty_TypeUpdateCallback.OnAfterUpdate(stage, ci_responsibleparty_type, frontCI_ResponsibleParty_Type)
	}
}

func (ci_responsibleparty_type *CI_ResponsibleParty_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_ResponsibleParty_TypeDeleteCallback != nil {
		var frontCI_ResponsibleParty_Type *CI_ResponsibleParty_Type
		if front != nil {
			frontCI_ResponsibleParty_Type, _ = front.(*CI_ResponsibleParty_Type)
		}
		stage.OnAfterCI_ResponsibleParty_TypeDeleteCallback.OnAfterDelete(stage, ci_responsibleparty_type, frontCI_ResponsibleParty_Type)
	}
}

func (ci_rolecode_propertytype *CI_RoleCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_RoleCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_RoleCode_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_rolecode_propertytype)
	}
}

func (ci_rolecode_propertytype *CI_RoleCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_RoleCode_PropertyTypeUpdateCallback != nil {
		var frontCI_RoleCode_PropertyType *CI_RoleCode_PropertyType
		if front != nil {
			frontCI_RoleCode_PropertyType, _ = front.(*CI_RoleCode_PropertyType)
		}
		stage.OnAfterCI_RoleCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_rolecode_propertytype, frontCI_RoleCode_PropertyType)
	}
}

func (ci_rolecode_propertytype *CI_RoleCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_RoleCode_PropertyTypeDeleteCallback != nil {
		var frontCI_RoleCode_PropertyType *CI_RoleCode_PropertyType
		if front != nil {
			frontCI_RoleCode_PropertyType, _ = front.(*CI_RoleCode_PropertyType)
		}
		stage.OnAfterCI_RoleCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_rolecode_propertytype, frontCI_RoleCode_PropertyType)
	}
}

func (ci_series_propertytype *CI_Series_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Series_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_Series_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_series_propertytype)
	}
}

func (ci_series_propertytype *CI_Series_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Series_PropertyTypeUpdateCallback != nil {
		var frontCI_Series_PropertyType *CI_Series_PropertyType
		if front != nil {
			frontCI_Series_PropertyType, _ = front.(*CI_Series_PropertyType)
		}
		stage.OnAfterCI_Series_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_series_propertytype, frontCI_Series_PropertyType)
	}
}

func (ci_series_propertytype *CI_Series_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Series_PropertyTypeDeleteCallback != nil {
		var frontCI_Series_PropertyType *CI_Series_PropertyType
		if front != nil {
			frontCI_Series_PropertyType, _ = front.(*CI_Series_PropertyType)
		}
		stage.OnAfterCI_Series_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_series_propertytype, frontCI_Series_PropertyType)
	}
}

func (ci_series_type *CI_Series_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Series_TypeCreateCallback != nil {
		stage.OnAfterCI_Series_TypeCreateCallback.OnAfterCreate(stage, ci_series_type)
	}
}

func (ci_series_type *CI_Series_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Series_TypeUpdateCallback != nil {
		var frontCI_Series_Type *CI_Series_Type
		if front != nil {
			frontCI_Series_Type, _ = front.(*CI_Series_Type)
		}
		stage.OnAfterCI_Series_TypeUpdateCallback.OnAfterUpdate(stage, ci_series_type, frontCI_Series_Type)
	}
}

func (ci_series_type *CI_Series_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Series_TypeDeleteCallback != nil {
		var frontCI_Series_Type *CI_Series_Type
		if front != nil {
			frontCI_Series_Type, _ = front.(*CI_Series_Type)
		}
		stage.OnAfterCI_Series_TypeDeleteCallback.OnAfterDelete(stage, ci_series_type, frontCI_Series_Type)
	}
}

func (ci_telephone_propertytype *CI_Telephone_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Telephone_PropertyTypeCreateCallback != nil {
		stage.OnAfterCI_Telephone_PropertyTypeCreateCallback.OnAfterCreate(stage, ci_telephone_propertytype)
	}
}

func (ci_telephone_propertytype *CI_Telephone_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Telephone_PropertyTypeUpdateCallback != nil {
		var frontCI_Telephone_PropertyType *CI_Telephone_PropertyType
		if front != nil {
			frontCI_Telephone_PropertyType, _ = front.(*CI_Telephone_PropertyType)
		}
		stage.OnAfterCI_Telephone_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ci_telephone_propertytype, frontCI_Telephone_PropertyType)
	}
}

func (ci_telephone_propertytype *CI_Telephone_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Telephone_PropertyTypeDeleteCallback != nil {
		var frontCI_Telephone_PropertyType *CI_Telephone_PropertyType
		if front != nil {
			frontCI_Telephone_PropertyType, _ = front.(*CI_Telephone_PropertyType)
		}
		stage.OnAfterCI_Telephone_PropertyTypeDeleteCallback.OnAfterDelete(stage, ci_telephone_propertytype, frontCI_Telephone_PropertyType)
	}
}

func (ci_telephone_type *CI_Telephone_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCI_Telephone_TypeCreateCallback != nil {
		stage.OnAfterCI_Telephone_TypeCreateCallback.OnAfterCreate(stage, ci_telephone_type)
	}
}

func (ci_telephone_type *CI_Telephone_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Telephone_TypeUpdateCallback != nil {
		var frontCI_Telephone_Type *CI_Telephone_Type
		if front != nil {
			frontCI_Telephone_Type, _ = front.(*CI_Telephone_Type)
		}
		stage.OnAfterCI_Telephone_TypeUpdateCallback.OnAfterUpdate(stage, ci_telephone_type, frontCI_Telephone_Type)
	}
}

func (ci_telephone_type *CI_Telephone_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCI_Telephone_TypeDeleteCallback != nil {
		var frontCI_Telephone_Type *CI_Telephone_Type
		if front != nil {
			frontCI_Telephone_Type, _ = front.(*CI_Telephone_Type)
		}
		stage.OnAfterCI_Telephone_TypeDeleteCallback.OnAfterDelete(stage, ci_telephone_type, frontCI_Telephone_Type)
	}
}

func (crspropertytype *CRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterCRSPropertyTypeCreateCallback.OnAfterCreate(stage, crspropertytype)
	}
}

func (crspropertytype *CRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCRSPropertyTypeUpdateCallback != nil {
		var frontCRSPropertyType *CRSPropertyType
		if front != nil {
			frontCRSPropertyType, _ = front.(*CRSPropertyType)
		}
		stage.OnAfterCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, crspropertytype, frontCRSPropertyType)
	}
}

func (crspropertytype *CRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCRSPropertyTypeDeleteCallback != nil {
		var frontCRSPropertyType *CRSPropertyType
		if front != nil {
			frontCRSPropertyType, _ = front.(*CRSPropertyType)
		}
		stage.OnAfterCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, crspropertytype, frontCRSPropertyType)
	}
}

func (cartesiancspropertytype *CartesianCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCartesianCSPropertyTypeCreateCallback != nil {
		stage.OnAfterCartesianCSPropertyTypeCreateCallback.OnAfterCreate(stage, cartesiancspropertytype)
	}
}

func (cartesiancspropertytype *CartesianCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCartesianCSPropertyTypeUpdateCallback != nil {
		var frontCartesianCSPropertyType *CartesianCSPropertyType
		if front != nil {
			frontCartesianCSPropertyType, _ = front.(*CartesianCSPropertyType)
		}
		stage.OnAfterCartesianCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, cartesiancspropertytype, frontCartesianCSPropertyType)
	}
}

func (cartesiancspropertytype *CartesianCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCartesianCSPropertyTypeDeleteCallback != nil {
		var frontCartesianCSPropertyType *CartesianCSPropertyType
		if front != nil {
			frontCartesianCSPropertyType, _ = front.(*CartesianCSPropertyType)
		}
		stage.OnAfterCartesianCSPropertyTypeDeleteCallback.OnAfterDelete(stage, cartesiancspropertytype, frontCartesianCSPropertyType)
	}
}

func (cartesiancstype *CartesianCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCartesianCSTypeCreateCallback != nil {
		stage.OnAfterCartesianCSTypeCreateCallback.OnAfterCreate(stage, cartesiancstype)
	}
}

func (cartesiancstype *CartesianCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCartesianCSTypeUpdateCallback != nil {
		var frontCartesianCSType *CartesianCSType
		if front != nil {
			frontCartesianCSType, _ = front.(*CartesianCSType)
		}
		stage.OnAfterCartesianCSTypeUpdateCallback.OnAfterUpdate(stage, cartesiancstype, frontCartesianCSType)
	}
}

func (cartesiancstype *CartesianCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCartesianCSTypeDeleteCallback != nil {
		var frontCartesianCSType *CartesianCSType
		if front != nil {
			frontCartesianCSType, _ = front.(*CartesianCSType)
		}
		stage.OnAfterCartesianCSTypeDeleteCallback.OnAfterDelete(stage, cartesiancstype, frontCartesianCSType)
	}
}

func (category *Category) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCategoryCreateCallback != nil {
		stage.OnAfterCategoryCreateCallback.OnAfterCreate(stage, category)
	}
}

func (category *Category) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCategoryUpdateCallback != nil {
		var frontCategory *Category
		if front != nil {
			frontCategory, _ = front.(*Category)
		}
		stage.OnAfterCategoryUpdateCallback.OnAfterUpdate(stage, category, frontCategory)
	}
}

func (category *Category) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCategoryDeleteCallback != nil {
		var frontCategory *Category
		if front != nil {
			frontCategory, _ = front.(*Category)
		}
		stage.OnAfterCategoryDeleteCallback.OnAfterDelete(stage, category, frontCategory)
	}
}

func (categoryextenttype *CategoryExtentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCategoryExtentTypeCreateCallback != nil {
		stage.OnAfterCategoryExtentTypeCreateCallback.OnAfterCreate(stage, categoryextenttype)
	}
}

func (categoryextenttype *CategoryExtentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCategoryExtentTypeUpdateCallback != nil {
		var frontCategoryExtentType *CategoryExtentType
		if front != nil {
			frontCategoryExtentType, _ = front.(*CategoryExtentType)
		}
		stage.OnAfterCategoryExtentTypeUpdateCallback.OnAfterUpdate(stage, categoryextenttype, frontCategoryExtentType)
	}
}

func (categoryextenttype *CategoryExtentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCategoryExtentTypeDeleteCallback != nil {
		var frontCategoryExtentType *CategoryExtentType
		if front != nil {
			frontCategoryExtentType, _ = front.(*CategoryExtentType)
		}
		stage.OnAfterCategoryExtentTypeDeleteCallback.OnAfterDelete(stage, categoryextenttype, frontCategoryExtentType)
	}
}

func (categorypropertytype *CategoryPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCategoryPropertyTypeCreateCallback != nil {
		stage.OnAfterCategoryPropertyTypeCreateCallback.OnAfterCreate(stage, categorypropertytype)
	}
}

func (categorypropertytype *CategoryPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCategoryPropertyTypeUpdateCallback != nil {
		var frontCategoryPropertyType *CategoryPropertyType
		if front != nil {
			frontCategoryPropertyType, _ = front.(*CategoryPropertyType)
		}
		stage.OnAfterCategoryPropertyTypeUpdateCallback.OnAfterUpdate(stage, categorypropertytype, frontCategoryPropertyType)
	}
}

func (categorypropertytype *CategoryPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCategoryPropertyTypeDeleteCallback != nil {
		var frontCategoryPropertyType *CategoryPropertyType
		if front != nil {
			frontCategoryPropertyType, _ = front.(*CategoryPropertyType)
		}
		stage.OnAfterCategoryPropertyTypeDeleteCallback.OnAfterDelete(stage, categorypropertytype, frontCategoryPropertyType)
	}
}

func (characterstring_propertytype *CharacterString_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCharacterString_PropertyTypeCreateCallback != nil {
		stage.OnAfterCharacterString_PropertyTypeCreateCallback.OnAfterCreate(stage, characterstring_propertytype)
	}
}

func (characterstring_propertytype *CharacterString_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCharacterString_PropertyTypeUpdateCallback != nil {
		var frontCharacterString_PropertyType *CharacterString_PropertyType
		if front != nil {
			frontCharacterString_PropertyType, _ = front.(*CharacterString_PropertyType)
		}
		stage.OnAfterCharacterString_PropertyTypeUpdateCallback.OnAfterUpdate(stage, characterstring_propertytype, frontCharacterString_PropertyType)
	}
}

func (characterstring_propertytype *CharacterString_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCharacterString_PropertyTypeDeleteCallback != nil {
		var frontCharacterString_PropertyType *CharacterString_PropertyType
		if front != nil {
			frontCharacterString_PropertyType, _ = front.(*CharacterString_PropertyType)
		}
		stage.OnAfterCharacterString_PropertyTypeDeleteCallback.OnAfterDelete(stage, characterstring_propertytype, frontCharacterString_PropertyType)
	}
}

func (circlebycenterpointtype *CircleByCenterPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleByCenterPointTypeCreateCallback != nil {
		stage.OnAfterCircleByCenterPointTypeCreateCallback.OnAfterCreate(stage, circlebycenterpointtype)
	}
}

func (circlebycenterpointtype *CircleByCenterPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleByCenterPointTypeUpdateCallback != nil {
		var frontCircleByCenterPointType *CircleByCenterPointType
		if front != nil {
			frontCircleByCenterPointType, _ = front.(*CircleByCenterPointType)
		}
		stage.OnAfterCircleByCenterPointTypeUpdateCallback.OnAfterUpdate(stage, circlebycenterpointtype, frontCircleByCenterPointType)
	}
}

func (circlebycenterpointtype *CircleByCenterPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleByCenterPointTypeDeleteCallback != nil {
		var frontCircleByCenterPointType *CircleByCenterPointType
		if front != nil {
			frontCircleByCenterPointType, _ = front.(*CircleByCenterPointType)
		}
		stage.OnAfterCircleByCenterPointTypeDeleteCallback.OnAfterDelete(stage, circlebycenterpointtype, frontCircleByCenterPointType)
	}
}

func (circletype *CircleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleTypeCreateCallback != nil {
		stage.OnAfterCircleTypeCreateCallback.OnAfterCreate(stage, circletype)
	}
}

func (circletype *CircleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleTypeUpdateCallback != nil {
		var frontCircleType *CircleType
		if front != nil {
			frontCircleType, _ = front.(*CircleType)
		}
		stage.OnAfterCircleTypeUpdateCallback.OnAfterUpdate(stage, circletype, frontCircleType)
	}
}

func (circletype *CircleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleTypeDeleteCallback != nil {
		var frontCircleType *CircleType
		if front != nil {
			frontCircleType, _ = front.(*CircleType)
		}
		stage.OnAfterCircleTypeDeleteCallback.OnAfterDelete(stage, circletype, frontCircleType)
	}
}

func (clothoidtype *ClothoidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterClothoidTypeCreateCallback != nil {
		stage.OnAfterClothoidTypeCreateCallback.OnAfterCreate(stage, clothoidtype)
	}
}

func (clothoidtype *ClothoidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClothoidTypeUpdateCallback != nil {
		var frontClothoidType *ClothoidType
		if front != nil {
			frontClothoidType, _ = front.(*ClothoidType)
		}
		stage.OnAfterClothoidTypeUpdateCallback.OnAfterUpdate(stage, clothoidtype, frontClothoidType)
	}
}

func (clothoidtype *ClothoidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClothoidTypeDeleteCallback != nil {
		var frontClothoidType *ClothoidType
		if front != nil {
			frontClothoidType, _ = front.(*ClothoidType)
		}
		stage.OnAfterClothoidTypeDeleteCallback.OnAfterDelete(stage, clothoidtype, frontClothoidType)
	}
}

func (compositecurvetype *CompositeCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompositeCurveTypeCreateCallback != nil {
		stage.OnAfterCompositeCurveTypeCreateCallback.OnAfterCreate(stage, compositecurvetype)
	}
}

func (compositecurvetype *CompositeCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeCurveTypeUpdateCallback != nil {
		var frontCompositeCurveType *CompositeCurveType
		if front != nil {
			frontCompositeCurveType, _ = front.(*CompositeCurveType)
		}
		stage.OnAfterCompositeCurveTypeUpdateCallback.OnAfterUpdate(stage, compositecurvetype, frontCompositeCurveType)
	}
}

func (compositecurvetype *CompositeCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeCurveTypeDeleteCallback != nil {
		var frontCompositeCurveType *CompositeCurveType
		if front != nil {
			frontCompositeCurveType, _ = front.(*CompositeCurveType)
		}
		stage.OnAfterCompositeCurveTypeDeleteCallback.OnAfterDelete(stage, compositecurvetype, frontCompositeCurveType)
	}
}

func (compositesolidtype *CompositeSolidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompositeSolidTypeCreateCallback != nil {
		stage.OnAfterCompositeSolidTypeCreateCallback.OnAfterCreate(stage, compositesolidtype)
	}
}

func (compositesolidtype *CompositeSolidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeSolidTypeUpdateCallback != nil {
		var frontCompositeSolidType *CompositeSolidType
		if front != nil {
			frontCompositeSolidType, _ = front.(*CompositeSolidType)
		}
		stage.OnAfterCompositeSolidTypeUpdateCallback.OnAfterUpdate(stage, compositesolidtype, frontCompositeSolidType)
	}
}

func (compositesolidtype *CompositeSolidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeSolidTypeDeleteCallback != nil {
		var frontCompositeSolidType *CompositeSolidType
		if front != nil {
			frontCompositeSolidType, _ = front.(*CompositeSolidType)
		}
		stage.OnAfterCompositeSolidTypeDeleteCallback.OnAfterDelete(stage, compositesolidtype, frontCompositeSolidType)
	}
}

func (compositesurfacetype *CompositeSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompositeSurfaceTypeCreateCallback != nil {
		stage.OnAfterCompositeSurfaceTypeCreateCallback.OnAfterCreate(stage, compositesurfacetype)
	}
}

func (compositesurfacetype *CompositeSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeSurfaceTypeUpdateCallback != nil {
		var frontCompositeSurfaceType *CompositeSurfaceType
		if front != nil {
			frontCompositeSurfaceType, _ = front.(*CompositeSurfaceType)
		}
		stage.OnAfterCompositeSurfaceTypeUpdateCallback.OnAfterUpdate(stage, compositesurfacetype, frontCompositeSurfaceType)
	}
}

func (compositesurfacetype *CompositeSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeSurfaceTypeDeleteCallback != nil {
		var frontCompositeSurfaceType *CompositeSurfaceType
		if front != nil {
			frontCompositeSurfaceType, _ = front.(*CompositeSurfaceType)
		}
		stage.OnAfterCompositeSurfaceTypeDeleteCallback.OnAfterDelete(stage, compositesurfacetype, frontCompositeSurfaceType)
	}
}

func (compositevaluetype *CompositeValueType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompositeValueTypeCreateCallback != nil {
		stage.OnAfterCompositeValueTypeCreateCallback.OnAfterCreate(stage, compositevaluetype)
	}
}

func (compositevaluetype *CompositeValueType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeValueTypeUpdateCallback != nil {
		var frontCompositeValueType *CompositeValueType
		if front != nil {
			frontCompositeValueType, _ = front.(*CompositeValueType)
		}
		stage.OnAfterCompositeValueTypeUpdateCallback.OnAfterUpdate(stage, compositevaluetype, frontCompositeValueType)
	}
}

func (compositevaluetype *CompositeValueType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompositeValueTypeDeleteCallback != nil {
		var frontCompositeValueType *CompositeValueType
		if front != nil {
			frontCompositeValueType, _ = front.(*CompositeValueType)
		}
		stage.OnAfterCompositeValueTypeDeleteCallback.OnAfterDelete(stage, compositevaluetype, frontCompositeValueType)
	}
}

func (compoundcrspropertytype *CompoundCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompoundCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterCompoundCRSPropertyTypeCreateCallback.OnAfterCreate(stage, compoundcrspropertytype)
	}
}

func (compoundcrspropertytype *CompoundCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompoundCRSPropertyTypeUpdateCallback != nil {
		var frontCompoundCRSPropertyType *CompoundCRSPropertyType
		if front != nil {
			frontCompoundCRSPropertyType, _ = front.(*CompoundCRSPropertyType)
		}
		stage.OnAfterCompoundCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, compoundcrspropertytype, frontCompoundCRSPropertyType)
	}
}

func (compoundcrspropertytype *CompoundCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompoundCRSPropertyTypeDeleteCallback != nil {
		var frontCompoundCRSPropertyType *CompoundCRSPropertyType
		if front != nil {
			frontCompoundCRSPropertyType, _ = front.(*CompoundCRSPropertyType)
		}
		stage.OnAfterCompoundCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, compoundcrspropertytype, frontCompoundCRSPropertyType)
	}
}

func (compoundcrstype *CompoundCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCompoundCRSTypeCreateCallback != nil {
		stage.OnAfterCompoundCRSTypeCreateCallback.OnAfterCreate(stage, compoundcrstype)
	}
}

func (compoundcrstype *CompoundCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompoundCRSTypeUpdateCallback != nil {
		var frontCompoundCRSType *CompoundCRSType
		if front != nil {
			frontCompoundCRSType, _ = front.(*CompoundCRSType)
		}
		stage.OnAfterCompoundCRSTypeUpdateCallback.OnAfterUpdate(stage, compoundcrstype, frontCompoundCRSType)
	}
}

func (compoundcrstype *CompoundCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCompoundCRSTypeDeleteCallback != nil {
		var frontCompoundCRSType *CompoundCRSType
		if front != nil {
			frontCompoundCRSType, _ = front.(*CompoundCRSType)
		}
		stage.OnAfterCompoundCRSTypeDeleteCallback.OnAfterDelete(stage, compoundcrstype, frontCompoundCRSType)
	}
}

func (concatenatedoperationpropertytype *ConcatenatedOperationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConcatenatedOperationPropertyTypeCreateCallback != nil {
		stage.OnAfterConcatenatedOperationPropertyTypeCreateCallback.OnAfterCreate(stage, concatenatedoperationpropertytype)
	}
}

func (concatenatedoperationpropertytype *ConcatenatedOperationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcatenatedOperationPropertyTypeUpdateCallback != nil {
		var frontConcatenatedOperationPropertyType *ConcatenatedOperationPropertyType
		if front != nil {
			frontConcatenatedOperationPropertyType, _ = front.(*ConcatenatedOperationPropertyType)
		}
		stage.OnAfterConcatenatedOperationPropertyTypeUpdateCallback.OnAfterUpdate(stage, concatenatedoperationpropertytype, frontConcatenatedOperationPropertyType)
	}
}

func (concatenatedoperationpropertytype *ConcatenatedOperationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcatenatedOperationPropertyTypeDeleteCallback != nil {
		var frontConcatenatedOperationPropertyType *ConcatenatedOperationPropertyType
		if front != nil {
			frontConcatenatedOperationPropertyType, _ = front.(*ConcatenatedOperationPropertyType)
		}
		stage.OnAfterConcatenatedOperationPropertyTypeDeleteCallback.OnAfterDelete(stage, concatenatedoperationpropertytype, frontConcatenatedOperationPropertyType)
	}
}

func (concatenatedoperationtype *ConcatenatedOperationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConcatenatedOperationTypeCreateCallback != nil {
		stage.OnAfterConcatenatedOperationTypeCreateCallback.OnAfterCreate(stage, concatenatedoperationtype)
	}
}

func (concatenatedoperationtype *ConcatenatedOperationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcatenatedOperationTypeUpdateCallback != nil {
		var frontConcatenatedOperationType *ConcatenatedOperationType
		if front != nil {
			frontConcatenatedOperationType, _ = front.(*ConcatenatedOperationType)
		}
		stage.OnAfterConcatenatedOperationTypeUpdateCallback.OnAfterUpdate(stage, concatenatedoperationtype, frontConcatenatedOperationType)
	}
}

func (concatenatedoperationtype *ConcatenatedOperationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcatenatedOperationTypeDeleteCallback != nil {
		var frontConcatenatedOperationType *ConcatenatedOperationType
		if front != nil {
			frontConcatenatedOperationType, _ = front.(*ConcatenatedOperationType)
		}
		stage.OnAfterConcatenatedOperationTypeDeleteCallback.OnAfterDelete(stage, concatenatedoperationtype, frontConcatenatedOperationType)
	}
}

func (conetype *ConeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConeTypeCreateCallback != nil {
		stage.OnAfterConeTypeCreateCallback.OnAfterCreate(stage, conetype)
	}
}

func (conetype *ConeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConeTypeUpdateCallback != nil {
		var frontConeType *ConeType
		if front != nil {
			frontConeType, _ = front.(*ConeType)
		}
		stage.OnAfterConeTypeUpdateCallback.OnAfterUpdate(stage, conetype, frontConeType)
	}
}

func (conetype *ConeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConeTypeDeleteCallback != nil {
		var frontConeType *ConeType
		if front != nil {
			frontConeType, _ = front.(*ConeType)
		}
		stage.OnAfterConeTypeDeleteCallback.OnAfterDelete(stage, conetype, frontConeType)
	}
}

func (conventionalunittype *ConventionalUnitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConventionalUnitTypeCreateCallback != nil {
		stage.OnAfterConventionalUnitTypeCreateCallback.OnAfterCreate(stage, conventionalunittype)
	}
}

func (conventionalunittype *ConventionalUnitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConventionalUnitTypeUpdateCallback != nil {
		var frontConventionalUnitType *ConventionalUnitType
		if front != nil {
			frontConventionalUnitType, _ = front.(*ConventionalUnitType)
		}
		stage.OnAfterConventionalUnitTypeUpdateCallback.OnAfterUpdate(stage, conventionalunittype, frontConventionalUnitType)
	}
}

func (conventionalunittype *ConventionalUnitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConventionalUnitTypeDeleteCallback != nil {
		var frontConventionalUnitType *ConventionalUnitType
		if front != nil {
			frontConventionalUnitType, _ = front.(*ConventionalUnitType)
		}
		stage.OnAfterConventionalUnitTypeDeleteCallback.OnAfterDelete(stage, conventionalunittype, frontConventionalUnitType)
	}
}

func (conversionpropertytype *ConversionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConversionPropertyTypeCreateCallback != nil {
		stage.OnAfterConversionPropertyTypeCreateCallback.OnAfterCreate(stage, conversionpropertytype)
	}
}

func (conversionpropertytype *ConversionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConversionPropertyTypeUpdateCallback != nil {
		var frontConversionPropertyType *ConversionPropertyType
		if front != nil {
			frontConversionPropertyType, _ = front.(*ConversionPropertyType)
		}
		stage.OnAfterConversionPropertyTypeUpdateCallback.OnAfterUpdate(stage, conversionpropertytype, frontConversionPropertyType)
	}
}

func (conversionpropertytype *ConversionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConversionPropertyTypeDeleteCallback != nil {
		var frontConversionPropertyType *ConversionPropertyType
		if front != nil {
			frontConversionPropertyType, _ = front.(*ConversionPropertyType)
		}
		stage.OnAfterConversionPropertyTypeDeleteCallback.OnAfterDelete(stage, conversionpropertytype, frontConversionPropertyType)
	}
}

func (conversiontopreferredunittype *ConversionToPreferredUnitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConversionToPreferredUnitTypeCreateCallback != nil {
		stage.OnAfterConversionToPreferredUnitTypeCreateCallback.OnAfterCreate(stage, conversiontopreferredunittype)
	}
}

func (conversiontopreferredunittype *ConversionToPreferredUnitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConversionToPreferredUnitTypeUpdateCallback != nil {
		var frontConversionToPreferredUnitType *ConversionToPreferredUnitType
		if front != nil {
			frontConversionToPreferredUnitType, _ = front.(*ConversionToPreferredUnitType)
		}
		stage.OnAfterConversionToPreferredUnitTypeUpdateCallback.OnAfterUpdate(stage, conversiontopreferredunittype, frontConversionToPreferredUnitType)
	}
}

func (conversiontopreferredunittype *ConversionToPreferredUnitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConversionToPreferredUnitTypeDeleteCallback != nil {
		var frontConversionToPreferredUnitType *ConversionToPreferredUnitType
		if front != nil {
			frontConversionToPreferredUnitType, _ = front.(*ConversionToPreferredUnitType)
		}
		stage.OnAfterConversionToPreferredUnitTypeDeleteCallback.OnAfterDelete(stage, conversiontopreferredunittype, frontConversionToPreferredUnitType)
	}
}

func (conversiontype *ConversionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConversionTypeCreateCallback != nil {
		stage.OnAfterConversionTypeCreateCallback.OnAfterCreate(stage, conversiontype)
	}
}

func (conversiontype *ConversionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConversionTypeUpdateCallback != nil {
		var frontConversionType *ConversionType
		if front != nil {
			frontConversionType, _ = front.(*ConversionType)
		}
		stage.OnAfterConversionTypeUpdateCallback.OnAfterUpdate(stage, conversiontype, frontConversionType)
	}
}

func (conversiontype *ConversionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConversionTypeDeleteCallback != nil {
		var frontConversionType *ConversionType
		if front != nil {
			frontConversionType, _ = front.(*ConversionType)
		}
		stage.OnAfterConversionTypeDeleteCallback.OnAfterDelete(stage, conversiontype, frontConversionType)
	}
}

func (coordinateoperationaccuracy *CoordinateOperationAccuracy) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoordinateOperationAccuracyCreateCallback != nil {
		stage.OnAfterCoordinateOperationAccuracyCreateCallback.OnAfterCreate(stage, coordinateoperationaccuracy)
	}
}

func (coordinateoperationaccuracy *CoordinateOperationAccuracy) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateOperationAccuracyUpdateCallback != nil {
		var frontCoordinateOperationAccuracy *CoordinateOperationAccuracy
		if front != nil {
			frontCoordinateOperationAccuracy, _ = front.(*CoordinateOperationAccuracy)
		}
		stage.OnAfterCoordinateOperationAccuracyUpdateCallback.OnAfterUpdate(stage, coordinateoperationaccuracy, frontCoordinateOperationAccuracy)
	}
}

func (coordinateoperationaccuracy *CoordinateOperationAccuracy) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateOperationAccuracyDeleteCallback != nil {
		var frontCoordinateOperationAccuracy *CoordinateOperationAccuracy
		if front != nil {
			frontCoordinateOperationAccuracy, _ = front.(*CoordinateOperationAccuracy)
		}
		stage.OnAfterCoordinateOperationAccuracyDeleteCallback.OnAfterDelete(stage, coordinateoperationaccuracy, frontCoordinateOperationAccuracy)
	}
}

func (coordinateoperationpropertytype *CoordinateOperationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoordinateOperationPropertyTypeCreateCallback != nil {
		stage.OnAfterCoordinateOperationPropertyTypeCreateCallback.OnAfterCreate(stage, coordinateoperationpropertytype)
	}
}

func (coordinateoperationpropertytype *CoordinateOperationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateOperationPropertyTypeUpdateCallback != nil {
		var frontCoordinateOperationPropertyType *CoordinateOperationPropertyType
		if front != nil {
			frontCoordinateOperationPropertyType, _ = front.(*CoordinateOperationPropertyType)
		}
		stage.OnAfterCoordinateOperationPropertyTypeUpdateCallback.OnAfterUpdate(stage, coordinateoperationpropertytype, frontCoordinateOperationPropertyType)
	}
}

func (coordinateoperationpropertytype *CoordinateOperationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateOperationPropertyTypeDeleteCallback != nil {
		var frontCoordinateOperationPropertyType *CoordinateOperationPropertyType
		if front != nil {
			frontCoordinateOperationPropertyType, _ = front.(*CoordinateOperationPropertyType)
		}
		stage.OnAfterCoordinateOperationPropertyTypeDeleteCallback.OnAfterDelete(stage, coordinateoperationpropertytype, frontCoordinateOperationPropertyType)
	}
}

func (coordinatesystemaxispropertytype *CoordinateSystemAxisPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoordinateSystemAxisPropertyTypeCreateCallback != nil {
		stage.OnAfterCoordinateSystemAxisPropertyTypeCreateCallback.OnAfterCreate(stage, coordinatesystemaxispropertytype)
	}
}

func (coordinatesystemaxispropertytype *CoordinateSystemAxisPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateSystemAxisPropertyTypeUpdateCallback != nil {
		var frontCoordinateSystemAxisPropertyType *CoordinateSystemAxisPropertyType
		if front != nil {
			frontCoordinateSystemAxisPropertyType, _ = front.(*CoordinateSystemAxisPropertyType)
		}
		stage.OnAfterCoordinateSystemAxisPropertyTypeUpdateCallback.OnAfterUpdate(stage, coordinatesystemaxispropertytype, frontCoordinateSystemAxisPropertyType)
	}
}

func (coordinatesystemaxispropertytype *CoordinateSystemAxisPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateSystemAxisPropertyTypeDeleteCallback != nil {
		var frontCoordinateSystemAxisPropertyType *CoordinateSystemAxisPropertyType
		if front != nil {
			frontCoordinateSystemAxisPropertyType, _ = front.(*CoordinateSystemAxisPropertyType)
		}
		stage.OnAfterCoordinateSystemAxisPropertyTypeDeleteCallback.OnAfterDelete(stage, coordinatesystemaxispropertytype, frontCoordinateSystemAxisPropertyType)
	}
}

func (coordinatesystemaxistype *CoordinateSystemAxisType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoordinateSystemAxisTypeCreateCallback != nil {
		stage.OnAfterCoordinateSystemAxisTypeCreateCallback.OnAfterCreate(stage, coordinatesystemaxistype)
	}
}

func (coordinatesystemaxistype *CoordinateSystemAxisType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateSystemAxisTypeUpdateCallback != nil {
		var frontCoordinateSystemAxisType *CoordinateSystemAxisType
		if front != nil {
			frontCoordinateSystemAxisType, _ = front.(*CoordinateSystemAxisType)
		}
		stage.OnAfterCoordinateSystemAxisTypeUpdateCallback.OnAfterUpdate(stage, coordinatesystemaxistype, frontCoordinateSystemAxisType)
	}
}

func (coordinatesystemaxistype *CoordinateSystemAxisType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateSystemAxisTypeDeleteCallback != nil {
		var frontCoordinateSystemAxisType *CoordinateSystemAxisType
		if front != nil {
			frontCoordinateSystemAxisType, _ = front.(*CoordinateSystemAxisType)
		}
		stage.OnAfterCoordinateSystemAxisTypeDeleteCallback.OnAfterDelete(stage, coordinatesystemaxistype, frontCoordinateSystemAxisType)
	}
}

func (coordinatesystempropertytype *CoordinateSystemPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoordinateSystemPropertyTypeCreateCallback != nil {
		stage.OnAfterCoordinateSystemPropertyTypeCreateCallback.OnAfterCreate(stage, coordinatesystempropertytype)
	}
}

func (coordinatesystempropertytype *CoordinateSystemPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateSystemPropertyTypeUpdateCallback != nil {
		var frontCoordinateSystemPropertyType *CoordinateSystemPropertyType
		if front != nil {
			frontCoordinateSystemPropertyType, _ = front.(*CoordinateSystemPropertyType)
		}
		stage.OnAfterCoordinateSystemPropertyTypeUpdateCallback.OnAfterUpdate(stage, coordinatesystempropertytype, frontCoordinateSystemPropertyType)
	}
}

func (coordinatesystempropertytype *CoordinateSystemPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoordinateSystemPropertyTypeDeleteCallback != nil {
		var frontCoordinateSystemPropertyType *CoordinateSystemPropertyType
		if front != nil {
			frontCoordinateSystemPropertyType, _ = front.(*CoordinateSystemPropertyType)
		}
		stage.OnAfterCoordinateSystemPropertyTypeDeleteCallback.OnAfterDelete(stage, coordinatesystempropertytype, frontCoordinateSystemPropertyType)
	}
}

func (count *Count) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCountCreateCallback != nil {
		stage.OnAfterCountCreateCallback.OnAfterCreate(stage, count)
	}
}

func (count *Count) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCountUpdateCallback != nil {
		var frontCount *Count
		if front != nil {
			frontCount, _ = front.(*Count)
		}
		stage.OnAfterCountUpdateCallback.OnAfterUpdate(stage, count, frontCount)
	}
}

func (count *Count) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCountDeleteCallback != nil {
		var frontCount *Count
		if front != nil {
			frontCount, _ = front.(*Count)
		}
		stage.OnAfterCountDeleteCallback.OnAfterDelete(stage, count, frontCount)
	}
}

func (countpropertytype *CountPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCountPropertyTypeCreateCallback != nil {
		stage.OnAfterCountPropertyTypeCreateCallback.OnAfterCreate(stage, countpropertytype)
	}
}

func (countpropertytype *CountPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCountPropertyTypeUpdateCallback != nil {
		var frontCountPropertyType *CountPropertyType
		if front != nil {
			frontCountPropertyType, _ = front.(*CountPropertyType)
		}
		stage.OnAfterCountPropertyTypeUpdateCallback.OnAfterUpdate(stage, countpropertytype, frontCountPropertyType)
	}
}

func (countpropertytype *CountPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCountPropertyTypeDeleteCallback != nil {
		var frontCountPropertyType *CountPropertyType
		if front != nil {
			frontCountPropertyType, _ = front.(*CountPropertyType)
		}
		stage.OnAfterCountPropertyTypeDeleteCallback.OnAfterDelete(stage, countpropertytype, frontCountPropertyType)
	}
}

func (country_propertytype *Country_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCountry_PropertyTypeCreateCallback != nil {
		stage.OnAfterCountry_PropertyTypeCreateCallback.OnAfterCreate(stage, country_propertytype)
	}
}

func (country_propertytype *Country_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCountry_PropertyTypeUpdateCallback != nil {
		var frontCountry_PropertyType *Country_PropertyType
		if front != nil {
			frontCountry_PropertyType, _ = front.(*Country_PropertyType)
		}
		stage.OnAfterCountry_PropertyTypeUpdateCallback.OnAfterUpdate(stage, country_propertytype, frontCountry_PropertyType)
	}
}

func (country_propertytype *Country_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCountry_PropertyTypeDeleteCallback != nil {
		var frontCountry_PropertyType *Country_PropertyType
		if front != nil {
			frontCountry_PropertyType, _ = front.(*Country_PropertyType)
		}
		stage.OnAfterCountry_PropertyTypeDeleteCallback.OnAfterDelete(stage, country_propertytype, frontCountry_PropertyType)
	}
}

func (coveragefunctiontype *CoverageFunctionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCoverageFunctionTypeCreateCallback != nil {
		stage.OnAfterCoverageFunctionTypeCreateCallback.OnAfterCreate(stage, coveragefunctiontype)
	}
}

func (coveragefunctiontype *CoverageFunctionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoverageFunctionTypeUpdateCallback != nil {
		var frontCoverageFunctionType *CoverageFunctionType
		if front != nil {
			frontCoverageFunctionType, _ = front.(*CoverageFunctionType)
		}
		stage.OnAfterCoverageFunctionTypeUpdateCallback.OnAfterUpdate(stage, coveragefunctiontype, frontCoverageFunctionType)
	}
}

func (coveragefunctiontype *CoverageFunctionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCoverageFunctionTypeDeleteCallback != nil {
		var frontCoverageFunctionType *CoverageFunctionType
		if front != nil {
			frontCoverageFunctionType, _ = front.(*CoverageFunctionType)
		}
		stage.OnAfterCoverageFunctionTypeDeleteCallback.OnAfterDelete(stage, coveragefunctiontype, frontCoverageFunctionType)
	}
}

func (cubicsplinetype *CubicSplineType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCubicSplineTypeCreateCallback != nil {
		stage.OnAfterCubicSplineTypeCreateCallback.OnAfterCreate(stage, cubicsplinetype)
	}
}

func (cubicsplinetype *CubicSplineType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCubicSplineTypeUpdateCallback != nil {
		var frontCubicSplineType *CubicSplineType
		if front != nil {
			frontCubicSplineType, _ = front.(*CubicSplineType)
		}
		stage.OnAfterCubicSplineTypeUpdateCallback.OnAfterUpdate(stage, cubicsplinetype, frontCubicSplineType)
	}
}

func (cubicsplinetype *CubicSplineType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCubicSplineTypeDeleteCallback != nil {
		var frontCubicSplineType *CubicSplineType
		if front != nil {
			frontCubicSplineType, _ = front.(*CubicSplineType)
		}
		stage.OnAfterCubicSplineTypeDeleteCallback.OnAfterDelete(stage, cubicsplinetype, frontCubicSplineType)
	}
}

func (curvearraypropertytype *CurveArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterCurveArrayPropertyTypeCreateCallback.OnAfterCreate(stage, curvearraypropertytype)
	}
}

func (curvearraypropertytype *CurveArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveArrayPropertyTypeUpdateCallback != nil {
		var frontCurveArrayPropertyType *CurveArrayPropertyType
		if front != nil {
			frontCurveArrayPropertyType, _ = front.(*CurveArrayPropertyType)
		}
		stage.OnAfterCurveArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, curvearraypropertytype, frontCurveArrayPropertyType)
	}
}

func (curvearraypropertytype *CurveArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveArrayPropertyTypeDeleteCallback != nil {
		var frontCurveArrayPropertyType *CurveArrayPropertyType
		if front != nil {
			frontCurveArrayPropertyType, _ = front.(*CurveArrayPropertyType)
		}
		stage.OnAfterCurveArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, curvearraypropertytype, frontCurveArrayPropertyType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterCurvePropertyTypeCreateCallback.OnAfterCreate(stage, curvepropertytype)
	}
}

func (curvepropertytype *CurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeUpdateCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, curvepropertytype, frontCurvePropertyType)
	}
}

func (curvepropertytype *CurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurvePropertyTypeDeleteCallback != nil {
		var frontCurvePropertyType *CurvePropertyType
		if front != nil {
			frontCurvePropertyType, _ = front.(*CurvePropertyType)
		}
		stage.OnAfterCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, curvepropertytype, frontCurvePropertyType)
	}
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveSegmentArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterCurveSegmentArrayPropertyTypeCreateCallback.OnAfterCreate(stage, curvesegmentarraypropertytype)
	}
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveSegmentArrayPropertyTypeUpdateCallback != nil {
		var frontCurveSegmentArrayPropertyType *CurveSegmentArrayPropertyType
		if front != nil {
			frontCurveSegmentArrayPropertyType, _ = front.(*CurveSegmentArrayPropertyType)
		}
		stage.OnAfterCurveSegmentArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, curvesegmentarraypropertytype, frontCurveSegmentArrayPropertyType)
	}
}

func (curvesegmentarraypropertytype *CurveSegmentArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveSegmentArrayPropertyTypeDeleteCallback != nil {
		var frontCurveSegmentArrayPropertyType *CurveSegmentArrayPropertyType
		if front != nil {
			frontCurveSegmentArrayPropertyType, _ = front.(*CurveSegmentArrayPropertyType)
		}
		stage.OnAfterCurveSegmentArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, curvesegmentarraypropertytype, frontCurveSegmentArrayPropertyType)
	}
}

func (curvetype *CurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCurveTypeCreateCallback != nil {
		stage.OnAfterCurveTypeCreateCallback.OnAfterCreate(stage, curvetype)
	}
}

func (curvetype *CurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeUpdateCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeUpdateCallback.OnAfterUpdate(stage, curvetype, frontCurveType)
	}
}

func (curvetype *CurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCurveTypeDeleteCallback != nil {
		var frontCurveType *CurveType
		if front != nil {
			frontCurveType, _ = front.(*CurveType)
		}
		stage.OnAfterCurveTypeDeleteCallback.OnAfterDelete(stage, curvetype, frontCurveType)
	}
}

func (cylindertype *CylinderType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCylinderTypeCreateCallback != nil {
		stage.OnAfterCylinderTypeCreateCallback.OnAfterCreate(stage, cylindertype)
	}
}

func (cylindertype *CylinderType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylinderTypeUpdateCallback != nil {
		var frontCylinderType *CylinderType
		if front != nil {
			frontCylinderType, _ = front.(*CylinderType)
		}
		stage.OnAfterCylinderTypeUpdateCallback.OnAfterUpdate(stage, cylindertype, frontCylinderType)
	}
}

func (cylindertype *CylinderType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylinderTypeDeleteCallback != nil {
		var frontCylinderType *CylinderType
		if front != nil {
			frontCylinderType, _ = front.(*CylinderType)
		}
		stage.OnAfterCylinderTypeDeleteCallback.OnAfterDelete(stage, cylindertype, frontCylinderType)
	}
}

func (cylindricalcspropertytype *CylindricalCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCylindricalCSPropertyTypeCreateCallback != nil {
		stage.OnAfterCylindricalCSPropertyTypeCreateCallback.OnAfterCreate(stage, cylindricalcspropertytype)
	}
}

func (cylindricalcspropertytype *CylindricalCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylindricalCSPropertyTypeUpdateCallback != nil {
		var frontCylindricalCSPropertyType *CylindricalCSPropertyType
		if front != nil {
			frontCylindricalCSPropertyType, _ = front.(*CylindricalCSPropertyType)
		}
		stage.OnAfterCylindricalCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, cylindricalcspropertytype, frontCylindricalCSPropertyType)
	}
}

func (cylindricalcspropertytype *CylindricalCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylindricalCSPropertyTypeDeleteCallback != nil {
		var frontCylindricalCSPropertyType *CylindricalCSPropertyType
		if front != nil {
			frontCylindricalCSPropertyType, _ = front.(*CylindricalCSPropertyType)
		}
		stage.OnAfterCylindricalCSPropertyTypeDeleteCallback.OnAfterDelete(stage, cylindricalcspropertytype, frontCylindricalCSPropertyType)
	}
}

func (cylindricalcstype *CylindricalCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCylindricalCSTypeCreateCallback != nil {
		stage.OnAfterCylindricalCSTypeCreateCallback.OnAfterCreate(stage, cylindricalcstype)
	}
}

func (cylindricalcstype *CylindricalCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylindricalCSTypeUpdateCallback != nil {
		var frontCylindricalCSType *CylindricalCSType
		if front != nil {
			frontCylindricalCSType, _ = front.(*CylindricalCSType)
		}
		stage.OnAfterCylindricalCSTypeUpdateCallback.OnAfterUpdate(stage, cylindricalcstype, frontCylindricalCSType)
	}
}

func (cylindricalcstype *CylindricalCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCylindricalCSTypeDeleteCallback != nil {
		var frontCylindricalCSType *CylindricalCSType
		if front != nil {
			frontCylindricalCSType, _ = front.(*CylindricalCSType)
		}
		stage.OnAfterCylindricalCSTypeDeleteCallback.OnAfterDelete(stage, cylindricalcstype, frontCylindricalCSType)
	}
}

func (dmsangletype *DMSAngleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDMSAngleTypeCreateCallback != nil {
		stage.OnAfterDMSAngleTypeCreateCallback.OnAfterCreate(stage, dmsangletype)
	}
}

func (dmsangletype *DMSAngleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMSAngleTypeUpdateCallback != nil {
		var frontDMSAngleType *DMSAngleType
		if front != nil {
			frontDMSAngleType, _ = front.(*DMSAngleType)
		}
		stage.OnAfterDMSAngleTypeUpdateCallback.OnAfterUpdate(stage, dmsangletype, frontDMSAngleType)
	}
}

func (dmsangletype *DMSAngleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDMSAngleTypeDeleteCallback != nil {
		var frontDMSAngleType *DMSAngleType
		if front != nil {
			frontDMSAngleType, _ = front.(*DMSAngleType)
		}
		stage.OnAfterDMSAngleTypeDeleteCallback.OnAfterDelete(stage, dmsangletype, frontDMSAngleType)
	}
}

func (dq_absoluteexternalpositionalaccuracy_propertytype *DQ_AbsoluteExternalPositionalAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_absoluteexternalpositionalaccuracy_propertytype)
	}
}

func (dq_absoluteexternalpositionalaccuracy_propertytype *DQ_AbsoluteExternalPositionalAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_AbsoluteExternalPositionalAccuracy_PropertyType *DQ_AbsoluteExternalPositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_AbsoluteExternalPositionalAccuracy_PropertyType, _ = front.(*DQ_AbsoluteExternalPositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_absoluteexternalpositionalaccuracy_propertytype, frontDQ_AbsoluteExternalPositionalAccuracy_PropertyType)
	}
}

func (dq_absoluteexternalpositionalaccuracy_propertytype *DQ_AbsoluteExternalPositionalAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_AbsoluteExternalPositionalAccuracy_PropertyType *DQ_AbsoluteExternalPositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_AbsoluteExternalPositionalAccuracy_PropertyType, _ = front.(*DQ_AbsoluteExternalPositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_absoluteexternalpositionalaccuracy_propertytype, frontDQ_AbsoluteExternalPositionalAccuracy_PropertyType)
	}
}

func (dq_absoluteexternalpositionalaccuracy_type *DQ_AbsoluteExternalPositionalAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_TypeCreateCallback != nil {
		stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_TypeCreateCallback.OnAfterCreate(stage, dq_absoluteexternalpositionalaccuracy_type)
	}
}

func (dq_absoluteexternalpositionalaccuracy_type *DQ_AbsoluteExternalPositionalAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_TypeUpdateCallback != nil {
		var frontDQ_AbsoluteExternalPositionalAccuracy_Type *DQ_AbsoluteExternalPositionalAccuracy_Type
		if front != nil {
			frontDQ_AbsoluteExternalPositionalAccuracy_Type, _ = front.(*DQ_AbsoluteExternalPositionalAccuracy_Type)
		}
		stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, dq_absoluteexternalpositionalaccuracy_type, frontDQ_AbsoluteExternalPositionalAccuracy_Type)
	}
}

func (dq_absoluteexternalpositionalaccuracy_type *DQ_AbsoluteExternalPositionalAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_TypeDeleteCallback != nil {
		var frontDQ_AbsoluteExternalPositionalAccuracy_Type *DQ_AbsoluteExternalPositionalAccuracy_Type
		if front != nil {
			frontDQ_AbsoluteExternalPositionalAccuracy_Type, _ = front.(*DQ_AbsoluteExternalPositionalAccuracy_Type)
		}
		stage.OnAfterDQ_AbsoluteExternalPositionalAccuracy_TypeDeleteCallback.OnAfterDelete(stage, dq_absoluteexternalpositionalaccuracy_type, frontDQ_AbsoluteExternalPositionalAccuracy_Type)
	}
}

func (dq_accuracyofatimemeasurement_propertytype *DQ_AccuracyOfATimeMeasurement_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_AccuracyOfATimeMeasurement_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_AccuracyOfATimeMeasurement_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_accuracyofatimemeasurement_propertytype)
	}
}

func (dq_accuracyofatimemeasurement_propertytype *DQ_AccuracyOfATimeMeasurement_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AccuracyOfATimeMeasurement_PropertyTypeUpdateCallback != nil {
		var frontDQ_AccuracyOfATimeMeasurement_PropertyType *DQ_AccuracyOfATimeMeasurement_PropertyType
		if front != nil {
			frontDQ_AccuracyOfATimeMeasurement_PropertyType, _ = front.(*DQ_AccuracyOfATimeMeasurement_PropertyType)
		}
		stage.OnAfterDQ_AccuracyOfATimeMeasurement_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_accuracyofatimemeasurement_propertytype, frontDQ_AccuracyOfATimeMeasurement_PropertyType)
	}
}

func (dq_accuracyofatimemeasurement_propertytype *DQ_AccuracyOfATimeMeasurement_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AccuracyOfATimeMeasurement_PropertyTypeDeleteCallback != nil {
		var frontDQ_AccuracyOfATimeMeasurement_PropertyType *DQ_AccuracyOfATimeMeasurement_PropertyType
		if front != nil {
			frontDQ_AccuracyOfATimeMeasurement_PropertyType, _ = front.(*DQ_AccuracyOfATimeMeasurement_PropertyType)
		}
		stage.OnAfterDQ_AccuracyOfATimeMeasurement_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_accuracyofatimemeasurement_propertytype, frontDQ_AccuracyOfATimeMeasurement_PropertyType)
	}
}

func (dq_accuracyofatimemeasurement_type *DQ_AccuracyOfATimeMeasurement_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_AccuracyOfATimeMeasurement_TypeCreateCallback != nil {
		stage.OnAfterDQ_AccuracyOfATimeMeasurement_TypeCreateCallback.OnAfterCreate(stage, dq_accuracyofatimemeasurement_type)
	}
}

func (dq_accuracyofatimemeasurement_type *DQ_AccuracyOfATimeMeasurement_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AccuracyOfATimeMeasurement_TypeUpdateCallback != nil {
		var frontDQ_AccuracyOfATimeMeasurement_Type *DQ_AccuracyOfATimeMeasurement_Type
		if front != nil {
			frontDQ_AccuracyOfATimeMeasurement_Type, _ = front.(*DQ_AccuracyOfATimeMeasurement_Type)
		}
		stage.OnAfterDQ_AccuracyOfATimeMeasurement_TypeUpdateCallback.OnAfterUpdate(stage, dq_accuracyofatimemeasurement_type, frontDQ_AccuracyOfATimeMeasurement_Type)
	}
}

func (dq_accuracyofatimemeasurement_type *DQ_AccuracyOfATimeMeasurement_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_AccuracyOfATimeMeasurement_TypeDeleteCallback != nil {
		var frontDQ_AccuracyOfATimeMeasurement_Type *DQ_AccuracyOfATimeMeasurement_Type
		if front != nil {
			frontDQ_AccuracyOfATimeMeasurement_Type, _ = front.(*DQ_AccuracyOfATimeMeasurement_Type)
		}
		stage.OnAfterDQ_AccuracyOfATimeMeasurement_TypeDeleteCallback.OnAfterDelete(stage, dq_accuracyofatimemeasurement_type, frontDQ_AccuracyOfATimeMeasurement_Type)
	}
}

func (dq_completenesscommission_propertytype *DQ_CompletenessCommission_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_CompletenessCommission_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_CompletenessCommission_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_completenesscommission_propertytype)
	}
}

func (dq_completenesscommission_propertytype *DQ_CompletenessCommission_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessCommission_PropertyTypeUpdateCallback != nil {
		var frontDQ_CompletenessCommission_PropertyType *DQ_CompletenessCommission_PropertyType
		if front != nil {
			frontDQ_CompletenessCommission_PropertyType, _ = front.(*DQ_CompletenessCommission_PropertyType)
		}
		stage.OnAfterDQ_CompletenessCommission_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_completenesscommission_propertytype, frontDQ_CompletenessCommission_PropertyType)
	}
}

func (dq_completenesscommission_propertytype *DQ_CompletenessCommission_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessCommission_PropertyTypeDeleteCallback != nil {
		var frontDQ_CompletenessCommission_PropertyType *DQ_CompletenessCommission_PropertyType
		if front != nil {
			frontDQ_CompletenessCommission_PropertyType, _ = front.(*DQ_CompletenessCommission_PropertyType)
		}
		stage.OnAfterDQ_CompletenessCommission_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_completenesscommission_propertytype, frontDQ_CompletenessCommission_PropertyType)
	}
}

func (dq_completenesscommission_type *DQ_CompletenessCommission_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_CompletenessCommission_TypeCreateCallback != nil {
		stage.OnAfterDQ_CompletenessCommission_TypeCreateCallback.OnAfterCreate(stage, dq_completenesscommission_type)
	}
}

func (dq_completenesscommission_type *DQ_CompletenessCommission_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessCommission_TypeUpdateCallback != nil {
		var frontDQ_CompletenessCommission_Type *DQ_CompletenessCommission_Type
		if front != nil {
			frontDQ_CompletenessCommission_Type, _ = front.(*DQ_CompletenessCommission_Type)
		}
		stage.OnAfterDQ_CompletenessCommission_TypeUpdateCallback.OnAfterUpdate(stage, dq_completenesscommission_type, frontDQ_CompletenessCommission_Type)
	}
}

func (dq_completenesscommission_type *DQ_CompletenessCommission_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessCommission_TypeDeleteCallback != nil {
		var frontDQ_CompletenessCommission_Type *DQ_CompletenessCommission_Type
		if front != nil {
			frontDQ_CompletenessCommission_Type, _ = front.(*DQ_CompletenessCommission_Type)
		}
		stage.OnAfterDQ_CompletenessCommission_TypeDeleteCallback.OnAfterDelete(stage, dq_completenesscommission_type, frontDQ_CompletenessCommission_Type)
	}
}

func (dq_completenessomission_propertytype *DQ_CompletenessOmission_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_CompletenessOmission_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_CompletenessOmission_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_completenessomission_propertytype)
	}
}

func (dq_completenessomission_propertytype *DQ_CompletenessOmission_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessOmission_PropertyTypeUpdateCallback != nil {
		var frontDQ_CompletenessOmission_PropertyType *DQ_CompletenessOmission_PropertyType
		if front != nil {
			frontDQ_CompletenessOmission_PropertyType, _ = front.(*DQ_CompletenessOmission_PropertyType)
		}
		stage.OnAfterDQ_CompletenessOmission_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_completenessomission_propertytype, frontDQ_CompletenessOmission_PropertyType)
	}
}

func (dq_completenessomission_propertytype *DQ_CompletenessOmission_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessOmission_PropertyTypeDeleteCallback != nil {
		var frontDQ_CompletenessOmission_PropertyType *DQ_CompletenessOmission_PropertyType
		if front != nil {
			frontDQ_CompletenessOmission_PropertyType, _ = front.(*DQ_CompletenessOmission_PropertyType)
		}
		stage.OnAfterDQ_CompletenessOmission_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_completenessomission_propertytype, frontDQ_CompletenessOmission_PropertyType)
	}
}

func (dq_completenessomission_type *DQ_CompletenessOmission_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_CompletenessOmission_TypeCreateCallback != nil {
		stage.OnAfterDQ_CompletenessOmission_TypeCreateCallback.OnAfterCreate(stage, dq_completenessomission_type)
	}
}

func (dq_completenessomission_type *DQ_CompletenessOmission_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessOmission_TypeUpdateCallback != nil {
		var frontDQ_CompletenessOmission_Type *DQ_CompletenessOmission_Type
		if front != nil {
			frontDQ_CompletenessOmission_Type, _ = front.(*DQ_CompletenessOmission_Type)
		}
		stage.OnAfterDQ_CompletenessOmission_TypeUpdateCallback.OnAfterUpdate(stage, dq_completenessomission_type, frontDQ_CompletenessOmission_Type)
	}
}

func (dq_completenessomission_type *DQ_CompletenessOmission_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_CompletenessOmission_TypeDeleteCallback != nil {
		var frontDQ_CompletenessOmission_Type *DQ_CompletenessOmission_Type
		if front != nil {
			frontDQ_CompletenessOmission_Type, _ = front.(*DQ_CompletenessOmission_Type)
		}
		stage.OnAfterDQ_CompletenessOmission_TypeDeleteCallback.OnAfterDelete(stage, dq_completenessomission_type, frontDQ_CompletenessOmission_Type)
	}
}

func (dq_completeness_propertytype *DQ_Completeness_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_Completeness_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_Completeness_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_completeness_propertytype)
	}
}

func (dq_completeness_propertytype *DQ_Completeness_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Completeness_PropertyTypeUpdateCallback != nil {
		var frontDQ_Completeness_PropertyType *DQ_Completeness_PropertyType
		if front != nil {
			frontDQ_Completeness_PropertyType, _ = front.(*DQ_Completeness_PropertyType)
		}
		stage.OnAfterDQ_Completeness_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_completeness_propertytype, frontDQ_Completeness_PropertyType)
	}
}

func (dq_completeness_propertytype *DQ_Completeness_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Completeness_PropertyTypeDeleteCallback != nil {
		var frontDQ_Completeness_PropertyType *DQ_Completeness_PropertyType
		if front != nil {
			frontDQ_Completeness_PropertyType, _ = front.(*DQ_Completeness_PropertyType)
		}
		stage.OnAfterDQ_Completeness_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_completeness_propertytype, frontDQ_Completeness_PropertyType)
	}
}

func (dq_conceptualconsistency_propertytype *DQ_ConceptualConsistency_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_ConceptualConsistency_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_ConceptualConsistency_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_conceptualconsistency_propertytype)
	}
}

func (dq_conceptualconsistency_propertytype *DQ_ConceptualConsistency_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConceptualConsistency_PropertyTypeUpdateCallback != nil {
		var frontDQ_ConceptualConsistency_PropertyType *DQ_ConceptualConsistency_PropertyType
		if front != nil {
			frontDQ_ConceptualConsistency_PropertyType, _ = front.(*DQ_ConceptualConsistency_PropertyType)
		}
		stage.OnAfterDQ_ConceptualConsistency_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_conceptualconsistency_propertytype, frontDQ_ConceptualConsistency_PropertyType)
	}
}

func (dq_conceptualconsistency_propertytype *DQ_ConceptualConsistency_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConceptualConsistency_PropertyTypeDeleteCallback != nil {
		var frontDQ_ConceptualConsistency_PropertyType *DQ_ConceptualConsistency_PropertyType
		if front != nil {
			frontDQ_ConceptualConsistency_PropertyType, _ = front.(*DQ_ConceptualConsistency_PropertyType)
		}
		stage.OnAfterDQ_ConceptualConsistency_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_conceptualconsistency_propertytype, frontDQ_ConceptualConsistency_PropertyType)
	}
}

func (dq_conceptualconsistency_type *DQ_ConceptualConsistency_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_ConceptualConsistency_TypeCreateCallback != nil {
		stage.OnAfterDQ_ConceptualConsistency_TypeCreateCallback.OnAfterCreate(stage, dq_conceptualconsistency_type)
	}
}

func (dq_conceptualconsistency_type *DQ_ConceptualConsistency_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConceptualConsistency_TypeUpdateCallback != nil {
		var frontDQ_ConceptualConsistency_Type *DQ_ConceptualConsistency_Type
		if front != nil {
			frontDQ_ConceptualConsistency_Type, _ = front.(*DQ_ConceptualConsistency_Type)
		}
		stage.OnAfterDQ_ConceptualConsistency_TypeUpdateCallback.OnAfterUpdate(stage, dq_conceptualconsistency_type, frontDQ_ConceptualConsistency_Type)
	}
}

func (dq_conceptualconsistency_type *DQ_ConceptualConsistency_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConceptualConsistency_TypeDeleteCallback != nil {
		var frontDQ_ConceptualConsistency_Type *DQ_ConceptualConsistency_Type
		if front != nil {
			frontDQ_ConceptualConsistency_Type, _ = front.(*DQ_ConceptualConsistency_Type)
		}
		stage.OnAfterDQ_ConceptualConsistency_TypeDeleteCallback.OnAfterDelete(stage, dq_conceptualconsistency_type, frontDQ_ConceptualConsistency_Type)
	}
}

func (dq_conformanceresult_propertytype *DQ_ConformanceResult_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_ConformanceResult_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_ConformanceResult_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_conformanceresult_propertytype)
	}
}

func (dq_conformanceresult_propertytype *DQ_ConformanceResult_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConformanceResult_PropertyTypeUpdateCallback != nil {
		var frontDQ_ConformanceResult_PropertyType *DQ_ConformanceResult_PropertyType
		if front != nil {
			frontDQ_ConformanceResult_PropertyType, _ = front.(*DQ_ConformanceResult_PropertyType)
		}
		stage.OnAfterDQ_ConformanceResult_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_conformanceresult_propertytype, frontDQ_ConformanceResult_PropertyType)
	}
}

func (dq_conformanceresult_propertytype *DQ_ConformanceResult_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConformanceResult_PropertyTypeDeleteCallback != nil {
		var frontDQ_ConformanceResult_PropertyType *DQ_ConformanceResult_PropertyType
		if front != nil {
			frontDQ_ConformanceResult_PropertyType, _ = front.(*DQ_ConformanceResult_PropertyType)
		}
		stage.OnAfterDQ_ConformanceResult_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_conformanceresult_propertytype, frontDQ_ConformanceResult_PropertyType)
	}
}

func (dq_conformanceresult_type *DQ_ConformanceResult_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_ConformanceResult_TypeCreateCallback != nil {
		stage.OnAfterDQ_ConformanceResult_TypeCreateCallback.OnAfterCreate(stage, dq_conformanceresult_type)
	}
}

func (dq_conformanceresult_type *DQ_ConformanceResult_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConformanceResult_TypeUpdateCallback != nil {
		var frontDQ_ConformanceResult_Type *DQ_ConformanceResult_Type
		if front != nil {
			frontDQ_ConformanceResult_Type, _ = front.(*DQ_ConformanceResult_Type)
		}
		stage.OnAfterDQ_ConformanceResult_TypeUpdateCallback.OnAfterUpdate(stage, dq_conformanceresult_type, frontDQ_ConformanceResult_Type)
	}
}

func (dq_conformanceresult_type *DQ_ConformanceResult_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ConformanceResult_TypeDeleteCallback != nil {
		var frontDQ_ConformanceResult_Type *DQ_ConformanceResult_Type
		if front != nil {
			frontDQ_ConformanceResult_Type, _ = front.(*DQ_ConformanceResult_Type)
		}
		stage.OnAfterDQ_ConformanceResult_TypeDeleteCallback.OnAfterDelete(stage, dq_conformanceresult_type, frontDQ_ConformanceResult_Type)
	}
}

func (dq_dataquality_propertytype *DQ_DataQuality_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_DataQuality_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_DataQuality_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_dataquality_propertytype)
	}
}

func (dq_dataquality_propertytype *DQ_DataQuality_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DataQuality_PropertyTypeUpdateCallback != nil {
		var frontDQ_DataQuality_PropertyType *DQ_DataQuality_PropertyType
		if front != nil {
			frontDQ_DataQuality_PropertyType, _ = front.(*DQ_DataQuality_PropertyType)
		}
		stage.OnAfterDQ_DataQuality_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_dataquality_propertytype, frontDQ_DataQuality_PropertyType)
	}
}

func (dq_dataquality_propertytype *DQ_DataQuality_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DataQuality_PropertyTypeDeleteCallback != nil {
		var frontDQ_DataQuality_PropertyType *DQ_DataQuality_PropertyType
		if front != nil {
			frontDQ_DataQuality_PropertyType, _ = front.(*DQ_DataQuality_PropertyType)
		}
		stage.OnAfterDQ_DataQuality_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_dataquality_propertytype, frontDQ_DataQuality_PropertyType)
	}
}

func (dq_dataquality_type *DQ_DataQuality_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_DataQuality_TypeCreateCallback != nil {
		stage.OnAfterDQ_DataQuality_TypeCreateCallback.OnAfterCreate(stage, dq_dataquality_type)
	}
}

func (dq_dataquality_type *DQ_DataQuality_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DataQuality_TypeUpdateCallback != nil {
		var frontDQ_DataQuality_Type *DQ_DataQuality_Type
		if front != nil {
			frontDQ_DataQuality_Type, _ = front.(*DQ_DataQuality_Type)
		}
		stage.OnAfterDQ_DataQuality_TypeUpdateCallback.OnAfterUpdate(stage, dq_dataquality_type, frontDQ_DataQuality_Type)
	}
}

func (dq_dataquality_type *DQ_DataQuality_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DataQuality_TypeDeleteCallback != nil {
		var frontDQ_DataQuality_Type *DQ_DataQuality_Type
		if front != nil {
			frontDQ_DataQuality_Type, _ = front.(*DQ_DataQuality_Type)
		}
		stage.OnAfterDQ_DataQuality_TypeDeleteCallback.OnAfterDelete(stage, dq_dataquality_type, frontDQ_DataQuality_Type)
	}
}

func (dq_domainconsistency_propertytype *DQ_DomainConsistency_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_DomainConsistency_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_DomainConsistency_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_domainconsistency_propertytype)
	}
}

func (dq_domainconsistency_propertytype *DQ_DomainConsistency_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DomainConsistency_PropertyTypeUpdateCallback != nil {
		var frontDQ_DomainConsistency_PropertyType *DQ_DomainConsistency_PropertyType
		if front != nil {
			frontDQ_DomainConsistency_PropertyType, _ = front.(*DQ_DomainConsistency_PropertyType)
		}
		stage.OnAfterDQ_DomainConsistency_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_domainconsistency_propertytype, frontDQ_DomainConsistency_PropertyType)
	}
}

func (dq_domainconsistency_propertytype *DQ_DomainConsistency_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DomainConsistency_PropertyTypeDeleteCallback != nil {
		var frontDQ_DomainConsistency_PropertyType *DQ_DomainConsistency_PropertyType
		if front != nil {
			frontDQ_DomainConsistency_PropertyType, _ = front.(*DQ_DomainConsistency_PropertyType)
		}
		stage.OnAfterDQ_DomainConsistency_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_domainconsistency_propertytype, frontDQ_DomainConsistency_PropertyType)
	}
}

func (dq_domainconsistency_type *DQ_DomainConsistency_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_DomainConsistency_TypeCreateCallback != nil {
		stage.OnAfterDQ_DomainConsistency_TypeCreateCallback.OnAfterCreate(stage, dq_domainconsistency_type)
	}
}

func (dq_domainconsistency_type *DQ_DomainConsistency_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DomainConsistency_TypeUpdateCallback != nil {
		var frontDQ_DomainConsistency_Type *DQ_DomainConsistency_Type
		if front != nil {
			frontDQ_DomainConsistency_Type, _ = front.(*DQ_DomainConsistency_Type)
		}
		stage.OnAfterDQ_DomainConsistency_TypeUpdateCallback.OnAfterUpdate(stage, dq_domainconsistency_type, frontDQ_DomainConsistency_Type)
	}
}

func (dq_domainconsistency_type *DQ_DomainConsistency_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_DomainConsistency_TypeDeleteCallback != nil {
		var frontDQ_DomainConsistency_Type *DQ_DomainConsistency_Type
		if front != nil {
			frontDQ_DomainConsistency_Type, _ = front.(*DQ_DomainConsistency_Type)
		}
		stage.OnAfterDQ_DomainConsistency_TypeDeleteCallback.OnAfterDelete(stage, dq_domainconsistency_type, frontDQ_DomainConsistency_Type)
	}
}

func (dq_element_propertytype *DQ_Element_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_Element_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_Element_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_element_propertytype)
	}
}

func (dq_element_propertytype *DQ_Element_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Element_PropertyTypeUpdateCallback != nil {
		var frontDQ_Element_PropertyType *DQ_Element_PropertyType
		if front != nil {
			frontDQ_Element_PropertyType, _ = front.(*DQ_Element_PropertyType)
		}
		stage.OnAfterDQ_Element_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_element_propertytype, frontDQ_Element_PropertyType)
	}
}

func (dq_element_propertytype *DQ_Element_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Element_PropertyTypeDeleteCallback != nil {
		var frontDQ_Element_PropertyType *DQ_Element_PropertyType
		if front != nil {
			frontDQ_Element_PropertyType, _ = front.(*DQ_Element_PropertyType)
		}
		stage.OnAfterDQ_Element_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_element_propertytype, frontDQ_Element_PropertyType)
	}
}

func (dq_evaluationmethodtypecode_propertytype *DQ_EvaluationMethodTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_EvaluationMethodTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_EvaluationMethodTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_evaluationmethodtypecode_propertytype)
	}
}

func (dq_evaluationmethodtypecode_propertytype *DQ_EvaluationMethodTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_EvaluationMethodTypeCode_PropertyTypeUpdateCallback != nil {
		var frontDQ_EvaluationMethodTypeCode_PropertyType *DQ_EvaluationMethodTypeCode_PropertyType
		if front != nil {
			frontDQ_EvaluationMethodTypeCode_PropertyType, _ = front.(*DQ_EvaluationMethodTypeCode_PropertyType)
		}
		stage.OnAfterDQ_EvaluationMethodTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_evaluationmethodtypecode_propertytype, frontDQ_EvaluationMethodTypeCode_PropertyType)
	}
}

func (dq_evaluationmethodtypecode_propertytype *DQ_EvaluationMethodTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_EvaluationMethodTypeCode_PropertyTypeDeleteCallback != nil {
		var frontDQ_EvaluationMethodTypeCode_PropertyType *DQ_EvaluationMethodTypeCode_PropertyType
		if front != nil {
			frontDQ_EvaluationMethodTypeCode_PropertyType, _ = front.(*DQ_EvaluationMethodTypeCode_PropertyType)
		}
		stage.OnAfterDQ_EvaluationMethodTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_evaluationmethodtypecode_propertytype, frontDQ_EvaluationMethodTypeCode_PropertyType)
	}
}

func (dq_formatconsistency_propertytype *DQ_FormatConsistency_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_FormatConsistency_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_FormatConsistency_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_formatconsistency_propertytype)
	}
}

func (dq_formatconsistency_propertytype *DQ_FormatConsistency_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_FormatConsistency_PropertyTypeUpdateCallback != nil {
		var frontDQ_FormatConsistency_PropertyType *DQ_FormatConsistency_PropertyType
		if front != nil {
			frontDQ_FormatConsistency_PropertyType, _ = front.(*DQ_FormatConsistency_PropertyType)
		}
		stage.OnAfterDQ_FormatConsistency_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_formatconsistency_propertytype, frontDQ_FormatConsistency_PropertyType)
	}
}

func (dq_formatconsistency_propertytype *DQ_FormatConsistency_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_FormatConsistency_PropertyTypeDeleteCallback != nil {
		var frontDQ_FormatConsistency_PropertyType *DQ_FormatConsistency_PropertyType
		if front != nil {
			frontDQ_FormatConsistency_PropertyType, _ = front.(*DQ_FormatConsistency_PropertyType)
		}
		stage.OnAfterDQ_FormatConsistency_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_formatconsistency_propertytype, frontDQ_FormatConsistency_PropertyType)
	}
}

func (dq_formatconsistency_type *DQ_FormatConsistency_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_FormatConsistency_TypeCreateCallback != nil {
		stage.OnAfterDQ_FormatConsistency_TypeCreateCallback.OnAfterCreate(stage, dq_formatconsistency_type)
	}
}

func (dq_formatconsistency_type *DQ_FormatConsistency_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_FormatConsistency_TypeUpdateCallback != nil {
		var frontDQ_FormatConsistency_Type *DQ_FormatConsistency_Type
		if front != nil {
			frontDQ_FormatConsistency_Type, _ = front.(*DQ_FormatConsistency_Type)
		}
		stage.OnAfterDQ_FormatConsistency_TypeUpdateCallback.OnAfterUpdate(stage, dq_formatconsistency_type, frontDQ_FormatConsistency_Type)
	}
}

func (dq_formatconsistency_type *DQ_FormatConsistency_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_FormatConsistency_TypeDeleteCallback != nil {
		var frontDQ_FormatConsistency_Type *DQ_FormatConsistency_Type
		if front != nil {
			frontDQ_FormatConsistency_Type, _ = front.(*DQ_FormatConsistency_Type)
		}
		stage.OnAfterDQ_FormatConsistency_TypeDeleteCallback.OnAfterDelete(stage, dq_formatconsistency_type, frontDQ_FormatConsistency_Type)
	}
}

func (dq_griddeddatapositionalaccuracy_propertytype *DQ_GriddedDataPositionalAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_GriddedDataPositionalAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_GriddedDataPositionalAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_griddeddatapositionalaccuracy_propertytype)
	}
}

func (dq_griddeddatapositionalaccuracy_propertytype *DQ_GriddedDataPositionalAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_GriddedDataPositionalAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_GriddedDataPositionalAccuracy_PropertyType *DQ_GriddedDataPositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_GriddedDataPositionalAccuracy_PropertyType, _ = front.(*DQ_GriddedDataPositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_GriddedDataPositionalAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_griddeddatapositionalaccuracy_propertytype, frontDQ_GriddedDataPositionalAccuracy_PropertyType)
	}
}

func (dq_griddeddatapositionalaccuracy_propertytype *DQ_GriddedDataPositionalAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_GriddedDataPositionalAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_GriddedDataPositionalAccuracy_PropertyType *DQ_GriddedDataPositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_GriddedDataPositionalAccuracy_PropertyType, _ = front.(*DQ_GriddedDataPositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_GriddedDataPositionalAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_griddeddatapositionalaccuracy_propertytype, frontDQ_GriddedDataPositionalAccuracy_PropertyType)
	}
}

func (dq_griddeddatapositionalaccuracy_type *DQ_GriddedDataPositionalAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_GriddedDataPositionalAccuracy_TypeCreateCallback != nil {
		stage.OnAfterDQ_GriddedDataPositionalAccuracy_TypeCreateCallback.OnAfterCreate(stage, dq_griddeddatapositionalaccuracy_type)
	}
}

func (dq_griddeddatapositionalaccuracy_type *DQ_GriddedDataPositionalAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_GriddedDataPositionalAccuracy_TypeUpdateCallback != nil {
		var frontDQ_GriddedDataPositionalAccuracy_Type *DQ_GriddedDataPositionalAccuracy_Type
		if front != nil {
			frontDQ_GriddedDataPositionalAccuracy_Type, _ = front.(*DQ_GriddedDataPositionalAccuracy_Type)
		}
		stage.OnAfterDQ_GriddedDataPositionalAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, dq_griddeddatapositionalaccuracy_type, frontDQ_GriddedDataPositionalAccuracy_Type)
	}
}

func (dq_griddeddatapositionalaccuracy_type *DQ_GriddedDataPositionalAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_GriddedDataPositionalAccuracy_TypeDeleteCallback != nil {
		var frontDQ_GriddedDataPositionalAccuracy_Type *DQ_GriddedDataPositionalAccuracy_Type
		if front != nil {
			frontDQ_GriddedDataPositionalAccuracy_Type, _ = front.(*DQ_GriddedDataPositionalAccuracy_Type)
		}
		stage.OnAfterDQ_GriddedDataPositionalAccuracy_TypeDeleteCallback.OnAfterDelete(stage, dq_griddeddatapositionalaccuracy_type, frontDQ_GriddedDataPositionalAccuracy_Type)
	}
}

func (dq_logicalconsistency_propertytype *DQ_LogicalConsistency_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_LogicalConsistency_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_LogicalConsistency_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_logicalconsistency_propertytype)
	}
}

func (dq_logicalconsistency_propertytype *DQ_LogicalConsistency_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_LogicalConsistency_PropertyTypeUpdateCallback != nil {
		var frontDQ_LogicalConsistency_PropertyType *DQ_LogicalConsistency_PropertyType
		if front != nil {
			frontDQ_LogicalConsistency_PropertyType, _ = front.(*DQ_LogicalConsistency_PropertyType)
		}
		stage.OnAfterDQ_LogicalConsistency_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_logicalconsistency_propertytype, frontDQ_LogicalConsistency_PropertyType)
	}
}

func (dq_logicalconsistency_propertytype *DQ_LogicalConsistency_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_LogicalConsistency_PropertyTypeDeleteCallback != nil {
		var frontDQ_LogicalConsistency_PropertyType *DQ_LogicalConsistency_PropertyType
		if front != nil {
			frontDQ_LogicalConsistency_PropertyType, _ = front.(*DQ_LogicalConsistency_PropertyType)
		}
		stage.OnAfterDQ_LogicalConsistency_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_logicalconsistency_propertytype, frontDQ_LogicalConsistency_PropertyType)
	}
}

func (dq_nonquantitativeattributeaccuracy_propertytype *DQ_NonQuantitativeAttributeAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_nonquantitativeattributeaccuracy_propertytype)
	}
}

func (dq_nonquantitativeattributeaccuracy_propertytype *DQ_NonQuantitativeAttributeAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_NonQuantitativeAttributeAccuracy_PropertyType *DQ_NonQuantitativeAttributeAccuracy_PropertyType
		if front != nil {
			frontDQ_NonQuantitativeAttributeAccuracy_PropertyType, _ = front.(*DQ_NonQuantitativeAttributeAccuracy_PropertyType)
		}
		stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_nonquantitativeattributeaccuracy_propertytype, frontDQ_NonQuantitativeAttributeAccuracy_PropertyType)
	}
}

func (dq_nonquantitativeattributeaccuracy_propertytype *DQ_NonQuantitativeAttributeAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_NonQuantitativeAttributeAccuracy_PropertyType *DQ_NonQuantitativeAttributeAccuracy_PropertyType
		if front != nil {
			frontDQ_NonQuantitativeAttributeAccuracy_PropertyType, _ = front.(*DQ_NonQuantitativeAttributeAccuracy_PropertyType)
		}
		stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_nonquantitativeattributeaccuracy_propertytype, frontDQ_NonQuantitativeAttributeAccuracy_PropertyType)
	}
}

func (dq_nonquantitativeattributeaccuracy_type *DQ_NonQuantitativeAttributeAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_TypeCreateCallback != nil {
		stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_TypeCreateCallback.OnAfterCreate(stage, dq_nonquantitativeattributeaccuracy_type)
	}
}

func (dq_nonquantitativeattributeaccuracy_type *DQ_NonQuantitativeAttributeAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_TypeUpdateCallback != nil {
		var frontDQ_NonQuantitativeAttributeAccuracy_Type *DQ_NonQuantitativeAttributeAccuracy_Type
		if front != nil {
			frontDQ_NonQuantitativeAttributeAccuracy_Type, _ = front.(*DQ_NonQuantitativeAttributeAccuracy_Type)
		}
		stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, dq_nonquantitativeattributeaccuracy_type, frontDQ_NonQuantitativeAttributeAccuracy_Type)
	}
}

func (dq_nonquantitativeattributeaccuracy_type *DQ_NonQuantitativeAttributeAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_TypeDeleteCallback != nil {
		var frontDQ_NonQuantitativeAttributeAccuracy_Type *DQ_NonQuantitativeAttributeAccuracy_Type
		if front != nil {
			frontDQ_NonQuantitativeAttributeAccuracy_Type, _ = front.(*DQ_NonQuantitativeAttributeAccuracy_Type)
		}
		stage.OnAfterDQ_NonQuantitativeAttributeAccuracy_TypeDeleteCallback.OnAfterDelete(stage, dq_nonquantitativeattributeaccuracy_type, frontDQ_NonQuantitativeAttributeAccuracy_Type)
	}
}

func (dq_positionalaccuracy_propertytype *DQ_PositionalAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_PositionalAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_PositionalAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_positionalaccuracy_propertytype)
	}
}

func (dq_positionalaccuracy_propertytype *DQ_PositionalAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_PositionalAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_PositionalAccuracy_PropertyType *DQ_PositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_PositionalAccuracy_PropertyType, _ = front.(*DQ_PositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_PositionalAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_positionalaccuracy_propertytype, frontDQ_PositionalAccuracy_PropertyType)
	}
}

func (dq_positionalaccuracy_propertytype *DQ_PositionalAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_PositionalAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_PositionalAccuracy_PropertyType *DQ_PositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_PositionalAccuracy_PropertyType, _ = front.(*DQ_PositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_PositionalAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_positionalaccuracy_propertytype, frontDQ_PositionalAccuracy_PropertyType)
	}
}

func (dq_quantitativeattributeaccuracy_propertytype *DQ_QuantitativeAttributeAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_QuantitativeAttributeAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_QuantitativeAttributeAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_quantitativeattributeaccuracy_propertytype)
	}
}

func (dq_quantitativeattributeaccuracy_propertytype *DQ_QuantitativeAttributeAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeAttributeAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_QuantitativeAttributeAccuracy_PropertyType *DQ_QuantitativeAttributeAccuracy_PropertyType
		if front != nil {
			frontDQ_QuantitativeAttributeAccuracy_PropertyType, _ = front.(*DQ_QuantitativeAttributeAccuracy_PropertyType)
		}
		stage.OnAfterDQ_QuantitativeAttributeAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_quantitativeattributeaccuracy_propertytype, frontDQ_QuantitativeAttributeAccuracy_PropertyType)
	}
}

func (dq_quantitativeattributeaccuracy_propertytype *DQ_QuantitativeAttributeAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeAttributeAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_QuantitativeAttributeAccuracy_PropertyType *DQ_QuantitativeAttributeAccuracy_PropertyType
		if front != nil {
			frontDQ_QuantitativeAttributeAccuracy_PropertyType, _ = front.(*DQ_QuantitativeAttributeAccuracy_PropertyType)
		}
		stage.OnAfterDQ_QuantitativeAttributeAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_quantitativeattributeaccuracy_propertytype, frontDQ_QuantitativeAttributeAccuracy_PropertyType)
	}
}

func (dq_quantitativeattributeaccuracy_type *DQ_QuantitativeAttributeAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_QuantitativeAttributeAccuracy_TypeCreateCallback != nil {
		stage.OnAfterDQ_QuantitativeAttributeAccuracy_TypeCreateCallback.OnAfterCreate(stage, dq_quantitativeattributeaccuracy_type)
	}
}

func (dq_quantitativeattributeaccuracy_type *DQ_QuantitativeAttributeAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeAttributeAccuracy_TypeUpdateCallback != nil {
		var frontDQ_QuantitativeAttributeAccuracy_Type *DQ_QuantitativeAttributeAccuracy_Type
		if front != nil {
			frontDQ_QuantitativeAttributeAccuracy_Type, _ = front.(*DQ_QuantitativeAttributeAccuracy_Type)
		}
		stage.OnAfterDQ_QuantitativeAttributeAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, dq_quantitativeattributeaccuracy_type, frontDQ_QuantitativeAttributeAccuracy_Type)
	}
}

func (dq_quantitativeattributeaccuracy_type *DQ_QuantitativeAttributeAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeAttributeAccuracy_TypeDeleteCallback != nil {
		var frontDQ_QuantitativeAttributeAccuracy_Type *DQ_QuantitativeAttributeAccuracy_Type
		if front != nil {
			frontDQ_QuantitativeAttributeAccuracy_Type, _ = front.(*DQ_QuantitativeAttributeAccuracy_Type)
		}
		stage.OnAfterDQ_QuantitativeAttributeAccuracy_TypeDeleteCallback.OnAfterDelete(stage, dq_quantitativeattributeaccuracy_type, frontDQ_QuantitativeAttributeAccuracy_Type)
	}
}

func (dq_quantitativeresult_propertytype *DQ_QuantitativeResult_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_QuantitativeResult_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_QuantitativeResult_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_quantitativeresult_propertytype)
	}
}

func (dq_quantitativeresult_propertytype *DQ_QuantitativeResult_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeResult_PropertyTypeUpdateCallback != nil {
		var frontDQ_QuantitativeResult_PropertyType *DQ_QuantitativeResult_PropertyType
		if front != nil {
			frontDQ_QuantitativeResult_PropertyType, _ = front.(*DQ_QuantitativeResult_PropertyType)
		}
		stage.OnAfterDQ_QuantitativeResult_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_quantitativeresult_propertytype, frontDQ_QuantitativeResult_PropertyType)
	}
}

func (dq_quantitativeresult_propertytype *DQ_QuantitativeResult_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeResult_PropertyTypeDeleteCallback != nil {
		var frontDQ_QuantitativeResult_PropertyType *DQ_QuantitativeResult_PropertyType
		if front != nil {
			frontDQ_QuantitativeResult_PropertyType, _ = front.(*DQ_QuantitativeResult_PropertyType)
		}
		stage.OnAfterDQ_QuantitativeResult_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_quantitativeresult_propertytype, frontDQ_QuantitativeResult_PropertyType)
	}
}

func (dq_quantitativeresult_type *DQ_QuantitativeResult_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_QuantitativeResult_TypeCreateCallback != nil {
		stage.OnAfterDQ_QuantitativeResult_TypeCreateCallback.OnAfterCreate(stage, dq_quantitativeresult_type)
	}
}

func (dq_quantitativeresult_type *DQ_QuantitativeResult_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeResult_TypeUpdateCallback != nil {
		var frontDQ_QuantitativeResult_Type *DQ_QuantitativeResult_Type
		if front != nil {
			frontDQ_QuantitativeResult_Type, _ = front.(*DQ_QuantitativeResult_Type)
		}
		stage.OnAfterDQ_QuantitativeResult_TypeUpdateCallback.OnAfterUpdate(stage, dq_quantitativeresult_type, frontDQ_QuantitativeResult_Type)
	}
}

func (dq_quantitativeresult_type *DQ_QuantitativeResult_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_QuantitativeResult_TypeDeleteCallback != nil {
		var frontDQ_QuantitativeResult_Type *DQ_QuantitativeResult_Type
		if front != nil {
			frontDQ_QuantitativeResult_Type, _ = front.(*DQ_QuantitativeResult_Type)
		}
		stage.OnAfterDQ_QuantitativeResult_TypeDeleteCallback.OnAfterDelete(stage, dq_quantitativeresult_type, frontDQ_QuantitativeResult_Type)
	}
}

func (dq_relativeinternalpositionalaccuracy_propertytype *DQ_RelativeInternalPositionalAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_RelativeInternalPositionalAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_RelativeInternalPositionalAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_relativeinternalpositionalaccuracy_propertytype)
	}
}

func (dq_relativeinternalpositionalaccuracy_propertytype *DQ_RelativeInternalPositionalAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_RelativeInternalPositionalAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_RelativeInternalPositionalAccuracy_PropertyType *DQ_RelativeInternalPositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_RelativeInternalPositionalAccuracy_PropertyType, _ = front.(*DQ_RelativeInternalPositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_RelativeInternalPositionalAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_relativeinternalpositionalaccuracy_propertytype, frontDQ_RelativeInternalPositionalAccuracy_PropertyType)
	}
}

func (dq_relativeinternalpositionalaccuracy_propertytype *DQ_RelativeInternalPositionalAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_RelativeInternalPositionalAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_RelativeInternalPositionalAccuracy_PropertyType *DQ_RelativeInternalPositionalAccuracy_PropertyType
		if front != nil {
			frontDQ_RelativeInternalPositionalAccuracy_PropertyType, _ = front.(*DQ_RelativeInternalPositionalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_RelativeInternalPositionalAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_relativeinternalpositionalaccuracy_propertytype, frontDQ_RelativeInternalPositionalAccuracy_PropertyType)
	}
}

func (dq_relativeinternalpositionalaccuracy_type *DQ_RelativeInternalPositionalAccuracy_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_RelativeInternalPositionalAccuracy_TypeCreateCallback != nil {
		stage.OnAfterDQ_RelativeInternalPositionalAccuracy_TypeCreateCallback.OnAfterCreate(stage, dq_relativeinternalpositionalaccuracy_type)
	}
}

func (dq_relativeinternalpositionalaccuracy_type *DQ_RelativeInternalPositionalAccuracy_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_RelativeInternalPositionalAccuracy_TypeUpdateCallback != nil {
		var frontDQ_RelativeInternalPositionalAccuracy_Type *DQ_RelativeInternalPositionalAccuracy_Type
		if front != nil {
			frontDQ_RelativeInternalPositionalAccuracy_Type, _ = front.(*DQ_RelativeInternalPositionalAccuracy_Type)
		}
		stage.OnAfterDQ_RelativeInternalPositionalAccuracy_TypeUpdateCallback.OnAfterUpdate(stage, dq_relativeinternalpositionalaccuracy_type, frontDQ_RelativeInternalPositionalAccuracy_Type)
	}
}

func (dq_relativeinternalpositionalaccuracy_type *DQ_RelativeInternalPositionalAccuracy_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_RelativeInternalPositionalAccuracy_TypeDeleteCallback != nil {
		var frontDQ_RelativeInternalPositionalAccuracy_Type *DQ_RelativeInternalPositionalAccuracy_Type
		if front != nil {
			frontDQ_RelativeInternalPositionalAccuracy_Type, _ = front.(*DQ_RelativeInternalPositionalAccuracy_Type)
		}
		stage.OnAfterDQ_RelativeInternalPositionalAccuracy_TypeDeleteCallback.OnAfterDelete(stage, dq_relativeinternalpositionalaccuracy_type, frontDQ_RelativeInternalPositionalAccuracy_Type)
	}
}

func (dq_result_propertytype *DQ_Result_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_Result_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_Result_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_result_propertytype)
	}
}

func (dq_result_propertytype *DQ_Result_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Result_PropertyTypeUpdateCallback != nil {
		var frontDQ_Result_PropertyType *DQ_Result_PropertyType
		if front != nil {
			frontDQ_Result_PropertyType, _ = front.(*DQ_Result_PropertyType)
		}
		stage.OnAfterDQ_Result_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_result_propertytype, frontDQ_Result_PropertyType)
	}
}

func (dq_result_propertytype *DQ_Result_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Result_PropertyTypeDeleteCallback != nil {
		var frontDQ_Result_PropertyType *DQ_Result_PropertyType
		if front != nil {
			frontDQ_Result_PropertyType, _ = front.(*DQ_Result_PropertyType)
		}
		stage.OnAfterDQ_Result_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_result_propertytype, frontDQ_Result_PropertyType)
	}
}

func (dq_scope_propertytype *DQ_Scope_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_Scope_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_Scope_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_scope_propertytype)
	}
}

func (dq_scope_propertytype *DQ_Scope_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Scope_PropertyTypeUpdateCallback != nil {
		var frontDQ_Scope_PropertyType *DQ_Scope_PropertyType
		if front != nil {
			frontDQ_Scope_PropertyType, _ = front.(*DQ_Scope_PropertyType)
		}
		stage.OnAfterDQ_Scope_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_scope_propertytype, frontDQ_Scope_PropertyType)
	}
}

func (dq_scope_propertytype *DQ_Scope_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Scope_PropertyTypeDeleteCallback != nil {
		var frontDQ_Scope_PropertyType *DQ_Scope_PropertyType
		if front != nil {
			frontDQ_Scope_PropertyType, _ = front.(*DQ_Scope_PropertyType)
		}
		stage.OnAfterDQ_Scope_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_scope_propertytype, frontDQ_Scope_PropertyType)
	}
}

func (dq_scope_type *DQ_Scope_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_Scope_TypeCreateCallback != nil {
		stage.OnAfterDQ_Scope_TypeCreateCallback.OnAfterCreate(stage, dq_scope_type)
	}
}

func (dq_scope_type *DQ_Scope_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Scope_TypeUpdateCallback != nil {
		var frontDQ_Scope_Type *DQ_Scope_Type
		if front != nil {
			frontDQ_Scope_Type, _ = front.(*DQ_Scope_Type)
		}
		stage.OnAfterDQ_Scope_TypeUpdateCallback.OnAfterUpdate(stage, dq_scope_type, frontDQ_Scope_Type)
	}
}

func (dq_scope_type *DQ_Scope_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_Scope_TypeDeleteCallback != nil {
		var frontDQ_Scope_Type *DQ_Scope_Type
		if front != nil {
			frontDQ_Scope_Type, _ = front.(*DQ_Scope_Type)
		}
		stage.OnAfterDQ_Scope_TypeDeleteCallback.OnAfterDelete(stage, dq_scope_type, frontDQ_Scope_Type)
	}
}

func (dq_temporalaccuracy_propertytype *DQ_TemporalAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_TemporalAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_TemporalAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_temporalaccuracy_propertytype)
	}
}

func (dq_temporalaccuracy_propertytype *DQ_TemporalAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_TemporalAccuracy_PropertyType *DQ_TemporalAccuracy_PropertyType
		if front != nil {
			frontDQ_TemporalAccuracy_PropertyType, _ = front.(*DQ_TemporalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_TemporalAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_temporalaccuracy_propertytype, frontDQ_TemporalAccuracy_PropertyType)
	}
}

func (dq_temporalaccuracy_propertytype *DQ_TemporalAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_TemporalAccuracy_PropertyType *DQ_TemporalAccuracy_PropertyType
		if front != nil {
			frontDQ_TemporalAccuracy_PropertyType, _ = front.(*DQ_TemporalAccuracy_PropertyType)
		}
		stage.OnAfterDQ_TemporalAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_temporalaccuracy_propertytype, frontDQ_TemporalAccuracy_PropertyType)
	}
}

func (dq_temporalconsistency_propertytype *DQ_TemporalConsistency_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_TemporalConsistency_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_TemporalConsistency_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_temporalconsistency_propertytype)
	}
}

func (dq_temporalconsistency_propertytype *DQ_TemporalConsistency_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalConsistency_PropertyTypeUpdateCallback != nil {
		var frontDQ_TemporalConsistency_PropertyType *DQ_TemporalConsistency_PropertyType
		if front != nil {
			frontDQ_TemporalConsistency_PropertyType, _ = front.(*DQ_TemporalConsistency_PropertyType)
		}
		stage.OnAfterDQ_TemporalConsistency_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_temporalconsistency_propertytype, frontDQ_TemporalConsistency_PropertyType)
	}
}

func (dq_temporalconsistency_propertytype *DQ_TemporalConsistency_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalConsistency_PropertyTypeDeleteCallback != nil {
		var frontDQ_TemporalConsistency_PropertyType *DQ_TemporalConsistency_PropertyType
		if front != nil {
			frontDQ_TemporalConsistency_PropertyType, _ = front.(*DQ_TemporalConsistency_PropertyType)
		}
		stage.OnAfterDQ_TemporalConsistency_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_temporalconsistency_propertytype, frontDQ_TemporalConsistency_PropertyType)
	}
}

func (dq_temporalconsistency_type *DQ_TemporalConsistency_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_TemporalConsistency_TypeCreateCallback != nil {
		stage.OnAfterDQ_TemporalConsistency_TypeCreateCallback.OnAfterCreate(stage, dq_temporalconsistency_type)
	}
}

func (dq_temporalconsistency_type *DQ_TemporalConsistency_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalConsistency_TypeUpdateCallback != nil {
		var frontDQ_TemporalConsistency_Type *DQ_TemporalConsistency_Type
		if front != nil {
			frontDQ_TemporalConsistency_Type, _ = front.(*DQ_TemporalConsistency_Type)
		}
		stage.OnAfterDQ_TemporalConsistency_TypeUpdateCallback.OnAfterUpdate(stage, dq_temporalconsistency_type, frontDQ_TemporalConsistency_Type)
	}
}

func (dq_temporalconsistency_type *DQ_TemporalConsistency_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalConsistency_TypeDeleteCallback != nil {
		var frontDQ_TemporalConsistency_Type *DQ_TemporalConsistency_Type
		if front != nil {
			frontDQ_TemporalConsistency_Type, _ = front.(*DQ_TemporalConsistency_Type)
		}
		stage.OnAfterDQ_TemporalConsistency_TypeDeleteCallback.OnAfterDelete(stage, dq_temporalconsistency_type, frontDQ_TemporalConsistency_Type)
	}
}

func (dq_temporalvalidity_propertytype *DQ_TemporalValidity_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_TemporalValidity_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_TemporalValidity_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_temporalvalidity_propertytype)
	}
}

func (dq_temporalvalidity_propertytype *DQ_TemporalValidity_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalValidity_PropertyTypeUpdateCallback != nil {
		var frontDQ_TemporalValidity_PropertyType *DQ_TemporalValidity_PropertyType
		if front != nil {
			frontDQ_TemporalValidity_PropertyType, _ = front.(*DQ_TemporalValidity_PropertyType)
		}
		stage.OnAfterDQ_TemporalValidity_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_temporalvalidity_propertytype, frontDQ_TemporalValidity_PropertyType)
	}
}

func (dq_temporalvalidity_propertytype *DQ_TemporalValidity_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalValidity_PropertyTypeDeleteCallback != nil {
		var frontDQ_TemporalValidity_PropertyType *DQ_TemporalValidity_PropertyType
		if front != nil {
			frontDQ_TemporalValidity_PropertyType, _ = front.(*DQ_TemporalValidity_PropertyType)
		}
		stage.OnAfterDQ_TemporalValidity_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_temporalvalidity_propertytype, frontDQ_TemporalValidity_PropertyType)
	}
}

func (dq_temporalvalidity_type *DQ_TemporalValidity_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_TemporalValidity_TypeCreateCallback != nil {
		stage.OnAfterDQ_TemporalValidity_TypeCreateCallback.OnAfterCreate(stage, dq_temporalvalidity_type)
	}
}

func (dq_temporalvalidity_type *DQ_TemporalValidity_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalValidity_TypeUpdateCallback != nil {
		var frontDQ_TemporalValidity_Type *DQ_TemporalValidity_Type
		if front != nil {
			frontDQ_TemporalValidity_Type, _ = front.(*DQ_TemporalValidity_Type)
		}
		stage.OnAfterDQ_TemporalValidity_TypeUpdateCallback.OnAfterUpdate(stage, dq_temporalvalidity_type, frontDQ_TemporalValidity_Type)
	}
}

func (dq_temporalvalidity_type *DQ_TemporalValidity_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TemporalValidity_TypeDeleteCallback != nil {
		var frontDQ_TemporalValidity_Type *DQ_TemporalValidity_Type
		if front != nil {
			frontDQ_TemporalValidity_Type, _ = front.(*DQ_TemporalValidity_Type)
		}
		stage.OnAfterDQ_TemporalValidity_TypeDeleteCallback.OnAfterDelete(stage, dq_temporalvalidity_type, frontDQ_TemporalValidity_Type)
	}
}

func (dq_thematicaccuracy_propertytype *DQ_ThematicAccuracy_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_ThematicAccuracy_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_ThematicAccuracy_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_thematicaccuracy_propertytype)
	}
}

func (dq_thematicaccuracy_propertytype *DQ_ThematicAccuracy_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ThematicAccuracy_PropertyTypeUpdateCallback != nil {
		var frontDQ_ThematicAccuracy_PropertyType *DQ_ThematicAccuracy_PropertyType
		if front != nil {
			frontDQ_ThematicAccuracy_PropertyType, _ = front.(*DQ_ThematicAccuracy_PropertyType)
		}
		stage.OnAfterDQ_ThematicAccuracy_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_thematicaccuracy_propertytype, frontDQ_ThematicAccuracy_PropertyType)
	}
}

func (dq_thematicaccuracy_propertytype *DQ_ThematicAccuracy_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ThematicAccuracy_PropertyTypeDeleteCallback != nil {
		var frontDQ_ThematicAccuracy_PropertyType *DQ_ThematicAccuracy_PropertyType
		if front != nil {
			frontDQ_ThematicAccuracy_PropertyType, _ = front.(*DQ_ThematicAccuracy_PropertyType)
		}
		stage.OnAfterDQ_ThematicAccuracy_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_thematicaccuracy_propertytype, frontDQ_ThematicAccuracy_PropertyType)
	}
}

func (dq_thematicclassificationcorrectness_propertytype *DQ_ThematicClassificationCorrectness_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_ThematicClassificationCorrectness_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_ThematicClassificationCorrectness_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_thematicclassificationcorrectness_propertytype)
	}
}

func (dq_thematicclassificationcorrectness_propertytype *DQ_ThematicClassificationCorrectness_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ThematicClassificationCorrectness_PropertyTypeUpdateCallback != nil {
		var frontDQ_ThematicClassificationCorrectness_PropertyType *DQ_ThematicClassificationCorrectness_PropertyType
		if front != nil {
			frontDQ_ThematicClassificationCorrectness_PropertyType, _ = front.(*DQ_ThematicClassificationCorrectness_PropertyType)
		}
		stage.OnAfterDQ_ThematicClassificationCorrectness_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_thematicclassificationcorrectness_propertytype, frontDQ_ThematicClassificationCorrectness_PropertyType)
	}
}

func (dq_thematicclassificationcorrectness_propertytype *DQ_ThematicClassificationCorrectness_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ThematicClassificationCorrectness_PropertyTypeDeleteCallback != nil {
		var frontDQ_ThematicClassificationCorrectness_PropertyType *DQ_ThematicClassificationCorrectness_PropertyType
		if front != nil {
			frontDQ_ThematicClassificationCorrectness_PropertyType, _ = front.(*DQ_ThematicClassificationCorrectness_PropertyType)
		}
		stage.OnAfterDQ_ThematicClassificationCorrectness_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_thematicclassificationcorrectness_propertytype, frontDQ_ThematicClassificationCorrectness_PropertyType)
	}
}

func (dq_thematicclassificationcorrectness_type *DQ_ThematicClassificationCorrectness_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_ThematicClassificationCorrectness_TypeCreateCallback != nil {
		stage.OnAfterDQ_ThematicClassificationCorrectness_TypeCreateCallback.OnAfterCreate(stage, dq_thematicclassificationcorrectness_type)
	}
}

func (dq_thematicclassificationcorrectness_type *DQ_ThematicClassificationCorrectness_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ThematicClassificationCorrectness_TypeUpdateCallback != nil {
		var frontDQ_ThematicClassificationCorrectness_Type *DQ_ThematicClassificationCorrectness_Type
		if front != nil {
			frontDQ_ThematicClassificationCorrectness_Type, _ = front.(*DQ_ThematicClassificationCorrectness_Type)
		}
		stage.OnAfterDQ_ThematicClassificationCorrectness_TypeUpdateCallback.OnAfterUpdate(stage, dq_thematicclassificationcorrectness_type, frontDQ_ThematicClassificationCorrectness_Type)
	}
}

func (dq_thematicclassificationcorrectness_type *DQ_ThematicClassificationCorrectness_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_ThematicClassificationCorrectness_TypeDeleteCallback != nil {
		var frontDQ_ThematicClassificationCorrectness_Type *DQ_ThematicClassificationCorrectness_Type
		if front != nil {
			frontDQ_ThematicClassificationCorrectness_Type, _ = front.(*DQ_ThematicClassificationCorrectness_Type)
		}
		stage.OnAfterDQ_ThematicClassificationCorrectness_TypeDeleteCallback.OnAfterDelete(stage, dq_thematicclassificationcorrectness_type, frontDQ_ThematicClassificationCorrectness_Type)
	}
}

func (dq_topologicalconsistency_propertytype *DQ_TopologicalConsistency_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_TopologicalConsistency_PropertyTypeCreateCallback != nil {
		stage.OnAfterDQ_TopologicalConsistency_PropertyTypeCreateCallback.OnAfterCreate(stage, dq_topologicalconsistency_propertytype)
	}
}

func (dq_topologicalconsistency_propertytype *DQ_TopologicalConsistency_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TopologicalConsistency_PropertyTypeUpdateCallback != nil {
		var frontDQ_TopologicalConsistency_PropertyType *DQ_TopologicalConsistency_PropertyType
		if front != nil {
			frontDQ_TopologicalConsistency_PropertyType, _ = front.(*DQ_TopologicalConsistency_PropertyType)
		}
		stage.OnAfterDQ_TopologicalConsistency_PropertyTypeUpdateCallback.OnAfterUpdate(stage, dq_topologicalconsistency_propertytype, frontDQ_TopologicalConsistency_PropertyType)
	}
}

func (dq_topologicalconsistency_propertytype *DQ_TopologicalConsistency_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TopologicalConsistency_PropertyTypeDeleteCallback != nil {
		var frontDQ_TopologicalConsistency_PropertyType *DQ_TopologicalConsistency_PropertyType
		if front != nil {
			frontDQ_TopologicalConsistency_PropertyType, _ = front.(*DQ_TopologicalConsistency_PropertyType)
		}
		stage.OnAfterDQ_TopologicalConsistency_PropertyTypeDeleteCallback.OnAfterDelete(stage, dq_topologicalconsistency_propertytype, frontDQ_TopologicalConsistency_PropertyType)
	}
}

func (dq_topologicalconsistency_type *DQ_TopologicalConsistency_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDQ_TopologicalConsistency_TypeCreateCallback != nil {
		stage.OnAfterDQ_TopologicalConsistency_TypeCreateCallback.OnAfterCreate(stage, dq_topologicalconsistency_type)
	}
}

func (dq_topologicalconsistency_type *DQ_TopologicalConsistency_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TopologicalConsistency_TypeUpdateCallback != nil {
		var frontDQ_TopologicalConsistency_Type *DQ_TopologicalConsistency_Type
		if front != nil {
			frontDQ_TopologicalConsistency_Type, _ = front.(*DQ_TopologicalConsistency_Type)
		}
		stage.OnAfterDQ_TopologicalConsistency_TypeUpdateCallback.OnAfterUpdate(stage, dq_topologicalconsistency_type, frontDQ_TopologicalConsistency_Type)
	}
}

func (dq_topologicalconsistency_type *DQ_TopologicalConsistency_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDQ_TopologicalConsistency_TypeDeleteCallback != nil {
		var frontDQ_TopologicalConsistency_Type *DQ_TopologicalConsistency_Type
		if front != nil {
			frontDQ_TopologicalConsistency_Type, _ = front.(*DQ_TopologicalConsistency_Type)
		}
		stage.OnAfterDQ_TopologicalConsistency_TypeDeleteCallback.OnAfterDelete(stage, dq_topologicalconsistency_type, frontDQ_TopologicalConsistency_Type)
	}
}

func (ds_associationtypecode_propertytype *DS_AssociationTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_AssociationTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_AssociationTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_associationtypecode_propertytype)
	}
}

func (ds_associationtypecode_propertytype *DS_AssociationTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_AssociationTypeCode_PropertyTypeUpdateCallback != nil {
		var frontDS_AssociationTypeCode_PropertyType *DS_AssociationTypeCode_PropertyType
		if front != nil {
			frontDS_AssociationTypeCode_PropertyType, _ = front.(*DS_AssociationTypeCode_PropertyType)
		}
		stage.OnAfterDS_AssociationTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_associationtypecode_propertytype, frontDS_AssociationTypeCode_PropertyType)
	}
}

func (ds_associationtypecode_propertytype *DS_AssociationTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_AssociationTypeCode_PropertyTypeDeleteCallback != nil {
		var frontDS_AssociationTypeCode_PropertyType *DS_AssociationTypeCode_PropertyType
		if front != nil {
			frontDS_AssociationTypeCode_PropertyType, _ = front.(*DS_AssociationTypeCode_PropertyType)
		}
		stage.OnAfterDS_AssociationTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_associationtypecode_propertytype, frontDS_AssociationTypeCode_PropertyType)
	}
}

func (ds_association_propertytype *DS_Association_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Association_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_Association_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_association_propertytype)
	}
}

func (ds_association_propertytype *DS_Association_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Association_PropertyTypeUpdateCallback != nil {
		var frontDS_Association_PropertyType *DS_Association_PropertyType
		if front != nil {
			frontDS_Association_PropertyType, _ = front.(*DS_Association_PropertyType)
		}
		stage.OnAfterDS_Association_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_association_propertytype, frontDS_Association_PropertyType)
	}
}

func (ds_association_propertytype *DS_Association_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Association_PropertyTypeDeleteCallback != nil {
		var frontDS_Association_PropertyType *DS_Association_PropertyType
		if front != nil {
			frontDS_Association_PropertyType, _ = front.(*DS_Association_PropertyType)
		}
		stage.OnAfterDS_Association_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_association_propertytype, frontDS_Association_PropertyType)
	}
}

func (ds_association_type *DS_Association_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_Association_TypeCreateCallback != nil {
		stage.OnAfterDS_Association_TypeCreateCallback.OnAfterCreate(stage, ds_association_type)
	}
}

func (ds_association_type *DS_Association_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Association_TypeUpdateCallback != nil {
		var frontDS_Association_Type *DS_Association_Type
		if front != nil {
			frontDS_Association_Type, _ = front.(*DS_Association_Type)
		}
		stage.OnAfterDS_Association_TypeUpdateCallback.OnAfterUpdate(stage, ds_association_type, frontDS_Association_Type)
	}
}

func (ds_association_type *DS_Association_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_Association_TypeDeleteCallback != nil {
		var frontDS_Association_Type *DS_Association_Type
		if front != nil {
			frontDS_Association_Type, _ = front.(*DS_Association_Type)
		}
		stage.OnAfterDS_Association_TypeDeleteCallback.OnAfterDelete(stage, ds_association_type, frontDS_Association_Type)
	}
}

func (ds_initiativetypecode_propertytype *DS_InitiativeTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDS_InitiativeTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterDS_InitiativeTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, ds_initiativetypecode_propertytype)
	}
}

func (ds_initiativetypecode_propertytype *DS_InitiativeTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_InitiativeTypeCode_PropertyTypeUpdateCallback != nil {
		var frontDS_InitiativeTypeCode_PropertyType *DS_InitiativeTypeCode_PropertyType
		if front != nil {
			frontDS_InitiativeTypeCode_PropertyType, _ = front.(*DS_InitiativeTypeCode_PropertyType)
		}
		stage.OnAfterDS_InitiativeTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ds_initiativetypecode_propertytype, frontDS_InitiativeTypeCode_PropertyType)
	}
}

func (ds_initiativetypecode_propertytype *DS_InitiativeTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDS_InitiativeTypeCode_PropertyTypeDeleteCallback != nil {
		var frontDS_InitiativeTypeCode_PropertyType *DS_InitiativeTypeCode_PropertyType
		if front != nil {
			frontDS_InitiativeTypeCode_PropertyType, _ = front.(*DS_InitiativeTypeCode_PropertyType)
		}
		stage.OnAfterDS_InitiativeTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, ds_initiativetypecode_propertytype, frontDS_InitiativeTypeCode_PropertyType)
	}
}

func (datablocktype *DataBlockType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataBlockTypeCreateCallback != nil {
		stage.OnAfterDataBlockTypeCreateCallback.OnAfterCreate(stage, datablocktype)
	}
}

func (datablocktype *DataBlockType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataBlockTypeUpdateCallback != nil {
		var frontDataBlockType *DataBlockType
		if front != nil {
			frontDataBlockType, _ = front.(*DataBlockType)
		}
		stage.OnAfterDataBlockTypeUpdateCallback.OnAfterUpdate(stage, datablocktype, frontDataBlockType)
	}
}

func (datablocktype *DataBlockType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataBlockTypeDeleteCallback != nil {
		var frontDataBlockType *DataBlockType
		if front != nil {
			frontDataBlockType, _ = front.(*DataBlockType)
		}
		stage.OnAfterDataBlockTypeDeleteCallback.OnAfterDelete(stage, datablocktype, frontDataBlockType)
	}
}

func (datetime_propertytype *DateTime_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDateTime_PropertyTypeCreateCallback != nil {
		stage.OnAfterDateTime_PropertyTypeCreateCallback.OnAfterCreate(stage, datetime_propertytype)
	}
}

func (datetime_propertytype *DateTime_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateTime_PropertyTypeUpdateCallback != nil {
		var frontDateTime_PropertyType *DateTime_PropertyType
		if front != nil {
			frontDateTime_PropertyType, _ = front.(*DateTime_PropertyType)
		}
		stage.OnAfterDateTime_PropertyTypeUpdateCallback.OnAfterUpdate(stage, datetime_propertytype, frontDateTime_PropertyType)
	}
}

func (datetime_propertytype *DateTime_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDateTime_PropertyTypeDeleteCallback != nil {
		var frontDateTime_PropertyType *DateTime_PropertyType
		if front != nil {
			frontDateTime_PropertyType, _ = front.(*DateTime_PropertyType)
		}
		stage.OnAfterDateTime_PropertyTypeDeleteCallback.OnAfterDelete(stage, datetime_propertytype, frontDateTime_PropertyType)
	}
}

func (date_propertytype *Date_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDate_PropertyTypeCreateCallback != nil {
		stage.OnAfterDate_PropertyTypeCreateCallback.OnAfterCreate(stage, date_propertytype)
	}
}

func (date_propertytype *Date_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDate_PropertyTypeUpdateCallback != nil {
		var frontDate_PropertyType *Date_PropertyType
		if front != nil {
			frontDate_PropertyType, _ = front.(*Date_PropertyType)
		}
		stage.OnAfterDate_PropertyTypeUpdateCallback.OnAfterUpdate(stage, date_propertytype, frontDate_PropertyType)
	}
}

func (date_propertytype *Date_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDate_PropertyTypeDeleteCallback != nil {
		var frontDate_PropertyType *Date_PropertyType
		if front != nil {
			frontDate_PropertyType, _ = front.(*Date_PropertyType)
		}
		stage.OnAfterDate_PropertyTypeDeleteCallback.OnAfterDelete(stage, date_propertytype, frontDate_PropertyType)
	}
}

func (datumpropertytype *DatumPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDatumPropertyTypeCreateCallback != nil {
		stage.OnAfterDatumPropertyTypeCreateCallback.OnAfterCreate(stage, datumpropertytype)
	}
}

func (datumpropertytype *DatumPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDatumPropertyTypeUpdateCallback != nil {
		var frontDatumPropertyType *DatumPropertyType
		if front != nil {
			frontDatumPropertyType, _ = front.(*DatumPropertyType)
		}
		stage.OnAfterDatumPropertyTypeUpdateCallback.OnAfterUpdate(stage, datumpropertytype, frontDatumPropertyType)
	}
}

func (datumpropertytype *DatumPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDatumPropertyTypeDeleteCallback != nil {
		var frontDatumPropertyType *DatumPropertyType
		if front != nil {
			frontDatumPropertyType, _ = front.(*DatumPropertyType)
		}
		stage.OnAfterDatumPropertyTypeDeleteCallback.OnAfterDelete(stage, datumpropertytype, frontDatumPropertyType)
	}
}

func (decimal_propertytype *Decimal_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDecimal_PropertyTypeCreateCallback != nil {
		stage.OnAfterDecimal_PropertyTypeCreateCallback.OnAfterCreate(stage, decimal_propertytype)
	}
}

func (decimal_propertytype *Decimal_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDecimal_PropertyTypeUpdateCallback != nil {
		var frontDecimal_PropertyType *Decimal_PropertyType
		if front != nil {
			frontDecimal_PropertyType, _ = front.(*Decimal_PropertyType)
		}
		stage.OnAfterDecimal_PropertyTypeUpdateCallback.OnAfterUpdate(stage, decimal_propertytype, frontDecimal_PropertyType)
	}
}

func (decimal_propertytype *Decimal_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDecimal_PropertyTypeDeleteCallback != nil {
		var frontDecimal_PropertyType *Decimal_PropertyType
		if front != nil {
			frontDecimal_PropertyType, _ = front.(*Decimal_PropertyType)
		}
		stage.OnAfterDecimal_PropertyTypeDeleteCallback.OnAfterDelete(stage, decimal_propertytype, frontDecimal_PropertyType)
	}
}

func (definitionbasetype *DefinitionBaseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDefinitionBaseTypeCreateCallback != nil {
		stage.OnAfterDefinitionBaseTypeCreateCallback.OnAfterCreate(stage, definitionbasetype)
	}
}

func (definitionbasetype *DefinitionBaseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefinitionBaseTypeUpdateCallback != nil {
		var frontDefinitionBaseType *DefinitionBaseType
		if front != nil {
			frontDefinitionBaseType, _ = front.(*DefinitionBaseType)
		}
		stage.OnAfterDefinitionBaseTypeUpdateCallback.OnAfterUpdate(stage, definitionbasetype, frontDefinitionBaseType)
	}
}

func (definitionbasetype *DefinitionBaseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefinitionBaseTypeDeleteCallback != nil {
		var frontDefinitionBaseType *DefinitionBaseType
		if front != nil {
			frontDefinitionBaseType, _ = front.(*DefinitionBaseType)
		}
		stage.OnAfterDefinitionBaseTypeDeleteCallback.OnAfterDelete(stage, definitionbasetype, frontDefinitionBaseType)
	}
}

func (definitionproxytype *DefinitionProxyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDefinitionProxyTypeCreateCallback != nil {
		stage.OnAfterDefinitionProxyTypeCreateCallback.OnAfterCreate(stage, definitionproxytype)
	}
}

func (definitionproxytype *DefinitionProxyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefinitionProxyTypeUpdateCallback != nil {
		var frontDefinitionProxyType *DefinitionProxyType
		if front != nil {
			frontDefinitionProxyType, _ = front.(*DefinitionProxyType)
		}
		stage.OnAfterDefinitionProxyTypeUpdateCallback.OnAfterUpdate(stage, definitionproxytype, frontDefinitionProxyType)
	}
}

func (definitionproxytype *DefinitionProxyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefinitionProxyTypeDeleteCallback != nil {
		var frontDefinitionProxyType *DefinitionProxyType
		if front != nil {
			frontDefinitionProxyType, _ = front.(*DefinitionProxyType)
		}
		stage.OnAfterDefinitionProxyTypeDeleteCallback.OnAfterDelete(stage, definitionproxytype, frontDefinitionProxyType)
	}
}

func (definitiontype *DefinitionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDefinitionTypeCreateCallback != nil {
		stage.OnAfterDefinitionTypeCreateCallback.OnAfterCreate(stage, definitiontype)
	}
}

func (definitiontype *DefinitionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefinitionTypeUpdateCallback != nil {
		var frontDefinitionType *DefinitionType
		if front != nil {
			frontDefinitionType, _ = front.(*DefinitionType)
		}
		stage.OnAfterDefinitionTypeUpdateCallback.OnAfterUpdate(stage, definitiontype, frontDefinitionType)
	}
}

func (definitiontype *DefinitionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDefinitionTypeDeleteCallback != nil {
		var frontDefinitionType *DefinitionType
		if front != nil {
			frontDefinitionType, _ = front.(*DefinitionType)
		}
		stage.OnAfterDefinitionTypeDeleteCallback.OnAfterDelete(stage, definitiontype, frontDefinitionType)
	}
}

func (degreestype *DegreesType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDegreesTypeCreateCallback != nil {
		stage.OnAfterDegreesTypeCreateCallback.OnAfterCreate(stage, degreestype)
	}
}

func (degreestype *DegreesType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreesTypeUpdateCallback != nil {
		var frontDegreesType *DegreesType
		if front != nil {
			frontDegreesType, _ = front.(*DegreesType)
		}
		stage.OnAfterDegreesTypeUpdateCallback.OnAfterUpdate(stage, degreestype, frontDegreesType)
	}
}

func (degreestype *DegreesType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDegreesTypeDeleteCallback != nil {
		var frontDegreesType *DegreesType
		if front != nil {
			frontDegreesType, _ = front.(*DegreesType)
		}
		stage.OnAfterDegreesTypeDeleteCallback.OnAfterDelete(stage, degreestype, frontDegreesType)
	}
}

func (derivationunittermtype *DerivationUnitTermType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDerivationUnitTermTypeCreateCallback != nil {
		stage.OnAfterDerivationUnitTermTypeCreateCallback.OnAfterCreate(stage, derivationunittermtype)
	}
}

func (derivationunittermtype *DerivationUnitTermType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivationUnitTermTypeUpdateCallback != nil {
		var frontDerivationUnitTermType *DerivationUnitTermType
		if front != nil {
			frontDerivationUnitTermType, _ = front.(*DerivationUnitTermType)
		}
		stage.OnAfterDerivationUnitTermTypeUpdateCallback.OnAfterUpdate(stage, derivationunittermtype, frontDerivationUnitTermType)
	}
}

func (derivationunittermtype *DerivationUnitTermType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivationUnitTermTypeDeleteCallback != nil {
		var frontDerivationUnitTermType *DerivationUnitTermType
		if front != nil {
			frontDerivationUnitTermType, _ = front.(*DerivationUnitTermType)
		}
		stage.OnAfterDerivationUnitTermTypeDeleteCallback.OnAfterDelete(stage, derivationunittermtype, frontDerivationUnitTermType)
	}
}

func (derivedcrspropertytype *DerivedCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDerivedCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterDerivedCRSPropertyTypeCreateCallback.OnAfterCreate(stage, derivedcrspropertytype)
	}
}

func (derivedcrspropertytype *DerivedCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedCRSPropertyTypeUpdateCallback != nil {
		var frontDerivedCRSPropertyType *DerivedCRSPropertyType
		if front != nil {
			frontDerivedCRSPropertyType, _ = front.(*DerivedCRSPropertyType)
		}
		stage.OnAfterDerivedCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, derivedcrspropertytype, frontDerivedCRSPropertyType)
	}
}

func (derivedcrspropertytype *DerivedCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedCRSPropertyTypeDeleteCallback != nil {
		var frontDerivedCRSPropertyType *DerivedCRSPropertyType
		if front != nil {
			frontDerivedCRSPropertyType, _ = front.(*DerivedCRSPropertyType)
		}
		stage.OnAfterDerivedCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, derivedcrspropertytype, frontDerivedCRSPropertyType)
	}
}

func (derivedcrstype *DerivedCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDerivedCRSTypeCreateCallback != nil {
		stage.OnAfterDerivedCRSTypeCreateCallback.OnAfterCreate(stage, derivedcrstype)
	}
}

func (derivedcrstype *DerivedCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedCRSTypeUpdateCallback != nil {
		var frontDerivedCRSType *DerivedCRSType
		if front != nil {
			frontDerivedCRSType, _ = front.(*DerivedCRSType)
		}
		stage.OnAfterDerivedCRSTypeUpdateCallback.OnAfterUpdate(stage, derivedcrstype, frontDerivedCRSType)
	}
}

func (derivedcrstype *DerivedCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedCRSTypeDeleteCallback != nil {
		var frontDerivedCRSType *DerivedCRSType
		if front != nil {
			frontDerivedCRSType, _ = front.(*DerivedCRSType)
		}
		stage.OnAfterDerivedCRSTypeDeleteCallback.OnAfterDelete(stage, derivedcrstype, frontDerivedCRSType)
	}
}

func (derivedunittype *DerivedUnitType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDerivedUnitTypeCreateCallback != nil {
		stage.OnAfterDerivedUnitTypeCreateCallback.OnAfterCreate(stage, derivedunittype)
	}
}

func (derivedunittype *DerivedUnitType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedUnitTypeUpdateCallback != nil {
		var frontDerivedUnitType *DerivedUnitType
		if front != nil {
			frontDerivedUnitType, _ = front.(*DerivedUnitType)
		}
		stage.OnAfterDerivedUnitTypeUpdateCallback.OnAfterUpdate(stage, derivedunittype, frontDerivedUnitType)
	}
}

func (derivedunittype *DerivedUnitType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDerivedUnitTypeDeleteCallback != nil {
		var frontDerivedUnitType *DerivedUnitType
		if front != nil {
			frontDerivedUnitType, _ = front.(*DerivedUnitType)
		}
		stage.OnAfterDerivedUnitTypeDeleteCallback.OnAfterDelete(stage, derivedunittype, frontDerivedUnitType)
	}
}

func (dictionaryentrytype *DictionaryEntryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDictionaryEntryTypeCreateCallback != nil {
		stage.OnAfterDictionaryEntryTypeCreateCallback.OnAfterCreate(stage, dictionaryentrytype)
	}
}

func (dictionaryentrytype *DictionaryEntryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDictionaryEntryTypeUpdateCallback != nil {
		var frontDictionaryEntryType *DictionaryEntryType
		if front != nil {
			frontDictionaryEntryType, _ = front.(*DictionaryEntryType)
		}
		stage.OnAfterDictionaryEntryTypeUpdateCallback.OnAfterUpdate(stage, dictionaryentrytype, frontDictionaryEntryType)
	}
}

func (dictionaryentrytype *DictionaryEntryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDictionaryEntryTypeDeleteCallback != nil {
		var frontDictionaryEntryType *DictionaryEntryType
		if front != nil {
			frontDictionaryEntryType, _ = front.(*DictionaryEntryType)
		}
		stage.OnAfterDictionaryEntryTypeDeleteCallback.OnAfterDelete(stage, dictionaryentrytype, frontDictionaryEntryType)
	}
}

func (dictionarytype *DictionaryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDictionaryTypeCreateCallback != nil {
		stage.OnAfterDictionaryTypeCreateCallback.OnAfterCreate(stage, dictionarytype)
	}
}

func (dictionarytype *DictionaryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDictionaryTypeUpdateCallback != nil {
		var frontDictionaryType *DictionaryType
		if front != nil {
			frontDictionaryType, _ = front.(*DictionaryType)
		}
		stage.OnAfterDictionaryTypeUpdateCallback.OnAfterUpdate(stage, dictionarytype, frontDictionaryType)
	}
}

func (dictionarytype *DictionaryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDictionaryTypeDeleteCallback != nil {
		var frontDictionaryType *DictionaryType
		if front != nil {
			frontDictionaryType, _ = front.(*DictionaryType)
		}
		stage.OnAfterDictionaryTypeDeleteCallback.OnAfterDelete(stage, dictionarytype, frontDictionaryType)
	}
}

func (directpositionlisttype *DirectPositionListType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectPositionListTypeCreateCallback != nil {
		stage.OnAfterDirectPositionListTypeCreateCallback.OnAfterCreate(stage, directpositionlisttype)
	}
}

func (directpositionlisttype *DirectPositionListType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionListTypeUpdateCallback != nil {
		var frontDirectPositionListType *DirectPositionListType
		if front != nil {
			frontDirectPositionListType, _ = front.(*DirectPositionListType)
		}
		stage.OnAfterDirectPositionListTypeUpdateCallback.OnAfterUpdate(stage, directpositionlisttype, frontDirectPositionListType)
	}
}

func (directpositionlisttype *DirectPositionListType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionListTypeDeleteCallback != nil {
		var frontDirectPositionListType *DirectPositionListType
		if front != nil {
			frontDirectPositionListType, _ = front.(*DirectPositionListType)
		}
		stage.OnAfterDirectPositionListTypeDeleteCallback.OnAfterDelete(stage, directpositionlisttype, frontDirectPositionListType)
	}
}

func (directpositiontype *DirectPositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectPositionTypeCreateCallback != nil {
		stage.OnAfterDirectPositionTypeCreateCallback.OnAfterCreate(stage, directpositiontype)
	}
}

func (directpositiontype *DirectPositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionTypeUpdateCallback != nil {
		var frontDirectPositionType *DirectPositionType
		if front != nil {
			frontDirectPositionType, _ = front.(*DirectPositionType)
		}
		stage.OnAfterDirectPositionTypeUpdateCallback.OnAfterUpdate(stage, directpositiontype, frontDirectPositionType)
	}
}

func (directpositiontype *DirectPositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectPositionTypeDeleteCallback != nil {
		var frontDirectPositionType *DirectPositionType
		if front != nil {
			frontDirectPositionType, _ = front.(*DirectPositionType)
		}
		stage.OnAfterDirectPositionTypeDeleteCallback.OnAfterDelete(stage, directpositiontype, frontDirectPositionType)
	}
}

func (directededgepropertytype *DirectedEdgePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectedEdgePropertyTypeCreateCallback != nil {
		stage.OnAfterDirectedEdgePropertyTypeCreateCallback.OnAfterCreate(stage, directededgepropertytype)
	}
}

func (directededgepropertytype *DirectedEdgePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedEdgePropertyTypeUpdateCallback != nil {
		var frontDirectedEdgePropertyType *DirectedEdgePropertyType
		if front != nil {
			frontDirectedEdgePropertyType, _ = front.(*DirectedEdgePropertyType)
		}
		stage.OnAfterDirectedEdgePropertyTypeUpdateCallback.OnAfterUpdate(stage, directededgepropertytype, frontDirectedEdgePropertyType)
	}
}

func (directededgepropertytype *DirectedEdgePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedEdgePropertyTypeDeleteCallback != nil {
		var frontDirectedEdgePropertyType *DirectedEdgePropertyType
		if front != nil {
			frontDirectedEdgePropertyType, _ = front.(*DirectedEdgePropertyType)
		}
		stage.OnAfterDirectedEdgePropertyTypeDeleteCallback.OnAfterDelete(stage, directededgepropertytype, frontDirectedEdgePropertyType)
	}
}

func (directedfacepropertytype *DirectedFacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectedFacePropertyTypeCreateCallback != nil {
		stage.OnAfterDirectedFacePropertyTypeCreateCallback.OnAfterCreate(stage, directedfacepropertytype)
	}
}

func (directedfacepropertytype *DirectedFacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedFacePropertyTypeUpdateCallback != nil {
		var frontDirectedFacePropertyType *DirectedFacePropertyType
		if front != nil {
			frontDirectedFacePropertyType, _ = front.(*DirectedFacePropertyType)
		}
		stage.OnAfterDirectedFacePropertyTypeUpdateCallback.OnAfterUpdate(stage, directedfacepropertytype, frontDirectedFacePropertyType)
	}
}

func (directedfacepropertytype *DirectedFacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedFacePropertyTypeDeleteCallback != nil {
		var frontDirectedFacePropertyType *DirectedFacePropertyType
		if front != nil {
			frontDirectedFacePropertyType, _ = front.(*DirectedFacePropertyType)
		}
		stage.OnAfterDirectedFacePropertyTypeDeleteCallback.OnAfterDelete(stage, directedfacepropertytype, frontDirectedFacePropertyType)
	}
}

func (directednodepropertytype *DirectedNodePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectedNodePropertyTypeCreateCallback != nil {
		stage.OnAfterDirectedNodePropertyTypeCreateCallback.OnAfterCreate(stage, directednodepropertytype)
	}
}

func (directednodepropertytype *DirectedNodePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedNodePropertyTypeUpdateCallback != nil {
		var frontDirectedNodePropertyType *DirectedNodePropertyType
		if front != nil {
			frontDirectedNodePropertyType, _ = front.(*DirectedNodePropertyType)
		}
		stage.OnAfterDirectedNodePropertyTypeUpdateCallback.OnAfterUpdate(stage, directednodepropertytype, frontDirectedNodePropertyType)
	}
}

func (directednodepropertytype *DirectedNodePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedNodePropertyTypeDeleteCallback != nil {
		var frontDirectedNodePropertyType *DirectedNodePropertyType
		if front != nil {
			frontDirectedNodePropertyType, _ = front.(*DirectedNodePropertyType)
		}
		stage.OnAfterDirectedNodePropertyTypeDeleteCallback.OnAfterDelete(stage, directednodepropertytype, frontDirectedNodePropertyType)
	}
}

func (directedobservationatdistancetype *DirectedObservationAtDistanceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectedObservationAtDistanceTypeCreateCallback != nil {
		stage.OnAfterDirectedObservationAtDistanceTypeCreateCallback.OnAfterCreate(stage, directedobservationatdistancetype)
	}
}

func (directedobservationatdistancetype *DirectedObservationAtDistanceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedObservationAtDistanceTypeUpdateCallback != nil {
		var frontDirectedObservationAtDistanceType *DirectedObservationAtDistanceType
		if front != nil {
			frontDirectedObservationAtDistanceType, _ = front.(*DirectedObservationAtDistanceType)
		}
		stage.OnAfterDirectedObservationAtDistanceTypeUpdateCallback.OnAfterUpdate(stage, directedobservationatdistancetype, frontDirectedObservationAtDistanceType)
	}
}

func (directedobservationatdistancetype *DirectedObservationAtDistanceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedObservationAtDistanceTypeDeleteCallback != nil {
		var frontDirectedObservationAtDistanceType *DirectedObservationAtDistanceType
		if front != nil {
			frontDirectedObservationAtDistanceType, _ = front.(*DirectedObservationAtDistanceType)
		}
		stage.OnAfterDirectedObservationAtDistanceTypeDeleteCallback.OnAfterDelete(stage, directedobservationatdistancetype, frontDirectedObservationAtDistanceType)
	}
}

func (directedobservationtype *DirectedObservationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectedObservationTypeCreateCallback != nil {
		stage.OnAfterDirectedObservationTypeCreateCallback.OnAfterCreate(stage, directedobservationtype)
	}
}

func (directedobservationtype *DirectedObservationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedObservationTypeUpdateCallback != nil {
		var frontDirectedObservationType *DirectedObservationType
		if front != nil {
			frontDirectedObservationType, _ = front.(*DirectedObservationType)
		}
		stage.OnAfterDirectedObservationTypeUpdateCallback.OnAfterUpdate(stage, directedobservationtype, frontDirectedObservationType)
	}
}

func (directedobservationtype *DirectedObservationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedObservationTypeDeleteCallback != nil {
		var frontDirectedObservationType *DirectedObservationType
		if front != nil {
			frontDirectedObservationType, _ = front.(*DirectedObservationType)
		}
		stage.OnAfterDirectedObservationTypeDeleteCallback.OnAfterDelete(stage, directedobservationtype, frontDirectedObservationType)
	}
}

func (directedtoposolidpropertytype *DirectedTopoSolidPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectedTopoSolidPropertyTypeCreateCallback != nil {
		stage.OnAfterDirectedTopoSolidPropertyTypeCreateCallback.OnAfterCreate(stage, directedtoposolidpropertytype)
	}
}

func (directedtoposolidpropertytype *DirectedTopoSolidPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedTopoSolidPropertyTypeUpdateCallback != nil {
		var frontDirectedTopoSolidPropertyType *DirectedTopoSolidPropertyType
		if front != nil {
			frontDirectedTopoSolidPropertyType, _ = front.(*DirectedTopoSolidPropertyType)
		}
		stage.OnAfterDirectedTopoSolidPropertyTypeUpdateCallback.OnAfterUpdate(stage, directedtoposolidpropertytype, frontDirectedTopoSolidPropertyType)
	}
}

func (directedtoposolidpropertytype *DirectedTopoSolidPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectedTopoSolidPropertyTypeDeleteCallback != nil {
		var frontDirectedTopoSolidPropertyType *DirectedTopoSolidPropertyType
		if front != nil {
			frontDirectedTopoSolidPropertyType, _ = front.(*DirectedTopoSolidPropertyType)
		}
		stage.OnAfterDirectedTopoSolidPropertyTypeDeleteCallback.OnAfterDelete(stage, directedtoposolidpropertytype, frontDirectedTopoSolidPropertyType)
	}
}

func (directiondescriptiontype *DirectionDescriptionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionDescriptionTypeCreateCallback != nil {
		stage.OnAfterDirectionDescriptionTypeCreateCallback.OnAfterCreate(stage, directiondescriptiontype)
	}
}

func (directiondescriptiontype *DirectionDescriptionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionDescriptionTypeUpdateCallback != nil {
		var frontDirectionDescriptionType *DirectionDescriptionType
		if front != nil {
			frontDirectionDescriptionType, _ = front.(*DirectionDescriptionType)
		}
		stage.OnAfterDirectionDescriptionTypeUpdateCallback.OnAfterUpdate(stage, directiondescriptiontype, frontDirectionDescriptionType)
	}
}

func (directiondescriptiontype *DirectionDescriptionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionDescriptionTypeDeleteCallback != nil {
		var frontDirectionDescriptionType *DirectionDescriptionType
		if front != nil {
			frontDirectionDescriptionType, _ = front.(*DirectionDescriptionType)
		}
		stage.OnAfterDirectionDescriptionTypeDeleteCallback.OnAfterDelete(stage, directiondescriptiontype, frontDirectionDescriptionType)
	}
}

func (directionpropertytype *DirectionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionPropertyTypeCreateCallback != nil {
		stage.OnAfterDirectionPropertyTypeCreateCallback.OnAfterCreate(stage, directionpropertytype)
	}
}

func (directionpropertytype *DirectionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionPropertyTypeUpdateCallback != nil {
		var frontDirectionPropertyType *DirectionPropertyType
		if front != nil {
			frontDirectionPropertyType, _ = front.(*DirectionPropertyType)
		}
		stage.OnAfterDirectionPropertyTypeUpdateCallback.OnAfterUpdate(stage, directionpropertytype, frontDirectionPropertyType)
	}
}

func (directionpropertytype *DirectionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionPropertyTypeDeleteCallback != nil {
		var frontDirectionPropertyType *DirectionPropertyType
		if front != nil {
			frontDirectionPropertyType, _ = front.(*DirectionPropertyType)
		}
		stage.OnAfterDirectionPropertyTypeDeleteCallback.OnAfterDelete(stage, directionpropertytype, frontDirectionPropertyType)
	}
}

func (directionvectortype *DirectionVectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDirectionVectorTypeCreateCallback != nil {
		stage.OnAfterDirectionVectorTypeCreateCallback.OnAfterCreate(stage, directionvectortype)
	}
}

func (directionvectortype *DirectionVectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionVectorTypeUpdateCallback != nil {
		var frontDirectionVectorType *DirectionVectorType
		if front != nil {
			frontDirectionVectorType, _ = front.(*DirectionVectorType)
		}
		stage.OnAfterDirectionVectorTypeUpdateCallback.OnAfterUpdate(stage, directionvectortype, frontDirectionVectorType)
	}
}

func (directionvectortype *DirectionVectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDirectionVectorTypeDeleteCallback != nil {
		var frontDirectionVectorType *DirectionVectorType
		if front != nil {
			frontDirectionVectorType, _ = front.(*DirectionVectorType)
		}
		stage.OnAfterDirectionVectorTypeDeleteCallback.OnAfterDelete(stage, directionvectortype, frontDirectionVectorType)
	}
}

func (discretecoveragetype *DiscreteCoverageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiscreteCoverageTypeCreateCallback != nil {
		stage.OnAfterDiscreteCoverageTypeCreateCallback.OnAfterCreate(stage, discretecoveragetype)
	}
}

func (discretecoveragetype *DiscreteCoverageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiscreteCoverageTypeUpdateCallback != nil {
		var frontDiscreteCoverageType *DiscreteCoverageType
		if front != nil {
			frontDiscreteCoverageType, _ = front.(*DiscreteCoverageType)
		}
		stage.OnAfterDiscreteCoverageTypeUpdateCallback.OnAfterUpdate(stage, discretecoveragetype, frontDiscreteCoverageType)
	}
}

func (discretecoveragetype *DiscreteCoverageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiscreteCoverageTypeDeleteCallback != nil {
		var frontDiscreteCoverageType *DiscreteCoverageType
		if front != nil {
			frontDiscreteCoverageType, _ = front.(*DiscreteCoverageType)
		}
		stage.OnAfterDiscreteCoverageTypeDeleteCallback.OnAfterDelete(stage, discretecoveragetype, frontDiscreteCoverageType)
	}
}

func (distance_propertytype *Distance_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDistance_PropertyTypeCreateCallback != nil {
		stage.OnAfterDistance_PropertyTypeCreateCallback.OnAfterCreate(stage, distance_propertytype)
	}
}

func (distance_propertytype *Distance_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistance_PropertyTypeUpdateCallback != nil {
		var frontDistance_PropertyType *Distance_PropertyType
		if front != nil {
			frontDistance_PropertyType, _ = front.(*Distance_PropertyType)
		}
		stage.OnAfterDistance_PropertyTypeUpdateCallback.OnAfterUpdate(stage, distance_propertytype, frontDistance_PropertyType)
	}
}

func (distance_propertytype *Distance_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDistance_PropertyTypeDeleteCallback != nil {
		var frontDistance_PropertyType *Distance_PropertyType
		if front != nil {
			frontDistance_PropertyType, _ = front.(*Distance_PropertyType)
		}
		stage.OnAfterDistance_PropertyTypeDeleteCallback.OnAfterDelete(stage, distance_propertytype, frontDistance_PropertyType)
	}
}

func (domainofvalidity *DomainOfValidity) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDomainOfValidityCreateCallback != nil {
		stage.OnAfterDomainOfValidityCreateCallback.OnAfterCreate(stage, domainofvalidity)
	}
}

func (domainofvalidity *DomainOfValidity) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDomainOfValidityUpdateCallback != nil {
		var frontDomainOfValidity *DomainOfValidity
		if front != nil {
			frontDomainOfValidity, _ = front.(*DomainOfValidity)
		}
		stage.OnAfterDomainOfValidityUpdateCallback.OnAfterUpdate(stage, domainofvalidity, frontDomainOfValidity)
	}
}

func (domainofvalidity *DomainOfValidity) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDomainOfValidityDeleteCallback != nil {
		var frontDomainOfValidity *DomainOfValidity
		if front != nil {
			frontDomainOfValidity, _ = front.(*DomainOfValidity)
		}
		stage.OnAfterDomainOfValidityDeleteCallback.OnAfterDelete(stage, domainofvalidity, frontDomainOfValidity)
	}
}

func (domainsettype *DomainSetType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDomainSetTypeCreateCallback != nil {
		stage.OnAfterDomainSetTypeCreateCallback.OnAfterCreate(stage, domainsettype)
	}
}

func (domainsettype *DomainSetType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDomainSetTypeUpdateCallback != nil {
		var frontDomainSetType *DomainSetType
		if front != nil {
			frontDomainSetType, _ = front.(*DomainSetType)
		}
		stage.OnAfterDomainSetTypeUpdateCallback.OnAfterUpdate(stage, domainsettype, frontDomainSetType)
	}
}

func (domainsettype *DomainSetType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDomainSetTypeDeleteCallback != nil {
		var frontDomainSetType *DomainSetType
		if front != nil {
			frontDomainSetType, _ = front.(*DomainSetType)
		}
		stage.OnAfterDomainSetTypeDeleteCallback.OnAfterDelete(stage, domainsettype, frontDomainSetType)
	}
}

func (dynamicfeaturecollectiontype *DynamicFeatureCollectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDynamicFeatureCollectionTypeCreateCallback != nil {
		stage.OnAfterDynamicFeatureCollectionTypeCreateCallback.OnAfterCreate(stage, dynamicfeaturecollectiontype)
	}
}

func (dynamicfeaturecollectiontype *DynamicFeatureCollectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureCollectionTypeUpdateCallback != nil {
		var frontDynamicFeatureCollectionType *DynamicFeatureCollectionType
		if front != nil {
			frontDynamicFeatureCollectionType, _ = front.(*DynamicFeatureCollectionType)
		}
		stage.OnAfterDynamicFeatureCollectionTypeUpdateCallback.OnAfterUpdate(stage, dynamicfeaturecollectiontype, frontDynamicFeatureCollectionType)
	}
}

func (dynamicfeaturecollectiontype *DynamicFeatureCollectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureCollectionTypeDeleteCallback != nil {
		var frontDynamicFeatureCollectionType *DynamicFeatureCollectionType
		if front != nil {
			frontDynamicFeatureCollectionType, _ = front.(*DynamicFeatureCollectionType)
		}
		stage.OnAfterDynamicFeatureCollectionTypeDeleteCallback.OnAfterDelete(stage, dynamicfeaturecollectiontype, frontDynamicFeatureCollectionType)
	}
}

func (dynamicfeaturemembertype *DynamicFeatureMemberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDynamicFeatureMemberTypeCreateCallback != nil {
		stage.OnAfterDynamicFeatureMemberTypeCreateCallback.OnAfterCreate(stage, dynamicfeaturemembertype)
	}
}

func (dynamicfeaturemembertype *DynamicFeatureMemberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureMemberTypeUpdateCallback != nil {
		var frontDynamicFeatureMemberType *DynamicFeatureMemberType
		if front != nil {
			frontDynamicFeatureMemberType, _ = front.(*DynamicFeatureMemberType)
		}
		stage.OnAfterDynamicFeatureMemberTypeUpdateCallback.OnAfterUpdate(stage, dynamicfeaturemembertype, frontDynamicFeatureMemberType)
	}
}

func (dynamicfeaturemembertype *DynamicFeatureMemberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureMemberTypeDeleteCallback != nil {
		var frontDynamicFeatureMemberType *DynamicFeatureMemberType
		if front != nil {
			frontDynamicFeatureMemberType, _ = front.(*DynamicFeatureMemberType)
		}
		stage.OnAfterDynamicFeatureMemberTypeDeleteCallback.OnAfterDelete(stage, dynamicfeaturemembertype, frontDynamicFeatureMemberType)
	}
}

func (dynamicfeaturetype *DynamicFeatureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDynamicFeatureTypeCreateCallback != nil {
		stage.OnAfterDynamicFeatureTypeCreateCallback.OnAfterCreate(stage, dynamicfeaturetype)
	}
}

func (dynamicfeaturetype *DynamicFeatureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureTypeUpdateCallback != nil {
		var frontDynamicFeatureType *DynamicFeatureType
		if front != nil {
			frontDynamicFeatureType, _ = front.(*DynamicFeatureType)
		}
		stage.OnAfterDynamicFeatureTypeUpdateCallback.OnAfterUpdate(stage, dynamicfeaturetype, frontDynamicFeatureType)
	}
}

func (dynamicfeaturetype *DynamicFeatureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDynamicFeatureTypeDeleteCallback != nil {
		var frontDynamicFeatureType *DynamicFeatureType
		if front != nil {
			frontDynamicFeatureType, _ = front.(*DynamicFeatureType)
		}
		stage.OnAfterDynamicFeatureTypeDeleteCallback.OnAfterDelete(stage, dynamicfeaturetype, frontDynamicFeatureType)
	}
}

func (ex_boundingpolygon_propertytype *EX_BoundingPolygon_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_BoundingPolygon_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_BoundingPolygon_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_boundingpolygon_propertytype)
	}
}

func (ex_boundingpolygon_propertytype *EX_BoundingPolygon_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_BoundingPolygon_PropertyTypeUpdateCallback != nil {
		var frontEX_BoundingPolygon_PropertyType *EX_BoundingPolygon_PropertyType
		if front != nil {
			frontEX_BoundingPolygon_PropertyType, _ = front.(*EX_BoundingPolygon_PropertyType)
		}
		stage.OnAfterEX_BoundingPolygon_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_boundingpolygon_propertytype, frontEX_BoundingPolygon_PropertyType)
	}
}

func (ex_boundingpolygon_propertytype *EX_BoundingPolygon_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_BoundingPolygon_PropertyTypeDeleteCallback != nil {
		var frontEX_BoundingPolygon_PropertyType *EX_BoundingPolygon_PropertyType
		if front != nil {
			frontEX_BoundingPolygon_PropertyType, _ = front.(*EX_BoundingPolygon_PropertyType)
		}
		stage.OnAfterEX_BoundingPolygon_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_boundingpolygon_propertytype, frontEX_BoundingPolygon_PropertyType)
	}
}

func (ex_boundingpolygon_type *EX_BoundingPolygon_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_BoundingPolygon_TypeCreateCallback != nil {
		stage.OnAfterEX_BoundingPolygon_TypeCreateCallback.OnAfterCreate(stage, ex_boundingpolygon_type)
	}
}

func (ex_boundingpolygon_type *EX_BoundingPolygon_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_BoundingPolygon_TypeUpdateCallback != nil {
		var frontEX_BoundingPolygon_Type *EX_BoundingPolygon_Type
		if front != nil {
			frontEX_BoundingPolygon_Type, _ = front.(*EX_BoundingPolygon_Type)
		}
		stage.OnAfterEX_BoundingPolygon_TypeUpdateCallback.OnAfterUpdate(stage, ex_boundingpolygon_type, frontEX_BoundingPolygon_Type)
	}
}

func (ex_boundingpolygon_type *EX_BoundingPolygon_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_BoundingPolygon_TypeDeleteCallback != nil {
		var frontEX_BoundingPolygon_Type *EX_BoundingPolygon_Type
		if front != nil {
			frontEX_BoundingPolygon_Type, _ = front.(*EX_BoundingPolygon_Type)
		}
		stage.OnAfterEX_BoundingPolygon_TypeDeleteCallback.OnAfterDelete(stage, ex_boundingpolygon_type, frontEX_BoundingPolygon_Type)
	}
}

func (ex_extent_propertytype *EX_Extent_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_Extent_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_Extent_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_extent_propertytype)
	}
}

func (ex_extent_propertytype *EX_Extent_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_Extent_PropertyTypeUpdateCallback != nil {
		var frontEX_Extent_PropertyType *EX_Extent_PropertyType
		if front != nil {
			frontEX_Extent_PropertyType, _ = front.(*EX_Extent_PropertyType)
		}
		stage.OnAfterEX_Extent_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_extent_propertytype, frontEX_Extent_PropertyType)
	}
}

func (ex_extent_propertytype *EX_Extent_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_Extent_PropertyTypeDeleteCallback != nil {
		var frontEX_Extent_PropertyType *EX_Extent_PropertyType
		if front != nil {
			frontEX_Extent_PropertyType, _ = front.(*EX_Extent_PropertyType)
		}
		stage.OnAfterEX_Extent_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_extent_propertytype, frontEX_Extent_PropertyType)
	}
}

func (ex_extent_type *EX_Extent_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_Extent_TypeCreateCallback != nil {
		stage.OnAfterEX_Extent_TypeCreateCallback.OnAfterCreate(stage, ex_extent_type)
	}
}

func (ex_extent_type *EX_Extent_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_Extent_TypeUpdateCallback != nil {
		var frontEX_Extent_Type *EX_Extent_Type
		if front != nil {
			frontEX_Extent_Type, _ = front.(*EX_Extent_Type)
		}
		stage.OnAfterEX_Extent_TypeUpdateCallback.OnAfterUpdate(stage, ex_extent_type, frontEX_Extent_Type)
	}
}

func (ex_extent_type *EX_Extent_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_Extent_TypeDeleteCallback != nil {
		var frontEX_Extent_Type *EX_Extent_Type
		if front != nil {
			frontEX_Extent_Type, _ = front.(*EX_Extent_Type)
		}
		stage.OnAfterEX_Extent_TypeDeleteCallback.OnAfterDelete(stage, ex_extent_type, frontEX_Extent_Type)
	}
}

func (ex_geographicboundingbox_propertytype *EX_GeographicBoundingBox_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_GeographicBoundingBox_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_GeographicBoundingBox_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_geographicboundingbox_propertytype)
	}
}

func (ex_geographicboundingbox_propertytype *EX_GeographicBoundingBox_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicBoundingBox_PropertyTypeUpdateCallback != nil {
		var frontEX_GeographicBoundingBox_PropertyType *EX_GeographicBoundingBox_PropertyType
		if front != nil {
			frontEX_GeographicBoundingBox_PropertyType, _ = front.(*EX_GeographicBoundingBox_PropertyType)
		}
		stage.OnAfterEX_GeographicBoundingBox_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_geographicboundingbox_propertytype, frontEX_GeographicBoundingBox_PropertyType)
	}
}

func (ex_geographicboundingbox_propertytype *EX_GeographicBoundingBox_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicBoundingBox_PropertyTypeDeleteCallback != nil {
		var frontEX_GeographicBoundingBox_PropertyType *EX_GeographicBoundingBox_PropertyType
		if front != nil {
			frontEX_GeographicBoundingBox_PropertyType, _ = front.(*EX_GeographicBoundingBox_PropertyType)
		}
		stage.OnAfterEX_GeographicBoundingBox_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_geographicboundingbox_propertytype, frontEX_GeographicBoundingBox_PropertyType)
	}
}

func (ex_geographicboundingbox_type *EX_GeographicBoundingBox_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_GeographicBoundingBox_TypeCreateCallback != nil {
		stage.OnAfterEX_GeographicBoundingBox_TypeCreateCallback.OnAfterCreate(stage, ex_geographicboundingbox_type)
	}
}

func (ex_geographicboundingbox_type *EX_GeographicBoundingBox_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicBoundingBox_TypeUpdateCallback != nil {
		var frontEX_GeographicBoundingBox_Type *EX_GeographicBoundingBox_Type
		if front != nil {
			frontEX_GeographicBoundingBox_Type, _ = front.(*EX_GeographicBoundingBox_Type)
		}
		stage.OnAfterEX_GeographicBoundingBox_TypeUpdateCallback.OnAfterUpdate(stage, ex_geographicboundingbox_type, frontEX_GeographicBoundingBox_Type)
	}
}

func (ex_geographicboundingbox_type *EX_GeographicBoundingBox_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicBoundingBox_TypeDeleteCallback != nil {
		var frontEX_GeographicBoundingBox_Type *EX_GeographicBoundingBox_Type
		if front != nil {
			frontEX_GeographicBoundingBox_Type, _ = front.(*EX_GeographicBoundingBox_Type)
		}
		stage.OnAfterEX_GeographicBoundingBox_TypeDeleteCallback.OnAfterDelete(stage, ex_geographicboundingbox_type, frontEX_GeographicBoundingBox_Type)
	}
}

func (ex_geographicdescription_propertytype *EX_GeographicDescription_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_GeographicDescription_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_GeographicDescription_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_geographicdescription_propertytype)
	}
}

func (ex_geographicdescription_propertytype *EX_GeographicDescription_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicDescription_PropertyTypeUpdateCallback != nil {
		var frontEX_GeographicDescription_PropertyType *EX_GeographicDescription_PropertyType
		if front != nil {
			frontEX_GeographicDescription_PropertyType, _ = front.(*EX_GeographicDescription_PropertyType)
		}
		stage.OnAfterEX_GeographicDescription_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_geographicdescription_propertytype, frontEX_GeographicDescription_PropertyType)
	}
}

func (ex_geographicdescription_propertytype *EX_GeographicDescription_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicDescription_PropertyTypeDeleteCallback != nil {
		var frontEX_GeographicDescription_PropertyType *EX_GeographicDescription_PropertyType
		if front != nil {
			frontEX_GeographicDescription_PropertyType, _ = front.(*EX_GeographicDescription_PropertyType)
		}
		stage.OnAfterEX_GeographicDescription_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_geographicdescription_propertytype, frontEX_GeographicDescription_PropertyType)
	}
}

func (ex_geographicdescription_type *EX_GeographicDescription_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_GeographicDescription_TypeCreateCallback != nil {
		stage.OnAfterEX_GeographicDescription_TypeCreateCallback.OnAfterCreate(stage, ex_geographicdescription_type)
	}
}

func (ex_geographicdescription_type *EX_GeographicDescription_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicDescription_TypeUpdateCallback != nil {
		var frontEX_GeographicDescription_Type *EX_GeographicDescription_Type
		if front != nil {
			frontEX_GeographicDescription_Type, _ = front.(*EX_GeographicDescription_Type)
		}
		stage.OnAfterEX_GeographicDescription_TypeUpdateCallback.OnAfterUpdate(stage, ex_geographicdescription_type, frontEX_GeographicDescription_Type)
	}
}

func (ex_geographicdescription_type *EX_GeographicDescription_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicDescription_TypeDeleteCallback != nil {
		var frontEX_GeographicDescription_Type *EX_GeographicDescription_Type
		if front != nil {
			frontEX_GeographicDescription_Type, _ = front.(*EX_GeographicDescription_Type)
		}
		stage.OnAfterEX_GeographicDescription_TypeDeleteCallback.OnAfterDelete(stage, ex_geographicdescription_type, frontEX_GeographicDescription_Type)
	}
}

func (ex_geographicextent_propertytype *EX_GeographicExtent_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_GeographicExtent_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_GeographicExtent_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_geographicextent_propertytype)
	}
}

func (ex_geographicextent_propertytype *EX_GeographicExtent_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicExtent_PropertyTypeUpdateCallback != nil {
		var frontEX_GeographicExtent_PropertyType *EX_GeographicExtent_PropertyType
		if front != nil {
			frontEX_GeographicExtent_PropertyType, _ = front.(*EX_GeographicExtent_PropertyType)
		}
		stage.OnAfterEX_GeographicExtent_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_geographicextent_propertytype, frontEX_GeographicExtent_PropertyType)
	}
}

func (ex_geographicextent_propertytype *EX_GeographicExtent_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_GeographicExtent_PropertyTypeDeleteCallback != nil {
		var frontEX_GeographicExtent_PropertyType *EX_GeographicExtent_PropertyType
		if front != nil {
			frontEX_GeographicExtent_PropertyType, _ = front.(*EX_GeographicExtent_PropertyType)
		}
		stage.OnAfterEX_GeographicExtent_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_geographicextent_propertytype, frontEX_GeographicExtent_PropertyType)
	}
}

func (ex_spatialtemporalextent_propertytype *EX_SpatialTemporalExtent_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_SpatialTemporalExtent_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_SpatialTemporalExtent_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_spatialtemporalextent_propertytype)
	}
}

func (ex_spatialtemporalextent_propertytype *EX_SpatialTemporalExtent_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_SpatialTemporalExtent_PropertyTypeUpdateCallback != nil {
		var frontEX_SpatialTemporalExtent_PropertyType *EX_SpatialTemporalExtent_PropertyType
		if front != nil {
			frontEX_SpatialTemporalExtent_PropertyType, _ = front.(*EX_SpatialTemporalExtent_PropertyType)
		}
		stage.OnAfterEX_SpatialTemporalExtent_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_spatialtemporalextent_propertytype, frontEX_SpatialTemporalExtent_PropertyType)
	}
}

func (ex_spatialtemporalextent_propertytype *EX_SpatialTemporalExtent_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_SpatialTemporalExtent_PropertyTypeDeleteCallback != nil {
		var frontEX_SpatialTemporalExtent_PropertyType *EX_SpatialTemporalExtent_PropertyType
		if front != nil {
			frontEX_SpatialTemporalExtent_PropertyType, _ = front.(*EX_SpatialTemporalExtent_PropertyType)
		}
		stage.OnAfterEX_SpatialTemporalExtent_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_spatialtemporalextent_propertytype, frontEX_SpatialTemporalExtent_PropertyType)
	}
}

func (ex_spatialtemporalextent_type *EX_SpatialTemporalExtent_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_SpatialTemporalExtent_TypeCreateCallback != nil {
		stage.OnAfterEX_SpatialTemporalExtent_TypeCreateCallback.OnAfterCreate(stage, ex_spatialtemporalextent_type)
	}
}

func (ex_spatialtemporalextent_type *EX_SpatialTemporalExtent_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_SpatialTemporalExtent_TypeUpdateCallback != nil {
		var frontEX_SpatialTemporalExtent_Type *EX_SpatialTemporalExtent_Type
		if front != nil {
			frontEX_SpatialTemporalExtent_Type, _ = front.(*EX_SpatialTemporalExtent_Type)
		}
		stage.OnAfterEX_SpatialTemporalExtent_TypeUpdateCallback.OnAfterUpdate(stage, ex_spatialtemporalextent_type, frontEX_SpatialTemporalExtent_Type)
	}
}

func (ex_spatialtemporalextent_type *EX_SpatialTemporalExtent_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_SpatialTemporalExtent_TypeDeleteCallback != nil {
		var frontEX_SpatialTemporalExtent_Type *EX_SpatialTemporalExtent_Type
		if front != nil {
			frontEX_SpatialTemporalExtent_Type, _ = front.(*EX_SpatialTemporalExtent_Type)
		}
		stage.OnAfterEX_SpatialTemporalExtent_TypeDeleteCallback.OnAfterDelete(stage, ex_spatialtemporalextent_type, frontEX_SpatialTemporalExtent_Type)
	}
}

func (ex_temporalextent_propertytype *EX_TemporalExtent_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_TemporalExtent_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_TemporalExtent_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_temporalextent_propertytype)
	}
}

func (ex_temporalextent_propertytype *EX_TemporalExtent_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_TemporalExtent_PropertyTypeUpdateCallback != nil {
		var frontEX_TemporalExtent_PropertyType *EX_TemporalExtent_PropertyType
		if front != nil {
			frontEX_TemporalExtent_PropertyType, _ = front.(*EX_TemporalExtent_PropertyType)
		}
		stage.OnAfterEX_TemporalExtent_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_temporalextent_propertytype, frontEX_TemporalExtent_PropertyType)
	}
}

func (ex_temporalextent_propertytype *EX_TemporalExtent_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_TemporalExtent_PropertyTypeDeleteCallback != nil {
		var frontEX_TemporalExtent_PropertyType *EX_TemporalExtent_PropertyType
		if front != nil {
			frontEX_TemporalExtent_PropertyType, _ = front.(*EX_TemporalExtent_PropertyType)
		}
		stage.OnAfterEX_TemporalExtent_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_temporalextent_propertytype, frontEX_TemporalExtent_PropertyType)
	}
}

func (ex_temporalextent_type *EX_TemporalExtent_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_TemporalExtent_TypeCreateCallback != nil {
		stage.OnAfterEX_TemporalExtent_TypeCreateCallback.OnAfterCreate(stage, ex_temporalextent_type)
	}
}

func (ex_temporalextent_type *EX_TemporalExtent_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_TemporalExtent_TypeUpdateCallback != nil {
		var frontEX_TemporalExtent_Type *EX_TemporalExtent_Type
		if front != nil {
			frontEX_TemporalExtent_Type, _ = front.(*EX_TemporalExtent_Type)
		}
		stage.OnAfterEX_TemporalExtent_TypeUpdateCallback.OnAfterUpdate(stage, ex_temporalextent_type, frontEX_TemporalExtent_Type)
	}
}

func (ex_temporalextent_type *EX_TemporalExtent_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_TemporalExtent_TypeDeleteCallback != nil {
		var frontEX_TemporalExtent_Type *EX_TemporalExtent_Type
		if front != nil {
			frontEX_TemporalExtent_Type, _ = front.(*EX_TemporalExtent_Type)
		}
		stage.OnAfterEX_TemporalExtent_TypeDeleteCallback.OnAfterDelete(stage, ex_temporalextent_type, frontEX_TemporalExtent_Type)
	}
}

func (ex_verticalextent_propertytype *EX_VerticalExtent_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_VerticalExtent_PropertyTypeCreateCallback != nil {
		stage.OnAfterEX_VerticalExtent_PropertyTypeCreateCallback.OnAfterCreate(stage, ex_verticalextent_propertytype)
	}
}

func (ex_verticalextent_propertytype *EX_VerticalExtent_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_VerticalExtent_PropertyTypeUpdateCallback != nil {
		var frontEX_VerticalExtent_PropertyType *EX_VerticalExtent_PropertyType
		if front != nil {
			frontEX_VerticalExtent_PropertyType, _ = front.(*EX_VerticalExtent_PropertyType)
		}
		stage.OnAfterEX_VerticalExtent_PropertyTypeUpdateCallback.OnAfterUpdate(stage, ex_verticalextent_propertytype, frontEX_VerticalExtent_PropertyType)
	}
}

func (ex_verticalextent_propertytype *EX_VerticalExtent_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_VerticalExtent_PropertyTypeDeleteCallback != nil {
		var frontEX_VerticalExtent_PropertyType *EX_VerticalExtent_PropertyType
		if front != nil {
			frontEX_VerticalExtent_PropertyType, _ = front.(*EX_VerticalExtent_PropertyType)
		}
		stage.OnAfterEX_VerticalExtent_PropertyTypeDeleteCallback.OnAfterDelete(stage, ex_verticalextent_propertytype, frontEX_VerticalExtent_PropertyType)
	}
}

func (ex_verticalextent_type *EX_VerticalExtent_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEX_VerticalExtent_TypeCreateCallback != nil {
		stage.OnAfterEX_VerticalExtent_TypeCreateCallback.OnAfterCreate(stage, ex_verticalextent_type)
	}
}

func (ex_verticalextent_type *EX_VerticalExtent_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_VerticalExtent_TypeUpdateCallback != nil {
		var frontEX_VerticalExtent_Type *EX_VerticalExtent_Type
		if front != nil {
			frontEX_VerticalExtent_Type, _ = front.(*EX_VerticalExtent_Type)
		}
		stage.OnAfterEX_VerticalExtent_TypeUpdateCallback.OnAfterUpdate(stage, ex_verticalextent_type, frontEX_VerticalExtent_Type)
	}
}

func (ex_verticalextent_type *EX_VerticalExtent_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEX_VerticalExtent_TypeDeleteCallback != nil {
		var frontEX_VerticalExtent_Type *EX_VerticalExtent_Type
		if front != nil {
			frontEX_VerticalExtent_Type, _ = front.(*EX_VerticalExtent_Type)
		}
		stage.OnAfterEX_VerticalExtent_TypeDeleteCallback.OnAfterDelete(stage, ex_verticalextent_type, frontEX_VerticalExtent_Type)
	}
}

func (edgetype *EdgeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEdgeTypeCreateCallback != nil {
		stage.OnAfterEdgeTypeCreateCallback.OnAfterCreate(stage, edgetype)
	}
}

func (edgetype *EdgeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEdgeTypeUpdateCallback != nil {
		var frontEdgeType *EdgeType
		if front != nil {
			frontEdgeType, _ = front.(*EdgeType)
		}
		stage.OnAfterEdgeTypeUpdateCallback.OnAfterUpdate(stage, edgetype, frontEdgeType)
	}
}

func (edgetype *EdgeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEdgeTypeDeleteCallback != nil {
		var frontEdgeType *EdgeType
		if front != nil {
			frontEdgeType, _ = front.(*EdgeType)
		}
		stage.OnAfterEdgeTypeDeleteCallback.OnAfterDelete(stage, edgetype, frontEdgeType)
	}
}

func (ellipsoidpropertytype *EllipsoidPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEllipsoidPropertyTypeCreateCallback != nil {
		stage.OnAfterEllipsoidPropertyTypeCreateCallback.OnAfterCreate(stage, ellipsoidpropertytype)
	}
}

func (ellipsoidpropertytype *EllipsoidPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidPropertyTypeUpdateCallback != nil {
		var frontEllipsoidPropertyType *EllipsoidPropertyType
		if front != nil {
			frontEllipsoidPropertyType, _ = front.(*EllipsoidPropertyType)
		}
		stage.OnAfterEllipsoidPropertyTypeUpdateCallback.OnAfterUpdate(stage, ellipsoidpropertytype, frontEllipsoidPropertyType)
	}
}

func (ellipsoidpropertytype *EllipsoidPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidPropertyTypeDeleteCallback != nil {
		var frontEllipsoidPropertyType *EllipsoidPropertyType
		if front != nil {
			frontEllipsoidPropertyType, _ = front.(*EllipsoidPropertyType)
		}
		stage.OnAfterEllipsoidPropertyTypeDeleteCallback.OnAfterDelete(stage, ellipsoidpropertytype, frontEllipsoidPropertyType)
	}
}

func (ellipsoidtype *EllipsoidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEllipsoidTypeCreateCallback != nil {
		stage.OnAfterEllipsoidTypeCreateCallback.OnAfterCreate(stage, ellipsoidtype)
	}
}

func (ellipsoidtype *EllipsoidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidTypeUpdateCallback != nil {
		var frontEllipsoidType *EllipsoidType
		if front != nil {
			frontEllipsoidType, _ = front.(*EllipsoidType)
		}
		stage.OnAfterEllipsoidTypeUpdateCallback.OnAfterUpdate(stage, ellipsoidtype, frontEllipsoidType)
	}
}

func (ellipsoidtype *EllipsoidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidTypeDeleteCallback != nil {
		var frontEllipsoidType *EllipsoidType
		if front != nil {
			frontEllipsoidType, _ = front.(*EllipsoidType)
		}
		stage.OnAfterEllipsoidTypeDeleteCallback.OnAfterDelete(stage, ellipsoidtype, frontEllipsoidType)
	}
}

func (ellipsoidalcspropertytype *EllipsoidalCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEllipsoidalCSPropertyTypeCreateCallback != nil {
		stage.OnAfterEllipsoidalCSPropertyTypeCreateCallback.OnAfterCreate(stage, ellipsoidalcspropertytype)
	}
}

func (ellipsoidalcspropertytype *EllipsoidalCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidalCSPropertyTypeUpdateCallback != nil {
		var frontEllipsoidalCSPropertyType *EllipsoidalCSPropertyType
		if front != nil {
			frontEllipsoidalCSPropertyType, _ = front.(*EllipsoidalCSPropertyType)
		}
		stage.OnAfterEllipsoidalCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, ellipsoidalcspropertytype, frontEllipsoidalCSPropertyType)
	}
}

func (ellipsoidalcspropertytype *EllipsoidalCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidalCSPropertyTypeDeleteCallback != nil {
		var frontEllipsoidalCSPropertyType *EllipsoidalCSPropertyType
		if front != nil {
			frontEllipsoidalCSPropertyType, _ = front.(*EllipsoidalCSPropertyType)
		}
		stage.OnAfterEllipsoidalCSPropertyTypeDeleteCallback.OnAfterDelete(stage, ellipsoidalcspropertytype, frontEllipsoidalCSPropertyType)
	}
}

func (ellipsoidalcstype *EllipsoidalCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEllipsoidalCSTypeCreateCallback != nil {
		stage.OnAfterEllipsoidalCSTypeCreateCallback.OnAfterCreate(stage, ellipsoidalcstype)
	}
}

func (ellipsoidalcstype *EllipsoidalCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidalCSTypeUpdateCallback != nil {
		var frontEllipsoidalCSType *EllipsoidalCSType
		if front != nil {
			frontEllipsoidalCSType, _ = front.(*EllipsoidalCSType)
		}
		stage.OnAfterEllipsoidalCSTypeUpdateCallback.OnAfterUpdate(stage, ellipsoidalcstype, frontEllipsoidalCSType)
	}
}

func (ellipsoidalcstype *EllipsoidalCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEllipsoidalCSTypeDeleteCallback != nil {
		var frontEllipsoidalCSType *EllipsoidalCSType
		if front != nil {
			frontEllipsoidalCSType, _ = front.(*EllipsoidalCSType)
		}
		stage.OnAfterEllipsoidalCSTypeDeleteCallback.OnAfterDelete(stage, ellipsoidalcstype, frontEllipsoidalCSType)
	}
}

func (engineeringcrspropertytype *EngineeringCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEngineeringCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterEngineeringCRSPropertyTypeCreateCallback.OnAfterCreate(stage, engineeringcrspropertytype)
	}
}

func (engineeringcrspropertytype *EngineeringCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringCRSPropertyTypeUpdateCallback != nil {
		var frontEngineeringCRSPropertyType *EngineeringCRSPropertyType
		if front != nil {
			frontEngineeringCRSPropertyType, _ = front.(*EngineeringCRSPropertyType)
		}
		stage.OnAfterEngineeringCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, engineeringcrspropertytype, frontEngineeringCRSPropertyType)
	}
}

func (engineeringcrspropertytype *EngineeringCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringCRSPropertyTypeDeleteCallback != nil {
		var frontEngineeringCRSPropertyType *EngineeringCRSPropertyType
		if front != nil {
			frontEngineeringCRSPropertyType, _ = front.(*EngineeringCRSPropertyType)
		}
		stage.OnAfterEngineeringCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, engineeringcrspropertytype, frontEngineeringCRSPropertyType)
	}
}

func (engineeringcrstype *EngineeringCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEngineeringCRSTypeCreateCallback != nil {
		stage.OnAfterEngineeringCRSTypeCreateCallback.OnAfterCreate(stage, engineeringcrstype)
	}
}

func (engineeringcrstype *EngineeringCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringCRSTypeUpdateCallback != nil {
		var frontEngineeringCRSType *EngineeringCRSType
		if front != nil {
			frontEngineeringCRSType, _ = front.(*EngineeringCRSType)
		}
		stage.OnAfterEngineeringCRSTypeUpdateCallback.OnAfterUpdate(stage, engineeringcrstype, frontEngineeringCRSType)
	}
}

func (engineeringcrstype *EngineeringCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringCRSTypeDeleteCallback != nil {
		var frontEngineeringCRSType *EngineeringCRSType
		if front != nil {
			frontEngineeringCRSType, _ = front.(*EngineeringCRSType)
		}
		stage.OnAfterEngineeringCRSTypeDeleteCallback.OnAfterDelete(stage, engineeringcrstype, frontEngineeringCRSType)
	}
}

func (engineeringdatumpropertytype *EngineeringDatumPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEngineeringDatumPropertyTypeCreateCallback != nil {
		stage.OnAfterEngineeringDatumPropertyTypeCreateCallback.OnAfterCreate(stage, engineeringdatumpropertytype)
	}
}

func (engineeringdatumpropertytype *EngineeringDatumPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringDatumPropertyTypeUpdateCallback != nil {
		var frontEngineeringDatumPropertyType *EngineeringDatumPropertyType
		if front != nil {
			frontEngineeringDatumPropertyType, _ = front.(*EngineeringDatumPropertyType)
		}
		stage.OnAfterEngineeringDatumPropertyTypeUpdateCallback.OnAfterUpdate(stage, engineeringdatumpropertytype, frontEngineeringDatumPropertyType)
	}
}

func (engineeringdatumpropertytype *EngineeringDatumPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringDatumPropertyTypeDeleteCallback != nil {
		var frontEngineeringDatumPropertyType *EngineeringDatumPropertyType
		if front != nil {
			frontEngineeringDatumPropertyType, _ = front.(*EngineeringDatumPropertyType)
		}
		stage.OnAfterEngineeringDatumPropertyTypeDeleteCallback.OnAfterDelete(stage, engineeringdatumpropertytype, frontEngineeringDatumPropertyType)
	}
}

func (engineeringdatumtype *EngineeringDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEngineeringDatumTypeCreateCallback != nil {
		stage.OnAfterEngineeringDatumTypeCreateCallback.OnAfterCreate(stage, engineeringdatumtype)
	}
}

func (engineeringdatumtype *EngineeringDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringDatumTypeUpdateCallback != nil {
		var frontEngineeringDatumType *EngineeringDatumType
		if front != nil {
			frontEngineeringDatumType, _ = front.(*EngineeringDatumType)
		}
		stage.OnAfterEngineeringDatumTypeUpdateCallback.OnAfterUpdate(stage, engineeringdatumtype, frontEngineeringDatumType)
	}
}

func (engineeringdatumtype *EngineeringDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEngineeringDatumTypeDeleteCallback != nil {
		var frontEngineeringDatumType *EngineeringDatumType
		if front != nil {
			frontEngineeringDatumType, _ = front.(*EngineeringDatumType)
		}
		stage.OnAfterEngineeringDatumTypeDeleteCallback.OnAfterDelete(stage, engineeringdatumtype, frontEngineeringDatumType)
	}
}

func (envelopetype *EnvelopeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEnvelopeTypeCreateCallback != nil {
		stage.OnAfterEnvelopeTypeCreateCallback.OnAfterCreate(stage, envelopetype)
	}
}

func (envelopetype *EnvelopeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnvelopeTypeUpdateCallback != nil {
		var frontEnvelopeType *EnvelopeType
		if front != nil {
			frontEnvelopeType, _ = front.(*EnvelopeType)
		}
		stage.OnAfterEnvelopeTypeUpdateCallback.OnAfterUpdate(stage, envelopetype, frontEnvelopeType)
	}
}

func (envelopetype *EnvelopeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnvelopeTypeDeleteCallback != nil {
		var frontEnvelopeType *EnvelopeType
		if front != nil {
			frontEnvelopeType, _ = front.(*EnvelopeType)
		}
		stage.OnAfterEnvelopeTypeDeleteCallback.OnAfterDelete(stage, envelopetype, frontEnvelopeType)
	}
}

func (envelopewithtimeperiodtype *EnvelopeWithTimePeriodType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEnvelopeWithTimePeriodTypeCreateCallback != nil {
		stage.OnAfterEnvelopeWithTimePeriodTypeCreateCallback.OnAfterCreate(stage, envelopewithtimeperiodtype)
	}
}

func (envelopewithtimeperiodtype *EnvelopeWithTimePeriodType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnvelopeWithTimePeriodTypeUpdateCallback != nil {
		var frontEnvelopeWithTimePeriodType *EnvelopeWithTimePeriodType
		if front != nil {
			frontEnvelopeWithTimePeriodType, _ = front.(*EnvelopeWithTimePeriodType)
		}
		stage.OnAfterEnvelopeWithTimePeriodTypeUpdateCallback.OnAfterUpdate(stage, envelopewithtimeperiodtype, frontEnvelopeWithTimePeriodType)
	}
}

func (envelopewithtimeperiodtype *EnvelopeWithTimePeriodType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEnvelopeWithTimePeriodTypeDeleteCallback != nil {
		var frontEnvelopeWithTimePeriodType *EnvelopeWithTimePeriodType
		if front != nil {
			frontEnvelopeWithTimePeriodType, _ = front.(*EnvelopeWithTimePeriodType)
		}
		stage.OnAfterEnvelopeWithTimePeriodTypeDeleteCallback.OnAfterDelete(stage, envelopewithtimeperiodtype, frontEnvelopeWithTimePeriodType)
	}
}

func (faceortoposolidpropertytype *FaceOrTopoSolidPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFaceOrTopoSolidPropertyTypeCreateCallback != nil {
		stage.OnAfterFaceOrTopoSolidPropertyTypeCreateCallback.OnAfterCreate(stage, faceortoposolidpropertytype)
	}
}

func (faceortoposolidpropertytype *FaceOrTopoSolidPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFaceOrTopoSolidPropertyTypeUpdateCallback != nil {
		var frontFaceOrTopoSolidPropertyType *FaceOrTopoSolidPropertyType
		if front != nil {
			frontFaceOrTopoSolidPropertyType, _ = front.(*FaceOrTopoSolidPropertyType)
		}
		stage.OnAfterFaceOrTopoSolidPropertyTypeUpdateCallback.OnAfterUpdate(stage, faceortoposolidpropertytype, frontFaceOrTopoSolidPropertyType)
	}
}

func (faceortoposolidpropertytype *FaceOrTopoSolidPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFaceOrTopoSolidPropertyTypeDeleteCallback != nil {
		var frontFaceOrTopoSolidPropertyType *FaceOrTopoSolidPropertyType
		if front != nil {
			frontFaceOrTopoSolidPropertyType, _ = front.(*FaceOrTopoSolidPropertyType)
		}
		stage.OnAfterFaceOrTopoSolidPropertyTypeDeleteCallback.OnAfterDelete(stage, faceortoposolidpropertytype, frontFaceOrTopoSolidPropertyType)
	}
}

func (facetype *FaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFaceTypeCreateCallback != nil {
		stage.OnAfterFaceTypeCreateCallback.OnAfterCreate(stage, facetype)
	}
}

func (facetype *FaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFaceTypeUpdateCallback != nil {
		var frontFaceType *FaceType
		if front != nil {
			frontFaceType, _ = front.(*FaceType)
		}
		stage.OnAfterFaceTypeUpdateCallback.OnAfterUpdate(stage, facetype, frontFaceType)
	}
}

func (facetype *FaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFaceTypeDeleteCallback != nil {
		var frontFaceType *FaceType
		if front != nil {
			frontFaceType, _ = front.(*FaceType)
		}
		stage.OnAfterFaceTypeDeleteCallback.OnAfterDelete(stage, facetype, frontFaceType)
	}
}

func (featurearraypropertytype *FeatureArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFeatureArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterFeatureArrayPropertyTypeCreateCallback.OnAfterCreate(stage, featurearraypropertytype)
	}
}

func (featurearraypropertytype *FeatureArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureArrayPropertyTypeUpdateCallback != nil {
		var frontFeatureArrayPropertyType *FeatureArrayPropertyType
		if front != nil {
			frontFeatureArrayPropertyType, _ = front.(*FeatureArrayPropertyType)
		}
		stage.OnAfterFeatureArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, featurearraypropertytype, frontFeatureArrayPropertyType)
	}
}

func (featurearraypropertytype *FeatureArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureArrayPropertyTypeDeleteCallback != nil {
		var frontFeatureArrayPropertyType *FeatureArrayPropertyType
		if front != nil {
			frontFeatureArrayPropertyType, _ = front.(*FeatureArrayPropertyType)
		}
		stage.OnAfterFeatureArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, featurearraypropertytype, frontFeatureArrayPropertyType)
	}
}

func (featurecollectiontype *FeatureCollectionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFeatureCollectionTypeCreateCallback != nil {
		stage.OnAfterFeatureCollectionTypeCreateCallback.OnAfterCreate(stage, featurecollectiontype)
	}
}

func (featurecollectiontype *FeatureCollectionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureCollectionTypeUpdateCallback != nil {
		var frontFeatureCollectionType *FeatureCollectionType
		if front != nil {
			frontFeatureCollectionType, _ = front.(*FeatureCollectionType)
		}
		stage.OnAfterFeatureCollectionTypeUpdateCallback.OnAfterUpdate(stage, featurecollectiontype, frontFeatureCollectionType)
	}
}

func (featurecollectiontype *FeatureCollectionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureCollectionTypeDeleteCallback != nil {
		var frontFeatureCollectionType *FeatureCollectionType
		if front != nil {
			frontFeatureCollectionType, _ = front.(*FeatureCollectionType)
		}
		stage.OnAfterFeatureCollectionTypeDeleteCallback.OnAfterDelete(stage, featurecollectiontype, frontFeatureCollectionType)
	}
}

func (featurepropertytype *FeaturePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFeaturePropertyTypeCreateCallback != nil {
		stage.OnAfterFeaturePropertyTypeCreateCallback.OnAfterCreate(stage, featurepropertytype)
	}
}

func (featurepropertytype *FeaturePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeaturePropertyTypeUpdateCallback != nil {
		var frontFeaturePropertyType *FeaturePropertyType
		if front != nil {
			frontFeaturePropertyType, _ = front.(*FeaturePropertyType)
		}
		stage.OnAfterFeaturePropertyTypeUpdateCallback.OnAfterUpdate(stage, featurepropertytype, frontFeaturePropertyType)
	}
}

func (featurepropertytype *FeaturePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeaturePropertyTypeDeleteCallback != nil {
		var frontFeaturePropertyType *FeaturePropertyType
		if front != nil {
			frontFeaturePropertyType, _ = front.(*FeaturePropertyType)
		}
		stage.OnAfterFeaturePropertyTypeDeleteCallback.OnAfterDelete(stage, featurepropertytype, frontFeaturePropertyType)
	}
}

func (filetype *FileType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFileTypeCreateCallback != nil {
		stage.OnAfterFileTypeCreateCallback.OnAfterCreate(stage, filetype)
	}
}

func (filetype *FileType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileTypeUpdateCallback != nil {
		var frontFileType *FileType
		if front != nil {
			frontFileType, _ = front.(*FileType)
		}
		stage.OnAfterFileTypeUpdateCallback.OnAfterUpdate(stage, filetype, frontFileType)
	}
}

func (filetype *FileType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFileTypeDeleteCallback != nil {
		var frontFileType *FileType
		if front != nil {
			frontFileType, _ = front.(*FileType)
		}
		stage.OnAfterFileTypeDeleteCallback.OnAfterDelete(stage, filetype, frontFileType)
	}
}

func (formulacitation *FormulaCitation) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormulaCitationCreateCallback != nil {
		stage.OnAfterFormulaCitationCreateCallback.OnAfterCreate(stage, formulacitation)
	}
}

func (formulacitation *FormulaCitation) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormulaCitationUpdateCallback != nil {
		var frontFormulaCitation *FormulaCitation
		if front != nil {
			frontFormulaCitation, _ = front.(*FormulaCitation)
		}
		stage.OnAfterFormulaCitationUpdateCallback.OnAfterUpdate(stage, formulacitation, frontFormulaCitation)
	}
}

func (formulacitation *FormulaCitation) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormulaCitationDeleteCallback != nil {
		var frontFormulaCitation *FormulaCitation
		if front != nil {
			frontFormulaCitation, _ = front.(*FormulaCitation)
		}
		stage.OnAfterFormulaCitationDeleteCallback.OnAfterDelete(stage, formulacitation, frontFormulaCitation)
	}
}

func (formulatype *FormulaType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFormulaTypeCreateCallback != nil {
		stage.OnAfterFormulaTypeCreateCallback.OnAfterCreate(stage, formulatype)
	}
}

func (formulatype *FormulaType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormulaTypeUpdateCallback != nil {
		var frontFormulaType *FormulaType
		if front != nil {
			frontFormulaType, _ = front.(*FormulaType)
		}
		stage.OnAfterFormulaTypeUpdateCallback.OnAfterUpdate(stage, formulatype, frontFormulaType)
	}
}

func (formulatype *FormulaType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFormulaTypeDeleteCallback != nil {
		var frontFormulaType *FormulaType
		if front != nil {
			frontFormulaType, _ = front.(*FormulaType)
		}
		stage.OnAfterFormulaTypeDeleteCallback.OnAfterDelete(stage, formulatype, frontFormulaType)
	}
}

func (gm_object_propertytype *GM_Object_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGM_Object_PropertyTypeCreateCallback != nil {
		stage.OnAfterGM_Object_PropertyTypeCreateCallback.OnAfterCreate(stage, gm_object_propertytype)
	}
}

func (gm_object_propertytype *GM_Object_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGM_Object_PropertyTypeUpdateCallback != nil {
		var frontGM_Object_PropertyType *GM_Object_PropertyType
		if front != nil {
			frontGM_Object_PropertyType, _ = front.(*GM_Object_PropertyType)
		}
		stage.OnAfterGM_Object_PropertyTypeUpdateCallback.OnAfterUpdate(stage, gm_object_propertytype, frontGM_Object_PropertyType)
	}
}

func (gm_object_propertytype *GM_Object_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGM_Object_PropertyTypeDeleteCallback != nil {
		var frontGM_Object_PropertyType *GM_Object_PropertyType
		if front != nil {
			frontGM_Object_PropertyType, _ = front.(*GM_Object_PropertyType)
		}
		stage.OnAfterGM_Object_PropertyTypeDeleteCallback.OnAfterDelete(stage, gm_object_propertytype, frontGM_Object_PropertyType)
	}
}

func (gm_point_propertytype *GM_Point_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGM_Point_PropertyTypeCreateCallback != nil {
		stage.OnAfterGM_Point_PropertyTypeCreateCallback.OnAfterCreate(stage, gm_point_propertytype)
	}
}

func (gm_point_propertytype *GM_Point_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGM_Point_PropertyTypeUpdateCallback != nil {
		var frontGM_Point_PropertyType *GM_Point_PropertyType
		if front != nil {
			frontGM_Point_PropertyType, _ = front.(*GM_Point_PropertyType)
		}
		stage.OnAfterGM_Point_PropertyTypeUpdateCallback.OnAfterUpdate(stage, gm_point_propertytype, frontGM_Point_PropertyType)
	}
}

func (gm_point_propertytype *GM_Point_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGM_Point_PropertyTypeDeleteCallback != nil {
		var frontGM_Point_PropertyType *GM_Point_PropertyType
		if front != nil {
			frontGM_Point_PropertyType, _ = front.(*GM_Point_PropertyType)
		}
		stage.OnAfterGM_Point_PropertyTypeDeleteCallback.OnAfterDelete(stage, gm_point_propertytype, frontGM_Point_PropertyType)
	}
}

func (generalconversionpropertytype *GeneralConversionPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeneralConversionPropertyTypeCreateCallback != nil {
		stage.OnAfterGeneralConversionPropertyTypeCreateCallback.OnAfterCreate(stage, generalconversionpropertytype)
	}
}

func (generalconversionpropertytype *GeneralConversionPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeneralConversionPropertyTypeUpdateCallback != nil {
		var frontGeneralConversionPropertyType *GeneralConversionPropertyType
		if front != nil {
			frontGeneralConversionPropertyType, _ = front.(*GeneralConversionPropertyType)
		}
		stage.OnAfterGeneralConversionPropertyTypeUpdateCallback.OnAfterUpdate(stage, generalconversionpropertytype, frontGeneralConversionPropertyType)
	}
}

func (generalconversionpropertytype *GeneralConversionPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeneralConversionPropertyTypeDeleteCallback != nil {
		var frontGeneralConversionPropertyType *GeneralConversionPropertyType
		if front != nil {
			frontGeneralConversionPropertyType, _ = front.(*GeneralConversionPropertyType)
		}
		stage.OnAfterGeneralConversionPropertyTypeDeleteCallback.OnAfterDelete(stage, generalconversionpropertytype, frontGeneralConversionPropertyType)
	}
}

func (generaltransformationpropertytype *GeneralTransformationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeneralTransformationPropertyTypeCreateCallback != nil {
		stage.OnAfterGeneralTransformationPropertyTypeCreateCallback.OnAfterCreate(stage, generaltransformationpropertytype)
	}
}

func (generaltransformationpropertytype *GeneralTransformationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeneralTransformationPropertyTypeUpdateCallback != nil {
		var frontGeneralTransformationPropertyType *GeneralTransformationPropertyType
		if front != nil {
			frontGeneralTransformationPropertyType, _ = front.(*GeneralTransformationPropertyType)
		}
		stage.OnAfterGeneralTransformationPropertyTypeUpdateCallback.OnAfterUpdate(stage, generaltransformationpropertytype, frontGeneralTransformationPropertyType)
	}
}

func (generaltransformationpropertytype *GeneralTransformationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeneralTransformationPropertyTypeDeleteCallback != nil {
		var frontGeneralTransformationPropertyType *GeneralTransformationPropertyType
		if front != nil {
			frontGeneralTransformationPropertyType, _ = front.(*GeneralTransformationPropertyType)
		}
		stage.OnAfterGeneralTransformationPropertyTypeDeleteCallback.OnAfterDelete(stage, generaltransformationpropertytype, frontGeneralTransformationPropertyType)
	}
}

func (genericmetadatatype *GenericMetaDataType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGenericMetaDataTypeCreateCallback != nil {
		stage.OnAfterGenericMetaDataTypeCreateCallback.OnAfterCreate(stage, genericmetadatatype)
	}
}

func (genericmetadatatype *GenericMetaDataType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGenericMetaDataTypeUpdateCallback != nil {
		var frontGenericMetaDataType *GenericMetaDataType
		if front != nil {
			frontGenericMetaDataType, _ = front.(*GenericMetaDataType)
		}
		stage.OnAfterGenericMetaDataTypeUpdateCallback.OnAfterUpdate(stage, genericmetadatatype, frontGenericMetaDataType)
	}
}

func (genericmetadatatype *GenericMetaDataType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGenericMetaDataTypeDeleteCallback != nil {
		var frontGenericMetaDataType *GenericMetaDataType
		if front != nil {
			frontGenericMetaDataType, _ = front.(*GenericMetaDataType)
		}
		stage.OnAfterGenericMetaDataTypeDeleteCallback.OnAfterDelete(stage, genericmetadatatype, frontGenericMetaDataType)
	}
}

func (genericname_propertytype *GenericName_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGenericName_PropertyTypeCreateCallback != nil {
		stage.OnAfterGenericName_PropertyTypeCreateCallback.OnAfterCreate(stage, genericname_propertytype)
	}
}

func (genericname_propertytype *GenericName_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGenericName_PropertyTypeUpdateCallback != nil {
		var frontGenericName_PropertyType *GenericName_PropertyType
		if front != nil {
			frontGenericName_PropertyType, _ = front.(*GenericName_PropertyType)
		}
		stage.OnAfterGenericName_PropertyTypeUpdateCallback.OnAfterUpdate(stage, genericname_propertytype, frontGenericName_PropertyType)
	}
}

func (genericname_propertytype *GenericName_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGenericName_PropertyTypeDeleteCallback != nil {
		var frontGenericName_PropertyType *GenericName_PropertyType
		if front != nil {
			frontGenericName_PropertyType, _ = front.(*GenericName_PropertyType)
		}
		stage.OnAfterGenericName_PropertyTypeDeleteCallback.OnAfterDelete(stage, genericname_propertytype, frontGenericName_PropertyType)
	}
}

func (geocentriccrspropertytype *GeocentricCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeocentricCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterGeocentricCRSPropertyTypeCreateCallback.OnAfterCreate(stage, geocentriccrspropertytype)
	}
}

func (geocentriccrspropertytype *GeocentricCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeocentricCRSPropertyTypeUpdateCallback != nil {
		var frontGeocentricCRSPropertyType *GeocentricCRSPropertyType
		if front != nil {
			frontGeocentricCRSPropertyType, _ = front.(*GeocentricCRSPropertyType)
		}
		stage.OnAfterGeocentricCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, geocentriccrspropertytype, frontGeocentricCRSPropertyType)
	}
}

func (geocentriccrspropertytype *GeocentricCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeocentricCRSPropertyTypeDeleteCallback != nil {
		var frontGeocentricCRSPropertyType *GeocentricCRSPropertyType
		if front != nil {
			frontGeocentricCRSPropertyType, _ = front.(*GeocentricCRSPropertyType)
		}
		stage.OnAfterGeocentricCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, geocentriccrspropertytype, frontGeocentricCRSPropertyType)
	}
}

func (geocentriccrstype *GeocentricCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeocentricCRSTypeCreateCallback != nil {
		stage.OnAfterGeocentricCRSTypeCreateCallback.OnAfterCreate(stage, geocentriccrstype)
	}
}

func (geocentriccrstype *GeocentricCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeocentricCRSTypeUpdateCallback != nil {
		var frontGeocentricCRSType *GeocentricCRSType
		if front != nil {
			frontGeocentricCRSType, _ = front.(*GeocentricCRSType)
		}
		stage.OnAfterGeocentricCRSTypeUpdateCallback.OnAfterUpdate(stage, geocentriccrstype, frontGeocentricCRSType)
	}
}

func (geocentriccrstype *GeocentricCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeocentricCRSTypeDeleteCallback != nil {
		var frontGeocentricCRSType *GeocentricCRSType
		if front != nil {
			frontGeocentricCRSType, _ = front.(*GeocentricCRSType)
		}
		stage.OnAfterGeocentricCRSTypeDeleteCallback.OnAfterDelete(stage, geocentriccrstype, frontGeocentricCRSType)
	}
}

func (geodesicstringtype *GeodesicStringType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodesicStringTypeCreateCallback != nil {
		stage.OnAfterGeodesicStringTypeCreateCallback.OnAfterCreate(stage, geodesicstringtype)
	}
}

func (geodesicstringtype *GeodesicStringType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicStringTypeUpdateCallback != nil {
		var frontGeodesicStringType *GeodesicStringType
		if front != nil {
			frontGeodesicStringType, _ = front.(*GeodesicStringType)
		}
		stage.OnAfterGeodesicStringTypeUpdateCallback.OnAfterUpdate(stage, geodesicstringtype, frontGeodesicStringType)
	}
}

func (geodesicstringtype *GeodesicStringType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicStringTypeDeleteCallback != nil {
		var frontGeodesicStringType *GeodesicStringType
		if front != nil {
			frontGeodesicStringType, _ = front.(*GeodesicStringType)
		}
		stage.OnAfterGeodesicStringTypeDeleteCallback.OnAfterDelete(stage, geodesicstringtype, frontGeodesicStringType)
	}
}

func (geodesictype *GeodesicType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodesicTypeCreateCallback != nil {
		stage.OnAfterGeodesicTypeCreateCallback.OnAfterCreate(stage, geodesictype)
	}
}

func (geodesictype *GeodesicType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicTypeUpdateCallback != nil {
		var frontGeodesicType *GeodesicType
		if front != nil {
			frontGeodesicType, _ = front.(*GeodesicType)
		}
		stage.OnAfterGeodesicTypeUpdateCallback.OnAfterUpdate(stage, geodesictype, frontGeodesicType)
	}
}

func (geodesictype *GeodesicType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodesicTypeDeleteCallback != nil {
		var frontGeodesicType *GeodesicType
		if front != nil {
			frontGeodesicType, _ = front.(*GeodesicType)
		}
		stage.OnAfterGeodesicTypeDeleteCallback.OnAfterDelete(stage, geodesictype, frontGeodesicType)
	}
}

func (geodeticcrspropertytype *GeodeticCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodeticCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterGeodeticCRSPropertyTypeCreateCallback.OnAfterCreate(stage, geodeticcrspropertytype)
	}
}

func (geodeticcrspropertytype *GeodeticCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticCRSPropertyTypeUpdateCallback != nil {
		var frontGeodeticCRSPropertyType *GeodeticCRSPropertyType
		if front != nil {
			frontGeodeticCRSPropertyType, _ = front.(*GeodeticCRSPropertyType)
		}
		stage.OnAfterGeodeticCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, geodeticcrspropertytype, frontGeodeticCRSPropertyType)
	}
}

func (geodeticcrspropertytype *GeodeticCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticCRSPropertyTypeDeleteCallback != nil {
		var frontGeodeticCRSPropertyType *GeodeticCRSPropertyType
		if front != nil {
			frontGeodeticCRSPropertyType, _ = front.(*GeodeticCRSPropertyType)
		}
		stage.OnAfterGeodeticCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, geodeticcrspropertytype, frontGeodeticCRSPropertyType)
	}
}

func (geodeticcrstype *GeodeticCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodeticCRSTypeCreateCallback != nil {
		stage.OnAfterGeodeticCRSTypeCreateCallback.OnAfterCreate(stage, geodeticcrstype)
	}
}

func (geodeticcrstype *GeodeticCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticCRSTypeUpdateCallback != nil {
		var frontGeodeticCRSType *GeodeticCRSType
		if front != nil {
			frontGeodeticCRSType, _ = front.(*GeodeticCRSType)
		}
		stage.OnAfterGeodeticCRSTypeUpdateCallback.OnAfterUpdate(stage, geodeticcrstype, frontGeodeticCRSType)
	}
}

func (geodeticcrstype *GeodeticCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticCRSTypeDeleteCallback != nil {
		var frontGeodeticCRSType *GeodeticCRSType
		if front != nil {
			frontGeodeticCRSType, _ = front.(*GeodeticCRSType)
		}
		stage.OnAfterGeodeticCRSTypeDeleteCallback.OnAfterDelete(stage, geodeticcrstype, frontGeodeticCRSType)
	}
}

func (geodeticdatumpropertytype *GeodeticDatumPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodeticDatumPropertyTypeCreateCallback != nil {
		stage.OnAfterGeodeticDatumPropertyTypeCreateCallback.OnAfterCreate(stage, geodeticdatumpropertytype)
	}
}

func (geodeticdatumpropertytype *GeodeticDatumPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticDatumPropertyTypeUpdateCallback != nil {
		var frontGeodeticDatumPropertyType *GeodeticDatumPropertyType
		if front != nil {
			frontGeodeticDatumPropertyType, _ = front.(*GeodeticDatumPropertyType)
		}
		stage.OnAfterGeodeticDatumPropertyTypeUpdateCallback.OnAfterUpdate(stage, geodeticdatumpropertytype, frontGeodeticDatumPropertyType)
	}
}

func (geodeticdatumpropertytype *GeodeticDatumPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticDatumPropertyTypeDeleteCallback != nil {
		var frontGeodeticDatumPropertyType *GeodeticDatumPropertyType
		if front != nil {
			frontGeodeticDatumPropertyType, _ = front.(*GeodeticDatumPropertyType)
		}
		stage.OnAfterGeodeticDatumPropertyTypeDeleteCallback.OnAfterDelete(stage, geodeticdatumpropertytype, frontGeodeticDatumPropertyType)
	}
}

func (geodeticdatumtype *GeodeticDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeodeticDatumTypeCreateCallback != nil {
		stage.OnAfterGeodeticDatumTypeCreateCallback.OnAfterCreate(stage, geodeticdatumtype)
	}
}

func (geodeticdatumtype *GeodeticDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticDatumTypeUpdateCallback != nil {
		var frontGeodeticDatumType *GeodeticDatumType
		if front != nil {
			frontGeodeticDatumType, _ = front.(*GeodeticDatumType)
		}
		stage.OnAfterGeodeticDatumTypeUpdateCallback.OnAfterUpdate(stage, geodeticdatumtype, frontGeodeticDatumType)
	}
}

func (geodeticdatumtype *GeodeticDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeodeticDatumTypeDeleteCallback != nil {
		var frontGeodeticDatumType *GeodeticDatumType
		if front != nil {
			frontGeodeticDatumType, _ = front.(*GeodeticDatumType)
		}
		stage.OnAfterGeodeticDatumTypeDeleteCallback.OnAfterDelete(stage, geodeticdatumtype, frontGeodeticDatumType)
	}
}

func (geographiccrspropertytype *GeographicCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeographicCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterGeographicCRSPropertyTypeCreateCallback.OnAfterCreate(stage, geographiccrspropertytype)
	}
}

func (geographiccrspropertytype *GeographicCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeographicCRSPropertyTypeUpdateCallback != nil {
		var frontGeographicCRSPropertyType *GeographicCRSPropertyType
		if front != nil {
			frontGeographicCRSPropertyType, _ = front.(*GeographicCRSPropertyType)
		}
		stage.OnAfterGeographicCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, geographiccrspropertytype, frontGeographicCRSPropertyType)
	}
}

func (geographiccrspropertytype *GeographicCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeographicCRSPropertyTypeDeleteCallback != nil {
		var frontGeographicCRSPropertyType *GeographicCRSPropertyType
		if front != nil {
			frontGeographicCRSPropertyType, _ = front.(*GeographicCRSPropertyType)
		}
		stage.OnAfterGeographicCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, geographiccrspropertytype, frontGeographicCRSPropertyType)
	}
}

func (geographiccrstype *GeographicCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeographicCRSTypeCreateCallback != nil {
		stage.OnAfterGeographicCRSTypeCreateCallback.OnAfterCreate(stage, geographiccrstype)
	}
}

func (geographiccrstype *GeographicCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeographicCRSTypeUpdateCallback != nil {
		var frontGeographicCRSType *GeographicCRSType
		if front != nil {
			frontGeographicCRSType, _ = front.(*GeographicCRSType)
		}
		stage.OnAfterGeographicCRSTypeUpdateCallback.OnAfterUpdate(stage, geographiccrstype, frontGeographicCRSType)
	}
}

func (geographiccrstype *GeographicCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeographicCRSTypeDeleteCallback != nil {
		var frontGeographicCRSType *GeographicCRSType
		if front != nil {
			frontGeographicCRSType, _ = front.(*GeographicCRSType)
		}
		stage.OnAfterGeographicCRSTypeDeleteCallback.OnAfterDelete(stage, geographiccrstype, frontGeographicCRSType)
	}
}

func (geometriccomplexpropertytype *GeometricComplexPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeometricComplexPropertyTypeCreateCallback != nil {
		stage.OnAfterGeometricComplexPropertyTypeCreateCallback.OnAfterCreate(stage, geometriccomplexpropertytype)
	}
}

func (geometriccomplexpropertytype *GeometricComplexPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometricComplexPropertyTypeUpdateCallback != nil {
		var frontGeometricComplexPropertyType *GeometricComplexPropertyType
		if front != nil {
			frontGeometricComplexPropertyType, _ = front.(*GeometricComplexPropertyType)
		}
		stage.OnAfterGeometricComplexPropertyTypeUpdateCallback.OnAfterUpdate(stage, geometriccomplexpropertytype, frontGeometricComplexPropertyType)
	}
}

func (geometriccomplexpropertytype *GeometricComplexPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometricComplexPropertyTypeDeleteCallback != nil {
		var frontGeometricComplexPropertyType *GeometricComplexPropertyType
		if front != nil {
			frontGeometricComplexPropertyType, _ = front.(*GeometricComplexPropertyType)
		}
		stage.OnAfterGeometricComplexPropertyTypeDeleteCallback.OnAfterDelete(stage, geometriccomplexpropertytype, frontGeometricComplexPropertyType)
	}
}

func (geometriccomplextype *GeometricComplexType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeometricComplexTypeCreateCallback != nil {
		stage.OnAfterGeometricComplexTypeCreateCallback.OnAfterCreate(stage, geometriccomplextype)
	}
}

func (geometriccomplextype *GeometricComplexType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometricComplexTypeUpdateCallback != nil {
		var frontGeometricComplexType *GeometricComplexType
		if front != nil {
			frontGeometricComplexType, _ = front.(*GeometricComplexType)
		}
		stage.OnAfterGeometricComplexTypeUpdateCallback.OnAfterUpdate(stage, geometriccomplextype, frontGeometricComplexType)
	}
}

func (geometriccomplextype *GeometricComplexType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometricComplexTypeDeleteCallback != nil {
		var frontGeometricComplexType *GeometricComplexType
		if front != nil {
			frontGeometricComplexType, _ = front.(*GeometricComplexType)
		}
		stage.OnAfterGeometricComplexTypeDeleteCallback.OnAfterDelete(stage, geometriccomplextype, frontGeometricComplexType)
	}
}

func (geometricprimitivepropertytype *GeometricPrimitivePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeometricPrimitivePropertyTypeCreateCallback != nil {
		stage.OnAfterGeometricPrimitivePropertyTypeCreateCallback.OnAfterCreate(stage, geometricprimitivepropertytype)
	}
}

func (geometricprimitivepropertytype *GeometricPrimitivePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometricPrimitivePropertyTypeUpdateCallback != nil {
		var frontGeometricPrimitivePropertyType *GeometricPrimitivePropertyType
		if front != nil {
			frontGeometricPrimitivePropertyType, _ = front.(*GeometricPrimitivePropertyType)
		}
		stage.OnAfterGeometricPrimitivePropertyTypeUpdateCallback.OnAfterUpdate(stage, geometricprimitivepropertytype, frontGeometricPrimitivePropertyType)
	}
}

func (geometricprimitivepropertytype *GeometricPrimitivePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometricPrimitivePropertyTypeDeleteCallback != nil {
		var frontGeometricPrimitivePropertyType *GeometricPrimitivePropertyType
		if front != nil {
			frontGeometricPrimitivePropertyType, _ = front.(*GeometricPrimitivePropertyType)
		}
		stage.OnAfterGeometricPrimitivePropertyTypeDeleteCallback.OnAfterDelete(stage, geometricprimitivepropertytype, frontGeometricPrimitivePropertyType)
	}
}

func (geometryarraypropertytype *GeometryArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeometryArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterGeometryArrayPropertyTypeCreateCallback.OnAfterCreate(stage, geometryarraypropertytype)
	}
}

func (geometryarraypropertytype *GeometryArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometryArrayPropertyTypeUpdateCallback != nil {
		var frontGeometryArrayPropertyType *GeometryArrayPropertyType
		if front != nil {
			frontGeometryArrayPropertyType, _ = front.(*GeometryArrayPropertyType)
		}
		stage.OnAfterGeometryArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, geometryarraypropertytype, frontGeometryArrayPropertyType)
	}
}

func (geometryarraypropertytype *GeometryArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometryArrayPropertyTypeDeleteCallback != nil {
		var frontGeometryArrayPropertyType *GeometryArrayPropertyType
		if front != nil {
			frontGeometryArrayPropertyType, _ = front.(*GeometryArrayPropertyType)
		}
		stage.OnAfterGeometryArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, geometryarraypropertytype, frontGeometryArrayPropertyType)
	}
}

func (geometrypropertytype *GeometryPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGeometryPropertyTypeCreateCallback != nil {
		stage.OnAfterGeometryPropertyTypeCreateCallback.OnAfterCreate(stage, geometrypropertytype)
	}
}

func (geometrypropertytype *GeometryPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometryPropertyTypeUpdateCallback != nil {
		var frontGeometryPropertyType *GeometryPropertyType
		if front != nil {
			frontGeometryPropertyType, _ = front.(*GeometryPropertyType)
		}
		stage.OnAfterGeometryPropertyTypeUpdateCallback.OnAfterUpdate(stage, geometrypropertytype, frontGeometryPropertyType)
	}
}

func (geometrypropertytype *GeometryPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGeometryPropertyTypeDeleteCallback != nil {
		var frontGeometryPropertyType *GeometryPropertyType
		if front != nil {
			frontGeometryPropertyType, _ = front.(*GeometryPropertyType)
		}
		stage.OnAfterGeometryPropertyTypeDeleteCallback.OnAfterDelete(stage, geometrypropertytype, frontGeometryPropertyType)
	}
}

func (gridenvelopetype *GridEnvelopeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGridEnvelopeTypeCreateCallback != nil {
		stage.OnAfterGridEnvelopeTypeCreateCallback.OnAfterCreate(stage, gridenvelopetype)
	}
}

func (gridenvelopetype *GridEnvelopeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridEnvelopeTypeUpdateCallback != nil {
		var frontGridEnvelopeType *GridEnvelopeType
		if front != nil {
			frontGridEnvelopeType, _ = front.(*GridEnvelopeType)
		}
		stage.OnAfterGridEnvelopeTypeUpdateCallback.OnAfterUpdate(stage, gridenvelopetype, frontGridEnvelopeType)
	}
}

func (gridenvelopetype *GridEnvelopeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridEnvelopeTypeDeleteCallback != nil {
		var frontGridEnvelopeType *GridEnvelopeType
		if front != nil {
			frontGridEnvelopeType, _ = front.(*GridEnvelopeType)
		}
		stage.OnAfterGridEnvelopeTypeDeleteCallback.OnAfterDelete(stage, gridenvelopetype, frontGridEnvelopeType)
	}
}

func (gridfunctiontype *GridFunctionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGridFunctionTypeCreateCallback != nil {
		stage.OnAfterGridFunctionTypeCreateCallback.OnAfterCreate(stage, gridfunctiontype)
	}
}

func (gridfunctiontype *GridFunctionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridFunctionTypeUpdateCallback != nil {
		var frontGridFunctionType *GridFunctionType
		if front != nil {
			frontGridFunctionType, _ = front.(*GridFunctionType)
		}
		stage.OnAfterGridFunctionTypeUpdateCallback.OnAfterUpdate(stage, gridfunctiontype, frontGridFunctionType)
	}
}

func (gridfunctiontype *GridFunctionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridFunctionTypeDeleteCallback != nil {
		var frontGridFunctionType *GridFunctionType
		if front != nil {
			frontGridFunctionType, _ = front.(*GridFunctionType)
		}
		stage.OnAfterGridFunctionTypeDeleteCallback.OnAfterDelete(stage, gridfunctiontype, frontGridFunctionType)
	}
}

func (gridlimitstype *GridLimitsType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGridLimitsTypeCreateCallback != nil {
		stage.OnAfterGridLimitsTypeCreateCallback.OnAfterCreate(stage, gridlimitstype)
	}
}

func (gridlimitstype *GridLimitsType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridLimitsTypeUpdateCallback != nil {
		var frontGridLimitsType *GridLimitsType
		if front != nil {
			frontGridLimitsType, _ = front.(*GridLimitsType)
		}
		stage.OnAfterGridLimitsTypeUpdateCallback.OnAfterUpdate(stage, gridlimitstype, frontGridLimitsType)
	}
}

func (gridlimitstype *GridLimitsType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridLimitsTypeDeleteCallback != nil {
		var frontGridLimitsType *GridLimitsType
		if front != nil {
			frontGridLimitsType, _ = front.(*GridLimitsType)
		}
		stage.OnAfterGridLimitsTypeDeleteCallback.OnAfterDelete(stage, gridlimitstype, frontGridLimitsType)
	}
}

func (gridtype *GridType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGridTypeCreateCallback != nil {
		stage.OnAfterGridTypeCreateCallback.OnAfterCreate(stage, gridtype)
	}
}

func (gridtype *GridType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridTypeUpdateCallback != nil {
		var frontGridType *GridType
		if front != nil {
			frontGridType, _ = front.(*GridType)
		}
		stage.OnAfterGridTypeUpdateCallback.OnAfterUpdate(stage, gridtype, frontGridType)
	}
}

func (gridtype *GridType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridTypeDeleteCallback != nil {
		var frontGridType *GridType
		if front != nil {
			frontGridType, _ = front.(*GridType)
		}
		stage.OnAfterGridTypeDeleteCallback.OnAfterDelete(stage, gridtype, frontGridType)
	}
}

func (historypropertytype *HistoryPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterHistoryPropertyTypeCreateCallback != nil {
		stage.OnAfterHistoryPropertyTypeCreateCallback.OnAfterCreate(stage, historypropertytype)
	}
}

func (historypropertytype *HistoryPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHistoryPropertyTypeUpdateCallback != nil {
		var frontHistoryPropertyType *HistoryPropertyType
		if front != nil {
			frontHistoryPropertyType, _ = front.(*HistoryPropertyType)
		}
		stage.OnAfterHistoryPropertyTypeUpdateCallback.OnAfterUpdate(stage, historypropertytype, frontHistoryPropertyType)
	}
}

func (historypropertytype *HistoryPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterHistoryPropertyTypeDeleteCallback != nil {
		var frontHistoryPropertyType *HistoryPropertyType
		if front != nil {
			frontHistoryPropertyType, _ = front.(*HistoryPropertyType)
		}
		stage.OnAfterHistoryPropertyTypeDeleteCallback.OnAfterDelete(stage, historypropertytype, frontHistoryPropertyType)
	}
}

func (identifiedobjecttype *IdentifiedObjectType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterIdentifiedObjectTypeCreateCallback != nil {
		stage.OnAfterIdentifiedObjectTypeCreateCallback.OnAfterCreate(stage, identifiedobjecttype)
	}
}

func (identifiedobjecttype *IdentifiedObjectType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIdentifiedObjectTypeUpdateCallback != nil {
		var frontIdentifiedObjectType *IdentifiedObjectType
		if front != nil {
			frontIdentifiedObjectType, _ = front.(*IdentifiedObjectType)
		}
		stage.OnAfterIdentifiedObjectTypeUpdateCallback.OnAfterUpdate(stage, identifiedobjecttype, frontIdentifiedObjectType)
	}
}

func (identifiedobjecttype *IdentifiedObjectType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIdentifiedObjectTypeDeleteCallback != nil {
		var frontIdentifiedObjectType *IdentifiedObjectType
		if front != nil {
			frontIdentifiedObjectType, _ = front.(*IdentifiedObjectType)
		}
		stage.OnAfterIdentifiedObjectTypeDeleteCallback.OnAfterDelete(stage, identifiedobjecttype, frontIdentifiedObjectType)
	}
}

func (imagecrspropertytype *ImageCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterImageCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterImageCRSPropertyTypeCreateCallback.OnAfterCreate(stage, imagecrspropertytype)
	}
}

func (imagecrspropertytype *ImageCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageCRSPropertyTypeUpdateCallback != nil {
		var frontImageCRSPropertyType *ImageCRSPropertyType
		if front != nil {
			frontImageCRSPropertyType, _ = front.(*ImageCRSPropertyType)
		}
		stage.OnAfterImageCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, imagecrspropertytype, frontImageCRSPropertyType)
	}
}

func (imagecrspropertytype *ImageCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageCRSPropertyTypeDeleteCallback != nil {
		var frontImageCRSPropertyType *ImageCRSPropertyType
		if front != nil {
			frontImageCRSPropertyType, _ = front.(*ImageCRSPropertyType)
		}
		stage.OnAfterImageCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, imagecrspropertytype, frontImageCRSPropertyType)
	}
}

func (imagecrstype *ImageCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterImageCRSTypeCreateCallback != nil {
		stage.OnAfterImageCRSTypeCreateCallback.OnAfterCreate(stage, imagecrstype)
	}
}

func (imagecrstype *ImageCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageCRSTypeUpdateCallback != nil {
		var frontImageCRSType *ImageCRSType
		if front != nil {
			frontImageCRSType, _ = front.(*ImageCRSType)
		}
		stage.OnAfterImageCRSTypeUpdateCallback.OnAfterUpdate(stage, imagecrstype, frontImageCRSType)
	}
}

func (imagecrstype *ImageCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageCRSTypeDeleteCallback != nil {
		var frontImageCRSType *ImageCRSType
		if front != nil {
			frontImageCRSType, _ = front.(*ImageCRSType)
		}
		stage.OnAfterImageCRSTypeDeleteCallback.OnAfterDelete(stage, imagecrstype, frontImageCRSType)
	}
}

func (imagedatumpropertytype *ImageDatumPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterImageDatumPropertyTypeCreateCallback != nil {
		stage.OnAfterImageDatumPropertyTypeCreateCallback.OnAfterCreate(stage, imagedatumpropertytype)
	}
}

func (imagedatumpropertytype *ImageDatumPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageDatumPropertyTypeUpdateCallback != nil {
		var frontImageDatumPropertyType *ImageDatumPropertyType
		if front != nil {
			frontImageDatumPropertyType, _ = front.(*ImageDatumPropertyType)
		}
		stage.OnAfterImageDatumPropertyTypeUpdateCallback.OnAfterUpdate(stage, imagedatumpropertytype, frontImageDatumPropertyType)
	}
}

func (imagedatumpropertytype *ImageDatumPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageDatumPropertyTypeDeleteCallback != nil {
		var frontImageDatumPropertyType *ImageDatumPropertyType
		if front != nil {
			frontImageDatumPropertyType, _ = front.(*ImageDatumPropertyType)
		}
		stage.OnAfterImageDatumPropertyTypeDeleteCallback.OnAfterDelete(stage, imagedatumpropertytype, frontImageDatumPropertyType)
	}
}

func (imagedatumtype *ImageDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterImageDatumTypeCreateCallback != nil {
		stage.OnAfterImageDatumTypeCreateCallback.OnAfterCreate(stage, imagedatumtype)
	}
}

func (imagedatumtype *ImageDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageDatumTypeUpdateCallback != nil {
		var frontImageDatumType *ImageDatumType
		if front != nil {
			frontImageDatumType, _ = front.(*ImageDatumType)
		}
		stage.OnAfterImageDatumTypeUpdateCallback.OnAfterUpdate(stage, imagedatumtype, frontImageDatumType)
	}
}

func (imagedatumtype *ImageDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterImageDatumTypeDeleteCallback != nil {
		var frontImageDatumType *ImageDatumType
		if front != nil {
			frontImageDatumType, _ = front.(*ImageDatumType)
		}
		stage.OnAfterImageDatumTypeDeleteCallback.OnAfterDelete(stage, imagedatumtype, frontImageDatumType)
	}
}

func (indirectentrytype *IndirectEntryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterIndirectEntryTypeCreateCallback != nil {
		stage.OnAfterIndirectEntryTypeCreateCallback.OnAfterCreate(stage, indirectentrytype)
	}
}

func (indirectentrytype *IndirectEntryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIndirectEntryTypeUpdateCallback != nil {
		var frontIndirectEntryType *IndirectEntryType
		if front != nil {
			frontIndirectEntryType, _ = front.(*IndirectEntryType)
		}
		stage.OnAfterIndirectEntryTypeUpdateCallback.OnAfterUpdate(stage, indirectentrytype, frontIndirectEntryType)
	}
}

func (indirectentrytype *IndirectEntryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterIndirectEntryTypeDeleteCallback != nil {
		var frontIndirectEntryType *IndirectEntryType
		if front != nil {
			frontIndirectEntryType, _ = front.(*IndirectEntryType)
		}
		stage.OnAfterIndirectEntryTypeDeleteCallback.OnAfterDelete(stage, indirectentrytype, frontIndirectEntryType)
	}
}

func (inlinepropertytype *InlinePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInlinePropertyTypeCreateCallback != nil {
		stage.OnAfterInlinePropertyTypeCreateCallback.OnAfterCreate(stage, inlinepropertytype)
	}
}

func (inlinepropertytype *InlinePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInlinePropertyTypeUpdateCallback != nil {
		var frontInlinePropertyType *InlinePropertyType
		if front != nil {
			frontInlinePropertyType, _ = front.(*InlinePropertyType)
		}
		stage.OnAfterInlinePropertyTypeUpdateCallback.OnAfterUpdate(stage, inlinepropertytype, frontInlinePropertyType)
	}
}

func (inlinepropertytype *InlinePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInlinePropertyTypeDeleteCallback != nil {
		var frontInlinePropertyType *InlinePropertyType
		if front != nil {
			frontInlinePropertyType, _ = front.(*InlinePropertyType)
		}
		stage.OnAfterInlinePropertyTypeDeleteCallback.OnAfterDelete(stage, inlinepropertytype, frontInlinePropertyType)
	}
}

func (integer_propertytype *Integer_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInteger_PropertyTypeCreateCallback != nil {
		stage.OnAfterInteger_PropertyTypeCreateCallback.OnAfterCreate(stage, integer_propertytype)
	}
}

func (integer_propertytype *Integer_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInteger_PropertyTypeUpdateCallback != nil {
		var frontInteger_PropertyType *Integer_PropertyType
		if front != nil {
			frontInteger_PropertyType, _ = front.(*Integer_PropertyType)
		}
		stage.OnAfterInteger_PropertyTypeUpdateCallback.OnAfterUpdate(stage, integer_propertytype, frontInteger_PropertyType)
	}
}

func (integer_propertytype *Integer_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInteger_PropertyTypeDeleteCallback != nil {
		var frontInteger_PropertyType *Integer_PropertyType
		if front != nil {
			frontInteger_PropertyType, _ = front.(*Integer_PropertyType)
		}
		stage.OnAfterInteger_PropertyTypeDeleteCallback.OnAfterDelete(stage, integer_propertytype, frontInteger_PropertyType)
	}
}

func (knotpropertytype *KnotPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKnotPropertyTypeCreateCallback != nil {
		stage.OnAfterKnotPropertyTypeCreateCallback.OnAfterCreate(stage, knotpropertytype)
	}
}

func (knotpropertytype *KnotPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKnotPropertyTypeUpdateCallback != nil {
		var frontKnotPropertyType *KnotPropertyType
		if front != nil {
			frontKnotPropertyType, _ = front.(*KnotPropertyType)
		}
		stage.OnAfterKnotPropertyTypeUpdateCallback.OnAfterUpdate(stage, knotpropertytype, frontKnotPropertyType)
	}
}

func (knotpropertytype *KnotPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKnotPropertyTypeDeleteCallback != nil {
		var frontKnotPropertyType *KnotPropertyType
		if front != nil {
			frontKnotPropertyType, _ = front.(*KnotPropertyType)
		}
		stage.OnAfterKnotPropertyTypeDeleteCallback.OnAfterDelete(stage, knotpropertytype, frontKnotPropertyType)
	}
}

func (knottype *KnotType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKnotTypeCreateCallback != nil {
		stage.OnAfterKnotTypeCreateCallback.OnAfterCreate(stage, knottype)
	}
}

func (knottype *KnotType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKnotTypeUpdateCallback != nil {
		var frontKnotType *KnotType
		if front != nil {
			frontKnotType, _ = front.(*KnotType)
		}
		stage.OnAfterKnotTypeUpdateCallback.OnAfterUpdate(stage, knottype, frontKnotType)
	}
}

func (knottype *KnotType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKnotTypeDeleteCallback != nil {
		var frontKnotType *KnotType
		if front != nil {
			frontKnotType, _ = front.(*KnotType)
		}
		stage.OnAfterKnotTypeDeleteCallback.OnAfterDelete(stage, knottype, frontKnotType)
	}
}

func (li_lineage_propertytype *LI_Lineage_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLI_Lineage_PropertyTypeCreateCallback != nil {
		stage.OnAfterLI_Lineage_PropertyTypeCreateCallback.OnAfterCreate(stage, li_lineage_propertytype)
	}
}

func (li_lineage_propertytype *LI_Lineage_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Lineage_PropertyTypeUpdateCallback != nil {
		var frontLI_Lineage_PropertyType *LI_Lineage_PropertyType
		if front != nil {
			frontLI_Lineage_PropertyType, _ = front.(*LI_Lineage_PropertyType)
		}
		stage.OnAfterLI_Lineage_PropertyTypeUpdateCallback.OnAfterUpdate(stage, li_lineage_propertytype, frontLI_Lineage_PropertyType)
	}
}

func (li_lineage_propertytype *LI_Lineage_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Lineage_PropertyTypeDeleteCallback != nil {
		var frontLI_Lineage_PropertyType *LI_Lineage_PropertyType
		if front != nil {
			frontLI_Lineage_PropertyType, _ = front.(*LI_Lineage_PropertyType)
		}
		stage.OnAfterLI_Lineage_PropertyTypeDeleteCallback.OnAfterDelete(stage, li_lineage_propertytype, frontLI_Lineage_PropertyType)
	}
}

func (li_lineage_type *LI_Lineage_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLI_Lineage_TypeCreateCallback != nil {
		stage.OnAfterLI_Lineage_TypeCreateCallback.OnAfterCreate(stage, li_lineage_type)
	}
}

func (li_lineage_type *LI_Lineage_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Lineage_TypeUpdateCallback != nil {
		var frontLI_Lineage_Type *LI_Lineage_Type
		if front != nil {
			frontLI_Lineage_Type, _ = front.(*LI_Lineage_Type)
		}
		stage.OnAfterLI_Lineage_TypeUpdateCallback.OnAfterUpdate(stage, li_lineage_type, frontLI_Lineage_Type)
	}
}

func (li_lineage_type *LI_Lineage_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Lineage_TypeDeleteCallback != nil {
		var frontLI_Lineage_Type *LI_Lineage_Type
		if front != nil {
			frontLI_Lineage_Type, _ = front.(*LI_Lineage_Type)
		}
		stage.OnAfterLI_Lineage_TypeDeleteCallback.OnAfterDelete(stage, li_lineage_type, frontLI_Lineage_Type)
	}
}

func (li_processstep_propertytype *LI_ProcessStep_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLI_ProcessStep_PropertyTypeCreateCallback != nil {
		stage.OnAfterLI_ProcessStep_PropertyTypeCreateCallback.OnAfterCreate(stage, li_processstep_propertytype)
	}
}

func (li_processstep_propertytype *LI_ProcessStep_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_ProcessStep_PropertyTypeUpdateCallback != nil {
		var frontLI_ProcessStep_PropertyType *LI_ProcessStep_PropertyType
		if front != nil {
			frontLI_ProcessStep_PropertyType, _ = front.(*LI_ProcessStep_PropertyType)
		}
		stage.OnAfterLI_ProcessStep_PropertyTypeUpdateCallback.OnAfterUpdate(stage, li_processstep_propertytype, frontLI_ProcessStep_PropertyType)
	}
}

func (li_processstep_propertytype *LI_ProcessStep_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_ProcessStep_PropertyTypeDeleteCallback != nil {
		var frontLI_ProcessStep_PropertyType *LI_ProcessStep_PropertyType
		if front != nil {
			frontLI_ProcessStep_PropertyType, _ = front.(*LI_ProcessStep_PropertyType)
		}
		stage.OnAfterLI_ProcessStep_PropertyTypeDeleteCallback.OnAfterDelete(stage, li_processstep_propertytype, frontLI_ProcessStep_PropertyType)
	}
}

func (li_processstep_type *LI_ProcessStep_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLI_ProcessStep_TypeCreateCallback != nil {
		stage.OnAfterLI_ProcessStep_TypeCreateCallback.OnAfterCreate(stage, li_processstep_type)
	}
}

func (li_processstep_type *LI_ProcessStep_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_ProcessStep_TypeUpdateCallback != nil {
		var frontLI_ProcessStep_Type *LI_ProcessStep_Type
		if front != nil {
			frontLI_ProcessStep_Type, _ = front.(*LI_ProcessStep_Type)
		}
		stage.OnAfterLI_ProcessStep_TypeUpdateCallback.OnAfterUpdate(stage, li_processstep_type, frontLI_ProcessStep_Type)
	}
}

func (li_processstep_type *LI_ProcessStep_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_ProcessStep_TypeDeleteCallback != nil {
		var frontLI_ProcessStep_Type *LI_ProcessStep_Type
		if front != nil {
			frontLI_ProcessStep_Type, _ = front.(*LI_ProcessStep_Type)
		}
		stage.OnAfterLI_ProcessStep_TypeDeleteCallback.OnAfterDelete(stage, li_processstep_type, frontLI_ProcessStep_Type)
	}
}

func (li_source_propertytype *LI_Source_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLI_Source_PropertyTypeCreateCallback != nil {
		stage.OnAfterLI_Source_PropertyTypeCreateCallback.OnAfterCreate(stage, li_source_propertytype)
	}
}

func (li_source_propertytype *LI_Source_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Source_PropertyTypeUpdateCallback != nil {
		var frontLI_Source_PropertyType *LI_Source_PropertyType
		if front != nil {
			frontLI_Source_PropertyType, _ = front.(*LI_Source_PropertyType)
		}
		stage.OnAfterLI_Source_PropertyTypeUpdateCallback.OnAfterUpdate(stage, li_source_propertytype, frontLI_Source_PropertyType)
	}
}

func (li_source_propertytype *LI_Source_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Source_PropertyTypeDeleteCallback != nil {
		var frontLI_Source_PropertyType *LI_Source_PropertyType
		if front != nil {
			frontLI_Source_PropertyType, _ = front.(*LI_Source_PropertyType)
		}
		stage.OnAfterLI_Source_PropertyTypeDeleteCallback.OnAfterDelete(stage, li_source_propertytype, frontLI_Source_PropertyType)
	}
}

func (li_source_type *LI_Source_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLI_Source_TypeCreateCallback != nil {
		stage.OnAfterLI_Source_TypeCreateCallback.OnAfterCreate(stage, li_source_type)
	}
}

func (li_source_type *LI_Source_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Source_TypeUpdateCallback != nil {
		var frontLI_Source_Type *LI_Source_Type
		if front != nil {
			frontLI_Source_Type, _ = front.(*LI_Source_Type)
		}
		stage.OnAfterLI_Source_TypeUpdateCallback.OnAfterUpdate(stage, li_source_type, frontLI_Source_Type)
	}
}

func (li_source_type *LI_Source_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLI_Source_TypeDeleteCallback != nil {
		var frontLI_Source_Type *LI_Source_Type
		if front != nil {
			frontLI_Source_Type, _ = front.(*LI_Source_Type)
		}
		stage.OnAfterLI_Source_TypeDeleteCallback.OnAfterDelete(stage, li_source_type, frontLI_Source_Type)
	}
}

func (languagecode_propertytype *LanguageCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLanguageCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterLanguageCode_PropertyTypeCreateCallback.OnAfterCreate(stage, languagecode_propertytype)
	}
}

func (languagecode_propertytype *LanguageCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLanguageCode_PropertyTypeUpdateCallback != nil {
		var frontLanguageCode_PropertyType *LanguageCode_PropertyType
		if front != nil {
			frontLanguageCode_PropertyType, _ = front.(*LanguageCode_PropertyType)
		}
		stage.OnAfterLanguageCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, languagecode_propertytype, frontLanguageCode_PropertyType)
	}
}

func (languagecode_propertytype *LanguageCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLanguageCode_PropertyTypeDeleteCallback != nil {
		var frontLanguageCode_PropertyType *LanguageCode_PropertyType
		if front != nil {
			frontLanguageCode_PropertyType, _ = front.(*LanguageCode_PropertyType)
		}
		stage.OnAfterLanguageCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, languagecode_propertytype, frontLanguageCode_PropertyType)
	}
}

func (length_propertytype *Length_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLength_PropertyTypeCreateCallback != nil {
		stage.OnAfterLength_PropertyTypeCreateCallback.OnAfterCreate(stage, length_propertytype)
	}
}

func (length_propertytype *Length_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLength_PropertyTypeUpdateCallback != nil {
		var frontLength_PropertyType *Length_PropertyType
		if front != nil {
			frontLength_PropertyType, _ = front.(*Length_PropertyType)
		}
		stage.OnAfterLength_PropertyTypeUpdateCallback.OnAfterUpdate(stage, length_propertytype, frontLength_PropertyType)
	}
}

func (length_propertytype *Length_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLength_PropertyTypeDeleteCallback != nil {
		var frontLength_PropertyType *Length_PropertyType
		if front != nil {
			frontLength_PropertyType, _ = front.(*Length_PropertyType)
		}
		stage.OnAfterLength_PropertyTypeDeleteCallback.OnAfterDelete(stage, length_propertytype, frontLength_PropertyType)
	}
}

func (linestringsegmentarraypropertytype *LineStringSegmentArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLineStringSegmentArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterLineStringSegmentArrayPropertyTypeCreateCallback.OnAfterCreate(stage, linestringsegmentarraypropertytype)
	}
}

func (linestringsegmentarraypropertytype *LineStringSegmentArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringSegmentArrayPropertyTypeUpdateCallback != nil {
		var frontLineStringSegmentArrayPropertyType *LineStringSegmentArrayPropertyType
		if front != nil {
			frontLineStringSegmentArrayPropertyType, _ = front.(*LineStringSegmentArrayPropertyType)
		}
		stage.OnAfterLineStringSegmentArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, linestringsegmentarraypropertytype, frontLineStringSegmentArrayPropertyType)
	}
}

func (linestringsegmentarraypropertytype *LineStringSegmentArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringSegmentArrayPropertyTypeDeleteCallback != nil {
		var frontLineStringSegmentArrayPropertyType *LineStringSegmentArrayPropertyType
		if front != nil {
			frontLineStringSegmentArrayPropertyType, _ = front.(*LineStringSegmentArrayPropertyType)
		}
		stage.OnAfterLineStringSegmentArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, linestringsegmentarraypropertytype, frontLineStringSegmentArrayPropertyType)
	}
}

func (linestringsegmenttype *LineStringSegmentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLineStringSegmentTypeCreateCallback != nil {
		stage.OnAfterLineStringSegmentTypeCreateCallback.OnAfterCreate(stage, linestringsegmenttype)
	}
}

func (linestringsegmenttype *LineStringSegmentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringSegmentTypeUpdateCallback != nil {
		var frontLineStringSegmentType *LineStringSegmentType
		if front != nil {
			frontLineStringSegmentType, _ = front.(*LineStringSegmentType)
		}
		stage.OnAfterLineStringSegmentTypeUpdateCallback.OnAfterUpdate(stage, linestringsegmenttype, frontLineStringSegmentType)
	}
}

func (linestringsegmenttype *LineStringSegmentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringSegmentTypeDeleteCallback != nil {
		var frontLineStringSegmentType *LineStringSegmentType
		if front != nil {
			frontLineStringSegmentType, _ = front.(*LineStringSegmentType)
		}
		stage.OnAfterLineStringSegmentTypeDeleteCallback.OnAfterDelete(stage, linestringsegmenttype, frontLineStringSegmentType)
	}
}

func (linestringtype *LineStringType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLineStringTypeCreateCallback != nil {
		stage.OnAfterLineStringTypeCreateCallback.OnAfterCreate(stage, linestringtype)
	}
}

func (linestringtype *LineStringType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringTypeUpdateCallback != nil {
		var frontLineStringType *LineStringType
		if front != nil {
			frontLineStringType, _ = front.(*LineStringType)
		}
		stage.OnAfterLineStringTypeUpdateCallback.OnAfterUpdate(stage, linestringtype, frontLineStringType)
	}
}

func (linestringtype *LineStringType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLineStringTypeDeleteCallback != nil {
		var frontLineStringType *LineStringType
		if front != nil {
			frontLineStringType, _ = front.(*LineStringType)
		}
		stage.OnAfterLineStringTypeDeleteCallback.OnAfterDelete(stage, linestringtype, frontLineStringType)
	}
}

func (linearcspropertytype *LinearCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinearCSPropertyTypeCreateCallback != nil {
		stage.OnAfterLinearCSPropertyTypeCreateCallback.OnAfterCreate(stage, linearcspropertytype)
	}
}

func (linearcspropertytype *LinearCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearCSPropertyTypeUpdateCallback != nil {
		var frontLinearCSPropertyType *LinearCSPropertyType
		if front != nil {
			frontLinearCSPropertyType, _ = front.(*LinearCSPropertyType)
		}
		stage.OnAfterLinearCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, linearcspropertytype, frontLinearCSPropertyType)
	}
}

func (linearcspropertytype *LinearCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearCSPropertyTypeDeleteCallback != nil {
		var frontLinearCSPropertyType *LinearCSPropertyType
		if front != nil {
			frontLinearCSPropertyType, _ = front.(*LinearCSPropertyType)
		}
		stage.OnAfterLinearCSPropertyTypeDeleteCallback.OnAfterDelete(stage, linearcspropertytype, frontLinearCSPropertyType)
	}
}

func (linearcstype *LinearCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinearCSTypeCreateCallback != nil {
		stage.OnAfterLinearCSTypeCreateCallback.OnAfterCreate(stage, linearcstype)
	}
}

func (linearcstype *LinearCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearCSTypeUpdateCallback != nil {
		var frontLinearCSType *LinearCSType
		if front != nil {
			frontLinearCSType, _ = front.(*LinearCSType)
		}
		stage.OnAfterLinearCSTypeUpdateCallback.OnAfterUpdate(stage, linearcstype, frontLinearCSType)
	}
}

func (linearcstype *LinearCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearCSTypeDeleteCallback != nil {
		var frontLinearCSType *LinearCSType
		if front != nil {
			frontLinearCSType, _ = front.(*LinearCSType)
		}
		stage.OnAfterLinearCSTypeDeleteCallback.OnAfterDelete(stage, linearcstype, frontLinearCSType)
	}
}

func (linearringpropertytype *LinearRingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinearRingPropertyTypeCreateCallback != nil {
		stage.OnAfterLinearRingPropertyTypeCreateCallback.OnAfterCreate(stage, linearringpropertytype)
	}
}

func (linearringpropertytype *LinearRingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearRingPropertyTypeUpdateCallback != nil {
		var frontLinearRingPropertyType *LinearRingPropertyType
		if front != nil {
			frontLinearRingPropertyType, _ = front.(*LinearRingPropertyType)
		}
		stage.OnAfterLinearRingPropertyTypeUpdateCallback.OnAfterUpdate(stage, linearringpropertytype, frontLinearRingPropertyType)
	}
}

func (linearringpropertytype *LinearRingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearRingPropertyTypeDeleteCallback != nil {
		var frontLinearRingPropertyType *LinearRingPropertyType
		if front != nil {
			frontLinearRingPropertyType, _ = front.(*LinearRingPropertyType)
		}
		stage.OnAfterLinearRingPropertyTypeDeleteCallback.OnAfterDelete(stage, linearringpropertytype, frontLinearRingPropertyType)
	}
}

func (linearringtype *LinearRingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLinearRingTypeCreateCallback != nil {
		stage.OnAfterLinearRingTypeCreateCallback.OnAfterCreate(stage, linearringtype)
	}
}

func (linearringtype *LinearRingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearRingTypeUpdateCallback != nil {
		var frontLinearRingType *LinearRingType
		if front != nil {
			frontLinearRingType, _ = front.(*LinearRingType)
		}
		stage.OnAfterLinearRingTypeUpdateCallback.OnAfterUpdate(stage, linearringtype, frontLinearRingType)
	}
}

func (linearringtype *LinearRingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLinearRingTypeDeleteCallback != nil {
		var frontLinearRingType *LinearRingType
		if front != nil {
			frontLinearRingType, _ = front.(*LinearRingType)
		}
		stage.OnAfterLinearRingTypeDeleteCallback.OnAfterDelete(stage, linearringtype, frontLinearRingType)
	}
}

func (localname_propertytype *LocalName_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocalName_PropertyTypeCreateCallback != nil {
		stage.OnAfterLocalName_PropertyTypeCreateCallback.OnAfterCreate(stage, localname_propertytype)
	}
}

func (localname_propertytype *LocalName_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalName_PropertyTypeUpdateCallback != nil {
		var frontLocalName_PropertyType *LocalName_PropertyType
		if front != nil {
			frontLocalName_PropertyType, _ = front.(*LocalName_PropertyType)
		}
		stage.OnAfterLocalName_PropertyTypeUpdateCallback.OnAfterUpdate(stage, localname_propertytype, frontLocalName_PropertyType)
	}
}

func (localname_propertytype *LocalName_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalName_PropertyTypeDeleteCallback != nil {
		var frontLocalName_PropertyType *LocalName_PropertyType
		if front != nil {
			frontLocalName_PropertyType, _ = front.(*LocalName_PropertyType)
		}
		stage.OnAfterLocalName_PropertyTypeDeleteCallback.OnAfterDelete(stage, localname_propertytype, frontLocalName_PropertyType)
	}
}

func (localisedcharacterstring_propertytype *LocalisedCharacterString_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocalisedCharacterString_PropertyTypeCreateCallback != nil {
		stage.OnAfterLocalisedCharacterString_PropertyTypeCreateCallback.OnAfterCreate(stage, localisedcharacterstring_propertytype)
	}
}

func (localisedcharacterstring_propertytype *LocalisedCharacterString_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalisedCharacterString_PropertyTypeUpdateCallback != nil {
		var frontLocalisedCharacterString_PropertyType *LocalisedCharacterString_PropertyType
		if front != nil {
			frontLocalisedCharacterString_PropertyType, _ = front.(*LocalisedCharacterString_PropertyType)
		}
		stage.OnAfterLocalisedCharacterString_PropertyTypeUpdateCallback.OnAfterUpdate(stage, localisedcharacterstring_propertytype, frontLocalisedCharacterString_PropertyType)
	}
}

func (localisedcharacterstring_propertytype *LocalisedCharacterString_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalisedCharacterString_PropertyTypeDeleteCallback != nil {
		var frontLocalisedCharacterString_PropertyType *LocalisedCharacterString_PropertyType
		if front != nil {
			frontLocalisedCharacterString_PropertyType, _ = front.(*LocalisedCharacterString_PropertyType)
		}
		stage.OnAfterLocalisedCharacterString_PropertyTypeDeleteCallback.OnAfterDelete(stage, localisedcharacterstring_propertytype, frontLocalisedCharacterString_PropertyType)
	}
}

func (localisedcharacterstring_type *LocalisedCharacterString_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocalisedCharacterString_TypeCreateCallback != nil {
		stage.OnAfterLocalisedCharacterString_TypeCreateCallback.OnAfterCreate(stage, localisedcharacterstring_type)
	}
}

func (localisedcharacterstring_type *LocalisedCharacterString_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalisedCharacterString_TypeUpdateCallback != nil {
		var frontLocalisedCharacterString_Type *LocalisedCharacterString_Type
		if front != nil {
			frontLocalisedCharacterString_Type, _ = front.(*LocalisedCharacterString_Type)
		}
		stage.OnAfterLocalisedCharacterString_TypeUpdateCallback.OnAfterUpdate(stage, localisedcharacterstring_type, frontLocalisedCharacterString_Type)
	}
}

func (localisedcharacterstring_type *LocalisedCharacterString_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocalisedCharacterString_TypeDeleteCallback != nil {
		var frontLocalisedCharacterString_Type *LocalisedCharacterString_Type
		if front != nil {
			frontLocalisedCharacterString_Type, _ = front.(*LocalisedCharacterString_Type)
		}
		stage.OnAfterLocalisedCharacterString_TypeDeleteCallback.OnAfterDelete(stage, localisedcharacterstring_type, frontLocalisedCharacterString_Type)
	}
}

func (locationpropertytype *LocationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocationPropertyTypeCreateCallback != nil {
		stage.OnAfterLocationPropertyTypeCreateCallback.OnAfterCreate(stage, locationpropertytype)
	}
}

func (locationpropertytype *LocationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocationPropertyTypeUpdateCallback != nil {
		var frontLocationPropertyType *LocationPropertyType
		if front != nil {
			frontLocationPropertyType, _ = front.(*LocationPropertyType)
		}
		stage.OnAfterLocationPropertyTypeUpdateCallback.OnAfterUpdate(stage, locationpropertytype, frontLocationPropertyType)
	}
}

func (locationpropertytype *LocationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocationPropertyTypeDeleteCallback != nil {
		var frontLocationPropertyType *LocationPropertyType
		if front != nil {
			frontLocationPropertyType, _ = front.(*LocationPropertyType)
		}
		stage.OnAfterLocationPropertyTypeDeleteCallback.OnAfterDelete(stage, locationpropertytype, frontLocationPropertyType)
	}
}

func (md_aggregateinformation_propertytype *MD_AggregateInformation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_AggregateInformation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_AggregateInformation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_aggregateinformation_propertytype)
	}
}

func (md_aggregateinformation_propertytype *MD_AggregateInformation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_AggregateInformation_PropertyTypeUpdateCallback != nil {
		var frontMD_AggregateInformation_PropertyType *MD_AggregateInformation_PropertyType
		if front != nil {
			frontMD_AggregateInformation_PropertyType, _ = front.(*MD_AggregateInformation_PropertyType)
		}
		stage.OnAfterMD_AggregateInformation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_aggregateinformation_propertytype, frontMD_AggregateInformation_PropertyType)
	}
}

func (md_aggregateinformation_propertytype *MD_AggregateInformation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_AggregateInformation_PropertyTypeDeleteCallback != nil {
		var frontMD_AggregateInformation_PropertyType *MD_AggregateInformation_PropertyType
		if front != nil {
			frontMD_AggregateInformation_PropertyType, _ = front.(*MD_AggregateInformation_PropertyType)
		}
		stage.OnAfterMD_AggregateInformation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_aggregateinformation_propertytype, frontMD_AggregateInformation_PropertyType)
	}
}

func (md_aggregateinformation_type *MD_AggregateInformation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_AggregateInformation_TypeCreateCallback != nil {
		stage.OnAfterMD_AggregateInformation_TypeCreateCallback.OnAfterCreate(stage, md_aggregateinformation_type)
	}
}

func (md_aggregateinformation_type *MD_AggregateInformation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_AggregateInformation_TypeUpdateCallback != nil {
		var frontMD_AggregateInformation_Type *MD_AggregateInformation_Type
		if front != nil {
			frontMD_AggregateInformation_Type, _ = front.(*MD_AggregateInformation_Type)
		}
		stage.OnAfterMD_AggregateInformation_TypeUpdateCallback.OnAfterUpdate(stage, md_aggregateinformation_type, frontMD_AggregateInformation_Type)
	}
}

func (md_aggregateinformation_type *MD_AggregateInformation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_AggregateInformation_TypeDeleteCallback != nil {
		var frontMD_AggregateInformation_Type *MD_AggregateInformation_Type
		if front != nil {
			frontMD_AggregateInformation_Type, _ = front.(*MD_AggregateInformation_Type)
		}
		stage.OnAfterMD_AggregateInformation_TypeDeleteCallback.OnAfterDelete(stage, md_aggregateinformation_type, frontMD_AggregateInformation_Type)
	}
}

func (md_browsegraphic_propertytype *MD_BrowseGraphic_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_BrowseGraphic_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_BrowseGraphic_PropertyTypeCreateCallback.OnAfterCreate(stage, md_browsegraphic_propertytype)
	}
}

func (md_browsegraphic_propertytype *MD_BrowseGraphic_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_BrowseGraphic_PropertyTypeUpdateCallback != nil {
		var frontMD_BrowseGraphic_PropertyType *MD_BrowseGraphic_PropertyType
		if front != nil {
			frontMD_BrowseGraphic_PropertyType, _ = front.(*MD_BrowseGraphic_PropertyType)
		}
		stage.OnAfterMD_BrowseGraphic_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_browsegraphic_propertytype, frontMD_BrowseGraphic_PropertyType)
	}
}

func (md_browsegraphic_propertytype *MD_BrowseGraphic_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_BrowseGraphic_PropertyTypeDeleteCallback != nil {
		var frontMD_BrowseGraphic_PropertyType *MD_BrowseGraphic_PropertyType
		if front != nil {
			frontMD_BrowseGraphic_PropertyType, _ = front.(*MD_BrowseGraphic_PropertyType)
		}
		stage.OnAfterMD_BrowseGraphic_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_browsegraphic_propertytype, frontMD_BrowseGraphic_PropertyType)
	}
}

func (md_browsegraphic_type *MD_BrowseGraphic_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_BrowseGraphic_TypeCreateCallback != nil {
		stage.OnAfterMD_BrowseGraphic_TypeCreateCallback.OnAfterCreate(stage, md_browsegraphic_type)
	}
}

func (md_browsegraphic_type *MD_BrowseGraphic_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_BrowseGraphic_TypeUpdateCallback != nil {
		var frontMD_BrowseGraphic_Type *MD_BrowseGraphic_Type
		if front != nil {
			frontMD_BrowseGraphic_Type, _ = front.(*MD_BrowseGraphic_Type)
		}
		stage.OnAfterMD_BrowseGraphic_TypeUpdateCallback.OnAfterUpdate(stage, md_browsegraphic_type, frontMD_BrowseGraphic_Type)
	}
}

func (md_browsegraphic_type *MD_BrowseGraphic_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_BrowseGraphic_TypeDeleteCallback != nil {
		var frontMD_BrowseGraphic_Type *MD_BrowseGraphic_Type
		if front != nil {
			frontMD_BrowseGraphic_Type, _ = front.(*MD_BrowseGraphic_Type)
		}
		stage.OnAfterMD_BrowseGraphic_TypeDeleteCallback.OnAfterDelete(stage, md_browsegraphic_type, frontMD_BrowseGraphic_Type)
	}
}

func (md_charactersetcode_propertytype *MD_CharacterSetCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_CharacterSetCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_CharacterSetCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_charactersetcode_propertytype)
	}
}

func (md_charactersetcode_propertytype *MD_CharacterSetCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CharacterSetCode_PropertyTypeUpdateCallback != nil {
		var frontMD_CharacterSetCode_PropertyType *MD_CharacterSetCode_PropertyType
		if front != nil {
			frontMD_CharacterSetCode_PropertyType, _ = front.(*MD_CharacterSetCode_PropertyType)
		}
		stage.OnAfterMD_CharacterSetCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_charactersetcode_propertytype, frontMD_CharacterSetCode_PropertyType)
	}
}

func (md_charactersetcode_propertytype *MD_CharacterSetCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_CharacterSetCode_PropertyTypeDeleteCallback != nil {
		var frontMD_CharacterSetCode_PropertyType *MD_CharacterSetCode_PropertyType
		if front != nil {
			frontMD_CharacterSetCode_PropertyType, _ = front.(*MD_CharacterSetCode_PropertyType)
		}
		stage.OnAfterMD_CharacterSetCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_charactersetcode_propertytype, frontMD_CharacterSetCode_PropertyType)
	}
}

func (md_classificationcode_propertytype *MD_ClassificationCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ClassificationCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ClassificationCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_classificationcode_propertytype)
	}
}

func (md_classificationcode_propertytype *MD_ClassificationCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ClassificationCode_PropertyTypeUpdateCallback != nil {
		var frontMD_ClassificationCode_PropertyType *MD_ClassificationCode_PropertyType
		if front != nil {
			frontMD_ClassificationCode_PropertyType, _ = front.(*MD_ClassificationCode_PropertyType)
		}
		stage.OnAfterMD_ClassificationCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_classificationcode_propertytype, frontMD_ClassificationCode_PropertyType)
	}
}

func (md_classificationcode_propertytype *MD_ClassificationCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ClassificationCode_PropertyTypeDeleteCallback != nil {
		var frontMD_ClassificationCode_PropertyType *MD_ClassificationCode_PropertyType
		if front != nil {
			frontMD_ClassificationCode_PropertyType, _ = front.(*MD_ClassificationCode_PropertyType)
		}
		stage.OnAfterMD_ClassificationCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_classificationcode_propertytype, frontMD_ClassificationCode_PropertyType)
	}
}

func (md_constraints_propertytype *MD_Constraints_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Constraints_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Constraints_PropertyTypeCreateCallback.OnAfterCreate(stage, md_constraints_propertytype)
	}
}

func (md_constraints_propertytype *MD_Constraints_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Constraints_PropertyTypeUpdateCallback != nil {
		var frontMD_Constraints_PropertyType *MD_Constraints_PropertyType
		if front != nil {
			frontMD_Constraints_PropertyType, _ = front.(*MD_Constraints_PropertyType)
		}
		stage.OnAfterMD_Constraints_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_constraints_propertytype, frontMD_Constraints_PropertyType)
	}
}

func (md_constraints_propertytype *MD_Constraints_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Constraints_PropertyTypeDeleteCallback != nil {
		var frontMD_Constraints_PropertyType *MD_Constraints_PropertyType
		if front != nil {
			frontMD_Constraints_PropertyType, _ = front.(*MD_Constraints_PropertyType)
		}
		stage.OnAfterMD_Constraints_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_constraints_propertytype, frontMD_Constraints_PropertyType)
	}
}

func (md_constraints_type *MD_Constraints_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Constraints_TypeCreateCallback != nil {
		stage.OnAfterMD_Constraints_TypeCreateCallback.OnAfterCreate(stage, md_constraints_type)
	}
}

func (md_constraints_type *MD_Constraints_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Constraints_TypeUpdateCallback != nil {
		var frontMD_Constraints_Type *MD_Constraints_Type
		if front != nil {
			frontMD_Constraints_Type, _ = front.(*MD_Constraints_Type)
		}
		stage.OnAfterMD_Constraints_TypeUpdateCallback.OnAfterUpdate(stage, md_constraints_type, frontMD_Constraints_Type)
	}
}

func (md_constraints_type *MD_Constraints_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Constraints_TypeDeleteCallback != nil {
		var frontMD_Constraints_Type *MD_Constraints_Type
		if front != nil {
			frontMD_Constraints_Type, _ = front.(*MD_Constraints_Type)
		}
		stage.OnAfterMD_Constraints_TypeDeleteCallback.OnAfterDelete(stage, md_constraints_type, frontMD_Constraints_Type)
	}
}

func (md_dataidentification_propertytype *MD_DataIdentification_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_DataIdentification_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_DataIdentification_PropertyTypeCreateCallback.OnAfterCreate(stage, md_dataidentification_propertytype)
	}
}

func (md_dataidentification_propertytype *MD_DataIdentification_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DataIdentification_PropertyTypeUpdateCallback != nil {
		var frontMD_DataIdentification_PropertyType *MD_DataIdentification_PropertyType
		if front != nil {
			frontMD_DataIdentification_PropertyType, _ = front.(*MD_DataIdentification_PropertyType)
		}
		stage.OnAfterMD_DataIdentification_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_dataidentification_propertytype, frontMD_DataIdentification_PropertyType)
	}
}

func (md_dataidentification_propertytype *MD_DataIdentification_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DataIdentification_PropertyTypeDeleteCallback != nil {
		var frontMD_DataIdentification_PropertyType *MD_DataIdentification_PropertyType
		if front != nil {
			frontMD_DataIdentification_PropertyType, _ = front.(*MD_DataIdentification_PropertyType)
		}
		stage.OnAfterMD_DataIdentification_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_dataidentification_propertytype, frontMD_DataIdentification_PropertyType)
	}
}

func (md_dataidentification_type *MD_DataIdentification_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_DataIdentification_TypeCreateCallback != nil {
		stage.OnAfterMD_DataIdentification_TypeCreateCallback.OnAfterCreate(stage, md_dataidentification_type)
	}
}

func (md_dataidentification_type *MD_DataIdentification_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DataIdentification_TypeUpdateCallback != nil {
		var frontMD_DataIdentification_Type *MD_DataIdentification_Type
		if front != nil {
			frontMD_DataIdentification_Type, _ = front.(*MD_DataIdentification_Type)
		}
		stage.OnAfterMD_DataIdentification_TypeUpdateCallback.OnAfterUpdate(stage, md_dataidentification_type, frontMD_DataIdentification_Type)
	}
}

func (md_dataidentification_type *MD_DataIdentification_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DataIdentification_TypeDeleteCallback != nil {
		var frontMD_DataIdentification_Type *MD_DataIdentification_Type
		if front != nil {
			frontMD_DataIdentification_Type, _ = front.(*MD_DataIdentification_Type)
		}
		stage.OnAfterMD_DataIdentification_TypeDeleteCallback.OnAfterDelete(stage, md_dataidentification_type, frontMD_DataIdentification_Type)
	}
}

func (md_digitaltransferoptions_propertytype *MD_DigitalTransferOptions_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_DigitalTransferOptions_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_DigitalTransferOptions_PropertyTypeCreateCallback.OnAfterCreate(stage, md_digitaltransferoptions_propertytype)
	}
}

func (md_digitaltransferoptions_propertytype *MD_DigitalTransferOptions_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DigitalTransferOptions_PropertyTypeUpdateCallback != nil {
		var frontMD_DigitalTransferOptions_PropertyType *MD_DigitalTransferOptions_PropertyType
		if front != nil {
			frontMD_DigitalTransferOptions_PropertyType, _ = front.(*MD_DigitalTransferOptions_PropertyType)
		}
		stage.OnAfterMD_DigitalTransferOptions_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_digitaltransferoptions_propertytype, frontMD_DigitalTransferOptions_PropertyType)
	}
}

func (md_digitaltransferoptions_propertytype *MD_DigitalTransferOptions_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DigitalTransferOptions_PropertyTypeDeleteCallback != nil {
		var frontMD_DigitalTransferOptions_PropertyType *MD_DigitalTransferOptions_PropertyType
		if front != nil {
			frontMD_DigitalTransferOptions_PropertyType, _ = front.(*MD_DigitalTransferOptions_PropertyType)
		}
		stage.OnAfterMD_DigitalTransferOptions_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_digitaltransferoptions_propertytype, frontMD_DigitalTransferOptions_PropertyType)
	}
}

func (md_digitaltransferoptions_type *MD_DigitalTransferOptions_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_DigitalTransferOptions_TypeCreateCallback != nil {
		stage.OnAfterMD_DigitalTransferOptions_TypeCreateCallback.OnAfterCreate(stage, md_digitaltransferoptions_type)
	}
}

func (md_digitaltransferoptions_type *MD_DigitalTransferOptions_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DigitalTransferOptions_TypeUpdateCallback != nil {
		var frontMD_DigitalTransferOptions_Type *MD_DigitalTransferOptions_Type
		if front != nil {
			frontMD_DigitalTransferOptions_Type, _ = front.(*MD_DigitalTransferOptions_Type)
		}
		stage.OnAfterMD_DigitalTransferOptions_TypeUpdateCallback.OnAfterUpdate(stage, md_digitaltransferoptions_type, frontMD_DigitalTransferOptions_Type)
	}
}

func (md_digitaltransferoptions_type *MD_DigitalTransferOptions_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DigitalTransferOptions_TypeDeleteCallback != nil {
		var frontMD_DigitalTransferOptions_Type *MD_DigitalTransferOptions_Type
		if front != nil {
			frontMD_DigitalTransferOptions_Type, _ = front.(*MD_DigitalTransferOptions_Type)
		}
		stage.OnAfterMD_DigitalTransferOptions_TypeDeleteCallback.OnAfterDelete(stage, md_digitaltransferoptions_type, frontMD_DigitalTransferOptions_Type)
	}
}

func (md_distributionunits_propertytype *MD_DistributionUnits_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_DistributionUnits_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_DistributionUnits_PropertyTypeCreateCallback.OnAfterCreate(stage, md_distributionunits_propertytype)
	}
}

func (md_distributionunits_propertytype *MD_DistributionUnits_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DistributionUnits_PropertyTypeUpdateCallback != nil {
		var frontMD_DistributionUnits_PropertyType *MD_DistributionUnits_PropertyType
		if front != nil {
			frontMD_DistributionUnits_PropertyType, _ = front.(*MD_DistributionUnits_PropertyType)
		}
		stage.OnAfterMD_DistributionUnits_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_distributionunits_propertytype, frontMD_DistributionUnits_PropertyType)
	}
}

func (md_distributionunits_propertytype *MD_DistributionUnits_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DistributionUnits_PropertyTypeDeleteCallback != nil {
		var frontMD_DistributionUnits_PropertyType *MD_DistributionUnits_PropertyType
		if front != nil {
			frontMD_DistributionUnits_PropertyType, _ = front.(*MD_DistributionUnits_PropertyType)
		}
		stage.OnAfterMD_DistributionUnits_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_distributionunits_propertytype, frontMD_DistributionUnits_PropertyType)
	}
}

func (md_distribution_propertytype *MD_Distribution_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Distribution_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Distribution_PropertyTypeCreateCallback.OnAfterCreate(stage, md_distribution_propertytype)
	}
}

func (md_distribution_propertytype *MD_Distribution_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distribution_PropertyTypeUpdateCallback != nil {
		var frontMD_Distribution_PropertyType *MD_Distribution_PropertyType
		if front != nil {
			frontMD_Distribution_PropertyType, _ = front.(*MD_Distribution_PropertyType)
		}
		stage.OnAfterMD_Distribution_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_distribution_propertytype, frontMD_Distribution_PropertyType)
	}
}

func (md_distribution_propertytype *MD_Distribution_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distribution_PropertyTypeDeleteCallback != nil {
		var frontMD_Distribution_PropertyType *MD_Distribution_PropertyType
		if front != nil {
			frontMD_Distribution_PropertyType, _ = front.(*MD_Distribution_PropertyType)
		}
		stage.OnAfterMD_Distribution_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_distribution_propertytype, frontMD_Distribution_PropertyType)
	}
}

func (md_distribution_type *MD_Distribution_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Distribution_TypeCreateCallback != nil {
		stage.OnAfterMD_Distribution_TypeCreateCallback.OnAfterCreate(stage, md_distribution_type)
	}
}

func (md_distribution_type *MD_Distribution_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distribution_TypeUpdateCallback != nil {
		var frontMD_Distribution_Type *MD_Distribution_Type
		if front != nil {
			frontMD_Distribution_Type, _ = front.(*MD_Distribution_Type)
		}
		stage.OnAfterMD_Distribution_TypeUpdateCallback.OnAfterUpdate(stage, md_distribution_type, frontMD_Distribution_Type)
	}
}

func (md_distribution_type *MD_Distribution_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distribution_TypeDeleteCallback != nil {
		var frontMD_Distribution_Type *MD_Distribution_Type
		if front != nil {
			frontMD_Distribution_Type, _ = front.(*MD_Distribution_Type)
		}
		stage.OnAfterMD_Distribution_TypeDeleteCallback.OnAfterDelete(stage, md_distribution_type, frontMD_Distribution_Type)
	}
}

func (md_distributor_propertytype *MD_Distributor_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Distributor_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Distributor_PropertyTypeCreateCallback.OnAfterCreate(stage, md_distributor_propertytype)
	}
}

func (md_distributor_propertytype *MD_Distributor_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distributor_PropertyTypeUpdateCallback != nil {
		var frontMD_Distributor_PropertyType *MD_Distributor_PropertyType
		if front != nil {
			frontMD_Distributor_PropertyType, _ = front.(*MD_Distributor_PropertyType)
		}
		stage.OnAfterMD_Distributor_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_distributor_propertytype, frontMD_Distributor_PropertyType)
	}
}

func (md_distributor_propertytype *MD_Distributor_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distributor_PropertyTypeDeleteCallback != nil {
		var frontMD_Distributor_PropertyType *MD_Distributor_PropertyType
		if front != nil {
			frontMD_Distributor_PropertyType, _ = front.(*MD_Distributor_PropertyType)
		}
		stage.OnAfterMD_Distributor_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_distributor_propertytype, frontMD_Distributor_PropertyType)
	}
}

func (md_distributor_type *MD_Distributor_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Distributor_TypeCreateCallback != nil {
		stage.OnAfterMD_Distributor_TypeCreateCallback.OnAfterCreate(stage, md_distributor_type)
	}
}

func (md_distributor_type *MD_Distributor_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distributor_TypeUpdateCallback != nil {
		var frontMD_Distributor_Type *MD_Distributor_Type
		if front != nil {
			frontMD_Distributor_Type, _ = front.(*MD_Distributor_Type)
		}
		stage.OnAfterMD_Distributor_TypeUpdateCallback.OnAfterUpdate(stage, md_distributor_type, frontMD_Distributor_Type)
	}
}

func (md_distributor_type *MD_Distributor_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Distributor_TypeDeleteCallback != nil {
		var frontMD_Distributor_Type *MD_Distributor_Type
		if front != nil {
			frontMD_Distributor_Type, _ = front.(*MD_Distributor_Type)
		}
		stage.OnAfterMD_Distributor_TypeDeleteCallback.OnAfterDelete(stage, md_distributor_type, frontMD_Distributor_Type)
	}
}

func (md_format_propertytype *MD_Format_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Format_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Format_PropertyTypeCreateCallback.OnAfterCreate(stage, md_format_propertytype)
	}
}

func (md_format_propertytype *MD_Format_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Format_PropertyTypeUpdateCallback != nil {
		var frontMD_Format_PropertyType *MD_Format_PropertyType
		if front != nil {
			frontMD_Format_PropertyType, _ = front.(*MD_Format_PropertyType)
		}
		stage.OnAfterMD_Format_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_format_propertytype, frontMD_Format_PropertyType)
	}
}

func (md_format_propertytype *MD_Format_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Format_PropertyTypeDeleteCallback != nil {
		var frontMD_Format_PropertyType *MD_Format_PropertyType
		if front != nil {
			frontMD_Format_PropertyType, _ = front.(*MD_Format_PropertyType)
		}
		stage.OnAfterMD_Format_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_format_propertytype, frontMD_Format_PropertyType)
	}
}

func (md_format_type *MD_Format_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Format_TypeCreateCallback != nil {
		stage.OnAfterMD_Format_TypeCreateCallback.OnAfterCreate(stage, md_format_type)
	}
}

func (md_format_type *MD_Format_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Format_TypeUpdateCallback != nil {
		var frontMD_Format_Type *MD_Format_Type
		if front != nil {
			frontMD_Format_Type, _ = front.(*MD_Format_Type)
		}
		stage.OnAfterMD_Format_TypeUpdateCallback.OnAfterUpdate(stage, md_format_type, frontMD_Format_Type)
	}
}

func (md_format_type *MD_Format_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Format_TypeDeleteCallback != nil {
		var frontMD_Format_Type *MD_Format_Type
		if front != nil {
			frontMD_Format_Type, _ = front.(*MD_Format_Type)
		}
		stage.OnAfterMD_Format_TypeDeleteCallback.OnAfterDelete(stage, md_format_type, frontMD_Format_Type)
	}
}

func (md_identification_propertytype *MD_Identification_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Identification_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Identification_PropertyTypeCreateCallback.OnAfterCreate(stage, md_identification_propertytype)
	}
}

func (md_identification_propertytype *MD_Identification_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Identification_PropertyTypeUpdateCallback != nil {
		var frontMD_Identification_PropertyType *MD_Identification_PropertyType
		if front != nil {
			frontMD_Identification_PropertyType, _ = front.(*MD_Identification_PropertyType)
		}
		stage.OnAfterMD_Identification_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_identification_propertytype, frontMD_Identification_PropertyType)
	}
}

func (md_identification_propertytype *MD_Identification_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Identification_PropertyTypeDeleteCallback != nil {
		var frontMD_Identification_PropertyType *MD_Identification_PropertyType
		if front != nil {
			frontMD_Identification_PropertyType, _ = front.(*MD_Identification_PropertyType)
		}
		stage.OnAfterMD_Identification_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_identification_propertytype, frontMD_Identification_PropertyType)
	}
}

func (md_identifier_propertytype *MD_Identifier_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Identifier_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Identifier_PropertyTypeCreateCallback.OnAfterCreate(stage, md_identifier_propertytype)
	}
}

func (md_identifier_propertytype *MD_Identifier_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Identifier_PropertyTypeUpdateCallback != nil {
		var frontMD_Identifier_PropertyType *MD_Identifier_PropertyType
		if front != nil {
			frontMD_Identifier_PropertyType, _ = front.(*MD_Identifier_PropertyType)
		}
		stage.OnAfterMD_Identifier_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_identifier_propertytype, frontMD_Identifier_PropertyType)
	}
}

func (md_identifier_propertytype *MD_Identifier_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Identifier_PropertyTypeDeleteCallback != nil {
		var frontMD_Identifier_PropertyType *MD_Identifier_PropertyType
		if front != nil {
			frontMD_Identifier_PropertyType, _ = front.(*MD_Identifier_PropertyType)
		}
		stage.OnAfterMD_Identifier_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_identifier_propertytype, frontMD_Identifier_PropertyType)
	}
}

func (md_identifier_type *MD_Identifier_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Identifier_TypeCreateCallback != nil {
		stage.OnAfterMD_Identifier_TypeCreateCallback.OnAfterCreate(stage, md_identifier_type)
	}
}

func (md_identifier_type *MD_Identifier_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Identifier_TypeUpdateCallback != nil {
		var frontMD_Identifier_Type *MD_Identifier_Type
		if front != nil {
			frontMD_Identifier_Type, _ = front.(*MD_Identifier_Type)
		}
		stage.OnAfterMD_Identifier_TypeUpdateCallback.OnAfterUpdate(stage, md_identifier_type, frontMD_Identifier_Type)
	}
}

func (md_identifier_type *MD_Identifier_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Identifier_TypeDeleteCallback != nil {
		var frontMD_Identifier_Type *MD_Identifier_Type
		if front != nil {
			frontMD_Identifier_Type, _ = front.(*MD_Identifier_Type)
		}
		stage.OnAfterMD_Identifier_TypeDeleteCallback.OnAfterDelete(stage, md_identifier_type, frontMD_Identifier_Type)
	}
}

func (md_keywordtypecode_propertytype *MD_KeywordTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_KeywordTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_KeywordTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_keywordtypecode_propertytype)
	}
}

func (md_keywordtypecode_propertytype *MD_KeywordTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_KeywordTypeCode_PropertyTypeUpdateCallback != nil {
		var frontMD_KeywordTypeCode_PropertyType *MD_KeywordTypeCode_PropertyType
		if front != nil {
			frontMD_KeywordTypeCode_PropertyType, _ = front.(*MD_KeywordTypeCode_PropertyType)
		}
		stage.OnAfterMD_KeywordTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_keywordtypecode_propertytype, frontMD_KeywordTypeCode_PropertyType)
	}
}

func (md_keywordtypecode_propertytype *MD_KeywordTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_KeywordTypeCode_PropertyTypeDeleteCallback != nil {
		var frontMD_KeywordTypeCode_PropertyType *MD_KeywordTypeCode_PropertyType
		if front != nil {
			frontMD_KeywordTypeCode_PropertyType, _ = front.(*MD_KeywordTypeCode_PropertyType)
		}
		stage.OnAfterMD_KeywordTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_keywordtypecode_propertytype, frontMD_KeywordTypeCode_PropertyType)
	}
}

func (md_keywords_propertytype *MD_Keywords_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Keywords_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Keywords_PropertyTypeCreateCallback.OnAfterCreate(stage, md_keywords_propertytype)
	}
}

func (md_keywords_propertytype *MD_Keywords_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Keywords_PropertyTypeUpdateCallback != nil {
		var frontMD_Keywords_PropertyType *MD_Keywords_PropertyType
		if front != nil {
			frontMD_Keywords_PropertyType, _ = front.(*MD_Keywords_PropertyType)
		}
		stage.OnAfterMD_Keywords_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_keywords_propertytype, frontMD_Keywords_PropertyType)
	}
}

func (md_keywords_propertytype *MD_Keywords_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Keywords_PropertyTypeDeleteCallback != nil {
		var frontMD_Keywords_PropertyType *MD_Keywords_PropertyType
		if front != nil {
			frontMD_Keywords_PropertyType, _ = front.(*MD_Keywords_PropertyType)
		}
		stage.OnAfterMD_Keywords_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_keywords_propertytype, frontMD_Keywords_PropertyType)
	}
}

func (md_keywords_type *MD_Keywords_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Keywords_TypeCreateCallback != nil {
		stage.OnAfterMD_Keywords_TypeCreateCallback.OnAfterCreate(stage, md_keywords_type)
	}
}

func (md_keywords_type *MD_Keywords_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Keywords_TypeUpdateCallback != nil {
		var frontMD_Keywords_Type *MD_Keywords_Type
		if front != nil {
			frontMD_Keywords_Type, _ = front.(*MD_Keywords_Type)
		}
		stage.OnAfterMD_Keywords_TypeUpdateCallback.OnAfterUpdate(stage, md_keywords_type, frontMD_Keywords_Type)
	}
}

func (md_keywords_type *MD_Keywords_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Keywords_TypeDeleteCallback != nil {
		var frontMD_Keywords_Type *MD_Keywords_Type
		if front != nil {
			frontMD_Keywords_Type, _ = front.(*MD_Keywords_Type)
		}
		stage.OnAfterMD_Keywords_TypeDeleteCallback.OnAfterDelete(stage, md_keywords_type, frontMD_Keywords_Type)
	}
}

func (md_legalconstraints_propertytype *MD_LegalConstraints_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_LegalConstraints_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_LegalConstraints_PropertyTypeCreateCallback.OnAfterCreate(stage, md_legalconstraints_propertytype)
	}
}

func (md_legalconstraints_propertytype *MD_LegalConstraints_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_LegalConstraints_PropertyTypeUpdateCallback != nil {
		var frontMD_LegalConstraints_PropertyType *MD_LegalConstraints_PropertyType
		if front != nil {
			frontMD_LegalConstraints_PropertyType, _ = front.(*MD_LegalConstraints_PropertyType)
		}
		stage.OnAfterMD_LegalConstraints_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_legalconstraints_propertytype, frontMD_LegalConstraints_PropertyType)
	}
}

func (md_legalconstraints_propertytype *MD_LegalConstraints_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_LegalConstraints_PropertyTypeDeleteCallback != nil {
		var frontMD_LegalConstraints_PropertyType *MD_LegalConstraints_PropertyType
		if front != nil {
			frontMD_LegalConstraints_PropertyType, _ = front.(*MD_LegalConstraints_PropertyType)
		}
		stage.OnAfterMD_LegalConstraints_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_legalconstraints_propertytype, frontMD_LegalConstraints_PropertyType)
	}
}

func (md_legalconstraints_type *MD_LegalConstraints_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_LegalConstraints_TypeCreateCallback != nil {
		stage.OnAfterMD_LegalConstraints_TypeCreateCallback.OnAfterCreate(stage, md_legalconstraints_type)
	}
}

func (md_legalconstraints_type *MD_LegalConstraints_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_LegalConstraints_TypeUpdateCallback != nil {
		var frontMD_LegalConstraints_Type *MD_LegalConstraints_Type
		if front != nil {
			frontMD_LegalConstraints_Type, _ = front.(*MD_LegalConstraints_Type)
		}
		stage.OnAfterMD_LegalConstraints_TypeUpdateCallback.OnAfterUpdate(stage, md_legalconstraints_type, frontMD_LegalConstraints_Type)
	}
}

func (md_legalconstraints_type *MD_LegalConstraints_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_LegalConstraints_TypeDeleteCallback != nil {
		var frontMD_LegalConstraints_Type *MD_LegalConstraints_Type
		if front != nil {
			frontMD_LegalConstraints_Type, _ = front.(*MD_LegalConstraints_Type)
		}
		stage.OnAfterMD_LegalConstraints_TypeDeleteCallback.OnAfterDelete(stage, md_legalconstraints_type, frontMD_LegalConstraints_Type)
	}
}

func (md_maintenancefrequencycode_propertytype *MD_MaintenanceFrequencyCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_MaintenanceFrequencyCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_MaintenanceFrequencyCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_maintenancefrequencycode_propertytype)
	}
}

func (md_maintenancefrequencycode_propertytype *MD_MaintenanceFrequencyCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MaintenanceFrequencyCode_PropertyTypeUpdateCallback != nil {
		var frontMD_MaintenanceFrequencyCode_PropertyType *MD_MaintenanceFrequencyCode_PropertyType
		if front != nil {
			frontMD_MaintenanceFrequencyCode_PropertyType, _ = front.(*MD_MaintenanceFrequencyCode_PropertyType)
		}
		stage.OnAfterMD_MaintenanceFrequencyCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_maintenancefrequencycode_propertytype, frontMD_MaintenanceFrequencyCode_PropertyType)
	}
}

func (md_maintenancefrequencycode_propertytype *MD_MaintenanceFrequencyCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MaintenanceFrequencyCode_PropertyTypeDeleteCallback != nil {
		var frontMD_MaintenanceFrequencyCode_PropertyType *MD_MaintenanceFrequencyCode_PropertyType
		if front != nil {
			frontMD_MaintenanceFrequencyCode_PropertyType, _ = front.(*MD_MaintenanceFrequencyCode_PropertyType)
		}
		stage.OnAfterMD_MaintenanceFrequencyCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_maintenancefrequencycode_propertytype, frontMD_MaintenanceFrequencyCode_PropertyType)
	}
}

func (md_maintenanceinformation_propertytype *MD_MaintenanceInformation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_MaintenanceInformation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_MaintenanceInformation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_maintenanceinformation_propertytype)
	}
}

func (md_maintenanceinformation_propertytype *MD_MaintenanceInformation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MaintenanceInformation_PropertyTypeUpdateCallback != nil {
		var frontMD_MaintenanceInformation_PropertyType *MD_MaintenanceInformation_PropertyType
		if front != nil {
			frontMD_MaintenanceInformation_PropertyType, _ = front.(*MD_MaintenanceInformation_PropertyType)
		}
		stage.OnAfterMD_MaintenanceInformation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_maintenanceinformation_propertytype, frontMD_MaintenanceInformation_PropertyType)
	}
}

func (md_maintenanceinformation_propertytype *MD_MaintenanceInformation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MaintenanceInformation_PropertyTypeDeleteCallback != nil {
		var frontMD_MaintenanceInformation_PropertyType *MD_MaintenanceInformation_PropertyType
		if front != nil {
			frontMD_MaintenanceInformation_PropertyType, _ = front.(*MD_MaintenanceInformation_PropertyType)
		}
		stage.OnAfterMD_MaintenanceInformation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_maintenanceinformation_propertytype, frontMD_MaintenanceInformation_PropertyType)
	}
}

func (md_maintenanceinformation_type *MD_MaintenanceInformation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_MaintenanceInformation_TypeCreateCallback != nil {
		stage.OnAfterMD_MaintenanceInformation_TypeCreateCallback.OnAfterCreate(stage, md_maintenanceinformation_type)
	}
}

func (md_maintenanceinformation_type *MD_MaintenanceInformation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MaintenanceInformation_TypeUpdateCallback != nil {
		var frontMD_MaintenanceInformation_Type *MD_MaintenanceInformation_Type
		if front != nil {
			frontMD_MaintenanceInformation_Type, _ = front.(*MD_MaintenanceInformation_Type)
		}
		stage.OnAfterMD_MaintenanceInformation_TypeUpdateCallback.OnAfterUpdate(stage, md_maintenanceinformation_type, frontMD_MaintenanceInformation_Type)
	}
}

func (md_maintenanceinformation_type *MD_MaintenanceInformation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MaintenanceInformation_TypeDeleteCallback != nil {
		var frontMD_MaintenanceInformation_Type *MD_MaintenanceInformation_Type
		if front != nil {
			frontMD_MaintenanceInformation_Type, _ = front.(*MD_MaintenanceInformation_Type)
		}
		stage.OnAfterMD_MaintenanceInformation_TypeDeleteCallback.OnAfterDelete(stage, md_maintenanceinformation_type, frontMD_MaintenanceInformation_Type)
	}
}

func (md_mediumformatcode_propertytype *MD_MediumFormatCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_MediumFormatCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_MediumFormatCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_mediumformatcode_propertytype)
	}
}

func (md_mediumformatcode_propertytype *MD_MediumFormatCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MediumFormatCode_PropertyTypeUpdateCallback != nil {
		var frontMD_MediumFormatCode_PropertyType *MD_MediumFormatCode_PropertyType
		if front != nil {
			frontMD_MediumFormatCode_PropertyType, _ = front.(*MD_MediumFormatCode_PropertyType)
		}
		stage.OnAfterMD_MediumFormatCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_mediumformatcode_propertytype, frontMD_MediumFormatCode_PropertyType)
	}
}

func (md_mediumformatcode_propertytype *MD_MediumFormatCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MediumFormatCode_PropertyTypeDeleteCallback != nil {
		var frontMD_MediumFormatCode_PropertyType *MD_MediumFormatCode_PropertyType
		if front != nil {
			frontMD_MediumFormatCode_PropertyType, _ = front.(*MD_MediumFormatCode_PropertyType)
		}
		stage.OnAfterMD_MediumFormatCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_mediumformatcode_propertytype, frontMD_MediumFormatCode_PropertyType)
	}
}

func (md_mediumnamecode_propertytype *MD_MediumNameCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_MediumNameCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_MediumNameCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_mediumnamecode_propertytype)
	}
}

func (md_mediumnamecode_propertytype *MD_MediumNameCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MediumNameCode_PropertyTypeUpdateCallback != nil {
		var frontMD_MediumNameCode_PropertyType *MD_MediumNameCode_PropertyType
		if front != nil {
			frontMD_MediumNameCode_PropertyType, _ = front.(*MD_MediumNameCode_PropertyType)
		}
		stage.OnAfterMD_MediumNameCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_mediumnamecode_propertytype, frontMD_MediumNameCode_PropertyType)
	}
}

func (md_mediumnamecode_propertytype *MD_MediumNameCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MediumNameCode_PropertyTypeDeleteCallback != nil {
		var frontMD_MediumNameCode_PropertyType *MD_MediumNameCode_PropertyType
		if front != nil {
			frontMD_MediumNameCode_PropertyType, _ = front.(*MD_MediumNameCode_PropertyType)
		}
		stage.OnAfterMD_MediumNameCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_mediumnamecode_propertytype, frontMD_MediumNameCode_PropertyType)
	}
}

func (md_medium_propertytype *MD_Medium_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Medium_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Medium_PropertyTypeCreateCallback.OnAfterCreate(stage, md_medium_propertytype)
	}
}

func (md_medium_propertytype *MD_Medium_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Medium_PropertyTypeUpdateCallback != nil {
		var frontMD_Medium_PropertyType *MD_Medium_PropertyType
		if front != nil {
			frontMD_Medium_PropertyType, _ = front.(*MD_Medium_PropertyType)
		}
		stage.OnAfterMD_Medium_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_medium_propertytype, frontMD_Medium_PropertyType)
	}
}

func (md_medium_propertytype *MD_Medium_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Medium_PropertyTypeDeleteCallback != nil {
		var frontMD_Medium_PropertyType *MD_Medium_PropertyType
		if front != nil {
			frontMD_Medium_PropertyType, _ = front.(*MD_Medium_PropertyType)
		}
		stage.OnAfterMD_Medium_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_medium_propertytype, frontMD_Medium_PropertyType)
	}
}

func (md_medium_type *MD_Medium_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Medium_TypeCreateCallback != nil {
		stage.OnAfterMD_Medium_TypeCreateCallback.OnAfterCreate(stage, md_medium_type)
	}
}

func (md_medium_type *MD_Medium_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Medium_TypeUpdateCallback != nil {
		var frontMD_Medium_Type *MD_Medium_Type
		if front != nil {
			frontMD_Medium_Type, _ = front.(*MD_Medium_Type)
		}
		stage.OnAfterMD_Medium_TypeUpdateCallback.OnAfterUpdate(stage, md_medium_type, frontMD_Medium_Type)
	}
}

func (md_medium_type *MD_Medium_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Medium_TypeDeleteCallback != nil {
		var frontMD_Medium_Type *MD_Medium_Type
		if front != nil {
			frontMD_Medium_Type, _ = front.(*MD_Medium_Type)
		}
		stage.OnAfterMD_Medium_TypeDeleteCallback.OnAfterDelete(stage, md_medium_type, frontMD_Medium_Type)
	}
}

func (md_progresscode_propertytype *MD_ProgressCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ProgressCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ProgressCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_progresscode_propertytype)
	}
}

func (md_progresscode_propertytype *MD_ProgressCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ProgressCode_PropertyTypeUpdateCallback != nil {
		var frontMD_ProgressCode_PropertyType *MD_ProgressCode_PropertyType
		if front != nil {
			frontMD_ProgressCode_PropertyType, _ = front.(*MD_ProgressCode_PropertyType)
		}
		stage.OnAfterMD_ProgressCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_progresscode_propertytype, frontMD_ProgressCode_PropertyType)
	}
}

func (md_progresscode_propertytype *MD_ProgressCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ProgressCode_PropertyTypeDeleteCallback != nil {
		var frontMD_ProgressCode_PropertyType *MD_ProgressCode_PropertyType
		if front != nil {
			frontMD_ProgressCode_PropertyType, _ = front.(*MD_ProgressCode_PropertyType)
		}
		stage.OnAfterMD_ProgressCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_progresscode_propertytype, frontMD_ProgressCode_PropertyType)
	}
}

func (md_referencesystem_propertytype *MD_ReferenceSystem_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ReferenceSystem_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ReferenceSystem_PropertyTypeCreateCallback.OnAfterCreate(stage, md_referencesystem_propertytype)
	}
}

func (md_referencesystem_propertytype *MD_ReferenceSystem_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ReferenceSystem_PropertyTypeUpdateCallback != nil {
		var frontMD_ReferenceSystem_PropertyType *MD_ReferenceSystem_PropertyType
		if front != nil {
			frontMD_ReferenceSystem_PropertyType, _ = front.(*MD_ReferenceSystem_PropertyType)
		}
		stage.OnAfterMD_ReferenceSystem_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_referencesystem_propertytype, frontMD_ReferenceSystem_PropertyType)
	}
}

func (md_referencesystem_propertytype *MD_ReferenceSystem_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ReferenceSystem_PropertyTypeDeleteCallback != nil {
		var frontMD_ReferenceSystem_PropertyType *MD_ReferenceSystem_PropertyType
		if front != nil {
			frontMD_ReferenceSystem_PropertyType, _ = front.(*MD_ReferenceSystem_PropertyType)
		}
		stage.OnAfterMD_ReferenceSystem_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_referencesystem_propertytype, frontMD_ReferenceSystem_PropertyType)
	}
}

func (md_referencesystem_type *MD_ReferenceSystem_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ReferenceSystem_TypeCreateCallback != nil {
		stage.OnAfterMD_ReferenceSystem_TypeCreateCallback.OnAfterCreate(stage, md_referencesystem_type)
	}
}

func (md_referencesystem_type *MD_ReferenceSystem_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ReferenceSystem_TypeUpdateCallback != nil {
		var frontMD_ReferenceSystem_Type *MD_ReferenceSystem_Type
		if front != nil {
			frontMD_ReferenceSystem_Type, _ = front.(*MD_ReferenceSystem_Type)
		}
		stage.OnAfterMD_ReferenceSystem_TypeUpdateCallback.OnAfterUpdate(stage, md_referencesystem_type, frontMD_ReferenceSystem_Type)
	}
}

func (md_referencesystem_type *MD_ReferenceSystem_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ReferenceSystem_TypeDeleteCallback != nil {
		var frontMD_ReferenceSystem_Type *MD_ReferenceSystem_Type
		if front != nil {
			frontMD_ReferenceSystem_Type, _ = front.(*MD_ReferenceSystem_Type)
		}
		stage.OnAfterMD_ReferenceSystem_TypeDeleteCallback.OnAfterDelete(stage, md_referencesystem_type, frontMD_ReferenceSystem_Type)
	}
}

func (md_representativefraction_propertytype *MD_RepresentativeFraction_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_RepresentativeFraction_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_RepresentativeFraction_PropertyTypeCreateCallback.OnAfterCreate(stage, md_representativefraction_propertytype)
	}
}

func (md_representativefraction_propertytype *MD_RepresentativeFraction_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RepresentativeFraction_PropertyTypeUpdateCallback != nil {
		var frontMD_RepresentativeFraction_PropertyType *MD_RepresentativeFraction_PropertyType
		if front != nil {
			frontMD_RepresentativeFraction_PropertyType, _ = front.(*MD_RepresentativeFraction_PropertyType)
		}
		stage.OnAfterMD_RepresentativeFraction_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_representativefraction_propertytype, frontMD_RepresentativeFraction_PropertyType)
	}
}

func (md_representativefraction_propertytype *MD_RepresentativeFraction_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RepresentativeFraction_PropertyTypeDeleteCallback != nil {
		var frontMD_RepresentativeFraction_PropertyType *MD_RepresentativeFraction_PropertyType
		if front != nil {
			frontMD_RepresentativeFraction_PropertyType, _ = front.(*MD_RepresentativeFraction_PropertyType)
		}
		stage.OnAfterMD_RepresentativeFraction_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_representativefraction_propertytype, frontMD_RepresentativeFraction_PropertyType)
	}
}

func (md_representativefraction_type *MD_RepresentativeFraction_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_RepresentativeFraction_TypeCreateCallback != nil {
		stage.OnAfterMD_RepresentativeFraction_TypeCreateCallback.OnAfterCreate(stage, md_representativefraction_type)
	}
}

func (md_representativefraction_type *MD_RepresentativeFraction_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RepresentativeFraction_TypeUpdateCallback != nil {
		var frontMD_RepresentativeFraction_Type *MD_RepresentativeFraction_Type
		if front != nil {
			frontMD_RepresentativeFraction_Type, _ = front.(*MD_RepresentativeFraction_Type)
		}
		stage.OnAfterMD_RepresentativeFraction_TypeUpdateCallback.OnAfterUpdate(stage, md_representativefraction_type, frontMD_RepresentativeFraction_Type)
	}
}

func (md_representativefraction_type *MD_RepresentativeFraction_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RepresentativeFraction_TypeDeleteCallback != nil {
		var frontMD_RepresentativeFraction_Type *MD_RepresentativeFraction_Type
		if front != nil {
			frontMD_RepresentativeFraction_Type, _ = front.(*MD_RepresentativeFraction_Type)
		}
		stage.OnAfterMD_RepresentativeFraction_TypeDeleteCallback.OnAfterDelete(stage, md_representativefraction_type, frontMD_RepresentativeFraction_Type)
	}
}

func (md_resolution_propertytype *MD_Resolution_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Resolution_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Resolution_PropertyTypeCreateCallback.OnAfterCreate(stage, md_resolution_propertytype)
	}
}

func (md_resolution_propertytype *MD_Resolution_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Resolution_PropertyTypeUpdateCallback != nil {
		var frontMD_Resolution_PropertyType *MD_Resolution_PropertyType
		if front != nil {
			frontMD_Resolution_PropertyType, _ = front.(*MD_Resolution_PropertyType)
		}
		stage.OnAfterMD_Resolution_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_resolution_propertytype, frontMD_Resolution_PropertyType)
	}
}

func (md_resolution_propertytype *MD_Resolution_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Resolution_PropertyTypeDeleteCallback != nil {
		var frontMD_Resolution_PropertyType *MD_Resolution_PropertyType
		if front != nil {
			frontMD_Resolution_PropertyType, _ = front.(*MD_Resolution_PropertyType)
		}
		stage.OnAfterMD_Resolution_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_resolution_propertytype, frontMD_Resolution_PropertyType)
	}
}

func (md_resolution_type *MD_Resolution_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Resolution_TypeCreateCallback != nil {
		stage.OnAfterMD_Resolution_TypeCreateCallback.OnAfterCreate(stage, md_resolution_type)
	}
}

func (md_resolution_type *MD_Resolution_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Resolution_TypeUpdateCallback != nil {
		var frontMD_Resolution_Type *MD_Resolution_Type
		if front != nil {
			frontMD_Resolution_Type, _ = front.(*MD_Resolution_Type)
		}
		stage.OnAfterMD_Resolution_TypeUpdateCallback.OnAfterUpdate(stage, md_resolution_type, frontMD_Resolution_Type)
	}
}

func (md_resolution_type *MD_Resolution_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Resolution_TypeDeleteCallback != nil {
		var frontMD_Resolution_Type *MD_Resolution_Type
		if front != nil {
			frontMD_Resolution_Type, _ = front.(*MD_Resolution_Type)
		}
		stage.OnAfterMD_Resolution_TypeDeleteCallback.OnAfterDelete(stage, md_resolution_type, frontMD_Resolution_Type)
	}
}

func (md_restrictioncode_propertytype *MD_RestrictionCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_RestrictionCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_RestrictionCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_restrictioncode_propertytype)
	}
}

func (md_restrictioncode_propertytype *MD_RestrictionCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RestrictionCode_PropertyTypeUpdateCallback != nil {
		var frontMD_RestrictionCode_PropertyType *MD_RestrictionCode_PropertyType
		if front != nil {
			frontMD_RestrictionCode_PropertyType, _ = front.(*MD_RestrictionCode_PropertyType)
		}
		stage.OnAfterMD_RestrictionCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_restrictioncode_propertytype, frontMD_RestrictionCode_PropertyType)
	}
}

func (md_restrictioncode_propertytype *MD_RestrictionCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_RestrictionCode_PropertyTypeDeleteCallback != nil {
		var frontMD_RestrictionCode_PropertyType *MD_RestrictionCode_PropertyType
		if front != nil {
			frontMD_RestrictionCode_PropertyType, _ = front.(*MD_RestrictionCode_PropertyType)
		}
		stage.OnAfterMD_RestrictionCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_restrictioncode_propertytype, frontMD_RestrictionCode_PropertyType)
	}
}

func (md_scopecode_propertytype *MD_ScopeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ScopeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ScopeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_scopecode_propertytype)
	}
}

func (md_scopecode_propertytype *MD_ScopeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ScopeCode_PropertyTypeUpdateCallback != nil {
		var frontMD_ScopeCode_PropertyType *MD_ScopeCode_PropertyType
		if front != nil {
			frontMD_ScopeCode_PropertyType, _ = front.(*MD_ScopeCode_PropertyType)
		}
		stage.OnAfterMD_ScopeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_scopecode_propertytype, frontMD_ScopeCode_PropertyType)
	}
}

func (md_scopecode_propertytype *MD_ScopeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ScopeCode_PropertyTypeDeleteCallback != nil {
		var frontMD_ScopeCode_PropertyType *MD_ScopeCode_PropertyType
		if front != nil {
			frontMD_ScopeCode_PropertyType, _ = front.(*MD_ScopeCode_PropertyType)
		}
		stage.OnAfterMD_ScopeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_scopecode_propertytype, frontMD_ScopeCode_PropertyType)
	}
}

func (md_scopedescription_propertytype *MD_ScopeDescription_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ScopeDescription_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ScopeDescription_PropertyTypeCreateCallback.OnAfterCreate(stage, md_scopedescription_propertytype)
	}
}

func (md_scopedescription_propertytype *MD_ScopeDescription_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ScopeDescription_PropertyTypeUpdateCallback != nil {
		var frontMD_ScopeDescription_PropertyType *MD_ScopeDescription_PropertyType
		if front != nil {
			frontMD_ScopeDescription_PropertyType, _ = front.(*MD_ScopeDescription_PropertyType)
		}
		stage.OnAfterMD_ScopeDescription_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_scopedescription_propertytype, frontMD_ScopeDescription_PropertyType)
	}
}

func (md_scopedescription_propertytype *MD_ScopeDescription_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ScopeDescription_PropertyTypeDeleteCallback != nil {
		var frontMD_ScopeDescription_PropertyType *MD_ScopeDescription_PropertyType
		if front != nil {
			frontMD_ScopeDescription_PropertyType, _ = front.(*MD_ScopeDescription_PropertyType)
		}
		stage.OnAfterMD_ScopeDescription_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_scopedescription_propertytype, frontMD_ScopeDescription_PropertyType)
	}
}

func (md_scopedescription_type *MD_ScopeDescription_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ScopeDescription_TypeCreateCallback != nil {
		stage.OnAfterMD_ScopeDescription_TypeCreateCallback.OnAfterCreate(stage, md_scopedescription_type)
	}
}

func (md_scopedescription_type *MD_ScopeDescription_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ScopeDescription_TypeUpdateCallback != nil {
		var frontMD_ScopeDescription_Type *MD_ScopeDescription_Type
		if front != nil {
			frontMD_ScopeDescription_Type, _ = front.(*MD_ScopeDescription_Type)
		}
		stage.OnAfterMD_ScopeDescription_TypeUpdateCallback.OnAfterUpdate(stage, md_scopedescription_type, frontMD_ScopeDescription_Type)
	}
}

func (md_scopedescription_type *MD_ScopeDescription_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ScopeDescription_TypeDeleteCallback != nil {
		var frontMD_ScopeDescription_Type *MD_ScopeDescription_Type
		if front != nil {
			frontMD_ScopeDescription_Type, _ = front.(*MD_ScopeDescription_Type)
		}
		stage.OnAfterMD_ScopeDescription_TypeDeleteCallback.OnAfterDelete(stage, md_scopedescription_type, frontMD_ScopeDescription_Type)
	}
}

func (md_securityconstraints_propertytype *MD_SecurityConstraints_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_SecurityConstraints_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_SecurityConstraints_PropertyTypeCreateCallback.OnAfterCreate(stage, md_securityconstraints_propertytype)
	}
}

func (md_securityconstraints_propertytype *MD_SecurityConstraints_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SecurityConstraints_PropertyTypeUpdateCallback != nil {
		var frontMD_SecurityConstraints_PropertyType *MD_SecurityConstraints_PropertyType
		if front != nil {
			frontMD_SecurityConstraints_PropertyType, _ = front.(*MD_SecurityConstraints_PropertyType)
		}
		stage.OnAfterMD_SecurityConstraints_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_securityconstraints_propertytype, frontMD_SecurityConstraints_PropertyType)
	}
}

func (md_securityconstraints_propertytype *MD_SecurityConstraints_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SecurityConstraints_PropertyTypeDeleteCallback != nil {
		var frontMD_SecurityConstraints_PropertyType *MD_SecurityConstraints_PropertyType
		if front != nil {
			frontMD_SecurityConstraints_PropertyType, _ = front.(*MD_SecurityConstraints_PropertyType)
		}
		stage.OnAfterMD_SecurityConstraints_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_securityconstraints_propertytype, frontMD_SecurityConstraints_PropertyType)
	}
}

func (md_securityconstraints_type *MD_SecurityConstraints_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_SecurityConstraints_TypeCreateCallback != nil {
		stage.OnAfterMD_SecurityConstraints_TypeCreateCallback.OnAfterCreate(stage, md_securityconstraints_type)
	}
}

func (md_securityconstraints_type *MD_SecurityConstraints_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SecurityConstraints_TypeUpdateCallback != nil {
		var frontMD_SecurityConstraints_Type *MD_SecurityConstraints_Type
		if front != nil {
			frontMD_SecurityConstraints_Type, _ = front.(*MD_SecurityConstraints_Type)
		}
		stage.OnAfterMD_SecurityConstraints_TypeUpdateCallback.OnAfterUpdate(stage, md_securityconstraints_type, frontMD_SecurityConstraints_Type)
	}
}

func (md_securityconstraints_type *MD_SecurityConstraints_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SecurityConstraints_TypeDeleteCallback != nil {
		var frontMD_SecurityConstraints_Type *MD_SecurityConstraints_Type
		if front != nil {
			frontMD_SecurityConstraints_Type, _ = front.(*MD_SecurityConstraints_Type)
		}
		stage.OnAfterMD_SecurityConstraints_TypeDeleteCallback.OnAfterDelete(stage, md_securityconstraints_type, frontMD_SecurityConstraints_Type)
	}
}

func (md_serviceidentification_propertytype *MD_ServiceIdentification_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ServiceIdentification_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ServiceIdentification_PropertyTypeCreateCallback.OnAfterCreate(stage, md_serviceidentification_propertytype)
	}
}

func (md_serviceidentification_propertytype *MD_ServiceIdentification_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ServiceIdentification_PropertyTypeUpdateCallback != nil {
		var frontMD_ServiceIdentification_PropertyType *MD_ServiceIdentification_PropertyType
		if front != nil {
			frontMD_ServiceIdentification_PropertyType, _ = front.(*MD_ServiceIdentification_PropertyType)
		}
		stage.OnAfterMD_ServiceIdentification_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_serviceidentification_propertytype, frontMD_ServiceIdentification_PropertyType)
	}
}

func (md_serviceidentification_propertytype *MD_ServiceIdentification_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ServiceIdentification_PropertyTypeDeleteCallback != nil {
		var frontMD_ServiceIdentification_PropertyType *MD_ServiceIdentification_PropertyType
		if front != nil {
			frontMD_ServiceIdentification_PropertyType, _ = front.(*MD_ServiceIdentification_PropertyType)
		}
		stage.OnAfterMD_ServiceIdentification_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_serviceidentification_propertytype, frontMD_ServiceIdentification_PropertyType)
	}
}

func (md_serviceidentification_type *MD_ServiceIdentification_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ServiceIdentification_TypeCreateCallback != nil {
		stage.OnAfterMD_ServiceIdentification_TypeCreateCallback.OnAfterCreate(stage, md_serviceidentification_type)
	}
}

func (md_serviceidentification_type *MD_ServiceIdentification_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ServiceIdentification_TypeUpdateCallback != nil {
		var frontMD_ServiceIdentification_Type *MD_ServiceIdentification_Type
		if front != nil {
			frontMD_ServiceIdentification_Type, _ = front.(*MD_ServiceIdentification_Type)
		}
		stage.OnAfterMD_ServiceIdentification_TypeUpdateCallback.OnAfterUpdate(stage, md_serviceidentification_type, frontMD_ServiceIdentification_Type)
	}
}

func (md_serviceidentification_type *MD_ServiceIdentification_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ServiceIdentification_TypeDeleteCallback != nil {
		var frontMD_ServiceIdentification_Type *MD_ServiceIdentification_Type
		if front != nil {
			frontMD_ServiceIdentification_Type, _ = front.(*MD_ServiceIdentification_Type)
		}
		stage.OnAfterMD_ServiceIdentification_TypeDeleteCallback.OnAfterDelete(stage, md_serviceidentification_type, frontMD_ServiceIdentification_Type)
	}
}

func (md_spatialrepresentationtypecode_propertytype *MD_SpatialRepresentationTypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_SpatialRepresentationTypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_SpatialRepresentationTypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_spatialrepresentationtypecode_propertytype)
	}
}

func (md_spatialrepresentationtypecode_propertytype *MD_SpatialRepresentationTypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SpatialRepresentationTypeCode_PropertyTypeUpdateCallback != nil {
		var frontMD_SpatialRepresentationTypeCode_PropertyType *MD_SpatialRepresentationTypeCode_PropertyType
		if front != nil {
			frontMD_SpatialRepresentationTypeCode_PropertyType, _ = front.(*MD_SpatialRepresentationTypeCode_PropertyType)
		}
		stage.OnAfterMD_SpatialRepresentationTypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_spatialrepresentationtypecode_propertytype, frontMD_SpatialRepresentationTypeCode_PropertyType)
	}
}

func (md_spatialrepresentationtypecode_propertytype *MD_SpatialRepresentationTypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_SpatialRepresentationTypeCode_PropertyTypeDeleteCallback != nil {
		var frontMD_SpatialRepresentationTypeCode_PropertyType *MD_SpatialRepresentationTypeCode_PropertyType
		if front != nil {
			frontMD_SpatialRepresentationTypeCode_PropertyType, _ = front.(*MD_SpatialRepresentationTypeCode_PropertyType)
		}
		stage.OnAfterMD_SpatialRepresentationTypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_spatialrepresentationtypecode_propertytype, frontMD_SpatialRepresentationTypeCode_PropertyType)
	}
}

func (md_standardorderprocess_propertytype *MD_StandardOrderProcess_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_StandardOrderProcess_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_StandardOrderProcess_PropertyTypeCreateCallback.OnAfterCreate(stage, md_standardorderprocess_propertytype)
	}
}

func (md_standardorderprocess_propertytype *MD_StandardOrderProcess_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_StandardOrderProcess_PropertyTypeUpdateCallback != nil {
		var frontMD_StandardOrderProcess_PropertyType *MD_StandardOrderProcess_PropertyType
		if front != nil {
			frontMD_StandardOrderProcess_PropertyType, _ = front.(*MD_StandardOrderProcess_PropertyType)
		}
		stage.OnAfterMD_StandardOrderProcess_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_standardorderprocess_propertytype, frontMD_StandardOrderProcess_PropertyType)
	}
}

func (md_standardorderprocess_propertytype *MD_StandardOrderProcess_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_StandardOrderProcess_PropertyTypeDeleteCallback != nil {
		var frontMD_StandardOrderProcess_PropertyType *MD_StandardOrderProcess_PropertyType
		if front != nil {
			frontMD_StandardOrderProcess_PropertyType, _ = front.(*MD_StandardOrderProcess_PropertyType)
		}
		stage.OnAfterMD_StandardOrderProcess_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_standardorderprocess_propertytype, frontMD_StandardOrderProcess_PropertyType)
	}
}

func (md_standardorderprocess_type *MD_StandardOrderProcess_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_StandardOrderProcess_TypeCreateCallback != nil {
		stage.OnAfterMD_StandardOrderProcess_TypeCreateCallback.OnAfterCreate(stage, md_standardorderprocess_type)
	}
}

func (md_standardorderprocess_type *MD_StandardOrderProcess_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_StandardOrderProcess_TypeUpdateCallback != nil {
		var frontMD_StandardOrderProcess_Type *MD_StandardOrderProcess_Type
		if front != nil {
			frontMD_StandardOrderProcess_Type, _ = front.(*MD_StandardOrderProcess_Type)
		}
		stage.OnAfterMD_StandardOrderProcess_TypeUpdateCallback.OnAfterUpdate(stage, md_standardorderprocess_type, frontMD_StandardOrderProcess_Type)
	}
}

func (md_standardorderprocess_type *MD_StandardOrderProcess_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_StandardOrderProcess_TypeDeleteCallback != nil {
		var frontMD_StandardOrderProcess_Type *MD_StandardOrderProcess_Type
		if front != nil {
			frontMD_StandardOrderProcess_Type, _ = front.(*MD_StandardOrderProcess_Type)
		}
		stage.OnAfterMD_StandardOrderProcess_TypeDeleteCallback.OnAfterDelete(stage, md_standardorderprocess_type, frontMD_StandardOrderProcess_Type)
	}
}

func (md_topiccategorycode_propertytype *MD_TopicCategoryCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_TopicCategoryCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_TopicCategoryCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_topiccategorycode_propertytype)
	}
}

func (md_topiccategorycode_propertytype *MD_TopicCategoryCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_TopicCategoryCode_PropertyTypeUpdateCallback != nil {
		var frontMD_TopicCategoryCode_PropertyType *MD_TopicCategoryCode_PropertyType
		if front != nil {
			frontMD_TopicCategoryCode_PropertyType, _ = front.(*MD_TopicCategoryCode_PropertyType)
		}
		stage.OnAfterMD_TopicCategoryCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_topiccategorycode_propertytype, frontMD_TopicCategoryCode_PropertyType)
	}
}

func (md_topiccategorycode_propertytype *MD_TopicCategoryCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_TopicCategoryCode_PropertyTypeDeleteCallback != nil {
		var frontMD_TopicCategoryCode_PropertyType *MD_TopicCategoryCode_PropertyType
		if front != nil {
			frontMD_TopicCategoryCode_PropertyType, _ = front.(*MD_TopicCategoryCode_PropertyType)
		}
		stage.OnAfterMD_TopicCategoryCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_topiccategorycode_propertytype, frontMD_TopicCategoryCode_PropertyType)
	}
}

func (md_usage_propertytype *MD_Usage_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Usage_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_Usage_PropertyTypeCreateCallback.OnAfterCreate(stage, md_usage_propertytype)
	}
}

func (md_usage_propertytype *MD_Usage_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Usage_PropertyTypeUpdateCallback != nil {
		var frontMD_Usage_PropertyType *MD_Usage_PropertyType
		if front != nil {
			frontMD_Usage_PropertyType, _ = front.(*MD_Usage_PropertyType)
		}
		stage.OnAfterMD_Usage_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_usage_propertytype, frontMD_Usage_PropertyType)
	}
}

func (md_usage_propertytype *MD_Usage_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Usage_PropertyTypeDeleteCallback != nil {
		var frontMD_Usage_PropertyType *MD_Usage_PropertyType
		if front != nil {
			frontMD_Usage_PropertyType, _ = front.(*MD_Usage_PropertyType)
		}
		stage.OnAfterMD_Usage_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_usage_propertytype, frontMD_Usage_PropertyType)
	}
}

func (md_usage_type *MD_Usage_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_Usage_TypeCreateCallback != nil {
		stage.OnAfterMD_Usage_TypeCreateCallback.OnAfterCreate(stage, md_usage_type)
	}
}

func (md_usage_type *MD_Usage_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Usage_TypeUpdateCallback != nil {
		var frontMD_Usage_Type *MD_Usage_Type
		if front != nil {
			frontMD_Usage_Type, _ = front.(*MD_Usage_Type)
		}
		stage.OnAfterMD_Usage_TypeUpdateCallback.OnAfterUpdate(stage, md_usage_type, frontMD_Usage_Type)
	}
}

func (md_usage_type *MD_Usage_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_Usage_TypeDeleteCallback != nil {
		var frontMD_Usage_Type *MD_Usage_Type
		if front != nil {
			frontMD_Usage_Type, _ = front.(*MD_Usage_Type)
		}
		stage.OnAfterMD_Usage_TypeDeleteCallback.OnAfterDelete(stage, md_usage_type, frontMD_Usage_Type)
	}
}

func (mappingruletype *MappingRuleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMappingRuleTypeCreateCallback != nil {
		stage.OnAfterMappingRuleTypeCreateCallback.OnAfterCreate(stage, mappingruletype)
	}
}

func (mappingruletype *MappingRuleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMappingRuleTypeUpdateCallback != nil {
		var frontMappingRuleType *MappingRuleType
		if front != nil {
			frontMappingRuleType, _ = front.(*MappingRuleType)
		}
		stage.OnAfterMappingRuleTypeUpdateCallback.OnAfterUpdate(stage, mappingruletype, frontMappingRuleType)
	}
}

func (mappingruletype *MappingRuleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMappingRuleTypeDeleteCallback != nil {
		var frontMappingRuleType *MappingRuleType
		if front != nil {
			frontMappingRuleType, _ = front.(*MappingRuleType)
		}
		stage.OnAfterMappingRuleTypeDeleteCallback.OnAfterDelete(stage, mappingruletype, frontMappingRuleType)
	}
}

func (measure_propertytype *Measure_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMeasure_PropertyTypeCreateCallback != nil {
		stage.OnAfterMeasure_PropertyTypeCreateCallback.OnAfterCreate(stage, measure_propertytype)
	}
}

func (measure_propertytype *Measure_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_PropertyTypeUpdateCallback != nil {
		var frontMeasure_PropertyType *Measure_PropertyType
		if front != nil {
			frontMeasure_PropertyType, _ = front.(*Measure_PropertyType)
		}
		stage.OnAfterMeasure_PropertyTypeUpdateCallback.OnAfterUpdate(stage, measure_propertytype, frontMeasure_PropertyType)
	}
}

func (measure_propertytype *Measure_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMeasure_PropertyTypeDeleteCallback != nil {
		var frontMeasure_PropertyType *Measure_PropertyType
		if front != nil {
			frontMeasure_PropertyType, _ = front.(*Measure_PropertyType)
		}
		stage.OnAfterMeasure_PropertyTypeDeleteCallback.OnAfterDelete(stage, measure_propertytype, frontMeasure_PropertyType)
	}
}

func (membername_propertytype *MemberName_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMemberName_PropertyTypeCreateCallback != nil {
		stage.OnAfterMemberName_PropertyTypeCreateCallback.OnAfterCreate(stage, membername_propertytype)
	}
}

func (membername_propertytype *MemberName_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMemberName_PropertyTypeUpdateCallback != nil {
		var frontMemberName_PropertyType *MemberName_PropertyType
		if front != nil {
			frontMemberName_PropertyType, _ = front.(*MemberName_PropertyType)
		}
		stage.OnAfterMemberName_PropertyTypeUpdateCallback.OnAfterUpdate(stage, membername_propertytype, frontMemberName_PropertyType)
	}
}

func (membername_propertytype *MemberName_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMemberName_PropertyTypeDeleteCallback != nil {
		var frontMemberName_PropertyType *MemberName_PropertyType
		if front != nil {
			frontMemberName_PropertyType, _ = front.(*MemberName_PropertyType)
		}
		stage.OnAfterMemberName_PropertyTypeDeleteCallback.OnAfterDelete(stage, membername_propertytype, frontMemberName_PropertyType)
	}
}

func (membername_type *MemberName_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMemberName_TypeCreateCallback != nil {
		stage.OnAfterMemberName_TypeCreateCallback.OnAfterCreate(stage, membername_type)
	}
}

func (membername_type *MemberName_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMemberName_TypeUpdateCallback != nil {
		var frontMemberName_Type *MemberName_Type
		if front != nil {
			frontMemberName_Type, _ = front.(*MemberName_Type)
		}
		stage.OnAfterMemberName_TypeUpdateCallback.OnAfterUpdate(stage, membername_type, frontMemberName_Type)
	}
}

func (membername_type *MemberName_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMemberName_TypeDeleteCallback != nil {
		var frontMemberName_Type *MemberName_Type
		if front != nil {
			frontMemberName_Type, _ = front.(*MemberName_Type)
		}
		stage.OnAfterMemberName_TypeDeleteCallback.OnAfterDelete(stage, membername_type, frontMemberName_Type)
	}
}

func (metadatapropertytype *MetaDataPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMetaDataPropertyTypeCreateCallback != nil {
		stage.OnAfterMetaDataPropertyTypeCreateCallback.OnAfterCreate(stage, metadatapropertytype)
	}
}

func (metadatapropertytype *MetaDataPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetaDataPropertyTypeUpdateCallback != nil {
		var frontMetaDataPropertyType *MetaDataPropertyType
		if front != nil {
			frontMetaDataPropertyType, _ = front.(*MetaDataPropertyType)
		}
		stage.OnAfterMetaDataPropertyTypeUpdateCallback.OnAfterUpdate(stage, metadatapropertytype, frontMetaDataPropertyType)
	}
}

func (metadatapropertytype *MetaDataPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMetaDataPropertyTypeDeleteCallback != nil {
		var frontMetaDataPropertyType *MetaDataPropertyType
		if front != nil {
			frontMetaDataPropertyType, _ = front.(*MetaDataPropertyType)
		}
		stage.OnAfterMetaDataPropertyTypeDeleteCallback.OnAfterDelete(stage, metadatapropertytype, frontMetaDataPropertyType)
	}
}

func (movingobjectstatustype *MovingObjectStatusType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMovingObjectStatusTypeCreateCallback != nil {
		stage.OnAfterMovingObjectStatusTypeCreateCallback.OnAfterCreate(stage, movingobjectstatustype)
	}
}

func (movingobjectstatustype *MovingObjectStatusType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMovingObjectStatusTypeUpdateCallback != nil {
		var frontMovingObjectStatusType *MovingObjectStatusType
		if front != nil {
			frontMovingObjectStatusType, _ = front.(*MovingObjectStatusType)
		}
		stage.OnAfterMovingObjectStatusTypeUpdateCallback.OnAfterUpdate(stage, movingobjectstatustype, frontMovingObjectStatusType)
	}
}

func (movingobjectstatustype *MovingObjectStatusType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMovingObjectStatusTypeDeleteCallback != nil {
		var frontMovingObjectStatusType *MovingObjectStatusType
		if front != nil {
			frontMovingObjectStatusType, _ = front.(*MovingObjectStatusType)
		}
		stage.OnAfterMovingObjectStatusTypeDeleteCallback.OnAfterDelete(stage, movingobjectstatustype, frontMovingObjectStatusType)
	}
}

func (multicurvepropertytype *MultiCurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterMultiCurvePropertyTypeCreateCallback.OnAfterCreate(stage, multicurvepropertytype)
	}
}

func (multicurvepropertytype *MultiCurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiCurvePropertyTypeUpdateCallback != nil {
		var frontMultiCurvePropertyType *MultiCurvePropertyType
		if front != nil {
			frontMultiCurvePropertyType, _ = front.(*MultiCurvePropertyType)
		}
		stage.OnAfterMultiCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, multicurvepropertytype, frontMultiCurvePropertyType)
	}
}

func (multicurvepropertytype *MultiCurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiCurvePropertyTypeDeleteCallback != nil {
		var frontMultiCurvePropertyType *MultiCurvePropertyType
		if front != nil {
			frontMultiCurvePropertyType, _ = front.(*MultiCurvePropertyType)
		}
		stage.OnAfterMultiCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, multicurvepropertytype, frontMultiCurvePropertyType)
	}
}

func (multicurvetype *MultiCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiCurveTypeCreateCallback != nil {
		stage.OnAfterMultiCurveTypeCreateCallback.OnAfterCreate(stage, multicurvetype)
	}
}

func (multicurvetype *MultiCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiCurveTypeUpdateCallback != nil {
		var frontMultiCurveType *MultiCurveType
		if front != nil {
			frontMultiCurveType, _ = front.(*MultiCurveType)
		}
		stage.OnAfterMultiCurveTypeUpdateCallback.OnAfterUpdate(stage, multicurvetype, frontMultiCurveType)
	}
}

func (multicurvetype *MultiCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiCurveTypeDeleteCallback != nil {
		var frontMultiCurveType *MultiCurveType
		if front != nil {
			frontMultiCurveType, _ = front.(*MultiCurveType)
		}
		stage.OnAfterMultiCurveTypeDeleteCallback.OnAfterDelete(stage, multicurvetype, frontMultiCurveType)
	}
}

func (multigeometrypropertytype *MultiGeometryPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiGeometryPropertyTypeCreateCallback != nil {
		stage.OnAfterMultiGeometryPropertyTypeCreateCallback.OnAfterCreate(stage, multigeometrypropertytype)
	}
}

func (multigeometrypropertytype *MultiGeometryPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiGeometryPropertyTypeUpdateCallback != nil {
		var frontMultiGeometryPropertyType *MultiGeometryPropertyType
		if front != nil {
			frontMultiGeometryPropertyType, _ = front.(*MultiGeometryPropertyType)
		}
		stage.OnAfterMultiGeometryPropertyTypeUpdateCallback.OnAfterUpdate(stage, multigeometrypropertytype, frontMultiGeometryPropertyType)
	}
}

func (multigeometrypropertytype *MultiGeometryPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiGeometryPropertyTypeDeleteCallback != nil {
		var frontMultiGeometryPropertyType *MultiGeometryPropertyType
		if front != nil {
			frontMultiGeometryPropertyType, _ = front.(*MultiGeometryPropertyType)
		}
		stage.OnAfterMultiGeometryPropertyTypeDeleteCallback.OnAfterDelete(stage, multigeometrypropertytype, frontMultiGeometryPropertyType)
	}
}

func (multigeometrytype *MultiGeometryType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiGeometryTypeCreateCallback != nil {
		stage.OnAfterMultiGeometryTypeCreateCallback.OnAfterCreate(stage, multigeometrytype)
	}
}

func (multigeometrytype *MultiGeometryType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiGeometryTypeUpdateCallback != nil {
		var frontMultiGeometryType *MultiGeometryType
		if front != nil {
			frontMultiGeometryType, _ = front.(*MultiGeometryType)
		}
		stage.OnAfterMultiGeometryTypeUpdateCallback.OnAfterUpdate(stage, multigeometrytype, frontMultiGeometryType)
	}
}

func (multigeometrytype *MultiGeometryType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiGeometryTypeDeleteCallback != nil {
		var frontMultiGeometryType *MultiGeometryType
		if front != nil {
			frontMultiGeometryType, _ = front.(*MultiGeometryType)
		}
		stage.OnAfterMultiGeometryTypeDeleteCallback.OnAfterDelete(stage, multigeometrytype, frontMultiGeometryType)
	}
}

func (multipointpropertytype *MultiPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiPointPropertyTypeCreateCallback != nil {
		stage.OnAfterMultiPointPropertyTypeCreateCallback.OnAfterCreate(stage, multipointpropertytype)
	}
}

func (multipointpropertytype *MultiPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiPointPropertyTypeUpdateCallback != nil {
		var frontMultiPointPropertyType *MultiPointPropertyType
		if front != nil {
			frontMultiPointPropertyType, _ = front.(*MultiPointPropertyType)
		}
		stage.OnAfterMultiPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, multipointpropertytype, frontMultiPointPropertyType)
	}
}

func (multipointpropertytype *MultiPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiPointPropertyTypeDeleteCallback != nil {
		var frontMultiPointPropertyType *MultiPointPropertyType
		if front != nil {
			frontMultiPointPropertyType, _ = front.(*MultiPointPropertyType)
		}
		stage.OnAfterMultiPointPropertyTypeDeleteCallback.OnAfterDelete(stage, multipointpropertytype, frontMultiPointPropertyType)
	}
}

func (multipointtype *MultiPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiPointTypeCreateCallback != nil {
		stage.OnAfterMultiPointTypeCreateCallback.OnAfterCreate(stage, multipointtype)
	}
}

func (multipointtype *MultiPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiPointTypeUpdateCallback != nil {
		var frontMultiPointType *MultiPointType
		if front != nil {
			frontMultiPointType, _ = front.(*MultiPointType)
		}
		stage.OnAfterMultiPointTypeUpdateCallback.OnAfterUpdate(stage, multipointtype, frontMultiPointType)
	}
}

func (multipointtype *MultiPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiPointTypeDeleteCallback != nil {
		var frontMultiPointType *MultiPointType
		if front != nil {
			frontMultiPointType, _ = front.(*MultiPointType)
		}
		stage.OnAfterMultiPointTypeDeleteCallback.OnAfterDelete(stage, multipointtype, frontMultiPointType)
	}
}

func (multisolidpropertytype *MultiSolidPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiSolidPropertyTypeCreateCallback != nil {
		stage.OnAfterMultiSolidPropertyTypeCreateCallback.OnAfterCreate(stage, multisolidpropertytype)
	}
}

func (multisolidpropertytype *MultiSolidPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSolidPropertyTypeUpdateCallback != nil {
		var frontMultiSolidPropertyType *MultiSolidPropertyType
		if front != nil {
			frontMultiSolidPropertyType, _ = front.(*MultiSolidPropertyType)
		}
		stage.OnAfterMultiSolidPropertyTypeUpdateCallback.OnAfterUpdate(stage, multisolidpropertytype, frontMultiSolidPropertyType)
	}
}

func (multisolidpropertytype *MultiSolidPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSolidPropertyTypeDeleteCallback != nil {
		var frontMultiSolidPropertyType *MultiSolidPropertyType
		if front != nil {
			frontMultiSolidPropertyType, _ = front.(*MultiSolidPropertyType)
		}
		stage.OnAfterMultiSolidPropertyTypeDeleteCallback.OnAfterDelete(stage, multisolidpropertytype, frontMultiSolidPropertyType)
	}
}

func (multisolidtype *MultiSolidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiSolidTypeCreateCallback != nil {
		stage.OnAfterMultiSolidTypeCreateCallback.OnAfterCreate(stage, multisolidtype)
	}
}

func (multisolidtype *MultiSolidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSolidTypeUpdateCallback != nil {
		var frontMultiSolidType *MultiSolidType
		if front != nil {
			frontMultiSolidType, _ = front.(*MultiSolidType)
		}
		stage.OnAfterMultiSolidTypeUpdateCallback.OnAfterUpdate(stage, multisolidtype, frontMultiSolidType)
	}
}

func (multisolidtype *MultiSolidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSolidTypeDeleteCallback != nil {
		var frontMultiSolidType *MultiSolidType
		if front != nil {
			frontMultiSolidType, _ = front.(*MultiSolidType)
		}
		stage.OnAfterMultiSolidTypeDeleteCallback.OnAfterDelete(stage, multisolidtype, frontMultiSolidType)
	}
}

func (multisurfacepropertytype *MultiSurfacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiSurfacePropertyTypeCreateCallback != nil {
		stage.OnAfterMultiSurfacePropertyTypeCreateCallback.OnAfterCreate(stage, multisurfacepropertytype)
	}
}

func (multisurfacepropertytype *MultiSurfacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSurfacePropertyTypeUpdateCallback != nil {
		var frontMultiSurfacePropertyType *MultiSurfacePropertyType
		if front != nil {
			frontMultiSurfacePropertyType, _ = front.(*MultiSurfacePropertyType)
		}
		stage.OnAfterMultiSurfacePropertyTypeUpdateCallback.OnAfterUpdate(stage, multisurfacepropertytype, frontMultiSurfacePropertyType)
	}
}

func (multisurfacepropertytype *MultiSurfacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSurfacePropertyTypeDeleteCallback != nil {
		var frontMultiSurfacePropertyType *MultiSurfacePropertyType
		if front != nil {
			frontMultiSurfacePropertyType, _ = front.(*MultiSurfacePropertyType)
		}
		stage.OnAfterMultiSurfacePropertyTypeDeleteCallback.OnAfterDelete(stage, multisurfacepropertytype, frontMultiSurfacePropertyType)
	}
}

func (multisurfacetype *MultiSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiSurfaceTypeCreateCallback != nil {
		stage.OnAfterMultiSurfaceTypeCreateCallback.OnAfterCreate(stage, multisurfacetype)
	}
}

func (multisurfacetype *MultiSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSurfaceTypeUpdateCallback != nil {
		var frontMultiSurfaceType *MultiSurfaceType
		if front != nil {
			frontMultiSurfaceType, _ = front.(*MultiSurfaceType)
		}
		stage.OnAfterMultiSurfaceTypeUpdateCallback.OnAfterUpdate(stage, multisurfacetype, frontMultiSurfaceType)
	}
}

func (multisurfacetype *MultiSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiSurfaceTypeDeleteCallback != nil {
		var frontMultiSurfaceType *MultiSurfaceType
		if front != nil {
			frontMultiSurfaceType, _ = front.(*MultiSurfaceType)
		}
		stage.OnAfterMultiSurfaceTypeDeleteCallback.OnAfterDelete(stage, multisurfacetype, frontMultiSurfaceType)
	}
}

func (multiplicityrange_propertytype *MultiplicityRange_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiplicityRange_PropertyTypeCreateCallback != nil {
		stage.OnAfterMultiplicityRange_PropertyTypeCreateCallback.OnAfterCreate(stage, multiplicityrange_propertytype)
	}
}

func (multiplicityrange_propertytype *MultiplicityRange_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicityRange_PropertyTypeUpdateCallback != nil {
		var frontMultiplicityRange_PropertyType *MultiplicityRange_PropertyType
		if front != nil {
			frontMultiplicityRange_PropertyType, _ = front.(*MultiplicityRange_PropertyType)
		}
		stage.OnAfterMultiplicityRange_PropertyTypeUpdateCallback.OnAfterUpdate(stage, multiplicityrange_propertytype, frontMultiplicityRange_PropertyType)
	}
}

func (multiplicityrange_propertytype *MultiplicityRange_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicityRange_PropertyTypeDeleteCallback != nil {
		var frontMultiplicityRange_PropertyType *MultiplicityRange_PropertyType
		if front != nil {
			frontMultiplicityRange_PropertyType, _ = front.(*MultiplicityRange_PropertyType)
		}
		stage.OnAfterMultiplicityRange_PropertyTypeDeleteCallback.OnAfterDelete(stage, multiplicityrange_propertytype, frontMultiplicityRange_PropertyType)
	}
}

func (multiplicityrange_type *MultiplicityRange_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiplicityRange_TypeCreateCallback != nil {
		stage.OnAfterMultiplicityRange_TypeCreateCallback.OnAfterCreate(stage, multiplicityrange_type)
	}
}

func (multiplicityrange_type *MultiplicityRange_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicityRange_TypeUpdateCallback != nil {
		var frontMultiplicityRange_Type *MultiplicityRange_Type
		if front != nil {
			frontMultiplicityRange_Type, _ = front.(*MultiplicityRange_Type)
		}
		stage.OnAfterMultiplicityRange_TypeUpdateCallback.OnAfterUpdate(stage, multiplicityrange_type, frontMultiplicityRange_Type)
	}
}

func (multiplicityrange_type *MultiplicityRange_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicityRange_TypeDeleteCallback != nil {
		var frontMultiplicityRange_Type *MultiplicityRange_Type
		if front != nil {
			frontMultiplicityRange_Type, _ = front.(*MultiplicityRange_Type)
		}
		stage.OnAfterMultiplicityRange_TypeDeleteCallback.OnAfterDelete(stage, multiplicityrange_type, frontMultiplicityRange_Type)
	}
}

func (multiplicity_propertytype *Multiplicity_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiplicity_PropertyTypeCreateCallback != nil {
		stage.OnAfterMultiplicity_PropertyTypeCreateCallback.OnAfterCreate(stage, multiplicity_propertytype)
	}
}

func (multiplicity_propertytype *Multiplicity_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicity_PropertyTypeUpdateCallback != nil {
		var frontMultiplicity_PropertyType *Multiplicity_PropertyType
		if front != nil {
			frontMultiplicity_PropertyType, _ = front.(*Multiplicity_PropertyType)
		}
		stage.OnAfterMultiplicity_PropertyTypeUpdateCallback.OnAfterUpdate(stage, multiplicity_propertytype, frontMultiplicity_PropertyType)
	}
}

func (multiplicity_propertytype *Multiplicity_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicity_PropertyTypeDeleteCallback != nil {
		var frontMultiplicity_PropertyType *Multiplicity_PropertyType
		if front != nil {
			frontMultiplicity_PropertyType, _ = front.(*Multiplicity_PropertyType)
		}
		stage.OnAfterMultiplicity_PropertyTypeDeleteCallback.OnAfterDelete(stage, multiplicity_propertytype, frontMultiplicity_PropertyType)
	}
}

func (multiplicity_type *Multiplicity_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMultiplicity_TypeCreateCallback != nil {
		stage.OnAfterMultiplicity_TypeCreateCallback.OnAfterCreate(stage, multiplicity_type)
	}
}

func (multiplicity_type *Multiplicity_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicity_TypeUpdateCallback != nil {
		var frontMultiplicity_Type *Multiplicity_Type
		if front != nil {
			frontMultiplicity_Type, _ = front.(*Multiplicity_Type)
		}
		stage.OnAfterMultiplicity_TypeUpdateCallback.OnAfterUpdate(stage, multiplicity_type, frontMultiplicity_Type)
	}
}

func (multiplicity_type *Multiplicity_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMultiplicity_TypeDeleteCallback != nil {
		var frontMultiplicity_Type *Multiplicity_Type
		if front != nil {
			frontMultiplicity_Type, _ = front.(*Multiplicity_Type)
		}
		stage.OnAfterMultiplicity_TypeDeleteCallback.OnAfterDelete(stage, multiplicity_type, frontMultiplicity_Type)
	}
}

func (nodeoredgepropertytype *NodeOrEdgePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNodeOrEdgePropertyTypeCreateCallback != nil {
		stage.OnAfterNodeOrEdgePropertyTypeCreateCallback.OnAfterCreate(stage, nodeoredgepropertytype)
	}
}

func (nodeoredgepropertytype *NodeOrEdgePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeOrEdgePropertyTypeUpdateCallback != nil {
		var frontNodeOrEdgePropertyType *NodeOrEdgePropertyType
		if front != nil {
			frontNodeOrEdgePropertyType, _ = front.(*NodeOrEdgePropertyType)
		}
		stage.OnAfterNodeOrEdgePropertyTypeUpdateCallback.OnAfterUpdate(stage, nodeoredgepropertytype, frontNodeOrEdgePropertyType)
	}
}

func (nodeoredgepropertytype *NodeOrEdgePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeOrEdgePropertyTypeDeleteCallback != nil {
		var frontNodeOrEdgePropertyType *NodeOrEdgePropertyType
		if front != nil {
			frontNodeOrEdgePropertyType, _ = front.(*NodeOrEdgePropertyType)
		}
		stage.OnAfterNodeOrEdgePropertyTypeDeleteCallback.OnAfterDelete(stage, nodeoredgepropertytype, frontNodeOrEdgePropertyType)
	}
}

func (nodepropertytype *NodePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNodePropertyTypeCreateCallback != nil {
		stage.OnAfterNodePropertyTypeCreateCallback.OnAfterCreate(stage, nodepropertytype)
	}
}

func (nodepropertytype *NodePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodePropertyTypeUpdateCallback != nil {
		var frontNodePropertyType *NodePropertyType
		if front != nil {
			frontNodePropertyType, _ = front.(*NodePropertyType)
		}
		stage.OnAfterNodePropertyTypeUpdateCallback.OnAfterUpdate(stage, nodepropertytype, frontNodePropertyType)
	}
}

func (nodepropertytype *NodePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodePropertyTypeDeleteCallback != nil {
		var frontNodePropertyType *NodePropertyType
		if front != nil {
			frontNodePropertyType, _ = front.(*NodePropertyType)
		}
		stage.OnAfterNodePropertyTypeDeleteCallback.OnAfterDelete(stage, nodepropertytype, frontNodePropertyType)
	}
}

func (nodetype *NodeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNodeTypeCreateCallback != nil {
		stage.OnAfterNodeTypeCreateCallback.OnAfterCreate(stage, nodetype)
	}
}

func (nodetype *NodeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeTypeUpdateCallback != nil {
		var frontNodeType *NodeType
		if front != nil {
			frontNodeType, _ = front.(*NodeType)
		}
		stage.OnAfterNodeTypeUpdateCallback.OnAfterUpdate(stage, nodetype, frontNodeType)
	}
}

func (nodetype *NodeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeTypeDeleteCallback != nil {
		var frontNodeType *NodeType
		if front != nil {
			frontNodeType, _ = front.(*NodeType)
		}
		stage.OnAfterNodeTypeDeleteCallback.OnAfterDelete(stage, nodetype, frontNodeType)
	}
}

func (number_propertytype *Number_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNumber_PropertyTypeCreateCallback != nil {
		stage.OnAfterNumber_PropertyTypeCreateCallback.OnAfterCreate(stage, number_propertytype)
	}
}

func (number_propertytype *Number_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumber_PropertyTypeUpdateCallback != nil {
		var frontNumber_PropertyType *Number_PropertyType
		if front != nil {
			frontNumber_PropertyType, _ = front.(*Number_PropertyType)
		}
		stage.OnAfterNumber_PropertyTypeUpdateCallback.OnAfterUpdate(stage, number_propertytype, frontNumber_PropertyType)
	}
}

func (number_propertytype *Number_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNumber_PropertyTypeDeleteCallback != nil {
		var frontNumber_PropertyType *Number_PropertyType
		if front != nil {
			frontNumber_PropertyType, _ = front.(*Number_PropertyType)
		}
		stage.OnAfterNumber_PropertyTypeDeleteCallback.OnAfterDelete(stage, number_propertytype, frontNumber_PropertyType)
	}
}

func (obliquecartesiancspropertytype *ObliqueCartesianCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObliqueCartesianCSPropertyTypeCreateCallback != nil {
		stage.OnAfterObliqueCartesianCSPropertyTypeCreateCallback.OnAfterCreate(stage, obliquecartesiancspropertytype)
	}
}

func (obliquecartesiancspropertytype *ObliqueCartesianCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObliqueCartesianCSPropertyTypeUpdateCallback != nil {
		var frontObliqueCartesianCSPropertyType *ObliqueCartesianCSPropertyType
		if front != nil {
			frontObliqueCartesianCSPropertyType, _ = front.(*ObliqueCartesianCSPropertyType)
		}
		stage.OnAfterObliqueCartesianCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, obliquecartesiancspropertytype, frontObliqueCartesianCSPropertyType)
	}
}

func (obliquecartesiancspropertytype *ObliqueCartesianCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObliqueCartesianCSPropertyTypeDeleteCallback != nil {
		var frontObliqueCartesianCSPropertyType *ObliqueCartesianCSPropertyType
		if front != nil {
			frontObliqueCartesianCSPropertyType, _ = front.(*ObliqueCartesianCSPropertyType)
		}
		stage.OnAfterObliqueCartesianCSPropertyTypeDeleteCallback.OnAfterDelete(stage, obliquecartesiancspropertytype, frontObliqueCartesianCSPropertyType)
	}
}

func (obliquecartesiancstype *ObliqueCartesianCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObliqueCartesianCSTypeCreateCallback != nil {
		stage.OnAfterObliqueCartesianCSTypeCreateCallback.OnAfterCreate(stage, obliquecartesiancstype)
	}
}

func (obliquecartesiancstype *ObliqueCartesianCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObliqueCartesianCSTypeUpdateCallback != nil {
		var frontObliqueCartesianCSType *ObliqueCartesianCSType
		if front != nil {
			frontObliqueCartesianCSType, _ = front.(*ObliqueCartesianCSType)
		}
		stage.OnAfterObliqueCartesianCSTypeUpdateCallback.OnAfterUpdate(stage, obliquecartesiancstype, frontObliqueCartesianCSType)
	}
}

func (obliquecartesiancstype *ObliqueCartesianCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObliqueCartesianCSTypeDeleteCallback != nil {
		var frontObliqueCartesianCSType *ObliqueCartesianCSType
		if front != nil {
			frontObliqueCartesianCSType, _ = front.(*ObliqueCartesianCSType)
		}
		stage.OnAfterObliqueCartesianCSTypeDeleteCallback.OnAfterDelete(stage, obliquecartesiancstype, frontObliqueCartesianCSType)
	}
}

func (observationtype *ObservationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObservationTypeCreateCallback != nil {
		stage.OnAfterObservationTypeCreateCallback.OnAfterCreate(stage, observationtype)
	}
}

func (observationtype *ObservationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObservationTypeUpdateCallback != nil {
		var frontObservationType *ObservationType
		if front != nil {
			frontObservationType, _ = front.(*ObservationType)
		}
		stage.OnAfterObservationTypeUpdateCallback.OnAfterUpdate(stage, observationtype, frontObservationType)
	}
}

func (observationtype *ObservationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObservationTypeDeleteCallback != nil {
		var frontObservationType *ObservationType
		if front != nil {
			frontObservationType, _ = front.(*ObservationType)
		}
		stage.OnAfterObservationTypeDeleteCallback.OnAfterDelete(stage, observationtype, frontObservationType)
	}
}

func (offsetcurvetype *OffsetCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOffsetCurveTypeCreateCallback != nil {
		stage.OnAfterOffsetCurveTypeCreateCallback.OnAfterCreate(stage, offsetcurvetype)
	}
}

func (offsetcurvetype *OffsetCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOffsetCurveTypeUpdateCallback != nil {
		var frontOffsetCurveType *OffsetCurveType
		if front != nil {
			frontOffsetCurveType, _ = front.(*OffsetCurveType)
		}
		stage.OnAfterOffsetCurveTypeUpdateCallback.OnAfterUpdate(stage, offsetcurvetype, frontOffsetCurveType)
	}
}

func (offsetcurvetype *OffsetCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOffsetCurveTypeDeleteCallback != nil {
		var frontOffsetCurveType *OffsetCurveType
		if front != nil {
			frontOffsetCurveType, _ = front.(*OffsetCurveType)
		}
		stage.OnAfterOffsetCurveTypeDeleteCallback.OnAfterDelete(stage, offsetcurvetype, frontOffsetCurveType)
	}
}

func (operationmethodpropertytype *OperationMethodPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOperationMethodPropertyTypeCreateCallback != nil {
		stage.OnAfterOperationMethodPropertyTypeCreateCallback.OnAfterCreate(stage, operationmethodpropertytype)
	}
}

func (operationmethodpropertytype *OperationMethodPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationMethodPropertyTypeUpdateCallback != nil {
		var frontOperationMethodPropertyType *OperationMethodPropertyType
		if front != nil {
			frontOperationMethodPropertyType, _ = front.(*OperationMethodPropertyType)
		}
		stage.OnAfterOperationMethodPropertyTypeUpdateCallback.OnAfterUpdate(stage, operationmethodpropertytype, frontOperationMethodPropertyType)
	}
}

func (operationmethodpropertytype *OperationMethodPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationMethodPropertyTypeDeleteCallback != nil {
		var frontOperationMethodPropertyType *OperationMethodPropertyType
		if front != nil {
			frontOperationMethodPropertyType, _ = front.(*OperationMethodPropertyType)
		}
		stage.OnAfterOperationMethodPropertyTypeDeleteCallback.OnAfterDelete(stage, operationmethodpropertytype, frontOperationMethodPropertyType)
	}
}

func (operationmethodtype *OperationMethodType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOperationMethodTypeCreateCallback != nil {
		stage.OnAfterOperationMethodTypeCreateCallback.OnAfterCreate(stage, operationmethodtype)
	}
}

func (operationmethodtype *OperationMethodType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationMethodTypeUpdateCallback != nil {
		var frontOperationMethodType *OperationMethodType
		if front != nil {
			frontOperationMethodType, _ = front.(*OperationMethodType)
		}
		stage.OnAfterOperationMethodTypeUpdateCallback.OnAfterUpdate(stage, operationmethodtype, frontOperationMethodType)
	}
}

func (operationmethodtype *OperationMethodType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationMethodTypeDeleteCallback != nil {
		var frontOperationMethodType *OperationMethodType
		if front != nil {
			frontOperationMethodType, _ = front.(*OperationMethodType)
		}
		stage.OnAfterOperationMethodTypeDeleteCallback.OnAfterDelete(stage, operationmethodtype, frontOperationMethodType)
	}
}

func (operationparametergrouppropertytype *OperationParameterGroupPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOperationParameterGroupPropertyTypeCreateCallback != nil {
		stage.OnAfterOperationParameterGroupPropertyTypeCreateCallback.OnAfterCreate(stage, operationparametergrouppropertytype)
	}
}

func (operationparametergrouppropertytype *OperationParameterGroupPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterGroupPropertyTypeUpdateCallback != nil {
		var frontOperationParameterGroupPropertyType *OperationParameterGroupPropertyType
		if front != nil {
			frontOperationParameterGroupPropertyType, _ = front.(*OperationParameterGroupPropertyType)
		}
		stage.OnAfterOperationParameterGroupPropertyTypeUpdateCallback.OnAfterUpdate(stage, operationparametergrouppropertytype, frontOperationParameterGroupPropertyType)
	}
}

func (operationparametergrouppropertytype *OperationParameterGroupPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterGroupPropertyTypeDeleteCallback != nil {
		var frontOperationParameterGroupPropertyType *OperationParameterGroupPropertyType
		if front != nil {
			frontOperationParameterGroupPropertyType, _ = front.(*OperationParameterGroupPropertyType)
		}
		stage.OnAfterOperationParameterGroupPropertyTypeDeleteCallback.OnAfterDelete(stage, operationparametergrouppropertytype, frontOperationParameterGroupPropertyType)
	}
}

func (operationparametergrouptype *OperationParameterGroupType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOperationParameterGroupTypeCreateCallback != nil {
		stage.OnAfterOperationParameterGroupTypeCreateCallback.OnAfterCreate(stage, operationparametergrouptype)
	}
}

func (operationparametergrouptype *OperationParameterGroupType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterGroupTypeUpdateCallback != nil {
		var frontOperationParameterGroupType *OperationParameterGroupType
		if front != nil {
			frontOperationParameterGroupType, _ = front.(*OperationParameterGroupType)
		}
		stage.OnAfterOperationParameterGroupTypeUpdateCallback.OnAfterUpdate(stage, operationparametergrouptype, frontOperationParameterGroupType)
	}
}

func (operationparametergrouptype *OperationParameterGroupType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterGroupTypeDeleteCallback != nil {
		var frontOperationParameterGroupType *OperationParameterGroupType
		if front != nil {
			frontOperationParameterGroupType, _ = front.(*OperationParameterGroupType)
		}
		stage.OnAfterOperationParameterGroupTypeDeleteCallback.OnAfterDelete(stage, operationparametergrouptype, frontOperationParameterGroupType)
	}
}

func (operationparameterpropertytype *OperationParameterPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOperationParameterPropertyTypeCreateCallback != nil {
		stage.OnAfterOperationParameterPropertyTypeCreateCallback.OnAfterCreate(stage, operationparameterpropertytype)
	}
}

func (operationparameterpropertytype *OperationParameterPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterPropertyTypeUpdateCallback != nil {
		var frontOperationParameterPropertyType *OperationParameterPropertyType
		if front != nil {
			frontOperationParameterPropertyType, _ = front.(*OperationParameterPropertyType)
		}
		stage.OnAfterOperationParameterPropertyTypeUpdateCallback.OnAfterUpdate(stage, operationparameterpropertytype, frontOperationParameterPropertyType)
	}
}

func (operationparameterpropertytype *OperationParameterPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterPropertyTypeDeleteCallback != nil {
		var frontOperationParameterPropertyType *OperationParameterPropertyType
		if front != nil {
			frontOperationParameterPropertyType, _ = front.(*OperationParameterPropertyType)
		}
		stage.OnAfterOperationParameterPropertyTypeDeleteCallback.OnAfterDelete(stage, operationparameterpropertytype, frontOperationParameterPropertyType)
	}
}

func (operationparametertype *OperationParameterType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOperationParameterTypeCreateCallback != nil {
		stage.OnAfterOperationParameterTypeCreateCallback.OnAfterCreate(stage, operationparametertype)
	}
}

func (operationparametertype *OperationParameterType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterTypeUpdateCallback != nil {
		var frontOperationParameterType *OperationParameterType
		if front != nil {
			frontOperationParameterType, _ = front.(*OperationParameterType)
		}
		stage.OnAfterOperationParameterTypeUpdateCallback.OnAfterUpdate(stage, operationparametertype, frontOperationParameterType)
	}
}

func (operationparametertype *OperationParameterType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationParameterTypeDeleteCallback != nil {
		var frontOperationParameterType *OperationParameterType
		if front != nil {
			frontOperationParameterType, _ = front.(*OperationParameterType)
		}
		stage.OnAfterOperationParameterTypeDeleteCallback.OnAfterDelete(stage, operationparametertype, frontOperationParameterType)
	}
}

func (operationpropertytype *OperationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOperationPropertyTypeCreateCallback != nil {
		stage.OnAfterOperationPropertyTypeCreateCallback.OnAfterCreate(stage, operationpropertytype)
	}
}

func (operationpropertytype *OperationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationPropertyTypeUpdateCallback != nil {
		var frontOperationPropertyType *OperationPropertyType
		if front != nil {
			frontOperationPropertyType, _ = front.(*OperationPropertyType)
		}
		stage.OnAfterOperationPropertyTypeUpdateCallback.OnAfterUpdate(stage, operationpropertytype, frontOperationPropertyType)
	}
}

func (operationpropertytype *OperationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOperationPropertyTypeDeleteCallback != nil {
		var frontOperationPropertyType *OperationPropertyType
		if front != nil {
			frontOperationPropertyType, _ = front.(*OperationPropertyType)
		}
		stage.OnAfterOperationPropertyTypeDeleteCallback.OnAfterDelete(stage, operationpropertytype, frontOperationPropertyType)
	}
}

func (orientablecurvetype *OrientableCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrientableCurveTypeCreateCallback != nil {
		stage.OnAfterOrientableCurveTypeCreateCallback.OnAfterCreate(stage, orientablecurvetype)
	}
}

func (orientablecurvetype *OrientableCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrientableCurveTypeUpdateCallback != nil {
		var frontOrientableCurveType *OrientableCurveType
		if front != nil {
			frontOrientableCurveType, _ = front.(*OrientableCurveType)
		}
		stage.OnAfterOrientableCurveTypeUpdateCallback.OnAfterUpdate(stage, orientablecurvetype, frontOrientableCurveType)
	}
}

func (orientablecurvetype *OrientableCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrientableCurveTypeDeleteCallback != nil {
		var frontOrientableCurveType *OrientableCurveType
		if front != nil {
			frontOrientableCurveType, _ = front.(*OrientableCurveType)
		}
		stage.OnAfterOrientableCurveTypeDeleteCallback.OnAfterDelete(stage, orientablecurvetype, frontOrientableCurveType)
	}
}

func (orientablesurfacetype *OrientableSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOrientableSurfaceTypeCreateCallback != nil {
		stage.OnAfterOrientableSurfaceTypeCreateCallback.OnAfterCreate(stage, orientablesurfacetype)
	}
}

func (orientablesurfacetype *OrientableSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrientableSurfaceTypeUpdateCallback != nil {
		var frontOrientableSurfaceType *OrientableSurfaceType
		if front != nil {
			frontOrientableSurfaceType, _ = front.(*OrientableSurfaceType)
		}
		stage.OnAfterOrientableSurfaceTypeUpdateCallback.OnAfterUpdate(stage, orientablesurfacetype, frontOrientableSurfaceType)
	}
}

func (orientablesurfacetype *OrientableSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOrientableSurfaceTypeDeleteCallback != nil {
		var frontOrientableSurfaceType *OrientableSurfaceType
		if front != nil {
			frontOrientableSurfaceType, _ = front.(*OrientableSurfaceType)
		}
		stage.OnAfterOrientableSurfaceTypeDeleteCallback.OnAfterDelete(stage, orientablesurfacetype, frontOrientableSurfaceType)
	}
}

func (pt_freetext_propertytype *PT_FreeText_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPT_FreeText_PropertyTypeCreateCallback != nil {
		stage.OnAfterPT_FreeText_PropertyTypeCreateCallback.OnAfterCreate(stage, pt_freetext_propertytype)
	}
}

func (pt_freetext_propertytype *PT_FreeText_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_FreeText_PropertyTypeUpdateCallback != nil {
		var frontPT_FreeText_PropertyType *PT_FreeText_PropertyType
		if front != nil {
			frontPT_FreeText_PropertyType, _ = front.(*PT_FreeText_PropertyType)
		}
		stage.OnAfterPT_FreeText_PropertyTypeUpdateCallback.OnAfterUpdate(stage, pt_freetext_propertytype, frontPT_FreeText_PropertyType)
	}
}

func (pt_freetext_propertytype *PT_FreeText_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_FreeText_PropertyTypeDeleteCallback != nil {
		var frontPT_FreeText_PropertyType *PT_FreeText_PropertyType
		if front != nil {
			frontPT_FreeText_PropertyType, _ = front.(*PT_FreeText_PropertyType)
		}
		stage.OnAfterPT_FreeText_PropertyTypeDeleteCallback.OnAfterDelete(stage, pt_freetext_propertytype, frontPT_FreeText_PropertyType)
	}
}

func (pt_freetext_type *PT_FreeText_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPT_FreeText_TypeCreateCallback != nil {
		stage.OnAfterPT_FreeText_TypeCreateCallback.OnAfterCreate(stage, pt_freetext_type)
	}
}

func (pt_freetext_type *PT_FreeText_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_FreeText_TypeUpdateCallback != nil {
		var frontPT_FreeText_Type *PT_FreeText_Type
		if front != nil {
			frontPT_FreeText_Type, _ = front.(*PT_FreeText_Type)
		}
		stage.OnAfterPT_FreeText_TypeUpdateCallback.OnAfterUpdate(stage, pt_freetext_type, frontPT_FreeText_Type)
	}
}

func (pt_freetext_type *PT_FreeText_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_FreeText_TypeDeleteCallback != nil {
		var frontPT_FreeText_Type *PT_FreeText_Type
		if front != nil {
			frontPT_FreeText_Type, _ = front.(*PT_FreeText_Type)
		}
		stage.OnAfterPT_FreeText_TypeDeleteCallback.OnAfterDelete(stage, pt_freetext_type, frontPT_FreeText_Type)
	}
}

func (pt_localecontainer_propertytype *PT_LocaleContainer_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPT_LocaleContainer_PropertyTypeCreateCallback != nil {
		stage.OnAfterPT_LocaleContainer_PropertyTypeCreateCallback.OnAfterCreate(stage, pt_localecontainer_propertytype)
	}
}

func (pt_localecontainer_propertytype *PT_LocaleContainer_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_LocaleContainer_PropertyTypeUpdateCallback != nil {
		var frontPT_LocaleContainer_PropertyType *PT_LocaleContainer_PropertyType
		if front != nil {
			frontPT_LocaleContainer_PropertyType, _ = front.(*PT_LocaleContainer_PropertyType)
		}
		stage.OnAfterPT_LocaleContainer_PropertyTypeUpdateCallback.OnAfterUpdate(stage, pt_localecontainer_propertytype, frontPT_LocaleContainer_PropertyType)
	}
}

func (pt_localecontainer_propertytype *PT_LocaleContainer_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_LocaleContainer_PropertyTypeDeleteCallback != nil {
		var frontPT_LocaleContainer_PropertyType *PT_LocaleContainer_PropertyType
		if front != nil {
			frontPT_LocaleContainer_PropertyType, _ = front.(*PT_LocaleContainer_PropertyType)
		}
		stage.OnAfterPT_LocaleContainer_PropertyTypeDeleteCallback.OnAfterDelete(stage, pt_localecontainer_propertytype, frontPT_LocaleContainer_PropertyType)
	}
}

func (pt_localecontainer_type *PT_LocaleContainer_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPT_LocaleContainer_TypeCreateCallback != nil {
		stage.OnAfterPT_LocaleContainer_TypeCreateCallback.OnAfterCreate(stage, pt_localecontainer_type)
	}
}

func (pt_localecontainer_type *PT_LocaleContainer_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_LocaleContainer_TypeUpdateCallback != nil {
		var frontPT_LocaleContainer_Type *PT_LocaleContainer_Type
		if front != nil {
			frontPT_LocaleContainer_Type, _ = front.(*PT_LocaleContainer_Type)
		}
		stage.OnAfterPT_LocaleContainer_TypeUpdateCallback.OnAfterUpdate(stage, pt_localecontainer_type, frontPT_LocaleContainer_Type)
	}
}

func (pt_localecontainer_type *PT_LocaleContainer_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_LocaleContainer_TypeDeleteCallback != nil {
		var frontPT_LocaleContainer_Type *PT_LocaleContainer_Type
		if front != nil {
			frontPT_LocaleContainer_Type, _ = front.(*PT_LocaleContainer_Type)
		}
		stage.OnAfterPT_LocaleContainer_TypeDeleteCallback.OnAfterDelete(stage, pt_localecontainer_type, frontPT_LocaleContainer_Type)
	}
}

func (pt_locale_propertytype *PT_Locale_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPT_Locale_PropertyTypeCreateCallback != nil {
		stage.OnAfterPT_Locale_PropertyTypeCreateCallback.OnAfterCreate(stage, pt_locale_propertytype)
	}
}

func (pt_locale_propertytype *PT_Locale_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_Locale_PropertyTypeUpdateCallback != nil {
		var frontPT_Locale_PropertyType *PT_Locale_PropertyType
		if front != nil {
			frontPT_Locale_PropertyType, _ = front.(*PT_Locale_PropertyType)
		}
		stage.OnAfterPT_Locale_PropertyTypeUpdateCallback.OnAfterUpdate(stage, pt_locale_propertytype, frontPT_Locale_PropertyType)
	}
}

func (pt_locale_propertytype *PT_Locale_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_Locale_PropertyTypeDeleteCallback != nil {
		var frontPT_Locale_PropertyType *PT_Locale_PropertyType
		if front != nil {
			frontPT_Locale_PropertyType, _ = front.(*PT_Locale_PropertyType)
		}
		stage.OnAfterPT_Locale_PropertyTypeDeleteCallback.OnAfterDelete(stage, pt_locale_propertytype, frontPT_Locale_PropertyType)
	}
}

func (pt_locale_type *PT_Locale_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPT_Locale_TypeCreateCallback != nil {
		stage.OnAfterPT_Locale_TypeCreateCallback.OnAfterCreate(stage, pt_locale_type)
	}
}

func (pt_locale_type *PT_Locale_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_Locale_TypeUpdateCallback != nil {
		var frontPT_Locale_Type *PT_Locale_Type
		if front != nil {
			frontPT_Locale_Type, _ = front.(*PT_Locale_Type)
		}
		stage.OnAfterPT_Locale_TypeUpdateCallback.OnAfterUpdate(stage, pt_locale_type, frontPT_Locale_Type)
	}
}

func (pt_locale_type *PT_Locale_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPT_Locale_TypeDeleteCallback != nil {
		var frontPT_Locale_Type *PT_Locale_Type
		if front != nil {
			frontPT_Locale_Type, _ = front.(*PT_Locale_Type)
		}
		stage.OnAfterPT_Locale_TypeDeleteCallback.OnAfterDelete(stage, pt_locale_type, frontPT_Locale_Type)
	}
}

func (parametervaluegrouptype *ParameterValueGroupType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParameterValueGroupTypeCreateCallback != nil {
		stage.OnAfterParameterValueGroupTypeCreateCallback.OnAfterCreate(stage, parametervaluegrouptype)
	}
}

func (parametervaluegrouptype *ParameterValueGroupType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterValueGroupTypeUpdateCallback != nil {
		var frontParameterValueGroupType *ParameterValueGroupType
		if front != nil {
			frontParameterValueGroupType, _ = front.(*ParameterValueGroupType)
		}
		stage.OnAfterParameterValueGroupTypeUpdateCallback.OnAfterUpdate(stage, parametervaluegrouptype, frontParameterValueGroupType)
	}
}

func (parametervaluegrouptype *ParameterValueGroupType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterValueGroupTypeDeleteCallback != nil {
		var frontParameterValueGroupType *ParameterValueGroupType
		if front != nil {
			frontParameterValueGroupType, _ = front.(*ParameterValueGroupType)
		}
		stage.OnAfterParameterValueGroupTypeDeleteCallback.OnAfterDelete(stage, parametervaluegrouptype, frontParameterValueGroupType)
	}
}

func (parametervaluetype *ParameterValueType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParameterValueTypeCreateCallback != nil {
		stage.OnAfterParameterValueTypeCreateCallback.OnAfterCreate(stage, parametervaluetype)
	}
}

func (parametervaluetype *ParameterValueType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterValueTypeUpdateCallback != nil {
		var frontParameterValueType *ParameterValueType
		if front != nil {
			frontParameterValueType, _ = front.(*ParameterValueType)
		}
		stage.OnAfterParameterValueTypeUpdateCallback.OnAfterUpdate(stage, parametervaluetype, frontParameterValueType)
	}
}

func (parametervaluetype *ParameterValueType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParameterValueTypeDeleteCallback != nil {
		var frontParameterValueType *ParameterValueType
		if front != nil {
			frontParameterValueType, _ = front.(*ParameterValueType)
		}
		stage.OnAfterParameterValueTypeDeleteCallback.OnAfterDelete(stage, parametervaluetype, frontParameterValueType)
	}
}

func (passthroughoperationpropertytype *PassThroughOperationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassThroughOperationPropertyTypeCreateCallback != nil {
		stage.OnAfterPassThroughOperationPropertyTypeCreateCallback.OnAfterCreate(stage, passthroughoperationpropertytype)
	}
}

func (passthroughoperationpropertytype *PassThroughOperationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassThroughOperationPropertyTypeUpdateCallback != nil {
		var frontPassThroughOperationPropertyType *PassThroughOperationPropertyType
		if front != nil {
			frontPassThroughOperationPropertyType, _ = front.(*PassThroughOperationPropertyType)
		}
		stage.OnAfterPassThroughOperationPropertyTypeUpdateCallback.OnAfterUpdate(stage, passthroughoperationpropertytype, frontPassThroughOperationPropertyType)
	}
}

func (passthroughoperationpropertytype *PassThroughOperationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassThroughOperationPropertyTypeDeleteCallback != nil {
		var frontPassThroughOperationPropertyType *PassThroughOperationPropertyType
		if front != nil {
			frontPassThroughOperationPropertyType, _ = front.(*PassThroughOperationPropertyType)
		}
		stage.OnAfterPassThroughOperationPropertyTypeDeleteCallback.OnAfterDelete(stage, passthroughoperationpropertytype, frontPassThroughOperationPropertyType)
	}
}

func (passthroughoperationtype *PassThroughOperationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPassThroughOperationTypeCreateCallback != nil {
		stage.OnAfterPassThroughOperationTypeCreateCallback.OnAfterCreate(stage, passthroughoperationtype)
	}
}

func (passthroughoperationtype *PassThroughOperationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassThroughOperationTypeUpdateCallback != nil {
		var frontPassThroughOperationType *PassThroughOperationType
		if front != nil {
			frontPassThroughOperationType, _ = front.(*PassThroughOperationType)
		}
		stage.OnAfterPassThroughOperationTypeUpdateCallback.OnAfterUpdate(stage, passthroughoperationtype, frontPassThroughOperationType)
	}
}

func (passthroughoperationtype *PassThroughOperationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPassThroughOperationTypeDeleteCallback != nil {
		var frontPassThroughOperationType *PassThroughOperationType
		if front != nil {
			frontPassThroughOperationType, _ = front.(*PassThroughOperationType)
		}
		stage.OnAfterPassThroughOperationTypeDeleteCallback.OnAfterDelete(stage, passthroughoperationtype, frontPassThroughOperationType)
	}
}

func (pointarraypropertytype *PointArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterPointArrayPropertyTypeCreateCallback.OnAfterCreate(stage, pointarraypropertytype)
	}
}

func (pointarraypropertytype *PointArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointArrayPropertyTypeUpdateCallback != nil {
		var frontPointArrayPropertyType *PointArrayPropertyType
		if front != nil {
			frontPointArrayPropertyType, _ = front.(*PointArrayPropertyType)
		}
		stage.OnAfterPointArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, pointarraypropertytype, frontPointArrayPropertyType)
	}
}

func (pointarraypropertytype *PointArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointArrayPropertyTypeDeleteCallback != nil {
		var frontPointArrayPropertyType *PointArrayPropertyType
		if front != nil {
			frontPointArrayPropertyType, _ = front.(*PointArrayPropertyType)
		}
		stage.OnAfterPointArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, pointarraypropertytype, frontPointArrayPropertyType)
	}
}

func (pointpropertytype *PointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointPropertyTypeCreateCallback != nil {
		stage.OnAfterPointPropertyTypeCreateCallback.OnAfterCreate(stage, pointpropertytype)
	}
}

func (pointpropertytype *PointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointPropertyTypeUpdateCallback != nil {
		var frontPointPropertyType *PointPropertyType
		if front != nil {
			frontPointPropertyType, _ = front.(*PointPropertyType)
		}
		stage.OnAfterPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, pointpropertytype, frontPointPropertyType)
	}
}

func (pointpropertytype *PointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointPropertyTypeDeleteCallback != nil {
		var frontPointPropertyType *PointPropertyType
		if front != nil {
			frontPointPropertyType, _ = front.(*PointPropertyType)
		}
		stage.OnAfterPointPropertyTypeDeleteCallback.OnAfterDelete(stage, pointpropertytype, frontPointPropertyType)
	}
}

func (pointtype *PointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointTypeCreateCallback != nil {
		stage.OnAfterPointTypeCreateCallback.OnAfterCreate(stage, pointtype)
	}
}

func (pointtype *PointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeUpdateCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeUpdateCallback.OnAfterUpdate(stage, pointtype, frontPointType)
	}
}

func (pointtype *PointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointTypeDeleteCallback != nil {
		var frontPointType *PointType
		if front != nil {
			frontPointType, _ = front.(*PointType)
		}
		stage.OnAfterPointTypeDeleteCallback.OnAfterDelete(stage, pointtype, frontPointType)
	}
}

func (polarcspropertytype *PolarCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolarCSPropertyTypeCreateCallback != nil {
		stage.OnAfterPolarCSPropertyTypeCreateCallback.OnAfterCreate(stage, polarcspropertytype)
	}
}

func (polarcspropertytype *PolarCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolarCSPropertyTypeUpdateCallback != nil {
		var frontPolarCSPropertyType *PolarCSPropertyType
		if front != nil {
			frontPolarCSPropertyType, _ = front.(*PolarCSPropertyType)
		}
		stage.OnAfterPolarCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, polarcspropertytype, frontPolarCSPropertyType)
	}
}

func (polarcspropertytype *PolarCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolarCSPropertyTypeDeleteCallback != nil {
		var frontPolarCSPropertyType *PolarCSPropertyType
		if front != nil {
			frontPolarCSPropertyType, _ = front.(*PolarCSPropertyType)
		}
		stage.OnAfterPolarCSPropertyTypeDeleteCallback.OnAfterDelete(stage, polarcspropertytype, frontPolarCSPropertyType)
	}
}

func (polarcstype *PolarCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolarCSTypeCreateCallback != nil {
		stage.OnAfterPolarCSTypeCreateCallback.OnAfterCreate(stage, polarcstype)
	}
}

func (polarcstype *PolarCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolarCSTypeUpdateCallback != nil {
		var frontPolarCSType *PolarCSType
		if front != nil {
			frontPolarCSType, _ = front.(*PolarCSType)
		}
		stage.OnAfterPolarCSTypeUpdateCallback.OnAfterUpdate(stage, polarcstype, frontPolarCSType)
	}
}

func (polarcstype *PolarCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolarCSTypeDeleteCallback != nil {
		var frontPolarCSType *PolarCSType
		if front != nil {
			frontPolarCSType, _ = front.(*PolarCSType)
		}
		stage.OnAfterPolarCSTypeDeleteCallback.OnAfterDelete(stage, polarcstype, frontPolarCSType)
	}
}

func (polygonpatchtype *PolygonPatchType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolygonPatchTypeCreateCallback != nil {
		stage.OnAfterPolygonPatchTypeCreateCallback.OnAfterCreate(stage, polygonpatchtype)
	}
}

func (polygonpatchtype *PolygonPatchType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonPatchTypeUpdateCallback != nil {
		var frontPolygonPatchType *PolygonPatchType
		if front != nil {
			frontPolygonPatchType, _ = front.(*PolygonPatchType)
		}
		stage.OnAfterPolygonPatchTypeUpdateCallback.OnAfterUpdate(stage, polygonpatchtype, frontPolygonPatchType)
	}
}

func (polygonpatchtype *PolygonPatchType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonPatchTypeDeleteCallback != nil {
		var frontPolygonPatchType *PolygonPatchType
		if front != nil {
			frontPolygonPatchType, _ = front.(*PolygonPatchType)
		}
		stage.OnAfterPolygonPatchTypeDeleteCallback.OnAfterDelete(stage, polygonpatchtype, frontPolygonPatchType)
	}
}

func (polygontype *PolygonType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPolygonTypeCreateCallback != nil {
		stage.OnAfterPolygonTypeCreateCallback.OnAfterCreate(stage, polygontype)
	}
}

func (polygontype *PolygonType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonTypeUpdateCallback != nil {
		var frontPolygonType *PolygonType
		if front != nil {
			frontPolygonType, _ = front.(*PolygonType)
		}
		stage.OnAfterPolygonTypeUpdateCallback.OnAfterUpdate(stage, polygontype, frontPolygonType)
	}
}

func (polygontype *PolygonType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPolygonTypeDeleteCallback != nil {
		var frontPolygonType *PolygonType
		if front != nil {
			frontPolygonType, _ = front.(*PolygonType)
		}
		stage.OnAfterPolygonTypeDeleteCallback.OnAfterDelete(stage, polygontype, frontPolygonType)
	}
}

func (primemeridianpropertytype *PrimeMeridianPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrimeMeridianPropertyTypeCreateCallback != nil {
		stage.OnAfterPrimeMeridianPropertyTypeCreateCallback.OnAfterCreate(stage, primemeridianpropertytype)
	}
}

func (primemeridianpropertytype *PrimeMeridianPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimeMeridianPropertyTypeUpdateCallback != nil {
		var frontPrimeMeridianPropertyType *PrimeMeridianPropertyType
		if front != nil {
			frontPrimeMeridianPropertyType, _ = front.(*PrimeMeridianPropertyType)
		}
		stage.OnAfterPrimeMeridianPropertyTypeUpdateCallback.OnAfterUpdate(stage, primemeridianpropertytype, frontPrimeMeridianPropertyType)
	}
}

func (primemeridianpropertytype *PrimeMeridianPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimeMeridianPropertyTypeDeleteCallback != nil {
		var frontPrimeMeridianPropertyType *PrimeMeridianPropertyType
		if front != nil {
			frontPrimeMeridianPropertyType, _ = front.(*PrimeMeridianPropertyType)
		}
		stage.OnAfterPrimeMeridianPropertyTypeDeleteCallback.OnAfterDelete(stage, primemeridianpropertytype, frontPrimeMeridianPropertyType)
	}
}

func (primemeridiantype *PrimeMeridianType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPrimeMeridianTypeCreateCallback != nil {
		stage.OnAfterPrimeMeridianTypeCreateCallback.OnAfterCreate(stage, primemeridiantype)
	}
}

func (primemeridiantype *PrimeMeridianType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimeMeridianTypeUpdateCallback != nil {
		var frontPrimeMeridianType *PrimeMeridianType
		if front != nil {
			frontPrimeMeridianType, _ = front.(*PrimeMeridianType)
		}
		stage.OnAfterPrimeMeridianTypeUpdateCallback.OnAfterUpdate(stage, primemeridiantype, frontPrimeMeridianType)
	}
}

func (primemeridiantype *PrimeMeridianType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPrimeMeridianTypeDeleteCallback != nil {
		var frontPrimeMeridianType *PrimeMeridianType
		if front != nil {
			frontPrimeMeridianType, _ = front.(*PrimeMeridianType)
		}
		stage.OnAfterPrimeMeridianTypeDeleteCallback.OnAfterDelete(stage, primemeridiantype, frontPrimeMeridianType)
	}
}

func (prioritylocationpropertytype *PriorityLocationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPriorityLocationPropertyTypeCreateCallback != nil {
		stage.OnAfterPriorityLocationPropertyTypeCreateCallback.OnAfterCreate(stage, prioritylocationpropertytype)
	}
}

func (prioritylocationpropertytype *PriorityLocationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPriorityLocationPropertyTypeUpdateCallback != nil {
		var frontPriorityLocationPropertyType *PriorityLocationPropertyType
		if front != nil {
			frontPriorityLocationPropertyType, _ = front.(*PriorityLocationPropertyType)
		}
		stage.OnAfterPriorityLocationPropertyTypeUpdateCallback.OnAfterUpdate(stage, prioritylocationpropertytype, frontPriorityLocationPropertyType)
	}
}

func (prioritylocationpropertytype *PriorityLocationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPriorityLocationPropertyTypeDeleteCallback != nil {
		var frontPriorityLocationPropertyType *PriorityLocationPropertyType
		if front != nil {
			frontPriorityLocationPropertyType, _ = front.(*PriorityLocationPropertyType)
		}
		stage.OnAfterPriorityLocationPropertyTypeDeleteCallback.OnAfterDelete(stage, prioritylocationpropertytype, frontPriorityLocationPropertyType)
	}
}

func (procedurepropertytype *ProcedurePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProcedurePropertyTypeCreateCallback != nil {
		stage.OnAfterProcedurePropertyTypeCreateCallback.OnAfterCreate(stage, procedurepropertytype)
	}
}

func (procedurepropertytype *ProcedurePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedurePropertyTypeUpdateCallback != nil {
		var frontProcedurePropertyType *ProcedurePropertyType
		if front != nil {
			frontProcedurePropertyType, _ = front.(*ProcedurePropertyType)
		}
		stage.OnAfterProcedurePropertyTypeUpdateCallback.OnAfterUpdate(stage, procedurepropertytype, frontProcedurePropertyType)
	}
}

func (procedurepropertytype *ProcedurePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProcedurePropertyTypeDeleteCallback != nil {
		var frontProcedurePropertyType *ProcedurePropertyType
		if front != nil {
			frontProcedurePropertyType, _ = front.(*ProcedurePropertyType)
		}
		stage.OnAfterProcedurePropertyTypeDeleteCallback.OnAfterDelete(stage, procedurepropertytype, frontProcedurePropertyType)
	}
}

func (projectedcrspropertytype *ProjectedCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProjectedCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterProjectedCRSPropertyTypeCreateCallback.OnAfterCreate(stage, projectedcrspropertytype)
	}
}

func (projectedcrspropertytype *ProjectedCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProjectedCRSPropertyTypeUpdateCallback != nil {
		var frontProjectedCRSPropertyType *ProjectedCRSPropertyType
		if front != nil {
			frontProjectedCRSPropertyType, _ = front.(*ProjectedCRSPropertyType)
		}
		stage.OnAfterProjectedCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, projectedcrspropertytype, frontProjectedCRSPropertyType)
	}
}

func (projectedcrspropertytype *ProjectedCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProjectedCRSPropertyTypeDeleteCallback != nil {
		var frontProjectedCRSPropertyType *ProjectedCRSPropertyType
		if front != nil {
			frontProjectedCRSPropertyType, _ = front.(*ProjectedCRSPropertyType)
		}
		stage.OnAfterProjectedCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, projectedcrspropertytype, frontProjectedCRSPropertyType)
	}
}

func (projectedcrstype *ProjectedCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProjectedCRSTypeCreateCallback != nil {
		stage.OnAfterProjectedCRSTypeCreateCallback.OnAfterCreate(stage, projectedcrstype)
	}
}

func (projectedcrstype *ProjectedCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProjectedCRSTypeUpdateCallback != nil {
		var frontProjectedCRSType *ProjectedCRSType
		if front != nil {
			frontProjectedCRSType, _ = front.(*ProjectedCRSType)
		}
		stage.OnAfterProjectedCRSTypeUpdateCallback.OnAfterUpdate(stage, projectedcrstype, frontProjectedCRSType)
	}
}

func (projectedcrstype *ProjectedCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProjectedCRSTypeDeleteCallback != nil {
		var frontProjectedCRSType *ProjectedCRSType
		if front != nil {
			frontProjectedCRSType, _ = front.(*ProjectedCRSType)
		}
		stage.OnAfterProjectedCRSTypeDeleteCallback.OnAfterDelete(stage, projectedcrstype, frontProjectedCRSType)
	}
}

func (quantity *Quantity) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterQuantityCreateCallback != nil {
		stage.OnAfterQuantityCreateCallback.OnAfterCreate(stage, quantity)
	}
}

func (quantity *Quantity) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterQuantityUpdateCallback != nil {
		var frontQuantity *Quantity
		if front != nil {
			frontQuantity, _ = front.(*Quantity)
		}
		stage.OnAfterQuantityUpdateCallback.OnAfterUpdate(stage, quantity, frontQuantity)
	}
}

func (quantity *Quantity) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterQuantityDeleteCallback != nil {
		var frontQuantity *Quantity
		if front != nil {
			frontQuantity, _ = front.(*Quantity)
		}
		stage.OnAfterQuantityDeleteCallback.OnAfterDelete(stage, quantity, frontQuantity)
	}
}

func (quantityextenttype *QuantityExtentType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterQuantityExtentTypeCreateCallback != nil {
		stage.OnAfterQuantityExtentTypeCreateCallback.OnAfterCreate(stage, quantityextenttype)
	}
}

func (quantityextenttype *QuantityExtentType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterQuantityExtentTypeUpdateCallback != nil {
		var frontQuantityExtentType *QuantityExtentType
		if front != nil {
			frontQuantityExtentType, _ = front.(*QuantityExtentType)
		}
		stage.OnAfterQuantityExtentTypeUpdateCallback.OnAfterUpdate(stage, quantityextenttype, frontQuantityExtentType)
	}
}

func (quantityextenttype *QuantityExtentType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterQuantityExtentTypeDeleteCallback != nil {
		var frontQuantityExtentType *QuantityExtentType
		if front != nil {
			frontQuantityExtentType, _ = front.(*QuantityExtentType)
		}
		stage.OnAfterQuantityExtentTypeDeleteCallback.OnAfterDelete(stage, quantityextenttype, frontQuantityExtentType)
	}
}

func (quantitypropertytype *QuantityPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterQuantityPropertyTypeCreateCallback != nil {
		stage.OnAfterQuantityPropertyTypeCreateCallback.OnAfterCreate(stage, quantitypropertytype)
	}
}

func (quantitypropertytype *QuantityPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterQuantityPropertyTypeUpdateCallback != nil {
		var frontQuantityPropertyType *QuantityPropertyType
		if front != nil {
			frontQuantityPropertyType, _ = front.(*QuantityPropertyType)
		}
		stage.OnAfterQuantityPropertyTypeUpdateCallback.OnAfterUpdate(stage, quantitypropertytype, frontQuantityPropertyType)
	}
}

func (quantitypropertytype *QuantityPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterQuantityPropertyTypeDeleteCallback != nil {
		var frontQuantityPropertyType *QuantityPropertyType
		if front != nil {
			frontQuantityPropertyType, _ = front.(*QuantityPropertyType)
		}
		stage.OnAfterQuantityPropertyTypeDeleteCallback.OnAfterDelete(stage, quantitypropertytype, frontQuantityPropertyType)
	}
}

func (rs_identifier_propertytype *RS_Identifier_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRS_Identifier_PropertyTypeCreateCallback != nil {
		stage.OnAfterRS_Identifier_PropertyTypeCreateCallback.OnAfterCreate(stage, rs_identifier_propertytype)
	}
}

func (rs_identifier_propertytype *RS_Identifier_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRS_Identifier_PropertyTypeUpdateCallback != nil {
		var frontRS_Identifier_PropertyType *RS_Identifier_PropertyType
		if front != nil {
			frontRS_Identifier_PropertyType, _ = front.(*RS_Identifier_PropertyType)
		}
		stage.OnAfterRS_Identifier_PropertyTypeUpdateCallback.OnAfterUpdate(stage, rs_identifier_propertytype, frontRS_Identifier_PropertyType)
	}
}

func (rs_identifier_propertytype *RS_Identifier_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRS_Identifier_PropertyTypeDeleteCallback != nil {
		var frontRS_Identifier_PropertyType *RS_Identifier_PropertyType
		if front != nil {
			frontRS_Identifier_PropertyType, _ = front.(*RS_Identifier_PropertyType)
		}
		stage.OnAfterRS_Identifier_PropertyTypeDeleteCallback.OnAfterDelete(stage, rs_identifier_propertytype, frontRS_Identifier_PropertyType)
	}
}

func (rs_identifier_type *RS_Identifier_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRS_Identifier_TypeCreateCallback != nil {
		stage.OnAfterRS_Identifier_TypeCreateCallback.OnAfterCreate(stage, rs_identifier_type)
	}
}

func (rs_identifier_type *RS_Identifier_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRS_Identifier_TypeUpdateCallback != nil {
		var frontRS_Identifier_Type *RS_Identifier_Type
		if front != nil {
			frontRS_Identifier_Type, _ = front.(*RS_Identifier_Type)
		}
		stage.OnAfterRS_Identifier_TypeUpdateCallback.OnAfterUpdate(stage, rs_identifier_type, frontRS_Identifier_Type)
	}
}

func (rs_identifier_type *RS_Identifier_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRS_Identifier_TypeDeleteCallback != nil {
		var frontRS_Identifier_Type *RS_Identifier_Type
		if front != nil {
			frontRS_Identifier_Type, _ = front.(*RS_Identifier_Type)
		}
		stage.OnAfterRS_Identifier_TypeDeleteCallback.OnAfterDelete(stage, rs_identifier_type, frontRS_Identifier_Type)
	}
}

func (rs_referencesystem_propertytype *RS_ReferenceSystem_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRS_ReferenceSystem_PropertyTypeCreateCallback != nil {
		stage.OnAfterRS_ReferenceSystem_PropertyTypeCreateCallback.OnAfterCreate(stage, rs_referencesystem_propertytype)
	}
}

func (rs_referencesystem_propertytype *RS_ReferenceSystem_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRS_ReferenceSystem_PropertyTypeUpdateCallback != nil {
		var frontRS_ReferenceSystem_PropertyType *RS_ReferenceSystem_PropertyType
		if front != nil {
			frontRS_ReferenceSystem_PropertyType, _ = front.(*RS_ReferenceSystem_PropertyType)
		}
		stage.OnAfterRS_ReferenceSystem_PropertyTypeUpdateCallback.OnAfterUpdate(stage, rs_referencesystem_propertytype, frontRS_ReferenceSystem_PropertyType)
	}
}

func (rs_referencesystem_propertytype *RS_ReferenceSystem_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRS_ReferenceSystem_PropertyTypeDeleteCallback != nil {
		var frontRS_ReferenceSystem_PropertyType *RS_ReferenceSystem_PropertyType
		if front != nil {
			frontRS_ReferenceSystem_PropertyType, _ = front.(*RS_ReferenceSystem_PropertyType)
		}
		stage.OnAfterRS_ReferenceSystem_PropertyTypeDeleteCallback.OnAfterDelete(stage, rs_referencesystem_propertytype, frontRS_ReferenceSystem_PropertyType)
	}
}

func (rangesettype *RangeSetType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRangeSetTypeCreateCallback != nil {
		stage.OnAfterRangeSetTypeCreateCallback.OnAfterCreate(stage, rangesettype)
	}
}

func (rangesettype *RangeSetType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRangeSetTypeUpdateCallback != nil {
		var frontRangeSetType *RangeSetType
		if front != nil {
			frontRangeSetType, _ = front.(*RangeSetType)
		}
		stage.OnAfterRangeSetTypeUpdateCallback.OnAfterUpdate(stage, rangesettype, frontRangeSetType)
	}
}

func (rangesettype *RangeSetType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRangeSetTypeDeleteCallback != nil {
		var frontRangeSetType *RangeSetType
		if front != nil {
			frontRangeSetType, _ = front.(*RangeSetType)
		}
		stage.OnAfterRangeSetTypeDeleteCallback.OnAfterDelete(stage, rangesettype, frontRangeSetType)
	}
}

func (real_propertytype *Real_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterReal_PropertyTypeCreateCallback != nil {
		stage.OnAfterReal_PropertyTypeCreateCallback.OnAfterCreate(stage, real_propertytype)
	}
}

func (real_propertytype *Real_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReal_PropertyTypeUpdateCallback != nil {
		var frontReal_PropertyType *Real_PropertyType
		if front != nil {
			frontReal_PropertyType, _ = front.(*Real_PropertyType)
		}
		stage.OnAfterReal_PropertyTypeUpdateCallback.OnAfterUpdate(stage, real_propertytype, frontReal_PropertyType)
	}
}

func (real_propertytype *Real_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReal_PropertyTypeDeleteCallback != nil {
		var frontReal_PropertyType *Real_PropertyType
		if front != nil {
			frontReal_PropertyType, _ = front.(*Real_PropertyType)
		}
		stage.OnAfterReal_PropertyTypeDeleteCallback.OnAfterDelete(stage, real_propertytype, frontReal_PropertyType)
	}
}

func (recordtype_propertytype *RecordType_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRecordType_PropertyTypeCreateCallback != nil {
		stage.OnAfterRecordType_PropertyTypeCreateCallback.OnAfterCreate(stage, recordtype_propertytype)
	}
}

func (recordtype_propertytype *RecordType_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRecordType_PropertyTypeUpdateCallback != nil {
		var frontRecordType_PropertyType *RecordType_PropertyType
		if front != nil {
			frontRecordType_PropertyType, _ = front.(*RecordType_PropertyType)
		}
		stage.OnAfterRecordType_PropertyTypeUpdateCallback.OnAfterUpdate(stage, recordtype_propertytype, frontRecordType_PropertyType)
	}
}

func (recordtype_propertytype *RecordType_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRecordType_PropertyTypeDeleteCallback != nil {
		var frontRecordType_PropertyType *RecordType_PropertyType
		if front != nil {
			frontRecordType_PropertyType, _ = front.(*RecordType_PropertyType)
		}
		stage.OnAfterRecordType_PropertyTypeDeleteCallback.OnAfterDelete(stage, recordtype_propertytype, frontRecordType_PropertyType)
	}
}

func (recordtype_type *RecordType_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRecordType_TypeCreateCallback != nil {
		stage.OnAfterRecordType_TypeCreateCallback.OnAfterCreate(stage, recordtype_type)
	}
}

func (recordtype_type *RecordType_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRecordType_TypeUpdateCallback != nil {
		var frontRecordType_Type *RecordType_Type
		if front != nil {
			frontRecordType_Type, _ = front.(*RecordType_Type)
		}
		stage.OnAfterRecordType_TypeUpdateCallback.OnAfterUpdate(stage, recordtype_type, frontRecordType_Type)
	}
}

func (recordtype_type *RecordType_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRecordType_TypeDeleteCallback != nil {
		var frontRecordType_Type *RecordType_Type
		if front != nil {
			frontRecordType_Type, _ = front.(*RecordType_Type)
		}
		stage.OnAfterRecordType_TypeDeleteCallback.OnAfterDelete(stage, recordtype_type, frontRecordType_Type)
	}
}

func (record_propertytype *Record_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRecord_PropertyTypeCreateCallback != nil {
		stage.OnAfterRecord_PropertyTypeCreateCallback.OnAfterCreate(stage, record_propertytype)
	}
}

func (record_propertytype *Record_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRecord_PropertyTypeUpdateCallback != nil {
		var frontRecord_PropertyType *Record_PropertyType
		if front != nil {
			frontRecord_PropertyType, _ = front.(*Record_PropertyType)
		}
		stage.OnAfterRecord_PropertyTypeUpdateCallback.OnAfterUpdate(stage, record_propertytype, frontRecord_PropertyType)
	}
}

func (record_propertytype *Record_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRecord_PropertyTypeDeleteCallback != nil {
		var frontRecord_PropertyType *Record_PropertyType
		if front != nil {
			frontRecord_PropertyType, _ = front.(*Record_PropertyType)
		}
		stage.OnAfterRecord_PropertyTypeDeleteCallback.OnAfterDelete(stage, record_propertytype, frontRecord_PropertyType)
	}
}

func (rectangletype *RectangleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectangleTypeCreateCallback != nil {
		stage.OnAfterRectangleTypeCreateCallback.OnAfterCreate(stage, rectangletype)
	}
}

func (rectangletype *RectangleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectangleTypeUpdateCallback != nil {
		var frontRectangleType *RectangleType
		if front != nil {
			frontRectangleType, _ = front.(*RectangleType)
		}
		stage.OnAfterRectangleTypeUpdateCallback.OnAfterUpdate(stage, rectangletype, frontRectangleType)
	}
}

func (rectangletype *RectangleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectangleTypeDeleteCallback != nil {
		var frontRectangleType *RectangleType
		if front != nil {
			frontRectangleType, _ = front.(*RectangleType)
		}
		stage.OnAfterRectangleTypeDeleteCallback.OnAfterDelete(stage, rectangletype, frontRectangleType)
	}
}

func (rectifiedgridtype *RectifiedGridType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRectifiedGridTypeCreateCallback != nil {
		stage.OnAfterRectifiedGridTypeCreateCallback.OnAfterCreate(stage, rectifiedgridtype)
	}
}

func (rectifiedgridtype *RectifiedGridType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectifiedGridTypeUpdateCallback != nil {
		var frontRectifiedGridType *RectifiedGridType
		if front != nil {
			frontRectifiedGridType, _ = front.(*RectifiedGridType)
		}
		stage.OnAfterRectifiedGridTypeUpdateCallback.OnAfterUpdate(stage, rectifiedgridtype, frontRectifiedGridType)
	}
}

func (rectifiedgridtype *RectifiedGridType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRectifiedGridTypeDeleteCallback != nil {
		var frontRectifiedGridType *RectifiedGridType
		if front != nil {
			frontRectifiedGridType, _ = front.(*RectifiedGridType)
		}
		stage.OnAfterRectifiedGridTypeDeleteCallback.OnAfterDelete(stage, rectifiedgridtype, frontRectifiedGridType)
	}
}

func (referencetype *ReferenceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterReferenceTypeCreateCallback != nil {
		stage.OnAfterReferenceTypeCreateCallback.OnAfterCreate(stage, referencetype)
	}
}

func (referencetype *ReferenceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReferenceTypeUpdateCallback != nil {
		var frontReferenceType *ReferenceType
		if front != nil {
			frontReferenceType, _ = front.(*ReferenceType)
		}
		stage.OnAfterReferenceTypeUpdateCallback.OnAfterUpdate(stage, referencetype, frontReferenceType)
	}
}

func (referencetype *ReferenceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterReferenceTypeDeleteCallback != nil {
		var frontReferenceType *ReferenceType
		if front != nil {
			frontReferenceType, _ = front.(*ReferenceType)
		}
		stage.OnAfterReferenceTypeDeleteCallback.OnAfterDelete(stage, referencetype, frontReferenceType)
	}
}

func (relatedtimetype *RelatedTimeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRelatedTimeTypeCreateCallback != nil {
		stage.OnAfterRelatedTimeTypeCreateCallback.OnAfterCreate(stage, relatedtimetype)
	}
}

func (relatedtimetype *RelatedTimeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRelatedTimeTypeUpdateCallback != nil {
		var frontRelatedTimeType *RelatedTimeType
		if front != nil {
			frontRelatedTimeType, _ = front.(*RelatedTimeType)
		}
		stage.OnAfterRelatedTimeTypeUpdateCallback.OnAfterUpdate(stage, relatedtimetype, frontRelatedTimeType)
	}
}

func (relatedtimetype *RelatedTimeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRelatedTimeTypeDeleteCallback != nil {
		var frontRelatedTimeType *RelatedTimeType
		if front != nil {
			frontRelatedTimeType, _ = front.(*RelatedTimeType)
		}
		stage.OnAfterRelatedTimeTypeDeleteCallback.OnAfterDelete(stage, relatedtimetype, frontRelatedTimeType)
	}
}

func (resulttype *ResultType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResultTypeCreateCallback != nil {
		stage.OnAfterResultTypeCreateCallback.OnAfterCreate(stage, resulttype)
	}
}

func (resulttype *ResultType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResultTypeUpdateCallback != nil {
		var frontResultType *ResultType
		if front != nil {
			frontResultType, _ = front.(*ResultType)
		}
		stage.OnAfterResultTypeUpdateCallback.OnAfterUpdate(stage, resulttype, frontResultType)
	}
}

func (resulttype *ResultType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResultTypeDeleteCallback != nil {
		var frontResultType *ResultType
		if front != nil {
			frontResultType, _ = front.(*ResultType)
		}
		stage.OnAfterResultTypeDeleteCallback.OnAfterDelete(stage, resulttype, frontResultType)
	}
}

func (ringpropertytype *RingPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRingPropertyTypeCreateCallback != nil {
		stage.OnAfterRingPropertyTypeCreateCallback.OnAfterCreate(stage, ringpropertytype)
	}
}

func (ringpropertytype *RingPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRingPropertyTypeUpdateCallback != nil {
		var frontRingPropertyType *RingPropertyType
		if front != nil {
			frontRingPropertyType, _ = front.(*RingPropertyType)
		}
		stage.OnAfterRingPropertyTypeUpdateCallback.OnAfterUpdate(stage, ringpropertytype, frontRingPropertyType)
	}
}

func (ringpropertytype *RingPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRingPropertyTypeDeleteCallback != nil {
		var frontRingPropertyType *RingPropertyType
		if front != nil {
			frontRingPropertyType, _ = front.(*RingPropertyType)
		}
		stage.OnAfterRingPropertyTypeDeleteCallback.OnAfterDelete(stage, ringpropertytype, frontRingPropertyType)
	}
}

func (ringtype *RingType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRingTypeCreateCallback != nil {
		stage.OnAfterRingTypeCreateCallback.OnAfterCreate(stage, ringtype)
	}
}

func (ringtype *RingType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRingTypeUpdateCallback != nil {
		var frontRingType *RingType
		if front != nil {
			frontRingType, _ = front.(*RingType)
		}
		stage.OnAfterRingTypeUpdateCallback.OnAfterUpdate(stage, ringtype, frontRingType)
	}
}

func (ringtype *RingType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRingTypeDeleteCallback != nil {
		var frontRingType *RingType
		if front != nil {
			frontRingType, _ = front.(*RingType)
		}
		stage.OnAfterRingTypeDeleteCallback.OnAfterDelete(stage, ringtype, frontRingType)
	}
}

func (sc_crs_propertytype *SC_CRS_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSC_CRS_PropertyTypeCreateCallback != nil {
		stage.OnAfterSC_CRS_PropertyTypeCreateCallback.OnAfterCreate(stage, sc_crs_propertytype)
	}
}

func (sc_crs_propertytype *SC_CRS_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSC_CRS_PropertyTypeUpdateCallback != nil {
		var frontSC_CRS_PropertyType *SC_CRS_PropertyType
		if front != nil {
			frontSC_CRS_PropertyType, _ = front.(*SC_CRS_PropertyType)
		}
		stage.OnAfterSC_CRS_PropertyTypeUpdateCallback.OnAfterUpdate(stage, sc_crs_propertytype, frontSC_CRS_PropertyType)
	}
}

func (sc_crs_propertytype *SC_CRS_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSC_CRS_PropertyTypeDeleteCallback != nil {
		var frontSC_CRS_PropertyType *SC_CRS_PropertyType
		if front != nil {
			frontSC_CRS_PropertyType, _ = front.(*SC_CRS_PropertyType)
		}
		stage.OnAfterSC_CRS_PropertyTypeDeleteCallback.OnAfterDelete(stage, sc_crs_propertytype, frontSC_CRS_PropertyType)
	}
}

func (scale_propertytype *Scale_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScale_PropertyTypeCreateCallback != nil {
		stage.OnAfterScale_PropertyTypeCreateCallback.OnAfterCreate(stage, scale_propertytype)
	}
}

func (scale_propertytype *Scale_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScale_PropertyTypeUpdateCallback != nil {
		var frontScale_PropertyType *Scale_PropertyType
		if front != nil {
			frontScale_PropertyType, _ = front.(*Scale_PropertyType)
		}
		stage.OnAfterScale_PropertyTypeUpdateCallback.OnAfterUpdate(stage, scale_propertytype, frontScale_PropertyType)
	}
}

func (scale_propertytype *Scale_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScale_PropertyTypeDeleteCallback != nil {
		var frontScale_PropertyType *Scale_PropertyType
		if front != nil {
			frontScale_PropertyType, _ = front.(*Scale_PropertyType)
		}
		stage.OnAfterScale_PropertyTypeDeleteCallback.OnAfterDelete(stage, scale_propertytype, frontScale_PropertyType)
	}
}

func (scopedname_propertytype *ScopedName_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterScopedName_PropertyTypeCreateCallback != nil {
		stage.OnAfterScopedName_PropertyTypeCreateCallback.OnAfterCreate(stage, scopedname_propertytype)
	}
}

func (scopedname_propertytype *ScopedName_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScopedName_PropertyTypeUpdateCallback != nil {
		var frontScopedName_PropertyType *ScopedName_PropertyType
		if front != nil {
			frontScopedName_PropertyType, _ = front.(*ScopedName_PropertyType)
		}
		stage.OnAfterScopedName_PropertyTypeUpdateCallback.OnAfterUpdate(stage, scopedname_propertytype, frontScopedName_PropertyType)
	}
}

func (scopedname_propertytype *ScopedName_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterScopedName_PropertyTypeDeleteCallback != nil {
		var frontScopedName_PropertyType *ScopedName_PropertyType
		if front != nil {
			frontScopedName_PropertyType, _ = front.(*ScopedName_PropertyType)
		}
		stage.OnAfterScopedName_PropertyTypeDeleteCallback.OnAfterDelete(stage, scopedname_propertytype, frontScopedName_PropertyType)
	}
}

func (seconddefiningparameter *SecondDefiningParameter) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSecondDefiningParameterCreateCallback != nil {
		stage.OnAfterSecondDefiningParameterCreateCallback.OnAfterCreate(stage, seconddefiningparameter)
	}
}

func (seconddefiningparameter *SecondDefiningParameter) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondDefiningParameterUpdateCallback != nil {
		var frontSecondDefiningParameter *SecondDefiningParameter
		if front != nil {
			frontSecondDefiningParameter, _ = front.(*SecondDefiningParameter)
		}
		stage.OnAfterSecondDefiningParameterUpdateCallback.OnAfterUpdate(stage, seconddefiningparameter, frontSecondDefiningParameter)
	}
}

func (seconddefiningparameter *SecondDefiningParameter) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondDefiningParameterDeleteCallback != nil {
		var frontSecondDefiningParameter *SecondDefiningParameter
		if front != nil {
			frontSecondDefiningParameter, _ = front.(*SecondDefiningParameter)
		}
		stage.OnAfterSecondDefiningParameterDeleteCallback.OnAfterDelete(stage, seconddefiningparameter, frontSecondDefiningParameter)
	}
}

func (seconddefiningparameter_1 *SecondDefiningParameter_1) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSecondDefiningParameter_1CreateCallback != nil {
		stage.OnAfterSecondDefiningParameter_1CreateCallback.OnAfterCreate(stage, seconddefiningparameter_1)
	}
}

func (seconddefiningparameter_1 *SecondDefiningParameter_1) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondDefiningParameter_1UpdateCallback != nil {
		var frontSecondDefiningParameter_1 *SecondDefiningParameter_1
		if front != nil {
			frontSecondDefiningParameter_1, _ = front.(*SecondDefiningParameter_1)
		}
		stage.OnAfterSecondDefiningParameter_1UpdateCallback.OnAfterUpdate(stage, seconddefiningparameter_1, frontSecondDefiningParameter_1)
	}
}

func (seconddefiningparameter_1 *SecondDefiningParameter_1) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSecondDefiningParameter_1DeleteCallback != nil {
		var frontSecondDefiningParameter_1 *SecondDefiningParameter_1
		if front != nil {
			frontSecondDefiningParameter_1, _ = front.(*SecondDefiningParameter_1)
		}
		stage.OnAfterSecondDefiningParameter_1DeleteCallback.OnAfterDelete(stage, seconddefiningparameter_1, frontSecondDefiningParameter_1)
	}
}

func (sequenceruletype *SequenceRuleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSequenceRuleTypeCreateCallback != nil {
		stage.OnAfterSequenceRuleTypeCreateCallback.OnAfterCreate(stage, sequenceruletype)
	}
}

func (sequenceruletype *SequenceRuleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSequenceRuleTypeUpdateCallback != nil {
		var frontSequenceRuleType *SequenceRuleType
		if front != nil {
			frontSequenceRuleType, _ = front.(*SequenceRuleType)
		}
		stage.OnAfterSequenceRuleTypeUpdateCallback.OnAfterUpdate(stage, sequenceruletype, frontSequenceRuleType)
	}
}

func (sequenceruletype *SequenceRuleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSequenceRuleTypeDeleteCallback != nil {
		var frontSequenceRuleType *SequenceRuleType
		if front != nil {
			frontSequenceRuleType, _ = front.(*SequenceRuleType)
		}
		stage.OnAfterSequenceRuleTypeDeleteCallback.OnAfterDelete(stage, sequenceruletype, frontSequenceRuleType)
	}
}

func (shellpropertytype *ShellPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShellPropertyTypeCreateCallback != nil {
		stage.OnAfterShellPropertyTypeCreateCallback.OnAfterCreate(stage, shellpropertytype)
	}
}

func (shellpropertytype *ShellPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShellPropertyTypeUpdateCallback != nil {
		var frontShellPropertyType *ShellPropertyType
		if front != nil {
			frontShellPropertyType, _ = front.(*ShellPropertyType)
		}
		stage.OnAfterShellPropertyTypeUpdateCallback.OnAfterUpdate(stage, shellpropertytype, frontShellPropertyType)
	}
}

func (shellpropertytype *ShellPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShellPropertyTypeDeleteCallback != nil {
		var frontShellPropertyType *ShellPropertyType
		if front != nil {
			frontShellPropertyType, _ = front.(*ShellPropertyType)
		}
		stage.OnAfterShellPropertyTypeDeleteCallback.OnAfterDelete(stage, shellpropertytype, frontShellPropertyType)
	}
}

func (shelltype *ShellType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShellTypeCreateCallback != nil {
		stage.OnAfterShellTypeCreateCallback.OnAfterCreate(stage, shelltype)
	}
}

func (shelltype *ShellType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShellTypeUpdateCallback != nil {
		var frontShellType *ShellType
		if front != nil {
			frontShellType, _ = front.(*ShellType)
		}
		stage.OnAfterShellTypeUpdateCallback.OnAfterUpdate(stage, shelltype, frontShellType)
	}
}

func (shelltype *ShellType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShellTypeDeleteCallback != nil {
		var frontShellType *ShellType
		if front != nil {
			frontShellType, _ = front.(*ShellType)
		}
		stage.OnAfterShellTypeDeleteCallback.OnAfterDelete(stage, shelltype, frontShellType)
	}
}

func (singlecrspropertytype *SingleCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSingleCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterSingleCRSPropertyTypeCreateCallback.OnAfterCreate(stage, singlecrspropertytype)
	}
}

func (singlecrspropertytype *SingleCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSingleCRSPropertyTypeUpdateCallback != nil {
		var frontSingleCRSPropertyType *SingleCRSPropertyType
		if front != nil {
			frontSingleCRSPropertyType, _ = front.(*SingleCRSPropertyType)
		}
		stage.OnAfterSingleCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, singlecrspropertytype, frontSingleCRSPropertyType)
	}
}

func (singlecrspropertytype *SingleCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSingleCRSPropertyTypeDeleteCallback != nil {
		var frontSingleCRSPropertyType *SingleCRSPropertyType
		if front != nil {
			frontSingleCRSPropertyType, _ = front.(*SingleCRSPropertyType)
		}
		stage.OnAfterSingleCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, singlecrspropertytype, frontSingleCRSPropertyType)
	}
}

func (singleoperationpropertytype *SingleOperationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSingleOperationPropertyTypeCreateCallback != nil {
		stage.OnAfterSingleOperationPropertyTypeCreateCallback.OnAfterCreate(stage, singleoperationpropertytype)
	}
}

func (singleoperationpropertytype *SingleOperationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSingleOperationPropertyTypeUpdateCallback != nil {
		var frontSingleOperationPropertyType *SingleOperationPropertyType
		if front != nil {
			frontSingleOperationPropertyType, _ = front.(*SingleOperationPropertyType)
		}
		stage.OnAfterSingleOperationPropertyTypeUpdateCallback.OnAfterUpdate(stage, singleoperationpropertytype, frontSingleOperationPropertyType)
	}
}

func (singleoperationpropertytype *SingleOperationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSingleOperationPropertyTypeDeleteCallback != nil {
		var frontSingleOperationPropertyType *SingleOperationPropertyType
		if front != nil {
			frontSingleOperationPropertyType, _ = front.(*SingleOperationPropertyType)
		}
		stage.OnAfterSingleOperationPropertyTypeDeleteCallback.OnAfterDelete(stage, singleoperationpropertytype, frontSingleOperationPropertyType)
	}
}

func (solidarraypropertytype *SolidArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSolidArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterSolidArrayPropertyTypeCreateCallback.OnAfterCreate(stage, solidarraypropertytype)
	}
}

func (solidarraypropertytype *SolidArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSolidArrayPropertyTypeUpdateCallback != nil {
		var frontSolidArrayPropertyType *SolidArrayPropertyType
		if front != nil {
			frontSolidArrayPropertyType, _ = front.(*SolidArrayPropertyType)
		}
		stage.OnAfterSolidArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, solidarraypropertytype, frontSolidArrayPropertyType)
	}
}

func (solidarraypropertytype *SolidArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSolidArrayPropertyTypeDeleteCallback != nil {
		var frontSolidArrayPropertyType *SolidArrayPropertyType
		if front != nil {
			frontSolidArrayPropertyType, _ = front.(*SolidArrayPropertyType)
		}
		stage.OnAfterSolidArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, solidarraypropertytype, frontSolidArrayPropertyType)
	}
}

func (solidpropertytype *SolidPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSolidPropertyTypeCreateCallback != nil {
		stage.OnAfterSolidPropertyTypeCreateCallback.OnAfterCreate(stage, solidpropertytype)
	}
}

func (solidpropertytype *SolidPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSolidPropertyTypeUpdateCallback != nil {
		var frontSolidPropertyType *SolidPropertyType
		if front != nil {
			frontSolidPropertyType, _ = front.(*SolidPropertyType)
		}
		stage.OnAfterSolidPropertyTypeUpdateCallback.OnAfterUpdate(stage, solidpropertytype, frontSolidPropertyType)
	}
}

func (solidpropertytype *SolidPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSolidPropertyTypeDeleteCallback != nil {
		var frontSolidPropertyType *SolidPropertyType
		if front != nil {
			frontSolidPropertyType, _ = front.(*SolidPropertyType)
		}
		stage.OnAfterSolidPropertyTypeDeleteCallback.OnAfterDelete(stage, solidpropertytype, frontSolidPropertyType)
	}
}

func (solidtype *SolidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSolidTypeCreateCallback != nil {
		stage.OnAfterSolidTypeCreateCallback.OnAfterCreate(stage, solidtype)
	}
}

func (solidtype *SolidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSolidTypeUpdateCallback != nil {
		var frontSolidType *SolidType
		if front != nil {
			frontSolidType, _ = front.(*SolidType)
		}
		stage.OnAfterSolidTypeUpdateCallback.OnAfterUpdate(stage, solidtype, frontSolidType)
	}
}

func (solidtype *SolidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSolidTypeDeleteCallback != nil {
		var frontSolidType *SolidType
		if front != nil {
			frontSolidType, _ = front.(*SolidType)
		}
		stage.OnAfterSolidTypeDeleteCallback.OnAfterDelete(stage, solidtype, frontSolidType)
	}
}

func (spheretype *SphereType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSphereTypeCreateCallback != nil {
		stage.OnAfterSphereTypeCreateCallback.OnAfterCreate(stage, spheretype)
	}
}

func (spheretype *SphereType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphereTypeUpdateCallback != nil {
		var frontSphereType *SphereType
		if front != nil {
			frontSphereType, _ = front.(*SphereType)
		}
		stage.OnAfterSphereTypeUpdateCallback.OnAfterUpdate(stage, spheretype, frontSphereType)
	}
}

func (spheretype *SphereType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphereTypeDeleteCallback != nil {
		var frontSphereType *SphereType
		if front != nil {
			frontSphereType, _ = front.(*SphereType)
		}
		stage.OnAfterSphereTypeDeleteCallback.OnAfterDelete(stage, spheretype, frontSphereType)
	}
}

func (sphericalcspropertytype *SphericalCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSphericalCSPropertyTypeCreateCallback != nil {
		stage.OnAfterSphericalCSPropertyTypeCreateCallback.OnAfterCreate(stage, sphericalcspropertytype)
	}
}

func (sphericalcspropertytype *SphericalCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphericalCSPropertyTypeUpdateCallback != nil {
		var frontSphericalCSPropertyType *SphericalCSPropertyType
		if front != nil {
			frontSphericalCSPropertyType, _ = front.(*SphericalCSPropertyType)
		}
		stage.OnAfterSphericalCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, sphericalcspropertytype, frontSphericalCSPropertyType)
	}
}

func (sphericalcspropertytype *SphericalCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphericalCSPropertyTypeDeleteCallback != nil {
		var frontSphericalCSPropertyType *SphericalCSPropertyType
		if front != nil {
			frontSphericalCSPropertyType, _ = front.(*SphericalCSPropertyType)
		}
		stage.OnAfterSphericalCSPropertyTypeDeleteCallback.OnAfterDelete(stage, sphericalcspropertytype, frontSphericalCSPropertyType)
	}
}

func (sphericalcstype *SphericalCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSphericalCSTypeCreateCallback != nil {
		stage.OnAfterSphericalCSTypeCreateCallback.OnAfterCreate(stage, sphericalcstype)
	}
}

func (sphericalcstype *SphericalCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphericalCSTypeUpdateCallback != nil {
		var frontSphericalCSType *SphericalCSType
		if front != nil {
			frontSphericalCSType, _ = front.(*SphericalCSType)
		}
		stage.OnAfterSphericalCSTypeUpdateCallback.OnAfterUpdate(stage, sphericalcstype, frontSphericalCSType)
	}
}

func (sphericalcstype *SphericalCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSphericalCSTypeDeleteCallback != nil {
		var frontSphericalCSType *SphericalCSType
		if front != nil {
			frontSphericalCSType, _ = front.(*SphericalCSType)
		}
		stage.OnAfterSphericalCSTypeDeleteCallback.OnAfterDelete(stage, sphericalcstype, frontSphericalCSType)
	}
}

func (stringorreftype *StringOrRefType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStringOrRefTypeCreateCallback != nil {
		stage.OnAfterStringOrRefTypeCreateCallback.OnAfterCreate(stage, stringorreftype)
	}
}

func (stringorreftype *StringOrRefType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStringOrRefTypeUpdateCallback != nil {
		var frontStringOrRefType *StringOrRefType
		if front != nil {
			frontStringOrRefType, _ = front.(*StringOrRefType)
		}
		stage.OnAfterStringOrRefTypeUpdateCallback.OnAfterUpdate(stage, stringorreftype, frontStringOrRefType)
	}
}

func (stringorreftype *StringOrRefType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStringOrRefTypeDeleteCallback != nil {
		var frontStringOrRefType *StringOrRefType
		if front != nil {
			frontStringOrRefType, _ = front.(*StringOrRefType)
		}
		stage.OnAfterStringOrRefTypeDeleteCallback.OnAfterDelete(stage, stringorreftype, frontStringOrRefType)
	}
}

func (surfacearraypropertytype *SurfaceArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterSurfaceArrayPropertyTypeCreateCallback.OnAfterCreate(stage, surfacearraypropertytype)
	}
}

func (surfacearraypropertytype *SurfaceArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceArrayPropertyTypeUpdateCallback != nil {
		var frontSurfaceArrayPropertyType *SurfaceArrayPropertyType
		if front != nil {
			frontSurfaceArrayPropertyType, _ = front.(*SurfaceArrayPropertyType)
		}
		stage.OnAfterSurfaceArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacearraypropertytype, frontSurfaceArrayPropertyType)
	}
}

func (surfacearraypropertytype *SurfaceArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceArrayPropertyTypeDeleteCallback != nil {
		var frontSurfaceArrayPropertyType *SurfaceArrayPropertyType
		if front != nil {
			frontSurfaceArrayPropertyType, _ = front.(*SurfaceArrayPropertyType)
		}
		stage.OnAfterSurfaceArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, surfacearraypropertytype, frontSurfaceArrayPropertyType)
	}
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfacePatchArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterSurfacePatchArrayPropertyTypeCreateCallback.OnAfterCreate(stage, surfacepatcharraypropertytype)
	}
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePatchArrayPropertyTypeUpdateCallback != nil {
		var frontSurfacePatchArrayPropertyType *SurfacePatchArrayPropertyType
		if front != nil {
			frontSurfacePatchArrayPropertyType, _ = front.(*SurfacePatchArrayPropertyType)
		}
		stage.OnAfterSurfacePatchArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacepatcharraypropertytype, frontSurfacePatchArrayPropertyType)
	}
}

func (surfacepatcharraypropertytype *SurfacePatchArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePatchArrayPropertyTypeDeleteCallback != nil {
		var frontSurfacePatchArrayPropertyType *SurfacePatchArrayPropertyType
		if front != nil {
			frontSurfacePatchArrayPropertyType, _ = front.(*SurfacePatchArrayPropertyType)
		}
		stage.OnAfterSurfacePatchArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, surfacepatcharraypropertytype, frontSurfacePatchArrayPropertyType)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfacePropertyTypeCreateCallback != nil {
		stage.OnAfterSurfacePropertyTypeCreateCallback.OnAfterCreate(stage, surfacepropertytype)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePropertyTypeUpdateCallback != nil {
		var frontSurfacePropertyType *SurfacePropertyType
		if front != nil {
			frontSurfacePropertyType, _ = front.(*SurfacePropertyType)
		}
		stage.OnAfterSurfacePropertyTypeUpdateCallback.OnAfterUpdate(stage, surfacepropertytype, frontSurfacePropertyType)
	}
}

func (surfacepropertytype *SurfacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfacePropertyTypeDeleteCallback != nil {
		var frontSurfacePropertyType *SurfacePropertyType
		if front != nil {
			frontSurfacePropertyType, _ = front.(*SurfacePropertyType)
		}
		stage.OnAfterSurfacePropertyTypeDeleteCallback.OnAfterDelete(stage, surfacepropertytype, frontSurfacePropertyType)
	}
}

func (surfacetype *SurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSurfaceTypeCreateCallback != nil {
		stage.OnAfterSurfaceTypeCreateCallback.OnAfterCreate(stage, surfacetype)
	}
}

func (surfacetype *SurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeUpdateCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeUpdateCallback.OnAfterUpdate(stage, surfacetype, frontSurfaceType)
	}
}

func (surfacetype *SurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSurfaceTypeDeleteCallback != nil {
		var frontSurfaceType *SurfaceType
		if front != nil {
			frontSurfaceType, _ = front.(*SurfaceType)
		}
		stage.OnAfterSurfaceTypeDeleteCallback.OnAfterDelete(stage, surfacetype, frontSurfaceType)
	}
}

func (tm_periodduration_propertytype *TM_PeriodDuration_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTM_PeriodDuration_PropertyTypeCreateCallback != nil {
		stage.OnAfterTM_PeriodDuration_PropertyTypeCreateCallback.OnAfterCreate(stage, tm_periodduration_propertytype)
	}
}

func (tm_periodduration_propertytype *TM_PeriodDuration_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTM_PeriodDuration_PropertyTypeUpdateCallback != nil {
		var frontTM_PeriodDuration_PropertyType *TM_PeriodDuration_PropertyType
		if front != nil {
			frontTM_PeriodDuration_PropertyType, _ = front.(*TM_PeriodDuration_PropertyType)
		}
		stage.OnAfterTM_PeriodDuration_PropertyTypeUpdateCallback.OnAfterUpdate(stage, tm_periodduration_propertytype, frontTM_PeriodDuration_PropertyType)
	}
}

func (tm_periodduration_propertytype *TM_PeriodDuration_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTM_PeriodDuration_PropertyTypeDeleteCallback != nil {
		var frontTM_PeriodDuration_PropertyType *TM_PeriodDuration_PropertyType
		if front != nil {
			frontTM_PeriodDuration_PropertyType, _ = front.(*TM_PeriodDuration_PropertyType)
		}
		stage.OnAfterTM_PeriodDuration_PropertyTypeDeleteCallback.OnAfterDelete(stage, tm_periodduration_propertytype, frontTM_PeriodDuration_PropertyType)
	}
}

func (tm_primitive_propertytype *TM_Primitive_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTM_Primitive_PropertyTypeCreateCallback != nil {
		stage.OnAfterTM_Primitive_PropertyTypeCreateCallback.OnAfterCreate(stage, tm_primitive_propertytype)
	}
}

func (tm_primitive_propertytype *TM_Primitive_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTM_Primitive_PropertyTypeUpdateCallback != nil {
		var frontTM_Primitive_PropertyType *TM_Primitive_PropertyType
		if front != nil {
			frontTM_Primitive_PropertyType, _ = front.(*TM_Primitive_PropertyType)
		}
		stage.OnAfterTM_Primitive_PropertyTypeUpdateCallback.OnAfterUpdate(stage, tm_primitive_propertytype, frontTM_Primitive_PropertyType)
	}
}

func (tm_primitive_propertytype *TM_Primitive_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTM_Primitive_PropertyTypeDeleteCallback != nil {
		var frontTM_Primitive_PropertyType *TM_Primitive_PropertyType
		if front != nil {
			frontTM_Primitive_PropertyType, _ = front.(*TM_Primitive_PropertyType)
		}
		stage.OnAfterTM_Primitive_PropertyTypeDeleteCallback.OnAfterDelete(stage, tm_primitive_propertytype, frontTM_Primitive_PropertyType)
	}
}

func (targetpropertytype *TargetPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTargetPropertyTypeCreateCallback != nil {
		stage.OnAfterTargetPropertyTypeCreateCallback.OnAfterCreate(stage, targetpropertytype)
	}
}

func (targetpropertytype *TargetPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTargetPropertyTypeUpdateCallback != nil {
		var frontTargetPropertyType *TargetPropertyType
		if front != nil {
			frontTargetPropertyType, _ = front.(*TargetPropertyType)
		}
		stage.OnAfterTargetPropertyTypeUpdateCallback.OnAfterUpdate(stage, targetpropertytype, frontTargetPropertyType)
	}
}

func (targetpropertytype *TargetPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTargetPropertyTypeDeleteCallback != nil {
		var frontTargetPropertyType *TargetPropertyType
		if front != nil {
			frontTargetPropertyType, _ = front.(*TargetPropertyType)
		}
		stage.OnAfterTargetPropertyTypeDeleteCallback.OnAfterDelete(stage, targetpropertytype, frontTargetPropertyType)
	}
}

func (temporalcrspropertytype *TemporalCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTemporalCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterTemporalCRSPropertyTypeCreateCallback.OnAfterCreate(stage, temporalcrspropertytype)
	}
}

func (temporalcrspropertytype *TemporalCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCRSPropertyTypeUpdateCallback != nil {
		var frontTemporalCRSPropertyType *TemporalCRSPropertyType
		if front != nil {
			frontTemporalCRSPropertyType, _ = front.(*TemporalCRSPropertyType)
		}
		stage.OnAfterTemporalCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, temporalcrspropertytype, frontTemporalCRSPropertyType)
	}
}

func (temporalcrspropertytype *TemporalCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCRSPropertyTypeDeleteCallback != nil {
		var frontTemporalCRSPropertyType *TemporalCRSPropertyType
		if front != nil {
			frontTemporalCRSPropertyType, _ = front.(*TemporalCRSPropertyType)
		}
		stage.OnAfterTemporalCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, temporalcrspropertytype, frontTemporalCRSPropertyType)
	}
}

func (temporalcrstype *TemporalCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTemporalCRSTypeCreateCallback != nil {
		stage.OnAfterTemporalCRSTypeCreateCallback.OnAfterCreate(stage, temporalcrstype)
	}
}

func (temporalcrstype *TemporalCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCRSTypeUpdateCallback != nil {
		var frontTemporalCRSType *TemporalCRSType
		if front != nil {
			frontTemporalCRSType, _ = front.(*TemporalCRSType)
		}
		stage.OnAfterTemporalCRSTypeUpdateCallback.OnAfterUpdate(stage, temporalcrstype, frontTemporalCRSType)
	}
}

func (temporalcrstype *TemporalCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCRSTypeDeleteCallback != nil {
		var frontTemporalCRSType *TemporalCRSType
		if front != nil {
			frontTemporalCRSType, _ = front.(*TemporalCRSType)
		}
		stage.OnAfterTemporalCRSTypeDeleteCallback.OnAfterDelete(stage, temporalcrstype, frontTemporalCRSType)
	}
}

func (temporalcspropertytype *TemporalCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTemporalCSPropertyTypeCreateCallback != nil {
		stage.OnAfterTemporalCSPropertyTypeCreateCallback.OnAfterCreate(stage, temporalcspropertytype)
	}
}

func (temporalcspropertytype *TemporalCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCSPropertyTypeUpdateCallback != nil {
		var frontTemporalCSPropertyType *TemporalCSPropertyType
		if front != nil {
			frontTemporalCSPropertyType, _ = front.(*TemporalCSPropertyType)
		}
		stage.OnAfterTemporalCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, temporalcspropertytype, frontTemporalCSPropertyType)
	}
}

func (temporalcspropertytype *TemporalCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCSPropertyTypeDeleteCallback != nil {
		var frontTemporalCSPropertyType *TemporalCSPropertyType
		if front != nil {
			frontTemporalCSPropertyType, _ = front.(*TemporalCSPropertyType)
		}
		stage.OnAfterTemporalCSPropertyTypeDeleteCallback.OnAfterDelete(stage, temporalcspropertytype, frontTemporalCSPropertyType)
	}
}

func (temporalcstype *TemporalCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTemporalCSTypeCreateCallback != nil {
		stage.OnAfterTemporalCSTypeCreateCallback.OnAfterCreate(stage, temporalcstype)
	}
}

func (temporalcstype *TemporalCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCSTypeUpdateCallback != nil {
		var frontTemporalCSType *TemporalCSType
		if front != nil {
			frontTemporalCSType, _ = front.(*TemporalCSType)
		}
		stage.OnAfterTemporalCSTypeUpdateCallback.OnAfterUpdate(stage, temporalcstype, frontTemporalCSType)
	}
}

func (temporalcstype *TemporalCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalCSTypeDeleteCallback != nil {
		var frontTemporalCSType *TemporalCSType
		if front != nil {
			frontTemporalCSType, _ = front.(*TemporalCSType)
		}
		stage.OnAfterTemporalCSTypeDeleteCallback.OnAfterDelete(stage, temporalcstype, frontTemporalCSType)
	}
}

func (temporaldatumbasetype *TemporalDatumBaseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTemporalDatumBaseTypeCreateCallback != nil {
		stage.OnAfterTemporalDatumBaseTypeCreateCallback.OnAfterCreate(stage, temporaldatumbasetype)
	}
}

func (temporaldatumbasetype *TemporalDatumBaseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalDatumBaseTypeUpdateCallback != nil {
		var frontTemporalDatumBaseType *TemporalDatumBaseType
		if front != nil {
			frontTemporalDatumBaseType, _ = front.(*TemporalDatumBaseType)
		}
		stage.OnAfterTemporalDatumBaseTypeUpdateCallback.OnAfterUpdate(stage, temporaldatumbasetype, frontTemporalDatumBaseType)
	}
}

func (temporaldatumbasetype *TemporalDatumBaseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalDatumBaseTypeDeleteCallback != nil {
		var frontTemporalDatumBaseType *TemporalDatumBaseType
		if front != nil {
			frontTemporalDatumBaseType, _ = front.(*TemporalDatumBaseType)
		}
		stage.OnAfterTemporalDatumBaseTypeDeleteCallback.OnAfterDelete(stage, temporaldatumbasetype, frontTemporalDatumBaseType)
	}
}

func (temporaldatumpropertytype *TemporalDatumPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTemporalDatumPropertyTypeCreateCallback != nil {
		stage.OnAfterTemporalDatumPropertyTypeCreateCallback.OnAfterCreate(stage, temporaldatumpropertytype)
	}
}

func (temporaldatumpropertytype *TemporalDatumPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalDatumPropertyTypeUpdateCallback != nil {
		var frontTemporalDatumPropertyType *TemporalDatumPropertyType
		if front != nil {
			frontTemporalDatumPropertyType, _ = front.(*TemporalDatumPropertyType)
		}
		stage.OnAfterTemporalDatumPropertyTypeUpdateCallback.OnAfterUpdate(stage, temporaldatumpropertytype, frontTemporalDatumPropertyType)
	}
}

func (temporaldatumpropertytype *TemporalDatumPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalDatumPropertyTypeDeleteCallback != nil {
		var frontTemporalDatumPropertyType *TemporalDatumPropertyType
		if front != nil {
			frontTemporalDatumPropertyType, _ = front.(*TemporalDatumPropertyType)
		}
		stage.OnAfterTemporalDatumPropertyTypeDeleteCallback.OnAfterDelete(stage, temporaldatumpropertytype, frontTemporalDatumPropertyType)
	}
}

func (temporaldatumtype *TemporalDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTemporalDatumTypeCreateCallback != nil {
		stage.OnAfterTemporalDatumTypeCreateCallback.OnAfterCreate(stage, temporaldatumtype)
	}
}

func (temporaldatumtype *TemporalDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalDatumTypeUpdateCallback != nil {
		var frontTemporalDatumType *TemporalDatumType
		if front != nil {
			frontTemporalDatumType, _ = front.(*TemporalDatumType)
		}
		stage.OnAfterTemporalDatumTypeUpdateCallback.OnAfterUpdate(stage, temporaldatumtype, frontTemporalDatumType)
	}
}

func (temporaldatumtype *TemporalDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTemporalDatumTypeDeleteCallback != nil {
		var frontTemporalDatumType *TemporalDatumType
		if front != nil {
			frontTemporalDatumType, _ = front.(*TemporalDatumType)
		}
		stage.OnAfterTemporalDatumTypeDeleteCallback.OnAfterDelete(stage, temporaldatumtype, frontTemporalDatumType)
	}
}

func (timecspropertytype *TimeCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCSPropertyTypeCreateCallback != nil {
		stage.OnAfterTimeCSPropertyTypeCreateCallback.OnAfterCreate(stage, timecspropertytype)
	}
}

func (timecspropertytype *TimeCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCSPropertyTypeUpdateCallback != nil {
		var frontTimeCSPropertyType *TimeCSPropertyType
		if front != nil {
			frontTimeCSPropertyType, _ = front.(*TimeCSPropertyType)
		}
		stage.OnAfterTimeCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, timecspropertytype, frontTimeCSPropertyType)
	}
}

func (timecspropertytype *TimeCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCSPropertyTypeDeleteCallback != nil {
		var frontTimeCSPropertyType *TimeCSPropertyType
		if front != nil {
			frontTimeCSPropertyType, _ = front.(*TimeCSPropertyType)
		}
		stage.OnAfterTimeCSPropertyTypeDeleteCallback.OnAfterDelete(stage, timecspropertytype, frontTimeCSPropertyType)
	}
}

func (timecstype *TimeCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCSTypeCreateCallback != nil {
		stage.OnAfterTimeCSTypeCreateCallback.OnAfterCreate(stage, timecstype)
	}
}

func (timecstype *TimeCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCSTypeUpdateCallback != nil {
		var frontTimeCSType *TimeCSType
		if front != nil {
			frontTimeCSType, _ = front.(*TimeCSType)
		}
		stage.OnAfterTimeCSTypeUpdateCallback.OnAfterUpdate(stage, timecstype, frontTimeCSType)
	}
}

func (timecstype *TimeCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCSTypeDeleteCallback != nil {
		var frontTimeCSType *TimeCSType
		if front != nil {
			frontTimeCSType, _ = front.(*TimeCSType)
		}
		stage.OnAfterTimeCSTypeDeleteCallback.OnAfterDelete(stage, timecstype, frontTimeCSType)
	}
}

func (timecalendarerapropertytype *TimeCalendarEraPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCalendarEraPropertyTypeCreateCallback != nil {
		stage.OnAfterTimeCalendarEraPropertyTypeCreateCallback.OnAfterCreate(stage, timecalendarerapropertytype)
	}
}

func (timecalendarerapropertytype *TimeCalendarEraPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarEraPropertyTypeUpdateCallback != nil {
		var frontTimeCalendarEraPropertyType *TimeCalendarEraPropertyType
		if front != nil {
			frontTimeCalendarEraPropertyType, _ = front.(*TimeCalendarEraPropertyType)
		}
		stage.OnAfterTimeCalendarEraPropertyTypeUpdateCallback.OnAfterUpdate(stage, timecalendarerapropertytype, frontTimeCalendarEraPropertyType)
	}
}

func (timecalendarerapropertytype *TimeCalendarEraPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarEraPropertyTypeDeleteCallback != nil {
		var frontTimeCalendarEraPropertyType *TimeCalendarEraPropertyType
		if front != nil {
			frontTimeCalendarEraPropertyType, _ = front.(*TimeCalendarEraPropertyType)
		}
		stage.OnAfterTimeCalendarEraPropertyTypeDeleteCallback.OnAfterDelete(stage, timecalendarerapropertytype, frontTimeCalendarEraPropertyType)
	}
}

func (timecalendareratype *TimeCalendarEraType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCalendarEraTypeCreateCallback != nil {
		stage.OnAfterTimeCalendarEraTypeCreateCallback.OnAfterCreate(stage, timecalendareratype)
	}
}

func (timecalendareratype *TimeCalendarEraType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarEraTypeUpdateCallback != nil {
		var frontTimeCalendarEraType *TimeCalendarEraType
		if front != nil {
			frontTimeCalendarEraType, _ = front.(*TimeCalendarEraType)
		}
		stage.OnAfterTimeCalendarEraTypeUpdateCallback.OnAfterUpdate(stage, timecalendareratype, frontTimeCalendarEraType)
	}
}

func (timecalendareratype *TimeCalendarEraType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarEraTypeDeleteCallback != nil {
		var frontTimeCalendarEraType *TimeCalendarEraType
		if front != nil {
			frontTimeCalendarEraType, _ = front.(*TimeCalendarEraType)
		}
		stage.OnAfterTimeCalendarEraTypeDeleteCallback.OnAfterDelete(stage, timecalendareratype, frontTimeCalendarEraType)
	}
}

func (timecalendarpropertytype *TimeCalendarPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCalendarPropertyTypeCreateCallback != nil {
		stage.OnAfterTimeCalendarPropertyTypeCreateCallback.OnAfterCreate(stage, timecalendarpropertytype)
	}
}

func (timecalendarpropertytype *TimeCalendarPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarPropertyTypeUpdateCallback != nil {
		var frontTimeCalendarPropertyType *TimeCalendarPropertyType
		if front != nil {
			frontTimeCalendarPropertyType, _ = front.(*TimeCalendarPropertyType)
		}
		stage.OnAfterTimeCalendarPropertyTypeUpdateCallback.OnAfterUpdate(stage, timecalendarpropertytype, frontTimeCalendarPropertyType)
	}
}

func (timecalendarpropertytype *TimeCalendarPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarPropertyTypeDeleteCallback != nil {
		var frontTimeCalendarPropertyType *TimeCalendarPropertyType
		if front != nil {
			frontTimeCalendarPropertyType, _ = front.(*TimeCalendarPropertyType)
		}
		stage.OnAfterTimeCalendarPropertyTypeDeleteCallback.OnAfterDelete(stage, timecalendarpropertytype, frontTimeCalendarPropertyType)
	}
}

func (timecalendartype *TimeCalendarType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCalendarTypeCreateCallback != nil {
		stage.OnAfterTimeCalendarTypeCreateCallback.OnAfterCreate(stage, timecalendartype)
	}
}

func (timecalendartype *TimeCalendarType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarTypeUpdateCallback != nil {
		var frontTimeCalendarType *TimeCalendarType
		if front != nil {
			frontTimeCalendarType, _ = front.(*TimeCalendarType)
		}
		stage.OnAfterTimeCalendarTypeUpdateCallback.OnAfterUpdate(stage, timecalendartype, frontTimeCalendarType)
	}
}

func (timecalendartype *TimeCalendarType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCalendarTypeDeleteCallback != nil {
		var frontTimeCalendarType *TimeCalendarType
		if front != nil {
			frontTimeCalendarType, _ = front.(*TimeCalendarType)
		}
		stage.OnAfterTimeCalendarTypeDeleteCallback.OnAfterDelete(stage, timecalendartype, frontTimeCalendarType)
	}
}

func (timeclockpropertytype *TimeClockPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeClockPropertyTypeCreateCallback != nil {
		stage.OnAfterTimeClockPropertyTypeCreateCallback.OnAfterCreate(stage, timeclockpropertytype)
	}
}

func (timeclockpropertytype *TimeClockPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeClockPropertyTypeUpdateCallback != nil {
		var frontTimeClockPropertyType *TimeClockPropertyType
		if front != nil {
			frontTimeClockPropertyType, _ = front.(*TimeClockPropertyType)
		}
		stage.OnAfterTimeClockPropertyTypeUpdateCallback.OnAfterUpdate(stage, timeclockpropertytype, frontTimeClockPropertyType)
	}
}

func (timeclockpropertytype *TimeClockPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeClockPropertyTypeDeleteCallback != nil {
		var frontTimeClockPropertyType *TimeClockPropertyType
		if front != nil {
			frontTimeClockPropertyType, _ = front.(*TimeClockPropertyType)
		}
		stage.OnAfterTimeClockPropertyTypeDeleteCallback.OnAfterDelete(stage, timeclockpropertytype, frontTimeClockPropertyType)
	}
}

func (timeclocktype *TimeClockType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeClockTypeCreateCallback != nil {
		stage.OnAfterTimeClockTypeCreateCallback.OnAfterCreate(stage, timeclocktype)
	}
}

func (timeclocktype *TimeClockType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeClockTypeUpdateCallback != nil {
		var frontTimeClockType *TimeClockType
		if front != nil {
			frontTimeClockType, _ = front.(*TimeClockType)
		}
		stage.OnAfterTimeClockTypeUpdateCallback.OnAfterUpdate(stage, timeclocktype, frontTimeClockType)
	}
}

func (timeclocktype *TimeClockType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeClockTypeDeleteCallback != nil {
		var frontTimeClockType *TimeClockType
		if front != nil {
			frontTimeClockType, _ = front.(*TimeClockType)
		}
		stage.OnAfterTimeClockTypeDeleteCallback.OnAfterDelete(stage, timeclocktype, frontTimeClockType)
	}
}

func (timecoordinatesystemtype *TimeCoordinateSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeCoordinateSystemTypeCreateCallback != nil {
		stage.OnAfterTimeCoordinateSystemTypeCreateCallback.OnAfterCreate(stage, timecoordinatesystemtype)
	}
}

func (timecoordinatesystemtype *TimeCoordinateSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCoordinateSystemTypeUpdateCallback != nil {
		var frontTimeCoordinateSystemType *TimeCoordinateSystemType
		if front != nil {
			frontTimeCoordinateSystemType, _ = front.(*TimeCoordinateSystemType)
		}
		stage.OnAfterTimeCoordinateSystemTypeUpdateCallback.OnAfterUpdate(stage, timecoordinatesystemtype, frontTimeCoordinateSystemType)
	}
}

func (timecoordinatesystemtype *TimeCoordinateSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeCoordinateSystemTypeDeleteCallback != nil {
		var frontTimeCoordinateSystemType *TimeCoordinateSystemType
		if front != nil {
			frontTimeCoordinateSystemType, _ = front.(*TimeCoordinateSystemType)
		}
		stage.OnAfterTimeCoordinateSystemTypeDeleteCallback.OnAfterDelete(stage, timecoordinatesystemtype, frontTimeCoordinateSystemType)
	}
}

func (timeedgepropertytype *TimeEdgePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeEdgePropertyTypeCreateCallback != nil {
		stage.OnAfterTimeEdgePropertyTypeCreateCallback.OnAfterCreate(stage, timeedgepropertytype)
	}
}

func (timeedgepropertytype *TimeEdgePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeEdgePropertyTypeUpdateCallback != nil {
		var frontTimeEdgePropertyType *TimeEdgePropertyType
		if front != nil {
			frontTimeEdgePropertyType, _ = front.(*TimeEdgePropertyType)
		}
		stage.OnAfterTimeEdgePropertyTypeUpdateCallback.OnAfterUpdate(stage, timeedgepropertytype, frontTimeEdgePropertyType)
	}
}

func (timeedgepropertytype *TimeEdgePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeEdgePropertyTypeDeleteCallback != nil {
		var frontTimeEdgePropertyType *TimeEdgePropertyType
		if front != nil {
			frontTimeEdgePropertyType, _ = front.(*TimeEdgePropertyType)
		}
		stage.OnAfterTimeEdgePropertyTypeDeleteCallback.OnAfterDelete(stage, timeedgepropertytype, frontTimeEdgePropertyType)
	}
}

func (timeedgetype *TimeEdgeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeEdgeTypeCreateCallback != nil {
		stage.OnAfterTimeEdgeTypeCreateCallback.OnAfterCreate(stage, timeedgetype)
	}
}

func (timeedgetype *TimeEdgeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeEdgeTypeUpdateCallback != nil {
		var frontTimeEdgeType *TimeEdgeType
		if front != nil {
			frontTimeEdgeType, _ = front.(*TimeEdgeType)
		}
		stage.OnAfterTimeEdgeTypeUpdateCallback.OnAfterUpdate(stage, timeedgetype, frontTimeEdgeType)
	}
}

func (timeedgetype *TimeEdgeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeEdgeTypeDeleteCallback != nil {
		var frontTimeEdgeType *TimeEdgeType
		if front != nil {
			frontTimeEdgeType, _ = front.(*TimeEdgeType)
		}
		stage.OnAfterTimeEdgeTypeDeleteCallback.OnAfterDelete(stage, timeedgetype, frontTimeEdgeType)
	}
}

func (timeinstantpropertytype *TimeInstantPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeInstantPropertyTypeCreateCallback != nil {
		stage.OnAfterTimeInstantPropertyTypeCreateCallback.OnAfterCreate(stage, timeinstantpropertytype)
	}
}

func (timeinstantpropertytype *TimeInstantPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeInstantPropertyTypeUpdateCallback != nil {
		var frontTimeInstantPropertyType *TimeInstantPropertyType
		if front != nil {
			frontTimeInstantPropertyType, _ = front.(*TimeInstantPropertyType)
		}
		stage.OnAfterTimeInstantPropertyTypeUpdateCallback.OnAfterUpdate(stage, timeinstantpropertytype, frontTimeInstantPropertyType)
	}
}

func (timeinstantpropertytype *TimeInstantPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeInstantPropertyTypeDeleteCallback != nil {
		var frontTimeInstantPropertyType *TimeInstantPropertyType
		if front != nil {
			frontTimeInstantPropertyType, _ = front.(*TimeInstantPropertyType)
		}
		stage.OnAfterTimeInstantPropertyTypeDeleteCallback.OnAfterDelete(stage, timeinstantpropertytype, frontTimeInstantPropertyType)
	}
}

func (timeinstanttype *TimeInstantType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeInstantTypeCreateCallback != nil {
		stage.OnAfterTimeInstantTypeCreateCallback.OnAfterCreate(stage, timeinstanttype)
	}
}

func (timeinstanttype *TimeInstantType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeInstantTypeUpdateCallback != nil {
		var frontTimeInstantType *TimeInstantType
		if front != nil {
			frontTimeInstantType, _ = front.(*TimeInstantType)
		}
		stage.OnAfterTimeInstantTypeUpdateCallback.OnAfterUpdate(stage, timeinstanttype, frontTimeInstantType)
	}
}

func (timeinstanttype *TimeInstantType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeInstantTypeDeleteCallback != nil {
		var frontTimeInstantType *TimeInstantType
		if front != nil {
			frontTimeInstantType, _ = front.(*TimeInstantType)
		}
		stage.OnAfterTimeInstantTypeDeleteCallback.OnAfterDelete(stage, timeinstanttype, frontTimeInstantType)
	}
}

func (timeintervallengthtype *TimeIntervalLengthType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeIntervalLengthTypeCreateCallback != nil {
		stage.OnAfterTimeIntervalLengthTypeCreateCallback.OnAfterCreate(stage, timeintervallengthtype)
	}
}

func (timeintervallengthtype *TimeIntervalLengthType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeIntervalLengthTypeUpdateCallback != nil {
		var frontTimeIntervalLengthType *TimeIntervalLengthType
		if front != nil {
			frontTimeIntervalLengthType, _ = front.(*TimeIntervalLengthType)
		}
		stage.OnAfterTimeIntervalLengthTypeUpdateCallback.OnAfterUpdate(stage, timeintervallengthtype, frontTimeIntervalLengthType)
	}
}

func (timeintervallengthtype *TimeIntervalLengthType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeIntervalLengthTypeDeleteCallback != nil {
		var frontTimeIntervalLengthType *TimeIntervalLengthType
		if front != nil {
			frontTimeIntervalLengthType, _ = front.(*TimeIntervalLengthType)
		}
		stage.OnAfterTimeIntervalLengthTypeDeleteCallback.OnAfterDelete(stage, timeintervallengthtype, frontTimeIntervalLengthType)
	}
}

func (timenodepropertytype *TimeNodePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeNodePropertyTypeCreateCallback != nil {
		stage.OnAfterTimeNodePropertyTypeCreateCallback.OnAfterCreate(stage, timenodepropertytype)
	}
}

func (timenodepropertytype *TimeNodePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeNodePropertyTypeUpdateCallback != nil {
		var frontTimeNodePropertyType *TimeNodePropertyType
		if front != nil {
			frontTimeNodePropertyType, _ = front.(*TimeNodePropertyType)
		}
		stage.OnAfterTimeNodePropertyTypeUpdateCallback.OnAfterUpdate(stage, timenodepropertytype, frontTimeNodePropertyType)
	}
}

func (timenodepropertytype *TimeNodePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeNodePropertyTypeDeleteCallback != nil {
		var frontTimeNodePropertyType *TimeNodePropertyType
		if front != nil {
			frontTimeNodePropertyType, _ = front.(*TimeNodePropertyType)
		}
		stage.OnAfterTimeNodePropertyTypeDeleteCallback.OnAfterDelete(stage, timenodepropertytype, frontTimeNodePropertyType)
	}
}

func (timenodetype *TimeNodeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeNodeTypeCreateCallback != nil {
		stage.OnAfterTimeNodeTypeCreateCallback.OnAfterCreate(stage, timenodetype)
	}
}

func (timenodetype *TimeNodeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeNodeTypeUpdateCallback != nil {
		var frontTimeNodeType *TimeNodeType
		if front != nil {
			frontTimeNodeType, _ = front.(*TimeNodeType)
		}
		stage.OnAfterTimeNodeTypeUpdateCallback.OnAfterUpdate(stage, timenodetype, frontTimeNodeType)
	}
}

func (timenodetype *TimeNodeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeNodeTypeDeleteCallback != nil {
		var frontTimeNodeType *TimeNodeType
		if front != nil {
			frontTimeNodeType, _ = front.(*TimeNodeType)
		}
		stage.OnAfterTimeNodeTypeDeleteCallback.OnAfterDelete(stage, timenodetype, frontTimeNodeType)
	}
}

func (timeordinalerapropertytype *TimeOrdinalEraPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeOrdinalEraPropertyTypeCreateCallback != nil {
		stage.OnAfterTimeOrdinalEraPropertyTypeCreateCallback.OnAfterCreate(stage, timeordinalerapropertytype)
	}
}

func (timeordinalerapropertytype *TimeOrdinalEraPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeOrdinalEraPropertyTypeUpdateCallback != nil {
		var frontTimeOrdinalEraPropertyType *TimeOrdinalEraPropertyType
		if front != nil {
			frontTimeOrdinalEraPropertyType, _ = front.(*TimeOrdinalEraPropertyType)
		}
		stage.OnAfterTimeOrdinalEraPropertyTypeUpdateCallback.OnAfterUpdate(stage, timeordinalerapropertytype, frontTimeOrdinalEraPropertyType)
	}
}

func (timeordinalerapropertytype *TimeOrdinalEraPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeOrdinalEraPropertyTypeDeleteCallback != nil {
		var frontTimeOrdinalEraPropertyType *TimeOrdinalEraPropertyType
		if front != nil {
			frontTimeOrdinalEraPropertyType, _ = front.(*TimeOrdinalEraPropertyType)
		}
		stage.OnAfterTimeOrdinalEraPropertyTypeDeleteCallback.OnAfterDelete(stage, timeordinalerapropertytype, frontTimeOrdinalEraPropertyType)
	}
}

func (timeordinaleratype *TimeOrdinalEraType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeOrdinalEraTypeCreateCallback != nil {
		stage.OnAfterTimeOrdinalEraTypeCreateCallback.OnAfterCreate(stage, timeordinaleratype)
	}
}

func (timeordinaleratype *TimeOrdinalEraType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeOrdinalEraTypeUpdateCallback != nil {
		var frontTimeOrdinalEraType *TimeOrdinalEraType
		if front != nil {
			frontTimeOrdinalEraType, _ = front.(*TimeOrdinalEraType)
		}
		stage.OnAfterTimeOrdinalEraTypeUpdateCallback.OnAfterUpdate(stage, timeordinaleratype, frontTimeOrdinalEraType)
	}
}

func (timeordinaleratype *TimeOrdinalEraType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeOrdinalEraTypeDeleteCallback != nil {
		var frontTimeOrdinalEraType *TimeOrdinalEraType
		if front != nil {
			frontTimeOrdinalEraType, _ = front.(*TimeOrdinalEraType)
		}
		stage.OnAfterTimeOrdinalEraTypeDeleteCallback.OnAfterDelete(stage, timeordinaleratype, frontTimeOrdinalEraType)
	}
}

func (timeordinalreferencesystemtype *TimeOrdinalReferenceSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeOrdinalReferenceSystemTypeCreateCallback != nil {
		stage.OnAfterTimeOrdinalReferenceSystemTypeCreateCallback.OnAfterCreate(stage, timeordinalreferencesystemtype)
	}
}

func (timeordinalreferencesystemtype *TimeOrdinalReferenceSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeOrdinalReferenceSystemTypeUpdateCallback != nil {
		var frontTimeOrdinalReferenceSystemType *TimeOrdinalReferenceSystemType
		if front != nil {
			frontTimeOrdinalReferenceSystemType, _ = front.(*TimeOrdinalReferenceSystemType)
		}
		stage.OnAfterTimeOrdinalReferenceSystemTypeUpdateCallback.OnAfterUpdate(stage, timeordinalreferencesystemtype, frontTimeOrdinalReferenceSystemType)
	}
}

func (timeordinalreferencesystemtype *TimeOrdinalReferenceSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeOrdinalReferenceSystemTypeDeleteCallback != nil {
		var frontTimeOrdinalReferenceSystemType *TimeOrdinalReferenceSystemType
		if front != nil {
			frontTimeOrdinalReferenceSystemType, _ = front.(*TimeOrdinalReferenceSystemType)
		}
		stage.OnAfterTimeOrdinalReferenceSystemTypeDeleteCallback.OnAfterDelete(stage, timeordinalreferencesystemtype, frontTimeOrdinalReferenceSystemType)
	}
}

func (timeperiodpropertytype *TimePeriodPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimePeriodPropertyTypeCreateCallback != nil {
		stage.OnAfterTimePeriodPropertyTypeCreateCallback.OnAfterCreate(stage, timeperiodpropertytype)
	}
}

func (timeperiodpropertytype *TimePeriodPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePeriodPropertyTypeUpdateCallback != nil {
		var frontTimePeriodPropertyType *TimePeriodPropertyType
		if front != nil {
			frontTimePeriodPropertyType, _ = front.(*TimePeriodPropertyType)
		}
		stage.OnAfterTimePeriodPropertyTypeUpdateCallback.OnAfterUpdate(stage, timeperiodpropertytype, frontTimePeriodPropertyType)
	}
}

func (timeperiodpropertytype *TimePeriodPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePeriodPropertyTypeDeleteCallback != nil {
		var frontTimePeriodPropertyType *TimePeriodPropertyType
		if front != nil {
			frontTimePeriodPropertyType, _ = front.(*TimePeriodPropertyType)
		}
		stage.OnAfterTimePeriodPropertyTypeDeleteCallback.OnAfterDelete(stage, timeperiodpropertytype, frontTimePeriodPropertyType)
	}
}

func (timeperiodtype *TimePeriodType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimePeriodTypeCreateCallback != nil {
		stage.OnAfterTimePeriodTypeCreateCallback.OnAfterCreate(stage, timeperiodtype)
	}
}

func (timeperiodtype *TimePeriodType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePeriodTypeUpdateCallback != nil {
		var frontTimePeriodType *TimePeriodType
		if front != nil {
			frontTimePeriodType, _ = front.(*TimePeriodType)
		}
		stage.OnAfterTimePeriodTypeUpdateCallback.OnAfterUpdate(stage, timeperiodtype, frontTimePeriodType)
	}
}

func (timeperiodtype *TimePeriodType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePeriodTypeDeleteCallback != nil {
		var frontTimePeriodType *TimePeriodType
		if front != nil {
			frontTimePeriodType, _ = front.(*TimePeriodType)
		}
		stage.OnAfterTimePeriodTypeDeleteCallback.OnAfterDelete(stage, timeperiodtype, frontTimePeriodType)
	}
}

func (timepositiontype *TimePositionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimePositionTypeCreateCallback != nil {
		stage.OnAfterTimePositionTypeCreateCallback.OnAfterCreate(stage, timepositiontype)
	}
}

func (timepositiontype *TimePositionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePositionTypeUpdateCallback != nil {
		var frontTimePositionType *TimePositionType
		if front != nil {
			frontTimePositionType, _ = front.(*TimePositionType)
		}
		stage.OnAfterTimePositionTypeUpdateCallback.OnAfterUpdate(stage, timepositiontype, frontTimePositionType)
	}
}

func (timepositiontype *TimePositionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePositionTypeDeleteCallback != nil {
		var frontTimePositionType *TimePositionType
		if front != nil {
			frontTimePositionType, _ = front.(*TimePositionType)
		}
		stage.OnAfterTimePositionTypeDeleteCallback.OnAfterDelete(stage, timepositiontype, frontTimePositionType)
	}
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimePrimitivePropertyTypeCreateCallback != nil {
		stage.OnAfterTimePrimitivePropertyTypeCreateCallback.OnAfterCreate(stage, timeprimitivepropertytype)
	}
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePrimitivePropertyTypeUpdateCallback != nil {
		var frontTimePrimitivePropertyType *TimePrimitivePropertyType
		if front != nil {
			frontTimePrimitivePropertyType, _ = front.(*TimePrimitivePropertyType)
		}
		stage.OnAfterTimePrimitivePropertyTypeUpdateCallback.OnAfterUpdate(stage, timeprimitivepropertytype, frontTimePrimitivePropertyType)
	}
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimePrimitivePropertyTypeDeleteCallback != nil {
		var frontTimePrimitivePropertyType *TimePrimitivePropertyType
		if front != nil {
			frontTimePrimitivePropertyType, _ = front.(*TimePrimitivePropertyType)
		}
		stage.OnAfterTimePrimitivePropertyTypeDeleteCallback.OnAfterDelete(stage, timeprimitivepropertytype, frontTimePrimitivePropertyType)
	}
}

func (timereferencesystemtype *TimeReferenceSystemType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeReferenceSystemTypeCreateCallback != nil {
		stage.OnAfterTimeReferenceSystemTypeCreateCallback.OnAfterCreate(stage, timereferencesystemtype)
	}
}

func (timereferencesystemtype *TimeReferenceSystemType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeReferenceSystemTypeUpdateCallback != nil {
		var frontTimeReferenceSystemType *TimeReferenceSystemType
		if front != nil {
			frontTimeReferenceSystemType, _ = front.(*TimeReferenceSystemType)
		}
		stage.OnAfterTimeReferenceSystemTypeUpdateCallback.OnAfterUpdate(stage, timereferencesystemtype, frontTimeReferenceSystemType)
	}
}

func (timereferencesystemtype *TimeReferenceSystemType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeReferenceSystemTypeDeleteCallback != nil {
		var frontTimeReferenceSystemType *TimeReferenceSystemType
		if front != nil {
			frontTimeReferenceSystemType, _ = front.(*TimeReferenceSystemType)
		}
		stage.OnAfterTimeReferenceSystemTypeDeleteCallback.OnAfterDelete(stage, timereferencesystemtype, frontTimeReferenceSystemType)
	}
}

func (timetopologycomplexpropertytype *TimeTopologyComplexPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeTopologyComplexPropertyTypeCreateCallback != nil {
		stage.OnAfterTimeTopologyComplexPropertyTypeCreateCallback.OnAfterCreate(stage, timetopologycomplexpropertytype)
	}
}

func (timetopologycomplexpropertytype *TimeTopologyComplexPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTopologyComplexPropertyTypeUpdateCallback != nil {
		var frontTimeTopologyComplexPropertyType *TimeTopologyComplexPropertyType
		if front != nil {
			frontTimeTopologyComplexPropertyType, _ = front.(*TimeTopologyComplexPropertyType)
		}
		stage.OnAfterTimeTopologyComplexPropertyTypeUpdateCallback.OnAfterUpdate(stage, timetopologycomplexpropertytype, frontTimeTopologyComplexPropertyType)
	}
}

func (timetopologycomplexpropertytype *TimeTopologyComplexPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTopologyComplexPropertyTypeDeleteCallback != nil {
		var frontTimeTopologyComplexPropertyType *TimeTopologyComplexPropertyType
		if front != nil {
			frontTimeTopologyComplexPropertyType, _ = front.(*TimeTopologyComplexPropertyType)
		}
		stage.OnAfterTimeTopologyComplexPropertyTypeDeleteCallback.OnAfterDelete(stage, timetopologycomplexpropertytype, frontTimeTopologyComplexPropertyType)
	}
}

func (timetopologycomplextype *TimeTopologyComplexType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeTopologyComplexTypeCreateCallback != nil {
		stage.OnAfterTimeTopologyComplexTypeCreateCallback.OnAfterCreate(stage, timetopologycomplextype)
	}
}

func (timetopologycomplextype *TimeTopologyComplexType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTopologyComplexTypeUpdateCallback != nil {
		var frontTimeTopologyComplexType *TimeTopologyComplexType
		if front != nil {
			frontTimeTopologyComplexType, _ = front.(*TimeTopologyComplexType)
		}
		stage.OnAfterTimeTopologyComplexTypeUpdateCallback.OnAfterUpdate(stage, timetopologycomplextype, frontTimeTopologyComplexType)
	}
}

func (timetopologycomplextype *TimeTopologyComplexType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTopologyComplexTypeDeleteCallback != nil {
		var frontTimeTopologyComplexType *TimeTopologyComplexType
		if front != nil {
			frontTimeTopologyComplexType, _ = front.(*TimeTopologyComplexType)
		}
		stage.OnAfterTimeTopologyComplexTypeDeleteCallback.OnAfterDelete(stage, timetopologycomplextype, frontTimeTopologyComplexType)
	}
}

func (timetopologyprimitivepropertytype *TimeTopologyPrimitivePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTimeTopologyPrimitivePropertyTypeCreateCallback != nil {
		stage.OnAfterTimeTopologyPrimitivePropertyTypeCreateCallback.OnAfterCreate(stage, timetopologyprimitivepropertytype)
	}
}

func (timetopologyprimitivepropertytype *TimeTopologyPrimitivePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTopologyPrimitivePropertyTypeUpdateCallback != nil {
		var frontTimeTopologyPrimitivePropertyType *TimeTopologyPrimitivePropertyType
		if front != nil {
			frontTimeTopologyPrimitivePropertyType, _ = front.(*TimeTopologyPrimitivePropertyType)
		}
		stage.OnAfterTimeTopologyPrimitivePropertyTypeUpdateCallback.OnAfterUpdate(stage, timetopologyprimitivepropertytype, frontTimeTopologyPrimitivePropertyType)
	}
}

func (timetopologyprimitivepropertytype *TimeTopologyPrimitivePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTimeTopologyPrimitivePropertyTypeDeleteCallback != nil {
		var frontTimeTopologyPrimitivePropertyType *TimeTopologyPrimitivePropertyType
		if front != nil {
			frontTimeTopologyPrimitivePropertyType, _ = front.(*TimeTopologyPrimitivePropertyType)
		}
		stage.OnAfterTimeTopologyPrimitivePropertyTypeDeleteCallback.OnAfterDelete(stage, timetopologyprimitivepropertytype, frontTimeTopologyPrimitivePropertyType)
	}
}

func (tintype *TinType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTinTypeCreateCallback != nil {
		stage.OnAfterTinTypeCreateCallback.OnAfterCreate(stage, tintype)
	}
}

func (tintype *TinType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTinTypeUpdateCallback != nil {
		var frontTinType *TinType
		if front != nil {
			frontTinType, _ = front.(*TinType)
		}
		stage.OnAfterTinTypeUpdateCallback.OnAfterUpdate(stage, tintype, frontTinType)
	}
}

func (tintype *TinType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTinTypeDeleteCallback != nil {
		var frontTinType *TinType
		if front != nil {
			frontTinType, _ = front.(*TinType)
		}
		stage.OnAfterTinTypeDeleteCallback.OnAfterDelete(stage, tintype, frontTinType)
	}
}

func (topocomplexpropertytype *TopoComplexPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoComplexPropertyTypeCreateCallback != nil {
		stage.OnAfterTopoComplexPropertyTypeCreateCallback.OnAfterCreate(stage, topocomplexpropertytype)
	}
}

func (topocomplexpropertytype *TopoComplexPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoComplexPropertyTypeUpdateCallback != nil {
		var frontTopoComplexPropertyType *TopoComplexPropertyType
		if front != nil {
			frontTopoComplexPropertyType, _ = front.(*TopoComplexPropertyType)
		}
		stage.OnAfterTopoComplexPropertyTypeUpdateCallback.OnAfterUpdate(stage, topocomplexpropertytype, frontTopoComplexPropertyType)
	}
}

func (topocomplexpropertytype *TopoComplexPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoComplexPropertyTypeDeleteCallback != nil {
		var frontTopoComplexPropertyType *TopoComplexPropertyType
		if front != nil {
			frontTopoComplexPropertyType, _ = front.(*TopoComplexPropertyType)
		}
		stage.OnAfterTopoComplexPropertyTypeDeleteCallback.OnAfterDelete(stage, topocomplexpropertytype, frontTopoComplexPropertyType)
	}
}

func (topocomplextype *TopoComplexType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoComplexTypeCreateCallback != nil {
		stage.OnAfterTopoComplexTypeCreateCallback.OnAfterCreate(stage, topocomplextype)
	}
}

func (topocomplextype *TopoComplexType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoComplexTypeUpdateCallback != nil {
		var frontTopoComplexType *TopoComplexType
		if front != nil {
			frontTopoComplexType, _ = front.(*TopoComplexType)
		}
		stage.OnAfterTopoComplexTypeUpdateCallback.OnAfterUpdate(stage, topocomplextype, frontTopoComplexType)
	}
}

func (topocomplextype *TopoComplexType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoComplexTypeDeleteCallback != nil {
		var frontTopoComplexType *TopoComplexType
		if front != nil {
			frontTopoComplexType, _ = front.(*TopoComplexType)
		}
		stage.OnAfterTopoComplexTypeDeleteCallback.OnAfterDelete(stage, topocomplextype, frontTopoComplexType)
	}
}

func (topocurvepropertytype *TopoCurvePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoCurvePropertyTypeCreateCallback != nil {
		stage.OnAfterTopoCurvePropertyTypeCreateCallback.OnAfterCreate(stage, topocurvepropertytype)
	}
}

func (topocurvepropertytype *TopoCurvePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoCurvePropertyTypeUpdateCallback != nil {
		var frontTopoCurvePropertyType *TopoCurvePropertyType
		if front != nil {
			frontTopoCurvePropertyType, _ = front.(*TopoCurvePropertyType)
		}
		stage.OnAfterTopoCurvePropertyTypeUpdateCallback.OnAfterUpdate(stage, topocurvepropertytype, frontTopoCurvePropertyType)
	}
}

func (topocurvepropertytype *TopoCurvePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoCurvePropertyTypeDeleteCallback != nil {
		var frontTopoCurvePropertyType *TopoCurvePropertyType
		if front != nil {
			frontTopoCurvePropertyType, _ = front.(*TopoCurvePropertyType)
		}
		stage.OnAfterTopoCurvePropertyTypeDeleteCallback.OnAfterDelete(stage, topocurvepropertytype, frontTopoCurvePropertyType)
	}
}

func (topocurvetype *TopoCurveType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoCurveTypeCreateCallback != nil {
		stage.OnAfterTopoCurveTypeCreateCallback.OnAfterCreate(stage, topocurvetype)
	}
}

func (topocurvetype *TopoCurveType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoCurveTypeUpdateCallback != nil {
		var frontTopoCurveType *TopoCurveType
		if front != nil {
			frontTopoCurveType, _ = front.(*TopoCurveType)
		}
		stage.OnAfterTopoCurveTypeUpdateCallback.OnAfterUpdate(stage, topocurvetype, frontTopoCurveType)
	}
}

func (topocurvetype *TopoCurveType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoCurveTypeDeleteCallback != nil {
		var frontTopoCurveType *TopoCurveType
		if front != nil {
			frontTopoCurveType, _ = front.(*TopoCurveType)
		}
		stage.OnAfterTopoCurveTypeDeleteCallback.OnAfterDelete(stage, topocurvetype, frontTopoCurveType)
	}
}

func (topopointpropertytype *TopoPointPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoPointPropertyTypeCreateCallback != nil {
		stage.OnAfterTopoPointPropertyTypeCreateCallback.OnAfterCreate(stage, topopointpropertytype)
	}
}

func (topopointpropertytype *TopoPointPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPointPropertyTypeUpdateCallback != nil {
		var frontTopoPointPropertyType *TopoPointPropertyType
		if front != nil {
			frontTopoPointPropertyType, _ = front.(*TopoPointPropertyType)
		}
		stage.OnAfterTopoPointPropertyTypeUpdateCallback.OnAfterUpdate(stage, topopointpropertytype, frontTopoPointPropertyType)
	}
}

func (topopointpropertytype *TopoPointPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPointPropertyTypeDeleteCallback != nil {
		var frontTopoPointPropertyType *TopoPointPropertyType
		if front != nil {
			frontTopoPointPropertyType, _ = front.(*TopoPointPropertyType)
		}
		stage.OnAfterTopoPointPropertyTypeDeleteCallback.OnAfterDelete(stage, topopointpropertytype, frontTopoPointPropertyType)
	}
}

func (topopointtype *TopoPointType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoPointTypeCreateCallback != nil {
		stage.OnAfterTopoPointTypeCreateCallback.OnAfterCreate(stage, topopointtype)
	}
}

func (topopointtype *TopoPointType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPointTypeUpdateCallback != nil {
		var frontTopoPointType *TopoPointType
		if front != nil {
			frontTopoPointType, _ = front.(*TopoPointType)
		}
		stage.OnAfterTopoPointTypeUpdateCallback.OnAfterUpdate(stage, topopointtype, frontTopoPointType)
	}
}

func (topopointtype *TopoPointType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPointTypeDeleteCallback != nil {
		var frontTopoPointType *TopoPointType
		if front != nil {
			frontTopoPointType, _ = front.(*TopoPointType)
		}
		stage.OnAfterTopoPointTypeDeleteCallback.OnAfterDelete(stage, topopointtype, frontTopoPointType)
	}
}

func (topoprimitivearrayassociationtype *TopoPrimitiveArrayAssociationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoPrimitiveArrayAssociationTypeCreateCallback != nil {
		stage.OnAfterTopoPrimitiveArrayAssociationTypeCreateCallback.OnAfterCreate(stage, topoprimitivearrayassociationtype)
	}
}

func (topoprimitivearrayassociationtype *TopoPrimitiveArrayAssociationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPrimitiveArrayAssociationTypeUpdateCallback != nil {
		var frontTopoPrimitiveArrayAssociationType *TopoPrimitiveArrayAssociationType
		if front != nil {
			frontTopoPrimitiveArrayAssociationType, _ = front.(*TopoPrimitiveArrayAssociationType)
		}
		stage.OnAfterTopoPrimitiveArrayAssociationTypeUpdateCallback.OnAfterUpdate(stage, topoprimitivearrayassociationtype, frontTopoPrimitiveArrayAssociationType)
	}
}

func (topoprimitivearrayassociationtype *TopoPrimitiveArrayAssociationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPrimitiveArrayAssociationTypeDeleteCallback != nil {
		var frontTopoPrimitiveArrayAssociationType *TopoPrimitiveArrayAssociationType
		if front != nil {
			frontTopoPrimitiveArrayAssociationType, _ = front.(*TopoPrimitiveArrayAssociationType)
		}
		stage.OnAfterTopoPrimitiveArrayAssociationTypeDeleteCallback.OnAfterDelete(stage, topoprimitivearrayassociationtype, frontTopoPrimitiveArrayAssociationType)
	}
}

func (topoprimitivemembertype *TopoPrimitiveMemberType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoPrimitiveMemberTypeCreateCallback != nil {
		stage.OnAfterTopoPrimitiveMemberTypeCreateCallback.OnAfterCreate(stage, topoprimitivemembertype)
	}
}

func (topoprimitivemembertype *TopoPrimitiveMemberType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPrimitiveMemberTypeUpdateCallback != nil {
		var frontTopoPrimitiveMemberType *TopoPrimitiveMemberType
		if front != nil {
			frontTopoPrimitiveMemberType, _ = front.(*TopoPrimitiveMemberType)
		}
		stage.OnAfterTopoPrimitiveMemberTypeUpdateCallback.OnAfterUpdate(stage, topoprimitivemembertype, frontTopoPrimitiveMemberType)
	}
}

func (topoprimitivemembertype *TopoPrimitiveMemberType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoPrimitiveMemberTypeDeleteCallback != nil {
		var frontTopoPrimitiveMemberType *TopoPrimitiveMemberType
		if front != nil {
			frontTopoPrimitiveMemberType, _ = front.(*TopoPrimitiveMemberType)
		}
		stage.OnAfterTopoPrimitiveMemberTypeDeleteCallback.OnAfterDelete(stage, topoprimitivemembertype, frontTopoPrimitiveMemberType)
	}
}

func (toposolidpropertytype *TopoSolidPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoSolidPropertyTypeCreateCallback != nil {
		stage.OnAfterTopoSolidPropertyTypeCreateCallback.OnAfterCreate(stage, toposolidpropertytype)
	}
}

func (toposolidpropertytype *TopoSolidPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSolidPropertyTypeUpdateCallback != nil {
		var frontTopoSolidPropertyType *TopoSolidPropertyType
		if front != nil {
			frontTopoSolidPropertyType, _ = front.(*TopoSolidPropertyType)
		}
		stage.OnAfterTopoSolidPropertyTypeUpdateCallback.OnAfterUpdate(stage, toposolidpropertytype, frontTopoSolidPropertyType)
	}
}

func (toposolidpropertytype *TopoSolidPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSolidPropertyTypeDeleteCallback != nil {
		var frontTopoSolidPropertyType *TopoSolidPropertyType
		if front != nil {
			frontTopoSolidPropertyType, _ = front.(*TopoSolidPropertyType)
		}
		stage.OnAfterTopoSolidPropertyTypeDeleteCallback.OnAfterDelete(stage, toposolidpropertytype, frontTopoSolidPropertyType)
	}
}

func (toposolidtype *TopoSolidType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoSolidTypeCreateCallback != nil {
		stage.OnAfterTopoSolidTypeCreateCallback.OnAfterCreate(stage, toposolidtype)
	}
}

func (toposolidtype *TopoSolidType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSolidTypeUpdateCallback != nil {
		var frontTopoSolidType *TopoSolidType
		if front != nil {
			frontTopoSolidType, _ = front.(*TopoSolidType)
		}
		stage.OnAfterTopoSolidTypeUpdateCallback.OnAfterUpdate(stage, toposolidtype, frontTopoSolidType)
	}
}

func (toposolidtype *TopoSolidType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSolidTypeDeleteCallback != nil {
		var frontTopoSolidType *TopoSolidType
		if front != nil {
			frontTopoSolidType, _ = front.(*TopoSolidType)
		}
		stage.OnAfterTopoSolidTypeDeleteCallback.OnAfterDelete(stage, toposolidtype, frontTopoSolidType)
	}
}

func (toposurfacepropertytype *TopoSurfacePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoSurfacePropertyTypeCreateCallback != nil {
		stage.OnAfterTopoSurfacePropertyTypeCreateCallback.OnAfterCreate(stage, toposurfacepropertytype)
	}
}

func (toposurfacepropertytype *TopoSurfacePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSurfacePropertyTypeUpdateCallback != nil {
		var frontTopoSurfacePropertyType *TopoSurfacePropertyType
		if front != nil {
			frontTopoSurfacePropertyType, _ = front.(*TopoSurfacePropertyType)
		}
		stage.OnAfterTopoSurfacePropertyTypeUpdateCallback.OnAfterUpdate(stage, toposurfacepropertytype, frontTopoSurfacePropertyType)
	}
}

func (toposurfacepropertytype *TopoSurfacePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSurfacePropertyTypeDeleteCallback != nil {
		var frontTopoSurfacePropertyType *TopoSurfacePropertyType
		if front != nil {
			frontTopoSurfacePropertyType, _ = front.(*TopoSurfacePropertyType)
		}
		stage.OnAfterTopoSurfacePropertyTypeDeleteCallback.OnAfterDelete(stage, toposurfacepropertytype, frontTopoSurfacePropertyType)
	}
}

func (toposurfacetype *TopoSurfaceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoSurfaceTypeCreateCallback != nil {
		stage.OnAfterTopoSurfaceTypeCreateCallback.OnAfterCreate(stage, toposurfacetype)
	}
}

func (toposurfacetype *TopoSurfaceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSurfaceTypeUpdateCallback != nil {
		var frontTopoSurfaceType *TopoSurfaceType
		if front != nil {
			frontTopoSurfaceType, _ = front.(*TopoSurfaceType)
		}
		stage.OnAfterTopoSurfaceTypeUpdateCallback.OnAfterUpdate(stage, toposurfacetype, frontTopoSurfaceType)
	}
}

func (toposurfacetype *TopoSurfaceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoSurfaceTypeDeleteCallback != nil {
		var frontTopoSurfaceType *TopoSurfaceType
		if front != nil {
			frontTopoSurfaceType, _ = front.(*TopoSurfaceType)
		}
		stage.OnAfterTopoSurfaceTypeDeleteCallback.OnAfterDelete(stage, toposurfacetype, frontTopoSurfaceType)
	}
}

func (topovolumepropertytype *TopoVolumePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoVolumePropertyTypeCreateCallback != nil {
		stage.OnAfterTopoVolumePropertyTypeCreateCallback.OnAfterCreate(stage, topovolumepropertytype)
	}
}

func (topovolumepropertytype *TopoVolumePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoVolumePropertyTypeUpdateCallback != nil {
		var frontTopoVolumePropertyType *TopoVolumePropertyType
		if front != nil {
			frontTopoVolumePropertyType, _ = front.(*TopoVolumePropertyType)
		}
		stage.OnAfterTopoVolumePropertyTypeUpdateCallback.OnAfterUpdate(stage, topovolumepropertytype, frontTopoVolumePropertyType)
	}
}

func (topovolumepropertytype *TopoVolumePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoVolumePropertyTypeDeleteCallback != nil {
		var frontTopoVolumePropertyType *TopoVolumePropertyType
		if front != nil {
			frontTopoVolumePropertyType, _ = front.(*TopoVolumePropertyType)
		}
		stage.OnAfterTopoVolumePropertyTypeDeleteCallback.OnAfterDelete(stage, topovolumepropertytype, frontTopoVolumePropertyType)
	}
}

func (topovolumetype *TopoVolumeType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopoVolumeTypeCreateCallback != nil {
		stage.OnAfterTopoVolumeTypeCreateCallback.OnAfterCreate(stage, topovolumetype)
	}
}

func (topovolumetype *TopoVolumeType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoVolumeTypeUpdateCallback != nil {
		var frontTopoVolumeType *TopoVolumeType
		if front != nil {
			frontTopoVolumeType, _ = front.(*TopoVolumeType)
		}
		stage.OnAfterTopoVolumeTypeUpdateCallback.OnAfterUpdate(stage, topovolumetype, frontTopoVolumeType)
	}
}

func (topovolumetype *TopoVolumeType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopoVolumeTypeDeleteCallback != nil {
		var frontTopoVolumeType *TopoVolumeType
		if front != nil {
			frontTopoVolumeType, _ = front.(*TopoVolumeType)
		}
		stage.OnAfterTopoVolumeTypeDeleteCallback.OnAfterDelete(stage, topovolumetype, frontTopoVolumeType)
	}
}

func (transformationpropertytype *TransformationPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTransformationPropertyTypeCreateCallback != nil {
		stage.OnAfterTransformationPropertyTypeCreateCallback.OnAfterCreate(stage, transformationpropertytype)
	}
}

func (transformationpropertytype *TransformationPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransformationPropertyTypeUpdateCallback != nil {
		var frontTransformationPropertyType *TransformationPropertyType
		if front != nil {
			frontTransformationPropertyType, _ = front.(*TransformationPropertyType)
		}
		stage.OnAfterTransformationPropertyTypeUpdateCallback.OnAfterUpdate(stage, transformationpropertytype, frontTransformationPropertyType)
	}
}

func (transformationpropertytype *TransformationPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransformationPropertyTypeDeleteCallback != nil {
		var frontTransformationPropertyType *TransformationPropertyType
		if front != nil {
			frontTransformationPropertyType, _ = front.(*TransformationPropertyType)
		}
		stage.OnAfterTransformationPropertyTypeDeleteCallback.OnAfterDelete(stage, transformationpropertytype, frontTransformationPropertyType)
	}
}

func (transformationtype *TransformationType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTransformationTypeCreateCallback != nil {
		stage.OnAfterTransformationTypeCreateCallback.OnAfterCreate(stage, transformationtype)
	}
}

func (transformationtype *TransformationType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransformationTypeUpdateCallback != nil {
		var frontTransformationType *TransformationType
		if front != nil {
			frontTransformationType, _ = front.(*TransformationType)
		}
		stage.OnAfterTransformationTypeUpdateCallback.OnAfterUpdate(stage, transformationtype, frontTransformationType)
	}
}

func (transformationtype *TransformationType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTransformationTypeDeleteCallback != nil {
		var frontTransformationType *TransformationType
		if front != nil {
			frontTransformationType, _ = front.(*TransformationType)
		}
		stage.OnAfterTransformationTypeDeleteCallback.OnAfterDelete(stage, transformationtype, frontTransformationType)
	}
}

func (triangletype *TriangleType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTriangleTypeCreateCallback != nil {
		stage.OnAfterTriangleTypeCreateCallback.OnAfterCreate(stage, triangletype)
	}
}

func (triangletype *TriangleType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTriangleTypeUpdateCallback != nil {
		var frontTriangleType *TriangleType
		if front != nil {
			frontTriangleType, _ = front.(*TriangleType)
		}
		stage.OnAfterTriangleTypeUpdateCallback.OnAfterUpdate(stage, triangletype, frontTriangleType)
	}
}

func (triangletype *TriangleType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTriangleTypeDeleteCallback != nil {
		var frontTriangleType *TriangleType
		if front != nil {
			frontTriangleType, _ = front.(*TriangleType)
		}
		stage.OnAfterTriangleTypeDeleteCallback.OnAfterDelete(stage, triangletype, frontTriangleType)
	}
}

func (typename_propertytype *TypeName_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTypeName_PropertyTypeCreateCallback != nil {
		stage.OnAfterTypeName_PropertyTypeCreateCallback.OnAfterCreate(stage, typename_propertytype)
	}
}

func (typename_propertytype *TypeName_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTypeName_PropertyTypeUpdateCallback != nil {
		var frontTypeName_PropertyType *TypeName_PropertyType
		if front != nil {
			frontTypeName_PropertyType, _ = front.(*TypeName_PropertyType)
		}
		stage.OnAfterTypeName_PropertyTypeUpdateCallback.OnAfterUpdate(stage, typename_propertytype, frontTypeName_PropertyType)
	}
}

func (typename_propertytype *TypeName_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTypeName_PropertyTypeDeleteCallback != nil {
		var frontTypeName_PropertyType *TypeName_PropertyType
		if front != nil {
			frontTypeName_PropertyType, _ = front.(*TypeName_PropertyType)
		}
		stage.OnAfterTypeName_PropertyTypeDeleteCallback.OnAfterDelete(stage, typename_propertytype, frontTypeName_PropertyType)
	}
}

func (typename_type *TypeName_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTypeName_TypeCreateCallback != nil {
		stage.OnAfterTypeName_TypeCreateCallback.OnAfterCreate(stage, typename_type)
	}
}

func (typename_type *TypeName_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTypeName_TypeUpdateCallback != nil {
		var frontTypeName_Type *TypeName_Type
		if front != nil {
			frontTypeName_Type, _ = front.(*TypeName_Type)
		}
		stage.OnAfterTypeName_TypeUpdateCallback.OnAfterUpdate(stage, typename_type, frontTypeName_Type)
	}
}

func (typename_type *TypeName_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTypeName_TypeDeleteCallback != nil {
		var frontTypeName_Type *TypeName_Type
		if front != nil {
			frontTypeName_Type, _ = front.(*TypeName_Type)
		}
		stage.OnAfterTypeName_TypeDeleteCallback.OnAfterDelete(stage, typename_type, frontTypeName_Type)
	}
}

func (url_propertytype *URL_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterURL_PropertyTypeCreateCallback != nil {
		stage.OnAfterURL_PropertyTypeCreateCallback.OnAfterCreate(stage, url_propertytype)
	}
}

func (url_propertytype *URL_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterURL_PropertyTypeUpdateCallback != nil {
		var frontURL_PropertyType *URL_PropertyType
		if front != nil {
			frontURL_PropertyType, _ = front.(*URL_PropertyType)
		}
		stage.OnAfterURL_PropertyTypeUpdateCallback.OnAfterUpdate(stage, url_propertytype, frontURL_PropertyType)
	}
}

func (url_propertytype *URL_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterURL_PropertyTypeDeleteCallback != nil {
		var frontURL_PropertyType *URL_PropertyType
		if front != nil {
			frontURL_PropertyType, _ = front.(*URL_PropertyType)
		}
		stage.OnAfterURL_PropertyTypeDeleteCallback.OnAfterDelete(stage, url_propertytype, frontURL_PropertyType)
	}
}

func (unitdefinitiontype *UnitDefinitionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitDefinitionTypeCreateCallback != nil {
		stage.OnAfterUnitDefinitionTypeCreateCallback.OnAfterCreate(stage, unitdefinitiontype)
	}
}

func (unitdefinitiontype *UnitDefinitionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitDefinitionTypeUpdateCallback != nil {
		var frontUnitDefinitionType *UnitDefinitionType
		if front != nil {
			frontUnitDefinitionType, _ = front.(*UnitDefinitionType)
		}
		stage.OnAfterUnitDefinitionTypeUpdateCallback.OnAfterUpdate(stage, unitdefinitiontype, frontUnitDefinitionType)
	}
}

func (unitdefinitiontype *UnitDefinitionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitDefinitionTypeDeleteCallback != nil {
		var frontUnitDefinitionType *UnitDefinitionType
		if front != nil {
			frontUnitDefinitionType, _ = front.(*UnitDefinitionType)
		}
		stage.OnAfterUnitDefinitionTypeDeleteCallback.OnAfterDelete(stage, unitdefinitiontype, frontUnitDefinitionType)
	}
}

func (unitofmeasuretype *UnitOfMeasureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitOfMeasureTypeCreateCallback != nil {
		stage.OnAfterUnitOfMeasureTypeCreateCallback.OnAfterCreate(stage, unitofmeasuretype)
	}
}

func (unitofmeasuretype *UnitOfMeasureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitOfMeasureTypeUpdateCallback != nil {
		var frontUnitOfMeasureType *UnitOfMeasureType
		if front != nil {
			frontUnitOfMeasureType, _ = front.(*UnitOfMeasureType)
		}
		stage.OnAfterUnitOfMeasureTypeUpdateCallback.OnAfterUpdate(stage, unitofmeasuretype, frontUnitOfMeasureType)
	}
}

func (unitofmeasuretype *UnitOfMeasureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitOfMeasureTypeDeleteCallback != nil {
		var frontUnitOfMeasureType *UnitOfMeasureType
		if front != nil {
			frontUnitOfMeasureType, _ = front.(*UnitOfMeasureType)
		}
		stage.OnAfterUnitOfMeasureTypeDeleteCallback.OnAfterDelete(stage, unitofmeasuretype, frontUnitOfMeasureType)
	}
}

func (unitofmeasure_propertytype *UnitOfMeasure_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnitOfMeasure_PropertyTypeCreateCallback != nil {
		stage.OnAfterUnitOfMeasure_PropertyTypeCreateCallback.OnAfterCreate(stage, unitofmeasure_propertytype)
	}
}

func (unitofmeasure_propertytype *UnitOfMeasure_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitOfMeasure_PropertyTypeUpdateCallback != nil {
		var frontUnitOfMeasure_PropertyType *UnitOfMeasure_PropertyType
		if front != nil {
			frontUnitOfMeasure_PropertyType, _ = front.(*UnitOfMeasure_PropertyType)
		}
		stage.OnAfterUnitOfMeasure_PropertyTypeUpdateCallback.OnAfterUpdate(stage, unitofmeasure_propertytype, frontUnitOfMeasure_PropertyType)
	}
}

func (unitofmeasure_propertytype *UnitOfMeasure_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnitOfMeasure_PropertyTypeDeleteCallback != nil {
		var frontUnitOfMeasure_PropertyType *UnitOfMeasure_PropertyType
		if front != nil {
			frontUnitOfMeasure_PropertyType, _ = front.(*UnitOfMeasure_PropertyType)
		}
		stage.OnAfterUnitOfMeasure_PropertyTypeDeleteCallback.OnAfterDelete(stage, unitofmeasure_propertytype, frontUnitOfMeasure_PropertyType)
	}
}

func (unlimitedinteger_propertytype *UnlimitedInteger_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnlimitedInteger_PropertyTypeCreateCallback != nil {
		stage.OnAfterUnlimitedInteger_PropertyTypeCreateCallback.OnAfterCreate(stage, unlimitedinteger_propertytype)
	}
}

func (unlimitedinteger_propertytype *UnlimitedInteger_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnlimitedInteger_PropertyTypeUpdateCallback != nil {
		var frontUnlimitedInteger_PropertyType *UnlimitedInteger_PropertyType
		if front != nil {
			frontUnlimitedInteger_PropertyType, _ = front.(*UnlimitedInteger_PropertyType)
		}
		stage.OnAfterUnlimitedInteger_PropertyTypeUpdateCallback.OnAfterUpdate(stage, unlimitedinteger_propertytype, frontUnlimitedInteger_PropertyType)
	}
}

func (unlimitedinteger_propertytype *UnlimitedInteger_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnlimitedInteger_PropertyTypeDeleteCallback != nil {
		var frontUnlimitedInteger_PropertyType *UnlimitedInteger_PropertyType
		if front != nil {
			frontUnlimitedInteger_PropertyType, _ = front.(*UnlimitedInteger_PropertyType)
		}
		stage.OnAfterUnlimitedInteger_PropertyTypeDeleteCallback.OnAfterDelete(stage, unlimitedinteger_propertytype, frontUnlimitedInteger_PropertyType)
	}
}

func (unlimitedinteger_type *UnlimitedInteger_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUnlimitedInteger_TypeCreateCallback != nil {
		stage.OnAfterUnlimitedInteger_TypeCreateCallback.OnAfterCreate(stage, unlimitedinteger_type)
	}
}

func (unlimitedinteger_type *UnlimitedInteger_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnlimitedInteger_TypeUpdateCallback != nil {
		var frontUnlimitedInteger_Type *UnlimitedInteger_Type
		if front != nil {
			frontUnlimitedInteger_Type, _ = front.(*UnlimitedInteger_Type)
		}
		stage.OnAfterUnlimitedInteger_TypeUpdateCallback.OnAfterUpdate(stage, unlimitedinteger_type, frontUnlimitedInteger_Type)
	}
}

func (unlimitedinteger_type *UnlimitedInteger_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUnlimitedInteger_TypeDeleteCallback != nil {
		var frontUnlimitedInteger_Type *UnlimitedInteger_Type
		if front != nil {
			frontUnlimitedInteger_Type, _ = front.(*UnlimitedInteger_Type)
		}
		stage.OnAfterUnlimitedInteger_TypeDeleteCallback.OnAfterDelete(stage, unlimitedinteger_type, frontUnlimitedInteger_Type)
	}
}

func (uomangle_propertytype *UomAngle_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUomAngle_PropertyTypeCreateCallback != nil {
		stage.OnAfterUomAngle_PropertyTypeCreateCallback.OnAfterCreate(stage, uomangle_propertytype)
	}
}

func (uomangle_propertytype *UomAngle_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomAngle_PropertyTypeUpdateCallback != nil {
		var frontUomAngle_PropertyType *UomAngle_PropertyType
		if front != nil {
			frontUomAngle_PropertyType, _ = front.(*UomAngle_PropertyType)
		}
		stage.OnAfterUomAngle_PropertyTypeUpdateCallback.OnAfterUpdate(stage, uomangle_propertytype, frontUomAngle_PropertyType)
	}
}

func (uomangle_propertytype *UomAngle_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomAngle_PropertyTypeDeleteCallback != nil {
		var frontUomAngle_PropertyType *UomAngle_PropertyType
		if front != nil {
			frontUomAngle_PropertyType, _ = front.(*UomAngle_PropertyType)
		}
		stage.OnAfterUomAngle_PropertyTypeDeleteCallback.OnAfterDelete(stage, uomangle_propertytype, frontUomAngle_PropertyType)
	}
}

func (uomarea_propertytype *UomArea_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUomArea_PropertyTypeCreateCallback != nil {
		stage.OnAfterUomArea_PropertyTypeCreateCallback.OnAfterCreate(stage, uomarea_propertytype)
	}
}

func (uomarea_propertytype *UomArea_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomArea_PropertyTypeUpdateCallback != nil {
		var frontUomArea_PropertyType *UomArea_PropertyType
		if front != nil {
			frontUomArea_PropertyType, _ = front.(*UomArea_PropertyType)
		}
		stage.OnAfterUomArea_PropertyTypeUpdateCallback.OnAfterUpdate(stage, uomarea_propertytype, frontUomArea_PropertyType)
	}
}

func (uomarea_propertytype *UomArea_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomArea_PropertyTypeDeleteCallback != nil {
		var frontUomArea_PropertyType *UomArea_PropertyType
		if front != nil {
			frontUomArea_PropertyType, _ = front.(*UomArea_PropertyType)
		}
		stage.OnAfterUomArea_PropertyTypeDeleteCallback.OnAfterDelete(stage, uomarea_propertytype, frontUomArea_PropertyType)
	}
}

func (uomlength_propertytype *UomLength_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUomLength_PropertyTypeCreateCallback != nil {
		stage.OnAfterUomLength_PropertyTypeCreateCallback.OnAfterCreate(stage, uomlength_propertytype)
	}
}

func (uomlength_propertytype *UomLength_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomLength_PropertyTypeUpdateCallback != nil {
		var frontUomLength_PropertyType *UomLength_PropertyType
		if front != nil {
			frontUomLength_PropertyType, _ = front.(*UomLength_PropertyType)
		}
		stage.OnAfterUomLength_PropertyTypeUpdateCallback.OnAfterUpdate(stage, uomlength_propertytype, frontUomLength_PropertyType)
	}
}

func (uomlength_propertytype *UomLength_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomLength_PropertyTypeDeleteCallback != nil {
		var frontUomLength_PropertyType *UomLength_PropertyType
		if front != nil {
			frontUomLength_PropertyType, _ = front.(*UomLength_PropertyType)
		}
		stage.OnAfterUomLength_PropertyTypeDeleteCallback.OnAfterDelete(stage, uomlength_propertytype, frontUomLength_PropertyType)
	}
}

func (uomscale_propertytype *UomScale_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUomScale_PropertyTypeCreateCallback != nil {
		stage.OnAfterUomScale_PropertyTypeCreateCallback.OnAfterCreate(stage, uomscale_propertytype)
	}
}

func (uomscale_propertytype *UomScale_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomScale_PropertyTypeUpdateCallback != nil {
		var frontUomScale_PropertyType *UomScale_PropertyType
		if front != nil {
			frontUomScale_PropertyType, _ = front.(*UomScale_PropertyType)
		}
		stage.OnAfterUomScale_PropertyTypeUpdateCallback.OnAfterUpdate(stage, uomscale_propertytype, frontUomScale_PropertyType)
	}
}

func (uomscale_propertytype *UomScale_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomScale_PropertyTypeDeleteCallback != nil {
		var frontUomScale_PropertyType *UomScale_PropertyType
		if front != nil {
			frontUomScale_PropertyType, _ = front.(*UomScale_PropertyType)
		}
		stage.OnAfterUomScale_PropertyTypeDeleteCallback.OnAfterDelete(stage, uomscale_propertytype, frontUomScale_PropertyType)
	}
}

func (uomtime_propertytype *UomTime_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUomTime_PropertyTypeCreateCallback != nil {
		stage.OnAfterUomTime_PropertyTypeCreateCallback.OnAfterCreate(stage, uomtime_propertytype)
	}
}

func (uomtime_propertytype *UomTime_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomTime_PropertyTypeUpdateCallback != nil {
		var frontUomTime_PropertyType *UomTime_PropertyType
		if front != nil {
			frontUomTime_PropertyType, _ = front.(*UomTime_PropertyType)
		}
		stage.OnAfterUomTime_PropertyTypeUpdateCallback.OnAfterUpdate(stage, uomtime_propertytype, frontUomTime_PropertyType)
	}
}

func (uomtime_propertytype *UomTime_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomTime_PropertyTypeDeleteCallback != nil {
		var frontUomTime_PropertyType *UomTime_PropertyType
		if front != nil {
			frontUomTime_PropertyType, _ = front.(*UomTime_PropertyType)
		}
		stage.OnAfterUomTime_PropertyTypeDeleteCallback.OnAfterDelete(stage, uomtime_propertytype, frontUomTime_PropertyType)
	}
}

func (uomvelocity_propertytype *UomVelocity_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUomVelocity_PropertyTypeCreateCallback != nil {
		stage.OnAfterUomVelocity_PropertyTypeCreateCallback.OnAfterCreate(stage, uomvelocity_propertytype)
	}
}

func (uomvelocity_propertytype *UomVelocity_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomVelocity_PropertyTypeUpdateCallback != nil {
		var frontUomVelocity_PropertyType *UomVelocity_PropertyType
		if front != nil {
			frontUomVelocity_PropertyType, _ = front.(*UomVelocity_PropertyType)
		}
		stage.OnAfterUomVelocity_PropertyTypeUpdateCallback.OnAfterUpdate(stage, uomvelocity_propertytype, frontUomVelocity_PropertyType)
	}
}

func (uomvelocity_propertytype *UomVelocity_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomVelocity_PropertyTypeDeleteCallback != nil {
		var frontUomVelocity_PropertyType *UomVelocity_PropertyType
		if front != nil {
			frontUomVelocity_PropertyType, _ = front.(*UomVelocity_PropertyType)
		}
		stage.OnAfterUomVelocity_PropertyTypeDeleteCallback.OnAfterDelete(stage, uomvelocity_propertytype, frontUomVelocity_PropertyType)
	}
}

func (uomvolume_propertytype *UomVolume_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUomVolume_PropertyTypeCreateCallback != nil {
		stage.OnAfterUomVolume_PropertyTypeCreateCallback.OnAfterCreate(stage, uomvolume_propertytype)
	}
}

func (uomvolume_propertytype *UomVolume_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomVolume_PropertyTypeUpdateCallback != nil {
		var frontUomVolume_PropertyType *UomVolume_PropertyType
		if front != nil {
			frontUomVolume_PropertyType, _ = front.(*UomVolume_PropertyType)
		}
		stage.OnAfterUomVolume_PropertyTypeUpdateCallback.OnAfterUpdate(stage, uomvolume_propertytype, frontUomVolume_PropertyType)
	}
}

func (uomvolume_propertytype *UomVolume_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUomVolume_PropertyTypeDeleteCallback != nil {
		var frontUomVolume_PropertyType *UomVolume_PropertyType
		if front != nil {
			frontUomVolume_PropertyType, _ = front.(*UomVolume_PropertyType)
		}
		stage.OnAfterUomVolume_PropertyTypeDeleteCallback.OnAfterDelete(stage, uomvolume_propertytype, frontUomVolume_PropertyType)
	}
}

func (userdefinedcspropertytype *UserDefinedCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUserDefinedCSPropertyTypeCreateCallback != nil {
		stage.OnAfterUserDefinedCSPropertyTypeCreateCallback.OnAfterCreate(stage, userdefinedcspropertytype)
	}
}

func (userdefinedcspropertytype *UserDefinedCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserDefinedCSPropertyTypeUpdateCallback != nil {
		var frontUserDefinedCSPropertyType *UserDefinedCSPropertyType
		if front != nil {
			frontUserDefinedCSPropertyType, _ = front.(*UserDefinedCSPropertyType)
		}
		stage.OnAfterUserDefinedCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, userdefinedcspropertytype, frontUserDefinedCSPropertyType)
	}
}

func (userdefinedcspropertytype *UserDefinedCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserDefinedCSPropertyTypeDeleteCallback != nil {
		var frontUserDefinedCSPropertyType *UserDefinedCSPropertyType
		if front != nil {
			frontUserDefinedCSPropertyType, _ = front.(*UserDefinedCSPropertyType)
		}
		stage.OnAfterUserDefinedCSPropertyTypeDeleteCallback.OnAfterDelete(stage, userdefinedcspropertytype, frontUserDefinedCSPropertyType)
	}
}

func (userdefinedcstype *UserDefinedCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterUserDefinedCSTypeCreateCallback != nil {
		stage.OnAfterUserDefinedCSTypeCreateCallback.OnAfterCreate(stage, userdefinedcstype)
	}
}

func (userdefinedcstype *UserDefinedCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserDefinedCSTypeUpdateCallback != nil {
		var frontUserDefinedCSType *UserDefinedCSType
		if front != nil {
			frontUserDefinedCSType, _ = front.(*UserDefinedCSType)
		}
		stage.OnAfterUserDefinedCSTypeUpdateCallback.OnAfterUpdate(stage, userdefinedcstype, frontUserDefinedCSType)
	}
}

func (userdefinedcstype *UserDefinedCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterUserDefinedCSTypeDeleteCallback != nil {
		var frontUserDefinedCSType *UserDefinedCSType
		if front != nil {
			frontUserDefinedCSType, _ = front.(*UserDefinedCSType)
		}
		stage.OnAfterUserDefinedCSTypeDeleteCallback.OnAfterDelete(stage, userdefinedcstype, frontUserDefinedCSType)
	}
}

func (valuearraypropertytype *ValueArrayPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValueArrayPropertyTypeCreateCallback != nil {
		stage.OnAfterValueArrayPropertyTypeCreateCallback.OnAfterCreate(stage, valuearraypropertytype)
	}
}

func (valuearraypropertytype *ValueArrayPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValueArrayPropertyTypeUpdateCallback != nil {
		var frontValueArrayPropertyType *ValueArrayPropertyType
		if front != nil {
			frontValueArrayPropertyType, _ = front.(*ValueArrayPropertyType)
		}
		stage.OnAfterValueArrayPropertyTypeUpdateCallback.OnAfterUpdate(stage, valuearraypropertytype, frontValueArrayPropertyType)
	}
}

func (valuearraypropertytype *ValueArrayPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValueArrayPropertyTypeDeleteCallback != nil {
		var frontValueArrayPropertyType *ValueArrayPropertyType
		if front != nil {
			frontValueArrayPropertyType, _ = front.(*ValueArrayPropertyType)
		}
		stage.OnAfterValueArrayPropertyTypeDeleteCallback.OnAfterDelete(stage, valuearraypropertytype, frontValueArrayPropertyType)
	}
}

func (valuearraytype *ValueArrayType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValueArrayTypeCreateCallback != nil {
		stage.OnAfterValueArrayTypeCreateCallback.OnAfterCreate(stage, valuearraytype)
	}
}

func (valuearraytype *ValueArrayType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValueArrayTypeUpdateCallback != nil {
		var frontValueArrayType *ValueArrayType
		if front != nil {
			frontValueArrayType, _ = front.(*ValueArrayType)
		}
		stage.OnAfterValueArrayTypeUpdateCallback.OnAfterUpdate(stage, valuearraytype, frontValueArrayType)
	}
}

func (valuearraytype *ValueArrayType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValueArrayTypeDeleteCallback != nil {
		var frontValueArrayType *ValueArrayType
		if front != nil {
			frontValueArrayType, _ = front.(*ValueArrayType)
		}
		stage.OnAfterValueArrayTypeDeleteCallback.OnAfterDelete(stage, valuearraytype, frontValueArrayType)
	}
}

func (valuepropertytype *ValuePropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterValuePropertyTypeCreateCallback != nil {
		stage.OnAfterValuePropertyTypeCreateCallback.OnAfterCreate(stage, valuepropertytype)
	}
}

func (valuepropertytype *ValuePropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValuePropertyTypeUpdateCallback != nil {
		var frontValuePropertyType *ValuePropertyType
		if front != nil {
			frontValuePropertyType, _ = front.(*ValuePropertyType)
		}
		stage.OnAfterValuePropertyTypeUpdateCallback.OnAfterUpdate(stage, valuepropertytype, frontValuePropertyType)
	}
}

func (valuepropertytype *ValuePropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterValuePropertyTypeDeleteCallback != nil {
		var frontValuePropertyType *ValuePropertyType
		if front != nil {
			frontValuePropertyType, _ = front.(*ValuePropertyType)
		}
		stage.OnAfterValuePropertyTypeDeleteCallback.OnAfterDelete(stage, valuepropertytype, frontValuePropertyType)
	}
}

func (vectortype *VectorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVectorTypeCreateCallback != nil {
		stage.OnAfterVectorTypeCreateCallback.OnAfterCreate(stage, vectortype)
	}
}

func (vectortype *VectorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVectorTypeUpdateCallback != nil {
		var frontVectorType *VectorType
		if front != nil {
			frontVectorType, _ = front.(*VectorType)
		}
		stage.OnAfterVectorTypeUpdateCallback.OnAfterUpdate(stage, vectortype, frontVectorType)
	}
}

func (vectortype *VectorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVectorTypeDeleteCallback != nil {
		var frontVectorType *VectorType
		if front != nil {
			frontVectorType, _ = front.(*VectorType)
		}
		stage.OnAfterVectorTypeDeleteCallback.OnAfterDelete(stage, vectortype, frontVectorType)
	}
}

func (verticalcrspropertytype *VerticalCRSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalCRSPropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalCRSPropertyTypeCreateCallback.OnAfterCreate(stage, verticalcrspropertytype)
	}
}

func (verticalcrspropertytype *VerticalCRSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCRSPropertyTypeUpdateCallback != nil {
		var frontVerticalCRSPropertyType *VerticalCRSPropertyType
		if front != nil {
			frontVerticalCRSPropertyType, _ = front.(*VerticalCRSPropertyType)
		}
		stage.OnAfterVerticalCRSPropertyTypeUpdateCallback.OnAfterUpdate(stage, verticalcrspropertytype, frontVerticalCRSPropertyType)
	}
}

func (verticalcrspropertytype *VerticalCRSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCRSPropertyTypeDeleteCallback != nil {
		var frontVerticalCRSPropertyType *VerticalCRSPropertyType
		if front != nil {
			frontVerticalCRSPropertyType, _ = front.(*VerticalCRSPropertyType)
		}
		stage.OnAfterVerticalCRSPropertyTypeDeleteCallback.OnAfterDelete(stage, verticalcrspropertytype, frontVerticalCRSPropertyType)
	}
}

func (verticalcrstype *VerticalCRSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalCRSTypeCreateCallback != nil {
		stage.OnAfterVerticalCRSTypeCreateCallback.OnAfterCreate(stage, verticalcrstype)
	}
}

func (verticalcrstype *VerticalCRSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCRSTypeUpdateCallback != nil {
		var frontVerticalCRSType *VerticalCRSType
		if front != nil {
			frontVerticalCRSType, _ = front.(*VerticalCRSType)
		}
		stage.OnAfterVerticalCRSTypeUpdateCallback.OnAfterUpdate(stage, verticalcrstype, frontVerticalCRSType)
	}
}

func (verticalcrstype *VerticalCRSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCRSTypeDeleteCallback != nil {
		var frontVerticalCRSType *VerticalCRSType
		if front != nil {
			frontVerticalCRSType, _ = front.(*VerticalCRSType)
		}
		stage.OnAfterVerticalCRSTypeDeleteCallback.OnAfterDelete(stage, verticalcrstype, frontVerticalCRSType)
	}
}

func (verticalcspropertytype *VerticalCSPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalCSPropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalCSPropertyTypeCreateCallback.OnAfterCreate(stage, verticalcspropertytype)
	}
}

func (verticalcspropertytype *VerticalCSPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCSPropertyTypeUpdateCallback != nil {
		var frontVerticalCSPropertyType *VerticalCSPropertyType
		if front != nil {
			frontVerticalCSPropertyType, _ = front.(*VerticalCSPropertyType)
		}
		stage.OnAfterVerticalCSPropertyTypeUpdateCallback.OnAfterUpdate(stage, verticalcspropertytype, frontVerticalCSPropertyType)
	}
}

func (verticalcspropertytype *VerticalCSPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCSPropertyTypeDeleteCallback != nil {
		var frontVerticalCSPropertyType *VerticalCSPropertyType
		if front != nil {
			frontVerticalCSPropertyType, _ = front.(*VerticalCSPropertyType)
		}
		stage.OnAfterVerticalCSPropertyTypeDeleteCallback.OnAfterDelete(stage, verticalcspropertytype, frontVerticalCSPropertyType)
	}
}

func (verticalcstype *VerticalCSType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalCSTypeCreateCallback != nil {
		stage.OnAfterVerticalCSTypeCreateCallback.OnAfterCreate(stage, verticalcstype)
	}
}

func (verticalcstype *VerticalCSType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCSTypeUpdateCallback != nil {
		var frontVerticalCSType *VerticalCSType
		if front != nil {
			frontVerticalCSType, _ = front.(*VerticalCSType)
		}
		stage.OnAfterVerticalCSTypeUpdateCallback.OnAfterUpdate(stage, verticalcstype, frontVerticalCSType)
	}
}

func (verticalcstype *VerticalCSType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalCSTypeDeleteCallback != nil {
		var frontVerticalCSType *VerticalCSType
		if front != nil {
			frontVerticalCSType, _ = front.(*VerticalCSType)
		}
		stage.OnAfterVerticalCSTypeDeleteCallback.OnAfterDelete(stage, verticalcstype, frontVerticalCSType)
	}
}

func (verticaldatumpropertytype *VerticalDatumPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalDatumPropertyTypeCreateCallback != nil {
		stage.OnAfterVerticalDatumPropertyTypeCreateCallback.OnAfterCreate(stage, verticaldatumpropertytype)
	}
}

func (verticaldatumpropertytype *VerticalDatumPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalDatumPropertyTypeUpdateCallback != nil {
		var frontVerticalDatumPropertyType *VerticalDatumPropertyType
		if front != nil {
			frontVerticalDatumPropertyType, _ = front.(*VerticalDatumPropertyType)
		}
		stage.OnAfterVerticalDatumPropertyTypeUpdateCallback.OnAfterUpdate(stage, verticaldatumpropertytype, frontVerticalDatumPropertyType)
	}
}

func (verticaldatumpropertytype *VerticalDatumPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalDatumPropertyTypeDeleteCallback != nil {
		var frontVerticalDatumPropertyType *VerticalDatumPropertyType
		if front != nil {
			frontVerticalDatumPropertyType, _ = front.(*VerticalDatumPropertyType)
		}
		stage.OnAfterVerticalDatumPropertyTypeDeleteCallback.OnAfterDelete(stage, verticaldatumpropertytype, frontVerticalDatumPropertyType)
	}
}

func (verticaldatumtype *VerticalDatumType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalDatumTypeCreateCallback != nil {
		stage.OnAfterVerticalDatumTypeCreateCallback.OnAfterCreate(stage, verticaldatumtype)
	}
}

func (verticaldatumtype *VerticalDatumType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalDatumTypeUpdateCallback != nil {
		var frontVerticalDatumType *VerticalDatumType
		if front != nil {
			frontVerticalDatumType, _ = front.(*VerticalDatumType)
		}
		stage.OnAfterVerticalDatumTypeUpdateCallback.OnAfterUpdate(stage, verticaldatumtype, frontVerticalDatumType)
	}
}

func (verticaldatumtype *VerticalDatumType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalDatumTypeDeleteCallback != nil {
		var frontVerticalDatumType *VerticalDatumType
		if front != nil {
			frontVerticalDatumType, _ = front.(*VerticalDatumType)
		}
		stage.OnAfterVerticalDatumTypeDeleteCallback.OnAfterDelete(stage, verticaldatumtype, frontVerticalDatumType)
	}
}

