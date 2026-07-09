package models

func (stager *Stager) enforceOrphanShapeRemove() (needCommit bool) {
	stage := stager.stage

	// Sets of referenced shapes
	refAxes := make(map[*AxesShape]bool)
	refPlantCirc := make(map[*PlantCircumferenceShape]bool)
	refGridPath := make(map[*GridPathShape]bool)
	refCircleGrid := make(map[*CircleGridShape]bool)
	refNextCircle := make(map[*NextCircleShape]bool)
	refExplanation := make(map[*ExplanationTextShape]bool)
	refGrowthVector := make(map[*GrowthVectorShape]bool)

	refInitialGrid := make(map[*InitialRhombusGridShape]bool)
	refRotatedGrid := make(map[*RotatedRhombusGridShape]bool)
	refGrowthCurveGrid := make(map[*GrowthCurveRhombusGridShape]bool)
	refVectorGrid := make(map[*PerpendicularVectorGrid]bool)

	refInitialShape := make(map[*InitialRhombusShape]bool)
	refRhombusShape := make(map[*RhombusShape]bool)
	refRotatedShape := make(map[*RotatedRhombusShape]bool)
	refGrowthCurveShape := make(map[*GrowthCurveRhombusShape]bool)
	refPerpendicularVector := make(map[*PerpendicularVector]bool)

	// Collect referenced shapes from all plants
	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		if plant.AxesShape != nil {
			refAxes[plant.AxesShape] = true
		}
		if plant.ReferenceRhombus != nil {
			refRhombusShape[plant.ReferenceRhombus] = true
		}
		if plant.PlantCircumferenceShape != nil {
			refPlantCirc[plant.PlantCircumferenceShape] = true
		}
		if plant.GridPathShape != nil {
			refGridPath[plant.GridPathShape] = true
		}
		if plant.ExplanationTextShape != nil {
			refExplanation[plant.ExplanationTextShape] = true
		}
		if plant.RotatedReferenceRhombus != nil {
			refRhombusShape[plant.RotatedReferenceRhombus] = true
		}
		if plant.RotatedPlantCircumferenceShape != nil {
			refPlantCirc[plant.RotatedPlantCircumferenceShape] = true
		}
		if plant.RotatedGridPathShape != nil {
			refGridPath[plant.RotatedGridPathShape] = true
		}
		if plant.GrowthVectorShape != nil {
			refGrowthVector[plant.GrowthVectorShape] = true
		}

		if plant.InitialRhombusGridShape != nil {
			refInitialGrid[plant.InitialRhombusGridShape] = true
			for _, shape := range plant.InitialRhombusGridShape.InitialRhombusShapes {
				if shape != nil {
					refInitialShape[shape] = true
				}
			}
		}

		if plant.RotatedRhombusGridShape2 != nil {
			refRotatedGrid[plant.RotatedRhombusGridShape2] = true
			for _, shape := range plant.RotatedRhombusGridShape2.RotatedRhombusShapes {
				if shape != nil {
					refRotatedShape[shape] = true
				}
			}
		}

		if plant.GrowthCurveRhombusGridShape != nil {
			refGrowthCurveGrid[plant.GrowthCurveRhombusGridShape] = true
			for _, shape := range plant.GrowthCurveRhombusGridShape.GrowthCurveRhombusShapes {
				if shape != nil {
					refGrowthCurveShape[shape] = true
				}
			}
		}

		if plant.PerpendicularVectorGrid != nil {
			refVectorGrid[plant.PerpendicularVectorGrid] = true
			for _, vec := range plant.PerpendicularVectorGrid.PerpendicularVectors {
				if vec != nil {
					refPerpendicularVector[vec] = true
				}
			}
		}
	}

	// Unstage unreferenced shapes
	for shape := range *GetGongstructInstancesSetFromPointerType[*AxesShape](stage) {
		if !refAxes[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*PlantCircumferenceShape](stage) {
		if !refPlantCirc[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*GridPathShape](stage) {
		if !refGridPath[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*CircleGridShape](stage) {
		if !refCircleGrid[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*NextCircleShape](stage) {
		if !refNextCircle[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*ExplanationTextShape](stage) {
		if !refExplanation[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthVectorShape](stage) {
		if !refGrowthVector[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}

	for grid := range *GetGongstructInstancesSetFromPointerType[*InitialRhombusGridShape](stage) {
		if !refInitialGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}
	for grid := range *GetGongstructInstancesSetFromPointerType[*RotatedRhombusGridShape](stage) {
		if !refRotatedGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}
	for grid := range *GetGongstructInstancesSetFromPointerType[*GrowthCurveRhombusGridShape](stage) {
		if !refGrowthCurveGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}
	for grid := range *GetGongstructInstancesSetFromPointerType[*PerpendicularVectorGrid](stage) {
		if !refVectorGrid[grid] {
			grid.Unstage(stage)
			needCommit = true
		}
	}

	for shape := range *GetGongstructInstancesSetFromPointerType[*RhombusShape](stage) {
		if !refRhombusShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*InitialRhombusShape](stage) {
		if !refInitialShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*RotatedRhombusShape](stage) {
		if !refRotatedShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for shape := range *GetGongstructInstancesSetFromPointerType[*GrowthCurveRhombusShape](stage) {
		if !refGrowthCurveShape[shape] {
			shape.Unstage(stage)
			needCommit = true
		}
	}
	for vec := range *GetGongstructInstancesSetFromPointerType[*PerpendicularVector](stage) {
		if !refPerpendicularVector[vec] {
			vec.Unstage(stage)
			needCommit = true
		}
	}

	return
}
