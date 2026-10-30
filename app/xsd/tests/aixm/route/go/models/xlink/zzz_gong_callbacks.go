// generated code - do not edit
package xlink

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
func (xlink_schemaelement *Xlink_SchemaElement) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXlink_SchemaElementCreateCallback != nil {
		stage.OnAfterXlink_SchemaElementCreateCallback.OnAfterCreate(stage, xlink_schemaelement)
	}
}

func (xlink_schemaelement *Xlink_SchemaElement) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXlink_SchemaElementUpdateCallback != nil {
		var frontXlink_SchemaElement *Xlink_SchemaElement
		if front != nil {
			frontXlink_SchemaElement, _ = front.(*Xlink_SchemaElement)
		}
		stage.OnAfterXlink_SchemaElementUpdateCallback.OnAfterUpdate(stage, xlink_schemaelement, frontXlink_SchemaElement)
	}
}

func (xlink_schemaelement *Xlink_SchemaElement) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXlink_SchemaElementDeleteCallback != nil {
		var frontXlink_SchemaElement *Xlink_SchemaElement
		if front != nil {
			frontXlink_SchemaElement, _ = front.(*Xlink_SchemaElement)
		}
		stage.OnAfterXlink_SchemaElementDeleteCallback.OnAfterDelete(stage, xlink_schemaelement, frontXlink_SchemaElement)
	}
}

