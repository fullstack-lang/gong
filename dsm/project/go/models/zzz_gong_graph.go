// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteProductShape:
		ok = stage.IsStagedNoteProductShape(target)

	case *NoteResourceShape:
		ok = stage.IsStagedNoteResourceShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteTaskShape:
		ok = stage.IsStagedNoteTaskShape(target)

	case *Product:
		ok = stage.IsStagedProduct(target)

	case *ProductCompositionShape:
		ok = stage.IsStagedProductCompositionShape(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Resource:
		ok = stage.IsStagedResource(target)

	case *ResourceCompositionShape:
		ok = stage.IsStagedResourceCompositionShape(target)

	case *ResourceShape:
		ok = stage.IsStagedResourceShape(target)

	case *ResourceTaskShape:
		ok = stage.IsStagedResourceTaskShape(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	case *TaskCompositionShape:
		ok = stage.IsStagedTaskCompositionShape(target)

	case *TaskGroup:
		ok = stage.IsStagedTaskGroup(target)

	case *TaskGroupShape:
		ok = stage.IsStagedTaskGroupShape(target)

	case *TaskInputShape:
		ok = stage.IsStagedTaskInputShape(target)

	case *TaskOutputShape:
		ok = stage.IsStagedTaskOutputShape(target)

	case *TaskShape:
		ok = stage.IsStagedTaskShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NoteProductShape:
		ok = stage.IsStagedNoteProductShape(target)

	case *NoteResourceShape:
		ok = stage.IsStagedNoteResourceShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *NoteTaskShape:
		ok = stage.IsStagedNoteTaskShape(target)

	case *Product:
		ok = stage.IsStagedProduct(target)

	case *ProductCompositionShape:
		ok = stage.IsStagedProductCompositionShape(target)

	case *ProductShape:
		ok = stage.IsStagedProductShape(target)

	case *Resource:
		ok = stage.IsStagedResource(target)

	case *ResourceCompositionShape:
		ok = stage.IsStagedResourceCompositionShape(target)

	case *ResourceShape:
		ok = stage.IsStagedResourceShape(target)

	case *ResourceTaskShape:
		ok = stage.IsStagedResourceTaskShape(target)

	case *Task:
		ok = stage.IsStagedTask(target)

	case *TaskCompositionShape:
		ok = stage.IsStagedTaskCompositionShape(target)

	case *TaskGroup:
		ok = stage.IsStagedTaskGroup(target)

	case *TaskGroupShape:
		ok = stage.IsStagedTaskGroupShape(target)

	case *TaskInputShape:
		ok = stage.IsStagedTaskInputShape(target)

	case *TaskOutputShape:
		ok = stage.IsStagedTaskOutputShape(target)

	case *TaskShape:
		ok = stage.IsStagedTaskShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNoteProductShape(noteproductshape *NoteProductShape) (ok bool) {

	_, ok = stage.NoteProductShapes[noteproductshape]

	return
}

func (stage *Stage) IsStagedNoteResourceShape(noteresourceshape *NoteResourceShape) (ok bool) {

	_, ok = stage.NoteResourceShapes[noteresourceshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteTaskShape(notetaskshape *NoteTaskShape) (ok bool) {

	_, ok = stage.NoteTaskShapes[notetaskshape]

	return
}

func (stage *Stage) IsStagedProduct(product *Product) (ok bool) {

	_, ok = stage.Products[product]

	return
}

func (stage *Stage) IsStagedProductCompositionShape(productcompositionshape *ProductCompositionShape) (ok bool) {

	_, ok = stage.ProductCompositionShapes[productcompositionshape]

	return
}

func (stage *Stage) IsStagedProductShape(productshape *ProductShape) (ok bool) {

	_, ok = stage.ProductShapes[productshape]

	return
}

func (stage *Stage) IsStagedResource(resource *Resource) (ok bool) {

	_, ok = stage.Resources[resource]

	return
}

func (stage *Stage) IsStagedResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape) (ok bool) {

	_, ok = stage.ResourceCompositionShapes[resourcecompositionshape]

	return
}

func (stage *Stage) IsStagedResourceShape(resourceshape *ResourceShape) (ok bool) {

	_, ok = stage.ResourceShapes[resourceshape]

	return
}

func (stage *Stage) IsStagedResourceTaskShape(resourcetaskshape *ResourceTaskShape) (ok bool) {

	_, ok = stage.ResourceTaskShapes[resourcetaskshape]

	return
}

func (stage *Stage) IsStagedTask(task *Task) (ok bool) {

	_, ok = stage.Tasks[task]

	return
}

func (stage *Stage) IsStagedTaskCompositionShape(taskcompositionshape *TaskCompositionShape) (ok bool) {

	_, ok = stage.TaskCompositionShapes[taskcompositionshape]

	return
}

func (stage *Stage) IsStagedTaskGroup(taskgroup *TaskGroup) (ok bool) {

	_, ok = stage.TaskGroups[taskgroup]

	return
}

func (stage *Stage) IsStagedTaskGroupShape(taskgroupshape *TaskGroupShape) (ok bool) {

	_, ok = stage.TaskGroupShapes[taskgroupshape]

	return
}

func (stage *Stage) IsStagedTaskInputShape(taskinputshape *TaskInputShape) (ok bool) {

	_, ok = stage.TaskInputShapes[taskinputshape]

	return
}

func (stage *Stage) IsStagedTaskOutputShape(taskoutputshape *TaskOutputShape) (ok bool) {

	_, ok = stage.TaskOutputShapes[taskoutputshape]

	return
}

func (stage *Stage) IsStagedTaskShape(taskshape *TaskShape) (ok bool) {

	_, ok = stage.TaskShapes[taskshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Note:
		stage.StageBranchNote(target)

	case *NoteProductShape:
		stage.StageBranchNoteProductShape(target)

	case *NoteResourceShape:
		stage.StageBranchNoteResourceShape(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *NoteTaskShape:
		stage.StageBranchNoteTaskShape(target)

	case *Product:
		stage.StageBranchProduct(target)

	case *ProductCompositionShape:
		stage.StageBranchProductCompositionShape(target)

	case *ProductShape:
		stage.StageBranchProductShape(target)

	case *Resource:
		stage.StageBranchResource(target)

	case *ResourceCompositionShape:
		stage.StageBranchResourceCompositionShape(target)

	case *ResourceShape:
		stage.StageBranchResourceShape(target)

	case *ResourceTaskShape:
		stage.StageBranchResourceTaskShape(target)

	case *Task:
		stage.StageBranchTask(target)

	case *TaskCompositionShape:
		stage.StageBranchTaskCompositionShape(target)

	case *TaskGroup:
		stage.StageBranchTaskGroup(target)

	case *TaskGroupShape:
		stage.StageBranchTaskGroupShape(target)

	case *TaskInputShape:
		stage.StageBranchTaskInputShape(target)

	case *TaskOutputShape:
		stage.StageBranchTaskOutputShape(target)

	case *TaskShape:
		stage.StageBranchTaskShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if IsStaged(stage, diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagram.Product_Shapes {
		StageBranch(stage, _productshape)
	}
	for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
		StageBranch(stage, _product)
	}
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		StageBranch(stage, _productcompositionshape)
	}
	for _, _taskshape := range diagram.Task_Shapes {
		StageBranch(stage, _taskshape)
	}
	for _, _task := range diagram.TasksWhoseNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _taskgroupshape := range diagram.TaskGroupShapes {
		StageBranch(stage, _taskgroupshape)
	}
	for _, _taskgroup := range diagram.TaskGroupsWhoseNodeIsExpanded {
		StageBranch(stage, _taskgroup)
	}
	for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
		StageBranch(stage, _taskcompositionshape)
	}
	for _, _taskinputshape := range diagram.TaskInputShapes {
		StageBranch(stage, _taskinputshape)
	}
	for _, _taskoutputshape := range diagram.TaskOutputShapes {
		StageBranch(stage, _taskoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		StageBranch(stage, _noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}
	for _, _noteproductshape := range diagram.NoteProductShapes {
		StageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		StageBranch(stage, _notetaskshape)
	}
	for _, _noteresourceshape := range diagram.NoteResourceShapes {
		StageBranch(stage, _noteresourceshape)
	}
	for _, _resourceshape := range diagram.Resource_Shapes {
		StageBranch(stage, _resourceshape)
	}
	for _, _resource := range diagram.ResourcesWhoseNodeIsExpanded {
		StageBranch(stage, _resource)
	}
	for _, _resourcecompositionshape := range diagram.ResourceComposition_Shapes {
		StageBranch(stage, _resourcecompositionshape)
	}
	for _, _resourcetaskshape := range diagram.ResourceTaskShapes {
		StageBranch(stage, _resourcetaskshape)
	}

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}
	for _, _product := range library.RootProducts {
		StageBranch(stage, _product)
	}
	for _, _task := range library.RootTasks {
		StageBranch(stage, _task)
	}
	for _, _taskgroup := range library.RootTaskGroups {
		StageBranch(stage, _taskgroup)
	}
	for _, _resource := range library.RootResources {
		StageBranch(stage, _resource)
	}
	for _, _note := range library.Notes {
		StageBranch(stage, _note)
	}
	for _, _diagram := range library.Diagrams {
		StageBranch(stage, _diagram)
	}

}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range note.Products {
		StageBranch(stage, _product)
	}
	for _, _task := range note.Tasks {
		StageBranch(stage, _task)
	}
	for _, _resource := range note.Resources {
		StageBranch(stage, _resource)
	}

}

func (stage *Stage) StageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if IsStaged(stage, noteproductshape) {
		return
	}

	noteproductshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		StageBranch(stage, noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		StageBranch(stage, noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteResourceShape(noteresourceshape *NoteResourceShape) {

	// check if instance is already staged
	if IsStaged(stage, noteresourceshape) {
		return
	}

	noteresourceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteresourceshape.Note != nil {
		StageBranch(stage, noteresourceshape.Note)
	}
	if noteresourceshape.Resource != nil {
		StageBranch(stage, noteresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		StageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if IsStaged(stage, notetaskshape) {
		return
	}

	notetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		StageBranch(stage, notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		StageBranch(stage, notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchProduct(product *Product) {

	// check if instance is already staged
	if IsStaged(stage, product) {
		return
	}

	product.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if product.ReferencedProduct != nil {
		StageBranch(stage, product.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range product.SubProducts {
		StageBranch(stage, _product)
	}

}

func (stage *Stage) StageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, productcompositionshape) {
		return
	}

	productcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		StageBranch(stage, productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if IsStaged(stage, productshape) {
		return
	}

	productshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		StageBranch(stage, productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchResource(resource *Resource) {

	// check if instance is already staged
	if IsStaged(stage, resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resource.ReferencedResource != nil {
		StageBranch(stage, resource.ReferencedResource)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range resource.Tasks {
		StageBranch(stage, _task)
	}
	for _, _resource := range resource.SubResources {
		StageBranch(stage, _resource)
	}

}

func (stage *Stage) StageBranchResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, resourcecompositionshape) {
		return
	}

	resourcecompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcecompositionshape.Resource != nil {
		StageBranch(stage, resourcecompositionshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchResourceShape(resourceshape *ResourceShape) {

	// check if instance is already staged
	if IsStaged(stage, resourceshape) {
		return
	}

	resourceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourceshape.Resource != nil {
		StageBranch(stage, resourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchResourceTaskShape(resourcetaskshape *ResourceTaskShape) {

	// check if instance is already staged
	if IsStaged(stage, resourcetaskshape) {
		return
	}

	resourcetaskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcetaskshape.Resource != nil {
		StageBranch(stage, resourcetaskshape.Resource)
	}
	if resourcetaskshape.Task != nil {
		StageBranch(stage, resourcetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTask(task *Task) {

	// check if instance is already staged
	if IsStaged(stage, task) {
		return
	}

	task.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if task.ReferencedTask != nil {
		StageBranch(stage, task.ReferencedTask)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range task.Predecessors {
		StageBranch(stage, _task)
	}
	for _, _product := range task.Inputs {
		StageBranch(stage, _product)
	}
	for _, _product := range task.Outputs {
		StageBranch(stage, _product)
	}
	for _, _taskgroup := range task.TaskGroupsToDisplay {
		StageBranch(stage, _taskgroup)
	}
	for _, _task := range task.SubTasks {
		StageBranch(stage, _task)
	}

}

func (stage *Stage) StageBranchTaskCompositionShape(taskcompositionshape *TaskCompositionShape) {

	// check if instance is already staged
	if IsStaged(stage, taskcompositionshape) {
		return
	}

	taskcompositionshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshape.Task != nil {
		StageBranch(stage, taskcompositionshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskGroup(taskgroup *TaskGroup) {

	// check if instance is already staged
	if IsStaged(stage, taskgroup) {
		return
	}

	taskgroup.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskgroup.Tasks {
		StageBranch(stage, _task)
	}

}

func (stage *Stage) StageBranchTaskGroupShape(taskgroupshape *TaskGroupShape) {

	// check if instance is already staged
	if IsStaged(stage, taskgroupshape) {
		return
	}

	taskgroupshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskgroupshape.TaskGroup != nil {
		StageBranch(stage, taskgroupshape.TaskGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskInputShape(taskinputshape *TaskInputShape) {

	// check if instance is already staged
	if IsStaged(stage, taskinputshape) {
		return
	}

	taskinputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshape.Product != nil {
		StageBranch(stage, taskinputshape.Product)
	}
	if taskinputshape.Task != nil {
		StageBranch(stage, taskinputshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskOutputShape(taskoutputshape *TaskOutputShape) {

	// check if instance is already staged
	if IsStaged(stage, taskoutputshape) {
		return
	}

	taskoutputshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshape.Task != nil {
		StageBranch(stage, taskoutputshape.Task)
	}
	if taskoutputshape.Product != nil {
		StageBranch(stage, taskoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if IsStaged(stage, taskshape) {
		return
	}

	taskshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		StageBranch(stage, taskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteProductShape:
		toT := CopyBranchNoteProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteResourceShape:
		toT := CopyBranchNoteResourceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteTaskShape:
		toT := CopyBranchNoteTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Product:
		toT := CopyBranchProduct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductCompositionShape:
		toT := CopyBranchProductCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ProductShape:
		toT := CopyBranchProductShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Resource:
		toT := CopyBranchResource(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ResourceCompositionShape:
		toT := CopyBranchResourceCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ResourceShape:
		toT := CopyBranchResourceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ResourceTaskShape:
		toT := CopyBranchResourceTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Task:
		toT := CopyBranchTask(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskCompositionShape:
		toT := CopyBranchTaskCompositionShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskGroup:
		toT := CopyBranchTaskGroup(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskGroupShape:
		toT := CopyBranchTaskGroupShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskInputShape:
		toT := CopyBranchTaskInputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskOutputShape:
		toT := CopyBranchTaskOutputShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TaskShape:
		toT := CopyBranchTaskShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.CopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagramFrom.Product_Shapes {
		diagramTo.Product_Shapes = append(diagramTo.Product_Shapes, CopyBranchProductShape(mapOrigCopy, _productshape))
	}
	for _, _product := range diagramFrom.ProductsWhoseNodeIsExpanded {
		diagramTo.ProductsWhoseNodeIsExpanded = append(diagramTo.ProductsWhoseNodeIsExpanded, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _productcompositionshape := range diagramFrom.ProductComposition_Shapes {
		diagramTo.ProductComposition_Shapes = append(diagramTo.ProductComposition_Shapes, CopyBranchProductCompositionShape(mapOrigCopy, _productcompositionshape))
	}
	for _, _taskshape := range diagramFrom.Task_Shapes {
		diagramTo.Task_Shapes = append(diagramTo.Task_Shapes, CopyBranchTaskShape(mapOrigCopy, _taskshape))
	}
	for _, _task := range diagramFrom.TasksWhoseNodeIsExpanded {
		diagramTo.TasksWhoseNodeIsExpanded = append(diagramTo.TasksWhoseNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramFrom.TasksWhoseInputNodeIsExpanded {
		diagramTo.TasksWhoseInputNodeIsExpanded = append(diagramTo.TasksWhoseInputNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramFrom.TasksWhoseOutputNodeIsExpanded {
		diagramTo.TasksWhoseOutputNodeIsExpanded = append(diagramTo.TasksWhoseOutputNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskgroupshape := range diagramFrom.TaskGroupShapes {
		diagramTo.TaskGroupShapes = append(diagramTo.TaskGroupShapes, CopyBranchTaskGroupShape(mapOrigCopy, _taskgroupshape))
	}
	for _, _taskgroup := range diagramFrom.TaskGroupsWhoseNodeIsExpanded {
		diagramTo.TaskGroupsWhoseNodeIsExpanded = append(diagramTo.TaskGroupsWhoseNodeIsExpanded, CopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}
	for _, _taskcompositionshape := range diagramFrom.TaskComposition_Shapes {
		diagramTo.TaskComposition_Shapes = append(diagramTo.TaskComposition_Shapes, CopyBranchTaskCompositionShape(mapOrigCopy, _taskcompositionshape))
	}
	for _, _taskinputshape := range diagramFrom.TaskInputShapes {
		diagramTo.TaskInputShapes = append(diagramTo.TaskInputShapes, CopyBranchTaskInputShape(mapOrigCopy, _taskinputshape))
	}
	for _, _taskoutputshape := range diagramFrom.TaskOutputShapes {
		diagramTo.TaskOutputShapes = append(diagramTo.TaskOutputShapes, CopyBranchTaskOutputShape(mapOrigCopy, _taskoutputshape))
	}
	for _, _noteshape := range diagramFrom.Note_Shapes {
		diagramTo.Note_Shapes = append(diagramTo.Note_Shapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramFrom.NotesWhoseNodeIsExpanded {
		diagramTo.NotesWhoseNodeIsExpanded = append(diagramTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _noteproductshape := range diagramFrom.NoteProductShapes {
		diagramTo.NoteProductShapes = append(diagramTo.NoteProductShapes, CopyBranchNoteProductShape(mapOrigCopy, _noteproductshape))
	}
	for _, _notetaskshape := range diagramFrom.NoteTaskShapes {
		diagramTo.NoteTaskShapes = append(diagramTo.NoteTaskShapes, CopyBranchNoteTaskShape(mapOrigCopy, _notetaskshape))
	}
	for _, _noteresourceshape := range diagramFrom.NoteResourceShapes {
		diagramTo.NoteResourceShapes = append(diagramTo.NoteResourceShapes, CopyBranchNoteResourceShape(mapOrigCopy, _noteresourceshape))
	}
	for _, _resourceshape := range diagramFrom.Resource_Shapes {
		diagramTo.Resource_Shapes = append(diagramTo.Resource_Shapes, CopyBranchResourceShape(mapOrigCopy, _resourceshape))
	}
	for _, _resource := range diagramFrom.ResourcesWhoseNodeIsExpanded {
		diagramTo.ResourcesWhoseNodeIsExpanded = append(diagramTo.ResourcesWhoseNodeIsExpanded, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _resourcecompositionshape := range diagramFrom.ResourceComposition_Shapes {
		diagramTo.ResourceComposition_Shapes = append(diagramTo.ResourceComposition_Shapes, CopyBranchResourceCompositionShape(mapOrigCopy, _resourcecompositionshape))
	}
	for _, _resourcetaskshape := range diagramFrom.ResourceTaskShapes {
		diagramTo.ResourceTaskShapes = append(diagramTo.ResourceTaskShapes, CopyBranchResourceTaskShape(mapOrigCopy, _resourcetaskshape))
	}

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _product := range libraryFrom.RootProducts {
		libraryTo.RootProducts = append(libraryTo.RootProducts, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _task := range libraryFrom.RootTasks {
		libraryTo.RootTasks = append(libraryTo.RootTasks, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskgroup := range libraryFrom.RootTaskGroups {
		libraryTo.RootTaskGroups = append(libraryTo.RootTaskGroups, CopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}
	for _, _resource := range libraryFrom.RootResources {
		libraryTo.RootResources = append(libraryTo.RootResources, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _note := range libraryFrom.Notes {
		libraryTo.Notes = append(libraryTo.Notes, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _diagram := range libraryFrom.Diagrams {
		libraryTo.Diagrams = append(libraryTo.Diagrams, CopyBranchDiagram(mapOrigCopy, _diagram))
	}

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range noteFrom.Products {
		noteTo.Products = append(noteTo.Products, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _task := range noteFrom.Tasks {
		noteTo.Tasks = append(noteTo.Tasks, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _resource := range noteFrom.Resources {
		noteTo.Resources = append(noteTo.Resources, CopyBranchResource(mapOrigCopy, _resource))
	}

	return
}

func CopyBranchNoteProductShape(mapOrigCopy map[any]any, noteproductshapeFrom *NoteProductShape) (noteproductshapeTo *NoteProductShape) {

	// noteproductshapeFrom has already been copied
	if _noteproductshapeTo, ok := mapOrigCopy[noteproductshapeFrom]; ok {
		noteproductshapeTo = _noteproductshapeTo.(*NoteProductShape)
		return
	}

	noteproductshapeTo = new(NoteProductShape)
	mapOrigCopy[noteproductshapeFrom] = noteproductshapeTo
	noteproductshapeFrom.CopyBasicFields(noteproductshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshapeFrom.Note != nil {
		noteproductshapeTo.Note = CopyBranchNote(mapOrigCopy, noteproductshapeFrom.Note)
	}
	if noteproductshapeFrom.Product != nil {
		noteproductshapeTo.Product = CopyBranchProduct(mapOrigCopy, noteproductshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteResourceShape(mapOrigCopy map[any]any, noteresourceshapeFrom *NoteResourceShape) (noteresourceshapeTo *NoteResourceShape) {

	// noteresourceshapeFrom has already been copied
	if _noteresourceshapeTo, ok := mapOrigCopy[noteresourceshapeFrom]; ok {
		noteresourceshapeTo = _noteresourceshapeTo.(*NoteResourceShape)
		return
	}

	noteresourceshapeTo = new(NoteResourceShape)
	mapOrigCopy[noteresourceshapeFrom] = noteresourceshapeTo
	noteresourceshapeFrom.CopyBasicFields(noteresourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteresourceshapeFrom.Note != nil {
		noteresourceshapeTo.Note = CopyBranchNote(mapOrigCopy, noteresourceshapeFrom.Note)
	}
	if noteresourceshapeFrom.Resource != nil {
		noteresourceshapeTo.Resource = CopyBranchResource(mapOrigCopy, noteresourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.CopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = CopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteTaskShape(mapOrigCopy map[any]any, notetaskshapeFrom *NoteTaskShape) (notetaskshapeTo *NoteTaskShape) {

	// notetaskshapeFrom has already been copied
	if _notetaskshapeTo, ok := mapOrigCopy[notetaskshapeFrom]; ok {
		notetaskshapeTo = _notetaskshapeTo.(*NoteTaskShape)
		return
	}

	notetaskshapeTo = new(NoteTaskShape)
	mapOrigCopy[notetaskshapeFrom] = notetaskshapeTo
	notetaskshapeFrom.CopyBasicFields(notetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshapeFrom.Note != nil {
		notetaskshapeTo.Note = CopyBranchNote(mapOrigCopy, notetaskshapeFrom.Note)
	}
	if notetaskshapeFrom.Task != nil {
		notetaskshapeTo.Task = CopyBranchTask(mapOrigCopy, notetaskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProduct(mapOrigCopy map[any]any, productFrom *Product) (productTo *Product) {

	// productFrom has already been copied
	if _productTo, ok := mapOrigCopy[productFrom]; ok {
		productTo = _productTo.(*Product)
		return
	}

	productTo = new(Product)
	mapOrigCopy[productFrom] = productTo
	productFrom.CopyBasicFields(productTo)

	//insertion point for the staging of instances referenced by pointers
	if productFrom.ReferencedProduct != nil {
		productTo.ReferencedProduct = CopyBranchProduct(mapOrigCopy, productFrom.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range productFrom.SubProducts {
		productTo.SubProducts = append(productTo.SubProducts, CopyBranchProduct(mapOrigCopy, _product))
	}

	return
}

func CopyBranchProductCompositionShape(mapOrigCopy map[any]any, productcompositionshapeFrom *ProductCompositionShape) (productcompositionshapeTo *ProductCompositionShape) {

	// productcompositionshapeFrom has already been copied
	if _productcompositionshapeTo, ok := mapOrigCopy[productcompositionshapeFrom]; ok {
		productcompositionshapeTo = _productcompositionshapeTo.(*ProductCompositionShape)
		return
	}

	productcompositionshapeTo = new(ProductCompositionShape)
	mapOrigCopy[productcompositionshapeFrom] = productcompositionshapeTo
	productcompositionshapeFrom.CopyBasicFields(productcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshapeFrom.Product != nil {
		productcompositionshapeTo.Product = CopyBranchProduct(mapOrigCopy, productcompositionshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchProductShape(mapOrigCopy map[any]any, productshapeFrom *ProductShape) (productshapeTo *ProductShape) {

	// productshapeFrom has already been copied
	if _productshapeTo, ok := mapOrigCopy[productshapeFrom]; ok {
		productshapeTo = _productshapeTo.(*ProductShape)
		return
	}

	productshapeTo = new(ProductShape)
	mapOrigCopy[productshapeFrom] = productshapeTo
	productshapeFrom.CopyBasicFields(productshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if productshapeFrom.Product != nil {
		productshapeTo.Product = CopyBranchProduct(mapOrigCopy, productshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {

	// resourceFrom has already been copied
	if _resourceTo, ok := mapOrigCopy[resourceFrom]; ok {
		resourceTo = _resourceTo.(*Resource)
		return
	}

	resourceTo = new(Resource)
	mapOrigCopy[resourceFrom] = resourceTo
	resourceFrom.CopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers
	if resourceFrom.ReferencedResource != nil {
		resourceTo.ReferencedResource = CopyBranchResource(mapOrigCopy, resourceFrom.ReferencedResource)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range resourceFrom.Tasks {
		resourceTo.Tasks = append(resourceTo.Tasks, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _resource := range resourceFrom.SubResources {
		resourceTo.SubResources = append(resourceTo.SubResources, CopyBranchResource(mapOrigCopy, _resource))
	}

	return
}

func CopyBranchResourceCompositionShape(mapOrigCopy map[any]any, resourcecompositionshapeFrom *ResourceCompositionShape) (resourcecompositionshapeTo *ResourceCompositionShape) {

	// resourcecompositionshapeFrom has already been copied
	if _resourcecompositionshapeTo, ok := mapOrigCopy[resourcecompositionshapeFrom]; ok {
		resourcecompositionshapeTo = _resourcecompositionshapeTo.(*ResourceCompositionShape)
		return
	}

	resourcecompositionshapeTo = new(ResourceCompositionShape)
	mapOrigCopy[resourcecompositionshapeFrom] = resourcecompositionshapeTo
	resourcecompositionshapeFrom.CopyBasicFields(resourcecompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourcecompositionshapeFrom.Resource != nil {
		resourcecompositionshapeTo.Resource = CopyBranchResource(mapOrigCopy, resourcecompositionshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchResourceShape(mapOrigCopy map[any]any, resourceshapeFrom *ResourceShape) (resourceshapeTo *ResourceShape) {

	// resourceshapeFrom has already been copied
	if _resourceshapeTo, ok := mapOrigCopy[resourceshapeFrom]; ok {
		resourceshapeTo = _resourceshapeTo.(*ResourceShape)
		return
	}

	resourceshapeTo = new(ResourceShape)
	mapOrigCopy[resourceshapeFrom] = resourceshapeTo
	resourceshapeFrom.CopyBasicFields(resourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourceshapeFrom.Resource != nil {
		resourceshapeTo.Resource = CopyBranchResource(mapOrigCopy, resourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchResourceTaskShape(mapOrigCopy map[any]any, resourcetaskshapeFrom *ResourceTaskShape) (resourcetaskshapeTo *ResourceTaskShape) {

	// resourcetaskshapeFrom has already been copied
	if _resourcetaskshapeTo, ok := mapOrigCopy[resourcetaskshapeFrom]; ok {
		resourcetaskshapeTo = _resourcetaskshapeTo.(*ResourceTaskShape)
		return
	}

	resourcetaskshapeTo = new(ResourceTaskShape)
	mapOrigCopy[resourcetaskshapeFrom] = resourcetaskshapeTo
	resourcetaskshapeFrom.CopyBasicFields(resourcetaskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if resourcetaskshapeFrom.Resource != nil {
		resourcetaskshapeTo.Resource = CopyBranchResource(mapOrigCopy, resourcetaskshapeFrom.Resource)
	}
	if resourcetaskshapeFrom.Task != nil {
		resourcetaskshapeTo.Task = CopyBranchTask(mapOrigCopy, resourcetaskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTask(mapOrigCopy map[any]any, taskFrom *Task) (taskTo *Task) {

	// taskFrom has already been copied
	if _taskTo, ok := mapOrigCopy[taskFrom]; ok {
		taskTo = _taskTo.(*Task)
		return
	}

	taskTo = new(Task)
	mapOrigCopy[taskFrom] = taskTo
	taskFrom.CopyBasicFields(taskTo)

	//insertion point for the staging of instances referenced by pointers
	if taskFrom.ReferencedTask != nil {
		taskTo.ReferencedTask = CopyBranchTask(mapOrigCopy, taskFrom.ReferencedTask)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskFrom.Predecessors {
		taskTo.Predecessors = append(taskTo.Predecessors, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _product := range taskFrom.Inputs {
		taskTo.Inputs = append(taskTo.Inputs, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _product := range taskFrom.Outputs {
		taskTo.Outputs = append(taskTo.Outputs, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _taskgroup := range taskFrom.TaskGroupsToDisplay {
		taskTo.TaskGroupsToDisplay = append(taskTo.TaskGroupsToDisplay, CopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}
	for _, _task := range taskFrom.SubTasks {
		taskTo.SubTasks = append(taskTo.SubTasks, CopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func CopyBranchTaskCompositionShape(mapOrigCopy map[any]any, taskcompositionshapeFrom *TaskCompositionShape) (taskcompositionshapeTo *TaskCompositionShape) {

	// taskcompositionshapeFrom has already been copied
	if _taskcompositionshapeTo, ok := mapOrigCopy[taskcompositionshapeFrom]; ok {
		taskcompositionshapeTo = _taskcompositionshapeTo.(*TaskCompositionShape)
		return
	}

	taskcompositionshapeTo = new(TaskCompositionShape)
	mapOrigCopy[taskcompositionshapeFrom] = taskcompositionshapeTo
	taskcompositionshapeFrom.CopyBasicFields(taskcompositionshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshapeFrom.Task != nil {
		taskcompositionshapeTo.Task = CopyBranchTask(mapOrigCopy, taskcompositionshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskGroup(mapOrigCopy map[any]any, taskgroupFrom *TaskGroup) (taskgroupTo *TaskGroup) {

	// taskgroupFrom has already been copied
	if _taskgroupTo, ok := mapOrigCopy[taskgroupFrom]; ok {
		taskgroupTo = _taskgroupTo.(*TaskGroup)
		return
	}

	taskgroupTo = new(TaskGroup)
	mapOrigCopy[taskgroupFrom] = taskgroupTo
	taskgroupFrom.CopyBasicFields(taskgroupTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskgroupFrom.Tasks {
		taskgroupTo.Tasks = append(taskgroupTo.Tasks, CopyBranchTask(mapOrigCopy, _task))
	}

	return
}

func CopyBranchTaskGroupShape(mapOrigCopy map[any]any, taskgroupshapeFrom *TaskGroupShape) (taskgroupshapeTo *TaskGroupShape) {

	// taskgroupshapeFrom has already been copied
	if _taskgroupshapeTo, ok := mapOrigCopy[taskgroupshapeFrom]; ok {
		taskgroupshapeTo = _taskgroupshapeTo.(*TaskGroupShape)
		return
	}

	taskgroupshapeTo = new(TaskGroupShape)
	mapOrigCopy[taskgroupshapeFrom] = taskgroupshapeTo
	taskgroupshapeFrom.CopyBasicFields(taskgroupshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskgroupshapeFrom.TaskGroup != nil {
		taskgroupshapeTo.TaskGroup = CopyBranchTaskGroup(mapOrigCopy, taskgroupshapeFrom.TaskGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskInputShape(mapOrigCopy map[any]any, taskinputshapeFrom *TaskInputShape) (taskinputshapeTo *TaskInputShape) {

	// taskinputshapeFrom has already been copied
	if _taskinputshapeTo, ok := mapOrigCopy[taskinputshapeFrom]; ok {
		taskinputshapeTo = _taskinputshapeTo.(*TaskInputShape)
		return
	}

	taskinputshapeTo = new(TaskInputShape)
	mapOrigCopy[taskinputshapeFrom] = taskinputshapeTo
	taskinputshapeFrom.CopyBasicFields(taskinputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshapeFrom.Product != nil {
		taskinputshapeTo.Product = CopyBranchProduct(mapOrigCopy, taskinputshapeFrom.Product)
	}
	if taskinputshapeFrom.Task != nil {
		taskinputshapeTo.Task = CopyBranchTask(mapOrigCopy, taskinputshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskOutputShape(mapOrigCopy map[any]any, taskoutputshapeFrom *TaskOutputShape) (taskoutputshapeTo *TaskOutputShape) {

	// taskoutputshapeFrom has already been copied
	if _taskoutputshapeTo, ok := mapOrigCopy[taskoutputshapeFrom]; ok {
		taskoutputshapeTo = _taskoutputshapeTo.(*TaskOutputShape)
		return
	}

	taskoutputshapeTo = new(TaskOutputShape)
	mapOrigCopy[taskoutputshapeFrom] = taskoutputshapeTo
	taskoutputshapeFrom.CopyBasicFields(taskoutputshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshapeFrom.Task != nil {
		taskoutputshapeTo.Task = CopyBranchTask(mapOrigCopy, taskoutputshapeFrom.Task)
	}
	if taskoutputshapeFrom.Product != nil {
		taskoutputshapeTo.Product = CopyBranchProduct(mapOrigCopy, taskoutputshapeFrom.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchTaskShape(mapOrigCopy map[any]any, taskshapeFrom *TaskShape) (taskshapeTo *TaskShape) {

	// taskshapeFrom has already been copied
	if _taskshapeTo, ok := mapOrigCopy[taskshapeFrom]; ok {
		taskshapeTo = _taskshapeTo.(*TaskShape)
		return
	}

	taskshapeTo = new(TaskShape)
	mapOrigCopy[taskshapeFrom] = taskshapeTo
	taskshapeFrom.CopyBasicFields(taskshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if taskshapeFrom.Task != nil {
		taskshapeTo.Task = CopyBranchTask(mapOrigCopy, taskshapeFrom.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *NoteProductShape:
		stage.UnstageBranchNoteProductShape(target)

	case *NoteResourceShape:
		stage.UnstageBranchNoteResourceShape(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *NoteTaskShape:
		stage.UnstageBranchNoteTaskShape(target)

	case *Product:
		stage.UnstageBranchProduct(target)

	case *ProductCompositionShape:
		stage.UnstageBranchProductCompositionShape(target)

	case *ProductShape:
		stage.UnstageBranchProductShape(target)

	case *Resource:
		stage.UnstageBranchResource(target)

	case *ResourceCompositionShape:
		stage.UnstageBranchResourceCompositionShape(target)

	case *ResourceShape:
		stage.UnstageBranchResourceShape(target)

	case *ResourceTaskShape:
		stage.UnstageBranchResourceTaskShape(target)

	case *Task:
		stage.UnstageBranchTask(target)

	case *TaskCompositionShape:
		stage.UnstageBranchTaskCompositionShape(target)

	case *TaskGroup:
		stage.UnstageBranchTaskGroup(target)

	case *TaskGroupShape:
		stage.UnstageBranchTaskGroupShape(target)

	case *TaskInputShape:
		stage.UnstageBranchTaskInputShape(target)

	case *TaskOutputShape:
		stage.UnstageBranchTaskOutputShape(target)

	case *TaskShape:
		stage.UnstageBranchTaskShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !IsStaged(stage, diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagram.Product_Shapes {
		UnstageBranch(stage, _productshape)
	}
	for _, _product := range diagram.ProductsWhoseNodeIsExpanded {
		UnstageBranch(stage, _product)
	}
	for _, _productcompositionshape := range diagram.ProductComposition_Shapes {
		UnstageBranch(stage, _productcompositionshape)
	}
	for _, _taskshape := range diagram.Task_Shapes {
		UnstageBranch(stage, _taskshape)
	}
	for _, _task := range diagram.TasksWhoseNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseInputNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range diagram.TasksWhoseOutputNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _taskgroupshape := range diagram.TaskGroupShapes {
		UnstageBranch(stage, _taskgroupshape)
	}
	for _, _taskgroup := range diagram.TaskGroupsWhoseNodeIsExpanded {
		UnstageBranch(stage, _taskgroup)
	}
	for _, _taskcompositionshape := range diagram.TaskComposition_Shapes {
		UnstageBranch(stage, _taskcompositionshape)
	}
	for _, _taskinputshape := range diagram.TaskInputShapes {
		UnstageBranch(stage, _taskinputshape)
	}
	for _, _taskoutputshape := range diagram.TaskOutputShapes {
		UnstageBranch(stage, _taskoutputshape)
	}
	for _, _noteshape := range diagram.Note_Shapes {
		UnstageBranch(stage, _noteshape)
	}
	for _, _note := range diagram.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}
	for _, _noteproductshape := range diagram.NoteProductShapes {
		UnstageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagram.NoteTaskShapes {
		UnstageBranch(stage, _notetaskshape)
	}
	for _, _noteresourceshape := range diagram.NoteResourceShapes {
		UnstageBranch(stage, _noteresourceshape)
	}
	for _, _resourceshape := range diagram.Resource_Shapes {
		UnstageBranch(stage, _resourceshape)
	}
	for _, _resource := range diagram.ResourcesWhoseNodeIsExpanded {
		UnstageBranch(stage, _resource)
	}
	for _, _resourcecompositionshape := range diagram.ResourceComposition_Shapes {
		UnstageBranch(stage, _resourcecompositionshape)
	}
	for _, _resourcetaskshape := range diagram.ResourceTaskShapes {
		UnstageBranch(stage, _resourcetaskshape)
	}

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}
	for _, _product := range library.RootProducts {
		UnstageBranch(stage, _product)
	}
	for _, _task := range library.RootTasks {
		UnstageBranch(stage, _task)
	}
	for _, _taskgroup := range library.RootTaskGroups {
		UnstageBranch(stage, _taskgroup)
	}
	for _, _resource := range library.RootResources {
		UnstageBranch(stage, _resource)
	}
	for _, _note := range library.Notes {
		UnstageBranch(stage, _note)
	}
	for _, _diagram := range library.Diagrams {
		UnstageBranch(stage, _diagram)
	}

}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range note.Products {
		UnstageBranch(stage, _product)
	}
	for _, _task := range note.Tasks {
		UnstageBranch(stage, _task)
	}
	for _, _resource := range note.Resources {
		UnstageBranch(stage, _resource)
	}

}

func (stage *Stage) UnstageBranchNoteProductShape(noteproductshape *NoteProductShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteproductshape) {
		return
	}

	noteproductshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteproductshape.Note != nil {
		UnstageBranch(stage, noteproductshape.Note)
	}
	if noteproductshape.Product != nil {
		UnstageBranch(stage, noteproductshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteResourceShape(noteresourceshape *NoteResourceShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteresourceshape) {
		return
	}

	noteresourceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteresourceshape.Note != nil {
		UnstageBranch(stage, noteresourceshape.Note)
	}
	if noteresourceshape.Resource != nil {
		UnstageBranch(stage, noteresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		UnstageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteTaskShape(notetaskshape *NoteTaskShape) {

	// check if instance is already staged
	if !IsStaged(stage, notetaskshape) {
		return
	}

	notetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notetaskshape.Note != nil {
		UnstageBranch(stage, notetaskshape.Note)
	}
	if notetaskshape.Task != nil {
		UnstageBranch(stage, notetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchProduct(product *Product) {

	// check if instance is already staged
	if !IsStaged(stage, product) {
		return
	}

	product.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if product.ReferencedProduct != nil {
		UnstageBranch(stage, product.ReferencedProduct)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _product := range product.SubProducts {
		UnstageBranch(stage, _product)
	}

}

func (stage *Stage) UnstageBranchProductCompositionShape(productcompositionshape *ProductCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, productcompositionshape) {
		return
	}

	productcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productcompositionshape.Product != nil {
		UnstageBranch(stage, productcompositionshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchProductShape(productshape *ProductShape) {

	// check if instance is already staged
	if !IsStaged(stage, productshape) {
		return
	}

	productshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if productshape.Product != nil {
		UnstageBranch(stage, productshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchResource(resource *Resource) {

	// check if instance is already staged
	if !IsStaged(stage, resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resource.ReferencedResource != nil {
		UnstageBranch(stage, resource.ReferencedResource)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range resource.Tasks {
		UnstageBranch(stage, _task)
	}
	for _, _resource := range resource.SubResources {
		UnstageBranch(stage, _resource)
	}

}

func (stage *Stage) UnstageBranchResourceCompositionShape(resourcecompositionshape *ResourceCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, resourcecompositionshape) {
		return
	}

	resourcecompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcecompositionshape.Resource != nil {
		UnstageBranch(stage, resourcecompositionshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchResourceShape(resourceshape *ResourceShape) {

	// check if instance is already staged
	if !IsStaged(stage, resourceshape) {
		return
	}

	resourceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourceshape.Resource != nil {
		UnstageBranch(stage, resourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchResourceTaskShape(resourcetaskshape *ResourceTaskShape) {

	// check if instance is already staged
	if !IsStaged(stage, resourcetaskshape) {
		return
	}

	resourcetaskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if resourcetaskshape.Resource != nil {
		UnstageBranch(stage, resourcetaskshape.Resource)
	}
	if resourcetaskshape.Task != nil {
		UnstageBranch(stage, resourcetaskshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTask(task *Task) {

	// check if instance is already staged
	if !IsStaged(stage, task) {
		return
	}

	task.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if task.ReferencedTask != nil {
		UnstageBranch(stage, task.ReferencedTask)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range task.Predecessors {
		UnstageBranch(stage, _task)
	}
	for _, _product := range task.Inputs {
		UnstageBranch(stage, _product)
	}
	for _, _product := range task.Outputs {
		UnstageBranch(stage, _product)
	}
	for _, _taskgroup := range task.TaskGroupsToDisplay {
		UnstageBranch(stage, _taskgroup)
	}
	for _, _task := range task.SubTasks {
		UnstageBranch(stage, _task)
	}

}

func (stage *Stage) UnstageBranchTaskCompositionShape(taskcompositionshape *TaskCompositionShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskcompositionshape) {
		return
	}

	taskcompositionshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskcompositionshape.Task != nil {
		UnstageBranch(stage, taskcompositionshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskGroup(taskgroup *TaskGroup) {

	// check if instance is already staged
	if !IsStaged(stage, taskgroup) {
		return
	}

	taskgroup.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _task := range taskgroup.Tasks {
		UnstageBranch(stage, _task)
	}

}

func (stage *Stage) UnstageBranchTaskGroupShape(taskgroupshape *TaskGroupShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskgroupshape) {
		return
	}

	taskgroupshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskgroupshape.TaskGroup != nil {
		UnstageBranch(stage, taskgroupshape.TaskGroup)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskInputShape(taskinputshape *TaskInputShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskinputshape) {
		return
	}

	taskinputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskinputshape.Product != nil {
		UnstageBranch(stage, taskinputshape.Product)
	}
	if taskinputshape.Task != nil {
		UnstageBranch(stage, taskinputshape.Task)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskOutputShape(taskoutputshape *TaskOutputShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskoutputshape) {
		return
	}

	taskoutputshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskoutputshape.Task != nil {
		UnstageBranch(stage, taskoutputshape.Task)
	}
	if taskoutputshape.Product != nil {
		UnstageBranch(stage, taskoutputshape.Product)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchTaskShape(taskshape *TaskShape) {

	// check if instance is already staged
	if !IsStaged(stage, taskshape) {
		return
	}

	taskshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if taskshape.Task != nil {
		UnstageBranch(stage, taskshape.Task)
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
	reference.TaskGroupsToDisplay = reference.TaskGroupsToDisplay[:0]
	for _, _b := range instance.TaskGroupsToDisplay {
		reference.TaskGroupsToDisplay = append(reference.TaskGroupsToDisplay, stage.TaskGroups_reference[_b])
	}
	reference.SubTasks = reference.SubTasks[:0]
	for _, _b := range instance.SubTasks {
		reference.SubTasks = append(reference.SubTasks, stage.Tasks_reference[_b])
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
	var _TaskGroupsToDisplay []*TaskGroup
	for _, _reference := range reference.TaskGroupsToDisplay {
		if _instance, ok := stage.TaskGroups_instance[_reference]; ok {
			_TaskGroupsToDisplay = append(_TaskGroupsToDisplay, _instance)
		}
	}
	reference.TaskGroupsToDisplay = _TaskGroupsToDisplay
	var _SubTasks []*Task
	for _, _reference := range reference.SubTasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_SubTasks = append(_SubTasks, _instance)
		}
	}
	reference.SubTasks = _SubTasks
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
	if diagram.DrawVerticalTimeLines != diagramOther.DrawVerticalTimeLines {
		diffs = append(diffs, diagram.GongMarshallField(stage, "DrawVerticalTimeLines"))
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
		ops := Diff(stage, diagram, diagramOther, "Product_Shapes", diagramOther.Product_Shapes, diagram.Product_Shapes)
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
		ops := Diff(stage, diagram, diagramOther, "ProductsWhoseNodeIsExpanded", diagramOther.ProductsWhoseNodeIsExpanded, diagram.ProductsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagram, diagramOther, "ProductComposition_Shapes", diagramOther.ProductComposition_Shapes, diagram.ProductComposition_Shapes)
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
		ops := Diff(stage, diagram, diagramOther, "Task_Shapes", diagramOther.Task_Shapes, diagram.Task_Shapes)
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
		ops := Diff(stage, diagram, diagramOther, "TasksWhoseNodeIsExpanded", diagramOther.TasksWhoseNodeIsExpanded, diagram.TasksWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagram, diagramOther, "TasksWhoseInputNodeIsExpanded", diagramOther.TasksWhoseInputNodeIsExpanded, diagram.TasksWhoseInputNodeIsExpanded)
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
		ops := Diff(stage, diagram, diagramOther, "TasksWhoseOutputNodeIsExpanded", diagramOther.TasksWhoseOutputNodeIsExpanded, diagram.TasksWhoseOutputNodeIsExpanded)
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
		ops := Diff(stage, diagram, diagramOther, "TaskGroupShapes", diagramOther.TaskGroupShapes, diagram.TaskGroupShapes)
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
		ops := Diff(stage, diagram, diagramOther, "TaskGroupsWhoseNodeIsExpanded", diagramOther.TaskGroupsWhoseNodeIsExpanded, diagram.TaskGroupsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagram, diagramOther, "TaskComposition_Shapes", diagramOther.TaskComposition_Shapes, diagram.TaskComposition_Shapes)
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
		ops := Diff(stage, diagram, diagramOther, "TaskInputShapes", diagramOther.TaskInputShapes, diagram.TaskInputShapes)
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
		ops := Diff(stage, diagram, diagramOther, "TaskOutputShapes", diagramOther.TaskOutputShapes, diagram.TaskOutputShapes)
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
		ops := Diff(stage, diagram, diagramOther, "Note_Shapes", diagramOther.Note_Shapes, diagram.Note_Shapes)
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
		ops := Diff(stage, diagram, diagramOther, "NotesWhoseNodeIsExpanded", diagramOther.NotesWhoseNodeIsExpanded, diagram.NotesWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagram, diagramOther, "NoteProductShapes", diagramOther.NoteProductShapes, diagram.NoteProductShapes)
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
		ops := Diff(stage, diagram, diagramOther, "NoteTaskShapes", diagramOther.NoteTaskShapes, diagram.NoteTaskShapes)
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
		ops := Diff(stage, diagram, diagramOther, "NoteResourceShapes", diagramOther.NoteResourceShapes, diagram.NoteResourceShapes)
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
		ops := Diff(stage, diagram, diagramOther, "Resource_Shapes", diagramOther.Resource_Shapes, diagram.Resource_Shapes)
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
		ops := Diff(stage, diagram, diagramOther, "ResourcesWhoseNodeIsExpanded", diagramOther.ResourcesWhoseNodeIsExpanded, diagram.ResourcesWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagram, diagramOther, "ResourceComposition_Shapes", diagramOther.ResourceComposition_Shapes, diagram.ResourceComposition_Shapes)
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
		ops := Diff(stage, diagram, diagramOther, "ResourceTaskShapes", diagramOther.ResourceTaskShapes, diagram.ResourceTaskShapes)
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
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
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
		ops := Diff(stage, library, libraryOther, "RootProducts", libraryOther.RootProducts, library.RootProducts)
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
		ops := Diff(stage, library, libraryOther, "RootTasks", libraryOther.RootTasks, library.RootTasks)
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
		ops := Diff(stage, library, libraryOther, "RootTaskGroups", libraryOther.RootTaskGroups, library.RootTaskGroups)
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
		ops := Diff(stage, library, libraryOther, "RootResources", libraryOther.RootResources, library.RootResources)
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
		ops := Diff(stage, library, libraryOther, "Notes", libraryOther.Notes, library.Notes)
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
		ops := Diff(stage, library, libraryOther, "Diagrams", libraryOther.Diagrams, library.Diagrams)
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
		ops := Diff(stage, note, noteOther, "Products", noteOther.Products, note.Products)
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
		ops := Diff(stage, note, noteOther, "Tasks", noteOther.Tasks, note.Tasks)
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
		ops := Diff(stage, note, noteOther, "Resources", noteOther.Resources, note.Resources)
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
		ops := Diff(stage, product, productOther, "SubProducts", productOther.SubProducts, product.SubProducts)
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
		ops := Diff(stage, resource, resourceOther, "Tasks", resourceOther.Tasks, resource.Tasks)
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
		ops := Diff(stage, resource, resourceOther, "SubResources", resourceOther.SubResources, resource.SubResources)
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
		ops := Diff(stage, task, taskOther, "Predecessors", taskOther.Predecessors, task.Predecessors)
		diffs = append(diffs, ops)
	}
	if task.IsStartDateComputedFromPredecessors != taskOther.IsStartDateComputedFromPredecessors {
		diffs = append(diffs, task.GongMarshallField(stage, "IsStartDateComputedFromPredecessors"))
	}
	if task.IsMilestone != taskOther.IsMilestone {
		diffs = append(diffs, task.GongMarshallField(stage, "IsMilestone"))
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
		ops := Diff(stage, task, taskOther, "Inputs", taskOther.Inputs, task.Inputs)
		diffs = append(diffs, ops)
	}
	if task.IsInputsNodeExpanded != taskOther.IsInputsNodeExpanded {
		diffs = append(diffs, task.GongMarshallField(stage, "IsInputsNodeExpanded"))
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
		ops := Diff(stage, task, taskOther, "Outputs", taskOther.Outputs, task.Outputs)
		diffs = append(diffs, ops)
	}
	if task.IsOutputsNodeExpanded != taskOther.IsOutputsNodeExpanded {
		diffs = append(diffs, task.GongMarshallField(stage, "IsOutputsNodeExpanded"))
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
		ops := Diff(stage, task, taskOther, "TaskGroupsToDisplay", taskOther.TaskGroupsToDisplay, task.TaskGroupsToDisplay)
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
		ops := Diff(stage, task, taskOther, "SubTasks", taskOther.SubTasks, task.SubTasks)
		diffs = append(diffs, ops)
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
		ops := Diff(stage, taskgroup, taskgroupOther, "Tasks", taskgroupOther.Tasks, taskgroup.Tasks)
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

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
