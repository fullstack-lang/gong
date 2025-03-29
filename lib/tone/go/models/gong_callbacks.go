// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Freqency:
		if stage.OnAfterFreqencyCreateCallback != nil {
			stage.OnAfterFreqencyCreateCallback.OnAfterCreate(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteCreateCallback != nil {
			stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, target)
		}
	case *Player:
		if stage.OnAfterPlayerCreateCallback != nil {
			stage.OnAfterPlayerCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Freqency:
		newTarget := any(new).(*Freqency)
		if stage.OnAfterFreqencyUpdateCallback != nil {
			stage.OnAfterFreqencyUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Note:
		newTarget := any(new).(*Note)
		if stage.OnAfterNoteUpdateCallback != nil {
			stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Player:
		newTarget := any(new).(*Player)
		if stage.OnAfterPlayerUpdateCallback != nil {
			stage.OnAfterPlayerUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Freqency:
		if stage.OnAfterFreqencyDeleteCallback != nil {
			staged := any(staged).(*Freqency)
			stage.OnAfterFreqencyDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Note:
		if stage.OnAfterNoteDeleteCallback != nil {
			staged := any(staged).(*Note)
			stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Player:
		if stage.OnAfterPlayerDeleteCallback != nil {
			staged := any(staged).(*Player)
			stage.OnAfterPlayerDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Freqency:
		if stage.OnAfterFreqencyReadCallback != nil {
			stage.OnAfterFreqencyReadCallback.OnAfterRead(stage, target)
		}
	case *Note:
		if stage.OnAfterNoteReadCallback != nil {
			stage.OnAfterNoteReadCallback.OnAfterRead(stage, target)
		}
	case *Player:
		if stage.OnAfterPlayerReadCallback != nil {
			stage.OnAfterPlayerReadCallback.OnAfterRead(stage, target)
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
	case *Freqency:
		stage.OnAfterFreqencyUpdateCallback = any(callback).(OnAfterUpdateInterface[Freqency])
	
	case *Note:
		stage.OnAfterNoteUpdateCallback = any(callback).(OnAfterUpdateInterface[Note])
	
	case *Player:
		stage.OnAfterPlayerUpdateCallback = any(callback).(OnAfterUpdateInterface[Player])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Freqency:
		stage.OnAfterFreqencyCreateCallback = any(callback).(OnAfterCreateInterface[Freqency])
	
	case *Note:
		stage.OnAfterNoteCreateCallback = any(callback).(OnAfterCreateInterface[Note])
	
	case *Player:
		stage.OnAfterPlayerCreateCallback = any(callback).(OnAfterCreateInterface[Player])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Freqency:
		stage.OnAfterFreqencyDeleteCallback = any(callback).(OnAfterDeleteInterface[Freqency])
	
	case *Note:
		stage.OnAfterNoteDeleteCallback = any(callback).(OnAfterDeleteInterface[Note])
	
	case *Player:
		stage.OnAfterPlayerDeleteCallback = any(callback).(OnAfterDeleteInterface[Player])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Freqency:
		stage.OnAfterFreqencyReadCallback = any(callback).(OnAfterReadInterface[Freqency])
	
	case *Note:
		stage.OnAfterNoteReadCallback = any(callback).(OnAfterReadInterface[Note])
	
	case *Player:
		stage.OnAfterPlayerReadCallback = any(callback).(OnAfterReadInterface[Player])
	
	}
}
