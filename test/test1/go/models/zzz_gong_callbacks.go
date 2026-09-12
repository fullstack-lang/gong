// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
