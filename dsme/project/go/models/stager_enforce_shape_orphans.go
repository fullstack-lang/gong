package models

func (stager *Stager) enforceShapeOrphans() (needCommit bool) {
	// 1. collect all shapes that are attached to a diagram
	reachableProductShapes := make(map[*ProductShape]struct{})
	reachableCompositionShapes := make(map[*CompositionShape]struct{})
	reachableTaskShapes := make(map[*TaskShape]struct{})

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		for _, shape := range diagram.Product_Shapes {
			reachableProductShapes[shape] = struct{}{}
		}
		for _, shape := range diagram.Composition_Shapes {
			reachableCompositionShapes[shape] = struct{}{}
		}
		for _, shape := range diagram.Task_Shapes {
			reachableTaskShapes[shape] = struct{}{}
		}
	}

	// 2. unstage shapes that are not attached to a diagram
	for _, shape := range GetGongstrucsSorted[*ProductShape](stager.stage) {
		if _, ok := reachableProductShapes[shape]; !ok {
			shape.Unstage(stager.stage)
			needCommit = true
		}
	}

	for _, shape := range GetGongstrucsSorted[*CompositionShape](stager.stage) {
		if _, ok := reachableCompositionShapes[shape]; !ok {
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

	return
}
