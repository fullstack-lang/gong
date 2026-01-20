package models

// enforceDuplicateRemove parses all associations between AbstractType
// and deletes duplicates within these associations
func (stager *Stager) enforceDuplicateRemove() (needCommit bool) {
	stage := stager.stage

	if removeDuplicatesSlice(&stager.root.Projects) {
		needCommit = true
	}

	for project := range stage.Projects {
		if removeDuplicatesSlice(&project.RootProducts) {
			needCommit = true
		}
		if removeDuplicatesSlice(&project.RootTasks) {
			needCommit = true
		}
		if removeDuplicatesSlice(&project.Notes) {
			needCommit = true
		}
	}

	for product := range stage.Products {
		if removeDuplicatesSlice(&product.SubProducts) {
			needCommit = true
		}
	}

	for task := range stage.Tasks {
		if removeDuplicatesSlice(&task.SubTasks) {
			needCommit = true
		}
		if removeDuplicatesSlice(&task.Inputs) {
			needCommit = true
		}
		if removeDuplicatesSlice(&task.Outputs) {
			needCommit = true
		}
	}

	for note := range stage.Notes {
		if removeDuplicatesSlice(&note.Products) {
			needCommit = true
		}
		if removeDuplicatesSlice(&note.Tasks) {
			needCommit = true
		}
	}

	return
}

// removeDuplicatesSlice is a generic helper that removes duplicate pointers from a slice.
// It keeps the first occurrence of each element and preserves order.
// It returns true if any duplicates were removed.
func removeDuplicatesSlice[T comparable](slicePtr *[]T) bool {
	keys := make(map[T]bool)
	list := []T{}
	changed := false
	for _, entry := range *slicePtr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		} else {
			changed = true
		}
	}
	if changed {
		*slicePtr = list
	}
	return changed
}
