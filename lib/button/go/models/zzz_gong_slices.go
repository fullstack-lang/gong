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
	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct ButtonToggle
	// insertion point per field

	// Compute reverse map for named struct Group
	// insertion point per field
	stage.Group_Buttons_reverseMap = make(map[*Button]*Group)
	for group := range stage.Groups {
		_ = group
		for _, _button := range group.Buttons {
			stage.Group_Buttons_reverseMap[_button] = group
		}
	}

	// Compute reverse map for named struct GroupToogle
	// insertion point per field
	stage.GroupToogle_ButtonToggles_reverseMap = make(map[*ButtonToggle]*GroupToogle)
	for grouptoogle := range stage.GroupToogles {
		_ = grouptoogle
		for _, _buttontoggle := range grouptoogle.ButtonToggles {
			stage.GroupToogle_ButtonToggles_reverseMap[_buttontoggle] = grouptoogle
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
	stage.Layout_GroupToogles_reverseMap = make(map[*GroupToogle]*Layout)
	for layout := range stage.Layouts {
		_ = layout
		for _, _grouptoogle := range layout.GroupToogles {
			stage.Layout_GroupToogles_reverseMap[_grouptoogle] = layout
		}
	}

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.ButtonToggles {
		res = append(res, instance)
	}

	for instance := range stage.Groups {
		res = append(res, instance)
	}

	for instance := range stage.GroupToogles {
		res = append(res, instance)
	}

	for instance := range stage.Layouts {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (button *Button) GongCopy() GongstructIF {
	newInstance := *button
	return &newInstance
}

func (buttontoggle *ButtonToggle) GongCopy() GongstructIF {
	newInstance := *buttontoggle
	return &newInstance
}

func (group *Group) GongCopy() GongstructIF {
	newInstance := *group
	return &newInstance
}

func (grouptoogle *GroupToogle) GongCopy() GongstructIF {
	newInstance := *grouptoogle
	return &newInstance
}

func (layout *Layout) GongCopy() GongstructIF {
	newInstance := *layout
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
	var buttons_newInstances []*Button
	var buttons_deletedInstances []*Button

	// parse all staged instances and check if they have a reference
	for button := range stage.Buttons {
		if ref, ok := stage.Buttons_reference[button]; !ok {
			buttons_newInstances = append(buttons_newInstances, button)
			newInstancesStmt += button.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := button.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := button.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", button.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for button := range stage.Buttons_reference {
		if _, ok := stage.Buttons[button]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, button)
			deletedInstancesStmt += button.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(buttons_newInstances)
	lenDeletedInstances += len(buttons_deletedInstances)
	var buttontoggles_newInstances []*ButtonToggle
	var buttontoggles_deletedInstances []*ButtonToggle

	// parse all staged instances and check if they have a reference
	for buttontoggle := range stage.ButtonToggles {
		if ref, ok := stage.ButtonToggles_reference[buttontoggle]; !ok {
			buttontoggles_newInstances = append(buttontoggles_newInstances, buttontoggle)
			newInstancesStmt += buttontoggle.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := buttontoggle.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := buttontoggle.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", buttontoggle.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for buttontoggle := range stage.ButtonToggles_reference {
		if _, ok := stage.ButtonToggles[buttontoggle]; !ok {
			buttontoggles_deletedInstances = append(buttontoggles_deletedInstances, buttontoggle)
			deletedInstancesStmt += buttontoggle.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(buttontoggles_newInstances)
	lenDeletedInstances += len(buttontoggles_deletedInstances)
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
	var grouptoogles_newInstances []*GroupToogle
	var grouptoogles_deletedInstances []*GroupToogle

	// parse all staged instances and check if they have a reference
	for grouptoogle := range stage.GroupToogles {
		if ref, ok := stage.GroupToogles_reference[grouptoogle]; !ok {
			grouptoogles_newInstances = append(grouptoogles_newInstances, grouptoogle)
			newInstancesStmt += grouptoogle.GongMarshallIdentifier(stage)
			fieldInitializers, pointersInitializations := grouptoogle.GongMarshallAllFields(stage)
			fieldsEditStmt += fieldInitializers
			fieldsEditStmt += pointersInitializations
		} else {
			diffs := grouptoogle.GongDiff(stage, ref)
			if len(diffs) > 0 {
				fieldsEditStmt += fmt.Sprintf("\n\t// modifications for instance \"%s\"", grouptoogle.GetName())
				for _, diff := range diffs {
					fieldsEditStmt += diff
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for grouptoogle := range stage.GroupToogles_reference {
		if _, ok := stage.GroupToogles[grouptoogle]; !ok {
			grouptoogles_deletedInstances = append(grouptoogles_deletedInstances, grouptoogle)
			deletedInstancesStmt += grouptoogle.GongMarshallUnstaging(stage)
		}
	}

	lenNewInstances += len(grouptoogles_newInstances)
	lenDeletedInstances += len(grouptoogles_deletedInstances)
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
	stage.Buttons_reference = make(map[*Button]*Button)
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance.GongCopy().(*Button)
	}

	stage.ButtonToggles_reference = make(map[*ButtonToggle]*ButtonToggle)
	for instance := range stage.ButtonToggles {
		stage.ButtonToggles_reference[instance] = instance.GongCopy().(*ButtonToggle)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
	}

	stage.GroupToogles_reference = make(map[*GroupToogle]*GroupToogle)
	for instance := range stage.GroupToogles {
		stage.GroupToogles_reference[instance] = instance.GongCopy().(*GroupToogle)
	}

	stage.Layouts_reference = make(map[*Layout]*Layout)
	for instance := range stage.Layouts {
		stage.Layouts_reference[instance] = instance.GongCopy().(*Layout)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (button *Button) GongGetOrder(stage *Stage) uint {
	return stage.ButtonMap_Staged_Order[button]
}

func (buttontoggle *ButtonToggle) GongGetOrder(stage *Stage) uint {
	return stage.ButtonToggleMap_Staged_Order[buttontoggle]
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	return stage.GroupMap_Staged_Order[group]
}

func (grouptoogle *GroupToogle) GongGetOrder(stage *Stage) uint {
	return stage.GroupToogleMap_Staged_Order[grouptoogle]
}

func (layout *Layout) GongGetOrder(stage *Stage) uint {
	return stage.LayoutMap_Staged_Order[layout]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

func (buttontoggle *ButtonToggle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", buttontoggle.GongGetGongstructName(), buttontoggle.GongGetOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

func (grouptoogle *GroupToogle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", grouptoogle.GongGetGongstructName(), grouptoogle.GongGetOrder(stage))
}

func (layout *Layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layout.GongGetGongstructName(), layout.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (button *Button) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Button")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", button.Name)
	return
}
func (buttontoggle *ButtonToggle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", buttontoggle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ButtonToggle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", buttontoggle.Name)
	return
}
func (group *Group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", group.Name)
	return
}
func (grouptoogle *GroupToogle) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", grouptoogle.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GroupToogle")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", grouptoogle.Name)
	return
}
func (layout *Layout) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layout.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Layout")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", layout.Name)
	return
}

// insertion point for unstaging
func (button *Button) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetIdentifier(stage))
	return
}
func (buttontoggle *ButtonToggle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", buttontoggle.GongGetIdentifier(stage))
	return
}
func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetIdentifier(stage))
	return
}
func (grouptoogle *GroupToogle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", grouptoogle.GongGetIdentifier(stage))
	return
}
func (layout *Layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layout.GongGetIdentifier(stage))
	return
}
