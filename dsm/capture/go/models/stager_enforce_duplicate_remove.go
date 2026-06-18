package models

import "time"

// enforceDuplicateRemove parses all associations between AbstractType
// and deletes duplicates within these associations
func (stager *Stager) enforceDuplicateRemove() (needCommit bool) {
	stage := stager.stage

	needCommit = removeDuplicatesSlice(stager, &stager.GetRootLibrary().SubLibraries) || needCommit

	for diagram := range *GetGongstructInstancesSetFromPointerType[*Diagram](stage) {
		needCommit = removeDuplicatesSlice(stager, &diagram.Product_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ProductComposition_Shapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagram.Requirement_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.Concept_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.DeliverableConceptShapes) || needCommit

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

	for deliverable := range *GetGongstructInstancesSetFromPointerType[*Deliverable](stage) {
		needCommit = removeDuplicatesSlice(stager, &deliverable.SubProducts) || needCommit
		needCommit = removeDuplicatesSlice(stager, &deliverable.Concepts) || needCommit
	}

	for concern := range *GetGongstructInstancesSetFromPointerType[*Concern](stage) {
		needCommit = removeDuplicatesSlice(stager, &concern.SubConcerns) || needCommit
		needCommit = removeDuplicatesSlice(stager, &concern.Inputs) || needCommit
		needCommit = removeDuplicatesSlice(stager, &concern.Outputs) || needCommit
	}

	for note := range *GetGongstructInstancesSetFromPointerType[*Note](stage) {
		needCommit = removeDuplicatesSlice(stager, &note.Products) || needCommit
		needCommit = removeDuplicatesSlice(stager, &note.Tasks) || needCommit
	}

	for stakeholder := range *GetGongstructInstancesSetFromPointerType[*Stakeholder](stage) {
		needCommit = removeDuplicatesSlice(stager, &stakeholder.SubStakeholders) || needCommit
		needCommit = removeDuplicatesSlice(stager, &stakeholder.Concerns) || needCommit
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
