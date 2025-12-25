package models

func (stager *Stager) enforceTaskInputOutputProjectConsistency() (needCommit bool) {
	stage := stager.stage

	taskToProject := make(map[*Task]*Project)
	productToProject := make(map[*Product]*Project)

	for project := range stage.Projects {
		for _, task := range project.RootTasks {
			mapTaskToProject(task, project, taskToProject)
		}
		for _, product := range project.RootProducts {
			mapProductToProject(product, project, productToProject)
		}
	}

	for task := range stage.Tasks {

		project, ok := taskToProject[task]
		if !ok {
			continue
		}

		// check inputs
		// filter in place
		n := 0
		for _, input := range task.Inputs {
			if productToProject[input] == project {
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
			if productToProject[output] == project {
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

func mapTaskToProject(task *Task, project *Project, taskToProject map[*Task]*Project) {
	taskToProject[task] = project
	for _, subTask := range task.SubTasks {
		mapTaskToProject(subTask, project, taskToProject)
	}
}

func mapProductToProject(product *Product, project *Project, productToProject map[*Product]*Project) {
	productToProject[product] = project
	for _, subProduct := range product.SubProducts {
		mapProductToProject(subProduct, project, productToProject)
	}
}
