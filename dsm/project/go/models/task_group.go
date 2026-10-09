package models

import "slices"

// GetTasks returns all tasks belonging to this TaskGroup,
// taking into account both taskGroup.Tasks and task.TaskGroups.
func (taskGroup *TaskGroup) GetTasks(stage *Stage) []*Task {
	if taskGroup == nil {
		return nil
	}
	var tasks []*Task
	taskSet := make(map[*Task]struct{})

	for _, task := range taskGroup.Tasks {
		if task == nil {
			continue
		}
		if _, exists := taskSet[task]; !exists {
			taskSet[task] = struct{}{}
			tasks = append(tasks, task)
		}
	}

	if stage != nil {
		for _, task := range stage.GetInstancesSorted[*Task]() {
			if _, exists := taskSet[task]; !exists {
				if slices.Contains(task.TaskGroups, taskGroup) {
					taskSet[task] = struct{}{}
					tasks = append(tasks, task)
				}
			}
		}
	}

	return tasks
}

func (stager *Stager) getTasksOfTaskGroup(taskGroup *TaskGroup) []*Task {
	if stager == nil {
		return taskGroup.GetTasks(nil)
	}
	return taskGroup.GetTasks(stager.stage)
}
