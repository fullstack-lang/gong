package models

import "slices"

// enforceUniqueProductTaskInProjects ensures that a Product or a Task is not referenced
// by more than one Project.
//
// The rule is "First Come, First Served": the first Project (in the order of root.Projects)
// that claims a Product/Task keeps it. Subsequent Projects lose the reference.
func (stager *Stager) enforceUniqueProductTaskInProjects() (needCommit bool) {

	root := stager.root

	// map to remember which project has already claimed a product
	productToProject := make(map[*Product]*Project)

	// map to remember which project has already claimed a task
	taskToProject := make(map[*Task]*Project)

	for _, project := range root.Projects {

		//
		// Products
		//
		var verifyProductUniqueness func(products *[]*Product)
		verifyProductUniqueness = func(products *[]*Product) {

			// traverse in reverse order to allow safe deletion
			for i := len(*products) - 1; i >= 0; i-- {
				product := (*products)[i]

				if owner, ok := productToProject[product]; ok {
					if owner != project {
						// The product is already owned by another project
						// We remove the link from the current project
						*products = slices.Delete(*products, i, i+1)
						needCommit = true
					}
					// If owner == project, it means the product is referenced twice
					// within the same project. We do nothing here as this function
					// only enforces uniqueness *across* projects.
				} else {
					productToProject[product] = project
					verifyProductUniqueness(&product.SubProducts)
				}
			}
		}

		verifyProductUniqueness(&project.RootProducts)

		//
		// Tasks
		//
		var verifyTaskUniqueness func(tasks *[]*Task)
		verifyTaskUniqueness = func(tasks *[]*Task) {

			// traverse in reverse order to allow safe deletion
			for i := len(*tasks) - 1; i >= 0; i-- {
				task := (*tasks)[i]

				if owner, ok := taskToProject[task]; ok {
					if owner != project {
						// The task is already owned by another project
						// We remove the link from the current project
						*tasks = slices.Delete(*tasks, i, i+1)
						needCommit = true
					}
					// If owner == project, it means the task is referenced twice
					// within the same project. We do nothing here as this function
					// only enforces uniqueness *across* projects.
				} else {
					taskToProject[task] = project
					verifyTaskUniqueness(&task.SubTasks)
				}
			}
		}

		verifyTaskUniqueness(&project.RootTasks)
	}

	return
}
