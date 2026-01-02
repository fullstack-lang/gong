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
	// Compute reverse map for named struct Checkbox
	// insertion point per field

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_Sliders_reverseMap = make(map[*Slider]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _slider := range group.Sliders {
			stage.Group_Sliders_reverseMap[_slider] = group
		}
	}
	stage.Group_Checkboxes_reverseMap = make(map[*Checkbox]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _checkbox := range group.Checkboxes {
			stage.Group_Checkboxes_reverseMap[_checkbox] = group
		}
	}

	// Compute reverse map for named struct Layout
	// insertion point per field
	stage.Layout_Groups_reverseMap = make(map[*Group]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _group := range layout.Groups {
			stage.Layout_Groups_reverseMap[_group] = layout
		}
	}

	// Compute reverse map for named struct Slider
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Checkboxs {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.Layouts {
		res = append(res, instance)
	}

	for instance := range stage.Sliders {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (checkbox *Checkbox) GongCopy() GongstructIF {
	newInstance := *checkbox
	return &newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := *group
	return &newInstance
}

func (layout *Layout) GongCopy() GongstructIF {
	newInstance := *layout
	return &newInstance
}

func (slider *Slider) GongCopy() GongstructIF {
	newInstance := *slider
	return &newInstance
}

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var pointersInitializesStatements string

	// insertion point per named struct
	var checkboxs_newInstances []*Checkbox
	var checkboxs_deletedInstances []*Checkbox

	// parse all staged instances and check if they have a reference
	for checkbox := range stage.Checkboxs {
		if ref, ok := stage.Checkboxs_reference[checkbox]; !ok {
			checkboxs_newInstances = append(checkboxs_newInstances, checkbox)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Checkbox "+checkbox.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					checkbox.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := checkbox.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := checkbox.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Checkbox \""+checkbox.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							checkbox.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for checkbox := range stage.Checkboxs_reference {
		if _, ok := stage.Checkboxs[checkbox]; !ok {
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, checkbox)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Checkbox "+checkbox.Name,
				)
			}
		}
	}

	lenNewInstances += len(checkboxs_newInstances)
	lenDeletedInstances += len(checkboxs_deletedInstances)
	var groups_newInstances []*Group
	var groups_deletedInstances []*Group

	// parse all staged instances and check if they have a reference
	for group := range stage.Groups {
		if ref, ok := stage.Groups_reference[group]; !ok {
			groups_newInstances = append(groups_newInstances, group)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Group "+group.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					group.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := group.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Group \""+group.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							group.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for group := range stage.Groups_reference {
		if _, ok := stage.Groups[group]; !ok {
			groups_deletedInstances = append(groups_deletedInstances, group)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Group "+group.Name,
				)
			}
		}
	}

	lenNewInstances += len(groups_newInstances)
	lenDeletedInstances += len(groups_deletedInstances)
	var layouts_newInstances []*Layout
	var layouts_deletedInstances []*Layout

	// parse all staged instances and check if they have a reference
	for layout := range stage.Layouts {
		if ref, ok := stage.Layouts_reference[layout]; !ok {
			layouts_newInstances = append(layouts_newInstances, layout)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Layout "+layout.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					layout.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := layout.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := layout.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Layout \""+layout.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							layout.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for layout := range stage.Layouts_reference {
		if _, ok := stage.Layouts[layout]; !ok {
			layouts_deletedInstances = append(layouts_deletedInstances, layout)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Layout "+layout.Name,
				)
			}
		}
	}

	lenNewInstances += len(layouts_newInstances)
	lenDeletedInstances += len(layouts_deletedInstances)
	var sliders_newInstances []*Slider
	var sliders_deletedInstances []*Slider

	// parse all staged instances and check if they have a reference
	for slider := range stage.Sliders {
		if ref, ok := stage.Sliders_reference[slider]; !ok {
			sliders_newInstances = append(sliders_newInstances, slider)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Slider "+slider.Name,
				)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					slider.GongMarshallIdentifier(stage),
				)
				basicFieldInitializers, pointersInitializations := slider.GongMarshallAllFields(stage)
				stage.GetProbeIF().AddNotification(
					time.Now(),
					basicFieldInitializers,
				)
				pointersInitializesStatements += pointersInitializations
			}
		} else {
			diffs := slider.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Slider \""+slider.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
					for _, diff := range diffs {
						stage.GetProbeIF().AddNotification(
							time.Now(),
							slider.GongMarshallField(stage, diff),
						)
					}
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for slider := range stage.Sliders_reference {
		if _, ok := stage.Sliders[slider]; !ok {
			sliders_deletedInstances = append(sliders_deletedInstances, slider)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Slider "+slider.Name,
				)
			}
		}
	}

	lenNewInstances += len(sliders_newInstances)
	lenDeletedInstances += len(sliders_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}

	if pointersInitializesStatements != "" {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().AddNotification(
				time.Now(),
				pointersInitializesStatements,
			)
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.Checkboxs_reference = make(map[*Checkbox]*Checkbox)
	for instance := range stage.Checkboxs {
		stage.Checkboxs_reference[instance] = instance.GongCopy().(*Checkbox)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
	}

	stage.Layouts_reference = make(map[*Layout]*Layout)
	for instance := range stage.Layouts {
		stage.Layouts_reference[instance] = instance.GongCopy().(*Layout)
	}

	stage.Sliders_reference = make(map[*Slider]*Slider)
	for instance := range stage.Sliders {
		stage.Sliders_reference[instance] = instance.GongCopy().(*Slider)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (checkbox *Checkbox) GongGetOrder(stage *Stage) uint {
	return stage.CheckboxMap_Staged_Order[checkbox]
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return stage.GroupMap_Staged_Order[group]
}

func (layout *Layout) GongGetOrder(stage *Stage) uint {
	return stage.LayoutMap_Staged_Order[layout]
}

func (slider *Slider) GongGetOrder(stage *Stage) uint {
	return stage.SliderMap_Staged_Order[slider]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (checkbox *Checkbox) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

func (layout *Layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layout.GongGetGongstructName(), layout.GongGetOrder(stage))
}

func (slider *Slider) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (checkbox *Checkbox) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Checkbox")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", checkbox.Name)
	return
}
func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
	return
}
func (layout *Layout) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layout.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Layout")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", layout.Name)
	return
}
func (slider *Slider) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = IdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slider.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slider")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", slider.Name)
	return
}
