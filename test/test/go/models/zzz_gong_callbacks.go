// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Astruct:
		if stage.OnAfterAstructCreateCallback != nil {
			stage.OnAfterAstructCreateCallback.OnAfterCreate(stage, target)
		}
	case *AstructBstruct2Use:
		if stage.OnAfterAstructBstruct2UseCreateCallback != nil {
			stage.OnAfterAstructBstruct2UseCreateCallback.OnAfterCreate(stage, target)
		}
	case *AstructBstructUse:
		if stage.OnAfterAstructBstructUseCreateCallback != nil {
			stage.OnAfterAstructBstructUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bstruct:
		if stage.OnAfterBstructCreateCallback != nil {
			stage.OnAfterBstructCreateCallback.OnAfterCreate(stage, target)
		}
	case *Dstruct:
		if stage.OnAfterDstructCreateCallback != nil {
			stage.OnAfterDstructCreateCallback.OnAfterCreate(stage, target)
		}
	case *F0123456789012345678901234567890:
		if stage.OnAfterF0123456789012345678901234567890CreateCallback != nil {
			stage.OnAfterF0123456789012345678901234567890CreateCallback.OnAfterCreate(stage, target)
		}
	case *Gstruct:
		if stage.OnAfterGstructCreateCallback != nil {
			stage.OnAfterGstructCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Astruct:
		newTarget := any(new).(*Astruct)
		if stage.OnAfterAstructUpdateCallback != nil {
			stage.OnAfterAstructUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AstructBstruct2Use:
		newTarget := any(new).(*AstructBstruct2Use)
		if stage.OnAfterAstructBstruct2UseUpdateCallback != nil {
			stage.OnAfterAstructBstruct2UseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AstructBstructUse:
		newTarget := any(new).(*AstructBstructUse)
		if stage.OnAfterAstructBstructUseUpdateCallback != nil {
			stage.OnAfterAstructBstructUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bstruct:
		newTarget := any(new).(*Bstruct)
		if stage.OnAfterBstructUpdateCallback != nil {
			stage.OnAfterBstructUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Dstruct:
		newTarget := any(new).(*Dstruct)
		if stage.OnAfterDstructUpdateCallback != nil {
			stage.OnAfterDstructUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *F0123456789012345678901234567890:
		newTarget := any(new).(*F0123456789012345678901234567890)
		if stage.OnAfterF0123456789012345678901234567890UpdateCallback != nil {
			stage.OnAfterF0123456789012345678901234567890UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Gstruct:
		newTarget := any(new).(*Gstruct)
		if stage.OnAfterGstructUpdateCallback != nil {
			stage.OnAfterGstructUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Astruct:
		if stage.OnAfterAstructDeleteCallback != nil {
			staged := any(staged).(*Astruct)
			stage.OnAfterAstructDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AstructBstruct2Use:
		if stage.OnAfterAstructBstruct2UseDeleteCallback != nil {
			staged := any(staged).(*AstructBstruct2Use)
			stage.OnAfterAstructBstruct2UseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AstructBstructUse:
		if stage.OnAfterAstructBstructUseDeleteCallback != nil {
			staged := any(staged).(*AstructBstructUse)
			stage.OnAfterAstructBstructUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bstruct:
		if stage.OnAfterBstructDeleteCallback != nil {
			staged := any(staged).(*Bstruct)
			stage.OnAfterBstructDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Dstruct:
		if stage.OnAfterDstructDeleteCallback != nil {
			staged := any(staged).(*Dstruct)
			stage.OnAfterDstructDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *F0123456789012345678901234567890:
		if stage.OnAfterF0123456789012345678901234567890DeleteCallback != nil {
			staged := any(staged).(*F0123456789012345678901234567890)
			stage.OnAfterF0123456789012345678901234567890DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Gstruct:
		if stage.OnAfterGstructDeleteCallback != nil {
			staged := any(staged).(*Gstruct)
			stage.OnAfterGstructDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Astruct:
		if stage.OnAfterAstructReadCallback != nil {
			stage.OnAfterAstructReadCallback.OnAfterRead(stage, target)
		}
	case *AstructBstruct2Use:
		if stage.OnAfterAstructBstruct2UseReadCallback != nil {
			stage.OnAfterAstructBstruct2UseReadCallback.OnAfterRead(stage, target)
		}
	case *AstructBstructUse:
		if stage.OnAfterAstructBstructUseReadCallback != nil {
			stage.OnAfterAstructBstructUseReadCallback.OnAfterRead(stage, target)
		}
	case *Bstruct:
		if stage.OnAfterBstructReadCallback != nil {
			stage.OnAfterBstructReadCallback.OnAfterRead(stage, target)
		}
	case *Dstruct:
		if stage.OnAfterDstructReadCallback != nil {
			stage.OnAfterDstructReadCallback.OnAfterRead(stage, target)
		}
	case *F0123456789012345678901234567890:
		if stage.OnAfterF0123456789012345678901234567890ReadCallback != nil {
			stage.OnAfterF0123456789012345678901234567890ReadCallback.OnAfterRead(stage, target)
		}
	case *Gstruct:
		if stage.OnAfterGstructReadCallback != nil {
			stage.OnAfterGstructReadCallback.OnAfterRead(stage, target)
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
	case *Astruct:
		stage.OnAfterAstructUpdateCallback = any(callback).(OnAfterUpdateInterface[Astruct])
	
	case *AstructBstruct2Use:
		stage.OnAfterAstructBstruct2UseUpdateCallback = any(callback).(OnAfterUpdateInterface[AstructBstruct2Use])
	
	case *AstructBstructUse:
		stage.OnAfterAstructBstructUseUpdateCallback = any(callback).(OnAfterUpdateInterface[AstructBstructUse])
	
	case *Bstruct:
		stage.OnAfterBstructUpdateCallback = any(callback).(OnAfterUpdateInterface[Bstruct])
	
	case *Dstruct:
		stage.OnAfterDstructUpdateCallback = any(callback).(OnAfterUpdateInterface[Dstruct])
	
	case *F0123456789012345678901234567890:
		stage.OnAfterF0123456789012345678901234567890UpdateCallback = any(callback).(OnAfterUpdateInterface[F0123456789012345678901234567890])
	
	case *Gstruct:
		stage.OnAfterGstructUpdateCallback = any(callback).(OnAfterUpdateInterface[Gstruct])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Astruct:
		stage.OnAfterAstructCreateCallback = any(callback).(OnAfterCreateInterface[Astruct])
	
	case *AstructBstruct2Use:
		stage.OnAfterAstructBstruct2UseCreateCallback = any(callback).(OnAfterCreateInterface[AstructBstruct2Use])
	
	case *AstructBstructUse:
		stage.OnAfterAstructBstructUseCreateCallback = any(callback).(OnAfterCreateInterface[AstructBstructUse])
	
	case *Bstruct:
		stage.OnAfterBstructCreateCallback = any(callback).(OnAfterCreateInterface[Bstruct])
	
	case *Dstruct:
		stage.OnAfterDstructCreateCallback = any(callback).(OnAfterCreateInterface[Dstruct])
	
	case *F0123456789012345678901234567890:
		stage.OnAfterF0123456789012345678901234567890CreateCallback = any(callback).(OnAfterCreateInterface[F0123456789012345678901234567890])
	
	case *Gstruct:
		stage.OnAfterGstructCreateCallback = any(callback).(OnAfterCreateInterface[Gstruct])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Astruct:
		stage.OnAfterAstructDeleteCallback = any(callback).(OnAfterDeleteInterface[Astruct])
	
	case *AstructBstruct2Use:
		stage.OnAfterAstructBstruct2UseDeleteCallback = any(callback).(OnAfterDeleteInterface[AstructBstruct2Use])
	
	case *AstructBstructUse:
		stage.OnAfterAstructBstructUseDeleteCallback = any(callback).(OnAfterDeleteInterface[AstructBstructUse])
	
	case *Bstruct:
		stage.OnAfterBstructDeleteCallback = any(callback).(OnAfterDeleteInterface[Bstruct])
	
	case *Dstruct:
		stage.OnAfterDstructDeleteCallback = any(callback).(OnAfterDeleteInterface[Dstruct])
	
	case *F0123456789012345678901234567890:
		stage.OnAfterF0123456789012345678901234567890DeleteCallback = any(callback).(OnAfterDeleteInterface[F0123456789012345678901234567890])
	
	case *Gstruct:
		stage.OnAfterGstructDeleteCallback = any(callback).(OnAfterDeleteInterface[Gstruct])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Astruct:
		stage.OnAfterAstructReadCallback = any(callback).(OnAfterReadInterface[Astruct])
	
	case *AstructBstruct2Use:
		stage.OnAfterAstructBstruct2UseReadCallback = any(callback).(OnAfterReadInterface[AstructBstruct2Use])
	
	case *AstructBstructUse:
		stage.OnAfterAstructBstructUseReadCallback = any(callback).(OnAfterReadInterface[AstructBstructUse])
	
	case *Bstruct:
		stage.OnAfterBstructReadCallback = any(callback).(OnAfterReadInterface[Bstruct])
	
	case *Dstruct:
		stage.OnAfterDstructReadCallback = any(callback).(OnAfterReadInterface[Dstruct])
	
	case *F0123456789012345678901234567890:
		stage.OnAfterF0123456789012345678901234567890ReadCallback = any(callback).(OnAfterReadInterface[F0123456789012345678901234567890])
	
	case *Gstruct:
		stage.OnAfterGstructReadCallback = any(callback).(OnAfterReadInterface[Gstruct])
	
	}
}
