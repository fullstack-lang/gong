package models

import "fmt"

func ensureDiagramShape[T PointerToGongstruct](stager *Stager, diagramName string, shapePtr *T, shapeName string, newShape func() T) bool {
	var zero T
	if *shapePtr == zero {
		s := newShape()
		s.SetName(diagramName + "-" + shapeName)
		s.StageVoid(stager.stage)
		*shapePtr = s
		stager.logAndNotify(fmt.Sprintf("Diagram %s: created missing %s", diagramName, shapeName))
		return true
	}
	return false
}

func (stager *Stager) enforceDiagramShapes() bool {
	modified := false
	stage := stager.stage

	for diagram := range *stage.GetInstancesSet[*TubeVase3DDiagram]() {
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Rendered3DShape, "Rendered3DShape", func() *Rendered3DShape { return new(Rendered3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TorusStackShape, "TorusStackShape", func() *TorusStackShape { return new(TorusStackShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.VerticalTorusStackShape, "VerticalTorusStackShape", func() *VerticalTorusStackShape { return new(VerticalTorusStackShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.PartiallyRotatedTorusShape, "PartiallyRotatedTorusShape", func() *PartiallyRotatedTorusShape { return new(PartiallyRotatedTorusShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.StackOfPartiallyRotatedTorusShape, "StackOfPartiallyRotatedTorusShape", func() *StackOfPartiallyRotatedTorusShape { return new(StackOfPartiallyRotatedTorusShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.PointsAndLines3DShape, "PointsAndLines3DShape", func() *PointsAndLines3DShape { return new(PointsAndLines3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.SampledPoints3DShape, "SampledPoints3DShape", func() *SampledPoints3DShape { return new(SampledPoints3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.OriginalPoints3DShape, "OriginalPoints3DShape", func() *OriginalPoints3DShape { return new(OriginalPoints3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Angle0Shape, "Angle0Shape", func() *Angle0Shape { return new(Angle0Shape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.KeyHole3DShape, "KeyHole3DShape", func() *KeyHole3DShape { return new(KeyHole3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Key3DShape, "Key3DShape", func() *Key3DShape { return new(Key3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.VolumeKey3DShape, "VolumeKey3DShape", func() *VolumeKey3DShape { return new(VolumeKey3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TorusEdge3DShape, "TorusEdge3DShape", func() *TorusEdge3DShape { return new(TorusEdge3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TiledFloor3DShape, "TiledFloor3DShape", func() *TiledFloor3DShape { return new(TiledFloor3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TopCurvePlane1Shape, "TopCurvePlane1Shape", func() *TopCurvePlane1Shape { return new(TopCurvePlane1Shape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.BottomCurvePlane1Shape, "BottomCurvePlane1Shape", func() *BottomCurvePlane1Shape { return new(BottomCurvePlane1Shape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TopCurvePlane2Shape, "TopCurvePlane2Shape", func() *TopCurvePlane2Shape { return new(TopCurvePlane2Shape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.BottomCurvePlane2Shape, "BottomCurvePlane2Shape", func() *BottomCurvePlane2Shape { return new(BottomCurvePlane2Shape) }) || modified
	}

	for diagram := range *stage.GetInstancesSet[*Stool3DDiagram]() {
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Rendered3DShape, "Rendered3DShape", func() *Rendered3DShape { return new(Rendered3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.SeatTopCurveShape, "SeatTopCurveShape", func() *SeatTopCurveShape { return new(SeatTopCurveShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.RotatedSeatTopCurveShape, "RotatedSeatTopCurveShape", func() *PartiallyRotatedSeatTopCurveShape { return new(PartiallyRotatedSeatTopCurveShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.SeatBottomCurveShape, "SeatBottomCurveShape", func() *SeatBottomCurveShape { return new(SeatBottomCurveShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.RotatedSeatBottomCurveShape, "RotatedSeatBottomCurveShape", func() *PartiallyRotatedSeatBottomCurveShape { return new(PartiallyRotatedSeatBottomCurveShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Torus3DShape, "Torus3DShape", func() *Torus3DShape { return new(Torus3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.RotatedTorusShape, "RotatedTorusShape", func() *PartiallyRotatedTorusShape { return new(PartiallyRotatedTorusShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.SampledPoints3DShape, "SampledPoints3DShape", func() *SampledPoints3DShape { return new(SampledPoints3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.RotatedSampledPoints3DShape, "RotatedSampledPoints3DShape", func() *RotatedSampledPoints3DShape { return new(RotatedSampledPoints3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.EyeSampledPoints3DShape, "EyeSampledPoints3DShape", func() *EyeSampledPoints3DShape { return new(EyeSampledPoints3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.EyeCornersSampledPoints3DShape, "EyeCornersSampledPoints3DShape", func() *EyeCornersSampledPoints3DShape { return new(EyeCornersSampledPoints3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Eye3DShape, "Eye3DShape", func() *Eye3DShape { return new(Eye3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.EyeSeatBottomCurveShape, "EyeSeatBottomCurveShape", func() *EyeSeatBottomCurveShape { return new(EyeSeatBottomCurveShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.EyeStoolBottomCurveShape, "EyeStoolBottomCurveShape", func() *EyeStoolBottomCurveShape { return new(EyeStoolBottomCurveShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Seat3DShape, "Seat3DShape", func() *Seat3DShape { return new(Seat3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.EyeVolume3DShape, "EyeVolume3DShape", func() *EyeVolume3DShape { return new(EyeVolume3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.SeatAndLegs3DShape, "SeatAndLegs3DShape", func() *SeatAndLegs3DShape { return new(SeatAndLegs3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.RotatedSeatAndLegs3DShape, "RotatedSeatAndLegs3DShape", func() *RotatedSeatAndLegs3DShape { return new(RotatedSeatAndLegs3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TiledFloor3DShape, "TiledFloor3DShape", func() *TiledFloor3DShape { return new(TiledFloor3DShape) }) || modified
	}

	for diagram := range *stage.GetInstancesSet[*Clock3DDiagram]() {
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Rendered3DShape, "Rendered3DShape", func() *Rendered3DShape { return new(Rendered3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.ClockTopCurveShape, "ClockTopCurveShape", func() *ClockTopCurveShape { return new(ClockTopCurveShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Torus3DShape, "Torus3DShape", func() *Torus3DShape { return new(Torus3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.SampledPoints3DShape, "SampledPoints3DShape", func() *SampledPoints3DShape { return new(SampledPoints3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TiledFloor3DShape, "TiledFloor3DShape", func() *TiledFloor3DShape { return new(TiledFloor3DShape) }) || modified
	}

	for diagram := range *stage.GetInstancesSet[*Plant3DDiagram]() {
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Rendered3DShape, "Rendered3DShape", func() *Rendered3DShape { return new(Rendered3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.StemCylinder3DShape, "StemCylinder3DShape", func() *StemCylinder3DShape { return &StemCylinder3DShape{Transparency: 0.35} }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.ParastichyNCurves3DShape, "ParastichyNCurves3DShape", func() *ParastichyNCurves3DShape { return new(ParastichyNCurves3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.ParastichyMCurves3DShape, "ParastichyMCurves3DShape", func() *ParastichyMCurves3DShape { return new(ParastichyMCurves3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.CutLine3DShape, "CutLine3DShape", func() *CutLine3DShape { return new(CutLine3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Circumference3DShape, "Circumference3DShape", func() *Circumference3DShape { return new(Circumference3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.TiledFloor3DShape, "TiledFloor3DShape", func() *TiledFloor3DShape { return new(TiledFloor3DShape) }) || modified
		modified = ensureDiagramShape(stager, diagram.Name, &diagram.Leaves3DShape, "Leaves3DShape", func() *Leaves3DShape { return new(Leaves3DShape) }) || modified
	}

	return modified
}
