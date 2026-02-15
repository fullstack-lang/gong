// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteProductShape:
		if stage.OnAfterNoteProductShapeCreateCallback != nil {
			stage.OnAfterNoteProductShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeCreateCallback != nil {
			stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeCreateCallback != nil {
			stage.OnAfterNoteTaskShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Product:
		if stage.OnAfterProductCreateCallback != nil {
			stage.OnAfterProductCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProductCompositionShape:
		if stage.OnAfterProductCompositionShapeCreateCallback != nil {
			stage.OnAfterProductCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ProductShape:
		if stage.OnAfterProductShapeCreateCallback != nil {
			stage.OnAfterProductShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Project:
		if stage.OnAfterProjectCreateCallback != nil {
			stage.OnAfterProjectCreateCallback.OnAfterCreate(stage, target)
		}
	case *Resource:
		if stage.OnAfterResourceCreateCallback != nil {
			stage.OnAfterResourceCreateCallback.OnAfterCreate(stage, target)
		}
	case *ResourceCompositionShape:
		if stage.OnAfterResourceCompositionShapeCreateCallback != nil {
			stage.OnAfterResourceCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ResourceShape:
		if stage.OnAfterResourceShapeCreateCallback != nil {
			stage.OnAfterResourceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ResourceTaskShape:
		if stage.OnAfterResourceTaskShapeCreateCallback != nil {
			stage.OnAfterResourceTaskShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Root:
		if stage.OnAfterRootCreateCallback != nil {
			stage.OnAfterRootCreateCallback.OnAfterCreate(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskCreateCallback != nil {
			stage.OnAfterTaskCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskCompositionShape:
		if stage.OnAfterTaskCompositionShapeCreateCallback != nil {
			stage.OnAfterTaskCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskInputShape:
		if stage.OnAfterTaskInputShapeCreateCallback != nil {
			stage.OnAfterTaskInputShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskOutputShape:
		if stage.OnAfterTaskOutputShapeCreateCallback != nil {
			stage.OnAfterTaskOutputShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeCreateCallback != nil {
			stage.OnAfterTaskShapeCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note:
		newTarget := any(new).(*Note)
		if stage.OnAfterNoteUpdateCallback != nil {
			stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteProductShape:
		newTarget := any(new).(*NoteProductShape)
		if stage.OnAfterNoteProductShapeUpdateCallback != nil {
			stage.OnAfterNoteProductShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteShape:
		newTarget := any(new).(*NoteShape)
		if stage.OnAfterNoteShapeUpdateCallback != nil {
			stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NoteTaskShape:
		newTarget := any(new).(*NoteTaskShape)
		if stage.OnAfterNoteTaskShapeUpdateCallback != nil {
			stage.OnAfterNoteTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Product:
		newTarget := any(new).(*Product)
		if stage.OnAfterProductUpdateCallback != nil {
			stage.OnAfterProductUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ProductCompositionShape:
		newTarget := any(new).(*ProductCompositionShape)
		if stage.OnAfterProductCompositionShapeUpdateCallback != nil {
			stage.OnAfterProductCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ProductShape:
		newTarget := any(new).(*ProductShape)
		if stage.OnAfterProductShapeUpdateCallback != nil {
			stage.OnAfterProductShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Project:
		newTarget := any(new).(*Project)
		if stage.OnAfterProjectUpdateCallback != nil {
			stage.OnAfterProjectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Resource:
		newTarget := any(new).(*Resource)
		if stage.OnAfterResourceUpdateCallback != nil {
			stage.OnAfterResourceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ResourceCompositionShape:
		newTarget := any(new).(*ResourceCompositionShape)
		if stage.OnAfterResourceCompositionShapeUpdateCallback != nil {
			stage.OnAfterResourceCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ResourceShape:
		newTarget := any(new).(*ResourceShape)
		if stage.OnAfterResourceShapeUpdateCallback != nil {
			stage.OnAfterResourceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ResourceTaskShape:
		newTarget := any(new).(*ResourceTaskShape)
		if stage.OnAfterResourceTaskShapeUpdateCallback != nil {
			stage.OnAfterResourceTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Root:
		newTarget := any(new).(*Root)
		if stage.OnAfterRootUpdateCallback != nil {
			stage.OnAfterRootUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Task:
		newTarget := any(new).(*Task)
		if stage.OnAfterTaskUpdateCallback != nil {
			stage.OnAfterTaskUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskCompositionShape:
		newTarget := any(new).(*TaskCompositionShape)
		if stage.OnAfterTaskCompositionShapeUpdateCallback != nil {
			stage.OnAfterTaskCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskInputShape:
		newTarget := any(new).(*TaskInputShape)
		if stage.OnAfterTaskInputShapeUpdateCallback != nil {
			stage.OnAfterTaskInputShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskOutputShape:
		newTarget := any(new).(*TaskOutputShape)
		if stage.OnAfterTaskOutputShapeUpdateCallback != nil {
			stage.OnAfterTaskOutputShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *TaskShape:
		newTarget := any(new).(*TaskShape)
		if stage.OnAfterTaskShapeUpdateCallback != nil {
			stage.OnAfterTaskShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note:
		if stage.OnAfterNoteDeleteCallback != nil {
			staged := any(staged).(*Note)
			stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteProductShape:
		if stage.OnAfterNoteProductShapeDeleteCallback != nil {
			staged := any(staged).(*NoteProductShape)
			stage.OnAfterNoteProductShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeDeleteCallback != nil {
			staged := any(staged).(*NoteShape)
			stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeDeleteCallback != nil {
			staged := any(staged).(*NoteTaskShape)
			stage.OnAfterNoteTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Product:
		if stage.OnAfterProductDeleteCallback != nil {
			staged := any(staged).(*Product)
			stage.OnAfterProductDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ProductCompositionShape:
		if stage.OnAfterProductCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*ProductCompositionShape)
			stage.OnAfterProductCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ProductShape:
		if stage.OnAfterProductShapeDeleteCallback != nil {
			staged := any(staged).(*ProductShape)
			stage.OnAfterProductShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Project:
		if stage.OnAfterProjectDeleteCallback != nil {
			staged := any(staged).(*Project)
			stage.OnAfterProjectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Resource:
		if stage.OnAfterResourceDeleteCallback != nil {
			staged := any(staged).(*Resource)
			stage.OnAfterResourceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ResourceCompositionShape:
		if stage.OnAfterResourceCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*ResourceCompositionShape)
			stage.OnAfterResourceCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ResourceShape:
		if stage.OnAfterResourceShapeDeleteCallback != nil {
			staged := any(staged).(*ResourceShape)
			stage.OnAfterResourceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ResourceTaskShape:
		if stage.OnAfterResourceTaskShapeDeleteCallback != nil {
			staged := any(staged).(*ResourceTaskShape)
			stage.OnAfterResourceTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Root:
		if stage.OnAfterRootDeleteCallback != nil {
			staged := any(staged).(*Root)
			stage.OnAfterRootDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Task:
		if stage.OnAfterTaskDeleteCallback != nil {
			staged := any(staged).(*Task)
			stage.OnAfterTaskDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskCompositionShape:
		if stage.OnAfterTaskCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*TaskCompositionShape)
			stage.OnAfterTaskCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskInputShape:
		if stage.OnAfterTaskInputShapeDeleteCallback != nil {
			staged := any(staged).(*TaskInputShape)
			stage.OnAfterTaskInputShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskOutputShape:
		if stage.OnAfterTaskOutputShapeDeleteCallback != nil {
			staged := any(staged).(*TaskOutputShape)
			stage.OnAfterTaskOutputShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeDeleteCallback != nil {
			staged := any(staged).(*TaskShape)
			stage.OnAfterTaskShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Diagram:
		if stage.OnAfterDiagramReadCallback != nil {
			stage.OnAfterDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *NoteProductShape:
		if stage.OnAfterNoteProductShapeReadCallback != nil {
			stage.OnAfterNoteProductShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteShape:
		if stage.OnAfterNoteShapeReadCallback != nil {
			stage.OnAfterNoteShapeReadCallback.OnAfterRead(stage, target)
		}
	case *NoteTaskShape:
		if stage.OnAfterNoteTaskShapeReadCallback != nil {
			stage.OnAfterNoteTaskShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Product:
		if stage.OnAfterProductReadCallback != nil {
			stage.OnAfterProductReadCallback.OnAfterRead(stage, target)
		}
	case *ProductCompositionShape:
		if stage.OnAfterProductCompositionShapeReadCallback != nil {
			stage.OnAfterProductCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ProductShape:
		if stage.OnAfterProductShapeReadCallback != nil {
			stage.OnAfterProductShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Project:
		if stage.OnAfterProjectReadCallback != nil {
			stage.OnAfterProjectReadCallback.OnAfterRead(stage, target)
		}
	case *Resource:
		if stage.OnAfterResourceReadCallback != nil {
			stage.OnAfterResourceReadCallback.OnAfterRead(stage, target)
		}
	case *ResourceCompositionShape:
		if stage.OnAfterResourceCompositionShapeReadCallback != nil {
			stage.OnAfterResourceCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ResourceShape:
		if stage.OnAfterResourceShapeReadCallback != nil {
			stage.OnAfterResourceShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ResourceTaskShape:
		if stage.OnAfterResourceTaskShapeReadCallback != nil {
			stage.OnAfterResourceTaskShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Root:
		if stage.OnAfterRootReadCallback != nil {
			stage.OnAfterRootReadCallback.OnAfterRead(stage, target)
		}
	case *Task:
		if stage.OnAfterTaskReadCallback != nil {
			stage.OnAfterTaskReadCallback.OnAfterRead(stage, target)
		}
	case *TaskCompositionShape:
		if stage.OnAfterTaskCompositionShapeReadCallback != nil {
			stage.OnAfterTaskCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TaskInputShape:
		if stage.OnAfterTaskInputShapeReadCallback != nil {
			stage.OnAfterTaskInputShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TaskOutputShape:
		if stage.OnAfterTaskOutputShapeReadCallback != nil {
			stage.OnAfterTaskOutputShapeReadCallback.OnAfterRead(stage, target)
		}
	case *TaskShape:
		if stage.OnAfterTaskShapeReadCallback != nil {
			stage.OnAfterTaskShapeReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Diagram:
		stage.OnAfterDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Diagram])
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NoteTaskShape])
	case *Product:
		stage.OnAfterProductUpdateCallback = any(callback).(OnAfterUpdateInterface[Product])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProductShape])
	case *Project:
		stage.OnAfterProjectUpdateCallback = any(callback).(OnAfterUpdateInterface[Project])
	case *Resource:
		stage.OnAfterResourceUpdateCallback = any(callback).(OnAfterUpdateInterface[Resource])
	case *ResourceCompositionShape:
		stage.OnAfterResourceCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ResourceCompositionShape])
	case *ResourceShape:
		stage.OnAfterResourceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ResourceShape])
	case *ResourceTaskShape:
		stage.OnAfterResourceTaskShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ResourceTaskShape])
	case *Root:
		stage.OnAfterRootUpdateCallback = any(callback).(OnAfterUpdateInterface[Root])
	case *Task:
		stage.OnAfterTaskUpdateCallback = any(callback).(OnAfterUpdateInterface[Task])
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskCompositionShape])
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskInputShape])
	case *TaskOutputShape:
		stage.OnAfterTaskOutputShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskOutputShape])
	case *TaskShape:
		stage.OnAfterTaskShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskShape])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Diagram:
		stage.OnAfterDiagramCreateCallback = any(callback).(OnAfterCreateInterface[Diagram])
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeCreateCallback = any(callback).(OnAfterCreateInterface[NoteTaskShape])
	case *Product:
		stage.OnAfterProductCreateCallback = any(callback).(OnAfterCreateInterface[Product])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProductShape])
	case *Project:
		stage.OnAfterProjectCreateCallback = any(callback).(OnAfterCreateInterface[Project])
	case *Resource:
		stage.OnAfterResourceCreateCallback = any(callback).(OnAfterCreateInterface[Resource])
	case *ResourceCompositionShape:
		stage.OnAfterResourceCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[ResourceCompositionShape])
	case *ResourceShape:
		stage.OnAfterResourceShapeCreateCallback = any(callback).(OnAfterCreateInterface[ResourceShape])
	case *ResourceTaskShape:
		stage.OnAfterResourceTaskShapeCreateCallback = any(callback).(OnAfterCreateInterface[ResourceTaskShape])
	case *Root:
		stage.OnAfterRootCreateCallback = any(callback).(OnAfterCreateInterface[Root])
	case *Task:
		stage.OnAfterTaskCreateCallback = any(callback).(OnAfterCreateInterface[Task])
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskCompositionShape])
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskInputShape])
	case *TaskOutputShape:
		stage.OnAfterTaskOutputShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskOutputShape])
	case *TaskShape:
		stage.OnAfterTaskShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskShape])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Diagram:
		stage.OnAfterDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Diagram])
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NoteTaskShape])
	case *Product:
		stage.OnAfterProductDeleteCallback = any(callback).(OnAfterDeleteInterface[Product])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProductShape])
	case *Project:
		stage.OnAfterProjectDeleteCallback = any(callback).(OnAfterDeleteInterface[Project])
	case *Resource:
		stage.OnAfterResourceDeleteCallback = any(callback).(OnAfterDeleteInterface[Resource])
	case *ResourceCompositionShape:
		stage.OnAfterResourceCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ResourceCompositionShape])
	case *ResourceShape:
		stage.OnAfterResourceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ResourceShape])
	case *ResourceTaskShape:
		stage.OnAfterResourceTaskShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ResourceTaskShape])
	case *Root:
		stage.OnAfterRootDeleteCallback = any(callback).(OnAfterDeleteInterface[Root])
	case *Task:
		stage.OnAfterTaskDeleteCallback = any(callback).(OnAfterDeleteInterface[Task])
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskCompositionShape])
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskInputShape])
	case *TaskOutputShape:
		stage.OnAfterTaskOutputShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskOutputShape])
	case *TaskShape:
		stage.OnAfterTaskShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskShape])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Diagram:
		stage.OnAfterDiagramReadCallback = any(callback).(OnAfterReadInterface[Diagram])
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	case *NoteProductShape:
		stage.OnAfterNoteProductShapeReadCallback = any(callback).(OnAfterReadInterface[NoteProductShape])
	case *NoteShape:
		stage.OnAfterNoteShapeReadCallback = any(callback).(OnAfterReadInterface[NoteShape])
	case *NoteTaskShape:
		stage.OnAfterNoteTaskShapeReadCallback = any(callback).(OnAfterReadInterface[NoteTaskShape])
	case *Product:
		stage.OnAfterProductReadCallback = any(callback).(OnAfterReadInterface[Product])
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[ProductCompositionShape])
	case *ProductShape:
		stage.OnAfterProductShapeReadCallback = any(callback).(OnAfterReadInterface[ProductShape])
	case *Project:
		stage.OnAfterProjectReadCallback = any(callback).(OnAfterReadInterface[Project])
	case *Resource:
		stage.OnAfterResourceReadCallback = any(callback).(OnAfterReadInterface[Resource])
	case *ResourceCompositionShape:
		stage.OnAfterResourceCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[ResourceCompositionShape])
	case *ResourceShape:
		stage.OnAfterResourceShapeReadCallback = any(callback).(OnAfterReadInterface[ResourceShape])
	case *ResourceTaskShape:
		stage.OnAfterResourceTaskShapeReadCallback = any(callback).(OnAfterReadInterface[ResourceTaskShape])
	case *Root:
		stage.OnAfterRootReadCallback = any(callback).(OnAfterReadInterface[Root])
	case *Task:
		stage.OnAfterTaskReadCallback = any(callback).(OnAfterReadInterface[Task])
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[TaskCompositionShape])
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeReadCallback = any(callback).(OnAfterReadInterface[TaskInputShape])
	case *TaskOutputShape:
		stage.OnAfterTaskOutputShapeReadCallback = any(callback).(OnAfterReadInterface[TaskOutputShape])
	case *TaskShape:
		stage.OnAfterTaskShapeReadCallback = any(callback).(OnAfterReadInterface[TaskShape])
	}
}
