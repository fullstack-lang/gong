// generated code - do not edit
package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct GongBasicField
	// insertion point per field

	// Compute reverse map for named struct GongEnum
	// insertion point per field
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
	stage.GongNote_Links_reverseMap = make(map[*GongLink]*GongNote)
	for gongnote := range stage.GongNotes {
		_ = gongnote
		for _, _gonglink := range gongnote.Links {
			stage.GongNote_Links_reverseMap[_gonglink] = gongnote
		}
	}

	// Compute reverse map for named struct GongStruct
	// insertion point per field
	stage.GongStruct_GongBasicFields_reverseMap = make(map[*GongBasicField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _gongbasicfield := range gongstruct.GongBasicFields {
			stage.GongStruct_GongBasicFields_reverseMap[_gongbasicfield] = gongstruct
		}
	}
	stage.GongStruct_GongTimeFields_reverseMap = make(map[*GongTimeField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _gongtimefield := range gongstruct.GongTimeFields {
			stage.GongStruct_GongTimeFields_reverseMap[_gongtimefield] = gongstruct
		}
	}
	stage.GongStruct_PointerToGongStructFields_reverseMap = make(map[*PointerToGongStructField]*GongStruct)
	for gongstruct := range stage.GongStructs {
		_ = gongstruct
		for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
			stage.GongStruct_PointerToGongStructFields_reverseMap[_pointertogongstructfield] = gongstruct
		}
	}
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

	// end of insertion point per named struct
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

// insertion point per named struct
func (gongbasicfield *GongBasicField) GongCopy() GongstructIF {
	newInstance := *gongbasicfield
	return &newInstance
}

func (gongenum *GongEnum) GongCopy() GongstructIF {
	newInstance := *gongenum
	return &newInstance
}

func (gongenumvalue *GongEnumValue) GongCopy() GongstructIF {
	newInstance := *gongenumvalue
	return &newInstance
}

func (gonglink *GongLink) GongCopy() GongstructIF {
	newInstance := *gonglink
	return &newInstance
}

func (gongnote *GongNote) GongCopy() GongstructIF {
	newInstance := *gongnote
	return &newInstance
}

func (gongstruct *GongStruct) GongCopy() GongstructIF {
	newInstance := *gongstruct
	return &newInstance
}

func (gongtimefield *GongTimeField) GongCopy() GongstructIF {
	newInstance := *gongtimefield
	return &newInstance
}

func (metareference *MetaReference) GongCopy() GongstructIF {
	newInstance := *metareference
	return &newInstance
}

func (modelpkg *ModelPkg) GongCopy() GongstructIF {
	newInstance := *modelpkg
	return &newInstance
}

func (pointertogongstructfield *PointerToGongStructField) GongCopy() GongstructIF {
	newInstance := *pointertogongstructfield
	return &newInstance
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongCopy() GongstructIF {
	newInstance := *sliceofpointertogongstructfield
	return &newInstance
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var gongbasicfields_newInstances []*GongBasicField
	var gongbasicfields_deletedInstances []*GongBasicField

	// parse all staged instances and check if they have a reference
	for gongbasicfield := range stage.GongBasicFields {
		if ref, ok := stage.GongBasicFields_reference[gongbasicfield]; !ok {
			gongbasicfields_newInstances = append(gongbasicfields_newInstances, gongbasicfield)
			newInstancesSlice = append(newInstancesSlice, gongbasicfield.GongMarshallIdentifier(stage))
			if stage.GongBasicFields_referenceOrder == nil {
				stage.GongBasicFields_referenceOrder = make(map[*GongBasicField]uint)
			}
			stage.GongBasicFields_referenceOrder[gongbasicfield] = stage.GongBasicFieldMap_Staged_Order[gongbasicfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongbasicfield.GongMarshallUnstaging(stage))
			delete(stage.GongBasicFields_referenceOrder, gongbasicfield)
			fieldInitializers, pointersInitializations := gongbasicfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongBasicFieldMap_Staged_Order[ref] = stage.GongBasicFieldMap_Staged_Order[gongbasicfield]
			diffs := gongbasicfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongbasicfield)
			delete(stage.GongBasicFieldMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongbasicfield.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.GongBasicFields_reference {
		if _, ok := stage.GongBasicFields[ref]; !ok {
			gongbasicfields_deletedInstances = append(gongbasicfields_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongbasicfields_newInstances)
	lenDeletedInstances += len(gongbasicfields_deletedInstances)
	var gongenums_newInstances []*GongEnum
	var gongenums_deletedInstances []*GongEnum

	// parse all staged instances and check if they have a reference
	for gongenum := range stage.GongEnums {
		if ref, ok := stage.GongEnums_reference[gongenum]; !ok {
			gongenums_newInstances = append(gongenums_newInstances, gongenum)
			newInstancesSlice = append(newInstancesSlice, gongenum.GongMarshallIdentifier(stage))
			if stage.GongEnums_referenceOrder == nil {
				stage.GongEnums_referenceOrder = make(map[*GongEnum]uint)
			}
			stage.GongEnums_referenceOrder[gongenum] = stage.GongEnumMap_Staged_Order[gongenum]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongenum.GongMarshallUnstaging(stage))
			delete(stage.GongEnums_referenceOrder, gongenum)
			fieldInitializers, pointersInitializations := gongenum.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongEnumMap_Staged_Order[ref] = stage.GongEnumMap_Staged_Order[gongenum]
			diffs := gongenum.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongenum)
			delete(stage.GongEnumMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongenum.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.GongEnums_reference {
		if _, ok := stage.GongEnums[ref]; !ok {
			gongenums_deletedInstances = append(gongenums_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongenums_newInstances)
	lenDeletedInstances += len(gongenums_deletedInstances)
	var gongenumvalues_newInstances []*GongEnumValue
	var gongenumvalues_deletedInstances []*GongEnumValue

	// parse all staged instances and check if they have a reference
	for gongenumvalue := range stage.GongEnumValues {
		if ref, ok := stage.GongEnumValues_reference[gongenumvalue]; !ok {
			gongenumvalues_newInstances = append(gongenumvalues_newInstances, gongenumvalue)
			newInstancesSlice = append(newInstancesSlice, gongenumvalue.GongMarshallIdentifier(stage))
			if stage.GongEnumValues_referenceOrder == nil {
				stage.GongEnumValues_referenceOrder = make(map[*GongEnumValue]uint)
			}
			stage.GongEnumValues_referenceOrder[gongenumvalue] = stage.GongEnumValueMap_Staged_Order[gongenumvalue]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongenumvalue.GongMarshallUnstaging(stage))
			delete(stage.GongEnumValues_referenceOrder, gongenumvalue)
			fieldInitializers, pointersInitializations := gongenumvalue.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongEnumValueMap_Staged_Order[ref] = stage.GongEnumValueMap_Staged_Order[gongenumvalue]
			diffs := gongenumvalue.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongenumvalue)
			delete(stage.GongEnumValueMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongenumvalue.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.GongEnumValues_reference {
		if _, ok := stage.GongEnumValues[ref]; !ok {
			gongenumvalues_deletedInstances = append(gongenumvalues_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongenumvalues_newInstances)
	lenDeletedInstances += len(gongenumvalues_deletedInstances)
	var gonglinks_newInstances []*GongLink
	var gonglinks_deletedInstances []*GongLink

	// parse all staged instances and check if they have a reference
	for gonglink := range stage.GongLinks {
		if ref, ok := stage.GongLinks_reference[gonglink]; !ok {
			gonglinks_newInstances = append(gonglinks_newInstances, gonglink)
			newInstancesSlice = append(newInstancesSlice, gonglink.GongMarshallIdentifier(stage))
			if stage.GongLinks_referenceOrder == nil {
				stage.GongLinks_referenceOrder = make(map[*GongLink]uint)
			}
			stage.GongLinks_referenceOrder[gonglink] = stage.GongLinkMap_Staged_Order[gonglink]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gonglink.GongMarshallUnstaging(stage))
			delete(stage.GongLinks_referenceOrder, gonglink)
			fieldInitializers, pointersInitializations := gonglink.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongLinkMap_Staged_Order[ref] = stage.GongLinkMap_Staged_Order[gonglink]
			diffs := gonglink.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gonglink)
			delete(stage.GongLinkMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gonglink.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.GongLinks_reference {
		if _, ok := stage.GongLinks[ref]; !ok {
			gonglinks_deletedInstances = append(gonglinks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gonglinks_newInstances)
	lenDeletedInstances += len(gonglinks_deletedInstances)
	var gongnotes_newInstances []*GongNote
	var gongnotes_deletedInstances []*GongNote

	// parse all staged instances and check if they have a reference
	for gongnote := range stage.GongNotes {
		if ref, ok := stage.GongNotes_reference[gongnote]; !ok {
			gongnotes_newInstances = append(gongnotes_newInstances, gongnote)
			newInstancesSlice = append(newInstancesSlice, gongnote.GongMarshallIdentifier(stage))
			if stage.GongNotes_referenceOrder == nil {
				stage.GongNotes_referenceOrder = make(map[*GongNote]uint)
			}
			stage.GongNotes_referenceOrder[gongnote] = stage.GongNoteMap_Staged_Order[gongnote]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongnote.GongMarshallUnstaging(stage))
			delete(stage.GongNotes_referenceOrder, gongnote)
			fieldInitializers, pointersInitializations := gongnote.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongNoteMap_Staged_Order[ref] = stage.GongNoteMap_Staged_Order[gongnote]
			diffs := gongnote.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongnote)
			delete(stage.GongNoteMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongnote.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.GongNotes_reference {
		if _, ok := stage.GongNotes[ref]; !ok {
			gongnotes_deletedInstances = append(gongnotes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongnotes_newInstances)
	lenDeletedInstances += len(gongnotes_deletedInstances)
	var gongstructs_newInstances []*GongStruct
	var gongstructs_deletedInstances []*GongStruct

	// parse all staged instances and check if they have a reference
	for gongstruct := range stage.GongStructs {
		if ref, ok := stage.GongStructs_reference[gongstruct]; !ok {
			gongstructs_newInstances = append(gongstructs_newInstances, gongstruct)
			newInstancesSlice = append(newInstancesSlice, gongstruct.GongMarshallIdentifier(stage))
			if stage.GongStructs_referenceOrder == nil {
				stage.GongStructs_referenceOrder = make(map[*GongStruct]uint)
			}
			stage.GongStructs_referenceOrder[gongstruct] = stage.GongStructMap_Staged_Order[gongstruct]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongstruct.GongMarshallUnstaging(stage))
			delete(stage.GongStructs_referenceOrder, gongstruct)
			fieldInitializers, pointersInitializations := gongstruct.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongStructMap_Staged_Order[ref] = stage.GongStructMap_Staged_Order[gongstruct]
			diffs := gongstruct.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongstruct)
			delete(stage.GongStructMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongstruct.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.GongStructs_reference {
		if _, ok := stage.GongStructs[ref]; !ok {
			gongstructs_deletedInstances = append(gongstructs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongstructs_newInstances)
	lenDeletedInstances += len(gongstructs_deletedInstances)
	var gongtimefields_newInstances []*GongTimeField
	var gongtimefields_deletedInstances []*GongTimeField

	// parse all staged instances and check if they have a reference
	for gongtimefield := range stage.GongTimeFields {
		if ref, ok := stage.GongTimeFields_reference[gongtimefield]; !ok {
			gongtimefields_newInstances = append(gongtimefields_newInstances, gongtimefield)
			newInstancesSlice = append(newInstancesSlice, gongtimefield.GongMarshallIdentifier(stage))
			if stage.GongTimeFields_referenceOrder == nil {
				stage.GongTimeFields_referenceOrder = make(map[*GongTimeField]uint)
			}
			stage.GongTimeFields_referenceOrder[gongtimefield] = stage.GongTimeFieldMap_Staged_Order[gongtimefield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongtimefield.GongMarshallUnstaging(stage))
			delete(stage.GongTimeFields_referenceOrder, gongtimefield)
			fieldInitializers, pointersInitializations := gongtimefield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongTimeFieldMap_Staged_Order[ref] = stage.GongTimeFieldMap_Staged_Order[gongtimefield]
			diffs := gongtimefield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongtimefield)
			delete(stage.GongTimeFieldMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongtimefield.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.GongTimeFields_reference {
		if _, ok := stage.GongTimeFields[ref]; !ok {
			gongtimefields_deletedInstances = append(gongtimefields_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongtimefields_newInstances)
	lenDeletedInstances += len(gongtimefields_deletedInstances)
	var metareferences_newInstances []*MetaReference
	var metareferences_deletedInstances []*MetaReference

	// parse all staged instances and check if they have a reference
	for metareference := range stage.MetaReferences {
		if ref, ok := stage.MetaReferences_reference[metareference]; !ok {
			metareferences_newInstances = append(metareferences_newInstances, metareference)
			newInstancesSlice = append(newInstancesSlice, metareference.GongMarshallIdentifier(stage))
			if stage.MetaReferences_referenceOrder == nil {
				stage.MetaReferences_referenceOrder = make(map[*MetaReference]uint)
			}
			stage.MetaReferences_referenceOrder[metareference] = stage.MetaReferenceMap_Staged_Order[metareference]
			newInstancesReverseSlice = append(newInstancesReverseSlice, metareference.GongMarshallUnstaging(stage))
			delete(stage.MetaReferences_referenceOrder, metareference)
			fieldInitializers, pointersInitializations := metareference.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MetaReferenceMap_Staged_Order[ref] = stage.MetaReferenceMap_Staged_Order[metareference]
			diffs := metareference.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, metareference)
			delete(stage.MetaReferenceMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", metareference.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.MetaReferences_reference {
		if _, ok := stage.MetaReferences[ref]; !ok {
			metareferences_deletedInstances = append(metareferences_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(metareferences_newInstances)
	lenDeletedInstances += len(metareferences_deletedInstances)
	var modelpkgs_newInstances []*ModelPkg
	var modelpkgs_deletedInstances []*ModelPkg

	// parse all staged instances and check if they have a reference
	for modelpkg := range stage.ModelPkgs {
		if ref, ok := stage.ModelPkgs_reference[modelpkg]; !ok {
			modelpkgs_newInstances = append(modelpkgs_newInstances, modelpkg)
			newInstancesSlice = append(newInstancesSlice, modelpkg.GongMarshallIdentifier(stage))
			if stage.ModelPkgs_referenceOrder == nil {
				stage.ModelPkgs_referenceOrder = make(map[*ModelPkg]uint)
			}
			stage.ModelPkgs_referenceOrder[modelpkg] = stage.ModelPkgMap_Staged_Order[modelpkg]
			newInstancesReverseSlice = append(newInstancesReverseSlice, modelpkg.GongMarshallUnstaging(stage))
			delete(stage.ModelPkgs_referenceOrder, modelpkg)
			fieldInitializers, pointersInitializations := modelpkg.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ModelPkgMap_Staged_Order[ref] = stage.ModelPkgMap_Staged_Order[modelpkg]
			diffs := modelpkg.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, modelpkg)
			delete(stage.ModelPkgMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", modelpkg.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ModelPkgs_reference {
		if _, ok := stage.ModelPkgs[ref]; !ok {
			modelpkgs_deletedInstances = append(modelpkgs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(modelpkgs_newInstances)
	lenDeletedInstances += len(modelpkgs_deletedInstances)
	var pointertogongstructfields_newInstances []*PointerToGongStructField
	var pointertogongstructfields_deletedInstances []*PointerToGongStructField

	// parse all staged instances and check if they have a reference
	for pointertogongstructfield := range stage.PointerToGongStructFields {
		if ref, ok := stage.PointerToGongStructFields_reference[pointertogongstructfield]; !ok {
			pointertogongstructfields_newInstances = append(pointertogongstructfields_newInstances, pointertogongstructfield)
			newInstancesSlice = append(newInstancesSlice, pointertogongstructfield.GongMarshallIdentifier(stage))
			if stage.PointerToGongStructFields_referenceOrder == nil {
				stage.PointerToGongStructFields_referenceOrder = make(map[*PointerToGongStructField]uint)
			}
			stage.PointerToGongStructFields_referenceOrder[pointertogongstructfield] = stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, pointertogongstructfield.GongMarshallUnstaging(stage))
			delete(stage.PointerToGongStructFields_referenceOrder, pointertogongstructfield)
			fieldInitializers, pointersInitializations := pointertogongstructfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PointerToGongStructFieldMap_Staged_Order[ref] = stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfield]
			diffs := pointertogongstructfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, pointertogongstructfield)
			delete(stage.PointerToGongStructFieldMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", pointertogongstructfield.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.PointerToGongStructFields_reference {
		if _, ok := stage.PointerToGongStructFields[ref]; !ok {
			pointertogongstructfields_deletedInstances = append(pointertogongstructfields_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(pointertogongstructfields_newInstances)
	lenDeletedInstances += len(pointertogongstructfields_deletedInstances)
	var sliceofpointertogongstructfields_newInstances []*SliceOfPointerToGongStructField
	var sliceofpointertogongstructfields_deletedInstances []*SliceOfPointerToGongStructField

	// parse all staged instances and check if they have a reference
	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields {
		if ref, ok := stage.SliceOfPointerToGongStructFields_reference[sliceofpointertogongstructfield]; !ok {
			sliceofpointertogongstructfields_newInstances = append(sliceofpointertogongstructfields_newInstances, sliceofpointertogongstructfield)
			newInstancesSlice = append(newInstancesSlice, sliceofpointertogongstructfield.GongMarshallIdentifier(stage))
			if stage.SliceOfPointerToGongStructFields_referenceOrder == nil {
				stage.SliceOfPointerToGongStructFields_referenceOrder = make(map[*SliceOfPointerToGongStructField]uint)
			}
			stage.SliceOfPointerToGongStructFields_referenceOrder[sliceofpointertogongstructfield] = stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, sliceofpointertogongstructfield.GongMarshallUnstaging(stage))
			delete(stage.SliceOfPointerToGongStructFields_referenceOrder, sliceofpointertogongstructfield)
			fieldInitializers, pointersInitializations := sliceofpointertogongstructfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SliceOfPointerToGongStructFieldMap_Staged_Order[ref] = stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfield]
			diffs := sliceofpointertogongstructfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, sliceofpointertogongstructfield)
			delete(stage.SliceOfPointerToGongStructFieldMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", sliceofpointertogongstructfield.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SliceOfPointerToGongStructFields_reference {
		if _, ok := stage.SliceOfPointerToGongStructFields[ref]; !ok {
			sliceofpointertogongstructfields_deletedInstances = append(sliceofpointertogongstructfields_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(sliceofpointertogongstructfields_newInstances)
	lenDeletedInstances += len(sliceofpointertogongstructfields_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.GongBasicFields_reference = make(map[*GongBasicField]*GongBasicField)
	stage.GongBasicFields_referenceOrder = make(map[*GongBasicField]uint) // diff Unstage needs the reference order
	for instance := range stage.GongBasicFields {
		stage.GongBasicFields_reference[instance] = instance.GongCopy().(*GongBasicField)
		stage.GongBasicFields_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongEnums_reference = make(map[*GongEnum]*GongEnum)
	stage.GongEnums_referenceOrder = make(map[*GongEnum]uint) // diff Unstage needs the reference order
	for instance := range stage.GongEnums {
		stage.GongEnums_reference[instance] = instance.GongCopy().(*GongEnum)
		stage.GongEnums_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongEnumValues_reference = make(map[*GongEnumValue]*GongEnumValue)
	stage.GongEnumValues_referenceOrder = make(map[*GongEnumValue]uint) // diff Unstage needs the reference order
	for instance := range stage.GongEnumValues {
		stage.GongEnumValues_reference[instance] = instance.GongCopy().(*GongEnumValue)
		stage.GongEnumValues_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongLinks_reference = make(map[*GongLink]*GongLink)
	stage.GongLinks_referenceOrder = make(map[*GongLink]uint) // diff Unstage needs the reference order
	for instance := range stage.GongLinks {
		stage.GongLinks_reference[instance] = instance.GongCopy().(*GongLink)
		stage.GongLinks_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongNotes_reference = make(map[*GongNote]*GongNote)
	stage.GongNotes_referenceOrder = make(map[*GongNote]uint) // diff Unstage needs the reference order
	for instance := range stage.GongNotes {
		stage.GongNotes_reference[instance] = instance.GongCopy().(*GongNote)
		stage.GongNotes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongStructs_reference = make(map[*GongStruct]*GongStruct)
	stage.GongStructs_referenceOrder = make(map[*GongStruct]uint) // diff Unstage needs the reference order
	for instance := range stage.GongStructs {
		stage.GongStructs_reference[instance] = instance.GongCopy().(*GongStruct)
		stage.GongStructs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongTimeFields_reference = make(map[*GongTimeField]*GongTimeField)
	stage.GongTimeFields_referenceOrder = make(map[*GongTimeField]uint) // diff Unstage needs the reference order
	for instance := range stage.GongTimeFields {
		stage.GongTimeFields_reference[instance] = instance.GongCopy().(*GongTimeField)
		stage.GongTimeFields_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.MetaReferences_reference = make(map[*MetaReference]*MetaReference)
	stage.MetaReferences_referenceOrder = make(map[*MetaReference]uint) // diff Unstage needs the reference order
	for instance := range stage.MetaReferences {
		stage.MetaReferences_reference[instance] = instance.GongCopy().(*MetaReference)
		stage.MetaReferences_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ModelPkgs_reference = make(map[*ModelPkg]*ModelPkg)
	stage.ModelPkgs_referenceOrder = make(map[*ModelPkg]uint) // diff Unstage needs the reference order
	for instance := range stage.ModelPkgs {
		stage.ModelPkgs_reference[instance] = instance.GongCopy().(*ModelPkg)
		stage.ModelPkgs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.PointerToGongStructFields_reference = make(map[*PointerToGongStructField]*PointerToGongStructField)
	stage.PointerToGongStructFields_referenceOrder = make(map[*PointerToGongStructField]uint) // diff Unstage needs the reference order
	for instance := range stage.PointerToGongStructFields {
		stage.PointerToGongStructFields_reference[instance] = instance.GongCopy().(*PointerToGongStructField)
		stage.PointerToGongStructFields_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SliceOfPointerToGongStructFields_reference = make(map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField)
	stage.SliceOfPointerToGongStructFields_referenceOrder = make(map[*SliceOfPointerToGongStructField]uint) // diff Unstage needs the reference order
	for instance := range stage.SliceOfPointerToGongStructFields {
		stage.SliceOfPointerToGongStructFields_reference[instance] = instance.GongCopy().(*SliceOfPointerToGongStructField)
		stage.SliceOfPointerToGongStructFields_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongBasicFieldMap_Staged_Order[gongbasicfield]; ok {
		return order
	}
	return stage.GongBasicFields_referenceOrder[gongbasicfield]
}

func (gongbasicfield *GongBasicField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongBasicFields_referenceOrder[gongbasicfield]
}

func (gongenum *GongEnum) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnumMap_Staged_Order[gongenum]; ok {
		return order
	}
	return stage.GongEnums_referenceOrder[gongenum]
}

func (gongenum *GongEnum) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongEnums_referenceOrder[gongenum]
}

func (gongenumvalue *GongEnumValue) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnumValueMap_Staged_Order[gongenumvalue]; ok {
		return order
	}
	return stage.GongEnumValues_referenceOrder[gongenumvalue]
}

func (gongenumvalue *GongEnumValue) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongEnumValues_referenceOrder[gongenumvalue]
}

func (gonglink *GongLink) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongLinkMap_Staged_Order[gonglink]; ok {
		return order
	}
	return stage.GongLinks_referenceOrder[gonglink]
}

func (gonglink *GongLink) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongLinks_referenceOrder[gonglink]
}

func (gongnote *GongNote) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongNoteMap_Staged_Order[gongnote]; ok {
		return order
	}
	return stage.GongNotes_referenceOrder[gongnote]
}

func (gongnote *GongNote) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongNotes_referenceOrder[gongnote]
}

func (gongstruct *GongStruct) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongStructMap_Staged_Order[gongstruct]; ok {
		return order
	}
	return stage.GongStructs_referenceOrder[gongstruct]
}

func (gongstruct *GongStruct) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongStructs_referenceOrder[gongstruct]
}

func (gongtimefield *GongTimeField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongTimeFieldMap_Staged_Order[gongtimefield]; ok {
		return order
	}
	return stage.GongTimeFields_referenceOrder[gongtimefield]
}

func (gongtimefield *GongTimeField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongTimeFields_referenceOrder[gongtimefield]
}

func (metareference *MetaReference) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MetaReferenceMap_Staged_Order[metareference]; ok {
		return order
	}
	return stage.MetaReferences_referenceOrder[metareference]
}

func (metareference *MetaReference) GongGetReferenceOrder(stage *Stage) uint {
	return stage.MetaReferences_referenceOrder[metareference]
}

func (modelpkg *ModelPkg) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ModelPkgMap_Staged_Order[modelpkg]; ok {
		return order
	}
	return stage.ModelPkgs_referenceOrder[modelpkg]
}

func (modelpkg *ModelPkg) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ModelPkgs_referenceOrder[modelpkg]
}

func (pointertogongstructfield *PointerToGongStructField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfield]; ok {
		return order
	}
	return stage.PointerToGongStructFields_referenceOrder[pointertogongstructfield]
}

func (pointertogongstructfield *PointerToGongStructField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.PointerToGongStructFields_referenceOrder[pointertogongstructfield]
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfield]; ok {
		return order
	}
	return stage.SliceOfPointerToGongStructFields_referenceOrder[sliceofpointertogongstructfield]
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SliceOfPointerToGongStructFields_referenceOrder[sliceofpointertogongstructfield]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongbasicfield.GongGetGongstructName(), gongbasicfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongbasicfield *GongBasicField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongbasicfield.GongGetGongstructName(), gongbasicfield.GongGetReferenceOrder(stage))
}

func (gongenum *GongEnum) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenum.GongGetGongstructName(), gongenum.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenum *GongEnum) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenum.GongGetGongstructName(), gongenum.GongGetReferenceOrder(stage))
}

func (gongenumvalue *GongEnumValue) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalue.GongGetGongstructName(), gongenumvalue.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumvalue *GongEnumValue) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalue.GongGetGongstructName(), gongenumvalue.GongGetReferenceOrder(stage))
}

func (gonglink *GongLink) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gonglink.GongGetGongstructName(), gonglink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gonglink *GongLink) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gonglink.GongGetGongstructName(), gonglink.GongGetReferenceOrder(stage))
}

func (gongnote *GongNote) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnote.GongGetGongstructName(), gongnote.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnote *GongNote) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnote.GongGetGongstructName(), gongnote.GongGetReferenceOrder(stage))
}

func (gongstruct *GongStruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstruct.GongGetGongstructName(), gongstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongstruct *GongStruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstruct.GongGetGongstructName(), gongstruct.GongGetReferenceOrder(stage))
}

func (gongtimefield *GongTimeField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongtimefield.GongGetGongstructName(), gongtimefield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongtimefield *GongTimeField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongtimefield.GongGetGongstructName(), gongtimefield.GongGetReferenceOrder(stage))
}

func (metareference *MetaReference) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metareference.GongGetGongstructName(), metareference.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metareference *MetaReference) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metareference.GongGetGongstructName(), metareference.GongGetReferenceOrder(stage))
}

func (modelpkg *ModelPkg) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", modelpkg.GongGetGongstructName(), modelpkg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (modelpkg *ModelPkg) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", modelpkg.GongGetGongstructName(), modelpkg.GongGetReferenceOrder(stage))
}

func (pointertogongstructfield *PointerToGongStructField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pointertogongstructfield.GongGetGongstructName(), pointertogongstructfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointertogongstructfield *PointerToGongStructField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pointertogongstructfield.GongGetGongstructName(), pointertogongstructfield.GongGetReferenceOrder(stage))
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sliceofpointertogongstructfield.GongGetGongstructName(), sliceofpointertogongstructfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sliceofpointertogongstructfield.GongGetGongstructName(), sliceofpointertogongstructfield.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongBasicField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongbasicfield.Name)
	return
}

func (gongenum *GongEnum) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnum")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongenum.Name)
	return
}

func (gongenumvalue *GongEnumValue) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumValue")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongenumvalue.Name)
	return
}

func (gonglink *GongLink) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongLink")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gonglink.Name)
	return
}

func (gongnote *GongNote) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNote")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongnote.Name)
	return
}

func (gongstruct *GongStruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongStruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongstruct.Name)
	return
}

func (gongtimefield *GongTimeField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongTimeField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongtimefield.Name)
	return
}

func (metareference *MetaReference) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metareference.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MetaReference")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", metareference.Name)
	return
}

func (modelpkg *ModelPkg) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ModelPkg")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", modelpkg.Name)
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PointerToGongStructField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", pointertogongstructfield.Name)
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SliceOfPointerToGongStructField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", sliceofpointertogongstructfield.Name)
	return
}

// insertion point for unstaging
func (gongbasicfield *GongBasicField) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongbasicfield.GongGetReferenceIdentifier(stage))
	return
}

func (gongenum *GongEnum) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenum.GongGetReferenceIdentifier(stage))
	return
}

func (gongenumvalue *GongEnumValue) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumvalue.GongGetReferenceIdentifier(stage))
	return
}

func (gonglink *GongLink) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gonglink.GongGetReferenceIdentifier(stage))
	return
}

func (gongnote *GongNote) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnote.GongGetReferenceIdentifier(stage))
	return
}

func (gongstruct *GongStruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongstruct.GongGetReferenceIdentifier(stage))
	return
}

func (gongtimefield *GongTimeField) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongtimefield.GongGetReferenceIdentifier(stage))
	return
}

func (metareference *MetaReference) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metareference.GongGetReferenceIdentifier(stage))
	return
}

func (modelpkg *ModelPkg) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", modelpkg.GongGetReferenceIdentifier(stage))
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pointertogongstructfield.GongGetReferenceIdentifier(stage))
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sliceofpointertogongstructfield.GongGetReferenceIdentifier(stage))
	return
}

// end of template
