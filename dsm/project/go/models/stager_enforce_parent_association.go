package models

// enforceParentAssociation computes the parent field for Product, Task, and Resource
// based on their respective sub-elements.
// It first clears existing parent relationships, then rebuilds them
// by iterating through the composition lists in the stage.
func (stager *Stager) enforceParentAssociation() {
	stage := stager.stage

	// 1. Reset and compute for Products
	for _, product := range stage.GetInstancesSorted[*Product]() {
		product.parentProduct = nil
	}
	for _, product := range stage.GetInstancesSorted[*Product]() {
		for _, subProduct := range product.SubProducts {
			if subProduct != nil {
				subProduct.parentProduct = product
			}
		}
	}

	// 2. Reset and compute for Tasks
	for _, task := range stage.GetInstancesSorted[*Task]() {
		task.parentTask = nil
	}
	for _, task := range stage.GetInstancesSorted[*Task]() {
		for _, subTask := range task.SubTasks {
			if subTask != nil {
				subTask.parentTask = task
			}
		}
	}

	// 3. Reset and compute for Resources
	for _, resource := range stage.GetInstancesSorted[*Resource]() {
		resource.parentResource = nil
	}
	for _, resource := range stage.GetInstancesSorted[*Resource]() {
		for _, subResource := range resource.SubResources {
			if subResource != nil {
				subResource.parentResource = resource
			}
		}
	}
}
