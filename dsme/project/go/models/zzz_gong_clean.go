// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}

	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by CompositionShape
func (compositionshape *CompositionShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	compositionshape.Product = GongCleanPointer(stage, compositionshape.Product)
}

// Clean garbage collect unstaged instances that are referenced by Diagram
func (diagram *Diagram) GongClean(stage *Stage) {
	// insertion point per field
	diagram.Product_Shapes = GongCleanSlice(stage, diagram.Product_Shapes)
	diagram.ProductsWhoseNodeIsExpanded = GongCleanSlice(stage, diagram.ProductsWhoseNodeIsExpanded)
	diagram.Composition_Shapes = GongCleanSlice(stage, diagram.Composition_Shapes)
	diagram.Task_Shapes = GongCleanSlice(stage, diagram.Task_Shapes)
	diagram.TasksWhoseNodeIsExpanded = GongCleanSlice(stage, diagram.TasksWhoseNodeIsExpanded)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Product
func (product *Product) GongClean(stage *Stage) {
	// insertion point per field
	product.SubProducts = GongCleanSlice(stage, product.SubProducts)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by ProductShape
func (productshape *ProductShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	productshape.Product = GongCleanPointer(stage, productshape.Product)
}

// Clean garbage collect unstaged instances that are referenced by Project
func (project *Project) GongClean(stage *Stage) {
	// insertion point per field
	project.RootProducts = GongCleanSlice(stage, project.RootProducts)
	project.RootTasks = GongCleanSlice(stage, project.RootTasks)
	project.Diagrams = GongCleanSlice(stage, project.Diagrams)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Root
func (root *Root) GongClean(stage *Stage) {
	// insertion point per field
	root.Projects = GongCleanSlice(stage, root.Projects)
	root.OrphanedProducts = GongCleanSlice(stage, root.OrphanedProducts)
	root.OrphanedTasks = GongCleanSlice(stage, root.OrphanedTasks)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Task
func (task *Task) GongClean(stage *Stage) {
	// insertion point per field
	task.SubTasks = GongCleanSlice(stage, task.SubTasks)
	task.Inputs = GongCleanSlice(stage, task.Inputs)
	task.Outputs = GongCleanSlice(stage, task.Outputs)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by TaskShape
func (taskshape *TaskShape) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	taskshape.Task = GongCleanPointer(stage, taskshape.Task)
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
