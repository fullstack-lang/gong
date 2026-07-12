// generated code - do not edit
package models

// insertion point
func (inst *AxesShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *CircleGridShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *ExplanationTextShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GridPathShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GrowthCurveBezierShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GrowthCurveBezierShapeGrid":
		switch reverseField.Fieldname {
		case "GrowthCurveBezierShapes":
			if _growthcurvebeziershapegrid, ok := stage.GrowthCurveBezierShapeGrid_GrowthCurveBezierShapes_reverseMap[inst]; ok {
				res = _growthcurvebeziershapegrid.Name
			}
		}
	}
	return
}

func (inst *GrowthCurveBezierShapeGrid) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GrowthCurveRhombusGridShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *GrowthCurveRhombusShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "GrowthCurveRhombusGridShape":
		switch reverseField.Fieldname {
		case "GrowthCurveRhombusShapes":
			if _growthcurverhombusgridshape, ok := stage.GrowthCurveRhombusGridShape_GrowthCurveRhombusShapes_reverseMap[inst]; ok {
				res = _growthcurverhombusgridshape.Name
			}
		}
	}
	return
}

func (inst *GrowthVectorShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *InitialRhombusGridShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *InitialRhombusShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "InitialRhombusGridShape":
		switch reverseField.Fieldname {
		case "InitialRhombusShapes":
			if _initialrhombusgridshape, ok := stage.InitialRhombusGridShape_InitialRhombusShapes_reverseMap[inst]; ok {
				res = _initialrhombusgridshape.Name
			}
		}
	}
	return
}

func (inst *Library) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			if _library, ok := stage.Library_SubLibraries_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *NextCircleShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *PerpendicularVector) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "PerpendicularVectorGrid":
		switch reverseField.Fieldname {
		case "PerpendicularVectors":
			if _perpendicularvectorgrid, ok := stage.PerpendicularVectorGrid_PerpendicularVectors_reverseMap[inst]; ok {
				res = _perpendicularvectorgrid.Name
			}
		}
	}
	return
}

func (inst *PerpendicularVectorGrid) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *PerpendicularVectorGridHalfway) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *PerpendicularVectorHalfway) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "PerpendicularVectorGridHalfway":
		switch reverseField.Fieldname {
		case "PerpendicularVectorHalfways":
			if _perpendicularvectorgridhalfway, ok := stage.PerpendicularVectorGridHalfway_PerpendicularVectorHalfways_reverseMap[inst]; ok {
				res = _perpendicularvectorgridhalfway.Name
			}
		}
	}
	return
}

func (inst *Plant) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Plants":
			if _library, ok := stage.Library_Plants_reverseMap[inst]; ok {
				res = _library.Name
			}
		}
	}
	return
}

func (inst *PlantCircumferenceShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *PlantDiagram) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Plant":
		switch reverseField.Fieldname {
		case "PlantDiagrams":
			if _plant, ok := stage.Plant_PlantDiagrams_reverseMap[inst]; ok {
				res = _plant.Name
			}
		}
	}
	return
}

func (inst *Rendered3DShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *RhombusShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *RotatedRhombusGridShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *RotatedRhombusShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "RotatedRhombusGridShape":
		switch reverseField.Fieldname {
		case "RotatedRhombusShapes":
			if _rotatedrhombusgridshape, ok := stage.RotatedRhombusGridShape_RotatedRhombusShapes_reverseMap[inst]; ok {
				res = _rotatedrhombusgridshape.Name
			}
		}
	}
	return
}

func (inst *StackGrowthCurveBezierShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "StackOfGrowthCurve":
		switch reverseField.Fieldname {
		case "StackGrowthCurveBezierShapes":
			if _stackofgrowthcurve, ok := stage.StackOfGrowthCurve_StackGrowthCurveBezierShapes_reverseMap[inst]; ok {
				res = _stackofgrowthcurve.Name
			}
		}
	}
	return
}

func (inst *StackOfGrowthCurve) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *StartArcShape) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "StartArcShapeGrid":
		switch reverseField.Fieldname {
		case "StartArcShapes":
			if _startarcshapegrid, ok := stage.StartArcShapeGrid_StartArcShapes_reverseMap[inst]; ok {
				res = _startarcshapegrid.Name
			}
		}
	}
	return
}

func (inst *StartArcShapeGrid) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

// insertion point
func (inst *AxesShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *CircleGridShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *ExplanationTextShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GridPathShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GrowthCurveBezierShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GrowthCurveBezierShapeGrid":
		switch reverseField.Fieldname {
		case "GrowthCurveBezierShapes":
			res = stage.GrowthCurveBezierShapeGrid_GrowthCurveBezierShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *GrowthCurveBezierShapeGrid) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GrowthCurveRhombusGridShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *GrowthCurveRhombusShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "GrowthCurveRhombusGridShape":
		switch reverseField.Fieldname {
		case "GrowthCurveRhombusShapes":
			res = stage.GrowthCurveRhombusGridShape_GrowthCurveRhombusShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *GrowthVectorShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *InitialRhombusGridShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *InitialRhombusShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "InitialRhombusGridShape":
		switch reverseField.Fieldname {
		case "InitialRhombusShapes":
			res = stage.InitialRhombusGridShape_InitialRhombusShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *Library) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "SubLibraries":
			res = stage.Library_SubLibraries_reverseMap[inst]
		}
	}
	return res
}

func (inst *NextCircleShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *PerpendicularVector) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "PerpendicularVectorGrid":
		switch reverseField.Fieldname {
		case "PerpendicularVectors":
			res = stage.PerpendicularVectorGrid_PerpendicularVectors_reverseMap[inst]
		}
	}
	return res
}

func (inst *PerpendicularVectorGrid) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *PerpendicularVectorGridHalfway) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *PerpendicularVectorHalfway) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "PerpendicularVectorGridHalfway":
		switch reverseField.Fieldname {
		case "PerpendicularVectorHalfways":
			res = stage.PerpendicularVectorGridHalfway_PerpendicularVectorHalfways_reverseMap[inst]
		}
	}
	return res
}

func (inst *Plant) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Library":
		switch reverseField.Fieldname {
		case "Plants":
			res = stage.Library_Plants_reverseMap[inst]
		}
	}
	return res
}

func (inst *PlantCircumferenceShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *PlantDiagram) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Plant":
		switch reverseField.Fieldname {
		case "PlantDiagrams":
			res = stage.Plant_PlantDiagrams_reverseMap[inst]
		}
	}
	return res
}

func (inst *Rendered3DShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *RhombusShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *RotatedRhombusGridShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *RotatedRhombusShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "RotatedRhombusGridShape":
		switch reverseField.Fieldname {
		case "RotatedRhombusShapes":
			res = stage.RotatedRhombusGridShape_RotatedRhombusShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *StackGrowthCurveBezierShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "StackOfGrowthCurve":
		switch reverseField.Fieldname {
		case "StackGrowthCurveBezierShapes":
			res = stage.StackOfGrowthCurve_StackGrowthCurveBezierShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *StackOfGrowthCurve) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *StartArcShape) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "StartArcShapeGrid":
		switch reverseField.Fieldname {
		case "StartArcShapes":
			res = stage.StartArcShapeGrid_StartArcShapes_reverseMap[inst]
		}
	}
	return res
}

func (inst *StartArcShapeGrid) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}
