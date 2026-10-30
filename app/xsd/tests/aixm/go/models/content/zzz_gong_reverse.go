// generated code - do not edit
package content

// insertion point
func (inst *AbstractMD_ContentInformation_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Band_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_Band_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ContentInformation_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_CoverageContentTypeCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_CoverageDescription_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_CoverageDescription_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_FeatureCatalogueDescription_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_FeatureCatalogueDescription_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ImageDescription_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ImageDescription_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_ImagingConditionCode_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}

func (inst *MD_RangeDimension_PropertyType) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "MD_CoverageDescription_Type":
		switch reverseField.Fieldname {
		case "Dimension":
			if _md_coveragedescription_type, ok := stage.MD_CoverageDescription_Type_Dimension_reverseMap[inst]; ok {
				res = _md_coveragedescription_type.Name
			}
		}
	case "MD_ImageDescription_Type":
		switch reverseField.Fieldname {
		case "Dimension":
			if _md_imagedescription_type, ok := stage.MD_ImageDescription_Type_Dimension_reverseMap[inst]; ok {
				res = _md_imagedescription_type.Name
			}
		}
	}
	return
}

func (inst *MD_RangeDimension_Type) GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) (res string) {
	res = ""
	return
}
