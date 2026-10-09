package models

import "slices"

// enforceTaskGroupTasks updates the TaskGroup.Tasks association to match
// the Task.TaskGroups association, while preserving the existing order of
// tasks within TaskGroup.Tasks for display order management.
func (stager *Stager) enforceTaskGroupTasks() (needCommit bool) {
	stage := stager.stage

	// 1. Backward compatibility: if a task was defined via taskGroup.Tasks
	// and has task.TaskGroups == nil (uninitialized), initialize task.TaskGroups
	// from the taskGroup.
	for _, taskGroup := range stage.GetInstancesSorted[*TaskGroup]() {
		for _, task := range taskGroup.Tasks {
			if task != nil && task.TaskGroups == nil {
				task.TaskGroups = append(task.TaskGroups, taskGroup)
				needCommit = true
			}
		}
	}

	// 2. Map each taskGroup to the set of tasks that reference it via task.TaskGroups
	taskGroupExpectedTasks := make(map[*TaskGroup]map[*Task]struct{})
	for _, taskGroup := range stage.GetInstancesSorted[*TaskGroup]() {
		taskGroupExpectedTasks[taskGroup] = make(map[*Task]struct{})
	}

	for _, task := range stage.GetInstancesSorted[*Task]() {
		for _, taskGroup := range task.TaskGroups {
			if taskGroup == nil {
				continue
			}
			if expectedMap, ok := taskGroupExpectedTasks[taskGroup]; ok {
				expectedMap[task] = struct{}{}
			}
		}
	}

	// 3. For each taskGroup, update taskGroup.Tasks to match expectedTasks
	// while preserving existing manual display ordering.
	for _, taskGroup := range stage.GetInstancesSorted[*TaskGroup]() {
		expectedTasks := taskGroupExpectedTasks[taskGroup]

		seen := make(map[*Task]struct{})
		var newTasks []*Task

		// Keep existing tasks in taskGroup.Tasks that are still expected,
		// preserving their manual display order and filtering out duplicates/nil.
		for _, task := range taskGroup.Tasks {
			if task == nil {
				continue
			}
			if _, expected := expectedTasks[task]; expected {
				if _, alreadySeen := seen[task]; !alreadySeen {
					seen[task] = struct{}{}
					newTasks = append(newTasks, task)
				}
			}
		}

		// Append any newly associated tasks that were not previously in taskGroup.Tasks
		for _, task := range stage.GetInstancesSorted[*Task]() {
			if _, expected := expectedTasks[task]; expected {
				if _, alreadySeen := seen[task]; !alreadySeen {
					seen[task] = struct{}{}
					newTasks = append(newTasks, task)
				}
			}
		}

		// Update taskGroup.Tasks if slice content or order changed
		if !slices.Equal(taskGroup.Tasks, newTasks) {
			taskGroup.Tasks = newTasks
			needCommit = true
		}
	}

	return
}
