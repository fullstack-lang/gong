// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (diagram *Diagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	return diagram.GongIsStaged(stage)
}

func (library *Library) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	return library.GongIsStaged(stage)
}

func (note *Note) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	return note.GongIsStaged(stage)
}

func (noteproductshape *NoteProductShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteProductShapes[noteproductshape]

	return
}

func (stage *Stage) IsStagedNoteProductShape(noteproductshape *NoteProductShape) (ok bool) {

	return noteproductshape.GongIsStaged(stage)
}

func (noteresourceshape *NoteResourceShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteResourceShapes[noteresourceshape]

	return
}

func (stage *Stage) IsStagedNoteResourceShape(noteresourceshape *NoteResourceShape) (ok bool) {

	return noteresourceshape.GongIsStaged(stage)
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	return noteshape.GongIsStaged(stage)
}

func (notetaskshape *NoteTaskShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteTaskShapes[notetaskshape]

	return
}

func (stage *Stage) IsStagedNoteTaskShape(notetaskshape *NoteTaskShape) (ok bool) {

	return notetaskshape.GongIsStaged(stage)
}

func (product *Product) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Products[product]

	return
}

func (stage *Stage) IsStagedProduct(product *Product) (ok bool) {

	return product.GongIsStaged(stage)
}

func (productcompositionshape *ProductCompositionShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ProductCompositionShapes[productcompositionshape]

	return
}

func (stage *Stage) IsStagedProductCompositionShape(productcompositionshape *ProductCompositionShape) (ok bool) {

	return productcompositionshape.GongIsStaged(stage)
}

func (productreferenceshape *ProductReferenceShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ProductReferenceShapes[productreferenceshape]

	return
}

func (stage *Stage) IsStagedProductReferenceShape(productreferenceshape *ProductReferenceShape) (ok bool) {

	return productreferenceshape.GongIsStaged(stage)
}

func (productshape *ProductShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ProductShapes[productshape]

	return
}

func (stage *Stage) IsStagedProductShape(productshape *ProductShape) (ok bool) {

	return productshape.GongIsStaged(stage)
}

func (resource *Resource) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Resources[resource]

	return
}

func (stage *Stage) IsStagedResource(resource *Resource) (ok bool) {

	return resource.GongIsStaged(stage)
}

func (resourcecompositionshape *ResourceCompositionShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ResourceCompositionShapes[resourcecompositionshape]

	return
}

func (stage *Stage) IsStagedResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape) (ok bool) {

	return resourcecompositionshape.GongIsStaged(stage)
}

func (resourceshape *ResourceShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ResourceShapes[resourceshape]

	return
}

func (stage *Stage) IsStagedResourceShape(resourceshape *ResourceShape) (ok bool) {

	return resourceshape.GongIsStaged(stage)
}

func (resourcetaskshape *ResourceTaskShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ResourceTaskShapes[resourcetaskshape]

	return
}

func (stage *Stage) IsStagedResourceTaskShape(resourcetaskshape *ResourceTaskShape) (ok bool) {

	return resourcetaskshape.GongIsStaged(stage)
}

func (task *Task) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Tasks[task]

	return
}

func (stage *Stage) IsStagedTask(task *Task) (ok bool) {

	return task.GongIsStaged(stage)
}

func (taskcompositionshape *TaskCompositionShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskCompositionShapes[taskcompositionshape]

	return
}

func (stage *Stage) IsStagedTaskCompositionShape(taskcompositionshape *TaskCompositionShape) (ok bool) {

	return taskcompositionshape.GongIsStaged(stage)
}

func (taskgroup *TaskGroup) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskGroups[taskgroup]

	return
}

func (stage *Stage) IsStagedTaskGroup(taskgroup *TaskGroup) (ok bool) {

	return taskgroup.GongIsStaged(stage)
}

func (taskgroupshape *TaskGroupShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskGroupShapes[taskgroupshape]

	return
}

func (stage *Stage) IsStagedTaskGroupShape(taskgroupshape *TaskGroupShape) (ok bool) {

	return taskgroupshape.GongIsStaged(stage)
}

func (taskinputshape *TaskInputShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskInputShapes[taskinputshape]

	return
}

func (stage *Stage) IsStagedTaskInputShape(taskinputshape *TaskInputShape) (ok bool) {

	return taskinputshape.GongIsStaged(stage)
}

func (taskoutputshape *TaskOutputShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskOutputShapes[taskoutputshape]

	return
}

func (stage *Stage) IsStagedTaskOutputShape(taskoutputshape *TaskOutputShape) (ok bool) {

	return taskoutputshape.GongIsStaged(stage)
}

func (taskpredecessorshape *TaskPredecessorShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskPredecessorShapes[taskpredecessorshape]

	return
}

func (stage *Stage) IsStagedTaskPredecessorShape(taskpredecessorshape *TaskPredecessorShape) (ok bool) {

	return taskpredecessorshape.GongIsStaged(stage)
}

func (taskshape *TaskShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.TaskShapes[taskshape]

	return
}

func (stage *Stage) IsStagedTaskShape(taskshape *TaskShape) (ok bool) {

	return taskshape.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (diagram *Diagram) GongStageBranch(stage *Stage) {
	stage.StageBranchDiagram(diagram)
}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if stage.IsStaged(diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagram.Product_Shapes {
		stage.StageBranch(_productshape)
	}
	for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
		stage.StageBranch(_product)
	}
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		stage.StageBranch(_productcompositionshape)
	}
	for _, _productreferenceshape := range diagram.ProductReference_Shapes {
		stage.StageBranch(_productreferenceshape)
	}
	for _, _taskshape := range diagram.Task_Shapes {
		stage.StageBranch(_taskshape)
	}
	for _, _task := range diagram.TasksWhoseNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _task := range diagram.TasksWhosePredecessorNodeIsExpanded {
		stage.StageBranch(_task)
	}
	for _, _taskgroupshape := range diagram.TaskGroupShapes {
		stage.StageBranch(_taskgroupshape)
	}
	for _, _taskgroup := range diagram.TaskGroupsWhoseNodeIsExpanded {
		stage.StageBranch(_taskgroup)
	}
	for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
		stage.StageBranch(_taskcompositionshape)
	}
	for _, _taskinputshape := range diagram.TaskInputShapes {
		stage.StageBranch(_taskinputshape)
	}
	for _, _taskoutputshape := range diagram.TaskOutputShapes {
		stage.StageBranch(_taskoutputshape)
	}
	for _, _taskpredecessorshape := range diagram.TaskPredecessorShapes {
		stage.StageBranch(_taskpredecessorshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		stage.StageBranch(_noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}
	for _, _noteproductshape := range diagram.NoteProductShapes {
		stage.StageBranch(_noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		stage.StageBranch(_notetaskshape)
	}
	for _, _noteresourceshape := range diagram.NoteResourceShapes {
		stage.StageBranch(_noteresourceshape)
	}
	for _, _resourceshape := range diagram.Resource_Shapes {
		stage.StageBranch(_resourceshape)
	}
	for _, _resource := range diagram.ResourcesWhoseNodeIsExpanded {
		stage.StageBranch(_resource)
	}
	for _, _resourcecompositionshape := range diagram.ResourceComposition_Shapes {
		stage.StageBranch(_resourcecompositionshape)
	}
	for _, _resourcetaskshape := range diagram.ResourceTaskShapes {
		stage.StageBranch(_resourcetaskshape)
	}

}

func (library *Library) GongStageBranch(stage *Stage) {
	stage.StageBranchLibrary(library)
}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}
	for _, _product := range library.RootProducts {
		stage.StageBranch(_product)
	}
	for _, _task := range library.RootTasks {
		stage.StageBranch(_task)
	}
	for _, _taskgroup := range library.RootTaskGroups {
		stage.StageBranch(_taskgroup)
	}
	for _, _resource := range library.RootResources {
		stage.StageBranch(_resource)
	}
	for _, _note := range library.Notes {
		stage.StageBranch(_note)
	}
	for _, _diagram := range library.Diagrams {
		stage.StageBranch(_diagram)
	}

}

func (note *Note) GongStageBranch(stage *Stage) {
	stage.StageBranchNote(note)
}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range note.Products {
		stage.StageBranch(_product)
	}
	for _, _task := range note.Tasks {
		stage.StageBranch(_task)
	}
	for _, _resource := range note.Resources {
		stage.StageBranch(_resource)
	}

}

func (noteproductshape *NoteProductShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteProductShape(noteproductshape)
}

func (stage *Stage) StageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if stage.IsStaged(noteproductshape) {
		return
	}

	noteproductshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		stage.StageBranch(noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		stage.StageBranch(noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteresourceshape *NoteResourceShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteResourceShape(noteresourceshape)
}

func (stage *Stage) StageBranchNoteResourceShape(noteresourceshape *NoteResourceShape) {

	// check if instance is already staged
	if stage.IsStaged(noteresourceshape) {
		return
	}

	noteresourceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteresourceshape.Note != nil {
		stage.StageBranch(noteresourceshape.Note)
	}
	if noteresourceshape.Resource != nil {
		stage.StageBranch(noteresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteShape(noteshape)
}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if stage.IsStaged(noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.StageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notetaskshape *NoteTaskShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteTaskShape(notetaskshape)
}

func (stage *Stage) StageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if stage.IsStaged(notetaskshape) {
		return
	}

	notetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		stage.StageBranch(notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		stage.StageBranch(notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (product *Product) GongStageBranch(stage *Stage) {
	stage.StageBranchProduct(product)
}

func (stage *Stage) StageBranchProduct(product *Product) {

	// check if instance is already staged
	if stage.IsStaged(product) {
		return
	}

	product.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if product.ReferencedProduct != nil {
		stage.StageBranch(product.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range product.SubProducts {
		stage.StageBranch(_product)
	}

}

func (productcompositionshape *ProductCompositionShape) GongStageBranch(stage *Stage) {
	stage.StageBranchProductCompositionShape(productcompositionshape)
}

func (stage *Stage) StageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if stage.IsStaged(productcompositionshape) {
		return
	}

	productcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		stage.StageBranch(productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (productreferenceshape *ProductReferenceShape) GongStageBranch(stage *Stage) {
	stage.StageBranchProductReferenceShape(productreferenceshape)
}

func (stage *Stage) StageBranchProductReferenceShape(productreferenceshape *ProductReferenceShape) {

	// check if instance is already staged
	if stage.IsStaged(productreferenceshape) {
		return
	}

	productreferenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productreferenceshape.Product != nil {
		stage.StageBranch(productreferenceshape.Product)
	}
	if productreferenceshape.ReferencedProduct != nil {
		stage.StageBranch(productreferenceshape.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (productshape *ProductShape) GongStageBranch(stage *Stage) {
	stage.StageBranchProductShape(productshape)
}

func (stage *Stage) StageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if stage.IsStaged(productshape) {
		return
	}

	productshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		stage.StageBranch(productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resource *Resource) GongStageBranch(stage *Stage) {
	stage.StageBranchResource(resource)
}

func (stage *Stage) StageBranchResource(resource *Resource) {

	// check if instance is already staged
	if stage.IsStaged(resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resource.ReferencedResource != nil {
		stage.StageBranch(resource.ReferencedResource)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range resource.Tasks {
		stage.StageBranch(_task)
	}
	for _, _resource := range resource.SubResources {
		stage.StageBranch(_resource)
	}

}

func (resourcecompositionshape *ResourceCompositionShape) GongStageBranch(stage *Stage) {
	stage.StageBranchResourceCompositionShape(resourcecompositionshape)
}

func (stage *Stage) StageBranchResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape) {

	// check if instance is already staged
	if stage.IsStaged(resourcecompositionshape) {
		return
	}

	resourcecompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcecompositionshape.Resource != nil {
		stage.StageBranch(resourcecompositionshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resourceshape *ResourceShape) GongStageBranch(stage *Stage) {
	stage.StageBranchResourceShape(resourceshape)
}

func (stage *Stage) StageBranchResourceShape(resourceshape *ResourceShape) {

	// check if instance is already staged
	if stage.IsStaged(resourceshape) {
		return
	}

	resourceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourceshape.Resource != nil {
		stage.StageBranch(resourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resourcetaskshape *ResourceTaskShape) GongStageBranch(stage *Stage) {
	stage.StageBranchResourceTaskShape(resourcetaskshape)
}

func (stage *Stage) StageBranchResourceTaskShape(resourcetaskshape *ResourceTaskShape) {

	// check if instance is already staged
	if stage.IsStaged(resourcetaskshape) {
		return
	}

	resourcetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcetaskshape.Resource != nil {
		stage.StageBranch(resourcetaskshape.Resource)
	}
	if resourcetaskshape.Task != nil {
		stage.StageBranch(resourcetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (task *Task) GongStageBranch(stage *Stage) {
	stage.StageBranchTask(task)
}

func (stage *Stage) StageBranchTask(task *Task) {

	// check if instance is already staged
	if stage.IsStaged(task) {
		return
	}

	task.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if task.ReferencedTask != nil {
		stage.StageBranch(task.ReferencedTask)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range task.Predecessors {
		stage.StageBranch(_task)
	}
	for _, _product := range task.Inputs {
		stage.StageBranch(_product)
	}
	for _, _product := range task.Outputs {
		stage.StageBranch(_product)
	}
	for _, _task := range task.SubTasks {
		stage.StageBranch(_task)
	}
	for _, _taskgroup := range task.TaskGroupsToDisplay {
		stage.StageBranch(_taskgroup)
	}

}

func (taskcompositionshape *TaskCompositionShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskCompositionShape(taskcompositionshape)
}

func (stage *Stage) StageBranchTaskCompositionShape(taskcompositionshape *TaskCompositionShape) {

	// check if instance is already staged
	if stage.IsStaged(taskcompositionshape) {
		return
	}

	taskcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshape.Task != nil {
		stage.StageBranch(taskcompositionshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskgroup *TaskGroup) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskGroup(taskgroup)
}

func (stage *Stage) StageBranchTaskGroup(taskgroup *TaskGroup) {

	// check if instance is already staged
	if stage.IsStaged(taskgroup) {
		return
	}

	taskgroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskgroup.Tasks {
		stage.StageBranch(_task)
	}

}

func (taskgroupshape *TaskGroupShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskGroupShape(taskgroupshape)
}

func (stage *Stage) StageBranchTaskGroupShape(taskgroupshape *TaskGroupShape) {

	// check if instance is already staged
	if stage.IsStaged(taskgroupshape) {
		return
	}

	taskgroupshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskgroupshape.TaskGroup != nil {
		stage.StageBranch(taskgroupshape.TaskGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskinputshape *TaskInputShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskInputShape(taskinputshape)
}

func (stage *Stage) StageBranchTaskInputShape(taskinputshape *TaskInputShape) {

	// check if instance is already staged
	if stage.IsStaged(taskinputshape) {
		return
	}

	taskinputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshape.Product != nil {
		stage.StageBranch(taskinputshape.Product)
	}
	if taskinputshape.Task != nil {
		stage.StageBranch(taskinputshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskoutputshape *TaskOutputShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskOutputShape(taskoutputshape)
}

func (stage *Stage) StageBranchTaskOutputShape(taskoutputshape *TaskOutputShape) {

	// check if instance is already staged
	if stage.IsStaged(taskoutputshape) {
		return
	}

	taskoutputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshape.Task != nil {
		stage.StageBranch(taskoutputshape.Task)
	}
	if taskoutputshape.Product != nil {
		stage.StageBranch(taskoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskpredecessorshape *TaskPredecessorShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskPredecessorShape(taskpredecessorshape)
}

func (stage *Stage) StageBranchTaskPredecessorShape(taskpredecessorshape *TaskPredecessorShape) {

	// check if instance is already staged
	if stage.IsStaged(taskpredecessorshape) {
		return
	}

	taskpredecessorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskpredecessorshape.Predecessor != nil {
		stage.StageBranch(taskpredecessorshape.Predecessor)
	}
	if taskpredecessorshape.Task != nil {
		stage.StageBranch(taskpredecessorshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskshape *TaskShape) GongStageBranch(stage *Stage) {
	stage.StageBranchTaskShape(taskshape)
}

func (stage *Stage) StageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if stage.IsStaged(taskshape) {
		return
	}

	taskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		stage.StageBranch(taskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Diagram:
		toT := GongCopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteProductShape:
		toT := GongCopyBranchNoteProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteResourceShape:
		toT := GongCopyBranchNoteResourceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := GongCopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteTaskShape:
		toT := GongCopyBranchNoteTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Product:
		toT := GongCopyBranchProduct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductCompositionShape:
		toT := GongCopyBranchProductCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductReferenceShape:
		toT := GongCopyBranchProductReferenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductShape:
		toT := GongCopyBranchProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Resource:
		toT := GongCopyBranchResource(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ResourceCompositionShape:
		toT := GongCopyBranchResourceCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ResourceShape:
		toT := GongCopyBranchResourceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ResourceTaskShape:
		toT := GongCopyBranchResourceTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Task:
		toT := GongCopyBranchTask(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskCompositionShape:
		toT := GongCopyBranchTaskCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskGroup:
		toT := GongCopyBranchTaskGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskGroupShape:
		toT := GongCopyBranchTaskGroupShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskInputShape:
		toT := GongCopyBranchTaskInputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskOutputShape:
		toT := GongCopyBranchTaskOutputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskPredecessorShape:
		toT := GongCopyBranchTaskPredecessorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskShape:
		toT := GongCopyBranchTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.GongCopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagramFrom.Product_Shapes {
		diagramTo.Product_Shapes = append(diagramTo.Product_Shapes, GongCopyBranchProductShape(mapOrigCopy, _productshape))
	}
	for _, _product := range diagramFrom.ProductsWhoseNodeIsExpanded {
		diagramTo.ProductsWhoseNodeIsExpanded = append(diagramTo.ProductsWhoseNodeIsExpanded, GongCopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _productcompositionshape := range diagramFrom.ProductComposition_Shapes {
		diagramTo.ProductComposition_Shapes = append(diagramTo.ProductComposition_Shapes, GongCopyBranchProductCompositionShape(mapOrigCopy, _productcompositionshape))
	}
	for _, _productreferenceshape := range diagramFrom.ProductReference_Shapes {
		diagramTo.ProductReference_Shapes = append(diagramTo.ProductReference_Shapes, GongCopyBranchProductReferenceShape(mapOrigCopy, _productreferenceshape))
	}
	for _, _taskshape := range diagramFrom.Task_Shapes {
		diagramTo.Task_Shapes = append(diagramTo.Task_Shapes, GongCopyBranchTaskShape(mapOrigCopy, _taskshape))
	}
	for _, _task := range diagramFrom.TasksWhoseNodeIsExpanded {
		diagramTo.TasksWhoseNodeIsExpanded = append(diagramTo.TasksWhoseNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramFrom.TasksWhoseInputNodeIsExpanded {
		diagramTo.TasksWhoseInputNodeIsExpanded = append(diagramTo.TasksWhoseInputNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramFrom.TasksWhoseOutputNodeIsExpanded {
		diagramTo.TasksWhoseOutputNodeIsExpanded = append(diagramTo.TasksWhoseOutputNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramFrom.TasksWhosePredecessorNodeIsExpanded {
		diagramTo.TasksWhosePredecessorNodeIsExpanded = append(diagramTo.TasksWhosePredecessorNodeIsExpanded, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskgroupshape := range diagramFrom.TaskGroupShapes {
		diagramTo.TaskGroupShapes = append(diagramTo.TaskGroupShapes, GongCopyBranchTaskGroupShape(mapOrigCopy, _taskgroupshape))
	}
	for _, _taskgroup := range diagramFrom.TaskGroupsWhoseNodeIsExpanded {
		diagramTo.TaskGroupsWhoseNodeIsExpanded = append(diagramTo.TaskGroupsWhoseNodeIsExpanded, GongCopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}
	for _, _taskcompositionshape := range diagramFrom.TaskComposition_Shapes {
		diagramTo.TaskComposition_Shapes = append(diagramTo.TaskComposition_Shapes, GongCopyBranchTaskCompositionShape(mapOrigCopy, _taskcompositionshape))
	}
	for _, _taskinputshape := range diagramFrom.TaskInputShapes {
		diagramTo.TaskInputShapes = append(diagramTo.TaskInputShapes, GongCopyBranchTaskInputShape(mapOrigCopy, _taskinputshape))
	}
	for _, _taskoutputshape := range diagramFrom.TaskOutputShapes {
		diagramTo.TaskOutputShapes = append(diagramTo.TaskOutputShapes, GongCopyBranchTaskOutputShape(mapOrigCopy, _taskoutputshape))
	}
	for _, _taskpredecessorshape := range diagramFrom.TaskPredecessorShapes {
		diagramTo.TaskPredecessorShapes = append(diagramTo.TaskPredecessorShapes, GongCopyBranchTaskPredecessorShape(mapOrigCopy, _taskpredecessorshape))
	}
	for _, _noteshape := range diagramFrom.Note_Shapes {
		diagramTo.Note_Shapes = append(diagramTo.Note_Shapes, GongCopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramFrom.NotesWhoseNodeIsExpanded {
		diagramTo.NotesWhoseNodeIsExpanded = append(diagramTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _noteproductshape := range diagramFrom.NoteProductShapes {
		diagramTo.NoteProductShapes = append(diagramTo.NoteProductShapes, GongCopyBranchNoteProductShape(mapOrigCopy, _noteproductshape))
	}
	for _, _notetaskshape := range diagramFrom.NoteTaskShapes {
		diagramTo.NoteTaskShapes = append(diagramTo.NoteTaskShapes, GongCopyBranchNoteTaskShape(mapOrigCopy, _notetaskshape))
	}
	for _, _noteresourceshape := range diagramFrom.NoteResourceShapes {
		diagramTo.NoteResourceShapes = append(diagramTo.NoteResourceShapes, GongCopyBranchNoteResourceShape(mapOrigCopy, _noteresourceshape))
	}
	for _, _resourceshape := range diagramFrom.Resource_Shapes {
		diagramTo.Resource_Shapes = append(diagramTo.Resource_Shapes, GongCopyBranchResourceShape(mapOrigCopy, _resourceshape))
	}
	for _, _resource := range diagramFrom.ResourcesWhoseNodeIsExpanded {
		diagramTo.ResourcesWhoseNodeIsExpanded = append(diagramTo.ResourcesWhoseNodeIsExpanded, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _resourcecompositionshape := range diagramFrom.ResourceComposition_Shapes {
		diagramTo.ResourceComposition_Shapes = append(diagramTo.ResourceComposition_Shapes, GongCopyBranchResourceCompositionShape(mapOrigCopy, _resourcecompositionshape))
	}
	for _, _resourcetaskshape := range diagramFrom.ResourceTaskShapes {
		diagramTo.ResourceTaskShapes = append(diagramTo.ResourceTaskShapes, GongCopyBranchResourceTaskShape(mapOrigCopy, _resourcetaskshape))
	}

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _product := range libraryFrom.RootProducts {
		libraryTo.RootProducts = append(libraryTo.RootProducts, GongCopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _task := range libraryFrom.RootTasks {
		libraryTo.RootTasks = append(libraryTo.RootTasks, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskgroup := range libraryFrom.RootTaskGroups {
		libraryTo.RootTaskGroups = append(libraryTo.RootTaskGroups, GongCopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}
	for _, _resource := range libraryFrom.RootResources {
		libraryTo.RootResources = append(libraryTo.RootResources, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _note := range libraryFrom.Notes {
		libraryTo.Notes = append(libraryTo.Notes, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _diagram := range libraryFrom.Diagrams {
		libraryTo.Diagrams = append(libraryTo.Diagrams, GongCopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range noteFrom.Products {
		noteTo.Products = append(noteTo.Products, GongCopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _task := range noteFrom.Tasks {
		noteTo.Tasks = append(noteTo.Tasks, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _resource := range noteFrom.Resources {
		noteTo.Resources = append(noteTo.Resources, GongCopyBranchResource(mapOrigCopy, _resource))
	}

	return
}

func GongCopyBranchNoteProductShape(mapOrigCopy map[any]any, noteproductshapeFrom *NoteProductShape) (noteproductshapeTo *NoteProductShape) {

	// noteproductshapeFrom has already been copied
	if _noteproductshapeTo, ok := mapOrigCopy[noteproductshapeFrom]; ok {
		noteproductshapeTo = _noteproductshapeTo.(*NoteProductShape)
		return
	}

	noteproductshapeTo = new(NoteProductShape)
	mapOrigCopy[noteproductshapeFrom] = noteproductshapeTo
	noteproductshapeFrom.GongCopyBasicFields(noteproductshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshapeFrom.Note != nil {
		noteproductshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteproductshapeFrom.Note)
	}
	if noteproductshapeFrom.Product != nil {
		noteproductshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, noteproductshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteResourceShape(mapOrigCopy map[any]any, noteresourceshapeFrom *NoteResourceShape) (noteresourceshapeTo *NoteResourceShape) {

	// noteresourceshapeFrom has already been copied
	if _noteresourceshapeTo, ok := mapOrigCopy[noteresourceshapeFrom]; ok {
		noteresourceshapeTo = _noteresourceshapeTo.(*NoteResourceShape)
		return
	}

	noteresourceshapeTo = new(NoteResourceShape)
	mapOrigCopy[noteresourceshapeFrom] = noteresourceshapeTo
	noteresourceshapeFrom.GongCopyBasicFields(noteresourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteresourceshapeFrom.Note != nil {
		noteresourceshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteresourceshapeFrom.Note)
	}
	if noteresourceshapeFrom.Resource != nil {
		noteresourceshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, noteresourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {

	// notetaskshapeFrom has already been copied
	if _notetaskshapeTo, ok := mapOrigCopy[notetaskshapeFrom]; ok {
		notetaskshapeTo = _notetaskshapeTo.(*NoteTaskShape)
		return
	}

	notetaskshapeTo = new(NoteTaskShape)
	mapOrigCopy[notetaskshapeFrom] = notetaskshapeTo
	notetaskshapeFrom.GongCopyBasicFields(notetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshapeFrom.Note != nil {
		notetaskshapeTo.Note = GongCopyBranchNote(mapOrigCopy, notetaskshapeFrom.Note)
	}
	if notetaskshapeFrom.Task != nil {
		notetaskshapeTo.Task = GongCopyBranchTask(mapOrigCopy, notetaskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchProduct(mapOrigCopy map[any]any, productFrom *Product) (productTo *Product) {

	// productFrom has already been copied
	if _productTo, ok := mapOrigCopy[productFrom]; ok {
		productTo = _productTo.(*Product)
		return
	}

	productTo = new(Product)
	mapOrigCopy[productFrom] = productTo
	productFrom.GongCopyBasicFields(productTo)

	//insertion point for the staging of instances referenced by pointers
	if productFrom.ReferencedProduct != nil {
		productTo.ReferencedProduct = GongCopyBranchProduct(mapOrigCopy, productFrom.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range productFrom.SubProducts {
		productTo.SubProducts = append(productTo.SubProducts, GongCopyBranchProduct(mapOrigCopy, _product))
	}

	return
}

func GongCopyBranchProductCompositionShape(mapOrigCopy map[any]any, productcompositionshapeFrom *ProductCompositionShape) (productcompositionshapeTo *ProductCompositionShape) {

	// productcompositionshapeFrom has already been copied
	if _productcompositionshapeTo, ok := mapOrigCopy[productcompositionshapeFrom]; ok {
		productcompositionshapeTo = _productcompositionshapeTo.(*ProductCompositionShape)
		return
	}

	productcompositionshapeTo = new(ProductCompositionShape)
	mapOrigCopy[productcompositionshapeFrom] = productcompositionshapeTo
	productcompositionshapeFrom.GongCopyBasicFields(productcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshapeFrom.Product != nil {
		productcompositionshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, productcompositionshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchProductReferenceShape(mapOrigCopy map[any]any, productreferenceshapeFrom *ProductReferenceShape) (productreferenceshapeTo *ProductReferenceShape) {

	// productreferenceshapeFrom has already been copied
	if _productreferenceshapeTo, ok := mapOrigCopy[productreferenceshapeFrom]; ok {
		productreferenceshapeTo = _productreferenceshapeTo.(*ProductReferenceShape)
		return
	}

	productreferenceshapeTo = new(ProductReferenceShape)
	mapOrigCopy[productreferenceshapeFrom] = productreferenceshapeTo
	productreferenceshapeFrom.GongCopyBasicFields(productreferenceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productreferenceshapeFrom.Product != nil {
		productreferenceshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, productreferenceshapeFrom.Product)
	}
	if productreferenceshapeFrom.ReferencedProduct != nil {
		productreferenceshapeTo.ReferencedProduct = GongCopyBranchProduct(mapOrigCopy, productreferenceshapeFrom.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchProductShape(mapOrigCopy map[any]any, productshapeFrom *ProductShape) (productshapeTo *ProductShape) {

	// productshapeFrom has already been copied
	if _productshapeTo, ok := mapOrigCopy[productshapeFrom]; ok {
		productshapeTo = _productshapeTo.(*ProductShape)
		return
	}

	productshapeTo = new(ProductShape)
	mapOrigCopy[productshapeFrom] = productshapeTo
	productshapeFrom.GongCopyBasicFields(productshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productshapeFrom.Product != nil {
		productshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, productshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {

	// resourceFrom has already been copied
	if _resourceTo, ok := mapOrigCopy[resourceFrom]; ok {
		resourceTo = _resourceTo.(*Resource)
		return
	}

	resourceTo = new(Resource)
	mapOrigCopy[resourceFrom] = resourceTo
	resourceFrom.GongCopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers
	if resourceFrom.ReferencedResource != nil {
		resourceTo.ReferencedResource = GongCopyBranchResource(mapOrigCopy, resourceFrom.ReferencedResource)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range resourceFrom.Tasks {
		resourceTo.Tasks = append(resourceTo.Tasks, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _resource := range resourceFrom.SubResources {
		resourceTo.SubResources = append(resourceTo.SubResources, GongCopyBranchResource(mapOrigCopy, _resource))
	}

	return
}

func GongCopyBranchResourceCompositionShape(mapOrigCopy map[any]any, resourcecompositionshapeFrom *ResourceCompositionShape) (resourcecompositionshapeTo *ResourceCompositionShape) {

	// resourcecompositionshapeFrom has already been copied
	if _resourcecompositionshapeTo, ok := mapOrigCopy[resourcecompositionshapeFrom]; ok {
		resourcecompositionshapeTo = _resourcecompositionshapeTo.(*ResourceCompositionShape)
		return
	}

	resourcecompositionshapeTo = new(ResourceCompositionShape)
	mapOrigCopy[resourcecompositionshapeFrom] = resourcecompositionshapeTo
	resourcecompositionshapeFrom.GongCopyBasicFields(resourcecompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourcecompositionshapeFrom.Resource != nil {
		resourcecompositionshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, resourcecompositionshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResourceShape(mapOrigCopy map[any]any, resourceshapeFrom *ResourceShape) (resourceshapeTo *ResourceShape) {

	// resourceshapeFrom has already been copied
	if _resourceshapeTo, ok := mapOrigCopy[resourceshapeFrom]; ok {
		resourceshapeTo = _resourceshapeTo.(*ResourceShape)
		return
	}

	resourceshapeTo = new(ResourceShape)
	mapOrigCopy[resourceshapeFrom] = resourceshapeTo
	resourceshapeFrom.GongCopyBasicFields(resourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourceshapeFrom.Resource != nil {
		resourceshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, resourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResourceTaskShape(mapOrigCopy map[any]any, resourcetaskshapeFrom *ResourceTaskShape) (resourcetaskshapeTo *ResourceTaskShape) {

	// resourcetaskshapeFrom has already been copied
	if _resourcetaskshapeTo, ok := mapOrigCopy[resourcetaskshapeFrom]; ok {
		resourcetaskshapeTo = _resourcetaskshapeTo.(*ResourceTaskShape)
		return
	}

	resourcetaskshapeTo = new(ResourceTaskShape)
	mapOrigCopy[resourcetaskshapeFrom] = resourcetaskshapeTo
	resourcetaskshapeFrom.GongCopyBasicFields(resourcetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourcetaskshapeFrom.Resource != nil {
		resourcetaskshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, resourcetaskshapeFrom.Resource)
	}
	if resourcetaskshapeFrom.Task != nil {
		resourcetaskshapeTo.Task = GongCopyBranchTask(mapOrigCopy, resourcetaskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTask(mapOrigCopy map[any]any, taskFrom *Task) (taskTo *Task) {

	// taskFrom has already been copied
	if _taskTo, ok := mapOrigCopy[taskFrom]; ok {
		taskTo = _taskTo.(*Task)
		return
	}

	taskTo = new(Task)
	mapOrigCopy[taskFrom] = taskTo
	taskFrom.GongCopyBasicFields(taskTo)

	//insertion point for the staging of instances referenced by pointers
	if taskFrom.ReferencedTask != nil {
		taskTo.ReferencedTask = GongCopyBranchTask(mapOrigCopy, taskFrom.ReferencedTask)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskFrom.Predecessors {
		taskTo.Predecessors = append(taskTo.Predecessors, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _product := range taskFrom.Inputs {
		taskTo.Inputs = append(taskTo.Inputs, GongCopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _product := range taskFrom.Outputs {
		taskTo.Outputs = append(taskTo.Outputs, GongCopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _task := range taskFrom.SubTasks {
		taskTo.SubTasks = append(taskTo.SubTasks, GongCopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskgroup := range taskFrom.TaskGroupsToDisplay {
		taskTo.TaskGroupsToDisplay = append(taskTo.TaskGroupsToDisplay, GongCopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}

	return
}

func GongCopyBranchTaskCompositionShape(mapOrigCopy map[any]any, taskcompositionshapeFrom *TaskCompositionShape) (taskcompositionshapeTo *TaskCompositionShape) {

	// taskcompositionshapeFrom has already been copied
	if _taskcompositionshapeTo, ok := mapOrigCopy[taskcompositionshapeFrom]; ok {
		taskcompositionshapeTo = _taskcompositionshapeTo.(*TaskCompositionShape)
		return
	}

	taskcompositionshapeTo = new(TaskCompositionShape)
	mapOrigCopy[taskcompositionshapeFrom] = taskcompositionshapeTo
	taskcompositionshapeFrom.GongCopyBasicFields(taskcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshapeFrom.Task != nil {
		taskcompositionshapeTo.Task = GongCopyBranchTask(mapOrigCopy, taskcompositionshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskGroup(mapOrigCopy map[any]any, taskgroupFrom *TaskGroup) (taskgroupTo *TaskGroup) {

	// taskgroupFrom has already been copied
	if _taskgroupTo, ok := mapOrigCopy[taskgroupFrom]; ok {
		taskgroupTo = _taskgroupTo.(*TaskGroup)
		return
	}

	taskgroupTo = new(TaskGroup)
	mapOrigCopy[taskgroupFrom] = taskgroupTo
	taskgroupFrom.GongCopyBasicFields(taskgroupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskgroupFrom.Tasks {
		taskgroupTo.Tasks = append(taskgroupTo.Tasks, GongCopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func GongCopyBranchTaskGroupShape(mapOrigCopy map[any]any, taskgroupshapeFrom *TaskGroupShape) (taskgroupshapeTo *TaskGroupShape) {

	// taskgroupshapeFrom has already been copied
	if _taskgroupshapeTo, ok := mapOrigCopy[taskgroupshapeFrom]; ok {
		taskgroupshapeTo = _taskgroupshapeTo.(*TaskGroupShape)
		return
	}

	taskgroupshapeTo = new(TaskGroupShape)
	mapOrigCopy[taskgroupshapeFrom] = taskgroupshapeTo
	taskgroupshapeFrom.GongCopyBasicFields(taskgroupshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskgroupshapeFrom.TaskGroup != nil {
		taskgroupshapeTo.TaskGroup = GongCopyBranchTaskGroup(mapOrigCopy, taskgroupshapeFrom.TaskGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskInputShape(mapOrigCopy map[any]any, taskinputshapeFrom *TaskInputShape) (taskinputshapeTo *TaskInputShape) {

	// taskinputshapeFrom has already been copied
	if _taskinputshapeTo, ok := mapOrigCopy[taskinputshapeFrom]; ok {
		taskinputshapeTo = _taskinputshapeTo.(*TaskInputShape)
		return
	}

	taskinputshapeTo = new(TaskInputShape)
	mapOrigCopy[taskinputshapeFrom] = taskinputshapeTo
	taskinputshapeFrom.GongCopyBasicFields(taskinputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshapeFrom.Product != nil {
		taskinputshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, taskinputshapeFrom.Product)
	}
	if taskinputshapeFrom.Task != nil {
		taskinputshapeTo.Task = GongCopyBranchTask(mapOrigCopy, taskinputshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskOutputShape(mapOrigCopy map[any]any, taskoutputshapeFrom *TaskOutputShape) (taskoutputshapeTo *TaskOutputShape) {

	// taskoutputshapeFrom has already been copied
	if _taskoutputshapeTo, ok := mapOrigCopy[taskoutputshapeFrom]; ok {
		taskoutputshapeTo = _taskoutputshapeTo.(*TaskOutputShape)
		return
	}

	taskoutputshapeTo = new(TaskOutputShape)
	mapOrigCopy[taskoutputshapeFrom] = taskoutputshapeTo
	taskoutputshapeFrom.GongCopyBasicFields(taskoutputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshapeFrom.Task != nil {
		taskoutputshapeTo.Task = GongCopyBranchTask(mapOrigCopy, taskoutputshapeFrom.Task)
	}
	if taskoutputshapeFrom.Product != nil {
		taskoutputshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, taskoutputshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskPredecessorShape(mapOrigCopy map[any]any, taskpredecessorshapeFrom *TaskPredecessorShape) (taskpredecessorshapeTo *TaskPredecessorShape) {

	// taskpredecessorshapeFrom has already been copied
	if _taskpredecessorshapeTo, ok := mapOrigCopy[taskpredecessorshapeFrom]; ok {
		taskpredecessorshapeTo = _taskpredecessorshapeTo.(*TaskPredecessorShape)
		return
	}

	taskpredecessorshapeTo = new(TaskPredecessorShape)
	mapOrigCopy[taskpredecessorshapeFrom] = taskpredecessorshapeTo
	taskpredecessorshapeFrom.GongCopyBasicFields(taskpredecessorshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskpredecessorshapeFrom.Predecessor != nil {
		taskpredecessorshapeTo.Predecessor = GongCopyBranchTask(mapOrigCopy, taskpredecessorshapeFrom.Predecessor)
	}
	if taskpredecessorshapeFrom.Task != nil {
		taskpredecessorshapeTo.Task = GongCopyBranchTask(mapOrigCopy, taskpredecessorshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskShape(mapOrigCopy map[any]any, taskshapeFrom *TaskShape) (taskshapeTo *TaskShape) {

	// taskshapeFrom has already been copied
	if _taskshapeTo, ok := mapOrigCopy[taskshapeFrom]; ok {
		taskshapeTo = _taskshapeTo.(*TaskShape)
		return
	}

	taskshapeTo = new(TaskShape)
	mapOrigCopy[taskshapeFrom] = taskshapeTo
	taskshapeFrom.GongCopyBasicFields(taskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskshapeFrom.Task != nil {
		taskshapeTo.Task = GongCopyBranchTask(mapOrigCopy, taskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (diagram *Diagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDiagram(diagram)
}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !stage.IsStaged(diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagram.Product_Shapes {
		stage.UnstageBranch(_productshape)
	}
	for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
		stage.UnstageBranch(_product)
	}
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		stage.UnstageBranch(_productcompositionshape)
	}
	for _, _productreferenceshape := range diagram.ProductReference_Shapes {
		stage.UnstageBranch(_productreferenceshape)
	}
	for _, _taskshape := range diagram.Task_Shapes {
		stage.UnstageBranch(_taskshape)
	}
	for _, _task := range diagram.TasksWhoseNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _task := range diagram.TasksWhosePredecessorNodeIsExpanded {
		stage.UnstageBranch(_task)
	}
	for _, _taskgroupshape := range diagram.TaskGroupShapes {
		stage.UnstageBranch(_taskgroupshape)
	}
	for _, _taskgroup := range diagram.TaskGroupsWhoseNodeIsExpanded {
		stage.UnstageBranch(_taskgroup)
	}
	for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
		stage.UnstageBranch(_taskcompositionshape)
	}
	for _, _taskinputshape := range diagram.TaskInputShapes {
		stage.UnstageBranch(_taskinputshape)
	}
	for _, _taskoutputshape := range diagram.TaskOutputShapes {
		stage.UnstageBranch(_taskoutputshape)
	}
	for _, _taskpredecessorshape := range diagram.TaskPredecessorShapes {
		stage.UnstageBranch(_taskpredecessorshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		stage.UnstageBranch(_noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}
	for _, _noteproductshape := range diagram.NoteProductShapes {
		stage.UnstageBranch(_noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		stage.UnstageBranch(_notetaskshape)
	}
	for _, _noteresourceshape := range diagram.NoteResourceShapes {
		stage.UnstageBranch(_noteresourceshape)
	}
	for _, _resourceshape := range diagram.Resource_Shapes {
		stage.UnstageBranch(_resourceshape)
	}
	for _, _resource := range diagram.ResourcesWhoseNodeIsExpanded {
		stage.UnstageBranch(_resource)
	}
	for _, _resourcecompositionshape := range diagram.ResourceComposition_Shapes {
		stage.UnstageBranch(_resourcecompositionshape)
	}
	for _, _resourcetaskshape := range diagram.ResourceTaskShapes {
		stage.UnstageBranch(_resourcetaskshape)
	}

}

func (library *Library) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLibrary(library)
}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}
	for _, _product := range library.RootProducts {
		stage.UnstageBranch(_product)
	}
	for _, _task := range library.RootTasks {
		stage.UnstageBranch(_task)
	}
	for _, _taskgroup := range library.RootTaskGroups {
		stage.UnstageBranch(_taskgroup)
	}
	for _, _resource := range library.RootResources {
		stage.UnstageBranch(_resource)
	}
	for _, _note := range library.Notes {
		stage.UnstageBranch(_note)
	}
	for _, _diagram := range library.Diagrams {
		stage.UnstageBranch(_diagram)
	}

}

func (note *Note) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNote(note)
}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range note.Products {
		stage.UnstageBranch(_product)
	}
	for _, _task := range note.Tasks {
		stage.UnstageBranch(_task)
	}
	for _, _resource := range note.Resources {
		stage.UnstageBranch(_resource)
	}

}

func (noteproductshape *NoteProductShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteProductShape(noteproductshape)
}

func (stage *Stage) UnstageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if !stage.IsStaged(noteproductshape) {
		return
	}

	noteproductshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		stage.UnstageBranch(noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		stage.UnstageBranch(noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteresourceshape *NoteResourceShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteResourceShape(noteresourceshape)
}

func (stage *Stage) UnstageBranchNoteResourceShape(noteresourceshape *NoteResourceShape) {

	// check if instance is already staged
	if !stage.IsStaged(noteresourceshape) {
		return
	}

	noteresourceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteresourceshape.Note != nil {
		stage.UnstageBranch(noteresourceshape.Note)
	}
	if noteresourceshape.Resource != nil {
		stage.UnstageBranch(noteresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteShape(noteshape)
}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !stage.IsStaged(noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.UnstageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (notetaskshape *NoteTaskShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteTaskShape(notetaskshape)
}

func (stage *Stage) UnstageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if !stage.IsStaged(notetaskshape) {
		return
	}

	notetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		stage.UnstageBranch(notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		stage.UnstageBranch(notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (product *Product) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchProduct(product)
}

func (stage *Stage) UnstageBranchProduct(product *Product) {

	// check if instance is already staged
	if !stage.IsStaged(product) {
		return
	}

	product.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if product.ReferencedProduct != nil {
		stage.UnstageBranch(product.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range product.SubProducts {
		stage.UnstageBranch(_product)
	}

}

func (productcompositionshape *ProductCompositionShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchProductCompositionShape(productcompositionshape)
}

func (stage *Stage) UnstageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if !stage.IsStaged(productcompositionshape) {
		return
	}

	productcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		stage.UnstageBranch(productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (productreferenceshape *ProductReferenceShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchProductReferenceShape(productreferenceshape)
}

func (stage *Stage) UnstageBranchProductReferenceShape(productreferenceshape *ProductReferenceShape) {

	// check if instance is already staged
	if !stage.IsStaged(productreferenceshape) {
		return
	}

	productreferenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productreferenceshape.Product != nil {
		stage.UnstageBranch(productreferenceshape.Product)
	}
	if productreferenceshape.ReferencedProduct != nil {
		stage.UnstageBranch(productreferenceshape.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (productshape *ProductShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchProductShape(productshape)
}

func (stage *Stage) UnstageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if !stage.IsStaged(productshape) {
		return
	}

	productshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		stage.UnstageBranch(productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resource *Resource) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchResource(resource)
}

func (stage *Stage) UnstageBranchResource(resource *Resource) {

	// check if instance is already staged
	if !stage.IsStaged(resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resource.ReferencedResource != nil {
		stage.UnstageBranch(resource.ReferencedResource)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range resource.Tasks {
		stage.UnstageBranch(_task)
	}
	for _, _resource := range resource.SubResources {
		stage.UnstageBranch(_resource)
	}

}

func (resourcecompositionshape *ResourceCompositionShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchResourceCompositionShape(resourcecompositionshape)
}

func (stage *Stage) UnstageBranchResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape) {

	// check if instance is already staged
	if !stage.IsStaged(resourcecompositionshape) {
		return
	}

	resourcecompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcecompositionshape.Resource != nil {
		stage.UnstageBranch(resourcecompositionshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resourceshape *ResourceShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchResourceShape(resourceshape)
}

func (stage *Stage) UnstageBranchResourceShape(resourceshape *ResourceShape) {

	// check if instance is already staged
	if !stage.IsStaged(resourceshape) {
		return
	}

	resourceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourceshape.Resource != nil {
		stage.UnstageBranch(resourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resourcetaskshape *ResourceTaskShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchResourceTaskShape(resourcetaskshape)
}

func (stage *Stage) UnstageBranchResourceTaskShape(resourcetaskshape *ResourceTaskShape) {

	// check if instance is already staged
	if !stage.IsStaged(resourcetaskshape) {
		return
	}

	resourcetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcetaskshape.Resource != nil {
		stage.UnstageBranch(resourcetaskshape.Resource)
	}
	if resourcetaskshape.Task != nil {
		stage.UnstageBranch(resourcetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (task *Task) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTask(task)
}

func (stage *Stage) UnstageBranchTask(task *Task) {

	// check if instance is already staged
	if !stage.IsStaged(task) {
		return
	}

	task.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if task.ReferencedTask != nil {
		stage.UnstageBranch(task.ReferencedTask)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range task.Predecessors {
		stage.UnstageBranch(_task)
	}
	for _, _product := range task.Inputs {
		stage.UnstageBranch(_product)
	}
	for _, _product := range task.Outputs {
		stage.UnstageBranch(_product)
	}
	for _, _task := range task.SubTasks {
		stage.UnstageBranch(_task)
	}
	for _, _taskgroup := range task.TaskGroupsToDisplay {
		stage.UnstageBranch(_taskgroup)
	}

}

func (taskcompositionshape *TaskCompositionShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskCompositionShape(taskcompositionshape)
}

func (stage *Stage) UnstageBranchTaskCompositionShape(taskcompositionshape *TaskCompositionShape) {

	// check if instance is already staged
	if !stage.IsStaged(taskcompositionshape) {
		return
	}

	taskcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshape.Task != nil {
		stage.UnstageBranch(taskcompositionshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskgroup *TaskGroup) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskGroup(taskgroup)
}

func (stage *Stage) UnstageBranchTaskGroup(taskgroup *TaskGroup) {

	// check if instance is already staged
	if !stage.IsStaged(taskgroup) {
		return
	}

	taskgroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskgroup.Tasks {
		stage.UnstageBranch(_task)
	}

}

func (taskgroupshape *TaskGroupShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskGroupShape(taskgroupshape)
}

func (stage *Stage) UnstageBranchTaskGroupShape(taskgroupshape *TaskGroupShape) {

	// check if instance is already staged
	if !stage.IsStaged(taskgroupshape) {
		return
	}

	taskgroupshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskgroupshape.TaskGroup != nil {
		stage.UnstageBranch(taskgroupshape.TaskGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskinputshape *TaskInputShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskInputShape(taskinputshape)
}

func (stage *Stage) UnstageBranchTaskInputShape(taskinputshape *TaskInputShape) {

	// check if instance is already staged
	if !stage.IsStaged(taskinputshape) {
		return
	}

	taskinputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshape.Product != nil {
		stage.UnstageBranch(taskinputshape.Product)
	}
	if taskinputshape.Task != nil {
		stage.UnstageBranch(taskinputshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskoutputshape *TaskOutputShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskOutputShape(taskoutputshape)
}

func (stage *Stage) UnstageBranchTaskOutputShape(taskoutputshape *TaskOutputShape) {

	// check if instance is already staged
	if !stage.IsStaged(taskoutputshape) {
		return
	}

	taskoutputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshape.Task != nil {
		stage.UnstageBranch(taskoutputshape.Task)
	}
	if taskoutputshape.Product != nil {
		stage.UnstageBranch(taskoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskpredecessorshape *TaskPredecessorShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskPredecessorShape(taskpredecessorshape)
}

func (stage *Stage) UnstageBranchTaskPredecessorShape(taskpredecessorshape *TaskPredecessorShape) {

	// check if instance is already staged
	if !stage.IsStaged(taskpredecessorshape) {
		return
	}

	taskpredecessorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskpredecessorshape.Predecessor != nil {
		stage.UnstageBranch(taskpredecessorshape.Predecessor)
	}
	if taskpredecessorshape.Task != nil {
		stage.UnstageBranch(taskpredecessorshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (taskshape *TaskShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchTaskShape(taskshape)
}

func (stage *Stage) UnstageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if !stage.IsStaged(taskshape) {
		return
	}

	taskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		stage.UnstageBranch(taskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Product_Shapes = reference.Product_Shapes[:0]
	for _, _b := range instance.Product_Shapes {
		reference.Product_Shapes = append(reference.Product_Shapes, stage.ProductShapes_reference[_b])
	}
	reference.ProductsWhoseNodeIsExpanded = reference.ProductsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ProductsWhoseNodeIsExpanded {
		reference.ProductsWhoseNodeIsExpanded = append(reference.ProductsWhoseNodeIsExpanded, stage.Products_reference[_b])
	}
	reference.ProductComposition_Shapes = reference.ProductComposition_Shapes[:0]
	for _, _b := range instance.ProductComposition_Shapes {
		reference.ProductComposition_Shapes = append(reference.ProductComposition_Shapes, stage.ProductCompositionShapes_reference[_b])
	}
	reference.ProductReference_Shapes = reference.ProductReference_Shapes[:0]
	for _, _b := range instance.ProductReference_Shapes {
		reference.ProductReference_Shapes = append(reference.ProductReference_Shapes, stage.ProductReferenceShapes_reference[_b])
	}
	reference.Task_Shapes = reference.Task_Shapes[:0]
	for _, _b := range instance.Task_Shapes {
		reference.Task_Shapes = append(reference.Task_Shapes, stage.TaskShapes_reference[_b])
	}
	reference.TasksWhoseNodeIsExpanded = reference.TasksWhoseNodeIsExpanded[:0]
	for _, _b := range instance.TasksWhoseNodeIsExpanded {
		reference.TasksWhoseNodeIsExpanded = append(reference.TasksWhoseNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TasksWhoseInputNodeIsExpanded = reference.TasksWhoseInputNodeIsExpanded[:0]
	for _, _b := range instance.TasksWhoseInputNodeIsExpanded {
		reference.TasksWhoseInputNodeIsExpanded = append(reference.TasksWhoseInputNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TasksWhoseOutputNodeIsExpanded = reference.TasksWhoseOutputNodeIsExpanded[:0]
	for _, _b := range instance.TasksWhoseOutputNodeIsExpanded {
		reference.TasksWhoseOutputNodeIsExpanded = append(reference.TasksWhoseOutputNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TasksWhosePredecessorNodeIsExpanded = reference.TasksWhosePredecessorNodeIsExpanded[:0]
	for _, _b := range instance.TasksWhosePredecessorNodeIsExpanded {
		reference.TasksWhosePredecessorNodeIsExpanded = append(reference.TasksWhosePredecessorNodeIsExpanded, stage.Tasks_reference[_b])
	}
	reference.TaskGroupShapes = reference.TaskGroupShapes[:0]
	for _, _b := range instance.TaskGroupShapes {
		reference.TaskGroupShapes = append(reference.TaskGroupShapes, stage.TaskGroupShapes_reference[_b])
	}
	reference.TaskGroupsWhoseNodeIsExpanded = reference.TaskGroupsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.TaskGroupsWhoseNodeIsExpanded {
		reference.TaskGroupsWhoseNodeIsExpanded = append(reference.TaskGroupsWhoseNodeIsExpanded, stage.TaskGroups_reference[_b])
	}
	reference.TaskComposition_Shapes = reference.TaskComposition_Shapes[:0]
	for _, _b := range instance.TaskComposition_Shapes {
		reference.TaskComposition_Shapes = append(reference.TaskComposition_Shapes, stage.TaskCompositionShapes_reference[_b])
	}
	reference.TaskInputShapes = reference.TaskInputShapes[:0]
	for _, _b := range instance.TaskInputShapes {
		reference.TaskInputShapes = append(reference.TaskInputShapes, stage.TaskInputShapes_reference[_b])
	}
	reference.TaskOutputShapes = reference.TaskOutputShapes[:0]
	for _, _b := range instance.TaskOutputShapes {
		reference.TaskOutputShapes = append(reference.TaskOutputShapes, stage.TaskOutputShapes_reference[_b])
	}
	reference.TaskPredecessorShapes = reference.TaskPredecessorShapes[:0]
	for _, _b := range instance.TaskPredecessorShapes {
		reference.TaskPredecessorShapes = append(reference.TaskPredecessorShapes, stage.TaskPredecessorShapes_reference[_b])
	}
	reference.Note_Shapes = reference.Note_Shapes[:0]
	for _, _b := range instance.Note_Shapes {
		reference.Note_Shapes = append(reference.Note_Shapes, stage.NoteShapes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
	reference.NoteProductShapes = reference.NoteProductShapes[:0]
	for _, _b := range instance.NoteProductShapes {
		reference.NoteProductShapes = append(reference.NoteProductShapes, stage.NoteProductShapes_reference[_b])
	}
	reference.NoteTaskShapes = reference.NoteTaskShapes[:0]
	for _, _b := range instance.NoteTaskShapes {
		reference.NoteTaskShapes = append(reference.NoteTaskShapes, stage.NoteTaskShapes_reference[_b])
	}
	reference.NoteResourceShapes = reference.NoteResourceShapes[:0]
	for _, _b := range instance.NoteResourceShapes {
		reference.NoteResourceShapes = append(reference.NoteResourceShapes, stage.NoteResourceShapes_reference[_b])
	}
	reference.Resource_Shapes = reference.Resource_Shapes[:0]
	for _, _b := range instance.Resource_Shapes {
		reference.Resource_Shapes = append(reference.Resource_Shapes, stage.ResourceShapes_reference[_b])
	}
	reference.ResourcesWhoseNodeIsExpanded = reference.ResourcesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ResourcesWhoseNodeIsExpanded {
		reference.ResourcesWhoseNodeIsExpanded = append(reference.ResourcesWhoseNodeIsExpanded, stage.Resources_reference[_b])
	}
	reference.ResourceComposition_Shapes = reference.ResourceComposition_Shapes[:0]
	for _, _b := range instance.ResourceComposition_Shapes {
		reference.ResourceComposition_Shapes = append(reference.ResourceComposition_Shapes, stage.ResourceCompositionShapes_reference[_b])
	}
	reference.ResourceTaskShapes = reference.ResourceTaskShapes[:0]
	for _, _b := range instance.ResourceTaskShapes {
		reference.ResourceTaskShapes = append(reference.ResourceTaskShapes, stage.ResourceTaskShapes_reference[_b])
	}
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.RootProducts = reference.RootProducts[:0]
	for _, _b := range instance.RootProducts {
		reference.RootProducts = append(reference.RootProducts, stage.Products_reference[_b])
	}
	reference.RootTasks = reference.RootTasks[:0]
	for _, _b := range instance.RootTasks {
		reference.RootTasks = append(reference.RootTasks, stage.Tasks_reference[_b])
	}
	reference.RootTaskGroups = reference.RootTaskGroups[:0]
	for _, _b := range instance.RootTaskGroups {
		reference.RootTaskGroups = append(reference.RootTaskGroups, stage.TaskGroups_reference[_b])
	}
	reference.RootResources = reference.RootResources[:0]
	for _, _b := range instance.RootResources {
		reference.RootResources = append(reference.RootResources, stage.Resources_reference[_b])
	}
	reference.Notes = reference.Notes[:0]
	for _, _b := range instance.Notes {
		reference.Notes = append(reference.Notes, stage.Notes_reference[_b])
	}
	reference.Diagrams = reference.Diagrams[:0]
	for _, _b := range instance.Diagrams {
		reference.Diagrams = append(reference.Diagrams, stage.Diagrams_reference[_b])
	}
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Products = reference.Products[:0]
	for _, _b := range instance.Products {
		reference.Products = append(reference.Products, stage.Products_reference[_b])
	}
	reference.Tasks = reference.Tasks[:0]
	for _, _b := range instance.Tasks {
		reference.Tasks = append(reference.Tasks, stage.Tasks_reference[_b])
	}
	reference.Resources = reference.Resources[:0]
	for _, _b := range instance.Resources {
		reference.Resources = append(reference.Resources, stage.Resources_reference[_b])
	}
}

func (reference *NoteProductShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteProductShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Product != nil {
		reference.Product = stage.Products_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteResourceShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteResourceShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Resource != nil {
		reference.Resource = stage.Resources_reference[instance.Resource]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteTaskShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

func (reference *Product) GongReconstructPointersFromReferences(stage *Stage, instance *Product) {
	// insertion point for pointers field
	if instance.ReferencedProduct != nil {
		reference.ReferencedProduct = stage.Products_reference[instance.ReferencedProduct]
	}
	// insertion point for slice of pointers field
	reference.SubProducts = reference.SubProducts[:0]
	for _, _b := range instance.SubProducts {
		reference.SubProducts = append(reference.SubProducts, stage.Products_reference[_b])
	}
}

func (reference *ProductCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductCompositionShape) {
	// insertion point for pointers field
	if instance.Product != nil {
		reference.Product = stage.Products_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *ProductReferenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductReferenceShape) {
	// insertion point for pointers field
	if instance.Product != nil {
		reference.Product = stage.Products_reference[instance.Product]
	}
	if instance.ReferencedProduct != nil {
		reference.ReferencedProduct = stage.Products_reference[instance.ReferencedProduct]
	}
	// insertion point for slice of pointers field
}

func (reference *ProductShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductShape) {
	// insertion point for pointers field
	if instance.Product != nil {
		reference.Product = stage.Products_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *Resource) GongReconstructPointersFromReferences(stage *Stage, instance *Resource) {
	// insertion point for pointers field
	if instance.ReferencedResource != nil {
		reference.ReferencedResource = stage.Resources_reference[instance.ReferencedResource]
	}
	// insertion point for slice of pointers field
	reference.Tasks = reference.Tasks[:0]
	for _, _b := range instance.Tasks {
		reference.Tasks = append(reference.Tasks, stage.Tasks_reference[_b])
	}
	reference.SubResources = reference.SubResources[:0]
	for _, _b := range instance.SubResources {
		reference.SubResources = append(reference.SubResources, stage.Resources_reference[_b])
	}
}

func (reference *ResourceCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ResourceCompositionShape) {
	// insertion point for pointers field
	if instance.Resource != nil {
		reference.Resource = stage.Resources_reference[instance.Resource]
	}
	// insertion point for slice of pointers field
}

func (reference *ResourceShape) GongReconstructPointersFromReferences(stage *Stage, instance *ResourceShape) {
	// insertion point for pointers field
	if instance.Resource != nil {
		reference.Resource = stage.Resources_reference[instance.Resource]
	}
	// insertion point for slice of pointers field
}

func (reference *ResourceTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *ResourceTaskShape) {
	// insertion point for pointers field
	if instance.Resource != nil {
		reference.Resource = stage.Resources_reference[instance.Resource]
	}
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

func (reference *Task) GongReconstructPointersFromReferences(stage *Stage, instance *Task) {
	// insertion point for pointers field
	if instance.ReferencedTask != nil {
		reference.ReferencedTask = stage.Tasks_reference[instance.ReferencedTask]
	}
	// insertion point for slice of pointers field
	reference.Predecessors = reference.Predecessors[:0]
	for _, _b := range instance.Predecessors {
		reference.Predecessors = append(reference.Predecessors, stage.Tasks_reference[_b])
	}
	reference.Inputs = reference.Inputs[:0]
	for _, _b := range instance.Inputs {
		reference.Inputs = append(reference.Inputs, stage.Products_reference[_b])
	}
	reference.Outputs = reference.Outputs[:0]
	for _, _b := range instance.Outputs {
		reference.Outputs = append(reference.Outputs, stage.Products_reference[_b])
	}
	reference.SubTasks = reference.SubTasks[:0]
	for _, _b := range instance.SubTasks {
		reference.SubTasks = append(reference.SubTasks, stage.Tasks_reference[_b])
	}
	reference.TaskGroupsToDisplay = reference.TaskGroupsToDisplay[:0]
	for _, _b := range instance.TaskGroupsToDisplay {
		reference.TaskGroupsToDisplay = append(reference.TaskGroupsToDisplay, stage.TaskGroups_reference[_b])
	}
}

func (reference *TaskCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskCompositionShape) {
	// insertion point for pointers field
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

func (reference *TaskGroup) GongReconstructPointersFromReferences(stage *Stage, instance *TaskGroup) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Tasks = reference.Tasks[:0]
	for _, _b := range instance.Tasks {
		reference.Tasks = append(reference.Tasks, stage.Tasks_reference[_b])
	}
}

func (reference *TaskGroupShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskGroupShape) {
	// insertion point for pointers field
	if instance.TaskGroup != nil {
		reference.TaskGroup = stage.TaskGroups_reference[instance.TaskGroup]
	}
	// insertion point for slice of pointers field
}

func (reference *TaskInputShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskInputShape) {
	// insertion point for pointers field
	if instance.Product != nil {
		reference.Product = stage.Products_reference[instance.Product]
	}
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

func (reference *TaskOutputShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskOutputShape) {
	// insertion point for pointers field
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	if instance.Product != nil {
		reference.Product = stage.Products_reference[instance.Product]
	}
	// insertion point for slice of pointers field
}

func (reference *TaskPredecessorShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskPredecessorShape) {
	// insertion point for pointers field
	if instance.Predecessor != nil {
		reference.Predecessor = stage.Tasks_reference[instance.Predecessor]
	}
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

func (reference *TaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskShape) {
	// insertion point for pointers field
	if instance.Task != nil {
		reference.Task = stage.Tasks_reference[instance.Task]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Product_Shapes []*ProductShape
	for _, _reference := range reference.Product_Shapes {
		if _instance, ok := stage.ProductShapes_instance[_reference]; ok {
			_Product_Shapes = append(_Product_Shapes, _instance)
		}
	}
	reference.Product_Shapes = _Product_Shapes
	var _ProductsWhoseNodeIsExpanded []*Product
	for _, _reference := range reference.ProductsWhoseNodeIsExpanded {
		if _instance, ok := stage.Products_instance[_reference]; ok {
			_ProductsWhoseNodeIsExpanded = append(_ProductsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ProductsWhoseNodeIsExpanded = _ProductsWhoseNodeIsExpanded
	var _ProductComposition_Shapes []*ProductCompositionShape
	for _, _reference := range reference.ProductComposition_Shapes {
		if _instance, ok := stage.ProductCompositionShapes_instance[_reference]; ok {
			_ProductComposition_Shapes = append(_ProductComposition_Shapes, _instance)
		}
	}
	reference.ProductComposition_Shapes = _ProductComposition_Shapes
	var _ProductReference_Shapes []*ProductReferenceShape
	for _, _reference := range reference.ProductReference_Shapes {
		if _instance, ok := stage.ProductReferenceShapes_instance[_reference]; ok {
			_ProductReference_Shapes = append(_ProductReference_Shapes, _instance)
		}
	}
	reference.ProductReference_Shapes = _ProductReference_Shapes
	var _Task_Shapes []*TaskShape
	for _, _reference := range reference.Task_Shapes {
		if _instance, ok := stage.TaskShapes_instance[_reference]; ok {
			_Task_Shapes = append(_Task_Shapes, _instance)
		}
	}
	reference.Task_Shapes = _Task_Shapes
	var _TasksWhoseNodeIsExpanded []*Task
	for _, _reference := range reference.TasksWhoseNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TasksWhoseNodeIsExpanded = append(_TasksWhoseNodeIsExpanded, _instance)
		}
	}
	reference.TasksWhoseNodeIsExpanded = _TasksWhoseNodeIsExpanded
	var _TasksWhoseInputNodeIsExpanded []*Task
	for _, _reference := range reference.TasksWhoseInputNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TasksWhoseInputNodeIsExpanded = append(_TasksWhoseInputNodeIsExpanded, _instance)
		}
	}
	reference.TasksWhoseInputNodeIsExpanded = _TasksWhoseInputNodeIsExpanded
	var _TasksWhoseOutputNodeIsExpanded []*Task
	for _, _reference := range reference.TasksWhoseOutputNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TasksWhoseOutputNodeIsExpanded = append(_TasksWhoseOutputNodeIsExpanded, _instance)
		}
	}
	reference.TasksWhoseOutputNodeIsExpanded = _TasksWhoseOutputNodeIsExpanded
	var _TasksWhosePredecessorNodeIsExpanded []*Task
	for _, _reference := range reference.TasksWhosePredecessorNodeIsExpanded {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_TasksWhosePredecessorNodeIsExpanded = append(_TasksWhosePredecessorNodeIsExpanded, _instance)
		}
	}
	reference.TasksWhosePredecessorNodeIsExpanded = _TasksWhosePredecessorNodeIsExpanded
	var _TaskGroupShapes []*TaskGroupShape
	for _, _reference := range reference.TaskGroupShapes {
		if _instance, ok := stage.TaskGroupShapes_instance[_reference]; ok {
			_TaskGroupShapes = append(_TaskGroupShapes, _instance)
		}
	}
	reference.TaskGroupShapes = _TaskGroupShapes
	var _TaskGroupsWhoseNodeIsExpanded []*TaskGroup
	for _, _reference := range reference.TaskGroupsWhoseNodeIsExpanded {
		if _instance, ok := stage.TaskGroups_instance[_reference]; ok {
			_TaskGroupsWhoseNodeIsExpanded = append(_TaskGroupsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.TaskGroupsWhoseNodeIsExpanded = _TaskGroupsWhoseNodeIsExpanded
	var _TaskComposition_Shapes []*TaskCompositionShape
	for _, _reference := range reference.TaskComposition_Shapes {
		if _instance, ok := stage.TaskCompositionShapes_instance[_reference]; ok {
			_TaskComposition_Shapes = append(_TaskComposition_Shapes, _instance)
		}
	}
	reference.TaskComposition_Shapes = _TaskComposition_Shapes
	var _TaskInputShapes []*TaskInputShape
	for _, _reference := range reference.TaskInputShapes {
		if _instance, ok := stage.TaskInputShapes_instance[_reference]; ok {
			_TaskInputShapes = append(_TaskInputShapes, _instance)
		}
	}
	reference.TaskInputShapes = _TaskInputShapes
	var _TaskOutputShapes []*TaskOutputShape
	for _, _reference := range reference.TaskOutputShapes {
		if _instance, ok := stage.TaskOutputShapes_instance[_reference]; ok {
			_TaskOutputShapes = append(_TaskOutputShapes, _instance)
		}
	}
	reference.TaskOutputShapes = _TaskOutputShapes
	var _TaskPredecessorShapes []*TaskPredecessorShape
	for _, _reference := range reference.TaskPredecessorShapes {
		if _instance, ok := stage.TaskPredecessorShapes_instance[_reference]; ok {
			_TaskPredecessorShapes = append(_TaskPredecessorShapes, _instance)
		}
	}
	reference.TaskPredecessorShapes = _TaskPredecessorShapes
	var _Note_Shapes []*NoteShape
	for _, _reference := range reference.Note_Shapes {
		if _instance, ok := stage.NoteShapes_instance[_reference]; ok {
			_Note_Shapes = append(_Note_Shapes, _instance)
		}
	}
	reference.Note_Shapes = _Note_Shapes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
	var _NoteProductShapes []*NoteProductShape
	for _, _reference := range reference.NoteProductShapes {
		if _instance, ok := stage.NoteProductShapes_instance[_reference]; ok {
			_NoteProductShapes = append(_NoteProductShapes, _instance)
		}
	}
	reference.NoteProductShapes = _NoteProductShapes
	var _NoteTaskShapes []*NoteTaskShape
	for _, _reference := range reference.NoteTaskShapes {
		if _instance, ok := stage.NoteTaskShapes_instance[_reference]; ok {
			_NoteTaskShapes = append(_NoteTaskShapes, _instance)
		}
	}
	reference.NoteTaskShapes = _NoteTaskShapes
	var _NoteResourceShapes []*NoteResourceShape
	for _, _reference := range reference.NoteResourceShapes {
		if _instance, ok := stage.NoteResourceShapes_instance[_reference]; ok {
			_NoteResourceShapes = append(_NoteResourceShapes, _instance)
		}
	}
	reference.NoteResourceShapes = _NoteResourceShapes
	var _Resource_Shapes []*ResourceShape
	for _, _reference := range reference.Resource_Shapes {
		if _instance, ok := stage.ResourceShapes_instance[_reference]; ok {
			_Resource_Shapes = append(_Resource_Shapes, _instance)
		}
	}
	reference.Resource_Shapes = _Resource_Shapes
	var _ResourcesWhoseNodeIsExpanded []*Resource
	for _, _reference := range reference.ResourcesWhoseNodeIsExpanded {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_ResourcesWhoseNodeIsExpanded = append(_ResourcesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ResourcesWhoseNodeIsExpanded = _ResourcesWhoseNodeIsExpanded
	var _ResourceComposition_Shapes []*ResourceCompositionShape
	for _, _reference := range reference.ResourceComposition_Shapes {
		if _instance, ok := stage.ResourceCompositionShapes_instance[_reference]; ok {
			_ResourceComposition_Shapes = append(_ResourceComposition_Shapes, _instance)
		}
	}
	reference.ResourceComposition_Shapes = _ResourceComposition_Shapes
	var _ResourceTaskShapes []*ResourceTaskShape
	for _, _reference := range reference.ResourceTaskShapes {
		if _instance, ok := stage.ResourceTaskShapes_instance[_reference]; ok {
			_ResourceTaskShapes = append(_ResourceTaskShapes, _instance)
		}
	}
	reference.ResourceTaskShapes = _ResourceTaskShapes
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _RootProducts []*Product
	for _, _reference := range reference.RootProducts {
		if _instance, ok := stage.Products_instance[_reference]; ok {
			_RootProducts = append(_RootProducts, _instance)
		}
	}
	reference.RootProducts = _RootProducts
	var _RootTasks []*Task
	for _, _reference := range reference.RootTasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_RootTasks = append(_RootTasks, _instance)
		}
	}
	reference.RootTasks = _RootTasks
	var _RootTaskGroups []*TaskGroup
	for _, _reference := range reference.RootTaskGroups {
		if _instance, ok := stage.TaskGroups_instance[_reference]; ok {
			_RootTaskGroups = append(_RootTaskGroups, _instance)
		}
	}
	reference.RootTaskGroups = _RootTaskGroups
	var _RootResources []*Resource
	for _, _reference := range reference.RootResources {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_RootResources = append(_RootResources, _instance)
		}
	}
	reference.RootResources = _RootResources
	var _Notes []*Note
	for _, _reference := range reference.Notes {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_Notes = append(_Notes, _instance)
		}
	}
	reference.Notes = _Notes
	var _Diagrams []*Diagram
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Products []*Product
	for _, _reference := range reference.Products {
		if _instance, ok := stage.Products_instance[_reference]; ok {
			_Products = append(_Products, _instance)
		}
	}
	reference.Products = _Products
	var _Tasks []*Task
	for _, _reference := range reference.Tasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_Tasks = append(_Tasks, _instance)
		}
	}
	reference.Tasks = _Tasks
	var _Resources []*Resource
	for _, _reference := range reference.Resources {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_Resources = append(_Resources, _instance)
		}
	}
	reference.Resources = _Resources
}

func (reference *NoteProductShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Resource; _reference != nil {
		reference.Resource = nil
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			reference.Resource = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Product) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ReferencedProduct; _reference != nil {
		reference.ReferencedProduct = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.ReferencedProduct = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _SubProducts []*Product
	for _, _reference := range reference.SubProducts {
		if _instance, ok := stage.Products_instance[_reference]; ok {
			_SubProducts = append(_SubProducts, _instance)
		}
	}
	reference.SubProducts = _SubProducts
}

func (reference *ProductCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ProductReferenceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	if _reference := reference.ReferencedProduct; _reference != nil {
		reference.ReferencedProduct = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.ReferencedProduct = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ProductShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Resource) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ReferencedResource; _reference != nil {
		reference.ReferencedResource = nil
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			reference.ReferencedResource = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Tasks []*Task
	for _, _reference := range reference.Tasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_Tasks = append(_Tasks, _instance)
		}
	}
	reference.Tasks = _Tasks
	var _SubResources []*Resource
	for _, _reference := range reference.SubResources {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_SubResources = append(_SubResources, _instance)
		}
	}
	reference.SubResources = _SubResources
}

func (reference *ResourceCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Resource; _reference != nil {
		reference.Resource = nil
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			reference.Resource = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Resource; _reference != nil {
		reference.Resource = nil
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			reference.Resource = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ResourceTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Resource; _reference != nil {
		reference.Resource = nil
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			reference.Resource = _instance
		}
	}
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Task) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ReferencedTask; _reference != nil {
		reference.ReferencedTask = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.ReferencedTask = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Predecessors []*Task
	for _, _reference := range reference.Predecessors {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_Predecessors = append(_Predecessors, _instance)
		}
	}
	reference.Predecessors = _Predecessors
	var _Inputs []*Product
	for _, _reference := range reference.Inputs {
		if _instance, ok := stage.Products_instance[_reference]; ok {
			_Inputs = append(_Inputs, _instance)
		}
	}
	reference.Inputs = _Inputs
	var _Outputs []*Product
	for _, _reference := range reference.Outputs {
		if _instance, ok := stage.Products_instance[_reference]; ok {
			_Outputs = append(_Outputs, _instance)
		}
	}
	reference.Outputs = _Outputs
	var _SubTasks []*Task
	for _, _reference := range reference.SubTasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_SubTasks = append(_SubTasks, _instance)
		}
	}
	reference.SubTasks = _SubTasks
	var _TaskGroupsToDisplay []*TaskGroup
	for _, _reference := range reference.TaskGroupsToDisplay {
		if _instance, ok := stage.TaskGroups_instance[_reference]; ok {
			_TaskGroupsToDisplay = append(_TaskGroupsToDisplay, _instance)
		}
	}
	reference.TaskGroupsToDisplay = _TaskGroupsToDisplay
}

func (reference *TaskCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TaskGroup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Tasks []*Task
	for _, _reference := range reference.Tasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_Tasks = append(_Tasks, _instance)
		}
	}
	reference.Tasks = _Tasks
}

func (reference *TaskGroupShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.TaskGroup; _reference != nil {
		reference.TaskGroup = nil
		if _instance, ok := stage.TaskGroups_instance[_reference]; ok {
			reference.TaskGroup = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TaskInputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TaskOutputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	if _reference := reference.Product; _reference != nil {
		reference.Product = nil
		if _instance, ok := stage.Products_instance[_reference]; ok {
			reference.Product = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TaskPredecessorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Predecessor; _reference != nil {
		reference.Predecessor = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Predecessor = _instance
		}
	}
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *TaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Task; _reference != nil {
		reference.Task = nil
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			reference.Task = _instance
		}
	}
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagram *Diagram) GongDiff(stage *Stage, diagramOther *Diagram) (diffs []string) {
	// insertion point for field diffs
	if diagram.Name != diagramOther.Name {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Name"))
	}
	if diagram.DefaultBoxWidth != diagramOther.DefaultBoxWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagram.DefaultBoxHeigth != diagramOther.DefaultBoxHeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DefaultBoxHeigth"))
	}
	if diagram.DateFormat != diagramOther.DateFormat {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DateFormat"))
	}
	if diagram.Width != diagramOther.Width {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Width"))
	}
	if diagram.Height != diagramOther.Height {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Height"))
	}
	if diagram.IsTimeDiagram != diagramOther.IsTimeDiagram {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsTimeDiagram"))
	}
	if diagram.ComputedStart != diagramOther.ComputedStart {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ComputedStart"))
	}
	if diagram.ComputedEnd != diagramOther.ComputedEnd {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ComputedEnd"))
	}
	if diagram.ComputedDuration != diagramOther.ComputedDuration {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ComputedDuration"))
	}
	if diagram.DrawVerticalTimeLines != diagramOther.DrawVerticalTimeLines {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DrawVerticalTimeLines"))
	}
	if diagram.HideWeekendsPeriod != diagramOther.HideWeekendsPeriod {
		diffs = append(diffs, diagram.GongMarshallField(stage, "HideWeekendsPeriod"))
	}
	if diagram.UseManualStartAndEndDates != diagramOther.UseManualStartAndEndDates {
		diffs = append(diffs, diagram.GongMarshallField(stage, "UseManualStartAndEndDates"))
	}
	if diagram.ManualStart != diagramOther.ManualStart {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ManualStart"))
	}
	if diagram.ManualEnd != diagramOther.ManualEnd {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ManualEnd"))
	}
	if diagram.TimeStep != diagramOther.TimeStep {
		diffs = append(diffs, diagram.GongMarshallField(stage, "TimeStep"))
	}
	if diagram.TimeStepScale != diagramOther.TimeStepScale {
		diffs = append(diffs, diagram.GongMarshallField(stage, "TimeStepScale"))
	}
	if diagram.LaneHeight != diagramOther.LaneHeight {
		diffs = append(diffs, diagram.GongMarshallField(stage, "LaneHeight"))
	}
	if diagram.RatioBarToLaneHeight != diagramOther.RatioBarToLaneHeight {
		diffs = append(diffs, diagram.GongMarshallField(stage, "RatioBarToLaneHeight"))
	}
	if diagram.YTopMargin != diagramOther.YTopMargin {
		diffs = append(diffs, diagram.GongMarshallField(stage, "YTopMargin"))
	}
	if diagram.XLeftText != diagramOther.XLeftText {
		diffs = append(diffs, diagram.GongMarshallField(stage, "XLeftText"))
	}
	if diagram.TextHeight != diagramOther.TextHeight {
		diffs = append(diffs, diagram.GongMarshallField(stage, "TextHeight"))
	}
	if diagram.XLeftLanes != diagramOther.XLeftLanes {
		diffs = append(diffs, diagram.GongMarshallField(stage, "XLeftLanes"))
	}
	if diagram.XRightMargin != diagramOther.XRightMargin {
		diffs = append(diffs, diagram.GongMarshallField(stage, "XRightMargin"))
	}
	if diagram.ArrowLengthToTheRightOfStartBar != diagramOther.ArrowLengthToTheRightOfStartBar {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArrowLengthToTheRightOfStartBar"))
	}
	if diagram.ArrowTipLenght != diagramOther.ArrowTipLenght {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArrowTipLenght"))
	}
	if diagram.TimeLine_Color != diagramOther.TimeLine_Color {
		diffs = append(diffs, diagram.GongMarshallField(stage, "TimeLine_Color"))
	}
	if diagram.TimeLine_FillOpacity != diagramOther.TimeLine_FillOpacity {
		diffs = append(diffs, diagram.GongMarshallField(stage, "TimeLine_FillOpacity"))
	}
	if diagram.TimeLine_Stroke != diagramOther.TimeLine_Stroke {
		diffs = append(diffs, diagram.GongMarshallField(stage, "TimeLine_Stroke"))
	}
	if diagram.TimeLine_StrokeWidth != diagramOther.TimeLine_StrokeWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "TimeLine_StrokeWidth"))
	}
	if diagram.Group_Stroke != diagramOther.Group_Stroke {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Group_Stroke"))
	}
	if diagram.Group_StrokeWidth != diagramOther.Group_StrokeWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Group_StrokeWidth"))
	}
	if diagram.Group_StrokeDashArray != diagramOther.Group_StrokeDashArray {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Group_StrokeDashArray"))
	}
	if diagram.DateYOffset != diagramOther.DateYOffset {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DateYOffset"))
	}
	if diagram.AlignOnStartEndOnYearStart != diagramOther.AlignOnStartEndOnYearStart {
		diffs = append(diffs, diagram.GongMarshallField(stage, "AlignOnStartEndOnYearStart"))
	}
	if diagram.ComputedPrefix != diagramOther.ComputedPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagram.IsExpanded != diagramOther.IsExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsExpanded"))
	}
	if diagram.IsChecked != diagramOther.IsChecked {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsChecked"))
	}
	if diagram.IsEditable_ != diagramOther.IsEditable_ {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsEditable_"))
	}
	if diagram.IsShowPrefix != diagramOther.IsShowPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsShowPrefix"))
	}
	if diagram.IsInAutoLayoutMode != diagramOther.IsInAutoLayoutMode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsInAutoLayoutMode"))
	}
	Product_ShapesDifferent := false
	if len(diagram.Product_Shapes) != len(diagramOther.Product_Shapes) {
		Product_ShapesDifferent = true
	} else {
		for i := range diagram.Product_Shapes {
			if (diagram.Product_Shapes[i] == nil) != (diagramOther.Product_Shapes[i] == nil) {
				Product_ShapesDifferent = true
				break
			} else if diagram.Product_Shapes[i] != nil && diagramOther.Product_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Product_Shapes[i] != diagramOther.Product_Shapes[i] {
					Product_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Product_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"Product_Shapes",
			len(diagramOther.Product_Shapes),
			len(diagram.Product_Shapes),
			func(i, j int) bool {
				return diagramOther.Product_Shapes[i] == diagram.Product_Shapes[j]
			},
			func(j int) string {
				return diagram.Product_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ProductsWhoseNodeIsExpandedDifferent := false
	if len(diagram.ProductsWhoseNodeIsExpanded) != len(diagramOther.ProductsWhoseNodeIsExpanded) {
		ProductsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ProductsWhoseNodeIsExpanded {
			if (diagram.ProductsWhoseNodeIsExpanded[i] == nil) != (diagramOther.ProductsWhoseNodeIsExpanded[i] == nil) {
				ProductsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ProductsWhoseNodeIsExpanded[i] != nil && diagramOther.ProductsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ProductsWhoseNodeIsExpanded[i] != diagramOther.ProductsWhoseNodeIsExpanded[i] {
					ProductsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ProductsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"ProductsWhoseNodeIsExpanded",
			len(diagramOther.ProductsWhoseNodeIsExpanded),
			len(diagram.ProductsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.ProductsWhoseNodeIsExpanded[i] == diagram.ProductsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.ProductsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagram.IsPBSNodeExpanded != diagramOther.IsPBSNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsPBSNodeExpanded"))
	}
	ProductComposition_ShapesDifferent := false
	if len(diagram.ProductComposition_Shapes) != len(diagramOther.ProductComposition_Shapes) {
		ProductComposition_ShapesDifferent = true
	} else {
		for i := range diagram.ProductComposition_Shapes {
			if (diagram.ProductComposition_Shapes[i] == nil) != (diagramOther.ProductComposition_Shapes[i] == nil) {
				ProductComposition_ShapesDifferent = true
				break
			} else if diagram.ProductComposition_Shapes[i] != nil && diagramOther.ProductComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ProductComposition_Shapes[i] != diagramOther.ProductComposition_Shapes[i] {
					ProductComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ProductComposition_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"ProductComposition_Shapes",
			len(diagramOther.ProductComposition_Shapes),
			len(diagram.ProductComposition_Shapes),
			func(i, j int) bool {
				return diagramOther.ProductComposition_Shapes[i] == diagram.ProductComposition_Shapes[j]
			},
			func(j int) string {
				return diagram.ProductComposition_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ProductReference_ShapesDifferent := false
	if len(diagram.ProductReference_Shapes) != len(diagramOther.ProductReference_Shapes) {
		ProductReference_ShapesDifferent = true
	} else {
		for i := range diagram.ProductReference_Shapes {
			if (diagram.ProductReference_Shapes[i] == nil) != (diagramOther.ProductReference_Shapes[i] == nil) {
				ProductReference_ShapesDifferent = true
				break
			} else if diagram.ProductReference_Shapes[i] != nil && diagramOther.ProductReference_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ProductReference_Shapes[i] != diagramOther.ProductReference_Shapes[i] {
					ProductReference_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ProductReference_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"ProductReference_Shapes",
			len(diagramOther.ProductReference_Shapes),
			len(diagram.ProductReference_Shapes),
			func(i, j int) bool {
				return diagramOther.ProductReference_Shapes[i] == diagram.ProductReference_Shapes[j]
			},
			func(j int) string {
				return diagram.ProductReference_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagram.IsWBSNodeExpanded != diagramOther.IsWBSNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsWBSNodeExpanded"))
	}
	Task_ShapesDifferent := false
	if len(diagram.Task_Shapes) != len(diagramOther.Task_Shapes) {
		Task_ShapesDifferent = true
	} else {
		for i := range diagram.Task_Shapes {
			if (diagram.Task_Shapes[i] == nil) != (diagramOther.Task_Shapes[i] == nil) {
				Task_ShapesDifferent = true
				break
			} else if diagram.Task_Shapes[i] != nil && diagramOther.Task_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Task_Shapes[i] != diagramOther.Task_Shapes[i] {
					Task_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Task_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"Task_Shapes",
			len(diagramOther.Task_Shapes),
			len(diagram.Task_Shapes),
			func(i, j int) bool {
				return diagramOther.Task_Shapes[i] == diagram.Task_Shapes[j]
			},
			func(j int) string {
				return diagram.Task_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TasksWhoseNodeIsExpandedDifferent := false
	if len(diagram.TasksWhoseNodeIsExpanded) != len(diagramOther.TasksWhoseNodeIsExpanded) {
		TasksWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TasksWhoseNodeIsExpanded {
			if (diagram.TasksWhoseNodeIsExpanded[i] == nil) != (diagramOther.TasksWhoseNodeIsExpanded[i] == nil) {
				TasksWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.TasksWhoseNodeIsExpanded[i] != nil && diagramOther.TasksWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.TasksWhoseNodeIsExpanded[i] != diagramOther.TasksWhoseNodeIsExpanded[i] {
					TasksWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"TasksWhoseNodeIsExpanded",
			len(diagramOther.TasksWhoseNodeIsExpanded),
			len(diagram.TasksWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.TasksWhoseNodeIsExpanded[i] == diagram.TasksWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.TasksWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TasksWhoseInputNodeIsExpandedDifferent := false
	if len(diagram.TasksWhoseInputNodeIsExpanded) != len(diagramOther.TasksWhoseInputNodeIsExpanded) {
		TasksWhoseInputNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TasksWhoseInputNodeIsExpanded {
			if (diagram.TasksWhoseInputNodeIsExpanded[i] == nil) != (diagramOther.TasksWhoseInputNodeIsExpanded[i] == nil) {
				TasksWhoseInputNodeIsExpandedDifferent = true
				break
			} else if diagram.TasksWhoseInputNodeIsExpanded[i] != nil && diagramOther.TasksWhoseInputNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.TasksWhoseInputNodeIsExpanded[i] != diagramOther.TasksWhoseInputNodeIsExpanded[i] {
					TasksWhoseInputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseInputNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"TasksWhoseInputNodeIsExpanded",
			len(diagramOther.TasksWhoseInputNodeIsExpanded),
			len(diagram.TasksWhoseInputNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.TasksWhoseInputNodeIsExpanded[i] == diagram.TasksWhoseInputNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.TasksWhoseInputNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TasksWhoseOutputNodeIsExpandedDifferent := false
	if len(diagram.TasksWhoseOutputNodeIsExpanded) != len(diagramOther.TasksWhoseOutputNodeIsExpanded) {
		TasksWhoseOutputNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TasksWhoseOutputNodeIsExpanded {
			if (diagram.TasksWhoseOutputNodeIsExpanded[i] == nil) != (diagramOther.TasksWhoseOutputNodeIsExpanded[i] == nil) {
				TasksWhoseOutputNodeIsExpandedDifferent = true
				break
			} else if diagram.TasksWhoseOutputNodeIsExpanded[i] != nil && diagramOther.TasksWhoseOutputNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.TasksWhoseOutputNodeIsExpanded[i] != diagramOther.TasksWhoseOutputNodeIsExpanded[i] {
					TasksWhoseOutputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseOutputNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"TasksWhoseOutputNodeIsExpanded",
			len(diagramOther.TasksWhoseOutputNodeIsExpanded),
			len(diagram.TasksWhoseOutputNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.TasksWhoseOutputNodeIsExpanded[i] == diagram.TasksWhoseOutputNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.TasksWhoseOutputNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TasksWhosePredecessorNodeIsExpandedDifferent := false
	if len(diagram.TasksWhosePredecessorNodeIsExpanded) != len(diagramOther.TasksWhosePredecessorNodeIsExpanded) {
		TasksWhosePredecessorNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TasksWhosePredecessorNodeIsExpanded {
			if (diagram.TasksWhosePredecessorNodeIsExpanded[i] == nil) != (diagramOther.TasksWhosePredecessorNodeIsExpanded[i] == nil) {
				TasksWhosePredecessorNodeIsExpandedDifferent = true
				break
			} else if diagram.TasksWhosePredecessorNodeIsExpanded[i] != nil && diagramOther.TasksWhosePredecessorNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.TasksWhosePredecessorNodeIsExpanded[i] != diagramOther.TasksWhosePredecessorNodeIsExpanded[i] {
					TasksWhosePredecessorNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhosePredecessorNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"TasksWhosePredecessorNodeIsExpanded",
			len(diagramOther.TasksWhosePredecessorNodeIsExpanded),
			len(diagram.TasksWhosePredecessorNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.TasksWhosePredecessorNodeIsExpanded[i] == diagram.TasksWhosePredecessorNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.TasksWhosePredecessorNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagram.IsTaskGroupsNodeExpanded != diagramOther.IsTaskGroupsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsTaskGroupsNodeExpanded"))
	}
	TaskGroupShapesDifferent := false
	if len(diagram.TaskGroupShapes) != len(diagramOther.TaskGroupShapes) {
		TaskGroupShapesDifferent = true
	} else {
		for i := range diagram.TaskGroupShapes {
			if (diagram.TaskGroupShapes[i] == nil) != (diagramOther.TaskGroupShapes[i] == nil) {
				TaskGroupShapesDifferent = true
				break
			} else if diagram.TaskGroupShapes[i] != nil && diagramOther.TaskGroupShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.TaskGroupShapes[i] != diagramOther.TaskGroupShapes[i] {
					TaskGroupShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskGroupShapesDifferent {
		ops := stage.Diff(
			diagram,
			"TaskGroupShapes",
			len(diagramOther.TaskGroupShapes),
			len(diagram.TaskGroupShapes),
			func(i, j int) bool {
				return diagramOther.TaskGroupShapes[i] == diagram.TaskGroupShapes[j]
			},
			func(j int) string {
				return diagram.TaskGroupShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskGroupsWhoseNodeIsExpandedDifferent := false
	if len(diagram.TaskGroupsWhoseNodeIsExpanded) != len(diagramOther.TaskGroupsWhoseNodeIsExpanded) {
		TaskGroupsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.TaskGroupsWhoseNodeIsExpanded {
			if (diagram.TaskGroupsWhoseNodeIsExpanded[i] == nil) != (diagramOther.TaskGroupsWhoseNodeIsExpanded[i] == nil) {
				TaskGroupsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.TaskGroupsWhoseNodeIsExpanded[i] != nil && diagramOther.TaskGroupsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.TaskGroupsWhoseNodeIsExpanded[i] != diagramOther.TaskGroupsWhoseNodeIsExpanded[i] {
					TaskGroupsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TaskGroupsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"TaskGroupsWhoseNodeIsExpanded",
			len(diagramOther.TaskGroupsWhoseNodeIsExpanded),
			len(diagram.TaskGroupsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.TaskGroupsWhoseNodeIsExpanded[i] == diagram.TaskGroupsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.TaskGroupsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskComposition_ShapesDifferent := false
	if len(diagram.TaskComposition_Shapes) != len(diagramOther.TaskComposition_Shapes) {
		TaskComposition_ShapesDifferent = true
	} else {
		for i := range diagram.TaskComposition_Shapes {
			if (diagram.TaskComposition_Shapes[i] == nil) != (diagramOther.TaskComposition_Shapes[i] == nil) {
				TaskComposition_ShapesDifferent = true
				break
			} else if diagram.TaskComposition_Shapes[i] != nil && diagramOther.TaskComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.TaskComposition_Shapes[i] != diagramOther.TaskComposition_Shapes[i] {
					TaskComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskComposition_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"TaskComposition_Shapes",
			len(diagramOther.TaskComposition_Shapes),
			len(diagram.TaskComposition_Shapes),
			func(i, j int) bool {
				return diagramOther.TaskComposition_Shapes[i] == diagram.TaskComposition_Shapes[j]
			},
			func(j int) string {
				return diagram.TaskComposition_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskInputShapesDifferent := false
	if len(diagram.TaskInputShapes) != len(diagramOther.TaskInputShapes) {
		TaskInputShapesDifferent = true
	} else {
		for i := range diagram.TaskInputShapes {
			if (diagram.TaskInputShapes[i] == nil) != (diagramOther.TaskInputShapes[i] == nil) {
				TaskInputShapesDifferent = true
				break
			} else if diagram.TaskInputShapes[i] != nil && diagramOther.TaskInputShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.TaskInputShapes[i] != diagramOther.TaskInputShapes[i] {
					TaskInputShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskInputShapesDifferent {
		ops := stage.Diff(
			diagram,
			"TaskInputShapes",
			len(diagramOther.TaskInputShapes),
			len(diagram.TaskInputShapes),
			func(i, j int) bool {
				return diagramOther.TaskInputShapes[i] == diagram.TaskInputShapes[j]
			},
			func(j int) string {
				return diagram.TaskInputShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskOutputShapesDifferent := false
	if len(diagram.TaskOutputShapes) != len(diagramOther.TaskOutputShapes) {
		TaskOutputShapesDifferent = true
	} else {
		for i := range diagram.TaskOutputShapes {
			if (diagram.TaskOutputShapes[i] == nil) != (diagramOther.TaskOutputShapes[i] == nil) {
				TaskOutputShapesDifferent = true
				break
			} else if diagram.TaskOutputShapes[i] != nil && diagramOther.TaskOutputShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.TaskOutputShapes[i] != diagramOther.TaskOutputShapes[i] {
					TaskOutputShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskOutputShapesDifferent {
		ops := stage.Diff(
			diagram,
			"TaskOutputShapes",
			len(diagramOther.TaskOutputShapes),
			len(diagram.TaskOutputShapes),
			func(i, j int) bool {
				return diagramOther.TaskOutputShapes[i] == diagram.TaskOutputShapes[j]
			},
			func(j int) string {
				return diagram.TaskOutputShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TaskPredecessorShapesDifferent := false
	if len(diagram.TaskPredecessorShapes) != len(diagramOther.TaskPredecessorShapes) {
		TaskPredecessorShapesDifferent = true
	} else {
		for i := range diagram.TaskPredecessorShapes {
			if (diagram.TaskPredecessorShapes[i] == nil) != (diagramOther.TaskPredecessorShapes[i] == nil) {
				TaskPredecessorShapesDifferent = true
				break
			} else if diagram.TaskPredecessorShapes[i] != nil && diagramOther.TaskPredecessorShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.TaskPredecessorShapes[i] != diagramOther.TaskPredecessorShapes[i] {
					TaskPredecessorShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskPredecessorShapesDifferent {
		ops := stage.Diff(
			diagram,
			"TaskPredecessorShapes",
			len(diagramOther.TaskPredecessorShapes),
			len(diagram.TaskPredecessorShapes),
			func(i, j int) bool {
				return diagramOther.TaskPredecessorShapes[i] == diagram.TaskPredecessorShapes[j]
			},
			func(j int) string {
				return diagram.TaskPredecessorShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Note_ShapesDifferent := false
	if len(diagram.Note_Shapes) != len(diagramOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagram.Note_Shapes {
			if (diagram.Note_Shapes[i] == nil) != (diagramOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagram.Note_Shapes[i] != nil && diagramOther.Note_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Note_Shapes[i] != diagramOther.Note_Shapes[i] {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"Note_Shapes",
			len(diagramOther.Note_Shapes),
			len(diagram.Note_Shapes),
			func(i, j int) bool {
				return diagramOther.Note_Shapes[i] == diagram.Note_Shapes[j]
			},
			func(j int) string {
				return diagram.Note_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagram.NotesWhoseNodeIsExpanded) != len(diagramOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.NotesWhoseNodeIsExpanded {
			if (diagram.NotesWhoseNodeIsExpanded[i] == nil) != (diagramOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.NotesWhoseNodeIsExpanded[i] != nil && diagramOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.NotesWhoseNodeIsExpanded[i] != diagramOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"NotesWhoseNodeIsExpanded",
			len(diagramOther.NotesWhoseNodeIsExpanded),
			len(diagram.NotesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.NotesWhoseNodeIsExpanded[i] == diagram.NotesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.NotesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagram.IsNotesNodeExpanded != diagramOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NoteProductShapesDifferent := false
	if len(diagram.NoteProductShapes) != len(diagramOther.NoteProductShapes) {
		NoteProductShapesDifferent = true
	} else {
		for i := range diagram.NoteProductShapes {
			if (diagram.NoteProductShapes[i] == nil) != (diagramOther.NoteProductShapes[i] == nil) {
				NoteProductShapesDifferent = true
				break
			} else if diagram.NoteProductShapes[i] != nil && diagramOther.NoteProductShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteProductShapes[i] != diagramOther.NoteProductShapes[i] {
					NoteProductShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteProductShapesDifferent {
		ops := stage.Diff(
			diagram,
			"NoteProductShapes",
			len(diagramOther.NoteProductShapes),
			len(diagram.NoteProductShapes),
			func(i, j int) bool {
				return diagramOther.NoteProductShapes[i] == diagram.NoteProductShapes[j]
			},
			func(j int) string {
				return diagram.NoteProductShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NoteTaskShapesDifferent := false
	if len(diagram.NoteTaskShapes) != len(diagramOther.NoteTaskShapes) {
		NoteTaskShapesDifferent = true
	} else {
		for i := range diagram.NoteTaskShapes {
			if (diagram.NoteTaskShapes[i] == nil) != (diagramOther.NoteTaskShapes[i] == nil) {
				NoteTaskShapesDifferent = true
				break
			} else if diagram.NoteTaskShapes[i] != nil && diagramOther.NoteTaskShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteTaskShapes[i] != diagramOther.NoteTaskShapes[i] {
					NoteTaskShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteTaskShapesDifferent {
		ops := stage.Diff(
			diagram,
			"NoteTaskShapes",
			len(diagramOther.NoteTaskShapes),
			len(diagram.NoteTaskShapes),
			func(i, j int) bool {
				return diagramOther.NoteTaskShapes[i] == diagram.NoteTaskShapes[j]
			},
			func(j int) string {
				return diagram.NoteTaskShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NoteResourceShapesDifferent := false
	if len(diagram.NoteResourceShapes) != len(diagramOther.NoteResourceShapes) {
		NoteResourceShapesDifferent = true
	} else {
		for i := range diagram.NoteResourceShapes {
			if (diagram.NoteResourceShapes[i] == nil) != (diagramOther.NoteResourceShapes[i] == nil) {
				NoteResourceShapesDifferent = true
				break
			} else if diagram.NoteResourceShapes[i] != nil && diagramOther.NoteResourceShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.NoteResourceShapes[i] != diagramOther.NoteResourceShapes[i] {
					NoteResourceShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteResourceShapesDifferent {
		ops := stage.Diff(
			diagram,
			"NoteResourceShapes",
			len(diagramOther.NoteResourceShapes),
			len(diagram.NoteResourceShapes),
			func(i, j int) bool {
				return diagramOther.NoteResourceShapes[i] == diagram.NoteResourceShapes[j]
			},
			func(j int) string {
				return diagram.NoteResourceShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Resource_ShapesDifferent := false
	if len(diagram.Resource_Shapes) != len(diagramOther.Resource_Shapes) {
		Resource_ShapesDifferent = true
	} else {
		for i := range diagram.Resource_Shapes {
			if (diagram.Resource_Shapes[i] == nil) != (diagramOther.Resource_Shapes[i] == nil) {
				Resource_ShapesDifferent = true
				break
			} else if diagram.Resource_Shapes[i] != nil && diagramOther.Resource_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Resource_Shapes[i] != diagramOther.Resource_Shapes[i] {
					Resource_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Resource_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"Resource_Shapes",
			len(diagramOther.Resource_Shapes),
			len(diagram.Resource_Shapes),
			func(i, j int) bool {
				return diagramOther.Resource_Shapes[i] == diagram.Resource_Shapes[j]
			},
			func(j int) string {
				return diagram.Resource_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ResourcesWhoseNodeIsExpandedDifferent := false
	if len(diagram.ResourcesWhoseNodeIsExpanded) != len(diagramOther.ResourcesWhoseNodeIsExpanded) {
		ResourcesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagram.ResourcesWhoseNodeIsExpanded {
			if (diagram.ResourcesWhoseNodeIsExpanded[i] == nil) != (diagramOther.ResourcesWhoseNodeIsExpanded[i] == nil) {
				ResourcesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagram.ResourcesWhoseNodeIsExpanded[i] != nil && diagramOther.ResourcesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagram.ResourcesWhoseNodeIsExpanded[i] != diagramOther.ResourcesWhoseNodeIsExpanded[i] {
					ResourcesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ResourcesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagram,
			"ResourcesWhoseNodeIsExpanded",
			len(diagramOther.ResourcesWhoseNodeIsExpanded),
			len(diagram.ResourcesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramOther.ResourcesWhoseNodeIsExpanded[i] == diagram.ResourcesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagram.ResourcesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagram.IsResourcesNodeExpanded != diagramOther.IsResourcesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	ResourceComposition_ShapesDifferent := false
	if len(diagram.ResourceComposition_Shapes) != len(diagramOther.ResourceComposition_Shapes) {
		ResourceComposition_ShapesDifferent = true
	} else {
		for i := range diagram.ResourceComposition_Shapes {
			if (diagram.ResourceComposition_Shapes[i] == nil) != (diagramOther.ResourceComposition_Shapes[i] == nil) {
				ResourceComposition_ShapesDifferent = true
				break
			} else if diagram.ResourceComposition_Shapes[i] != nil && diagramOther.ResourceComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ResourceComposition_Shapes[i] != diagramOther.ResourceComposition_Shapes[i] {
					ResourceComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ResourceComposition_ShapesDifferent {
		ops := stage.Diff(
			diagram,
			"ResourceComposition_Shapes",
			len(diagramOther.ResourceComposition_Shapes),
			len(diagram.ResourceComposition_Shapes),
			func(i, j int) bool {
				return diagramOther.ResourceComposition_Shapes[i] == diagram.ResourceComposition_Shapes[j]
			},
			func(j int) string {
				return diagram.ResourceComposition_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ResourceTaskShapesDifferent := false
	if len(diagram.ResourceTaskShapes) != len(diagramOther.ResourceTaskShapes) {
		ResourceTaskShapesDifferent = true
	} else {
		for i := range diagram.ResourceTaskShapes {
			if (diagram.ResourceTaskShapes[i] == nil) != (diagramOther.ResourceTaskShapes[i] == nil) {
				ResourceTaskShapesDifferent = true
				break
			} else if diagram.ResourceTaskShapes[i] != nil && diagramOther.ResourceTaskShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ResourceTaskShapes[i] != diagramOther.ResourceTaskShapes[i] {
					ResourceTaskShapesDifferent = true
					break
				}
			}
		}
	}
	if ResourceTaskShapesDifferent {
		ops := stage.Diff(
			diagram,
			"ResourceTaskShapes",
			len(diagramOther.ResourceTaskShapes),
			len(diagram.ResourceTaskShapes),
			func(i, j int) bool {
				return diagramOther.ResourceTaskShapes[i] == diagram.ResourceTaskShapes[j]
			},
			func(j int) string {
				return diagram.ResourceTaskShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := stage.Diff(
			library,
			"SubLibraries",
			len(libraryOther.SubLibraries),
			len(library.SubLibraries),
			func(i, j int) bool {
				return libraryOther.SubLibraries[i] == library.SubLibraries[j]
			},
			func(j int) string {
				return library.SubLibraries[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	RootProductsDifferent := false
	if len(library.RootProducts) != len(libraryOther.RootProducts) {
		RootProductsDifferent = true
	} else {
		for i := range library.RootProducts {
			if (library.RootProducts[i] == nil) != (libraryOther.RootProducts[i] == nil) {
				RootProductsDifferent = true
				break
			} else if library.RootProducts[i] != nil && libraryOther.RootProducts[i] != nil {
				// this is a pointer comparaison
				if library.RootProducts[i] != libraryOther.RootProducts[i] {
					RootProductsDifferent = true
					break
				}
			}
		}
	}
	if RootProductsDifferent {
		ops := stage.Diff(
			library,
			"RootProducts",
			len(libraryOther.RootProducts),
			len(library.RootProducts),
			func(i, j int) bool {
				return libraryOther.RootProducts[i] == library.RootProducts[j]
			},
			func(j int) string {
				return library.RootProducts[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootTasksDifferent := false
	if len(library.RootTasks) != len(libraryOther.RootTasks) {
		RootTasksDifferent = true
	} else {
		for i := range library.RootTasks {
			if (library.RootTasks[i] == nil) != (libraryOther.RootTasks[i] == nil) {
				RootTasksDifferent = true
				break
			} else if library.RootTasks[i] != nil && libraryOther.RootTasks[i] != nil {
				// this is a pointer comparaison
				if library.RootTasks[i] != libraryOther.RootTasks[i] {
					RootTasksDifferent = true
					break
				}
			}
		}
	}
	if RootTasksDifferent {
		ops := stage.Diff(
			library,
			"RootTasks",
			len(libraryOther.RootTasks),
			len(library.RootTasks),
			func(i, j int) bool {
				return libraryOther.RootTasks[i] == library.RootTasks[j]
			},
			func(j int) string {
				return library.RootTasks[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootTaskGroupsDifferent := false
	if len(library.RootTaskGroups) != len(libraryOther.RootTaskGroups) {
		RootTaskGroupsDifferent = true
	} else {
		for i := range library.RootTaskGroups {
			if (library.RootTaskGroups[i] == nil) != (libraryOther.RootTaskGroups[i] == nil) {
				RootTaskGroupsDifferent = true
				break
			} else if library.RootTaskGroups[i] != nil && libraryOther.RootTaskGroups[i] != nil {
				// this is a pointer comparaison
				if library.RootTaskGroups[i] != libraryOther.RootTaskGroups[i] {
					RootTaskGroupsDifferent = true
					break
				}
			}
		}
	}
	if RootTaskGroupsDifferent {
		ops := stage.Diff(
			library,
			"RootTaskGroups",
			len(libraryOther.RootTaskGroups),
			len(library.RootTaskGroups),
			func(i, j int) bool {
				return libraryOther.RootTaskGroups[i] == library.RootTaskGroups[j]
			},
			func(j int) string {
				return library.RootTaskGroups[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootResourcesDifferent := false
	if len(library.RootResources) != len(libraryOther.RootResources) {
		RootResourcesDifferent = true
	} else {
		for i := range library.RootResources {
			if (library.RootResources[i] == nil) != (libraryOther.RootResources[i] == nil) {
				RootResourcesDifferent = true
				break
			} else if library.RootResources[i] != nil && libraryOther.RootResources[i] != nil {
				// this is a pointer comparaison
				if library.RootResources[i] != libraryOther.RootResources[i] {
					RootResourcesDifferent = true
					break
				}
			}
		}
	}
	if RootResourcesDifferent {
		ops := stage.Diff(
			library,
			"RootResources",
			len(libraryOther.RootResources),
			len(library.RootResources),
			func(i, j int) bool {
				return libraryOther.RootResources[i] == library.RootResources[j]
			},
			func(j int) string {
				return library.RootResources[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NotesDifferent := false
	if len(library.Notes) != len(libraryOther.Notes) {
		NotesDifferent = true
	} else {
		for i := range library.Notes {
			if (library.Notes[i] == nil) != (libraryOther.Notes[i] == nil) {
				NotesDifferent = true
				break
			} else if library.Notes[i] != nil && libraryOther.Notes[i] != nil {
				// this is a pointer comparaison
				if library.Notes[i] != libraryOther.Notes[i] {
					NotesDifferent = true
					break
				}
			}
		}
	}
	if NotesDifferent {
		ops := stage.Diff(
			library,
			"Notes",
			len(libraryOther.Notes),
			len(library.Notes),
			func(i, j int) bool {
				return libraryOther.Notes[i] == library.Notes[j]
			},
			func(j int) string {
				return library.Notes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DiagramsDifferent := false
	if len(library.Diagrams) != len(libraryOther.Diagrams) {
		DiagramsDifferent = true
	} else {
		for i := range library.Diagrams {
			if (library.Diagrams[i] == nil) != (libraryOther.Diagrams[i] == nil) {
				DiagramsDifferent = true
				break
			} else if library.Diagrams[i] != nil && libraryOther.Diagrams[i] != nil {
				// this is a pointer comparaison
				if library.Diagrams[i] != libraryOther.Diagrams[i] {
					DiagramsDifferent = true
					break
				}
			}
		}
	}
	if DiagramsDifferent {
		ops := stage.Diff(
			library,
			"Diagrams",
			len(libraryOther.Diagrams),
			len(library.Diagrams),
			func(i, j int) bool {
				return libraryOther.Diagrams[i] == library.Diagrams[j]
			},
			func(j int) string {
				return library.Diagrams[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.ComputedPrefix != noteOther.ComputedPrefix {
		diffs = append(diffs, note.GongMarshallField(stage, "ComputedPrefix"))
	}
	if note.IsExpanded != noteOther.IsExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsExpanded"))
	}
	if note.LayoutDirection != noteOther.LayoutDirection {
		diffs = append(diffs, note.GongMarshallField(stage, "LayoutDirection"))
	}
	ProductsDifferent := false
	if len(note.Products) != len(noteOther.Products) {
		ProductsDifferent = true
	} else {
		for i := range note.Products {
			if (note.Products[i] == nil) != (noteOther.Products[i] == nil) {
				ProductsDifferent = true
				break
			} else if note.Products[i] != nil && noteOther.Products[i] != nil {
				// this is a pointer comparaison
				if note.Products[i] != noteOther.Products[i] {
					ProductsDifferent = true
					break
				}
			}
		}
	}
	if ProductsDifferent {
		ops := stage.Diff(
			note,
			"Products",
			len(noteOther.Products),
			len(note.Products),
			func(i, j int) bool {
				return noteOther.Products[i] == note.Products[j]
			},
			func(j int) string {
				return note.Products[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	TasksDifferent := false
	if len(note.Tasks) != len(noteOther.Tasks) {
		TasksDifferent = true
	} else {
		for i := range note.Tasks {
			if (note.Tasks[i] == nil) != (noteOther.Tasks[i] == nil) {
				TasksDifferent = true
				break
			} else if note.Tasks[i] != nil && noteOther.Tasks[i] != nil {
				// this is a pointer comparaison
				if note.Tasks[i] != noteOther.Tasks[i] {
					TasksDifferent = true
					break
				}
			}
		}
	}
	if TasksDifferent {
		ops := stage.Diff(
			note,
			"Tasks",
			len(noteOther.Tasks),
			len(note.Tasks),
			func(i, j int) bool {
				return noteOther.Tasks[i] == note.Tasks[j]
			},
			func(j int) string {
				return note.Tasks[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ResourcesDifferent := false
	if len(note.Resources) != len(noteOther.Resources) {
		ResourcesDifferent = true
	} else {
		for i := range note.Resources {
			if (note.Resources[i] == nil) != (noteOther.Resources[i] == nil) {
				ResourcesDifferent = true
				break
			} else if note.Resources[i] != nil && noteOther.Resources[i] != nil {
				// this is a pointer comparaison
				if note.Resources[i] != noteOther.Resources[i] {
					ResourcesDifferent = true
					break
				}
			}
		}
	}
	if ResourcesDifferent {
		ops := stage.Diff(
			note,
			"Resources",
			len(noteOther.Resources),
			len(note.Resources),
			func(i, j int) bool {
				return noteOther.Resources[i] == note.Resources[j]
			},
			func(j int) string {
				return note.Resources[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteproductshape *NoteProductShape) GongDiff(stage *Stage, noteproductshapeOther *NoteProductShape) (diffs []string) {
	// insertion point for field diffs
	if noteproductshape.Name != noteproductshapeOther.Name {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Name"))
	}
	if (noteproductshape.Note == nil) != (noteproductshapeOther.Note == nil) {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Note"))
	} else if noteproductshape.Note != nil && noteproductshapeOther.Note != nil {
		if noteproductshape.Note != noteproductshapeOther.Note {
			diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Note"))
		}
	}
	if (noteproductshape.Product == nil) != (noteproductshapeOther.Product == nil) {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Product"))
	} else if noteproductshape.Product != nil && noteproductshapeOther.Product != nil {
		if noteproductshape.Product != noteproductshapeOther.Product {
			diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Product"))
		}
	}
	if noteproductshape.StartRatio != noteproductshapeOther.StartRatio {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "StartRatio"))
	}
	if noteproductshape.EndRatio != noteproductshapeOther.EndRatio {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "EndRatio"))
	}
	if noteproductshape.StartOrientation != noteproductshapeOther.StartOrientation {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "StartOrientation"))
	}
	if noteproductshape.EndOrientation != noteproductshapeOther.EndOrientation {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "EndOrientation"))
	}
	if noteproductshape.CornerOffsetRatio != noteproductshapeOther.CornerOffsetRatio {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if noteproductshape.IsHidden != noteproductshapeOther.IsHidden {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteresourceshape *NoteResourceShape) GongDiff(stage *Stage, noteresourceshapeOther *NoteResourceShape) (diffs []string) {
	// insertion point for field diffs
	if noteresourceshape.Name != noteresourceshapeOther.Name {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "Name"))
	}
	if (noteresourceshape.Note == nil) != (noteresourceshapeOther.Note == nil) {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "Note"))
	} else if noteresourceshape.Note != nil && noteresourceshapeOther.Note != nil {
		if noteresourceshape.Note != noteresourceshapeOther.Note {
			diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "Note"))
		}
	}
	if (noteresourceshape.Resource == nil) != (noteresourceshapeOther.Resource == nil) {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "Resource"))
	} else if noteresourceshape.Resource != nil && noteresourceshapeOther.Resource != nil {
		if noteresourceshape.Resource != noteresourceshapeOther.Resource {
			diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "Resource"))
		}
	}
	if noteresourceshape.StartRatio != noteresourceshapeOther.StartRatio {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "StartRatio"))
	}
	if noteresourceshape.EndRatio != noteresourceshapeOther.EndRatio {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "EndRatio"))
	}
	if noteresourceshape.StartOrientation != noteresourceshapeOther.StartOrientation {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "StartOrientation"))
	}
	if noteresourceshape.EndOrientation != noteresourceshapeOther.EndOrientation {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "EndOrientation"))
	}
	if noteresourceshape.CornerOffsetRatio != noteresourceshapeOther.CornerOffsetRatio {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if noteresourceshape.IsHidden != noteresourceshapeOther.IsHidden {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteshape *NoteShape) GongDiff(stage *Stage, noteshapeOther *NoteShape) (diffs []string) {
	// insertion point for field diffs
	if noteshape.Name != noteshapeOther.Name {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Name"))
	}
	if (noteshape.Note == nil) != (noteshapeOther.Note == nil) {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
	} else if noteshape.Note != nil && noteshapeOther.Note != nil {
		if noteshape.Note != noteshapeOther.Note {
			diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
		}
	}
	if noteshape.OverideLayoutDirection != noteshapeOther.OverideLayoutDirection {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "OverideLayoutDirection"))
	}
	if noteshape.LayoutDirection != noteshapeOther.LayoutDirection {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "LayoutDirection"))
	}
	if noteshape.X != noteshapeOther.X {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "X"))
	}
	if noteshape.Y != noteshapeOther.Y {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Y"))
	}
	if noteshape.Width != noteshapeOther.Width {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Width"))
	}
	if noteshape.Height != noteshapeOther.Height {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Height"))
	}
	if noteshape.IsHidden != noteshapeOther.IsHidden {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notetaskshape *NoteTaskShape) GongDiff(stage *Stage, notetaskshapeOther *NoteTaskShape) (diffs []string) {
	// insertion point for field diffs
	if notetaskshape.Name != notetaskshapeOther.Name {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Name"))
	}
	if (notetaskshape.Note == nil) != (notetaskshapeOther.Note == nil) {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
	} else if notetaskshape.Note != nil && notetaskshapeOther.Note != nil {
		if notetaskshape.Note != notetaskshapeOther.Note {
			diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
		}
	}
	if (notetaskshape.Task == nil) != (notetaskshapeOther.Task == nil) {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
	} else if notetaskshape.Task != nil && notetaskshapeOther.Task != nil {
		if notetaskshape.Task != notetaskshapeOther.Task {
			diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
		}
	}
	if notetaskshape.StartRatio != notetaskshapeOther.StartRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "StartRatio"))
	}
	if notetaskshape.EndRatio != notetaskshapeOther.EndRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "EndRatio"))
	}
	if notetaskshape.StartOrientation != notetaskshapeOther.StartOrientation {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notetaskshape.EndOrientation != notetaskshapeOther.EndOrientation {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notetaskshape.CornerOffsetRatio != notetaskshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notetaskshape.IsHidden != notetaskshapeOther.IsHidden {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (product *Product) GongDiff(stage *Stage, productOther *Product) (diffs []string) {
	// insertion point for field diffs
	if product.Name != productOther.Name {
		diffs = append(diffs, product.GongMarshallField(stage, "Name"))
	}
	if product.Description != productOther.Description {
		diffs = append(diffs, product.GongMarshallField(stage, "Description"))
	}
	SubProductsDifferent := false
	if len(product.SubProducts) != len(productOther.SubProducts) {
		SubProductsDifferent = true
	} else {
		for i := range product.SubProducts {
			if (product.SubProducts[i] == nil) != (productOther.SubProducts[i] == nil) {
				SubProductsDifferent = true
				break
			} else if product.SubProducts[i] != nil && productOther.SubProducts[i] != nil {
				// this is a pointer comparaison
				if product.SubProducts[i] != productOther.SubProducts[i] {
					SubProductsDifferent = true
					break
				}
			}
		}
	}
	if SubProductsDifferent {
		ops := stage.Diff(
			product,
			"SubProducts",
			len(productOther.SubProducts),
			len(product.SubProducts),
			func(i, j int) bool {
				return productOther.SubProducts[i] == product.SubProducts[j]
			},
			func(j int) string {
				return product.SubProducts[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if product.IsProducersNodeExpanded != productOther.IsProducersNodeExpanded {
		diffs = append(diffs, product.GongMarshallField(stage, "IsProducersNodeExpanded"))
	}
	if product.IsConsumersNodeExpanded != productOther.IsConsumersNodeExpanded {
		diffs = append(diffs, product.GongMarshallField(stage, "IsConsumersNodeExpanded"))
	}
	if product.IsImport != productOther.IsImport {
		diffs = append(diffs, product.GongMarshallField(stage, "IsImport"))
	}
	if (product.ReferencedProduct == nil) != (productOther.ReferencedProduct == nil) {
		diffs = append(diffs, product.GongMarshallField(stage, "ReferencedProduct"))
	} else if product.ReferencedProduct != nil && productOther.ReferencedProduct != nil {
		if product.ReferencedProduct != productOther.ReferencedProduct {
			diffs = append(diffs, product.GongMarshallField(stage, "ReferencedProduct"))
		}
	}
	if product.ComputedPrefix != productOther.ComputedPrefix {
		diffs = append(diffs, product.GongMarshallField(stage, "ComputedPrefix"))
	}
	if product.IsExpanded != productOther.IsExpanded {
		diffs = append(diffs, product.GongMarshallField(stage, "IsExpanded"))
	}
	if product.LayoutDirection != productOther.LayoutDirection {
		diffs = append(diffs, product.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (productcompositionshape *ProductCompositionShape) GongDiff(stage *Stage, productcompositionshapeOther *ProductCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if productcompositionshape.Name != productcompositionshapeOther.Name {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "Name"))
	}
	if (productcompositionshape.Product == nil) != (productcompositionshapeOther.Product == nil) {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "Product"))
	} else if productcompositionshape.Product != nil && productcompositionshapeOther.Product != nil {
		if productcompositionshape.Product != productcompositionshapeOther.Product {
			diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "Product"))
		}
	}
	if productcompositionshape.StartRatio != productcompositionshapeOther.StartRatio {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if productcompositionshape.EndRatio != productcompositionshapeOther.EndRatio {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if productcompositionshape.StartOrientation != productcompositionshapeOther.StartOrientation {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if productcompositionshape.EndOrientation != productcompositionshapeOther.EndOrientation {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if productcompositionshape.CornerOffsetRatio != productcompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if productcompositionshape.IsHidden != productcompositionshapeOther.IsHidden {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (productreferenceshape *ProductReferenceShape) GongDiff(stage *Stage, productreferenceshapeOther *ProductReferenceShape) (diffs []string) {
	// insertion point for field diffs
	if productreferenceshape.Name != productreferenceshapeOther.Name {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "Name"))
	}
	if (productreferenceshape.Product == nil) != (productreferenceshapeOther.Product == nil) {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "Product"))
	} else if productreferenceshape.Product != nil && productreferenceshapeOther.Product != nil {
		if productreferenceshape.Product != productreferenceshapeOther.Product {
			diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "Product"))
		}
	}
	if (productreferenceshape.ReferencedProduct == nil) != (productreferenceshapeOther.ReferencedProduct == nil) {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "ReferencedProduct"))
	} else if productreferenceshape.ReferencedProduct != nil && productreferenceshapeOther.ReferencedProduct != nil {
		if productreferenceshape.ReferencedProduct != productreferenceshapeOther.ReferencedProduct {
			diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "ReferencedProduct"))
		}
	}
	if productreferenceshape.StartRatio != productreferenceshapeOther.StartRatio {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "StartRatio"))
	}
	if productreferenceshape.EndRatio != productreferenceshapeOther.EndRatio {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "EndRatio"))
	}
	if productreferenceshape.StartOrientation != productreferenceshapeOther.StartOrientation {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "StartOrientation"))
	}
	if productreferenceshape.EndOrientation != productreferenceshapeOther.EndOrientation {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "EndOrientation"))
	}
	if productreferenceshape.CornerOffsetRatio != productreferenceshapeOther.CornerOffsetRatio {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if productreferenceshape.IsHidden != productreferenceshapeOther.IsHidden {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (productshape *ProductShape) GongDiff(stage *Stage, productshapeOther *ProductShape) (diffs []string) {
	// insertion point for field diffs
	if productshape.Name != productshapeOther.Name {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Name"))
	}
	if (productshape.Product == nil) != (productshapeOther.Product == nil) {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Product"))
	} else if productshape.Product != nil && productshapeOther.Product != nil {
		if productshape.Product != productshapeOther.Product {
			diffs = append(diffs, productshape.GongMarshallField(stage, "Product"))
		}
	}
	if productshape.IsShowType != productshapeOther.IsShowType {
		diffs = append(diffs, productshape.GongMarshallField(stage, "IsShowType"))
	}
	if productshape.OverideLayoutDirection != productshapeOther.OverideLayoutDirection {
		diffs = append(diffs, productshape.GongMarshallField(stage, "OverideLayoutDirection"))
	}
	if productshape.LayoutDirection != productshapeOther.LayoutDirection {
		diffs = append(diffs, productshape.GongMarshallField(stage, "LayoutDirection"))
	}
	if productshape.X != productshapeOther.X {
		diffs = append(diffs, productshape.GongMarshallField(stage, "X"))
	}
	if productshape.Y != productshapeOther.Y {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Y"))
	}
	if productshape.Width != productshapeOther.Width {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Width"))
	}
	if productshape.Height != productshapeOther.Height {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Height"))
	}
	if productshape.IsHidden != productshapeOther.IsHidden {
		diffs = append(diffs, productshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (resource *Resource) GongDiff(stage *Stage, resourceOther *Resource) (diffs []string) {
	// insertion point for field diffs
	if resource.Name != resourceOther.Name {
		diffs = append(diffs, resource.GongMarshallField(stage, "Name"))
	}
	if resource.Description != resourceOther.Description {
		diffs = append(diffs, resource.GongMarshallField(stage, "Description"))
	}
	TasksDifferent := false
	if len(resource.Tasks) != len(resourceOther.Tasks) {
		TasksDifferent = true
	} else {
		for i := range resource.Tasks {
			if (resource.Tasks[i] == nil) != (resourceOther.Tasks[i] == nil) {
				TasksDifferent = true
				break
			} else if resource.Tasks[i] != nil && resourceOther.Tasks[i] != nil {
				// this is a pointer comparaison
				if resource.Tasks[i] != resourceOther.Tasks[i] {
					TasksDifferent = true
					break
				}
			}
		}
	}
	if TasksDifferent {
		ops := stage.Diff(
			resource,
			"Tasks",
			len(resourceOther.Tasks),
			len(resource.Tasks),
			func(i, j int) bool {
				return resourceOther.Tasks[i] == resource.Tasks[j]
			},
			func(j int) string {
				return resource.Tasks[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	SubResourcesDifferent := false
	if len(resource.SubResources) != len(resourceOther.SubResources) {
		SubResourcesDifferent = true
	} else {
		for i := range resource.SubResources {
			if (resource.SubResources[i] == nil) != (resourceOther.SubResources[i] == nil) {
				SubResourcesDifferent = true
				break
			} else if resource.SubResources[i] != nil && resourceOther.SubResources[i] != nil {
				// this is a pointer comparaison
				if resource.SubResources[i] != resourceOther.SubResources[i] {
					SubResourcesDifferent = true
					break
				}
			}
		}
	}
	if SubResourcesDifferent {
		ops := stage.Diff(
			resource,
			"SubResources",
			len(resourceOther.SubResources),
			len(resource.SubResources),
			func(i, j int) bool {
				return resourceOther.SubResources[i] == resource.SubResources[j]
			},
			func(j int) string {
				return resource.SubResources[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if resource.ComputedPrefix != resourceOther.ComputedPrefix {
		diffs = append(diffs, resource.GongMarshallField(stage, "ComputedPrefix"))
	}
	if resource.IsExpanded != resourceOther.IsExpanded {
		diffs = append(diffs, resource.GongMarshallField(stage, "IsExpanded"))
	}
	if resource.LayoutDirection != resourceOther.LayoutDirection {
		diffs = append(diffs, resource.GongMarshallField(stage, "LayoutDirection"))
	}
	if resource.IsImport != resourceOther.IsImport {
		diffs = append(diffs, resource.GongMarshallField(stage, "IsImport"))
	}
	if (resource.ReferencedResource == nil) != (resourceOther.ReferencedResource == nil) {
		diffs = append(diffs, resource.GongMarshallField(stage, "ReferencedResource"))
	} else if resource.ReferencedResource != nil && resourceOther.ReferencedResource != nil {
		if resource.ReferencedResource != resourceOther.ReferencedResource {
			diffs = append(diffs, resource.GongMarshallField(stage, "ReferencedResource"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (resourcecompositionshape *ResourceCompositionShape) GongDiff(stage *Stage, resourcecompositionshapeOther *ResourceCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if resourcecompositionshape.Name != resourcecompositionshapeOther.Name {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "Name"))
	}
	if (resourcecompositionshape.Resource == nil) != (resourcecompositionshapeOther.Resource == nil) {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "Resource"))
	} else if resourcecompositionshape.Resource != nil && resourcecompositionshapeOther.Resource != nil {
		if resourcecompositionshape.Resource != resourcecompositionshapeOther.Resource {
			diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "Resource"))
		}
	}
	if resourcecompositionshape.StartRatio != resourcecompositionshapeOther.StartRatio {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if resourcecompositionshape.EndRatio != resourcecompositionshapeOther.EndRatio {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if resourcecompositionshape.StartOrientation != resourcecompositionshapeOther.StartOrientation {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if resourcecompositionshape.EndOrientation != resourcecompositionshapeOther.EndOrientation {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if resourcecompositionshape.CornerOffsetRatio != resourcecompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if resourcecompositionshape.IsHidden != resourcecompositionshapeOther.IsHidden {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (resourceshape *ResourceShape) GongDiff(stage *Stage, resourceshapeOther *ResourceShape) (diffs []string) {
	// insertion point for field diffs
	if resourceshape.Name != resourceshapeOther.Name {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "Name"))
	}
	if (resourceshape.Resource == nil) != (resourceshapeOther.Resource == nil) {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "Resource"))
	} else if resourceshape.Resource != nil && resourceshapeOther.Resource != nil {
		if resourceshape.Resource != resourceshapeOther.Resource {
			diffs = append(diffs, resourceshape.GongMarshallField(stage, "Resource"))
		}
	}
	if resourceshape.OverideLayoutDirection != resourceshapeOther.OverideLayoutDirection {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "OverideLayoutDirection"))
	}
	if resourceshape.LayoutDirection != resourceshapeOther.LayoutDirection {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "LayoutDirection"))
	}
	if resourceshape.X != resourceshapeOther.X {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "X"))
	}
	if resourceshape.Y != resourceshapeOther.Y {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "Y"))
	}
	if resourceshape.Width != resourceshapeOther.Width {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "Width"))
	}
	if resourceshape.Height != resourceshapeOther.Height {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "Height"))
	}
	if resourceshape.IsHidden != resourceshapeOther.IsHidden {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (resourcetaskshape *ResourceTaskShape) GongDiff(stage *Stage, resourcetaskshapeOther *ResourceTaskShape) (diffs []string) {
	// insertion point for field diffs
	if resourcetaskshape.Name != resourcetaskshapeOther.Name {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "Name"))
	}
	if (resourcetaskshape.Resource == nil) != (resourcetaskshapeOther.Resource == nil) {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "Resource"))
	} else if resourcetaskshape.Resource != nil && resourcetaskshapeOther.Resource != nil {
		if resourcetaskshape.Resource != resourcetaskshapeOther.Resource {
			diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "Resource"))
		}
	}
	if (resourcetaskshape.Task == nil) != (resourcetaskshapeOther.Task == nil) {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "Task"))
	} else if resourcetaskshape.Task != nil && resourcetaskshapeOther.Task != nil {
		if resourcetaskshape.Task != resourcetaskshapeOther.Task {
			diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "Task"))
		}
	}
	if resourcetaskshape.StartRatio != resourcetaskshapeOther.StartRatio {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "StartRatio"))
	}
	if resourcetaskshape.EndRatio != resourcetaskshapeOther.EndRatio {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "EndRatio"))
	}
	if resourcetaskshape.StartOrientation != resourcetaskshapeOther.StartOrientation {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "StartOrientation"))
	}
	if resourcetaskshape.EndOrientation != resourcetaskshapeOther.EndOrientation {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "EndOrientation"))
	}
	if resourcetaskshape.CornerOffsetRatio != resourcetaskshapeOther.CornerOffsetRatio {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if resourcetaskshape.IsHidden != resourcetaskshapeOther.IsHidden {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (task *Task) GongDiff(stage *Stage, taskOther *Task) (diffs []string) {
	// insertion point for field diffs
	if task.Name != taskOther.Name {
		diffs = append(diffs, task.GongMarshallField(stage, "Name"))
	}
	if task.Description != taskOther.Description {
		diffs = append(diffs, task.GongMarshallField(stage, "Description"))
	}
	if task.Start != taskOther.Start {
		diffs = append(diffs, task.GongMarshallField(stage, "Start"))
	}
	if task.End != taskOther.End {
		diffs = append(diffs, task.GongMarshallField(stage, "End"))
	}
	if task.IsAllDay != taskOther.IsAllDay {
		diffs = append(diffs, task.GongMarshallField(stage, "IsAllDay"))
	}
	if task.IsMilestone != taskOther.IsMilestone {
		diffs = append(diffs, task.GongMarshallField(stage, "IsMilestone"))
	}
	PredecessorsDifferent := false
	if len(task.Predecessors) != len(taskOther.Predecessors) {
		PredecessorsDifferent = true
	} else {
		for i := range task.Predecessors {
			if (task.Predecessors[i] == nil) != (taskOther.Predecessors[i] == nil) {
				PredecessorsDifferent = true
				break
			} else if task.Predecessors[i] != nil && taskOther.Predecessors[i] != nil {
				// this is a pointer comparaison
				if task.Predecessors[i] != taskOther.Predecessors[i] {
					PredecessorsDifferent = true
					break
				}
			}
		}
	}
	if PredecessorsDifferent {
		ops := stage.Diff(
			task,
			"Predecessors",
			len(taskOther.Predecessors),
			len(task.Predecessors),
			func(i, j int) bool {
				return taskOther.Predecessors[i] == task.Predecessors[j]
			},
			func(j int) string {
				return task.Predecessors[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if task.DependencyType != taskOther.DependencyType {
		diffs = append(diffs, task.GongMarshallField(stage, "DependencyType"))
	}
	if task.DependencyDurationYears != taskOther.DependencyDurationYears {
		diffs = append(diffs, task.GongMarshallField(stage, "DependencyDurationYears"))
	}
	if task.DependencyDurationMonths != taskOther.DependencyDurationMonths {
		diffs = append(diffs, task.GongMarshallField(stage, "DependencyDurationMonths"))
	}
	if task.DependencyDurationWeeks != taskOther.DependencyDurationWeeks {
		diffs = append(diffs, task.GongMarshallField(stage, "DependencyDurationWeeks"))
	}
	if task.DependencyDurationDays != taskOther.DependencyDurationDays {
		diffs = append(diffs, task.GongMarshallField(stage, "DependencyDurationDays"))
	}
	if task.DependencyDurationHours != taskOther.DependencyDurationHours {
		diffs = append(diffs, task.GongMarshallField(stage, "DependencyDurationHours"))
	}
	if task.DurationYears != taskOther.DurationYears {
		diffs = append(diffs, task.GongMarshallField(stage, "DurationYears"))
	}
	if task.DurationMonths != taskOther.DurationMonths {
		diffs = append(diffs, task.GongMarshallField(stage, "DurationMonths"))
	}
	if task.DurationWeeks != taskOther.DurationWeeks {
		diffs = append(diffs, task.GongMarshallField(stage, "DurationWeeks"))
	}
	if task.DurationDays != taskOther.DurationDays {
		diffs = append(diffs, task.GongMarshallField(stage, "DurationDays"))
	}
	if task.DurationHours != taskOther.DurationHours {
		diffs = append(diffs, task.GongMarshallField(stage, "DurationHours"))
	}
	if task.IsEndDateComputedFromDuration != taskOther.IsEndDateComputedFromDuration {
		diffs = append(diffs, task.GongMarshallField(stage, "IsEndDateComputedFromDuration"))
	}
	InputsDifferent := false
	if len(task.Inputs) != len(taskOther.Inputs) {
		InputsDifferent = true
	} else {
		for i := range task.Inputs {
			if (task.Inputs[i] == nil) != (taskOther.Inputs[i] == nil) {
				InputsDifferent = true
				break
			} else if task.Inputs[i] != nil && taskOther.Inputs[i] != nil {
				// this is a pointer comparaison
				if task.Inputs[i] != taskOther.Inputs[i] {
					InputsDifferent = true
					break
				}
			}
		}
	}
	if InputsDifferent {
		ops := stage.Diff(
			task,
			"Inputs",
			len(taskOther.Inputs),
			len(task.Inputs),
			func(i, j int) bool {
				return taskOther.Inputs[i] == task.Inputs[j]
			},
			func(j int) string {
				return task.Inputs[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	OutputsDifferent := false
	if len(task.Outputs) != len(taskOther.Outputs) {
		OutputsDifferent = true
	} else {
		for i := range task.Outputs {
			if (task.Outputs[i] == nil) != (taskOther.Outputs[i] == nil) {
				OutputsDifferent = true
				break
			} else if task.Outputs[i] != nil && taskOther.Outputs[i] != nil {
				// this is a pointer comparaison
				if task.Outputs[i] != taskOther.Outputs[i] {
					OutputsDifferent = true
					break
				}
			}
		}
	}
	if OutputsDifferent {
		ops := stage.Diff(
			task,
			"Outputs",
			len(taskOther.Outputs),
			len(task.Outputs),
			func(i, j int) bool {
				return taskOther.Outputs[i] == task.Outputs[j]
			},
			func(j int) string {
				return task.Outputs[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	SubTasksDifferent := false
	if len(task.SubTasks) != len(taskOther.SubTasks) {
		SubTasksDifferent = true
	} else {
		for i := range task.SubTasks {
			if (task.SubTasks[i] == nil) != (taskOther.SubTasks[i] == nil) {
				SubTasksDifferent = true
				break
			} else if task.SubTasks[i] != nil && taskOther.SubTasks[i] != nil {
				// this is a pointer comparaison
				if task.SubTasks[i] != taskOther.SubTasks[i] {
					SubTasksDifferent = true
					break
				}
			}
		}
	}
	if SubTasksDifferent {
		ops := stage.Diff(
			task,
			"SubTasks",
			len(taskOther.SubTasks),
			len(task.SubTasks),
			func(i, j int) bool {
				return taskOther.SubTasks[i] == task.SubTasks[j]
			},
			func(j int) string {
				return task.SubTasks[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if task.IsWithCompletion != taskOther.IsWithCompletion {
		diffs = append(diffs, task.GongMarshallField(stage, "IsWithCompletion"))
	}
	if task.Completion != taskOther.Completion {
		diffs = append(diffs, task.GongMarshallField(stage, "Completion"))
	}
	if task.DisplayVerticalBar != taskOther.DisplayVerticalBar {
		diffs = append(diffs, task.GongMarshallField(stage, "DisplayVerticalBar"))
	}
	TaskGroupsToDisplayDifferent := false
	if len(task.TaskGroupsToDisplay) != len(taskOther.TaskGroupsToDisplay) {
		TaskGroupsToDisplayDifferent = true
	} else {
		for i := range task.TaskGroupsToDisplay {
			if (task.TaskGroupsToDisplay[i] == nil) != (taskOther.TaskGroupsToDisplay[i] == nil) {
				TaskGroupsToDisplayDifferent = true
				break
			} else if task.TaskGroupsToDisplay[i] != nil && taskOther.TaskGroupsToDisplay[i] != nil {
				// this is a pointer comparaison
				if task.TaskGroupsToDisplay[i] != taskOther.TaskGroupsToDisplay[i] {
					TaskGroupsToDisplayDifferent = true
					break
				}
			}
		}
	}
	if TaskGroupsToDisplayDifferent {
		ops := stage.Diff(
			task,
			"TaskGroupsToDisplay",
			len(taskOther.TaskGroupsToDisplay),
			len(task.TaskGroupsToDisplay),
			func(i, j int) bool {
				return taskOther.TaskGroupsToDisplay[i] == task.TaskGroupsToDisplay[j]
			},
			func(j int) string {
				return task.TaskGroupsToDisplay[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if task.TextPosition != taskOther.TextPosition {
		diffs = append(diffs, task.GongMarshallField(stage, "TextPosition"))
	}
	if task.XOffset != taskOther.XOffset {
		diffs = append(diffs, task.GongMarshallField(stage, "XOffset"))
	}
	if task.YOffset != taskOther.YOffset {
		diffs = append(diffs, task.GongMarshallField(stage, "YOffset"))
	}
	if task.IsImport != taskOther.IsImport {
		diffs = append(diffs, task.GongMarshallField(stage, "IsImport"))
	}
	if (task.ReferencedTask == nil) != (taskOther.ReferencedTask == nil) {
		diffs = append(diffs, task.GongMarshallField(stage, "ReferencedTask"))
	} else if task.ReferencedTask != nil && taskOther.ReferencedTask != nil {
		if task.ReferencedTask != taskOther.ReferencedTask {
			diffs = append(diffs, task.GongMarshallField(stage, "ReferencedTask"))
		}
	}
	if task.IsInputsNodeExpanded != taskOther.IsInputsNodeExpanded {
		diffs = append(diffs, task.GongMarshallField(stage, "IsInputsNodeExpanded"))
	}
	if task.IsOutputsNodeExpanded != taskOther.IsOutputsNodeExpanded {
		diffs = append(diffs, task.GongMarshallField(stage, "IsOutputsNodeExpanded"))
	}
	if task.ComputedPrefix != taskOther.ComputedPrefix {
		diffs = append(diffs, task.GongMarshallField(stage, "ComputedPrefix"))
	}
	if task.IsExpanded != taskOther.IsExpanded {
		diffs = append(diffs, task.GongMarshallField(stage, "IsExpanded"))
	}
	if task.LayoutDirection != taskOther.LayoutDirection {
		diffs = append(diffs, task.GongMarshallField(stage, "LayoutDirection"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskcompositionshape *TaskCompositionShape) GongDiff(stage *Stage, taskcompositionshapeOther *TaskCompositionShape) (diffs []string) {
	// insertion point for field diffs
	if taskcompositionshape.Name != taskcompositionshapeOther.Name {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "Name"))
	}
	if (taskcompositionshape.Task == nil) != (taskcompositionshapeOther.Task == nil) {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "Task"))
	} else if taskcompositionshape.Task != nil && taskcompositionshapeOther.Task != nil {
		if taskcompositionshape.Task != taskcompositionshapeOther.Task {
			diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "Task"))
		}
	}
	if taskcompositionshape.StartRatio != taskcompositionshapeOther.StartRatio {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "StartRatio"))
	}
	if taskcompositionshape.EndRatio != taskcompositionshapeOther.EndRatio {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "EndRatio"))
	}
	if taskcompositionshape.StartOrientation != taskcompositionshapeOther.StartOrientation {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "StartOrientation"))
	}
	if taskcompositionshape.EndOrientation != taskcompositionshapeOther.EndOrientation {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "EndOrientation"))
	}
	if taskcompositionshape.CornerOffsetRatio != taskcompositionshapeOther.CornerOffsetRatio {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if taskcompositionshape.IsHidden != taskcompositionshapeOther.IsHidden {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskgroup *TaskGroup) GongDiff(stage *Stage, taskgroupOther *TaskGroup) (diffs []string) {
	// insertion point for field diffs
	if taskgroup.Name != taskgroupOther.Name {
		diffs = append(diffs, taskgroup.GongMarshallField(stage, "Name"))
	}
	if taskgroup.ComputedPrefix != taskgroupOther.ComputedPrefix {
		diffs = append(diffs, taskgroup.GongMarshallField(stage, "ComputedPrefix"))
	}
	if taskgroup.IsExpanded != taskgroupOther.IsExpanded {
		diffs = append(diffs, taskgroup.GongMarshallField(stage, "IsExpanded"))
	}
	TasksDifferent := false
	if len(taskgroup.Tasks) != len(taskgroupOther.Tasks) {
		TasksDifferent = true
	} else {
		for i := range taskgroup.Tasks {
			if (taskgroup.Tasks[i] == nil) != (taskgroupOther.Tasks[i] == nil) {
				TasksDifferent = true
				break
			} else if taskgroup.Tasks[i] != nil && taskgroupOther.Tasks[i] != nil {
				// this is a pointer comparaison
				if taskgroup.Tasks[i] != taskgroupOther.Tasks[i] {
					TasksDifferent = true
					break
				}
			}
		}
	}
	if TasksDifferent {
		ops := stage.Diff(
			taskgroup,
			"Tasks",
			len(taskgroupOther.Tasks),
			len(taskgroup.Tasks),
			func(i, j int) bool {
				return taskgroupOther.Tasks[i] == taskgroup.Tasks[j]
			},
			func(j int) string {
				return taskgroup.Tasks[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskgroupshape *TaskGroupShape) GongDiff(stage *Stage, taskgroupshapeOther *TaskGroupShape) (diffs []string) {
	// insertion point for field diffs
	if taskgroupshape.Name != taskgroupshapeOther.Name {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "Name"))
	}
	if (taskgroupshape.TaskGroup == nil) != (taskgroupshapeOther.TaskGroup == nil) {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "TaskGroup"))
	} else if taskgroupshape.TaskGroup != nil && taskgroupshapeOther.TaskGroup != nil {
		if taskgroupshape.TaskGroup != taskgroupshapeOther.TaskGroup {
			diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "TaskGroup"))
		}
	}
	if taskgroupshape.X != taskgroupshapeOther.X {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "X"))
	}
	if taskgroupshape.Y != taskgroupshapeOther.Y {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "Y"))
	}
	if taskgroupshape.Width != taskgroupshapeOther.Width {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "Width"))
	}
	if taskgroupshape.Height != taskgroupshapeOther.Height {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "Height"))
	}
	if taskgroupshape.IsHidden != taskgroupshapeOther.IsHidden {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskinputshape *TaskInputShape) GongDiff(stage *Stage, taskinputshapeOther *TaskInputShape) (diffs []string) {
	// insertion point for field diffs
	if taskinputshape.Name != taskinputshapeOther.Name {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "Name"))
	}
	if (taskinputshape.Product == nil) != (taskinputshapeOther.Product == nil) {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "Product"))
	} else if taskinputshape.Product != nil && taskinputshapeOther.Product != nil {
		if taskinputshape.Product != taskinputshapeOther.Product {
			diffs = append(diffs, taskinputshape.GongMarshallField(stage, "Product"))
		}
	}
	if (taskinputshape.Task == nil) != (taskinputshapeOther.Task == nil) {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "Task"))
	} else if taskinputshape.Task != nil && taskinputshapeOther.Task != nil {
		if taskinputshape.Task != taskinputshapeOther.Task {
			diffs = append(diffs, taskinputshape.GongMarshallField(stage, "Task"))
		}
	}
	if taskinputshape.StartRatio != taskinputshapeOther.StartRatio {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "StartRatio"))
	}
	if taskinputshape.EndRatio != taskinputshapeOther.EndRatio {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "EndRatio"))
	}
	if taskinputshape.StartOrientation != taskinputshapeOther.StartOrientation {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "StartOrientation"))
	}
	if taskinputshape.EndOrientation != taskinputshapeOther.EndOrientation {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "EndOrientation"))
	}
	if taskinputshape.CornerOffsetRatio != taskinputshapeOther.CornerOffsetRatio {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if taskinputshape.IsHidden != taskinputshapeOther.IsHidden {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskoutputshape *TaskOutputShape) GongDiff(stage *Stage, taskoutputshapeOther *TaskOutputShape) (diffs []string) {
	// insertion point for field diffs
	if taskoutputshape.Name != taskoutputshapeOther.Name {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "Name"))
	}
	if (taskoutputshape.Task == nil) != (taskoutputshapeOther.Task == nil) {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "Task"))
	} else if taskoutputshape.Task != nil && taskoutputshapeOther.Task != nil {
		if taskoutputshape.Task != taskoutputshapeOther.Task {
			diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "Task"))
		}
	}
	if (taskoutputshape.Product == nil) != (taskoutputshapeOther.Product == nil) {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "Product"))
	} else if taskoutputshape.Product != nil && taskoutputshapeOther.Product != nil {
		if taskoutputshape.Product != taskoutputshapeOther.Product {
			diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "Product"))
		}
	}
	if taskoutputshape.StartRatio != taskoutputshapeOther.StartRatio {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "StartRatio"))
	}
	if taskoutputshape.EndRatio != taskoutputshapeOther.EndRatio {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "EndRatio"))
	}
	if taskoutputshape.StartOrientation != taskoutputshapeOther.StartOrientation {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "StartOrientation"))
	}
	if taskoutputshape.EndOrientation != taskoutputshapeOther.EndOrientation {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "EndOrientation"))
	}
	if taskoutputshape.CornerOffsetRatio != taskoutputshapeOther.CornerOffsetRatio {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if taskoutputshape.IsHidden != taskoutputshapeOther.IsHidden {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskpredecessorshape *TaskPredecessorShape) GongDiff(stage *Stage, taskpredecessorshapeOther *TaskPredecessorShape) (diffs []string) {
	// insertion point for field diffs
	if taskpredecessorshape.Name != taskpredecessorshapeOther.Name {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "Name"))
	}
	if (taskpredecessorshape.Predecessor == nil) != (taskpredecessorshapeOther.Predecessor == nil) {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "Predecessor"))
	} else if taskpredecessorshape.Predecessor != nil && taskpredecessorshapeOther.Predecessor != nil {
		if taskpredecessorshape.Predecessor != taskpredecessorshapeOther.Predecessor {
			diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "Predecessor"))
		}
	}
	if (taskpredecessorshape.Task == nil) != (taskpredecessorshapeOther.Task == nil) {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "Task"))
	} else if taskpredecessorshape.Task != nil && taskpredecessorshapeOther.Task != nil {
		if taskpredecessorshape.Task != taskpredecessorshapeOther.Task {
			diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "Task"))
		}
	}
	if taskpredecessorshape.StartRatio != taskpredecessorshapeOther.StartRatio {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "StartRatio"))
	}
	if taskpredecessorshape.EndRatio != taskpredecessorshapeOther.EndRatio {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "EndRatio"))
	}
	if taskpredecessorshape.StartOrientation != taskpredecessorshapeOther.StartOrientation {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "StartOrientation"))
	}
	if taskpredecessorshape.EndOrientation != taskpredecessorshapeOther.EndOrientation {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "EndOrientation"))
	}
	if taskpredecessorshape.CornerOffsetRatio != taskpredecessorshapeOther.CornerOffsetRatio {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if taskpredecessorshape.IsHidden != taskpredecessorshapeOther.IsHidden {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (taskshape *TaskShape) GongDiff(stage *Stage, taskshapeOther *TaskShape) (diffs []string) {
	// insertion point for field diffs
	if taskshape.Name != taskshapeOther.Name {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Name"))
	}
	if (taskshape.Task == nil) != (taskshapeOther.Task == nil) {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Task"))
	} else if taskshape.Task != nil && taskshapeOther.Task != nil {
		if taskshape.Task != taskshapeOther.Task {
			diffs = append(diffs, taskshape.GongMarshallField(stage, "Task"))
		}
	}
	if taskshape.IsShowDate != taskshapeOther.IsShowDate {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "IsShowDate"))
	}
	if taskshape.OverideLayoutDirection != taskshapeOther.OverideLayoutDirection {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "OverideLayoutDirection"))
	}
	if taskshape.LayoutDirection != taskshapeOther.LayoutDirection {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "LayoutDirection"))
	}
	if taskshape.X != taskshapeOther.X {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "X"))
	}
	if taskshape.Y != taskshapeOther.Y {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Y"))
	}
	if taskshape.Width != taskshapeOther.Width {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Width"))
	}
	if taskshape.Height != taskshapeOther.Height {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Height"))
	}
	if taskshape.IsHidden != taskshapeOther.IsHidden {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
