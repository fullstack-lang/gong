// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct Astruct
	for astruct := range stage.Astructs {
		_ = astruct
		// insertion point per field
		var _Anarrayofb []*Bstruct
		for _, _bstruct := range astruct.Anarrayofb {
			if IsStaged(stage, _bstruct) {
			 	_Anarrayofb = append(_Anarrayofb, _bstruct)
			}
		}
		astruct.Anarrayofb = _Anarrayofb
		var _Dstruct4s []*Dstruct
		for _, _dstruct := range astruct.Dstruct4s {
			if IsStaged(stage, _dstruct) {
			 	_Dstruct4s = append(_Dstruct4s, _dstruct)
			}
		}
		astruct.Dstruct4s = _Dstruct4s
		var _Anarrayofa []*Astruct
		for _, _astruct := range astruct.Anarrayofa {
			if IsStaged(stage, _astruct) {
			 	_Anarrayofa = append(_Anarrayofa, _astruct)
			}
		}
		astruct.Anarrayofa = _Anarrayofa
		var _Anotherarrayofb []*Bstruct
		for _, _bstruct := range astruct.Anotherarrayofb {
			if IsStaged(stage, _bstruct) {
			 	_Anotherarrayofb = append(_Anotherarrayofb, _bstruct)
			}
		}
		astruct.Anotherarrayofb = _Anotherarrayofb
		var _AnarrayofbUse []*AstructBstructUse
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			if IsStaged(stage, _astructbstructuse) {
			 	_AnarrayofbUse = append(_AnarrayofbUse, _astructbstructuse)
			}
		}
		astruct.AnarrayofbUse = _AnarrayofbUse
		var _Anarrayofb2Use []*AstructBstruct2Use
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			if IsStaged(stage, _astructbstruct2use) {
			 	_Anarrayofb2Use = append(_Anarrayofb2Use, _astructbstruct2use)
			}
		}
		astruct.Anarrayofb2Use = _Anarrayofb2Use
		// insertion point per field
		if !IsStaged(stage, astruct.Associationtob) {
			astruct.Associationtob = nil
		}
		if !IsStaged(stage, astruct.Anotherassociationtob_2) {
			astruct.Anotherassociationtob_2 = nil
		}
		if !IsStaged(stage, astruct.Bstruct) {
			astruct.Bstruct = nil
		}
		if !IsStaged(stage, astruct.Bstruct2) {
			astruct.Bstruct2 = nil
		}
		if !IsStaged(stage, astruct.Dstruct) {
			astruct.Dstruct = nil
		}
		if !IsStaged(stage, astruct.Dstruct2) {
			astruct.Dstruct2 = nil
		}
		if !IsStaged(stage, astruct.Dstruct3) {
			astruct.Dstruct3 = nil
		}
		if !IsStaged(stage, astruct.Dstruct4) {
			astruct.Dstruct4 = nil
		}
		if !IsStaged(stage, astruct.AnAstruct) {
			astruct.AnAstruct = nil
		}
	}

	// Compute reverse map for named struct AstructBstruct2Use
	for astructbstruct2use := range stage.AstructBstruct2Uses {
		_ = astructbstruct2use
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, astructbstruct2use.Bstrcut2) {
			astructbstruct2use.Bstrcut2 = nil
		}
	}

	// Compute reverse map for named struct AstructBstructUse
	for astructbstructuse := range stage.AstructBstructUses {
		_ = astructbstructuse
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, astructbstructuse.Bstruct2) {
			astructbstructuse.Bstruct2 = nil
		}
	}

	// Compute reverse map for named struct Bstruct
	for bstruct := range stage.Bstructs {
		_ = bstruct
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Dstruct
	for dstruct := range stage.Dstructs {
		_ = dstruct
		// insertion point per field
		var _Anarrayofb []*Bstruct
		for _, _bstruct := range dstruct.Anarrayofb {
			if IsStaged(stage, _bstruct) {
			 	_Anarrayofb = append(_Anarrayofb, _bstruct)
			}
		}
		dstruct.Anarrayofb = _Anarrayofb
		var _Gstructs []*Gstruct
		for _, _gstruct := range dstruct.Gstructs {
			if IsStaged(stage, _gstruct) {
			 	_Gstructs = append(_Gstructs, _gstruct)
			}
		}
		dstruct.Gstructs = _Gstructs
		// insertion point per field
		if !IsStaged(stage, dstruct.Gstruct) {
			dstruct.Gstruct = nil
		}
	}

	// Compute reverse map for named struct F0123456789012345678901234567890
	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s {
		_ = f0123456789012345678901234567890
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Gstruct
	for gstruct := range stage.Gstructs {
		_ = gstruct
		// insertion point per field
		// insertion point per field
	}

}
