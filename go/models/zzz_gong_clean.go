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

// GongCleanSlice is a backward-compatible forwarder to stage.CleanSlice.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	return stage.CleanSlice(slice)
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

// GongCleanPointer is a backward-compatible forwarder to stage.CleanPointer.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	return stage.CleanPointer(element)
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by GongBasicField
func (gongbasicfield *GongBasicField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &gongbasicfield.GongEnum) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GongEnum
func (gongenum *GongEnum) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongenum.GongEnumValues) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongEnumValue
func (gongenumvalue *GongEnumValue) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongLink
func (gonglink *GongLink) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongNote
func (gongnote *GongNote) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongnote.Links) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongStruct
func (gongstruct *GongStruct) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongstruct.GongBasicFields) || modified
	modified = GongCleanSlice(stage, &gongstruct.GongTimeFields) || modified
	modified = GongCleanSlice(stage, &gongstruct.PointerToGongStructFields) || modified
	modified = GongCleanSlice(stage, &gongstruct.SliceOfPointerToGongStructFields) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongTimeField
func (gongtimefield *GongTimeField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by MetaReference
func (metareference *MetaReference) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ModelPkg
func (modelpkg *ModelPkg) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PointerToGongStructField
func (pointertogongstructfield *PointerToGongStructField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &pointertogongstructfield.GongStruct) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SliceOfPointerToGongStructField
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &sliceofpointertogongstructfield.GongStruct) || modified
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
