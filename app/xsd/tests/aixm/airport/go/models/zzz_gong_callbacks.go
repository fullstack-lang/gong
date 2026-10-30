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
func (aixmbasicmessage *AIXMBasicMessage) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAIXMBasicMessageCreateCallback != nil {
		stage.OnAfterAIXMBasicMessageCreateCallback.OnAfterCreate(stage, aixmbasicmessage)
	}
}

func (aixmbasicmessage *AIXMBasicMessage) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAIXMBasicMessageUpdateCallback != nil {
		var frontAIXMBasicMessage *AIXMBasicMessage
		if front != nil {
			frontAIXMBasicMessage, _ = front.(*AIXMBasicMessage)
		}
		stage.OnAfterAIXMBasicMessageUpdateCallback.OnAfterUpdate(stage, aixmbasicmessage, frontAIXMBasicMessage)
	}
}

func (aixmbasicmessage *AIXMBasicMessage) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAIXMBasicMessageDeleteCallback != nil {
		var frontAIXMBasicMessage *AIXMBasicMessage
		if front != nil {
			frontAIXMBasicMessage, _ = front.(*AIXMBasicMessage)
		}
		stage.OnAfterAIXMBasicMessageDeleteCallback.OnAfterDelete(stage, aixmbasicmessage, frontAIXMBasicMessage)
	}
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAIXMBasicMessageTypeCreateCallback != nil {
		stage.OnAfterAIXMBasicMessageTypeCreateCallback.OnAfterCreate(stage, aixmbasicmessagetype)
	}
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAIXMBasicMessageTypeUpdateCallback != nil {
		var frontAIXMBasicMessageType *AIXMBasicMessageType
		if front != nil {
			frontAIXMBasicMessageType, _ = front.(*AIXMBasicMessageType)
		}
		stage.OnAfterAIXMBasicMessageTypeUpdateCallback.OnAfterUpdate(stage, aixmbasicmessagetype, frontAIXMBasicMessageType)
	}
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAIXMBasicMessageTypeDeleteCallback != nil {
		var frontAIXMBasicMessageType *AIXMBasicMessageType
		if front != nil {
			frontAIXMBasicMessageType, _ = front.(*AIXMBasicMessageType)
		}
		stage.OnAfterAIXMBasicMessageTypeDeleteCallback.OnAfterDelete(stage, aixmbasicmessagetype, frontAIXMBasicMessageType)
	}
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBasicMessageMemberAIXMPropertyTypeCreateCallback != nil {
		stage.OnAfterBasicMessageMemberAIXMPropertyTypeCreateCallback.OnAfterCreate(stage, basicmessagememberaixmpropertytype)
	}
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBasicMessageMemberAIXMPropertyTypeUpdateCallback != nil {
		var frontBasicMessageMemberAIXMPropertyType *BasicMessageMemberAIXMPropertyType
		if front != nil {
			frontBasicMessageMemberAIXMPropertyType, _ = front.(*BasicMessageMemberAIXMPropertyType)
		}
		stage.OnAfterBasicMessageMemberAIXMPropertyTypeUpdateCallback.OnAfterUpdate(stage, basicmessagememberaixmpropertytype, frontBasicMessageMemberAIXMPropertyType)
	}
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBasicMessageMemberAIXMPropertyTypeDeleteCallback != nil {
		var frontBasicMessageMemberAIXMPropertyType *BasicMessageMemberAIXMPropertyType
		if front != nil {
			frontBasicMessageMemberAIXMPropertyType, _ = front.(*BasicMessageMemberAIXMPropertyType)
		}
		stage.OnAfterBasicMessageMemberAIXMPropertyTypeDeleteCallback.OnAfterDelete(stage, basicmessagememberaixmpropertytype, frontBasicMessageMemberAIXMPropertyType)
	}
}

