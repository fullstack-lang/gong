// new file
package models

import "slices"

// enforceTaskCompositionShapes iterates through all diagrams and removes any
// TaskCompositionShape whose associated Task or parent Task is not present
// in the diagram's Task_Shapes. This ensures that links are only displayed
// when both connected nodes are visible.
func (stager *Stager) enforceTaskCompositionShapes() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		// map of tasks present in the diagram
		set_Task := make(map[*Task]struct{})
		for _, taskShape := range diagram.Task_Shapes {
			if taskShape.Task != nil {
				set_Task[taskShape.Task] = struct{}{}
			}
		}

		// check that the composition shape is valid
		for i := len(diagram.TaskComposition_Shapes) - 1; i >= 0; i-- {
			compositionShape := diagram.TaskComposition_Shapes[i]
			task := compositionShape.Task

			// check if the task is present in the diagram
			// if not, remove the composition shape
			if _, ok := set_Task[task]; !ok {
				diagram.TaskComposition_Shapes = slices.Delete(diagram.TaskComposition_Shapes, i, i+1)
				compositionShape.Unstage(stager.stage)
				needCommit = true
				continue
			}

			// check if the parent task is present in the diagram
			// if not, remove the composition shape
			parentTask := task.parentTask
			if _, ok := set_Task[parentTask]; !ok {
				diagram.TaskComposition_Shapes = slices.Delete(diagram.TaskComposition_Shapes, i, i+1)
				compositionShape.Unstage(stager.stage)
				needCommit = true
				continue
			}
		}
	}
	return
}
