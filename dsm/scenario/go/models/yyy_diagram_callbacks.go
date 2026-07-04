// generated code (do not edit)
package models

import (
	"log"
	"math/rand/v2"
	"slices"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func addAssociationShapeToDiagram[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType // the association concrete type shape
	},
	ACT_ Gongstruct](stager *Stager, start ATstart, end ATend, shapes *[]ACT) {
	stage := stager.stage
	_ = stage

	compositionShape := newConcreteAssociation(start, end, shapes)
	compositionShape.StageVoid(stager.stage)
}

func onAddAssociationShape[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct,
](
	stager *Stager, start ATstart, end ATend, shapes *[]ACT,
) func() {
	return func() {
		addAssociationShapeToDiagram(stager, start, end, shapes)
		stager.stage.Commit()
	}
}

func onRemoveAssociationShape[
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](stager *Stager, compositionShape ACT, shapes *[]ACT) func() {
	return func() {
		compositionShape.UnstageVoid(stager.stage)
		idx := slices.Index(*shapes, compositionShape)
		*shapes = slices.Delete(*shapes, idx, idx+1)
		stager.stage.Commit()
	}
}

func setCallbacksElementInDiagram[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
	CT interface {
		*CT_
		RectShapeInterface
		ConcreteType
	},
	CT_ Gongstruct,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](
	stager *Stager,
	node *tree.Node,
	diagram DiagramIF,
	element AT,
	parentElement AT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap map[AT]CT,
	compositionShape ACT,
	compositionShapes *[]ACT,
) {
	node.OnIsCheckedChanged = func(isChecked bool) {
		shape := shapesMap[element]
		if isChecked {
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(element, diagram, shapes, stager, node.ClientOnY)
			shapesMap[element] = shape

			if parentElement != nil {
				if parentShape, ok := shapesMap[parentElement]; ok {
					if compositionShapes != nil {
						addAssociationShapeToDiagram(stager, parentElement, element, compositionShapes)

						var siblingsCount int
						for _, compShape := range *compositionShapes {
							if compShape.GetAbstractStartElement() == parentElement {
								siblingsCount++
							}
						}

						var isHorizontal bool
						if pShape, ok := any(parentShape).(LayoutConcreteType); ok {
							if pShape.GetConcreteLayoutDirection() == Horizontal {
								isHorizontal = true
							}
						}

						if isHorizontal {
							shape.SetX(parentShape.GetX() + parentShape.GetWidth()/2.0 + 50.0)
							shape.SetY(parentShape.GetY() + parentShape.GetHeight() + 50.0 + float64(siblingsCount-1)*parentShape.GetHeight()*1.2)

							if len(*compositionShapes) > 0 {
								newCompositionShape := (*compositionShapes)[len(*compositionShapes)-1]
								newCompositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
								newCompositionShape.SetEndOrientation(ORIENTATION_HORIZONTAL)
								newCompositionShape.SetCornerOffsetRatio(1.5)
							}
						} else {
							var isGrandParentHorizontal bool
							for _, compShape := range *compositionShapes {
								if compShape.GetAbstractEndElement() == any(parentElement).(AbstractType) {
									grandParentElement := compShape.GetAbstractStartElement()
									if grandParentElement != nil {
										if grandParentShape, ok := shapesMap[grandParentElement.(AT)]; ok {
											if gpShape, ok := any(grandParentShape).(LayoutConcreteType); ok {
												if gpShape.GetConcreteLayoutDirection() == Horizontal {
													isGrandParentHorizontal = true
												}
											}
										}
									}
								}
							}

							if isGrandParentHorizontal {
								shape.SetX(parentShape.GetX() + parentShape.GetWidth()/2.0 + 50.0 + float64(siblingsCount-1)*parentShape.GetWidth()*1.2)
							} else {
								shape.SetX(parentShape.GetX() + float64(siblingsCount-1)*parentShape.GetWidth()*1.2)
							}
							shape.SetY(parentShape.GetY() + parentShape.GetHeight()*2.0)

							if len(*compositionShapes) > 0 {
								newCompositionShape := (*compositionShapes)[len(*compositionShapes)-1]
								newCompositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
								newCompositionShape.SetEndOrientation(ORIENTATION_VERTICAL)
								ratio := (shape.GetY() - parentShape.GetY()) / parentShape.GetHeight()
								newCompositionShape.SetCornerOffsetRatio((ratio-1.0)/2.0 + 1.0)
							}
						}
					}
				}
			}
			stager.stage.Commit()
		} else {
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stager.stage)
			idx := slices.Index(*shapes, shape)
			*shapes = slices.Delete(*shapes, idx, idx+1)
			stager.stage.Commit()
		}
	}

	node.OnIsExpandedChange = func(isExpanded bool) {
		if isExpanded {
			if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
				*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
			}
		} else {
			if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
				*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
			}
		}
		stager.stage.CommitWithSuspendedCallbacks()
	}

	node.OnNameChange = func(newName string) {
		element.SetName(newName)
		element.SetIsInRenameMode(false)
		stager.stage.Commit()
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
	}

	node.OnIsSecondCheckboxCheckedChanged = func(isChecked bool) {
		if isChecked {
			compositionShape := newConcreteAssociation(parentElement, element, compositionShapes)
			if parentShape, ok := shapesMap[parentElement]; ok && shapesMap[element] != nil {
				shape := shapesMap[element]
				ratio := (shape.GetY() - parentShape.GetY()) / parentShape.GetHeight()
				compositionShape.SetCornerOffsetRatio((ratio-1.0)/2.0 + 1.0)
			}
			compositionShape.StageVoid(stager.stage)
			stager.stage.Commit()
		} else {
			compositionShape.UnstageVoid(stager.stage)
			idx := slices.Index(*compositionShapes, compositionShape)
			*compositionShapes = slices.Delete(*compositionShapes, idx, idx+1)
			stager.stage.Commit()
		}
	}

	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.CommitWithSuspendedCallbacks()
	}
}

// setCallbacksElementInDiagramWithoutLink handles the update event for an element in the diagram without a link
func setCallbacksElementInDiagramWithoutLink[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
	ParentAT interface {
		*ParentAT_
		AbstractType
	},
	ParentAT_ Gongstruct,
	CT interface {
		*CT_
		RectShapeInterface
		ConcreteType
	},
	CT_ Gongstruct](
	stager *Stager,
	node *tree.Node,
	diagram DiagramIF,
	element AT,
	parentElement ParentAT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap map[AT]CT,
) {
	node.OnIsCheckedChanged = func(isChecked bool) {
		shape := shapesMap[element]
		if isChecked {
			if shape != nil {
				log.Panic("adding a shape to an already existing shape")
			}
			shape = newShapeToDiagram(element, diagram, shapes, stager, node.ClientOnY)
			shapesMap[element] = shape
			stager.stage.Commit()
		} else {
			if shape == nil {
				log.Panic("remove a non existing shape to diagram")
			}
			shape.UnstageVoid(stager.stage)
			idx := slices.Index(*shapes, shape)
			*shapes = slices.Delete(*shapes, idx, idx+1)
			stager.stage.Commit()
		}
	}

	node.OnIsExpandedChange = func(isExpanded bool) {
		if isExpanded {
			if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
				*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
			}
		} else {
			if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
				*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
			}
		}
		stager.stage.Commit()
	}

	node.OnNameChange = func(newName string) {
		element.SetName(newName)
		element.SetIsInRenameMode(false)
		stager.stage.Commit()
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
	}

	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
	}
}

func setCallbacksExpandableNode[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
](
	stager *Stager,
	node *tree.Node,
	element AT,
	elementsWhoseNodeIsExpanded *[]AT,
) {
	node.OnIsExpandedChange = func(isExpanded bool) {
		if isExpanded {
			if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
				*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
			}
		} else {
			if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
				*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
			}
		}
		stager.stage.Commit()
	}

	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
	}
}

func newShapeToDiagram[AT AbstractType, CT interface {
	*S
	RectShapeInterface
	ConcreteType
}, S Gongstruct](
	abstractElement AT,
	diagram DiagramIF,
	shapes *[]CT,
	stager *Stager,
	clientOnY float64,
) CT {
	shape := CT(new(S))
	shape.StageVoid(stager.stage)
	shape.SetAbstractElement(abstractElement)
	shape.SetName(abstractElement.GetName() + "-" + diagram.GetName())
	shape.SetHeight(diagram.GetDefaultBoxHeigth())
	shape.SetWidth(diagram.GetDefaultBoxWidth())

	zoom := stager.GetSvgObject().Zoom
	if zoom == 0 {
		zoom = 1.0
	}
	panX := stager.GetSvgObject().PanX
	panY := stager.GetSvgObject().PanY

	shape.SetX((100+panX)/zoom + rand.Float64()*100.0)

	if clientOnY != 0 {
		shape.SetY((clientOnY + panY) / zoom)
	} else {
		shape.SetY((100+panY)/zoom + rand.Float64()*100.0)
	}
	*shapes = append(*shapes, shape)

	return shape
}

func svgRect[CT interface {
	*CT_
	RectShapeInterface
	ConcreteType
}, CT_ Gongstruct](
	stager *Stager,
	diagram DiagramIF, // might be ni
	shape CT,
	layer *svg.Layer,
) *svg.Rect {
	root := stager.GetRootLibrary()

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

		rect.OnSelect = onSelectRectElement(stager, abstractElement)
		rect.OnMove = onMoveRectElement(stager, shape, true)
		rect.OnResize = onResizeRectElement(stager, shape)
		// for allowing later Stage() on the rect shape
		shape.SetReceiver(shape)
	}

	title := new(svg.RectAnchoredText)
	{
		title.Name = abstractElement.GetName()

		content := shape.GetAbstractElement().GetName()
		if diagram.GetIsShowPrefix() {
			content = abstractElement.GetComputedPrefix() + " " + content
		}

		if rect.Width > 0 {
			content = strutils.WrapStringPreservingNewlines(content, int(rect.Width/root.NbPixPerCharacter))
		}
		title.Content = content
		title.Stroke = svg.Black.ToString()
		title.StrokeWidth = 1
		title.StrokeOpacity = 1
		title.Color = svg.Black.ToString()
		title.FillOpacity = 1

		title.FontSize = "16px"
		title.X_Offset = 0
		title.Y_Offset = 0
		title.RectAnchorType = svg.RECT_CENTER_MIDDLE
		title.TextAnchorType = svg.TEXT_ANCHOR_CENTER

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, title)
	}

	layer.Rects = append(layer.Rects, rect)
	return rect
}

func onSelectRectElement[AT AbstractType](
	stager *Stager,
	abstractElement AT,
) func() {
	return func() {
		stager.stage.CommitWithSuspendedCallbacks()
		stager.probeForm.FillUpFormFromGongstruct(abstractElement, GetPointerToGongstructName[AT]())
		// update the tree because it contains the undo/redo calls
		stager.ux_tree()
	}
}

func onMoveRectElement[CT interface {
	*CT_
	RectShapeInterface
	ConcreteType
}, CT_ Gongstruct](
	stager *Stager,
	rectShape CT,
	suspendCallbacksOnPositionDiff bool,
) func(x, y float64) {
	return func(x, y float64) {
		rectShape.SetX(x)
		rectShape.SetY(y)

		if suspendCallbacksOnPositionDiff {
			// Issue #7, this will allow multiple rect to be moved together
			stager.stage.CommitWithSuspendedCallbacks()
			// update the tree because it contains the undo/redo calls
			stager.ux_tree()
		} else {
			// if the position is different, no need to generates
			// the SVG updates with the semantic rules updates.
			stager.stage.CommitWithSuspendedCallbacks()
		}
	}
}

func onResizeRectElement[CT interface {
	*CT_
	RectShapeInterface
	ConcreteType
}, CT_ Gongstruct](
	stager *Stager,
	rectShape CT,
) func(x, y, width, height float64) {
	return func(x, y, width, height float64) {
		rectShape.SetX(x)
		rectShape.SetY(y)
		rectShape.SetWidth(width)
		rectShape.SetHeight(height)

		stager.stage.Commit()
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
) (link *svg.Link) {
	if startRect == nil || endRect == nil {
		return
	}

	link = new(svg.Link)

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
	link.CornerRadius = 5
	if isDashed {
		link.StrokeDashArray = "5 5"
	}

	link.OnSelect = func() {
		stager.stage.CommitWithSuspendedCallbacks()
		stager.probeForm.FillUpFormFromGongstruct(productOfInterest, GetPointerToGongstructName[AT]())
	}

	link.OnChange = func(updatedLink *svg.Link) {
		shape.SetStartRatio(updatedLink.StartRatio)
		shape.SetEndRatio(updatedLink.EndRatio)
		shape.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
		shape.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
		shape.SetEndOrientation(OrientationType(updatedLink.EndOrientation))

		stager.stage.Commit()
	}

	layer.Links = append(layer.Links, link)

	return
}

func svgAssociationLinkAsCT[AT AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
		ConcreteType
	},
	ACT_ Gongstruct](
	stager *Stager,
	startRect *svg.Rect,
	endRect *svg.Rect,
	shape ACT,
	layer *svg.Layer,
	isDashed bool,
) (link *svg.Link) {
	if startRect == nil || endRect == nil {
		return
	}

	link = new(svg.Link)

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
	link.CornerRadius = 5
	if isDashed {
		link.StrokeDashArray = "5 5"
	}

	link.OnSelect = func() {
		stager.stage.CommitWithSuspendedCallbacks()
		stager.probeForm.FillUpFormFromGongstruct(shape.GetAbstractElement().(AT), GetPointerToGongstructName[AT]())
	}

	link.OnChange = func(updatedLink *svg.Link) {
		shape.SetStartRatio(updatedLink.StartRatio)
		shape.SetEndRatio(updatedLink.EndRatio)
		shape.SetCornerOffsetRatio(updatedLink.CornerOffsetRatio)
		shape.SetStartOrientation(OrientationType(updatedLink.StartOrientation))
		shape.SetEndOrientation(OrientationType(updatedLink.EndOrientation))

		stager.stage.Commit()
	}

	layer.Links = append(layer.Links, link)

	return
}
