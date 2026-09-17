// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagram.Product_Shapes) || modified
	modified = stage.CleanSlice(&diagram.ProductsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ProductComposition_Shapes) || modified
	modified = stage.CleanSlice(&diagram.Task_Shapes) || modified
	modified = stage.CleanSlice(&diagram.TasksWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.TasksWhoseInputNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.TasksWhoseOutputNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.TasksWhosePredecessorNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.TaskGroupShapes) || modified
	modified = stage.CleanSlice(&diagram.TaskGroupsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.TaskComposition_Shapes) || modified
	modified = stage.CleanSlice(&diagram.TaskInputShapes) || modified
	modified = stage.CleanSlice(&diagram.TaskOutputShapes) || modified
	modified = stage.CleanSlice(&diagram.TaskPredecessorShapes) || modified
	modified = stage.CleanSlice(&diagram.Note_Shapes) || modified
	modified = stage.CleanSlice(&diagram.NotesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.NoteProductShapes) || modified
	modified = stage.CleanSlice(&diagram.NoteTaskShapes) || modified
	modified = stage.CleanSlice(&diagram.NoteResourceShapes) || modified
	modified = stage.CleanSlice(&diagram.Resource_Shapes) || modified
	modified = stage.CleanSlice(&diagram.ResourcesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagram.ResourceComposition_Shapes) || modified
	modified = stage.CleanSlice(&diagram.ResourceTaskShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	modified = stage.CleanSlice(&library.RootProducts) || modified
	modified = stage.CleanSlice(&library.RootTasks) || modified
	modified = stage.CleanSlice(&library.RootTaskGroups) || modified
	modified = stage.CleanSlice(&library.RootResources) || modified
	modified = stage.CleanSlice(&library.Notes) || modified
	modified = stage.CleanSlice(&library.Diagrams) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&note.Products) || modified
	modified = stage.CleanSlice(&note.Tasks) || modified
	modified = stage.CleanSlice(&note.Resources) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteProductShape
func (noteproductshape *NoteProductShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteproductshape.Note) || modified
	modified = stage.CleanPointer(&noteproductshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteResourceShape
func (noteresourceshape *NoteResourceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteresourceshape.Note) || modified
	modified = stage.CleanPointer(&noteresourceshape.Resource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteTaskShape
func (notetaskshape *NoteTaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&notetaskshape.Note) || modified
	modified = stage.CleanPointer(&notetaskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Product
func (product *Product) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&product.SubProducts) || modified
	// insertion point per field
	modified = stage.CleanPointer(&product.ReferencedProduct) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProductCompositionShape
func (productcompositionshape *ProductCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&productcompositionshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProductShape
func (productshape *ProductShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&productshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Resource
func (resource *Resource) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&resource.Tasks) || modified
	modified = stage.CleanSlice(&resource.SubResources) || modified
	// insertion point per field
	modified = stage.CleanPointer(&resource.ReferencedResource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ResourceCompositionShape
func (resourcecompositionshape *ResourceCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&resourcecompositionshape.Resource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ResourceShape
func (resourceshape *ResourceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&resourceshape.Resource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ResourceTaskShape
func (resourcetaskshape *ResourceTaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&resourcetaskshape.Resource) || modified
	modified = stage.CleanPointer(&resourcetaskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Task
func (task *Task) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&task.Predecessors) || modified
	modified = stage.CleanSlice(&task.Inputs) || modified
	modified = stage.CleanSlice(&task.Outputs) || modified
	modified = stage.CleanSlice(&task.SubTasks) || modified
	modified = stage.CleanSlice(&task.TaskGroupsToDisplay) || modified
	// insertion point per field
	modified = stage.CleanPointer(&task.ReferencedTask) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskCompositionShape
func (taskcompositionshape *TaskCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taskcompositionshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskGroup
func (taskgroup *TaskGroup) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&taskgroup.Tasks) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskGroupShape
func (taskgroupshape *TaskGroupShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taskgroupshape.TaskGroup) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskInputShape
func (taskinputshape *TaskInputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taskinputshape.Product) || modified
	modified = stage.CleanPointer(&taskinputshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskOutputShape
func (taskoutputshape *TaskOutputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taskoutputshape.Task) || modified
	modified = stage.CleanPointer(&taskoutputshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskPredecessorShape
func (taskpredecessorshape *TaskPredecessorShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taskpredecessorshape.Predecessor) || modified
	modified = stage.CleanPointer(&taskpredecessorshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskShape
func (taskshape *TaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&taskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
