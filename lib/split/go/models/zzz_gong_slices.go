// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
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
	// Compute reverse map for named struct AsSplit
	// insertion point per field
	stage.AsSplit_AsSplitAreas_reverseMap = make(map[*AsSplitArea]*AsSplit)
	for assplit := range stage.AsSplits {
		_ = assplit
		for _, _assplitarea := range assplit.AsSplitAreas {
			stage.AsSplit_AsSplitAreas_reverseMap[_assplitarea] = assplit
		}
	}

	// Compute reverse map for named struct View
	// insertion point per field
	stage.View_RootAsSplitAreas_reverseMap = make(map[*AsSplitArea]*View)
	for view := range stage.Views {
		_ = view
		for _, _assplitarea := range view.RootAsSplitAreas {
			stage.View_RootAsSplitAreas_reverseMap[_assplitarea] = view
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AsSplits)

	res = __gong__appendInstances(res, stage.AsSplitAreas)

	res = __gong__appendInstances(res, stage.Buttons)

	res = __gong__appendInstances(res, stage.Cursors)

	res = __gong__appendInstances(res, stage.FavIcons)

	res = __gong__appendInstances(res, stage.Forms)

	res = __gong__appendInstances(res, stage.Loads)

	res = __gong__appendInstances(res, stage.LogoOnTheLefts)

	res = __gong__appendInstances(res, stage.LogoOnTheRights)

	res = __gong__appendInstances(res, stage.Markdowns)

	res = __gong__appendInstances(res, stage.Sliders)

	res = __gong__appendInstances(res, stage.Splits)

	res = __gong__appendInstances(res, stage.Svgs)

	res = __gong__appendInstances(res, stage.Tables)

	res = __gong__appendInstances(res, stage.Threejss)

	res = __gong__appendInstances(res, stage.Titles)

	res = __gong__appendInstances(res, stage.Tones)

	res = __gong__appendInstances(res, stage.Trees)

	res = __gong__appendInstances(res, stage.Views)

	res = __gong__appendInstances(res, stage.Xlsxs)

	return
}

// insertion point per named struct
func (assplit *AsSplit) GongCopy() GongstructIF {
	newInstance := new(AsSplit)
	assplit.GongCopyBasicFields(newInstance)
	return newInstance
}

func (assplitarea *AsSplitArea) GongCopy() GongstructIF {
	newInstance := new(AsSplitArea)
	assplitarea.GongCopyBasicFields(newInstance)
	return newInstance
}

func (button *Button) GongCopy() GongstructIF {
	newInstance := new(Button)
	button.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cursor *Cursor) GongCopy() GongstructIF {
	newInstance := new(Cursor)
	cursor.GongCopyBasicFields(newInstance)
	return newInstance
}

func (favicon *FavIcon) GongCopy() GongstructIF {
	newInstance := new(FavIcon)
	favicon.GongCopyBasicFields(newInstance)
	return newInstance
}

func (form *Form) GongCopy() GongstructIF {
	newInstance := new(Form)
	form.GongCopyBasicFields(newInstance)
	return newInstance
}

func (load *Load) GongCopy() GongstructIF {
	newInstance := new(Load)
	load.GongCopyBasicFields(newInstance)
	return newInstance
}

func (logoontheleft *LogoOnTheLeft) GongCopy() GongstructIF {
	newInstance := new(LogoOnTheLeft)
	logoontheleft.GongCopyBasicFields(newInstance)
	return newInstance
}

func (logoontheright *LogoOnTheRight) GongCopy() GongstructIF {
	newInstance := new(LogoOnTheRight)
	logoontheright.GongCopyBasicFields(newInstance)
	return newInstance
}

func (markdown *Markdown) GongCopy() GongstructIF {
	newInstance := new(Markdown)
	markdown.GongCopyBasicFields(newInstance)
	return newInstance
}

func (slider *Slider) GongCopy() GongstructIF {
	newInstance := new(Slider)
	slider.GongCopyBasicFields(newInstance)
	return newInstance
}

func (split *Split) GongCopy() GongstructIF {
	newInstance := new(Split)
	split.GongCopyBasicFields(newInstance)
	return newInstance
}

func (svg *Svg) GongCopy() GongstructIF {
	newInstance := new(Svg)
	svg.GongCopyBasicFields(newInstance)
	return newInstance
}

func (table *Table) GongCopy() GongstructIF {
	newInstance := new(Table)
	table.GongCopyBasicFields(newInstance)
	return newInstance
}

func (threejs *Threejs) GongCopy() GongstructIF {
	newInstance := new(Threejs)
	threejs.GongCopyBasicFields(newInstance)
	return newInstance
}

func (title *Title) GongCopy() GongstructIF {
	newInstance := new(Title)
	title.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tone *Tone) GongCopy() GongstructIF {
	newInstance := new(Tone)
	tone.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tree *Tree) GongCopy() GongstructIF {
	newInstance := new(Tree)
	tree.GongCopyBasicFields(newInstance)
	return newInstance
}

func (view *View) GongCopy() GongstructIF {
	newInstance := new(View)
	view.GongCopyBasicFields(newInstance)
	return newInstance
}

func (xlsx *Xlsx) GongCopy() GongstructIF {
	newInstance := new(Xlsx)
	xlsx.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (assplit *AsSplit) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, assplit)
}

func (assplitarea *AsSplitArea) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, assplitarea)
}

func (button *Button) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, button)
}

func (cursor *Cursor) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cursor)
}

func (favicon *FavIcon) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, favicon)
}

func (form *Form) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, form)
}

func (load *Load) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, load)
}

func (logoontheleft *LogoOnTheLeft) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, logoontheleft)
}

func (logoontheright *LogoOnTheRight) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, logoontheright)
}

func (markdown *Markdown) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, markdown)
}

func (slider *Slider) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, slider)
}

func (split *Split) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, split)
}

func (svg *Svg) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, svg)
}

func (table *Table) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, table)
}

func (threejs *Threejs) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, threejs)
}

func (title *Title) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, title)
}

func (tone *Tone) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tone)
}

func (tree *Tree) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tree)
}

func (view *View) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, view)
}

func (xlsx *Xlsx) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, xlsx)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.AsSplits,
		stage.AsSplit_stagedOrder,
		stage.AsSplits_reference,
		&stage.AsSplits_referenceOrder,
		stage.AsSplits_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.AsSplitAreas,
		stage.AsSplitArea_stagedOrder,
		stage.AsSplitAreas_reference,
		&stage.AsSplitAreas_referenceOrder,
		stage.AsSplitAreas_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Buttons,
		stage.Button_stagedOrder,
		stage.Buttons_reference,
		&stage.Buttons_referenceOrder,
		stage.Buttons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Cursors,
		stage.Cursor_stagedOrder,
		stage.Cursors_reference,
		&stage.Cursors_referenceOrder,
		stage.Cursors_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.FavIcons,
		stage.FavIcon_stagedOrder,
		stage.FavIcons_reference,
		&stage.FavIcons_referenceOrder,
		stage.FavIcons_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Forms,
		stage.Form_stagedOrder,
		stage.Forms_reference,
		&stage.Forms_referenceOrder,
		stage.Forms_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Loads,
		stage.Load_stagedOrder,
		stage.Loads_reference,
		&stage.Loads_referenceOrder,
		stage.Loads_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LogoOnTheLefts,
		stage.LogoOnTheLeft_stagedOrder,
		stage.LogoOnTheLefts_reference,
		&stage.LogoOnTheLefts_referenceOrder,
		stage.LogoOnTheLefts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LogoOnTheRights,
		stage.LogoOnTheRight_stagedOrder,
		stage.LogoOnTheRights_reference,
		&stage.LogoOnTheRights_referenceOrder,
		stage.LogoOnTheRights_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Markdowns,
		stage.Markdown_stagedOrder,
		stage.Markdowns_reference,
		&stage.Markdowns_referenceOrder,
		stage.Markdowns_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Sliders,
		stage.Slider_stagedOrder,
		stage.Sliders_reference,
		&stage.Sliders_referenceOrder,
		stage.Sliders_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Splits,
		stage.Split_stagedOrder,
		stage.Splits_reference,
		&stage.Splits_referenceOrder,
		stage.Splits_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Svgs,
		stage.Svg_stagedOrder,
		stage.Svgs_reference,
		&stage.Svgs_referenceOrder,
		stage.Svgs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tables,
		stage.Table_stagedOrder,
		stage.Tables_reference,
		&stage.Tables_referenceOrder,
		stage.Tables_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Threejss,
		stage.Threejs_stagedOrder,
		stage.Threejss_reference,
		&stage.Threejss_referenceOrder,
		stage.Threejss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Titles,
		stage.Title_stagedOrder,
		stage.Titles_reference,
		&stage.Titles_referenceOrder,
		stage.Titles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tones,
		stage.Tone_stagedOrder,
		stage.Tones_reference,
		&stage.Tones_referenceOrder,
		stage.Tones_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Trees,
		stage.Tree_stagedOrder,
		stage.Trees_reference,
		&stage.Trees_referenceOrder,
		stage.Trees_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Views,
		stage.View_stagedOrder,
		stage.Views_reference,
		&stage.Views_referenceOrder,
		stage.Views_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Xlsxs,
		stage.Xlsx_stagedOrder,
		stage.Xlsxs_reference,
		&stage.Xlsxs_referenceOrder,
		stage.Xlsxs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.AsSplits, &stage.AsSplits_reference, &stage.AsSplits_referenceOrder, &stage.AsSplits_instance)

	__gong__computeReferencePass1(stage, stage.AsSplitAreas, &stage.AsSplitAreas_reference, &stage.AsSplitAreas_referenceOrder, &stage.AsSplitAreas_instance)

	__gong__computeReferencePass1(stage, stage.Buttons, &stage.Buttons_reference, &stage.Buttons_referenceOrder, &stage.Buttons_instance)

	__gong__computeReferencePass1(stage, stage.Cursors, &stage.Cursors_reference, &stage.Cursors_referenceOrder, &stage.Cursors_instance)

	__gong__computeReferencePass1(stage, stage.FavIcons, &stage.FavIcons_reference, &stage.FavIcons_referenceOrder, &stage.FavIcons_instance)

	__gong__computeReferencePass1(stage, stage.Forms, &stage.Forms_reference, &stage.Forms_referenceOrder, &stage.Forms_instance)

	__gong__computeReferencePass1(stage, stage.Loads, &stage.Loads_reference, &stage.Loads_referenceOrder, &stage.Loads_instance)

	__gong__computeReferencePass1(stage, stage.LogoOnTheLefts, &stage.LogoOnTheLefts_reference, &stage.LogoOnTheLefts_referenceOrder, &stage.LogoOnTheLefts_instance)

	__gong__computeReferencePass1(stage, stage.LogoOnTheRights, &stage.LogoOnTheRights_reference, &stage.LogoOnTheRights_referenceOrder, &stage.LogoOnTheRights_instance)

	__gong__computeReferencePass1(stage, stage.Markdowns, &stage.Markdowns_reference, &stage.Markdowns_referenceOrder, &stage.Markdowns_instance)

	__gong__computeReferencePass1(stage, stage.Sliders, &stage.Sliders_reference, &stage.Sliders_referenceOrder, &stage.Sliders_instance)

	__gong__computeReferencePass1(stage, stage.Splits, &stage.Splits_reference, &stage.Splits_referenceOrder, &stage.Splits_instance)

	__gong__computeReferencePass1(stage, stage.Svgs, &stage.Svgs_reference, &stage.Svgs_referenceOrder, &stage.Svgs_instance)

	__gong__computeReferencePass1(stage, stage.Tables, &stage.Tables_reference, &stage.Tables_referenceOrder, &stage.Tables_instance)

	__gong__computeReferencePass1(stage, stage.Threejss, &stage.Threejss_reference, &stage.Threejss_referenceOrder, &stage.Threejss_instance)

	__gong__computeReferencePass1(stage, stage.Titles, &stage.Titles_reference, &stage.Titles_referenceOrder, &stage.Titles_instance)

	__gong__computeReferencePass1(stage, stage.Tones, &stage.Tones_reference, &stage.Tones_referenceOrder, &stage.Tones_instance)

	__gong__computeReferencePass1(stage, stage.Trees, &stage.Trees_reference, &stage.Trees_referenceOrder, &stage.Trees_instance)

	__gong__computeReferencePass1(stage, stage.Views, &stage.Views_reference, &stage.Views_referenceOrder, &stage.Views_instance)

	__gong__computeReferencePass1(stage, stage.Xlsxs, &stage.Xlsxs_reference, &stage.Xlsxs_referenceOrder, &stage.Xlsxs_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AsSplits, stage.AsSplits_reference, stage)

	__gong__computeReferencePass2(stage.AsSplitAreas, stage.AsSplitAreas_reference, stage)

	__gong__computeReferencePass2(stage.Buttons, stage.Buttons_reference, stage)

	__gong__computeReferencePass2(stage.Cursors, stage.Cursors_reference, stage)

	__gong__computeReferencePass2(stage.FavIcons, stage.FavIcons_reference, stage)

	__gong__computeReferencePass2(stage.Forms, stage.Forms_reference, stage)

	__gong__computeReferencePass2(stage.Loads, stage.Loads_reference, stage)

	__gong__computeReferencePass2(stage.LogoOnTheLefts, stage.LogoOnTheLefts_reference, stage)

	__gong__computeReferencePass2(stage.LogoOnTheRights, stage.LogoOnTheRights_reference, stage)

	__gong__computeReferencePass2(stage.Markdowns, stage.Markdowns_reference, stage)

	__gong__computeReferencePass2(stage.Sliders, stage.Sliders_reference, stage)

	__gong__computeReferencePass2(stage.Splits, stage.Splits_reference, stage)

	__gong__computeReferencePass2(stage.Svgs, stage.Svgs_reference, stage)

	__gong__computeReferencePass2(stage.Tables, stage.Tables_reference, stage)

	__gong__computeReferencePass2(stage.Threejss, stage.Threejss_reference, stage)

	__gong__computeReferencePass2(stage.Titles, stage.Titles_reference, stage)

	__gong__computeReferencePass2(stage.Tones, stage.Tones_reference, stage)

	__gong__computeReferencePass2(stage.Trees, stage.Trees_reference, stage)

	__gong__computeReferencePass2(stage.Views, stage.Views_reference, stage)

	__gong__computeReferencePass2(stage.Xlsxs, stage.Xlsxs_reference, stage)

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
	return __gong__getOrder(stage.AsSplit_stagedOrder, stage.AsSplits_referenceOrder, assplit, "AsSplit")
}

func (assplitarea *AsSplitArea) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AsSplitArea_stagedOrder, stage.AsSplitAreas_referenceOrder, assplitarea, "AsSplitArea")
}

func (button *Button) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Button_stagedOrder, stage.Buttons_referenceOrder, button, "Button")
}

func (cursor *Cursor) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Cursor_stagedOrder, stage.Cursors_referenceOrder, cursor, "Cursor")
}

func (favicon *FavIcon) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.FavIcon_stagedOrder, stage.FavIcons_referenceOrder, favicon, "FavIcon")
}

func (form *Form) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Form_stagedOrder, stage.Forms_referenceOrder, form, "Form")
}

func (load *Load) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Load_stagedOrder, stage.Loads_referenceOrder, load, "Load")
}

func (logoontheleft *LogoOnTheLeft) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LogoOnTheLeft_stagedOrder, stage.LogoOnTheLefts_referenceOrder, logoontheleft, "LogoOnTheLeft")
}

func (logoontheright *LogoOnTheRight) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LogoOnTheRight_stagedOrder, stage.LogoOnTheRights_referenceOrder, logoontheright, "LogoOnTheRight")
}

func (markdown *Markdown) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Markdown_stagedOrder, stage.Markdowns_referenceOrder, markdown, "Markdown")
}

func (slider *Slider) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Slider_stagedOrder, stage.Sliders_referenceOrder, slider, "Slider")
}

func (split *Split) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Split_stagedOrder, stage.Splits_referenceOrder, split, "Split")
}

func (svg *Svg) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Svg_stagedOrder, stage.Svgs_referenceOrder, svg, "Svg")
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Table_stagedOrder, stage.Tables_referenceOrder, table, "Table")
}

func (threejs *Threejs) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Threejs_stagedOrder, stage.Threejss_referenceOrder, threejs, "Threejs")
}

func (title *Title) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Title_stagedOrder, stage.Titles_referenceOrder, title, "Title")
}

func (tone *Tone) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tone_stagedOrder, stage.Tones_referenceOrder, tone, "Tone")
}

func (tree *Tree) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tree_stagedOrder, stage.Trees_referenceOrder, tree, "Tree")
}

func (view *View) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.View_stagedOrder, stage.Views_referenceOrder, view, "View")
}

func (xlsx *Xlsx) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Xlsx_stagedOrder, stage.Xlsxs_referenceOrder, xlsx, "Xlsx")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (assplit *AsSplit) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(assplit, assplit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assplit *AsSplit) GongGetReferenceIdentifier(stage *Stage) string {
	return assplit.GongGetIdentifier(stage)
}

func (assplitarea *AsSplitArea) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(assplitarea, assplitarea.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assplitarea *AsSplitArea) GongGetReferenceIdentifier(stage *Stage) string {
	return assplitarea.GongGetIdentifier(stage)
}

func (button *Button) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(button, button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return button.GongGetIdentifier(stage)
}

func (cursor *Cursor) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cursor, cursor.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cursor *Cursor) GongGetReferenceIdentifier(stage *Stage) string {
	return cursor.GongGetIdentifier(stage)
}

func (favicon *FavIcon) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(favicon, favicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (favicon *FavIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return favicon.GongGetIdentifier(stage)
}

func (form *Form) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(form, form.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (form *Form) GongGetReferenceIdentifier(stage *Stage) string {
	return form.GongGetIdentifier(stage)
}

func (load *Load) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(load, load.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (load *Load) GongGetReferenceIdentifier(stage *Stage) string {
	return load.GongGetIdentifier(stage)
}

func (logoontheleft *LogoOnTheLeft) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(logoontheleft, logoontheleft.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (logoontheleft *LogoOnTheLeft) GongGetReferenceIdentifier(stage *Stage) string {
	return logoontheleft.GongGetIdentifier(stage)
}

func (logoontheright *LogoOnTheRight) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(logoontheright, logoontheright.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (logoontheright *LogoOnTheRight) GongGetReferenceIdentifier(stage *Stage) string {
	return logoontheright.GongGetIdentifier(stage)
}

func (markdown *Markdown) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(markdown, markdown.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (markdown *Markdown) GongGetReferenceIdentifier(stage *Stage) string {
	return markdown.GongGetIdentifier(stage)
}

func (slider *Slider) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(slider, slider.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slider *Slider) GongGetReferenceIdentifier(stage *Stage) string {
	return slider.GongGetIdentifier(stage)
}

func (split *Split) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(split, split.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (split *Split) GongGetReferenceIdentifier(stage *Stage) string {
	return split.GongGetIdentifier(stage)
}

func (svg *Svg) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(svg, svg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svg *Svg) GongGetReferenceIdentifier(stage *Stage) string {
	return svg.GongGetIdentifier(stage)
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(table, table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return table.GongGetIdentifier(stage)
}

func (threejs *Threejs) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(threejs, threejs.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (threejs *Threejs) GongGetReferenceIdentifier(stage *Stage) string {
	return threejs.GongGetIdentifier(stage)
}

func (title *Title) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(title, title.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (title *Title) GongGetReferenceIdentifier(stage *Stage) string {
	return title.GongGetIdentifier(stage)
}

func (tone *Tone) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tone, tone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tone *Tone) GongGetReferenceIdentifier(stage *Stage) string {
	return tone.GongGetIdentifier(stage)
}

func (tree *Tree) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tree, tree.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tree *Tree) GongGetReferenceIdentifier(stage *Stage) string {
	return tree.GongGetIdentifier(stage)
}

func (view *View) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(view, view.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (view *View) GongGetReferenceIdentifier(stage *Stage) string {
	return view.GongGetIdentifier(stage)
}

func (xlsx *Xlsx) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(xlsx, xlsx.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlsx *Xlsx) GongGetReferenceIdentifier(stage *Stage) string {
	return xlsx.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (assplit *AsSplit) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(assplit.GongGetIdentifier(stage), "AsSplit", assplit.Name)
}

func (assplitarea *AsSplitArea) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(assplitarea.GongGetIdentifier(stage), "AsSplitArea", assplitarea.Name)
}

func (button *Button) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(button.GongGetIdentifier(stage), "Button", button.Name)
}

func (cursor *Cursor) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cursor.GongGetIdentifier(stage), "Cursor", cursor.Name)
}

func (favicon *FavIcon) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(favicon.GongGetIdentifier(stage), "FavIcon", favicon.Name)
}

func (form *Form) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(form.GongGetIdentifier(stage), "Form", form.Name)
}

func (load *Load) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(load.GongGetIdentifier(stage), "Load", load.Name)
}

func (logoontheleft *LogoOnTheLeft) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(logoontheleft.GongGetIdentifier(stage), "LogoOnTheLeft", logoontheleft.Name)
}

func (logoontheright *LogoOnTheRight) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(logoontheright.GongGetIdentifier(stage), "LogoOnTheRight", logoontheright.Name)
}

func (markdown *Markdown) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(markdown.GongGetIdentifier(stage), "Markdown", markdown.Name)
}

func (slider *Slider) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(slider.GongGetIdentifier(stage), "Slider", slider.Name)
}

func (split *Split) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(split.GongGetIdentifier(stage), "Split", split.Name)
}

func (svg *Svg) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(svg.GongGetIdentifier(stage), "Svg", svg.Name)
}

func (table *Table) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(table.GongGetIdentifier(stage), "Table", table.Name)
}

func (threejs *Threejs) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(threejs.GongGetIdentifier(stage), "Threejs", threejs.Name)
}

func (title *Title) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(title.GongGetIdentifier(stage), "Title", title.Name)
}

func (tone *Tone) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tone.GongGetIdentifier(stage), "Tone", tone.Name)
}

func (tree *Tree) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tree.GongGetIdentifier(stage), "Tree", tree.Name)
}

func (view *View) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(view.GongGetIdentifier(stage), "View", view.Name)
}

func (xlsx *Xlsx) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(xlsx.GongGetIdentifier(stage), "Xlsx", xlsx.Name)
}

// insertion point for unstaging
func (assplit *AsSplit) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(assplit.GongGetReferenceIdentifier(stage))
}

func (assplitarea *AsSplitArea) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(assplitarea.GongGetReferenceIdentifier(stage))
}

func (button *Button) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(button.GongGetReferenceIdentifier(stage))
}

func (cursor *Cursor) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cursor.GongGetReferenceIdentifier(stage))
}

func (favicon *FavIcon) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(favicon.GongGetReferenceIdentifier(stage))
}

func (form *Form) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(form.GongGetReferenceIdentifier(stage))
}

func (load *Load) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(load.GongGetReferenceIdentifier(stage))
}

func (logoontheleft *LogoOnTheLeft) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(logoontheleft.GongGetReferenceIdentifier(stage))
}

func (logoontheright *LogoOnTheRight) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(logoontheright.GongGetReferenceIdentifier(stage))
}

func (markdown *Markdown) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(markdown.GongGetReferenceIdentifier(stage))
}

func (slider *Slider) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(slider.GongGetReferenceIdentifier(stage))
}

func (split *Split) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(split.GongGetReferenceIdentifier(stage))
}

func (svg *Svg) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(svg.GongGetReferenceIdentifier(stage))
}

func (table *Table) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(table.GongGetReferenceIdentifier(stage))
}

func (threejs *Threejs) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(threejs.GongGetReferenceIdentifier(stage))
}

func (title *Title) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(title.GongGetReferenceIdentifier(stage))
}

func (tone *Tone) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tone.GongGetReferenceIdentifier(stage))
}

func (tree *Tree) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tree.GongGetReferenceIdentifier(stage))
}

func (view *View) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(view.GongGetReferenceIdentifier(stage))
}

func (xlsx *Xlsx) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(xlsx.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
