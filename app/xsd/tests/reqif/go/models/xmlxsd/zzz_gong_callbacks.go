// generated code - do not edit
package xmlxsd

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
func (xmlxsd_schemaelement *Xmlxsd_SchemaElement) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterXmlxsd_SchemaElementCreateCallback != nil {
		stage.OnAfterXmlxsd_SchemaElementCreateCallback.OnAfterCreate(stage, xmlxsd_schemaelement)
	}
}

func (xmlxsd_schemaelement *Xmlxsd_SchemaElement) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXmlxsd_SchemaElementUpdateCallback != nil {
		var frontXmlxsd_SchemaElement *Xmlxsd_SchemaElement
		if front != nil {
			frontXmlxsd_SchemaElement, _ = front.(*Xmlxsd_SchemaElement)
		}
		stage.OnAfterXmlxsd_SchemaElementUpdateCallback.OnAfterUpdate(stage, xmlxsd_schemaelement, frontXmlxsd_SchemaElement)
	}
}

func (xmlxsd_schemaelement *Xmlxsd_SchemaElement) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterXmlxsd_SchemaElementDeleteCallback != nil {
		var frontXmlxsd_SchemaElement *Xmlxsd_SchemaElement
		if front != nil {
			frontXmlxsd_SchemaElement, _ = front.(*Xmlxsd_SchemaElement)
		}
		stage.OnAfterXmlxsd_SchemaElementDeleteCallback.OnAfterDelete(stage, xmlxsd_schemaelement, frontXmlxsd_SchemaElement)
	}
}

