package models

import (
	"fmt"
	"math"
	"slices"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// An Influence is between one Artist/Movement/ArtefactType and another
//
// since gong does not yet support interface, one have to mutliply source/target types
type Influence struct {
	Name string

	SourceMovement     *Movement
	SourceArtefactType *ArtefactType
	SourceArtist       *Artist

	source ArtElement

	TargetMovement     *Movement
	TargetArtefactType *ArtefactType
	TargetArtist       *Artist

	target ArtElement

	// hypothetical, some influences are with ashed lines
	// For instance, Marchine Art to Brancusi (indeed)
	IsHypothtical bool
}

type InfluenceShape struct {
	Name      string
	Influence *Influence

	ControlPointShapes []*ControlPointShape
}

type ControlPointShape struct {
	Name string

	// Control Points are defined relative to their
	// closest rect between the start rect and the end rect
	X_Relative, Y_Relative      float64
	IsStartShapeTheClosestShape bool
}

func (shape *InfluenceShape) GetArtElement() *Influence {
	return shape.Influence
}

type InfluenceShapeProxy struct {
	influenceShape *InfluenceShape
	stager         *Stager
}

func (p *InfluenceShapeProxy) LinkUpdated(updatedLink *svg.Link) {

	point := &svg.Point{
		X: updatedLink.MouseX,
		Y: updatedLink.MouseY,
	}
	controlPoint := PointToControlPoint(point, updatedLink)
	startRect := true
	if controlPoint.ClosestRect == updatedLink.End {
		startRect = false
	}
	controlPoint.Name = "Control Point Shape in " + p.influenceShape.GetName() + " " + fmt.Sprintf("%d", len(p.influenceShape.ControlPointShapes))

	// creates a control point
	controlPointShape := (&ControlPointShape{
		Name:                        controlPoint.Name,
		X_Relative:                  controlPoint.X_Relative,
		Y_Relative:                  controlPoint.Y_Relative,
		IsStartShapeTheClosestShape: startRect,
	}).Stage(p.stager.stage)
	p.influenceShape.ControlPointShapes = append(p.influenceShape.ControlPointShapes, controlPointShape)

	p.stager.stage.Commit()
}

type ControlPointShapeProxy struct {
	stager            *Stager
	influenceShape    *InfluenceShape
	controlPointShape *ControlPointShape
}

func (p *ControlPointShapeProxy) ControlPointUpdated(controlPoint *svg.ControlPoint) {

	if areClose(p.controlPointShape.X_Relative, controlPoint.X_Relative,
		p.controlPointShape.Y_Relative, controlPoint.Y_Relative, threshold) {
		idx := slices.Index(p.influenceShape.ControlPointShapes, p.controlPointShape)
		p.influenceShape.ControlPointShapes = slices.Delete(p.influenceShape.ControlPointShapes, idx, idx+1)
	} else {
		p.controlPointShape.X_Relative = controlPoint.X_Relative
		p.controlPointShape.Y_Relative = controlPoint.Y_Relative

		if controlPoint.ClosestRect.GetName() != p.influenceShape.Influence.source.GetName() {
			p.controlPointShape.IsStartShapeTheClosestShape = false
		} else {
			p.controlPointShape.IsStartShapeTheClosestShape = true
		}
	}

	p.stager.stage.Commit()
}

/**
 * areClose checks if the absolute distances between (x1, x2) and (y1, y2)
 * are both below a given threshold.
 */
func areClose(x1, x2, y1, y2, threshold float64) bool {
	// Calculate the absolute distance for x and y pairs
	xDistance := math.Abs(x1 - x2)
	yDistance := math.Abs(y1 - y2)

	// Return true only if BOTH distances are less than the threshold
	return xDistance < threshold && yDistance < threshold
}

const threshold = 0.05

type InfluenceNodeProxy struct {
	stager    *Stager
	diagram   *Diagram
	influence *Influence
	node      *tree.Node
}

// OnAfterUpdate implements models.NodeImplInterface.
func (d *InfluenceNodeProxy) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {

	if frontNode.IsChecked && !stagedNode.IsChecked {
		influenceShape := &InfluenceShape{
			Influence: d.influence,
		}
		influenceShape.Stage(d.stager.stage)
		d.diagram.InfluenceShapes = append(d.diagram.InfluenceShapes, influenceShape)
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {

		for idx, shape := range d.diagram.InfluenceShapes {
			if shape.Influence == d.influence {
				shape.Unstage(d.stager.stage)
				d.diagram.InfluenceShapes = slices.Delete(d.diagram.InfluenceShapes, idx, idx+1)
				continue
			}
		}
	}

	d.stager.stage.Commit()
}

// PointToControlPoint converts a Point (absolute coords) to a ControlPoint (relative coords)
// based on the provided Link.
// This function assumes link.Start and link.End are not nil.
func PointToControlPoint(point *svg.Point, link *svg.Link) svg.ControlPoint {
	// This will panic if link.Start or link.End is nil.
	distToStart := calculateDistance(point.X, point.Y, link.Start.X, link.Start.Y)
	distToEnd := calculateDistance(point.X, point.Y, link.End.X, link.End.Y)

	var closestRect *svg.Rect
	if distToStart > distToEnd {
		closestRect = link.End
	} else {
		closestRect = link.Start
	}

	// Add a small epsilon to prevent division by zero, just like in the TS code.
	xRel := (point.X - closestRect.X) / (closestRect.Width + 0.0001)
	yRel := (point.Y - closestRect.Y) / (closestRect.Height + 0.0001)

	return svg.ControlPoint{
		X_Relative:  xRel,
		Y_Relative:  yRel,
		ClosestRect: closestRect,
	}
}

func ControlPointToPoint(controlPoint *svg.ControlPoint) *svg.Point {
	// Go automatically dereferences the pointer (controlPoint.ClosestRect.X)
	// This will panic if ClosestRect is nil, mirroring the TS '!' assertion.
	x := controlPoint.ClosestRect.X + controlPoint.X_Relative*controlPoint.ClosestRect.Width
	y := controlPoint.ClosestRect.Y + controlPoint.Y_Relative*controlPoint.ClosestRect.Height

	return &svg.Point{X: x, Y: y}
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
