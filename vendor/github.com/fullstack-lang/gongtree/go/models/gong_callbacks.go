package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonCreateCallback != nil {
			stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeCreateCallback != nil {
			stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeCreateCallback != nil {
			stage.OnAfterTreeCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Button:
		newTarget := any(new).(*Button)
		if stage.OnAfterButtonUpdateCallback != nil {
			stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Node:
		newTarget := any(new).(*Node)
		if stage.OnAfterNodeUpdateCallback != nil {
			stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tree:
		newTarget := any(new).(*Tree)
		if stage.OnAfterTreeUpdateCallback != nil {
			stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonDeleteCallback != nil {
			staged := any(staged).(*Button)
			stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Node:
		if stage.OnAfterNodeDeleteCallback != nil {
			staged := any(staged).(*Node)
			stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tree:
		if stage.OnAfterTreeDeleteCallback != nil {
			staged := any(staged).(*Tree)
			stage.OnAfterTreeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonReadCallback != nil {
			stage.OnAfterButtonReadCallback.OnAfterRead(stage, target)
		}
	case *Node:
		if stage.OnAfterNodeReadCallback != nil {
			stage.OnAfterNodeReadCallback.OnAfterRead(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeReadCallback != nil {
			stage.OnAfterTreeReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonUpdateCallback = any(callback).(OnAfterUpdateInterface[Button])
	
	case *Node:
		stage.OnAfterNodeUpdateCallback = any(callback).(OnAfterUpdateInterface[Node])
	
	case *Tree:
		stage.OnAfterTreeUpdateCallback = any(callback).(OnAfterUpdateInterface[Tree])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonCreateCallback = any(callback).(OnAfterCreateInterface[Button])
	
	case *Node:
		stage.OnAfterNodeCreateCallback = any(callback).(OnAfterCreateInterface[Node])
	
	case *Tree:
		stage.OnAfterTreeCreateCallback = any(callback).(OnAfterCreateInterface[Tree])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonDeleteCallback = any(callback).(OnAfterDeleteInterface[Button])
	
	case *Node:
		stage.OnAfterNodeDeleteCallback = any(callback).(OnAfterDeleteInterface[Node])
	
	case *Tree:
		stage.OnAfterTreeDeleteCallback = any(callback).(OnAfterDeleteInterface[Tree])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Button:
		stage.OnAfterButtonReadCallback = any(callback).(OnAfterReadInterface[Button])
	
	case *Node:
		stage.OnAfterNodeReadCallback = any(callback).(OnAfterReadInterface[Node])
	
	case *Tree:
		stage.OnAfterTreeReadCallback = any(callback).(OnAfterReadInterface[Tree])
	
	}
}
