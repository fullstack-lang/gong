package models

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableProductShapes := make(map[*ProductShape]struct{})
	reachableProductCompositionShapes := make(map[*ProductCompositionShape]struct{})
	reachableTaskShapes := make(map[*TaskShape]struct{})
	reachableTaskCompositionShapes := make(map[*TaskCompositionShape]struct{})
	reachableTaskInputShapes := make(map[*TaskInputShape]struct{})
	reachableTaskOutputShapes := make(map[*TaskOutputShape]struct{})
	reachableNoteShapes := make(map[*NoteShape]struct{})
	reachableNoteProductShapes := make(map[*NoteProductShape]struct{})
	reachableNoteTaskShapes := make(map[*NoteTaskShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		collectShapes(diagram.Product_Shapes, reachableProductShapes)
		collectShapes(diagram.ProductComposition_Shapes, reachableProductCompositionShapes)
		collectShapes(diagram.Task_Shapes, reachableTaskShapes)
		collectShapes(diagram.TaskComposition_Shapes, reachableTaskCompositionShapes)
		collectShapes(diagram.TaskInputShapes, reachableTaskInputShapes)
		collectShapes(diagram.TaskOutputShapes, reachableTaskOutputShapes)
		collectShapes(diagram.Note_Shapes, reachableNoteShapes)
		collectShapes(diagram.NoteProductShapes, reachableNoteProductShapes)
		collectShapes(diagram.NoteTaskShapes, reachableNoteTaskShapes)
	}

	// 2. unstage shapes that are not attached to a diagram
	needCommit = unstageOrphans(stager.stage, reachableProductShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableProductCompositionShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableTaskShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableTaskCompositionShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableTaskInputShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableTaskOutputShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableNoteShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableNoteProductShapes) || needCommit
	needCommit = unstageOrphans(stager.stage, reachableNoteTaskShapes) || needCommit

	return
}

func collectShapes[T comparable](shapes []T, reachable map[T]struct{}) {
	for _, shape := range shapes {
		reachable[shape] = struct{}{}
	}
}

func unstageOrphans[T PointerToGongstruct](stage *Stage, reachable map[T]struct{}) (needCommit bool) {
	for _, shape := range GetGongstrucsSorted[T](stage) {
		if _, ok := reachable[shape]; !ok {
			shape.UnstageVoid(stage)
			needCommit = true
		}
	}
	return
}
