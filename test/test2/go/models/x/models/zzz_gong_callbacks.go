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
func (submodel *SubModel) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSubModelCreateCallback != nil {
		stage.OnAfterSubModelCreateCallback.OnAfterCreate(stage, submodel)
	}
}

func (submodel *SubModel) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSubModelUpdateCallback != nil {
		var frontSubModel *SubModel
		if front != nil {
			frontSubModel, _ = front.(*SubModel)
		}
		stage.OnAfterSubModelUpdateCallback.OnAfterUpdate(stage, submodel, frontSubModel)
	}
}

func (submodel *SubModel) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSubModelDeleteCallback != nil {
		var frontSubModel *SubModel
		if front != nil {
			frontSubModel, _ = front.(*SubModel)
		}
		stage.OnAfterSubModelDeleteCallback.OnAfterDelete(stage, submodel, frontSubModel)
	}
}

