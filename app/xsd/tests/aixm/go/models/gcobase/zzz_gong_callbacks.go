// generated code - do not edit
package gcobase

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
func (abstractobject_type *AbstractObject_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractObject_TypeCreateCallback != nil {
		stage.OnAfterAbstractObject_TypeCreateCallback.OnAfterCreate(stage, abstractobject_type)
	}
}

func (abstractobject_type *AbstractObject_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractObject_TypeUpdateCallback != nil {
		var frontAbstractObject_Type *AbstractObject_Type
		if front != nil {
			frontAbstractObject_Type, _ = front.(*AbstractObject_Type)
		}
		stage.OnAfterAbstractObject_TypeUpdateCallback.OnAfterUpdate(stage, abstractobject_type, frontAbstractObject_Type)
	}
}

func (abstractobject_type *AbstractObject_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractObject_TypeDeleteCallback != nil {
		var frontAbstractObject_Type *AbstractObject_Type
		if front != nil {
			frontAbstractObject_Type, _ = front.(*AbstractObject_Type)
		}
		stage.OnAfterAbstractObject_TypeDeleteCallback.OnAfterDelete(stage, abstractobject_type, frontAbstractObject_Type)
	}
}

func (codelistvalue_type *CodeListValue_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterCodeListValue_TypeCreateCallback != nil {
		stage.OnAfterCodeListValue_TypeCreateCallback.OnAfterCreate(stage, codelistvalue_type)
	}
}

func (codelistvalue_type *CodeListValue_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeListValue_TypeUpdateCallback != nil {
		var frontCodeListValue_Type *CodeListValue_Type
		if front != nil {
			frontCodeListValue_Type, _ = front.(*CodeListValue_Type)
		}
		stage.OnAfterCodeListValue_TypeUpdateCallback.OnAfterUpdate(stage, codelistvalue_type, frontCodeListValue_Type)
	}
}

func (codelistvalue_type *CodeListValue_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterCodeListValue_TypeDeleteCallback != nil {
		var frontCodeListValue_Type *CodeListValue_Type
		if front != nil {
			frontCodeListValue_Type, _ = front.(*CodeListValue_Type)
		}
		stage.OnAfterCodeListValue_TypeDeleteCallback.OnAfterDelete(stage, codelistvalue_type, frontCodeListValue_Type)
	}
}

func (objectreference_propertytype *ObjectReference_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterObjectReference_PropertyTypeCreateCallback != nil {
		stage.OnAfterObjectReference_PropertyTypeCreateCallback.OnAfterCreate(stage, objectreference_propertytype)
	}
}

func (objectreference_propertytype *ObjectReference_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObjectReference_PropertyTypeUpdateCallback != nil {
		var frontObjectReference_PropertyType *ObjectReference_PropertyType
		if front != nil {
			frontObjectReference_PropertyType, _ = front.(*ObjectReference_PropertyType)
		}
		stage.OnAfterObjectReference_PropertyTypeUpdateCallback.OnAfterUpdate(stage, objectreference_propertytype, frontObjectReference_PropertyType)
	}
}

func (objectreference_propertytype *ObjectReference_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterObjectReference_PropertyTypeDeleteCallback != nil {
		var frontObjectReference_PropertyType *ObjectReference_PropertyType
		if front != nil {
			frontObjectReference_PropertyType, _ = front.(*ObjectReference_PropertyType)
		}
		stage.OnAfterObjectReference_PropertyTypeDeleteCallback.OnAfterDelete(stage, objectreference_propertytype, frontObjectReference_PropertyType)
	}
}

