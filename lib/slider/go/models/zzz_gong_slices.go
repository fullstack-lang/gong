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
	var checkboxs_newInstances []*Checkbox
	var checkboxs_deletedInstances []*Checkbox

	// parse all staged instances and check if they have a reference
	for checkbox := range stage.Checkboxs {
		if ref, ok := stage.Checkboxs_reference[checkbox]; !ok {
			checkboxs_newInstances = append(checkboxs_newInstances, checkbox)
			newInstancesStmt += checkbox.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := checkbox.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := checkbox.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", checkbox.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for checkbox := range stage.Checkboxs_reference {
		if _, ok := stage.Checkboxs[checkbox]; !ok {
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, checkbox)
			deletedInstancesStmt += checkbox.GongMarshallUnstaging(stage)
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
			newInstancesStmt += group.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := group.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", group.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for group := range stage.Groups_reference {
		if _, ok := stage.Groups[group]; !ok {
			groups_deletedInstances = append(groups_deletedInstances, group)
			deletedInstancesStmt += group.GongMarshallUnstaging(stage)
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
			newInstancesStmt += layout.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := layout.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := layout.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", layout.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for layout := range stage.Layouts_reference {
		if _, ok := stage.Layouts[layout]; !ok {
			layouts_deletedInstances = append(layouts_deletedInstances, layout)
			deletedInstancesStmt += layout.GongMarshallUnstaging(stage)
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
			newInstancesStmt += slider.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := slider.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := slider.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", slider.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for slider := range stage.Sliders_reference {
		if _, ok := stage.Sliders[slider]; !ok {
			sliders_deletedInstances = append(sliders_deletedInstances, slider)
			deletedInstancesStmt += slider.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(sliders_newInstances)
	lenDeletedInstances += len(sliders_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		notif := newInstancesStmt+fieldsEditStmt+deletedInstancesStmt
		notif += fmt.Sprintf("\n\t// %s", time.Now().Format(time.RFC3339Nano))
		notif += fmt.Sprintf("\n\tstage.Commit()")
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
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Checkbox")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", checkbox.Name)
	return
}
func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
	return
}
func (layout *Layout) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layout.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Layout")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", layout.Name)
	return
}
func (slider *Slider) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slider.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slider")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", slider.Name)
	return
}

// insertion point for unstaging
func (checkbox *Checkbox) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
	return
}
func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	return
}
func (layout *Layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layout.GongGetIdentifier(stage))
	return
}
func (slider *Slider) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slider.GongGetIdentifier(stage))
	return
}
