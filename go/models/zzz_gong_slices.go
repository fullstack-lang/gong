// generated code - do not edit
package models

import (
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

	// insertion point per named struct
	var gongbasicfields_newInstances []*GongBasicField
	var gongbasicfields_deletedInstances []*GongBasicField

	// parse all staged instances and check if they have a reference
	for gongbasicfield := range stage.GongBasicFields {
		if ref, ok := stage.GongBasicFields_reference[gongbasicfield]; !ok {
			gongbasicfields_newInstances = append(gongbasicfields_newInstances, gongbasicfield)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GongBasicField "+gongbasicfield.Name,
				)
			}
		} else {
			diffs := gongbasicfield.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GongBasicField \""+gongbasicfield.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongbasicfield := range stage.GongBasicFields_reference {
		if _, ok := stage.GongBasicFields[gongbasicfield]; !ok {
			gongbasicfields_deletedInstances = append(gongbasicfields_deletedInstances, gongbasicfield)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GongBasicField "+gongbasicfield.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GongEnum "+gongenum.Name,
				)
			}
		} else {
			diffs := gongenum.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GongEnum \""+gongenum.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongenum := range stage.GongEnums_reference {
		if _, ok := stage.GongEnums[gongenum]; !ok {
			gongenums_deletedInstances = append(gongenums_deletedInstances, gongenum)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GongEnum "+gongenum.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GongEnumValue "+gongenumvalue.Name,
				)
			}
		} else {
			diffs := gongenumvalue.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GongEnumValue \""+gongenumvalue.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongenumvalue := range stage.GongEnumValues_reference {
		if _, ok := stage.GongEnumValues[gongenumvalue]; !ok {
			gongenumvalues_deletedInstances = append(gongenumvalues_deletedInstances, gongenumvalue)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GongEnumValue "+gongenumvalue.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GongLink "+gonglink.Name,
				)
			}
		} else {
			diffs := gonglink.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GongLink \""+gonglink.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gonglink := range stage.GongLinks_reference {
		if _, ok := stage.GongLinks[gonglink]; !ok {
			gonglinks_deletedInstances = append(gonglinks_deletedInstances, gonglink)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GongLink "+gonglink.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GongNote "+gongnote.Name,
				)
			}
		} else {
			diffs := gongnote.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GongNote \""+gongnote.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongnote := range stage.GongNotes_reference {
		if _, ok := stage.GongNotes[gongnote]; !ok {
			gongnotes_deletedInstances = append(gongnotes_deletedInstances, gongnote)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GongNote "+gongnote.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GongStruct "+gongstruct.Name,
				)
			}
		} else {
			diffs := gongstruct.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GongStruct \""+gongstruct.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongstruct := range stage.GongStructs_reference {
		if _, ok := stage.GongStructs[gongstruct]; !ok {
			gongstructs_deletedInstances = append(gongstructs_deletedInstances, gongstruct)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GongStruct "+gongstruct.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of GongTimeField "+gongtimefield.Name,
				)
			}
		} else {
			diffs := gongtimefield.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of GongTimeField \""+gongtimefield.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongtimefield := range stage.GongTimeFields_reference {
		if _, ok := stage.GongTimeFields[gongtimefield]; !ok {
			gongtimefields_deletedInstances = append(gongtimefields_deletedInstances, gongtimefield)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of GongTimeField "+gongtimefield.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of MetaReference "+metareference.Name,
				)
			}
		} else {
			diffs := metareference.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of MetaReference \""+metareference.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for metareference := range stage.MetaReferences_reference {
		if _, ok := stage.MetaReferences[metareference]; !ok {
			metareferences_deletedInstances = append(metareferences_deletedInstances, metareference)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of MetaReference "+metareference.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of ModelPkg "+modelpkg.Name,
				)
			}
		} else {
			diffs := modelpkg.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of ModelPkg \""+modelpkg.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for modelpkg := range stage.ModelPkgs_reference {
		if _, ok := stage.ModelPkgs[modelpkg]; !ok {
			modelpkgs_deletedInstances = append(modelpkgs_deletedInstances, modelpkg)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of ModelPkg "+modelpkg.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of PointerToGongStructField "+pointertogongstructfield.Name,
				)
			}
		} else {
			diffs := pointertogongstructfield.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of PointerToGongStructField \""+pointertogongstructfield.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for pointertogongstructfield := range stage.PointerToGongStructFields_reference {
		if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; !ok {
			pointertogongstructfields_deletedInstances = append(pointertogongstructfields_deletedInstances, pointertogongstructfield)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of PointerToGongStructField "+pointertogongstructfield.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of SliceOfPointerToGongStructField "+sliceofpointertogongstructfield.Name,
				)
			}
		} else {
			diffs := sliceofpointertogongstructfield.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of SliceOfPointerToGongStructField \""+sliceofpointertogongstructfield.Name + "\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for sliceofpointertogongstructfield := range stage.SliceOfPointerToGongStructFields_reference {
		if _, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]; !ok {
			sliceofpointertogongstructfields_deletedInstances = append(sliceofpointertogongstructfields_deletedInstances, sliceofpointertogongstructfield)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of SliceOfPointerToGongStructField "+sliceofpointertogongstructfield.Name,
				)
			}
		}
	}

	lenNewInstances += len(sliceofpointertogongstructfields_newInstances)
	lenDeletedInstances += len(sliceofpointertogongstructfields_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.GongBasicFields_reference = make(map[*GongBasicField]*GongBasicField)
	for instance := range stage.GongBasicFields {
		stage.GongBasicFields_reference[instance] = instance.GongCopy().(*GongBasicField)
	}

	stage.GongEnums_reference = make(map[*GongEnum]*GongEnum)
	for instance := range stage.GongEnums {
		stage.GongEnums_reference[instance] = instance.GongCopy().(*GongEnum)
	}

	stage.GongEnumValues_reference = make(map[*GongEnumValue]*GongEnumValue)
	for instance := range stage.GongEnumValues {
		stage.GongEnumValues_reference[instance] = instance.GongCopy().(*GongEnumValue)
	}

	stage.GongLinks_reference = make(map[*GongLink]*GongLink)
	for instance := range stage.GongLinks {
		stage.GongLinks_reference[instance] = instance.GongCopy().(*GongLink)
	}

	stage.GongNotes_reference = make(map[*GongNote]*GongNote)
	for instance := range stage.GongNotes {
		stage.GongNotes_reference[instance] = instance.GongCopy().(*GongNote)
	}

	stage.GongStructs_reference = make(map[*GongStruct]*GongStruct)
	for instance := range stage.GongStructs {
		stage.GongStructs_reference[instance] = instance.GongCopy().(*GongStruct)
	}

	stage.GongTimeFields_reference = make(map[*GongTimeField]*GongTimeField)
	for instance := range stage.GongTimeFields {
		stage.GongTimeFields_reference[instance] = instance.GongCopy().(*GongTimeField)
	}

	stage.MetaReferences_reference = make(map[*MetaReference]*MetaReference)
	for instance := range stage.MetaReferences {
		stage.MetaReferences_reference[instance] = instance.GongCopy().(*MetaReference)
	}

	stage.ModelPkgs_reference = make(map[*ModelPkg]*ModelPkg)
	for instance := range stage.ModelPkgs {
		stage.ModelPkgs_reference[instance] = instance.GongCopy().(*ModelPkg)
	}

	stage.PointerToGongStructFields_reference = make(map[*PointerToGongStructField]*PointerToGongStructField)
	for instance := range stage.PointerToGongStructFields {
		stage.PointerToGongStructFields_reference[instance] = instance.GongCopy().(*PointerToGongStructField)
	}

	stage.SliceOfPointerToGongStructFields_reference = make(map[*SliceOfPointerToGongStructField]*SliceOfPointerToGongStructField)
	for instance := range stage.SliceOfPointerToGongStructFields {
		stage.SliceOfPointerToGongStructFields_reference[instance] = instance.GongCopy().(*SliceOfPointerToGongStructField)
	}

}
