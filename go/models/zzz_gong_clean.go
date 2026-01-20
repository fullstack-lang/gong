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
// Clean garbage collect unstaged instances that are referenced by GongBasicField
func (gongbasicfield *GongBasicField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &gongbasicfield.GongEnum)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by GongEnum
func (gongenum *GongEnum) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongenum.GongEnumValues)  || modified
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
	modified = GongCleanSlice(stage, &gongnote.Links)  || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by GongStruct
func (gongstruct *GongStruct) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &gongstruct.GongBasicFields)  || modified
	modified = GongCleanSlice(stage, &gongstruct.GongTimeFields)  || modified
	modified = GongCleanSlice(stage, &gongstruct.PointerToGongStructFields)  || modified
	modified = GongCleanSlice(stage, &gongstruct.SliceOfPointerToGongStructFields)  || modified
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
	modified = GongCleanPointer(stage, &pointertogongstructfield.GongStruct)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SliceOfPointerToGongStructField
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &sliceofpointertogongstructfield.GongStruct)  || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
