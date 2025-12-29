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

func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	// insertion point per named struct
	var assplits_newInstances []*AsSplit
	var assplits_deletedInstances []*AsSplit

	// parse all staged instances and check if they have a reference
	for assplit := range stage.AsSplits {
		if ref, ok := stage.AsSplits_reference[assplit]; !ok {
			assplits_newInstances = append(assplits_newInstances, assplit)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of AsSplit "+assplit.Name,
				)
			}
		} else {
			diffs := assplit.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of AsSplit \""+assplit.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for assplit := range stage.AsSplits_reference {
		if _, ok := stage.AsSplits[assplit]; !ok {
			assplits_deletedInstances = append(assplits_deletedInstances, assplit)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of AsSplit "+assplit.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of AsSplitArea "+assplitarea.Name,
				)
			}
		} else {
			diffs := assplitarea.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of AsSplitArea \""+assplitarea.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for assplitarea := range stage.AsSplitAreas_reference {
		if _, ok := stage.AsSplitAreas[assplitarea]; !ok {
			assplitareas_deletedInstances = append(assplitareas_deletedInstances, assplitarea)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of AsSplitArea "+assplitarea.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Button "+button.Name,
				)
			}
		} else {
			diffs := button.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Button \""+button.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for button := range stage.Buttons_reference {
		if _, ok := stage.Buttons[button]; !ok {
			buttons_deletedInstances = append(buttons_deletedInstances, button)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Button "+button.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Cursor "+cursor.Name,
				)
			}
		} else {
			diffs := cursor.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Cursor \""+cursor.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for cursor := range stage.Cursors_reference {
		if _, ok := stage.Cursors[cursor]; !ok {
			cursors_deletedInstances = append(cursors_deletedInstances, cursor)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Cursor "+cursor.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of FavIcon "+favicon.Name,
				)
			}
		} else {
			diffs := favicon.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of FavIcon \""+favicon.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for favicon := range stage.FavIcons_reference {
		if _, ok := stage.FavIcons[favicon]; !ok {
			favicons_deletedInstances = append(favicons_deletedInstances, favicon)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of FavIcon "+favicon.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Form "+form.Name,
				)
			}
		} else {
			diffs := form.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Form \""+form.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for form := range stage.Forms_reference {
		if _, ok := stage.Forms[form]; !ok {
			forms_deletedInstances = append(forms_deletedInstances, form)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Form "+form.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Load "+load.Name,
				)
			}
		} else {
			diffs := load.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Load \""+load.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for load := range stage.Loads_reference {
		if _, ok := stage.Loads[load]; !ok {
			loads_deletedInstances = append(loads_deletedInstances, load)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Load "+load.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of LogoOnTheLeft "+logoontheleft.Name,
				)
			}
		} else {
			diffs := logoontheleft.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of LogoOnTheLeft \""+logoontheleft.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for logoontheleft := range stage.LogoOnTheLefts_reference {
		if _, ok := stage.LogoOnTheLefts[logoontheleft]; !ok {
			logoonthelefts_deletedInstances = append(logoonthelefts_deletedInstances, logoontheleft)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of LogoOnTheLeft "+logoontheleft.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of LogoOnTheRight "+logoontheright.Name,
				)
			}
		} else {
			diffs := logoontheright.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of LogoOnTheRight \""+logoontheright.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for logoontheright := range stage.LogoOnTheRights_reference {
		if _, ok := stage.LogoOnTheRights[logoontheright]; !ok {
			logoontherights_deletedInstances = append(logoontherights_deletedInstances, logoontheright)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of LogoOnTheRight "+logoontheright.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Markdown "+markdown.Name,
				)
			}
		} else {
			diffs := markdown.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Markdown \""+markdown.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for markdown := range stage.Markdowns_reference {
		if _, ok := stage.Markdowns[markdown]; !ok {
			markdowns_deletedInstances = append(markdowns_deletedInstances, markdown)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Markdown "+markdown.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Slider "+slider.Name,
				)
			}
		} else {
			diffs := slider.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Slider \""+slider.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
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
	var splits_newInstances []*Split
	var splits_deletedInstances []*Split

	// parse all staged instances and check if they have a reference
	for split := range stage.Splits {
		if ref, ok := stage.Splits_reference[split]; !ok {
			splits_newInstances = append(splits_newInstances, split)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Split "+split.Name,
				)
			}
		} else {
			diffs := split.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Split \""+split.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for split := range stage.Splits_reference {
		if _, ok := stage.Splits[split]; !ok {
			splits_deletedInstances = append(splits_deletedInstances, split)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Split "+split.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Svg "+svg.Name,
				)
			}
		} else {
			diffs := svg.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Svg \""+svg.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for svg := range stage.Svgs_reference {
		if _, ok := stage.Svgs[svg]; !ok {
			svgs_deletedInstances = append(svgs_deletedInstances, svg)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Svg "+svg.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Table "+table.Name,
				)
			}
		} else {
			diffs := table.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Table \""+table.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for table := range stage.Tables_reference {
		if _, ok := stage.Tables[table]; !ok {
			tables_deletedInstances = append(tables_deletedInstances, table)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Table "+table.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Title "+title.Name,
				)
			}
		} else {
			diffs := title.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Title \""+title.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for title := range stage.Titles_reference {
		if _, ok := stage.Titles[title]; !ok {
			titles_deletedInstances = append(titles_deletedInstances, title)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Title "+title.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Tone "+tone.Name,
				)
			}
		} else {
			diffs := tone.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Tone \""+tone.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tone := range stage.Tones_reference {
		if _, ok := stage.Tones[tone]; !ok {
			tones_deletedInstances = append(tones_deletedInstances, tone)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Tone "+tone.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Tree "+tree.Name,
				)
			}
		} else {
			diffs := tree.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Tree \""+tree.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for tree := range stage.Trees_reference {
		if _, ok := stage.Trees[tree]; !ok {
			trees_deletedInstances = append(trees_deletedInstances, tree)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Tree "+tree.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of View "+view.Name,
				)
			}
		} else {
			diffs := view.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of View \""+view.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for view := range stage.Views_reference {
		if _, ok := stage.Views[view]; !ok {
			views_deletedInstances = append(views_deletedInstances, view)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of View "+view.Name,
				)
			}
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
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected new instance of Xlsx "+xlsx.Name,
				)
			}
		} else {
			diffs := xlsx.GongDiff(ref)
			if len(diffs) > 0 {
				if stage.GetProbeIF() != nil {
					stage.GetProbeIF().AddNotification(
						time.Now(),
						"Commit detected modified instance of Xlsx \""+xlsx.Name+"\" diffs on fields: \""+strings.Join(diffs, ", \"")+"\"",
					)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for xlsx := range stage.Xlsxs_reference {
		if _, ok := stage.Xlsxs[xlsx]; !ok {
			xlsxs_deletedInstances = append(xlsxs_deletedInstances, xlsx)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Commit detected deleted instance of Xlsx "+xlsx.Name,
				)
			}
		}
	}

	lenNewInstances += len(xlsxs_newInstances)
	lenDeletedInstances += len(xlsxs_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {
		// if stage.GetProbeIF() != nil {
		// 	stage.GetProbeIF().CommitNotificationTable()
		// }
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.AsSplits_reference = make(map[*AsSplit]*AsSplit)
	for instance := range stage.AsSplits {
		stage.AsSplits_reference[instance] = instance.GongCopy().(*AsSplit)
	}

	stage.AsSplitAreas_reference = make(map[*AsSplitArea]*AsSplitArea)
	for instance := range stage.AsSplitAreas {
		stage.AsSplitAreas_reference[instance] = instance.GongCopy().(*AsSplitArea)
	}

	stage.Buttons_reference = make(map[*Button]*Button)
	for instance := range stage.Buttons {
		stage.Buttons_reference[instance] = instance.GongCopy().(*Button)
	}

	stage.Cursors_reference = make(map[*Cursor]*Cursor)
	for instance := range stage.Cursors {
		stage.Cursors_reference[instance] = instance.GongCopy().(*Cursor)
	}

	stage.FavIcons_reference = make(map[*FavIcon]*FavIcon)
	for instance := range stage.FavIcons {
		stage.FavIcons_reference[instance] = instance.GongCopy().(*FavIcon)
	}

	stage.Forms_reference = make(map[*Form]*Form)
	for instance := range stage.Forms {
		stage.Forms_reference[instance] = instance.GongCopy().(*Form)
	}

	stage.Loads_reference = make(map[*Load]*Load)
	for instance := range stage.Loads {
		stage.Loads_reference[instance] = instance.GongCopy().(*Load)
	}

	stage.LogoOnTheLefts_reference = make(map[*LogoOnTheLeft]*LogoOnTheLeft)
	for instance := range stage.LogoOnTheLefts {
		stage.LogoOnTheLefts_reference[instance] = instance.GongCopy().(*LogoOnTheLeft)
	}

	stage.LogoOnTheRights_reference = make(map[*LogoOnTheRight]*LogoOnTheRight)
	for instance := range stage.LogoOnTheRights {
		stage.LogoOnTheRights_reference[instance] = instance.GongCopy().(*LogoOnTheRight)
	}

	stage.Markdowns_reference = make(map[*Markdown]*Markdown)
	for instance := range stage.Markdowns {
		stage.Markdowns_reference[instance] = instance.GongCopy().(*Markdown)
	}

	stage.Sliders_reference = make(map[*Slider]*Slider)
	for instance := range stage.Sliders {
		stage.Sliders_reference[instance] = instance.GongCopy().(*Slider)
	}

	stage.Splits_reference = make(map[*Split]*Split)
	for instance := range stage.Splits {
		stage.Splits_reference[instance] = instance.GongCopy().(*Split)
	}

	stage.Svgs_reference = make(map[*Svg]*Svg)
	for instance := range stage.Svgs {
		stage.Svgs_reference[instance] = instance.GongCopy().(*Svg)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	for instance := range stage.Tables {
		stage.Tables_reference[instance] = instance.GongCopy().(*Table)
	}

	stage.Titles_reference = make(map[*Title]*Title)
	for instance := range stage.Titles {
		stage.Titles_reference[instance] = instance.GongCopy().(*Title)
	}

	stage.Tones_reference = make(map[*Tone]*Tone)
	for instance := range stage.Tones {
		stage.Tones_reference[instance] = instance.GongCopy().(*Tone)
	}

	stage.Trees_reference = make(map[*Tree]*Tree)
	for instance := range stage.Trees {
		stage.Trees_reference[instance] = instance.GongCopy().(*Tree)
	}

	stage.Views_reference = make(map[*View]*View)
	for instance := range stage.Views {
		stage.Views_reference[instance] = instance.GongCopy().(*View)
	}

	stage.Xlsxs_reference = make(map[*Xlsx]*Xlsx)
	for instance := range stage.Xlsxs {
		stage.Xlsxs_reference[instance] = instance.GongCopy().(*Xlsx)
	}

}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (assplit *AsSplit) GongGetOrder(stage *Stage) uint {
	return stage.AsSplitMap_Staged_Order[assplit]
}

func (assplitarea *AsSplitArea) GongGetOrder(stage *Stage) uint {
	return stage.AsSplitAreaMap_Staged_Order[assplitarea]
}

func (button *Button) GongGetOrder(stage *Stage) uint {
	return stage.ButtonMap_Staged_Order[button]
}

func (cursor *Cursor) GongGetOrder(stage *Stage) uint {
	return stage.CursorMap_Staged_Order[cursor]
}

func (favicon *FavIcon) GongGetOrder(stage *Stage) uint {
	return stage.FavIconMap_Staged_Order[favicon]
}

func (form *Form) GongGetOrder(stage *Stage) uint {
	return stage.FormMap_Staged_Order[form]
}

func (load *Load) GongGetOrder(stage *Stage) uint {
	return stage.LoadMap_Staged_Order[load]
}

func (logoontheleft *LogoOnTheLeft) GongGetOrder(stage *Stage) uint {
	return stage.LogoOnTheLeftMap_Staged_Order[logoontheleft]
}

func (logoontheright *LogoOnTheRight) GongGetOrder(stage *Stage) uint {
	return stage.LogoOnTheRightMap_Staged_Order[logoontheright]
}

func (markdown *Markdown) GongGetOrder(stage *Stage) uint {
	return stage.MarkdownMap_Staged_Order[markdown]
}

func (slider *Slider) GongGetOrder(stage *Stage) uint {
	return stage.SliderMap_Staged_Order[slider]
}

func (split *Split) GongGetOrder(stage *Stage) uint {
	return stage.SplitMap_Staged_Order[split]
}

func (svg *Svg) GongGetOrder(stage *Stage) uint {
	return stage.SvgMap_Staged_Order[svg]
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	return stage.TableMap_Staged_Order[table]
}

func (title *Title) GongGetOrder(stage *Stage) uint {
	return stage.TitleMap_Staged_Order[title]
}

func (tone *Tone) GongGetOrder(stage *Stage) uint {
	return stage.ToneMap_Staged_Order[tone]
}

func (tree *Tree) GongGetOrder(stage *Stage) uint {
	return stage.TreeMap_Staged_Order[tree]
}

func (view *View) GongGetOrder(stage *Stage) uint {
	return stage.ViewMap_Staged_Order[view]
}

func (xlsx *Xlsx) GongGetOrder(stage *Stage) uint {
	return stage.XlsxMap_Staged_Order[xlsx]
}


// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (assplit *AsSplit) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplit.GongGetGongstructName(), assplit.GongGetOrder(stage))
}

func (assplitarea *AsSplitArea) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplitarea.GongGetGongstructName(), assplitarea.GongGetOrder(stage))
}

func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

func (cursor *Cursor) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cursor.GongGetGongstructName(), cursor.GongGetOrder(stage))
}

func (favicon *FavIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", favicon.GongGetGongstructName(), favicon.GongGetOrder(stage))
}

func (form *Form) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", form.GongGetGongstructName(), form.GongGetOrder(stage))
}

func (load *Load) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", load.GongGetGongstructName(), load.GongGetOrder(stage))
}

func (logoontheleft *LogoOnTheLeft) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheleft.GongGetGongstructName(), logoontheleft.GongGetOrder(stage))
}

func (logoontheright *LogoOnTheRight) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheright.GongGetGongstructName(), logoontheright.GongGetOrder(stage))
}

func (markdown *Markdown) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", markdown.GongGetGongstructName(), markdown.GongGetOrder(stage))
}

func (slider *Slider) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetOrder(stage))
}

func (split *Split) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", split.GongGetGongstructName(), split.GongGetOrder(stage))
}

func (svg *Svg) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

func (title *Title) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", title.GongGetGongstructName(), title.GongGetOrder(stage))
}

func (tone *Tone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tone.GongGetGongstructName(), tone.GongGetOrder(stage))
}

func (tree *Tree) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetOrder(stage))
}

func (view *View) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", view.GongGetGongstructName(), view.GongGetOrder(stage))
}

func (xlsx *Xlsx) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsx.GongGetGongstructName(), xlsx.GongGetOrder(stage))
}

