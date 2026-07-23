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
	case *GrowthCurve2DRibbon:
		if stage.OnAfterGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurve2DRibbonEndShape:
		if stage.OnAfterGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthCurve2DRibbonStartShape:
		if stage.OnAfterGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *MidArcVectorShape:
		if stage.OnAfterMidArcVectorShapeCreateCallback != nil {
			stage.OnAfterMidArcVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *MidArcVectorShapeGrid:
		if stage.OnAfterMidArcVectorShapeGridCreateCallback != nil {
			stage.OnAfterMidArcVectorShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbon:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbonEndShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbonStartShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectory:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2CreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2CreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeCreateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartiallyRotatedTorusShape:
		if stage.OnAfterPartiallyRotatedTorusShapeCreateCallback != nil {
			stage.OnAfterPartiallyRotatedTorusShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *RhombusStuff:
		if stage.OnAfterRhombusStuffCreateCallback != nil {
			stage.OnAfterRhombusStuffCreateCallback.OnAfterCreate(stage, target)
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
	case *ShiftedRightGrowthCurve2DRibbon:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfGrowthCurve2D:
		if stage.OnAfterStackOfGrowthCurve2DCreateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfGrowthCurve2DRibbon:
		if stage.OnAfterStackOfGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfPartiallyRotatedTorusShape:
		if stage.OnAfterStackOfPartiallyRotatedTorusShapeCreateCallback != nil {
			stage.OnAfterStackOfPartiallyRotatedTorusShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfRotatedGrowthCurve2D:
		if stage.OnAfterStackOfRotatedGrowthCurve2DCreateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackOfRotatedGrowthCurve2DRibbon:
		if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonCreateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StackRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeCreateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartArcShape:
		if stage.OnAfterStartArcShapeCreateCallback != nil {
			stage.OnAfterStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartArcShapeGrid:
		if stage.OnAfterStartArcShapeGridCreateCallback != nil {
			stage.OnAfterStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartHalfwayArcShape:
		if stage.OnAfterStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *StartHalfwayArcShapeGrid:
		if stage.OnAfterStartHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndArcShape:
		if stage.OnAfterTopEndArcShapeCreateCallback != nil {
			stage.OnAfterTopEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndArcShapeGrid:
		if stage.OnAfterTopEndArcShapeGridCreateCallback != nil {
			stage.OnAfterTopEndArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndHalfwayArcShape:
		if stage.OnAfterTopEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopEndHalfwayArcShapeGrid:
		if stage.OnAfterTopEndHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
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
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfGrowthCurve2D:
		if stage.OnAfterTopStackOfGrowthCurve2DCreateCallback != nil {
			stage.OnAfterTopStackOfGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2D:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DCreateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeCreateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeCreateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartArcShape:
		if stage.OnAfterTopStartArcShapeCreateCallback != nil {
			stage.OnAfterTopStartArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartArcShapeGrid:
		if stage.OnAfterTopStartArcShapeGridCreateCallback != nil {
			stage.OnAfterTopStartArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartHalfwayArcShape:
		if stage.OnAfterTopStartHalfwayArcShapeCreateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TopStartHalfwayArcShapeGrid:
		if stage.OnAfterTopStartHalfwayArcShapeGridCreateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, target)
		}
	case *TorusStackShape:
		if stage.OnAfterTorusStackShapeCreateCallback != nil {
			stage.OnAfterTorusStackShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *VerticalTorusStackShape:
		if stage.OnAfterVerticalTorusStackShapeCreateCallback != nil {
			stage.OnAfterVerticalTorusStackShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *GrowthCurve2DRibbon:
		newTarget := any(new).(*GrowthCurve2DRibbon)
		if stage.OnAfterGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*GrowthCurve2DRibbonEndShape)
		if stage.OnAfterGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*GrowthCurve2DRibbonStartShape)
		if stage.OnAfterGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *PartiallyGrowthCurve2DRibbon:
		newTarget := any(new).(*PartiallyGrowthCurve2DRibbon)
		if stage.OnAfterPartiallyGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DRibbonEndShape)
		if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DRibbonStartShape)
		if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectory:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectory)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1P2)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2UpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyGrowthCurve2DTrajectoryShape:
		newTarget := any(new).(*PartiallyGrowthCurve2DTrajectoryShape)
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeUpdateCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartiallyRotatedTorusShape:
		newTarget := any(new).(*PartiallyRotatedTorusShape)
		if stage.OnAfterPartiallyRotatedTorusShapeUpdateCallback != nil {
			stage.OnAfterPartiallyRotatedTorusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *RhombusStuff:
		newTarget := any(new).(*RhombusStuff)
		if stage.OnAfterRhombusStuffUpdateCallback != nil {
			stage.OnAfterRhombusStuffUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *ShiftedRightGrowthCurve2DRibbon:
		newTarget := any(new).(*ShiftedRightGrowthCurve2DRibbon)
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*ShiftedRightGrowthCurve2DRibbonEndShape)
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*ShiftedRightGrowthCurve2DRibbonStartShape)
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DEndHalfwayArcShape:
		newTarget := any(new).(*StackGrowthCurve2DEndHalfwayArcShape)
		if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*StackGrowthCurve2DRibbonEndShape)
		if stage.OnAfterStackGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*StackGrowthCurve2DRibbonStartShape)
		if stage.OnAfterStackGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackGrowthCurve2DStartHalfwayArcShape:
		newTarget := any(new).(*StackGrowthCurve2DStartHalfwayArcShape)
		if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfGrowthCurve2D:
		newTarget := any(new).(*StackOfGrowthCurve2D)
		if stage.OnAfterStackOfGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfGrowthCurve2DRibbon:
		newTarget := any(new).(*StackOfGrowthCurve2DRibbon)
		if stage.OnAfterStackOfGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfPartiallyRotatedTorusShape:
		newTarget := any(new).(*StackOfPartiallyRotatedTorusShape)
		if stage.OnAfterStackOfPartiallyRotatedTorusShapeUpdateCallback != nil {
			stage.OnAfterStackOfPartiallyRotatedTorusShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfRotatedGrowthCurve2D:
		newTarget := any(new).(*StackOfRotatedGrowthCurve2D)
		if stage.OnAfterStackOfRotatedGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackOfRotatedGrowthCurve2DRibbon:
		newTarget := any(new).(*StackOfRotatedGrowthCurve2DRibbon)
		if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonUpdateCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DEndArcShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DEndArcShape)
		if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DRibbonEndShape)
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DRibbonStartShape)
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StackRotatedGrowthCurve2DStartArcShape:
		newTarget := any(new).(*StackRotatedGrowthCurve2DStartArcShape)
		if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeUpdateCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *StartHalfwayArcShape:
		newTarget := any(new).(*StartHalfwayArcShape)
		if stage.OnAfterStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *StartHalfwayArcShapeGrid:
		newTarget := any(new).(*StartHalfwayArcShapeGrid)
		if stage.OnAfterStartHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterStartHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *TopEndHalfwayArcShape:
		newTarget := any(new).(*TopEndHalfwayArcShape)
		if stage.OnAfterTopEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopEndHalfwayArcShapeGrid:
		newTarget := any(new).(*TopEndHalfwayArcShapeGrid)
		if stage.OnAfterTopEndHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		newTarget := any(new).(*TopStackGrowthCurve2DEndHalfwayArcShape)
		if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		newTarget := any(new).(*TopStackGrowthCurve2DStartHalfwayArcShape)
		if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfGrowthCurve2D:
		newTarget := any(new).(*TopStackOfGrowthCurve2D)
		if stage.OnAfterTopStackOfGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterTopStackOfGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfRotatedGrowthCurve2D:
		newTarget := any(new).(*TopStackOfRotatedGrowthCurve2D)
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DUpdateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		newTarget := any(new).(*TopStackOfRotatedGrowthCurve2DEndArcShape)
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		newTarget := any(new).(*TopStackOfRotatedGrowthCurve2DStartArcShape)
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeUpdateCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *TopStartHalfwayArcShape:
		newTarget := any(new).(*TopStartHalfwayArcShape)
		if stage.OnAfterTopStartHalfwayArcShapeUpdateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TopStartHalfwayArcShapeGrid:
		newTarget := any(new).(*TopStartHalfwayArcShapeGrid)
		if stage.OnAfterTopStartHalfwayArcShapeGridUpdateCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TorusStackShape:
		newTarget := any(new).(*TorusStackShape)
		if stage.OnAfterTorusStackShapeUpdateCallback != nil {
			stage.OnAfterTorusStackShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *VerticalTorusStackShape:
		newTarget := any(new).(*VerticalTorusStackShape)
		if stage.OnAfterVerticalTorusStackShapeUpdateCallback != nil {
			stage.OnAfterVerticalTorusStackShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *GrowthCurve2DRibbon:
		if stage.OnAfterGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2DRibbon)
			stage.OnAfterGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurve2DRibbonEndShape:
		if stage.OnAfterGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2DRibbonEndShape)
			stage.OnAfterGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthCurve2DRibbonStartShape:
		if stage.OnAfterGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthCurve2DRibbonStartShape)
			stage.OnAfterGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *PartiallyGrowthCurve2DRibbon:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DRibbon)
			stage.OnAfterPartiallyGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DRibbonEndShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DRibbonEndShape)
			stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DRibbonStartShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DRibbonStartShape)
			stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectory:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectory)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2DeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1P2)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyGrowthCurve2DTrajectoryShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyGrowthCurve2DTrajectoryShape)
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartiallyRotatedTorusShape:
		if stage.OnAfterPartiallyRotatedTorusShapeDeleteCallback != nil {
			staged := any(staged).(*PartiallyRotatedTorusShape)
			stage.OnAfterPartiallyRotatedTorusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *RhombusStuff:
		if stage.OnAfterRhombusStuffDeleteCallback != nil {
			staged := any(staged).(*RhombusStuff)
			stage.OnAfterRhombusStuffDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *ShiftedRightGrowthCurve2DRibbon:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*ShiftedRightGrowthCurve2DRibbon)
			stage.OnAfterShiftedRightGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedRightGrowthCurve2DRibbonEndShape)
			stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*ShiftedRightGrowthCurve2DRibbonStartShape)
			stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DEndHalfwayArcShape)
			stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DRibbonEndShape)
			stage.OnAfterStackGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DRibbonStartShape)
			stage.OnAfterStackGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackGrowthCurve2DStartHalfwayArcShape)
			stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfGrowthCurve2D:
		if stage.OnAfterStackOfGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*StackOfGrowthCurve2D)
			stage.OnAfterStackOfGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfGrowthCurve2DRibbon:
		if stage.OnAfterStackOfGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*StackOfGrowthCurve2DRibbon)
			stage.OnAfterStackOfGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfPartiallyRotatedTorusShape:
		if stage.OnAfterStackOfPartiallyRotatedTorusShapeDeleteCallback != nil {
			staged := any(staged).(*StackOfPartiallyRotatedTorusShape)
			stage.OnAfterStackOfPartiallyRotatedTorusShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfRotatedGrowthCurve2D:
		if stage.OnAfterStackOfRotatedGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*StackOfRotatedGrowthCurve2D)
			stage.OnAfterStackOfRotatedGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackOfRotatedGrowthCurve2DRibbon:
		if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonDeleteCallback != nil {
			staged := any(staged).(*StackOfRotatedGrowthCurve2DRibbon)
			stage.OnAfterStackOfRotatedGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DEndArcShape)
			stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DRibbonEndShape)
			stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DRibbonStartShape)
			stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StackRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*StackRotatedGrowthCurve2DStartArcShape)
			stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *StartHalfwayArcShape:
		if stage.OnAfterStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*StartHalfwayArcShape)
			stage.OnAfterStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *StartHalfwayArcShapeGrid:
		if stage.OnAfterStartHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*StartHalfwayArcShapeGrid)
			stage.OnAfterStartHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *TopEndHalfwayArcShape:
		if stage.OnAfterTopEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopEndHalfwayArcShape)
			stage.OnAfterTopEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopEndHalfwayArcShapeGrid:
		if stage.OnAfterTopEndHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopEndHalfwayArcShapeGrid)
			stage.OnAfterTopEndHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackGrowthCurve2DEndHalfwayArcShape)
			stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackGrowthCurve2DStartHalfwayArcShape)
			stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfGrowthCurve2D:
		if stage.OnAfterTopStackOfGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*TopStackOfGrowthCurve2D)
			stage.OnAfterTopStackOfGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfRotatedGrowthCurve2D:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DDeleteCallback != nil {
			staged := any(staged).(*TopStackOfRotatedGrowthCurve2D)
			stage.OnAfterTopStackOfRotatedGrowthCurve2DDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackOfRotatedGrowthCurve2DEndArcShape)
			stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStackOfRotatedGrowthCurve2DStartArcShape)
			stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *TopStartHalfwayArcShape:
		if stage.OnAfterTopStartHalfwayArcShapeDeleteCallback != nil {
			staged := any(staged).(*TopStartHalfwayArcShape)
			stage.OnAfterTopStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TopStartHalfwayArcShapeGrid:
		if stage.OnAfterTopStartHalfwayArcShapeGridDeleteCallback != nil {
			staged := any(staged).(*TopStartHalfwayArcShapeGrid)
			stage.OnAfterTopStartHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TorusStackShape:
		if stage.OnAfterTorusStackShapeDeleteCallback != nil {
			staged := any(staged).(*TorusStackShape)
			stage.OnAfterTorusStackShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *VerticalTorusStackShape:
		if stage.OnAfterVerticalTorusStackShapeDeleteCallback != nil {
			staged := any(staged).(*VerticalTorusStackShape)
			stage.OnAfterVerticalTorusStackShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *GrowthCurve2DRibbon:
		if stage.OnAfterGrowthCurve2DRibbonReadCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthCurve2DRibbonEndShape:
		if stage.OnAfterGrowthCurve2DRibbonEndShapeReadCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonEndShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthCurve2DRibbonStartShape:
		if stage.OnAfterGrowthCurve2DRibbonStartShapeReadCallback != nil {
			stage.OnAfterGrowthCurve2DRibbonStartShapeReadCallback.OnAfterRead(stage, target)
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
	case *MidArcVectorShape:
		if stage.OnAfterMidArcVectorShapeReadCallback != nil {
			stage.OnAfterMidArcVectorShapeReadCallback.OnAfterRead(stage, target)
		}
	case *MidArcVectorShapeGrid:
		if stage.OnAfterMidArcVectorShapeGridReadCallback != nil {
			stage.OnAfterMidArcVectorShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbon:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbonEndShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DRibbonStartShape:
		if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectory:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2ReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2ReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyGrowthCurve2DTrajectoryShape:
		if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeReadCallback != nil {
			stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeReadCallback.OnAfterRead(stage, target)
		}
	case *PartiallyRotatedTorusShape:
		if stage.OnAfterPartiallyRotatedTorusShapeReadCallback != nil {
			stage.OnAfterPartiallyRotatedTorusShapeReadCallback.OnAfterRead(stage, target)
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
	case *RhombusStuff:
		if stage.OnAfterRhombusStuffReadCallback != nil {
			stage.OnAfterRhombusStuffReadCallback.OnAfterRead(stage, target)
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
	case *ShiftedRightGrowthCurve2DRibbon:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonReadCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeReadCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeReadCallback != nil {
			stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeReadCallback != nil {
			stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackGrowthCurve2DRibbonEndShapeReadCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonEndShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackGrowthCurve2DRibbonStartShapeReadCallback != nil {
			stage.OnAfterStackGrowthCurve2DRibbonStartShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeReadCallback != nil {
			stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackOfGrowthCurve2D:
		if stage.OnAfterStackOfGrowthCurve2DReadCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DReadCallback.OnAfterRead(stage, target)
		}
	case *StackOfGrowthCurve2DRibbon:
		if stage.OnAfterStackOfGrowthCurve2DRibbonReadCallback != nil {
			stage.OnAfterStackOfGrowthCurve2DRibbonReadCallback.OnAfterRead(stage, target)
		}
	case *StackOfPartiallyRotatedTorusShape:
		if stage.OnAfterStackOfPartiallyRotatedTorusShapeReadCallback != nil {
			stage.OnAfterStackOfPartiallyRotatedTorusShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackOfRotatedGrowthCurve2D:
		if stage.OnAfterStackOfRotatedGrowthCurve2DReadCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DReadCallback.OnAfterRead(stage, target)
		}
	case *StackOfRotatedGrowthCurve2DRibbon:
		if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonReadCallback != nil {
			stage.OnAfterStackOfRotatedGrowthCurve2DRibbonReadCallback.OnAfterRead(stage, target)
		}
	case *StackRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeReadCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeReadCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeReadCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StackRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeReadCallback != nil {
			stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StartArcShape:
		if stage.OnAfterStartArcShapeReadCallback != nil {
			stage.OnAfterStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StartArcShapeGrid:
		if stage.OnAfterStartArcShapeGridReadCallback != nil {
			stage.OnAfterStartArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *StartHalfwayArcShape:
		if stage.OnAfterStartHalfwayArcShapeReadCallback != nil {
			stage.OnAfterStartHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *StartHalfwayArcShapeGrid:
		if stage.OnAfterStartHalfwayArcShapeGridReadCallback != nil {
			stage.OnAfterStartHalfwayArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *TopEndArcShape:
		if stage.OnAfterTopEndArcShapeReadCallback != nil {
			stage.OnAfterTopEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopEndArcShapeGrid:
		if stage.OnAfterTopEndArcShapeGridReadCallback != nil {
			stage.OnAfterTopEndArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *TopEndHalfwayArcShape:
		if stage.OnAfterTopEndHalfwayArcShapeReadCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopEndHalfwayArcShapeGrid:
		if stage.OnAfterTopEndHalfwayArcShapeGridReadCallback != nil {
			stage.OnAfterTopEndHalfwayArcShapeGridReadCallback.OnAfterRead(stage, target)
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
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeReadCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeReadCallback != nil {
			stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackOfGrowthCurve2D:
		if stage.OnAfterTopStackOfGrowthCurve2DReadCallback != nil {
			stage.OnAfterTopStackOfGrowthCurve2DReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2D:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DReadCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeReadCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeReadCallback != nil {
			stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStartArcShape:
		if stage.OnAfterTopStartArcShapeReadCallback != nil {
			stage.OnAfterTopStartArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStartArcShapeGrid:
		if stage.OnAfterTopStartArcShapeGridReadCallback != nil {
			stage.OnAfterTopStartArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *TopStartHalfwayArcShape:
		if stage.OnAfterTopStartHalfwayArcShapeReadCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TopStartHalfwayArcShapeGrid:
		if stage.OnAfterTopStartHalfwayArcShapeGridReadCallback != nil {
			stage.OnAfterTopStartHalfwayArcShapeGridReadCallback.OnAfterRead(stage, target)
		}
	case *TorusStackShape:
		if stage.OnAfterTorusStackShapeReadCallback != nil {
			stage.OnAfterTorusStackShapeReadCallback.OnAfterRead(stage, target)
		}
	case *VerticalTorusStackShape:
		if stage.OnAfterVerticalTorusStackShapeReadCallback != nil {
			stage.OnAfterVerticalTorusStackShapeReadCallback.OnAfterRead(stage, target)
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
	case *GrowthCurve2DRibbon:
		stage.OnAfterGrowthCurve2DRibbonUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurve2DRibbon])
	case *GrowthCurve2DRibbonEndShape:
		stage.OnAfterGrowthCurve2DRibbonEndShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurve2DRibbonEndShape])
	case *GrowthCurve2DRibbonStartShape:
		stage.OnAfterGrowthCurve2DRibbonStartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthCurve2DRibbonStartShape])
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
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[MidArcVectorShapeGrid])
	case *PartiallyGrowthCurve2DRibbon:
		stage.OnAfterPartiallyGrowthCurve2DRibbonUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DRibbon])
	case *PartiallyGrowthCurve2DRibbonEndShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DRibbonEndShape])
	case *PartiallyGrowthCurve2DRibbonStartShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DRibbonStartShape])
	case *PartiallyGrowthCurve2DTrajectory:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectory])
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectoryP1CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2UpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectoryP1P2])
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape])
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectoryP1PointShape])
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectoryP2CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectoryP2PointShape])
	case *PartiallyGrowthCurve2DTrajectoryShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyGrowthCurve2DTrajectoryShape])
	case *PartiallyRotatedTorusShape:
		stage.OnAfterPartiallyRotatedTorusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartiallyRotatedTorusShape])
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
	case *RhombusStuff:
		stage.OnAfterRhombusStuffUpdateCallback = any(callback).(OnAfterUpdateInterface[RhombusStuff])
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
	case *ShiftedRightGrowthCurve2DRibbon:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedRightGrowthCurve2DRibbon])
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedRightGrowthCurve2DRibbonEndShape])
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ShiftedRightGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackGrowthCurve2DEndHalfwayArcShape])
	case *StackGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackGrowthCurve2DRibbonEndShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackGrowthCurve2DRibbonEndShape])
	case *StackGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackGrowthCurve2DRibbonStartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackGrowthCurve2DStartHalfwayArcShape])
	case *StackOfGrowthCurve2D:
		stage.OnAfterStackOfGrowthCurve2DUpdateCallback = any(callback).(OnAfterUpdateInterface[StackOfGrowthCurve2D])
	case *StackOfGrowthCurve2DRibbon:
		stage.OnAfterStackOfGrowthCurve2DRibbonUpdateCallback = any(callback).(OnAfterUpdateInterface[StackOfGrowthCurve2DRibbon])
	case *StackOfPartiallyRotatedTorusShape:
		stage.OnAfterStackOfPartiallyRotatedTorusShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackOfPartiallyRotatedTorusShape])
	case *StackOfRotatedGrowthCurve2D:
		stage.OnAfterStackOfRotatedGrowthCurve2DUpdateCallback = any(callback).(OnAfterUpdateInterface[StackOfRotatedGrowthCurve2D])
	case *StackOfRotatedGrowthCurve2DRibbon:
		stage.OnAfterStackOfRotatedGrowthCurve2DRibbonUpdateCallback = any(callback).(OnAfterUpdateInterface[StackOfRotatedGrowthCurve2DRibbon])
	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackRotatedGrowthCurve2DEndArcShape])
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackRotatedGrowthCurve2DRibbonEndShape])
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackRotatedGrowthCurve2DRibbonStartShape])
	case *StackRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StackRotatedGrowthCurve2DStartArcShape])
	case *StartArcShape:
		stage.OnAfterStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[StartArcShapeGrid])
	case *StartHalfwayArcShape:
		stage.OnAfterStartHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[StartHalfwayArcShape])
	case *StartHalfwayArcShapeGrid:
		stage.OnAfterStartHalfwayArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[StartHalfwayArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopEndArcShapeGrid])
	case *TopEndHalfwayArcShape:
		stage.OnAfterTopEndHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopEndHalfwayArcShape])
	case *TopEndHalfwayArcShapeGrid:
		stage.OnAfterTopEndHalfwayArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopEndHalfwayArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DUpdateCallback = any(callback).(OnAfterUpdateInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackGrowthCurve2DEndHalfwayArcShape])
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackGrowthCurve2DStartHalfwayArcShape])
	case *TopStackOfGrowthCurve2D:
		stage.OnAfterTopStackOfGrowthCurve2DUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackOfGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2D:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackOfRotatedGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackOfRotatedGrowthCurve2DEndArcShape])
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStackOfRotatedGrowthCurve2DStartArcShape])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStartArcShapeGrid])
	case *TopStartHalfwayArcShape:
		stage.OnAfterTopStartHalfwayArcShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStartHalfwayArcShape])
	case *TopStartHalfwayArcShapeGrid:
		stage.OnAfterTopStartHalfwayArcShapeGridUpdateCallback = any(callback).(OnAfterUpdateInterface[TopStartHalfwayArcShapeGrid])
	case *TorusStackShape:
		stage.OnAfterTorusStackShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TorusStackShape])
	case *VerticalTorusStackShape:
		stage.OnAfterVerticalTorusStackShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[VerticalTorusStackShape])
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
	case *GrowthCurve2DRibbon:
		stage.OnAfterGrowthCurve2DRibbonCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurve2DRibbon])
	case *GrowthCurve2DRibbonEndShape:
		stage.OnAfterGrowthCurve2DRibbonEndShapeCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurve2DRibbonEndShape])
	case *GrowthCurve2DRibbonStartShape:
		stage.OnAfterGrowthCurve2DRibbonStartShapeCreateCallback = any(callback).(OnAfterCreateInterface[GrowthCurve2DRibbonStartShape])
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
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[MidArcVectorShapeGrid])
	case *PartiallyGrowthCurve2DRibbon:
		stage.OnAfterPartiallyGrowthCurve2DRibbonCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DRibbon])
	case *PartiallyGrowthCurve2DRibbonEndShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DRibbonEndShape])
	case *PartiallyGrowthCurve2DRibbonStartShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DRibbonStartShape])
	case *PartiallyGrowthCurve2DTrajectory:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectory])
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectoryP1CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2CreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectoryP1P2])
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape])
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectoryP1PointShape])
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectoryP2CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectoryP2PointShape])
	case *PartiallyGrowthCurve2DTrajectoryShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyGrowthCurve2DTrajectoryShape])
	case *PartiallyRotatedTorusShape:
		stage.OnAfterPartiallyRotatedTorusShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartiallyRotatedTorusShape])
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
	case *RhombusStuff:
		stage.OnAfterRhombusStuffCreateCallback = any(callback).(OnAfterCreateInterface[RhombusStuff])
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
	case *ShiftedRightGrowthCurve2DRibbon:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedRightGrowthCurve2DRibbon])
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedRightGrowthCurve2DRibbonEndShape])
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeCreateCallback = any(callback).(OnAfterCreateInterface[ShiftedRightGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackGrowthCurve2DEndHalfwayArcShape])
	case *StackGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackGrowthCurve2DRibbonEndShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackGrowthCurve2DRibbonEndShape])
	case *StackGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackGrowthCurve2DRibbonStartShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackGrowthCurve2DStartHalfwayArcShape])
	case *StackOfGrowthCurve2D:
		stage.OnAfterStackOfGrowthCurve2DCreateCallback = any(callback).(OnAfterCreateInterface[StackOfGrowthCurve2D])
	case *StackOfGrowthCurve2DRibbon:
		stage.OnAfterStackOfGrowthCurve2DRibbonCreateCallback = any(callback).(OnAfterCreateInterface[StackOfGrowthCurve2DRibbon])
	case *StackOfPartiallyRotatedTorusShape:
		stage.OnAfterStackOfPartiallyRotatedTorusShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackOfPartiallyRotatedTorusShape])
	case *StackOfRotatedGrowthCurve2D:
		stage.OnAfterStackOfRotatedGrowthCurve2DCreateCallback = any(callback).(OnAfterCreateInterface[StackOfRotatedGrowthCurve2D])
	case *StackOfRotatedGrowthCurve2DRibbon:
		stage.OnAfterStackOfRotatedGrowthCurve2DRibbonCreateCallback = any(callback).(OnAfterCreateInterface[StackOfRotatedGrowthCurve2DRibbon])
	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackRotatedGrowthCurve2DEndArcShape])
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackRotatedGrowthCurve2DRibbonEndShape])
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackRotatedGrowthCurve2DRibbonStartShape])
	case *StackRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StackRotatedGrowthCurve2DStartArcShape])
	case *StartArcShape:
		stage.OnAfterStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[StartArcShapeGrid])
	case *StartHalfwayArcShape:
		stage.OnAfterStartHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[StartHalfwayArcShape])
	case *StartHalfwayArcShapeGrid:
		stage.OnAfterStartHalfwayArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[StartHalfwayArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopEndArcShapeGrid])
	case *TopEndHalfwayArcShape:
		stage.OnAfterTopEndHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopEndHalfwayArcShape])
	case *TopEndHalfwayArcShapeGrid:
		stage.OnAfterTopEndHalfwayArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopEndHalfwayArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DCreateCallback = any(callback).(OnAfterCreateInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStackGrowthCurve2DEndHalfwayArcShape])
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStackGrowthCurve2DStartHalfwayArcShape])
	case *TopStackOfGrowthCurve2D:
		stage.OnAfterTopStackOfGrowthCurve2DCreateCallback = any(callback).(OnAfterCreateInterface[TopStackOfGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2D:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DCreateCallback = any(callback).(OnAfterCreateInterface[TopStackOfRotatedGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStackOfRotatedGrowthCurve2DEndArcShape])
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStackOfRotatedGrowthCurve2DStartArcShape])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopStartArcShapeGrid])
	case *TopStartHalfwayArcShape:
		stage.OnAfterTopStartHalfwayArcShapeCreateCallback = any(callback).(OnAfterCreateInterface[TopStartHalfwayArcShape])
	case *TopStartHalfwayArcShapeGrid:
		stage.OnAfterTopStartHalfwayArcShapeGridCreateCallback = any(callback).(OnAfterCreateInterface[TopStartHalfwayArcShapeGrid])
	case *TorusStackShape:
		stage.OnAfterTorusStackShapeCreateCallback = any(callback).(OnAfterCreateInterface[TorusStackShape])
	case *VerticalTorusStackShape:
		stage.OnAfterVerticalTorusStackShapeCreateCallback = any(callback).(OnAfterCreateInterface[VerticalTorusStackShape])
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
	case *GrowthCurve2DRibbon:
		stage.OnAfterGrowthCurve2DRibbonDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurve2DRibbon])
	case *GrowthCurve2DRibbonEndShape:
		stage.OnAfterGrowthCurve2DRibbonEndShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurve2DRibbonEndShape])
	case *GrowthCurve2DRibbonStartShape:
		stage.OnAfterGrowthCurve2DRibbonStartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthCurve2DRibbonStartShape])
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
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[MidArcVectorShapeGrid])
	case *PartiallyGrowthCurve2DRibbon:
		stage.OnAfterPartiallyGrowthCurve2DRibbonDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DRibbon])
	case *PartiallyGrowthCurve2DRibbonEndShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DRibbonEndShape])
	case *PartiallyGrowthCurve2DRibbonStartShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DRibbonStartShape])
	case *PartiallyGrowthCurve2DTrajectory:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectory])
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectoryP1CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2DeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectoryP1P2])
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape])
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectoryP1PointShape])
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectoryP2CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectoryP2PointShape])
	case *PartiallyGrowthCurve2DTrajectoryShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyGrowthCurve2DTrajectoryShape])
	case *PartiallyRotatedTorusShape:
		stage.OnAfterPartiallyRotatedTorusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartiallyRotatedTorusShape])
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
	case *RhombusStuff:
		stage.OnAfterRhombusStuffDeleteCallback = any(callback).(OnAfterDeleteInterface[RhombusStuff])
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
	case *ShiftedRightGrowthCurve2DRibbon:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedRightGrowthCurve2DRibbon])
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedRightGrowthCurve2DRibbonEndShape])
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ShiftedRightGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackGrowthCurve2DEndHalfwayArcShape])
	case *StackGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackGrowthCurve2DRibbonEndShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackGrowthCurve2DRibbonEndShape])
	case *StackGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackGrowthCurve2DRibbonStartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackGrowthCurve2DStartHalfwayArcShape])
	case *StackOfGrowthCurve2D:
		stage.OnAfterStackOfGrowthCurve2DDeleteCallback = any(callback).(OnAfterDeleteInterface[StackOfGrowthCurve2D])
	case *StackOfGrowthCurve2DRibbon:
		stage.OnAfterStackOfGrowthCurve2DRibbonDeleteCallback = any(callback).(OnAfterDeleteInterface[StackOfGrowthCurve2DRibbon])
	case *StackOfPartiallyRotatedTorusShape:
		stage.OnAfterStackOfPartiallyRotatedTorusShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackOfPartiallyRotatedTorusShape])
	case *StackOfRotatedGrowthCurve2D:
		stage.OnAfterStackOfRotatedGrowthCurve2DDeleteCallback = any(callback).(OnAfterDeleteInterface[StackOfRotatedGrowthCurve2D])
	case *StackOfRotatedGrowthCurve2DRibbon:
		stage.OnAfterStackOfRotatedGrowthCurve2DRibbonDeleteCallback = any(callback).(OnAfterDeleteInterface[StackOfRotatedGrowthCurve2DRibbon])
	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackRotatedGrowthCurve2DEndArcShape])
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackRotatedGrowthCurve2DRibbonEndShape])
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackRotatedGrowthCurve2DRibbonStartShape])
	case *StackRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StackRotatedGrowthCurve2DStartArcShape])
	case *StartArcShape:
		stage.OnAfterStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[StartArcShapeGrid])
	case *StartHalfwayArcShape:
		stage.OnAfterStartHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[StartHalfwayArcShape])
	case *StartHalfwayArcShapeGrid:
		stage.OnAfterStartHalfwayArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[StartHalfwayArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopEndArcShapeGrid])
	case *TopEndHalfwayArcShape:
		stage.OnAfterTopEndHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopEndHalfwayArcShape])
	case *TopEndHalfwayArcShapeGrid:
		stage.OnAfterTopEndHalfwayArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopEndHalfwayArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DDeleteCallback = any(callback).(OnAfterDeleteInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackGrowthCurve2DEndHalfwayArcShape])
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackGrowthCurve2DStartHalfwayArcShape])
	case *TopStackOfGrowthCurve2D:
		stage.OnAfterTopStackOfGrowthCurve2DDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackOfGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2D:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackOfRotatedGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackOfRotatedGrowthCurve2DEndArcShape])
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStackOfRotatedGrowthCurve2DStartArcShape])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStartArcShapeGrid])
	case *TopStartHalfwayArcShape:
		stage.OnAfterTopStartHalfwayArcShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStartHalfwayArcShape])
	case *TopStartHalfwayArcShapeGrid:
		stage.OnAfterTopStartHalfwayArcShapeGridDeleteCallback = any(callback).(OnAfterDeleteInterface[TopStartHalfwayArcShapeGrid])
	case *TorusStackShape:
		stage.OnAfterTorusStackShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TorusStackShape])
	case *VerticalTorusStackShape:
		stage.OnAfterVerticalTorusStackShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[VerticalTorusStackShape])
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
	case *GrowthCurve2DRibbon:
		stage.OnAfterGrowthCurve2DRibbonReadCallback = any(callback).(OnAfterReadInterface[GrowthCurve2DRibbon])
	case *GrowthCurve2DRibbonEndShape:
		stage.OnAfterGrowthCurve2DRibbonEndShapeReadCallback = any(callback).(OnAfterReadInterface[GrowthCurve2DRibbonEndShape])
	case *GrowthCurve2DRibbonStartShape:
		stage.OnAfterGrowthCurve2DRibbonStartShapeReadCallback = any(callback).(OnAfterReadInterface[GrowthCurve2DRibbonStartShape])
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
	case *MidArcVectorShape:
		stage.OnAfterMidArcVectorShapeReadCallback = any(callback).(OnAfterReadInterface[MidArcVectorShape])
	case *MidArcVectorShapeGrid:
		stage.OnAfterMidArcVectorShapeGridReadCallback = any(callback).(OnAfterReadInterface[MidArcVectorShapeGrid])
	case *PartiallyGrowthCurve2DRibbon:
		stage.OnAfterPartiallyGrowthCurve2DRibbonReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DRibbon])
	case *PartiallyGrowthCurve2DRibbonEndShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DRibbonEndShape])
	case *PartiallyGrowthCurve2DRibbonStartShape:
		stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DRibbonStartShape])
	case *PartiallyGrowthCurve2DTrajectory:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectory])
	case *PartiallyGrowthCurve2DTrajectoryP1CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectoryP1CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP1P2:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2ReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectoryP1P2])
	case *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape])
	case *PartiallyGrowthCurve2DTrajectoryP1PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectoryP1PointShape])
	case *PartiallyGrowthCurve2DTrajectoryP2CurveShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectoryP2CurveShape])
	case *PartiallyGrowthCurve2DTrajectoryP2PointShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectoryP2PointShape])
	case *PartiallyGrowthCurve2DTrajectoryShape:
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyGrowthCurve2DTrajectoryShape])
	case *PartiallyRotatedTorusShape:
		stage.OnAfterPartiallyRotatedTorusShapeReadCallback = any(callback).(OnAfterReadInterface[PartiallyRotatedTorusShape])
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
	case *RhombusStuff:
		stage.OnAfterRhombusStuffReadCallback = any(callback).(OnAfterReadInterface[RhombusStuff])
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
	case *ShiftedRightGrowthCurve2DRibbon:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonReadCallback = any(callback).(OnAfterReadInterface[ShiftedRightGrowthCurve2DRibbon])
	case *ShiftedRightGrowthCurve2DRibbonEndShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeReadCallback = any(callback).(OnAfterReadInterface[ShiftedRightGrowthCurve2DRibbonEndShape])
	case *ShiftedRightGrowthCurve2DRibbonStartShape:
		stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeReadCallback = any(callback).(OnAfterReadInterface[ShiftedRightGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[StackGrowthCurve2DEndHalfwayArcShape])
	case *StackGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackGrowthCurve2DRibbonEndShapeReadCallback = any(callback).(OnAfterReadInterface[StackGrowthCurve2DRibbonEndShape])
	case *StackGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackGrowthCurve2DRibbonStartShapeReadCallback = any(callback).(OnAfterReadInterface[StackGrowthCurve2DRibbonStartShape])
	case *StackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[StackGrowthCurve2DStartHalfwayArcShape])
	case *StackOfGrowthCurve2D:
		stage.OnAfterStackOfGrowthCurve2DReadCallback = any(callback).(OnAfterReadInterface[StackOfGrowthCurve2D])
	case *StackOfGrowthCurve2DRibbon:
		stage.OnAfterStackOfGrowthCurve2DRibbonReadCallback = any(callback).(OnAfterReadInterface[StackOfGrowthCurve2DRibbon])
	case *StackOfPartiallyRotatedTorusShape:
		stage.OnAfterStackOfPartiallyRotatedTorusShapeReadCallback = any(callback).(OnAfterReadInterface[StackOfPartiallyRotatedTorusShape])
	case *StackOfRotatedGrowthCurve2D:
		stage.OnAfterStackOfRotatedGrowthCurve2DReadCallback = any(callback).(OnAfterReadInterface[StackOfRotatedGrowthCurve2D])
	case *StackOfRotatedGrowthCurve2DRibbon:
		stage.OnAfterStackOfRotatedGrowthCurve2DRibbonReadCallback = any(callback).(OnAfterReadInterface[StackOfRotatedGrowthCurve2DRibbon])
	case *StackRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[StackRotatedGrowthCurve2DEndArcShape])
	case *StackRotatedGrowthCurve2DRibbonEndShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeReadCallback = any(callback).(OnAfterReadInterface[StackRotatedGrowthCurve2DRibbonEndShape])
	case *StackRotatedGrowthCurve2DRibbonStartShape:
		stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeReadCallback = any(callback).(OnAfterReadInterface[StackRotatedGrowthCurve2DRibbonStartShape])
	case *StackRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[StackRotatedGrowthCurve2DStartArcShape])
	case *StartArcShape:
		stage.OnAfterStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[StartArcShape])
	case *StartArcShapeGrid:
		stage.OnAfterStartArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[StartArcShapeGrid])
	case *StartHalfwayArcShape:
		stage.OnAfterStartHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[StartHalfwayArcShape])
	case *StartHalfwayArcShapeGrid:
		stage.OnAfterStartHalfwayArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[StartHalfwayArcShapeGrid])
	case *TopEndArcShape:
		stage.OnAfterTopEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopEndArcShape])
	case *TopEndArcShapeGrid:
		stage.OnAfterTopEndArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopEndArcShapeGrid])
	case *TopEndHalfwayArcShape:
		stage.OnAfterTopEndHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopEndHalfwayArcShape])
	case *TopEndHalfwayArcShapeGrid:
		stage.OnAfterTopEndHalfwayArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopEndHalfwayArcShapeGrid])
	case *TopGrowthCurve2D:
		stage.OnAfterTopGrowthCurve2DReadCallback = any(callback).(OnAfterReadInterface[TopGrowthCurve2D])
	case *TopMidArcVectorShape:
		stage.OnAfterTopMidArcVectorShapeReadCallback = any(callback).(OnAfterReadInterface[TopMidArcVectorShape])
	case *TopMidArcVectorShapeGrid:
		stage.OnAfterTopMidArcVectorShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopMidArcVectorShapeGrid])
	case *TopStackGrowthCurve2DEndHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStackGrowthCurve2DEndHalfwayArcShape])
	case *TopStackGrowthCurve2DStartHalfwayArcShape:
		stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStackGrowthCurve2DStartHalfwayArcShape])
	case *TopStackOfGrowthCurve2D:
		stage.OnAfterTopStackOfGrowthCurve2DReadCallback = any(callback).(OnAfterReadInterface[TopStackOfGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2D:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DReadCallback = any(callback).(OnAfterReadInterface[TopStackOfRotatedGrowthCurve2D])
	case *TopStackOfRotatedGrowthCurve2DEndArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStackOfRotatedGrowthCurve2DEndArcShape])
	case *TopStackOfRotatedGrowthCurve2DStartArcShape:
		stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStackOfRotatedGrowthCurve2DStartArcShape])
	case *TopStartArcShape:
		stage.OnAfterTopStartArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStartArcShape])
	case *TopStartArcShapeGrid:
		stage.OnAfterTopStartArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopStartArcShapeGrid])
	case *TopStartHalfwayArcShape:
		stage.OnAfterTopStartHalfwayArcShapeReadCallback = any(callback).(OnAfterReadInterface[TopStartHalfwayArcShape])
	case *TopStartHalfwayArcShapeGrid:
		stage.OnAfterTopStartHalfwayArcShapeGridReadCallback = any(callback).(OnAfterReadInterface[TopStartHalfwayArcShapeGrid])
	case *TorusStackShape:
		stage.OnAfterTorusStackShapeReadCallback = any(callback).(OnAfterReadInterface[TorusStackShape])
	case *VerticalTorusStackShape:
		stage.OnAfterVerticalTorusStackShapeReadCallback = any(callback).(OnAfterReadInterface[VerticalTorusStackShape])
	}
}
