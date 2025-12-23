package models

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableProductShapes := make(map[*ProductShape]struct{})
	reachableProductCompositionShapes := make(map[*ProductCompositionShape]struct{})
	reachableTaskShapes := make(map[*TaskShape]struct{})
	reachableTaskCompositionShapes := make(map[*TaskCompositionShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		for _, shape := range diagram.Product_Shapes {
			reachableProductShapes[shape] = struct{}{}
		}
		for _, shape := range diagram.ProductComposition_Shapes {
			reachableProductCompositionShapes[shape] = struct{}{}
		}
		for _, shape := range diagram.Task_Shapes {
			reachableTaskShapes[shape] = struct{}{}
		}
		for _, shape := range diagram.TaskComposition_Shapes {
			reachableTaskCompositionShapes[shape] = struct{}{}
		}
	}

	// 2. unstage shapes that are not attached to a diagram
	for _, shape := range GetGongstrucsSorted[*ProductShape](stager.stage) {
		if _, ok := reachableProductShapes[shape]; !ok {
			shape.Unstage(stager.stage)
			needCommit = true
		}
	}

	for _, shape := range GetGongstrucsSorted[*ProductCompositionShape](stager.stage) {
		if _, ok := reachableProductCompositionShapes[shape]; !ok {
			shape.Unstage(stager.stage)
			needCommit = true
		}
	}

	for _, shape := range GetGongstrucsSorted[*TaskShape](stager.stage) {
		if _, ok := reachableTaskShapes[shape]; !ok {
			shape.Unstage(stager.stage)
			needCommit = true
		}
	}

	for _, shape := range GetGongstrucsSorted[*TaskCompositionShape](stager.stage) {
		if _, ok := reachableTaskCompositionShapes[shape]; !ok {
			shape.Unstage(stager.stage)
			needCommit = true
		}
	}

	return
}
