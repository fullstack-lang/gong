package models

func onCopyDiagram(stager *Stager, diagram *Diagram) func() {
	return func() {
		newDiagram := new(Diagram)
		newDiagram.Name = diagram.Name + " copy"
		newDiagram.IsEditable_ = true

		library := stager.stage.Library_Diagrams_reverseMap[diagram]
		Append(&library.Diagrams, newDiagram)

		for _, s := range diagram.Product_Shapes {
			newShape := s.GongCopy().(*ProductShape)
			newShape.Product = s.Product
			newDiagram.Product_Shapes = append(newDiagram.Product_Shapes, newShape)
		}

		for _, s := range diagram.ProductComposition_Shapes {
			newShape := s.GongCopy().(*ProductCompositionShape)
			newShape.Product = s.Product
			newDiagram.ProductComposition_Shapes = append(newDiagram.ProductComposition_Shapes, newShape)
		}

		for _, s := range diagram.Task_Shapes {
			newShape := s.GongCopy().(*TaskShape)
			newShape.Task = s.Task
			newDiagram.Task_Shapes = append(newDiagram.Task_Shapes, newShape)
		}

		for _, s := range diagram.TaskComposition_Shapes {
			newShape := s.GongCopy().(*TaskCompositionShape)
			newShape.Task = s.Task
			newDiagram.TaskComposition_Shapes = append(newDiagram.TaskComposition_Shapes, newShape)
		}

		for _, s := range diagram.TaskInputShapes {
			newShape := s.GongCopy().(*TaskInputShape)
			newShape.Task = s.Task
			newShape.Product = s.Product
			newDiagram.TaskInputShapes = append(newDiagram.TaskInputShapes, newShape)
		}

		for _, s := range diagram.TaskOutputShapes {
			newShape := s.GongCopy().(*TaskOutputShape)
			newShape.Task = s.Task
			newShape.Product = s.Product
			newDiagram.TaskOutputShapes = append(newDiagram.TaskOutputShapes, newShape)
		}

		for _, s := range diagram.Note_Shapes {
			newShape := s.GongCopy().(*NoteShape)
			newShape.Note = s.Note
			newDiagram.Note_Shapes = append(newDiagram.Note_Shapes, newShape)
		}

		for _, s := range diagram.NoteProductShapes {
			newShape := s.GongCopy().(*NoteProductShape)
			newShape.Note = s.Note
			newShape.Product = s.Product
			newDiagram.NoteProductShapes = append(newDiagram.NoteProductShapes, newShape)
		}

		for _, s := range diagram.NoteTaskShapes {
			newShape := s.GongCopy().(*NoteTaskShape)
			newShape.Note = s.Note
			newShape.Task = s.Task
			newDiagram.NoteTaskShapes = append(newDiagram.NoteTaskShapes, newShape)
		}

		for _, s := range diagram.NoteResourceShapes {
			newShape := s.GongCopy().(*NoteResourceShape)
			newShape.Note = s.Note
			newShape.Resource = s.Resource
			newDiagram.NoteResourceShapes = append(newDiagram.NoteResourceShapes, newShape)
		}

		for _, s := range diagram.Resource_Shapes {
			newShape := s.GongCopy().(*ResourceShape)
			newShape.Resource = s.Resource
			newDiagram.Resource_Shapes = append(newDiagram.Resource_Shapes, newShape)
		}

		for _, s := range diagram.ResourceComposition_Shapes {
			newShape := s.GongCopy().(*ResourceCompositionShape)
			newShape.Resource = s.Resource
			newDiagram.ResourceComposition_Shapes = append(newDiagram.ResourceComposition_Shapes, newShape)
		}

		for _, s := range diagram.ResourceTaskShapes {
			newShape := s.GongCopy().(*ResourceTaskShape)
			newShape.Resource = s.Resource
			newShape.Task = s.Task
			newDiagram.ResourceTaskShapes = append(newDiagram.ResourceTaskShapes, newShape)
		}

		StageBranch(stager.stage, newDiagram)
		stager.stage.Commit()
	}
}
