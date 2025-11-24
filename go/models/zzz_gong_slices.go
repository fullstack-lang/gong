// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct GongBasicField
	// insertion point per field

	// Compute reverse map for named struct GongEnum
	// insertion point per field
	clear(stage.GongEnum_GongEnumValues_reverseMap)
	stage.GongEnum_GongEnumValues_reverseMap = make(map[*GongEnumValue]*GongEnum)
	for gongenum := range stage.GongEnums {
		_ = gongenum
		for _, _gongenumvalue := range gongenum.GongEnumValues {
			stage.GongEnum_GongEnumValues_reverseMap[_gongenumvalue] = gongenum
		}
	}

	// Compute reverse map for named struct GongEnumValue
	// insertion point per field

	// Compute reverse map for named struct GongLink
	// insertion point per field

	// Compute reverse map for named struct GongNote
	// insertion point per field
	clear(stage.GongNote_Links_reverseMap)
	stage.GongNote_Links_reverseMap = make(map[*GongLink]*GongNote)
	for gongnote := range stage.GongNotes {
		_ = gongnote
		for _, _gonglink := range gongnote.Links {
			stage.GongNote_Links_reverseMap[_gonglink] = gongnote
		}
	}

	// Compute reverse map for named struct GongStruct
	// insertion point per field
	clear(stage.GongStruct_GongBasicFields_reverseMap)
	stage.GongStruct_GongBasicFields_reverseMap = make(map[*GongBasicField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			stage.GongStruct_GongBasicFields_reverseMap[_gongbasicfield] = gongstruct
		}
	}
	clear(stage.GongStruct_GongTimeFields_reverseMap)
	stage.GongStruct_GongTimeFields_reverseMap = make(map[*GongTimeField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _gongtimefield := range gongstruct.GongTimeFields {
			stage.GongStruct_GongTimeFields_reverseMap[_gongtimefield] = gongstruct
		}
	}
	clear(stage.GongStruct_PointerToGongStructFields_reverseMap)
	stage.GongStruct_PointerToGongStructFields_reverseMap = make(map[*PointerToGongStructField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			stage.GongStruct_PointerToGongStructFields_reverseMap[_pointertogongstructfield] = gongstruct
		}
	}
	clear(stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap)
	stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap = make(map[*SliceOfPointerToGongStructField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
			stage.GongStruct_SliceOfPointerToGongStructFields_reverseMap[_sliceofpointertogongstructfield] = gongstruct
		}
	}

	// Compute reverse map for named struct GongTimeField
	// insertion point per field

	// Compute reverse map for named struct MetaReference
	// insertion point per field

	// Compute reverse map for named struct ModelPkg
	// insertion point per field

	// Compute reverse map for named struct PointerToGongStructField
	// insertion point per field

	// Compute reverse map for named struct SliceOfPointerToGongStructField
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.GongBasicFields {
		res = append(res, instance)
	}

	for instance := range stage.GongEnums {
		res = append(res, instance)
	}

	for instance := range stage.GongEnumValues {
		res = append(res, instance)
	}

	for instance := range stage.GongLinks {
		res = append(res, instance)
	}

	for instance := range stage.GongNotes {
		res = append(res, instance)
	}

	for instance := range stage.GongStructs {
		res = append(res, instance)
	}

	for instance := range stage.GongTimeFields {
		res = append(res, instance)
	}

	for instance := range stage.MetaReferences {
		res = append(res, instance)
	}

	for instance := range stage.ModelPkgs {
		res = append(res, instance)
	}

	for instance := range stage.PointerToGongStructFields {
		res = append(res, instance)
	}

	for instance := range stage.SliceOfPointerToGongStructFields {
		res = append(res, instance)
	}


	return
}
