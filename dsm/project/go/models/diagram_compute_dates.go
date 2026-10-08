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

			taskEnd := task.End
			if task.IsAllDay && !task.IsMilestone {
				taskEnd = taskEnd.AddDate(0, 0, 1)
			}

			if firstTask {
				diagram.ComputedStart = task.Start
				diagram.ComputedEnd = taskEnd
				firstTask = false
			} else {
				if diagram.ComputedStart.After(task.Start) {
					diagram.ComputedStart = task.Start
				}
				if diagram.ComputedEnd.Before(taskEnd) {
					diagram.ComputedEnd = taskEnd
				}
			}
		}
	}

	if diagram.UseManualStartAndEndDates {
		diagram.ComputedStart = diagram.ManualStart
		diagram.ComputedEnd = diagram.ManualEnd
	}

	// align start on the beginning of the first time scale
	if diagram.AlignOnBeginningOfTimeScale {
		start := diagram.ComputedStart
		loc := start.Location()
		timeStepScale := diagram.TimeStepScale
		if timeStepScale == "" || timeStepScale == NONE {
			timeStepScale = MONTHS
		}
		switch timeStepScale {
		case YEARS:
			diagram.ComputedStart = time.Date(start.Year(), time.January, 1, 0, 0, 0, 0, loc)
		case MONTHS:
			diagram.ComputedStart = time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, loc)
		case WEEKS:
			weekday := int(start.Weekday())
			if weekday == 0 {
				weekday = 7
			}
			diagram.ComputedStart = time.Date(start.Year(), start.Month(), start.Day()-(weekday-1), 0, 0, 0, 0, loc)
		case DAYS:
			diagram.ComputedStart = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, loc)
		default:
			diagram.ComputedStart = time.Date(start.Year(), time.January, 1, 0, 0, 0, 0, loc)
		}
	}

	diagram.ComputedDuration = diagram.ComputedEnd.Sub(diagram.ComputedStart)
}
