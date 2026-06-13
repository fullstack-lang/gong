// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *DiagramHierarchy:
		ok = stage.IsStagedDiagramHierarchy(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Milestone:
		ok = stage.IsStagedMilestone(target)

	case *MilestoneShape:
		ok = stage.IsStagedMilestoneShape(target)

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
	case *DiagramHierarchy:
		ok = stage.IsStagedDiagramHierarchy(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Milestone:
		ok = stage.IsStagedMilestone(target)

	case *MilestoneShape:
		ok = stage.IsStagedMilestoneShape(target)

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
func (stage *Stage) IsStagedDiagramHierarchy(diagramhierarchy *DiagramHierarchy) (ok bool) {

	_, ok = stage.DiagramHierarchys[diagramhierarchy]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedMilestone(milestone *Milestone) (ok bool) {

	_, ok = stage.Milestones[milestone]

	return
}

func (stage *Stage) IsStagedMilestoneShape(milestoneshape *MilestoneShape) (ok bool) {

	_, ok = stage.MilestoneShapes[milestoneshape]

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
	case *DiagramHierarchy:
		stage.StageBranchDiagramHierarchy(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Milestone:
		stage.StageBranchMilestone(target)

	case *MilestoneShape:
		stage.StageBranchMilestoneShape(target)

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
func (stage *Stage) StageBranchDiagramHierarchy(diagramhierarchy *DiagramHierarchy) {

	// check if instance is already staged
	if IsStaged(stage, diagramhierarchy) {
		return
	}

	diagramhierarchy.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagramhierarchy.Product_Shapes {
		StageBranch(stage, _productshape)
	}
	for _, _product := range diagramhierarchy.ProductsWhoseNodeIsExpanded {
		StageBranch(stage, _product)
	}
	for _, _productcompositionshape := range diagramhierarchy.ProductComposition_Shapes {
		StageBranch(stage, _productcompositionshape)
	}
	for _, _taskshape := range diagramhierarchy.Task_Shapes {
		StageBranch(stage, _taskshape)
	}
	for _, _task := range diagramhierarchy.TasksWhoseNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range diagramhierarchy.TasksWhoseInputNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _task := range diagramhierarchy.TasksWhoseOutputNodeIsExpanded {
		StageBranch(stage, _task)
	}
	for _, _taskgroupshape := range diagramhierarchy.TaskGroupShapes {
		StageBranch(stage, _taskgroupshape)
	}
	for _, _taskgroup := range diagramhierarchy.TaskGroupsWhoseNodeIsExpanded {
		StageBranch(stage, _taskgroup)
	}
	for _, _milestoneshape := range diagramhierarchy.MilestoneShapes {
		StageBranch(stage, _milestoneshape)
	}
	for _, _milestone := range diagramhierarchy.MilestonesWhoseNodeIsExpanded {
		StageBranch(stage, _milestone)
	}
	for _, _taskcompositionshape := range diagramhierarchy.TaskComposition_Shapes {
		StageBranch(stage, _taskcompositionshape)
	}
	for _, _taskinputshape := range diagramhierarchy.TaskInputShapes {
		StageBranch(stage, _taskinputshape)
	}
	for _, _taskoutputshape := range diagramhierarchy.TaskOutputShapes {
		StageBranch(stage, _taskoutputshape)
	}
	for _, _noteshape := range diagramhierarchy.Note_Shapes {
		StageBranch(stage, _noteshape)
	}
	for _, _note := range diagramhierarchy.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}
	for _, _noteproductshape := range diagramhierarchy.NoteProductShapes {
		StageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagramhierarchy.NoteTaskShapes {
		StageBranch(stage, _notetaskshape)
	}
	for _, _noteresourceshape := range diagramhierarchy.NoteResourceShapes {
		StageBranch(stage, _noteresourceshape)
	}
	for _, _resourceshape := range diagramhierarchy.Resource_Shapes {
		StageBranch(stage, _resourceshape)
	}
	for _, _resource := range diagramhierarchy.ResourcesWhoseNodeIsExpanded {
		StageBranch(stage, _resource)
	}
	for _, _resourcecompositionshape := range diagramhierarchy.ResourceComposition_Shapes {
		StageBranch(stage, _resourcecompositionshape)
	}
	for _, _resourcetaskshape := range diagramhierarchy.ResourceTaskShapes {
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
	for _, _milestone := range library.RootMilestones {
		StageBranch(stage, _milestone)
	}
	for _, _resource := range library.RootResources {
		StageBranch(stage, _resource)
	}
	for _, _note := range library.Notes {
		StageBranch(stage, _note)
	}
	for _, _diagramhierarchy := range library.Diagrams {
		StageBranch(stage, _diagramhierarchy)
	}

}

func (stage *Stage) StageBranchMilestone(milestone *Milestone) {

	// check if instance is already staged
	if IsStaged(stage, milestone) {
		return
	}

	milestone.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _taskgroup := range milestone.TaskGroupsToDisplay {
		StageBranch(stage, _taskgroup)
	}

}

func (stage *Stage) StageBranchMilestoneShape(milestoneshape *MilestoneShape) {

	// check if instance is already staged
	if IsStaged(stage, milestoneshape) {
		return
	}

	milestoneshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if milestoneshape.Milestone != nil {
		StageBranch(stage, milestoneshape.Milestone)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	for _, _task := range task.SubTasks {
		StageBranch(stage, _task)
	}
	for _, _product := range task.Inputs {
		StageBranch(stage, _product)
	}
	for _, _product := range task.Outputs {
		StageBranch(stage, _product)
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
	case *DiagramHierarchy:
		toT := CopyBranchDiagramHierarchy(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Milestone:
		toT := CopyBranchMilestone(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MilestoneShape:
		toT := CopyBranchMilestoneShape(mapOrigCopy, fromT)
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
func CopyBranchDiagramHierarchy(mapOrigCopy map[any]any, diagramhierarchyFrom *DiagramHierarchy) (diagramhierarchyTo *DiagramHierarchy) {

	// diagramhierarchyFrom has already been copied
	if _diagramhierarchyTo, ok := mapOrigCopy[diagramhierarchyFrom]; ok {
		diagramhierarchyTo = _diagramhierarchyTo.(*DiagramHierarchy)
		return
	}

	diagramhierarchyTo = new(DiagramHierarchy)
	mapOrigCopy[diagramhierarchyFrom] = diagramhierarchyTo
	diagramhierarchyFrom.CopyBasicFields(diagramhierarchyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagramhierarchyFrom.Product_Shapes {
		diagramhierarchyTo.Product_Shapes = append(diagramhierarchyTo.Product_Shapes, CopyBranchProductShape(mapOrigCopy, _productshape))
	}
	for _, _product := range diagramhierarchyFrom.ProductsWhoseNodeIsExpanded {
		diagramhierarchyTo.ProductsWhoseNodeIsExpanded = append(diagramhierarchyTo.ProductsWhoseNodeIsExpanded, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _productcompositionshape := range diagramhierarchyFrom.ProductComposition_Shapes {
		diagramhierarchyTo.ProductComposition_Shapes = append(diagramhierarchyTo.ProductComposition_Shapes, CopyBranchProductCompositionShape(mapOrigCopy, _productcompositionshape))
	}
	for _, _taskshape := range diagramhierarchyFrom.Task_Shapes {
		diagramhierarchyTo.Task_Shapes = append(diagramhierarchyTo.Task_Shapes, CopyBranchTaskShape(mapOrigCopy, _taskshape))
	}
	for _, _task := range diagramhierarchyFrom.TasksWhoseNodeIsExpanded {
		diagramhierarchyTo.TasksWhoseNodeIsExpanded = append(diagramhierarchyTo.TasksWhoseNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramhierarchyFrom.TasksWhoseInputNodeIsExpanded {
		diagramhierarchyTo.TasksWhoseInputNodeIsExpanded = append(diagramhierarchyTo.TasksWhoseInputNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _task := range diagramhierarchyFrom.TasksWhoseOutputNodeIsExpanded {
		diagramhierarchyTo.TasksWhoseOutputNodeIsExpanded = append(diagramhierarchyTo.TasksWhoseOutputNodeIsExpanded, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _taskgroupshape := range diagramhierarchyFrom.TaskGroupShapes {
		diagramhierarchyTo.TaskGroupShapes = append(diagramhierarchyTo.TaskGroupShapes, CopyBranchTaskGroupShape(mapOrigCopy, _taskgroupshape))
	}
	for _, _taskgroup := range diagramhierarchyFrom.TaskGroupsWhoseNodeIsExpanded {
		diagramhierarchyTo.TaskGroupsWhoseNodeIsExpanded = append(diagramhierarchyTo.TaskGroupsWhoseNodeIsExpanded, CopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}
	for _, _milestoneshape := range diagramhierarchyFrom.MilestoneShapes {
		diagramhierarchyTo.MilestoneShapes = append(diagramhierarchyTo.MilestoneShapes, CopyBranchMilestoneShape(mapOrigCopy, _milestoneshape))
	}
	for _, _milestone := range diagramhierarchyFrom.MilestonesWhoseNodeIsExpanded {
		diagramhierarchyTo.MilestonesWhoseNodeIsExpanded = append(diagramhierarchyTo.MilestonesWhoseNodeIsExpanded, CopyBranchMilestone(mapOrigCopy, _milestone))
	}
	for _, _taskcompositionshape := range diagramhierarchyFrom.TaskComposition_Shapes {
		diagramhierarchyTo.TaskComposition_Shapes = append(diagramhierarchyTo.TaskComposition_Shapes, CopyBranchTaskCompositionShape(mapOrigCopy, _taskcompositionshape))
	}
	for _, _taskinputshape := range diagramhierarchyFrom.TaskInputShapes {
		diagramhierarchyTo.TaskInputShapes = append(diagramhierarchyTo.TaskInputShapes, CopyBranchTaskInputShape(mapOrigCopy, _taskinputshape))
	}
	for _, _taskoutputshape := range diagramhierarchyFrom.TaskOutputShapes {
		diagramhierarchyTo.TaskOutputShapes = append(diagramhierarchyTo.TaskOutputShapes, CopyBranchTaskOutputShape(mapOrigCopy, _taskoutputshape))
	}
	for _, _noteshape := range diagramhierarchyFrom.Note_Shapes {
		diagramhierarchyTo.Note_Shapes = append(diagramhierarchyTo.Note_Shapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramhierarchyFrom.NotesWhoseNodeIsExpanded {
		diagramhierarchyTo.NotesWhoseNodeIsExpanded = append(diagramhierarchyTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _noteproductshape := range diagramhierarchyFrom.NoteProductShapes {
		diagramhierarchyTo.NoteProductShapes = append(diagramhierarchyTo.NoteProductShapes, CopyBranchNoteProductShape(mapOrigCopy, _noteproductshape))
	}
	for _, _notetaskshape := range diagramhierarchyFrom.NoteTaskShapes {
		diagramhierarchyTo.NoteTaskShapes = append(diagramhierarchyTo.NoteTaskShapes, CopyBranchNoteTaskShape(mapOrigCopy, _notetaskshape))
	}
	for _, _noteresourceshape := range diagramhierarchyFrom.NoteResourceShapes {
		diagramhierarchyTo.NoteResourceShapes = append(diagramhierarchyTo.NoteResourceShapes, CopyBranchNoteResourceShape(mapOrigCopy, _noteresourceshape))
	}
	for _, _resourceshape := range diagramhierarchyFrom.Resource_Shapes {
		diagramhierarchyTo.Resource_Shapes = append(diagramhierarchyTo.Resource_Shapes, CopyBranchResourceShape(mapOrigCopy, _resourceshape))
	}
	for _, _resource := range diagramhierarchyFrom.ResourcesWhoseNodeIsExpanded {
		diagramhierarchyTo.ResourcesWhoseNodeIsExpanded = append(diagramhierarchyTo.ResourcesWhoseNodeIsExpanded, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _resourcecompositionshape := range diagramhierarchyFrom.ResourceComposition_Shapes {
		diagramhierarchyTo.ResourceComposition_Shapes = append(diagramhierarchyTo.ResourceComposition_Shapes, CopyBranchResourceCompositionShape(mapOrigCopy, _resourcecompositionshape))
	}
	for _, _resourcetaskshape := range diagramhierarchyFrom.ResourceTaskShapes {
		diagramhierarchyTo.ResourceTaskShapes = append(diagramhierarchyTo.ResourceTaskShapes, CopyBranchResourceTaskShape(mapOrigCopy, _resourcetaskshape))
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
	for _, _milestone := range libraryFrom.RootMilestones {
		libraryTo.RootMilestones = append(libraryTo.RootMilestones, CopyBranchMilestone(mapOrigCopy, _milestone))
	}
	for _, _resource := range libraryFrom.RootResources {
		libraryTo.RootResources = append(libraryTo.RootResources, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _note := range libraryFrom.Notes {
		libraryTo.Notes = append(libraryTo.Notes, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _diagramhierarchy := range libraryFrom.Diagrams {
		libraryTo.Diagrams = append(libraryTo.Diagrams, CopyBranchDiagramHierarchy(mapOrigCopy, _diagramhierarchy))
	}

	return
}

func CopyBranchMilestone(mapOrigCopy map[any]any, milestoneFrom *Milestone) (milestoneTo *Milestone) {

	// milestoneFrom has already been copied
	if _milestoneTo, ok := mapOrigCopy[milestoneFrom]; ok {
		milestoneTo = _milestoneTo.(*Milestone)
		return
	}

	milestoneTo = new(Milestone)
	mapOrigCopy[milestoneFrom] = milestoneTo
	milestoneFrom.CopyBasicFields(milestoneTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _taskgroup := range milestoneFrom.TaskGroupsToDisplay {
		milestoneTo.TaskGroupsToDisplay = append(milestoneTo.TaskGroupsToDisplay, CopyBranchTaskGroup(mapOrigCopy, _taskgroup))
	}

	return
}

func CopyBranchMilestoneShape(mapOrigCopy map[any]any, milestoneshapeFrom *MilestoneShape) (milestoneshapeTo *MilestoneShape) {

	// milestoneshapeFrom has already been copied
	if _milestoneshapeTo, ok := mapOrigCopy[milestoneshapeFrom]; ok {
		milestoneshapeTo = _milestoneshapeTo.(*MilestoneShape)
		return
	}

	milestoneshapeTo = new(MilestoneShape)
	mapOrigCopy[milestoneshapeFrom] = milestoneshapeTo
	milestoneshapeFrom.CopyBasicFields(milestoneshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if milestoneshapeFrom.Milestone != nil {
		milestoneshapeTo.Milestone = CopyBranchMilestone(mapOrigCopy, milestoneshapeFrom.Milestone)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	for _, _task := range taskFrom.SubTasks {
		taskTo.SubTasks = append(taskTo.SubTasks, CopyBranchTask(mapOrigCopy, _task))
	}
	for _, _product := range taskFrom.Inputs {
		taskTo.Inputs = append(taskTo.Inputs, CopyBranchProduct(mapOrigCopy, _product))
	}
	for _, _product := range taskFrom.Outputs {
		taskTo.Outputs = append(taskTo.Outputs, CopyBranchProduct(mapOrigCopy, _product))
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
	case *DiagramHierarchy:
		stage.UnstageBranchDiagramHierarchy(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Milestone:
		stage.UnstageBranchMilestone(target)

	case *MilestoneShape:
		stage.UnstageBranchMilestoneShape(target)

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
func (stage *Stage) UnstageBranchDiagramHierarchy(diagramhierarchy *DiagramHierarchy) {

	// check if instance is already staged
	if !IsStaged(stage, diagramhierarchy) {
		return
	}

	diagramhierarchy.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _productshape := range diagramhierarchy.Product_Shapes {
		UnstageBranch(stage, _productshape)
	}
	for _, _product := range diagramhierarchy.ProductsWhoseNodeIsExpanded {
		UnstageBranch(stage, _product)
	}
	for _, _productcompositionshape := range diagramhierarchy.ProductComposition_Shapes {
		UnstageBranch(stage, _productcompositionshape)
	}
	for _, _taskshape := range diagramhierarchy.Task_Shapes {
		UnstageBranch(stage, _taskshape)
	}
	for _, _task := range diagramhierarchy.TasksWhoseNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range diagramhierarchy.TasksWhoseInputNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _task := range diagramhierarchy.TasksWhoseOutputNodeIsExpanded {
		UnstageBranch(stage, _task)
	}
	for _, _taskgroupshape := range diagramhierarchy.TaskGroupShapes {
		UnstageBranch(stage, _taskgroupshape)
	}
	for _, _taskgroup := range diagramhierarchy.TaskGroupsWhoseNodeIsExpanded {
		UnstageBranch(stage, _taskgroup)
	}
	for _, _milestoneshape := range diagramhierarchy.MilestoneShapes {
		UnstageBranch(stage, _milestoneshape)
	}
	for _, _milestone := range diagramhierarchy.MilestonesWhoseNodeIsExpanded {
		UnstageBranch(stage, _milestone)
	}
	for _, _taskcompositionshape := range diagramhierarchy.TaskComposition_Shapes {
		UnstageBranch(stage, _taskcompositionshape)
	}
	for _, _taskinputshape := range diagramhierarchy.TaskInputShapes {
		UnstageBranch(stage, _taskinputshape)
	}
	for _, _taskoutputshape := range diagramhierarchy.TaskOutputShapes {
		UnstageBranch(stage, _taskoutputshape)
	}
	for _, _noteshape := range diagramhierarchy.Note_Shapes {
		UnstageBranch(stage, _noteshape)
	}
	for _, _note := range diagramhierarchy.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}
	for _, _noteproductshape := range diagramhierarchy.NoteProductShapes {
		UnstageBranch(stage, _noteproductshape)
	}
	for _, _notetaskshape := range diagramhierarchy.NoteTaskShapes {
		UnstageBranch(stage, _notetaskshape)
	}
	for _, _noteresourceshape := range diagramhierarchy.NoteResourceShapes {
		UnstageBranch(stage, _noteresourceshape)
	}
	for _, _resourceshape := range diagramhierarchy.Resource_Shapes {
		UnstageBranch(stage, _resourceshape)
	}
	for _, _resource := range diagramhierarchy.ResourcesWhoseNodeIsExpanded {
		UnstageBranch(stage, _resource)
	}
	for _, _resourcecompositionshape := range diagramhierarchy.ResourceComposition_Shapes {
		UnstageBranch(stage, _resourcecompositionshape)
	}
	for _, _resourcetaskshape := range diagramhierarchy.ResourceTaskShapes {
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
	for _, _milestone := range library.RootMilestones {
		UnstageBranch(stage, _milestone)
	}
	for _, _resource := range library.RootResources {
		UnstageBranch(stage, _resource)
	}
	for _, _note := range library.Notes {
		UnstageBranch(stage, _note)
	}
	for _, _diagramhierarchy := range library.Diagrams {
		UnstageBranch(stage, _diagramhierarchy)
	}

}

func (stage *Stage) UnstageBranchMilestone(milestone *Milestone) {

	// check if instance is already staged
	if !IsStaged(stage, milestone) {
		return
	}

	milestone.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _taskgroup := range milestone.TaskGroupsToDisplay {
		UnstageBranch(stage, _taskgroup)
	}

}

func (stage *Stage) UnstageBranchMilestoneShape(milestoneshape *MilestoneShape) {

	// check if instance is already staged
	if !IsStaged(stage, milestoneshape) {
		return
	}

	milestoneshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if milestoneshape.Milestone != nil {
		UnstageBranch(stage, milestoneshape.Milestone)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
	for _, _task := range task.SubTasks {
		UnstageBranch(stage, _task)
	}
	for _, _product := range task.Inputs {
		UnstageBranch(stage, _product)
	}
	for _, _product := range task.Outputs {
		UnstageBranch(stage, _product)
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
func (reference *DiagramHierarchy) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramHierarchy) {
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
	reference.MilestoneShapes = reference.MilestoneShapes[:0]
	for _, _b := range instance.MilestoneShapes {
		reference.MilestoneShapes = append(reference.MilestoneShapes, stage.MilestoneShapes_reference[_b])
	}
	reference.MilestonesWhoseNodeIsExpanded = reference.MilestonesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.MilestonesWhoseNodeIsExpanded {
		reference.MilestonesWhoseNodeIsExpanded = append(reference.MilestonesWhoseNodeIsExpanded, stage.Milestones_reference[_b])
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
	reference.RootMilestones = reference.RootMilestones[:0]
	for _, _b := range instance.RootMilestones {
		reference.RootMilestones = append(reference.RootMilestones, stage.Milestones_reference[_b])
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
		reference.Diagrams = append(reference.Diagrams, stage.DiagramHierarchys_reference[_b])
	}
}

func (reference *Milestone) GongReconstructPointersFromReferences(stage *Stage, instance *Milestone) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.TaskGroupsToDisplay = reference.TaskGroupsToDisplay[:0]
	for _, _b := range instance.TaskGroupsToDisplay {
		reference.TaskGroupsToDisplay = append(reference.TaskGroupsToDisplay, stage.TaskGroups_reference[_b])
	}
}

func (reference *MilestoneShape) GongReconstructPointersFromReferences(stage *Stage, instance *MilestoneShape) {
	// insertion point for pointers field
	if instance.Milestone != nil {
		reference.Milestone = stage.Milestones_reference[instance.Milestone]
	}
	// insertion point for slice of pointers field
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
	reference.SubTasks = reference.SubTasks[:0]
	for _, _b := range instance.SubTasks {
		reference.SubTasks = append(reference.SubTasks, stage.Tasks_reference[_b])
	}
	reference.Inputs = reference.Inputs[:0]
	for _, _b := range instance.Inputs {
		reference.Inputs = append(reference.Inputs, stage.Products_reference[_b])
	}
	reference.Outputs = reference.Outputs[:0]
	for _, _b := range instance.Outputs {
		reference.Outputs = append(reference.Outputs, stage.Products_reference[_b])
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
func (reference *DiagramHierarchy) GongReconstructPointersFromInstances(stage *Stage) {
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
	var _MilestoneShapes []*MilestoneShape
	for _, _reference := range reference.MilestoneShapes {
		if _instance, ok := stage.MilestoneShapes_instance[_reference]; ok {
			_MilestoneShapes = append(_MilestoneShapes, _instance)
		}
	}
	reference.MilestoneShapes = _MilestoneShapes
	var _MilestonesWhoseNodeIsExpanded []*Milestone
	for _, _reference := range reference.MilestonesWhoseNodeIsExpanded {
		if _instance, ok := stage.Milestones_instance[_reference]; ok {
			_MilestonesWhoseNodeIsExpanded = append(_MilestonesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.MilestonesWhoseNodeIsExpanded = _MilestonesWhoseNodeIsExpanded
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
	var _RootMilestones []*Milestone
	for _, _reference := range reference.RootMilestones {
		if _instance, ok := stage.Milestones_instance[_reference]; ok {
			_RootMilestones = append(_RootMilestones, _instance)
		}
	}
	reference.RootMilestones = _RootMilestones
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
	var _Diagrams []*DiagramHierarchy
	for _, _reference := range reference.Diagrams {
		if _instance, ok := stage.DiagramHierarchys_instance[_reference]; ok {
			_Diagrams = append(_Diagrams, _instance)
		}
	}
	reference.Diagrams = _Diagrams
}

func (reference *Milestone) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _TaskGroupsToDisplay []*TaskGroup
	for _, _reference := range reference.TaskGroupsToDisplay {
		if _instance, ok := stage.TaskGroups_instance[_reference]; ok {
			_TaskGroupsToDisplay = append(_TaskGroupsToDisplay, _instance)
		}
	}
	reference.TaskGroupsToDisplay = _TaskGroupsToDisplay
}

func (reference *MilestoneShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Milestone; _reference != nil {
		reference.Milestone = nil
		if _instance, ok := stage.Milestones_instance[_reference]; ok {
			reference.Milestone = _instance
		}
	}
	// insertion point for slice of pointers fields
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
	var _SubTasks []*Task
	for _, _reference := range reference.SubTasks {
		if _instance, ok := stage.Tasks_instance[_reference]; ok {
			_SubTasks = append(_SubTasks, _instance)
		}
	}
	reference.SubTasks = _SubTasks
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
func (diagramhierarchy *DiagramHierarchy) GongDiff(stage *Stage, diagramhierarchyOther *DiagramHierarchy) (diffs []string) {
	// insertion point for field diffs
	if diagramhierarchy.Name != diagramhierarchyOther.Name {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "Name"))
	}
	if diagramhierarchy.ComputedPrefix != diagramhierarchyOther.ComputedPrefix {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagramhierarchy.IsExpanded != diagramhierarchyOther.IsExpanded {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsExpanded"))
	}
	if diagramhierarchy.IsChecked != diagramhierarchyOther.IsChecked {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsChecked"))
	}
	if diagramhierarchy.IsEditable_ != diagramhierarchyOther.IsEditable_ {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsEditable_"))
	}
	if diagramhierarchy.IsShowPrefix != diagramhierarchyOther.IsShowPrefix {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsShowPrefix"))
	}
	if diagramhierarchy.DefaultBoxWidth != diagramhierarchyOther.DefaultBoxWidth {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagramhierarchy.DefaultBoxHeigth != diagramhierarchyOther.DefaultBoxHeigth {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "DefaultBoxHeigth"))
	}
	if diagramhierarchy.Width != diagramhierarchyOther.Width {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "Width"))
	}
	if diagramhierarchy.Height != diagramhierarchyOther.Height {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "Height"))
	}
	Product_ShapesDifferent := false
	if len(diagramhierarchy.Product_Shapes) != len(diagramhierarchyOther.Product_Shapes) {
		Product_ShapesDifferent = true
	} else {
		for i := range diagramhierarchy.Product_Shapes {
			if (diagramhierarchy.Product_Shapes[i] == nil) != (diagramhierarchyOther.Product_Shapes[i] == nil) {
				Product_ShapesDifferent = true
				break
			} else if diagramhierarchy.Product_Shapes[i] != nil && diagramhierarchyOther.Product_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.Product_Shapes[i] != diagramhierarchyOther.Product_Shapes[i] {
					Product_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Product_ShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "Product_Shapes", diagramhierarchyOther.Product_Shapes, diagramhierarchy.Product_Shapes)
		diffs = append(diffs, ops)
	}
	ProductsWhoseNodeIsExpandedDifferent := false
	if len(diagramhierarchy.ProductsWhoseNodeIsExpanded) != len(diagramhierarchyOther.ProductsWhoseNodeIsExpanded) {
		ProductsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.ProductsWhoseNodeIsExpanded {
			if (diagramhierarchy.ProductsWhoseNodeIsExpanded[i] == nil) != (diagramhierarchyOther.ProductsWhoseNodeIsExpanded[i] == nil) {
				ProductsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.ProductsWhoseNodeIsExpanded[i] != nil && diagramhierarchyOther.ProductsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.ProductsWhoseNodeIsExpanded[i] != diagramhierarchyOther.ProductsWhoseNodeIsExpanded[i] {
					ProductsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ProductsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "ProductsWhoseNodeIsExpanded", diagramhierarchyOther.ProductsWhoseNodeIsExpanded, diagramhierarchy.ProductsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.IsPBSNodeExpanded != diagramhierarchyOther.IsPBSNodeExpanded {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsPBSNodeExpanded"))
	}
	ProductComposition_ShapesDifferent := false
	if len(diagramhierarchy.ProductComposition_Shapes) != len(diagramhierarchyOther.ProductComposition_Shapes) {
		ProductComposition_ShapesDifferent = true
	} else {
		for i := range diagramhierarchy.ProductComposition_Shapes {
			if (diagramhierarchy.ProductComposition_Shapes[i] == nil) != (diagramhierarchyOther.ProductComposition_Shapes[i] == nil) {
				ProductComposition_ShapesDifferent = true
				break
			} else if diagramhierarchy.ProductComposition_Shapes[i] != nil && diagramhierarchyOther.ProductComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.ProductComposition_Shapes[i] != diagramhierarchyOther.ProductComposition_Shapes[i] {
					ProductComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ProductComposition_ShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "ProductComposition_Shapes", diagramhierarchyOther.ProductComposition_Shapes, diagramhierarchy.ProductComposition_Shapes)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.IsWBSNodeExpanded != diagramhierarchyOther.IsWBSNodeExpanded {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsWBSNodeExpanded"))
	}
	Task_ShapesDifferent := false
	if len(diagramhierarchy.Task_Shapes) != len(diagramhierarchyOther.Task_Shapes) {
		Task_ShapesDifferent = true
	} else {
		for i := range diagramhierarchy.Task_Shapes {
			if (diagramhierarchy.Task_Shapes[i] == nil) != (diagramhierarchyOther.Task_Shapes[i] == nil) {
				Task_ShapesDifferent = true
				break
			} else if diagramhierarchy.Task_Shapes[i] != nil && diagramhierarchyOther.Task_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.Task_Shapes[i] != diagramhierarchyOther.Task_Shapes[i] {
					Task_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Task_ShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "Task_Shapes", diagramhierarchyOther.Task_Shapes, diagramhierarchy.Task_Shapes)
		diffs = append(diffs, ops)
	}
	TasksWhoseNodeIsExpandedDifferent := false
	if len(diagramhierarchy.TasksWhoseNodeIsExpanded) != len(diagramhierarchyOther.TasksWhoseNodeIsExpanded) {
		TasksWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.TasksWhoseNodeIsExpanded {
			if (diagramhierarchy.TasksWhoseNodeIsExpanded[i] == nil) != (diagramhierarchyOther.TasksWhoseNodeIsExpanded[i] == nil) {
				TasksWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.TasksWhoseNodeIsExpanded[i] != nil && diagramhierarchyOther.TasksWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TasksWhoseNodeIsExpanded[i] != diagramhierarchyOther.TasksWhoseNodeIsExpanded[i] {
					TasksWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TasksWhoseNodeIsExpanded", diagramhierarchyOther.TasksWhoseNodeIsExpanded, diagramhierarchy.TasksWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	TasksWhoseInputNodeIsExpandedDifferent := false
	if len(diagramhierarchy.TasksWhoseInputNodeIsExpanded) != len(diagramhierarchyOther.TasksWhoseInputNodeIsExpanded) {
		TasksWhoseInputNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.TasksWhoseInputNodeIsExpanded {
			if (diagramhierarchy.TasksWhoseInputNodeIsExpanded[i] == nil) != (diagramhierarchyOther.TasksWhoseInputNodeIsExpanded[i] == nil) {
				TasksWhoseInputNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.TasksWhoseInputNodeIsExpanded[i] != nil && diagramhierarchyOther.TasksWhoseInputNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TasksWhoseInputNodeIsExpanded[i] != diagramhierarchyOther.TasksWhoseInputNodeIsExpanded[i] {
					TasksWhoseInputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseInputNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TasksWhoseInputNodeIsExpanded", diagramhierarchyOther.TasksWhoseInputNodeIsExpanded, diagramhierarchy.TasksWhoseInputNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	TasksWhoseOutputNodeIsExpandedDifferent := false
	if len(diagramhierarchy.TasksWhoseOutputNodeIsExpanded) != len(diagramhierarchyOther.TasksWhoseOutputNodeIsExpanded) {
		TasksWhoseOutputNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.TasksWhoseOutputNodeIsExpanded {
			if (diagramhierarchy.TasksWhoseOutputNodeIsExpanded[i] == nil) != (diagramhierarchyOther.TasksWhoseOutputNodeIsExpanded[i] == nil) {
				TasksWhoseOutputNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.TasksWhoseOutputNodeIsExpanded[i] != nil && diagramhierarchyOther.TasksWhoseOutputNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TasksWhoseOutputNodeIsExpanded[i] != diagramhierarchyOther.TasksWhoseOutputNodeIsExpanded[i] {
					TasksWhoseOutputNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TasksWhoseOutputNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TasksWhoseOutputNodeIsExpanded", diagramhierarchyOther.TasksWhoseOutputNodeIsExpanded, diagramhierarchy.TasksWhoseOutputNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.IsTaskGroupsNodeExpanded != diagramhierarchyOther.IsTaskGroupsNodeExpanded {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsTaskGroupsNodeExpanded"))
	}
	TaskGroupShapesDifferent := false
	if len(diagramhierarchy.TaskGroupShapes) != len(diagramhierarchyOther.TaskGroupShapes) {
		TaskGroupShapesDifferent = true
	} else {
		for i := range diagramhierarchy.TaskGroupShapes {
			if (diagramhierarchy.TaskGroupShapes[i] == nil) != (diagramhierarchyOther.TaskGroupShapes[i] == nil) {
				TaskGroupShapesDifferent = true
				break
			} else if diagramhierarchy.TaskGroupShapes[i] != nil && diagramhierarchyOther.TaskGroupShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TaskGroupShapes[i] != diagramhierarchyOther.TaskGroupShapes[i] {
					TaskGroupShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskGroupShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TaskGroupShapes", diagramhierarchyOther.TaskGroupShapes, diagramhierarchy.TaskGroupShapes)
		diffs = append(diffs, ops)
	}
	TaskGroupsWhoseNodeIsExpandedDifferent := false
	if len(diagramhierarchy.TaskGroupsWhoseNodeIsExpanded) != len(diagramhierarchyOther.TaskGroupsWhoseNodeIsExpanded) {
		TaskGroupsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.TaskGroupsWhoseNodeIsExpanded {
			if (diagramhierarchy.TaskGroupsWhoseNodeIsExpanded[i] == nil) != (diagramhierarchyOther.TaskGroupsWhoseNodeIsExpanded[i] == nil) {
				TaskGroupsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.TaskGroupsWhoseNodeIsExpanded[i] != nil && diagramhierarchyOther.TaskGroupsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TaskGroupsWhoseNodeIsExpanded[i] != diagramhierarchyOther.TaskGroupsWhoseNodeIsExpanded[i] {
					TaskGroupsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if TaskGroupsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TaskGroupsWhoseNodeIsExpanded", diagramhierarchyOther.TaskGroupsWhoseNodeIsExpanded, diagramhierarchy.TaskGroupsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.IsMilestonesNodeExpanded != diagramhierarchyOther.IsMilestonesNodeExpanded {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsMilestonesNodeExpanded"))
	}
	MilestoneShapesDifferent := false
	if len(diagramhierarchy.MilestoneShapes) != len(diagramhierarchyOther.MilestoneShapes) {
		MilestoneShapesDifferent = true
	} else {
		for i := range diagramhierarchy.MilestoneShapes {
			if (diagramhierarchy.MilestoneShapes[i] == nil) != (diagramhierarchyOther.MilestoneShapes[i] == nil) {
				MilestoneShapesDifferent = true
				break
			} else if diagramhierarchy.MilestoneShapes[i] != nil && diagramhierarchyOther.MilestoneShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.MilestoneShapes[i] != diagramhierarchyOther.MilestoneShapes[i] {
					MilestoneShapesDifferent = true
					break
				}
			}
		}
	}
	if MilestoneShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "MilestoneShapes", diagramhierarchyOther.MilestoneShapes, diagramhierarchy.MilestoneShapes)
		diffs = append(diffs, ops)
	}
	MilestonesWhoseNodeIsExpandedDifferent := false
	if len(diagramhierarchy.MilestonesWhoseNodeIsExpanded) != len(diagramhierarchyOther.MilestonesWhoseNodeIsExpanded) {
		MilestonesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.MilestonesWhoseNodeIsExpanded {
			if (diagramhierarchy.MilestonesWhoseNodeIsExpanded[i] == nil) != (diagramhierarchyOther.MilestonesWhoseNodeIsExpanded[i] == nil) {
				MilestonesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.MilestonesWhoseNodeIsExpanded[i] != nil && diagramhierarchyOther.MilestonesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.MilestonesWhoseNodeIsExpanded[i] != diagramhierarchyOther.MilestonesWhoseNodeIsExpanded[i] {
					MilestonesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if MilestonesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "MilestonesWhoseNodeIsExpanded", diagramhierarchyOther.MilestonesWhoseNodeIsExpanded, diagramhierarchy.MilestonesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.DateFormat != diagramhierarchyOther.DateFormat {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "DateFormat"))
	}
	TaskComposition_ShapesDifferent := false
	if len(diagramhierarchy.TaskComposition_Shapes) != len(diagramhierarchyOther.TaskComposition_Shapes) {
		TaskComposition_ShapesDifferent = true
	} else {
		for i := range diagramhierarchy.TaskComposition_Shapes {
			if (diagramhierarchy.TaskComposition_Shapes[i] == nil) != (diagramhierarchyOther.TaskComposition_Shapes[i] == nil) {
				TaskComposition_ShapesDifferent = true
				break
			} else if diagramhierarchy.TaskComposition_Shapes[i] != nil && diagramhierarchyOther.TaskComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TaskComposition_Shapes[i] != diagramhierarchyOther.TaskComposition_Shapes[i] {
					TaskComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskComposition_ShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TaskComposition_Shapes", diagramhierarchyOther.TaskComposition_Shapes, diagramhierarchy.TaskComposition_Shapes)
		diffs = append(diffs, ops)
	}
	TaskInputShapesDifferent := false
	if len(diagramhierarchy.TaskInputShapes) != len(diagramhierarchyOther.TaskInputShapes) {
		TaskInputShapesDifferent = true
	} else {
		for i := range diagramhierarchy.TaskInputShapes {
			if (diagramhierarchy.TaskInputShapes[i] == nil) != (diagramhierarchyOther.TaskInputShapes[i] == nil) {
				TaskInputShapesDifferent = true
				break
			} else if diagramhierarchy.TaskInputShapes[i] != nil && diagramhierarchyOther.TaskInputShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TaskInputShapes[i] != diagramhierarchyOther.TaskInputShapes[i] {
					TaskInputShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskInputShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TaskInputShapes", diagramhierarchyOther.TaskInputShapes, diagramhierarchy.TaskInputShapes)
		diffs = append(diffs, ops)
	}
	TaskOutputShapesDifferent := false
	if len(diagramhierarchy.TaskOutputShapes) != len(diagramhierarchyOther.TaskOutputShapes) {
		TaskOutputShapesDifferent = true
	} else {
		for i := range diagramhierarchy.TaskOutputShapes {
			if (diagramhierarchy.TaskOutputShapes[i] == nil) != (diagramhierarchyOther.TaskOutputShapes[i] == nil) {
				TaskOutputShapesDifferent = true
				break
			} else if diagramhierarchy.TaskOutputShapes[i] != nil && diagramhierarchyOther.TaskOutputShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.TaskOutputShapes[i] != diagramhierarchyOther.TaskOutputShapes[i] {
					TaskOutputShapesDifferent = true
					break
				}
			}
		}
	}
	if TaskOutputShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "TaskOutputShapes", diagramhierarchyOther.TaskOutputShapes, diagramhierarchy.TaskOutputShapes)
		diffs = append(diffs, ops)
	}
	Note_ShapesDifferent := false
	if len(diagramhierarchy.Note_Shapes) != len(diagramhierarchyOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagramhierarchy.Note_Shapes {
			if (diagramhierarchy.Note_Shapes[i] == nil) != (diagramhierarchyOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagramhierarchy.Note_Shapes[i] != nil && diagramhierarchyOther.Note_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.Note_Shapes[i] != diagramhierarchyOther.Note_Shapes[i] {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "Note_Shapes", diagramhierarchyOther.Note_Shapes, diagramhierarchy.Note_Shapes)
		diffs = append(diffs, ops)
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagramhierarchy.NotesWhoseNodeIsExpanded) != len(diagramhierarchyOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.NotesWhoseNodeIsExpanded {
			if (diagramhierarchy.NotesWhoseNodeIsExpanded[i] == nil) != (diagramhierarchyOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.NotesWhoseNodeIsExpanded[i] != nil && diagramhierarchyOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.NotesWhoseNodeIsExpanded[i] != diagramhierarchyOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "NotesWhoseNodeIsExpanded", diagramhierarchyOther.NotesWhoseNodeIsExpanded, diagramhierarchy.NotesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.IsNotesNodeExpanded != diagramhierarchyOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NoteProductShapesDifferent := false
	if len(diagramhierarchy.NoteProductShapes) != len(diagramhierarchyOther.NoteProductShapes) {
		NoteProductShapesDifferent = true
	} else {
		for i := range diagramhierarchy.NoteProductShapes {
			if (diagramhierarchy.NoteProductShapes[i] == nil) != (diagramhierarchyOther.NoteProductShapes[i] == nil) {
				NoteProductShapesDifferent = true
				break
			} else if diagramhierarchy.NoteProductShapes[i] != nil && diagramhierarchyOther.NoteProductShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.NoteProductShapes[i] != diagramhierarchyOther.NoteProductShapes[i] {
					NoteProductShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteProductShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "NoteProductShapes", diagramhierarchyOther.NoteProductShapes, diagramhierarchy.NoteProductShapes)
		diffs = append(diffs, ops)
	}
	NoteTaskShapesDifferent := false
	if len(diagramhierarchy.NoteTaskShapes) != len(diagramhierarchyOther.NoteTaskShapes) {
		NoteTaskShapesDifferent = true
	} else {
		for i := range diagramhierarchy.NoteTaskShapes {
			if (diagramhierarchy.NoteTaskShapes[i] == nil) != (diagramhierarchyOther.NoteTaskShapes[i] == nil) {
				NoteTaskShapesDifferent = true
				break
			} else if diagramhierarchy.NoteTaskShapes[i] != nil && diagramhierarchyOther.NoteTaskShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.NoteTaskShapes[i] != diagramhierarchyOther.NoteTaskShapes[i] {
					NoteTaskShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteTaskShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "NoteTaskShapes", diagramhierarchyOther.NoteTaskShapes, diagramhierarchy.NoteTaskShapes)
		diffs = append(diffs, ops)
	}
	NoteResourceShapesDifferent := false
	if len(diagramhierarchy.NoteResourceShapes) != len(diagramhierarchyOther.NoteResourceShapes) {
		NoteResourceShapesDifferent = true
	} else {
		for i := range diagramhierarchy.NoteResourceShapes {
			if (diagramhierarchy.NoteResourceShapes[i] == nil) != (diagramhierarchyOther.NoteResourceShapes[i] == nil) {
				NoteResourceShapesDifferent = true
				break
			} else if diagramhierarchy.NoteResourceShapes[i] != nil && diagramhierarchyOther.NoteResourceShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.NoteResourceShapes[i] != diagramhierarchyOther.NoteResourceShapes[i] {
					NoteResourceShapesDifferent = true
					break
				}
			}
		}
	}
	if NoteResourceShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "NoteResourceShapes", diagramhierarchyOther.NoteResourceShapes, diagramhierarchy.NoteResourceShapes)
		diffs = append(diffs, ops)
	}
	Resource_ShapesDifferent := false
	if len(diagramhierarchy.Resource_Shapes) != len(diagramhierarchyOther.Resource_Shapes) {
		Resource_ShapesDifferent = true
	} else {
		for i := range diagramhierarchy.Resource_Shapes {
			if (diagramhierarchy.Resource_Shapes[i] == nil) != (diagramhierarchyOther.Resource_Shapes[i] == nil) {
				Resource_ShapesDifferent = true
				break
			} else if diagramhierarchy.Resource_Shapes[i] != nil && diagramhierarchyOther.Resource_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.Resource_Shapes[i] != diagramhierarchyOther.Resource_Shapes[i] {
					Resource_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Resource_ShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "Resource_Shapes", diagramhierarchyOther.Resource_Shapes, diagramhierarchy.Resource_Shapes)
		diffs = append(diffs, ops)
	}
	ResourcesWhoseNodeIsExpandedDifferent := false
	if len(diagramhierarchy.ResourcesWhoseNodeIsExpanded) != len(diagramhierarchyOther.ResourcesWhoseNodeIsExpanded) {
		ResourcesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramhierarchy.ResourcesWhoseNodeIsExpanded {
			if (diagramhierarchy.ResourcesWhoseNodeIsExpanded[i] == nil) != (diagramhierarchyOther.ResourcesWhoseNodeIsExpanded[i] == nil) {
				ResourcesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramhierarchy.ResourcesWhoseNodeIsExpanded[i] != nil && diagramhierarchyOther.ResourcesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.ResourcesWhoseNodeIsExpanded[i] != diagramhierarchyOther.ResourcesWhoseNodeIsExpanded[i] {
					ResourcesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ResourcesWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "ResourcesWhoseNodeIsExpanded", diagramhierarchyOther.ResourcesWhoseNodeIsExpanded, diagramhierarchy.ResourcesWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.IsResourcesNodeExpanded != diagramhierarchyOther.IsResourcesNodeExpanded {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	ResourceComposition_ShapesDifferent := false
	if len(diagramhierarchy.ResourceComposition_Shapes) != len(diagramhierarchyOther.ResourceComposition_Shapes) {
		ResourceComposition_ShapesDifferent = true
	} else {
		for i := range diagramhierarchy.ResourceComposition_Shapes {
			if (diagramhierarchy.ResourceComposition_Shapes[i] == nil) != (diagramhierarchyOther.ResourceComposition_Shapes[i] == nil) {
				ResourceComposition_ShapesDifferent = true
				break
			} else if diagramhierarchy.ResourceComposition_Shapes[i] != nil && diagramhierarchyOther.ResourceComposition_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.ResourceComposition_Shapes[i] != diagramhierarchyOther.ResourceComposition_Shapes[i] {
					ResourceComposition_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ResourceComposition_ShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "ResourceComposition_Shapes", diagramhierarchyOther.ResourceComposition_Shapes, diagramhierarchy.ResourceComposition_Shapes)
		diffs = append(diffs, ops)
	}
	ResourceTaskShapesDifferent := false
	if len(diagramhierarchy.ResourceTaskShapes) != len(diagramhierarchyOther.ResourceTaskShapes) {
		ResourceTaskShapesDifferent = true
	} else {
		for i := range diagramhierarchy.ResourceTaskShapes {
			if (diagramhierarchy.ResourceTaskShapes[i] == nil) != (diagramhierarchyOther.ResourceTaskShapes[i] == nil) {
				ResourceTaskShapesDifferent = true
				break
			} else if diagramhierarchy.ResourceTaskShapes[i] != nil && diagramhierarchyOther.ResourceTaskShapes[i] != nil {
				// this is a pointer comparaison
				if diagramhierarchy.ResourceTaskShapes[i] != diagramhierarchyOther.ResourceTaskShapes[i] {
					ResourceTaskShapesDifferent = true
					break
				}
			}
		}
	}
	if ResourceTaskShapesDifferent {
		ops := Diff(stage, diagramhierarchy, diagramhierarchyOther, "ResourceTaskShapes", diagramhierarchyOther.ResourceTaskShapes, diagramhierarchy.ResourceTaskShapes)
		diffs = append(diffs, ops)
	}
	if diagramhierarchy.IsTimeDiagram != diagramhierarchyOther.IsTimeDiagram {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "IsTimeDiagram"))
	}
	if diagramhierarchy.ComputedStart != diagramhierarchyOther.ComputedStart {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ComputedStart"))
	}
	if diagramhierarchy.ComputedEnd != diagramhierarchyOther.ComputedEnd {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ComputedEnd"))
	}
	if diagramhierarchy.ComputedDuration != diagramhierarchyOther.ComputedDuration {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ComputedDuration"))
	}
	if diagramhierarchy.UseManualStartAndEndDates != diagramhierarchyOther.UseManualStartAndEndDates {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "UseManualStartAndEndDates"))
	}
	if diagramhierarchy.ManualStart != diagramhierarchyOther.ManualStart {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ManualStart"))
	}
	if diagramhierarchy.ManualEnd != diagramhierarchyOther.ManualEnd {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ManualEnd"))
	}
	if diagramhierarchy.TimeStep != diagramhierarchyOther.TimeStep {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "TimeStep"))
	}
	if diagramhierarchy.TimeStepScale != diagramhierarchyOther.TimeStepScale {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "TimeStepScale"))
	}
	if diagramhierarchy.LaneHeight != diagramhierarchyOther.LaneHeight {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "LaneHeight"))
	}
	if diagramhierarchy.RatioBarToLaneHeight != diagramhierarchyOther.RatioBarToLaneHeight {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "RatioBarToLaneHeight"))
	}
	if diagramhierarchy.YTopMargin != diagramhierarchyOther.YTopMargin {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "YTopMargin"))
	}
	if diagramhierarchy.XLeftText != diagramhierarchyOther.XLeftText {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "XLeftText"))
	}
	if diagramhierarchy.TextHeight != diagramhierarchyOther.TextHeight {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "TextHeight"))
	}
	if diagramhierarchy.XLeftLanes != diagramhierarchyOther.XLeftLanes {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "XLeftLanes"))
	}
	if diagramhierarchy.XRightMargin != diagramhierarchyOther.XRightMargin {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "XRightMargin"))
	}
	if diagramhierarchy.ArrowLengthToTheRightOfStartBar != diagramhierarchyOther.ArrowLengthToTheRightOfStartBar {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ArrowLengthToTheRightOfStartBar"))
	}
	if diagramhierarchy.ArrowTipLenght != diagramhierarchyOther.ArrowTipLenght {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "ArrowTipLenght"))
	}
	if diagramhierarchy.TimeLine_Color != diagramhierarchyOther.TimeLine_Color {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "TimeLine_Color"))
	}
	if diagramhierarchy.TimeLine_FillOpacity != diagramhierarchyOther.TimeLine_FillOpacity {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "TimeLine_FillOpacity"))
	}
	if diagramhierarchy.TimeLine_Stroke != diagramhierarchyOther.TimeLine_Stroke {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "TimeLine_Stroke"))
	}
	if diagramhierarchy.TimeLine_StrokeWidth != diagramhierarchyOther.TimeLine_StrokeWidth {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "TimeLine_StrokeWidth"))
	}
	if diagramhierarchy.DrawVerticalTimeLines != diagramhierarchyOther.DrawVerticalTimeLines {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "DrawVerticalTimeLines"))
	}
	if diagramhierarchy.Group_Stroke != diagramhierarchyOther.Group_Stroke {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "Group_Stroke"))
	}
	if diagramhierarchy.Group_StrokeWidth != diagramhierarchyOther.Group_StrokeWidth {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "Group_StrokeWidth"))
	}
	if diagramhierarchy.Group_StrokeDashArray != diagramhierarchyOther.Group_StrokeDashArray {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "Group_StrokeDashArray"))
	}
	if diagramhierarchy.DateYOffset != diagramhierarchyOther.DateYOffset {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "DateYOffset"))
	}
	if diagramhierarchy.AlignOnStartEndOnYearStart != diagramhierarchyOther.AlignOnStartEndOnYearStart {
		diffs = append(diffs, diagramhierarchy.GongMarshallField(stage, "AlignOnStartEndOnYearStart"))
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
	RootMilestonesDifferent := false
	if len(library.RootMilestones) != len(libraryOther.RootMilestones) {
		RootMilestonesDifferent = true
	} else {
		for i := range library.RootMilestones {
			if (library.RootMilestones[i] == nil) != (libraryOther.RootMilestones[i] == nil) {
				RootMilestonesDifferent = true
				break
			} else if library.RootMilestones[i] != nil && libraryOther.RootMilestones[i] != nil {
				// this is a pointer comparaison
				if library.RootMilestones[i] != libraryOther.RootMilestones[i] {
					RootMilestonesDifferent = true
					break
				}
			}
		}
	}
	if RootMilestonesDifferent {
		ops := Diff(stage, library, libraryOther, "RootMilestones", libraryOther.RootMilestones, library.RootMilestones)
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
func (milestone *Milestone) GongDiff(stage *Stage, milestoneOther *Milestone) (diffs []string) {
	// insertion point for field diffs
	if milestone.Name != milestoneOther.Name {
		diffs = append(diffs, milestone.GongMarshallField(stage, "Name"))
	}
	if milestone.ComputedPrefix != milestoneOther.ComputedPrefix {
		diffs = append(diffs, milestone.GongMarshallField(stage, "ComputedPrefix"))
	}
	if milestone.IsExpanded != milestoneOther.IsExpanded {
		diffs = append(diffs, milestone.GongMarshallField(stage, "IsExpanded"))
	}
	if milestone.Date != milestoneOther.Date {
		diffs = append(diffs, milestone.GongMarshallField(stage, "Date"))
	}
	if milestone.DisplayVerticalBar != milestoneOther.DisplayVerticalBar {
		diffs = append(diffs, milestone.GongMarshallField(stage, "DisplayVerticalBar"))
	}
	TaskGroupsToDisplayDifferent := false
	if len(milestone.TaskGroupsToDisplay) != len(milestoneOther.TaskGroupsToDisplay) {
		TaskGroupsToDisplayDifferent = true
	} else {
		for i := range milestone.TaskGroupsToDisplay {
			if (milestone.TaskGroupsToDisplay[i] == nil) != (milestoneOther.TaskGroupsToDisplay[i] == nil) {
				TaskGroupsToDisplayDifferent = true
				break
			} else if milestone.TaskGroupsToDisplay[i] != nil && milestoneOther.TaskGroupsToDisplay[i] != nil {
				// this is a pointer comparaison
				if milestone.TaskGroupsToDisplay[i] != milestoneOther.TaskGroupsToDisplay[i] {
					TaskGroupsToDisplayDifferent = true
					break
				}
			}
		}
	}
	if TaskGroupsToDisplayDifferent {
		ops := Diff(stage, milestone, milestoneOther, "TaskGroupsToDisplay", milestoneOther.TaskGroupsToDisplay, milestone.TaskGroupsToDisplay)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (milestoneshape *MilestoneShape) GongDiff(stage *Stage, milestoneshapeOther *MilestoneShape) (diffs []string) {
	// insertion point for field diffs
	if milestoneshape.Name != milestoneshapeOther.Name {
		diffs = append(diffs, milestoneshape.GongMarshallField(stage, "Name"))
	}
	if (milestoneshape.Milestone == nil) != (milestoneshapeOther.Milestone == nil) {
		diffs = append(diffs, milestoneshape.GongMarshallField(stage, "Milestone"))
	} else if milestoneshape.Milestone != nil && milestoneshapeOther.Milestone != nil {
		if milestoneshape.Milestone != milestoneshapeOther.Milestone {
			diffs = append(diffs, milestoneshape.GongMarshallField(stage, "Milestone"))
		}
	}
	if milestoneshape.X != milestoneshapeOther.X {
		diffs = append(diffs, milestoneshape.GongMarshallField(stage, "X"))
	}
	if milestoneshape.Y != milestoneshapeOther.Y {
		diffs = append(diffs, milestoneshape.GongMarshallField(stage, "Y"))
	}
	if milestoneshape.Width != milestoneshapeOther.Width {
		diffs = append(diffs, milestoneshape.GongMarshallField(stage, "Width"))
	}
	if milestoneshape.Height != milestoneshapeOther.Height {
		diffs = append(diffs, milestoneshape.GongMarshallField(stage, "Height"))
	}
	if milestoneshape.IsHidden != milestoneshapeOther.IsHidden {
		diffs = append(diffs, milestoneshape.GongMarshallField(stage, "IsHidden"))
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
	if product.ComputedPrefix != productOther.ComputedPrefix {
		diffs = append(diffs, product.GongMarshallField(stage, "ComputedPrefix"))
	}
	if product.IsExpanded != productOther.IsExpanded {
		diffs = append(diffs, product.GongMarshallField(stage, "IsExpanded"))
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
	if resource.ComputedPrefix != resourceOther.ComputedPrefix {
		diffs = append(diffs, resource.GongMarshallField(stage, "ComputedPrefix"))
	}
	if resource.IsExpanded != resourceOther.IsExpanded {
		diffs = append(diffs, resource.GongMarshallField(stage, "IsExpanded"))
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
	if task.ComputedPrefix != taskOther.ComputedPrefix {
		diffs = append(diffs, task.GongMarshallField(stage, "ComputedPrefix"))
	}
	if task.IsExpanded != taskOther.IsExpanded {
		diffs = append(diffs, task.GongMarshallField(stage, "IsExpanded"))
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
	if task.Start != taskOther.Start {
		diffs = append(diffs, task.GongMarshallField(stage, "Start"))
	}
	if task.End != taskOther.End {
		diffs = append(diffs, task.GongMarshallField(stage, "End"))
	}
	if task.Description != taskOther.Description {
		diffs = append(diffs, task.GongMarshallField(stage, "Description"))
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
