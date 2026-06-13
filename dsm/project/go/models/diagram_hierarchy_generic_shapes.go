package models

func (diagramHierarchy *DiagramHierarchy) GetNodes() (nodes []ConcreteType) {
	for _, s := range diagramHierarchy.Product_Shapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagramHierarchy.Task_Shapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagramHierarchy.Note_Shapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagramHierarchy.Resource_Shapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagramHierarchy.TaskGroupShapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagramHierarchy.MilestoneShapes {
		nodes = append(nodes, s)
	}
	return
}

func (diagramHierarchy *DiagramHierarchy) AddNode(node ConcreteType) {
	switch n := node.(type) {
	case *ProductShape:
		diagramHierarchy.Product_Shapes = append(diagramHierarchy.Product_Shapes, n)
	case *TaskShape:
		diagramHierarchy.Task_Shapes = append(diagramHierarchy.Task_Shapes, n)
	case *NoteShape:
		diagramHierarchy.Note_Shapes = append(diagramHierarchy.Note_Shapes, n)
	case *ResourceShape:
		diagramHierarchy.Resource_Shapes = append(diagramHierarchy.Resource_Shapes, n)
	case *TaskGroupShape:
		diagramHierarchy.TaskGroupShapes = append(diagramHierarchy.TaskGroupShapes, n)
	case *MilestoneShape:
		diagramHierarchy.MilestoneShapes = append(diagramHierarchy.MilestoneShapes, n)
	}
}

func (diagramHierarchy *DiagramHierarchy) GetLinks() (links []AssociationConcreteType) {
	for _, s := range diagramHierarchy.ProductComposition_Shapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.TaskComposition_Shapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.TaskInputShapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.TaskOutputShapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.NoteProductShapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.NoteTaskShapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.NoteResourceShapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.ResourceComposition_Shapes {
		links = append(links, s)
	}
	for _, s := range diagramHierarchy.ResourceTaskShapes {
		links = append(links, s)
	}
	return
}

func (diagramHierarchy *DiagramHierarchy) AddLink(link AssociationConcreteType) {
	switch l := link.(type) {
	case *ProductCompositionShape:
		diagramHierarchy.ProductComposition_Shapes = append(diagramHierarchy.ProductComposition_Shapes, l)
	case *TaskCompositionShape:
		diagramHierarchy.TaskComposition_Shapes = append(diagramHierarchy.TaskComposition_Shapes, l)
	case *TaskInputShape:
		diagramHierarchy.TaskInputShapes = append(diagramHierarchy.TaskInputShapes, l)
	case *TaskOutputShape:
		diagramHierarchy.TaskOutputShapes = append(diagramHierarchy.TaskOutputShapes, l)
	case *NoteProductShape:
		diagramHierarchy.NoteProductShapes = append(diagramHierarchy.NoteProductShapes, l)
	case *NoteTaskShape:
		diagramHierarchy.NoteTaskShapes = append(diagramHierarchy.NoteTaskShapes, l)
	case *NoteResourceShape:
		diagramHierarchy.NoteResourceShapes = append(diagramHierarchy.NoteResourceShapes, l)
	case *ResourceCompositionShape:
		diagramHierarchy.ResourceComposition_Shapes = append(diagramHierarchy.ResourceComposition_Shapes, l)
	case *ResourceTaskShape:
		diagramHierarchy.ResourceTaskShapes = append(diagramHierarchy.ResourceTaskShapes, l)
	}
}
