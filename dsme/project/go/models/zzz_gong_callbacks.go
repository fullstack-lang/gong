// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Product:
		if stage.OnAfterProductCreateCallback != nil {
			stage.OnAfterProductCreateCallback.OnAfterCreate(stage, target)
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
	case *Product:
		newTarget := any(new).(*Product)
		if stage.OnAfterProductUpdateCallback != nil {
			stage.OnAfterProductUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Product:
		if stage.OnAfterProductDeleteCallback != nil {
			staged := any(staged).(*Product)
			stage.OnAfterProductDeleteCallback.OnAfterDelete(stage, staged, front)
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
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Product:
		if stage.OnAfterProductReadCallback != nil {
			stage.OnAfterProductReadCallback.OnAfterRead(stage, target)
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
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Product:
		stage.OnAfterProductUpdateCallback = any(callback).(OnAfterUpdateInterface[Product])
	
	case *Project:
		stage.OnAfterProjectUpdateCallback = any(callback).(OnAfterUpdateInterface[Project])
	
	case *Root:
		stage.OnAfterRootUpdateCallback = any(callback).(OnAfterUpdateInterface[Root])
	
	case *Task:
		stage.OnAfterTaskUpdateCallback = any(callback).(OnAfterUpdateInterface[Task])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Product:
		stage.OnAfterProductCreateCallback = any(callback).(OnAfterCreateInterface[Product])
	
	case *Project:
		stage.OnAfterProjectCreateCallback = any(callback).(OnAfterCreateInterface[Project])
	
	case *Root:
		stage.OnAfterRootCreateCallback = any(callback).(OnAfterCreateInterface[Root])
	
	case *Task:
		stage.OnAfterTaskCreateCallback = any(callback).(OnAfterCreateInterface[Task])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Product:
		stage.OnAfterProductDeleteCallback = any(callback).(OnAfterDeleteInterface[Product])
	
	case *Project:
		stage.OnAfterProjectDeleteCallback = any(callback).(OnAfterDeleteInterface[Project])
	
	case *Root:
		stage.OnAfterRootDeleteCallback = any(callback).(OnAfterDeleteInterface[Root])
	
	case *Task:
		stage.OnAfterTaskDeleteCallback = any(callback).(OnAfterDeleteInterface[Task])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Product:
		stage.OnAfterProductReadCallback = any(callback).(OnAfterReadInterface[Product])
	
	case *Project:
		stage.OnAfterProjectReadCallback = any(callback).(OnAfterReadInterface[Project])
	
	case *Root:
		stage.OnAfterRootReadCallback = any(callback).(OnAfterReadInterface[Root])
	
	case *Task:
		stage.OnAfterTaskReadCallback = any(callback).(OnAfterReadInterface[Task])
	
	}
}
