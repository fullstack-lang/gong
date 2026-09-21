// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (diagram *Diagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramCreateCallback != nil {
		stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, diagram)
	}
}

func (diagram *Diagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramUpdateCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, diagram, frontDiagram)
	}
}

func (diagram *Diagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramDeleteCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, diagram, frontDiagram)
	}
}

func (library *Library) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLibraryCreateCallback != nil {
		stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, library)
	}
}

func (library *Library) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryUpdateCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, library, frontLibrary)
	}
}

func (library *Library) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryDeleteCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, library, frontLibrary)
	}
}

func (note *Note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteCreateCallback != nil {
		stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, note)
	}
}

func (note *Note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteUpdateCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, note, frontNote)
	}
}

func (note *Note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeleteCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, note, frontNote)
	}
}

func (noteproductshape *NoteProductShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteProductShapeCreateCallback != nil {
		stage.OnAfterNoteProductShapeCreateCallback.OnAfterCreate(stage, noteproductshape)
	}
}

func (noteproductshape *NoteProductShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteProductShapeUpdateCallback != nil {
		var frontNoteProductShape *NoteProductShape
		if front != nil {
			frontNoteProductShape, _ = front.(*NoteProductShape)
		}
		stage.OnAfterNoteProductShapeUpdateCallback.OnAfterUpdate(stage, noteproductshape, frontNoteProductShape)
	}
}

func (noteproductshape *NoteProductShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteProductShapeDeleteCallback != nil {
		var frontNoteProductShape *NoteProductShape
		if front != nil {
			frontNoteProductShape, _ = front.(*NoteProductShape)
		}
		stage.OnAfterNoteProductShapeDeleteCallback.OnAfterDelete(stage, noteproductshape, frontNoteProductShape)
	}
}

func (noteresourceshape *NoteResourceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteResourceShapeCreateCallback != nil {
		stage.OnAfterNoteResourceShapeCreateCallback.OnAfterCreate(stage, noteresourceshape)
	}
}

func (noteresourceshape *NoteResourceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteResourceShapeUpdateCallback != nil {
		var frontNoteResourceShape *NoteResourceShape
		if front != nil {
			frontNoteResourceShape, _ = front.(*NoteResourceShape)
		}
		stage.OnAfterNoteResourceShapeUpdateCallback.OnAfterUpdate(stage, noteresourceshape, frontNoteResourceShape)
	}
}

func (noteresourceshape *NoteResourceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteResourceShapeDeleteCallback != nil {
		var frontNoteResourceShape *NoteResourceShape
		if front != nil {
			frontNoteResourceShape, _ = front.(*NoteResourceShape)
		}
		stage.OnAfterNoteResourceShapeDeleteCallback.OnAfterDelete(stage, noteresourceshape, frontNoteResourceShape)
	}
}

func (noteshape *NoteShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteShapeCreateCallback != nil {
		stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, noteshape)
	}
}

func (noteshape *NoteShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeUpdateCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, noteshape, frontNoteShape)
	}
}

func (noteshape *NoteShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeDeleteCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, noteshape, frontNoteShape)
	}
}

func (notetaskshape *NoteTaskShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteTaskShapeCreateCallback != nil {
		stage.OnAfterNoteTaskShapeCreateCallback.OnAfterCreate(stage, notetaskshape)
	}
}

func (notetaskshape *NoteTaskShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTaskShapeUpdateCallback != nil {
		var frontNoteTaskShape *NoteTaskShape
		if front != nil {
			frontNoteTaskShape, _ = front.(*NoteTaskShape)
		}
		stage.OnAfterNoteTaskShapeUpdateCallback.OnAfterUpdate(stage, notetaskshape, frontNoteTaskShape)
	}
}

func (notetaskshape *NoteTaskShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTaskShapeDeleteCallback != nil {
		var frontNoteTaskShape *NoteTaskShape
		if front != nil {
			frontNoteTaskShape, _ = front.(*NoteTaskShape)
		}
		stage.OnAfterNoteTaskShapeDeleteCallback.OnAfterDelete(stage, notetaskshape, frontNoteTaskShape)
	}
}

func (product *Product) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProductCreateCallback != nil {
		stage.OnAfterProductCreateCallback.OnAfterCreate(stage, product)
	}
}

func (product *Product) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductUpdateCallback != nil {
		var frontProduct *Product
		if front != nil {
			frontProduct, _ = front.(*Product)
		}
		stage.OnAfterProductUpdateCallback.OnAfterUpdate(stage, product, frontProduct)
	}
}

func (product *Product) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductDeleteCallback != nil {
		var frontProduct *Product
		if front != nil {
			frontProduct, _ = front.(*Product)
		}
		stage.OnAfterProductDeleteCallback.OnAfterDelete(stage, product, frontProduct)
	}
}

func (productcompositionshape *ProductCompositionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProductCompositionShapeCreateCallback != nil {
		stage.OnAfterProductCompositionShapeCreateCallback.OnAfterCreate(stage, productcompositionshape)
	}
}

func (productcompositionshape *ProductCompositionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductCompositionShapeUpdateCallback != nil {
		var frontProductCompositionShape *ProductCompositionShape
		if front != nil {
			frontProductCompositionShape, _ = front.(*ProductCompositionShape)
		}
		stage.OnAfterProductCompositionShapeUpdateCallback.OnAfterUpdate(stage, productcompositionshape, frontProductCompositionShape)
	}
}

func (productcompositionshape *ProductCompositionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductCompositionShapeDeleteCallback != nil {
		var frontProductCompositionShape *ProductCompositionShape
		if front != nil {
			frontProductCompositionShape, _ = front.(*ProductCompositionShape)
		}
		stage.OnAfterProductCompositionShapeDeleteCallback.OnAfterDelete(stage, productcompositionshape, frontProductCompositionShape)
	}
}

func (productreferenceshape *ProductReferenceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProductReferenceShapeCreateCallback != nil {
		stage.OnAfterProductReferenceShapeCreateCallback.OnAfterCreate(stage, productreferenceshape)
	}
}

func (productreferenceshape *ProductReferenceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductReferenceShapeUpdateCallback != nil {
		var frontProductReferenceShape *ProductReferenceShape
		if front != nil {
			frontProductReferenceShape, _ = front.(*ProductReferenceShape)
		}
		stage.OnAfterProductReferenceShapeUpdateCallback.OnAfterUpdate(stage, productreferenceshape, frontProductReferenceShape)
	}
}

func (productreferenceshape *ProductReferenceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductReferenceShapeDeleteCallback != nil {
		var frontProductReferenceShape *ProductReferenceShape
		if front != nil {
			frontProductReferenceShape, _ = front.(*ProductReferenceShape)
		}
		stage.OnAfterProductReferenceShapeDeleteCallback.OnAfterDelete(stage, productreferenceshape, frontProductReferenceShape)
	}
}

func (productshape *ProductShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterProductShapeCreateCallback != nil {
		stage.OnAfterProductShapeCreateCallback.OnAfterCreate(stage, productshape)
	}
}

func (productshape *ProductShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductShapeUpdateCallback != nil {
		var frontProductShape *ProductShape
		if front != nil {
			frontProductShape, _ = front.(*ProductShape)
		}
		stage.OnAfterProductShapeUpdateCallback.OnAfterUpdate(stage, productshape, frontProductShape)
	}
}

func (productshape *ProductShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterProductShapeDeleteCallback != nil {
		var frontProductShape *ProductShape
		if front != nil {
			frontProductShape, _ = front.(*ProductShape)
		}
		stage.OnAfterProductShapeDeleteCallback.OnAfterDelete(stage, productshape, frontProductShape)
	}
}

func (resource *Resource) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResourceCreateCallback != nil {
		stage.OnAfterResourceCreateCallback.OnAfterCreate(stage, resource)
	}
}

func (resource *Resource) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceUpdateCallback != nil {
		var frontResource *Resource
		if front != nil {
			frontResource, _ = front.(*Resource)
		}
		stage.OnAfterResourceUpdateCallback.OnAfterUpdate(stage, resource, frontResource)
	}
}

func (resource *Resource) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceDeleteCallback != nil {
		var frontResource *Resource
		if front != nil {
			frontResource, _ = front.(*Resource)
		}
		stage.OnAfterResourceDeleteCallback.OnAfterDelete(stage, resource, frontResource)
	}
}

func (resourcecompositionshape *ResourceCompositionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResourceCompositionShapeCreateCallback != nil {
		stage.OnAfterResourceCompositionShapeCreateCallback.OnAfterCreate(stage, resourcecompositionshape)
	}
}

func (resourcecompositionshape *ResourceCompositionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceCompositionShapeUpdateCallback != nil {
		var frontResourceCompositionShape *ResourceCompositionShape
		if front != nil {
			frontResourceCompositionShape, _ = front.(*ResourceCompositionShape)
		}
		stage.OnAfterResourceCompositionShapeUpdateCallback.OnAfterUpdate(stage, resourcecompositionshape, frontResourceCompositionShape)
	}
}

func (resourcecompositionshape *ResourceCompositionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceCompositionShapeDeleteCallback != nil {
		var frontResourceCompositionShape *ResourceCompositionShape
		if front != nil {
			frontResourceCompositionShape, _ = front.(*ResourceCompositionShape)
		}
		stage.OnAfterResourceCompositionShapeDeleteCallback.OnAfterDelete(stage, resourcecompositionshape, frontResourceCompositionShape)
	}
}

func (resourceshape *ResourceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResourceShapeCreateCallback != nil {
		stage.OnAfterResourceShapeCreateCallback.OnAfterCreate(stage, resourceshape)
	}
}

func (resourceshape *ResourceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceShapeUpdateCallback != nil {
		var frontResourceShape *ResourceShape
		if front != nil {
			frontResourceShape, _ = front.(*ResourceShape)
		}
		stage.OnAfterResourceShapeUpdateCallback.OnAfterUpdate(stage, resourceshape, frontResourceShape)
	}
}

func (resourceshape *ResourceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceShapeDeleteCallback != nil {
		var frontResourceShape *ResourceShape
		if front != nil {
			frontResourceShape, _ = front.(*ResourceShape)
		}
		stage.OnAfterResourceShapeDeleteCallback.OnAfterDelete(stage, resourceshape, frontResourceShape)
	}
}

func (resourcetaskshape *ResourceTaskShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResourceTaskShapeCreateCallback != nil {
		stage.OnAfterResourceTaskShapeCreateCallback.OnAfterCreate(stage, resourcetaskshape)
	}
}

func (resourcetaskshape *ResourceTaskShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceTaskShapeUpdateCallback != nil {
		var frontResourceTaskShape *ResourceTaskShape
		if front != nil {
			frontResourceTaskShape, _ = front.(*ResourceTaskShape)
		}
		stage.OnAfterResourceTaskShapeUpdateCallback.OnAfterUpdate(stage, resourcetaskshape, frontResourceTaskShape)
	}
}

func (resourcetaskshape *ResourceTaskShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceTaskShapeDeleteCallback != nil {
		var frontResourceTaskShape *ResourceTaskShape
		if front != nil {
			frontResourceTaskShape, _ = front.(*ResourceTaskShape)
		}
		stage.OnAfterResourceTaskShapeDeleteCallback.OnAfterDelete(stage, resourcetaskshape, frontResourceTaskShape)
	}
}

func (task *Task) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskCreateCallback != nil {
		stage.OnAfterTaskCreateCallback.OnAfterCreate(stage, task)
	}
}

func (task *Task) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskUpdateCallback != nil {
		var frontTask *Task
		if front != nil {
			frontTask, _ = front.(*Task)
		}
		stage.OnAfterTaskUpdateCallback.OnAfterUpdate(stage, task, frontTask)
	}
}

func (task *Task) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskDeleteCallback != nil {
		var frontTask *Task
		if front != nil {
			frontTask, _ = front.(*Task)
		}
		stage.OnAfterTaskDeleteCallback.OnAfterDelete(stage, task, frontTask)
	}
}

func (taskcompositionshape *TaskCompositionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskCompositionShapeCreateCallback != nil {
		stage.OnAfterTaskCompositionShapeCreateCallback.OnAfterCreate(stage, taskcompositionshape)
	}
}

func (taskcompositionshape *TaskCompositionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskCompositionShapeUpdateCallback != nil {
		var frontTaskCompositionShape *TaskCompositionShape
		if front != nil {
			frontTaskCompositionShape, _ = front.(*TaskCompositionShape)
		}
		stage.OnAfterTaskCompositionShapeUpdateCallback.OnAfterUpdate(stage, taskcompositionshape, frontTaskCompositionShape)
	}
}

func (taskcompositionshape *TaskCompositionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskCompositionShapeDeleteCallback != nil {
		var frontTaskCompositionShape *TaskCompositionShape
		if front != nil {
			frontTaskCompositionShape, _ = front.(*TaskCompositionShape)
		}
		stage.OnAfterTaskCompositionShapeDeleteCallback.OnAfterDelete(stage, taskcompositionshape, frontTaskCompositionShape)
	}
}

func (taskgroup *TaskGroup) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskGroupCreateCallback != nil {
		stage.OnAfterTaskGroupCreateCallback.OnAfterCreate(stage, taskgroup)
	}
}

func (taskgroup *TaskGroup) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskGroupUpdateCallback != nil {
		var frontTaskGroup *TaskGroup
		if front != nil {
			frontTaskGroup, _ = front.(*TaskGroup)
		}
		stage.OnAfterTaskGroupUpdateCallback.OnAfterUpdate(stage, taskgroup, frontTaskGroup)
	}
}

func (taskgroup *TaskGroup) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskGroupDeleteCallback != nil {
		var frontTaskGroup *TaskGroup
		if front != nil {
			frontTaskGroup, _ = front.(*TaskGroup)
		}
		stage.OnAfterTaskGroupDeleteCallback.OnAfterDelete(stage, taskgroup, frontTaskGroup)
	}
}

func (taskgroupshape *TaskGroupShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskGroupShapeCreateCallback != nil {
		stage.OnAfterTaskGroupShapeCreateCallback.OnAfterCreate(stage, taskgroupshape)
	}
}

func (taskgroupshape *TaskGroupShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskGroupShapeUpdateCallback != nil {
		var frontTaskGroupShape *TaskGroupShape
		if front != nil {
			frontTaskGroupShape, _ = front.(*TaskGroupShape)
		}
		stage.OnAfterTaskGroupShapeUpdateCallback.OnAfterUpdate(stage, taskgroupshape, frontTaskGroupShape)
	}
}

func (taskgroupshape *TaskGroupShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskGroupShapeDeleteCallback != nil {
		var frontTaskGroupShape *TaskGroupShape
		if front != nil {
			frontTaskGroupShape, _ = front.(*TaskGroupShape)
		}
		stage.OnAfterTaskGroupShapeDeleteCallback.OnAfterDelete(stage, taskgroupshape, frontTaskGroupShape)
	}
}

func (taskinputshape *TaskInputShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskInputShapeCreateCallback != nil {
		stage.OnAfterTaskInputShapeCreateCallback.OnAfterCreate(stage, taskinputshape)
	}
}

func (taskinputshape *TaskInputShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskInputShapeUpdateCallback != nil {
		var frontTaskInputShape *TaskInputShape
		if front != nil {
			frontTaskInputShape, _ = front.(*TaskInputShape)
		}
		stage.OnAfterTaskInputShapeUpdateCallback.OnAfterUpdate(stage, taskinputshape, frontTaskInputShape)
	}
}

func (taskinputshape *TaskInputShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskInputShapeDeleteCallback != nil {
		var frontTaskInputShape *TaskInputShape
		if front != nil {
			frontTaskInputShape, _ = front.(*TaskInputShape)
		}
		stage.OnAfterTaskInputShapeDeleteCallback.OnAfterDelete(stage, taskinputshape, frontTaskInputShape)
	}
}

func (taskoutputshape *TaskOutputShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskOutputShapeCreateCallback != nil {
		stage.OnAfterTaskOutputShapeCreateCallback.OnAfterCreate(stage, taskoutputshape)
	}
}

func (taskoutputshape *TaskOutputShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskOutputShapeUpdateCallback != nil {
		var frontTaskOutputShape *TaskOutputShape
		if front != nil {
			frontTaskOutputShape, _ = front.(*TaskOutputShape)
		}
		stage.OnAfterTaskOutputShapeUpdateCallback.OnAfterUpdate(stage, taskoutputshape, frontTaskOutputShape)
	}
}

func (taskoutputshape *TaskOutputShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskOutputShapeDeleteCallback != nil {
		var frontTaskOutputShape *TaskOutputShape
		if front != nil {
			frontTaskOutputShape, _ = front.(*TaskOutputShape)
		}
		stage.OnAfterTaskOutputShapeDeleteCallback.OnAfterDelete(stage, taskoutputshape, frontTaskOutputShape)
	}
}

func (taskpredecessorshape *TaskPredecessorShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskPredecessorShapeCreateCallback != nil {
		stage.OnAfterTaskPredecessorShapeCreateCallback.OnAfterCreate(stage, taskpredecessorshape)
	}
}

func (taskpredecessorshape *TaskPredecessorShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskPredecessorShapeUpdateCallback != nil {
		var frontTaskPredecessorShape *TaskPredecessorShape
		if front != nil {
			frontTaskPredecessorShape, _ = front.(*TaskPredecessorShape)
		}
		stage.OnAfterTaskPredecessorShapeUpdateCallback.OnAfterUpdate(stage, taskpredecessorshape, frontTaskPredecessorShape)
	}
}

func (taskpredecessorshape *TaskPredecessorShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskPredecessorShapeDeleteCallback != nil {
		var frontTaskPredecessorShape *TaskPredecessorShape
		if front != nil {
			frontTaskPredecessorShape, _ = front.(*TaskPredecessorShape)
		}
		stage.OnAfterTaskPredecessorShapeDeleteCallback.OnAfterDelete(stage, taskpredecessorshape, frontTaskPredecessorShape)
	}
}

func (taskshape *TaskShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTaskShapeCreateCallback != nil {
		stage.OnAfterTaskShapeCreateCallback.OnAfterCreate(stage, taskshape)
	}
}

func (taskshape *TaskShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskShapeUpdateCallback != nil {
		var frontTaskShape *TaskShape
		if front != nil {
			frontTaskShape, _ = front.(*TaskShape)
		}
		stage.OnAfterTaskShapeUpdateCallback.OnAfterUpdate(stage, taskshape, frontTaskShape)
	}
}

func (taskshape *TaskShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTaskShapeDeleteCallback != nil {
		var frontTaskShape *TaskShape
		if front != nil {
			frontTaskShape, _ = front.(*TaskShape)
		}
		stage.OnAfterTaskShapeDeleteCallback.OnAfterDelete(stage, taskshape, frontTaskShape)
	}
}

