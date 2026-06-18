package models

// enforceTaskInputOutputLibraryConsistency ensures that the input and output
// deliverables of a task belong to the same library as the task itself.
func (stager *Stager) enforceTaskInputOutputLibraryConsistency() (needCommit bool) {
	stage := stager.stage

	taskToLibrary := make(map[*Concern]*Library)
	deliverableToLibrary := make(map[*Deliverable]*Library)

	for library := range *GetGongstructInstancesSetFromPointerType[*Library](stage) {
		for _, task := range library.RootConcerns {
			mapTaskToLibrary(task, library, taskToLibrary)
		}
		for _, deliverable := range library.RootDeliverables {
			mapDeliverableToLibrary(deliverable, library, deliverableToLibrary)
		}
	}

	for task := range *GetGongstructInstancesSetFromPointerType[*Concern](stage) {

		library, ok := taskToLibrary[task]
		if !ok {
			continue
		}

		// check inputs
		// filter in place
		n := 0
		for _, input := range task.Inputs {
			if deliverableToLibrary[input] == library {
				task.Inputs[n] = input
				n++
			} else {
				needCommit = true
			}
		}
		task.Inputs = task.Inputs[:n]

		// check outputs
		n = 0
		for _, output := range task.Outputs {
			if deliverableToLibrary[output] == library {
				task.Outputs[n] = output
				n++
			} else {
				needCommit = true
			}
		}
		task.Outputs = task.Outputs[:n]
	}

	return
}

func mapTaskToLibrary(task *Concern, library *Library, taskToLibrary map[*Concern]*Library) {
	taskToLibrary[task] = library
	for _, subTask := range task.SubConcerns {
		mapTaskToLibrary(subTask, library, taskToLibrary)
	}
}

func mapDeliverableToLibrary(deliverable *Deliverable, library *Library, deliverableToLibrary map[*Deliverable]*Library) {
	deliverableToLibrary[deliverable] = library
	for _, subDeliverable := range deliverable.SubDeliverables {
		mapDeliverableToLibrary(subDeliverable, library, deliverableToLibrary)
	}
}
