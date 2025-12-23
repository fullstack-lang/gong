package models

import (
	"fmt"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svgGenerateLink(
	startRect *svg.Rect,
	endRect *svg.Rect,
	compositionShape *LinkShape,
	subProduct *Product,
	layer *svg.Layer,
	isDashed bool,
) {
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
	link.StartOrientation = svg.OrientationType(compositionShape.StartOrientation)
	link.StartRatio = compositionShape.StartRatio

	link.End = endRect
	link.EndOrientation = svg.OrientationType(compositionShape.EndOrientation)
	link.EndRatio = compositionShape.EndRatio
	link.HasEndArrow = true
	link.EndArrowSize = 10

	link.CornerOffsetRatio = compositionShape.CornerOffsetRatio
	link.CornerRadius = 15
	if isDashed {
		link.StrokeDashArray = "5 5"
	}

	link.Impl = NewStateComposition_LinkShapeProxy(
		compositionShape,
		subProduct,
		stager,
	)

	layer.Links = append(layer.Links, link)
}

func NewStateComposition_LinkShapeProxy(
	stateShapeCompositionInterface LinkShapeInterface,
	subProduct *Product,
	stager *Stager,
) (proxy *StateComposition_LinkShapeProxy) {
	proxy = new(StateComposition_LinkShapeProxy)
	proxy.linkShapeInterface = stateShapeCompositionInterface
	proxy.subProduct = subProduct
	proxy.stager = stager
	return
}

type StateComposition_LinkShapeProxy struct {
	linkShapeInterface LinkShapeInterface
	subProduct         *Product
	stager             *Stager
}

// LinkUpdated implements models.LinkImplInterface.
func (p *StateComposition_LinkShapeProxy) LinkUpdated(updatedLink *svg.Link) {
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
		p.stager.probeForm.FillUpFormFromGongstruct(p.subProduct, GetPointerToGongstructName[*Product]())
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
