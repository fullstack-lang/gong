package models

import (
	"fmt"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svgGenerateLink(
	startRect *svg.Rect,
	endRect *svg.Rect,
	name string,
	messageName string,
	transitionShape *LinkShape,
	transition *Transition,
	layer *svg.Layer,
	isDashed bool) {

	if startRect == nil || endRect == nil {
		return
	}

	link := new(svg.Link)

	link.Name = startRect.Name + " to " + endRect.Name
	link.Stroke = svg.Black.ToString()
	link.StrokeWidth = 1.5
	link.StrokeOpacity = 1

	link.Type = svg.LINK_TYPE_FLOATING_ORTHOGONAL

	link.Start = startRect
	link.StartOrientation = svg.OrientationType(transitionShape.StartOrientation)
	link.StartRatio = transitionShape.StartRatio

	link.End = endRect
	link.EndOrientation = svg.OrientationType(transitionShape.EndOrientation)
	link.EndRatio = transitionShape.EndRatio
	link.HasEndArrow = true
	link.EndArrowSize = 10

	link.CornerOffsetRatio = transitionShape.CornerOffsetRatio
	link.CornerRadius = 15
	if isDashed {
		link.StrokeDashArray = "5 5"
	}

	link.Impl = NewStateTransition_LinkShapeProxy(
		transitionShape,
		transition,
		stager,
	)

	layer.Links = append(layer.Links, link)

	{
		linkAnchoredText := new(svg.LinkAnchoredText)
		linkAnchoredText.Name = link.Name
		linkAnchoredText.Stroke = svg.Black.ToString()
		linkAnchoredText.StrokeWidth = 1
		linkAnchoredText.StrokeOpacity = 1
		linkAnchoredText.Color = svg.Black.ToString()
		linkAnchoredText.FillOpacity = 1

		linkAnchoredText.Content = name
		linkAnchoredText.AutomaticLayout = true
		linkAnchoredText.LinkAnchorType = svg.LINK_LEFT_OR_TOP
		if messageName != "" {
			linkAnchoredText.Y_Offset = -8
		} else {
			if link.StartOrientation == svg.ORIENTATION_HORIZONTAL {
				linkAnchoredText.Y_Offset = +4
			} else {
				linkAnchoredText.Y_Offset = -4
			}
		}
		link.TextAtArrowStart = append(link.TextAtArrowStart, linkAnchoredText)
	}

	if messageName != "" {
		linkAnchoredText := new(svg.LinkAnchoredText)
		linkAnchoredText.Name = "\n" + messageName
		linkAnchoredText.Stroke = svg.Red.ToString()
		linkAnchoredText.StrokeWidth = 1
		linkAnchoredText.StrokeOpacity = 1
		linkAnchoredText.Color = svg.Red.ToString()
		linkAnchoredText.FillOpacity = 1

		linkAnchoredText.Content = "\n" + messageName
		linkAnchoredText.AutomaticLayout = true
		linkAnchoredText.LinkAnchorType = svg.LINK_LEFT_OR_TOP
		linkAnchoredText.X_Offset = 10
		linkAnchoredText.Y_Offset = 8
		link.TextAtArrowStart = append(link.TextAtArrowStart, linkAnchoredText)
	}
}

func NewStateTransition_LinkShapeProxy(
	stateShapeTransitionInterface LinkShapeInterface,
	transition *Transition,
	stager *Stager,
) (proxy *StateTransition_LinkShapeProxy) {
	proxy = new(StateTransition_LinkShapeProxy)
	proxy.linkShapeInterface = stateShapeTransitionInterface
	proxy.transition = transition
	proxy.stager = stager
	return
}

type StateTransition_LinkShapeProxy struct {
	linkShapeInterface LinkShapeInterface
	transition         *Transition
	stager             *Stager
}

// LinkUpdated implements models.LinkImplInterface.
func (p *StateTransition_LinkShapeProxy) LinkUpdated(updatedLink *svg.Link) {

	diff := p.linkShapeInterface.GetStartRatio() != updatedLink.StartRatio ||
		p.linkShapeInterface.GetEndRatio() != updatedLink.EndRatio ||
		p.linkShapeInterface.GetStartOrientation() != OrientationType(updatedLink.StartOrientation) ||
		p.linkShapeInterface.GetEndOrientation() != OrientationType(updatedLink.EndOrientation) ||
		p.linkShapeInterface.GetCornerOffsetRatio() != updatedLink.CornerOffsetRatio

	p.linkShapeInterface.SetStartRatio(updatedLink.StartRatio)
	p.linkShapeInterface.SetEndRatio(updatedLink.EndRatio)
	p.linkShapeInterface.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
	p.linkShapeInterface.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
	p.linkShapeInterface.SetEndOrientation(OrientationType(updatedLink.EndOrientation))

	if !diff {
		p.stager.stage.CommitWithSuspendedCallbacks()
		p.stager.probeForm.FillUpFormFromGongstruct(p.transition, "Transition")
	} else {
		p.stager.stage.Commit()
	}
}

/**
 * CheckLinkShapeDiff compares a LinkShapeInterface with an updated LinkShape struct.
 * It prints the specific fields that have changed and returns true if any
 * differences are found.
 */
func CheckLinkShapeDiff(iface LinkShapeInterface, updatedLink *svg.Link) bool {
	foundDiff := false

	if oldVal := iface.GetStartRatio(); oldVal != updatedLink.StartRatio {
		fmt.Printf("LinkShape Diff: Field=StartRatio, Old=%v, New=%v\n", oldVal, updatedLink.StartRatio)
		foundDiff = true
	}

	if oldVal := iface.GetEndRatio(); oldVal != updatedLink.EndRatio {
		fmt.Printf("LinkShape Diff: Field=EndRatio, Old=%v, New=%v\n", oldVal, updatedLink.EndRatio)
		foundDiff = true
	}

	if oldVal := iface.GetStartOrientation(); oldVal != OrientationType(updatedLink.StartOrientation) {
		fmt.Printf("LinkShape Diff: Field=StartOrientation, Old=%v, New=%v\n", oldVal, updatedLink.StartOrientation)
		foundDiff = true
	}

	if oldVal := iface.GetEndOrientation(); oldVal != OrientationType(updatedLink.EndOrientation) {
		fmt.Printf("LinkShape Diff: Field=EndOrientation, Old=%v, New=%v\n", oldVal, updatedLink.EndOrientation)
		foundDiff = true
	}

	if oldVal := iface.GetCornerOffsetRatio(); oldVal != updatedLink.CornerOffsetRatio {
		fmt.Printf("LinkShape Diff: Field=CornerOffsetRatio, Old=%v, New=%v\n", oldVal, updatedLink.CornerOffsetRatio)
		foundDiff = true
	}

	return foundDiff
}
