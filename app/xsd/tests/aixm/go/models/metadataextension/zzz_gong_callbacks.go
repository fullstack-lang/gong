// generated code - do not edit
package metadataextension

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
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_DatatypeCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_DatatypeCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_datatypecode_propertytype)
	}
}

func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DatatypeCode_PropertyTypeUpdateCallback != nil {
		var frontMD_DatatypeCode_PropertyType *MD_DatatypeCode_PropertyType
		if front != nil {
			frontMD_DatatypeCode_PropertyType, _ = front.(*MD_DatatypeCode_PropertyType)
		}
		stage.OnAfterMD_DatatypeCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_datatypecode_propertytype, frontMD_DatatypeCode_PropertyType)
	}
}

func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_DatatypeCode_PropertyTypeDeleteCallback != nil {
		var frontMD_DatatypeCode_PropertyType *MD_DatatypeCode_PropertyType
		if front != nil {
			frontMD_DatatypeCode_PropertyType, _ = front.(*MD_DatatypeCode_PropertyType)
		}
		stage.OnAfterMD_DatatypeCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_datatypecode_propertytype, frontMD_DatatypeCode_PropertyType)
	}
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ExtendedElementInformation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ExtendedElementInformation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_extendedelementinformation_propertytype)
	}
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ExtendedElementInformation_PropertyTypeUpdateCallback != nil {
		var frontMD_ExtendedElementInformation_PropertyType *MD_ExtendedElementInformation_PropertyType
		if front != nil {
			frontMD_ExtendedElementInformation_PropertyType, _ = front.(*MD_ExtendedElementInformation_PropertyType)
		}
		stage.OnAfterMD_ExtendedElementInformation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_extendedelementinformation_propertytype, frontMD_ExtendedElementInformation_PropertyType)
	}
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ExtendedElementInformation_PropertyTypeDeleteCallback != nil {
		var frontMD_ExtendedElementInformation_PropertyType *MD_ExtendedElementInformation_PropertyType
		if front != nil {
			frontMD_ExtendedElementInformation_PropertyType, _ = front.(*MD_ExtendedElementInformation_PropertyType)
		}
		stage.OnAfterMD_ExtendedElementInformation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_extendedelementinformation_propertytype, frontMD_ExtendedElementInformation_PropertyType)
	}
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ExtendedElementInformation_TypeCreateCallback != nil {
		stage.OnAfterMD_ExtendedElementInformation_TypeCreateCallback.OnAfterCreate(stage, md_extendedelementinformation_type)
	}
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ExtendedElementInformation_TypeUpdateCallback != nil {
		var frontMD_ExtendedElementInformation_Type *MD_ExtendedElementInformation_Type
		if front != nil {
			frontMD_ExtendedElementInformation_Type, _ = front.(*MD_ExtendedElementInformation_Type)
		}
		stage.OnAfterMD_ExtendedElementInformation_TypeUpdateCallback.OnAfterUpdate(stage, md_extendedelementinformation_type, frontMD_ExtendedElementInformation_Type)
	}
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ExtendedElementInformation_TypeDeleteCallback != nil {
		var frontMD_ExtendedElementInformation_Type *MD_ExtendedElementInformation_Type
		if front != nil {
			frontMD_ExtendedElementInformation_Type, _ = front.(*MD_ExtendedElementInformation_Type)
		}
		stage.OnAfterMD_ExtendedElementInformation_TypeDeleteCallback.OnAfterDelete(stage, md_extendedelementinformation_type, frontMD_ExtendedElementInformation_Type)
	}
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_MetadataExtensionInformation_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_MetadataExtensionInformation_PropertyTypeCreateCallback.OnAfterCreate(stage, md_metadataextensioninformation_propertytype)
	}
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MetadataExtensionInformation_PropertyTypeUpdateCallback != nil {
		var frontMD_MetadataExtensionInformation_PropertyType *MD_MetadataExtensionInformation_PropertyType
		if front != nil {
			frontMD_MetadataExtensionInformation_PropertyType, _ = front.(*MD_MetadataExtensionInformation_PropertyType)
		}
		stage.OnAfterMD_MetadataExtensionInformation_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_metadataextensioninformation_propertytype, frontMD_MetadataExtensionInformation_PropertyType)
	}
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MetadataExtensionInformation_PropertyTypeDeleteCallback != nil {
		var frontMD_MetadataExtensionInformation_PropertyType *MD_MetadataExtensionInformation_PropertyType
		if front != nil {
			frontMD_MetadataExtensionInformation_PropertyType, _ = front.(*MD_MetadataExtensionInformation_PropertyType)
		}
		stage.OnAfterMD_MetadataExtensionInformation_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_metadataextensioninformation_propertytype, frontMD_MetadataExtensionInformation_PropertyType)
	}
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_MetadataExtensionInformation_TypeCreateCallback != nil {
		stage.OnAfterMD_MetadataExtensionInformation_TypeCreateCallback.OnAfterCreate(stage, md_metadataextensioninformation_type)
	}
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MetadataExtensionInformation_TypeUpdateCallback != nil {
		var frontMD_MetadataExtensionInformation_Type *MD_MetadataExtensionInformation_Type
		if front != nil {
			frontMD_MetadataExtensionInformation_Type, _ = front.(*MD_MetadataExtensionInformation_Type)
		}
		stage.OnAfterMD_MetadataExtensionInformation_TypeUpdateCallback.OnAfterUpdate(stage, md_metadataextensioninformation_type, frontMD_MetadataExtensionInformation_Type)
	}
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_MetadataExtensionInformation_TypeDeleteCallback != nil {
		var frontMD_MetadataExtensionInformation_Type *MD_MetadataExtensionInformation_Type
		if front != nil {
			frontMD_MetadataExtensionInformation_Type, _ = front.(*MD_MetadataExtensionInformation_Type)
		}
		stage.OnAfterMD_MetadataExtensionInformation_TypeDeleteCallback.OnAfterDelete(stage, md_metadataextensioninformation_type, frontMD_MetadataExtensionInformation_Type)
	}
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMD_ObligationCode_PropertyTypeCreateCallback != nil {
		stage.OnAfterMD_ObligationCode_PropertyTypeCreateCallback.OnAfterCreate(stage, md_obligationcode_propertytype)
	}
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ObligationCode_PropertyTypeUpdateCallback != nil {
		var frontMD_ObligationCode_PropertyType *MD_ObligationCode_PropertyType
		if front != nil {
			frontMD_ObligationCode_PropertyType, _ = front.(*MD_ObligationCode_PropertyType)
		}
		stage.OnAfterMD_ObligationCode_PropertyTypeUpdateCallback.OnAfterUpdate(stage, md_obligationcode_propertytype, frontMD_ObligationCode_PropertyType)
	}
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMD_ObligationCode_PropertyTypeDeleteCallback != nil {
		var frontMD_ObligationCode_PropertyType *MD_ObligationCode_PropertyType
		if front != nil {
			frontMD_ObligationCode_PropertyType, _ = front.(*MD_ObligationCode_PropertyType)
		}
		stage.OnAfterMD_ObligationCode_PropertyTypeDeleteCallback.OnAfterDelete(stage, md_obligationcode_propertytype, frontMD_ObligationCode_PropertyType)
	}
}

