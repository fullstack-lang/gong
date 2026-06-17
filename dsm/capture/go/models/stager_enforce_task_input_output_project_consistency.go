package models

// enforceTaskInputOutputLibraryConsistency ensures that the input and output
// products of a task belong to the same library as the task itself.
func (stager *Stager) enforceTaskInputOutputLibraryConsistency() (needCommit bool) {
	stage := stager.stage

	taskToLibrary := make(map[*Concern]*Library)
	productToLibrary := make(map[*Deliverable]*Library)

	for library := range *GetGongstructInstancesSetFromPointerType[*Library](stage) {
		for _, task := range library.RootConcerns {
			mapTaskToLibrary(task, library, taskToLibrary)
		}
		for _, product := range library.RootDeliverables {
			mapProductToLibrary(product, library, productToLibrary)
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
			if productToLibrary[input] == library {
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
			if productToLibrary[output] == library {
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

func mapProductToLibrary(product *Deliverable, library *Library, productToLibrary map[*Deliverable]*Library) {
	productToLibrary[product] = library
	for _, subProduct := range product.SubProducts {
		mapProductToLibrary(subProduct, library, productToLibrary)
	}
}
