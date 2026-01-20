// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	*slice = cleanedSlice
	return len(cleanedSlice) != len(*slice)
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Astruct
func (astruct *Astruct) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &astruct.Anarrayofb) || modified
	modified = GongCleanSlice(stage, &astruct.Dstruct4s) || modified
	modified = GongCleanSlice(stage, &astruct.Anarrayofa) || modified
	modified = GongCleanSlice(stage, &astruct.Anotherarrayofb) || modified
	modified = GongCleanSlice(stage, &astruct.AnarrayofbUse) || modified
	modified = GongCleanSlice(stage, &astruct.Anarrayofb2Use) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &astruct.Associationtob) || modified
	modified = GongCleanPointer(stage, &astruct.Anotherassociationtob_2) || modified
	modified = GongCleanPointer(stage, &astruct.Bstruct) || modified
	modified = GongCleanPointer(stage, &astruct.Bstruct2) || modified
	modified = GongCleanPointer(stage, &astruct.Dstruct) || modified
	modified = GongCleanPointer(stage, &astruct.Dstruct2) || modified
	modified = GongCleanPointer(stage, &astruct.Dstruct3) || modified
	modified = GongCleanPointer(stage, &astruct.Dstruct4) || modified
	modified = GongCleanPointer(stage, &astruct.AnAstruct) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AstructBstruct2Use
func (astructbstruct2use *AstructBstruct2Use) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &astructbstruct2use.Bstrcut2) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AstructBstructUse
func (astructbstructuse *AstructBstructUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &astructbstructuse.Bstruct2) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Bstruct
func (bstruct *Bstruct) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Dstruct
func (dstruct *Dstruct) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &dstruct.Anarrayofb) || modified
	modified = GongCleanSlice(stage, &dstruct.Gstructs) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &dstruct.Gstruct) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by F0123456789012345678901234567890
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Gstruct
func (gstruct *Gstruct) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
