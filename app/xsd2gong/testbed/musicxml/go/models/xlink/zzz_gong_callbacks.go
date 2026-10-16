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
func (arctype *ArcType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArcTypeCreateCallback != nil {
		stage.OnAfterArcTypeCreateCallback.OnAfterCreate(stage, arctype)
	}
}

func (arctype *ArcType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcTypeUpdateCallback != nil {
		var frontArcType *ArcType
		if front != nil {
			frontArcType, _ = front.(*ArcType)
		}
		stage.OnAfterArcTypeUpdateCallback.OnAfterUpdate(stage, arctype, frontArcType)
	}
}

func (arctype *ArcType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArcTypeDeleteCallback != nil {
		var frontArcType *ArcType
		if front != nil {
			frontArcType, _ = front.(*ArcType)
		}
		stage.OnAfterArcTypeDeleteCallback.OnAfterDelete(stage, arctype, frontArcType)
	}
}

func (extended *Extended) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterExtendedCreateCallback != nil {
		stage.OnAfterExtendedCreateCallback.OnAfterCreate(stage, extended)
	}
}

func (extended *Extended) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtendedUpdateCallback != nil {
		var frontExtended *Extended
		if front != nil {
			frontExtended, _ = front.(*Extended)
		}
		stage.OnAfterExtendedUpdateCallback.OnAfterUpdate(stage, extended, frontExtended)
	}
}

func (extended *Extended) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExtendedDeleteCallback != nil {
		var frontExtended *Extended
		if front != nil {
			frontExtended, _ = front.(*Extended)
		}
		stage.OnAfterExtendedDeleteCallback.OnAfterDelete(stage, extended, frontExtended)
	}
}

func (locatortype *LocatorType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLocatorTypeCreateCallback != nil {
		stage.OnAfterLocatorTypeCreateCallback.OnAfterCreate(stage, locatortype)
	}
}

func (locatortype *LocatorType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocatorTypeUpdateCallback != nil {
		var frontLocatorType *LocatorType
		if front != nil {
			frontLocatorType, _ = front.(*LocatorType)
		}
		stage.OnAfterLocatorTypeUpdateCallback.OnAfterUpdate(stage, locatortype, frontLocatorType)
	}
}

func (locatortype *LocatorType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLocatorTypeDeleteCallback != nil {
		var frontLocatorType *LocatorType
		if front != nil {
			frontLocatorType, _ = front.(*LocatorType)
		}
		stage.OnAfterLocatorTypeDeleteCallback.OnAfterDelete(stage, locatortype, frontLocatorType)
	}
}

func (resourcetype *ResourceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResourceTypeCreateCallback != nil {
		stage.OnAfterResourceTypeCreateCallback.OnAfterCreate(stage, resourcetype)
	}
}

func (resourcetype *ResourceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceTypeUpdateCallback != nil {
		var frontResourceType *ResourceType
		if front != nil {
			frontResourceType, _ = front.(*ResourceType)
		}
		stage.OnAfterResourceTypeUpdateCallback.OnAfterUpdate(stage, resourcetype, frontResourceType)
	}
}

func (resourcetype *ResourceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceTypeDeleteCallback != nil {
		var frontResourceType *ResourceType
		if front != nil {
			frontResourceType, _ = front.(*ResourceType)
		}
		stage.OnAfterResourceTypeDeleteCallback.OnAfterDelete(stage, resourcetype, frontResourceType)
	}
}

func (simple *Simple) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSimpleCreateCallback != nil {
		stage.OnAfterSimpleCreateCallback.OnAfterCreate(stage, simple)
	}
}

func (simple *Simple) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSimpleUpdateCallback != nil {
		var frontSimple *Simple
		if front != nil {
			frontSimple, _ = front.(*Simple)
		}
		stage.OnAfterSimpleUpdateCallback.OnAfterUpdate(stage, simple, frontSimple)
	}
}

func (simple *Simple) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSimpleDeleteCallback != nil {
		var frontSimple *Simple
		if front != nil {
			frontSimple, _ = front.(*Simple)
		}
		stage.OnAfterSimpleDeleteCallback.OnAfterDelete(stage, simple, frontSimple)
	}
}

func (titleelttype *TitleEltType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTitleEltTypeCreateCallback != nil {
		stage.OnAfterTitleEltTypeCreateCallback.OnAfterCreate(stage, titleelttype)
	}
}

func (titleelttype *TitleEltType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTitleEltTypeUpdateCallback != nil {
		var frontTitleEltType *TitleEltType
		if front != nil {
			frontTitleEltType, _ = front.(*TitleEltType)
		}
		stage.OnAfterTitleEltTypeUpdateCallback.OnAfterUpdate(stage, titleelttype, frontTitleEltType)
	}
}

func (titleelttype *TitleEltType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTitleEltTypeDeleteCallback != nil {
		var frontTitleEltType *TitleEltType
		if front != nil {
			frontTitleEltType, _ = front.(*TitleEltType)
		}
		stage.OnAfterTitleEltTypeDeleteCallback.OnAfterDelete(stage, titleelttype, frontTitleEltType)
	}
}

