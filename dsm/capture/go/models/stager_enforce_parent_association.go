package models

// enforceParentAssociation computes the parent field for Deliverable, Task, and Resource
// based on their respective sub-elements.
// It first clears existing parent relationships, then rebuilds them
// by iterating through the composition lists in the stage.
func (stager *Stager) enforceParentAssociation() {
	stage := stager.stage

	// 1. Reset and compute for Deliverables
	for _, deliverable := range stage.GetInstancesSorted[*Deliverable]() {
		deliverable.parentDeliverable = nil
	}
	for _, deliverable := range stage.GetInstancesSorted[*Deliverable]() {
		for _, subDeliverable := range deliverable.SubDeliverables {
			if subDeliverable != nil {
				subDeliverable.parentDeliverable = deliverable
			}
		}
	}

	// 2. Reset and compute for Tasks
	for _, task := range stage.GetInstancesSorted[*Concern]() {
		task.parentConcern = nil
	}
	for _, task := range stage.GetInstancesSorted[*Concern]() {
		for _, subTask := range task.SubConcerns {
			if subTask != nil {
				subTask.parentConcern = task
			}
		}
	}

	// 3. Reset and compute for Resources
	for _, resource := range stage.GetInstancesSorted[*Stakeholder]() {
		resource.parentStakeholder = nil
	}
	for _, resource := range stage.GetInstancesSorted[*Stakeholder]() {
		for _, subResource := range resource.SubStakeholders {
			if subResource != nil {
				subResource.parentStakeholder = resource
			}
		}
	}
}
