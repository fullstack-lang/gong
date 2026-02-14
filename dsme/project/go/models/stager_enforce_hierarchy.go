package models

import (
	"fmt"
	"slices"
	"time"
)

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

	// helper for products
	visitedProducts := make(map[*Product]struct{})
	var verifyProductHierarchy func(products *[]*Product, parent *Product)
	verifyProductHierarchy = func(products *[]*Product, parent *Product) {
		for i := len(*products) - 1; i >= 0; i-- {
			product := (*products)[i]

			if _, ok := visitedProducts[product]; ok {
				// already present in the hierarchy, remove it
				*products = slices.Delete(*products, i, i+1)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Product %s is already present in the hierarchy, removing it", product.Name))
				needCommit = true
			} else {
				visitedProducts[product] = struct{}{}
				product.parentProduct = parent
				verifyProductHierarchy(&product.SubProducts, product)
			}
		}
	}

	// helper for tasks
	visitedTasks := make(map[*Task]struct{})
	var verifyTaskHierarchy func(tasks *[]*Task, parent *Task)
	verifyTaskHierarchy = func(tasks *[]*Task, parent *Task) {
		for i := len(*tasks) - 1; i >= 0; i-- {
			task := (*tasks)[i]

			if _, ok := visitedTasks[task]; ok {
				// already present in the hierarchy, remove it
				*tasks = slices.Delete(*tasks, i, i+1)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Task %s is already present in the hierarchy, removing it", task.Name))
				needCommit = true
			} else {
				visitedTasks[task] = struct{}{}
				task.parentTask = parent
				verifyTaskHierarchy(&task.SubTasks, task)
			}
		}
	}

	visitedResources := make(map[*Resource]struct{})
	var verifyResourceHierarchy func(resources *[]*Resource, parent *Resource)
	verifyResourceHierarchy = func(resources *[]*Resource, parent *Resource) {
		for i := len(*resources) - 1; i >= 0; i-- {
			resource := (*resources)[i]

			if resource == nil {
				// remove nil resource
				*resources = slices.Delete(*resources, i, i+1)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Nil resource found in the hierarchy, removing it"))
				needCommit = true
			} else if _, ok := visitedResources[resource]; ok {
				// already present in the hierarchy, remove it
				*resources = slices.Delete(*resources, i, i+1)
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Resource %s is already present in the hierarchy, removing it", resource.Name))
				needCommit = true
			} else {
				visitedResources[resource] = struct{}{}
				resource.parentResource = parent
				verifyResourceHierarchy(&resource.SubResources, resource)
			}
		}
	}

	for _, project := range root.Projects {
		verifyProductHierarchy(&project.RootProducts, nil)
		verifyTaskHierarchy(&project.RootTasks, nil)
	}

	return
}
