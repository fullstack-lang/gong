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

	// Compute reverse map for named struct Threejs
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

	// end of insertion point per named struct
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

	for instance := range stage.Threejss {
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
func (assplit *AsSplit) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(assplit).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(assplit), uint64(stage.GetOrder(assplit)))
	return
}

func (assplitarea *AsSplitArea) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(assplitarea).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(assplitarea), uint64(stage.GetOrder(assplitarea)))
	return
}

func (button *Button) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(button).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(button), uint64(stage.GetOrder(button)))
	return
}

func (cursor *Cursor) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cursor).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cursor), uint64(stage.GetOrder(cursor)))
	return
}

func (favicon *FavIcon) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(favicon).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(favicon), uint64(stage.GetOrder(favicon)))
	return
}

func (form *Form) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(form).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(form), uint64(stage.GetOrder(form)))
	return
}

func (load *Load) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(load).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(load), uint64(stage.GetOrder(load)))
	return
}

func (logoontheleft *LogoOnTheLeft) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(logoontheleft).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(logoontheleft), uint64(stage.GetOrder(logoontheleft)))
	return
}

func (logoontheright *LogoOnTheRight) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(logoontheright).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(logoontheright), uint64(stage.GetOrder(logoontheright)))
	return
}

func (markdown *Markdown) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(markdown).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(markdown), uint64(stage.GetOrder(markdown)))
	return
}

func (slider *Slider) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(slider).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(slider), uint64(stage.GetOrder(slider)))
	return
}

func (split *Split) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(split).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(split), uint64(stage.GetOrder(split)))
	return
}

func (svg *Svg) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(svg).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(svg), uint64(stage.GetOrder(svg)))
	return
}

func (table *Table) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(table).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(table), uint64(stage.GetOrder(table)))
	return
}

func (threejs *Threejs) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(threejs).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(threejs), uint64(stage.GetOrder(threejs)))
	return
}

func (title *Title) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(title).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(title), uint64(stage.GetOrder(title)))
	return
}

func (tone *Tone) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tone).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tone), uint64(stage.GetOrder(tone)))
	return
}

func (tree *Tree) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tree).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tree), uint64(stage.GetOrder(tree)))
	return
}

func (view *View) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(view).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(view), uint64(stage.GetOrder(view)))
	return
}

func (xlsx *Xlsx) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(xlsx).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(xlsx), uint64(stage.GetOrder(xlsx)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
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
	stage.AsSplits_reference = make(map[*AsSplit]*AsSplit)
	stage.AsSplits_referenceOrder = make(map[*AsSplit]uint) // diff Unstage needs the reference order
	stage.AsSplits_instance = make(map[*AsSplit]*AsSplit)
	for instance := range stage.AsSplits {
		_copy := instance.GongCopy().(*AsSplit)
		stage.AsSplits_reference[instance] = _copy
		stage.AsSplits_instance[_copy] = instance
		stage.AsSplits_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.AsSplitAreas_reference = make(map[*AsSplitArea]*AsSplitArea)
	stage.AsSplitAreas_referenceOrder = make(map[*AsSplitArea]uint) // diff Unstage needs the reference order
	stage.AsSplitAreas_instance = make(map[*AsSplitArea]*AsSplitArea)
	for instance := range stage.AsSplitAreas {
		_copy := instance.GongCopy().(*AsSplitArea)
		stage.AsSplitAreas_reference[instance] = _copy
		stage.AsSplitAreas_instance[_copy] = instance
		stage.AsSplitAreas_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Buttons_reference = make(map[*Button]*Button)
	stage.Buttons_referenceOrder = make(map[*Button]uint) // diff Unstage needs the reference order
	stage.Buttons_instance = make(map[*Button]*Button)
	for instance := range stage.Buttons {
		_copy := instance.GongCopy().(*Button)
		stage.Buttons_reference[instance] = _copy
		stage.Buttons_instance[_copy] = instance
		stage.Buttons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Cursors_reference = make(map[*Cursor]*Cursor)
	stage.Cursors_referenceOrder = make(map[*Cursor]uint) // diff Unstage needs the reference order
	stage.Cursors_instance = make(map[*Cursor]*Cursor)
	for instance := range stage.Cursors {
		_copy := instance.GongCopy().(*Cursor)
		stage.Cursors_reference[instance] = _copy
		stage.Cursors_instance[_copy] = instance
		stage.Cursors_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.FavIcons_reference = make(map[*FavIcon]*FavIcon)
	stage.FavIcons_referenceOrder = make(map[*FavIcon]uint) // diff Unstage needs the reference order
	stage.FavIcons_instance = make(map[*FavIcon]*FavIcon)
	for instance := range stage.FavIcons {
		_copy := instance.GongCopy().(*FavIcon)
		stage.FavIcons_reference[instance] = _copy
		stage.FavIcons_instance[_copy] = instance
		stage.FavIcons_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Forms_reference = make(map[*Form]*Form)
	stage.Forms_referenceOrder = make(map[*Form]uint) // diff Unstage needs the reference order
	stage.Forms_instance = make(map[*Form]*Form)
	for instance := range stage.Forms {
		_copy := instance.GongCopy().(*Form)
		stage.Forms_reference[instance] = _copy
		stage.Forms_instance[_copy] = instance
		stage.Forms_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Loads_reference = make(map[*Load]*Load)
	stage.Loads_referenceOrder = make(map[*Load]uint) // diff Unstage needs the reference order
	stage.Loads_instance = make(map[*Load]*Load)
	for instance := range stage.Loads {
		_copy := instance.GongCopy().(*Load)
		stage.Loads_reference[instance] = _copy
		stage.Loads_instance[_copy] = instance
		stage.Loads_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.LogoOnTheLefts_reference = make(map[*LogoOnTheLeft]*LogoOnTheLeft)
	stage.LogoOnTheLefts_referenceOrder = make(map[*LogoOnTheLeft]uint) // diff Unstage needs the reference order
	stage.LogoOnTheLefts_instance = make(map[*LogoOnTheLeft]*LogoOnTheLeft)
	for instance := range stage.LogoOnTheLefts {
		_copy := instance.GongCopy().(*LogoOnTheLeft)
		stage.LogoOnTheLefts_reference[instance] = _copy
		stage.LogoOnTheLefts_instance[_copy] = instance
		stage.LogoOnTheLefts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.LogoOnTheRights_reference = make(map[*LogoOnTheRight]*LogoOnTheRight)
	stage.LogoOnTheRights_referenceOrder = make(map[*LogoOnTheRight]uint) // diff Unstage needs the reference order
	stage.LogoOnTheRights_instance = make(map[*LogoOnTheRight]*LogoOnTheRight)
	for instance := range stage.LogoOnTheRights {
		_copy := instance.GongCopy().(*LogoOnTheRight)
		stage.LogoOnTheRights_reference[instance] = _copy
		stage.LogoOnTheRights_instance[_copy] = instance
		stage.LogoOnTheRights_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Markdowns_reference = make(map[*Markdown]*Markdown)
	stage.Markdowns_referenceOrder = make(map[*Markdown]uint) // diff Unstage needs the reference order
	stage.Markdowns_instance = make(map[*Markdown]*Markdown)
	for instance := range stage.Markdowns {
		_copy := instance.GongCopy().(*Markdown)
		stage.Markdowns_reference[instance] = _copy
		stage.Markdowns_instance[_copy] = instance
		stage.Markdowns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Sliders_reference = make(map[*Slider]*Slider)
	stage.Sliders_referenceOrder = make(map[*Slider]uint) // diff Unstage needs the reference order
	stage.Sliders_instance = make(map[*Slider]*Slider)
	for instance := range stage.Sliders {
		_copy := instance.GongCopy().(*Slider)
		stage.Sliders_reference[instance] = _copy
		stage.Sliders_instance[_copy] = instance
		stage.Sliders_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Splits_reference = make(map[*Split]*Split)
	stage.Splits_referenceOrder = make(map[*Split]uint) // diff Unstage needs the reference order
	stage.Splits_instance = make(map[*Split]*Split)
	for instance := range stage.Splits {
		_copy := instance.GongCopy().(*Split)
		stage.Splits_reference[instance] = _copy
		stage.Splits_instance[_copy] = instance
		stage.Splits_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Svgs_reference = make(map[*Svg]*Svg)
	stage.Svgs_referenceOrder = make(map[*Svg]uint) // diff Unstage needs the reference order
	stage.Svgs_instance = make(map[*Svg]*Svg)
	for instance := range stage.Svgs {
		_copy := instance.GongCopy().(*Svg)
		stage.Svgs_reference[instance] = _copy
		stage.Svgs_instance[_copy] = instance
		stage.Svgs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tables_reference = make(map[*Table]*Table)
	stage.Tables_referenceOrder = make(map[*Table]uint) // diff Unstage needs the reference order
	stage.Tables_instance = make(map[*Table]*Table)
	for instance := range stage.Tables {
		_copy := instance.GongCopy().(*Table)
		stage.Tables_reference[instance] = _copy
		stage.Tables_instance[_copy] = instance
		stage.Tables_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Threejss_reference = make(map[*Threejs]*Threejs)
	stage.Threejss_referenceOrder = make(map[*Threejs]uint) // diff Unstage needs the reference order
	stage.Threejss_instance = make(map[*Threejs]*Threejs)
	for instance := range stage.Threejss {
		_copy := instance.GongCopy().(*Threejs)
		stage.Threejss_reference[instance] = _copy
		stage.Threejss_instance[_copy] = instance
		stage.Threejss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Titles_reference = make(map[*Title]*Title)
	stage.Titles_referenceOrder = make(map[*Title]uint) // diff Unstage needs the reference order
	stage.Titles_instance = make(map[*Title]*Title)
	for instance := range stage.Titles {
		_copy := instance.GongCopy().(*Title)
		stage.Titles_reference[instance] = _copy
		stage.Titles_instance[_copy] = instance
		stage.Titles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tones_reference = make(map[*Tone]*Tone)
	stage.Tones_referenceOrder = make(map[*Tone]uint) // diff Unstage needs the reference order
	stage.Tones_instance = make(map[*Tone]*Tone)
	for instance := range stage.Tones {
		_copy := instance.GongCopy().(*Tone)
		stage.Tones_reference[instance] = _copy
		stage.Tones_instance[_copy] = instance
		stage.Tones_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Trees_reference = make(map[*Tree]*Tree)
	stage.Trees_referenceOrder = make(map[*Tree]uint) // diff Unstage needs the reference order
	stage.Trees_instance = make(map[*Tree]*Tree)
	for instance := range stage.Trees {
		_copy := instance.GongCopy().(*Tree)
		stage.Trees_reference[instance] = _copy
		stage.Trees_instance[_copy] = instance
		stage.Trees_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Views_reference = make(map[*View]*View)
	stage.Views_referenceOrder = make(map[*View]uint) // diff Unstage needs the reference order
	stage.Views_instance = make(map[*View]*View)
	for instance := range stage.Views {
		_copy := instance.GongCopy().(*View)
		stage.Views_reference[instance] = _copy
		stage.Views_instance[_copy] = instance
		stage.Views_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Xlsxs_reference = make(map[*Xlsx]*Xlsx)
	stage.Xlsxs_referenceOrder = make(map[*Xlsx]uint) // diff Unstage needs the reference order
	stage.Xlsxs_instance = make(map[*Xlsx]*Xlsx)
	for instance := range stage.Xlsxs {
		_copy := instance.GongCopy().(*Xlsx)
		stage.Xlsxs_reference[instance] = _copy
		stage.Xlsxs_instance[_copy] = instance
		stage.Xlsxs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.AsSplits {
		reference := stage.AsSplits_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AsSplitAreas {
		reference := stage.AsSplitAreas_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Buttons {
		reference := stage.Buttons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Cursors {
		reference := stage.Cursors_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.FavIcons {
		reference := stage.FavIcons_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Forms {
		reference := stage.Forms_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Loads {
		reference := stage.Loads_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.LogoOnTheLefts {
		reference := stage.LogoOnTheLefts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.LogoOnTheRights {
		reference := stage.LogoOnTheRights_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Markdowns {
		reference := stage.Markdowns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Sliders {
		reference := stage.Sliders_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Splits {
		reference := stage.Splits_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Svgs {
		reference := stage.Svgs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tables {
		reference := stage.Tables_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Threejss {
		reference := stage.Threejss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Titles {
		reference := stage.Titles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tones {
		reference := stage.Tones_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Trees {
		reference := stage.Trees_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Views {
		reference := stage.Views_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Xlsxs {
		reference := stage.Xlsxs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
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
	if order, ok := stage.AsSplit_stagedOrder[assplit]; ok {
		return order
	}
	if order, ok := stage.AsSplits_referenceOrder[assplit]; ok {
		return order
	} else {
		log.Printf("instance %p of type AsSplit was not staged and does not have a reference order", assplit)
		return 0
	}
}

func (assplitarea *AsSplitArea) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AsSplitArea_stagedOrder[assplitarea]; ok {
		return order
	}
	if order, ok := stage.AsSplitAreas_referenceOrder[assplitarea]; ok {
		return order
	} else {
		log.Printf("instance %p of type AsSplitArea was not staged and does not have a reference order", assplitarea)
		return 0
	}
}

func (button *Button) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Button_stagedOrder[button]; ok {
		return order
	}
	if order, ok := stage.Buttons_referenceOrder[button]; ok {
		return order
	} else {
		log.Printf("instance %p of type Button was not staged and does not have a reference order", button)
		return 0
	}
}

func (cursor *Cursor) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Cursor_stagedOrder[cursor]; ok {
		return order
	}
	if order, ok := stage.Cursors_referenceOrder[cursor]; ok {
		return order
	} else {
		log.Printf("instance %p of type Cursor was not staged and does not have a reference order", cursor)
		return 0
	}
}

func (favicon *FavIcon) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.FavIcon_stagedOrder[favicon]; ok {
		return order
	}
	if order, ok := stage.FavIcons_referenceOrder[favicon]; ok {
		return order
	} else {
		log.Printf("instance %p of type FavIcon was not staged and does not have a reference order", favicon)
		return 0
	}
}

func (form *Form) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Form_stagedOrder[form]; ok {
		return order
	}
	if order, ok := stage.Forms_referenceOrder[form]; ok {
		return order
	} else {
		log.Printf("instance %p of type Form was not staged and does not have a reference order", form)
		return 0
	}
}

func (load *Load) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Load_stagedOrder[load]; ok {
		return order
	}
	if order, ok := stage.Loads_referenceOrder[load]; ok {
		return order
	} else {
		log.Printf("instance %p of type Load was not staged and does not have a reference order", load)
		return 0
	}
}

func (logoontheleft *LogoOnTheLeft) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LogoOnTheLeft_stagedOrder[logoontheleft]; ok {
		return order
	}
	if order, ok := stage.LogoOnTheLefts_referenceOrder[logoontheleft]; ok {
		return order
	} else {
		log.Printf("instance %p of type LogoOnTheLeft was not staged and does not have a reference order", logoontheleft)
		return 0
	}
}

func (logoontheright *LogoOnTheRight) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LogoOnTheRight_stagedOrder[logoontheright]; ok {
		return order
	}
	if order, ok := stage.LogoOnTheRights_referenceOrder[logoontheright]; ok {
		return order
	} else {
		log.Printf("instance %p of type LogoOnTheRight was not staged and does not have a reference order", logoontheright)
		return 0
	}
}

func (markdown *Markdown) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Markdown_stagedOrder[markdown]; ok {
		return order
	}
	if order, ok := stage.Markdowns_referenceOrder[markdown]; ok {
		return order
	} else {
		log.Printf("instance %p of type Markdown was not staged and does not have a reference order", markdown)
		return 0
	}
}

func (slider *Slider) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Slider_stagedOrder[slider]; ok {
		return order
	}
	if order, ok := stage.Sliders_referenceOrder[slider]; ok {
		return order
	} else {
		log.Printf("instance %p of type Slider was not staged and does not have a reference order", slider)
		return 0
	}
}

func (split *Split) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Split_stagedOrder[split]; ok {
		return order
	}
	if order, ok := stage.Splits_referenceOrder[split]; ok {
		return order
	} else {
		log.Printf("instance %p of type Split was not staged and does not have a reference order", split)
		return 0
	}
}

func (svg *Svg) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Svg_stagedOrder[svg]; ok {
		return order
	}
	if order, ok := stage.Svgs_referenceOrder[svg]; ok {
		return order
	} else {
		log.Printf("instance %p of type Svg was not staged and does not have a reference order", svg)
		return 0
	}
}

func (table *Table) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Table_stagedOrder[table]; ok {
		return order
	}
	if order, ok := stage.Tables_referenceOrder[table]; ok {
		return order
	} else {
		log.Printf("instance %p of type Table was not staged and does not have a reference order", table)
		return 0
	}
}

func (threejs *Threejs) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Threejs_stagedOrder[threejs]; ok {
		return order
	}
	if order, ok := stage.Threejss_referenceOrder[threejs]; ok {
		return order
	} else {
		log.Printf("instance %p of type Threejs was not staged and does not have a reference order", threejs)
		return 0
	}
}

func (title *Title) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Title_stagedOrder[title]; ok {
		return order
	}
	if order, ok := stage.Titles_referenceOrder[title]; ok {
		return order
	} else {
		log.Printf("instance %p of type Title was not staged and does not have a reference order", title)
		return 0
	}
}

func (tone *Tone) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tone_stagedOrder[tone]; ok {
		return order
	}
	if order, ok := stage.Tones_referenceOrder[tone]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tone was not staged and does not have a reference order", tone)
		return 0
	}
}

func (tree *Tree) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tree_stagedOrder[tree]; ok {
		return order
	}
	if order, ok := stage.Trees_referenceOrder[tree]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tree was not staged and does not have a reference order", tree)
		return 0
	}
}

func (view *View) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.View_stagedOrder[view]; ok {
		return order
	}
	if order, ok := stage.Views_referenceOrder[view]; ok {
		return order
	} else {
		log.Printf("instance %p of type View was not staged and does not have a reference order", view)
		return 0
	}
}

func (xlsx *Xlsx) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Xlsx_stagedOrder[xlsx]; ok {
		return order
	}
	if order, ok := stage.Xlsxs_referenceOrder[xlsx]; ok {
		return order
	} else {
		log.Printf("instance %p of type Xlsx was not staged and does not have a reference order", xlsx)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", assplit.GongGetGongstructName(), assplit.GongGetOrder(stage))
}

func (assplitarea *AsSplitArea) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplitarea.GongGetGongstructName(), assplitarea.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assplitarea *AsSplitArea) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assplitarea.GongGetGongstructName(), assplitarea.GongGetOrder(stage))
}

func (button *Button) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (button *Button) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", button.GongGetGongstructName(), button.GongGetOrder(stage))
}

func (cursor *Cursor) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cursor.GongGetGongstructName(), cursor.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cursor *Cursor) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cursor.GongGetGongstructName(), cursor.GongGetOrder(stage))
}

func (favicon *FavIcon) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", favicon.GongGetGongstructName(), favicon.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (favicon *FavIcon) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", favicon.GongGetGongstructName(), favicon.GongGetOrder(stage))
}

func (form *Form) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", form.GongGetGongstructName(), form.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (form *Form) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", form.GongGetGongstructName(), form.GongGetOrder(stage))
}

func (load *Load) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", load.GongGetGongstructName(), load.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (load *Load) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", load.GongGetGongstructName(), load.GongGetOrder(stage))
}

func (logoontheleft *LogoOnTheLeft) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheleft.GongGetGongstructName(), logoontheleft.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (logoontheleft *LogoOnTheLeft) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheleft.GongGetGongstructName(), logoontheleft.GongGetOrder(stage))
}

func (logoontheright *LogoOnTheRight) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheright.GongGetGongstructName(), logoontheright.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (logoontheright *LogoOnTheRight) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", logoontheright.GongGetGongstructName(), logoontheright.GongGetOrder(stage))
}

func (markdown *Markdown) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", markdown.GongGetGongstructName(), markdown.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (markdown *Markdown) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", markdown.GongGetGongstructName(), markdown.GongGetOrder(stage))
}

func (slider *Slider) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slider *Slider) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slider.GongGetGongstructName(), slider.GongGetOrder(stage))
}

func (split *Split) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", split.GongGetGongstructName(), split.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (split *Split) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", split.GongGetGongstructName(), split.GongGetOrder(stage))
}

func (svg *Svg) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (svg *Svg) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", svg.GongGetGongstructName(), svg.GongGetOrder(stage))
}

func (table *Table) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (table *Table) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", table.GongGetGongstructName(), table.GongGetOrder(stage))
}

func (threejs *Threejs) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", threejs.GongGetGongstructName(), threejs.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (threejs *Threejs) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", threejs.GongGetGongstructName(), threejs.GongGetOrder(stage))
}

func (title *Title) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", title.GongGetGongstructName(), title.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (title *Title) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", title.GongGetGongstructName(), title.GongGetOrder(stage))
}

func (tone *Tone) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tone.GongGetGongstructName(), tone.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tone *Tone) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tone.GongGetGongstructName(), tone.GongGetOrder(stage))
}

func (tree *Tree) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tree *Tree) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tree.GongGetGongstructName(), tree.GongGetOrder(stage))
}

func (view *View) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", view.GongGetGongstructName(), view.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (view *View) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", view.GongGetGongstructName(), view.GongGetOrder(stage))
}

func (xlsx *Xlsx) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsx.GongGetGongstructName(), xlsx.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xlsx *Xlsx) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xlsx.GongGetGongstructName(), xlsx.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (assplit *AsSplit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assplit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(assplit.Name))
	return
}

func (assplitarea *AsSplitArea) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assplitarea.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AsSplitArea")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(assplitarea.Name))
	return
}

func (button *Button) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", button.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Button")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(button.Name))
	return
}

func (cursor *Cursor) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cursor.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cursor")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cursor.Name))
	return
}

func (favicon *FavIcon) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", favicon.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "FavIcon")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(favicon.Name))
	return
}

func (form *Form) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", form.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Form")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(form.Name))
	return
}

func (load *Load) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", load.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Load")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(load.Name))
	return
}

func (logoontheleft *LogoOnTheLeft) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", logoontheleft.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LogoOnTheLeft")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(logoontheleft.Name))
	return
}

func (logoontheright *LogoOnTheRight) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", logoontheright.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LogoOnTheRight")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(logoontheright.Name))
	return
}

func (markdown *Markdown) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", markdown.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Markdown")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(markdown.Name))
	return
}

func (slider *Slider) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slider.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slider")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(slider.Name))
	return
}

func (split *Split) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", split.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Split")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(split.Name))
	return
}

func (svg *Svg) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", svg.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Svg")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(svg.Name))
	return
}

func (table *Table) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", table.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Table")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(table.Name))
	return
}

func (threejs *Threejs) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", threejs.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Threejs")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(threejs.Name))
	return
}

func (title *Title) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", title.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Title")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(title.Name))
	return
}

func (tone *Tone) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tone.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tone")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tone.Name))
	return
}

func (tree *Tree) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tree.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tree")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tree.Name))
	return
}

func (view *View) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", view.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "View")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(view.Name))
	return
}

func (xlsx *Xlsx) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xlsx.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Xlsx")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(xlsx.Name))
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

func (threejs *Threejs) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", threejs.GongGetReferenceIdentifier(stage))
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

// end of template
