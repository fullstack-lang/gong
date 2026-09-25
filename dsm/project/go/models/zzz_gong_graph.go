// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (diagram *Diagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Diagrams[diagram]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (noteproductshape *NoteProductShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteProductShapes[noteproductshape]
	return ok
}

func (noteresourceshape *NoteResourceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteResourceShapes[noteresourceshape]
	return ok
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteShapes[noteshape]
	return ok
}

func (notetaskshape *NoteTaskShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteTaskShapes[notetaskshape]
	return ok
}

func (product *Product) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Products[product]
	return ok
}

func (productcompositionshape *ProductCompositionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ProductCompositionShapes[productcompositionshape]
	return ok
}

func (productreferenceshape *ProductReferenceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ProductReferenceShapes[productreferenceshape]
	return ok
}

func (productshape *ProductShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ProductShapes[productshape]
	return ok
}

func (resource *Resource) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Resources[resource]
	return ok
}

func (resourcecompositionshape *ResourceCompositionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ResourceCompositionShapes[resourcecompositionshape]
	return ok
}

func (resourceshape *ResourceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ResourceShapes[resourceshape]
	return ok
}

func (resourcetaskshape *ResourceTaskShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ResourceTaskShapes[resourcetaskshape]
	return ok
}

func (task *Task) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Tasks[task]
	return ok
}

func (taskcompositionshape *TaskCompositionShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskCompositionShapes[taskcompositionshape]
	return ok
}

func (taskgroup *TaskGroup) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskGroups[taskgroup]
	return ok
}

func (taskgroupshape *TaskGroupShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskGroupShapes[taskgroupshape]
	return ok
}

func (taskinputshape *TaskInputShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskInputShapes[taskinputshape]
	return ok
}

func (taskoutputshape *TaskOutputShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskOutputShapes[taskoutputshape]
	return ok
}

func (taskpredecessorshape *TaskPredecessorShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskPredecessorShapes[taskpredecessorshape]
	return ok
}

func (taskshape *TaskShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TaskShapes[taskshape]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (diagram *Diagram) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	diagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	noteproductshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteproductshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	noteresourceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteresourceshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	noteshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteshapeFrom)
	if alreadyCopied {
		return
	}
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {
	var alreadyCopied bool
	notetaskshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notetaskshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	productTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, productFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	productcompositionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, productcompositionshapeFrom)
	if alreadyCopied {
		return
	}
	productcompositionshapeFrom.GongCopyBasicFields(productcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshapeFrom.Product != nil {
		productcompositionshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, productcompositionshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchProductReferenceShape(mapOrigCopy map[any]any, productreferenceshapeFrom *ProductReferenceShape) (productreferenceshapeTo *ProductReferenceShape) {
	var alreadyCopied bool
	productreferenceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, productreferenceshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	productshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, productshapeFrom)
	if alreadyCopied {
		return
	}
	productshapeFrom.GongCopyBasicFields(productshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productshapeFrom.Product != nil {
		productshapeTo.Product = GongCopyBranchProduct(mapOrigCopy, productshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {
	var alreadyCopied bool
	resourceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, resourceFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	resourcecompositionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, resourcecompositionshapeFrom)
	if alreadyCopied {
		return
	}
	resourcecompositionshapeFrom.GongCopyBasicFields(resourcecompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourcecompositionshapeFrom.Resource != nil {
		resourcecompositionshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, resourcecompositionshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResourceShape(mapOrigCopy map[any]any, resourceshapeFrom *ResourceShape) (resourceshapeTo *ResourceShape) {
	var alreadyCopied bool
	resourceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, resourceshapeFrom)
	if alreadyCopied {
		return
	}
	resourceshapeFrom.GongCopyBasicFields(resourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourceshapeFrom.Resource != nil {
		resourceshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, resourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResourceTaskShape(mapOrigCopy map[any]any, resourcetaskshapeFrom *ResourceTaskShape) (resourcetaskshapeTo *ResourceTaskShape) {
	var alreadyCopied bool
	resourcetaskshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, resourcetaskshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	taskTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	taskcompositionshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskcompositionshapeFrom)
	if alreadyCopied {
		return
	}
	taskcompositionshapeFrom.GongCopyBasicFields(taskcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshapeFrom.Task != nil {
		taskcompositionshapeTo.Task = GongCopyBranchTask(mapOrigCopy, taskcompositionshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskGroup(mapOrigCopy map[any]any, taskgroupFrom *TaskGroup) (taskgroupTo *TaskGroup) {
	var alreadyCopied bool
	taskgroupTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskgroupFrom)
	if alreadyCopied {
		return
	}
	taskgroupFrom.GongCopyBasicFields(taskgroupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskgroupFrom.Tasks {
		taskgroupTo.Tasks = append(taskgroupTo.Tasks, GongCopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func GongCopyBranchTaskGroupShape(mapOrigCopy map[any]any, taskgroupshapeFrom *TaskGroupShape) (taskgroupshapeTo *TaskGroupShape) {
	var alreadyCopied bool
	taskgroupshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskgroupshapeFrom)
	if alreadyCopied {
		return
	}
	taskgroupshapeFrom.GongCopyBasicFields(taskgroupshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskgroupshapeFrom.TaskGroup != nil {
		taskgroupshapeTo.TaskGroup = GongCopyBranchTaskGroup(mapOrigCopy, taskgroupshapeFrom.TaskGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTaskInputShape(mapOrigCopy map[any]any, taskinputshapeFrom *TaskInputShape) (taskinputshapeTo *TaskInputShape) {
	var alreadyCopied bool
	taskinputshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskinputshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	taskoutputshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskoutputshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	taskpredecessorshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskpredecessorshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	taskshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, taskshapeFrom)
	if alreadyCopied {
		return
	}
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

// insertion point for unstage branch per struct
func (diagram *Diagram) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Product_Shapes, stage.ProductShapes_reference, instance.Product_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ProductsWhoseNodeIsExpanded, stage.Products_reference, instance.ProductsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ProductComposition_Shapes, stage.ProductCompositionShapes_reference, instance.ProductComposition_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ProductReference_Shapes, stage.ProductReferenceShapes_reference, instance.ProductReference_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Task_Shapes, stage.TaskShapes_reference, instance.Task_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TasksWhoseNodeIsExpanded, stage.Tasks_reference, instance.TasksWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TasksWhoseInputNodeIsExpanded, stage.Tasks_reference, instance.TasksWhoseInputNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TasksWhoseOutputNodeIsExpanded, stage.Tasks_reference, instance.TasksWhoseOutputNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TasksWhosePredecessorNodeIsExpanded, stage.Tasks_reference, instance.TasksWhosePredecessorNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskGroupShapes, stage.TaskGroupShapes_reference, instance.TaskGroupShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskGroupsWhoseNodeIsExpanded, stage.TaskGroups_reference, instance.TaskGroupsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskComposition_Shapes, stage.TaskCompositionShapes_reference, instance.TaskComposition_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskInputShapes, stage.TaskInputShapes_reference, instance.TaskInputShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskOutputShapes, stage.TaskOutputShapes_reference, instance.TaskOutputShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskPredecessorShapes, stage.TaskPredecessorShapes_reference, instance.TaskPredecessorShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note_Shapes, stage.NoteShapes_reference, instance.Note_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteProductShapes, stage.NoteProductShapes_reference, instance.NoteProductShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteTaskShapes, stage.NoteTaskShapes_reference, instance.NoteTaskShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NoteResourceShapes, stage.NoteResourceShapes_reference, instance.NoteResourceShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Resource_Shapes, stage.ResourceShapes_reference, instance.Resource_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ResourcesWhoseNodeIsExpanded, stage.Resources_reference, instance.ResourcesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ResourceComposition_Shapes, stage.ResourceCompositionShapes_reference, instance.ResourceComposition_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ResourceTaskShapes, stage.ResourceTaskShapes_reference, instance.ResourceTaskShapes)
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootProducts, stage.Products_reference, instance.RootProducts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootTasks, stage.Tasks_reference, instance.RootTasks)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootTaskGroups, stage.TaskGroups_reference, instance.RootTaskGroups)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootResources, stage.Resources_reference, instance.RootResources)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Notes, stage.Notes_reference, instance.Notes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Diagrams, stage.Diagrams_reference, instance.Diagrams)
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Products, stage.Products_reference, instance.Products)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tasks, stage.Tasks_reference, instance.Tasks)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Resources, stage.Resources_reference, instance.Resources)
}

func (reference *NoteProductShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteProductShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Product, stage.Products_reference, instance.Product)
	// insertion point for slice of pointers field
}

func (reference *NoteResourceShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteResourceShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Resource, stage.Resources_reference, instance.Resource)
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	// insertion point for slice of pointers field
}

func (reference *NoteTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteTaskShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

func (reference *Product) GongReconstructPointersFromReferences(stage *Stage, instance *Product) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ReferencedProduct, stage.Products_reference, instance.ReferencedProduct)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubProducts, stage.Products_reference, instance.SubProducts)
}

func (reference *ProductCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductCompositionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Product, stage.Products_reference, instance.Product)
	// insertion point for slice of pointers field
}

func (reference *ProductReferenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductReferenceShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Product, stage.Products_reference, instance.Product)
	__gong__reconstructPointer(&reference.ReferencedProduct, stage.Products_reference, instance.ReferencedProduct)
	// insertion point for slice of pointers field
}

func (reference *ProductShape) GongReconstructPointersFromReferences(stage *Stage, instance *ProductShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Product, stage.Products_reference, instance.Product)
	// insertion point for slice of pointers field
}

func (reference *Resource) GongReconstructPointersFromReferences(stage *Stage, instance *Resource) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ReferencedResource, stage.Resources_reference, instance.ReferencedResource)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tasks, stage.Tasks_reference, instance.Tasks)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubResources, stage.Resources_reference, instance.SubResources)
}

func (reference *ResourceCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *ResourceCompositionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Resource, stage.Resources_reference, instance.Resource)
	// insertion point for slice of pointers field
}

func (reference *ResourceShape) GongReconstructPointersFromReferences(stage *Stage, instance *ResourceShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Resource, stage.Resources_reference, instance.Resource)
	// insertion point for slice of pointers field
}

func (reference *ResourceTaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *ResourceTaskShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Resource, stage.Resources_reference, instance.Resource)
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

func (reference *Task) GongReconstructPointersFromReferences(stage *Stage, instance *Task) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ReferencedTask, stage.Tasks_reference, instance.ReferencedTask)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Predecessors, stage.Tasks_reference, instance.Predecessors)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Inputs, stage.Products_reference, instance.Inputs)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Outputs, stage.Products_reference, instance.Outputs)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubTasks, stage.Tasks_reference, instance.SubTasks)
	__gong__reconstructSliceOfPointersFromReferences(&reference.TaskGroupsToDisplay, stage.TaskGroups_reference, instance.TaskGroupsToDisplay)
}

func (reference *TaskCompositionShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskCompositionShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

func (reference *TaskGroup) GongReconstructPointersFromReferences(stage *Stage, instance *TaskGroup) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Tasks, stage.Tasks_reference, instance.Tasks)
}

func (reference *TaskGroupShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskGroupShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TaskGroup, stage.TaskGroups_reference, instance.TaskGroup)
	// insertion point for slice of pointers field
}

func (reference *TaskInputShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskInputShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Product, stage.Products_reference, instance.Product)
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

func (reference *TaskOutputShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskOutputShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	__gong__reconstructPointer(&reference.Product, stage.Products_reference, instance.Product)
	// insertion point for slice of pointers field
}

func (reference *TaskPredecessorShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskPredecessorShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Predecessor, stage.Tasks_reference, instance.Predecessor)
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

func (reference *TaskShape) GongReconstructPointersFromReferences(stage *Stage, instance *TaskShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Task, stage.Tasks_reference, instance.Task)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Product_Shapes, stage.ProductShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ProductsWhoseNodeIsExpanded, stage.Products_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ProductComposition_Shapes, stage.ProductCompositionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ProductReference_Shapes, stage.ProductReferenceShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Task_Shapes, stage.TaskShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TasksWhoseNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TasksWhoseInputNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TasksWhoseOutputNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TasksWhosePredecessorNodeIsExpanded, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskGroupShapes, stage.TaskGroupShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskGroupsWhoseNodeIsExpanded, stage.TaskGroups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskComposition_Shapes, stage.TaskCompositionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskInputShapes, stage.TaskInputShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskOutputShapes, stage.TaskOutputShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskPredecessorShapes, stage.TaskPredecessorShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note_Shapes, stage.NoteShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteProductShapes, stage.NoteProductShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteTaskShapes, stage.NoteTaskShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NoteResourceShapes, stage.NoteResourceShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Resource_Shapes, stage.ResourceShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ResourcesWhoseNodeIsExpanded, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ResourceComposition_Shapes, stage.ResourceCompositionShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ResourceTaskShapes, stage.ResourceTaskShapes_instance)
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootProducts, stage.Products_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootTasks, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootTaskGroups, stage.TaskGroups_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootResources, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Notes, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Diagrams, stage.Diagrams_instance)
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Products, stage.Products_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tasks, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Resources, stage.Resources_instance)
}

func (reference *NoteProductShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Product, stage.Products_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Resource, stage.Resources_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
	// insertion point for slice of pointers fields
}

func (reference *Product) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ReferencedProduct, stage.Products_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubProducts, stage.Products_instance)
}

func (reference *ProductCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Product, stage.Products_instance)
	// insertion point for slice of pointers fields
}

func (reference *ProductReferenceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Product, stage.Products_instance)
	__gong__reconstructPointerFromInstance(&reference.ReferencedProduct, stage.Products_instance)
	// insertion point for slice of pointers fields
}

func (reference *ProductShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Product, stage.Products_instance)
	// insertion point for slice of pointers fields
}

func (reference *Resource) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ReferencedResource, stage.Resources_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tasks, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubResources, stage.Resources_instance)
}

func (reference *ResourceCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Resource, stage.Resources_instance)
	// insertion point for slice of pointers fields
}

func (reference *ResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Resource, stage.Resources_instance)
	// insertion point for slice of pointers fields
}

func (reference *ResourceTaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Resource, stage.Resources_instance)
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
	// insertion point for slice of pointers fields
}

func (reference *Task) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ReferencedTask, stage.Tasks_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Predecessors, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Inputs, stage.Products_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Outputs, stage.Products_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubTasks, stage.Tasks_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.TaskGroupsToDisplay, stage.TaskGroups_instance)
}

func (reference *TaskCompositionShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
	// insertion point for slice of pointers fields
}

func (reference *TaskGroup) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Tasks, stage.Tasks_instance)
}

func (reference *TaskGroupShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TaskGroup, stage.TaskGroups_instance)
	// insertion point for slice of pointers fields
}

func (reference *TaskInputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Product, stage.Products_instance)
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
	// insertion point for slice of pointers fields
}

func (reference *TaskOutputShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
	__gong__reconstructPointerFromInstance(&reference.Product, stage.Products_instance)
	// insertion point for slice of pointers fields
}

func (reference *TaskPredecessorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Predecessor, stage.Tasks_instance)
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
	// insertion point for slice of pointers fields
}

func (reference *TaskShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Task, stage.Tasks_instance)
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
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Product_Shapes", diagramOther.Product_Shapes, diagram.Product_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ProductsWhoseNodeIsExpanded", diagramOther.ProductsWhoseNodeIsExpanded, diagram.ProductsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsPBSNodeExpanded != diagramOther.IsPBSNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsPBSNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ProductComposition_Shapes", diagramOther.ProductComposition_Shapes, diagram.ProductComposition_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ProductReference_Shapes", diagramOther.ProductReference_Shapes, diagram.ProductReference_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsWBSNodeExpanded != diagramOther.IsWBSNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsWBSNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Task_Shapes", diagramOther.Task_Shapes, diagram.Task_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TasksWhoseNodeIsExpanded", diagramOther.TasksWhoseNodeIsExpanded, diagram.TasksWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TasksWhoseInputNodeIsExpanded", diagramOther.TasksWhoseInputNodeIsExpanded, diagram.TasksWhoseInputNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TasksWhoseOutputNodeIsExpanded", diagramOther.TasksWhoseOutputNodeIsExpanded, diagram.TasksWhoseOutputNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TasksWhosePredecessorNodeIsExpanded", diagramOther.TasksWhosePredecessorNodeIsExpanded, diagram.TasksWhosePredecessorNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsTaskGroupsNodeExpanded != diagramOther.IsTaskGroupsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsTaskGroupsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TaskGroupShapes", diagramOther.TaskGroupShapes, diagram.TaskGroupShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TaskGroupsWhoseNodeIsExpanded", diagramOther.TaskGroupsWhoseNodeIsExpanded, diagram.TaskGroupsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TaskComposition_Shapes", diagramOther.TaskComposition_Shapes, diagram.TaskComposition_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TaskInputShapes", diagramOther.TaskInputShapes, diagram.TaskInputShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TaskOutputShapes", diagramOther.TaskOutputShapes, diagram.TaskOutputShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "TaskPredecessorShapes", diagramOther.TaskPredecessorShapes, diagram.TaskPredecessorShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Note_Shapes", diagramOther.Note_Shapes, diagram.Note_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NotesWhoseNodeIsExpanded", diagramOther.NotesWhoseNodeIsExpanded, diagram.NotesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsNotesNodeExpanded != diagramOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NoteProductShapes", diagramOther.NoteProductShapes, diagram.NoteProductShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NoteTaskShapes", diagramOther.NoteTaskShapes, diagram.NoteTaskShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "NoteResourceShapes", diagramOther.NoteResourceShapes, diagram.NoteResourceShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "Resource_Shapes", diagramOther.Resource_Shapes, diagram.Resource_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ResourcesWhoseNodeIsExpanded", diagramOther.ResourcesWhoseNodeIsExpanded, diagram.ResourcesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagram.IsResourcesNodeExpanded != diagramOther.IsResourcesNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ResourceComposition_Shapes", diagramOther.ResourceComposition_Shapes, diagram.ResourceComposition_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ResourceTaskShapes", diagramOther.ResourceTaskShapes, diagram.ResourceTaskShapes); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, library, "RootProducts", libraryOther.RootProducts, library.RootProducts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootTasks", libraryOther.RootTasks, library.RootTasks); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootTaskGroups", libraryOther.RootTaskGroups, library.RootTaskGroups); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootResources", libraryOther.RootResources, library.RootResources); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "Notes", libraryOther.Notes, library.Notes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "Diagrams", libraryOther.Diagrams, library.Diagrams); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, note, "Products", noteOther.Products, note.Products); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Tasks", noteOther.Tasks, note.Tasks); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Resources", noteOther.Resources, note.Resources); ops != "" {
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
	if noteproductshape.Note != noteproductshapeOther.Note {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Note"))
	}
	if noteproductshape.Product != noteproductshapeOther.Product {
		diffs = append(diffs, noteproductshape.GongMarshallField(stage, "Product"))
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
	if noteresourceshape.Note != noteresourceshapeOther.Note {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "Note"))
	}
	if noteresourceshape.Resource != noteresourceshapeOther.Resource {
		diffs = append(diffs, noteresourceshape.GongMarshallField(stage, "Resource"))
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
	if noteshape.Note != noteshapeOther.Note {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
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
	if notetaskshape.Note != notetaskshapeOther.Note {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Note"))
	}
	if notetaskshape.Task != notetaskshapeOther.Task {
		diffs = append(diffs, notetaskshape.GongMarshallField(stage, "Task"))
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
	if ops := __gong__diffSliceOfPointers(stage, product, "SubProducts", productOther.SubProducts, product.SubProducts); ops != "" {
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
	if product.ReferencedProduct != productOther.ReferencedProduct {
		diffs = append(diffs, product.GongMarshallField(stage, "ReferencedProduct"))
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
	if productcompositionshape.Product != productcompositionshapeOther.Product {
		diffs = append(diffs, productcompositionshape.GongMarshallField(stage, "Product"))
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
	if productreferenceshape.Product != productreferenceshapeOther.Product {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "Product"))
	}
	if productreferenceshape.ReferencedProduct != productreferenceshapeOther.ReferencedProduct {
		diffs = append(diffs, productreferenceshape.GongMarshallField(stage, "ReferencedProduct"))
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
	if productshape.Product != productshapeOther.Product {
		diffs = append(diffs, productshape.GongMarshallField(stage, "Product"))
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
	if ops := __gong__diffSliceOfPointers(stage, resource, "Tasks", resourceOther.Tasks, resource.Tasks); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, resource, "SubResources", resourceOther.SubResources, resource.SubResources); ops != "" {
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
	if resource.ReferencedResource != resourceOther.ReferencedResource {
		diffs = append(diffs, resource.GongMarshallField(stage, "ReferencedResource"))
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
	if resourcecompositionshape.Resource != resourcecompositionshapeOther.Resource {
		diffs = append(diffs, resourcecompositionshape.GongMarshallField(stage, "Resource"))
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
	if resourceshape.Resource != resourceshapeOther.Resource {
		diffs = append(diffs, resourceshape.GongMarshallField(stage, "Resource"))
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
	if resourcetaskshape.Resource != resourcetaskshapeOther.Resource {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "Resource"))
	}
	if resourcetaskshape.Task != resourcetaskshapeOther.Task {
		diffs = append(diffs, resourcetaskshape.GongMarshallField(stage, "Task"))
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
	if ops := __gong__diffSliceOfPointers(stage, task, "Predecessors", taskOther.Predecessors, task.Predecessors); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, task, "Inputs", taskOther.Inputs, task.Inputs); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, task, "Outputs", taskOther.Outputs, task.Outputs); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, task, "SubTasks", taskOther.SubTasks, task.SubTasks); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, task, "TaskGroupsToDisplay", taskOther.TaskGroupsToDisplay, task.TaskGroupsToDisplay); ops != "" {
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
	if task.ReferencedTask != taskOther.ReferencedTask {
		diffs = append(diffs, task.GongMarshallField(stage, "ReferencedTask"))
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
	if taskcompositionshape.Task != taskcompositionshapeOther.Task {
		diffs = append(diffs, taskcompositionshape.GongMarshallField(stage, "Task"))
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
	if ops := __gong__diffSliceOfPointers(stage, taskgroup, "Tasks", taskgroupOther.Tasks, taskgroup.Tasks); ops != "" {
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
	if taskgroupshape.TaskGroup != taskgroupshapeOther.TaskGroup {
		diffs = append(diffs, taskgroupshape.GongMarshallField(stage, "TaskGroup"))
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
	if taskinputshape.Product != taskinputshapeOther.Product {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "Product"))
	}
	if taskinputshape.Task != taskinputshapeOther.Task {
		diffs = append(diffs, taskinputshape.GongMarshallField(stage, "Task"))
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
	if taskoutputshape.Task != taskoutputshapeOther.Task {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "Task"))
	}
	if taskoutputshape.Product != taskoutputshapeOther.Product {
		diffs = append(diffs, taskoutputshape.GongMarshallField(stage, "Product"))
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
	if taskpredecessorshape.Predecessor != taskpredecessorshapeOther.Predecessor {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "Predecessor"))
	}
	if taskpredecessorshape.Task != taskpredecessorshapeOther.Task {
		diffs = append(diffs, taskpredecessorshape.GongMarshallField(stage, "Task"))
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
	if taskshape.Task != taskshapeOther.Task {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "Task"))
	}
	if taskshape.IsShowDate != taskshapeOther.IsShowDate {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "IsShowDate"))
	}
	if taskshape.VerticalOffset != taskshapeOther.VerticalOffset {
		diffs = append(diffs, taskshape.GongMarshallField(stage, "VerticalOffset"))
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

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
