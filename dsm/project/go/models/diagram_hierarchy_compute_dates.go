package models

import "time"

func (diagramHierarchy *DiagramHierarchy) ComputeStartAndEndDate() {
	firstTask := true
	for _, taskGroupShape := range diagramHierarchy.TaskGroupShapes {
		taskGroup := taskGroupShape.TaskGroup
		for _, task := range taskGroup.Tasks {
			_, ok := diagramHierarchy.map_Task_TaskShape[task]
			if !ok {
				continue
			}

			if firstTask {
				diagramHierarchy.ComputedStart = task.Start
				diagramHierarchy.ComputedEnd = task.End
				firstTask = false
			} else {
				if diagramHierarchy.ComputedStart.After(task.Start) {
					diagramHierarchy.ComputedStart = task.Start
				}
				if diagramHierarchy.ComputedEnd.Before(task.End) {
					diagramHierarchy.ComputedEnd = task.End
				}
			}
		}
	}

	if diagramHierarchy.UseManualStartAndEndDates {
		diagramHierarchy.ComputedStart = diagramHierarchy.ManualStart
		diagramHierarchy.ComputedEnd = diagramHierarchy.ManualEnd
	}

	// align start on the beginning of the year
	if diagramHierarchy.AlignOnStartEndOnYearStart {
		diagramHierarchy.ComputedStart = time.Date(diagramHierarchy.ComputedStart.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		diagramHierarchy.ComputedEnd = time.Date(diagramHierarchy.ComputedEnd.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	}

	diagramHierarchy.ComputedDuration = diagramHierarchy.ComputedEnd.Sub(diagramHierarchy.ComputedStart)
}
