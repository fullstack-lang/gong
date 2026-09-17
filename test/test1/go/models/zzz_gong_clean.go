// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T PointerToGongstruct](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T PointerToGongstruct](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by Astruct
func (astruct *Astruct) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&astruct.Anarrayofb) || modified
	modified = stage.CleanSlice(&astruct.Dstruct4s) || modified
	modified = stage.CleanSlice(&astruct.Anarrayofa) || modified
	modified = stage.CleanSlice(&astruct.Anotherarrayofb) || modified
	modified = stage.CleanSlice(&astruct.AnarrayofbUse) || modified
	modified = stage.CleanSlice(&astruct.Anarrayofb2Use) || modified
	// insertion point per field
	modified = stage.CleanPointer(&astruct.Associationtob) || modified
	modified = stage.CleanPointer(&astruct.Anotherassociationtob_2) || modified
	modified = stage.CleanPointer(&astruct.Bstruct) || modified
	modified = stage.CleanPointer(&astruct.Bstruct2) || modified
	modified = stage.CleanPointer(&astruct.Dstruct) || modified
	modified = stage.CleanPointer(&astruct.Dstruct2) || modified
	modified = stage.CleanPointer(&astruct.Dstruct3) || modified
	modified = stage.CleanPointer(&astruct.Dstruct4) || modified
	modified = stage.CleanPointer(&astruct.AnAstruct) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AstructBstruct2Use
func (astructbstruct2use *AstructBstruct2Use) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&astructbstruct2use.Bstrcut2) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AstructBstructUse
func (astructbstructuse *AstructBstructUse) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&astructbstructuse.Bstruct2) || modified
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
	modified = stage.CleanSlice(&dstruct.Anarrayofb) || modified
	modified = stage.CleanSlice(&dstruct.Gstructs) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dstruct.Gstruct) || modified
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
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
