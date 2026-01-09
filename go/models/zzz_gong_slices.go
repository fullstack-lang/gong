// generated code - do not edit
package models

import (
	"fmt"
	"strings"
	"time"
)

var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesStmt string
	_ = newInstancesStmt
	var fieldsEditStmt string
	_ = fieldsEditStmt
	var deletedInstancesStmt string
	_ = deletedInstancesStmt

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
			newInstancesStmt += gongbasicfield.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gongbasicfield.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gongbasicfield.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gongbasicfield.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongbasicfield := range stage.GongBasicFields_reference {
		if _, ok := stage.GongBasicFields[gongbasicfield]; !ok {
			gongbasicfields_deletedInstances = append(gongbasicfields_deletedInstances, gongbasicfield)
			deletedInstancesStmt += gongbasicfield.GongMarshallUnstaging(stage)
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
			newInstancesStmt += gongenum.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gongenum.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gongenum.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gongenum.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongenum := range stage.GongEnums_reference {
		if _, ok := stage.GongEnums[gongenum]; !ok {
			gongenums_deletedInstances = append(gongenums_deletedInstances, gongenum)
			deletedInstancesStmt += gongenum.GongMarshallUnstaging(stage)
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
			newInstancesStmt += gongenumvalue.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gongenumvalue.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gongenumvalue.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gongenumvalue.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongenumvalue := range stage.GongEnumValues_reference {
		if _, ok := stage.GongEnumValues[gongenumvalue]; !ok {
			gongenumvalues_deletedInstances = append(gongenumvalues_deletedInstances, gongenumvalue)
			deletedInstancesStmt += gongenumvalue.GongMarshallUnstaging(stage)
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
			newInstancesStmt += gonglink.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gonglink.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gonglink.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gonglink.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gonglink := range stage.GongLinks_reference {
		if _, ok := stage.GongLinks[gonglink]; !ok {
			gonglinks_deletedInstances = append(gonglinks_deletedInstances, gonglink)
			deletedInstancesStmt += gonglink.GongMarshallUnstaging(stage)
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
			newInstancesStmt += gongnote.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gongnote.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gongnote.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gongnote.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongnote := range stage.GongNotes_reference {
		if _, ok := stage.GongNotes[gongnote]; !ok {
			gongnotes_deletedInstances = append(gongnotes_deletedInstances, gongnote)
			deletedInstancesStmt += gongnote.GongMarshallUnstaging(stage)
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
			newInstancesStmt += gongstruct.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gongstruct.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gongstruct.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gongstruct.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongstruct := range stage.GongStructs_reference {
		if _, ok := stage.GongStructs[gongstruct]; !ok {
			gongstructs_deletedInstances = append(gongstructs_deletedInstances, gongstruct)
			deletedInstancesStmt += gongstruct.GongMarshallUnstaging(stage)
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
			newInstancesStmt += gongtimefield.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := gongtimefield.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := gongtimefield.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", gongtimefield.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongtimefield := range stage.GongTimeFields_reference {
		if _, ok := stage.GongTimeFields[gongtimefield]; !ok {
			gongtimefields_deletedInstances = append(gongtimefields_deletedInstances, gongtimefield)
			deletedInstancesStmt += gongtimefield.GongMarshallUnstaging(stage)
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
			newInstancesStmt += metareference.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := metareference.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := metareference.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", metareference.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for metareference := range stage.MetaReferences_reference {
		if _, ok := stage.MetaReferences[metareference]; !ok {
			metareferences_deletedInstances = append(metareferences_deletedInstances, metareference)
			deletedInstancesStmt += metareference.GongMarshallUnstaging(stage)
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
			newInstancesStmt += modelpkg.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := modelpkg.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := modelpkg.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", modelpkg.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for modelpkg := range stage.ModelPkgs_reference {
		if _, ok := stage.ModelPkgs[modelpkg]; !ok {
			modelpkgs_deletedInstances = append(modelpkgs_deletedInstances, modelpkg)
			deletedInstancesStmt += modelpkg.GongMarshallUnstaging(stage)
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
			newInstancesStmt += pointertogongstructfield.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := pointertogongstructfield.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := pointertogongstructfield.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", pointertogongstructfield.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for pointertogongstructfield := range stage.PointerToGongStructFields_reference {
		if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; !ok {
			pointertogongstructfields_deletedInstances = append(pointertogongstructfields_deletedInstances, pointertogongstructfield)
			deletedInstancesStmt += pointertogongstructfield.GongMarshallUnstaging(stage)
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
			newInstancesStmt += sliceofpointertogongstructfield.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := sliceofpointertogongstructfield.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := sliceofpointertogongstructfield.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", sliceofpointertogongstructfield.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields_reference {
		if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; !ok {
			sliceofpointertogongstructfields_deletedInstances = append(sliceofpointertogongstructfields_deletedInstances, sliceofpointertogongstructfield)
			deletedInstancesStmt += sliceofpointertogongstructfield.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(sliceofpointertogongstructfields_newInstances)
	lenDeletedInstances += len(sliceofpointertogongstructfields_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += "\n\tstage.Commit()"
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				notif,
			)
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

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

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (gongbasicfield *GongBasicField) GongGetOrder(stage *Stage) uint {
	return stage.GongBasicFieldMap_Staged_Order[gongbasicfield]
}

func (gongbasicfield *GongBasicField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongBasicFields_referenceOrder[gongbasicfield]
}

func (gongenum *GongEnum) GongGetOrder(stage *Stage) uint {
	return stage.GongEnumMap_Staged_Order[gongenum]
}

func (gongenum *GongEnum) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongEnums_referenceOrder[gongenum]
}

func (gongenumvalue *GongEnumValue) GongGetOrder(stage *Stage) uint {
	return stage.GongEnumValueMap_Staged_Order[gongenumvalue]
}

func (gongenumvalue *GongEnumValue) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongEnumValues_referenceOrder[gongenumvalue]
}

func (gonglink *GongLink) GongGetOrder(stage *Stage) uint {
	return stage.GongLinkMap_Staged_Order[gonglink]
}

func (gonglink *GongLink) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongLinks_referenceOrder[gonglink]
}

func (gongnote *GongNote) GongGetOrder(stage *Stage) uint {
	return stage.GongNoteMap_Staged_Order[gongnote]
}

func (gongnote *GongNote) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongNotes_referenceOrder[gongnote]
}

func (gongstruct *GongStruct) GongGetOrder(stage *Stage) uint {
	return stage.GongStructMap_Staged_Order[gongstruct]
}

func (gongstruct *GongStruct) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongStructs_referenceOrder[gongstruct]
}

func (gongtimefield *GongTimeField) GongGetOrder(stage *Stage) uint {
	return stage.GongTimeFieldMap_Staged_Order[gongtimefield]
}

func (gongtimefield *GongTimeField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongTimeFields_referenceOrder[gongtimefield]
}

func (metareference *MetaReference) GongGetOrder(stage *Stage) uint {
	return stage.MetaReferenceMap_Staged_Order[metareference]
}

func (metareference *MetaReference) GongGetReferenceOrder(stage *Stage) uint {
	return stage.MetaReferences_referenceOrder[metareference]
}

func (modelpkg *ModelPkg) GongGetOrder(stage *Stage) uint {
	return stage.ModelPkgMap_Staged_Order[modelpkg]
}

func (modelpkg *ModelPkg) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ModelPkgs_referenceOrder[modelpkg]
}

func (pointertogongstructfield *PointerToGongStructField) GongGetOrder(stage *Stage) uint {
	return stage.PointerToGongStructFieldMap_Staged_Order[pointertogongstructfield]
}

func (pointertogongstructfield *PointerToGongStructField) GongGetReferenceOrder(stage *Stage) uint {
	return stage.PointerToGongStructFields_referenceOrder[pointertogongstructfield]
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongGetOrder(stage *Stage) uint {
	return stage.SliceOfPointerToGongStructFieldMap_Staged_Order[sliceofpointertogongstructfield]
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
