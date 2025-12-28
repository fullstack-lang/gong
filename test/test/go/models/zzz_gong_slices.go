// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Astruct
	// insertion point per field
	stage.Astruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anarrayofb {
			stage.Astruct_Anarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	stage.Astruct_Dstruct4s_reverseMap = make(map[*Dstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _dstruct := range astruct.Dstruct4s {
			stage.Astruct_Dstruct4s_reverseMap[_dstruct] = astruct
		}
	}
	stage.Astruct_Anarrayofa_reverseMap = make(map[*Astruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astruct := range astruct.Anarrayofa {
			stage.Astruct_Anarrayofa_reverseMap[_astruct] = astruct
		}
	}
	stage.Astruct_Anotherarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anotherarrayofb {
			stage.Astruct_Anotherarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	stage.Astruct_AnarrayofbUse_reverseMap = make(map[*AstructBstructUse]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			stage.Astruct_AnarrayofbUse_reverseMap[_astructbstructuse] = astruct
		}
	}
	stage.Astruct_Anarrayofb2Use_reverseMap = make(map[*AstructBstruct2Use]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			stage.Astruct_Anarrayofb2Use_reverseMap[_astructbstruct2use] = astruct
		}
	}

	// Compute reverse map for named struct AstructBstruct2Use
	// insertion point per field

	// Compute reverse map for named struct AstructBstructUse
	// insertion point per field

	// Compute reverse map for named struct Bstruct
	// insertion point per field

	// Compute reverse map for named struct Dstruct
	// insertion point per field
	stage.Dstruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _bstruct := range dstruct.Anarrayofb {
			stage.Dstruct_Anarrayofb_reverseMap[_bstruct] = dstruct
		}
	}
	stage.Dstruct_Gstructs_reverseMap = make(map[*Gstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _gstruct := range dstruct.Gstructs {
			stage.Dstruct_Gstructs_reverseMap[_gstruct] = dstruct
		}
	}

	// Compute reverse map for named struct F0123456789012345678901234567890
	// insertion point per field

	// Compute reverse map for named struct Gstruct
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Astructs {
		res = append(res, instance)
	}

	for instance := range stage.AstructBstruct2Uses {
		res = append(res, instance)
	}

	for instance := range stage.AstructBstructUses {
		res = append(res, instance)
	}

	for instance := range stage.Bstructs {
		res = append(res, instance)
	}

	for instance := range stage.Dstructs {
		res = append(res, instance)
	}

	for instance := range stage.F0123456789012345678901234567890s {
		res = append(res, instance)
	}

	for instance := range stage.Gstructs {
		res = append(res, instance)
	}

	return
}


// insertion point per named struct
func (astruct *Astruct) GongCopy() GongstructIF {
	newInstance := *astruct
	return &newInstance
}

func (astructbstruct2use *AstructBstruct2Use) GongCopy() GongstructIF {
	newInstance := *astructbstruct2use
	return &newInstance
}

func (astructbstructuse *AstructBstructUse) GongCopy() GongstructIF {
	newInstance := *astructbstructuse
	return &newInstance
}

func (bstruct *Bstruct) GongCopy() GongstructIF {
	newInstance := *bstruct
	return &newInstance
}

func (dstruct *Dstruct) GongCopy() GongstructIF {
	newInstance := *dstruct
	return &newInstance
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongCopy() GongstructIF {
	newInstance := *f0123456789012345678901234567890
	return &newInstance
}

func (gstruct *Gstruct) GongCopy() GongstructIF {
	newInstance := *gstruct
	return &newInstance
}


func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenDeletedInstances int
	
	// insertion point per named struct
	var astructs_newInstances []*Astruct
	var astructs_deletedInstances []*Astruct

	// parse all staged instances and check if they have a reference
	for astruct := range stage.Astructs {
		if _, ok := stage.Astructs_reference[astruct]; !ok {
			astructs_newInstances = append(astructs_newInstances, astruct)
		}
	}

	// parse all reference instances and check if they are still staged
	for astruct := range stage.Astructs_reference {
		if _, ok := stage.Astructs[astruct]; !ok {
			astructs_deletedInstances = append(astructs_deletedInstances, astruct)
		}
	}

	lenNewInstances += len(astructs_newInstances)
	lenDeletedInstances += len(astructs_deletedInstances)
	var astructbstruct2uses_newInstances []*AstructBstruct2Use
	var astructbstruct2uses_deletedInstances []*AstructBstruct2Use

	// parse all staged instances and check if they have a reference
	for astructbstruct2use := range stage.AstructBstruct2Uses {
		if _, ok := stage.AstructBstruct2Uses_reference[astructbstruct2use]; !ok {
			astructbstruct2uses_newInstances = append(astructbstruct2uses_newInstances, astructbstruct2use)
		}
	}

	// parse all reference instances and check if they are still staged
	for astructbstruct2use := range stage.AstructBstruct2Uses_reference {
		if _, ok := stage.AstructBstruct2Uses[astructbstruct2use]; !ok {
			astructbstruct2uses_deletedInstances = append(astructbstruct2uses_deletedInstances, astructbstruct2use)
		}
	}

	lenNewInstances += len(astructbstruct2uses_newInstances)
	lenDeletedInstances += len(astructbstruct2uses_deletedInstances)
	var astructbstructuses_newInstances []*AstructBstructUse
	var astructbstructuses_deletedInstances []*AstructBstructUse

	// parse all staged instances and check if they have a reference
	for astructbstructuse := range stage.AstructBstructUses {
		if _, ok := stage.AstructBstructUses_reference[astructbstructuse]; !ok {
			astructbstructuses_newInstances = append(astructbstructuses_newInstances, astructbstructuse)
		}
	}

	// parse all reference instances and check if they are still staged
	for astructbstructuse := range stage.AstructBstructUses_reference {
		if _, ok := stage.AstructBstructUses[astructbstructuse]; !ok {
			astructbstructuses_deletedInstances = append(astructbstructuses_deletedInstances, astructbstructuse)
		}
	}

	lenNewInstances += len(astructbstructuses_newInstances)
	lenDeletedInstances += len(astructbstructuses_deletedInstances)
	var bstructs_newInstances []*Bstruct
	var bstructs_deletedInstances []*Bstruct

	// parse all staged instances and check if they have a reference
	for bstruct := range stage.Bstructs {
		if _, ok := stage.Bstructs_reference[bstruct]; !ok {
			bstructs_newInstances = append(bstructs_newInstances, bstruct)
		}
	}

	// parse all reference instances and check if they are still staged
	for bstruct := range stage.Bstructs_reference {
		if _, ok := stage.Bstructs[bstruct]; !ok {
			bstructs_deletedInstances = append(bstructs_deletedInstances, bstruct)
		}
	}

	lenNewInstances += len(bstructs_newInstances)
	lenDeletedInstances += len(bstructs_deletedInstances)
	var dstructs_newInstances []*Dstruct
	var dstructs_deletedInstances []*Dstruct

	// parse all staged instances and check if they have a reference
	for dstruct := range stage.Dstructs {
		if _, ok := stage.Dstructs_reference[dstruct]; !ok {
			dstructs_newInstances = append(dstructs_newInstances, dstruct)
		}
	}

	// parse all reference instances and check if they are still staged
	for dstruct := range stage.Dstructs_reference {
		if _, ok := stage.Dstructs[dstruct]; !ok {
			dstructs_deletedInstances = append(dstructs_deletedInstances, dstruct)
		}
	}

	lenNewInstances += len(dstructs_newInstances)
	lenDeletedInstances += len(dstructs_deletedInstances)
	var f0123456789012345678901234567890s_newInstances []*F0123456789012345678901234567890
	var f0123456789012345678901234567890s_deletedInstances []*F0123456789012345678901234567890

	// parse all staged instances and check if they have a reference
	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s {
		if _, ok := stage.F0123456789012345678901234567890s_reference[f0123456789012345678901234567890]; !ok {
			f0123456789012345678901234567890s_newInstances = append(f0123456789012345678901234567890s_newInstances, f0123456789012345678901234567890)
		}
	}

	// parse all reference instances and check if they are still staged
	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s_reference {
		if _, ok := stage.F0123456789012345678901234567890s[f0123456789012345678901234567890]; !ok {
			f0123456789012345678901234567890s_deletedInstances = append(f0123456789012345678901234567890s_deletedInstances, f0123456789012345678901234567890)
		}
	}

	lenNewInstances += len(f0123456789012345678901234567890s_newInstances)
	lenDeletedInstances += len(f0123456789012345678901234567890s_deletedInstances)
	var gstructs_newInstances []*Gstruct
	var gstructs_deletedInstances []*Gstruct

	// parse all staged instances and check if they have a reference
	for gstruct := range stage.Gstructs {
		if _, ok := stage.Gstructs_reference[gstruct]; !ok {
			gstructs_newInstances = append(gstructs_newInstances, gstruct)
		}
	}

	// parse all reference instances and check if they are still staged
	for gstruct := range stage.Gstructs_reference {
		if _, ok := stage.Gstructs[gstruct]; !ok {
			gstructs_deletedInstances = append(gstructs_deletedInstances, gstruct)
		}
	}

	lenNewInstances += len(gstructs_newInstances)
	lenDeletedInstances += len(gstructs_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Astructs_reference = make(map[*Astruct]*Astruct)
	for instance := range stage.Astructs {
		stage.Astructs_reference[instance] = instance
	}

	stage.AstructBstruct2Uses_reference = make(map[*AstructBstruct2Use]*AstructBstruct2Use)
	for instance := range stage.AstructBstruct2Uses {
		stage.AstructBstruct2Uses_reference[instance] = instance
	}

	stage.AstructBstructUses_reference = make(map[*AstructBstructUse]*AstructBstructUse)
	for instance := range stage.AstructBstructUses {
		stage.AstructBstructUses_reference[instance] = instance
	}

	stage.Bstructs_reference = make(map[*Bstruct]*Bstruct)
	for instance := range stage.Bstructs {
		stage.Bstructs_reference[instance] = instance
	}

	stage.Dstructs_reference = make(map[*Dstruct]*Dstruct)
	for instance := range stage.Dstructs {
		stage.Dstructs_reference[instance] = instance
	}

	stage.F0123456789012345678901234567890s_reference = make(map[*F0123456789012345678901234567890]*F0123456789012345678901234567890)
	for instance := range stage.F0123456789012345678901234567890s {
		stage.F0123456789012345678901234567890s_reference[instance] = instance
	}

	stage.Gstructs_reference = make(map[*Gstruct]*Gstruct)
	for instance := range stage.Gstructs {
		stage.Gstructs_reference[instance] = instance
	}

}
