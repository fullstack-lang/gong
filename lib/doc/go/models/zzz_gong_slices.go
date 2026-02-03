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
	// Compute reverse map for named struct AttributeShape
	// insertion point per field

	// Compute reverse map for named struct Classdiagram
	// insertion point per field
	stage.Classdiagram_GongStructShapes_reverseMap = make(map[*GongStructShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			stage.Classdiagram_GongStructShapes_reverseMap[_gongstructshape] = classdiagram
		}
	}
	stage.Classdiagram_GongEnumShapes_reverseMap = make(map[*GongEnumShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			stage.Classdiagram_GongEnumShapes_reverseMap[_gongenumshape] = classdiagram
		}
	}
	stage.Classdiagram_GongNoteShapes_reverseMap = make(map[*GongNoteShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			stage.Classdiagram_GongNoteShapes_reverseMap[_gongnoteshape] = classdiagram
		}
	}

	// Compute reverse map for named struct DiagramPackage
	// insertion point per field
	stage.DiagramPackage_Classdiagrams_reverseMap = make(map[*Classdiagram]*DiagramPackage)
	for diagrampackage := range stage.DiagramPackages {
		_ = diagrampackage
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			stage.DiagramPackage_Classdiagrams_reverseMap[_classdiagram] = diagrampackage
		}
	}

	// Compute reverse map for named struct GongEnumShape
	// insertion point per field
	stage.GongEnumShape_GongEnumValueShapes_reverseMap = make(map[*GongEnumValueShape]*GongEnumShape)
	for gongenumshape := range stage.GongEnumShapes {
		_ = gongenumshape
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			stage.GongEnumShape_GongEnumValueShapes_reverseMap[_gongenumvalueshape] = gongenumshape
		}
	}

	// Compute reverse map for named struct GongEnumValueShape
	// insertion point per field

	// Compute reverse map for named struct GongNoteLinkShape
	// insertion point per field

	// Compute reverse map for named struct GongNoteShape
	// insertion point per field
	stage.GongNoteShape_GongNoteLinkShapes_reverseMap = make(map[*GongNoteLinkShape]*GongNoteShape)
	for gongnoteshape := range stage.GongNoteShapes {
		_ = gongnoteshape
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			stage.GongNoteShape_GongNoteLinkShapes_reverseMap[_gongnotelinkshape] = gongnoteshape
		}
	}

	// Compute reverse map for named struct GongStructShape
	// insertion point per field
	stage.GongStructShape_AttributeShapes_reverseMap = make(map[*AttributeShape]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			stage.GongStructShape_AttributeShapes_reverseMap[_attributeshape] = gongstructshape
		}
	}
	stage.GongStructShape_LinkShapes_reverseMap = make(map[*LinkShape]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _linkshape := range gongstructshape.LinkShapes {
			stage.GongStructShape_LinkShapes_reverseMap[_linkshape] = gongstructshape
		}
	}

	// Compute reverse map for named struct LinkShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.AttributeShapes {
		res = append(res, instance)
	}

	for instance := range stage.Classdiagrams {
		res = append(res, instance)
	}

	for instance := range stage.DiagramPackages {
		res = append(res, instance)
	}

	for instance := range stage.GongEnumShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongEnumValueShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongNoteLinkShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongNoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongStructShapes {
		res = append(res, instance)
	}

	for instance := range stage.LinkShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (attributeshape *AttributeShape) GongCopy() GongstructIF {
	newInstance := *attributeshape
	return &newInstance
}

func (classdiagram *Classdiagram) GongCopy() GongstructIF {
	newInstance := *classdiagram
	return &newInstance
}

func (diagrampackage *DiagramPackage) GongCopy() GongstructIF {
	newInstance := *diagrampackage
	return &newInstance
}

func (gongenumshape *GongEnumShape) GongCopy() GongstructIF {
	newInstance := *gongenumshape
	return &newInstance
}

func (gongenumvalueshape *GongEnumValueShape) GongCopy() GongstructIF {
	newInstance := *gongenumvalueshape
	return &newInstance
}

func (gongnotelinkshape *GongNoteLinkShape) GongCopy() GongstructIF {
	newInstance := *gongnotelinkshape
	return &newInstance
}

func (gongnoteshape *GongNoteShape) GongCopy() GongstructIF {
	newInstance := *gongnoteshape
	return &newInstance
}

func (gongstructshape *GongStructShape) GongCopy() GongstructIF {
	newInstance := *gongstructshape
	return &newInstance
}

func (linkshape *LinkShape) GongCopy() GongstructIF {
	newInstance := *linkshape
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
	var attributeshapes_newInstances []*AttributeShape
	var attributeshapes_deletedInstances []*AttributeShape

	// parse all staged instances and check if they have a reference
	for attributeshape := range stage.AttributeShapes {
		if ref, ok := stage.AttributeShapes_reference[attributeshape]; !ok {
			attributeshapes_newInstances = append(attributeshapes_newInstances, attributeshape)
			newInstancesSlice = append(newInstancesSlice, attributeshape.GongMarshallIdentifier(stage))
			if stage.AttributeShapes_referenceOrder == nil {
				stage.AttributeShapes_referenceOrder = make(map[*AttributeShape]uint)
			}
			stage.AttributeShapes_referenceOrder[attributeshape] = stage.AttributeShapeMap_Staged_Order[attributeshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attributeshape.GongMarshallUnstaging(stage))
			delete(stage.AttributeShapes_referenceOrder, attributeshape)
			fieldInitializers, pointersInitializations := attributeshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AttributeShapeMap_Staged_Order[ref] = stage.AttributeShapeMap_Staged_Order[attributeshape]
			diffs := attributeshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attributeshape)
			delete(stage.AttributeShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attributeshape.GetName())
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
	for ref := range stage.AttributeShapes_reference {
		if _, ok := stage.AttributeShapes[ref]; !ok {
			attributeshapes_deletedInstances = append(attributeshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attributeshapes_newInstances)
	lenDeletedInstances += len(attributeshapes_deletedInstances)
	var classdiagrams_newInstances []*Classdiagram
	var classdiagrams_deletedInstances []*Classdiagram

	// parse all staged instances and check if they have a reference
	for classdiagram := range stage.Classdiagrams {
		if ref, ok := stage.Classdiagrams_reference[classdiagram]; !ok {
			classdiagrams_newInstances = append(classdiagrams_newInstances, classdiagram)
			newInstancesSlice = append(newInstancesSlice, classdiagram.GongMarshallIdentifier(stage))
			if stage.Classdiagrams_referenceOrder == nil {
				stage.Classdiagrams_referenceOrder = make(map[*Classdiagram]uint)
			}
			stage.Classdiagrams_referenceOrder[classdiagram] = stage.ClassdiagramMap_Staged_Order[classdiagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, classdiagram.GongMarshallUnstaging(stage))
			delete(stage.Classdiagrams_referenceOrder, classdiagram)
			fieldInitializers, pointersInitializations := classdiagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ClassdiagramMap_Staged_Order[ref] = stage.ClassdiagramMap_Staged_Order[classdiagram]
			diffs := classdiagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, classdiagram)
			delete(stage.ClassdiagramMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", classdiagram.GetName())
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
	for ref := range stage.Classdiagrams_reference {
		if _, ok := stage.Classdiagrams[ref]; !ok {
			classdiagrams_deletedInstances = append(classdiagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(classdiagrams_newInstances)
	lenDeletedInstances += len(classdiagrams_deletedInstances)
	var diagrampackages_newInstances []*DiagramPackage
	var diagrampackages_deletedInstances []*DiagramPackage

	// parse all staged instances and check if they have a reference
	for diagrampackage := range stage.DiagramPackages {
		if ref, ok := stage.DiagramPackages_reference[diagrampackage]; !ok {
			diagrampackages_newInstances = append(diagrampackages_newInstances, diagrampackage)
			newInstancesSlice = append(newInstancesSlice, diagrampackage.GongMarshallIdentifier(stage))
			if stage.DiagramPackages_referenceOrder == nil {
				stage.DiagramPackages_referenceOrder = make(map[*DiagramPackage]uint)
			}
			stage.DiagramPackages_referenceOrder[diagrampackage] = stage.DiagramPackageMap_Staged_Order[diagrampackage]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagrampackage.GongMarshallUnstaging(stage))
			delete(stage.DiagramPackages_referenceOrder, diagrampackage)
			fieldInitializers, pointersInitializations := diagrampackage.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramPackageMap_Staged_Order[ref] = stage.DiagramPackageMap_Staged_Order[diagrampackage]
			diffs := diagrampackage.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagrampackage)
			delete(stage.DiagramPackageMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", diagrampackage.GetName())
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
	for ref := range stage.DiagramPackages_reference {
		if _, ok := stage.DiagramPackages[ref]; !ok {
			diagrampackages_deletedInstances = append(diagrampackages_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagrampackages_newInstances)
	lenDeletedInstances += len(diagrampackages_deletedInstances)
	var gongenumshapes_newInstances []*GongEnumShape
	var gongenumshapes_deletedInstances []*GongEnumShape

	// parse all staged instances and check if they have a reference
	for gongenumshape := range stage.GongEnumShapes {
		if ref, ok := stage.GongEnumShapes_reference[gongenumshape]; !ok {
			gongenumshapes_newInstances = append(gongenumshapes_newInstances, gongenumshape)
			newInstancesSlice = append(newInstancesSlice, gongenumshape.GongMarshallIdentifier(stage))
			if stage.GongEnumShapes_referenceOrder == nil {
				stage.GongEnumShapes_referenceOrder = make(map[*GongEnumShape]uint)
			}
			stage.GongEnumShapes_referenceOrder[gongenumshape] = stage.GongEnumShapeMap_Staged_Order[gongenumshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongenumshape.GongMarshallUnstaging(stage))
			delete(stage.GongEnumShapes_referenceOrder, gongenumshape)
			fieldInitializers, pointersInitializations := gongenumshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongEnumShapeMap_Staged_Order[ref] = stage.GongEnumShapeMap_Staged_Order[gongenumshape]
			diffs := gongenumshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongenumshape)
			delete(stage.GongEnumShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongenumshape.GetName())
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
	for ref := range stage.GongEnumShapes_reference {
		if _, ok := stage.GongEnumShapes[ref]; !ok {
			gongenumshapes_deletedInstances = append(gongenumshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongenumshapes_newInstances)
	lenDeletedInstances += len(gongenumshapes_deletedInstances)
	var gongenumvalueshapes_newInstances []*GongEnumValueShape
	var gongenumvalueshapes_deletedInstances []*GongEnumValueShape

	// parse all staged instances and check if they have a reference
	for gongenumvalueshape := range stage.GongEnumValueShapes {
		if ref, ok := stage.GongEnumValueShapes_reference[gongenumvalueshape]; !ok {
			gongenumvalueshapes_newInstances = append(gongenumvalueshapes_newInstances, gongenumvalueshape)
			newInstancesSlice = append(newInstancesSlice, gongenumvalueshape.GongMarshallIdentifier(stage))
			if stage.GongEnumValueShapes_referenceOrder == nil {
				stage.GongEnumValueShapes_referenceOrder = make(map[*GongEnumValueShape]uint)
			}
			stage.GongEnumValueShapes_referenceOrder[gongenumvalueshape] = stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongenumvalueshape.GongMarshallUnstaging(stage))
			delete(stage.GongEnumValueShapes_referenceOrder, gongenumvalueshape)
			fieldInitializers, pointersInitializations := gongenumvalueshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongEnumValueShapeMap_Staged_Order[ref] = stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshape]
			diffs := gongenumvalueshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongenumvalueshape)
			delete(stage.GongEnumValueShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongenumvalueshape.GetName())
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
	for ref := range stage.GongEnumValueShapes_reference {
		if _, ok := stage.GongEnumValueShapes[ref]; !ok {
			gongenumvalueshapes_deletedInstances = append(gongenumvalueshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongenumvalueshapes_newInstances)
	lenDeletedInstances += len(gongenumvalueshapes_deletedInstances)
	var gongnotelinkshapes_newInstances []*GongNoteLinkShape
	var gongnotelinkshapes_deletedInstances []*GongNoteLinkShape

	// parse all staged instances and check if they have a reference
	for gongnotelinkshape := range stage.GongNoteLinkShapes {
		if ref, ok := stage.GongNoteLinkShapes_reference[gongnotelinkshape]; !ok {
			gongnotelinkshapes_newInstances = append(gongnotelinkshapes_newInstances, gongnotelinkshape)
			newInstancesSlice = append(newInstancesSlice, gongnotelinkshape.GongMarshallIdentifier(stage))
			if stage.GongNoteLinkShapes_referenceOrder == nil {
				stage.GongNoteLinkShapes_referenceOrder = make(map[*GongNoteLinkShape]uint)
			}
			stage.GongNoteLinkShapes_referenceOrder[gongnotelinkshape] = stage.GongNoteLinkShapeMap_Staged_Order[gongnotelinkshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongnotelinkshape.GongMarshallUnstaging(stage))
			delete(stage.GongNoteLinkShapes_referenceOrder, gongnotelinkshape)
			fieldInitializers, pointersInitializations := gongnotelinkshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongNoteLinkShapeMap_Staged_Order[ref] = stage.GongNoteLinkShapeMap_Staged_Order[gongnotelinkshape]
			diffs := gongnotelinkshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongnotelinkshape)
			delete(stage.GongNoteLinkShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongnotelinkshape.GetName())
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
	for ref := range stage.GongNoteLinkShapes_reference {
		if _, ok := stage.GongNoteLinkShapes[ref]; !ok {
			gongnotelinkshapes_deletedInstances = append(gongnotelinkshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongnotelinkshapes_newInstances)
	lenDeletedInstances += len(gongnotelinkshapes_deletedInstances)
	var gongnoteshapes_newInstances []*GongNoteShape
	var gongnoteshapes_deletedInstances []*GongNoteShape

	// parse all staged instances and check if they have a reference
	for gongnoteshape := range stage.GongNoteShapes {
		if ref, ok := stage.GongNoteShapes_reference[gongnoteshape]; !ok {
			gongnoteshapes_newInstances = append(gongnoteshapes_newInstances, gongnoteshape)
			newInstancesSlice = append(newInstancesSlice, gongnoteshape.GongMarshallIdentifier(stage))
			if stage.GongNoteShapes_referenceOrder == nil {
				stage.GongNoteShapes_referenceOrder = make(map[*GongNoteShape]uint)
			}
			stage.GongNoteShapes_referenceOrder[gongnoteshape] = stage.GongNoteShapeMap_Staged_Order[gongnoteshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongnoteshape.GongMarshallUnstaging(stage))
			delete(stage.GongNoteShapes_referenceOrder, gongnoteshape)
			fieldInitializers, pointersInitializations := gongnoteshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongNoteShapeMap_Staged_Order[ref] = stage.GongNoteShapeMap_Staged_Order[gongnoteshape]
			diffs := gongnoteshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongnoteshape)
			delete(stage.GongNoteShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongnoteshape.GetName())
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
	for ref := range stage.GongNoteShapes_reference {
		if _, ok := stage.GongNoteShapes[ref]; !ok {
			gongnoteshapes_deletedInstances = append(gongnoteshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongnoteshapes_newInstances)
	lenDeletedInstances += len(gongnoteshapes_deletedInstances)
	var gongstructshapes_newInstances []*GongStructShape
	var gongstructshapes_deletedInstances []*GongStructShape

	// parse all staged instances and check if they have a reference
	for gongstructshape := range stage.GongStructShapes {
		if ref, ok := stage.GongStructShapes_reference[gongstructshape]; !ok {
			gongstructshapes_newInstances = append(gongstructshapes_newInstances, gongstructshape)
			newInstancesSlice = append(newInstancesSlice, gongstructshape.GongMarshallIdentifier(stage))
			if stage.GongStructShapes_referenceOrder == nil {
				stage.GongStructShapes_referenceOrder = make(map[*GongStructShape]uint)
			}
			stage.GongStructShapes_referenceOrder[gongstructshape] = stage.GongStructShapeMap_Staged_Order[gongstructshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, gongstructshape.GongMarshallUnstaging(stage))
			delete(stage.GongStructShapes_referenceOrder, gongstructshape)
			fieldInitializers, pointersInitializations := gongstructshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GongStructShapeMap_Staged_Order[ref] = stage.GongStructShapeMap_Staged_Order[gongstructshape]
			diffs := gongstructshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, gongstructshape)
			delete(stage.GongStructShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", gongstructshape.GetName())
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
	for ref := range stage.GongStructShapes_reference {
		if _, ok := stage.GongStructShapes[ref]; !ok {
			gongstructshapes_deletedInstances = append(gongstructshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(gongstructshapes_newInstances)
	lenDeletedInstances += len(gongstructshapes_deletedInstances)
	var linkshapes_newInstances []*LinkShape
	var linkshapes_deletedInstances []*LinkShape

	// parse all staged instances and check if they have a reference
	for linkshape := range stage.LinkShapes {
		if ref, ok := stage.LinkShapes_reference[linkshape]; !ok {
			linkshapes_newInstances = append(linkshapes_newInstances, linkshape)
			newInstancesSlice = append(newInstancesSlice, linkshape.GongMarshallIdentifier(stage))
			if stage.LinkShapes_referenceOrder == nil {
				stage.LinkShapes_referenceOrder = make(map[*LinkShape]uint)
			}
			stage.LinkShapes_referenceOrder[linkshape] = stage.LinkShapeMap_Staged_Order[linkshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, linkshape.GongMarshallUnstaging(stage))
			delete(stage.LinkShapes_referenceOrder, linkshape)
			fieldInitializers, pointersInitializations := linkshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LinkShapeMap_Staged_Order[ref] = stage.LinkShapeMap_Staged_Order[linkshape]
			diffs := linkshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, linkshape)
			delete(stage.LinkShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", linkshape.GetName())
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
	for ref := range stage.LinkShapes_reference {
		if _, ok := stage.LinkShapes[ref]; !ok {
			linkshapes_deletedInstances = append(linkshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(linkshapes_newInstances)
	lenDeletedInstances += len(linkshapes_deletedInstances)

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
	stage.AttributeShapes_reference = make(map[*AttributeShape]*AttributeShape)
	stage.AttributeShapes_referenceOrder = make(map[*AttributeShape]uint) // diff Unstage needs the reference order
	for instance := range stage.AttributeShapes {
		stage.AttributeShapes_reference[instance] = instance.GongCopy().(*AttributeShape)
		stage.AttributeShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Classdiagrams_reference = make(map[*Classdiagram]*Classdiagram)
	stage.Classdiagrams_referenceOrder = make(map[*Classdiagram]uint) // diff Unstage needs the reference order
	for instance := range stage.Classdiagrams {
		stage.Classdiagrams_reference[instance] = instance.GongCopy().(*Classdiagram)
		stage.Classdiagrams_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DiagramPackages_reference = make(map[*DiagramPackage]*DiagramPackage)
	stage.DiagramPackages_referenceOrder = make(map[*DiagramPackage]uint) // diff Unstage needs the reference order
	for instance := range stage.DiagramPackages {
		stage.DiagramPackages_reference[instance] = instance.GongCopy().(*DiagramPackage)
		stage.DiagramPackages_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongEnumShapes_reference = make(map[*GongEnumShape]*GongEnumShape)
	stage.GongEnumShapes_referenceOrder = make(map[*GongEnumShape]uint) // diff Unstage needs the reference order
	for instance := range stage.GongEnumShapes {
		stage.GongEnumShapes_reference[instance] = instance.GongCopy().(*GongEnumShape)
		stage.GongEnumShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongEnumValueShapes_reference = make(map[*GongEnumValueShape]*GongEnumValueShape)
	stage.GongEnumValueShapes_referenceOrder = make(map[*GongEnumValueShape]uint) // diff Unstage needs the reference order
	for instance := range stage.GongEnumValueShapes {
		stage.GongEnumValueShapes_reference[instance] = instance.GongCopy().(*GongEnumValueShape)
		stage.GongEnumValueShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongNoteLinkShapes_reference = make(map[*GongNoteLinkShape]*GongNoteLinkShape)
	stage.GongNoteLinkShapes_referenceOrder = make(map[*GongNoteLinkShape]uint) // diff Unstage needs the reference order
	for instance := range stage.GongNoteLinkShapes {
		stage.GongNoteLinkShapes_reference[instance] = instance.GongCopy().(*GongNoteLinkShape)
		stage.GongNoteLinkShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongNoteShapes_reference = make(map[*GongNoteShape]*GongNoteShape)
	stage.GongNoteShapes_referenceOrder = make(map[*GongNoteShape]uint) // diff Unstage needs the reference order
	for instance := range stage.GongNoteShapes {
		stage.GongNoteShapes_reference[instance] = instance.GongCopy().(*GongNoteShape)
		stage.GongNoteShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GongStructShapes_reference = make(map[*GongStructShape]*GongStructShape)
	stage.GongStructShapes_referenceOrder = make(map[*GongStructShape]uint) // diff Unstage needs the reference order
	for instance := range stage.GongStructShapes {
		stage.GongStructShapes_reference[instance] = instance.GongCopy().(*GongStructShape)
		stage.GongStructShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.LinkShapes_reference = make(map[*LinkShape]*LinkShape)
	stage.LinkShapes_referenceOrder = make(map[*LinkShape]uint) // diff Unstage needs the reference order
	for instance := range stage.LinkShapes {
		stage.LinkShapes_reference[instance] = instance.GongCopy().(*LinkShape)
		stage.LinkShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
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
func (attributeshape *AttributeShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AttributeShapeMap_Staged_Order[attributeshape]; ok {
		return order
	}
	return stage.AttributeShapes_referenceOrder[attributeshape]
}

func (attributeshape *AttributeShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.AttributeShapes_referenceOrder[attributeshape]
}

func (classdiagram *Classdiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ClassdiagramMap_Staged_Order[classdiagram]; ok {
		return order
	}
	return stage.Classdiagrams_referenceOrder[classdiagram]
}

func (classdiagram *Classdiagram) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Classdiagrams_referenceOrder[classdiagram]
}

func (diagrampackage *DiagramPackage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramPackageMap_Staged_Order[diagrampackage]; ok {
		return order
	}
	return stage.DiagramPackages_referenceOrder[diagrampackage]
}

func (diagrampackage *DiagramPackage) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DiagramPackages_referenceOrder[diagrampackage]
}

func (gongenumshape *GongEnumShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnumShapeMap_Staged_Order[gongenumshape]; ok {
		return order
	}
	return stage.GongEnumShapes_referenceOrder[gongenumshape]
}

func (gongenumshape *GongEnumShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongEnumShapes_referenceOrder[gongenumshape]
}

func (gongenumvalueshape *GongEnumValueShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnumValueShapeMap_Staged_Order[gongenumvalueshape]; ok {
		return order
	}
	return stage.GongEnumValueShapes_referenceOrder[gongenumvalueshape]
}

func (gongenumvalueshape *GongEnumValueShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongEnumValueShapes_referenceOrder[gongenumvalueshape]
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongNoteLinkShapeMap_Staged_Order[gongnotelinkshape]; ok {
		return order
	}
	return stage.GongNoteLinkShapes_referenceOrder[gongnotelinkshape]
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongNoteLinkShapes_referenceOrder[gongnotelinkshape]
}

func (gongnoteshape *GongNoteShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongNoteShapeMap_Staged_Order[gongnoteshape]; ok {
		return order
	}
	return stage.GongNoteShapes_referenceOrder[gongnoteshape]
}

func (gongnoteshape *GongNoteShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongNoteShapes_referenceOrder[gongnoteshape]
}

func (gongstructshape *GongStructShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongStructShapeMap_Staged_Order[gongstructshape]; ok {
		return order
	}
	return stage.GongStructShapes_referenceOrder[gongstructshape]
}

func (gongstructshape *GongStructShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GongStructShapes_referenceOrder[gongstructshape]
}

func (linkshape *LinkShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkShapeMap_Staged_Order[linkshape]; ok {
		return order
	}
	return stage.LinkShapes_referenceOrder[linkshape]
}

func (linkshape *LinkShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.LinkShapes_referenceOrder[linkshape]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (attributeshape *AttributeShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributeshape.GongGetGongstructName(), attributeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributeshape *AttributeShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributeshape.GongGetGongstructName(), attributeshape.GongGetReferenceOrder(stage))
}

func (classdiagram *Classdiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", classdiagram.GongGetGongstructName(), classdiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (classdiagram *Classdiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", classdiagram.GongGetGongstructName(), classdiagram.GongGetReferenceOrder(stage))
}

func (diagrampackage *DiagramPackage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagrampackage.GongGetGongstructName(), diagrampackage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagrampackage *DiagramPackage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagrampackage.GongGetGongstructName(), diagrampackage.GongGetReferenceOrder(stage))
}

func (gongenumshape *GongEnumShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumshape.GongGetGongstructName(), gongenumshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumshape *GongEnumShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumshape.GongGetGongstructName(), gongenumshape.GongGetReferenceOrder(stage))
}

func (gongenumvalueshape *GongEnumValueShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalueshape.GongGetGongstructName(), gongenumvalueshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumvalueshape *GongEnumValueShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalueshape.GongGetGongstructName(), gongenumvalueshape.GongGetReferenceOrder(stage))
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnotelinkshape.GongGetGongstructName(), gongnotelinkshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnotelinkshape *GongNoteLinkShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnotelinkshape.GongGetGongstructName(), gongnotelinkshape.GongGetReferenceOrder(stage))
}

func (gongnoteshape *GongNoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnoteshape.GongGetGongstructName(), gongnoteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnoteshape *GongNoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnoteshape.GongGetGongstructName(), gongnoteshape.GongGetReferenceOrder(stage))
}

func (gongstructshape *GongStructShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstructshape.GongGetGongstructName(), gongstructshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongstructshape *GongStructShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstructshape.GongGetGongstructName(), gongstructshape.GongGetReferenceOrder(stage))
}

func (linkshape *LinkShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkshape.GongGetGongstructName(), linkshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkshape *LinkShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkshape.GongGetGongstructName(), linkshape.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (attributeshape *AttributeShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AttributeShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attributeshape.Name)
	return
}

func (classdiagram *Classdiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Classdiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", classdiagram.Name)
	return
}

func (diagrampackage *DiagramPackage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramPackage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", diagrampackage.Name)
	return
}

func (gongenumshape *GongEnumShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongenumshape.Name)
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumValueShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongenumvalueshape.Name)
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNoteLinkShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongnotelinkshape.Name)
	return
}

func (gongnoteshape *GongNoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongnoteshape.Name)
	return
}

func (gongstructshape *GongStructShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongStructShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", gongstructshape.Name)
	return
}

func (linkshape *LinkShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LinkShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", linkshape.Name)
	return
}

// insertion point for unstaging
func (attributeshape *AttributeShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributeshape.GongGetReferenceIdentifier(stage))
	return
}

func (classdiagram *Classdiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", classdiagram.GongGetReferenceIdentifier(stage))
	return
}

func (diagrampackage *DiagramPackage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagrampackage.GongGetReferenceIdentifier(stage))
	return
}

func (gongenumshape *GongEnumShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumvalueshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnotelinkshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongnoteshape *GongNoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnoteshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongstructshape *GongStructShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongstructshape.GongGetReferenceIdentifier(stage))
	return
}

func (linkshape *LinkShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkshape.GongGetReferenceIdentifier(stage))
	return
}

// end of template
