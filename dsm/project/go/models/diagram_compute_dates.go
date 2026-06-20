package models

import (
	"log"
	"time"
)

func (diagram *Diagram) computeStartAndEndDate() {
	firstTask := true
	for _, taskGroupShape := range diagram.TaskGroupShapes {
		taskGroup := taskGroupShape.TaskGroup
		if taskGroup == nil {
			log.Panic("TaskGroupShape has a no TaskGroup", taskGroupShape.Name)
			continue
		}
		for _, task := range taskGroup.Tasks {
			_, ok := diagram.map_Task_TaskShape[task]
			if !ok {
				continue
			}

			if firstTask {
				diagram.ComputedStart = task.Start
				diagram.ComputedEnd = task.End
				firstTask = false
			} else {
				if diagram.ComputedStart.After(task.Start) {
					diagram.ComputedStart = task.Start
				}
				if diagram.ComputedEnd.Before(task.End) {
					diagram.ComputedEnd = task.End
				}
			}
		}
	}

	if diagram.UseManualStartAndEndDates {
		diagram.ComputedStart = diagram.ManualStart
		diagram.ComputedEnd = diagram.ManualEnd
	}

	// align start on the beginning of the year
	if diagram.AlignOnStartEndOnYearStart {
		diagram.ComputedStart = time.Date(diagram.ComputedStart.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
		diagram.ComputedEnd = time.Date(diagram.ComputedEnd.Year(), time.December, 31, 0, 0, 0, 0, time.UTC)
	}

	diagram.ComputedDuration = diagram.ComputedEnd.Sub(diagram.ComputedStart)
}
