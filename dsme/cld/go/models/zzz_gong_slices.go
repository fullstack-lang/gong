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
	// Compute reverse map for named struct Category1
	// insertion point per field

	// Compute reverse map for named struct Category1Shape
	// insertion point per field

	// Compute reverse map for named struct Category2
	// insertion point per field

	// Compute reverse map for named struct Category2Shape
	// insertion point per field

	// Compute reverse map for named struct Category3
	// insertion point per field

	// Compute reverse map for named struct Category3Shape
	// insertion point per field

	// Compute reverse map for named struct ControlPointShape
	// insertion point per field

	// Compute reverse map for named struct Desk
	// insertion point per field

	// Compute reverse map for named struct Diagram
	// insertion point per field
	stage.Diagram_Category1Shapes_reverseMap = make(map[*Category1Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _category1shape := range diagram.Category1Shapes {
			stage.Diagram_Category1Shapes_reverseMap[_category1shape] = diagram
		}
	}
	stage.Diagram_Category2Shapes_reverseMap = make(map[*Category2Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _category2shape := range diagram.Category2Shapes {
			stage.Diagram_Category2Shapes_reverseMap[_category2shape] = diagram
		}
	}
	stage.Diagram_Category3Shapes_reverseMap = make(map[*Category3Shape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _category3shape := range diagram.Category3Shapes {
			stage.Diagram_Category3Shapes_reverseMap[_category3shape] = diagram
		}
	}
	stage.Diagram_InfluenceShapes_reverseMap = make(map[*InfluenceShape]*Diagram)
	for diagram := range stage.Diagrams {
		_ = diagram
		for _, _influenceshape := range diagram.InfluenceShapes {
			stage.Diagram_InfluenceShapes_reverseMap[_influenceshape] = diagram
		}
	}

	// Compute reverse map for named struct Influence
	// insertion point per field

	// Compute reverse map for named struct InfluenceShape
	// insertion point per field
	stage.InfluenceShape_ControlPointShapes_reverseMap = make(map[*ControlPointShape]*InfluenceShape)
	for influenceshape := range stage.InfluenceShapes {
		_ = influenceshape
		for _, _controlpointshape := range influenceshape.ControlPointShapes {
			stage.InfluenceShape_ControlPointShapes_reverseMap[_controlpointshape] = influenceshape
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Category1s {
		res = append(res, instance)
	}

	for instance := range stage.Category1Shapes {
		res = append(res, instance)
	}

	for instance := range stage.Category2s {
		res = append(res, instance)
	}

	for instance := range stage.Category2Shapes {
		res = append(res, instance)
	}

	for instance := range stage.Category3s {
		res = append(res, instance)
	}

	for instance := range stage.Category3Shapes {
		res = append(res, instance)
	}

	for instance := range stage.ControlPointShapes {
		res = append(res, instance)
	}

	for instance := range stage.Desks {
		res = append(res, instance)
	}

	for instance := range stage.Diagrams {
		res = append(res, instance)
	}

	for instance := range stage.Influences {
		res = append(res, instance)
	}

	for instance := range stage.InfluenceShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (category1 *Category1) GongCopy() GongstructIF {
	newInstance := *category1
	return &newInstance
}

func (category1shape *Category1Shape) GongCopy() GongstructIF {
	newInstance := *category1shape
	return &newInstance
}

func (category2 *Category2) GongCopy() GongstructIF {
	newInstance := *category2
	return &newInstance
}

func (category2shape *Category2Shape) GongCopy() GongstructIF {
	newInstance := *category2shape
	return &newInstance
}

func (category3 *Category3) GongCopy() GongstructIF {
	newInstance := *category3
	return &newInstance
}

func (category3shape *Category3Shape) GongCopy() GongstructIF {
	newInstance := *category3shape
	return &newInstance
}

func (controlpointshape *ControlPointShape) GongCopy() GongstructIF {
	newInstance := *controlpointshape
	return &newInstance
}

func (desk *Desk) GongCopy() GongstructIF {
	newInstance := *desk
	return &newInstance
}

func (diagram *Diagram) GongCopy() GongstructIF {
	newInstance := *diagram
	return &newInstance
}

func (influence *Influence) GongCopy() GongstructIF {
	newInstance := *influence
	return &newInstance
}

func (influenceshape *InfluenceShape) GongCopy() GongstructIF {
	newInstance := *influenceshape
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
	var category1s_newInstances []*Category1
	var category1s_deletedInstances []*Category1

	// parse all staged instances and check if they have a reference
	for category1 := range stage.Category1s {
		if ref, ok := stage.Category1s_reference[category1]; !ok {
			category1s_newInstances = append(category1s_newInstances, category1)
			newInstancesSlice = append(newInstancesSlice, category1.GongMarshallIdentifier(stage))
			if stage.Category1s_referenceOrder == nil {
				stage.Category1s_referenceOrder = make(map[*Category1]uint)
			}
			stage.Category1s_referenceOrder[category1] = stage.Category1Map_Staged_Order[category1]
			newInstancesReverseSlice = append(newInstancesReverseSlice, category1.GongMarshallUnstaging(stage))
			delete(stage.Category1s_referenceOrder, category1)
			fieldInitializers, pointersInitializations := category1.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Category1Map_Staged_Order[ref] = stage.Category1Map_Staged_Order[category1]
			diffs := category1.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, category1)
			delete(stage.Category1Map_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", category1.GetName())
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
	for ref := range stage.Category1s_reference {
		if _, ok := stage.Category1s[ref]; !ok {
			category1s_deletedInstances = append(category1s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(category1s_newInstances)
	lenDeletedInstances += len(category1s_deletedInstances)
	var category1shapes_newInstances []*Category1Shape
	var category1shapes_deletedInstances []*Category1Shape

	// parse all staged instances and check if they have a reference
	for category1shape := range stage.Category1Shapes {
		if ref, ok := stage.Category1Shapes_reference[category1shape]; !ok {
			category1shapes_newInstances = append(category1shapes_newInstances, category1shape)
			newInstancesSlice = append(newInstancesSlice, category1shape.GongMarshallIdentifier(stage))
			if stage.Category1Shapes_referenceOrder == nil {
				stage.Category1Shapes_referenceOrder = make(map[*Category1Shape]uint)
			}
			stage.Category1Shapes_referenceOrder[category1shape] = stage.Category1ShapeMap_Staged_Order[category1shape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, category1shape.GongMarshallUnstaging(stage))
			delete(stage.Category1Shapes_referenceOrder, category1shape)
			fieldInitializers, pointersInitializations := category1shape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Category1ShapeMap_Staged_Order[ref] = stage.Category1ShapeMap_Staged_Order[category1shape]
			diffs := category1shape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, category1shape)
			delete(stage.Category1ShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", category1shape.GetName())
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
	for ref := range stage.Category1Shapes_reference {
		if _, ok := stage.Category1Shapes[ref]; !ok {
			category1shapes_deletedInstances = append(category1shapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(category1shapes_newInstances)
	lenDeletedInstances += len(category1shapes_deletedInstances)
	var category2s_newInstances []*Category2
	var category2s_deletedInstances []*Category2

	// parse all staged instances and check if they have a reference
	for category2 := range stage.Category2s {
		if ref, ok := stage.Category2s_reference[category2]; !ok {
			category2s_newInstances = append(category2s_newInstances, category2)
			newInstancesSlice = append(newInstancesSlice, category2.GongMarshallIdentifier(stage))
			if stage.Category2s_referenceOrder == nil {
				stage.Category2s_referenceOrder = make(map[*Category2]uint)
			}
			stage.Category2s_referenceOrder[category2] = stage.Category2Map_Staged_Order[category2]
			newInstancesReverseSlice = append(newInstancesReverseSlice, category2.GongMarshallUnstaging(stage))
			delete(stage.Category2s_referenceOrder, category2)
			fieldInitializers, pointersInitializations := category2.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Category2Map_Staged_Order[ref] = stage.Category2Map_Staged_Order[category2]
			diffs := category2.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, category2)
			delete(stage.Category2Map_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", category2.GetName())
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
	for ref := range stage.Category2s_reference {
		if _, ok := stage.Category2s[ref]; !ok {
			category2s_deletedInstances = append(category2s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(category2s_newInstances)
	lenDeletedInstances += len(category2s_deletedInstances)
	var category2shapes_newInstances []*Category2Shape
	var category2shapes_deletedInstances []*Category2Shape

	// parse all staged instances and check if they have a reference
	for category2shape := range stage.Category2Shapes {
		if ref, ok := stage.Category2Shapes_reference[category2shape]; !ok {
			category2shapes_newInstances = append(category2shapes_newInstances, category2shape)
			newInstancesSlice = append(newInstancesSlice, category2shape.GongMarshallIdentifier(stage))
			if stage.Category2Shapes_referenceOrder == nil {
				stage.Category2Shapes_referenceOrder = make(map[*Category2Shape]uint)
			}
			stage.Category2Shapes_referenceOrder[category2shape] = stage.Category2ShapeMap_Staged_Order[category2shape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, category2shape.GongMarshallUnstaging(stage))
			delete(stage.Category2Shapes_referenceOrder, category2shape)
			fieldInitializers, pointersInitializations := category2shape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Category2ShapeMap_Staged_Order[ref] = stage.Category2ShapeMap_Staged_Order[category2shape]
			diffs := category2shape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, category2shape)
			delete(stage.Category2ShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", category2shape.GetName())
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
	for ref := range stage.Category2Shapes_reference {
		if _, ok := stage.Category2Shapes[ref]; !ok {
			category2shapes_deletedInstances = append(category2shapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(category2shapes_newInstances)
	lenDeletedInstances += len(category2shapes_deletedInstances)
	var category3s_newInstances []*Category3
	var category3s_deletedInstances []*Category3

	// parse all staged instances and check if they have a reference
	for category3 := range stage.Category3s {
		if ref, ok := stage.Category3s_reference[category3]; !ok {
			category3s_newInstances = append(category3s_newInstances, category3)
			newInstancesSlice = append(newInstancesSlice, category3.GongMarshallIdentifier(stage))
			if stage.Category3s_referenceOrder == nil {
				stage.Category3s_referenceOrder = make(map[*Category3]uint)
			}
			stage.Category3s_referenceOrder[category3] = stage.Category3Map_Staged_Order[category3]
			newInstancesReverseSlice = append(newInstancesReverseSlice, category3.GongMarshallUnstaging(stage))
			delete(stage.Category3s_referenceOrder, category3)
			fieldInitializers, pointersInitializations := category3.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Category3Map_Staged_Order[ref] = stage.Category3Map_Staged_Order[category3]
			diffs := category3.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, category3)
			delete(stage.Category3Map_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", category3.GetName())
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
	for ref := range stage.Category3s_reference {
		if _, ok := stage.Category3s[ref]; !ok {
			category3s_deletedInstances = append(category3s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(category3s_newInstances)
	lenDeletedInstances += len(category3s_deletedInstances)
	var category3shapes_newInstances []*Category3Shape
	var category3shapes_deletedInstances []*Category3Shape

	// parse all staged instances and check if they have a reference
	for category3shape := range stage.Category3Shapes {
		if ref, ok := stage.Category3Shapes_reference[category3shape]; !ok {
			category3shapes_newInstances = append(category3shapes_newInstances, category3shape)
			newInstancesSlice = append(newInstancesSlice, category3shape.GongMarshallIdentifier(stage))
			if stage.Category3Shapes_referenceOrder == nil {
				stage.Category3Shapes_referenceOrder = make(map[*Category3Shape]uint)
			}
			stage.Category3Shapes_referenceOrder[category3shape] = stage.Category3ShapeMap_Staged_Order[category3shape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, category3shape.GongMarshallUnstaging(stage))
			delete(stage.Category3Shapes_referenceOrder, category3shape)
			fieldInitializers, pointersInitializations := category3shape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Category3ShapeMap_Staged_Order[ref] = stage.Category3ShapeMap_Staged_Order[category3shape]
			diffs := category3shape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, category3shape)
			delete(stage.Category3ShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", category3shape.GetName())
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
	for ref := range stage.Category3Shapes_reference {
		if _, ok := stage.Category3Shapes[ref]; !ok {
			category3shapes_deletedInstances = append(category3shapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(category3shapes_newInstances)
	lenDeletedInstances += len(category3shapes_deletedInstances)
	var controlpointshapes_newInstances []*ControlPointShape
	var controlpointshapes_deletedInstances []*ControlPointShape

	// parse all staged instances and check if they have a reference
	for controlpointshape := range stage.ControlPointShapes {
		if ref, ok := stage.ControlPointShapes_reference[controlpointshape]; !ok {
			controlpointshapes_newInstances = append(controlpointshapes_newInstances, controlpointshape)
			newInstancesSlice = append(newInstancesSlice, controlpointshape.GongMarshallIdentifier(stage))
			if stage.ControlPointShapes_referenceOrder == nil {
				stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint)
			}
			stage.ControlPointShapes_referenceOrder[controlpointshape] = stage.ControlPointShapeMap_Staged_Order[controlpointshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, controlpointshape.GongMarshallUnstaging(stage))
			delete(stage.ControlPointShapes_referenceOrder, controlpointshape)
			fieldInitializers, pointersInitializations := controlpointshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ControlPointShapeMap_Staged_Order[ref] = stage.ControlPointShapeMap_Staged_Order[controlpointshape]
			diffs := controlpointshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, controlpointshape)
			delete(stage.ControlPointShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", controlpointshape.GetName())
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
	for ref := range stage.ControlPointShapes_reference {
		if _, ok := stage.ControlPointShapes[ref]; !ok {
			controlpointshapes_deletedInstances = append(controlpointshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(controlpointshapes_newInstances)
	lenDeletedInstances += len(controlpointshapes_deletedInstances)
	var desks_newInstances []*Desk
	var desks_deletedInstances []*Desk

	// parse all staged instances and check if they have a reference
	for desk := range stage.Desks {
		if ref, ok := stage.Desks_reference[desk]; !ok {
			desks_newInstances = append(desks_newInstances, desk)
			newInstancesSlice = append(newInstancesSlice, desk.GongMarshallIdentifier(stage))
			if stage.Desks_referenceOrder == nil {
				stage.Desks_referenceOrder = make(map[*Desk]uint)
			}
			stage.Desks_referenceOrder[desk] = stage.DeskMap_Staged_Order[desk]
			newInstancesReverseSlice = append(newInstancesReverseSlice, desk.GongMarshallUnstaging(stage))
			delete(stage.Desks_referenceOrder, desk)
			fieldInitializers, pointersInitializations := desk.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DeskMap_Staged_Order[ref] = stage.DeskMap_Staged_Order[desk]
			diffs := desk.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, desk)
			delete(stage.DeskMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", desk.GetName())
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
	for ref := range stage.Desks_reference {
		if _, ok := stage.Desks[ref]; !ok {
			desks_deletedInstances = append(desks_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(desks_newInstances)
	lenDeletedInstances += len(desks_deletedInstances)
	var diagrams_newInstances []*Diagram
	var diagrams_deletedInstances []*Diagram

	// parse all staged instances and check if they have a reference
	for diagram := range stage.Diagrams {
		if ref, ok := stage.Diagrams_reference[diagram]; !ok {
			diagrams_newInstances = append(diagrams_newInstances, diagram)
			newInstancesSlice = append(newInstancesSlice, diagram.GongMarshallIdentifier(stage))
			if stage.Diagrams_referenceOrder == nil {
				stage.Diagrams_referenceOrder = make(map[*Diagram]uint)
			}
			stage.Diagrams_referenceOrder[diagram] = stage.DiagramMap_Staged_Order[diagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagram.GongMarshallUnstaging(stage))
			delete(stage.Diagrams_referenceOrder, diagram)
			fieldInitializers, pointersInitializations := diagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramMap_Staged_Order[ref] = stage.DiagramMap_Staged_Order[diagram]
			diffs := diagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagram)
			delete(stage.DiagramMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", diagram.GetName())
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
	for ref := range stage.Diagrams_reference {
		if _, ok := stage.Diagrams[ref]; !ok {
			diagrams_deletedInstances = append(diagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagrams_newInstances)
	lenDeletedInstances += len(diagrams_deletedInstances)
	var influences_newInstances []*Influence
	var influences_deletedInstances []*Influence

	// parse all staged instances and check if they have a reference
	for influence := range stage.Influences {
		if ref, ok := stage.Influences_reference[influence]; !ok {
			influences_newInstances = append(influences_newInstances, influence)
			newInstancesSlice = append(newInstancesSlice, influence.GongMarshallIdentifier(stage))
			if stage.Influences_referenceOrder == nil {
				stage.Influences_referenceOrder = make(map[*Influence]uint)
			}
			stage.Influences_referenceOrder[influence] = stage.InfluenceMap_Staged_Order[influence]
			newInstancesReverseSlice = append(newInstancesReverseSlice, influence.GongMarshallUnstaging(stage))
			delete(stage.Influences_referenceOrder, influence)
			fieldInitializers, pointersInitializations := influence.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.InfluenceMap_Staged_Order[ref] = stage.InfluenceMap_Staged_Order[influence]
			diffs := influence.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, influence)
			delete(stage.InfluenceMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", influence.GetName())
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
	for ref := range stage.Influences_reference {
		if _, ok := stage.Influences[ref]; !ok {
			influences_deletedInstances = append(influences_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(influences_newInstances)
	lenDeletedInstances += len(influences_deletedInstances)
	var influenceshapes_newInstances []*InfluenceShape
	var influenceshapes_deletedInstances []*InfluenceShape

	// parse all staged instances and check if they have a reference
	for influenceshape := range stage.InfluenceShapes {
		if ref, ok := stage.InfluenceShapes_reference[influenceshape]; !ok {
			influenceshapes_newInstances = append(influenceshapes_newInstances, influenceshape)
			newInstancesSlice = append(newInstancesSlice, influenceshape.GongMarshallIdentifier(stage))
			if stage.InfluenceShapes_referenceOrder == nil {
				stage.InfluenceShapes_referenceOrder = make(map[*InfluenceShape]uint)
			}
			stage.InfluenceShapes_referenceOrder[influenceshape] = stage.InfluenceShapeMap_Staged_Order[influenceshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, influenceshape.GongMarshallUnstaging(stage))
			delete(stage.InfluenceShapes_referenceOrder, influenceshape)
			fieldInitializers, pointersInitializations := influenceshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.InfluenceShapeMap_Staged_Order[ref] = stage.InfluenceShapeMap_Staged_Order[influenceshape]
			diffs := influenceshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, influenceshape)
			delete(stage.InfluenceShapeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", influenceshape.GetName())
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
	for ref := range stage.InfluenceShapes_reference {
		if _, ok := stage.InfluenceShapes[ref]; !ok {
			influenceshapes_deletedInstances = append(influenceshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(influenceshapes_newInstances)
	lenDeletedInstances += len(influenceshapes_deletedInstances)

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
	stage.Category1s_reference = make(map[*Category1]*Category1)
	stage.Category1s_referenceOrder = make(map[*Category1]uint) // diff Unstage needs the reference order
	for instance := range stage.Category1s {
		stage.Category1s_reference[instance] = instance.GongCopy().(*Category1)
		stage.Category1s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Category1Shapes_reference = make(map[*Category1Shape]*Category1Shape)
	stage.Category1Shapes_referenceOrder = make(map[*Category1Shape]uint) // diff Unstage needs the reference order
	for instance := range stage.Category1Shapes {
		stage.Category1Shapes_reference[instance] = instance.GongCopy().(*Category1Shape)
		stage.Category1Shapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Category2s_reference = make(map[*Category2]*Category2)
	stage.Category2s_referenceOrder = make(map[*Category2]uint) // diff Unstage needs the reference order
	for instance := range stage.Category2s {
		stage.Category2s_reference[instance] = instance.GongCopy().(*Category2)
		stage.Category2s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Category2Shapes_reference = make(map[*Category2Shape]*Category2Shape)
	stage.Category2Shapes_referenceOrder = make(map[*Category2Shape]uint) // diff Unstage needs the reference order
	for instance := range stage.Category2Shapes {
		stage.Category2Shapes_reference[instance] = instance.GongCopy().(*Category2Shape)
		stage.Category2Shapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Category3s_reference = make(map[*Category3]*Category3)
	stage.Category3s_referenceOrder = make(map[*Category3]uint) // diff Unstage needs the reference order
	for instance := range stage.Category3s {
		stage.Category3s_reference[instance] = instance.GongCopy().(*Category3)
		stage.Category3s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Category3Shapes_reference = make(map[*Category3Shape]*Category3Shape)
	stage.Category3Shapes_referenceOrder = make(map[*Category3Shape]uint) // diff Unstage needs the reference order
	for instance := range stage.Category3Shapes {
		stage.Category3Shapes_reference[instance] = instance.GongCopy().(*Category3Shape)
		stage.Category3Shapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	stage.ControlPointShapes_referenceOrder = make(map[*ControlPointShape]uint) // diff Unstage needs the reference order
	for instance := range stage.ControlPointShapes {
		stage.ControlPointShapes_reference[instance] = instance.GongCopy().(*ControlPointShape)
		stage.ControlPointShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Desks_reference = make(map[*Desk]*Desk)
	stage.Desks_referenceOrder = make(map[*Desk]uint) // diff Unstage needs the reference order
	for instance := range stage.Desks {
		stage.Desks_reference[instance] = instance.GongCopy().(*Desk)
		stage.Desks_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	stage.Diagrams_referenceOrder = make(map[*Diagram]uint) // diff Unstage needs the reference order
	for instance := range stage.Diagrams {
		stage.Diagrams_reference[instance] = instance.GongCopy().(*Diagram)
		stage.Diagrams_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Influences_reference = make(map[*Influence]*Influence)
	stage.Influences_referenceOrder = make(map[*Influence]uint) // diff Unstage needs the reference order
	for instance := range stage.Influences {
		stage.Influences_reference[instance] = instance.GongCopy().(*Influence)
		stage.Influences_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.InfluenceShapes_reference = make(map[*InfluenceShape]*InfluenceShape)
	stage.InfluenceShapes_referenceOrder = make(map[*InfluenceShape]uint) // diff Unstage needs the reference order
	for instance := range stage.InfluenceShapes {
		stage.InfluenceShapes_reference[instance] = instance.GongCopy().(*InfluenceShape)
		stage.InfluenceShapes_referenceOrder[instance] = instance.GongGetOrder(stage)
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
func (category1 *Category1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Category1Map_Staged_Order[category1]; ok {
		return order
	}
	return stage.Category1s_referenceOrder[category1]
}

func (category1 *Category1) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Category1s_referenceOrder[category1]
}

func (category1shape *Category1Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Category1ShapeMap_Staged_Order[category1shape]; ok {
		return order
	}
	return stage.Category1Shapes_referenceOrder[category1shape]
}

func (category1shape *Category1Shape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Category1Shapes_referenceOrder[category1shape]
}

func (category2 *Category2) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Category2Map_Staged_Order[category2]; ok {
		return order
	}
	return stage.Category2s_referenceOrder[category2]
}

func (category2 *Category2) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Category2s_referenceOrder[category2]
}

func (category2shape *Category2Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Category2ShapeMap_Staged_Order[category2shape]; ok {
		return order
	}
	return stage.Category2Shapes_referenceOrder[category2shape]
}

func (category2shape *Category2Shape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Category2Shapes_referenceOrder[category2shape]
}

func (category3 *Category3) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Category3Map_Staged_Order[category3]; ok {
		return order
	}
	return stage.Category3s_referenceOrder[category3]
}

func (category3 *Category3) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Category3s_referenceOrder[category3]
}

func (category3shape *Category3Shape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Category3ShapeMap_Staged_Order[category3shape]; ok {
		return order
	}
	return stage.Category3Shapes_referenceOrder[category3shape]
}

func (category3shape *Category3Shape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Category3Shapes_referenceOrder[category3shape]
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ControlPointShapeMap_Staged_Order[controlpointshape]; ok {
		return order
	}
	return stage.ControlPointShapes_referenceOrder[controlpointshape]
}

func (controlpointshape *ControlPointShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ControlPointShapes_referenceOrder[controlpointshape]
}

func (desk *Desk) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DeskMap_Staged_Order[desk]; ok {
		return order
	}
	return stage.Desks_referenceOrder[desk]
}

func (desk *Desk) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Desks_referenceOrder[desk]
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramMap_Staged_Order[diagram]; ok {
		return order
	}
	return stage.Diagrams_referenceOrder[diagram]
}

func (diagram *Diagram) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Diagrams_referenceOrder[diagram]
}

func (influence *Influence) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.InfluenceMap_Staged_Order[influence]; ok {
		return order
	}
	return stage.Influences_referenceOrder[influence]
}

func (influence *Influence) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Influences_referenceOrder[influence]
}

func (influenceshape *InfluenceShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.InfluenceShapeMap_Staged_Order[influenceshape]; ok {
		return order
	}
	return stage.InfluenceShapes_referenceOrder[influenceshape]
}

func (influenceshape *InfluenceShape) GongGetReferenceOrder(stage *Stage) uint {
	return stage.InfluenceShapes_referenceOrder[influenceshape]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (category1 *Category1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category1.GongGetGongstructName(), category1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (category1 *Category1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category1.GongGetGongstructName(), category1.GongGetReferenceOrder(stage))
}

func (category1shape *Category1Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category1shape.GongGetGongstructName(), category1shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (category1shape *Category1Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category1shape.GongGetGongstructName(), category1shape.GongGetReferenceOrder(stage))
}

func (category2 *Category2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category2.GongGetGongstructName(), category2.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (category2 *Category2) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category2.GongGetGongstructName(), category2.GongGetReferenceOrder(stage))
}

func (category2shape *Category2Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category2shape.GongGetGongstructName(), category2shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (category2shape *Category2Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category2shape.GongGetGongstructName(), category2shape.GongGetReferenceOrder(stage))
}

func (category3 *Category3) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category3.GongGetGongstructName(), category3.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (category3 *Category3) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category3.GongGetGongstructName(), category3.GongGetReferenceOrder(stage))
}

func (category3shape *Category3Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category3shape.GongGetGongstructName(), category3shape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (category3shape *Category3Shape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category3shape.GongGetGongstructName(), category3shape.GongGetReferenceOrder(stage))
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (controlpointshape *ControlPointShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetReferenceOrder(stage))
}

func (desk *Desk) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", desk.GongGetGongstructName(), desk.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (desk *Desk) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", desk.GongGetGongstructName(), desk.GongGetReferenceOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagram *Diagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetReferenceOrder(stage))
}

func (influence *Influence) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influence.GongGetGongstructName(), influence.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influence *Influence) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influence.GongGetGongstructName(), influence.GongGetReferenceOrder(stage))
}

func (influenceshape *InfluenceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influenceshape.GongGetGongstructName(), influenceshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (influenceshape *InfluenceShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influenceshape.GongGetGongstructName(), influenceshape.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (category1 *Category1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category1.Name)
	return
}

func (category1shape *Category1Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category1shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category1Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category1shape.Name)
	return
}

func (category2 *Category2) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category2.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category2")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category2.Name)
	return
}

func (category2shape *Category2Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category2shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category2Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category2shape.Name)
	return
}

func (category3 *Category3) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category3.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category3")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category3.Name)
	return
}

func (category3shape *Category3Shape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category3shape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Category3Shape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", category3shape.Name)
	return
}

func (controlpointshape *ControlPointShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ControlPointShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", controlpointshape.Name)
	return
}

func (desk *Desk) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", desk.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Desk")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", desk.Name)
	return
}

func (diagram *Diagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Diagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", diagram.Name)
	return
}

func (influence *Influence) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influence.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Influence")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", influence.Name)
	return
}

func (influenceshape *InfluenceShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influenceshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "InfluenceShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", influenceshape.Name)
	return
}

// insertion point for unstaging
func (category1 *Category1) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category1.GongGetReferenceIdentifier(stage))
	return
}

func (category1shape *Category1Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category1shape.GongGetReferenceIdentifier(stage))
	return
}

func (category2 *Category2) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category2.GongGetReferenceIdentifier(stage))
	return
}

func (category2shape *Category2Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category2shape.GongGetReferenceIdentifier(stage))
	return
}

func (category3 *Category3) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category3.GongGetReferenceIdentifier(stage))
	return
}

func (category3shape *Category3Shape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", category3shape.GongGetReferenceIdentifier(stage))
	return
}

func (controlpointshape *ControlPointShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", controlpointshape.GongGetReferenceIdentifier(stage))
	return
}

func (desk *Desk) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", desk.GongGetReferenceIdentifier(stage))
	return
}

func (diagram *Diagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagram.GongGetReferenceIdentifier(stage))
	return
}

func (influence *Influence) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influence.GongGetReferenceIdentifier(stage))
	return
}

func (influenceshape *InfluenceShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", influenceshape.GongGetReferenceIdentifier(stage))
	return
}

// end of template
