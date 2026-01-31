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
	// Compute reverse map for named struct AsSplit
	// insertion point per field
	stage.AsSplit_AsSplitAreas_reverseMap = make(map[*AsSplitArea]*AsSplit)
	for assplit := range stage.AsSplits {
		_ = assplit
		for _, _assplitarea := range assplit.AsSplitAreas {
			stage.AsSplit_AsSplitAreas_reverseMap[_assplitarea] = assplit
		}
	}

	// Compute reverse map for named struct AsSplitArea
	// insertion point per field

	// Compute reverse map for named struct Button
	// insertion point per field

	// Compute reverse map for named struct Cursor
	// insertion point per field

	// Compute reverse map for named struct FavIcon
	// insertion point per field

	// Compute reverse map for named struct Form
	// insertion point per field

	// Compute reverse map for named struct Load
	// insertion point per field

	// Compute reverse map for named struct LogoOnTheLeft
	// insertion point per field

	// Compute reverse map for named struct LogoOnTheRight
	// insertion point per field

	// Compute reverse map for named struct Markdown
	// insertion point per field

	// Compute reverse map for named struct Slider
	// insertion point per field

	// Compute reverse map for named struct Split
	// insertion point per field

	// Compute reverse map for named struct Svg
	// insertion point per field

	// Compute reverse map for named struct Table
	// insertion point per field

	// Compute reverse map for named struct Title
	// insertion point per field

	// Compute reverse map for named struct Tone
	// insertion point per field

	// Compute reverse map for named struct Tree
	// insertion point per field

	// Compute reverse map for named struct View
	// insertion point per field
	stage.View_RootAsSplitAreas_reverseMap = make(map[*AsSplitArea]*View)
	for view := range stage.Views {
		_ = view
		for _, _assplitarea := range view.RootAsSplitAreas {
			stage.View_RootAsSplitAreas_reverseMap[_assplitarea] = view
		}
	}

	// Compute reverse map for named struct Xlsx
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.AsSplits {
		res = append(res, instance)
	}

	for instance := range stage.AsSplitAreas {
		res = append(res, instance)
	}

	for instance := range stage.Buttons {
		res = append(res, instance)
	}

	for instance := range stage.Cursors {
		res = append(res, instance)
	}

	for instance := range stage.FavIcons {
		res = append(res, instance)
	}

	for instance := range stage.Forms {
		res = append(res, instance)
	}

	for instance := range stage.Loads {
		res = append(res, instance)
	}

	for instance := range stage.LogoOnTheLefts {
		res = append(res, instance)
	}

	for instance := range stage.LogoOnTheRights {
		res = append(res, instance)
	}

	for instance := range stage.Markdowns {
		res = append(res, instance)
	}

	for instance := range stage.Sliders {
		res = append(res, instance)
	}

	for instance := range stage.Splits {
		res = append(res, instance)
	}

	for instance := range stage.Svgs {
		res = append(res, instance)
	}

	for instance := range stage.Tables {
		res = append(res, instance)
	}

	for instance := range stage.Titles {
		res = append(res, instance)
	}

	for instance := range stage.Tones {
		res = append(res, instance)
	}

	for instance := range stage.Trees {
		res = append(res, instance)
	}

	for instance := range stage.Views {
		res = append(res, instance)
	}

	for instance := range stage.Xlsxs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (assplit *AsSplit) GongCopy() GongstructIF {
	newInstance := *assplit
	return &newInstance
}

func (assplitarea *AsSplitArea) GongCopy() GongstructIF {
	newInstance := *assplitarea
	return &newInstance
}

func (button *Button) GongCopy() GongstructIF {
	newInstance := *button
	return &newInstance
}

func (cursor *Cursor) GongCopy() GongstructIF {
	newInstance := *cursor
	return &newInstance
}

func (favicon *FavIcon) GongCopy() GongstructIF {
	newInstance := *favicon
	return &newInstance
}

func (form *Form) GongCopy() GongstructIF {
	newInstance := *form
	return &newInstance
}

func (load *Load) GongCopy() GongstructIF {
	newInstance := *load
	return &newInstance
}

func (logoontheleft *LogoOnTheLeft) GongCopy() GongstructIF {
	newInstance := *logoontheleft
	return &newInstance
}

func (logoontheright *LogoOnTheRight) GongCopy() GongstructIF {
	newInstance := *logoontheright
	return &newInstance
}

func (markdown *Markdown) GongCopy() GongstructIF {
	newInstance := *markdown
	return &newInstance
}

func (slider *Slider) GongCopy() GongstructIF {
	newInstance := *slider
	return &newInstance
}

func (split *Split) GongCopy() GongstructIF {
	newInstance := *split
	return &newInstance
}

func (svg *Svg) GongCopy() GongstructIF {
	newInstance := *svg
	return &newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := *table
	return &newInstance
}

func (title *Title) GongCopy() GongstructIF {
	newInstance := *title
	return &newInstance
}

func (tone *Tone) GongCopy() GongstructIF {
	newInstance := *tone
	return &newInstance
}

func (tree *Tree) GongCopy() GongstructIF {
	newInstance := *tree
	return &newInstance
}

func (view *View) GongCopy() GongstructIF {
	newInstance := *view
	return &newInstance
}

func (xlsx *Xlsx) GongCopy() GongstructIF {
	newInstance := *xlsx
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
	var assplits_newInstances []*AsSplit
	var assplits_deletedInstances []*AsSplit

	// parse all staged instances and check if they have a reference
	for assplit := range stage.AsSplits {
		if ref, ok := stage.AsSplits_reference[assplit]; !ok {
			assplits_newInstances = append(assplits_newInstances, assplit)
			newInstancesSlice = append(newInstancesSlice, assplit.GongMarshallIdentifier(stage))
			if stage.AsSplits_referenceOrder == nil {
				stage.AsSplits_referenceOrder = make(map[*AsSplit]uint)
			}
			stage.AsSplits_referenceOrder[assplit] = stage.AsSplitMap_Staged_Order[assplit]
			newInstancesReverseSlice = append(newInstancesReverseSlice, assplit.GongMarshallUnstaging(stage))
			delete(stage.AsSplits_referenceOrder, assplit)
			fieldInitializers, pointersInitializations := assplit.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AsSplitMap_Staged_Order[ref] = stage.AsSplitMap_Staged_Order[assplit]
			diffs := assplit.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, assplit)
			delete(stage.AsSplitMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", assplit.GetName())
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
	for ref := range stage.AsSplits_reference {
		if _, ok := stage.AsSplits[ref]; !ok {
			assplits_deletedInstances = append(assplits_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(assplits_newInstances)
	lenDeletedInstances += len(assplits_deletedInstances)
	var assplitareas_newInstances []*AsSplitArea
	var assplitareas_deletedInstances []*AsSplitArea

	// parse all staged instances and check if they have a reference
	for assplitarea := range stage.AsSplitAreas {
		if ref, ok := stage.AsSplitAreas_reference[assplitarea]; !ok {
			assplitareas_newInstances = append(assplitareas_newInstances, assplitarea)
			newInstancesSlice = append(newInstancesSlice, assplitarea.GongMarshallIdentifier(stage))
			if stage.AsSplitAreas_referenceOrder == nil {
				stage.AsSplitAreas_referenceOrder = make(map[*AsSplitArea]uint)
			}
			stage.AsSplitAreas_referenceOrder[assplitarea] = stage.AsSplitAreaMap_Staged_Order[assplitarea]
			newInstancesReverseSlice = append(newInstancesReverseSlice, assplitarea.GongMarshallUnstaging(stage))
			delete(stage.AsSplitAreas_referenceOrder, assplitarea)
			fieldInitializers, pointersInitializations := assplitarea.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AsSplitAreaMap_Staged_Order[ref] = stage.AsSplitAreaMap_Staged_Order[assplitarea]
			diffs := assplitarea.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, assplitarea)
			delete(stage.AsSplitAreaMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", assplitarea.GetName())
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
	for ref := range stage.AsSplitAreas_reference {
		if _, ok := stage.AsSplitAreas[ref]; !ok {
			assplitareas_deletedInstances = append(assplitareas_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(assplitareas_newInstances)
	lenDeletedInstances += len(assplitareas_deletedInstances)
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
	var cursors_newInstances []*Cursor
	var cursors_deletedInstances []*Cursor

	// parse all staged instances and check if they have a reference
	for cursor := range stage.Cursors {
		if ref, ok := stage.Cursors_reference[cursor]; !ok {
			cursors_newInstances = append(cursors_newInstances, cursor)
			newInstancesSlice = append(newInstancesSlice, cursor.GongMarshallIdentifier(stage))
			if stage.Cursors_referenceOrder == nil {
				stage.Cursors_referenceOrder = make(map[*Cursor]uint)
			}
			stage.Cursors_referenceOrder[cursor] = stage.CursorMap_Staged_Order[cursor]
			newInstancesReverseSlice = append(newInstancesReverseSlice, cursor.GongMarshallUnstaging(stage))
			delete(stage.Cursors_referenceOrder, cursor)
			fieldInitializers, pointersInitializations := cursor.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.CursorMap_Staged_Order[ref] = stage.CursorMap_Staged_Order[cursor]
			diffs := cursor.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, cursor)
			delete(stage.CursorMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", cursor.GetName())
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
	for ref := range stage.Cursors_reference {
		if _, ok := stage.Cursors[ref]; !ok {
			cursors_deletedInstances = append(cursors_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(cursors_newInstances)
	lenDeletedInstances += len(cursors_deletedInstances)
	var favicons_newInstances []*FavIcon
	var favicons_deletedInstances []*FavIcon

	// parse all staged instances and check if they have a reference
	for favicon := range stage.FavIcons {
		if ref, ok := stage.FavIcons_reference[favicon]; !ok {
			favicons_newInstances = append(favicons_newInstances, favicon)
			newInstancesSlice = append(newInstancesSlice, favicon.GongMarshallIdentifier(stage))
			if stage.FavIcons_referenceOrder == nil {
				stage.FavIcons_referenceOrder = make(map[*FavIcon]uint)
			}
			stage.FavIcons_referenceOrder[favicon] = stage.FavIconMap_Staged_Order[favicon]
			newInstancesReverseSlice = append(newInstancesReverseSlice, favicon.GongMarshallUnstaging(stage))
			delete(stage.FavIcons_referenceOrder, favicon)
			fieldInitializers, pointersInitializations := favicon.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FavIconMap_Staged_Order[ref] = stage.FavIconMap_Staged_Order[favicon]
			diffs := favicon.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, favicon)
			delete(stage.FavIconMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", favicon.GetName())
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
	for ref := range stage.FavIcons_reference {
		if _, ok := stage.FavIcons[ref]; !ok {
			favicons_deletedInstances = append(favicons_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(favicons_newInstances)
	lenDeletedInstances += len(favicons_deletedInstances)
	var forms_newInstances []*Form
	var forms_deletedInstances []*Form

	// parse all staged instances and check if they have a reference
	for form := range stage.Forms {
		if ref, ok := stage.Forms_reference[form]; !ok {
			forms_newInstances = append(forms_newInstances, form)
			newInstancesSlice = append(newInstancesSlice, form.GongMarshallIdentifier(stage))
			if stage.Forms_referenceOrder == nil {
				stage.Forms_referenceOrder = make(map[*Form]uint)
			}
			stage.Forms_referenceOrder[form] = stage.FormMap_Staged_Order[form]
			newInstancesReverseSlice = append(newInstancesReverseSlice, form.GongMarshallUnstaging(stage))
			delete(stage.Forms_referenceOrder, form)
			fieldInitializers, pointersInitializations := form.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.FormMap_Staged_Order[ref] = stage.FormMap_Staged_Order[form]
			diffs := form.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, form)
			delete(stage.FormMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", form.GetName())
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
	for ref := range stage.Forms_reference {
		if _, ok := stage.Forms[ref]; !ok {
			forms_deletedInstances = append(forms_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(forms_newInstances)
	lenDeletedInstances += len(forms_deletedInstances)
	var loads_newInstances []*Load
	var loads_deletedInstances []*Load

	// parse all staged instances and check if they have a reference
	for load := range stage.Loads {
		if ref, ok := stage.Loads_reference[load]; !ok {
			loads_newInstances = append(loads_newInstances, load)
			newInstancesSlice = append(newInstancesSlice, load.GongMarshallIdentifier(stage))
			if stage.Loads_referenceOrder == nil {
				stage.Loads_referenceOrder = make(map[*Load]uint)
			}
			stage.Loads_referenceOrder[load] = stage.LoadMap_Staged_Order[load]
			newInstancesReverseSlice = append(newInstancesReverseSlice, load.GongMarshallUnstaging(stage))
			delete(stage.Loads_referenceOrder, load)
			fieldInitializers, pointersInitializations := load.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LoadMap_Staged_Order[ref] = stage.LoadMap_Staged_Order[load]
			diffs := load.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, load)
			delete(stage.LoadMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", load.GetName())
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
	for ref := range stage.Loads_reference {
		if _, ok := stage.Loads[ref]; !ok {
			loads_deletedInstances = append(loads_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(loads_newInstances)
	lenDeletedInstances += len(loads_deletedInstances)
	var logoonthelefts_newInstances []*LogoOnTheLeft
	var logoonthelefts_deletedInstances []*LogoOnTheLeft

	// parse all staged instances and check if they have a reference
	for logoontheleft := range stage.LogoOnTheLefts {
		if ref, ok := stage.LogoOnTheLefts_reference[logoontheleft]; !ok {
			logoonthelefts_newInstances = append(logoonthelefts_newInstances, logoontheleft)
			newInstancesSlice = append(newInstancesSlice, logoontheleft.GongMarshallIdentifier(stage))
			if stage.LogoOnTheLefts_referenceOrder == nil {
				stage.LogoOnTheLefts_referenceOrder = make(map[*LogoOnTheLeft]uint)
			}
			stage.LogoOnTheLefts_referenceOrder[logoontheleft] = stage.LogoOnTheLeftMap_Staged_Order[logoontheleft]
			newInstancesReverseSlice = append(newInstancesReverseSlice, logoontheleft.GongMarshallUnstaging(stage))
			delete(stage.LogoOnTheLefts_referenceOrder, logoontheleft)
			fieldInitializers, pointersInitializations := logoontheleft.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LogoOnTheLeftMap_Staged_Order[ref] = stage.LogoOnTheLeftMap_Staged_Order[logoontheleft]
			diffs := logoontheleft.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, logoontheleft)
			delete(stage.LogoOnTheLeftMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", logoontheleft.GetName())
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
	for ref := range stage.LogoOnTheLefts_reference {
		if _, ok := stage.LogoOnTheLefts[ref]; !ok {
			logoonthelefts_deletedInstances = append(logoonthelefts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(logoonthelefts_newInstances)
	lenDeletedInstances += len(logoonthelefts_deletedInstances)
	var logoontherights_newInstances []*LogoOnTheRight
	var logoontherights_deletedInstances []*LogoOnTheRight

	// parse all staged instances and check if they have a reference
	for logoontheright := range stage.LogoOnTheRights {
		if ref, ok := stage.LogoOnTheRights_reference[logoontheright]; !ok {
			logoontherights_newInstances = append(logoontherights_newInstances, logoontheright)
			newInstancesSlice = append(newInstancesSlice, logoontheright.GongMarshallIdentifier(stage))
			if stage.LogoOnTheRights_referenceOrder == nil {
				stage.LogoOnTheRights_referenceOrder = make(map[*LogoOnTheRight]uint)
			}
			stage.LogoOnTheRights_referenceOrder[logoontheright] = stage.LogoOnTheRightMap_Staged_Order[logoontheright]
			newInstancesReverseSlice = append(newInstancesReverseSlice, logoontheright.GongMarshallUnstaging(stage))
			delete(stage.LogoOnTheRights_referenceOrder, logoontheright)
			fieldInitializers, pointersInitializations := logoontheright.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LogoOnTheRightMap_Staged_Order[ref] = stage.LogoOnTheRightMap_Staged_Order[logoontheright]
			diffs := logoontheright.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, logoontheright)
			delete(stage.LogoOnTheRightMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", logoontheright.GetName())
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
	for ref := range stage.LogoOnTheRights_reference {
		if _, ok := stage.LogoOnTheRights[ref]; !ok {
			logoontherights_deletedInstances = append(logoontherights_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(logoontherights_newInstances)
	lenDeletedInstances += len(logoontherights_deletedInstances)
	var markdowns_newInstances []*Markdown
	var markdowns_deletedInstances []*Markdown

	// parse all staged instances and check if they have a reference
	for markdown := range stage.Markdowns {
		if ref, ok := stage.Markdowns_reference[markdown]; !ok {
			markdowns_newInstances = append(markdowns_newInstances, markdown)
			newInstancesSlice = append(newInstancesSlice, markdown.GongMarshallIdentifier(stage))
			if stage.Markdowns_referenceOrder == nil {
				stage.Markdowns_referenceOrder = make(map[*Markdown]uint)
			}
			stage.Markdowns_referenceOrder[markdown] = stage.MarkdownMap_Staged_Order[markdown]
			newInstancesReverseSlice = append(newInstancesReverseSlice, markdown.GongMarshallUnstaging(stage))
			delete(stage.Markdowns_referenceOrder, markdown)
			fieldInitializers, pointersInitializations := markdown.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.MarkdownMap_Staged_Order[ref] = stage.MarkdownMap_Staged_Order[markdown]
			diffs := markdown.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, markdown)
			delete(stage.MarkdownMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", markdown.GetName())
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
	for ref := range stage.Markdowns_reference {
		if _, ok := stage.Markdowns[ref]; !ok {
			markdowns_deletedInstances = append(markdowns_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(markdowns_newInstances)
	lenDeletedInstances += len(markdowns_deletedInstances)
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
	var splits_newInstances []*Split
	var splits_deletedInstances []*Split

	// parse all staged instances and check if they have a reference
	for split := range stage.Splits {
		if ref, ok := stage.Splits_reference[split]; !ok {
			splits_newInstances = append(splits_newInstances, split)
			newInstancesSlice = append(newInstancesSlice, split.GongMarshallIdentifier(stage))
			if stage.Splits_referenceOrder == nil {
				stage.Splits_referenceOrder = make(map[*Split]uint)
			}
			stage.Splits_referenceOrder[split] = stage.SplitMap_Staged_Order[split]
			newInstancesReverseSlice = append(newInstancesReverseSlice, split.GongMarshallUnstaging(stage))
			delete(stage.Splits_referenceOrder, split)
			fieldInitializers, pointersInitializations := split.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SplitMap_Staged_Order[ref] = stage.SplitMap_Staged_Order[split]
			diffs := split.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, split)
			delete(stage.SplitMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", split.GetName())
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
	for ref := range stage.Splits_reference {
		if _, ok := stage.Splits[ref]; !ok {
			splits_deletedInstances = append(splits_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(splits_newInstances)
	lenDeletedInstances += len(splits_deletedInstances)
	var svgs_newInstances []*Svg
	var svgs_deletedInstances []*Svg

	// parse all staged instances and check if they have a reference
	for svg := range stage.Svgs {
		if ref, ok := stage.Svgs_reference[svg]; !ok {
			svgs_newInstances = append(svgs_newInstances, svg)
			newInstancesSlice = append(newInstancesSlice, svg.GongMarshallIdentifier(stage))
			if stage.Svgs_referenceOrder == nil {
				stage.Svgs_referenceOrder = make(map[*Svg]uint)
			}
			stage.Svgs_referenceOrder[svg] = stage.SvgMap_Staged_Order[svg]
			newInstancesReverseSlice = append(newInstancesReverseSlice, svg.GongMarshallUnstaging(stage))
			delete(stage.Svgs_referenceOrder, svg)
			fieldInitializers, pointersInitializations := svg.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SvgMap_Staged_Order[ref] = stage.SvgMap_Staged_Order[svg]
			diffs := svg.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, svg)
			delete(stage.SvgMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", svg.GetName())
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
	for ref := range stage.Svgs_reference {
		if _, ok := stage.Svgs[ref]; !ok {
			svgs_deletedInstances = append(svgs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(svgs_newInstances)
	lenDeletedInstances += len(svgs_deletedInstances)
	var tables_newInstances []*Table
	var tables_deletedInstances []*Table

	// parse all staged instances and check if they have a reference
	for table := range stage.Tables {
		if ref, ok := stage.Tables_reference[table]; !ok {
			tables_newInstances = append(tables_newInstances, table)
			newInstancesSlice = append(newInstancesSlice, table.GongMarshallIdentifier(stage))
			if stage.Tables_referenceOrder == nil {
				stage.Tables_referenceOrder = make(map[*Table]uint)
			}
			stage.Tables_referenceOrder[table] = stage.TableMap_Staged_Order[table]
			newInstancesReverseSlice = append(newInstancesReverseSlice, table.GongMarshallUnstaging(stage))
			delete(stage.Tables_referenceOrder, table)
			fieldInitializers, pointersInitializations := table.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TableMap_Staged_Order[ref] = stage.TableMap_Staged_Order[table]
			diffs := table.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, table)
			delete(stage.TableMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", table.GetName())
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
	for ref := range stage.Tables_reference {
		if _, ok := stage.Tables[ref]; !ok {
			tables_deletedInstances = append(tables_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(tables_newInstances)
	lenDeletedInstances += len(tables_deletedInstances)
	var titles_newInstances []*Title
	var titles_deletedInstances []*Title

	// parse all staged instances and check if they have a reference
	for title := range stage.Titles {
		if ref, ok := stage.Titles_reference[title]; !ok {
			titles_newInstances = append(titles_newInstances, title)
			newInstancesSlice = append(newInstancesSlice, title.GongMarshallIdentifier(stage))
			if stage.Titles_referenceOrder == nil {
				stage.Titles_referenceOrder = make(map[*Title]uint)
			}
			stage.Titles_referenceOrder[title] = stage.TitleMap_Staged_Order[title]
			newInstancesReverseSlice = append(newInstancesReverseSlice, title.GongMarshallUnstaging(stage))
			delete(stage.Titles_referenceOrder, title)
			fieldInitializers, pointersInitializations := title.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TitleMap_Staged_Order[ref] = stage.TitleMap_Staged_Order[title]
			diffs := title.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, title)
			delete(stage.TitleMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", title.GetName())
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
	for ref := range stage.Titles_reference {
		if _, ok := stage.Titles[ref]; !ok {
			titles_deletedInstances = append(titles_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(titles_newInstances)
	lenDeletedInstances += len(titles_deletedInstances)
	var tones_newInstances []*Tone
	var tones_deletedInstances []*Tone

	// parse all staged instances and check if they have a reference
	for tone := range stage.Tones {
		if ref, ok := stage.Tones_reference[tone]; !ok {
			tones_newInstances = append(tones_newInstances, tone)
			newInstancesSlice = append(newInstancesSlice, tone.GongMarshallIdentifier(stage))
			if stage.Tones_referenceOrder == nil {
				stage.Tones_referenceOrder = make(map[*Tone]uint)
			}
			stage.Tones_referenceOrder[tone] = stage.ToneMap_Staged_Order[tone]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tone.GongMarshallUnstaging(stage))
			delete(stage.Tones_referenceOrder, tone)
			fieldInitializers, pointersInitializations := tone.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ToneMap_Staged_Order[ref] = stage.ToneMap_Staged_Order[tone]
			diffs := tone.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tone)
			delete(stage.ToneMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", tone.GetName())
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
	for ref := range stage.Tones_reference {
		if _, ok := stage.Tones[ref]; !ok {
			tones_deletedInstances = append(tones_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(tones_newInstances)
	lenDeletedInstances += len(tones_deletedInstances)
	var trees_newInstances []*Tree
	var trees_deletedInstances []*Tree

	// parse all staged instances and check if they have a reference
	for tree := range stage.Trees {
		if ref, ok := stage.Trees_reference[tree]; !ok {
			trees_newInstances = append(trees_newInstances, tree)
			newInstancesSlice = append(newInstancesSlice, tree.GongMarshallIdentifier(stage))
			if stage.Trees_referenceOrder == nil {
				stage.Trees_referenceOrder = make(map[*Tree]uint)
			}
			stage.Trees_referenceOrder[tree] = stage.TreeMap_Staged_Order[tree]
			newInstancesReverseSlice = append(newInstancesReverseSlice, tree.GongMarshallUnstaging(stage))
			delete(stage.Trees_referenceOrder, tree)
			fieldInitializers, pointersInitializations := tree.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.TreeMap_Staged_Order[ref] = stage.TreeMap_Staged_Order[tree]
			diffs := tree.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, tree)
			delete(stage.TreeMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", tree.GetName())
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
	for ref := range stage.Trees_reference {
		if _, ok := stage.Trees[ref]; !ok {
			trees_deletedInstances = append(trees_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(trees_newInstances)
	lenDeletedInstances += len(trees_deletedInstances)
	var views_newInstances []*View
	var views_deletedInstances []*View

	// parse all staged instances and check if they have a reference
	for view := range stage.Views {
		if ref, ok := stage.Views_reference[view]; !ok {
			views_newInstances = append(views_newInstances, view)
			newInstancesSlice = append(newInstancesSlice, view.GongMarshallIdentifier(stage))
			if stage.Views_referenceOrder == nil {
				stage.Views_referenceOrder = make(map[*View]uint)
			}
			stage.Views_referenceOrder[view] = stage.ViewMap_Staged_Order[view]
			newInstancesReverseSlice = append(newInstancesReverseSlice, view.GongMarshallUnstaging(stage))
			delete(stage.Views_referenceOrder, view)
			fieldInitializers, pointersInitializations := view.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ViewMap_Staged_Order[ref] = stage.ViewMap_Staged_Order[view]
			diffs := view.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, view)
			delete(stage.ViewMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", view.GetName())
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
	for ref := range stage.Views_reference {
		if _, ok := stage.Views[ref]; !ok {
			views_deletedInstances = append(views_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(views_newInstances)
	lenDeletedInstances += len(views_deletedInstances)
	var xlsxs_newInstances []*Xlsx
	var xlsxs_deletedInstances []*Xlsx

	// parse all staged instances and check if they have a reference
	for xlsx := range stage.Xlsxs {
		if ref, ok := stage.Xlsxs_reference[xlsx]; !ok {
			xlsxs_newInstances = append(xlsxs_newInstances, xlsx)
			newInstancesSlice = append(newInstancesSlice, xlsx.GongMarshallIdentifier(stage))
			if stage.Xlsxs_referenceOrder == nil {
				stage.Xlsxs_referenceOrder = make(map[*Xlsx]uint)
			}
			stage.Xlsxs_referenceOrder[xlsx] = stage.XlsxMap_Staged_Order[xlsx]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xlsx.GongMarshallUnstaging(stage))
			delete(stage.Xlsxs_referenceOrder, xlsx)
			fieldInitializers, pointersInitializations := xlsx.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XlsxMap_Staged_Order[ref] = stage.XlsxMap_Staged_Order[xlsx]
			diffs := xlsx.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xlsx)
			delete(stage.XlsxMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", xlsx.GetName())
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
	for ref := range stage.Xlsxs_reference {
		if _, ok := stage.Xlsxs[ref]; !ok {
			xlsxs_deletedInstances = append(xlsxs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(xlsxs_newInstances)
	lenDeletedInstances += len(xlsxs_deletedInstances)

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
	stage.AsSplits_reference = make(map[*AsSplit]*AsSplit)
	stage.AsSplits_referenceOrder = make(map[*AsSplit]uint) // diff Unstage needs the reference order
	for instance := range stage.AsSplits {
		stage.AsSplits_reference[instance] = instance.GongCopy().(*AsSplit)
		stage.AsSplits_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.AsSplitAreas_reference = make(map[*AsSplitArea]*AsSplitArea)
	stage.AsSplitAreas_referenceOrder = make(map[*AsSplitArea]uint) // diff Unstage needs the reference order
	for instance := range stage.AsSplitAreas {
		stage.AsSplitAreas_reference[instance] = instance.GongCopy().(*AsSplitArea)
		stage.AsSplitAreas_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Buttons_reference = make(map[*Button]*Button)
	stage.Buttons_referenceOrder = make(map[*Button]uint) // diff Unstage needs the reference order
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance.GongCopy().(*Button)
		stage.Buttons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Cursors_reference = make(map[*Cursor]*Cursor)
	stage.Cursors_referenceOrder = make(map[*Cursor]uint) // diff Unstage needs the reference order
	for instance := range stage.Cursors {
		stage.Cursors_reference[instance] = instance.GongCopy().(*Cursor)
		stage.Cursors_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.FavIcons_reference = make(map[*FavIcon]*FavIcon)
	stage.FavIcons_referenceOrder = make(map[*FavIcon]uint) // diff Unstage needs the reference order
	for instance := range stage.FavIcons {
		stage.FavIcons_reference[instance] = instance.GongCopy().(*FavIcon)
		stage.FavIcons_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Forms_reference = make(map[*Form]*Form)
	stage.Forms_referenceOrder = make(map[*Form]uint) // diff Unstage needs the reference order
	for instance := range stage.Forms {
		stage.Forms_reference[instance] = instance.GongCopy().(*Form)
		stage.Forms_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Loads_reference = make(map[*Load]*Load)
	stage.Loads_referenceOrder = make(map[*Load]uint) // diff Unstage needs the reference order
	for instance := range stage.Loads {
		stage.Loads_reference[instance] = instance.GongCopy().(*Load)
		stage.Loads_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.LogoOnTheLefts_reference = make(map[*LogoOnTheLeft]*LogoOnTheLeft)
	stage.LogoOnTheLefts_referenceOrder = make(map[*LogoOnTheLeft]uint) // diff Unstage needs the reference order
	for instance := range stage.LogoOnTheLefts {
		stage.LogoOnTheLefts_reference[instance] = instance.GongCopy().(*LogoOnTheLeft)
		stage.LogoOnTheLefts_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.LogoOnTheRights_reference = make(map[*LogoOnTheRight]*LogoOnTheRight)
	stage.LogoOnTheRights_referenceOrder = make(map[*LogoOnTheRight]uint) // diff Unstage needs the reference order
	for instance := range stage.LogoOnTheRights {
		stage.LogoOnTheRights_reference[instance] = instance.GongCopy().(*LogoOnTheRight)
		stage.LogoOnTheRights_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Markdowns_reference = make(map[*Markdown]*Markdown)
	stage.Markdowns_referenceOrder = make(map[*Markdown]uint) // diff Unstage needs the reference order
	for instance := range stage.Markdowns {
		stage.Markdowns_reference[instance] = instance.GongCopy().(*Markdown)
		stage.Markdowns_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Sliders_reference = make(map[*Slider]*Slider)
	stage.Sliders_referenceOrder = make(map[*Slider]uint) // diff Unstage needs the reference order
	for instance := range stage.Sliders {
		stage.Sliders_reference[instance] = instance.GongCopy().(*Slider)
		stage.Sliders_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Splits_reference = make(map[*Split]*Split)
	stage.Splits_referenceOrder = make(map[*Split]uint) // diff Unstage needs the reference order
	for instance := range stage.Splits {
		stage.Splits_reference[instance] = instance.GongCopy().(*Split)
		stage.Splits_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Svgs_reference = make(map[*Svg]*Svg)
	stage.Svgs_referenceOrder = make(map[*Svg]uint) // diff Unstage needs the reference order
	for instance := range stage.Svgs {
		stage.Svgs_reference[instance] = instance.GongCopy().(*Svg)
		stage.Svgs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	stage.Tables_referenceOrder = make(map[*Table]uint) // diff Unstage needs the reference order
	for instance := range stage.Tables {
		stage.Tables_reference[instance] = instance.GongCopy().(*Table)
		stage.Tables_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Titles_reference = make(map[*Title]*Title)
	stage.Titles_referenceOrder = make(map[*Title]uint) // diff Unstage needs the reference order
	for instance := range stage.Titles {
		stage.Titles_reference[instance] = instance.GongCopy().(*Title)
		stage.Titles_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Tones_reference = make(map[*Tone]*Tone)
	stage.Tones_referenceOrder = make(map[*Tone]uint) // diff Unstage needs the reference order
	for instance := range stage.Tones {
		stage.Tones_reference[instance] = instance.GongCopy().(*Tone)
		stage.Tones_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Trees_reference = make(map[*Tree]*Tree)
	stage.Trees_referenceOrder = make(map[*Tree]uint) // diff Unstage needs the reference order
	for instance := range stage.Trees {
		stage.Trees_reference[instance] = instance.GongCopy().(*Tree)
		stage.Trees_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Views_reference = make(map[*View]*View)
	stage.Views_referenceOrder = make(map[*View]uint) // diff Unstage needs the reference order
	for instance := range stage.Views {
		stage.Views_reference[instance] = instance.GongCopy().(*View)
		stage.Views_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Xlsxs_reference = make(map[*Xlsx]*Xlsx)
	stage.Xlsxs_referenceOrder = make(map[*Xlsx]uint) // diff Unstage needs the reference order
	for instance := range stage.Xlsxs {
		stage.Xlsxs_reference[instance] = instance.GongCopy().(*Xlsx)
		stage.Xlsxs_referenceOrder[instance] = instance.GongGetOrder(stage)
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
func (assplit *AsSplit) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AsSplitMap_Staged_Order[assplit]; ok {
		return order
	}
	return stage.AsSplits_referenceOrder[assplit]
}

func (assplit *AsSplit) GongGetReferenceOrder(stage *Stage) uint {
	return stage.AsSplits_referenceOrder[assplit]
}

func (assplitarea *AsSplitArea) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AsSplitAreaMap_Staged_Order[assplitarea]; ok {
		return order
	}
	return stage.AsSplitAreas_referenceOrder[assplitarea]
}

func (assplitarea *AsSplitArea) GongGetReferenceOrder(stage *Stage) uint {
	return stage.AsSplitAreas_referenceOrder[assplitarea]
}

func (button *Button) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ButtonMap_Staged_Order[button]; ok {
		return order
	}
	return stage.Buttons_referenceOrder[button]
}

func (button *Button) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Buttons_referenceOrder[button]
}

func (cursor *Cursor) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.CursorMap_Staged_Order[cursor]; ok {
		return order
	}
	return stage.Cursors_referenceOrder[cursor]
}

func (cursor *Cursor) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Cursors_referenceOrder[cursor]
}

func (favicon *FavIcon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FavIconMap_Staged_Order[favicon]; ok {
		return order
	}
	return stage.FavIcons_referenceOrder[favicon]
}

func (favicon *FavIcon) GongGetReferenceOrder(stage *Stage) uint {
	return stage.FavIcons_referenceOrder[favicon]
}

func (form *Form) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FormMap_Staged_Order[form]; ok {
		return order
	}
	return stage.Forms_referenceOrder[form]
}

func (form *Form) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Forms_referenceOrder[form]
}

func (load *Load) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LoadMap_Staged_Order[load]; ok {
		return order
	}
	return stage.Loads_referenceOrder[load]
}

func (load *Load) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Loads_referenceOrder[load]
}

func (logoontheleft *LogoOnTheLeft) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LogoOnTheLeftMap_Staged_Order[logoontheleft]; ok {
		return order
	}
	return stage.LogoOnTheLefts_referenceOrder[logoontheleft]
}

func (logoontheleft *LogoOnTheLeft) GongGetReferenceOrder(stage *Stage) uint {
	return stage.LogoOnTheLefts_referenceOrder[logoontheleft]
}

func (logoontheright *LogoOnTheRight) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LogoOnTheRightMap_Staged_Order[logoontheright]; ok {
		return order
	}
	return stage.LogoOnTheRights_referenceOrder[logoontheright]
}

func (logoontheright *LogoOnTheRight) GongGetReferenceOrder(stage *Stage) uint {
	return stage.LogoOnTheRights_referenceOrder[logoontheright]
}

func (markdown *Markdown) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.MarkdownMap_Staged_Order[markdown]; ok {
		return order
	}
	return stage.Markdowns_referenceOrder[markdown]
}

func (markdown *Markdown) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Markdowns_referenceOrder[markdown]
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

func (split *Split) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SplitMap_Staged_Order[split]; ok {
		return order
	}
	return stage.Splits_referenceOrder[split]
}

func (split *Split) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Splits_referenceOrder[split]
}

func (svg *Svg) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SvgMap_Staged_Order[svg]; ok {
		return order
	}
	return stage.Svgs_referenceOrder[svg]
}

func (svg *Svg) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Svgs_referenceOrder[svg]
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TableMap_Staged_Order[table]; ok {
		return order
	}
	return stage.Tables_referenceOrder[table]
}

func (table *Table) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Tables_referenceOrder[table]
}

func (title *Title) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TitleMap_Staged_Order[title]; ok {
		return order
	}
	return stage.Titles_referenceOrder[title]
}

func (title *Title) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Titles_referenceOrder[title]
}

func (tone *Tone) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ToneMap_Staged_Order[tone]; ok {
		return order
	}
	return stage.Tones_referenceOrder[tone]
}

func (tone *Tone) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Tones_referenceOrder[tone]
}

func (tree *Tree) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.TreeMap_Staged_Order[tree]; ok {
		return order
	}
	return stage.Trees_referenceOrder[tree]
}

func (tree *Tree) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Trees_referenceOrder[tree]
}

func (view *View) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ViewMap_Staged_Order[view]; ok {
		return order
	}
	return stage.Views_referenceOrder[view]
}

func (view *View) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Views_referenceOrder[view]
}

func (xlsx *Xlsx) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XlsxMap_Staged_Order[xlsx]; ok {
		return order
	}
	return stage.Xlsxs_referenceOrder[xlsx]
}

func (xlsx *Xlsx) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Xlsxs_referenceOrder[xlsx]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (assplit *AsSplit) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplit.GongGetGongstructName(), assplit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assplit *AsSplit) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplit.GongGetGongstructName(), assplit.GongGetReferenceOrder(stage))
}

func (assplitarea *AsSplitArea) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplitarea.GongGetGongstructName(), assplitarea.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assplitarea *AsSplitArea) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplitarea.GongGetGongstructName(), assplitarea.GongGetReferenceOrder(stage))
}

func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetReferenceOrder(stage))
}

func (cursor *Cursor) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cursor.GongGetGongstructName(), cursor.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cursor *Cursor) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cursor.GongGetGongstructName(), cursor.GongGetReferenceOrder(stage))
}

func (favicon *FavIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", favicon.GongGetGongstructName(), favicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (favicon *FavIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", favicon.GongGetGongstructName(), favicon.GongGetReferenceOrder(stage))
}

func (form *Form) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", form.GongGetGongstructName(), form.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (form *Form) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", form.GongGetGongstructName(), form.GongGetReferenceOrder(stage))
}

func (load *Load) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", load.GongGetGongstructName(), load.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (load *Load) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", load.GongGetGongstructName(), load.GongGetReferenceOrder(stage))
}

func (logoontheleft *LogoOnTheLeft) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheleft.GongGetGongstructName(), logoontheleft.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (logoontheleft *LogoOnTheLeft) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheleft.GongGetGongstructName(), logoontheleft.GongGetReferenceOrder(stage))
}

func (logoontheright *LogoOnTheRight) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheright.GongGetGongstructName(), logoontheright.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (logoontheright *LogoOnTheRight) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheright.GongGetGongstructName(), logoontheright.GongGetReferenceOrder(stage))
}

func (markdown *Markdown) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", markdown.GongGetGongstructName(), markdown.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (markdown *Markdown) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", markdown.GongGetGongstructName(), markdown.GongGetReferenceOrder(stage))
}

func (slider *Slider) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slider *Slider) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetReferenceOrder(stage))
}

func (split *Split) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", split.GongGetGongstructName(), split.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (split *Split) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", split.GongGetGongstructName(), split.GongGetReferenceOrder(stage))
}

func (svg *Svg) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svg *Svg) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetReferenceOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetReferenceOrder(stage))
}

func (title *Title) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", title.GongGetGongstructName(), title.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (title *Title) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", title.GongGetGongstructName(), title.GongGetReferenceOrder(stage))
}

func (tone *Tone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tone.GongGetGongstructName(), tone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tone *Tone) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tone.GongGetGongstructName(), tone.GongGetReferenceOrder(stage))
}

func (tree *Tree) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tree *Tree) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetReferenceOrder(stage))
}

func (view *View) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", view.GongGetGongstructName(), view.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (view *View) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", view.GongGetGongstructName(), view.GongGetReferenceOrder(stage))
}

func (xlsx *Xlsx) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsx.GongGetGongstructName(), xlsx.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlsx *Xlsx) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsx.GongGetGongstructName(), xlsx.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (assplit *AsSplit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assplit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", assplit.Name)
	return
}
func (assplitarea *AsSplitArea) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplitArea")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", assplitarea.Name)
	return
}
func (button *Button) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Button")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", button.Name)
	return
}
func (cursor *Cursor) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cursor.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cursor")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", cursor.Name)
	return
}
func (favicon *FavIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", favicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FavIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", favicon.Name)
	return
}
func (form *Form) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", form.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Form")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", form.Name)
	return
}
func (load *Load) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", load.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Load")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", load.Name)
	return
}
func (logoontheleft *LogoOnTheLeft) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LogoOnTheLeft")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", logoontheleft.Name)
	return
}
func (logoontheright *LogoOnTheRight) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LogoOnTheRight")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", logoontheright.Name)
	return
}
func (markdown *Markdown) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", markdown.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Markdown")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", markdown.Name)
	return
}
func (slider *Slider) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slider.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slider")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", slider.Name)
	return
}
func (split *Split) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", split.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Split")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", split.Name)
	return
}
func (svg *Svg) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svg.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Svg")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", svg.Name)
	return
}
func (table *Table) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", table.Name)
	return
}
func (title *Title) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", title.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Title")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", title.Name)
	return
}
func (tone *Tone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tone.Name)
	return
}
func (tree *Tree) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tree.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tree")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", tree.Name)
	return
}
func (view *View) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", view.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "View")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", view.Name)
	return
}
func (xlsx *Xlsx) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xlsx")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlsx.Name)
	return
}

// insertion point for unstaging
func (assplit *AsSplit) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assplit.GongGetReferenceIdentifier(stage))
	return
}
func (assplitarea *AsSplitArea) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assplitarea.GongGetReferenceIdentifier(stage))
	return
}
func (button *Button) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetReferenceIdentifier(stage))
	return
}
func (cursor *Cursor) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cursor.GongGetReferenceIdentifier(stage))
	return
}
func (favicon *FavIcon) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", favicon.GongGetReferenceIdentifier(stage))
	return
}
func (form *Form) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", form.GongGetReferenceIdentifier(stage))
	return
}
func (load *Load) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", load.GongGetReferenceIdentifier(stage))
	return
}
func (logoontheleft *LogoOnTheLeft) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", logoontheleft.GongGetReferenceIdentifier(stage))
	return
}
func (logoontheright *LogoOnTheRight) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", logoontheright.GongGetReferenceIdentifier(stage))
	return
}
func (markdown *Markdown) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", markdown.GongGetReferenceIdentifier(stage))
	return
}
func (slider *Slider) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slider.GongGetReferenceIdentifier(stage))
	return
}
func (split *Split) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", split.GongGetReferenceIdentifier(stage))
	return
}
func (svg *Svg) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svg.GongGetReferenceIdentifier(stage))
	return
}
func (table *Table) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetReferenceIdentifier(stage))
	return
}
func (title *Title) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", title.GongGetReferenceIdentifier(stage))
	return
}
func (tone *Tone) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tone.GongGetReferenceIdentifier(stage))
	return
}
func (tree *Tree) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tree.GongGetReferenceIdentifier(stage))
	return
}
func (view *View) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", view.GongGetReferenceIdentifier(stage))
	return
}
func (xlsx *Xlsx) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlsx.GongGetReferenceIdentifier(stage))
	return
}
