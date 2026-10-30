// generated code - do not edit
package refs

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
func (reftargettype *RefTargetType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRefTargetTypeCreateCallback != nil {
		stage.OnAfterRefTargetTypeCreateCallback.OnAfterCreate(stage, reftargettype)
	}
}

func (reftargettype *RefTargetType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRefTargetTypeUpdateCallback != nil {
		var frontRefTargetType *RefTargetType
		if front != nil {
			frontRefTargetType, _ = front.(*RefTargetType)
		}
		stage.OnAfterRefTargetTypeUpdateCallback.OnAfterUpdate(stage, reftargettype, frontRefTargetType)
	}
}

func (reftargettype *RefTargetType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRefTargetTypeDeleteCallback != nil {
		var frontRefTargetType *RefTargetType
		if front != nil {
			frontRefTargetType, _ = front.(*RefTargetType)
		}
		stage.OnAfterRefTargetTypeDeleteCallback.OnAfterDelete(stage, reftargettype, frontRefTargetType)
	}
}

