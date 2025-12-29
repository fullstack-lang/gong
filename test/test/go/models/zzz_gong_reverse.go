// generated code - do not edit
package models

// insertion point
func (inst *Astruct) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Anarrayofa":
			if _astruct, ok := stage.Astruct_Anarrayofa_reverseMap[inst]; ok {
				res = _astruct.Name
			}
		}
	}
	return
}

func (inst *AstructBstruct2Use) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Anarrayofb2Use":
			if _astruct, ok := stage.Astruct_Anarrayofb2Use_reverseMap[inst]; ok {
				res = _astruct.Name
			}
		}
	}
	return
}

func (inst *AstructBstructUse) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "AnarrayofbUse":
			if _astruct, ok := stage.Astruct_AnarrayofbUse_reverseMap[inst]; ok {
				res = _astruct.Name
			}
		}
	}
	return
}

func (inst *Bstruct) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Anarrayofb":
			if _astruct, ok := stage.Astruct_Anarrayofb_reverseMap[inst]; ok {
				res = _astruct.Name
			}
		case "Anotherarrayofb":
			if _astruct, ok := stage.Astruct_Anotherarrayofb_reverseMap[inst]; ok {
				res = _astruct.Name
			}
		}
	case "Dstruct":
		switch reverseField.Fieldname {
		case "Anarrayofb":
			if _dstruct, ok := stage.Dstruct_Anarrayofb_reverseMap[inst]; ok {
				res = _dstruct.Name
			}
		}
	}
	return
}

func (inst *Dstruct) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Dstruct4s":
			if _astruct, ok := stage.Astruct_Dstruct4s_reverseMap[inst]; ok {
				res = _astruct.Name
			}
		}
	}
	return
}

func (inst *F0123456789012345678901234567890) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	}
	return
}

func (inst *Gstruct) GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) (res string) {

	res = ""
	switch reverseField.GongstructName {
	// insertion point
	case "Dstruct":
		switch reverseField.Fieldname {
		case "Gstructs":
			if _dstruct, ok := stage.Dstruct_Gstructs_reverseMap[inst]; ok {
				res = _dstruct.Name
			}
		}
	}
	return
}

// insertion point
func (inst *Astruct) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Anarrayofa":
			res = stage.Astruct_Anarrayofa_reverseMap[inst]
		}
	}
	return res
}

func (inst *AstructBstruct2Use) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Anarrayofb2Use":
			res = stage.Astruct_Anarrayofb2Use_reverseMap[inst]
		}
	}
	return res
}

func (inst *AstructBstructUse) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "AnarrayofbUse":
			res = stage.Astruct_AnarrayofbUse_reverseMap[inst]
		}
	}
	return res
}

func (inst *Bstruct) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Anarrayofb":
			res = stage.Astruct_Anarrayofb_reverseMap[inst]
		case "Anotherarrayofb":
			res = stage.Astruct_Anotherarrayofb_reverseMap[inst]
		}
	case "Dstruct":
		switch reverseField.Fieldname {
		case "Anarrayofb":
			res = stage.Dstruct_Anarrayofb_reverseMap[inst]
		}
	}
	return res
}

func (inst *Dstruct) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Astruct":
		switch reverseField.Fieldname {
		case "Dstruct4s":
			res = stage.Astruct_Dstruct4s_reverseMap[inst]
		}
	}
	return res
}

func (inst *F0123456789012345678901234567890) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	}
	return res
}

func (inst *Gstruct) GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) (res GongstructIF) {

	res = nil
	switch reverseField.GongstructName {
	// insertion point
	case "Dstruct":
		switch reverseField.Fieldname {
		case "Gstructs":
			res = stage.Dstruct_Gstructs_reverseMap[inst]
		}
	}
	return res
}
