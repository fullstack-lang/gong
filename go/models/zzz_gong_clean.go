// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct GongBasicField
	for gongbasicfield := range stage.GongBasicFields {
		_ = gongbasicfield
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, gongbasicfield.GongEnum) {
			gongbasicfield.GongEnum = nil
		}
	}

	// Compute reverse map for named struct GongEnum
	for gongenum := range stage.GongEnums {
		_ = gongenum
		// insertion point per field
		var _GongEnumValues []*GongEnumValue
		for _, _gongenumvalue := range gongenum.GongEnumValues {
			if IsStaged(stage, _gongenumvalue) {
			 	_GongEnumValues = append(_GongEnumValues, _gongenumvalue)
			}
		}
		gongenum.GongEnumValues = _GongEnumValues
		// insertion point per field
	}

	// Compute reverse map for named struct GongEnumValue
	for gongenumvalue := range stage.GongEnumValues {
		_ = gongenumvalue
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct GongLink
	for gonglink := range stage.GongLinks {
		_ = gonglink
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct GongNote
	for gongnote := range stage.GongNotes {
		_ = gongnote
		// insertion point per field
		var _Links []*GongLink
		for _, _gonglink := range gongnote.Links {
			if IsStaged(stage, _gonglink) {
			 	_Links = append(_Links, _gonglink)
			}
		}
		gongnote.Links = _Links
		// insertion point per field
	}

	// Compute reverse map for named struct GongStruct
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		// insertion point per field
		var _GongBasicFields []*GongBasicField
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			if IsStaged(stage, _gongbasicfield) {
			 	_GongBasicFields = append(_GongBasicFields, _gongbasicfield)
			}
		}
		gongstruct.GongBasicFields = _GongBasicFields
		var _GongTimeFields []*GongTimeField
		for _, _gongtimefield := range gongstruct.GongTimeFields {
			if IsStaged(stage, _gongtimefield) {
			 	_GongTimeFields = append(_GongTimeFields, _gongtimefield)
			}
		}
		gongstruct.GongTimeFields = _GongTimeFields
		var _PointerToGongStructFields []*PointerToGongStructField
		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			if IsStaged(stage, _pointertogongstructfield) {
			 	_PointerToGongStructFields = append(_PointerToGongStructFields, _pointertogongstructfield)
			}
		}
		gongstruct.PointerToGongStructFields = _PointerToGongStructFields
		var _SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField
		for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
			if IsStaged(stage, _sliceofpointertogongstructfield) {
			 	_SliceOfPointerToGongStructFields = append(_SliceOfPointerToGongStructFields, _sliceofpointertogongstructfield)
			}
		}
		gongstruct.SliceOfPointerToGongStructFields = _SliceOfPointerToGongStructFields
		// insertion point per field
	}

	// Compute reverse map for named struct GongTimeField
	for gongtimefield := range stage.GongTimeFields {
		_ = gongtimefield
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct MetaReference
	for metareference := range stage.MetaReferences {
		_ = metareference
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct ModelPkg
	for modelpkg := range stage.ModelPkgs {
		_ = modelpkg
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct PointerToGongStructField
	for pointertogongstructfield := range stage.PointerToGongStructFields {
		_ = pointertogongstructfield
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, pointertogongstructfield.GongStruct) {
			pointertogongstructfield.GongStruct = nil
		}
	}

	// Compute reverse map for named struct SliceOfPointerToGongStructField
	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		_ = sliceofpointertogongstructfield
		// insertion point per field
		// insertion point per field
		if !IsStaged(stage, sliceofpointertogongstructfield.GongStruct) {
			sliceofpointertogongstructfield.GongStruct = nil
		}
	}

}
