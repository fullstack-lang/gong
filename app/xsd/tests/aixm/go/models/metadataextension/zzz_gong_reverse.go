// generated code - do not edit
package metadataextension

// insertion point
func (inst *MD_DatatypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ExtendedElementInformation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_MetadataExtensionInformation_Type":
		switch reverseField.Fieldname {
		case "ExtendedElementInformation":
			if _md_metadataextensioninformation_type, ok := stage.MD_MetadataExtensionInformation_Type_ExtendedElementInformation_reverseMap[inst]; ok {
				res = _md_metadataextensioninformation_type.Name
			}
		}
	}
	return
}

func (inst *MD_ExtendedElementInformation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_MetadataExtensionInformation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_MetadataExtensionInformation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ObligationCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
