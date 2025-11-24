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
// Clean garbage collect unstaged instances that are referenced by GongBasicField
func (gongbasicfield *GongBasicField) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	gongbasicfield.GongEnum = GongCleanPointer(stage, gongbasicfield.GongEnum)
}

// Clean garbage collect unstaged instances that are referenced by GongEnum
func (gongenum *GongEnum) GongClean(stage *Stage) {
	// insertion point per field
	gongenum.GongEnumValues = GongCleanSlice(stage, gongenum.GongEnumValues)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongEnumValue
func (gongenumvalue *GongEnumValue) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongLink
func (gonglink *GongLink) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongNote
func (gongnote *GongNote) GongClean(stage *Stage) {
	// insertion point per field
	gongnote.Links = GongCleanSlice(stage, gongnote.Links)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongStruct
func (gongstruct *GongStruct) GongClean(stage *Stage) {
	// insertion point per field
	gongstruct.GongBasicFields = GongCleanSlice(stage, gongstruct.GongBasicFields)
	gongstruct.GongTimeFields = GongCleanSlice(stage, gongstruct.GongTimeFields)
	gongstruct.PointerToGongStructFields = GongCleanSlice(stage, gongstruct.PointerToGongStructFields)
	gongstruct.SliceOfPointerToGongStructFields = GongCleanSlice(stage, gongstruct.SliceOfPointerToGongStructFields)
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by GongTimeField
func (gongtimefield *GongTimeField) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by MetaReference
func (metareference *MetaReference) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by ModelPkg
func (modelpkg *ModelPkg) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
}

// Clean garbage collect unstaged instances that are referenced by PointerToGongStructField
func (pointertogongstructfield *PointerToGongStructField) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	pointertogongstructfield.GongStruct = GongCleanPointer(stage, pointertogongstructfield.GongStruct)
}

// Clean garbage collect unstaged instances that are referenced by SliceOfPointerToGongStructField
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongClean(stage *Stage) {
	// insertion point per field
	// insertion point per field
	sliceofpointertogongstructfield.GongStruct = GongCleanPointer(stage, sliceofpointertogongstructfield.GongStruct)
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() {
	for _, instance := range stage.GetInstances() {
		instance.GongClean(stage)
	}
}
