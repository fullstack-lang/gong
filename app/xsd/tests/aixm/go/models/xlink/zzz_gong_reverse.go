// generated code - do not edit
package xlink

// insertion point
func (inst *ArcType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Extended) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *LocatorType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *ResourceType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *Simple) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *TitleEltType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "ArcType":
		switch reverseField.Fieldname {
		case "Title_1":
			if _arctype, ok := stage.ArcType_Title_1_reverseMap[inst]; ok {
				res = _arctype.Name
			}
		}
	case "LocatorType":
		switch reverseField.Fieldname {
		case "Title_1":
			if _locatortype, ok := stage.LocatorType_Title_1_reverseMap[inst]; ok {
				res = _locatortype.Name
			}
		}
	}
	return
}
