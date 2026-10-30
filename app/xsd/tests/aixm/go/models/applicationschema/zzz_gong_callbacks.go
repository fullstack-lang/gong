// generated code - do not edit
package applicationschema

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
func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ApplicationSchemaInformation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ApplicationSchemaInformation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_applicationschemainformation_propertytype)
	}
}

func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ApplicationSchemaInformation_PropertyTypeUpdateCallback != nil {
		var frontMD_ApplicationSchemaInformation_PropertyType *MD_ApplicationSchemaInformation_PropertyType
		if front != nil {
			frontMD_ApplicationSchemaInformation_PropertyType, _ = front.(*MD_ApplicationSchemaInformation_PropertyType)
		}
		stage.OnAfterMD_ApplicationSchemaInformation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_applicationschemainformation_propertytype, frontMD_ApplicationSchemaInformation_PropertyType)
	}
}

func (md_applicationschemainformation_propertytype *MD_ApplicationSchemaInformation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ApplicationSchemaInformation_PropertyTypeDeleteCallback != nil {
		var frontMD_ApplicationSchemaInformation_PropertyType *MD_ApplicationSchemaInformation_PropertyType
		if front != nil {
			frontMD_ApplicationSchemaInformation_PropertyType, _ = front.(*MD_ApplicationSchemaInformation_PropertyType)
		}
		stage.OnAfterMD_ApplicationSchemaInformation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_applicationschemainformation_propertytype, frontMD_ApplicationSchemaInformation_PropertyType)
	}
}

func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ApplicationSchemaInformation_TypeCreateCallback != nil {
		stage.OnAfterMD_ApplicationSchemaInformation_TypeCreateCallback.OnAfterCreate(stage, md_applicationschemainformation_type)
	}
}

func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ApplicationSchemaInformation_TypeUpdateCallback != nil {
		var frontMD_ApplicationSchemaInformation_Type *MD_ApplicationSchemaInformation_Type
		if front != nil {
			frontMD_ApplicationSchemaInformation_Type, _ = front.(*MD_ApplicationSchemaInformation_Type)
		}
		stage.OnAfterMD_ApplicationSchemaInformation_TypeUpdateCallback.OnAfterUpdate(stage, md_applicationschemainformation_type, frontMD_ApplicationSchemaInformation_Type)
	}
}

func (md_applicationschemainformation_type *MD_ApplicationSchemaInformation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ApplicationSchemaInformation_TypeDeleteCallback != nil {
		var frontMD_ApplicationSchemaInformation_Type *MD_ApplicationSchemaInformation_Type
		if front != nil {
			frontMD_ApplicationSchemaInformation_Type, _ = front.(*MD_ApplicationSchemaInformation_Type)
		}
		stage.OnAfterMD_ApplicationSchemaInformation_TypeDeleteCallback.OnAfterDelete(stage, md_applicationschemainformation_type, frontMD_ApplicationSchemaInformation_Type)
	}
}

