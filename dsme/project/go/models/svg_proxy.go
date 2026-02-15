package models

import (
	"log"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type svgProxy struct {
	stager  *Stager
	diagram *Diagram
}

// SVGUpdated implements SVGImplInterface.
func (p *svgProxy) SVGUpdated(updatedSVG *svg.SVG) {

	diagram := p.diagram

	if updatedSVG.DrawingState != svg.DRAWING_LINK {
		// in any cases, have the form editor set up with the instance
		p.stager.probeForm.FillUpFormFromGongstruct(p.diagram, "Diagram")
		return
	}

	type associationType string
	const (
		ASSOCIATION_TYPE_PRODUCT_COMPOSITION  associationType = "ProductComposition"
		ASSOCIATION_TYPE_TASK_COMPOSITION     associationType = "TaskComposition"
		ASSOCIATION_TYPE_TASK_INPUT           associationType = "TaskInput"
		ASSOCIATION_TYPE_TASK_OUTPUT          associationType = "TaskOutput"
		ASSOCIAITON_TYPE_NOTE_PRODUCT         associationType = "NoteProduct"
		ASSOCIAITON_TYPE_NOTE_TASK            associationType = "NoteTask"
		ASSOCATIONN_TYPE_RESOURCE_COMPOSITION associationType = "ResourceComposition"
		ASSOCIATION_TYPE_RESOURCE_TASK        associationType = "ResourceTask"
	)

	var assocType associationType
	var sourceAbstratctElement, targetAbstractElement AbstractType
	startRect := updatedSVG.StartRect
	endRect := updatedSVG.EndRect

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

		parentProduct.SubProducts = append(parentProduct.SubProducts, subProduct)
		onAddAssociationShape(p.stager, parentProduct, subProduct, &diagram.ProductComposition_Shapes)

	case ASSOCIATION_TYPE_TASK_COMPOSITION:
		subTask := targetAbstractElement.(*Task)
		parentTask := sourceAbstratctElement.(*Task)

		parentTask.SubTasks = append(parentTask.SubTasks, subTask)
		addAssociationShapeToDiagram(p.stager, parentTask, subTask, &diagram.TaskComposition_Shapes)

	case ASSOCIATION_TYPE_TASK_INPUT:
		product := sourceAbstratctElement.(*Product)
		task := targetAbstractElement.(*Task)

		task.Inputs = append(task.Inputs, product)
		addAssociationShapeToDiagram(p.stager, task, product, &diagram.TaskInputShapes)

	case ASSOCIATION_TYPE_TASK_OUTPUT:
		task := sourceAbstratctElement.(*Task)
		product := targetAbstractElement.(*Product)

		task.Outputs = append(task.Outputs, product)
		addAssociationShapeToDiagram(p.stager, task, product, &diagram.TaskOutputShapes)

	case ASSOCIAITON_TYPE_NOTE_PRODUCT:
		note := sourceAbstratctElement.(*Note)
		product := targetAbstractElement.(*Product)

		note.Products = append(note.Products, product)
		addAssociationShapeToDiagram(p.stager, note, product, &diagram.NoteProductShapes)

	case ASSOCIAITON_TYPE_NOTE_TASK:
		note := sourceAbstratctElement.(*Note)
		task := targetAbstractElement.(*Task)

		note.Tasks = append(note.Tasks, task)
		addAssociationShapeToDiagram(p.stager, note, task, &diagram.NoteTaskShapes)

	case ASSOCATIONN_TYPE_RESOURCE_COMPOSITION:
		subResource := targetAbstractElement.(*Resource)
		parentResource := sourceAbstratctElement.(*Resource)

		parentResource.SubResources = append(parentResource.SubResources, subResource)
		addAssociationShapeToDiagram(p.stager, parentResource, subResource, &diagram.ResourceComposition_Shapes)

	case ASSOCIATION_TYPE_RESOURCE_TASK:
		resource := sourceAbstratctElement.(*Resource)
		task := targetAbstractElement.(*Task)

		resource.Tasks = append(resource.Tasks, task)
		addAssociationShapeToDiagram(p.stager, resource, task, &diagram.ResourceTaskShapes)
	}

	if assocType == "" {
		log.Println("unsupported association")
		return
	}

	// commit to encode the result, this will generate a new SVG generation
	p.stager.stage.Commit()

}

func (p *svgProxy) OnAfterUpdate(stage *svg.Stage, stagedSVG, frontSVG *svg.SVG) {

	// log.Println("SVG updated", stagedSVG.GetName())

}
