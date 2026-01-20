// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	*slice = cleanedSlice
	return len(cleanedSlice) != len(*slice)
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagram.Product_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.ProductsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.ProductComposition_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.Task_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.TasksWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.TasksWhoseInputNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.TasksWhoseOutputNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.TaskComposition_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.TaskInputShapes) || modified
	modified = GongCleanSlice(stage, &diagram.TaskOutputShapes) || modified
	modified = GongCleanSlice(stage, &diagram.Note_Shapes) || modified
	modified = GongCleanSlice(stage, &diagram.NotesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagram.NoteProductShapes) || modified
	modified = GongCleanSlice(stage, &diagram.NoteTaskShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &note.Products) || modified
	modified = GongCleanSlice(stage, &note.Tasks) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteProductShape
func (noteproductshape *NoteProductShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteproductshape.Note) || modified
	modified = GongCleanPointer(stage, &noteproductshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteTaskShape
func (notetaskshape *NoteTaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &notetaskshape.Note) || modified
	modified = GongCleanPointer(stage, &notetaskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Product
func (product *Product) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &product.SubProducts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ProductCompositionShape
func (productcompositionshape *ProductCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &productcompositionshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ProductShape
func (productshape *ProductShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &productshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Project
func (project *Project) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &project.RootProducts) || modified
	modified = GongCleanSlice(stage, &project.RootTasks) || modified
	modified = GongCleanSlice(stage, &project.Diagrams) || modified
	modified = GongCleanSlice(stage, &project.Notes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Root
func (root *Root) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &root.Projects) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Task
func (task *Task) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &task.SubTasks) || modified
	modified = GongCleanSlice(stage, &task.Inputs) || modified
	modified = GongCleanSlice(stage, &task.Outputs) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskCompositionShape
func (taskcompositionshape *TaskCompositionShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &taskcompositionshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskInputShape
func (taskinputshape *TaskInputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &taskinputshape.Task) || modified
	modified = GongCleanPointer(stage, &taskinputshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskOutputShape
func (taskoutputshape *TaskOutputShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &taskoutputshape.Task) || modified
	modified = GongCleanPointer(stage, &taskoutputshape.Product) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by TaskShape
func (taskshape *TaskShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &taskshape.Task) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
