// generated code (do not edit)
package models

import (
	"log"
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func onUpdateElementInDiagram[
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
	diagram DiagramIF,
	element AT,
	parentElement AT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap map[AT]CT,
	compositionShape ACT,
	compositionShapes *[]ACT,
) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		// find the shape (if any)
		shape := shapesMap[element]

		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(element, diagram, shapes, stager.stage)
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
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stager.stage)
			idx := slices.Index(*shapes, shape)
			*shapes = slices.Delete(*shapes, idx, idx+1)
			stager.stage.Commit()
			return
		}

		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			if frontNode.IsExpanded {
				if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
					*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
				}
			} else {
				if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
					*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.CommitWithSuspendedCallbacks()
			return
		}

		if frontNode.Name != stagedNode.Name {
			element.SetName(frontNode.Name)
			element.SetIsInRenameMode(false)
			stager.stage.Commit()
			stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
			return
		}

		if frontNode.IsSecondCheckboxChecked && !stagedNode.IsSecondCheckboxChecked {
			compositionShape := newConcreteAssociation(parentElement, element, compositionShapes)
			if parentShape, ok := shapesMap[parentElement]; ok && shape != nil {
				ratio := (shape.GetY() - parentShape.GetY()) / parentShape.GetHeight()
				compositionShape.SetCornerOffsetRatio((ratio-1.0)/2.0 + 1.0)
			}
			compositionShape.StageVoid(stager.stage)
		}
		if !frontNode.IsSecondCheckboxChecked && stagedNode.IsSecondCheckboxChecked {
			compositionShape.UnstageVoid(stager.stage)
			idx := slices.Index(*compositionShapes, compositionShape)
			*compositionShapes = slices.Delete(*compositionShapes, idx, idx+1)
			stager.stage.Commit()
		}

		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.CommitWithSuspendedCallbacks()
	}
}

// onUpdateElementInDiagramWithoutLink handles the update event for an element in the diagram without a link
func onUpdateElementInDiagramWithoutLink[
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
	diagram DiagramIF,
	element AT,
	parentElement ParentAT,
	elementsWhoseNodeIsExpanded *[]AT,
	shapes *[]CT,
	shapesMap map[AT]CT,
) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		// find the shape (if any)
		shape := shapesMap[element]

		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already existing shape")
			}
			shape = newShapeToDiagram(element, diagram, shapes, stager.stage)
			shapesMap[element] = shape

			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to diagram")
			}
			shape.UnstageVoid(stager.stage)
			idx := slices.Index(*shapes, shape)
			*shapes = slices.Delete(*shapes, idx, idx+1)
			stager.stage.Commit()
			return
		}

		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			if frontNode.IsExpanded {
				if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
					*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
				}
			} else {
				if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
					*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}

		if frontNode.Name != stagedNode.Name {
			element.SetName(frontNode.Name)
			element.SetIsInRenameMode(false)
			stager.stage.Commit()
			stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
			return
		}

		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
	}
}
