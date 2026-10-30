// generated code - do not edit
package aixmabstractgmlobjecttypes

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
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMFeatureBaseTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMFeatureBaseTypeCreateCallback.OnAfterCreate(stage, abstractaixmfeaturebasetype)
	}
}

func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMFeatureBaseTypeUpdateCallback != nil {
		var frontAbstractAIXMFeatureBaseType *AbstractAIXMFeatureBaseType
		if front != nil {
			frontAbstractAIXMFeatureBaseType, _ = front.(*AbstractAIXMFeatureBaseType)
		}
		stage.OnAfterAbstractAIXMFeatureBaseTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmfeaturebasetype, frontAbstractAIXMFeatureBaseType)
	}
}

func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMFeatureBaseTypeDeleteCallback != nil {
		var frontAbstractAIXMFeatureBaseType *AbstractAIXMFeatureBaseType
		if front != nil {
			frontAbstractAIXMFeatureBaseType, _ = front.(*AbstractAIXMFeatureBaseType)
		}
		stage.OnAfterAbstractAIXMFeatureBaseTypeDeleteCallback.OnAfterDelete(stage, abstractaixmfeaturebasetype, frontAbstractAIXMFeatureBaseType)
	}
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMFeatureTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMFeatureTypeCreateCallback.OnAfterCreate(stage, abstractaixmfeaturetype)
	}
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMFeatureTypeUpdateCallback != nil {
		var frontAbstractAIXMFeatureType *AbstractAIXMFeatureType
		if front != nil {
			frontAbstractAIXMFeatureType, _ = front.(*AbstractAIXMFeatureType)
		}
		stage.OnAfterAbstractAIXMFeatureTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmfeaturetype, frontAbstractAIXMFeatureType)
	}
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMFeatureTypeDeleteCallback != nil {
		var frontAbstractAIXMFeatureType *AbstractAIXMFeatureType
		if front != nil {
			frontAbstractAIXMFeatureType, _ = front.(*AbstractAIXMFeatureType)
		}
		stage.OnAfterAbstractAIXMFeatureTypeDeleteCallback.OnAfterDelete(stage, abstractaixmfeaturetype, frontAbstractAIXMFeatureType)
	}
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMMessageBaseTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMMessageBaseTypeCreateCallback.OnAfterCreate(stage, abstractaixmmessagebasetype)
	}
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMMessageBaseTypeUpdateCallback != nil {
		var frontAbstractAIXMMessageBaseType *AbstractAIXMMessageBaseType
		if front != nil {
			frontAbstractAIXMMessageBaseType, _ = front.(*AbstractAIXMMessageBaseType)
		}
		stage.OnAfterAbstractAIXMMessageBaseTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmmessagebasetype, frontAbstractAIXMMessageBaseType)
	}
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMMessageBaseTypeDeleteCallback != nil {
		var frontAbstractAIXMMessageBaseType *AbstractAIXMMessageBaseType
		if front != nil {
			frontAbstractAIXMMessageBaseType, _ = front.(*AbstractAIXMMessageBaseType)
		}
		stage.OnAfterAbstractAIXMMessageBaseTypeDeleteCallback.OnAfterDelete(stage, abstractaixmmessagebasetype, frontAbstractAIXMMessageBaseType)
	}
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMMessageTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMMessageTypeCreateCallback.OnAfterCreate(stage, abstractaixmmessagetype)
	}
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMMessageTypeUpdateCallback != nil {
		var frontAbstractAIXMMessageType *AbstractAIXMMessageType
		if front != nil {
			frontAbstractAIXMMessageType, _ = front.(*AbstractAIXMMessageType)
		}
		stage.OnAfterAbstractAIXMMessageTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmmessagetype, frontAbstractAIXMMessageType)
	}
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMMessageTypeDeleteCallback != nil {
		var frontAbstractAIXMMessageType *AbstractAIXMMessageType
		if front != nil {
			frontAbstractAIXMMessageType, _ = front.(*AbstractAIXMMessageType)
		}
		stage.OnAfterAbstractAIXMMessageTypeDeleteCallback.OnAfterDelete(stage, abstractaixmmessagetype, frontAbstractAIXMMessageType)
	}
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMObjectTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMObjectTypeCreateCallback.OnAfterCreate(stage, abstractaixmobjecttype)
	}
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMObjectTypeUpdateCallback != nil {
		var frontAbstractAIXMObjectType *AbstractAIXMObjectType
		if front != nil {
			frontAbstractAIXMObjectType, _ = front.(*AbstractAIXMObjectType)
		}
		stage.OnAfterAbstractAIXMObjectTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmobjecttype, frontAbstractAIXMObjectType)
	}
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMObjectTypeDeleteCallback != nil {
		var frontAbstractAIXMObjectType *AbstractAIXMObjectType
		if front != nil {
			frontAbstractAIXMObjectType, _ = front.(*AbstractAIXMObjectType)
		}
		stage.OnAfterAbstractAIXMObjectTypeDeleteCallback.OnAfterDelete(stage, abstractaixmobjecttype, frontAbstractAIXMObjectType)
	}
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMPropertyTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMPropertyTypeCreateCallback.OnAfterCreate(stage, abstractaixmpropertytype)
	}
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMPropertyTypeUpdateCallback != nil {
		var frontAbstractAIXMPropertyType *AbstractAIXMPropertyType
		if front != nil {
			frontAbstractAIXMPropertyType, _ = front.(*AbstractAIXMPropertyType)
		}
		stage.OnAfterAbstractAIXMPropertyTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmpropertytype, frontAbstractAIXMPropertyType)
	}
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMPropertyTypeDeleteCallback != nil {
		var frontAbstractAIXMPropertyType *AbstractAIXMPropertyType
		if front != nil {
			frontAbstractAIXMPropertyType, _ = front.(*AbstractAIXMPropertyType)
		}
		stage.OnAfterAbstractAIXMPropertyTypeDeleteCallback.OnAfterDelete(stage, abstractaixmpropertytype, frontAbstractAIXMPropertyType)
	}
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMTimeSliceBaseTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMTimeSliceBaseTypeCreateCallback.OnAfterCreate(stage, abstractaixmtimeslicebasetype)
	}
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMTimeSliceBaseTypeUpdateCallback != nil {
		var frontAbstractAIXMTimeSliceBaseType *AbstractAIXMTimeSliceBaseType
		if front != nil {
			frontAbstractAIXMTimeSliceBaseType, _ = front.(*AbstractAIXMTimeSliceBaseType)
		}
		stage.OnAfterAbstractAIXMTimeSliceBaseTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmtimeslicebasetype, frontAbstractAIXMTimeSliceBaseType)
	}
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMTimeSliceBaseTypeDeleteCallback != nil {
		var frontAbstractAIXMTimeSliceBaseType *AbstractAIXMTimeSliceBaseType
		if front != nil {
			frontAbstractAIXMTimeSliceBaseType, _ = front.(*AbstractAIXMTimeSliceBaseType)
		}
		stage.OnAfterAbstractAIXMTimeSliceBaseTypeDeleteCallback.OnAfterDelete(stage, abstractaixmtimeslicebasetype, frontAbstractAIXMTimeSliceBaseType)
	}
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractAIXMTimeSliceTypeCreateCallback != nil {
		stage.OnAfterAbstractAIXMTimeSliceTypeCreateCallback.OnAfterCreate(stage, abstractaixmtimeslicetype)
	}
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMTimeSliceTypeUpdateCallback != nil {
		var frontAbstractAIXMTimeSliceType *AbstractAIXMTimeSliceType
		if front != nil {
			frontAbstractAIXMTimeSliceType, _ = front.(*AbstractAIXMTimeSliceType)
		}
		stage.OnAfterAbstractAIXMTimeSliceTypeUpdateCallback.OnAfterUpdate(stage, abstractaixmtimeslicetype, frontAbstractAIXMTimeSliceType)
	}
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractAIXMTimeSliceTypeDeleteCallback != nil {
		var frontAbstractAIXMTimeSliceType *AbstractAIXMTimeSliceType
		if front != nil {
			frontAbstractAIXMTimeSliceType, _ = front.(*AbstractAIXMTimeSliceType)
		}
		stage.OnAfterAbstractAIXMTimeSliceTypeDeleteCallback.OnAfterDelete(stage, abstractaixmtimeslicetype, frontAbstractAIXMTimeSliceType)
	}
}

func (abstractextensiontype *AbstractExtensionType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAbstractExtensionTypeCreateCallback != nil {
		stage.OnAfterAbstractExtensionTypeCreateCallback.OnAfterCreate(stage, abstractextensiontype)
	}
}

func (abstractextensiontype *AbstractExtensionType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractExtensionTypeUpdateCallback != nil {
		var frontAbstractExtensionType *AbstractExtensionType
		if front != nil {
			frontAbstractExtensionType, _ = front.(*AbstractExtensionType)
		}
		stage.OnAfterAbstractExtensionTypeUpdateCallback.OnAfterUpdate(stage, abstractextensiontype, frontAbstractExtensionType)
	}
}

func (abstractextensiontype *AbstractExtensionType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAbstractExtensionTypeDeleteCallback != nil {
		var frontAbstractExtensionType *AbstractExtensionType
		if front != nil {
			frontAbstractExtensionType, _ = front.(*AbstractExtensionType)
		}
		stage.OnAfterAbstractExtensionTypeDeleteCallback.OnAfterDelete(stage, abstractextensiontype, frontAbstractExtensionType)
	}
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFeatureMetadataPropertyTypeCreateCallback != nil {
		stage.OnAfterFeatureMetadataPropertyTypeCreateCallback.OnAfterCreate(stage, featuremetadatapropertytype)
	}
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureMetadataPropertyTypeUpdateCallback != nil {
		var frontFeatureMetadataPropertyType *FeatureMetadataPropertyType
		if front != nil {
			frontFeatureMetadataPropertyType, _ = front.(*FeatureMetadataPropertyType)
		}
		stage.OnAfterFeatureMetadataPropertyTypeUpdateCallback.OnAfterUpdate(stage, featuremetadatapropertytype, frontFeatureMetadataPropertyType)
	}
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureMetadataPropertyTypeDeleteCallback != nil {
		var frontFeatureMetadataPropertyType *FeatureMetadataPropertyType
		if front != nil {
			frontFeatureMetadataPropertyType, _ = front.(*FeatureMetadataPropertyType)
		}
		stage.OnAfterFeatureMetadataPropertyTypeDeleteCallback.OnAfterDelete(stage, featuremetadatapropertytype, frontFeatureMetadataPropertyType)
	}
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterFeatureTimeSliceMetadataPropertyTypeCreateCallback != nil {
		stage.OnAfterFeatureTimeSliceMetadataPropertyTypeCreateCallback.OnAfterCreate(stage, featuretimeslicemetadatapropertytype)
	}
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureTimeSliceMetadataPropertyTypeUpdateCallback != nil {
		var frontFeatureTimeSliceMetadataPropertyType *FeatureTimeSliceMetadataPropertyType
		if front != nil {
			frontFeatureTimeSliceMetadataPropertyType, _ = front.(*FeatureTimeSliceMetadataPropertyType)
		}
		stage.OnAfterFeatureTimeSliceMetadataPropertyTypeUpdateCallback.OnAfterUpdate(stage, featuretimeslicemetadatapropertytype, frontFeatureTimeSliceMetadataPropertyType)
	}
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterFeatureTimeSliceMetadataPropertyTypeDeleteCallback != nil {
		var frontFeatureTimeSliceMetadataPropertyType *FeatureTimeSliceMetadataPropertyType
		if front != nil {
			frontFeatureTimeSliceMetadataPropertyType, _ = front.(*FeatureTimeSliceMetadataPropertyType)
		}
		stage.OnAfterFeatureTimeSliceMetadataPropertyTypeDeleteCallback.OnAfterDelete(stage, featuretimeslicemetadatapropertytype, frontFeatureTimeSliceMetadataPropertyType)
	}
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMessageMetadataPropertyTypeCreateCallback != nil {
		stage.OnAfterMessageMetadataPropertyTypeCreateCallback.OnAfterCreate(stage, messagemetadatapropertytype)
	}
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageMetadataPropertyTypeUpdateCallback != nil {
		var frontMessageMetadataPropertyType *MessageMetadataPropertyType
		if front != nil {
			frontMessageMetadataPropertyType, _ = front.(*MessageMetadataPropertyType)
		}
		stage.OnAfterMessageMetadataPropertyTypeUpdateCallback.OnAfterUpdate(stage, messagemetadatapropertytype, frontMessageMetadataPropertyType)
	}
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMessageMetadataPropertyTypeDeleteCallback != nil {
		var frontMessageMetadataPropertyType *MessageMetadataPropertyType
		if front != nil {
			frontMessageMetadataPropertyType, _ = front.(*MessageMetadataPropertyType)
		}
		stage.OnAfterMessageMetadataPropertyTypeDeleteCallback.OnAfterDelete(stage, messagemetadatapropertytype, frontMessageMetadataPropertyType)
	}
}

