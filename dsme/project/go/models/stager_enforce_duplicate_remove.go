package models

import "time"

// enforceDuplicateRemove parses all associations between AbstractType
// and deletes duplicates within these associations
func (stager *Stager) enforceDuplicateRemove() (needCommit bool) {
	stage := stager.stage

	if removeDuplicatesSlice(stager, &stager.root.Projects) {
		needCommit = true
	}

	for diagram := range stage.Diagrams {
		if removeDuplicatesSlice(stager, &diagram.Product_Shapes) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &diagram.Task_Shapes) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &diagram.Note_Shapes) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &diagram.Resource_Shapes) {
			needCommit = true
		}

		if removeDuplicatesSlice(stager, &diagram.TaskComposition_Shapes) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &diagram.TaskInputShapes) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &diagram.TaskOutputShapes) {
			needCommit = true
		}
	}

	for project := range stage.Projects {
		if removeDuplicatesSlice(stager, &project.RootProducts) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &project.RootTasks) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &project.Notes) {
			needCommit = true
		}
	}

	for product := range stage.Products {
		if removeDuplicatesSlice(stager, &product.SubProducts) {
			needCommit = true
		}
	}

	for task := range stage.Tasks {
		if removeDuplicatesSlice(stager, &task.SubTasks) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &task.Inputs) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &task.Outputs) {
			needCommit = true
		}
	}

	for note := range stage.Notes {
		if removeDuplicatesSlice(stager, &note.Products) {
			needCommit = true
		}
		if removeDuplicatesSlice(stager, &note.Tasks) {
			needCommit = true
		}
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
