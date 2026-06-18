package models

import (
)

func onCopyDiagram(stager *Stager, diagram *Diagram) func() {
	return func() {
		newDiagram := new(Diagram)
		newDiagram.Name = diagram.Name + " copy"
		newDiagram.IsEditable_ = true

		library := stager.stage.Library_Diagrams_reverseMap[diagram]
		Append(&library.Diagrams, newDiagram)

		for _, s := range diagram.Deliverable_Shapes {
			newShape := s.GongCopy().(*DeliverableShape)
			newDiagram.Deliverable_Shapes = append(newDiagram.Deliverable_Shapes, newShape)
		}

		for _, s := range diagram.DeliverableComposition_Shapes {
			newShape := s.GongCopy().(*DeliverableCompositionShape)
			newDiagram.DeliverableComposition_Shapes = append(newDiagram.DeliverableComposition_Shapes, newShape)
		}

		for _, s := range diagram.Concern_Shapes {
			newShape := s.GongCopy().(*ConcernShape)
			newDiagram.Concern_Shapes = append(newDiagram.Concern_Shapes, newShape)
		}

		for _, s := range diagram.ConcernComposition_Shapes {
			newShape := s.GongCopy().(*ConcernCompositionShape)
			newDiagram.ConcernComposition_Shapes = append(newDiagram.ConcernComposition_Shapes, newShape)
		}

		for _, s := range diagram.ConcernInputShapes {
			newShape := s.GongCopy().(*ConcernInputShape)
			newDiagram.ConcernInputShapes = append(newDiagram.ConcernInputShapes, newShape)
		}

		for _, s := range diagram.ConcernOutputShapes {
			newShape := s.GongCopy().(*ConcernOutputShape)
			newDiagram.ConcernOutputShapes = append(newDiagram.ConcernOutputShapes, newShape)
		}

		for _, s := range diagram.Note_Shapes {
			newShape := s.GongCopy().(*NoteShape)
			newDiagram.Note_Shapes = append(newDiagram.Note_Shapes, newShape)
		}

		for _, s := range diagram.NoteDeliverableShapes {
			newShape := s.GongCopy().(*NoteDeliverableShape)
			newDiagram.NoteDeliverableShapes = append(newDiagram.NoteDeliverableShapes, newShape)
		}

		for _, s := range diagram.NoteTaskShapes {
			newShape := s.GongCopy().(*NoteTaskShape)
			newDiagram.NoteTaskShapes = append(newDiagram.NoteTaskShapes, newShape)
		}

		StageBranch(stager.stage, newDiagram)
		stager.stage.Commit()
	}
}
