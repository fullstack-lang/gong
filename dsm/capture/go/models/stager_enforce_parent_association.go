package models

// enforceParentAssociation computes the parent field for Product, Task, and Resource
// based on their respective sub-elements.
// It first clears existing parent relationships, then rebuilds them
// by iterating through the composition lists in the stage.
func (stager *Stager) enforceParentAssociation() {
	stage := stager.stage

	// 1. Reset and compute for Products
	for _, product := range GetGongstrucsSorted[*Deliverable](stage) {
		product.parentProduct = nil
	}
	for _, product := range GetGongstrucsSorted[*Deliverable](stage) {
		for _, subProduct := range product.SubProducts {
			if subProduct != nil {
				subProduct.parentProduct = product
			}
		}
	}

	// 2. Reset and compute for Tasks
	for _, task := range GetGongstrucsSorted[*Concern](stage) {
		task.parentConcern = nil
	}
	for _, task := range GetGongstrucsSorted[*Concern](stage) {
		for _, subTask := range task.SubConcerns {
			if subTask != nil {
				subTask.parentConcern = task
			}
		}
	}

	// 3. Reset and compute for Resources
	for _, resource := range GetGongstrucsSorted[*Stakeholder](stage) {
		resource.parentStakeholder = nil
	}
	for _, resource := range GetGongstrucsSorted[*Stakeholder](stage) {
		for _, subResource := range resource.SubStakeholders {
			if subResource != nil {
				subResource.parentStakeholder = resource
			}
		}
	}
}
