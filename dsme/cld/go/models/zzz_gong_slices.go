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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var category1s_newInstances []*Category1
	var category1s_deletedInstances []*Category1

	// parse all staged instances and check if they have a reference
	for category1 := range stage.Category1s {
		if ref, ok := stage.Category1s_reference[category1]; !ok {
			category1s_newInstances = append(category1s_newInstances, category1)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Category1 "+category1.Name,
				)
			}
		} else {
			diffs := category1.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Category1 \""+category1.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for category1 := range stage.Category1s_reference {
		if _, ok := stage.Category1s[category1]; !ok {
			category1s_deletedInstances = append(category1s_deletedInstances, category1)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Category1 "+category1.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Category1Shape "+category1shape.Name,
				)
			}
		} else {
			diffs := category1shape.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Category1Shape \""+category1shape.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for category1shape := range stage.Category1Shapes_reference {
		if _, ok := stage.Category1Shapes[category1shape]; !ok {
			category1shapes_deletedInstances = append(category1shapes_deletedInstances, category1shape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Category1Shape "+category1shape.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Category2 "+category2.Name,
				)
			}
		} else {
			diffs := category2.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Category2 \""+category2.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for category2 := range stage.Category2s_reference {
		if _, ok := stage.Category2s[category2]; !ok {
			category2s_deletedInstances = append(category2s_deletedInstances, category2)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Category2 "+category2.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Category2Shape "+category2shape.Name,
				)
			}
		} else {
			diffs := category2shape.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Category2Shape \""+category2shape.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for category2shape := range stage.Category2Shapes_reference {
		if _, ok := stage.Category2Shapes[category2shape]; !ok {
			category2shapes_deletedInstances = append(category2shapes_deletedInstances, category2shape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Category2Shape "+category2shape.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Category3 "+category3.Name,
				)
			}
		} else {
			diffs := category3.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Category3 \""+category3.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for category3 := range stage.Category3s_reference {
		if _, ok := stage.Category3s[category3]; !ok {
			category3s_deletedInstances = append(category3s_deletedInstances, category3)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Category3 "+category3.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Category3Shape "+category3shape.Name,
				)
			}
		} else {
			diffs := category3shape.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Category3Shape \""+category3shape.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for category3shape := range stage.Category3Shapes_reference {
		if _, ok := stage.Category3Shapes[category3shape]; !ok {
			category3shapes_deletedInstances = append(category3shapes_deletedInstances, category3shape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Category3Shape "+category3shape.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of ControlPointShape "+controlpointshape.Name,
				)
			}
		} else {
			diffs := controlpointshape.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of ControlPointShape \""+controlpointshape.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for controlpointshape := range stage.ControlPointShapes_reference {
		if _, ok := stage.ControlPointShapes[controlpointshape]; !ok {
			controlpointshapes_deletedInstances = append(controlpointshapes_deletedInstances, controlpointshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of ControlPointShape "+controlpointshape.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Desk "+desk.Name,
				)
			}
		} else {
			diffs := desk.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Desk \""+desk.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for desk := range stage.Desks_reference {
		if _, ok := stage.Desks[desk]; !ok {
			desks_deletedInstances = append(desks_deletedInstances, desk)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Desk "+desk.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Diagram "+diagram.Name,
				)
			}
		} else {
			diffs := diagram.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Diagram \""+diagram.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for diagram := range stage.Diagrams_reference {
		if _, ok := stage.Diagrams[diagram]; !ok {
			diagrams_deletedInstances = append(diagrams_deletedInstances, diagram)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Diagram "+diagram.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Influence "+influence.Name,
				)
			}
		} else {
			diffs := influence.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Influence \""+influence.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for influence := range stage.Influences_reference {
		if _, ok := stage.Influences[influence]; !ok {
			influences_deletedInstances = append(influences_deletedInstances, influence)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Influence "+influence.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of InfluenceShape "+influenceshape.Name,
				)
			}
		} else {
			diffs := influenceshape.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of InfluenceShape \""+influenceshape.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for influenceshape := range stage.InfluenceShapes_reference {
		if _, ok := stage.InfluenceShapes[influenceshape]; !ok {
			influenceshapes_deletedInstances = append(influenceshapes_deletedInstances, influenceshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of InfluenceShape "+influenceshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(influenceshapes_newInstances)
	lenDeletedInstances += len(influenceshapes_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Category1s_reference = make(map[*Category1]*Category1)
	for instance := range stage.Category1s {
		stage.Category1s_reference[instance] = instance.GongCopy().(*Category1)
	}

	stage.Category1Shapes_reference = make(map[*Category1Shape]*Category1Shape)
	for instance := range stage.Category1Shapes {
		stage.Category1Shapes_reference[instance] = instance.GongCopy().(*Category1Shape)
	}

	stage.Category2s_reference = make(map[*Category2]*Category2)
	for instance := range stage.Category2s {
		stage.Category2s_reference[instance] = instance.GongCopy().(*Category2)
	}

	stage.Category2Shapes_reference = make(map[*Category2Shape]*Category2Shape)
	for instance := range stage.Category2Shapes {
		stage.Category2Shapes_reference[instance] = instance.GongCopy().(*Category2Shape)
	}

	stage.Category3s_reference = make(map[*Category3]*Category3)
	for instance := range stage.Category3s {
		stage.Category3s_reference[instance] = instance.GongCopy().(*Category3)
	}

	stage.Category3Shapes_reference = make(map[*Category3Shape]*Category3Shape)
	for instance := range stage.Category3Shapes {
		stage.Category3Shapes_reference[instance] = instance.GongCopy().(*Category3Shape)
	}

	stage.ControlPointShapes_reference = make(map[*ControlPointShape]*ControlPointShape)
	for instance := range stage.ControlPointShapes {
		stage.ControlPointShapes_reference[instance] = instance.GongCopy().(*ControlPointShape)
	}

	stage.Desks_reference = make(map[*Desk]*Desk)
	for instance := range stage.Desks {
		stage.Desks_reference[instance] = instance.GongCopy().(*Desk)
	}

	stage.Diagrams_reference = make(map[*Diagram]*Diagram)
	for instance := range stage.Diagrams {
		stage.Diagrams_reference[instance] = instance.GongCopy().(*Diagram)
	}

	stage.Influences_reference = make(map[*Influence]*Influence)
	for instance := range stage.Influences {
		stage.Influences_reference[instance] = instance.GongCopy().(*Influence)
	}

	stage.InfluenceShapes_reference = make(map[*InfluenceShape]*InfluenceShape)
	for instance := range stage.InfluenceShapes {
		stage.InfluenceShapes_reference[instance] = instance.GongCopy().(*InfluenceShape)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (category1 *Category1) GongGetOrder(stage *Stage) uint {
	return stage.Category1Map_Staged_Order[category1]
}

func (category1shape *Category1Shape) GongGetOrder(stage *Stage) uint {
	return stage.Category1ShapeMap_Staged_Order[category1shape]
}

func (category2 *Category2) GongGetOrder(stage *Stage) uint {
	return stage.Category2Map_Staged_Order[category2]
}

func (category2shape *Category2Shape) GongGetOrder(stage *Stage) uint {
	return stage.Category2ShapeMap_Staged_Order[category2shape]
}

func (category3 *Category3) GongGetOrder(stage *Stage) uint {
	return stage.Category3Map_Staged_Order[category3]
}

func (category3shape *Category3Shape) GongGetOrder(stage *Stage) uint {
	return stage.Category3ShapeMap_Staged_Order[category3shape]
}

func (controlpointshape *ControlPointShape) GongGetOrder(stage *Stage) uint {
	return stage.ControlPointShapeMap_Staged_Order[controlpointshape]
}

func (desk *Desk) GongGetOrder(stage *Stage) uint {
	return stage.DeskMap_Staged_Order[desk]
}

func (diagram *Diagram) GongGetOrder(stage *Stage) uint {
	return stage.DiagramMap_Staged_Order[diagram]
}

func (influence *Influence) GongGetOrder(stage *Stage) uint {
	return stage.InfluenceMap_Staged_Order[influence]
}

func (influenceshape *InfluenceShape) GongGetOrder(stage *Stage) uint {
	return stage.InfluenceShapeMap_Staged_Order[influenceshape]
}


// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (category1 *Category1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category1.GongGetGongstructName(), category1.GongGetOrder(stage))
}

func (category1shape *Category1Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category1shape.GongGetGongstructName(), category1shape.GongGetOrder(stage))
}

func (category2 *Category2) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category2.GongGetGongstructName(), category2.GongGetOrder(stage))
}

func (category2shape *Category2Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category2shape.GongGetGongstructName(), category2shape.GongGetOrder(stage))
}

func (category3 *Category3) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category3.GongGetGongstructName(), category3.GongGetOrder(stage))
}

func (category3shape *Category3Shape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", category3shape.GongGetGongstructName(), category3shape.GongGetOrder(stage))
}

func (controlpointshape *ControlPointShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", controlpointshape.GongGetGongstructName(), controlpointshape.GongGetOrder(stage))
}

func (desk *Desk) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", desk.GongGetGongstructName(), desk.GongGetOrder(stage))
}

func (diagram *Diagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagram.GongGetGongstructName(), diagram.GongGetOrder(stage))
}

func (influence *Influence) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influence.GongGetGongstructName(), influence.GongGetOrder(stage))
}

func (influenceshape *InfluenceShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", influenceshape.GongGetGongstructName(), influenceshape.GongGetOrder(stage))
}

