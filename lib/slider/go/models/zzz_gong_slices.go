// generated code - do not edit
package models

import (
	"fmt"
	"sort"
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
	var checkboxs_newInstances []*Checkbox
	var checkboxs_deletedInstances []*Checkbox

	// parse all staged instances and check if they have a reference
	for checkbox := range stage.Checkboxs {
		if ref, ok := stage.Checkboxs_reference[checkbox]; !ok {
			checkboxs_newInstances = append(checkboxs_newInstances, checkbox)
			newInstancesSlice = append(newInstancesSlice, checkbox.GongMarshallIdentifier(stage))
			if stage.Checkboxs_referenceOrder == nil {
				stage.Checkboxs_referenceOrder = make(map[*Checkbox]uint)
			}
			stage.Checkboxs_referenceOrder[checkbox] = stage.CheckboxMap_Staged_Order[checkbox]
			newInstancesReverseSlice = append(newInstancesReverseSlice, checkbox.GongMarshallUnstaging(stage))
			delete(stage.Checkboxs_referenceOrder, checkbox)
			fieldInitializers, pointersInitializations := checkbox.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CheckboxMap_Staged_Order[ref] = stage.CheckboxMap_Staged_Order[checkbox]
			diffs := checkbox.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, checkbox)
			delete(stage.CheckboxMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", checkbox.GetName())
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
	for ref := range stage.Checkboxs_reference {
		if _, ok := stage.Checkboxs[ref]; !ok {
			checkboxs_deletedInstances = append(checkboxs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
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
	var sliders_newInstances []*Slider
	var sliders_deletedInstances []*Slider

	// parse all staged instances and check if they have a reference
	for slider := range stage.Sliders {
		if ref, ok := stage.Sliders_reference[slider]; !ok {
			sliders_newInstances = append(sliders_newInstances, slider)
			newInstancesSlice = append(newInstancesSlice, slider.GongMarshallIdentifier(stage))
			if stage.Sliders_referenceOrder == nil {
				stage.Sliders_referenceOrder = make(map[*Slider]uint)
			}
			stage.Sliders_referenceOrder[slider] = stage.SliderMap_Staged_Order[slider]
			newInstancesReverseSlice = append(newInstancesReverseSlice, slider.GongMarshallUnstaging(stage))
			delete(stage.Sliders_referenceOrder, slider)
			fieldInitializers, pointersInitializations := slider.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SliderMap_Staged_Order[ref] = stage.SliderMap_Staged_Order[slider]
			diffs := slider.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, slider)
			delete(stage.SliderMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", slider.GetName())
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
	for ref := range stage.Sliders_reference {
		if _, ok := stage.Sliders[ref]; !ok {
			sliders_deletedInstances = append(sliders_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(sliders_newInstances)
	lenDeletedInstances += len(sliders_deletedInstances)

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

		if stage.GetProbeIF() != nil {
			var mergedCommits string
			for _, commit := range stage.forwardCommits {
				mergedCommits += commit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Forward commits:\n"+
					mergedCommits,
			)

			var reverseMergedCommits string
			for _, reverserCommit := range stage.backwardCommits {
				reverseMergedCommits += reverserCommit
			}
			stage.GetProbeIF().AddNotification(
				time.Now(),
				"	// Backward commits:\n"+
					reverseMergedCommits,
			)

			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {

	// insertion point per named struct
	stage.Checkboxs_reference = make(map[*Checkbox]*Checkbox)
	stage.Checkboxs_referenceOrder = make(map[*Checkbox]uint) // diff Unstage needs the reference order
	for instance := range stage.Checkboxs {
		stage.Checkboxs_reference[instance] = instance.GongCopy().(*Checkbox)
		stage.Checkboxs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint) // diff Unstage needs the reference order
	for instance := range stage.Groups {
		stage.Groups_reference[instance] = instance.GongCopy().(*Group)
		stage.Groups_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Layouts_reference = make(map[*Layout]*Layout)
	stage.Layouts_referenceOrder = make(map[*Layout]uint) // diff Unstage needs the reference order
	for instance := range stage.Layouts {
		stage.Layouts_reference[instance] = instance.GongCopy().(*Layout)
		stage.Layouts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Sliders_reference = make(map[*Slider]*Slider)
	stage.Sliders_referenceOrder = make(map[*Slider]uint) // diff Unstage needs the reference order
	for instance := range stage.Sliders {
		stage.Sliders_reference[instance] = instance.GongCopy().(*Slider)
		stage.Sliders_referenceOrder[instance] = instance.GongGetOrder(stage)
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
func (checkbox *Checkbox) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CheckboxMap_Staged_Order[checkbox]; ok {
		return order
	}
	return stage.Checkboxs_referenceOrder[checkbox]
}

func (checkbox *Checkbox) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Checkboxs_referenceOrder[checkbox]
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

func (layout *Layout) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LayoutMap_Staged_Order[layout]; ok {
		return order
	}
	return stage.Layouts_referenceOrder[layout]
}

func (layout *Layout) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Layouts_referenceOrder[layout]
}

func (slider *Slider) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SliderMap_Staged_Order[slider]; ok {
		return order
	}
	return stage.Sliders_referenceOrder[slider]
}

func (slider *Slider) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Sliders_referenceOrder[slider]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (checkbox *Checkbox) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (checkbox *Checkbox) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", checkbox.GongGetGongstructName(), checkbox.GongGetReferenceOrder(stage))
}

func (group *Group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group *Group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group.GongGetGongstructName(), group.GongGetReferenceOrder(stage))
}

func (layout *Layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layout.GongGetGongstructName(), layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (layout *Layout) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", layout.GongGetGongstructName(), layout.GongGetReferenceOrder(stage))
}

func (slider *Slider) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slider *Slider) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetReferenceOrder(stage))
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
	decl = strings.ReplaceAll(decl, "{{Identifier}}", checkbox.GongGetReferenceIdentifier(stage))
	return
}
func (group *Group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group.GongGetReferenceIdentifier(stage))
	return
}
func (layout *Layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", layout.GongGetReferenceIdentifier(stage))
	return
}
func (slider *Slider) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slider.GongGetReferenceIdentifier(stage))
	return
}
