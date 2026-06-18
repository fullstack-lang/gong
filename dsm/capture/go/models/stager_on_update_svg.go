package models

import (
	"log"
	"slices"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

// SVGUpdated implements SVGImplInterface.
func (stager *Stager) onUpdateSVG(frontSVG *svg.SVG) {
	diagram := stager.diagram
	svgObject := stager.svgObject

	if svgObject.DrawingState == frontSVG.DrawingState {
		// in any cases, have the form editor set up with the instance
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
		return
	}

	// IMPORTANT : we are only interested when the updateSVG has finished drawing the connexion
	if frontSVG.DrawingState == svg.DRAWING_LINK {
		return
	}

	type associationType string
	const (
		ASSOCIATION_TYPE_PRODUCT_COMPOSITION associationType = "ProductComposition"

		ASSOCIATION_TYPE_TASK_COMPOSITION associationType = "TaskComposition"
		ASSOCIATION_TYPE_TASK_INPUT       associationType = "TaskInput"
		ASSOCIATION_TYPE_TASK_OUTPUT      associationType = "TaskOutput"

		ASSOCIAITON_TYPE_NOTE_PRODUCT  associationType = "NoteProduct"
		ASSOCIAITON_TYPE_NOTE_TASK     associationType = "NoteTask"
		ASSOCIAITON_TYPE_NOTE_RESOURCE associationType = "NoteResource"

		ASSOCATIONN_TYPE_RESOURCE_COMPOSITION associationType = "ResourceComposition"
		ASSOCIATION_TYPE_RESOURCE_TASK        associationType = "ResourceTask"

		ASSOCIATION_TYPE_DELIVERABLE_CONCEPT associationType = "DeliverableConcept"
	)

	var assocType associationType
	var sourceAbstratctElement, targetAbstractElement AbstractType
	startRect := frontSVG.StartRect
	endRect := frontSVG.EndRect

	// determine the association type from the source and target rects
	if productShape, ok := diagram.map_SvgRect_ProductShape[startRect]; ok {
		if subProductShape, ok := diagram.map_SvgRect_ProductShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_PRODUCT_COMPOSITION
			sourceAbstratctElement = productShape.Product
			targetAbstractElement = subProductShape.Product
		}
	}
	if taskShape, ok := diagram.map_SvgRect_ConcernShape[startRect]; ok {
		if subTaskShape, ok := diagram.map_SvgRect_ConcernShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_COMPOSITION
			sourceAbstratctElement = taskShape.Concern
			targetAbstractElement = subTaskShape.Concern
		}
	}
	if taskShape, ok := diagram.map_SvgRect_ConcernShape[startRect]; ok {
		if productShape, ok := diagram.map_SvgRect_ProductShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_OUTPUT
			sourceAbstratctElement = taskShape.Concern
			targetAbstractElement = productShape.Product
		}
	}
	if productShape, ok := diagram.map_SvgRect_ProductShape[startRect]; ok {
		if taskShape, ok := diagram.map_SvgRect_ConcernShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_INPUT
			sourceAbstratctElement = productShape.Product
			targetAbstractElement = taskShape.Concern
		}
	}
	if noteShape, ok := diagram.map_SvgRect_NoteShape[startRect]; ok {
		if productShape, ok := diagram.map_SvgRect_ProductShape[endRect]; ok {
			assocType = ASSOCIAITON_TYPE_NOTE_PRODUCT
			sourceAbstratctElement = noteShape.Note
			targetAbstractElement = productShape.Product
		}
	}
	if noteShape, ok := diagram.map_SvgRect_NoteShape[startRect]; ok {
		if taskShape, ok := diagram.map_SvgRect_ConcernShape[endRect]; ok {
			assocType = ASSOCIAITON_TYPE_NOTE_TASK
			sourceAbstratctElement = noteShape.Note
			targetAbstractElement = taskShape.Concern
		}
	}
	if noteShape, ok := diagram.map_SvgRect_NoteShape[startRect]; ok {
		if resourceShape, ok := diagram.map_SvgRect_StakeholderShape[endRect]; ok {
			assocType = ASSOCIAITON_TYPE_NOTE_RESOURCE
			sourceAbstratctElement = noteShape.Note
			targetAbstractElement = resourceShape.Stakeholder
		}
	}

	if resourceShape, ok := diagram.map_SvgRect_StakeholderShape[startRect]; ok {
		if subResourceShape, ok := diagram.map_SvgRect_StakeholderShape[endRect]; ok {
			assocType = ASSOCATIONN_TYPE_RESOURCE_COMPOSITION
			sourceAbstratctElement = resourceShape.Stakeholder
			targetAbstractElement = subResourceShape.Stakeholder
		}
	}
	if resourceShape, ok := diagram.map_SvgRect_StakeholderShape[startRect]; ok {
		if taskShape, ok := diagram.map_SvgRect_ConcernShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_RESOURCE_TASK
			sourceAbstratctElement = resourceShape.Stakeholder
			targetAbstractElement = taskShape.Concern
		}
	}
	if productShape, ok := diagram.map_SvgRect_ProductShape[startRect]; ok {
		if conceptShape, ok := diagram.map_SvgRect_ConceptShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_DELIVERABLE_CONCEPT
			sourceAbstratctElement = productShape.Product
			targetAbstractElement = conceptShape.Concept
		}
	}

	// create the association according to the association type
	switch assocType {
	case ASSOCIATION_TYPE_PRODUCT_COMPOSITION:
		subProduct := targetAbstractElement.(*Deliverable)
		parentProduct := sourceAbstratctElement.(*Deliverable)

		if subProduct.parentProduct != nil {
			oldParent := subProduct.parentProduct
			if idx := slices.Index(oldParent.SubProducts, subProduct); idx != -1 {
				oldParent.SubProducts = slices.Delete(oldParent.SubProducts, idx, idx+1)
			}
		} else {
			if idx := slices.Index(subProduct.GetOwningLibrary().RootDeliverables, subProduct); idx != -1 {
				subProduct.GetOwningLibrary().RootDeliverables = slices.Delete(subProduct.GetOwningLibrary().RootDeliverables, idx, idx+1)
			}
		}

		parentProduct.SubProducts = append(parentProduct.SubProducts, subProduct)
		subProduct.parentProduct = parentProduct
		addAssociationShapeToDiagram(stager, parentProduct, subProduct, &diagram.ProductComposition_Shapes)

	case ASSOCIATION_TYPE_TASK_COMPOSITION:
		subTask := targetAbstractElement.(*Concern)
		parentTask := sourceAbstratctElement.(*Concern)

		if subTask.parentConcern != nil {
			oldParent := subTask.parentConcern
			if idx := slices.Index(oldParent.SubConcerns, subTask); idx != -1 {
				oldParent.SubConcerns = slices.Delete(oldParent.SubConcerns, idx, idx+1)
			}
		} else {
			if idx := slices.Index(subTask.GetOwningLibrary().RootConcerns, subTask); idx != -1 {
				subTask.GetOwningLibrary().RootConcerns = slices.Delete(subTask.GetOwningLibrary().RootConcerns, idx, idx+1)
			}
		}

		subTask.parentConcern = parentTask
		parentTask.SubConcerns = append(parentTask.SubConcerns, subTask)
		addAssociationShapeToDiagram(stager, parentTask, subTask, &diagram.ConcernComposition_Shapes)

	case ASSOCIATION_TYPE_TASK_INPUT:
		product := sourceAbstratctElement.(*Deliverable)
		task := targetAbstractElement.(*Concern)

		task.Inputs = append(task.Inputs, product)
		addAssociationShapeToDiagram(stager, task, product, &diagram.ConcernInputShapes)

	case ASSOCIATION_TYPE_TASK_OUTPUT:
		task := sourceAbstratctElement.(*Concern)
		product := targetAbstractElement.(*Deliverable)

		task.Outputs = append(task.Outputs, product)
		addAssociationShapeToDiagram(stager, task, product, &diagram.ConcernOutputShapes)

	case ASSOCIAITON_TYPE_NOTE_PRODUCT:
		note := sourceAbstratctElement.(*Note)
		product := targetAbstractElement.(*Deliverable)

		note.Products = append(note.Products, product)
		addAssociationShapeToDiagram(stager, note, product, &diagram.NoteProductShapes)

	case ASSOCIAITON_TYPE_NOTE_TASK:
		note := sourceAbstratctElement.(*Note)
		task := targetAbstractElement.(*Concern)

		note.Tasks = append(note.Tasks, task)
		addAssociationShapeToDiagram(stager, note, task, &diagram.NoteTaskShapes)

	case ASSOCIAITON_TYPE_NOTE_RESOURCE:
		note := sourceAbstratctElement.(*Note)
		resource := targetAbstractElement.(*Stakeholder)

		note.Resources = append(note.Resources, resource)
		addAssociationShapeToDiagram(stager, note, resource, &diagram.NoteResourceShapes)

	case ASSOCATIONN_TYPE_RESOURCE_COMPOSITION:
		subResource := targetAbstractElement.(*Stakeholder)
		parentResource := sourceAbstratctElement.(*Stakeholder)

		if subResource.parentStakeholder != nil {
			oldParent := subResource.parentStakeholder
			if idx := slices.Index(oldParent.SubStakeholders, subResource); idx != -1 {
				oldParent.SubStakeholders = slices.Delete(oldParent.SubStakeholders, idx, idx+1)
			}
		} else {
			if idx := slices.Index(subResource.GetOwningLibrary().RootStakeholders, subResource); idx != -1 {
				subResource.GetOwningLibrary().RootStakeholders = slices.Delete(subResource.GetOwningLibrary().RootStakeholders, idx, idx+1)
			}
		}

		subResource.parentStakeholder = parentResource
		parentResource.SubStakeholders = append(parentResource.SubStakeholders, subResource)
		addAssociationShapeToDiagram(stager, parentResource, subResource, &diagram.ResourceComposition_Shapes)

	case ASSOCIATION_TYPE_RESOURCE_TASK:
		resource := sourceAbstratctElement.(*Stakeholder)
		task := targetAbstractElement.(*Concern)

		resource.Concerns = append(resource.Concerns, task)
		addAssociationShapeToDiagram(stager, resource, task, &diagram.StakeholderConcernShapes)

	case ASSOCIATION_TYPE_DELIVERABLE_CONCEPT:
		deliverable := sourceAbstratctElement.(*Deliverable)
		concept := targetAbstractElement.(*Concept)

		deliverable.Concepts = append(deliverable.Concepts, concept)
		addAssociationShapeToDiagram(stager, deliverable, concept, &diagram.DeliverableConceptShapes)
	}

	if assocType == "" {
		log.Println("unsupported association")
		return
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
