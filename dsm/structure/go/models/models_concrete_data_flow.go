package models

import (
	"fmt"
	"math"
	"slices"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

const controlPointDeletePixelThreshold = 5.0

type ControlPointShape struct {
	Name string

	X_Relative                  float64
	Y_Relative                  float64
	IsStartShapeTheClosestShape bool
}

type ControlPointShapeProxy struct {
	stager            *Stager
	dataFlowShape     *DataFlowShape
	controlPointShape *ControlPointShape
	startRect         *svg.Rect
	endRect           *svg.Rect
}

func (p *ControlPointShapeProxy) ControlPointUpdated(controlPoint *svg.ControlPoint) {
	if p.dataFlowShape == nil || p.controlPointShape == nil {
		return
	}

	oldPt := ControlPointShapeToPoint(p.controlPointShape, p.startRect, p.endRect)
	fallbackRect := p.startRect
	if !p.controlPointShape.IsStartShapeTheClosestShape {
		fallbackRect = p.endRect
	}
	newPt := ControlPointToPoint(controlPoint, fallbackRect)

	if math.Hypot(newPt.X-oldPt.X, newPt.Y-oldPt.Y) < controlPointDeletePixelThreshold {
		idx := slices.Index(p.dataFlowShape.ControlPointShapes, p.controlPointShape)
		if idx != -1 {
			p.dataFlowShape.ControlPointShapes = slices.Delete(p.dataFlowShape.ControlPointShapes, idx, idx+1)
		}
		p.controlPointShape.Unstage(p.stager.stage)
	} else {
		p.controlPointShape.X_Relative = controlPoint.X_Relative
		p.controlPointShape.Y_Relative = controlPoint.Y_Relative
		if controlPoint.ClosestRect != nil {
			p.controlPointShape.IsStartShapeTheClosestShape = controlPoint.ClosestRect == p.startRect
		}
	}

	p.stager.stage.Commit()
}

type LinkShapeProxy struct {
	stager        *Stager
	dataFlowShape *DataFlowShape
	startRect     *svg.Rect
	endRect       *svg.Rect
}

func (p *LinkShapeProxy) LinkUpdated(updatedLink *svg.Link) {
	if p.dataFlowShape == nil || p.startRect == nil || p.endRect == nil {
		return
	}

	point := &svg.Point{
		X: updatedLink.MouseX,
		Y: updatedLink.MouseY,
	}
	cp := PointToControlPoint(point, p.startRect, p.endRect)
	isStartClosest := cp.ClosestRect == p.startRect

	controlPointShape := (&ControlPointShape{
		Name: fmt.Sprintf("Control Point Shape in %s %d",
			p.dataFlowShape.Name, len(p.dataFlowShape.ControlPointShapes)),
		X_Relative:                  cp.X_Relative,
		Y_Relative:                  cp.Y_Relative,
		IsStartShapeTheClosestShape: isStartClosest,
	}).Stage(p.stager.stage)

	pathPoints := make([]svg.Point, 0, len(p.dataFlowShape.ControlPointShapes)+2)
	pathPoints = append(pathPoints, svg.Point{
		X: p.startRect.X + p.startRect.Width/2.0,
		Y: p.startRect.Y + p.startRect.Height/2.0,
	})
	for _, existing := range p.dataFlowShape.ControlPointShapes {
		pathPoints = append(pathPoints, ControlPointShapeToPoint(existing, p.startRect, p.endRect))
	}
	pathPoints = append(pathPoints, svg.Point{
		X: p.endRect.X + p.endRect.Width/2.0,
		Y: p.endRect.Y + p.endRect.Height/2.0,
	})

	bestIndex := 0
	minDistance := math.MaxFloat64
	for i := 0; i < len(pathPoints)-1; i++ {
		dist := distanceToSegment(point.X, point.Y, pathPoints[i].X, pathPoints[i].Y, pathPoints[i+1].X, pathPoints[i+1].Y)
		if dist < minDistance {
			minDistance = dist
			bestIndex = i
		}
	}

	p.dataFlowShape.ControlPointShapes = slices.Insert(p.dataFlowShape.ControlPointShapes, bestIndex, controlPointShape)

	p.stager.stage.Commit()

	if p.stager.probeForm != nil && p.dataFlowShape.DataFlow != nil {
		p.stager.probeForm.FillUpFormFromGongstruct(p.dataFlowShape.DataFlow, "DataFlow")
	}
	fillUpShapeAppearanceForm(p.stager, p.dataFlowShape)
}

func PointToControlPoint(point *svg.Point, startRect, endRect *svg.Rect) svg.ControlPoint {
	startCenterX := startRect.X + startRect.Width/2.0
	startCenterY := startRect.Y + startRect.Height/2.0
	endCenterX := endRect.X + endRect.Width/2.0
	endCenterY := endRect.Y + endRect.Height/2.0

	distToStart := math.Hypot(point.X-startCenterX, point.Y-startCenterY)
	distToEnd := math.Hypot(point.X-endCenterX, point.Y-endCenterY)

	closestRect := startRect
	if distToStart > distToEnd {
		closestRect = endRect
	}

	xRel := (point.X - closestRect.X) / (closestRect.Width + 0.0001)
	yRel := (point.Y - closestRect.Y) / (closestRect.Height + 0.0001)

	return svg.ControlPoint{
		X_Relative:  xRel,
		Y_Relative:  yRel,
		ClosestRect: closestRect,
	}
}

func ControlPointShapeToPoint(cp *ControlPointShape, startRect, endRect *svg.Rect) svg.Point {
	if cp == nil {
		return svg.Point{}
	}
	closestRect := startRect
	if !cp.IsStartShapeTheClosestShape {
		closestRect = endRect
	}
	if closestRect == nil {
		return svg.Point{}
	}
	return svg.Point{
		X: closestRect.X + cp.X_Relative*closestRect.Width,
		Y: closestRect.Y + cp.Y_Relative*closestRect.Height,
	}
}

func ControlPointToPoint(cp *svg.ControlPoint, fallbackRect *svg.Rect) svg.Point {
	if cp == nil {
		return svg.Point{}
	}
	rect := cp.ClosestRect
	if rect == nil {
		rect = fallbackRect
	}
	if rect == nil {
		return svg.Point{X: cp.X_Relative, Y: cp.Y_Relative}
	}
	return svg.Point{
		X: rect.X + cp.X_Relative*rect.Width,
		Y: rect.Y + cp.Y_Relative*rect.Height,
	}
}

func distanceToSegment(px, py, ax, ay, bx, by float64) float64 {
	l2 := (ax-bx)*(ax-bx) + (ay-by)*(ay-by)
	if l2 == 0 {
		return math.Hypot(px-ax, py-ay)
	}
	t := ((px-ax)*(bx-ax) + (py-ay)*(by-ay)) / l2
	t = math.Max(0, math.Min(1, t))
	projX := ax + t*(bx-ax)
	projY := ay + t*(by-ay)
	return math.Hypot(px-projX, py-projY)
}

type DataFlowShape struct {
	Name string

	DataFlow *DataFlow

	// ParentPartShape is non-nil when this DataFlowShape represents an internal DataFlow
	// of ParentPartShape.Part.Type.DataFlows.
	ParentPartShape *PartShape
	StartPortShape  *PortShape
	EndPortShape    *PortShape

	IsProtocolNodeExpanded bool

	LinkShape
	ControlPointShapes []*ControlPointShape

	dataShapes          []*DataShape
	protocolShape       *ProtocolShape
	protocolLayerShapes []*ProtocolLayerShape
}

func (s *DataFlowShape) GetAbstractElement() AbstractType {
	// very important to return nil, otherwisse, it will be a typed nil and the check of nil will not work
	if s.DataFlow == nil {
		return nil
	}
	return s.DataFlow
}

func (s *DataFlowShape) SetAbstractElement(a AbstractType) {
	s.DataFlow = a.(*DataFlow)
}

func (s *DataFlowShape) GetAbstractEndElement() AbstractType {
	if s.DataFlow == nil || (s.DataFlow.EndPort == nil && s.DataFlow.EndExternalPart == nil) {
		return nil
	}
	if s.DataFlow.EndPort != nil {
		return s.DataFlow.EndPort
	}
	if s.DataFlow.EndExternalPart != nil {
		return s.DataFlow.EndExternalPart
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractEndElement(abstractElement AbstractType) {
	if endPort, ok := abstractElement.(*Port); ok {
		s.DataFlow.EndPort = endPort
		return
	}
	if endExternalPart, ok := abstractElement.(*Part); ok {
		s.DataFlow.EndExternalPart = endExternalPart
		return
	}
}

func (s *DataFlowShape) GetAbstractStartElement() AbstractType {
	if s.DataFlow == nil || (s.DataFlow.StartPort == nil && s.DataFlow.StartExternalPart == nil) {
		return nil
	}
	if s.DataFlow.StartPort != nil {
		return s.DataFlow.StartPort
	}
	if s.DataFlow.StartExternalPart != nil {
		return s.DataFlow.StartExternalPart
	}
	return nil // can not happen
}

func (s *DataFlowShape) SetAbstractStartElement(abstractElement AbstractType) {
	if startPort, ok := abstractElement.(*Port); ok {
		s.DataFlow.StartPort = startPort
		return
	}
	if startExternalPart, ok := abstractElement.(*Part); ok {
		s.DataFlow.StartExternalPart = startExternalPart
		return
	}
}

var _ AssociationConcreteType = (*DataFlowShape)(nil)

var _ ConcreteType = (*DataFlowShape)(nil)
