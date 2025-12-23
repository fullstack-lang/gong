package models

import "slices"

// enforceHierarchy checks that for every [models.Product] and [models.Task] present
// in a DAG attached to the root, they are unique (belong to one PBS/WBS only once max)
//
// # It also computes the parentProduct and parentTask fields
//
// if a Product/Task has to be removed from a list, returns true
func (stager *Stager) enforceHierarchy() (needCommit bool) {
	root := stager.root

	// 1. Reset all parent pointers
	// not necessary as we will traverse the whole graph

	// 2. Traverse the graph and check for uniqueness
	visitedProducts := make(map[*Product]struct{})
	visitedTasks := make(map[*Task]struct{})

	// helper for products
	var verifyProductHierarchy func(products *[]*Product, parent *Product)
	verifyProductHierarchy = func(products *[]*Product, parent *Product) {
		for i := len(*products) - 1; i >= 0; i-- {
			product := (*products)[i]

			if _, ok := visitedProducts[product]; ok {
				// already present in the hierarchy, remove it
				*products = slices.Delete(*products, i, i+1)
				needCommit = true
			} else {
				visitedProducts[product] = struct{}{}
				product.parentProduct = parent
				verifyProductHierarchy(&product.SubProducts, product)
			}
		}
	}

	// helper for tasks
	var verifyTaskHierarchy func(tasks *[]*Task, parent *Task)
	verifyTaskHierarchy = func(tasks *[]*Task, parent *Task) {
		for i := len(*tasks) - 1; i >= 0; i-- {
			task := (*tasks)[i]

			if _, ok := visitedTasks[task]; ok {
				// already present in the hierarchy, remove it
				*tasks = slices.Delete(*tasks, i, i+1)
				needCommit = true
			} else {
				visitedTasks[task] = struct{}{}
				task.parentTask = parent
				verifyTaskHierarchy(&task.SubTasks, task)
			}
		}
	}

	for _, project := range root.Projects {
		verifyProductHierarchy(&project.RootProducts, nil)
		verifyTaskHierarchy(&project.RootTasks, nil)
	}

	verifyProductHierarchy(&root.OrphanedProducts, nil)
	verifyTaskHierarchy(&root.OrphanedTasks, nil)

	return
}
