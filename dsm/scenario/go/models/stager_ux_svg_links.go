package models

import (
	"fmt"
	"math"
	"slices"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// LinkImplActorStateShapeTransition satisfies the LinkImplInterface
type LinkImplActorStateShapeTransition struct {
	actorStateTransitionShape *ActorStateTransitionShape
	stage                     *Stage
}

func (stager *Stager) NewLinkImplActorStateShapeTransition(
	actorStateTransitionShape *ActorStateTransitionShape,
) *LinkImplActorStateShapeTransition {
	return &LinkImplActorStateShapeTransition{
		actorStateTransitionShape: actorStateTransitionShape,
		stage:                     stager.stage,
	}
}

func (linkImpl *LinkImplActorStateShapeTransition) LinkUpdated(updatedLink *gongsvg_models.Link) {
	point := &gongsvg_models.Point{
		X: updatedLink.MouseX,
		Y: updatedLink.MouseY,
	}
	controlPoint := PointToControlPoint(point, updatedLink)
	startRect := true
	if controlPoint.ClosestRect == updatedLink.End {
		startRect = false
	}
	controlPoint.Name = "Control Point Shape in " + linkImpl.actorStateTransitionShape.Name + " " + fmt.Sprintf("%d", len(linkImpl.actorStateTransitionShape.ControlPointShapes))

	// creates a control point
	controlPointShape := (&ControlPointShape{
		Name:                        controlPoint.Name,
		X_Relative:                  controlPoint.X_Relative,
		Y_Relative:                  controlPoint.Y_Relative,
		IsStartShapeTheClosestShape: startRect,
	}).Stage(linkImpl.stage)

	// heuristic to find the best index to insert the new control point
	var pathPoints []gongsvg_models.Point
	pathPoints = append(pathPoints, gongsvg_models.Point{
		X: updatedLink.Start.X + updatedLink.Start.Width/2.0,
		Y: updatedLink.Start.Y + updatedLink.Start.Height/2.0,
	})

	for _, cpShape := range linkImpl.actorStateTransitionShape.ControlPointShapes {
		closestRect := updatedLink.Start
		if !cpShape.IsStartShapeTheClosestShape {
			closestRect = updatedLink.End
		}
		x := closestRect.X + cpShape.X_Relative*closestRect.Width
		y := closestRect.Y + cpShape.Y_Relative*closestRect.Height
		pathPoints = append(pathPoints, gongsvg_models.Point{X: x, Y: y})
	}

	pathPoints = append(pathPoints, gongsvg_models.Point{
		X: updatedLink.End.X + updatedLink.End.Width/2.0,
		Y: updatedLink.End.Y + updatedLink.End.Height/2.0,
	})

	bestIndex := 0
	minDistance := math.MaxFloat64

	for i := 0; i < len(pathPoints)-1; i++ {
		p1 := pathPoints[i]
		p2 := pathPoints[i+1]
		dist := distanceToSegment(point.X, point.Y, p1.X, p1.Y, p2.X, p2.Y)
		if dist < minDistance {
			minDistance = dist
			bestIndex = i
		}
	}

	linkImpl.actorStateTransitionShape.ControlPointShapes = slices.Insert(linkImpl.actorStateTransitionShape.ControlPointShapes, bestIndex, controlPointShape)

	linkImpl.stage.Commit()
}

type ControlPointShapeProxy struct {
	stager                    *Stager
	actorStateTransitionShape *ActorStateTransitionShape
	controlPointShape         *ControlPointShape
}

func (p *ControlPointShapeProxy) ControlPointUpdated(controlPoint *gongsvg_models.ControlPoint) {
	if areClose(p.controlPointShape.X_Relative, controlPoint.X_Relative,
		p.controlPointShape.Y_Relative, controlPoint.Y_Relative, threshold) {
		idx := slices.Index(p.actorStateTransitionShape.ControlPointShapes, p.controlPointShape)
		p.actorStateTransitionShape.ControlPointShapes = slices.Delete(p.actorStateTransitionShape.ControlPointShapes, idx, idx+1)
	} else {
		p.controlPointShape.X_Relative = controlPoint.X_Relative
		p.controlPointShape.Y_Relative = controlPoint.Y_Relative

		if controlPoint.ClosestRect.GetName() != p.actorStateTransitionShape.Start.ActorState.Name {
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

// PointToControlPoint converts a Point (absolute coords) to a ControlPoint (relative coords)
// based on the provided Link.
// This function assumes link.Start and link.End are not nil.
func PointToControlPoint(point *gongsvg_models.Point, link *gongsvg_models.Link) gongsvg_models.ControlPoint {
	// This will panic if link.Start or link.End is nil.
	distToStart := calculateDistance(point.X, point.Y, link.Start.X, link.Start.Y)
	distToEnd := calculateDistance(point.X, point.Y, link.End.X, link.End.Y)

	var closestRect *gongsvg_models.Rect
	if distToStart > distToEnd {
		closestRect = link.End
	} else {
		closestRect = link.Start
	}

	// Add a small epsilon to prevent division by zero, just like in the TS code.
	xRel := (point.X - closestRect.X) / (closestRect.Width + 0.0001)
	yRel := (point.Y - closestRect.Y) / (closestRect.Height + 0.0001)

	return gongsvg_models.ControlPoint{
		X_Relative:  xRel,
		Y_Relative:  yRel,
		ClosestRect: closestRect,
	}
}

func ControlPointToPoint(controlPoint *gongsvg_models.ControlPoint) *gongsvg_models.Point {
	// Go automatically dereferences the pointer (controlPoint.ClosestRect.X)
	// This will panic if ClosestRect is nil, mirroring the TS '!' assertion.
	x := controlPoint.ClosestRect.X + controlPoint.X_Relative*controlPoint.ClosestRect.Width
	y := controlPoint.ClosestRect.Y + controlPoint.Y_Relative*controlPoint.ClosestRect.Height

	return &gongsvg_models.Point{X: x, Y: y}
}

func calculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func distanceToSegment(px, py, ax, ay, bx, by float64) float64 {
	l2 := (ax-bx)*(ax-bx) + (ay-by)*(ay-by)
	if l2 == 0 {
		return calculateDistance(px, py, ax, ay)
	}
	t := ((px-ax)*(bx-ax) + (py-ay)*(by-ay)) / l2
	t = math.Max(0, math.Min(1, t))
	projX := ax + t*(bx-ax)
	projY := ay + t*(by-ay)
	return calculateDistance(px, py, projX, projY)
}

func (stager *Stager) drawActorStateTransitionShapes(
	diagram *Diagram,
	actorStateLayer *gongsvg_models.Layer,
	map_ActorStateShape_Rect map[*ActorStateShape]*gongsvg_models.Rect,
) {
	var newActorStateTransitionShapes []*ActorStateTransitionShape
	for _, actorStateTransitionShape := range diagram.ActorStateTransitionShapes {

		// shape self destruct if not consistent
		if actorStateTransitionShape.ActorStateTransition == nil ||
			actorStateTransitionShape.Start == nil ||
			actorStateTransitionShape.End == nil ||
			actorStateTransitionShape.Start.ActorState == nil ||
			actorStateTransitionShape.End.ActorState == nil {
			actorStateTransitionShape.Unstage(stager.stage)
			stager.stage.Commit()
		} else {
			newActorStateTransitionShapes = append(newActorStateTransitionShapes, actorStateTransitionShape)
		}
	}
	diagram.ActorStateTransitionShapes = newActorStateTransitionShapes

	for _, actorStateTransitionShape := range diagram.ActorStateTransitionShapes {

		if actorStateTransitionShape.GetIsHidden() {
			continue
		}
		if actorStateTransitionShape.Start.GetIsHidden() || actorStateTransitionShape.End.GetIsHidden() {
			continue
		}

		transitionLink := new(gongsvg_models.Link)
		transitionLink.Name = actorStateTransitionShape.Name

		transitionLink.Start = map_ActorStateShape_Rect[actorStateTransitionShape.Start]
		transitionLink.End = map_ActorStateShape_Rect[actorStateTransitionShape.End]

		transitionLink.HasEndArrow = true
		transitionLink.EndArrowSize = 10
		transitionLink.Type = gongsvg_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS

		transitionLink.Stroke = "#7E57C2" // Soft purple
		transitionLink.StrokeWidth = 2.5
		transitionLink.StrokeOpacity = 1

		transitionLink.StartAnchorType = gongsvg_models.ANCHOR_CENTER
		transitionLink.EndAnchorType = gongsvg_models.ANCHOR_CENTER

		transitionLink.StartArrowOffset = 15 // Offset from the start shape
		transitionLink.EndArrowOffset = 15   // Offset from the end shape

		transitionLink.CornerRadius = 15 // Make the corners smooth
		transitionLink.IsBezierCurve = true

		for _, controlPointShape := range actorStateTransitionShape.ControlPointShapes {
			controlPoint := new(gongsvg_models.ControlPoint)
			controlPoint.Name = controlPointShape.Name
			controlPoint.X_Relative = controlPointShape.X_Relative
			controlPoint.Y_Relative = controlPointShape.Y_Relative
			if controlPointShape.IsStartShapeTheClosestShape {
				controlPoint.ClosestRect = transitionLink.Start
			} else {
				controlPoint.ClosestRect = transitionLink.End
			}

			if diagram.IsInDrawMode {
				controlPoint.Impl = &ControlPointShapeProxy{
					stager:                    stager,
					actorStateTransitionShape: actorStateTransitionShape,
					controlPointShape:         controlPointShape,
				}
			}

			transitionLink.ControlPoints = append(transitionLink.ControlPoints, controlPoint)
		}

		if diagram.IsInDrawMode {
			transitionLink.Impl = stager.NewLinkImplActorStateShapeTransition(actorStateTransitionShape)
		}

		actorStateLayer.Links = append(actorStateLayer.Links, transitionLink)
	}
}

func (stager *Stager) drawScenarioParameterToParameterLinks(
	diagram *Diagram,
	scenarioParameterLayer *gongsvg_models.Layer,
	map_ScenarioParameterShape_Rect map[*ParametersAggregateShape]*gongsvg_models.Rect,
	map_ParameterShape_Rect map[*ParameterShape]*gongsvg_models.Rect,
) {
	for _, scenarioParameterShape := range diagram.ScenarioParameterShapes {

		if scenarioParameterShape.ScenarioParameter == nil {
			scenarioParameterShape.Unstage(stager.stage)
			for idx, _scenarioParameterShape := range diagram.ScenarioParameterShapes {
				if scenarioParameterShape == _scenarioParameterShape {
					diagram.ScenarioParameterShapes = slices.Delete(diagram.ScenarioParameterShapes, idx, idx+1)
					break
				}
			}
			stager.stage.Commit()
			continue
		}

		if scenarioParameterShape.GetIsHidden() {
			continue
		}

		for _, parameter := range scenarioParameterShape.ScenarioParameter.Parameters {

			// find the relevant paramter shape
			var parameterShape *ParameterShape
			for _, _paparameterShape := range diagram.ParameterShapes {
				if _paparameterShape.Parameter == parameter {
					parameterShape = _paparameterShape
				}
			}
			if parameterShape == nil || parameterShape.GetIsHidden() {
				continue
			}

			dependencyLink := new(gongsvg_models.Link)
			dependencyLink.Name = scenarioParameterShape.Name

			dependencyLink.Start = map_ScenarioParameterShape_Rect[scenarioParameterShape]
			dependencyLink.End = map_ParameterShape_Rect[parameterShape]

			dependencyLink.HasEndArrow = true
			dependencyLink.EndArrowSize = 10
			dependencyLink.Type = gongsvg_models.LINK_TYPE_LINE_WITH_CONTROL_POINTS

			dependencyLink.Stroke = "#B0BEC5" // Soft blue-grey
			dependencyLink.StrokeDashArray = "5 3"
			dependencyLink.StrokeWidth = 1.5
			dependencyLink.StrokeOpacity = 1

			dependencyLink.StartAnchorType = gongsvg_models.ANCHOR_CENTER
			dependencyLink.EndAnchorType = gongsvg_models.ANCHOR_CENTER

			dependencyLink.StartArrowOffset = 15 // Offset from the start shape
			dependencyLink.EndArrowOffset = 15   // Offset from the end shape

			dependencyLink.CornerRadius = 20 // like barrgraph

			scenarioParameterLayer.Links = append(scenarioParameterLayer.Links, dependencyLink)
		}
	}
}
