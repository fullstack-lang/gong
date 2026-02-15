package models

import "time"

// enforceDuplicateRemove parses all associations between AbstractType
// and deletes duplicates within these associations
func (stager *Stager) enforceDuplicateRemove() (needCommit bool) {
	stage := stager.stage

	needCommit = removeDuplicatesSlice(stager, &stager.root.Projects) || needCommit

	for diagram := range stage.Diagrams {
		needCommit = removeDuplicatesSlice(stager, &diagram.Product_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ProductComposition_Shapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagram.Task_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.TaskComposition_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.TaskInputShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.TaskOutputShapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagram.Note_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.NoteProductShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.NoteTaskShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.NoteResourceShapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagram.Resource_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ResourceComposition_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagram.ResourceTaskShapes) || needCommit
	}

	for project := range stage.Projects {
		needCommit = removeDuplicatesSlice(stager, &project.RootProducts) || needCommit
		needCommit = removeDuplicatesSlice(stager, &project.RootTasks) || needCommit
		needCommit = removeDuplicatesSlice(stager, &project.RootResources) || needCommit
		needCommit = removeDuplicatesSlice(stager, &project.Notes) || needCommit
	}

	for product := range stage.Products {
		needCommit = removeDuplicatesSlice(stager, &product.SubProducts) || needCommit
	}

	for task := range stage.Tasks {
		needCommit = removeDuplicatesSlice(stager, &task.SubTasks) || needCommit
		needCommit = removeDuplicatesSlice(stager, &task.Inputs) || needCommit
		needCommit = removeDuplicatesSlice(stager, &task.Outputs) || needCommit
	}

	for note := range stage.Notes {
		needCommit = removeDuplicatesSlice(stager, &note.Products) || needCommit
		needCommit = removeDuplicatesSlice(stager, &note.Tasks) || needCommit
	}

	for resource := range stage.Resources {
		needCommit = removeDuplicatesSlice(stager, &resource.SubResources) || needCommit
		needCommit = removeDuplicatesSlice(stager, &resource.Tasks) || needCommit
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
