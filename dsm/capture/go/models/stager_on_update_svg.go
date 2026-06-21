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
		ASSOCIATION_TYPE_DELIVERABLE_COMPOSITION associationType = "DeliverableComposition"

		ASSOCIATION_TYPE_TASK_COMPOSITION associationType = "TaskComposition"
		ASSOCIATION_TYPE_TASK_INPUT       associationType = "TaskInput"
		ASSOCIATION_TYPE_TASK_OUTPUT      associationType = "TaskOutput"

		ASSOCIAITON_TYPE_NOTE_DELIVERABLE associationType = "NoteDeliverable"
		ASSOCIAITON_TYPE_NOTE_TASK        associationType = "NoteTask"
		ASSOCIAITON_TYPE_NOTE_RESOURCE    associationType = "NoteResource"

		ASSOCATIONN_TYPE_RESOURCE_COMPOSITION associationType = "ResourceComposition"
		ASSOCIATION_TYPE_RESOURCE_TASK        associationType = "ResourceTask"

		ASSOCIATION_TYPE_DELIVERABLE_CONCEPT associationType = "DeliverableConcept"
	)

	var assocType associationType
	var sourceAbstratctElement, targetAbstractElement AbstractType
	startRect := frontSVG.StartRect
	endRect := frontSVG.EndRect

	// determine the association type from the source and target rects
	if deliverableShape, ok := diagram.map_SvgRect_DeliverableShape[startRect]; ok {
		if subDeliverableShape, ok := diagram.map_SvgRect_DeliverableShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_DELIVERABLE_COMPOSITION
			sourceAbstratctElement = deliverableShape.Deliverable
			targetAbstractElement = subDeliverableShape.Deliverable
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
		if deliverableShape, ok := diagram.map_SvgRect_DeliverableShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_OUTPUT
			sourceAbstratctElement = taskShape.Concern
			targetAbstractElement = deliverableShape.Deliverable
		}
	}
	if deliverableShape, ok := diagram.map_SvgRect_DeliverableShape[startRect]; ok {
		if taskShape, ok := diagram.map_SvgRect_ConcernShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_INPUT
			sourceAbstratctElement = deliverableShape.Deliverable
			targetAbstractElement = taskShape.Concern
		}
	}
	if noteShape, ok := diagram.map_SvgRect_NoteShape[startRect]; ok {
		if deliverableShape, ok := diagram.map_SvgRect_DeliverableShape[endRect]; ok {
			assocType = ASSOCIAITON_TYPE_NOTE_DELIVERABLE
			sourceAbstratctElement = noteShape.Note
			targetAbstractElement = deliverableShape.Deliverable
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
	if deliverableShape, ok := diagram.map_SvgRect_DeliverableShape[startRect]; ok {
		if conceptShape, ok := diagram.map_SvgRect_ConceptShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_DELIVERABLE_CONCEPT
			sourceAbstratctElement = deliverableShape.Deliverable
			targetAbstractElement = conceptShape.Concept
		}
	}

	// create the association according to the association type
	switch assocType {
	case ASSOCIATION_TYPE_DELIVERABLE_COMPOSITION:
		subDeliverable := targetAbstractElement.(*Deliverable)
		parentDeliverable := sourceAbstratctElement.(*Deliverable)

		if subDeliverable.parentDeliverable != nil {
			oldParent := subDeliverable.parentDeliverable
			if idx := slices.Index(oldParent.SubDeliverables, subDeliverable); idx != -1 {
				oldParent.SubDeliverables = slices.Delete(oldParent.SubDeliverables, idx, idx+1)
			}
		} else {
			if idx := slices.Index(subDeliverable.GetOwningLibrary().RootDeliverables, subDeliverable); idx != -1 {
				subDeliverable.GetOwningLibrary().RootDeliverables = slices.Delete(subDeliverable.GetOwningLibrary().RootDeliverables, idx, idx+1)
			}
		}

		parentDeliverable.SubDeliverables = append(parentDeliverable.SubDeliverables, subDeliverable)
		subDeliverable.parentDeliverable = parentDeliverable
		addAssociationShapeToDiagram(stager, parentDeliverable, subDeliverable, &diagram.DeliverableComposition_Shapes)

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
		deliverable := sourceAbstratctElement.(*Deliverable)
		task := targetAbstractElement.(*Concern)

		task.Inputs = append(task.Inputs, deliverable)
		addAssociationShapeToDiagram(stager, task, deliverable, &diagram.ConcernInputShapes)

	case ASSOCIATION_TYPE_TASK_OUTPUT:
		task := sourceAbstratctElement.(*Concern)
		deliverable := targetAbstractElement.(*Deliverable)

		task.Outputs = append(task.Outputs, deliverable)
		addAssociationShapeToDiagram(stager, task, deliverable, &diagram.ConcernOutputShapes)

	case ASSOCIAITON_TYPE_NOTE_DELIVERABLE:
		note := sourceAbstratctElement.(*Note)
		deliverable := targetAbstractElement.(*Deliverable)

		note.Deliverables = append(note.Deliverables, deliverable)
		addAssociationShapeToDiagram(stager, note, deliverable, &diagram.NoteDeliverableShapes)

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
