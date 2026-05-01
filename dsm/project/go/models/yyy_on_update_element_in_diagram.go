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

						shape.SetX(parentShape.GetX() + float64(siblingsCount-1)*parentShape.GetWidth()*1.2)
						shape.SetY(parentShape.GetY() + parentShape.GetHeight()*2.0)
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

		if frontNode.IsSecondCheckboxChecked && !stagedNode.IsSecondCheckboxChecked {
			compositionShape := newConcreteAssociation(parentElement, element, compositionShapes)
			compositionShape.StageVoid(stager.stage)
		}
		if !frontNode.IsSecondCheckboxChecked && stagedNode.IsSecondCheckboxChecked {
			compositionShape.UnstageVoid(stager.stage)
			idx := slices.Index(*compositionShapes, compositionShape)
			*compositionShapes = slices.Delete(*compositionShapes, idx, idx+1)
			stager.stage.Commit()
		}

		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
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
			if any(shape) != nil {
				log.Panic("adding a shape to an already existing shape")
			}
			shape = newShapeToDiagram(element, diagram, shapes, stager.stage)

			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if any(shape) == nil {
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
