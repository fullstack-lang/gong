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
	if taskShape, ok := diagram.map_SvgRect_TaskShape[startRect]; ok {
		if subTaskShape, ok := diagram.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_COMPOSITION
			sourceAbstratctElement = taskShape.Task
			targetAbstractElement = subTaskShape.Task
		}
	}
	if taskShape, ok := diagram.map_SvgRect_TaskShape[startRect]; ok {
		if productShape, ok := diagram.map_SvgRect_ProductShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_OUTPUT
			sourceAbstratctElement = taskShape.Task
			targetAbstractElement = productShape.Product
		}
	}
	if productShape, ok := diagram.map_SvgRect_ProductShape[startRect]; ok {
		if taskShape, ok := diagram.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_TASK_INPUT
			sourceAbstratctElement = productShape.Product
			targetAbstractElement = taskShape.Task
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
		if taskShape, ok := diagram.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIAITON_TYPE_NOTE_TASK
			sourceAbstratctElement = noteShape.Note
			targetAbstractElement = taskShape.Task
		}
	}
	if noteShape, ok := diagram.map_SvgRect_NoteShape[startRect]; ok {
		if resourceShape, ok := diagram.map_SvgRect_ResourceShape[endRect]; ok {
			assocType = ASSOCIAITON_TYPE_NOTE_RESOURCE
			sourceAbstratctElement = noteShape.Note
			targetAbstractElement = resourceShape.Resource
		}
	}

	if resourceShape, ok := diagram.map_SvgRect_ResourceShape[startRect]; ok {
		if subResourceShape, ok := diagram.map_SvgRect_ResourceShape[endRect]; ok {
			assocType = ASSOCATIONN_TYPE_RESOURCE_COMPOSITION
			sourceAbstratctElement = resourceShape.Resource
			targetAbstractElement = subResourceShape.Resource
		}
	}
	if resourceShape, ok := diagram.map_SvgRect_ResourceShape[startRect]; ok {
		if taskShape, ok := diagram.map_SvgRect_TaskShape[endRect]; ok {
			assocType = ASSOCIATION_TYPE_RESOURCE_TASK
			sourceAbstratctElement = resourceShape.Resource
			targetAbstractElement = taskShape.Task
		}
	}

	// create the association according to the association type
	switch assocType {
	case ASSOCIATION_TYPE_PRODUCT_COMPOSITION:
		subProduct := targetAbstractElement.(*Product)
		parentProduct := sourceAbstratctElement.(*Product)

		if subProduct.parentProduct != nil {
			oldParent := subProduct.parentProduct
			if idx := slices.Index(oldParent.SubProducts, subProduct); idx != -1 {
				oldParent.SubProducts = slices.Delete(oldParent.SubProducts, idx, idx+1)
			}
		} else {
			if idx := slices.Index(subProduct.GetOwningLibrary().RootProducts, subProduct); idx != -1 {
				subProduct.GetOwningLibrary().RootProducts = slices.Delete(subProduct.GetOwningLibrary().RootProducts, idx, idx+1)
			}
		}

		parentProduct.SubProducts = append(parentProduct.SubProducts, subProduct)
		subProduct.parentProduct = parentProduct
		addAssociationShapeToDiagram(stager, parentProduct, subProduct, &diagram.ProductComposition_Shapes)

	case ASSOCIATION_TYPE_TASK_COMPOSITION:
		subTask := targetAbstractElement.(*Task)
		parentTask := sourceAbstratctElement.(*Task)

		if subTask.parentTask != nil {
			oldParent := subTask.parentTask
			if idx := slices.Index(oldParent.SubTasks, subTask); idx != -1 {
				oldParent.SubTasks = slices.Delete(oldParent.SubTasks, idx, idx+1)
			}
		} else {
			if idx := slices.Index(subTask.GetOwningLibrary().RootTasks, subTask); idx != -1 {
				subTask.GetOwningLibrary().RootTasks = slices.Delete(subTask.GetOwningLibrary().RootTasks, idx, idx+1)
			}
		}

		subTask.parentTask = parentTask
		parentTask.SubTasks = append(parentTask.SubTasks, subTask)
		addAssociationShapeToDiagram(stager, parentTask, subTask, &diagram.TaskComposition_Shapes)

	case ASSOCIATION_TYPE_TASK_INPUT:
		product := sourceAbstratctElement.(*Product)
		task := targetAbstractElement.(*Task)

		task.Inputs = append(task.Inputs, product)
		addAssociationShapeToDiagram(stager, task, product, &diagram.TaskInputShapes)

	case ASSOCIATION_TYPE_TASK_OUTPUT:
		task := sourceAbstratctElement.(*Task)
		product := targetAbstractElement.(*Product)

		task.Outputs = append(task.Outputs, product)
		addAssociationShapeToDiagram(stager, task, product, &diagram.TaskOutputShapes)

	case ASSOCIAITON_TYPE_NOTE_PRODUCT:
		note := sourceAbstratctElement.(*Note)
		product := targetAbstractElement.(*Product)

		note.Products = append(note.Products, product)
		addAssociationShapeToDiagram(stager, note, product, &diagram.NoteProductShapes)

	case ASSOCIAITON_TYPE_NOTE_TASK:
		note := sourceAbstratctElement.(*Note)
		task := targetAbstractElement.(*Task)

		note.Tasks = append(note.Tasks, task)
		addAssociationShapeToDiagram(stager, note, task, &diagram.NoteTaskShapes)

	case ASSOCIAITON_TYPE_NOTE_RESOURCE:
		note := sourceAbstratctElement.(*Note)
		resource := targetAbstractElement.(*Resource)

		note.Resources = append(note.Resources, resource)
		addAssociationShapeToDiagram(stager, note, resource, &diagram.NoteResourceShapes)

	case ASSOCATIONN_TYPE_RESOURCE_COMPOSITION:
		subResource := targetAbstractElement.(*Resource)
		parentResource := sourceAbstratctElement.(*Resource)

		if subResource.parentResource != nil {
			oldParent := subResource.parentResource
			if idx := slices.Index(oldParent.SubResources, subResource); idx != -1 {
				oldParent.SubResources = slices.Delete(oldParent.SubResources, idx, idx+1)
			}
		} else {
			if idx := slices.Index(subResource.GetOwningLibrary().RootResources, subResource); idx != -1 {
				subResource.GetOwningLibrary().RootResources = slices.Delete(subResource.GetOwningLibrary().RootResources, idx, idx+1)
			}
		}

		subResource.parentResource = parentResource
		parentResource.SubResources = append(parentResource.SubResources, subResource)
		addAssociationShapeToDiagram(stager, parentResource, subResource, &diagram.ResourceComposition_Shapes)

	case ASSOCIATION_TYPE_RESOURCE_TASK:
		resource := sourceAbstratctElement.(*Resource)
		task := targetAbstractElement.(*Task)

		resource.Tasks = append(resource.Tasks, task)
		addAssociationShapeToDiagram(stager, resource, task, &diagram.ResourceTaskShapes)
	}

	if assocType == "" {
		log.Println("unsupported association")
		return
	}

	// commit to encode the result, this will generate a new SVG generation
	stager.stage.Commit()
}
