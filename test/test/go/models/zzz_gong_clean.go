// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice []T) []T {
	if slice == nil {
		return nil
	}
    
	var cleanedSlice []T
	for _, element := range slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	return cleanedSlice
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element T) T {
	if !IsStagedPointerToGongstruct(stage, element) {
		var zero T
		return zero
	}
	return element
}

// Clean computes the reverse map, for all intances, for all clean to pointers field
func (stage *Stage) Clean() {
	// insertion point per named struct
	// clean up Astruct
	for astruct := range stage.Astructs {
		_ = astruct
		// insertion point per field
		astruct.Anarrayofb = GongCleanSlice(stage, astruct.Anarrayofb)
		astruct.Dstruct4s = GongCleanSlice(stage, astruct.Dstruct4s)
		astruct.Anarrayofa = GongCleanSlice(stage, astruct.Anarrayofa)
		astruct.Anotherarrayofb = GongCleanSlice(stage, astruct.Anotherarrayofb)
		astruct.AnarrayofbUse = GongCleanSlice(stage, astruct.AnarrayofbUse)
		astruct.Anarrayofb2Use = GongCleanSlice(stage, astruct.Anarrayofb2Use)
		// insertion point per field
		astruct.Associationtob = GongCleanPointer(stage, astruct.Associationtob)
		astruct.Anotherassociationtob_2 = GongCleanPointer(stage, astruct.Anotherassociationtob_2)
		astruct.Bstruct = GongCleanPointer(stage, astruct.Bstruct)
		astruct.Bstruct2 = GongCleanPointer(stage, astruct.Bstruct2)
		astruct.Dstruct = GongCleanPointer(stage, astruct.Dstruct)
		astruct.Dstruct2 = GongCleanPointer(stage, astruct.Dstruct2)
		astruct.Dstruct3 = GongCleanPointer(stage, astruct.Dstruct3)
		astruct.Dstruct4 = GongCleanPointer(stage, astruct.Dstruct4)
		astruct.AnAstruct = GongCleanPointer(stage, astruct.AnAstruct)
	}

	// clean up AstructBstruct2Use
	for astructbstruct2use := range stage.AstructBstruct2Uses {
		_ = astructbstruct2use
		// insertion point per field
		// insertion point per field
		astructbstruct2use.Bstrcut2 = GongCleanPointer(stage, astructbstruct2use.Bstrcut2)
	}

	// clean up AstructBstructUse
	for astructbstructuse := range stage.AstructBstructUses {
		_ = astructbstructuse
		// insertion point per field
		// insertion point per field
		astructbstructuse.Bstruct2 = GongCleanPointer(stage, astructbstructuse.Bstruct2)
	}

	// clean up Bstruct
	for bstruct := range stage.Bstructs {
		_ = bstruct
		// insertion point per field
		// insertion point per field
	}

	// clean up Dstruct
	for dstruct := range stage.Dstructs {
		_ = dstruct
		// insertion point per field
		dstruct.Anarrayofb = GongCleanSlice(stage, dstruct.Anarrayofb)
		dstruct.Gstructs = GongCleanSlice(stage, dstruct.Gstructs)
		// insertion point per field
		dstruct.Gstruct = GongCleanPointer(stage, dstruct.Gstruct)
	}

	// clean up F0123456789012345678901234567890
	for f0123456789012345678901234567890 := range stage.F0123456789012345678901234567890s {
		_ = f0123456789012345678901234567890
		// insertion point per field
		// insertion point per field
	}

	// clean up Gstruct
	for gstruct := range stage.Gstructs {
		_ = gstruct
		// insertion point per field
		// insertion point per field
	}

}
