// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ArcNormalVectorShape:
		if stage.OnAfterArcNormalVectorShapeCreateCallback != nil {
			stage.OnAfterArcNormalVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ArcNormalVectorShapeGrid:
		if stage.OnAfterArcNormalVectorShapeGridCreateCallback != nil {
			stage.OnAfterArcNormalVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *AxesShape:
		if stage.OnAfterAxesShapeCreateCallback != nil {
			stage.OnAfterAxesShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *BaseVectorShape:
		if stage.OnAfterBaseVectorShapeCreateCallback != nil {
			stage.OnAfterBaseVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *BaseVectorShapeGrid:
		if stage.OnAfterBaseVectorShapeGridCreateCallback != nil {
			stage.OnAfterBaseVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeCreateCallback != nil {
			stage.OnAfterCircleGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndArcShape:
		if stage.OnAfterEndArcShapeCreateCallback != nil {
			stage.OnAfterEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndArcShapeGrid:
		if stage.OnAfterEndArcShapeGridCreateCallback != nil {
			stage.OnAfterEndArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndHalfwayArcShape:
		if stage.OnAfterEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *EndHalfwayArcShapeGrid:
		if stage.OnAfterEndHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *ExplanationTextShape:
		if stage.OnAfterExplanationTextShapeCreateCallback != nil {
			stage.OnAfterExplanationTextShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeCreateCallback != nil {
			stage.OnAfterGridPathShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurve2D:
		if stage.OnAfterGrowthCurve2DCreateCallback != nil {
			stage.OnAfterGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
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
	case *HalfwayArcShape:
		if stage.OnAfterHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *HalfwayArcShapeGrid:
		if stage.OnAfterHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
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
	case *MidArcVectorShape:
		if stage.OnAfterMidArcVectorShapeCreateCallback != nil {
			stage.OnAfterMidArcVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *MidArcVectorShapeGrid:
		if stage.OnAfterMidArcVectorShapeGridCreateCallback != nil {
			stage.OnAfterMidArcVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
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
	case *PerpendicularVectorGridHalfway:
		if stage.OnAfterPerpendicularVectorGridHalfwayCreateCallback != nil {
			stage.OnAfterPerpendicularVectorGridHalfwayCreateCallback.OnAfterCreate(stage, target)
		}
	case *PerpendicularVectorHalfway:
		if stage.OnAfterPerpendicularVectorHalfwayCreateCallback != nil {
			stage.OnAfterPerpendicularVectorHalfwayCreateCallback.OnAfterCreate(stage, target)
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
	case *Rendered3DShape:
		if stage.OnAfterRendered3DShapeCreateCallback != nil {
			stage.OnAfterRendered3DShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *ShiftedBottomTopStartArcShape:
		if stage.OnAfterShiftedBottomTopStartArcShapeCreateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedBottomTopStartArcShapeGrid:
		if stage.OnAfterShiftedBottomTopStartArcShapeGridCreateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackNormalVector:
		if stage.OnAfterShiftedLeftStackNormalVectorCreateCallback != nil {
			stage.OnAfterShiftedLeftStackNormalVectorCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackOfGrowthCurve:
		if stage.OnAfterShiftedLeftStackOfGrowthCurveCreateCallback != nil {
			stage.OnAfterShiftedLeftStackOfGrowthCurveCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedLeftStackOfNormalVector:
		if stage.OnAfterShiftedLeftStackOfNormalVectorCreateCallback != nil {
			stage.OnAfterShiftedLeftStackOfNormalVectorCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurveEndArcShape:
		if stage.OnAfterStackGrowthCurveEndArcShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurveEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurveStartArcShape:
		if stage.OnAfterStackGrowthCurveStartArcShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurveStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfGrowthCurve:
		if stage.OnAfterStackOfGrowthCurveCreateCallback != nil {
			stage.OnAfterStackOfGrowthCurveCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartArcShape:
		if stage.OnAfterStartArcShapeCreateCallback != nil {
			stage.OnAfterStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartArcShapeGrid:
		if stage.OnAfterStartArcShapeGridCreateCallback != nil {
			stage.OnAfterStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndArcShape:
		if stage.OnAfterTopEndArcShapeCreateCallback != nil {
			stage.OnAfterTopEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndArcShapeGrid:
		if stage.OnAfterTopEndArcShapeGridCreateCallback != nil {
			stage.OnAfterTopEndArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopGrowthCurve2D:
		if stage.OnAfterTopGrowthCurve2DCreateCallback != nil {
			stage.OnAfterTopGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopMidArcVectorShape:
		if stage.OnAfterTopMidArcVectorShapeCreateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopMidArcVectorShapeGrid:
		if stage.OnAfterTopMidArcVectorShapeGridCreateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackGrowthCurveEndArcShape:
		if stage.OnAfterTopStackGrowthCurveEndArcShapeCreateCallback != nil {
			stage.OnAfterTopStackGrowthCurveEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackGrowthCurveStartArcShape:
		if stage.OnAfterTopStackGrowthCurveStartArcShapeCreateCallback != nil {
			stage.OnAfterTopStackGrowthCurveStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfGrowthCurve:
		if stage.OnAfterTopStackOfGrowthCurveCreateCallback != nil {
			stage.OnAfterTopStackOfGrowthCurveCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartArcShape:
		if stage.OnAfterTopStartArcShapeCreateCallback != nil {
			stage.OnAfterTopStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartArcShapeGrid:
		if stage.OnAfterTopStartArcShapeGridCreateCallback != nil {
			stage.OnAfterTopStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
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
	case *ArcNormalVectorShape:
		newTarget := any(new).(*ArcNormalVectorShape)
		if stage.OnAfterArcNormalVectorShapeUpdateCallback != nil {
			stage.OnAfterArcNormalVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ArcNormalVectorShapeGrid:
		newTarget := any(new).(*ArcNormalVectorShapeGrid)
		if stage.OnAfterArcNormalVectorShapeGridUpdateCallback != nil {
			stage.OnAfterArcNormalVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AxesShape:
		newTarget := any(new).(*AxesShape)
		if stage.OnAfterAxesShapeUpdateCallback != nil {
			stage.OnAfterAxesShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BaseVectorShape:
		newTarget := any(new).(*BaseVectorShape)
		if stage.OnAfterBaseVectorShapeUpdateCallback != nil {
			stage.OnAfterBaseVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *BaseVectorShapeGrid:
		newTarget := any(new).(*BaseVectorShapeGrid)
		if stage.OnAfterBaseVectorShapeGridUpdateCallback != nil {
			stage.OnAfterBaseVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CircleGridShape:
		newTarget := any(new).(*CircleGridShape)
		if stage.OnAfterCircleGridShapeUpdateCallback != nil {
			stage.OnAfterCircleGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndArcShape:
		newTarget := any(new).(*EndArcShape)
		if stage.OnAfterEndArcShapeUpdateCallback != nil {
			stage.OnAfterEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndArcShapeGrid:
		newTarget := any(new).(*EndArcShapeGrid)
		if stage.OnAfterEndArcShapeGridUpdateCallback != nil {
			stage.OnAfterEndArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndHalfwayArcShape:
		newTarget := any(new).(*EndHalfwayArcShape)
		if stage.OnAfterEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EndHalfwayArcShapeGrid:
		newTarget := any(new).(*EndHalfwayArcShapeGrid)
		if stage.OnAfterEndHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterEndHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *GrowthCurve2D:
		newTarget := any(new).(*GrowthCurve2D)
		if stage.OnAfterGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *HalfwayArcShape:
		newTarget := any(new).(*HalfwayArcShape)
		if stage.OnAfterHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *HalfwayArcShapeGrid:
		newTarget := any(new).(*HalfwayArcShapeGrid)
		if stage.OnAfterHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *MidArcVectorShape:
		newTarget := any(new).(*MidArcVectorShape)
		if stage.OnAfterMidArcVectorShapeUpdateCallback != nil {
			stage.OnAfterMidArcVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MidArcVectorShapeGrid:
		newTarget := any(new).(*MidArcVectorShapeGrid)
		if stage.OnAfterMidArcVectorShapeGridUpdateCallback != nil {
			stage.OnAfterMidArcVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *PerpendicularVectorGridHalfway:
		newTarget := any(new).(*PerpendicularVectorGridHalfway)
		if stage.OnAfterPerpendicularVectorGridHalfwayUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorGridHalfwayUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PerpendicularVectorHalfway:
		newTarget := any(new).(*PerpendicularVectorHalfway)
		if stage.OnAfterPerpendicularVectorHalfwayUpdateCallback != nil {
			stage.OnAfterPerpendicularVectorHalfwayUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Rendered3DShape:
		newTarget := any(new).(*Rendered3DShape)
		if stage.OnAfterRendered3DShapeUpdateCallback != nil {
			stage.OnAfterRendered3DShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *ShiftedBottomTopStartArcShape:
		newTarget := any(new).(*ShiftedBottomTopStartArcShape)
		if stage.OnAfterShiftedBottomTopStartArcShapeUpdateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedBottomTopStartArcShapeGrid:
		newTarget := any(new).(*ShiftedBottomTopStartArcShapeGrid)
		if stage.OnAfterShiftedBottomTopStartArcShapeGridUpdateCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		newTarget := any(new).(*ShiftedLeftStackGrowthCurveEndArcShape)
		if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		newTarget := any(new).(*ShiftedLeftStackGrowthCurveStartArcShape)
		if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackNormalVector:
		newTarget := any(new).(*ShiftedLeftStackNormalVector)
		if stage.OnAfterShiftedLeftStackNormalVectorUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackNormalVectorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackOfGrowthCurve:
		newTarget := any(new).(*ShiftedLeftStackOfGrowthCurve)
		if stage.OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedLeftStackOfNormalVector:
		newTarget := any(new).(*ShiftedLeftStackOfNormalVector)
		if stage.OnAfterShiftedLeftStackOfNormalVectorUpdateCallback != nil {
			stage.OnAfterShiftedLeftStackOfNormalVectorUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurveEndArcShape:
		newTarget := any(new).(*StackGrowthCurveEndArcShape)
		if stage.OnAfterStackGrowthCurveEndArcShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurveEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurveStartArcShape:
		newTarget := any(new).(*StackGrowthCurveStartArcShape)
		if stage.OnAfterStackGrowthCurveStartArcShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurveStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfGrowthCurve:
		newTarget := any(new).(*StackOfGrowthCurve)
		if stage.OnAfterStackOfGrowthCurveUpdateCallback != nil {
			stage.OnAfterStackOfGrowthCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StartArcShape:
		newTarget := any(new).(*StartArcShape)
		if stage.OnAfterStartArcShapeUpdateCallback != nil {
			stage.OnAfterStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StartArcShapeGrid:
		newTarget := any(new).(*StartArcShapeGrid)
		if stage.OnAfterStartArcShapeGridUpdateCallback != nil {
			stage.OnAfterStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopEndArcShape:
		newTarget := any(new).(*TopEndArcShape)
		if stage.OnAfterTopEndArcShapeUpdateCallback != nil {
			stage.OnAfterTopEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopEndArcShapeGrid:
		newTarget := any(new).(*TopEndArcShapeGrid)
		if stage.OnAfterTopEndArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopEndArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopGrowthCurve2D:
		newTarget := any(new).(*TopGrowthCurve2D)
		if stage.OnAfterTopGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterTopGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopMidArcVectorShape:
		newTarget := any(new).(*TopMidArcVectorShape)
		if stage.OnAfterTopMidArcVectorShapeUpdateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopMidArcVectorShapeGrid:
		newTarget := any(new).(*TopMidArcVectorShapeGrid)
		if stage.OnAfterTopMidArcVectorShapeGridUpdateCallback != nil {
			stage.OnAfterTopMidArcVectorShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackGrowthCurveEndArcShape:
		newTarget := any(new).(*TopStackGrowthCurveEndArcShape)
		if stage.OnAfterTopStackGrowthCurveEndArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackGrowthCurveEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackGrowthCurveStartArcShape:
		newTarget := any(new).(*TopStackGrowthCurveStartArcShape)
		if stage.OnAfterTopStackGrowthCurveStartArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackGrowthCurveStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfGrowthCurve:
		newTarget := any(new).(*TopStackOfGrowthCurve)
		if stage.OnAfterTopStackOfGrowthCurveUpdateCallback != nil {
			stage.OnAfterTopStackOfGrowthCurveUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStartArcShape:
		newTarget := any(new).(*TopStartArcShape)
		if stage.OnAfterTopStartArcShapeUpdateCallback != nil {
			stage.OnAfterTopStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStartArcShapeGrid:
		newTarget := any(new).(*TopStartArcShapeGrid)
		if stage.OnAfterTopStartArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *ArcNormalVectorShape:
		if stage.OnAfterArcNormalVectorShapeDeleteCallback != nil {
			staged := any(staged).(*ArcNormalVectorShape)
			stage.OnAfterArcNormalVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ArcNormalVectorShapeGrid:
		if stage.OnAfterArcNormalVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*ArcNormalVectorShapeGrid)
			stage.OnAfterArcNormalVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AxesShape:
		if stage.OnAfterAxesShapeDeleteCallback != nil {
			staged := any(staged).(*AxesShape)
			stage.OnAfterAxesShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BaseVectorShape:
		if stage.OnAfterBaseVectorShapeDeleteCallback != nil {
			staged := any(staged).(*BaseVectorShape)
			stage.OnAfterBaseVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *BaseVectorShapeGrid:
		if stage.OnAfterBaseVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*BaseVectorShapeGrid)
			stage.OnAfterBaseVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeDeleteCallback != nil {
			staged := any(staged).(*CircleGridShape)
			stage.OnAfterCircleGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndArcShape:
		if stage.OnAfterEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*EndArcShape)
			stage.OnAfterEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndArcShapeGrid:
		if stage.OnAfterEndArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*EndArcShapeGrid)
			stage.OnAfterEndArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndHalfwayArcShape:
		if stage.OnAfterEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*EndHalfwayArcShape)
			stage.OnAfterEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EndHalfwayArcShapeGrid:
		if stage.OnAfterEndHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*EndHalfwayArcShapeGrid)
			stage.OnAfterEndHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *GrowthCurve2D:
		if stage.OnAfterGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2D)
			stage.OnAfterGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *HalfwayArcShape:
		if stage.OnAfterHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*HalfwayArcShape)
			stage.OnAfterHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *HalfwayArcShapeGrid:
		if stage.OnAfterHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*HalfwayArcShapeGrid)
			stage.OnAfterHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *MidArcVectorShape:
		if stage.OnAfterMidArcVectorShapeDeleteCallback != nil {
			staged := any(staged).(*MidArcVectorShape)
			stage.OnAfterMidArcVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MidArcVectorShapeGrid:
		if stage.OnAfterMidArcVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*MidArcVectorShapeGrid)
			stage.OnAfterMidArcVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *PerpendicularVectorGridHalfway:
		if stage.OnAfterPerpendicularVectorGridHalfwayDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVectorGridHalfway)
			stage.OnAfterPerpendicularVectorGridHalfwayDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PerpendicularVectorHalfway:
		if stage.OnAfterPerpendicularVectorHalfwayDeleteCallback != nil {
			staged := any(staged).(*PerpendicularVectorHalfway)
			stage.OnAfterPerpendicularVectorHalfwayDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *Rendered3DShape:
		if stage.OnAfterRendered3DShapeDeleteCallback != nil {
			staged := any(staged).(*Rendered3DShape)
			stage.OnAfterRendered3DShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *ShiftedBottomTopStartArcShape:
		if stage.OnAfterShiftedBottomTopStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedBottomTopStartArcShape)
			stage.OnAfterShiftedBottomTopStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedBottomTopStartArcShapeGrid:
		if stage.OnAfterShiftedBottomTopStartArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*ShiftedBottomTopStartArcShapeGrid)
			stage.OnAfterShiftedBottomTopStartArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackGrowthCurveEndArcShape)
			stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackGrowthCurveStartArcShape)
			stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackNormalVector:
		if stage.OnAfterShiftedLeftStackNormalVectorDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackNormalVector)
			stage.OnAfterShiftedLeftStackNormalVectorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackOfGrowthCurve:
		if stage.OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackOfGrowthCurve)
			stage.OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedLeftStackOfNormalVector:
		if stage.OnAfterShiftedLeftStackOfNormalVectorDeleteCallback != nil {
			staged := any(staged).(*ShiftedLeftStackOfNormalVector)
			stage.OnAfterShiftedLeftStackOfNormalVectorDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurveEndArcShape:
		if stage.OnAfterStackGrowthCurveEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurveEndArcShape)
			stage.OnAfterStackGrowthCurveEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurveStartArcShape:
		if stage.OnAfterStackGrowthCurveStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurveStartArcShape)
			stage.OnAfterStackGrowthCurveStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfGrowthCurve:
		if stage.OnAfterStackOfGrowthCurveDeleteCallback != nil {
			staged := any(staged).(*StackOfGrowthCurve)
			stage.OnAfterStackOfGrowthCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StartArcShape:
		if stage.OnAfterStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*StartArcShape)
			stage.OnAfterStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StartArcShapeGrid:
		if stage.OnAfterStartArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*StartArcShapeGrid)
			stage.OnAfterStartArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopEndArcShape:
		if stage.OnAfterTopEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopEndArcShape)
			stage.OnAfterTopEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopEndArcShapeGrid:
		if stage.OnAfterTopEndArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopEndArcShapeGrid)
			stage.OnAfterTopEndArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopGrowthCurve2D:
		if stage.OnAfterTopGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*TopGrowthCurve2D)
			stage.OnAfterTopGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopMidArcVectorShape:
		if stage.OnAfterTopMidArcVectorShapeDeleteCallback != nil {
			staged := any(staged).(*TopMidArcVectorShape)
			stage.OnAfterTopMidArcVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopMidArcVectorShapeGrid:
		if stage.OnAfterTopMidArcVectorShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopMidArcVectorShapeGrid)
			stage.OnAfterTopMidArcVectorShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackGrowthCurveEndArcShape:
		if stage.OnAfterTopStackGrowthCurveEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackGrowthCurveEndArcShape)
			stage.OnAfterTopStackGrowthCurveEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackGrowthCurveStartArcShape:
		if stage.OnAfterTopStackGrowthCurveStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackGrowthCurveStartArcShape)
			stage.OnAfterTopStackGrowthCurveStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfGrowthCurve:
		if stage.OnAfterTopStackOfGrowthCurveDeleteCallback != nil {
			staged := any(staged).(*TopStackOfGrowthCurve)
			stage.OnAfterTopStackOfGrowthCurveDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStartArcShape:
		if stage.OnAfterTopStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStartArcShape)
			stage.OnAfterTopStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStartArcShapeGrid:
		if stage.OnAfterTopStartArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopStartArcShapeGrid)
			stage.OnAfterTopStartArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ArcNormalVectorShape:
		if stage.OnAfterArcNormalVectorShapeReadCallback != nil {
			stage.OnAfterArcNormalVectorShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ArcNormalVectorShapeGrid:
		if stage.OnAfterArcNormalVectorShapeGridReadCallback != nil {
			stage.OnAfterArcNormalVectorShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *AxesShape:
		if stage.OnAfterAxesShapeReadCallback != nil {
			stage.OnAfterAxesShapeReadCallback.OnAfterRead(stage, target)
		}
	case *BaseVectorShape:
		if stage.OnAfterBaseVectorShapeReadCallback != nil {
			stage.OnAfterBaseVectorShapeReadCallback.OnAfterRead(stage, target)
		}
	case *BaseVectorShapeGrid:
		if stage.OnAfterBaseVectorShapeGridReadCallback != nil {
			stage.OnAfterBaseVectorShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeReadCallback != nil {
			stage.OnAfterCircleGridShapeReadCallback.OnAfterRead(stage, target)
		}
	case *EndArcShape:
		if stage.OnAfterEndArcShapeReadCallback != nil {
			stage.OnAfterEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *EndArcShapeGrid:
		if stage.OnAfterEndArcShapeGridReadCallback != nil {
			stage.OnAfterEndArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *EndHalfwayArcShape:
		if stage.OnAfterEndHalfwayArcShapeReadCallback != nil {
			stage.OnAfterEndHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *EndHalfwayArcShapeGrid:
		if stage.OnAfterEndHalfwayArcShapeGridReadCallback != nil {
			stage.OnAfterEndHalfwayArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *ExplanationTextShape:
		if stage.OnAfterExplanationTextShapeReadCallback != nil {
			stage.OnAfterExplanationTextShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeReadCallback != nil {
			stage.OnAfterGridPathShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthCurve2D:
		if stage.OnAfterGrowthCurve2DReadCallback != nil {
			stage.OnAfterGrowthCurve2DReadCallback.OnAfterRead(stage, target)
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
	case *HalfwayArcShape:
		if stage.OnAfterHalfwayArcShapeReadCallback != nil {
			stage.OnAfterHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *HalfwayArcShapeGrid:
		if stage.OnAfterHalfwayArcShapeGridReadCallback != nil {
			stage.OnAfterHalfwayArcShapeGridReadCallback.OnAfterRead(stage, target)
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
	case *MidArcVectorShape:
		if stage.OnAfterMidArcVectorShapeReadCallback != nil {
			stage.OnAfterMidArcVectorShapeReadCallback.OnAfterRead(stage, target)
		}
	case *MidArcVectorShapeGrid:
		if stage.OnAfterMidArcVectorShapeGridReadCallback != nil {
			stage.OnAfterMidArcVectorShapeGridReadCallback.OnAfterRead(stage, target)
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
	case *PerpendicularVectorGridHalfway:
		if stage.OnAfterPerpendicularVectorGridHalfwayReadCallback != nil {
			stage.OnAfterPerpendicularVectorGridHalfwayReadCallback.OnAfterRead(stage, target)
		}
	case *PerpendicularVectorHalfway:
		if stage.OnAfterPerpendicularVectorHalfwayReadCallback != nil {
			stage.OnAfterPerpendicularVectorHalfwayReadCallback.OnAfterRead(stage, target)
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
	case *Rendered3DShape:
		if stage.OnAfterRendered3DShapeReadCallback != nil {
			stage.OnAfterRendered3DShapeReadCallback.OnAfterRead(stage, target)
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
	case *ShiftedBottomTopStartArcShape:
		if stage.OnAfterShiftedBottomTopStartArcShapeReadCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedBottomTopStartArcShapeGrid:
		if stage.OnAfterShiftedBottomTopStartArcShapeGridReadCallback != nil {
			stage.OnAfterShiftedBottomTopStartArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeReadCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeReadCallback != nil {
			stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedLeftStackNormalVector:
		if stage.OnAfterShiftedLeftStackNormalVectorReadCallback != nil {
			stage.OnAfterShiftedLeftStackNormalVectorReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedLeftStackOfGrowthCurve:
		if stage.OnAfterShiftedLeftStackOfGrowthCurveReadCallback != nil {
			stage.OnAfterShiftedLeftStackOfGrowthCurveReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedLeftStackOfNormalVector:
		if stage.OnAfterShiftedLeftStackOfNormalVectorReadCallback != nil {
			stage.OnAfterShiftedLeftStackOfNormalVectorReadCallback.OnAfterRead(stage, target)
		}
	case *StackGrowthCurveEndArcShape:
		if stage.OnAfterStackGrowthCurveEndArcShapeReadCallback != nil {
			stage.OnAfterStackGrowthCurveEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackGrowthCurveStartArcShape:
		if stage.OnAfterStackGrowthCurveStartArcShapeReadCallback != nil {
			stage.OnAfterStackGrowthCurveStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackOfGrowthCurve:
		if stage.OnAfterStackOfGrowthCurveReadCallback != nil {
			stage.OnAfterStackOfGrowthCurveReadCallback.OnAfterRead(stage, target)
		}
	case *StartArcShape:
		if stage.OnAfterStartArcShapeReadCallback != nil {
			stage.OnAfterStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StartArcShapeGrid:
		if stage.OnAfterStartArcShapeGridReadCallback != nil {
			stage.OnAfterStartArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *TopEndArcShape:
		if stage.OnAfterTopEndArcShapeReadCallback != nil {
			stage.OnAfterTopEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopEndArcShapeGrid:
		if stage.OnAfterTopEndArcShapeGridReadCallback != nil {
			stage.OnAfterTopEndArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *TopGrowthCurve2D:
		if stage.OnAfterTopGrowthCurve2DReadCallback != nil {
			stage.OnAfterTopGrowthCurve2DReadCallback.OnAfterRead(stage, target)
		}
	case *TopMidArcVectorShape:
		if stage.OnAfterTopMidArcVectorShapeReadCallback != nil {
			stage.OnAfterTopMidArcVectorShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopMidArcVectorShapeGrid:
		if stage.OnAfterTopMidArcVectorShapeGridReadCallback != nil {
			stage.OnAfterTopMidArcVectorShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackGrowthCurveEndArcShape:
		if stage.OnAfterTopStackGrowthCurveEndArcShapeReadCallback != nil {
			stage.OnAfterTopStackGrowthCurveEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackGrowthCurveStartArcShape:
		if stage.OnAfterTopStackGrowthCurveStartArcShapeReadCallback != nil {
			stage.OnAfterTopStackGrowthCurveStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackOfGrowthCurve:
		if stage.OnAfterTopStackOfGrowthCurveReadCallback != nil {
			stage.OnAfterTopStackOfGrowthCurveReadCallback.OnAfterRead(stage, target)
		}
	case *TopStartArcShape:
		if stage.OnAfterTopStartArcShapeReadCallback != nil {
			stage.OnAfterTopStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStartArcShapeGrid:
		if stage.OnAfterTopStartArcShapeGridReadCallback != nil {
			stage.OnAfterTopStartArcShapeGridReadCallback.OnAfterRead(stage, target)
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
	case *ArcNormalVectorShape:
		stage.OnAfterArcNormalVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ArcNormalVectorShape])
	case *ArcNormalVectorShapeGrid:
		stage.OnAfterArcNormalVectorShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[ArcNormalVectorShapeGrid])
	case *AxesShape:
		stage.OnAfterAxesShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[AxesShape])
	case *BaseVectorShape:
		stage.OnAfterBaseVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[BaseVectorShape])
	case *BaseVectorShapeGrid:
		stage.OnAfterBaseVectorShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[BaseVectorShapeGrid])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[CircleGridShape])
	case *EndArcShape:
		stage.OnAfterEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[EndArcShape])
	case *EndArcShapeGrid:
		stage.OnAfterEndArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[EndArcShapeGrid])
	case *EndHalfwayArcShape:
		stage.OnAfterEndHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[EndHalfwayArcShape])
	case *EndHalfwayArcShapeGrid:
		stage.OnAfterEndHalfwayArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[EndHalfwayArcShapeGrid])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GridPathShape])
	case *GrowthCurve2D:
		stage.OnAfterGrowthCurve2DUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurve2D])
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
	case *HalfwayArcShape:
		stage.OnAfterHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[HalfwayArcShape])
	case *HalfwayArcShapeGrid:
		stage.OnAfterHalfwayArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[HalfwayArcShapeGrid])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[MidArcVectorShapeGrid])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorUpdateCallback = any(callback).(OnAfterUpdateInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridUpdateCallback = any(callback).(OnAfterUpdateInterface[PerpendicularVectorGrid])
	case *PerpendicularVectorGridHalfway:
		stage.OnAfterPerpendicularVectorGridHalfwayUpdateCallback = any(callback).(OnAfterUpdateInterface[PerpendicularVectorGridHalfway])
	case *PerpendicularVectorHalfway:
		stage.OnAfterPerpendicularVectorHalfwayUpdateCallback = any(callback).(OnAfterUpdateInterface[PerpendicularVectorHalfway])
	case *Plant:
		stage.OnAfterPlantUpdateCallback = any(callback).(OnAfterUpdateInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[PlantDiagram])
	case *Rendered3DShape:
		stage.OnAfterRendered3DShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[Rendered3DShape])
	case *RhombusShape:
		stage.OnAfterRhombusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RotatedRhombusShape])
	case *ShiftedBottomTopStartArcShape:
		stage.OnAfterShiftedBottomTopStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedBottomTopStartArcShape])
	case *ShiftedBottomTopStartArcShapeGrid:
		stage.OnAfterShiftedBottomTopStartArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedBottomTopStartArcShapeGrid])
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedLeftStackGrowthCurveEndArcShape])
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedLeftStackGrowthCurveStartArcShape])
	case *ShiftedLeftStackNormalVector:
		stage.OnAfterShiftedLeftStackNormalVectorUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedLeftStackNormalVector])
	case *ShiftedLeftStackOfGrowthCurve:
		stage.OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedLeftStackOfGrowthCurve])
	case *ShiftedLeftStackOfNormalVector:
		stage.OnAfterShiftedLeftStackOfNormalVectorUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedLeftStackOfNormalVector])
	case *StackGrowthCurveEndArcShape:
		stage.OnAfterStackGrowthCurveEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackGrowthCurveEndArcShape])
	case *StackGrowthCurveStartArcShape:
		stage.OnAfterStackGrowthCurveStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackGrowthCurveStartArcShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveUpdateCallback = any(callback).(OnAfterUpdateInterface[StackOfGrowthCurve])
	case *StartArcShape:
		stage.OnAfterStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[StartArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopEndArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DUpdateCallback = any(callback).(OnAfterUpdateInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurveEndArcShape:
		stage.OnAfterTopStackGrowthCurveEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackGrowthCurveEndArcShape])
	case *TopStackGrowthCurveStartArcShape:
		stage.OnAfterTopStackGrowthCurveStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackGrowthCurveStartArcShape])
	case *TopStackOfGrowthCurve:
		stage.OnAfterTopStackOfGrowthCurveUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackOfGrowthCurve])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStartArcShapeGrid])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ArcNormalVectorShape:
		stage.OnAfterArcNormalVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[ArcNormalVectorShape])
	case *ArcNormalVectorShapeGrid:
		stage.OnAfterArcNormalVectorShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[ArcNormalVectorShapeGrid])
	case *AxesShape:
		stage.OnAfterAxesShapeCreateCallback = any(callback).(OnAfterCreateInterface[AxesShape])
	case *BaseVectorShape:
		stage.OnAfterBaseVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[BaseVectorShape])
	case *BaseVectorShapeGrid:
		stage.OnAfterBaseVectorShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[BaseVectorShapeGrid])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[CircleGridShape])
	case *EndArcShape:
		stage.OnAfterEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[EndArcShape])
	case *EndArcShapeGrid:
		stage.OnAfterEndArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[EndArcShapeGrid])
	case *EndHalfwayArcShape:
		stage.OnAfterEndHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[EndHalfwayArcShape])
	case *EndHalfwayArcShapeGrid:
		stage.OnAfterEndHalfwayArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[EndHalfwayArcShapeGrid])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeCreateCallback = any(callback).(OnAfterCreateInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeCreateCallback = any(callback).(OnAfterCreateInterface[GridPathShape])
	case *GrowthCurve2D:
		stage.OnAfterGrowthCurve2DCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurve2D])
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
	case *HalfwayArcShape:
		stage.OnAfterHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[HalfwayArcShape])
	case *HalfwayArcShapeGrid:
		stage.OnAfterHalfwayArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[HalfwayArcShapeGrid])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeCreateCallback = any(callback).(OnAfterCreateInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[MidArcVectorShapeGrid])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeCreateCallback = any(callback).(OnAfterCreateInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorCreateCallback = any(callback).(OnAfterCreateInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridCreateCallback = any(callback).(OnAfterCreateInterface[PerpendicularVectorGrid])
	case *PerpendicularVectorGridHalfway:
		stage.OnAfterPerpendicularVectorGridHalfwayCreateCallback = any(callback).(OnAfterCreateInterface[PerpendicularVectorGridHalfway])
	case *PerpendicularVectorHalfway:
		stage.OnAfterPerpendicularVectorHalfwayCreateCallback = any(callback).(OnAfterCreateInterface[PerpendicularVectorHalfway])
	case *Plant:
		stage.OnAfterPlantCreateCallback = any(callback).(OnAfterCreateInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeCreateCallback = any(callback).(OnAfterCreateInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramCreateCallback = any(callback).(OnAfterCreateInterface[PlantDiagram])
	case *Rendered3DShape:
		stage.OnAfterRendered3DShapeCreateCallback = any(callback).(OnAfterCreateInterface[Rendered3DShape])
	case *RhombusShape:
		stage.OnAfterRhombusShapeCreateCallback = any(callback).(OnAfterCreateInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeCreateCallback = any(callback).(OnAfterCreateInterface[RotatedRhombusShape])
	case *ShiftedBottomTopStartArcShape:
		stage.OnAfterShiftedBottomTopStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedBottomTopStartArcShape])
	case *ShiftedBottomTopStartArcShapeGrid:
		stage.OnAfterShiftedBottomTopStartArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedBottomTopStartArcShapeGrid])
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedLeftStackGrowthCurveEndArcShape])
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedLeftStackGrowthCurveStartArcShape])
	case *ShiftedLeftStackNormalVector:
		stage.OnAfterShiftedLeftStackNormalVectorCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedLeftStackNormalVector])
	case *ShiftedLeftStackOfGrowthCurve:
		stage.OnAfterShiftedLeftStackOfGrowthCurveCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedLeftStackOfGrowthCurve])
	case *ShiftedLeftStackOfNormalVector:
		stage.OnAfterShiftedLeftStackOfNormalVectorCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedLeftStackOfNormalVector])
	case *StackGrowthCurveEndArcShape:
		stage.OnAfterStackGrowthCurveEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackGrowthCurveEndArcShape])
	case *StackGrowthCurveStartArcShape:
		stage.OnAfterStackGrowthCurveStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackGrowthCurveStartArcShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveCreateCallback = any(callback).(OnAfterCreateInterface[StackOfGrowthCurve])
	case *StartArcShape:
		stage.OnAfterStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[StartArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopEndArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DCreateCallback = any(callback).(OnAfterCreateInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurveEndArcShape:
		stage.OnAfterTopStackGrowthCurveEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStackGrowthCurveEndArcShape])
	case *TopStackGrowthCurveStartArcShape:
		stage.OnAfterTopStackGrowthCurveStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStackGrowthCurveStartArcShape])
	case *TopStackOfGrowthCurve:
		stage.OnAfterTopStackOfGrowthCurveCreateCallback = any(callback).(OnAfterCreateInterface[TopStackOfGrowthCurve])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopStartArcShapeGrid])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ArcNormalVectorShape:
		stage.OnAfterArcNormalVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ArcNormalVectorShape])
	case *ArcNormalVectorShapeGrid:
		stage.OnAfterArcNormalVectorShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[ArcNormalVectorShapeGrid])
	case *AxesShape:
		stage.OnAfterAxesShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AxesShape])
	case *BaseVectorShape:
		stage.OnAfterBaseVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[BaseVectorShape])
	case *BaseVectorShapeGrid:
		stage.OnAfterBaseVectorShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[BaseVectorShapeGrid])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[CircleGridShape])
	case *EndArcShape:
		stage.OnAfterEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[EndArcShape])
	case *EndArcShapeGrid:
		stage.OnAfterEndArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[EndArcShapeGrid])
	case *EndHalfwayArcShape:
		stage.OnAfterEndHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[EndHalfwayArcShape])
	case *EndHalfwayArcShapeGrid:
		stage.OnAfterEndHalfwayArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[EndHalfwayArcShapeGrid])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GridPathShape])
	case *GrowthCurve2D:
		stage.OnAfterGrowthCurve2DDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurve2D])
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
	case *HalfwayArcShape:
		stage.OnAfterHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[HalfwayArcShape])
	case *HalfwayArcShapeGrid:
		stage.OnAfterHalfwayArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[HalfwayArcShapeGrid])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[MidArcVectorShapeGrid])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorDeleteCallback = any(callback).(OnAfterDeleteInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridDeleteCallback = any(callback).(OnAfterDeleteInterface[PerpendicularVectorGrid])
	case *PerpendicularVectorGridHalfway:
		stage.OnAfterPerpendicularVectorGridHalfwayDeleteCallback = any(callback).(OnAfterDeleteInterface[PerpendicularVectorGridHalfway])
	case *PerpendicularVectorHalfway:
		stage.OnAfterPerpendicularVectorHalfwayDeleteCallback = any(callback).(OnAfterDeleteInterface[PerpendicularVectorHalfway])
	case *Plant:
		stage.OnAfterPlantDeleteCallback = any(callback).(OnAfterDeleteInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[PlantDiagram])
	case *Rendered3DShape:
		stage.OnAfterRendered3DShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[Rendered3DShape])
	case *RhombusShape:
		stage.OnAfterRhombusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RotatedRhombusShape])
	case *ShiftedBottomTopStartArcShape:
		stage.OnAfterShiftedBottomTopStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedBottomTopStartArcShape])
	case *ShiftedBottomTopStartArcShapeGrid:
		stage.OnAfterShiftedBottomTopStartArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedBottomTopStartArcShapeGrid])
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedLeftStackGrowthCurveEndArcShape])
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedLeftStackGrowthCurveStartArcShape])
	case *ShiftedLeftStackNormalVector:
		stage.OnAfterShiftedLeftStackNormalVectorDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedLeftStackNormalVector])
	case *ShiftedLeftStackOfGrowthCurve:
		stage.OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedLeftStackOfGrowthCurve])
	case *ShiftedLeftStackOfNormalVector:
		stage.OnAfterShiftedLeftStackOfNormalVectorDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedLeftStackOfNormalVector])
	case *StackGrowthCurveEndArcShape:
		stage.OnAfterStackGrowthCurveEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackGrowthCurveEndArcShape])
	case *StackGrowthCurveStartArcShape:
		stage.OnAfterStackGrowthCurveStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackGrowthCurveStartArcShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveDeleteCallback = any(callback).(OnAfterDeleteInterface[StackOfGrowthCurve])
	case *StartArcShape:
		stage.OnAfterStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[StartArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopEndArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DDeleteCallback = any(callback).(OnAfterDeleteInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurveEndArcShape:
		stage.OnAfterTopStackGrowthCurveEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackGrowthCurveEndArcShape])
	case *TopStackGrowthCurveStartArcShape:
		stage.OnAfterTopStackGrowthCurveStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackGrowthCurveStartArcShape])
	case *TopStackOfGrowthCurve:
		stage.OnAfterTopStackOfGrowthCurveDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackOfGrowthCurve])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStartArcShapeGrid])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ArcNormalVectorShape:
		stage.OnAfterArcNormalVectorShapeReadCallback = any(callback).(OnAfterReadInterface[ArcNormalVectorShape])
	case *ArcNormalVectorShapeGrid:
		stage.OnAfterArcNormalVectorShapeGridReadCallback = any(callback).(OnAfterReadInterface[ArcNormalVectorShapeGrid])
	case *AxesShape:
		stage.OnAfterAxesShapeReadCallback = any(callback).(OnAfterReadInterface[AxesShape])
	case *BaseVectorShape:
		stage.OnAfterBaseVectorShapeReadCallback = any(callback).(OnAfterReadInterface[BaseVectorShape])
	case *BaseVectorShapeGrid:
		stage.OnAfterBaseVectorShapeGridReadCallback = any(callback).(OnAfterReadInterface[BaseVectorShapeGrid])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeReadCallback = any(callback).(OnAfterReadInterface[CircleGridShape])
	case *EndArcShape:
		stage.OnAfterEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[EndArcShape])
	case *EndArcShapeGrid:
		stage.OnAfterEndArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[EndArcShapeGrid])
	case *EndHalfwayArcShape:
		stage.OnAfterEndHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[EndHalfwayArcShape])
	case *EndHalfwayArcShapeGrid:
		stage.OnAfterEndHalfwayArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[EndHalfwayArcShapeGrid])
	case *ExplanationTextShape:
		stage.OnAfterExplanationTextShapeReadCallback = any(callback).(OnAfterReadInterface[ExplanationTextShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeReadCallback = any(callback).(OnAfterReadInterface[GridPathShape])
	case *GrowthCurve2D:
		stage.OnAfterGrowthCurve2DReadCallback = any(callback).(OnAfterReadInterface[GrowthCurve2D])
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
	case *HalfwayArcShape:
		stage.OnAfterHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[HalfwayArcShape])
	case *HalfwayArcShapeGrid:
		stage.OnAfterHalfwayArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[HalfwayArcShapeGrid])
	case *InitialRhombusGridShape:
		stage.OnAfterInitialRhombusGridShapeReadCallback = any(callback).(OnAfterReadInterface[InitialRhombusGridShape])
	case *InitialRhombusShape:
		stage.OnAfterInitialRhombusShapeReadCallback = any(callback).(OnAfterReadInterface[InitialRhombusShape])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeReadCallback = any(callback).(OnAfterReadInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridReadCallback = any(callback).(OnAfterReadInterface[MidArcVectorShapeGrid])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeReadCallback = any(callback).(OnAfterReadInterface[NextCircleShape])
	case *PerpendicularVector:
		stage.OnAfterPerpendicularVectorReadCallback = any(callback).(OnAfterReadInterface[PerpendicularVector])
	case *PerpendicularVectorGrid:
		stage.OnAfterPerpendicularVectorGridReadCallback = any(callback).(OnAfterReadInterface[PerpendicularVectorGrid])
	case *PerpendicularVectorGridHalfway:
		stage.OnAfterPerpendicularVectorGridHalfwayReadCallback = any(callback).(OnAfterReadInterface[PerpendicularVectorGridHalfway])
	case *PerpendicularVectorHalfway:
		stage.OnAfterPerpendicularVectorHalfwayReadCallback = any(callback).(OnAfterReadInterface[PerpendicularVectorHalfway])
	case *Plant:
		stage.OnAfterPlantReadCallback = any(callback).(OnAfterReadInterface[Plant])
	case *PlantCircumferenceShape:
		stage.OnAfterPlantCircumferenceShapeReadCallback = any(callback).(OnAfterReadInterface[PlantCircumferenceShape])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramReadCallback = any(callback).(OnAfterReadInterface[PlantDiagram])
	case *Rendered3DShape:
		stage.OnAfterRendered3DShapeReadCallback = any(callback).(OnAfterReadInterface[Rendered3DShape])
	case *RhombusShape:
		stage.OnAfterRhombusShapeReadCallback = any(callback).(OnAfterReadInterface[RhombusShape])
	case *RotatedRhombusGridShape:
		stage.OnAfterRotatedRhombusGridShapeReadCallback = any(callback).(OnAfterReadInterface[RotatedRhombusGridShape])
	case *RotatedRhombusShape:
		stage.OnAfterRotatedRhombusShapeReadCallback = any(callback).(OnAfterReadInterface[RotatedRhombusShape])
	case *ShiftedBottomTopStartArcShape:
		stage.OnAfterShiftedBottomTopStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[ShiftedBottomTopStartArcShape])
	case *ShiftedBottomTopStartArcShapeGrid:
		stage.OnAfterShiftedBottomTopStartArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[ShiftedBottomTopStartArcShapeGrid])
	case *ShiftedLeftStackGrowthCurveEndArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[ShiftedLeftStackGrowthCurveEndArcShape])
	case *ShiftedLeftStackGrowthCurveStartArcShape:
		stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[ShiftedLeftStackGrowthCurveStartArcShape])
	case *ShiftedLeftStackNormalVector:
		stage.OnAfterShiftedLeftStackNormalVectorReadCallback = any(callback).(OnAfterReadInterface[ShiftedLeftStackNormalVector])
	case *ShiftedLeftStackOfGrowthCurve:
		stage.OnAfterShiftedLeftStackOfGrowthCurveReadCallback = any(callback).(OnAfterReadInterface[ShiftedLeftStackOfGrowthCurve])
	case *ShiftedLeftStackOfNormalVector:
		stage.OnAfterShiftedLeftStackOfNormalVectorReadCallback = any(callback).(OnAfterReadInterface[ShiftedLeftStackOfNormalVector])
	case *StackGrowthCurveEndArcShape:
		stage.OnAfterStackGrowthCurveEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[StackGrowthCurveEndArcShape])
	case *StackGrowthCurveStartArcShape:
		stage.OnAfterStackGrowthCurveStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[StackGrowthCurveStartArcShape])
	case *StackOfGrowthCurve:
		stage.OnAfterStackOfGrowthCurveReadCallback = any(callback).(OnAfterReadInterface[StackOfGrowthCurve])
	case *StartArcShape:
		stage.OnAfterStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[StartArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopEndArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DReadCallback = any(callback).(OnAfterReadInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeReadCallback = any(callback).(OnAfterReadInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurveEndArcShape:
		stage.OnAfterTopStackGrowthCurveEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStackGrowthCurveEndArcShape])
	case *TopStackGrowthCurveStartArcShape:
		stage.OnAfterTopStackGrowthCurveStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStackGrowthCurveStartArcShape])
	case *TopStackOfGrowthCurve:
		stage.OnAfterTopStackOfGrowthCurveReadCallback = any(callback).(OnAfterReadInterface[TopStackOfGrowthCurve])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopStartArcShapeGrid])
	}
}
