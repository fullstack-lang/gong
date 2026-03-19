package models

import "slices"

// enforceUniquenessInLibraries ensures that a Product or a Task is not referenced
// by more than one Library.
//
// The rule is "First Come, First Served": the first Library (in the order of root.Libraries)
// that claims a Product/Task keeps it. Subsequent Libraries lose the reference.
func (stager *Stager) enforceUniquenessInLibraries() (needCommit bool) {
	root := stager.root

	// map to remember which library has already claimed a product
	stager.productToLibrary = make(map[*Product]*Library)

	// map to remember which library has already claimed a task
	stager.taskToLibrary = make(map[*Task]*Library)

	for _, library := range root.Libraries {

		//
		// Products
		//
		var verifyProductUniqueness func(products *[]*Product)
		verifyProductUniqueness = func(products *[]*Product) {
			// traverse in reverse order to allow safe deletion
			for i := len(*products) - 1; i >= 0; i-- {
				product := (*products)[i]

				if owner, ok := stager.productToLibrary[product]; ok {
					if owner != library {
						// The product is already owned by another library
						// We remove the link from the current library
						*products = slices.Delete(*products, i, i+1)
						needCommit = true
					}
					// If owner == library, it means the product is referenced twice
					// within the same library. We do nothing here as this function
					// only enforces uniqueness *across* projects.
				} else {
					stager.productToLibrary[product] = library
					verifyProductUniqueness(&product.SubProducts)
				}
			}
		}

		verifyProductUniqueness(&library.RootProducts)

		//
		// Tasks
		//
		var verifyTaskUniqueness func(tasks *[]*Task)
		verifyTaskUniqueness = func(tasks *[]*Task) {
			// traverse in reverse order to allow safe deletion
			for i := len(*tasks) - 1; i >= 0; i-- {
				task := (*tasks)[i]

				if owner, ok := stager.taskToLibrary[task]; ok {
					if owner != library {
						// The task is already owned by another library
						// We remove the link from the current library
						*tasks = slices.Delete(*tasks, i, i+1)
						needCommit = true
					}
					// If owner == library, it means the task is referenced twice
					// within the same library. We do nothing here as this function
					// only enforces uniqueness *across* projects.
				} else {
					stager.taskToLibrary[task] = library
					verifyTaskUniqueness(&task.SubTasks)
				}
			}
		}

		verifyTaskUniqueness(&library.RootTasks)
	}

	return
}
