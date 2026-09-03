package models

func (stager *Stager) enforceDiagramShapes() bool {
	modified := false
	stage := stager.stage

	for diagram := range *GetGongstructInstancesSetFromPointerType[*Vase3DDiagram](stage) {
		if diagram.Rendered3DShape == nil { diagram.Rendered3DShape = (&Rendered3DShape{Name: diagram.Name + "-Rendered3DShape"}).Stage(stage); modified = true }
		if diagram.TorusStackShape == nil { diagram.TorusStackShape = (&TorusStackShape{Name: diagram.Name + "-TorusStackShape"}).Stage(stage); modified = true }
		if diagram.VerticalTorusStackShape == nil { diagram.VerticalTorusStackShape = (&VerticalTorusStackShape{Name: diagram.Name + "-VerticalTorusStackShape"}).Stage(stage); modified = true }
		if diagram.PartiallyRotatedTorusShape == nil { diagram.PartiallyRotatedTorusShape = (&PartiallyRotatedTorusShape{Name: diagram.Name + "-PartiallyRotatedTorusShape"}).Stage(stage); modified = true }
		if diagram.StackOfPartiallyRotatedTorusShape == nil { diagram.StackOfPartiallyRotatedTorusShape = (&StackOfPartiallyRotatedTorusShape{Name: diagram.Name + "-StackOfPartiallyRotatedTorusShape"}).Stage(stage); modified = true }
		if diagram.PointsAndLines3DShape == nil { diagram.PointsAndLines3DShape = (&PointsAndLines3DShape{Name: diagram.Name + "-PointsAndLines3DShape"}).Stage(stage); modified = true }
		if diagram.SampledPoints3DShape == nil { diagram.SampledPoints3DShape = (&SampledPoints3DShape{Name: diagram.Name + "-SampledPoints3DShape"}).Stage(stage); modified = true }
		if diagram.OriginalPoints3DShape == nil { diagram.OriginalPoints3DShape = (&OriginalPoints3DShape{Name: diagram.Name + "-OriginalPoints3DShape"}).Stage(stage); modified = true }
		if diagram.Angle0Shape == nil { diagram.Angle0Shape = (&Angle0Shape{Name: diagram.Name + "-Angle0Shape"}).Stage(stage); modified = true }
		if diagram.KeyHole3DShape == nil { diagram.KeyHole3DShape = (&KeyHole3DShape{Name: diagram.Name + "-KeyHole3DShape"}).Stage(stage); modified = true }
		if diagram.Key3DShape == nil { diagram.Key3DShape = (&Key3DShape{Name: diagram.Name + "-Key3DShape"}).Stage(stage); modified = true }
		if diagram.VolumeKey3DShape == nil { diagram.VolumeKey3DShape = (&VolumeKey3DShape{Name: diagram.Name + "-VolumeKey3DShape"}).Stage(stage); modified = true }
		if diagram.TorusEdge3DShape == nil { diagram.TorusEdge3DShape = (&TorusEdge3DShape{Name: diagram.Name + "-TorusEdge3DShape"}).Stage(stage); modified = true }
		if diagram.TiledFloor3DShape == nil { diagram.TiledFloor3DShape = (&TiledFloor3DShape{Name: diagram.Name + "-TiledFloor3DShape"}).Stage(stage); modified = true }
	}

	for diagram := range *GetGongstructInstancesSetFromPointerType[*Stool3DDiagram](stage) {
		if diagram.Rendered3DShape == nil { diagram.Rendered3DShape = (&Rendered3DShape{Name: diagram.Name + "-Rendered3DShape"}).Stage(stage); modified = true }
		if diagram.SeatTopCurveShape == nil { diagram.SeatTopCurveShape = (&SeatTopCurveShape{Name: diagram.Name + "-SeatTopCurveShape"}).Stage(stage); modified = true }
		if diagram.RotatedSeatTopCurveShape == nil { diagram.RotatedSeatTopCurveShape = (&PartiallyRotatedSeatTopCurveShape{Name: diagram.Name + "-RotatedSeatTopCurveShape"}).Stage(stage); modified = true }
		if diagram.SeatBottomCurveShape == nil { diagram.SeatBottomCurveShape = (&SeatBottomCurveShape{Name: diagram.Name + "-SeatBottomCurveShape"}).Stage(stage); modified = true }
		if diagram.RotatedSeatBottomCurveShape == nil { diagram.RotatedSeatBottomCurveShape = (&PartiallyRotatedSeatBottomCurveShape{Name: diagram.Name + "-RotatedSeatBottomCurveShape"}).Stage(stage); modified = true }
		if diagram.Torus3DShape == nil { diagram.Torus3DShape = (&Torus3DShape{Name: diagram.Name + "-Torus3DShape"}).Stage(stage); modified = true }
		if diagram.RotatedTorusShape == nil { diagram.RotatedTorusShape = (&PartiallyRotatedTorusShape{Name: diagram.Name + "-RotatedTorusShape"}).Stage(stage); modified = true }
		if diagram.SampledPoints3DShape == nil { diagram.SampledPoints3DShape = (&SampledPoints3DShape{Name: diagram.Name + "-SampledPoints3DShape"}).Stage(stage); modified = true }
		if diagram.RotatedSampledPoints3DShape == nil { diagram.RotatedSampledPoints3DShape = (&RotatedSampledPoints3DShape{Name: diagram.Name + "-RotatedSampledPoints3DShape"}).Stage(stage); modified = true }
		if diagram.EyeSampledPoints3DShape == nil { diagram.EyeSampledPoints3DShape = (&EyeSampledPoints3DShape{Name: diagram.Name + "-EyeSampledPoints3DShape"}).Stage(stage); modified = true }
		if diagram.EyeCornersSampledPoints3DShape == nil { diagram.EyeCornersSampledPoints3DShape = (&EyeCornersSampledPoints3DShape{Name: diagram.Name + "-EyeCornersSampledPoints3DShape"}).Stage(stage); modified = true }
		if diagram.Eye3DShape == nil { diagram.Eye3DShape = (&Eye3DShape{Name: diagram.Name + "-Eye3DShape"}).Stage(stage); modified = true }
		if diagram.EyeSeatBottomCurveShape == nil { diagram.EyeSeatBottomCurveShape = (&EyeSeatBottomCurveShape{Name: diagram.Name + "-EyeSeatBottomCurveShape"}).Stage(stage); modified = true }
		if diagram.EyeStoolBottomCurveShape == nil { diagram.EyeStoolBottomCurveShape = (&EyeStoolBottomCurveShape{Name: diagram.Name + "-EyeStoolBottomCurveShape"}).Stage(stage); modified = true }
		if diagram.Seat3DShape == nil { diagram.Seat3DShape = (&Seat3DShape{Name: diagram.Name + "-Seat3DShape"}).Stage(stage); modified = true }
		if diagram.EyeVolume3DShape == nil { diagram.EyeVolume3DShape = (&EyeVolume3DShape{Name: diagram.Name + "-EyeVolume3DShape"}).Stage(stage); modified = true }
		if diagram.SeatAndLegs3DShape == nil { diagram.SeatAndLegs3DShape = (&SeatAndLegs3DShape{Name: diagram.Name + "-SeatAndLegs3DShape"}).Stage(stage); modified = true }
		if diagram.RotatedSeatAndLegs3DShape == nil { diagram.RotatedSeatAndLegs3DShape = (&RotatedSeatAndLegs3DShape{Name: diagram.Name + "-RotatedSeatAndLegs3DShape"}).Stage(stage); modified = true }
		if diagram.TiledFloor3DShape == nil { diagram.TiledFloor3DShape = (&TiledFloor3DShape{Name: diagram.Name + "-TiledFloor3DShape"}).Stage(stage); modified = true }
	}

	for diagram := range *GetGongstructInstancesSetFromPointerType[*Clock3DDiagram](stage) {
		if diagram.Rendered3DShape == nil { diagram.Rendered3DShape = (&Rendered3DShape{Name: diagram.Name + "-Rendered3DShape"}).Stage(stage); modified = true }
		if diagram.ClockTopCurveShape == nil { diagram.ClockTopCurveShape = (&ClockTopCurveShape{Name: diagram.Name + "-ClockTopCurveShape"}).Stage(stage); modified = true }
		if diagram.Torus3DShape == nil { diagram.Torus3DShape = (&Torus3DShape{Name: diagram.Name + "-Torus3DShape"}).Stage(stage); modified = true }
		if diagram.SampledPoints3DShape == nil { diagram.SampledPoints3DShape = (&SampledPoints3DShape{Name: diagram.Name + "-SampledPoints3DShape"}).Stage(stage); modified = true }
		if diagram.TiledFloor3DShape == nil { diagram.TiledFloor3DShape = (&TiledFloor3DShape{Name: diagram.Name + "-TiledFloor3DShape"}).Stage(stage); modified = true }
	}

	for diagram := range *GetGongstructInstancesSetFromPointerType[*Plant3DDiagram](stage) {
		if diagram.Rendered3DShape == nil { diagram.Rendered3DShape = (&Rendered3DShape{Name: diagram.Name + "-Rendered3DShape"}).Stage(stage); modified = true }
		if diagram.StemCylinder3DShape == nil { diagram.StemCylinder3DShape = (&StemCylinder3DShape{Name: diagram.Name + "-StemCylinder3DShape", Transparency: 0.35}).Stage(stage); modified = true }
		if diagram.ParastichyNCurves3DShape == nil { diagram.ParastichyNCurves3DShape = (&ParastichyNCurves3DShape{Name: diagram.Name + "-ParastichyNCurves3DShape"}).Stage(stage); modified = true }
		if diagram.ParastichyMCurves3DShape == nil { diagram.ParastichyMCurves3DShape = (&ParastichyMCurves3DShape{Name: diagram.Name + "-ParastichyMCurves3DShape"}).Stage(stage); modified = true }
		if diagram.CutLine3DShape == nil { diagram.CutLine3DShape = (&CutLine3DShape{Name: diagram.Name + "-CutLine3DShape"}).Stage(stage); modified = true }
		if diagram.Circumference3DShape == nil { diagram.Circumference3DShape = (&Circumference3DShape{Name: diagram.Name + "-Circumference3DShape"}).Stage(stage); modified = true }
		if diagram.TiledFloor3DShape == nil { diagram.TiledFloor3DShape = (&TiledFloor3DShape{Name: diagram.Name + "-TiledFloor3DShape"}).Stage(stage); modified = true }
		if diagram.Leaves3DShape == nil { diagram.Leaves3DShape = (&Leaves3DShape{Name: diagram.Name + "-Leaves3DShape"}).Stage(stage); modified = true }
	}

	return modified
}
