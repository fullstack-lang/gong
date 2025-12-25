package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/pkg/strutils"
)

func svgRect[CT interface {
	*CT_
	RectShapeInterface
	ConcreteType
}, CT_ Gongstruct](
	stager *Stager,
	diagram *Diagram, // might be ni
	shape CT,
	layer *svg.Layer,
) *svg.Rect {
	root := stager.root

	abstractElement := shape.GetAbstractElement()

	rect := new(svg.Rect)
	rect.Name = abstractElement.GetName()
	rect.Stroke = svg.Black.ToString()
	rect.StrokeWidth = 2
	rect.StrokeOpacity = 1
	rect.X = shape.GetX()
	rect.Y = shape.GetY()
	rect.Width = shape.GetWidth()
	rect.Height = shape.GetHeight()
	rect.RX = 15

	// rect is editable if diagram is not null

	if diagram.IsEditable() {
		rect.CanHaveBottomHandle = true
		rect.CanHaveLeftHandle = true
		rect.CanHaveRightHandle = true
		rect.CanHaveTopHandle = true

		rect.CanMoveHorizontaly = true
		rect.CanMoveVerticaly = true

		rect.Impl = &svg.FunctionalSvgRectProxy{
			OnUpdated: OnUpdateRectElement(stager, abstractElement, shape),
		}
		// for allowing later Stage() on the rect shape
		shape.SetReceiver(shape)
	}

	title := new(svg.RectAnchoredText)
	{
		title.Name = abstractElement.GetName()

		content := shape.GetAbstractElement().GetName()
		if diagram.ShowPrefix {
			content = abstractElement.GetComputedPrefix() + " " + content
		}

		if rect.Width > 0 {
			content = strutils.WrapString(content, int(rect.Width/root.NbPixPerCharacter))
		}
		title.Content = content
		title.Stroke = svg.Black.ToString()
		title.StrokeWidth = 1
		title.StrokeOpacity = 1
		title.Color = svg.Black.ToString()
		title.FillOpacity = 1

		title.FontSize = "16px"
		title.X_Offset = 0
		title.Y_Offset = 30
		title.RectAnchorType = svg.RECT_TOP
		title.TextAnchorType = svg.TEXT_ANCHOR_CENTER

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
	}

	layer.Rects = append(layer.Rects, rect)
	return rect
}

func OnUpdateRectElement[CT interface {
	*CT_
	RectShapeInterface
	ConcreteType
}, CT_ Gongstruct](
	stager *Stager,
	abstractElement AbstractType,
	rectShape CT,
) func(frontRect *svg.Rect) {
	return func(updatedRect *svg.Rect) {
		diffSize := rectShape.GetWidth() != updatedRect.Width ||
			rectShape.GetHeight() != updatedRect.Height

		diffPosition := rectShape.GetX() != updatedRect.X ||
			rectShape.GetY() != updatedRect.Y

		rectShape.SetX(updatedRect.X)
		rectShape.SetY(updatedRect.Y)
		rectShape.SetWidth(updatedRect.Width)
		rectShape.SetHeight(updatedRect.Height)

		if !diffSize && !diffPosition {
			stager.stage.CommitWithSuspendedCallbacks()
			stager.probeForm.FillUpFormFromGongstruct(abstractElement, GetPointerToGongstructName[AbstractType]())
		}

		if diffPosition {
			// Issue #7, this will allow multiple rect to be moved together
			stager.stage.CommitWithSuspendedCallbacks()
		}

		if diffSize {
			stager.stage.Commit()
		}
	}
}

func svgAssociationLink[AT AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](
	stager *Stager,
	startRect *svg.Rect,
	endRect *svg.Rect,
	shape ACT,
	productOfInterest AT,
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
	link.StartOrientation = svg.OrientationType(shape.GetStartOrientation())
	link.StartRatio = shape.GetStartRatio()

	link.End = endRect
	link.EndOrientation = svg.OrientationType(shape.GetEndOrientation())
	link.EndRatio = shape.GetEndRatio()
	link.HasEndArrow = true
	link.EndArrowSize = 10

	link.CornerOffsetRatio = shape.GetCornerOffsetRatio()
	link.CornerRadius = 15
	if isDashed {
		link.StrokeDashArray = "5 5"
	}

	link.Impl = &svg.FunctionalSvgLinkProxy{
		OnUpdated: func(updatedLink *svg.Link) {
			diff := shape.GetStartRatio() != updatedLink.StartRatio ||
				shape.GetEndRatio() != updatedLink.EndRatio ||
				shape.GetStartOrientation() != OrientationType(updatedLink.StartOrientation) ||
				shape.GetEndOrientation() != OrientationType(updatedLink.EndOrientation) ||
				shape.GetCornerOffsetRatio() != updatedLink.CornerOffsetRatio

			shape.SetStartRatio(updatedLink.StartRatio)
			shape.SetEndRatio(updatedLink.EndRatio)
			shape.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
			shape.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
			shape.SetEndOrientation(OrientationType(updatedLink.EndOrientation))

			if !diff {
				stager.stage.CommitWithSuspendedCallbacks()
				stager.probeForm.FillUpFormFromGongstruct(productOfInterest, GetPointerToGongstructName[AT]())
			} else {
				stager.stage.Commit()
			}
		},
	}

	layer.Links = append(layer.Links, link)
}
