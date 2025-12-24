package models

import "slices"

func (stager *Stager) enforceTaskInputOutputShapes() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		// map of tasks present in the diagram
		set_Task := make(map[*Task]struct{})
		for _, taskShape := range diagram.Task_Shapes {
			if taskShape.Task != nil {
				set_Task[taskShape.Task] = struct{}{}
			}
		}
		// map of products present in the diagram
		set_Product := make(map[*Product]struct{})
		for _, productShape := range diagram.Product_Shapes {
			if productShape.Product != nil {
				set_Product[productShape.Product] = struct{}{}
			}
		}

		// check that the composition shape is valid
		for i := len(diagram.TaskInputShapes) - 1; i >= 0; i-- {
			taskInputShape := diagram.TaskInputShapes[i]
			task := taskInputShape.Task
			product := taskInputShape.Product

			// check if the task is present in the diagram
			// if not, remove the composition shape
			if _, ok := set_Task[task]; !ok {
				diagram.TaskInputShapes = slices.Delete(diagram.TaskInputShapes, i, i+1)
				taskInputShape.Unstage(stager.stage)
				needCommit = true
				continue
			}

			// check if the product is present in the diagram
			// if not, remove the composition shape
			if _, ok := set_Product[product]; !ok {
				diagram.TaskInputShapes = slices.Delete(diagram.TaskInputShapes, i, i+1)
				taskInputShape.Unstage(stager.stage)
				needCommit = true
				continue
			}
		}

		for i := len(diagram.TaskOutputShapes) - 1; i >= 0; i-- {
			taskOutputShape := diagram.TaskOutputShapes[i]
			task := taskOutputShape.Task
			product := taskOutputShape.Product

			// check if the task is present in the diagram
			// if not, remove the composition shape
			if _, ok := set_Task[task]; !ok {
				diagram.TaskOutputShapes = slices.Delete(diagram.TaskOutputShapes, i, i+1)
				taskOutputShape.Unstage(stager.stage)
				needCommit = true
				continue
			}

			// check if the product is present in the diagram
			// if not, remove the composition shape
			if _, ok := set_Product[product]; !ok {
				diagram.TaskOutputShapes = slices.Delete(diagram.TaskOutputShapes, i, i+1)
				taskOutputShape.Unstage(stager.stage)
				needCommit = true
				continue
			}
		}
	}
	return
}
