// generated code - do not edit
package models

// insertion point
func (inst *AIXMBasicMessage) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *AIXMBasicMessageType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *BasicMessageMemberAIXMPropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "AIXMBasicMessage":
		switch reverseField.Fieldname {
		case "HasMember":
			if _aixmbasicmessage, ok := stage.AIXMBasicMessage_HasMember_reverseMap[inst]; ok {
				res = _aixmbasicmessage.Name
			}
		}
	case "AIXMBasicMessageType":
		switch reverseField.Fieldname {
		case "HasMember":
			if _aixmbasicmessagetype, ok := stage.AIXMBasicMessageType_HasMember_reverseMap[inst]; ok {
				res = _aixmbasicmessagetype.Name
			}
		}
	}
	return
}
