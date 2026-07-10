// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AxesShape:
		if stage.OnAfterAxesShapeCreateCallback != nil {
			stage.OnAfterAxesShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeCreateCallback != nil {
			stage.OnAfterCircleGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExplanationTextShape:
		if stage.OnAfterExplanationTextShapeCreateCallback != nil {
			stage.OnAfterExplanationTextShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeCreateCallback != nil {
			stage.OnAfterGridPathShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurveBezierShape:
		if stage.OnAfterGrowthCurveBezierShapeCreateCallback != nil {
			stage.OnAfterGrowthCurveBezierShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurveBezierShapeGrid:
		if stage.OnAfterGrowthCurveBezierShapeGridCreateCallback != nil {
			stage.OnAfterGrowthCurveBezierShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurveRhombusGridShape:
		if stage.OnAfterGrowthCurveRhombusGridShapeCreateCallback != nil {
			stage.OnAfterGrowthCurveRhombusGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurveRhombusShape:
		if stage.OnAfterGrowthCurveRhombusShapeCreateCallback != nil {
			stage.OnAfterGrowthCurveRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeCreateCallback != nil {
			stage.OnAfterGrowthVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *InitialRhombusGridShape:
		if stage.OnAfterInitialRhombusGridShapeCreateCallback != nil {
			stage.OnAfterInitialRhombusGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *InitialRhombusShape:
		if stage.OnAfterInitialRhombusShapeCreateCallback != nil {
			stage.OnAfterInitialRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *NextCircleShape:
		if stage.OnAfterNextCircleShapeCreateCallback != nil {
			stage.OnAfterNextCircleShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerpendicularVector:
		if stage.OnAfterPerpendicularVectorCreateCallback != nil {
			stage.OnAfterPerpendicularVectorCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerpendicularVectorGrid:
		if stage.OnAfterPerpendicularVectorGridCreateCallback != nil {
			stage.OnAfterPerpendicularVectorGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *Plant:
		if stage.OnAfterPlantCreateCallback != nil {
			stage.OnAfterPlantCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlantCircumferenceShape:
		if stage.OnAfterPlantCircumferenceShapeCreateCallback != nil {
			stage.OnAfterPlantCircumferenceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramCreateCallback != nil {
			stage.OnAfterPlantDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *RhombusShape:
		if stage.OnAfterRhombusShapeCreateCallback != nil {
			stage.OnAfterRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *RotatedRhombusGridShape:
		if stage.OnAfterRotatedRhombusGridShapeCreateCallback != nil {
			stage.OnAfterRotatedRhombusGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *RotatedRhombusShape:
		if stage.OnAfterRotatedRhombusShapeCreateCallback != nil {
			stage.OnAfterRotatedRhombusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurveBezierShape:
		if stage.OnAfterStackGrowthCurveBezierShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurveBezierShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfGrowthCurve:
		if stage.OnAfterStackOfGrowthCurveCreateCallback != nil {
			stage.OnAfterStackOfGrowthCurveCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AxesShape:
		newTarget := any(new).(*AxesShape)
		if stage.OnAfterAxesShapeUpdateCallback != nil {
			stage.OnAfterAxesShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CircleGridShape:
		newTarget := any(new).(*CircleGridShape)
		if stage.OnAfterCircleGridShapeUpdateCallback != nil {
			stage.OnAfterCircleGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ExplanationTextShape:
		newTarget := any(new).(*ExplanationTextShape)
		if stage.OnAfterExplanationTextShapeUpdateCallback != nil {
			stage.OnAfterExplanationTextShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GridPathShape:
		newTarget := any(new).(*GridPathShape)
		if stage.OnAfterGridPathShapeUpdateCallback != nil {
			stage.OnAfterGridPathShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurveBezierShape:
		newTarget := any(new).(*GrowthCurveBezierShape)
		if stage.OnAfterGrowthCurveBezierShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurveBezierShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurveBezierShapeGrid:
		newTarget := any(new).(*GrowthCurveBezierShapeGrid)
		if stage.OnAfterGrowthCurveBezierShapeGridUpdateCallback != nil {
			stage.OnAfterGrowthCurveBezierShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurveRhombusGridShape:
		newTarget := any(new).(*GrowthCurveRhombusGridShape)
		if stage.OnAfterGrowthCurveRhombusGridShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurveRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurveRhombusShape:
		newTarget := any(new).(*GrowthCurveRhombusShape)
		if stage.OnAfterGrowthCurveRhombusShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurveRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthVectorShape:
		newTarget := any(new).(*GrowthVectorShape)
		if stage.OnAfterGrowthVectorShapeUpdateCallback != nil {
			stage.OnAfterGrowthVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *InitialRhombusGridShape:
		newTarget := any(new).(*InitialRhombusGridShape)
		if stage.OnAfterInitialRhombusGridShapeUpdateCallback != nil {
			stage.OnAfterInitialRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *InitialRhombusShape:
		newTarget := any(new).(*InitialRhombusShape)
		if stage.OnAfterInitialRhombusShapeUpdateCallback != nil {
			stage.OnAfterInitialRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NextCircleShape:
		newTarget := any(new).(*NextCircleShape)
		if stage.OnAfterNextCircleShapeUpdateCallback != nil {
			stage.OnAfterNextCircleShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerpendicularVector:
		newTarget := any(new).(*PerpendicularVector)
		if stage.OnAfterPerpendicularVectorUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerpendicularVectorGrid:
		newTarget := any(new).(*PerpendicularVectorGrid)
		if stage.OnAfterPerpendicularVectorGridUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Plant:
		newTarget := any(new).(*Plant)
		if stage.OnAfterPlantUpdateCallback != nil {
			stage.OnAfterPlantUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PlantCircumferenceShape:
		newTarget := any(new).(*PlantCircumferenceShape)
		if stage.OnAfterPlantCircumferenceShapeUpdateCallback != nil {
			stage.OnAfterPlantCircumferenceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PlantDiagram:
		newTarget := any(new).(*PlantDiagram)
		if stage.OnAfterPlantDiagramUpdateCallback != nil {
			stage.OnAfterPlantDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RhombusShape:
		newTarget := any(new).(*RhombusShape)
		if stage.OnAfterRhombusShapeUpdateCallback != nil {
			stage.OnAfterRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RotatedRhombusGridShape:
		newTarget := any(new).(*RotatedRhombusGridShape)
		if stage.OnAfterRotatedRhombusGridShapeUpdateCallback != nil {
			stage.OnAfterRotatedRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RotatedRhombusShape:
		newTarget := any(new).(*RotatedRhombusShape)
		if stage.OnAfterRotatedRhombusShapeUpdateCallback != nil {
			stage.OnAfterRotatedRhombusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurveBezierShape:
		newTarget := any(new).(*StackGrowthCurveBezierShape)
		if stage.OnAfterStackGrowthCurveBezierShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurveBezierShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfGrowthCurve:
		newTarget := any(new).(*StackOfGrowthCurve)
		if stage.OnAfterStackOfGrowthCurveUpdateCallback != nil {
			stage.OnAfterStackOfGrowthCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AxesShape:
		if stage.OnAfterAxesShapeDeleteCallback != nil {
			staged := any(staged).(*AxesShape)
			stage.OnAfterAxesShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeDeleteCallback != nil {
			staged := any(staged).(*CircleGridShape)
			stage.OnAfterCircleGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ExplanationTextShape:
		if stage.OnAfterExplanationTextShapeDeleteCallback != nil {
			staged := any(staged).(*ExplanationTextShape)
			stage.OnAfterExplanationTextShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeDeleteCallback != nil {
			staged := any(staged).(*GridPathShape)
			stage.OnAfterGridPathShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurveBezierShape:
		if stage.OnAfterGrowthCurveBezierShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurveBezierShape)
			stage.OnAfterGrowthCurveBezierShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurveBezierShapeGrid:
		if stage.OnAfterGrowthCurveBezierShapeGridDeleteCallback != nil {
			staged := any(staged).(*GrowthCurveBezierShapeGrid)
			stage.OnAfterGrowthCurveBezierShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurveRhombusGridShape:
		if stage.OnAfterGrowthCurveRhombusGridShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurveRhombusGridShape)
			stage.OnAfterGrowthCurveRhombusGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurveRhombusShape:
		if stage.OnAfterGrowthCurveRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurveRhombusShape)
			stage.OnAfterGrowthCurveRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthVectorShape)
			stage.OnAfterGrowthVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *InitialRhombusGridShape:
		if stage.OnAfterInitialRhombusGridShapeDeleteCallback != nil {
			staged := any(staged).(*InitialRhombusGridShape)
			stage.OnAfterInitialRhombusGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *InitialRhombusShape:
		if stage.OnAfterInitialRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*InitialRhombusShape)
			stage.OnAfterInitialRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NextCircleShape:
		if stage.OnAfterNextCircleShapeDeleteCallback != nil {
			staged := any(staged).(*NextCircleShape)
			stage.OnAfterNextCircleShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerpendicularVector:
		if stage.OnAfterPerpendicularVectorDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVector)
			stage.OnAfterPerpendicularVectorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerpendicularVectorGrid:
		if stage.OnAfterPerpendicularVectorGridDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVectorGrid)
			stage.OnAfterPerpendicularVectorGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Plant:
		if stage.OnAfterPlantDeleteCallback != nil {
			staged := any(staged).(*Plant)
			stage.OnAfterPlantDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PlantCircumferenceShape:
		if stage.OnAfterPlantCircumferenceShapeDeleteCallback != nil {
			staged := any(staged).(*PlantCircumferenceShape)
			stage.OnAfterPlantCircumferenceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramDeleteCallback != nil {
			staged := any(staged).(*PlantDiagram)
			stage.OnAfterPlantDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RhombusShape:
		if stage.OnAfterRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*RhombusShape)
			stage.OnAfterRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RotatedRhombusGridShape:
		if stage.OnAfterRotatedRhombusGridShapeDeleteCallback != nil {
			staged := any(staged).(*RotatedRhombusGridShape)
			stage.OnAfterRotatedRhombusGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RotatedRhombusShape:
		if stage.OnAfterRotatedRhombusShapeDeleteCallback != nil {
			staged := any(staged).(*RotatedRhombusShape)
			stage.OnAfterRotatedRhombusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurveBezierShape:
		if stage.OnAfterStackGrowthCurveBezierShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurveBezierShape)
			stage.OnAfterStackGrowthCurveBezierShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfGrowthCurve:
		if stage.OnAfterStackOfGrowthCurveDeleteCallback != nil {
			staged := any(staged).(*StackOfGrowthCurve)
			stage.OnAfterStackOfGrowthCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AxesShape:
		if stage.OnAfterAxesShapeReadCallback != nil {
			stage.OnAfterAxesShapeReadCallback.OnAfterRead(stage, target)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeReadCallback != nil {
			stage.OnAfterCircleGridShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ExplanationTextShape:
		if stage.OnAfterExplanationTextShapeReadCallback != nil {
			stage.OnAfterExplanationTextShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeReadCallback != nil {
			stage.OnAfterGridPathShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthCurveBezierShape:
		if stage.OnAfterGrowthCurveBezierShapeReadCallback != nil {
			stage.OnAfterGrowthCurveBezierShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthCurveBezierShapeGrid:
		if stage.OnAfterGrowthCurveBezierShapeGridReadCallback != nil {
			stage.OnAfterGrowthCurveBezierShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthCurveRhombusGridShape:
		if stage.OnAfterGrowthCurveRhombusGridShapeReadCallback != nil {
			stage.OnAfterGrowthCurveRhombusGridShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthCurveRhombusShape:
		if stage.OnAfterGrowthCurveRhombusShapeReadCallback != nil {
			stage.OnAfterGrowthCurveRhombusShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeReadCallback != nil {
			stage.OnAfterGrowthVectorShapeReadCallback.OnAfterRead(stage, target)
		}
	case *InitialRhombusGridShape:
		if stage.OnAfterInitialRhombusGridShapeReadCallback != nil {
			stage.OnAfterInitialRhombusGridShapeReadCallback.OnAfterRead(stage, target)
		}
	case *InitialRhombusShape:
		if stage.OnAfterInitialRhombusShapeReadCallback != nil {
			stage.OnAfterInitialRhombusShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *NextCircleShape:
		if stage.OnAfterNextCircleShapeReadCallback != nil {
			stage.OnAfterNextCircleShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PerpendicularVector:
		if stage.OnAfterPerpendicularVectorReadCallback != nil {
			stage.OnAfterPerpendicularVectorReadCallback.OnAfterRead(stage, target)
		}
	case *PerpendicularVectorGrid:
		if stage.OnAfterPerpendicularVectorGridReadCallback != nil {
			stage.OnAfterPerpendicularVectorGridReadCallback.OnAfterRead(stage, target)
		}
	case *Plant:
		if stage.OnAfterPlantReadCallback != nil {
			stage.OnAfterPlantReadCallback.OnAfterRead(stage, target)
		}
	case *PlantCircumferenceShape:
		if stage.OnAfterPlantCircumferenceShapeReadCallback != nil {
			stage.OnAfterPlantCircumferenceShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramReadCallback != nil {
			stage.OnAfterPlantDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *RhombusShape:
		if stage.OnAfterRhombusShapeReadCallback != nil {
			stage.OnAfterRhombusShapeReadCallback.OnAfterRead(stage, target)
		}
	case *RotatedRhombusGridShape:
		if stage.OnAfterRotatedRhombusGridShapeReadCallback != nil {
			stage.OnAfterRotatedRhombusGridShapeReadCallback.OnAfterRead(stage, target)
		}
	case *RotatedRhombusShape:
		if stage.OnAfterRotatedRhombusShapeReadCallback != nil {
			stage.OnAfterRotatedRhombusShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackGrowthCurveBezierShape:
		if stage.OnAfterStackGrowthCurveBezierShapeReadCallback != nil {
			stage.OnAfterStackGrowthCurveBezierShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackOfGrowthCurve:
		if stage.OnAfterStackOfGrowthCurveReadCallback != nil {
			stage.OnAfterStackOfGrowthCurveReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[AxesShape])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[CircleGridShape])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GridPathShape])
	case *GrowthCurveBezierShape:
		stage.OnAfterGrowthCurveBezierShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurveBezierShape])
	case *GrowthCurveBezierShapeGrid:
		stage.OnAfterGrowthCurveBezierShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurveBezierShapeGrid])
	case *GrowthCurveRhombusGridShape:
		stage.OnAfterGrowthCurveRhombusGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurveRhombusGridShape])
	case *GrowthCurveRhombusShape:
		stage.OnAfterGrowthCurveRhombusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurveRhombusShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthVectorShape])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorUpdateCallback = any(callback).(OnAfterUpdateInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridUpdateCallback = any(callback).(OnAfterUpdateInterface[PerpendicularVectorGrid])
	case *Plant:
		stage.OnAfterPlantUpdateCallback = any(callback).(OnAfterUpdateInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[PlantDiagram])
	case *RhombusShape:
		stage.OnAfterRhombusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RotatedRhombusShape])
	case *StackGrowthCurveBezierShape:
		stage.OnAfterStackGrowthCurveBezierShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackGrowthCurveBezierShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveUpdateCallback = any(callback).(OnAfterUpdateInterface[StackOfGrowthCurve])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeCreateCallback = any(callback).(OnAfterCreateInterface[AxesShape])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[CircleGridShape])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeCreateCallback = any(callback).(OnAfterCreateInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeCreateCallback = any(callback).(OnAfterCreateInterface[GridPathShape])
	case *GrowthCurveBezierShape:
		stage.OnAfterGrowthCurveBezierShapeCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurveBezierShape])
	case *GrowthCurveBezierShapeGrid:
		stage.OnAfterGrowthCurveBezierShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurveBezierShapeGrid])
	case *GrowthCurveRhombusGridShape:
		stage.OnAfterGrowthCurveRhombusGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurveRhombusGridShape])
	case *GrowthCurveRhombusShape:
		stage.OnAfterGrowthCurveRhombusShapeCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurveRhombusShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[GrowthVectorShape])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeCreateCallback = any(callback).(OnAfterCreateInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeCreateCallback = any(callback).(OnAfterCreateInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorCreateCallback = any(callback).(OnAfterCreateInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridCreateCallback = any(callback).(OnAfterCreateInterface[PerpendicularVectorGrid])
	case *Plant:
		stage.OnAfterPlantCreateCallback = any(callback).(OnAfterCreateInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeCreateCallback = any(callback).(OnAfterCreateInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramCreateCallback = any(callback).(OnAfterCreateInterface[PlantDiagram])
	case *RhombusShape:
		stage.OnAfterRhombusShapeCreateCallback = any(callback).(OnAfterCreateInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeCreateCallback = any(callback).(OnAfterCreateInterface[RotatedRhombusShape])
	case *StackGrowthCurveBezierShape:
		stage.OnAfterStackGrowthCurveBezierShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackGrowthCurveBezierShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveCreateCallback = any(callback).(OnAfterCreateInterface[StackOfGrowthCurve])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AxesShape])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[CircleGridShape])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GridPathShape])
	case *GrowthCurveBezierShape:
		stage.OnAfterGrowthCurveBezierShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurveBezierShape])
	case *GrowthCurveBezierShapeGrid:
		stage.OnAfterGrowthCurveBezierShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurveBezierShapeGrid])
	case *GrowthCurveRhombusGridShape:
		stage.OnAfterGrowthCurveRhombusGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurveRhombusGridShape])
	case *GrowthCurveRhombusShape:
		stage.OnAfterGrowthCurveRhombusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurveRhombusShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthVectorShape])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorDeleteCallback = any(callback).(OnAfterDeleteInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridDeleteCallback = any(callback).(OnAfterDeleteInterface[PerpendicularVectorGrid])
	case *Plant:
		stage.OnAfterPlantDeleteCallback = any(callback).(OnAfterDeleteInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[PlantDiagram])
	case *RhombusShape:
		stage.OnAfterRhombusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RotatedRhombusShape])
	case *StackGrowthCurveBezierShape:
		stage.OnAfterStackGrowthCurveBezierShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackGrowthCurveBezierShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveDeleteCallback = any(callback).(OnAfterDeleteInterface[StackOfGrowthCurve])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeReadCallback = any(callback).(OnAfterReadInterface[AxesShape])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeReadCallback = any(callback).(OnAfterReadInterface[CircleGridShape])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeReadCallback = any(callback).(OnAfterReadInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeReadCallback = any(callback).(OnAfterReadInterface[GridPathShape])
	case *GrowthCurveBezierShape:
		stage.OnAfterGrowthCurveBezierShapeReadCallback = any(callback).(OnAfterReadInterface[GrowthCurveBezierShape])
	case *GrowthCurveBezierShapeGrid:
		stage.OnAfterGrowthCurveBezierShapeGridReadCallback = any(callback).(OnAfterReadInterface[GrowthCurveBezierShapeGrid])
	case *GrowthCurveRhombusGridShape:
		stage.OnAfterGrowthCurveRhombusGridShapeReadCallback = any(callback).(OnAfterReadInterface[GrowthCurveRhombusGridShape])
	case *GrowthCurveRhombusShape:
		stage.OnAfterGrowthCurveRhombusShapeReadCallback = any(callback).(OnAfterReadInterface[GrowthCurveRhombusShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeReadCallback = any(callback).(OnAfterReadInterface[GrowthVectorShape])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeReadCallback = any(callback).(OnAfterReadInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeReadCallback = any(callback).(OnAfterReadInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeReadCallback = any(callback).(OnAfterReadInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorReadCallback = any(callback).(OnAfterReadInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridReadCallback = any(callback).(OnAfterReadInterface[PerpendicularVectorGrid])
	case *Plant:
		stage.OnAfterPlantReadCallback = any(callback).(OnAfterReadInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeReadCallback = any(callback).(OnAfterReadInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramReadCallback = any(callback).(OnAfterReadInterface[PlantDiagram])
	case *RhombusShape:
		stage.OnAfterRhombusShapeReadCallback = any(callback).(OnAfterReadInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeReadCallback = any(callback).(OnAfterReadInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeReadCallback = any(callback).(OnAfterReadInterface[RotatedRhombusShape])
	case *StackGrowthCurveBezierShape:
		stage.OnAfterStackGrowthCurveBezierShapeReadCallback = any(callback).(OnAfterReadInterface[StackGrowthCurveBezierShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveReadCallback = any(callback).(OnAfterReadInterface[StackOfGrowthCurve])
	}
}
