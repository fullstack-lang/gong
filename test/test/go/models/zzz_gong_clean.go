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

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Astruct
func (astruct *Astruct) GongClean(stage *Stage) {
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

// Clean garbage collect unstaged instances that are referenced by AstructBstruct2Use
func (astructbstruct2use *AstructBstruct2Use) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	astructbstruct2use.Bstrcut2 = GongCleanPointer(stage, astructbstruct2use.Bstrcut2)
}

// Clean garbage collect unstaged instances that are referenced by AstructBstructUse
func (astructbstructuse *AstructBstructUse) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	astructbstructuse.Bstruct2 = GongCleanPointer(stage, astructbstructuse.Bstruct2)
}

// Clean garbage collect unstaged instances that are referenced by Bstruct
func (bstruct *Bstruct) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Dstruct
func (dstruct *Dstruct) GongClean(stage *Stage) {
	// insertion point per field
	dstruct.Anarrayofb = GongCleanSlice(stage, dstruct.Anarrayofb)
	dstruct.Gstructs = GongCleanSlice(stage, dstruct.Gstructs)
	// insertion point per field
	dstruct.Gstruct = GongCleanPointer(stage, dstruct.Gstruct)
}

// Clean garbage collect unstaged instances that are referenced by F0123456789012345678901234567890
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by Gstruct
func (gstruct *Gstruct) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
