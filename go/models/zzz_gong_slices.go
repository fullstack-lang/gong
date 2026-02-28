// generated code - do not edit
package models

import (
	"fmt"
	"log"
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
	newInstance := new(GongBasicField)
	gongbasicfield.CopyBasicFields(newInstance)
	return newInstance
}

func (gongenum *GongEnum) GongCopy() GongstructIF {
	newInstance := new(GongEnum)
	gongenum.CopyBasicFields(newInstance)
	return newInstance
}

func (gongenumvalue *GongEnumValue) GongCopy() GongstructIF {
	newInstance := new(GongEnumValue)
	gongenumvalue.CopyBasicFields(newInstance)
	return newInstance
}

func (gonglink *GongLink) GongCopy() GongstructIF {
	newInstance := new(GongLink)
	gonglink.CopyBasicFields(newInstance)
	return newInstance
}

func (gongnote *GongNote) GongCopy() GongstructIF {
	newInstance := new(GongNote)
	gongnote.CopyBasicFields(newInstance)
	return newInstance
}

func (gongstruct *GongStruct) GongCopy() GongstructIF {
	newInstance := new(GongStruct)
	gongstruct.CopyBasicFields(newInstance)
	return newInstance
}

func (gongtimefield *GongTimeField) GongCopy() GongstructIF {
	newInstance := new(GongTimeField)
	gongtimefield.CopyBasicFields(newInstance)
	return newInstance
}

func (metareference *MetaReference) GongCopy() GongstructIF {
	newInstance := new(MetaReference)
	metareference.CopyBasicFields(newInstance)
	return newInstance
}

func (modelpkg *ModelPkg) GongCopy() GongstructIF {
	newInstance := new(ModelPkg)
	modelpkg.CopyBasicFields(newInstance)
	return newInstance
}

func (pointertogongstructfield *PointerToGongStructField) GongCopy() GongstructIF {
	newInstance := new(PointerToGongStructField)
	pointertogongstructfield.CopyBasicFields(newInstance)
	return newInstance
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongCopy() GongstructIF {
	newInstance := new(SliceOfPointerToGongStructField)
	sliceofpointertogongstructfield.CopyBasicFields(newInstance)
	return newInstance
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
			stage.GongBasicFields_referenceOrder[gongbasicfield] = stage.GongBasicField_stagedOrder[gongbasicfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongbasicfield.GongMarshallUnstaging(stage))
			// delete(stage.GongBasicFields_referenceOrder, gongbasicfield)
			fieldInitializers, pointersInitializations := gongbasicfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongBasicField_stagedOrder[ref] = stage.GongBasicField_stagedOrder[gongbasicfield]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gongbasicfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongbasicfield)
			// delete(stage.GongBasicField_stagedOrder, ref)
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
	for _, ref := range stage.GongBasicFields_reference {
		instance := stage.GongBasicFields_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GongBasicFields[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.GongEnums_referenceOrder[gongenum] = stage.GongEnum_stagedOrder[gongenum]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongenum.GongMarshallUnstaging(stage))
			// delete(stage.GongEnums_referenceOrder, gongenum)
			fieldInitializers, pointersInitializations := gongenum.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongEnum_stagedOrder[ref] = stage.GongEnum_stagedOrder[gongenum]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gongenum.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongenum)
			// delete(stage.GongEnum_stagedOrder, ref)
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
	for _, ref := range stage.GongEnums_reference {
		instance := stage.GongEnums_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GongEnums[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.GongEnumValues_referenceOrder[gongenumvalue] = stage.GongEnumValue_stagedOrder[gongenumvalue]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongenumvalue.GongMarshallUnstaging(stage))
			// delete(stage.GongEnumValues_referenceOrder, gongenumvalue)
			fieldInitializers, pointersInitializations := gongenumvalue.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongEnumValue_stagedOrder[ref] = stage.GongEnumValue_stagedOrder[gongenumvalue]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gongenumvalue.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongenumvalue)
			// delete(stage.GongEnumValue_stagedOrder, ref)
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
	for _, ref := range stage.GongEnumValues_reference {
		instance := stage.GongEnumValues_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GongEnumValues[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.GongLinks_referenceOrder[gonglink] = stage.GongLink_stagedOrder[gonglink]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gonglink.GongMarshallUnstaging(stage))
			// delete(stage.GongLinks_referenceOrder, gonglink)
			fieldInitializers, pointersInitializations := gonglink.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongLink_stagedOrder[ref] = stage.GongLink_stagedOrder[gonglink]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gonglink.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gonglink)
			// delete(stage.GongLink_stagedOrder, ref)
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
	for _, ref := range stage.GongLinks_reference {
		instance := stage.GongLinks_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GongLinks[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.GongNotes_referenceOrder[gongnote] = stage.GongNote_stagedOrder[gongnote]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongnote.GongMarshallUnstaging(stage))
			// delete(stage.GongNotes_referenceOrder, gongnote)
			fieldInitializers, pointersInitializations := gongnote.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongNote_stagedOrder[ref] = stage.GongNote_stagedOrder[gongnote]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gongnote.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongnote)
			// delete(stage.GongNote_stagedOrder, ref)
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
	for _, ref := range stage.GongNotes_reference {
		instance := stage.GongNotes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GongNotes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.GongStructs_referenceOrder[gongstruct] = stage.GongStruct_stagedOrder[gongstruct]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongstruct.GongMarshallUnstaging(stage))
			// delete(stage.GongStructs_referenceOrder, gongstruct)
			fieldInitializers, pointersInitializations := gongstruct.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongStruct_stagedOrder[ref] = stage.GongStruct_stagedOrder[gongstruct]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gongstruct.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongstruct)
			// delete(stage.GongStruct_stagedOrder, ref)
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
	for _, ref := range stage.GongStructs_reference {
		instance := stage.GongStructs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GongStructs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.GongTimeFields_referenceOrder[gongtimefield] = stage.GongTimeField_stagedOrder[gongtimefield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongtimefield.GongMarshallUnstaging(stage))
			// delete(stage.GongTimeFields_referenceOrder, gongtimefield)
			fieldInitializers, pointersInitializations := gongtimefield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongTimeField_stagedOrder[ref] = stage.GongTimeField_stagedOrder[gongtimefield]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := gongtimefield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongtimefield)
			// delete(stage.GongTimeField_stagedOrder, ref)
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
	for _, ref := range stage.GongTimeFields_reference {
		instance := stage.GongTimeFields_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.GongTimeFields[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.MetaReferences_referenceOrder[metareference] = stage.MetaReference_stagedOrder[metareference]
			newInstancesReverseSlice = append(newInstancesReverseSlice, metareference.GongMarshallUnstaging(stage))
			// delete(stage.MetaReferences_referenceOrder, metareference)
			fieldInitializers, pointersInitializations := metareference.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MetaReference_stagedOrder[ref] = stage.MetaReference_stagedOrder[metareference]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := metareference.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, metareference)
			// delete(stage.MetaReference_stagedOrder, ref)
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
	for _, ref := range stage.MetaReferences_reference {
		instance := stage.MetaReferences_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.MetaReferences[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.ModelPkgs_referenceOrder[modelpkg] = stage.ModelPkg_stagedOrder[modelpkg]
			newInstancesReverseSlice = append(newInstancesReverseSlice, modelpkg.GongMarshallUnstaging(stage))
			// delete(stage.ModelPkgs_referenceOrder, modelpkg)
			fieldInitializers, pointersInitializations := modelpkg.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ModelPkg_stagedOrder[ref] = stage.ModelPkg_stagedOrder[modelpkg]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := modelpkg.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, modelpkg)
			// delete(stage.ModelPkg_stagedOrder, ref)
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
	for _, ref := range stage.ModelPkgs_reference {
		instance := stage.ModelPkgs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.ModelPkgs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.PointerToGongStructFields_referenceOrder[pointertogongstructfield] = stage.PointerToGongStructField_stagedOrder[pointertogongstructfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, pointertogongstructfield.GongMarshallUnstaging(stage))
			// delete(stage.PointerToGongStructFields_referenceOrder, pointertogongstructfield)
			fieldInitializers, pointersInitializations := pointertogongstructfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PointerToGongStructField_stagedOrder[ref] = stage.PointerToGongStructField_stagedOrder[pointertogongstructfield]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := pointertogongstructfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, pointertogongstructfield)
			// delete(stage.PointerToGongStructField_stagedOrder, ref)
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
	for _, ref := range stage.PointerToGongStructFields_reference {
		instance := stage.PointerToGongStructFields_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PointerToGongStructFields[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
			stage.SliceOfPointerToGongStructFields_referenceOrder[sliceofpointertogongstructfield] = stage.SliceOfPointerToGongStructField_stagedOrder[sliceofpointertogongstructfield]
			newInstancesReverseSlice = append(newInstancesReverseSlice, sliceofpointertogongstructfield.GongMarshallUnstaging(stage))
			// delete(stage.SliceOfPointerToGongStructFields_referenceOrder, sliceofpointertogongstructfield)
			fieldInitializers, pointersInitializations := sliceofpointertogongstructfield.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SliceOfPointerToGongStructField_stagedOrder[ref] = stage.SliceOfPointerToGongStructField_stagedOrder[sliceofpointertogongstructfield]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := sliceofpointertogongstructfield.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, sliceofpointertogongstructfield)
			// delete(stage.SliceOfPointerToGongStructField_stagedOrder, ref)
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
	for _, ref := range stage.SliceOfPointerToGongStructFields_reference {
		instance := stage.SliceOfPointerToGongStructFields_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SliceOfPointerToGongStructFields[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
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
	stage.GongBasicFields_instance = make(map[*GongBasicField]*GongBasicField)
	for instance := range stage.GongBasicFields {
		_copy := instance.GongCopy().(*GongBasicField)
		stage.GongBasicFields_reference[instance] = _copy
		stage.GongBasicFields_instance[_copy] = instance
		stage.GongBasicFields_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongEnums_reference = make(map[*GongEnum]*GongEnum)
	stage.GongEnums_referenceOrder = make(map[*GongEnum]uint) // diff Unstage needs the reference order
	stage.GongEnums_instance = make(map[*GongEnum]*GongEnum)
	for instance := range stage.GongEnums {
		_copy := instance.GongCopy().(*GongEnum)
		stage.GongEnums_reference[instance] = _copy
		stage.GongEnums_instance[_copy] = instance
		stage.GongEnums_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongEnumValues_reference = make(map[*GongEnumValue]*GongEnumValue)
	stage.GongEnumValues_referenceOrder = make(map[*GongEnumValue]uint) // diff Unstage needs the reference order
	stage.GongEnumValues_instance = make(map[*GongEnumValue]*GongEnumValue)
	for instance := range stage.GongEnumValues {
		_copy := instance.GongCopy().(*GongEnumValue)
		stage.GongEnumValues_reference[instance] = _copy
		stage.GongEnumValues_instance[_copy] = instance
		stage.GongEnumValues_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongLinks_reference = make(map[*GongLink]*GongLink)
	stage.GongLinks_referenceOrder = make(map[*GongLink]uint) // diff Unstage needs the reference order
	stage.GongLinks_instance = make(map[*GongLink]*GongLink)
	for instance := range stage.GongLinks {
		_copy := instance.GongCopy().(*GongLink)
		stage.GongLinks_reference[instance] = _copy
		stage.GongLinks_instance[_copy] = instance
		stage.GongLinks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongNotes_reference = make(map[*GongNote]*GongNote)
	stage.GongNotes_referenceOrder = make(map[*GongNote]uint) // diff Unstage needs the reference order
	stage.GongNotes_instance = make(map[*GongNote]*GongNote)
	for instance := range stage.GongNotes {
		_copy := instance.GongCopy().(*GongNote)
		stage.GongNotes_reference[instance] = _copy
		stage.GongNotes_instance[_copy] = instance
		stage.GongNotes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongStructs_reference = make(map[*GongStruct]*GongStruct)
	stage.GongStructs_referenceOrder = make(map[*GongStruct]uint) // diff Unstage needs the reference order
	stage.GongStructs_instance = make(map[*GongStruct]*GongStruct)
	for instance := range stage.GongStructs {
		_copy := instance.GongCopy().(*GongStruct)
		stage.GongStructs_reference[instance] = _copy
		stage.GongStructs_instance[_copy] = instance
		stage.GongStructs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongTimeFields_reference = make(map[*GongTimeField]*GongTimeField)
	stage.GongTimeFields_referenceOrder = make(map[*GongTimeField]uint) // diff Unstage needs the reference order
	stage.GongTimeFields_instance = make(map[*GongTimeField]*GongTimeField)
	for instance := range stage.GongTimeFields {
		_copy := instance.GongCopy().(*GongTimeField)
		stage.GongTimeFields_reference[instance] = _copy
		stage.GongTimeFields_instance[_copy] = instance
		stage.GongTimeFields_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.MetaReferences_reference = make(map[*MetaReference]*MetaReference)
	stage.MetaReferences_referenceOrder = make(map[*MetaReference]uint) // diff Unstage needs the reference order
	stage.MetaReferences_instance = make(map[*MetaReference]*MetaReference)
	for instance := range stage.MetaReferences {
		_copy := instance.GongCopy().(*MetaReference)
		stage.MetaReferences_reference[instance] = _copy
		stage.MetaReferences_instance[_copy] = instance
		stage.MetaReferences_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ModelPkgs_reference = make(map[*ModelPkg]*ModelPkg)
	stage.ModelPkgs_referenceOrder = make(map[*ModelPkg]uint) // diff Unstage needs the reference order
	stage.ModelPkgs_instance = make(map[*ModelPkg]*ModelPkg)
	for instance := range stage.ModelPkgs {
		_copy := instance.GongCopy().(*ModelPkg)
		stage.ModelPkgs_reference[instance] = _copy
		stage.ModelPkgs_instance[_copy] = instance
		stage.ModelPkgs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PointerToGongStructFields_reference = make(map[*PointerToGongStructField]*PointerToGongStructField)
	stage.PointerToGongStructFields_referenceOrder = make(map[*PointerToGongStructField]uint) // diff Unstage needs the reference order
	stage.PointerToGongStructFields_instance = make(map[*PointerToGongStructField]*PointerToGongStructField)
	for instance := range stage.PointerToGongStructFields {
		_copy := instance.GongCopy().(*PointerToGongStructField)
		stage.PointerToGongStructFields_reference[instance] = _copy
		stage.PointerToGongStructFields_instance[_copy] = instance
		stage.PointerToGongStructFields_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SliceOfPointerToGongStructFields_reference = make(map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField)
	stage.SliceOfPointerToGongStructFields_referenceOrder = make(map[*SliceOfPointerToGongStructField]uint) // diff Unstage needs the reference order
	stage.SliceOfPointerToGongStructFields_instance = make(map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField)
	for instance := range stage.SliceOfPointerToGongStructFields {
		_copy := instance.GongCopy().(*SliceOfPointerToGongStructField)
		stage.SliceOfPointerToGongStructFields_reference[instance] = _copy
		stage.SliceOfPointerToGongStructFields_instance[_copy] = instance
		stage.SliceOfPointerToGongStructFields_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.GongBasicFields {
		reference := stage.GongBasicFields_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongEnums {
		reference := stage.GongEnums_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongEnumValues {
		reference := stage.GongEnumValues_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongLinks {
		reference := stage.GongLinks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongNotes {
		reference := stage.GongNotes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongStructs {
		reference := stage.GongStructs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongTimeFields {
		reference := stage.GongTimeFields_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.MetaReferences {
		reference := stage.MetaReferences_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ModelPkgs {
		reference := stage.ModelPkgs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PointerToGongStructFields {
		reference := stage.PointerToGongStructFields_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SliceOfPointerToGongStructFields {
		reference := stage.SliceOfPointerToGongStructFields_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
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
	if order, ok := stage.GongBasicField_stagedOrder[gongbasicfield]; ok {
		return order
	}
	if order, ok := stage.GongBasicFields_referenceOrder[gongbasicfield]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongBasicField was not staged and does not have a reference order", gongbasicfield)
		return 0
	}
}

func (gongenum *GongEnum) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnum_stagedOrder[gongenum]; ok {
		return order
	}
	if order, ok := stage.GongEnums_referenceOrder[gongenum]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongEnum was not staged and does not have a reference order", gongenum)
		return 0
	}
}

func (gongenumvalue *GongEnumValue) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnumValue_stagedOrder[gongenumvalue]; ok {
		return order
	}
	if order, ok := stage.GongEnumValues_referenceOrder[gongenumvalue]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongEnumValue was not staged and does not have a reference order", gongenumvalue)
		return 0
	}
}

func (gonglink *GongLink) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongLink_stagedOrder[gonglink]; ok {
		return order
	}
	if order, ok := stage.GongLinks_referenceOrder[gonglink]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongLink was not staged and does not have a reference order", gonglink)
		return 0
	}
}

func (gongnote *GongNote) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongNote_stagedOrder[gongnote]; ok {
		return order
	}
	if order, ok := stage.GongNotes_referenceOrder[gongnote]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongNote was not staged and does not have a reference order", gongnote)
		return 0
	}
}

func (gongstruct *GongStruct) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongStruct_stagedOrder[gongstruct]; ok {
		return order
	}
	if order, ok := stage.GongStructs_referenceOrder[gongstruct]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongStruct was not staged and does not have a reference order", gongstruct)
		return 0
	}
}

func (gongtimefield *GongTimeField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongTimeField_stagedOrder[gongtimefield]; ok {
		return order
	}
	if order, ok := stage.GongTimeFields_referenceOrder[gongtimefield]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongTimeField was not staged and does not have a reference order", gongtimefield)
		return 0
	}
}

func (metareference *MetaReference) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MetaReference_stagedOrder[metareference]; ok {
		return order
	}
	if order, ok := stage.MetaReferences_referenceOrder[metareference]; ok {
		return order
	} else {
		log.Printf("instance %p of type MetaReference was not staged and does not have a reference order", metareference)
		return 0
	}
}

func (modelpkg *ModelPkg) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ModelPkg_stagedOrder[modelpkg]; ok {
		return order
	}
	if order, ok := stage.ModelPkgs_referenceOrder[modelpkg]; ok {
		return order
	} else {
		log.Printf("instance %p of type ModelPkg was not staged and does not have a reference order", modelpkg)
		return 0
	}
}

func (pointertogongstructfield *PointerToGongStructField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PointerToGongStructField_stagedOrder[pointertogongstructfield]; ok {
		return order
	}
	if order, ok := stage.PointerToGongStructFields_referenceOrder[pointertogongstructfield]; ok {
		return order
	} else {
		log.Printf("instance %p of type PointerToGongStructField was not staged and does not have a reference order", pointertogongstructfield)
		return 0
	}
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SliceOfPointerToGongStructField_stagedOrder[sliceofpointertogongstructfield]; ok {
		return order
	}
	if order, ok := stage.SliceOfPointerToGongStructFields_referenceOrder[sliceofpointertogongstructfield]; ok {
		return order
	} else {
		log.Printf("instance %p of type SliceOfPointerToGongStructField was not staged and does not have a reference order", sliceofpointertogongstructfield)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", gongbasicfield.GongGetGongstructName(), gongbasicfield.GongGetOrder(stage))
}

func (gongenum *GongEnum) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenum.GongGetGongstructName(), gongenum.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenum *GongEnum) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenum.GongGetGongstructName(), gongenum.GongGetOrder(stage))
}

func (gongenumvalue *GongEnumValue) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalue.GongGetGongstructName(), gongenumvalue.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumvalue *GongEnumValue) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalue.GongGetGongstructName(), gongenumvalue.GongGetOrder(stage))
}

func (gonglink *GongLink) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gonglink.GongGetGongstructName(), gonglink.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gonglink *GongLink) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gonglink.GongGetGongstructName(), gonglink.GongGetOrder(stage))
}

func (gongnote *GongNote) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnote.GongGetGongstructName(), gongnote.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnote *GongNote) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnote.GongGetGongstructName(), gongnote.GongGetOrder(stage))
}

func (gongstruct *GongStruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstruct.GongGetGongstructName(), gongstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongstruct *GongStruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstruct.GongGetGongstructName(), gongstruct.GongGetOrder(stage))
}

func (gongtimefield *GongTimeField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongtimefield.GongGetGongstructName(), gongtimefield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongtimefield *GongTimeField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongtimefield.GongGetGongstructName(), gongtimefield.GongGetOrder(stage))
}

func (metareference *MetaReference) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metareference.GongGetGongstructName(), metareference.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metareference *MetaReference) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metareference.GongGetGongstructName(), metareference.GongGetOrder(stage))
}

func (modelpkg *ModelPkg) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", modelpkg.GongGetGongstructName(), modelpkg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (modelpkg *ModelPkg) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", modelpkg.GongGetGongstructName(), modelpkg.GongGetOrder(stage))
}

func (pointertogongstructfield *PointerToGongStructField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pointertogongstructfield.GongGetGongstructName(), pointertogongstructfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pointertogongstructfield *PointerToGongStructField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pointertogongstructfield.GongGetGongstructName(), pointertogongstructfield.GongGetOrder(stage))
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sliceofpointertogongstructfield.GongGetGongstructName(), sliceofpointertogongstructfield.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sliceofpointertogongstructfield.GongGetGongstructName(), sliceofpointertogongstructfield.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongbasicfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongBasicField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gongbasicfield.Name))
	return
}

func (gongenum *GongEnum) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenum.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnum")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gongenum.Name))
	return
}

func (gongenumvalue *GongEnumValue) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumvalue.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumValue")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gongenumvalue.Name))
	return
}

func (gonglink *GongLink) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gonglink.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongLink")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gonglink.Name))
	return
}

func (gongnote *GongNote) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnote.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNote")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gongnote.Name))
	return
}

func (gongstruct *GongStruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongStruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gongstruct.Name))
	return
}

func (gongtimefield *GongTimeField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongtimefield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongTimeField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(gongtimefield.Name))
	return
}

func (metareference *MetaReference) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metareference.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "MetaReference")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(metareference.Name))
	return
}

func (modelpkg *ModelPkg) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", modelpkg.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ModelPkg")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(modelpkg.Name))
	return
}

func (pointertogongstructfield *PointerToGongStructField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pointertogongstructfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PointerToGongStructField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(pointertogongstructfield.Name))
	return
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sliceofpointertogongstructfield.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SliceOfPointerToGongStructField")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(sliceofpointertogongstructfield.Name))
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
