package models

import "time"

// enforceDuplicateRemove parses all associations between AbstractType
// and deletes duplicates within these associations
func (stager *Stager) enforceDuplicateRemove() (needCommit bool) {
	stage := stager.stage

	needCommit = removeDuplicatesSlice(stager, &stager.getRootLibrary().SubLibraries) || needCommit

	for diagramHierarchy := range *GetGongstructInstancesSetFromPointerType[*DiagramHierarchy](stage) {
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.Product_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.ProductComposition_Shapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.Task_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.TaskComposition_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.TaskInputShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.TaskOutputShapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.Note_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.NoteProductShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.NoteTaskShapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.NoteResourceShapes) || needCommit

		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.Resource_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.ResourceComposition_Shapes) || needCommit
		needCommit = removeDuplicatesSlice(stager, &diagramHierarchy.ResourceTaskShapes) || needCommit
	}

	for library := range *GetGongstructInstancesSetFromPointerType[*Library](stage) {
		needCommit = removeDuplicatesSlice(stager, &library.RootProducts) || needCommit
		needCommit = removeDuplicatesSlice(stager, &library.RootTasks) || needCommit
		needCommit = removeDuplicatesSlice(stager, &library.RootResources) || needCommit
		needCommit = removeDuplicatesSlice(stager, &library.Notes) || needCommit
	}

	for product := range *GetGongstructInstancesSetFromPointerType[*Product](stage) {
		needCommit = removeDuplicatesSlice(stager, &product.SubProducts) || needCommit
	}

	for task := range *GetGongstructInstancesSetFromPointerType[*Task](stage) {
		needCommit = removeDuplicatesSlice(stager, &task.SubTasks) || needCommit
		needCommit = removeDuplicatesSlice(stager, &task.Inputs) || needCommit
		needCommit = removeDuplicatesSlice(stager, &task.Outputs) || needCommit
	}

	for note := range *GetGongstructInstancesSetFromPointerType[*Note](stage) {
		needCommit = removeDuplicatesSlice(stager, &note.Products) || needCommit
		needCommit = removeDuplicatesSlice(stager, &note.Tasks) || needCommit
	}

	for resource := range *GetGongstructInstancesSetFromPointerType[*Resource](stage) {
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
