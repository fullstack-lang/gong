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
	
	case *Product:
		stage.OnAfterProductUpdateCallback = any(callback).(OnAfterUpdateInterface[Product])
	
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProductCompositionShape])
	
	case *ProductShape:
		stage.OnAfterProductShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectUpdateCallback = any(callback).(OnAfterUpdateInterface[Project])
	
	case *Root:
		stage.OnAfterRootUpdateCallback = any(callback).(OnAfterUpdateInterface[Root])
	
	case *Task:
		stage.OnAfterTaskUpdateCallback = any(callback).(OnAfterUpdateInterface[Task])
	
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskCompositionShape])
	
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskInputShape])
	
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
	
	case *Product:
		stage.OnAfterProductCreateCallback = any(callback).(OnAfterCreateInterface[Product])
	
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProductCompositionShape])
	
	case *ProductShape:
		stage.OnAfterProductShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectCreateCallback = any(callback).(OnAfterCreateInterface[Project])
	
	case *Root:
		stage.OnAfterRootCreateCallback = any(callback).(OnAfterCreateInterface[Root])
	
	case *Task:
		stage.OnAfterTaskCreateCallback = any(callback).(OnAfterCreateInterface[Task])
	
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskCompositionShape])
	
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskInputShape])
	
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
	
	case *Product:
		stage.OnAfterProductDeleteCallback = any(callback).(OnAfterDeleteInterface[Product])
	
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProductCompositionShape])
	
	case *ProductShape:
		stage.OnAfterProductShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectDeleteCallback = any(callback).(OnAfterDeleteInterface[Project])
	
	case *Root:
		stage.OnAfterRootDeleteCallback = any(callback).(OnAfterDeleteInterface[Root])
	
	case *Task:
		stage.OnAfterTaskDeleteCallback = any(callback).(OnAfterDeleteInterface[Task])
	
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskCompositionShape])
	
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskInputShape])
	
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
	
	case *Product:
		stage.OnAfterProductReadCallback = any(callback).(OnAfterReadInterface[Product])
	
	case *ProductCompositionShape:
		stage.OnAfterProductCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[ProductCompositionShape])
	
	case *ProductShape:
		stage.OnAfterProductShapeReadCallback = any(callback).(OnAfterReadInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectReadCallback = any(callback).(OnAfterReadInterface[Project])
	
	case *Root:
		stage.OnAfterRootReadCallback = any(callback).(OnAfterReadInterface[Root])
	
	case *Task:
		stage.OnAfterTaskReadCallback = any(callback).(OnAfterReadInterface[Task])
	
	case *TaskCompositionShape:
		stage.OnAfterTaskCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[TaskCompositionShape])
	
	case *TaskInputShape:
		stage.OnAfterTaskInputShapeReadCallback = any(callback).(OnAfterReadInterface[TaskInputShape])
	
	case *TaskShape:
		stage.OnAfterTaskShapeReadCallback = any(callback).(OnAfterReadInterface[TaskShape])
	
	}
}
