// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Astruct
	// insertion point per field
	clear(stage.Astruct_Anarrayofb_reverseMap)
	stage.Astruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anarrayofb {
			stage.Astruct_Anarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	clear(stage.Astruct_Dstruct4s_reverseMap)
	stage.Astruct_Dstruct4s_reverseMap = make(map[*Dstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _dstruct := range astruct.Dstruct4s {
			stage.Astruct_Dstruct4s_reverseMap[_dstruct] = astruct
		}
	}
	clear(stage.Astruct_Anarrayofa_reverseMap)
	stage.Astruct_Anarrayofa_reverseMap = make(map[*Astruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astruct := range astruct.Anarrayofa {
			stage.Astruct_Anarrayofa_reverseMap[_astruct] = astruct
		}
	}
	clear(stage.Astruct_Anotherarrayofb_reverseMap)
	stage.Astruct_Anotherarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anotherarrayofb {
			stage.Astruct_Anotherarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	clear(stage.Astruct_AnarrayofbUse_reverseMap)
	stage.Astruct_AnarrayofbUse_reverseMap = make(map[*AstructBstructUse]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			stage.Astruct_AnarrayofbUse_reverseMap[_astructbstructuse] = astruct
		}
	}
	clear(stage.Astruct_Anarrayofb2Use_reverseMap)
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
	clear(stage.Dstruct_Anarrayofb_reverseMap)
	stage.Dstruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _bstruct := range dstruct.Anarrayofb {
			stage.Dstruct_Anarrayofb_reverseMap[_bstruct] = dstruct
		}
	}
	clear(stage.Dstruct_Gstructs_reverseMap)
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
