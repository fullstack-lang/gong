package models

// enforceParentAssociation computes the parent field for Product, Task, and Resource
// based on their respective sub-elements.
// It first clears existing parent relationships, then rebuilds them
// by iterating through the composition lists in the stage.
func (stager *Stager) enforceParentAssociation() {
	stage := stager.stage

	// 1. Reset and compute for Products
	for _, product := range GetGongstrucsSorted[*Product](stage) {
		product.parentProduct = nil
	}
	for _, product := range GetGongstrucsSorted[*Product](stage) {
		for _, subProduct := range product.SubProducts {
			if subProduct != nil {
				subProduct.parentProduct = product
			}
		}
	}

	// 2. Reset and compute for Tasks
	for _, task := range GetGongstrucsSorted[*Task](stage) {
		task.parentTask = nil
	}
	for _, task := range GetGongstrucsSorted[*Task](stage) {
		for _, subTask := range task.SubTasks {
			if subTask != nil {
				subTask.parentTask = task
			}
		}
	}

	// 3. Reset and compute for Resources
	for _, resource := range GetGongstrucsSorted[*Resource](stage) {
		resource.parentResource = nil
	}
	for _, resource := range GetGongstrucsSorted[*Resource](stage) {
		for _, subResource := range resource.SubResources {
			if subResource != nil {
				subResource.parentResource = resource
			}
		}
	}
}
