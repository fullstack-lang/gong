// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *CompositionShape:
		if stage.OnAfterCompositionShapeCreateCallback != nil {
			stage.OnAfterCompositionShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Product:
		if stage.OnAfterProductCreateCallback != nil {
			stage.OnAfterProductCreateCallback.OnAfterCreate(stage, target)
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
	case *CompositionShape:
		newTarget := any(new).(*CompositionShape)
		if stage.OnAfterCompositionShapeUpdateCallback != nil {
			stage.OnAfterCompositionShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
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
	case *CompositionShape:
		if stage.OnAfterCompositionShapeDeleteCallback != nil {
			staged := any(staged).(*CompositionShape)
			stage.OnAfterCompositionShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
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
	case *CompositionShape:
		if stage.OnAfterCompositionShapeReadCallback != nil {
			stage.OnAfterCompositionShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramReadCallback != nil {
			stage.OnAfterDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Product:
		if stage.OnAfterProductReadCallback != nil {
			stage.OnAfterProductReadCallback.OnAfterRead(stage, target)
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
	case *CompositionShape:
		stage.OnAfterCompositionShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[CompositionShape])
	
	case *Diagram:
		stage.OnAfterDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Diagram])
	
	case *Product:
		stage.OnAfterProductUpdateCallback = any(callback).(OnAfterUpdateInterface[Product])
	
	case *ProductShape:
		stage.OnAfterProductShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectUpdateCallback = any(callback).(OnAfterUpdateInterface[Project])
	
	case *Root:
		stage.OnAfterRootUpdateCallback = any(callback).(OnAfterUpdateInterface[Root])
	
	case *Task:
		stage.OnAfterTaskUpdateCallback = any(callback).(OnAfterUpdateInterface[Task])
	
	case *TaskShape:
		stage.OnAfterTaskShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[TaskShape])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CompositionShape:
		stage.OnAfterCompositionShapeCreateCallback = any(callback).(OnAfterCreateInterface[CompositionShape])
	
	case *Diagram:
		stage.OnAfterDiagramCreateCallback = any(callback).(OnAfterCreateInterface[Diagram])
	
	case *Product:
		stage.OnAfterProductCreateCallback = any(callback).(OnAfterCreateInterface[Product])
	
	case *ProductShape:
		stage.OnAfterProductShapeCreateCallback = any(callback).(OnAfterCreateInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectCreateCallback = any(callback).(OnAfterCreateInterface[Project])
	
	case *Root:
		stage.OnAfterRootCreateCallback = any(callback).(OnAfterCreateInterface[Root])
	
	case *Task:
		stage.OnAfterTaskCreateCallback = any(callback).(OnAfterCreateInterface[Task])
	
	case *TaskShape:
		stage.OnAfterTaskShapeCreateCallback = any(callback).(OnAfterCreateInterface[TaskShape])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CompositionShape:
		stage.OnAfterCompositionShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[CompositionShape])
	
	case *Diagram:
		stage.OnAfterDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Diagram])
	
	case *Product:
		stage.OnAfterProductDeleteCallback = any(callback).(OnAfterDeleteInterface[Product])
	
	case *ProductShape:
		stage.OnAfterProductShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectDeleteCallback = any(callback).(OnAfterDeleteInterface[Project])
	
	case *Root:
		stage.OnAfterRootDeleteCallback = any(callback).(OnAfterDeleteInterface[Root])
	
	case *Task:
		stage.OnAfterTaskDeleteCallback = any(callback).(OnAfterDeleteInterface[Task])
	
	case *TaskShape:
		stage.OnAfterTaskShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[TaskShape])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *CompositionShape:
		stage.OnAfterCompositionShapeReadCallback = any(callback).(OnAfterReadInterface[CompositionShape])
	
	case *Diagram:
		stage.OnAfterDiagramReadCallback = any(callback).(OnAfterReadInterface[Diagram])
	
	case *Product:
		stage.OnAfterProductReadCallback = any(callback).(OnAfterReadInterface[Product])
	
	case *ProductShape:
		stage.OnAfterProductShapeReadCallback = any(callback).(OnAfterReadInterface[ProductShape])
	
	case *Project:
		stage.OnAfterProjectReadCallback = any(callback).(OnAfterReadInterface[Project])
	
	case *Root:
		stage.OnAfterRootReadCallback = any(callback).(OnAfterReadInterface[Root])
	
	case *Task:
		stage.OnAfterTaskReadCallback = any(callback).(OnAfterReadInterface[Task])
	
	case *TaskShape:
		stage.OnAfterTaskShapeReadCallback = any(callback).(OnAfterReadInterface[TaskShape])
	
	}
}
