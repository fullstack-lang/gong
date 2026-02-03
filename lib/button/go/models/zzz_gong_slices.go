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

	// end of insertion point per named struct
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
	var buttons_newInstances []*Button
	var buttons_deletedInstances []*Button

	// parse all staged instances and check if they have a reference
	for button := range stage.Buttons {
		if ref, ok := stage.Buttons_reference[button]; !ok {
			buttons_newInstances = append(buttons_newInstances, button)
			newInstancesSlice = append(newInstancesSlice, button.GongMarshallIdentifier(stage))
			if stage.Buttons_referenceOrder == nil {
				stage.Buttons_referenceOrder = make(map[*Button]uint)
			}
			stage.Buttons_referenceOrder[button] = stage.ButtonMap_Staged_Order[button]
			newInstancesReverseSlice = append(newInstancesReverseSlice, button.GongMarshallUnstaging(stage))
			delete(stage.Buttons_referenceOrder, button)
			fieldInitializers, pointersInitializations := button.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ButtonMap_Staged_Order[ref] = stage.ButtonMap_Staged_Order[button]
			diffs := button.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, button)
			delete(stage.ButtonMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", button.GetName())
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
	for ref := range stage.Buttons_reference {
		if _, ok := stage.Buttons[ref]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, buttontoggle.GongMarshallIdentifier(stage))
			if stage.ButtonToggles_referenceOrder == nil {
				stage.ButtonToggles_referenceOrder = make(map[*ButtonToggle]uint)
			}
			stage.ButtonToggles_referenceOrder[buttontoggle] = stage.ButtonToggleMap_Staged_Order[buttontoggle]
			newInstancesReverseSlice = append(newInstancesReverseSlice, buttontoggle.GongMarshallUnstaging(stage))
			delete(stage.ButtonToggles_referenceOrder, buttontoggle)
			fieldInitializers, pointersInitializations := buttontoggle.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ButtonToggleMap_Staged_Order[ref] = stage.ButtonToggleMap_Staged_Order[buttontoggle]
			diffs := buttontoggle.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, buttontoggle)
			delete(stage.ButtonToggleMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", buttontoggle.GetName())
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
	for ref := range stage.ButtonToggles_reference {
		if _, ok := stage.ButtonToggles[ref]; !ok {
			buttontoggles_deletedInstances = append(buttontoggles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, group.GongMarshallIdentifier(stage))
			if stage.Groups_referenceOrder == nil {
				stage.Groups_referenceOrder = make(map[*Group]uint)
			}
			stage.Groups_referenceOrder[group] = stage.GroupMap_Staged_Order[group]
			newInstancesReverseSlice = append(newInstancesReverseSlice, group.GongMarshallUnstaging(stage))
			delete(stage.Groups_referenceOrder, group)
			fieldInitializers, pointersInitializations := group.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GroupMap_Staged_Order[ref] = stage.GroupMap_Staged_Order[group]
			diffs := group.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, group)
			delete(stage.GroupMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", group.GetName())
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
	for ref := range stage.Groups_reference {
		if _, ok := stage.Groups[ref]; !ok {
			groups_deletedInstances = append(groups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, grouptoogle.GongMarshallIdentifier(stage))
			if stage.GroupToogles_referenceOrder == nil {
				stage.GroupToogles_referenceOrder = make(map[*GroupToogle]uint)
			}
			stage.GroupToogles_referenceOrder[grouptoogle] = stage.GroupToogleMap_Staged_Order[grouptoogle]
			newInstancesReverseSlice = append(newInstancesReverseSlice, grouptoogle.GongMarshallUnstaging(stage))
			delete(stage.GroupToogles_referenceOrder, grouptoogle)
			fieldInitializers, pointersInitializations := grouptoogle.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.GroupToogleMap_Staged_Order[ref] = stage.GroupToogleMap_Staged_Order[grouptoogle]
			diffs := grouptoogle.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, grouptoogle)
			delete(stage.GroupToogleMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", grouptoogle.GetName())
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
	for ref := range stage.GroupToogles_reference {
		if _, ok := stage.GroupToogles[ref]; !ok {
			grouptoogles_deletedInstances = append(grouptoogles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
			newInstancesSlice = append(newInstancesSlice, layout.GongMarshallIdentifier(stage))
			if stage.Layouts_referenceOrder == nil {
				stage.Layouts_referenceOrder = make(map[*Layout]uint)
			}
			stage.Layouts_referenceOrder[layout] = stage.LayoutMap_Staged_Order[layout]
			newInstancesReverseSlice = append(newInstancesReverseSlice, layout.GongMarshallUnstaging(stage))
			delete(stage.Layouts_referenceOrder, layout)
			fieldInitializers, pointersInitializations := layout.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LayoutMap_Staged_Order[ref] = stage.LayoutMap_Staged_Order[layout]
			diffs := layout.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, layout)
			delete(stage.LayoutMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", layout.GetName())
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
	for ref := range stage.Layouts_reference {
		if _, ok := stage.Layouts[ref]; !ok {
			layouts_deletedInstances = append(layouts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(layouts_newInstances)
	lenDeletedInstances += len(layouts_deletedInstances)

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
	stage.Buttons_reference = make(map[*Button]*Button)
	stage.Buttons_referenceOrder = make(map[*Button]uint) // diff Unstage needs the reference order
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance.GongCopy().(*Button)
		stage.Buttons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ButtonToggles_reference = make(map[*ButtonToggle]*ButtonToggle)
	stage.ButtonToggles_referenceOrder = make(map[*ButtonToggle]uint) // diff Unstage needs the reference order
	for instance := range stage.ButtonToggles {
		stage.ButtonToggles_reference[instance] = instance.GongCopy().(*ButtonToggle)
		stage.ButtonToggles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint) // diff Unstage needs the reference order
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
		stage.Groups_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.GroupToogles_reference = make(map[*GroupToogle]*GroupToogle)
	stage.GroupToogles_referenceOrder = make(map[*GroupToogle]uint) // diff Unstage needs the reference order
	for instance := range stage.GroupToogles {
		stage.GroupToogles_reference[instance] = instance.GongCopy().(*GroupToogle)
		stage.GroupToogles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Layouts_reference = make(map[*Layout]*Layout)
	stage.Layouts_referenceOrder = make(map[*Layout]uint) // diff Unstage needs the reference order
	for instance := range stage.Layouts {
		stage.Layouts_reference[instance] = instance.GongCopy().(*Layout)
		stage.Layouts_referenceOrder[instance] = instance.GongGetOrder(stage)
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
func (button *Button) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ButtonMap_Staged_Order[button]; ok {
		return order
	}
	return stage.Buttons_referenceOrder[button]
}

func (button *Button) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Buttons_referenceOrder[button]
}

func (buttontoggle *ButtonToggle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ButtonToggleMap_Staged_Order[buttontoggle]; ok {
		return order
	}
	return stage.ButtonToggles_referenceOrder[buttontoggle]
}

func (buttontoggle *ButtonToggle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ButtonToggles_referenceOrder[buttontoggle]
}

func (group *Group) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GroupMap_Staged_Order[group]; ok {
		return order
	}
	return stage.Groups_referenceOrder[group]
}

func (group *Group) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Groups_referenceOrder[group]
}

func (grouptoogle *GroupToogle) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GroupToogleMap_Staged_Order[grouptoogle]; ok {
		return order
	}
	return stage.GroupToogles_referenceOrder[grouptoogle]
}

func (grouptoogle *GroupToogle) GongGetReferenceOrder(stage *Stage) uint {
	return stage.GroupToogles_referenceOrder[grouptoogle]
}

func (layout *Layout) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LayoutMap_Staged_Order[layout]; ok {
		return order
	}
	return stage.Layouts_referenceOrder[layout]
}

func (layout *Layout) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Layouts_referenceOrder[layout]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetReferenceOrder(stage))
}

func (buttontoggle *ButtonToggle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", buttontoggle.GongGetGongstructName(), buttontoggle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (buttontoggle *ButtonToggle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", buttontoggle.GongGetGongstructName(), buttontoggle.GongGetReferenceOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetReferenceOrder(stage))
}

func (grouptoogle *GroupToogle) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", grouptoogle.GongGetGongstructName(), grouptoogle.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grouptoogle *GroupToogle) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", grouptoogle.GongGetGongstructName(), grouptoogle.GongGetReferenceOrder(stage))
}

func (layout *Layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layout.GongGetGongstructName(), layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layout *Layout) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layout.GongGetGongstructName(), layout.GongGetReferenceOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetReferenceIdentifier(stage))
	return
}

func (buttontoggle *ButtonToggle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", buttontoggle.GongGetReferenceIdentifier(stage))
	return
}

func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetReferenceIdentifier(stage))
	return
}

func (grouptoogle *GroupToogle) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", grouptoogle.GongGetReferenceIdentifier(stage))
	return
}

func (layout *Layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layout.GongGetReferenceIdentifier(stage))
	return
}

// end of template
