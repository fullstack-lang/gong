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
	// clean up GongBasicField
	for gongbasicfield := range stage.GongBasicFields {
		_ = gongbasicfield
		// insertion point per field
		// insertion point per field
		gongbasicfield.GongEnum = GongCleanPointer(stage, gongbasicfield.GongEnum)
	}

	// clean up GongEnum
	for gongenum := range stage.GongEnums {
		_ = gongenum
		// insertion point per field
		gongenum.GongEnumValues = GongCleanSlice(stage, gongenum.GongEnumValues)
		// insertion point per field
	}

	// clean up GongEnumValue
	for gongenumvalue := range stage.GongEnumValues {
		_ = gongenumvalue
		// insertion point per field
		// insertion point per field
	}

	// clean up GongLink
	for gonglink := range stage.GongLinks {
		_ = gonglink
		// insertion point per field
		// insertion point per field
	}

	// clean up GongNote
	for gongnote := range stage.GongNotes {
		_ = gongnote
		// insertion point per field
		gongnote.Links = GongCleanSlice(stage, gongnote.Links)
		// insertion point per field
	}

	// clean up GongStruct
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		// insertion point per field
		gongstruct.GongBasicFields = GongCleanSlice(stage, gongstruct.GongBasicFields)
		gongstruct.GongTimeFields = GongCleanSlice(stage, gongstruct.GongTimeFields)
		gongstruct.PointerToGongStructFields = GongCleanSlice(stage, gongstruct.PointerToGongStructFields)
		gongstruct.SliceOfPointerToGongStructFields = GongCleanSlice(stage, gongstruct.SliceOfPointerToGongStructFields)
		// insertion point per field
	}

	// clean up GongTimeField
	for gongtimefield := range stage.GongTimeFields {
		_ = gongtimefield
		// insertion point per field
		// insertion point per field
	}

	// clean up MetaReference
	for metareference := range stage.MetaReferences {
		_ = metareference
		// insertion point per field
		// insertion point per field
	}

	// clean up ModelPkg
	for modelpkg := range stage.ModelPkgs {
		_ = modelpkg
		// insertion point per field
		// insertion point per field
	}

	// clean up PointerToGongStructField
	for pointertogongstructfield := range stage.PointerToGongStructFields {
		_ = pointertogongstructfield
		// insertion point per field
		// insertion point per field
		pointertogongstructfield.GongStruct = GongCleanPointer(stage, pointertogongstructfield.GongStruct)
	}

	// clean up SliceOfPointerToGongStructField
	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		_ = sliceofpointertogongstructfield
		// insertion point per field
		// insertion point per field
		sliceofpointertogongstructfield.GongStruct = GongCleanPointer(stage, sliceofpointertogongstructfield.GongStruct)
	}

}
