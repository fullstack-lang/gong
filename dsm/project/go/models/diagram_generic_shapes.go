package models

func (diagram *Diagram) GetNodes() (nodes []ConcreteType) {
	for _, s := range diagram.Product_Shapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagram.Task_Shapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagram.Note_Shapes {
		nodes = append(nodes, s)
	}
	for _, s := range diagram.Resource_Shapes {
		nodes = append(nodes, s)
	}
	return
}

func (diagram *Diagram) AddNode(node ConcreteType) {
	switch n := node.(type) {
	case *ProductShape:
		diagram.Product_Shapes = append(diagram.Product_Shapes, n)
	case *TaskShape:
		diagram.Task_Shapes = append(diagram.Task_Shapes, n)
	case *NoteShape:
		diagram.Note_Shapes = append(diagram.Note_Shapes, n)
	case *ResourceShape:
		diagram.Resource_Shapes = append(diagram.Resource_Shapes, n)
	}
}

func (diagram *Diagram) GetLinks() (links []AssociationConcreteType) {
	for _, s := range diagram.ProductComposition_Shapes {
		links = append(links, s)
	}
	for _, s := range diagram.TaskComposition_Shapes {
		links = append(links, s)
	}
	for _, s := range diagram.TaskInputShapes {
		links = append(links, s)
	}
	for _, s := range diagram.TaskOutputShapes {
		links = append(links, s)
	}
	for _, s := range diagram.NoteProductShapes {
		links = append(links, s)
	}
	for _, s := range diagram.NoteTaskShapes {
		links = append(links, s)
	}
	for _, s := range diagram.NoteResourceShapes {
		links = append(links, s)
	}
	for _, s := range diagram.ResourceComposition_Shapes {
		links = append(links, s)
	}
	for _, s := range diagram.ResourceTaskShapes {
		links = append(links, s)
	}
	return
}

func (diagram *Diagram) AddLink(link AssociationConcreteType) {
	switch l := link.(type) {
	case *ProductCompositionShape:
		diagram.ProductComposition_Shapes = append(diagram.ProductComposition_Shapes, l)
	case *TaskCompositionShape:
		diagram.TaskComposition_Shapes = append(diagram.TaskComposition_Shapes, l)
	case *TaskInputShape:
		diagram.TaskInputShapes = append(diagram.TaskInputShapes, l)
	case *TaskOutputShape:
		diagram.TaskOutputShapes = append(diagram.TaskOutputShapes, l)
	case *NoteProductShape:
		diagram.NoteProductShapes = append(diagram.NoteProductShapes, l)
	case *NoteTaskShape:
		diagram.NoteTaskShapes = append(diagram.NoteTaskShapes, l)
	case *NoteResourceShape:
		diagram.NoteResourceShapes = append(diagram.NoteResourceShapes, l)
	case *ResourceCompositionShape:
		diagram.ResourceComposition_Shapes = append(diagram.ResourceComposition_Shapes, l)
	case *ResourceTaskShape:
		diagram.ResourceTaskShapes = append(diagram.ResourceTaskShapes, l)
	}
}
