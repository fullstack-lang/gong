// generated code - do not edit
package models

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
func (angle0shape *Angle0Shape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAngle0ShapeCreateCallback != nil {
		stage.OnAfterAngle0ShapeCreateCallback.OnAfterCreate(stage, angle0shape)
	}
}

func (angle0shape *Angle0Shape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngle0ShapeUpdateCallback != nil {
		var frontAngle0Shape *Angle0Shape
		if front != nil {
			frontAngle0Shape, _ = front.(*Angle0Shape)
		}
		stage.OnAfterAngle0ShapeUpdateCallback.OnAfterUpdate(stage, angle0shape, frontAngle0Shape)
	}
}

func (angle0shape *Angle0Shape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAngle0ShapeDeleteCallback != nil {
		var frontAngle0Shape *Angle0Shape
		if front != nil {
			frontAngle0Shape, _ = front.(*Angle0Shape)
		}
		stage.OnAfterAngle0ShapeDeleteCallback.OnAfterDelete(stage, angle0shape, frontAngle0Shape)
	}
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcNormalVectorShapeCreateCallback != nil {
		stage.OnAfterArcNormalVectorShapeCreateCallback.OnAfterCreate(stage, arcnormalvectorshape)
	}
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcNormalVectorShapeUpdateCallback != nil {
		var frontArcNormalVectorShape *ArcNormalVectorShape
		if front != nil {
			frontArcNormalVectorShape, _ = front.(*ArcNormalVectorShape)
		}
		stage.OnAfterArcNormalVectorShapeUpdateCallback.OnAfterUpdate(stage, arcnormalvectorshape, frontArcNormalVectorShape)
	}
}

func (arcnormalvectorshape *ArcNormalVectorShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcNormalVectorShapeDeleteCallback != nil {
		var frontArcNormalVectorShape *ArcNormalVectorShape
		if front != nil {
			frontArcNormalVectorShape, _ = front.(*ArcNormalVectorShape)
		}
		stage.OnAfterArcNormalVectorShapeDeleteCallback.OnAfterDelete(stage, arcnormalvectorshape, frontArcNormalVectorShape)
	}
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcNormalVectorShapeGridCreateCallback != nil {
		stage.OnAfterArcNormalVectorShapeGridCreateCallback.OnAfterCreate(stage, arcnormalvectorshapegrid)
	}
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcNormalVectorShapeGridUpdateCallback != nil {
		var frontArcNormalVectorShapeGrid *ArcNormalVectorShapeGrid
		if front != nil {
			frontArcNormalVectorShapeGrid, _ = front.(*ArcNormalVectorShapeGrid)
		}
		stage.OnAfterArcNormalVectorShapeGridUpdateCallback.OnAfterUpdate(stage, arcnormalvectorshapegrid, frontArcNormalVectorShapeGrid)
	}
}

func (arcnormalvectorshapegrid *ArcNormalVectorShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcNormalVectorShapeGridDeleteCallback != nil {
		var frontArcNormalVectorShapeGrid *ArcNormalVectorShapeGrid
		if front != nil {
			frontArcNormalVectorShapeGrid, _ = front.(*ArcNormalVectorShapeGrid)
		}
		stage.OnAfterArcNormalVectorShapeGridDeleteCallback.OnAfterDelete(stage, arcnormalvectorshapegrid, frontArcNormalVectorShapeGrid)
	}
}

func (axesshape *AxesShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAxesShapeCreateCallback != nil {
		stage.OnAfterAxesShapeCreateCallback.OnAfterCreate(stage, axesshape)
	}
}

func (axesshape *AxesShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAxesShapeUpdateCallback != nil {
		var frontAxesShape *AxesShape
		if front != nil {
			frontAxesShape, _ = front.(*AxesShape)
		}
		stage.OnAfterAxesShapeUpdateCallback.OnAfterUpdate(stage, axesshape, frontAxesShape)
	}
}

func (axesshape *AxesShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAxesShapeDeleteCallback != nil {
		var frontAxesShape *AxesShape
		if front != nil {
			frontAxesShape, _ = front.(*AxesShape)
		}
		stage.OnAfterAxesShapeDeleteCallback.OnAfterDelete(stage, axesshape, frontAxesShape)
	}
}

func (basevectorshape *BaseVectorShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBaseVectorShapeCreateCallback != nil {
		stage.OnAfterBaseVectorShapeCreateCallback.OnAfterCreate(stage, basevectorshape)
	}
}

func (basevectorshape *BaseVectorShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBaseVectorShapeUpdateCallback != nil {
		var frontBaseVectorShape *BaseVectorShape
		if front != nil {
			frontBaseVectorShape, _ = front.(*BaseVectorShape)
		}
		stage.OnAfterBaseVectorShapeUpdateCallback.OnAfterUpdate(stage, basevectorshape, frontBaseVectorShape)
	}
}

func (basevectorshape *BaseVectorShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBaseVectorShapeDeleteCallback != nil {
		var frontBaseVectorShape *BaseVectorShape
		if front != nil {
			frontBaseVectorShape, _ = front.(*BaseVectorShape)
		}
		stage.OnAfterBaseVectorShapeDeleteCallback.OnAfterDelete(stage, basevectorshape, frontBaseVectorShape)
	}
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBaseVectorShapeGridCreateCallback != nil {
		stage.OnAfterBaseVectorShapeGridCreateCallback.OnAfterCreate(stage, basevectorshapegrid)
	}
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBaseVectorShapeGridUpdateCallback != nil {
		var frontBaseVectorShapeGrid *BaseVectorShapeGrid
		if front != nil {
			frontBaseVectorShapeGrid, _ = front.(*BaseVectorShapeGrid)
		}
		stage.OnAfterBaseVectorShapeGridUpdateCallback.OnAfterUpdate(stage, basevectorshapegrid, frontBaseVectorShapeGrid)
	}
}

func (basevectorshapegrid *BaseVectorShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBaseVectorShapeGridDeleteCallback != nil {
		var frontBaseVectorShapeGrid *BaseVectorShapeGrid
		if front != nil {
			frontBaseVectorShapeGrid, _ = front.(*BaseVectorShapeGrid)
		}
		stage.OnAfterBaseVectorShapeGridDeleteCallback.OnAfterDelete(stage, basevectorshapegrid, frontBaseVectorShapeGrid)
	}
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBottomCurvePlane1ShapeCreateCallback != nil {
		stage.OnAfterBottomCurvePlane1ShapeCreateCallback.OnAfterCreate(stage, bottomcurveplane1shape)
	}
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBottomCurvePlane1ShapeUpdateCallback != nil {
		var frontBottomCurvePlane1Shape *BottomCurvePlane1Shape
		if front != nil {
			frontBottomCurvePlane1Shape, _ = front.(*BottomCurvePlane1Shape)
		}
		stage.OnAfterBottomCurvePlane1ShapeUpdateCallback.OnAfterUpdate(stage, bottomcurveplane1shape, frontBottomCurvePlane1Shape)
	}
}

func (bottomcurveplane1shape *BottomCurvePlane1Shape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBottomCurvePlane1ShapeDeleteCallback != nil {
		var frontBottomCurvePlane1Shape *BottomCurvePlane1Shape
		if front != nil {
			frontBottomCurvePlane1Shape, _ = front.(*BottomCurvePlane1Shape)
		}
		stage.OnAfterBottomCurvePlane1ShapeDeleteCallback.OnAfterDelete(stage, bottomcurveplane1shape, frontBottomCurvePlane1Shape)
	}
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBottomCurvePlane2ShapeCreateCallback != nil {
		stage.OnAfterBottomCurvePlane2ShapeCreateCallback.OnAfterCreate(stage, bottomcurveplane2shape)
	}
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBottomCurvePlane2ShapeUpdateCallback != nil {
		var frontBottomCurvePlane2Shape *BottomCurvePlane2Shape
		if front != nil {
			frontBottomCurvePlane2Shape, _ = front.(*BottomCurvePlane2Shape)
		}
		stage.OnAfterBottomCurvePlane2ShapeUpdateCallback.OnAfterUpdate(stage, bottomcurveplane2shape, frontBottomCurvePlane2Shape)
	}
}

func (bottomcurveplane2shape *BottomCurvePlane2Shape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBottomCurvePlane2ShapeDeleteCallback != nil {
		var frontBottomCurvePlane2Shape *BottomCurvePlane2Shape
		if front != nil {
			frontBottomCurvePlane2Shape, _ = front.(*BottomCurvePlane2Shape)
		}
		stage.OnAfterBottomCurvePlane2ShapeDeleteCallback.OnAfterDelete(stage, bottomcurveplane2shape, frontBottomCurvePlane2Shape)
	}
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterChosenP1P2PairShapeCreateCallback != nil {
		stage.OnAfterChosenP1P2PairShapeCreateCallback.OnAfterCreate(stage, chosenp1p2pairshape)
	}
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChosenP1P2PairShapeUpdateCallback != nil {
		var frontChosenP1P2PairShape *ChosenP1P2PairShape
		if front != nil {
			frontChosenP1P2PairShape, _ = front.(*ChosenP1P2PairShape)
		}
		stage.OnAfterChosenP1P2PairShapeUpdateCallback.OnAfterUpdate(stage, chosenp1p2pairshape, frontChosenP1P2PairShape)
	}
}

func (chosenp1p2pairshape *ChosenP1P2PairShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterChosenP1P2PairShapeDeleteCallback != nil {
		var frontChosenP1P2PairShape *ChosenP1P2PairShape
		if front != nil {
			frontChosenP1P2PairShape, _ = front.(*ChosenP1P2PairShape)
		}
		stage.OnAfterChosenP1P2PairShapeDeleteCallback.OnAfterDelete(stage, chosenp1p2pairshape, frontChosenP1P2PairShape)
	}
}

func (circlegridshape *CircleGridShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircleGridShapeCreateCallback != nil {
		stage.OnAfterCircleGridShapeCreateCallback.OnAfterCreate(stage, circlegridshape)
	}
}

func (circlegridshape *CircleGridShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleGridShapeUpdateCallback != nil {
		var frontCircleGridShape *CircleGridShape
		if front != nil {
			frontCircleGridShape, _ = front.(*CircleGridShape)
		}
		stage.OnAfterCircleGridShapeUpdateCallback.OnAfterUpdate(stage, circlegridshape, frontCircleGridShape)
	}
}

func (circlegridshape *CircleGridShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircleGridShapeDeleteCallback != nil {
		var frontCircleGridShape *CircleGridShape
		if front != nil {
			frontCircleGridShape, _ = front.(*CircleGridShape)
		}
		stage.OnAfterCircleGridShapeDeleteCallback.OnAfterDelete(stage, circlegridshape, frontCircleGridShape)
	}
}

func (circumference3dshape *Circumference3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCircumference3DShapeCreateCallback != nil {
		stage.OnAfterCircumference3DShapeCreateCallback.OnAfterCreate(stage, circumference3dshape)
	}
}

func (circumference3dshape *Circumference3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircumference3DShapeUpdateCallback != nil {
		var frontCircumference3DShape *Circumference3DShape
		if front != nil {
			frontCircumference3DShape, _ = front.(*Circumference3DShape)
		}
		stage.OnAfterCircumference3DShapeUpdateCallback.OnAfterUpdate(stage, circumference3dshape, frontCircumference3DShape)
	}
}

func (circumference3dshape *Circumference3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCircumference3DShapeDeleteCallback != nil {
		var frontCircumference3DShape *Circumference3DShape
		if front != nil {
			frontCircumference3DShape, _ = front.(*Circumference3DShape)
		}
		stage.OnAfterCircumference3DShapeDeleteCallback.OnAfterDelete(stage, circumference3dshape, frontCircumference3DShape)
	}
}

func (clock2ddiagram *Clock2DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterClock2DDiagramCreateCallback != nil {
		stage.OnAfterClock2DDiagramCreateCallback.OnAfterCreate(stage, clock2ddiagram)
	}
}

func (clock2ddiagram *Clock2DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClock2DDiagramUpdateCallback != nil {
		var frontClock2DDiagram *Clock2DDiagram
		if front != nil {
			frontClock2DDiagram, _ = front.(*Clock2DDiagram)
		}
		stage.OnAfterClock2DDiagramUpdateCallback.OnAfterUpdate(stage, clock2ddiagram, frontClock2DDiagram)
	}
}

func (clock2ddiagram *Clock2DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClock2DDiagramDeleteCallback != nil {
		var frontClock2DDiagram *Clock2DDiagram
		if front != nil {
			frontClock2DDiagram, _ = front.(*Clock2DDiagram)
		}
		stage.OnAfterClock2DDiagramDeleteCallback.OnAfterDelete(stage, clock2ddiagram, frontClock2DDiagram)
	}
}

func (clock3ddiagram *Clock3DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterClock3DDiagramCreateCallback != nil {
		stage.OnAfterClock3DDiagramCreateCallback.OnAfterCreate(stage, clock3ddiagram)
	}
}

func (clock3ddiagram *Clock3DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClock3DDiagramUpdateCallback != nil {
		var frontClock3DDiagram *Clock3DDiagram
		if front != nil {
			frontClock3DDiagram, _ = front.(*Clock3DDiagram)
		}
		stage.OnAfterClock3DDiagramUpdateCallback.OnAfterUpdate(stage, clock3ddiagram, frontClock3DDiagram)
	}
}

func (clock3ddiagram *Clock3DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClock3DDiagramDeleteCallback != nil {
		var frontClock3DDiagram *Clock3DDiagram
		if front != nil {
			frontClock3DDiagram, _ = front.(*Clock3DDiagram)
		}
		stage.OnAfterClock3DDiagramDeleteCallback.OnAfterDelete(stage, clock3ddiagram, frontClock3DDiagram)
	}
}

func (clockabstract *ClockAbstract) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterClockAbstractCreateCallback != nil {
		stage.OnAfterClockAbstractCreateCallback.OnAfterCreate(stage, clockabstract)
	}
}

func (clockabstract *ClockAbstract) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClockAbstractUpdateCallback != nil {
		var frontClockAbstract *ClockAbstract
		if front != nil {
			frontClockAbstract, _ = front.(*ClockAbstract)
		}
		stage.OnAfterClockAbstractUpdateCallback.OnAfterUpdate(stage, clockabstract, frontClockAbstract)
	}
}

func (clockabstract *ClockAbstract) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClockAbstractDeleteCallback != nil {
		var frontClockAbstract *ClockAbstract
		if front != nil {
			frontClockAbstract, _ = front.(*ClockAbstract)
		}
		stage.OnAfterClockAbstractDeleteCallback.OnAfterDelete(stage, clockabstract, frontClockAbstract)
	}
}

func (clocktopcurveshape *ClockTopCurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterClockTopCurveShapeCreateCallback != nil {
		stage.OnAfterClockTopCurveShapeCreateCallback.OnAfterCreate(stage, clocktopcurveshape)
	}
}

func (clocktopcurveshape *ClockTopCurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClockTopCurveShapeUpdateCallback != nil {
		var frontClockTopCurveShape *ClockTopCurveShape
		if front != nil {
			frontClockTopCurveShape, _ = front.(*ClockTopCurveShape)
		}
		stage.OnAfterClockTopCurveShapeUpdateCallback.OnAfterUpdate(stage, clocktopcurveshape, frontClockTopCurveShape)
	}
}

func (clocktopcurveshape *ClockTopCurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterClockTopCurveShapeDeleteCallback != nil {
		var frontClockTopCurveShape *ClockTopCurveShape
		if front != nil {
			frontClockTopCurveShape, _ = front.(*ClockTopCurveShape)
		}
		stage.OnAfterClockTopCurveShapeDeleteCallback.OnAfterDelete(stage, clocktopcurveshape, frontClockTopCurveShape)
	}
}

func (cutline3dshape *CutLine3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCutLine3DShapeCreateCallback != nil {
		stage.OnAfterCutLine3DShapeCreateCallback.OnAfterCreate(stage, cutline3dshape)
	}
}

func (cutline3dshape *CutLine3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCutLine3DShapeUpdateCallback != nil {
		var frontCutLine3DShape *CutLine3DShape
		if front != nil {
			frontCutLine3DShape, _ = front.(*CutLine3DShape)
		}
		stage.OnAfterCutLine3DShapeUpdateCallback.OnAfterUpdate(stage, cutline3dshape, frontCutLine3DShape)
	}
}

func (cutline3dshape *CutLine3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCutLine3DShapeDeleteCallback != nil {
		var frontCutLine3DShape *CutLine3DShape
		if front != nil {
			frontCutLine3DShape, _ = front.(*CutLine3DShape)
		}
		stage.OnAfterCutLine3DShapeDeleteCallback.OnAfterDelete(stage, cutline3dshape, frontCutLine3DShape)
	}
}

func (endarcshape *EndArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEndArcShapeCreateCallback != nil {
		stage.OnAfterEndArcShapeCreateCallback.OnAfterCreate(stage, endarcshape)
	}
}

func (endarcshape *EndArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndArcShapeUpdateCallback != nil {
		var frontEndArcShape *EndArcShape
		if front != nil {
			frontEndArcShape, _ = front.(*EndArcShape)
		}
		stage.OnAfterEndArcShapeUpdateCallback.OnAfterUpdate(stage, endarcshape, frontEndArcShape)
	}
}

func (endarcshape *EndArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndArcShapeDeleteCallback != nil {
		var frontEndArcShape *EndArcShape
		if front != nil {
			frontEndArcShape, _ = front.(*EndArcShape)
		}
		stage.OnAfterEndArcShapeDeleteCallback.OnAfterDelete(stage, endarcshape, frontEndArcShape)
	}
}

func (endarcshapegrid *EndArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEndArcShapeGridCreateCallback != nil {
		stage.OnAfterEndArcShapeGridCreateCallback.OnAfterCreate(stage, endarcshapegrid)
	}
}

func (endarcshapegrid *EndArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndArcShapeGridUpdateCallback != nil {
		var frontEndArcShapeGrid *EndArcShapeGrid
		if front != nil {
			frontEndArcShapeGrid, _ = front.(*EndArcShapeGrid)
		}
		stage.OnAfterEndArcShapeGridUpdateCallback.OnAfterUpdate(stage, endarcshapegrid, frontEndArcShapeGrid)
	}
}

func (endarcshapegrid *EndArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndArcShapeGridDeleteCallback != nil {
		var frontEndArcShapeGrid *EndArcShapeGrid
		if front != nil {
			frontEndArcShapeGrid, _ = front.(*EndArcShapeGrid)
		}
		stage.OnAfterEndArcShapeGridDeleteCallback.OnAfterDelete(stage, endarcshapegrid, frontEndArcShapeGrid)
	}
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEndHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, endhalfwayarcshape)
	}
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndHalfwayArcShapeUpdateCallback != nil {
		var frontEndHalfwayArcShape *EndHalfwayArcShape
		if front != nil {
			frontEndHalfwayArcShape, _ = front.(*EndHalfwayArcShape)
		}
		stage.OnAfterEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, endhalfwayarcshape, frontEndHalfwayArcShape)
	}
}

func (endhalfwayarcshape *EndHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndHalfwayArcShapeDeleteCallback != nil {
		var frontEndHalfwayArcShape *EndHalfwayArcShape
		if front != nil {
			frontEndHalfwayArcShape, _ = front.(*EndHalfwayArcShape)
		}
		stage.OnAfterEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, endhalfwayarcshape, frontEndHalfwayArcShape)
	}
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEndHalfwayArcShapeGridCreateCallback != nil {
		stage.OnAfterEndHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, endhalfwayarcshapegrid)
	}
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndHalfwayArcShapeGridUpdateCallback != nil {
		var frontEndHalfwayArcShapeGrid *EndHalfwayArcShapeGrid
		if front != nil {
			frontEndHalfwayArcShapeGrid, _ = front.(*EndHalfwayArcShapeGrid)
		}
		stage.OnAfterEndHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, endhalfwayarcshapegrid, frontEndHalfwayArcShapeGrid)
	}
}

func (endhalfwayarcshapegrid *EndHalfwayArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEndHalfwayArcShapeGridDeleteCallback != nil {
		var frontEndHalfwayArcShapeGrid *EndHalfwayArcShapeGrid
		if front != nil {
			frontEndHalfwayArcShapeGrid, _ = front.(*EndHalfwayArcShapeGrid)
		}
		stage.OnAfterEndHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, endhalfwayarcshapegrid, frontEndHalfwayArcShapeGrid)
	}
}

func (explanationtextshape *ExplanationTextShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterExplanationTextShapeCreateCallback != nil {
		stage.OnAfterExplanationTextShapeCreateCallback.OnAfterCreate(stage, explanationtextshape)
	}
}

func (explanationtextshape *ExplanationTextShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExplanationTextShapeUpdateCallback != nil {
		var frontExplanationTextShape *ExplanationTextShape
		if front != nil {
			frontExplanationTextShape, _ = front.(*ExplanationTextShape)
		}
		stage.OnAfterExplanationTextShapeUpdateCallback.OnAfterUpdate(stage, explanationtextshape, frontExplanationTextShape)
	}
}

func (explanationtextshape *ExplanationTextShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExplanationTextShapeDeleteCallback != nil {
		var frontExplanationTextShape *ExplanationTextShape
		if front != nil {
			frontExplanationTextShape, _ = front.(*ExplanationTextShape)
		}
		stage.OnAfterExplanationTextShapeDeleteCallback.OnAfterDelete(stage, explanationtextshape, frontExplanationTextShape)
	}
}

func (eye3dshape *Eye3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEye3DShapeCreateCallback != nil {
		stage.OnAfterEye3DShapeCreateCallback.OnAfterCreate(stage, eye3dshape)
	}
}

func (eye3dshape *Eye3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEye3DShapeUpdateCallback != nil {
		var frontEye3DShape *Eye3DShape
		if front != nil {
			frontEye3DShape, _ = front.(*Eye3DShape)
		}
		stage.OnAfterEye3DShapeUpdateCallback.OnAfterUpdate(stage, eye3dshape, frontEye3DShape)
	}
}

func (eye3dshape *Eye3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEye3DShapeDeleteCallback != nil {
		var frontEye3DShape *Eye3DShape
		if front != nil {
			frontEye3DShape, _ = front.(*Eye3DShape)
		}
		stage.OnAfterEye3DShapeDeleteCallback.OnAfterDelete(stage, eye3dshape, frontEye3DShape)
	}
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEyeCornersSampledPoints3DShapeCreateCallback != nil {
		stage.OnAfterEyeCornersSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, eyecornerssampledpoints3dshape)
	}
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeCornersSampledPoints3DShapeUpdateCallback != nil {
		var frontEyeCornersSampledPoints3DShape *EyeCornersSampledPoints3DShape
		if front != nil {
			frontEyeCornersSampledPoints3DShape, _ = front.(*EyeCornersSampledPoints3DShape)
		}
		stage.OnAfterEyeCornersSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, eyecornerssampledpoints3dshape, frontEyeCornersSampledPoints3DShape)
	}
}

func (eyecornerssampledpoints3dshape *EyeCornersSampledPoints3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeCornersSampledPoints3DShapeDeleteCallback != nil {
		var frontEyeCornersSampledPoints3DShape *EyeCornersSampledPoints3DShape
		if front != nil {
			frontEyeCornersSampledPoints3DShape, _ = front.(*EyeCornersSampledPoints3DShape)
		}
		stage.OnAfterEyeCornersSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, eyecornerssampledpoints3dshape, frontEyeCornersSampledPoints3DShape)
	}
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEyeSampledPoints3DShapeCreateCallback != nil {
		stage.OnAfterEyeSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, eyesampledpoints3dshape)
	}
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeSampledPoints3DShapeUpdateCallback != nil {
		var frontEyeSampledPoints3DShape *EyeSampledPoints3DShape
		if front != nil {
			frontEyeSampledPoints3DShape, _ = front.(*EyeSampledPoints3DShape)
		}
		stage.OnAfterEyeSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, eyesampledpoints3dshape, frontEyeSampledPoints3DShape)
	}
}

func (eyesampledpoints3dshape *EyeSampledPoints3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeSampledPoints3DShapeDeleteCallback != nil {
		var frontEyeSampledPoints3DShape *EyeSampledPoints3DShape
		if front != nil {
			frontEyeSampledPoints3DShape, _ = front.(*EyeSampledPoints3DShape)
		}
		stage.OnAfterEyeSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, eyesampledpoints3dshape, frontEyeSampledPoints3DShape)
	}
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEyeSeatBottomCurveShapeCreateCallback != nil {
		stage.OnAfterEyeSeatBottomCurveShapeCreateCallback.OnAfterCreate(stage, eyeseatbottomcurveshape)
	}
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeSeatBottomCurveShapeUpdateCallback != nil {
		var frontEyeSeatBottomCurveShape *EyeSeatBottomCurveShape
		if front != nil {
			frontEyeSeatBottomCurveShape, _ = front.(*EyeSeatBottomCurveShape)
		}
		stage.OnAfterEyeSeatBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, eyeseatbottomcurveshape, frontEyeSeatBottomCurveShape)
	}
}

func (eyeseatbottomcurveshape *EyeSeatBottomCurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeSeatBottomCurveShapeDeleteCallback != nil {
		var frontEyeSeatBottomCurveShape *EyeSeatBottomCurveShape
		if front != nil {
			frontEyeSeatBottomCurveShape, _ = front.(*EyeSeatBottomCurveShape)
		}
		stage.OnAfterEyeSeatBottomCurveShapeDeleteCallback.OnAfterDelete(stage, eyeseatbottomcurveshape, frontEyeSeatBottomCurveShape)
	}
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEyeStoolBottomCurveShapeCreateCallback != nil {
		stage.OnAfterEyeStoolBottomCurveShapeCreateCallback.OnAfterCreate(stage, eyestoolbottomcurveshape)
	}
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeStoolBottomCurveShapeUpdateCallback != nil {
		var frontEyeStoolBottomCurveShape *EyeStoolBottomCurveShape
		if front != nil {
			frontEyeStoolBottomCurveShape, _ = front.(*EyeStoolBottomCurveShape)
		}
		stage.OnAfterEyeStoolBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, eyestoolbottomcurveshape, frontEyeStoolBottomCurveShape)
	}
}

func (eyestoolbottomcurveshape *EyeStoolBottomCurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeStoolBottomCurveShapeDeleteCallback != nil {
		var frontEyeStoolBottomCurveShape *EyeStoolBottomCurveShape
		if front != nil {
			frontEyeStoolBottomCurveShape, _ = front.(*EyeStoolBottomCurveShape)
		}
		stage.OnAfterEyeStoolBottomCurveShapeDeleteCallback.OnAfterDelete(stage, eyestoolbottomcurveshape, frontEyeStoolBottomCurveShape)
	}
}

func (eyevolume3dshape *EyeVolume3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterEyeVolume3DShapeCreateCallback != nil {
		stage.OnAfterEyeVolume3DShapeCreateCallback.OnAfterCreate(stage, eyevolume3dshape)
	}
}

func (eyevolume3dshape *EyeVolume3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeVolume3DShapeUpdateCallback != nil {
		var frontEyeVolume3DShape *EyeVolume3DShape
		if front != nil {
			frontEyeVolume3DShape, _ = front.(*EyeVolume3DShape)
		}
		stage.OnAfterEyeVolume3DShapeUpdateCallback.OnAfterUpdate(stage, eyevolume3dshape, frontEyeVolume3DShape)
	}
}

func (eyevolume3dshape *EyeVolume3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterEyeVolume3DShapeDeleteCallback != nil {
		var frontEyeVolume3DShape *EyeVolume3DShape
		if front != nil {
			frontEyeVolume3DShape, _ = front.(*EyeVolume3DShape)
		}
		stage.OnAfterEyeVolume3DShapeDeleteCallback.OnAfterDelete(stage, eyevolume3dshape, frontEyeVolume3DShape)
	}
}

func (gridpathshape *GridPathShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGridPathShapeCreateCallback != nil {
		stage.OnAfterGridPathShapeCreateCallback.OnAfterCreate(stage, gridpathshape)
	}
}

func (gridpathshape *GridPathShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridPathShapeUpdateCallback != nil {
		var frontGridPathShape *GridPathShape
		if front != nil {
			frontGridPathShape, _ = front.(*GridPathShape)
		}
		stage.OnAfterGridPathShapeUpdateCallback.OnAfterUpdate(stage, gridpathshape, frontGridPathShape)
	}
}

func (gridpathshape *GridPathShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGridPathShapeDeleteCallback != nil {
		var frontGridPathShape *GridPathShape
		if front != nil {
			frontGridPathShape, _ = front.(*GridPathShape)
		}
		stage.OnAfterGridPathShapeDeleteCallback.OnAfterDelete(stage, gridpathshape, frontGridPathShape)
	}
}

func (growthcurve2d *GrowthCurve2D) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGrowthCurve2DCreateCallback != nil {
		stage.OnAfterGrowthCurve2DCreateCallback.OnAfterCreate(stage, growthcurve2d)
	}
}

func (growthcurve2d *GrowthCurve2D) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DUpdateCallback != nil {
		var frontGrowthCurve2D *GrowthCurve2D
		if front != nil {
			frontGrowthCurve2D, _ = front.(*GrowthCurve2D)
		}
		stage.OnAfterGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, growthcurve2d, frontGrowthCurve2D)
	}
}

func (growthcurve2d *GrowthCurve2D) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DDeleteCallback != nil {
		var frontGrowthCurve2D *GrowthCurve2D
		if front != nil {
			frontGrowthCurve2D, _ = front.(*GrowthCurve2D)
		}
		stage.OnAfterGrowthCurve2DDeleteCallback.OnAfterDelete(stage, growthcurve2d, frontGrowthCurve2D)
	}
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGrowthCurve2DRibbonCreateCallback != nil {
		stage.OnAfterGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, growthcurve2dribbon)
	}
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DRibbonUpdateCallback != nil {
		var frontGrowthCurve2DRibbon *GrowthCurve2DRibbon
		if front != nil {
			frontGrowthCurve2DRibbon, _ = front.(*GrowthCurve2DRibbon)
		}
		stage.OnAfterGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, growthcurve2dribbon, frontGrowthCurve2DRibbon)
	}
}

func (growthcurve2dribbon *GrowthCurve2DRibbon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DRibbonDeleteCallback != nil {
		var frontGrowthCurve2DRibbon *GrowthCurve2DRibbon
		if front != nil {
			frontGrowthCurve2DRibbon, _ = front.(*GrowthCurve2DRibbon)
		}
		stage.OnAfterGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, growthcurve2dribbon, frontGrowthCurve2DRibbon)
	}
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGrowthCurve2DRibbonEndShapeCreateCallback != nil {
		stage.OnAfterGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, growthcurve2dribbonendshape)
	}
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
		var frontGrowthCurve2DRibbonEndShape *GrowthCurve2DRibbonEndShape
		if front != nil {
			frontGrowthCurve2DRibbonEndShape, _ = front.(*GrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, growthcurve2dribbonendshape, frontGrowthCurve2DRibbonEndShape)
	}
}

func (growthcurve2dribbonendshape *GrowthCurve2DRibbonEndShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
		var frontGrowthCurve2DRibbonEndShape *GrowthCurve2DRibbonEndShape
		if front != nil {
			frontGrowthCurve2DRibbonEndShape, _ = front.(*GrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, growthcurve2dribbonendshape, frontGrowthCurve2DRibbonEndShape)
	}
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGrowthCurve2DRibbonStartShapeCreateCallback != nil {
		stage.OnAfterGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, growthcurve2dribbonstartshape)
	}
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
		var frontGrowthCurve2DRibbonStartShape *GrowthCurve2DRibbonStartShape
		if front != nil {
			frontGrowthCurve2DRibbonStartShape, _ = front.(*GrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, growthcurve2dribbonstartshape, frontGrowthCurve2DRibbonStartShape)
	}
}

func (growthcurve2dribbonstartshape *GrowthCurve2DRibbonStartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
		var frontGrowthCurve2DRibbonStartShape *GrowthCurve2DRibbonStartShape
		if front != nil {
			frontGrowthCurve2DRibbonStartShape, _ = front.(*GrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, growthcurve2dribbonstartshape, frontGrowthCurve2DRibbonStartShape)
	}
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGrowthCurveRhombusGridShapeCreateCallback != nil {
		stage.OnAfterGrowthCurveRhombusGridShapeCreateCallback.OnAfterCreate(stage, growthcurverhombusgridshape)
	}
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurveRhombusGridShapeUpdateCallback != nil {
		var frontGrowthCurveRhombusGridShape *GrowthCurveRhombusGridShape
		if front != nil {
			frontGrowthCurveRhombusGridShape, _ = front.(*GrowthCurveRhombusGridShape)
		}
		stage.OnAfterGrowthCurveRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, growthcurverhombusgridshape, frontGrowthCurveRhombusGridShape)
	}
}

func (growthcurverhombusgridshape *GrowthCurveRhombusGridShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurveRhombusGridShapeDeleteCallback != nil {
		var frontGrowthCurveRhombusGridShape *GrowthCurveRhombusGridShape
		if front != nil {
			frontGrowthCurveRhombusGridShape, _ = front.(*GrowthCurveRhombusGridShape)
		}
		stage.OnAfterGrowthCurveRhombusGridShapeDeleteCallback.OnAfterDelete(stage, growthcurverhombusgridshape, frontGrowthCurveRhombusGridShape)
	}
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGrowthCurveRhombusShapeCreateCallback != nil {
		stage.OnAfterGrowthCurveRhombusShapeCreateCallback.OnAfterCreate(stage, growthcurverhombusshape)
	}
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurveRhombusShapeUpdateCallback != nil {
		var frontGrowthCurveRhombusShape *GrowthCurveRhombusShape
		if front != nil {
			frontGrowthCurveRhombusShape, _ = front.(*GrowthCurveRhombusShape)
		}
		stage.OnAfterGrowthCurveRhombusShapeUpdateCallback.OnAfterUpdate(stage, growthcurverhombusshape, frontGrowthCurveRhombusShape)
	}
}

func (growthcurverhombusshape *GrowthCurveRhombusShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthCurveRhombusShapeDeleteCallback != nil {
		var frontGrowthCurveRhombusShape *GrowthCurveRhombusShape
		if front != nil {
			frontGrowthCurveRhombusShape, _ = front.(*GrowthCurveRhombusShape)
		}
		stage.OnAfterGrowthCurveRhombusShapeDeleteCallback.OnAfterDelete(stage, growthcurverhombusshape, frontGrowthCurveRhombusShape)
	}
}

func (growthvectorshape *GrowthVectorShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGrowthVectorShapeCreateCallback != nil {
		stage.OnAfterGrowthVectorShapeCreateCallback.OnAfterCreate(stage, growthvectorshape)
	}
}

func (growthvectorshape *GrowthVectorShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthVectorShapeUpdateCallback != nil {
		var frontGrowthVectorShape *GrowthVectorShape
		if front != nil {
			frontGrowthVectorShape, _ = front.(*GrowthVectorShape)
		}
		stage.OnAfterGrowthVectorShapeUpdateCallback.OnAfterUpdate(stage, growthvectorshape, frontGrowthVectorShape)
	}
}

func (growthvectorshape *GrowthVectorShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGrowthVectorShapeDeleteCallback != nil {
		var frontGrowthVectorShape *GrowthVectorShape
		if front != nil {
			frontGrowthVectorShape, _ = front.(*GrowthVectorShape)
		}
		stage.OnAfterGrowthVectorShapeDeleteCallback.OnAfterDelete(stage, growthvectorshape, frontGrowthVectorShape)
	}
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInitialRhombusGridShapeCreateCallback != nil {
		stage.OnAfterInitialRhombusGridShapeCreateCallback.OnAfterCreate(stage, initialrhombusgridshape)
	}
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialRhombusGridShapeUpdateCallback != nil {
		var frontInitialRhombusGridShape *InitialRhombusGridShape
		if front != nil {
			frontInitialRhombusGridShape, _ = front.(*InitialRhombusGridShape)
		}
		stage.OnAfterInitialRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, initialrhombusgridshape, frontInitialRhombusGridShape)
	}
}

func (initialrhombusgridshape *InitialRhombusGridShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialRhombusGridShapeDeleteCallback != nil {
		var frontInitialRhombusGridShape *InitialRhombusGridShape
		if front != nil {
			frontInitialRhombusGridShape, _ = front.(*InitialRhombusGridShape)
		}
		stage.OnAfterInitialRhombusGridShapeDeleteCallback.OnAfterDelete(stage, initialrhombusgridshape, frontInitialRhombusGridShape)
	}
}

func (initialrhombusshape *InitialRhombusShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInitialRhombusShapeCreateCallback != nil {
		stage.OnAfterInitialRhombusShapeCreateCallback.OnAfterCreate(stage, initialrhombusshape)
	}
}

func (initialrhombusshape *InitialRhombusShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialRhombusShapeUpdateCallback != nil {
		var frontInitialRhombusShape *InitialRhombusShape
		if front != nil {
			frontInitialRhombusShape, _ = front.(*InitialRhombusShape)
		}
		stage.OnAfterInitialRhombusShapeUpdateCallback.OnAfterUpdate(stage, initialrhombusshape, frontInitialRhombusShape)
	}
}

func (initialrhombusshape *InitialRhombusShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInitialRhombusShapeDeleteCallback != nil {
		var frontInitialRhombusShape *InitialRhombusShape
		if front != nil {
			frontInitialRhombusShape, _ = front.(*InitialRhombusShape)
		}
		stage.OnAfterInitialRhombusShapeDeleteCallback.OnAfterDelete(stage, initialrhombusshape, frontInitialRhombusShape)
	}
}

func (key3dshape *Key3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKey3DShapeCreateCallback != nil {
		stage.OnAfterKey3DShapeCreateCallback.OnAfterCreate(stage, key3dshape)
	}
}

func (key3dshape *Key3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKey3DShapeUpdateCallback != nil {
		var frontKey3DShape *Key3DShape
		if front != nil {
			frontKey3DShape, _ = front.(*Key3DShape)
		}
		stage.OnAfterKey3DShapeUpdateCallback.OnAfterUpdate(stage, key3dshape, frontKey3DShape)
	}
}

func (key3dshape *Key3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKey3DShapeDeleteCallback != nil {
		var frontKey3DShape *Key3DShape
		if front != nil {
			frontKey3DShape, _ = front.(*Key3DShape)
		}
		stage.OnAfterKey3DShapeDeleteCallback.OnAfterDelete(stage, key3dshape, frontKey3DShape)
	}
}

func (keyhole3dshape *KeyHole3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKeyHole3DShapeCreateCallback != nil {
		stage.OnAfterKeyHole3DShapeCreateCallback.OnAfterCreate(stage, keyhole3dshape)
	}
}

func (keyhole3dshape *KeyHole3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyHole3DShapeUpdateCallback != nil {
		var frontKeyHole3DShape *KeyHole3DShape
		if front != nil {
			frontKeyHole3DShape, _ = front.(*KeyHole3DShape)
		}
		stage.OnAfterKeyHole3DShapeUpdateCallback.OnAfterUpdate(stage, keyhole3dshape, frontKeyHole3DShape)
	}
}

func (keyhole3dshape *KeyHole3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyHole3DShapeDeleteCallback != nil {
		var frontKeyHole3DShape *KeyHole3DShape
		if front != nil {
			frontKeyHole3DShape, _ = front.(*KeyHole3DShape)
		}
		stage.OnAfterKeyHole3DShapeDeleteCallback.OnAfterDelete(stage, keyhole3dshape, frontKeyHole3DShape)
	}
}

func (keyholeshape *KeyHoleShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterKeyHoleShapeCreateCallback != nil {
		stage.OnAfterKeyHoleShapeCreateCallback.OnAfterCreate(stage, keyholeshape)
	}
}

func (keyholeshape *KeyHoleShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyHoleShapeUpdateCallback != nil {
		var frontKeyHoleShape *KeyHoleShape
		if front != nil {
			frontKeyHoleShape, _ = front.(*KeyHoleShape)
		}
		stage.OnAfterKeyHoleShapeUpdateCallback.OnAfterUpdate(stage, keyholeshape, frontKeyHoleShape)
	}
}

func (keyholeshape *KeyHoleShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterKeyHoleShapeDeleteCallback != nil {
		var frontKeyHoleShape *KeyHoleShape
		if front != nil {
			frontKeyHoleShape, _ = front.(*KeyHoleShape)
		}
		stage.OnAfterKeyHoleShapeDeleteCallback.OnAfterDelete(stage, keyholeshape, frontKeyHoleShape)
	}
}

func (leaves3dshape *Leaves3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLeaves3DShapeCreateCallback != nil {
		stage.OnAfterLeaves3DShapeCreateCallback.OnAfterCreate(stage, leaves3dshape)
	}
}

func (leaves3dshape *Leaves3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLeaves3DShapeUpdateCallback != nil {
		var frontLeaves3DShape *Leaves3DShape
		if front != nil {
			frontLeaves3DShape, _ = front.(*Leaves3DShape)
		}
		stage.OnAfterLeaves3DShapeUpdateCallback.OnAfterUpdate(stage, leaves3dshape, frontLeaves3DShape)
	}
}

func (leaves3dshape *Leaves3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLeaves3DShapeDeleteCallback != nil {
		var frontLeaves3DShape *Leaves3DShape
		if front != nil {
			frontLeaves3DShape, _ = front.(*Leaves3DShape)
		}
		stage.OnAfterLeaves3DShapeDeleteCallback.OnAfterDelete(stage, leaves3dshape, frontLeaves3DShape)
	}
}

func (library *Library) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLibraryCreateCallback != nil {
		stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, library)
	}
}

func (library *Library) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryUpdateCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, library, frontLibrary)
	}
}

func (library *Library) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryDeleteCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, library, frontLibrary)
	}
}

func (midarcvectorshape *MidArcVectorShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMidArcVectorShapeCreateCallback != nil {
		stage.OnAfterMidArcVectorShapeCreateCallback.OnAfterCreate(stage, midarcvectorshape)
	}
}

func (midarcvectorshape *MidArcVectorShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidArcVectorShapeUpdateCallback != nil {
		var frontMidArcVectorShape *MidArcVectorShape
		if front != nil {
			frontMidArcVectorShape, _ = front.(*MidArcVectorShape)
		}
		stage.OnAfterMidArcVectorShapeUpdateCallback.OnAfterUpdate(stage, midarcvectorshape, frontMidArcVectorShape)
	}
}

func (midarcvectorshape *MidArcVectorShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidArcVectorShapeDeleteCallback != nil {
		var frontMidArcVectorShape *MidArcVectorShape
		if front != nil {
			frontMidArcVectorShape, _ = front.(*MidArcVectorShape)
		}
		stage.OnAfterMidArcVectorShapeDeleteCallback.OnAfterDelete(stage, midarcvectorshape, frontMidArcVectorShape)
	}
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMidArcVectorShapeGridCreateCallback != nil {
		stage.OnAfterMidArcVectorShapeGridCreateCallback.OnAfterCreate(stage, midarcvectorshapegrid)
	}
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidArcVectorShapeGridUpdateCallback != nil {
		var frontMidArcVectorShapeGrid *MidArcVectorShapeGrid
		if front != nil {
			frontMidArcVectorShapeGrid, _ = front.(*MidArcVectorShapeGrid)
		}
		stage.OnAfterMidArcVectorShapeGridUpdateCallback.OnAfterUpdate(stage, midarcvectorshapegrid, frontMidArcVectorShapeGrid)
	}
}

func (midarcvectorshapegrid *MidArcVectorShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMidArcVectorShapeGridDeleteCallback != nil {
		var frontMidArcVectorShapeGrid *MidArcVectorShapeGrid
		if front != nil {
			frontMidArcVectorShapeGrid, _ = front.(*MidArcVectorShapeGrid)
		}
		stage.OnAfterMidArcVectorShapeGridDeleteCallback.OnAfterDelete(stage, midarcvectorshapegrid, frontMidArcVectorShapeGrid)
	}
}

func (musicabstract *MusicAbstract) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMusicAbstractCreateCallback != nil {
		stage.OnAfterMusicAbstractCreateCallback.OnAfterCreate(stage, musicabstract)
	}
}

func (musicabstract *MusicAbstract) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMusicAbstractUpdateCallback != nil {
		var frontMusicAbstract *MusicAbstract
		if front != nil {
			frontMusicAbstract, _ = front.(*MusicAbstract)
		}
		stage.OnAfterMusicAbstractUpdateCallback.OnAfterUpdate(stage, musicabstract, frontMusicAbstract)
	}
}

func (musicabstract *MusicAbstract) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMusicAbstractDeleteCallback != nil {
		var frontMusicAbstract *MusicAbstract
		if front != nil {
			frontMusicAbstract, _ = front.(*MusicAbstract)
		}
		stage.OnAfterMusicAbstractDeleteCallback.OnAfterDelete(stage, musicabstract, frontMusicAbstract)
	}
}

func (originalpoints3dshape *OriginalPoints3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterOriginalPoints3DShapeCreateCallback != nil {
		stage.OnAfterOriginalPoints3DShapeCreateCallback.OnAfterCreate(stage, originalpoints3dshape)
	}
}

func (originalpoints3dshape *OriginalPoints3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOriginalPoints3DShapeUpdateCallback != nil {
		var frontOriginalPoints3DShape *OriginalPoints3DShape
		if front != nil {
			frontOriginalPoints3DShape, _ = front.(*OriginalPoints3DShape)
		}
		stage.OnAfterOriginalPoints3DShapeUpdateCallback.OnAfterUpdate(stage, originalpoints3dshape, frontOriginalPoints3DShape)
	}
}

func (originalpoints3dshape *OriginalPoints3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterOriginalPoints3DShapeDeleteCallback != nil {
		var frontOriginalPoints3DShape *OriginalPoints3DShape
		if front != nil {
			frontOriginalPoints3DShape, _ = front.(*OriginalPoints3DShape)
		}
		stage.OnAfterOriginalPoints3DShapeDeleteCallback.OnAfterDelete(stage, originalpoints3dshape, frontOriginalPoints3DShape)
	}
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParastichyMCurves3DShapeCreateCallback != nil {
		stage.OnAfterParastichyMCurves3DShapeCreateCallback.OnAfterCreate(stage, parastichymcurves3dshape)
	}
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParastichyMCurves3DShapeUpdateCallback != nil {
		var frontParastichyMCurves3DShape *ParastichyMCurves3DShape
		if front != nil {
			frontParastichyMCurves3DShape, _ = front.(*ParastichyMCurves3DShape)
		}
		stage.OnAfterParastichyMCurves3DShapeUpdateCallback.OnAfterUpdate(stage, parastichymcurves3dshape, frontParastichyMCurves3DShape)
	}
}

func (parastichymcurves3dshape *ParastichyMCurves3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParastichyMCurves3DShapeDeleteCallback != nil {
		var frontParastichyMCurves3DShape *ParastichyMCurves3DShape
		if front != nil {
			frontParastichyMCurves3DShape, _ = front.(*ParastichyMCurves3DShape)
		}
		stage.OnAfterParastichyMCurves3DShapeDeleteCallback.OnAfterDelete(stage, parastichymcurves3dshape, frontParastichyMCurves3DShape)
	}
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterParastichyNCurves3DShapeCreateCallback != nil {
		stage.OnAfterParastichyNCurves3DShapeCreateCallback.OnAfterCreate(stage, parastichyncurves3dshape)
	}
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParastichyNCurves3DShapeUpdateCallback != nil {
		var frontParastichyNCurves3DShape *ParastichyNCurves3DShape
		if front != nil {
			frontParastichyNCurves3DShape, _ = front.(*ParastichyNCurves3DShape)
		}
		stage.OnAfterParastichyNCurves3DShapeUpdateCallback.OnAfterUpdate(stage, parastichyncurves3dshape, frontParastichyNCurves3DShape)
	}
}

func (parastichyncurves3dshape *ParastichyNCurves3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterParastichyNCurves3DShapeDeleteCallback != nil {
		var frontParastichyNCurves3DShape *ParastichyNCurves3DShape
		if front != nil {
			frontParastichyNCurves3DShape, _ = front.(*ParastichyNCurves3DShape)
		}
		stage.OnAfterParastichyNCurves3DShapeDeleteCallback.OnAfterDelete(stage, parastichyncurves3dshape, frontParastichyNCurves3DShape)
	}
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dribbon)
	}
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DRibbon *PartiallyGrowthCurve2DRibbon
		if front != nil {
			frontPartiallyGrowthCurve2DRibbon, _ = front.(*PartiallyGrowthCurve2DRibbon)
		}
		stage.OnAfterPartiallyGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dribbon, frontPartiallyGrowthCurve2DRibbon)
	}
}

func (partiallygrowthcurve2dribbon *PartiallyGrowthCurve2DRibbon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DRibbon *PartiallyGrowthCurve2DRibbon
		if front != nil {
			frontPartiallyGrowthCurve2DRibbon, _ = front.(*PartiallyGrowthCurve2DRibbon)
		}
		stage.OnAfterPartiallyGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dribbon, frontPartiallyGrowthCurve2DRibbon)
	}
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dribbonendshape)
	}
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DRibbonEndShape *PartiallyGrowthCurve2DRibbonEndShape
		if front != nil {
			frontPartiallyGrowthCurve2DRibbonEndShape, _ = front.(*PartiallyGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dribbonendshape, frontPartiallyGrowthCurve2DRibbonEndShape)
	}
}

func (partiallygrowthcurve2dribbonendshape *PartiallyGrowthCurve2DRibbonEndShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DRibbonEndShape *PartiallyGrowthCurve2DRibbonEndShape
		if front != nil {
			frontPartiallyGrowthCurve2DRibbonEndShape, _ = front.(*PartiallyGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dribbonendshape, frontPartiallyGrowthCurve2DRibbonEndShape)
	}
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dribbonstartshape)
	}
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DRibbonStartShape *PartiallyGrowthCurve2DRibbonStartShape
		if front != nil {
			frontPartiallyGrowthCurve2DRibbonStartShape, _ = front.(*PartiallyGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dribbonstartshape, frontPartiallyGrowthCurve2DRibbonStartShape)
	}
}

func (partiallygrowthcurve2dribbonstartshape *PartiallyGrowthCurve2DRibbonStartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DRibbonStartShape *PartiallyGrowthCurve2DRibbonStartShape
		if front != nil {
			frontPartiallyGrowthCurve2DRibbonStartShape, _ = front.(*PartiallyGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dribbonstartshape, frontPartiallyGrowthCurve2DRibbonStartShape)
	}
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectory)
	}
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectory *PartiallyGrowthCurve2DTrajectory
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectory, _ = front.(*PartiallyGrowthCurve2DTrajectory)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectory, frontPartiallyGrowthCurve2DTrajectory)
	}
}

func (partiallygrowthcurve2dtrajectory *PartiallyGrowthCurve2DTrajectory) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectory *PartiallyGrowthCurve2DTrajectory
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectory, _ = front.(*PartiallyGrowthCurve2DTrajectory)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectory, frontPartiallyGrowthCurve2DTrajectory)
	}
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectoryp1curveshape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1CurveShape *PartiallyGrowthCurve2DTrajectoryP1CurveShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1CurveShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectoryp1curveshape, frontPartiallyGrowthCurve2DTrajectoryP1CurveShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1curveshape *PartiallyGrowthCurve2DTrajectoryP1CurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1CurveShape *PartiallyGrowthCurve2DTrajectoryP1CurveShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1CurveShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1CurveShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1CurveShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectoryp1curveshape, frontPartiallyGrowthCurve2DTrajectoryP1CurveShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2CreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2CreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectoryp1p2)
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2UpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1P2 *PartiallyGrowthCurve2DTrajectoryP1P2
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1P2, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1P2)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2UpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectoryp1p2, frontPartiallyGrowthCurve2DTrajectoryP1P2)
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2 *PartiallyGrowthCurve2DTrajectoryP1P2) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2DeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1P2 *PartiallyGrowthCurve2DTrajectoryP1P2
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1P2, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1P2)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2DeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectoryp1p2, frontPartiallyGrowthCurve2DTrajectoryP1P2)
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectoryp1p2pairlineshape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectoryp1p2pairlineshape, frontPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1p2pairlineshape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape *PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1P2PairLineShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectoryp1p2pairlineshape, frontPartiallyGrowthCurve2DTrajectoryP1P2PairLineShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectoryp1pointshape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1PointShape *PartiallyGrowthCurve2DTrajectoryP1PointShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1PointShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectoryp1pointshape, frontPartiallyGrowthCurve2DTrajectoryP1PointShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp1pointshape *PartiallyGrowthCurve2DTrajectoryP1PointShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP1PointShape *PartiallyGrowthCurve2DTrajectoryP1PointShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP1PointShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP1PointShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP1PointShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectoryp1pointshape, frontPartiallyGrowthCurve2DTrajectoryP1PointShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectoryp2curveshape)
	}
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP2CurveShape *PartiallyGrowthCurve2DTrajectoryP2CurveShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP2CurveShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectoryp2curveshape, frontPartiallyGrowthCurve2DTrajectoryP2CurveShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp2curveshape *PartiallyGrowthCurve2DTrajectoryP2CurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP2CurveShape *PartiallyGrowthCurve2DTrajectoryP2CurveShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP2CurveShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP2CurveShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2CurveShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectoryp2curveshape, frontPartiallyGrowthCurve2DTrajectoryP2CurveShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectoryp2pointshape)
	}
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP2PointShape *PartiallyGrowthCurve2DTrajectoryP2PointShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP2PointShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectoryp2pointshape, frontPartiallyGrowthCurve2DTrajectoryP2PointShape)
	}
}

func (partiallygrowthcurve2dtrajectoryp2pointshape *PartiallyGrowthCurve2DTrajectoryP2PointShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryP2PointShape *PartiallyGrowthCurve2DTrajectoryP2PointShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryP2PointShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryP2PointShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryP2PointShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectoryp2pointshape, frontPartiallyGrowthCurve2DTrajectoryP2PointShape)
	}
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeCreateCallback != nil {
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeCreateCallback.OnAfterCreate(stage, partiallygrowthcurve2dtrajectoryshape)
	}
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeUpdateCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryShape *PartiallyGrowthCurve2DTrajectoryShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeUpdateCallback.OnAfterUpdate(stage, partiallygrowthcurve2dtrajectoryshape, frontPartiallyGrowthCurve2DTrajectoryShape)
	}
}

func (partiallygrowthcurve2dtrajectoryshape *PartiallyGrowthCurve2DTrajectoryShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeDeleteCallback != nil {
		var frontPartiallyGrowthCurve2DTrajectoryShape *PartiallyGrowthCurve2DTrajectoryShape
		if front != nil {
			frontPartiallyGrowthCurve2DTrajectoryShape, _ = front.(*PartiallyGrowthCurve2DTrajectoryShape)
		}
		stage.OnAfterPartiallyGrowthCurve2DTrajectoryShapeDeleteCallback.OnAfterDelete(stage, partiallygrowthcurve2dtrajectoryshape, frontPartiallyGrowthCurve2DTrajectoryShape)
	}
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyRotatedSeatBottomCurveShapeCreateCallback != nil {
		stage.OnAfterPartiallyRotatedSeatBottomCurveShapeCreateCallback.OnAfterCreate(stage, partiallyrotatedseatbottomcurveshape)
	}
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyRotatedSeatBottomCurveShapeUpdateCallback != nil {
		var frontPartiallyRotatedSeatBottomCurveShape *PartiallyRotatedSeatBottomCurveShape
		if front != nil {
			frontPartiallyRotatedSeatBottomCurveShape, _ = front.(*PartiallyRotatedSeatBottomCurveShape)
		}
		stage.OnAfterPartiallyRotatedSeatBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, partiallyrotatedseatbottomcurveshape, frontPartiallyRotatedSeatBottomCurveShape)
	}
}

func (partiallyrotatedseatbottomcurveshape *PartiallyRotatedSeatBottomCurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyRotatedSeatBottomCurveShapeDeleteCallback != nil {
		var frontPartiallyRotatedSeatBottomCurveShape *PartiallyRotatedSeatBottomCurveShape
		if front != nil {
			frontPartiallyRotatedSeatBottomCurveShape, _ = front.(*PartiallyRotatedSeatBottomCurveShape)
		}
		stage.OnAfterPartiallyRotatedSeatBottomCurveShapeDeleteCallback.OnAfterDelete(stage, partiallyrotatedseatbottomcurveshape, frontPartiallyRotatedSeatBottomCurveShape)
	}
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyRotatedSeatTopCurveShapeCreateCallback != nil {
		stage.OnAfterPartiallyRotatedSeatTopCurveShapeCreateCallback.OnAfterCreate(stage, partiallyrotatedseattopcurveshape)
	}
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyRotatedSeatTopCurveShapeUpdateCallback != nil {
		var frontPartiallyRotatedSeatTopCurveShape *PartiallyRotatedSeatTopCurveShape
		if front != nil {
			frontPartiallyRotatedSeatTopCurveShape, _ = front.(*PartiallyRotatedSeatTopCurveShape)
		}
		stage.OnAfterPartiallyRotatedSeatTopCurveShapeUpdateCallback.OnAfterUpdate(stage, partiallyrotatedseattopcurveshape, frontPartiallyRotatedSeatTopCurveShape)
	}
}

func (partiallyrotatedseattopcurveshape *PartiallyRotatedSeatTopCurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyRotatedSeatTopCurveShapeDeleteCallback != nil {
		var frontPartiallyRotatedSeatTopCurveShape *PartiallyRotatedSeatTopCurveShape
		if front != nil {
			frontPartiallyRotatedSeatTopCurveShape, _ = front.(*PartiallyRotatedSeatTopCurveShape)
		}
		stage.OnAfterPartiallyRotatedSeatTopCurveShapeDeleteCallback.OnAfterDelete(stage, partiallyrotatedseattopcurveshape, frontPartiallyRotatedSeatTopCurveShape)
	}
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartiallyRotatedTorusShapeCreateCallback != nil {
		stage.OnAfterPartiallyRotatedTorusShapeCreateCallback.OnAfterCreate(stage, partiallyrotatedtorusshape)
	}
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyRotatedTorusShapeUpdateCallback != nil {
		var frontPartiallyRotatedTorusShape *PartiallyRotatedTorusShape
		if front != nil {
			frontPartiallyRotatedTorusShape, _ = front.(*PartiallyRotatedTorusShape)
		}
		stage.OnAfterPartiallyRotatedTorusShapeUpdateCallback.OnAfterUpdate(stage, partiallyrotatedtorusshape, frontPartiallyRotatedTorusShape)
	}
}

func (partiallyrotatedtorusshape *PartiallyRotatedTorusShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartiallyRotatedTorusShapeDeleteCallback != nil {
		var frontPartiallyRotatedTorusShape *PartiallyRotatedTorusShape
		if front != nil {
			frontPartiallyRotatedTorusShape, _ = front.(*PartiallyRotatedTorusShape)
		}
		stage.OnAfterPartiallyRotatedTorusShapeDeleteCallback.OnAfterDelete(stage, partiallyrotatedtorusshape, frontPartiallyRotatedTorusShape)
	}
}

func (perpendicularvector *PerpendicularVector) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPerpendicularVectorCreateCallback != nil {
		stage.OnAfterPerpendicularVectorCreateCallback.OnAfterCreate(stage, perpendicularvector)
	}
}

func (perpendicularvector *PerpendicularVector) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorUpdateCallback != nil {
		var frontPerpendicularVector *PerpendicularVector
		if front != nil {
			frontPerpendicularVector, _ = front.(*PerpendicularVector)
		}
		stage.OnAfterPerpendicularVectorUpdateCallback.OnAfterUpdate(stage, perpendicularvector, frontPerpendicularVector)
	}
}

func (perpendicularvector *PerpendicularVector) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorDeleteCallback != nil {
		var frontPerpendicularVector *PerpendicularVector
		if front != nil {
			frontPerpendicularVector, _ = front.(*PerpendicularVector)
		}
		stage.OnAfterPerpendicularVectorDeleteCallback.OnAfterDelete(stage, perpendicularvector, frontPerpendicularVector)
	}
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPerpendicularVectorGridCreateCallback != nil {
		stage.OnAfterPerpendicularVectorGridCreateCallback.OnAfterCreate(stage, perpendicularvectorgrid)
	}
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorGridUpdateCallback != nil {
		var frontPerpendicularVectorGrid *PerpendicularVectorGrid
		if front != nil {
			frontPerpendicularVectorGrid, _ = front.(*PerpendicularVectorGrid)
		}
		stage.OnAfterPerpendicularVectorGridUpdateCallback.OnAfterUpdate(stage, perpendicularvectorgrid, frontPerpendicularVectorGrid)
	}
}

func (perpendicularvectorgrid *PerpendicularVectorGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorGridDeleteCallback != nil {
		var frontPerpendicularVectorGrid *PerpendicularVectorGrid
		if front != nil {
			frontPerpendicularVectorGrid, _ = front.(*PerpendicularVectorGrid)
		}
		stage.OnAfterPerpendicularVectorGridDeleteCallback.OnAfterDelete(stage, perpendicularvectorgrid, frontPerpendicularVectorGrid)
	}
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPerpendicularVectorGridHalfwayCreateCallback != nil {
		stage.OnAfterPerpendicularVectorGridHalfwayCreateCallback.OnAfterCreate(stage, perpendicularvectorgridhalfway)
	}
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorGridHalfwayUpdateCallback != nil {
		var frontPerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway
		if front != nil {
			frontPerpendicularVectorGridHalfway, _ = front.(*PerpendicularVectorGridHalfway)
		}
		stage.OnAfterPerpendicularVectorGridHalfwayUpdateCallback.OnAfterUpdate(stage, perpendicularvectorgridhalfway, frontPerpendicularVectorGridHalfway)
	}
}

func (perpendicularvectorgridhalfway *PerpendicularVectorGridHalfway) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorGridHalfwayDeleteCallback != nil {
		var frontPerpendicularVectorGridHalfway *PerpendicularVectorGridHalfway
		if front != nil {
			frontPerpendicularVectorGridHalfway, _ = front.(*PerpendicularVectorGridHalfway)
		}
		stage.OnAfterPerpendicularVectorGridHalfwayDeleteCallback.OnAfterDelete(stage, perpendicularvectorgridhalfway, frontPerpendicularVectorGridHalfway)
	}
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPerpendicularVectorHalfwayCreateCallback != nil {
		stage.OnAfterPerpendicularVectorHalfwayCreateCallback.OnAfterCreate(stage, perpendicularvectorhalfway)
	}
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorHalfwayUpdateCallback != nil {
		var frontPerpendicularVectorHalfway *PerpendicularVectorHalfway
		if front != nil {
			frontPerpendicularVectorHalfway, _ = front.(*PerpendicularVectorHalfway)
		}
		stage.OnAfterPerpendicularVectorHalfwayUpdateCallback.OnAfterUpdate(stage, perpendicularvectorhalfway, frontPerpendicularVectorHalfway)
	}
}

func (perpendicularvectorhalfway *PerpendicularVectorHalfway) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPerpendicularVectorHalfwayDeleteCallback != nil {
		var frontPerpendicularVectorHalfway *PerpendicularVectorHalfway
		if front != nil {
			frontPerpendicularVectorHalfway, _ = front.(*PerpendicularVectorHalfway)
		}
		stage.OnAfterPerpendicularVectorHalfwayDeleteCallback.OnAfterDelete(stage, perpendicularvectorhalfway, frontPerpendicularVectorHalfway)
	}
}

func (plant2ddiagram *Plant2DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlant2DDiagramCreateCallback != nil {
		stage.OnAfterPlant2DDiagramCreateCallback.OnAfterCreate(stage, plant2ddiagram)
	}
}

func (plant2ddiagram *Plant2DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlant2DDiagramUpdateCallback != nil {
		var frontPlant2DDiagram *Plant2DDiagram
		if front != nil {
			frontPlant2DDiagram, _ = front.(*Plant2DDiagram)
		}
		stage.OnAfterPlant2DDiagramUpdateCallback.OnAfterUpdate(stage, plant2ddiagram, frontPlant2DDiagram)
	}
}

func (plant2ddiagram *Plant2DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlant2DDiagramDeleteCallback != nil {
		var frontPlant2DDiagram *Plant2DDiagram
		if front != nil {
			frontPlant2DDiagram, _ = front.(*Plant2DDiagram)
		}
		stage.OnAfterPlant2DDiagramDeleteCallback.OnAfterDelete(stage, plant2ddiagram, frontPlant2DDiagram)
	}
}

func (plant3ddiagram *Plant3DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlant3DDiagramCreateCallback != nil {
		stage.OnAfterPlant3DDiagramCreateCallback.OnAfterCreate(stage, plant3ddiagram)
	}
}

func (plant3ddiagram *Plant3DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlant3DDiagramUpdateCallback != nil {
		var frontPlant3DDiagram *Plant3DDiagram
		if front != nil {
			frontPlant3DDiagram, _ = front.(*Plant3DDiagram)
		}
		stage.OnAfterPlant3DDiagramUpdateCallback.OnAfterUpdate(stage, plant3ddiagram, frontPlant3DDiagram)
	}
}

func (plant3ddiagram *Plant3DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlant3DDiagramDeleteCallback != nil {
		var frontPlant3DDiagram *Plant3DDiagram
		if front != nil {
			frontPlant3DDiagram, _ = front.(*Plant3DDiagram)
		}
		stage.OnAfterPlant3DDiagramDeleteCallback.OnAfterDelete(stage, plant3ddiagram, frontPlant3DDiagram)
	}
}

func (plantabstract *PlantAbstract) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlantAbstractCreateCallback != nil {
		stage.OnAfterPlantAbstractCreateCallback.OnAfterCreate(stage, plantabstract)
	}
}

func (plantabstract *PlantAbstract) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlantAbstractUpdateCallback != nil {
		var frontPlantAbstract *PlantAbstract
		if front != nil {
			frontPlantAbstract, _ = front.(*PlantAbstract)
		}
		stage.OnAfterPlantAbstractUpdateCallback.OnAfterUpdate(stage, plantabstract, frontPlantAbstract)
	}
}

func (plantabstract *PlantAbstract) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlantAbstractDeleteCallback != nil {
		var frontPlantAbstract *PlantAbstract
		if front != nil {
			frontPlantAbstract, _ = front.(*PlantAbstract)
		}
		stage.OnAfterPlantAbstractDeleteCallback.OnAfterDelete(stage, plantabstract, frontPlantAbstract)
	}
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlantCircumferenceShapeCreateCallback != nil {
		stage.OnAfterPlantCircumferenceShapeCreateCallback.OnAfterCreate(stage, plantcircumferenceshape)
	}
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlantCircumferenceShapeUpdateCallback != nil {
		var frontPlantCircumferenceShape *PlantCircumferenceShape
		if front != nil {
			frontPlantCircumferenceShape, _ = front.(*PlantCircumferenceShape)
		}
		stage.OnAfterPlantCircumferenceShapeUpdateCallback.OnAfterUpdate(stage, plantcircumferenceshape, frontPlantCircumferenceShape)
	}
}

func (plantcircumferenceshape *PlantCircumferenceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlantCircumferenceShapeDeleteCallback != nil {
		var frontPlantCircumferenceShape *PlantCircumferenceShape
		if front != nil {
			frontPlantCircumferenceShape, _ = front.(*PlantCircumferenceShape)
		}
		stage.OnAfterPlantCircumferenceShapeDeleteCallback.OnAfterDelete(stage, plantcircumferenceshape, frontPlantCircumferenceShape)
	}
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPointsAndLines3DShapeCreateCallback != nil {
		stage.OnAfterPointsAndLines3DShapeCreateCallback.OnAfterCreate(stage, pointsandlines3dshape)
	}
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointsAndLines3DShapeUpdateCallback != nil {
		var frontPointsAndLines3DShape *PointsAndLines3DShape
		if front != nil {
			frontPointsAndLines3DShape, _ = front.(*PointsAndLines3DShape)
		}
		stage.OnAfterPointsAndLines3DShapeUpdateCallback.OnAfterUpdate(stage, pointsandlines3dshape, frontPointsAndLines3DShape)
	}
}

func (pointsandlines3dshape *PointsAndLines3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPointsAndLines3DShapeDeleteCallback != nil {
		var frontPointsAndLines3DShape *PointsAndLines3DShape
		if front != nil {
			frontPointsAndLines3DShape, _ = front.(*PointsAndLines3DShape)
		}
		stage.OnAfterPointsAndLines3DShapeDeleteCallback.OnAfterDelete(stage, pointsandlines3dshape, frontPointsAndLines3DShape)
	}
}

func (pxshape *PxShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPxShapeCreateCallback != nil {
		stage.OnAfterPxShapeCreateCallback.OnAfterCreate(stage, pxshape)
	}
}

func (pxshape *PxShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPxShapeUpdateCallback != nil {
		var frontPxShape *PxShape
		if front != nil {
			frontPxShape, _ = front.(*PxShape)
		}
		stage.OnAfterPxShapeUpdateCallback.OnAfterUpdate(stage, pxshape, frontPxShape)
	}
}

func (pxshape *PxShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPxShapeDeleteCallback != nil {
		var frontPxShape *PxShape
		if front != nil {
			frontPxShape, _ = front.(*PxShape)
		}
		stage.OnAfterPxShapeDeleteCallback.OnAfterDelete(stage, pxshape, frontPxShape)
	}
}

func (rendered3dshape *Rendered3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRendered3DShapeCreateCallback != nil {
		stage.OnAfterRendered3DShapeCreateCallback.OnAfterCreate(stage, rendered3dshape)
	}
}

func (rendered3dshape *Rendered3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRendered3DShapeUpdateCallback != nil {
		var frontRendered3DShape *Rendered3DShape
		if front != nil {
			frontRendered3DShape, _ = front.(*Rendered3DShape)
		}
		stage.OnAfterRendered3DShapeUpdateCallback.OnAfterUpdate(stage, rendered3dshape, frontRendered3DShape)
	}
}

func (rendered3dshape *Rendered3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRendered3DShapeDeleteCallback != nil {
		var frontRendered3DShape *Rendered3DShape
		if front != nil {
			frontRendered3DShape, _ = front.(*Rendered3DShape)
		}
		stage.OnAfterRendered3DShapeDeleteCallback.OnAfterDelete(stage, rendered3dshape, frontRendered3DShape)
	}
}

func (rhombusshape *RhombusShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRhombusShapeCreateCallback != nil {
		stage.OnAfterRhombusShapeCreateCallback.OnAfterCreate(stage, rhombusshape)
	}
}

func (rhombusshape *RhombusShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRhombusShapeUpdateCallback != nil {
		var frontRhombusShape *RhombusShape
		if front != nil {
			frontRhombusShape, _ = front.(*RhombusShape)
		}
		stage.OnAfterRhombusShapeUpdateCallback.OnAfterUpdate(stage, rhombusshape, frontRhombusShape)
	}
}

func (rhombusshape *RhombusShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRhombusShapeDeleteCallback != nil {
		var frontRhombusShape *RhombusShape
		if front != nil {
			frontRhombusShape, _ = front.(*RhombusShape)
		}
		stage.OnAfterRhombusShapeDeleteCallback.OnAfterDelete(stage, rhombusshape, frontRhombusShape)
	}
}

func (rhombusstuff *RhombusStuff) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRhombusStuffCreateCallback != nil {
		stage.OnAfterRhombusStuffCreateCallback.OnAfterCreate(stage, rhombusstuff)
	}
}

func (rhombusstuff *RhombusStuff) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRhombusStuffUpdateCallback != nil {
		var frontRhombusStuff *RhombusStuff
		if front != nil {
			frontRhombusStuff, _ = front.(*RhombusStuff)
		}
		stage.OnAfterRhombusStuffUpdateCallback.OnAfterUpdate(stage, rhombusstuff, frontRhombusStuff)
	}
}

func (rhombusstuff *RhombusStuff) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRhombusStuffDeleteCallback != nil {
		var frontRhombusStuff *RhombusStuff
		if front != nil {
			frontRhombusStuff, _ = front.(*RhombusStuff)
		}
		stage.OnAfterRhombusStuffDeleteCallback.OnAfterDelete(stage, rhombusstuff, frontRhombusStuff)
	}
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRotatedRhombusGridShapeCreateCallback != nil {
		stage.OnAfterRotatedRhombusGridShapeCreateCallback.OnAfterCreate(stage, rotatedrhombusgridshape)
	}
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedRhombusGridShapeUpdateCallback != nil {
		var frontRotatedRhombusGridShape *RotatedRhombusGridShape
		if front != nil {
			frontRotatedRhombusGridShape, _ = front.(*RotatedRhombusGridShape)
		}
		stage.OnAfterRotatedRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, rotatedrhombusgridshape, frontRotatedRhombusGridShape)
	}
}

func (rotatedrhombusgridshape *RotatedRhombusGridShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedRhombusGridShapeDeleteCallback != nil {
		var frontRotatedRhombusGridShape *RotatedRhombusGridShape
		if front != nil {
			frontRotatedRhombusGridShape, _ = front.(*RotatedRhombusGridShape)
		}
		stage.OnAfterRotatedRhombusGridShapeDeleteCallback.OnAfterDelete(stage, rotatedrhombusgridshape, frontRotatedRhombusGridShape)
	}
}

func (rotatedrhombusshape *RotatedRhombusShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRotatedRhombusShapeCreateCallback != nil {
		stage.OnAfterRotatedRhombusShapeCreateCallback.OnAfterCreate(stage, rotatedrhombusshape)
	}
}

func (rotatedrhombusshape *RotatedRhombusShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedRhombusShapeUpdateCallback != nil {
		var frontRotatedRhombusShape *RotatedRhombusShape
		if front != nil {
			frontRotatedRhombusShape, _ = front.(*RotatedRhombusShape)
		}
		stage.OnAfterRotatedRhombusShapeUpdateCallback.OnAfterUpdate(stage, rotatedrhombusshape, frontRotatedRhombusShape)
	}
}

func (rotatedrhombusshape *RotatedRhombusShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedRhombusShapeDeleteCallback != nil {
		var frontRotatedRhombusShape *RotatedRhombusShape
		if front != nil {
			frontRotatedRhombusShape, _ = front.(*RotatedRhombusShape)
		}
		stage.OnAfterRotatedRhombusShapeDeleteCallback.OnAfterDelete(stage, rotatedrhombusshape, frontRotatedRhombusShape)
	}
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRotatedSampledPoints3DShapeCreateCallback != nil {
		stage.OnAfterRotatedSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, rotatedsampledpoints3dshape)
	}
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedSampledPoints3DShapeUpdateCallback != nil {
		var frontRotatedSampledPoints3DShape *RotatedSampledPoints3DShape
		if front != nil {
			frontRotatedSampledPoints3DShape, _ = front.(*RotatedSampledPoints3DShape)
		}
		stage.OnAfterRotatedSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, rotatedsampledpoints3dshape, frontRotatedSampledPoints3DShape)
	}
}

func (rotatedsampledpoints3dshape *RotatedSampledPoints3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedSampledPoints3DShapeDeleteCallback != nil {
		var frontRotatedSampledPoints3DShape *RotatedSampledPoints3DShape
		if front != nil {
			frontRotatedSampledPoints3DShape, _ = front.(*RotatedSampledPoints3DShape)
		}
		stage.OnAfterRotatedSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, rotatedsampledpoints3dshape, frontRotatedSampledPoints3DShape)
	}
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRotatedSeatAndLegs3DShapeCreateCallback != nil {
		stage.OnAfterRotatedSeatAndLegs3DShapeCreateCallback.OnAfterCreate(stage, rotatedseatandlegs3dshape)
	}
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedSeatAndLegs3DShapeUpdateCallback != nil {
		var frontRotatedSeatAndLegs3DShape *RotatedSeatAndLegs3DShape
		if front != nil {
			frontRotatedSeatAndLegs3DShape, _ = front.(*RotatedSeatAndLegs3DShape)
		}
		stage.OnAfterRotatedSeatAndLegs3DShapeUpdateCallback.OnAfterUpdate(stage, rotatedseatandlegs3dshape, frontRotatedSeatAndLegs3DShape)
	}
}

func (rotatedseatandlegs3dshape *RotatedSeatAndLegs3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRotatedSeatAndLegs3DShapeDeleteCallback != nil {
		var frontRotatedSeatAndLegs3DShape *RotatedSeatAndLegs3DShape
		if front != nil {
			frontRotatedSeatAndLegs3DShape, _ = front.(*RotatedSeatAndLegs3DShape)
		}
		stage.OnAfterRotatedSeatAndLegs3DShapeDeleteCallback.OnAfterDelete(stage, rotatedseatandlegs3dshape, frontRotatedSeatAndLegs3DShape)
	}
}

func (sampledpoints3dshape *SampledPoints3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSampledPoints3DShapeCreateCallback != nil {
		stage.OnAfterSampledPoints3DShapeCreateCallback.OnAfterCreate(stage, sampledpoints3dshape)
	}
}

func (sampledpoints3dshape *SampledPoints3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSampledPoints3DShapeUpdateCallback != nil {
		var frontSampledPoints3DShape *SampledPoints3DShape
		if front != nil {
			frontSampledPoints3DShape, _ = front.(*SampledPoints3DShape)
		}
		stage.OnAfterSampledPoints3DShapeUpdateCallback.OnAfterUpdate(stage, sampledpoints3dshape, frontSampledPoints3DShape)
	}
}

func (sampledpoints3dshape *SampledPoints3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSampledPoints3DShapeDeleteCallback != nil {
		var frontSampledPoints3DShape *SampledPoints3DShape
		if front != nil {
			frontSampledPoints3DShape, _ = front.(*SampledPoints3DShape)
		}
		stage.OnAfterSampledPoints3DShapeDeleteCallback.OnAfterDelete(stage, sampledpoints3dshape, frontSampledPoints3DShape)
	}
}

func (seat3dshape *Seat3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeat3DShapeCreateCallback != nil {
		stage.OnAfterSeat3DShapeCreateCallback.OnAfterCreate(stage, seat3dshape)
	}
}

func (seat3dshape *Seat3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeat3DShapeUpdateCallback != nil {
		var frontSeat3DShape *Seat3DShape
		if front != nil {
			frontSeat3DShape, _ = front.(*Seat3DShape)
		}
		stage.OnAfterSeat3DShapeUpdateCallback.OnAfterUpdate(stage, seat3dshape, frontSeat3DShape)
	}
}

func (seat3dshape *Seat3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeat3DShapeDeleteCallback != nil {
		var frontSeat3DShape *Seat3DShape
		if front != nil {
			frontSeat3DShape, _ = front.(*Seat3DShape)
		}
		stage.OnAfterSeat3DShapeDeleteCallback.OnAfterDelete(stage, seat3dshape, frontSeat3DShape)
	}
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeatAndLegs3DShapeCreateCallback != nil {
		stage.OnAfterSeatAndLegs3DShapeCreateCallback.OnAfterCreate(stage, seatandlegs3dshape)
	}
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeatAndLegs3DShapeUpdateCallback != nil {
		var frontSeatAndLegs3DShape *SeatAndLegs3DShape
		if front != nil {
			frontSeatAndLegs3DShape, _ = front.(*SeatAndLegs3DShape)
		}
		stage.OnAfterSeatAndLegs3DShapeUpdateCallback.OnAfterUpdate(stage, seatandlegs3dshape, frontSeatAndLegs3DShape)
	}
}

func (seatandlegs3dshape *SeatAndLegs3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeatAndLegs3DShapeDeleteCallback != nil {
		var frontSeatAndLegs3DShape *SeatAndLegs3DShape
		if front != nil {
			frontSeatAndLegs3DShape, _ = front.(*SeatAndLegs3DShape)
		}
		stage.OnAfterSeatAndLegs3DShapeDeleteCallback.OnAfterDelete(stage, seatandlegs3dshape, frontSeatAndLegs3DShape)
	}
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeatBottomCurveShapeCreateCallback != nil {
		stage.OnAfterSeatBottomCurveShapeCreateCallback.OnAfterCreate(stage, seatbottomcurveshape)
	}
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeatBottomCurveShapeUpdateCallback != nil {
		var frontSeatBottomCurveShape *SeatBottomCurveShape
		if front != nil {
			frontSeatBottomCurveShape, _ = front.(*SeatBottomCurveShape)
		}
		stage.OnAfterSeatBottomCurveShapeUpdateCallback.OnAfterUpdate(stage, seatbottomcurveshape, frontSeatBottomCurveShape)
	}
}

func (seatbottomcurveshape *SeatBottomCurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeatBottomCurveShapeDeleteCallback != nil {
		var frontSeatBottomCurveShape *SeatBottomCurveShape
		if front != nil {
			frontSeatBottomCurveShape, _ = front.(*SeatBottomCurveShape)
		}
		stage.OnAfterSeatBottomCurveShapeDeleteCallback.OnAfterDelete(stage, seatbottomcurveshape, frontSeatBottomCurveShape)
	}
}

func (seattopcurveshape *SeatTopCurveShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSeatTopCurveShapeCreateCallback != nil {
		stage.OnAfterSeatTopCurveShapeCreateCallback.OnAfterCreate(stage, seattopcurveshape)
	}
}

func (seattopcurveshape *SeatTopCurveShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeatTopCurveShapeUpdateCallback != nil {
		var frontSeatTopCurveShape *SeatTopCurveShape
		if front != nil {
			frontSeatTopCurveShape, _ = front.(*SeatTopCurveShape)
		}
		stage.OnAfterSeatTopCurveShapeUpdateCallback.OnAfterUpdate(stage, seattopcurveshape, frontSeatTopCurveShape)
	}
}

func (seattopcurveshape *SeatTopCurveShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSeatTopCurveShapeDeleteCallback != nil {
		var frontSeatTopCurveShape *SeatTopCurveShape
		if front != nil {
			frontSeatTopCurveShape, _ = front.(*SeatTopCurveShape)
		}
		stage.OnAfterSeatTopCurveShapeDeleteCallback.OnAfterDelete(stage, seattopcurveshape, frontSeatTopCurveShape)
	}
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedBottomTopStartArcShapeCreateCallback != nil {
		stage.OnAfterShiftedBottomTopStartArcShapeCreateCallback.OnAfterCreate(stage, shiftedbottomtopstartarcshape)
	}
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedBottomTopStartArcShapeUpdateCallback != nil {
		var frontShiftedBottomTopStartArcShape *ShiftedBottomTopStartArcShape
		if front != nil {
			frontShiftedBottomTopStartArcShape, _ = front.(*ShiftedBottomTopStartArcShape)
		}
		stage.OnAfterShiftedBottomTopStartArcShapeUpdateCallback.OnAfterUpdate(stage, shiftedbottomtopstartarcshape, frontShiftedBottomTopStartArcShape)
	}
}

func (shiftedbottomtopstartarcshape *ShiftedBottomTopStartArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedBottomTopStartArcShapeDeleteCallback != nil {
		var frontShiftedBottomTopStartArcShape *ShiftedBottomTopStartArcShape
		if front != nil {
			frontShiftedBottomTopStartArcShape, _ = front.(*ShiftedBottomTopStartArcShape)
		}
		stage.OnAfterShiftedBottomTopStartArcShapeDeleteCallback.OnAfterDelete(stage, shiftedbottomtopstartarcshape, frontShiftedBottomTopStartArcShape)
	}
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedBottomTopStartArcShapeGridCreateCallback != nil {
		stage.OnAfterShiftedBottomTopStartArcShapeGridCreateCallback.OnAfterCreate(stage, shiftedbottomtopstartarcshapegrid)
	}
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedBottomTopStartArcShapeGridUpdateCallback != nil {
		var frontShiftedBottomTopStartArcShapeGrid *ShiftedBottomTopStartArcShapeGrid
		if front != nil {
			frontShiftedBottomTopStartArcShapeGrid, _ = front.(*ShiftedBottomTopStartArcShapeGrid)
		}
		stage.OnAfterShiftedBottomTopStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, shiftedbottomtopstartarcshapegrid, frontShiftedBottomTopStartArcShapeGrid)
	}
}

func (shiftedbottomtopstartarcshapegrid *ShiftedBottomTopStartArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedBottomTopStartArcShapeGridDeleteCallback != nil {
		var frontShiftedBottomTopStartArcShapeGrid *ShiftedBottomTopStartArcShapeGrid
		if front != nil {
			frontShiftedBottomTopStartArcShapeGrid, _ = front.(*ShiftedBottomTopStartArcShapeGrid)
		}
		stage.OnAfterShiftedBottomTopStartArcShapeGridDeleteCallback.OnAfterDelete(stage, shiftedbottomtopstartarcshapegrid, frontShiftedBottomTopStartArcShapeGrid)
	}
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonCreateCallback != nil {
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, shiftedleftgrowthcurve2dribbon)
	}
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonUpdateCallback != nil {
		var frontShiftedLeftGrowthCurve2DRibbon *ShiftedLeftGrowthCurve2DRibbon
		if front != nil {
			frontShiftedLeftGrowthCurve2DRibbon, _ = front.(*ShiftedLeftGrowthCurve2DRibbon)
		}
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, shiftedleftgrowthcurve2dribbon, frontShiftedLeftGrowthCurve2DRibbon)
	}
}

func (shiftedleftgrowthcurve2dribbon *ShiftedLeftGrowthCurve2DRibbon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonDeleteCallback != nil {
		var frontShiftedLeftGrowthCurve2DRibbon *ShiftedLeftGrowthCurve2DRibbon
		if front != nil {
			frontShiftedLeftGrowthCurve2DRibbon, _ = front.(*ShiftedLeftGrowthCurve2DRibbon)
		}
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, shiftedleftgrowthcurve2dribbon, frontShiftedLeftGrowthCurve2DRibbon)
	}
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeCreateCallback != nil {
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, shiftedleftgrowthcurve2dribbonendshape)
	}
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
		var frontShiftedLeftGrowthCurve2DRibbonEndShape *ShiftedLeftGrowthCurve2DRibbonEndShape
		if front != nil {
			frontShiftedLeftGrowthCurve2DRibbonEndShape, _ = front.(*ShiftedLeftGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, shiftedleftgrowthcurve2dribbonendshape, frontShiftedLeftGrowthCurve2DRibbonEndShape)
	}
}

func (shiftedleftgrowthcurve2dribbonendshape *ShiftedLeftGrowthCurve2DRibbonEndShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
		var frontShiftedLeftGrowthCurve2DRibbonEndShape *ShiftedLeftGrowthCurve2DRibbonEndShape
		if front != nil {
			frontShiftedLeftGrowthCurve2DRibbonEndShape, _ = front.(*ShiftedLeftGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, shiftedleftgrowthcurve2dribbonendshape, frontShiftedLeftGrowthCurve2DRibbonEndShape)
	}
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeCreateCallback != nil {
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, shiftedleftgrowthcurve2dribbonstartshape)
	}
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
		var frontShiftedLeftGrowthCurve2DRibbonStartShape *ShiftedLeftGrowthCurve2DRibbonStartShape
		if front != nil {
			frontShiftedLeftGrowthCurve2DRibbonStartShape, _ = front.(*ShiftedLeftGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, shiftedleftgrowthcurve2dribbonstartshape, frontShiftedLeftGrowthCurve2DRibbonStartShape)
	}
}

func (shiftedleftgrowthcurve2dribbonstartshape *ShiftedLeftGrowthCurve2DRibbonStartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
		var frontShiftedLeftGrowthCurve2DRibbonStartShape *ShiftedLeftGrowthCurve2DRibbonStartShape
		if front != nil {
			frontShiftedLeftGrowthCurve2DRibbonStartShape, _ = front.(*ShiftedLeftGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterShiftedLeftGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, shiftedleftgrowthcurve2dribbonstartshape, frontShiftedLeftGrowthCurve2DRibbonStartShape)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonCreateCallback != nil {
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, shiftedleftpartiallygrowthcurve2dribbon)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonUpdateCallback != nil {
		var frontShiftedLeftPartiallyGrowthCurve2DRibbon *ShiftedLeftPartiallyGrowthCurve2DRibbon
		if front != nil {
			frontShiftedLeftPartiallyGrowthCurve2DRibbon, _ = front.(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
		}
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, shiftedleftpartiallygrowthcurve2dribbon, frontShiftedLeftPartiallyGrowthCurve2DRibbon)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbon *ShiftedLeftPartiallyGrowthCurve2DRibbon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonDeleteCallback != nil {
		var frontShiftedLeftPartiallyGrowthCurve2DRibbon *ShiftedLeftPartiallyGrowthCurve2DRibbon
		if front != nil {
			frontShiftedLeftPartiallyGrowthCurve2DRibbon, _ = front.(*ShiftedLeftPartiallyGrowthCurve2DRibbon)
		}
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, shiftedleftpartiallygrowthcurve2dribbon, frontShiftedLeftPartiallyGrowthCurve2DRibbon)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeCreateCallback != nil {
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, shiftedleftpartiallygrowthcurve2dribbonendshape)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
		var frontShiftedLeftPartiallyGrowthCurve2DRibbonEndShape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape
		if front != nil {
			frontShiftedLeftPartiallyGrowthCurve2DRibbonEndShape, _ = front.(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, shiftedleftpartiallygrowthcurve2dribbonendshape, frontShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonendshape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
		var frontShiftedLeftPartiallyGrowthCurve2DRibbonEndShape *ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape
		if front != nil {
			frontShiftedLeftPartiallyGrowthCurve2DRibbonEndShape, _ = front.(*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, shiftedleftpartiallygrowthcurve2dribbonendshape, frontShiftedLeftPartiallyGrowthCurve2DRibbonEndShape)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeCreateCallback != nil {
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, shiftedleftpartiallygrowthcurve2dribbonstartshape)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
		var frontShiftedLeftPartiallyGrowthCurve2DRibbonStartShape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape
		if front != nil {
			frontShiftedLeftPartiallyGrowthCurve2DRibbonStartShape, _ = front.(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, shiftedleftpartiallygrowthcurve2dribbonstartshape, frontShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	}
}

func (shiftedleftpartiallygrowthcurve2dribbonstartshape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
		var frontShiftedLeftPartiallyGrowthCurve2DRibbonStartShape *ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape
		if front != nil {
			frontShiftedLeftPartiallyGrowthCurve2DRibbonStartShape, _ = front.(*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterShiftedLeftPartiallyGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, shiftedleftpartiallygrowthcurve2dribbonstartshape, frontShiftedLeftPartiallyGrowthCurve2DRibbonStartShape)
	}
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback != nil {
		stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeCreateCallback.OnAfterCreate(stage, shiftedleftstackgrowthcurveendarcshape)
	}
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback != nil {
		var frontShiftedLeftStackGrowthCurveEndArcShape *ShiftedLeftStackGrowthCurveEndArcShape
		if front != nil {
			frontShiftedLeftStackGrowthCurveEndArcShape, _ = front.(*ShiftedLeftStackGrowthCurveEndArcShape)
		}
		stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeUpdateCallback.OnAfterUpdate(stage, shiftedleftstackgrowthcurveendarcshape, frontShiftedLeftStackGrowthCurveEndArcShape)
	}
}

func (shiftedleftstackgrowthcurveendarcshape *ShiftedLeftStackGrowthCurveEndArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback != nil {
		var frontShiftedLeftStackGrowthCurveEndArcShape *ShiftedLeftStackGrowthCurveEndArcShape
		if front != nil {
			frontShiftedLeftStackGrowthCurveEndArcShape, _ = front.(*ShiftedLeftStackGrowthCurveEndArcShape)
		}
		stage.OnAfterShiftedLeftStackGrowthCurveEndArcShapeDeleteCallback.OnAfterDelete(stage, shiftedleftstackgrowthcurveendarcshape, frontShiftedLeftStackGrowthCurveEndArcShape)
	}
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback != nil {
		stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeCreateCallback.OnAfterCreate(stage, shiftedleftstackgrowthcurvestartarcshape)
	}
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback != nil {
		var frontShiftedLeftStackGrowthCurveStartArcShape *ShiftedLeftStackGrowthCurveStartArcShape
		if front != nil {
			frontShiftedLeftStackGrowthCurveStartArcShape, _ = front.(*ShiftedLeftStackGrowthCurveStartArcShape)
		}
		stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeUpdateCallback.OnAfterUpdate(stage, shiftedleftstackgrowthcurvestartarcshape, frontShiftedLeftStackGrowthCurveStartArcShape)
	}
}

func (shiftedleftstackgrowthcurvestartarcshape *ShiftedLeftStackGrowthCurveStartArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback != nil {
		var frontShiftedLeftStackGrowthCurveStartArcShape *ShiftedLeftStackGrowthCurveStartArcShape
		if front != nil {
			frontShiftedLeftStackGrowthCurveStartArcShape, _ = front.(*ShiftedLeftStackGrowthCurveStartArcShape)
		}
		stage.OnAfterShiftedLeftStackGrowthCurveStartArcShapeDeleteCallback.OnAfterDelete(stage, shiftedleftstackgrowthcurvestartarcshape, frontShiftedLeftStackGrowthCurveStartArcShape)
	}
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftStackNormalVectorCreateCallback != nil {
		stage.OnAfterShiftedLeftStackNormalVectorCreateCallback.OnAfterCreate(stage, shiftedleftstacknormalvector)
	}
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackNormalVectorUpdateCallback != nil {
		var frontShiftedLeftStackNormalVector *ShiftedLeftStackNormalVector
		if front != nil {
			frontShiftedLeftStackNormalVector, _ = front.(*ShiftedLeftStackNormalVector)
		}
		stage.OnAfterShiftedLeftStackNormalVectorUpdateCallback.OnAfterUpdate(stage, shiftedleftstacknormalvector, frontShiftedLeftStackNormalVector)
	}
}

func (shiftedleftstacknormalvector *ShiftedLeftStackNormalVector) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackNormalVectorDeleteCallback != nil {
		var frontShiftedLeftStackNormalVector *ShiftedLeftStackNormalVector
		if front != nil {
			frontShiftedLeftStackNormalVector, _ = front.(*ShiftedLeftStackNormalVector)
		}
		stage.OnAfterShiftedLeftStackNormalVectorDeleteCallback.OnAfterDelete(stage, shiftedleftstacknormalvector, frontShiftedLeftStackNormalVector)
	}
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftStackOfGrowthCurveCreateCallback != nil {
		stage.OnAfterShiftedLeftStackOfGrowthCurveCreateCallback.OnAfterCreate(stage, shiftedleftstackofgrowthcurve)
	}
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback != nil {
		var frontShiftedLeftStackOfGrowthCurve *ShiftedLeftStackOfGrowthCurve
		if front != nil {
			frontShiftedLeftStackOfGrowthCurve, _ = front.(*ShiftedLeftStackOfGrowthCurve)
		}
		stage.OnAfterShiftedLeftStackOfGrowthCurveUpdateCallback.OnAfterUpdate(stage, shiftedleftstackofgrowthcurve, frontShiftedLeftStackOfGrowthCurve)
	}
}

func (shiftedleftstackofgrowthcurve *ShiftedLeftStackOfGrowthCurve) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback != nil {
		var frontShiftedLeftStackOfGrowthCurve *ShiftedLeftStackOfGrowthCurve
		if front != nil {
			frontShiftedLeftStackOfGrowthCurve, _ = front.(*ShiftedLeftStackOfGrowthCurve)
		}
		stage.OnAfterShiftedLeftStackOfGrowthCurveDeleteCallback.OnAfterDelete(stage, shiftedleftstackofgrowthcurve, frontShiftedLeftStackOfGrowthCurve)
	}
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedLeftStackOfNormalVectorCreateCallback != nil {
		stage.OnAfterShiftedLeftStackOfNormalVectorCreateCallback.OnAfterCreate(stage, shiftedleftstackofnormalvector)
	}
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackOfNormalVectorUpdateCallback != nil {
		var frontShiftedLeftStackOfNormalVector *ShiftedLeftStackOfNormalVector
		if front != nil {
			frontShiftedLeftStackOfNormalVector, _ = front.(*ShiftedLeftStackOfNormalVector)
		}
		stage.OnAfterShiftedLeftStackOfNormalVectorUpdateCallback.OnAfterUpdate(stage, shiftedleftstackofnormalvector, frontShiftedLeftStackOfNormalVector)
	}
}

func (shiftedleftstackofnormalvector *ShiftedLeftStackOfNormalVector) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedLeftStackOfNormalVectorDeleteCallback != nil {
		var frontShiftedLeftStackOfNormalVector *ShiftedLeftStackOfNormalVector
		if front != nil {
			frontShiftedLeftStackOfNormalVector, _ = front.(*ShiftedLeftStackOfNormalVector)
		}
		stage.OnAfterShiftedLeftStackOfNormalVectorDeleteCallback.OnAfterDelete(stage, shiftedleftstackofnormalvector, frontShiftedLeftStackOfNormalVector)
	}
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonCreateCallback != nil {
		stage.OnAfterShiftedRightGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, shiftedrightgrowthcurve2dribbon)
	}
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonUpdateCallback != nil {
		var frontShiftedRightGrowthCurve2DRibbon *ShiftedRightGrowthCurve2DRibbon
		if front != nil {
			frontShiftedRightGrowthCurve2DRibbon, _ = front.(*ShiftedRightGrowthCurve2DRibbon)
		}
		stage.OnAfterShiftedRightGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, shiftedrightgrowthcurve2dribbon, frontShiftedRightGrowthCurve2DRibbon)
	}
}

func (shiftedrightgrowthcurve2dribbon *ShiftedRightGrowthCurve2DRibbon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonDeleteCallback != nil {
		var frontShiftedRightGrowthCurve2DRibbon *ShiftedRightGrowthCurve2DRibbon
		if front != nil {
			frontShiftedRightGrowthCurve2DRibbon, _ = front.(*ShiftedRightGrowthCurve2DRibbon)
		}
		stage.OnAfterShiftedRightGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, shiftedrightgrowthcurve2dribbon, frontShiftedRightGrowthCurve2DRibbon)
	}
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeCreateCallback != nil {
		stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, shiftedrightgrowthcurve2dribbonendshape)
	}
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
		var frontShiftedRightGrowthCurve2DRibbonEndShape *ShiftedRightGrowthCurve2DRibbonEndShape
		if front != nil {
			frontShiftedRightGrowthCurve2DRibbonEndShape, _ = front.(*ShiftedRightGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, shiftedrightgrowthcurve2dribbonendshape, frontShiftedRightGrowthCurve2DRibbonEndShape)
	}
}

func (shiftedrightgrowthcurve2dribbonendshape *ShiftedRightGrowthCurve2DRibbonEndShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
		var frontShiftedRightGrowthCurve2DRibbonEndShape *ShiftedRightGrowthCurve2DRibbonEndShape
		if front != nil {
			frontShiftedRightGrowthCurve2DRibbonEndShape, _ = front.(*ShiftedRightGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterShiftedRightGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, shiftedrightgrowthcurve2dribbonendshape, frontShiftedRightGrowthCurve2DRibbonEndShape)
	}
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeCreateCallback != nil {
		stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, shiftedrightgrowthcurve2dribbonstartshape)
	}
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
		var frontShiftedRightGrowthCurve2DRibbonStartShape *ShiftedRightGrowthCurve2DRibbonStartShape
		if front != nil {
			frontShiftedRightGrowthCurve2DRibbonStartShape, _ = front.(*ShiftedRightGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, shiftedrightgrowthcurve2dribbonstartshape, frontShiftedRightGrowthCurve2DRibbonStartShape)
	}
}

func (shiftedrightgrowthcurve2dribbonstartshape *ShiftedRightGrowthCurve2DRibbonStartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
		var frontShiftedRightGrowthCurve2DRibbonStartShape *ShiftedRightGrowthCurve2DRibbonStartShape
		if front != nil {
			frontShiftedRightGrowthCurve2DRibbonStartShape, _ = front.(*ShiftedRightGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterShiftedRightGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, shiftedrightgrowthcurve2dribbonstartshape, frontShiftedRightGrowthCurve2DRibbonStartShape)
	}
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, stackgrowthcurve2dendhalfwayarcshape)
	}
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback != nil {
		var frontStackGrowthCurve2DEndHalfwayArcShape *StackGrowthCurve2DEndHalfwayArcShape
		if front != nil {
			frontStackGrowthCurve2DEndHalfwayArcShape, _ = front.(*StackGrowthCurve2DEndHalfwayArcShape)
		}
		stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, stackgrowthcurve2dendhalfwayarcshape, frontStackGrowthCurve2DEndHalfwayArcShape)
	}
}

func (stackgrowthcurve2dendhalfwayarcshape *StackGrowthCurve2DEndHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback != nil {
		var frontStackGrowthCurve2DEndHalfwayArcShape *StackGrowthCurve2DEndHalfwayArcShape
		if front != nil {
			frontStackGrowthCurve2DEndHalfwayArcShape, _ = front.(*StackGrowthCurve2DEndHalfwayArcShape)
		}
		stage.OnAfterStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, stackgrowthcurve2dendhalfwayarcshape, frontStackGrowthCurve2DEndHalfwayArcShape)
	}
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackGrowthCurve2DRibbonEndShapeCreateCallback != nil {
		stage.OnAfterStackGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, stackgrowthcurve2dribbonendshape)
	}
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
		var frontStackGrowthCurve2DRibbonEndShape *StackGrowthCurve2DRibbonEndShape
		if front != nil {
			frontStackGrowthCurve2DRibbonEndShape, _ = front.(*StackGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterStackGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, stackgrowthcurve2dribbonendshape, frontStackGrowthCurve2DRibbonEndShape)
	}
}

func (stackgrowthcurve2dribbonendshape *StackGrowthCurve2DRibbonEndShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
		var frontStackGrowthCurve2DRibbonEndShape *StackGrowthCurve2DRibbonEndShape
		if front != nil {
			frontStackGrowthCurve2DRibbonEndShape, _ = front.(*StackGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterStackGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, stackgrowthcurve2dribbonendshape, frontStackGrowthCurve2DRibbonEndShape)
	}
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackGrowthCurve2DRibbonStartShapeCreateCallback != nil {
		stage.OnAfterStackGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, stackgrowthcurve2dribbonstartshape)
	}
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
		var frontStackGrowthCurve2DRibbonStartShape *StackGrowthCurve2DRibbonStartShape
		if front != nil {
			frontStackGrowthCurve2DRibbonStartShape, _ = front.(*StackGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterStackGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, stackgrowthcurve2dribbonstartshape, frontStackGrowthCurve2DRibbonStartShape)
	}
}

func (stackgrowthcurve2dribbonstartshape *StackGrowthCurve2DRibbonStartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
		var frontStackGrowthCurve2DRibbonStartShape *StackGrowthCurve2DRibbonStartShape
		if front != nil {
			frontStackGrowthCurve2DRibbonStartShape, _ = front.(*StackGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterStackGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, stackgrowthcurve2dribbonstartshape, frontStackGrowthCurve2DRibbonStartShape)
	}
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, stackgrowthcurve2dstarthalfwayarcshape)
	}
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback != nil {
		var frontStackGrowthCurve2DStartHalfwayArcShape *StackGrowthCurve2DStartHalfwayArcShape
		if front != nil {
			frontStackGrowthCurve2DStartHalfwayArcShape, _ = front.(*StackGrowthCurve2DStartHalfwayArcShape)
		}
		stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, stackgrowthcurve2dstarthalfwayarcshape, frontStackGrowthCurve2DStartHalfwayArcShape)
	}
}

func (stackgrowthcurve2dstarthalfwayarcshape *StackGrowthCurve2DStartHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback != nil {
		var frontStackGrowthCurve2DStartHalfwayArcShape *StackGrowthCurve2DStartHalfwayArcShape
		if front != nil {
			frontStackGrowthCurve2DStartHalfwayArcShape, _ = front.(*StackGrowthCurve2DStartHalfwayArcShape)
		}
		stage.OnAfterStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, stackgrowthcurve2dstarthalfwayarcshape, frontStackGrowthCurve2DStartHalfwayArcShape)
	}
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackOfGrowthCurve2DCreateCallback != nil {
		stage.OnAfterStackOfGrowthCurve2DCreateCallback.OnAfterCreate(stage, stackofgrowthcurve2d)
	}
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfGrowthCurve2DUpdateCallback != nil {
		var frontStackOfGrowthCurve2D *StackOfGrowthCurve2D
		if front != nil {
			frontStackOfGrowthCurve2D, _ = front.(*StackOfGrowthCurve2D)
		}
		stage.OnAfterStackOfGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, stackofgrowthcurve2d, frontStackOfGrowthCurve2D)
	}
}

func (stackofgrowthcurve2d *StackOfGrowthCurve2D) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfGrowthCurve2DDeleteCallback != nil {
		var frontStackOfGrowthCurve2D *StackOfGrowthCurve2D
		if front != nil {
			frontStackOfGrowthCurve2D, _ = front.(*StackOfGrowthCurve2D)
		}
		stage.OnAfterStackOfGrowthCurve2DDeleteCallback.OnAfterDelete(stage, stackofgrowthcurve2d, frontStackOfGrowthCurve2D)
	}
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackOfGrowthCurve2DByGrowthVectorCreateCallback != nil {
		stage.OnAfterStackOfGrowthCurve2DByGrowthVectorCreateCallback.OnAfterCreate(stage, stackofgrowthcurve2dbygrowthvector)
	}
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfGrowthCurve2DByGrowthVectorUpdateCallback != nil {
		var frontStackOfGrowthCurve2DByGrowthVector *StackOfGrowthCurve2DByGrowthVector
		if front != nil {
			frontStackOfGrowthCurve2DByGrowthVector, _ = front.(*StackOfGrowthCurve2DByGrowthVector)
		}
		stage.OnAfterStackOfGrowthCurve2DByGrowthVectorUpdateCallback.OnAfterUpdate(stage, stackofgrowthcurve2dbygrowthvector, frontStackOfGrowthCurve2DByGrowthVector)
	}
}

func (stackofgrowthcurve2dbygrowthvector *StackOfGrowthCurve2DByGrowthVector) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfGrowthCurve2DByGrowthVectorDeleteCallback != nil {
		var frontStackOfGrowthCurve2DByGrowthVector *StackOfGrowthCurve2DByGrowthVector
		if front != nil {
			frontStackOfGrowthCurve2DByGrowthVector, _ = front.(*StackOfGrowthCurve2DByGrowthVector)
		}
		stage.OnAfterStackOfGrowthCurve2DByGrowthVectorDeleteCallback.OnAfterDelete(stage, stackofgrowthcurve2dbygrowthvector, frontStackOfGrowthCurve2DByGrowthVector)
	}
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackOfGrowthCurve2DRibbonCreateCallback != nil {
		stage.OnAfterStackOfGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, stackofgrowthcurve2dribbon)
	}
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfGrowthCurve2DRibbonUpdateCallback != nil {
		var frontStackOfGrowthCurve2DRibbon *StackOfGrowthCurve2DRibbon
		if front != nil {
			frontStackOfGrowthCurve2DRibbon, _ = front.(*StackOfGrowthCurve2DRibbon)
		}
		stage.OnAfterStackOfGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, stackofgrowthcurve2dribbon, frontStackOfGrowthCurve2DRibbon)
	}
}

func (stackofgrowthcurve2dribbon *StackOfGrowthCurve2DRibbon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfGrowthCurve2DRibbonDeleteCallback != nil {
		var frontStackOfGrowthCurve2DRibbon *StackOfGrowthCurve2DRibbon
		if front != nil {
			frontStackOfGrowthCurve2DRibbon, _ = front.(*StackOfGrowthCurve2DRibbon)
		}
		stage.OnAfterStackOfGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, stackofgrowthcurve2dribbon, frontStackOfGrowthCurve2DRibbon)
	}
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackOfPartiallyRotatedTorusShapeCreateCallback != nil {
		stage.OnAfterStackOfPartiallyRotatedTorusShapeCreateCallback.OnAfterCreate(stage, stackofpartiallyrotatedtorusshape)
	}
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfPartiallyRotatedTorusShapeUpdateCallback != nil {
		var frontStackOfPartiallyRotatedTorusShape *StackOfPartiallyRotatedTorusShape
		if front != nil {
			frontStackOfPartiallyRotatedTorusShape, _ = front.(*StackOfPartiallyRotatedTorusShape)
		}
		stage.OnAfterStackOfPartiallyRotatedTorusShapeUpdateCallback.OnAfterUpdate(stage, stackofpartiallyrotatedtorusshape, frontStackOfPartiallyRotatedTorusShape)
	}
}

func (stackofpartiallyrotatedtorusshape *StackOfPartiallyRotatedTorusShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfPartiallyRotatedTorusShapeDeleteCallback != nil {
		var frontStackOfPartiallyRotatedTorusShape *StackOfPartiallyRotatedTorusShape
		if front != nil {
			frontStackOfPartiallyRotatedTorusShape, _ = front.(*StackOfPartiallyRotatedTorusShape)
		}
		stage.OnAfterStackOfPartiallyRotatedTorusShapeDeleteCallback.OnAfterDelete(stage, stackofpartiallyrotatedtorusshape, frontStackOfPartiallyRotatedTorusShape)
	}
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackOfRotatedGrowthCurve2DCreateCallback != nil {
		stage.OnAfterStackOfRotatedGrowthCurve2DCreateCallback.OnAfterCreate(stage, stackofrotatedgrowthcurve2d)
	}
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfRotatedGrowthCurve2DUpdateCallback != nil {
		var frontStackOfRotatedGrowthCurve2D *StackOfRotatedGrowthCurve2D
		if front != nil {
			frontStackOfRotatedGrowthCurve2D, _ = front.(*StackOfRotatedGrowthCurve2D)
		}
		stage.OnAfterStackOfRotatedGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, stackofrotatedgrowthcurve2d, frontStackOfRotatedGrowthCurve2D)
	}
}

func (stackofrotatedgrowthcurve2d *StackOfRotatedGrowthCurve2D) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfRotatedGrowthCurve2DDeleteCallback != nil {
		var frontStackOfRotatedGrowthCurve2D *StackOfRotatedGrowthCurve2D
		if front != nil {
			frontStackOfRotatedGrowthCurve2D, _ = front.(*StackOfRotatedGrowthCurve2D)
		}
		stage.OnAfterStackOfRotatedGrowthCurve2DDeleteCallback.OnAfterDelete(stage, stackofrotatedgrowthcurve2d, frontStackOfRotatedGrowthCurve2D)
	}
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonCreateCallback != nil {
		stage.OnAfterStackOfRotatedGrowthCurve2DRibbonCreateCallback.OnAfterCreate(stage, stackofrotatedgrowthcurve2dribbon)
	}
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonUpdateCallback != nil {
		var frontStackOfRotatedGrowthCurve2DRibbon *StackOfRotatedGrowthCurve2DRibbon
		if front != nil {
			frontStackOfRotatedGrowthCurve2DRibbon, _ = front.(*StackOfRotatedGrowthCurve2DRibbon)
		}
		stage.OnAfterStackOfRotatedGrowthCurve2DRibbonUpdateCallback.OnAfterUpdate(stage, stackofrotatedgrowthcurve2dribbon, frontStackOfRotatedGrowthCurve2DRibbon)
	}
}

func (stackofrotatedgrowthcurve2dribbon *StackOfRotatedGrowthCurve2DRibbon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackOfRotatedGrowthCurve2DRibbonDeleteCallback != nil {
		var frontStackOfRotatedGrowthCurve2DRibbon *StackOfRotatedGrowthCurve2DRibbon
		if front != nil {
			frontStackOfRotatedGrowthCurve2DRibbon, _ = front.(*StackOfRotatedGrowthCurve2DRibbon)
		}
		stage.OnAfterStackOfRotatedGrowthCurve2DRibbonDeleteCallback.OnAfterDelete(stage, stackofrotatedgrowthcurve2dribbon, frontStackOfRotatedGrowthCurve2DRibbon)
	}
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeCreateCallback != nil {
		stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeCreateCallback.OnAfterCreate(stage, stackrotatedgrowthcurve2dendarcshape)
	}
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeUpdateCallback != nil {
		var frontStackRotatedGrowthCurve2DEndArcShape *StackRotatedGrowthCurve2DEndArcShape
		if front != nil {
			frontStackRotatedGrowthCurve2DEndArcShape, _ = front.(*StackRotatedGrowthCurve2DEndArcShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeUpdateCallback.OnAfterUpdate(stage, stackrotatedgrowthcurve2dendarcshape, frontStackRotatedGrowthCurve2DEndArcShape)
	}
}

func (stackrotatedgrowthcurve2dendarcshape *StackRotatedGrowthCurve2DEndArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeDeleteCallback != nil {
		var frontStackRotatedGrowthCurve2DEndArcShape *StackRotatedGrowthCurve2DEndArcShape
		if front != nil {
			frontStackRotatedGrowthCurve2DEndArcShape, _ = front.(*StackRotatedGrowthCurve2DEndArcShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DEndArcShapeDeleteCallback.OnAfterDelete(stage, stackrotatedgrowthcurve2dendarcshape, frontStackRotatedGrowthCurve2DEndArcShape)
	}
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeCreateCallback != nil {
		stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeCreateCallback.OnAfterCreate(stage, stackrotatedgrowthcurve2dribbonendshape)
	}
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeUpdateCallback != nil {
		var frontStackRotatedGrowthCurve2DRibbonEndShape *StackRotatedGrowthCurve2DRibbonEndShape
		if front != nil {
			frontStackRotatedGrowthCurve2DRibbonEndShape, _ = front.(*StackRotatedGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeUpdateCallback.OnAfterUpdate(stage, stackrotatedgrowthcurve2dribbonendshape, frontStackRotatedGrowthCurve2DRibbonEndShape)
	}
}

func (stackrotatedgrowthcurve2dribbonendshape *StackRotatedGrowthCurve2DRibbonEndShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeDeleteCallback != nil {
		var frontStackRotatedGrowthCurve2DRibbonEndShape *StackRotatedGrowthCurve2DRibbonEndShape
		if front != nil {
			frontStackRotatedGrowthCurve2DRibbonEndShape, _ = front.(*StackRotatedGrowthCurve2DRibbonEndShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DRibbonEndShapeDeleteCallback.OnAfterDelete(stage, stackrotatedgrowthcurve2dribbonendshape, frontStackRotatedGrowthCurve2DRibbonEndShape)
	}
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeCreateCallback != nil {
		stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeCreateCallback.OnAfterCreate(stage, stackrotatedgrowthcurve2dribbonstartshape)
	}
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeUpdateCallback != nil {
		var frontStackRotatedGrowthCurve2DRibbonStartShape *StackRotatedGrowthCurve2DRibbonStartShape
		if front != nil {
			frontStackRotatedGrowthCurve2DRibbonStartShape, _ = front.(*StackRotatedGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeUpdateCallback.OnAfterUpdate(stage, stackrotatedgrowthcurve2dribbonstartshape, frontStackRotatedGrowthCurve2DRibbonStartShape)
	}
}

func (stackrotatedgrowthcurve2dribbonstartshape *StackRotatedGrowthCurve2DRibbonStartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeDeleteCallback != nil {
		var frontStackRotatedGrowthCurve2DRibbonStartShape *StackRotatedGrowthCurve2DRibbonStartShape
		if front != nil {
			frontStackRotatedGrowthCurve2DRibbonStartShape, _ = front.(*StackRotatedGrowthCurve2DRibbonStartShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DRibbonStartShapeDeleteCallback.OnAfterDelete(stage, stackrotatedgrowthcurve2dribbonstartshape, frontStackRotatedGrowthCurve2DRibbonStartShape)
	}
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeCreateCallback != nil {
		stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeCreateCallback.OnAfterCreate(stage, stackrotatedgrowthcurve2dstartarcshape)
	}
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeUpdateCallback != nil {
		var frontStackRotatedGrowthCurve2DStartArcShape *StackRotatedGrowthCurve2DStartArcShape
		if front != nil {
			frontStackRotatedGrowthCurve2DStartArcShape, _ = front.(*StackRotatedGrowthCurve2DStartArcShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeUpdateCallback.OnAfterUpdate(stage, stackrotatedgrowthcurve2dstartarcshape, frontStackRotatedGrowthCurve2DStartArcShape)
	}
}

func (stackrotatedgrowthcurve2dstartarcshape *StackRotatedGrowthCurve2DStartArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeDeleteCallback != nil {
		var frontStackRotatedGrowthCurve2DStartArcShape *StackRotatedGrowthCurve2DStartArcShape
		if front != nil {
			frontStackRotatedGrowthCurve2DStartArcShape, _ = front.(*StackRotatedGrowthCurve2DStartArcShape)
		}
		stage.OnAfterStackRotatedGrowthCurve2DStartArcShapeDeleteCallback.OnAfterDelete(stage, stackrotatedgrowthcurve2dstartarcshape, frontStackRotatedGrowthCurve2DStartArcShape)
	}
}

func (startarcshape *StartArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStartArcShapeCreateCallback != nil {
		stage.OnAfterStartArcShapeCreateCallback.OnAfterCreate(stage, startarcshape)
	}
}

func (startarcshape *StartArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartArcShapeUpdateCallback != nil {
		var frontStartArcShape *StartArcShape
		if front != nil {
			frontStartArcShape, _ = front.(*StartArcShape)
		}
		stage.OnAfterStartArcShapeUpdateCallback.OnAfterUpdate(stage, startarcshape, frontStartArcShape)
	}
}

func (startarcshape *StartArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartArcShapeDeleteCallback != nil {
		var frontStartArcShape *StartArcShape
		if front != nil {
			frontStartArcShape, _ = front.(*StartArcShape)
		}
		stage.OnAfterStartArcShapeDeleteCallback.OnAfterDelete(stage, startarcshape, frontStartArcShape)
	}
}

func (startarcshapegrid *StartArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStartArcShapeGridCreateCallback != nil {
		stage.OnAfterStartArcShapeGridCreateCallback.OnAfterCreate(stage, startarcshapegrid)
	}
}

func (startarcshapegrid *StartArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartArcShapeGridUpdateCallback != nil {
		var frontStartArcShapeGrid *StartArcShapeGrid
		if front != nil {
			frontStartArcShapeGrid, _ = front.(*StartArcShapeGrid)
		}
		stage.OnAfterStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, startarcshapegrid, frontStartArcShapeGrid)
	}
}

func (startarcshapegrid *StartArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartArcShapeGridDeleteCallback != nil {
		var frontStartArcShapeGrid *StartArcShapeGrid
		if front != nil {
			frontStartArcShapeGrid, _ = front.(*StartArcShapeGrid)
		}
		stage.OnAfterStartArcShapeGridDeleteCallback.OnAfterDelete(stage, startarcshapegrid, frontStartArcShapeGrid)
	}
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStartHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, starthalfwayarcshape)
	}
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartHalfwayArcShapeUpdateCallback != nil {
		var frontStartHalfwayArcShape *StartHalfwayArcShape
		if front != nil {
			frontStartHalfwayArcShape, _ = front.(*StartHalfwayArcShape)
		}
		stage.OnAfterStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, starthalfwayarcshape, frontStartHalfwayArcShape)
	}
}

func (starthalfwayarcshape *StartHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartHalfwayArcShapeDeleteCallback != nil {
		var frontStartHalfwayArcShape *StartHalfwayArcShape
		if front != nil {
			frontStartHalfwayArcShape, _ = front.(*StartHalfwayArcShape)
		}
		stage.OnAfterStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, starthalfwayarcshape, frontStartHalfwayArcShape)
	}
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStartHalfwayArcShapeGridCreateCallback != nil {
		stage.OnAfterStartHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, starthalfwayarcshapegrid)
	}
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartHalfwayArcShapeGridUpdateCallback != nil {
		var frontStartHalfwayArcShapeGrid *StartHalfwayArcShapeGrid
		if front != nil {
			frontStartHalfwayArcShapeGrid, _ = front.(*StartHalfwayArcShapeGrid)
		}
		stage.OnAfterStartHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, starthalfwayarcshapegrid, frontStartHalfwayArcShapeGrid)
	}
}

func (starthalfwayarcshapegrid *StartHalfwayArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStartHalfwayArcShapeGridDeleteCallback != nil {
		var frontStartHalfwayArcShapeGrid *StartHalfwayArcShapeGrid
		if front != nil {
			frontStartHalfwayArcShapeGrid, _ = front.(*StartHalfwayArcShapeGrid)
		}
		stage.OnAfterStartHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, starthalfwayarcshapegrid, frontStartHalfwayArcShapeGrid)
	}
}

func (stemcylinder3dshape *StemCylinder3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStemCylinder3DShapeCreateCallback != nil {
		stage.OnAfterStemCylinder3DShapeCreateCallback.OnAfterCreate(stage, stemcylinder3dshape)
	}
}

func (stemcylinder3dshape *StemCylinder3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStemCylinder3DShapeUpdateCallback != nil {
		var frontStemCylinder3DShape *StemCylinder3DShape
		if front != nil {
			frontStemCylinder3DShape, _ = front.(*StemCylinder3DShape)
		}
		stage.OnAfterStemCylinder3DShapeUpdateCallback.OnAfterUpdate(stage, stemcylinder3dshape, frontStemCylinder3DShape)
	}
}

func (stemcylinder3dshape *StemCylinder3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStemCylinder3DShapeDeleteCallback != nil {
		var frontStemCylinder3DShape *StemCylinder3DShape
		if front != nil {
			frontStemCylinder3DShape, _ = front.(*StemCylinder3DShape)
		}
		stage.OnAfterStemCylinder3DShapeDeleteCallback.OnAfterDelete(stage, stemcylinder3dshape, frontStemCylinder3DShape)
	}
}

func (stool2ddiagram *Stool2DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStool2DDiagramCreateCallback != nil {
		stage.OnAfterStool2DDiagramCreateCallback.OnAfterCreate(stage, stool2ddiagram)
	}
}

func (stool2ddiagram *Stool2DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStool2DDiagramUpdateCallback != nil {
		var frontStool2DDiagram *Stool2DDiagram
		if front != nil {
			frontStool2DDiagram, _ = front.(*Stool2DDiagram)
		}
		stage.OnAfterStool2DDiagramUpdateCallback.OnAfterUpdate(stage, stool2ddiagram, frontStool2DDiagram)
	}
}

func (stool2ddiagram *Stool2DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStool2DDiagramDeleteCallback != nil {
		var frontStool2DDiagram *Stool2DDiagram
		if front != nil {
			frontStool2DDiagram, _ = front.(*Stool2DDiagram)
		}
		stage.OnAfterStool2DDiagramDeleteCallback.OnAfterDelete(stage, stool2ddiagram, frontStool2DDiagram)
	}
}

func (stool3ddiagram *Stool3DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStool3DDiagramCreateCallback != nil {
		stage.OnAfterStool3DDiagramCreateCallback.OnAfterCreate(stage, stool3ddiagram)
	}
}

func (stool3ddiagram *Stool3DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStool3DDiagramUpdateCallback != nil {
		var frontStool3DDiagram *Stool3DDiagram
		if front != nil {
			frontStool3DDiagram, _ = front.(*Stool3DDiagram)
		}
		stage.OnAfterStool3DDiagramUpdateCallback.OnAfterUpdate(stage, stool3ddiagram, frontStool3DDiagram)
	}
}

func (stool3ddiagram *Stool3DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStool3DDiagramDeleteCallback != nil {
		var frontStool3DDiagram *Stool3DDiagram
		if front != nil {
			frontStool3DDiagram, _ = front.(*Stool3DDiagram)
		}
		stage.OnAfterStool3DDiagramDeleteCallback.OnAfterDelete(stage, stool3ddiagram, frontStool3DDiagram)
	}
}

func (stoolabstract *StoolAbstract) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStoolAbstractCreateCallback != nil {
		stage.OnAfterStoolAbstractCreateCallback.OnAfterCreate(stage, stoolabstract)
	}
}

func (stoolabstract *StoolAbstract) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStoolAbstractUpdateCallback != nil {
		var frontStoolAbstract *StoolAbstract
		if front != nil {
			frontStoolAbstract, _ = front.(*StoolAbstract)
		}
		stage.OnAfterStoolAbstractUpdateCallback.OnAfterUpdate(stage, stoolabstract, frontStoolAbstract)
	}
}

func (stoolabstract *StoolAbstract) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStoolAbstractDeleteCallback != nil {
		var frontStoolAbstract *StoolAbstract
		if front != nil {
			frontStoolAbstract, _ = front.(*StoolAbstract)
		}
		stage.OnAfterStoolAbstractDeleteCallback.OnAfterDelete(stage, stoolabstract, frontStoolAbstract)
	}
}

func (tiledfloor3dshape *TiledFloor3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTiledFloor3DShapeCreateCallback != nil {
		stage.OnAfterTiledFloor3DShapeCreateCallback.OnAfterCreate(stage, tiledfloor3dshape)
	}
}

func (tiledfloor3dshape *TiledFloor3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTiledFloor3DShapeUpdateCallback != nil {
		var frontTiledFloor3DShape *TiledFloor3DShape
		if front != nil {
			frontTiledFloor3DShape, _ = front.(*TiledFloor3DShape)
		}
		stage.OnAfterTiledFloor3DShapeUpdateCallback.OnAfterUpdate(stage, tiledfloor3dshape, frontTiledFloor3DShape)
	}
}

func (tiledfloor3dshape *TiledFloor3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTiledFloor3DShapeDeleteCallback != nil {
		var frontTiledFloor3DShape *TiledFloor3DShape
		if front != nil {
			frontTiledFloor3DShape, _ = front.(*TiledFloor3DShape)
		}
		stage.OnAfterTiledFloor3DShapeDeleteCallback.OnAfterDelete(stage, tiledfloor3dshape, frontTiledFloor3DShape)
	}
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopCurvePlane1ShapeCreateCallback != nil {
		stage.OnAfterTopCurvePlane1ShapeCreateCallback.OnAfterCreate(stage, topcurveplane1shape)
	}
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopCurvePlane1ShapeUpdateCallback != nil {
		var frontTopCurvePlane1Shape *TopCurvePlane1Shape
		if front != nil {
			frontTopCurvePlane1Shape, _ = front.(*TopCurvePlane1Shape)
		}
		stage.OnAfterTopCurvePlane1ShapeUpdateCallback.OnAfterUpdate(stage, topcurveplane1shape, frontTopCurvePlane1Shape)
	}
}

func (topcurveplane1shape *TopCurvePlane1Shape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopCurvePlane1ShapeDeleteCallback != nil {
		var frontTopCurvePlane1Shape *TopCurvePlane1Shape
		if front != nil {
			frontTopCurvePlane1Shape, _ = front.(*TopCurvePlane1Shape)
		}
		stage.OnAfterTopCurvePlane1ShapeDeleteCallback.OnAfterDelete(stage, topcurveplane1shape, frontTopCurvePlane1Shape)
	}
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopCurvePlane2ShapeCreateCallback != nil {
		stage.OnAfterTopCurvePlane2ShapeCreateCallback.OnAfterCreate(stage, topcurveplane2shape)
	}
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopCurvePlane2ShapeUpdateCallback != nil {
		var frontTopCurvePlane2Shape *TopCurvePlane2Shape
		if front != nil {
			frontTopCurvePlane2Shape, _ = front.(*TopCurvePlane2Shape)
		}
		stage.OnAfterTopCurvePlane2ShapeUpdateCallback.OnAfterUpdate(stage, topcurveplane2shape, frontTopCurvePlane2Shape)
	}
}

func (topcurveplane2shape *TopCurvePlane2Shape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopCurvePlane2ShapeDeleteCallback != nil {
		var frontTopCurvePlane2Shape *TopCurvePlane2Shape
		if front != nil {
			frontTopCurvePlane2Shape, _ = front.(*TopCurvePlane2Shape)
		}
		stage.OnAfterTopCurvePlane2ShapeDeleteCallback.OnAfterDelete(stage, topcurveplane2shape, frontTopCurvePlane2Shape)
	}
}

func (topendarcshape *TopEndArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopEndArcShapeCreateCallback != nil {
		stage.OnAfterTopEndArcShapeCreateCallback.OnAfterCreate(stage, topendarcshape)
	}
}

func (topendarcshape *TopEndArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndArcShapeUpdateCallback != nil {
		var frontTopEndArcShape *TopEndArcShape
		if front != nil {
			frontTopEndArcShape, _ = front.(*TopEndArcShape)
		}
		stage.OnAfterTopEndArcShapeUpdateCallback.OnAfterUpdate(stage, topendarcshape, frontTopEndArcShape)
	}
}

func (topendarcshape *TopEndArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndArcShapeDeleteCallback != nil {
		var frontTopEndArcShape *TopEndArcShape
		if front != nil {
			frontTopEndArcShape, _ = front.(*TopEndArcShape)
		}
		stage.OnAfterTopEndArcShapeDeleteCallback.OnAfterDelete(stage, topendarcshape, frontTopEndArcShape)
	}
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopEndArcShapeGridCreateCallback != nil {
		stage.OnAfterTopEndArcShapeGridCreateCallback.OnAfterCreate(stage, topendarcshapegrid)
	}
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndArcShapeGridUpdateCallback != nil {
		var frontTopEndArcShapeGrid *TopEndArcShapeGrid
		if front != nil {
			frontTopEndArcShapeGrid, _ = front.(*TopEndArcShapeGrid)
		}
		stage.OnAfterTopEndArcShapeGridUpdateCallback.OnAfterUpdate(stage, topendarcshapegrid, frontTopEndArcShapeGrid)
	}
}

func (topendarcshapegrid *TopEndArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndArcShapeGridDeleteCallback != nil {
		var frontTopEndArcShapeGrid *TopEndArcShapeGrid
		if front != nil {
			frontTopEndArcShapeGrid, _ = front.(*TopEndArcShapeGrid)
		}
		stage.OnAfterTopEndArcShapeGridDeleteCallback.OnAfterDelete(stage, topendarcshapegrid, frontTopEndArcShapeGrid)
	}
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopEndHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterTopEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, topendhalfwayarcshape)
	}
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndHalfwayArcShapeUpdateCallback != nil {
		var frontTopEndHalfwayArcShape *TopEndHalfwayArcShape
		if front != nil {
			frontTopEndHalfwayArcShape, _ = front.(*TopEndHalfwayArcShape)
		}
		stage.OnAfterTopEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, topendhalfwayarcshape, frontTopEndHalfwayArcShape)
	}
}

func (topendhalfwayarcshape *TopEndHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndHalfwayArcShapeDeleteCallback != nil {
		var frontTopEndHalfwayArcShape *TopEndHalfwayArcShape
		if front != nil {
			frontTopEndHalfwayArcShape, _ = front.(*TopEndHalfwayArcShape)
		}
		stage.OnAfterTopEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, topendhalfwayarcshape, frontTopEndHalfwayArcShape)
	}
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopEndHalfwayArcShapeGridCreateCallback != nil {
		stage.OnAfterTopEndHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, topendhalfwayarcshapegrid)
	}
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndHalfwayArcShapeGridUpdateCallback != nil {
		var frontTopEndHalfwayArcShapeGrid *TopEndHalfwayArcShapeGrid
		if front != nil {
			frontTopEndHalfwayArcShapeGrid, _ = front.(*TopEndHalfwayArcShapeGrid)
		}
		stage.OnAfterTopEndHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, topendhalfwayarcshapegrid, frontTopEndHalfwayArcShapeGrid)
	}
}

func (topendhalfwayarcshapegrid *TopEndHalfwayArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopEndHalfwayArcShapeGridDeleteCallback != nil {
		var frontTopEndHalfwayArcShapeGrid *TopEndHalfwayArcShapeGrid
		if front != nil {
			frontTopEndHalfwayArcShapeGrid, _ = front.(*TopEndHalfwayArcShapeGrid)
		}
		stage.OnAfterTopEndHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, topendhalfwayarcshapegrid, frontTopEndHalfwayArcShapeGrid)
	}
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopGrowthCurve2DCreateCallback != nil {
		stage.OnAfterTopGrowthCurve2DCreateCallback.OnAfterCreate(stage, topgrowthcurve2d)
	}
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopGrowthCurve2DUpdateCallback != nil {
		var frontTopGrowthCurve2D *TopGrowthCurve2D
		if front != nil {
			frontTopGrowthCurve2D, _ = front.(*TopGrowthCurve2D)
		}
		stage.OnAfterTopGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, topgrowthcurve2d, frontTopGrowthCurve2D)
	}
}

func (topgrowthcurve2d *TopGrowthCurve2D) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopGrowthCurve2DDeleteCallback != nil {
		var frontTopGrowthCurve2D *TopGrowthCurve2D
		if front != nil {
			frontTopGrowthCurve2D, _ = front.(*TopGrowthCurve2D)
		}
		stage.OnAfterTopGrowthCurve2DDeleteCallback.OnAfterDelete(stage, topgrowthcurve2d, frontTopGrowthCurve2D)
	}
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopMidArcVectorShapeCreateCallback != nil {
		stage.OnAfterTopMidArcVectorShapeCreateCallback.OnAfterCreate(stage, topmidarcvectorshape)
	}
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopMidArcVectorShapeUpdateCallback != nil {
		var frontTopMidArcVectorShape *TopMidArcVectorShape
		if front != nil {
			frontTopMidArcVectorShape, _ = front.(*TopMidArcVectorShape)
		}
		stage.OnAfterTopMidArcVectorShapeUpdateCallback.OnAfterUpdate(stage, topmidarcvectorshape, frontTopMidArcVectorShape)
	}
}

func (topmidarcvectorshape *TopMidArcVectorShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopMidArcVectorShapeDeleteCallback != nil {
		var frontTopMidArcVectorShape *TopMidArcVectorShape
		if front != nil {
			frontTopMidArcVectorShape, _ = front.(*TopMidArcVectorShape)
		}
		stage.OnAfterTopMidArcVectorShapeDeleteCallback.OnAfterDelete(stage, topmidarcvectorshape, frontTopMidArcVectorShape)
	}
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopMidArcVectorShapeGridCreateCallback != nil {
		stage.OnAfterTopMidArcVectorShapeGridCreateCallback.OnAfterCreate(stage, topmidarcvectorshapegrid)
	}
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopMidArcVectorShapeGridUpdateCallback != nil {
		var frontTopMidArcVectorShapeGrid *TopMidArcVectorShapeGrid
		if front != nil {
			frontTopMidArcVectorShapeGrid, _ = front.(*TopMidArcVectorShapeGrid)
		}
		stage.OnAfterTopMidArcVectorShapeGridUpdateCallback.OnAfterUpdate(stage, topmidarcvectorshapegrid, frontTopMidArcVectorShapeGrid)
	}
}

func (topmidarcvectorshapegrid *TopMidArcVectorShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopMidArcVectorShapeGridDeleteCallback != nil {
		var frontTopMidArcVectorShapeGrid *TopMidArcVectorShapeGrid
		if front != nil {
			frontTopMidArcVectorShapeGrid, _ = front.(*TopMidArcVectorShapeGrid)
		}
		stage.OnAfterTopMidArcVectorShapeGridDeleteCallback.OnAfterDelete(stage, topmidarcvectorshapegrid, frontTopMidArcVectorShapeGrid)
	}
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeCreateCallback.OnAfterCreate(stage, topstackgrowthcurve2dendhalfwayarcshape)
	}
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback != nil {
		var frontTopStackGrowthCurve2DEndHalfwayArcShape *TopStackGrowthCurve2DEndHalfwayArcShape
		if front != nil {
			frontTopStackGrowthCurve2DEndHalfwayArcShape, _ = front.(*TopStackGrowthCurve2DEndHalfwayArcShape)
		}
		stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, topstackgrowthcurve2dendhalfwayarcshape, frontTopStackGrowthCurve2DEndHalfwayArcShape)
	}
}

func (topstackgrowthcurve2dendhalfwayarcshape *TopStackGrowthCurve2DEndHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback != nil {
		var frontTopStackGrowthCurve2DEndHalfwayArcShape *TopStackGrowthCurve2DEndHalfwayArcShape
		if front != nil {
			frontTopStackGrowthCurve2DEndHalfwayArcShape, _ = front.(*TopStackGrowthCurve2DEndHalfwayArcShape)
		}
		stage.OnAfterTopStackGrowthCurve2DEndHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, topstackgrowthcurve2dendhalfwayarcshape, frontTopStackGrowthCurve2DEndHalfwayArcShape)
	}
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, topstackgrowthcurve2dstarthalfwayarcshape)
	}
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback != nil {
		var frontTopStackGrowthCurve2DStartHalfwayArcShape *TopStackGrowthCurve2DStartHalfwayArcShape
		if front != nil {
			frontTopStackGrowthCurve2DStartHalfwayArcShape, _ = front.(*TopStackGrowthCurve2DStartHalfwayArcShape)
		}
		stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, topstackgrowthcurve2dstarthalfwayarcshape, frontTopStackGrowthCurve2DStartHalfwayArcShape)
	}
}

func (topstackgrowthcurve2dstarthalfwayarcshape *TopStackGrowthCurve2DStartHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback != nil {
		var frontTopStackGrowthCurve2DStartHalfwayArcShape *TopStackGrowthCurve2DStartHalfwayArcShape
		if front != nil {
			frontTopStackGrowthCurve2DStartHalfwayArcShape, _ = front.(*TopStackGrowthCurve2DStartHalfwayArcShape)
		}
		stage.OnAfterTopStackGrowthCurve2DStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, topstackgrowthcurve2dstarthalfwayarcshape, frontTopStackGrowthCurve2DStartHalfwayArcShape)
	}
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStackOfGrowthCurve2DCreateCallback != nil {
		stage.OnAfterTopStackOfGrowthCurve2DCreateCallback.OnAfterCreate(stage, topstackofgrowthcurve2d)
	}
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfGrowthCurve2DUpdateCallback != nil {
		var frontTopStackOfGrowthCurve2D *TopStackOfGrowthCurve2D
		if front != nil {
			frontTopStackOfGrowthCurve2D, _ = front.(*TopStackOfGrowthCurve2D)
		}
		stage.OnAfterTopStackOfGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, topstackofgrowthcurve2d, frontTopStackOfGrowthCurve2D)
	}
}

func (topstackofgrowthcurve2d *TopStackOfGrowthCurve2D) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfGrowthCurve2DDeleteCallback != nil {
		var frontTopStackOfGrowthCurve2D *TopStackOfGrowthCurve2D
		if front != nil {
			frontTopStackOfGrowthCurve2D, _ = front.(*TopStackOfGrowthCurve2D)
		}
		stage.OnAfterTopStackOfGrowthCurve2DDeleteCallback.OnAfterDelete(stage, topstackofgrowthcurve2d, frontTopStackOfGrowthCurve2D)
	}
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DCreateCallback != nil {
		stage.OnAfterTopStackOfRotatedGrowthCurve2DCreateCallback.OnAfterCreate(stage, topstackofrotatedgrowthcurve2d)
	}
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DUpdateCallback != nil {
		var frontTopStackOfRotatedGrowthCurve2D *TopStackOfRotatedGrowthCurve2D
		if front != nil {
			frontTopStackOfRotatedGrowthCurve2D, _ = front.(*TopStackOfRotatedGrowthCurve2D)
		}
		stage.OnAfterTopStackOfRotatedGrowthCurve2DUpdateCallback.OnAfterUpdate(stage, topstackofrotatedgrowthcurve2d, frontTopStackOfRotatedGrowthCurve2D)
	}
}

func (topstackofrotatedgrowthcurve2d *TopStackOfRotatedGrowthCurve2D) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DDeleteCallback != nil {
		var frontTopStackOfRotatedGrowthCurve2D *TopStackOfRotatedGrowthCurve2D
		if front != nil {
			frontTopStackOfRotatedGrowthCurve2D, _ = front.(*TopStackOfRotatedGrowthCurve2D)
		}
		stage.OnAfterTopStackOfRotatedGrowthCurve2DDeleteCallback.OnAfterDelete(stage, topstackofrotatedgrowthcurve2d, frontTopStackOfRotatedGrowthCurve2D)
	}
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeCreateCallback != nil {
		stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeCreateCallback.OnAfterCreate(stage, topstackofrotatedgrowthcurve2dendarcshape)
	}
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeUpdateCallback != nil {
		var frontTopStackOfRotatedGrowthCurve2DEndArcShape *TopStackOfRotatedGrowthCurve2DEndArcShape
		if front != nil {
			frontTopStackOfRotatedGrowthCurve2DEndArcShape, _ = front.(*TopStackOfRotatedGrowthCurve2DEndArcShape)
		}
		stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeUpdateCallback.OnAfterUpdate(stage, topstackofrotatedgrowthcurve2dendarcshape, frontTopStackOfRotatedGrowthCurve2DEndArcShape)
	}
}

func (topstackofrotatedgrowthcurve2dendarcshape *TopStackOfRotatedGrowthCurve2DEndArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeDeleteCallback != nil {
		var frontTopStackOfRotatedGrowthCurve2DEndArcShape *TopStackOfRotatedGrowthCurve2DEndArcShape
		if front != nil {
			frontTopStackOfRotatedGrowthCurve2DEndArcShape, _ = front.(*TopStackOfRotatedGrowthCurve2DEndArcShape)
		}
		stage.OnAfterTopStackOfRotatedGrowthCurve2DEndArcShapeDeleteCallback.OnAfterDelete(stage, topstackofrotatedgrowthcurve2dendarcshape, frontTopStackOfRotatedGrowthCurve2DEndArcShape)
	}
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeCreateCallback != nil {
		stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeCreateCallback.OnAfterCreate(stage, topstackofrotatedgrowthcurve2dstartarcshape)
	}
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeUpdateCallback != nil {
		var frontTopStackOfRotatedGrowthCurve2DStartArcShape *TopStackOfRotatedGrowthCurve2DStartArcShape
		if front != nil {
			frontTopStackOfRotatedGrowthCurve2DStartArcShape, _ = front.(*TopStackOfRotatedGrowthCurve2DStartArcShape)
		}
		stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeUpdateCallback.OnAfterUpdate(stage, topstackofrotatedgrowthcurve2dstartarcshape, frontTopStackOfRotatedGrowthCurve2DStartArcShape)
	}
}

func (topstackofrotatedgrowthcurve2dstartarcshape *TopStackOfRotatedGrowthCurve2DStartArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeDeleteCallback != nil {
		var frontTopStackOfRotatedGrowthCurve2DStartArcShape *TopStackOfRotatedGrowthCurve2DStartArcShape
		if front != nil {
			frontTopStackOfRotatedGrowthCurve2DStartArcShape, _ = front.(*TopStackOfRotatedGrowthCurve2DStartArcShape)
		}
		stage.OnAfterTopStackOfRotatedGrowthCurve2DStartArcShapeDeleteCallback.OnAfterDelete(stage, topstackofrotatedgrowthcurve2dstartarcshape, frontTopStackOfRotatedGrowthCurve2DStartArcShape)
	}
}

func (topstartarcshape *TopStartArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStartArcShapeCreateCallback != nil {
		stage.OnAfterTopStartArcShapeCreateCallback.OnAfterCreate(stage, topstartarcshape)
	}
}

func (topstartarcshape *TopStartArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartArcShapeUpdateCallback != nil {
		var frontTopStartArcShape *TopStartArcShape
		if front != nil {
			frontTopStartArcShape, _ = front.(*TopStartArcShape)
		}
		stage.OnAfterTopStartArcShapeUpdateCallback.OnAfterUpdate(stage, topstartarcshape, frontTopStartArcShape)
	}
}

func (topstartarcshape *TopStartArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartArcShapeDeleteCallback != nil {
		var frontTopStartArcShape *TopStartArcShape
		if front != nil {
			frontTopStartArcShape, _ = front.(*TopStartArcShape)
		}
		stage.OnAfterTopStartArcShapeDeleteCallback.OnAfterDelete(stage, topstartarcshape, frontTopStartArcShape)
	}
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStartArcShapeGridCreateCallback != nil {
		stage.OnAfterTopStartArcShapeGridCreateCallback.OnAfterCreate(stage, topstartarcshapegrid)
	}
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartArcShapeGridUpdateCallback != nil {
		var frontTopStartArcShapeGrid *TopStartArcShapeGrid
		if front != nil {
			frontTopStartArcShapeGrid, _ = front.(*TopStartArcShapeGrid)
		}
		stage.OnAfterTopStartArcShapeGridUpdateCallback.OnAfterUpdate(stage, topstartarcshapegrid, frontTopStartArcShapeGrid)
	}
}

func (topstartarcshapegrid *TopStartArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartArcShapeGridDeleteCallback != nil {
		var frontTopStartArcShapeGrid *TopStartArcShapeGrid
		if front != nil {
			frontTopStartArcShapeGrid, _ = front.(*TopStartArcShapeGrid)
		}
		stage.OnAfterTopStartArcShapeGridDeleteCallback.OnAfterDelete(stage, topstartarcshapegrid, frontTopStartArcShapeGrid)
	}
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStartHalfwayArcShapeCreateCallback != nil {
		stage.OnAfterTopStartHalfwayArcShapeCreateCallback.OnAfterCreate(stage, topstarthalfwayarcshape)
	}
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartHalfwayArcShapeUpdateCallback != nil {
		var frontTopStartHalfwayArcShape *TopStartHalfwayArcShape
		if front != nil {
			frontTopStartHalfwayArcShape, _ = front.(*TopStartHalfwayArcShape)
		}
		stage.OnAfterTopStartHalfwayArcShapeUpdateCallback.OnAfterUpdate(stage, topstarthalfwayarcshape, frontTopStartHalfwayArcShape)
	}
}

func (topstarthalfwayarcshape *TopStartHalfwayArcShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartHalfwayArcShapeDeleteCallback != nil {
		var frontTopStartHalfwayArcShape *TopStartHalfwayArcShape
		if front != nil {
			frontTopStartHalfwayArcShape, _ = front.(*TopStartHalfwayArcShape)
		}
		stage.OnAfterTopStartHalfwayArcShapeDeleteCallback.OnAfterDelete(stage, topstarthalfwayarcshape, frontTopStartHalfwayArcShape)
	}
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTopStartHalfwayArcShapeGridCreateCallback != nil {
		stage.OnAfterTopStartHalfwayArcShapeGridCreateCallback.OnAfterCreate(stage, topstarthalfwayarcshapegrid)
	}
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartHalfwayArcShapeGridUpdateCallback != nil {
		var frontTopStartHalfwayArcShapeGrid *TopStartHalfwayArcShapeGrid
		if front != nil {
			frontTopStartHalfwayArcShapeGrid, _ = front.(*TopStartHalfwayArcShapeGrid)
		}
		stage.OnAfterTopStartHalfwayArcShapeGridUpdateCallback.OnAfterUpdate(stage, topstarthalfwayarcshapegrid, frontTopStartHalfwayArcShapeGrid)
	}
}

func (topstarthalfwayarcshapegrid *TopStartHalfwayArcShapeGrid) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTopStartHalfwayArcShapeGridDeleteCallback != nil {
		var frontTopStartHalfwayArcShapeGrid *TopStartHalfwayArcShapeGrid
		if front != nil {
			frontTopStartHalfwayArcShapeGrid, _ = front.(*TopStartHalfwayArcShapeGrid)
		}
		stage.OnAfterTopStartHalfwayArcShapeGridDeleteCallback.OnAfterDelete(stage, topstarthalfwayarcshapegrid, frontTopStartHalfwayArcShapeGrid)
	}
}

func (torus3dshape *Torus3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTorus3DShapeCreateCallback != nil {
		stage.OnAfterTorus3DShapeCreateCallback.OnAfterCreate(stage, torus3dshape)
	}
}

func (torus3dshape *Torus3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorus3DShapeUpdateCallback != nil {
		var frontTorus3DShape *Torus3DShape
		if front != nil {
			frontTorus3DShape, _ = front.(*Torus3DShape)
		}
		stage.OnAfterTorus3DShapeUpdateCallback.OnAfterUpdate(stage, torus3dshape, frontTorus3DShape)
	}
}

func (torus3dshape *Torus3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorus3DShapeDeleteCallback != nil {
		var frontTorus3DShape *Torus3DShape
		if front != nil {
			frontTorus3DShape, _ = front.(*Torus3DShape)
		}
		stage.OnAfterTorus3DShapeDeleteCallback.OnAfterDelete(stage, torus3dshape, frontTorus3DShape)
	}
}

func (torusedge3dshape *TorusEdge3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTorusEdge3DShapeCreateCallback != nil {
		stage.OnAfterTorusEdge3DShapeCreateCallback.OnAfterCreate(stage, torusedge3dshape)
	}
}

func (torusedge3dshape *TorusEdge3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorusEdge3DShapeUpdateCallback != nil {
		var frontTorusEdge3DShape *TorusEdge3DShape
		if front != nil {
			frontTorusEdge3DShape, _ = front.(*TorusEdge3DShape)
		}
		stage.OnAfterTorusEdge3DShapeUpdateCallback.OnAfterUpdate(stage, torusedge3dshape, frontTorusEdge3DShape)
	}
}

func (torusedge3dshape *TorusEdge3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorusEdge3DShapeDeleteCallback != nil {
		var frontTorusEdge3DShape *TorusEdge3DShape
		if front != nil {
			frontTorusEdge3DShape, _ = front.(*TorusEdge3DShape)
		}
		stage.OnAfterTorusEdge3DShapeDeleteCallback.OnAfterDelete(stage, torusedge3dshape, frontTorusEdge3DShape)
	}
}

func (torusstackshape *TorusStackShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTorusStackShapeCreateCallback != nil {
		stage.OnAfterTorusStackShapeCreateCallback.OnAfterCreate(stage, torusstackshape)
	}
}

func (torusstackshape *TorusStackShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorusStackShapeUpdateCallback != nil {
		var frontTorusStackShape *TorusStackShape
		if front != nil {
			frontTorusStackShape, _ = front.(*TorusStackShape)
		}
		stage.OnAfterTorusStackShapeUpdateCallback.OnAfterUpdate(stage, torusstackshape, frontTorusStackShape)
	}
}

func (torusstackshape *TorusStackShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTorusStackShapeDeleteCallback != nil {
		var frontTorusStackShape *TorusStackShape
		if front != nil {
			frontTorusStackShape, _ = front.(*TorusStackShape)
		}
		stage.OnAfterTorusStackShapeDeleteCallback.OnAfterDelete(stage, torusstackshape, frontTorusStackShape)
	}
}

func (trapezevolume3dshape *TrapezeVolume3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTrapezeVolume3DShapeCreateCallback != nil {
		stage.OnAfterTrapezeVolume3DShapeCreateCallback.OnAfterCreate(stage, trapezevolume3dshape)
	}
}

func (trapezevolume3dshape *TrapezeVolume3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTrapezeVolume3DShapeUpdateCallback != nil {
		var frontTrapezeVolume3DShape *TrapezeVolume3DShape
		if front != nil {
			frontTrapezeVolume3DShape, _ = front.(*TrapezeVolume3DShape)
		}
		stage.OnAfterTrapezeVolume3DShapeUpdateCallback.OnAfterUpdate(stage, trapezevolume3dshape, frontTrapezeVolume3DShape)
	}
}

func (trapezevolume3dshape *TrapezeVolume3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTrapezeVolume3DShapeDeleteCallback != nil {
		var frontTrapezeVolume3DShape *TrapezeVolume3DShape
		if front != nil {
			frontTrapezeVolume3DShape, _ = front.(*TrapezeVolume3DShape)
		}
		stage.OnAfterTrapezeVolume3DShapeDeleteCallback.OnAfterDelete(stage, trapezevolume3dshape, frontTrapezeVolume3DShape)
	}
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTubeVase3DDiagramCreateCallback != nil {
		stage.OnAfterTubeVase3DDiagramCreateCallback.OnAfterCreate(stage, tubevase3ddiagram)
	}
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTubeVase3DDiagramUpdateCallback != nil {
		var frontTubeVase3DDiagram *TubeVase3DDiagram
		if front != nil {
			frontTubeVase3DDiagram, _ = front.(*TubeVase3DDiagram)
		}
		stage.OnAfterTubeVase3DDiagramUpdateCallback.OnAfterUpdate(stage, tubevase3ddiagram, frontTubeVase3DDiagram)
	}
}

func (tubevase3ddiagram *TubeVase3DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTubeVase3DDiagramDeleteCallback != nil {
		var frontTubeVase3DDiagram *TubeVase3DDiagram
		if front != nil {
			frontTubeVase3DDiagram, _ = front.(*TubeVase3DDiagram)
		}
		stage.OnAfterTubeVase3DDiagramDeleteCallback.OnAfterDelete(stage, tubevase3ddiagram, frontTubeVase3DDiagram)
	}
}

func (tubevaseabstract *TubeVaseAbstract) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTubeVaseAbstractCreateCallback != nil {
		stage.OnAfterTubeVaseAbstractCreateCallback.OnAfterCreate(stage, tubevaseabstract)
	}
}

func (tubevaseabstract *TubeVaseAbstract) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTubeVaseAbstractUpdateCallback != nil {
		var frontTubeVaseAbstract *TubeVaseAbstract
		if front != nil {
			frontTubeVaseAbstract, _ = front.(*TubeVaseAbstract)
		}
		stage.OnAfterTubeVaseAbstractUpdateCallback.OnAfterUpdate(stage, tubevaseabstract, frontTubeVaseAbstract)
	}
}

func (tubevaseabstract *TubeVaseAbstract) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTubeVaseAbstractDeleteCallback != nil {
		var frontTubeVaseAbstract *TubeVaseAbstract
		if front != nil {
			frontTubeVaseAbstract, _ = front.(*TubeVaseAbstract)
		}
		stage.OnAfterTubeVaseAbstractDeleteCallback.OnAfterDelete(stage, tubevaseabstract, frontTubeVaseAbstract)
	}
}

func (vase2ddiagram *Vase2DDiagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVase2DDiagramCreateCallback != nil {
		stage.OnAfterVase2DDiagramCreateCallback.OnAfterCreate(stage, vase2ddiagram)
	}
}

func (vase2ddiagram *Vase2DDiagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVase2DDiagramUpdateCallback != nil {
		var frontVase2DDiagram *Vase2DDiagram
		if front != nil {
			frontVase2DDiagram, _ = front.(*Vase2DDiagram)
		}
		stage.OnAfterVase2DDiagramUpdateCallback.OnAfterUpdate(stage, vase2ddiagram, frontVase2DDiagram)
	}
}

func (vase2ddiagram *Vase2DDiagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVase2DDiagramDeleteCallback != nil {
		var frontVase2DDiagram *Vase2DDiagram
		if front != nil {
			frontVase2DDiagram, _ = front.(*Vase2DDiagram)
		}
		stage.OnAfterVase2DDiagramDeleteCallback.OnAfterDelete(stage, vase2ddiagram, frontVase2DDiagram)
	}
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVerticalTorusStackShapeCreateCallback != nil {
		stage.OnAfterVerticalTorusStackShapeCreateCallback.OnAfterCreate(stage, verticaltorusstackshape)
	}
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalTorusStackShapeUpdateCallback != nil {
		var frontVerticalTorusStackShape *VerticalTorusStackShape
		if front != nil {
			frontVerticalTorusStackShape, _ = front.(*VerticalTorusStackShape)
		}
		stage.OnAfterVerticalTorusStackShapeUpdateCallback.OnAfterUpdate(stage, verticaltorusstackshape, frontVerticalTorusStackShape)
	}
}

func (verticaltorusstackshape *VerticalTorusStackShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVerticalTorusStackShapeDeleteCallback != nil {
		var frontVerticalTorusStackShape *VerticalTorusStackShape
		if front != nil {
			frontVerticalTorusStackShape, _ = front.(*VerticalTorusStackShape)
		}
		stage.OnAfterVerticalTorusStackShapeDeleteCallback.OnAfterDelete(stage, verticaltorusstackshape, frontVerticalTorusStackShape)
	}
}

func (volumekey3dshape *VolumeKey3DShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterVolumeKey3DShapeCreateCallback != nil {
		stage.OnAfterVolumeKey3DShapeCreateCallback.OnAfterCreate(stage, volumekey3dshape)
	}
}

func (volumekey3dshape *VolumeKey3DShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVolumeKey3DShapeUpdateCallback != nil {
		var frontVolumeKey3DShape *VolumeKey3DShape
		if front != nil {
			frontVolumeKey3DShape, _ = front.(*VolumeKey3DShape)
		}
		stage.OnAfterVolumeKey3DShapeUpdateCallback.OnAfterUpdate(stage, volumekey3dshape, frontVolumeKey3DShape)
	}
}

func (volumekey3dshape *VolumeKey3DShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterVolumeKey3DShapeDeleteCallback != nil {
		var frontVolumeKey3DShape *VolumeKey3DShape
		if front != nil {
			frontVolumeKey3DShape, _ = front.(*VolumeKey3DShape)
		}
		stage.OnAfterVolumeKey3DShapeDeleteCallback.OnAfterDelete(stage, volumekey3dshape, frontVolumeKey3DShape)
	}
}

