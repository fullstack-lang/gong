package models

import "time"

// enforceDuplicateRemove parses all associations between AbstractType
// and deletes duplicates within these associations
func (stager *Stager) enforceDuplicateRemove() (needCommit bool) {
	stage := stager.stage

	needCommit = removeDuplicatesSlice(stager, &stager.rootLibrary.SubLibraries) || needCommit

	for diagram := range *GetGongstructInstancesSetFromPointerType[*Diagram](stage) {
		needCommit = removeDuplicatesSlice(stager, &diagram.Product_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ProductComposition_Shapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagram.Concern_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ConcernComposition_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ConcernInputShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ConcernOutputShapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagram.Note_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.NoteProductShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.NoteTaskShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.NoteResourceShapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagram.Stakeholder_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ResourceComposition_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.StakeholderConcernShapes) || needCommit
	}

	for library := range *GetGongstructInstancesSetFromPointerType[*Library](stage) {
		needCommit = removeDuplicatesSlice(stager, &library.RootDeliverables) || needCommit
		needCommit = removeDuplicatesSlice(stager, &library.RootConcerns) || needCommit
		needCommit = removeDuplicatesSlice(stager, &library.RootStakeholders) || needCommit
		needCommit = removeDuplicatesSlice(stager, &library.Notes) || needCommit
	}

	for product := range *GetGongstructInstancesSetFromPointerType[*Deliverable](stage) {
		needCommit = removeDuplicatesSlice(stager, &product.SubProducts) || needCommit
	}

	for task := range *GetGongstructInstancesSetFromPointerType[*Concern](stage) {
		needCommit = removeDuplicatesSlice(stager, &task.SubConcerns) || needCommit
		needCommit = removeDuplicatesSlice(stager, &task.Inputs) || needCommit
		needCommit = removeDuplicatesSlice(stager, &task.Outputs) || needCommit
	}

	for note := range *GetGongstructInstancesSetFromPointerType[*Note](stage) {
		needCommit = removeDuplicatesSlice(stager, &note.Products) || needCommit
		needCommit = removeDuplicatesSlice(stager, &note.Tasks) || needCommit
	}

	for resource := range *GetGongstructInstancesSetFromPointerType[*Stakeholder](stage) {
		needCommit = removeDuplicatesSlice(stager, &resource.SubStakeholders) || needCommit
		needCommit = removeDuplicatesSlice(stager, &resource.Concerns) || needCommit
	}

	return
}

// removeDuplicatesSlice is a generic helper that removes duplicate pointers from a slice.
// It keeps the first occurrence of each element and preserves order.
// It returns true if any duplicates were removed.
func removeDuplicatesSlice[T PointerToGongstruct](stager *Stager, slicePtr *[]T) bool {
	keys := make(map[T]bool)
	list := []T{}
	changed := false
	for _, entry := range *slicePtr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		} else {
			changed = true
			stager.probeForm.AddNotification(time.Now(), "Removed duplicate entry from slice: "+entry.GetName())
		}
	}
	if changed {
		*slicePtr = list
	}
	return changed
}
