// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const ProbeTreeSidebarSuffix = "-sidebar"
const ProbeTableSuffix = "-table"
const ProbeFormSuffix = "-form"
const ProbeSplitSuffix = "-probe"

func (stage *StageStruct) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *StageStruct) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *StageStruct) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *StageStruct) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
}

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	name string

	// insertion point for definition of arrays registering instances
	AsSplits           map[*AsSplit]any
	AsSplits_mapString map[string]*AsSplit

	// insertion point for slice of pointers maps
	AsSplit_AsSplitAreas_reverseMap map[*AsSplitArea]*AsSplit

	OnAfterAsSplitCreateCallback OnAfterCreateInterface[AsSplit]
	OnAfterAsSplitUpdateCallback OnAfterUpdateInterface[AsSplit]
	OnAfterAsSplitDeleteCallback OnAfterDeleteInterface[AsSplit]
	OnAfterAsSplitReadCallback   OnAfterReadInterface[AsSplit]

	AsSplitAreas           map[*AsSplitArea]any
	AsSplitAreas_mapString map[string]*AsSplitArea

	// insertion point for slice of pointers maps
	OnAfterAsSplitAreaCreateCallback OnAfterCreateInterface[AsSplitArea]
	OnAfterAsSplitAreaUpdateCallback OnAfterUpdateInterface[AsSplitArea]
	OnAfterAsSplitAreaDeleteCallback OnAfterDeleteInterface[AsSplitArea]
	OnAfterAsSplitAreaReadCallback   OnAfterReadInterface[AsSplitArea]

	Buttons           map[*Button]any
	Buttons_mapString map[string]*Button

	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback OnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback OnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback OnAfterDeleteInterface[Button]
	OnAfterButtonReadCallback   OnAfterReadInterface[Button]

	Cursors           map[*Cursor]any
	Cursors_mapString map[string]*Cursor

	// insertion point for slice of pointers maps
	OnAfterCursorCreateCallback OnAfterCreateInterface[Cursor]
	OnAfterCursorUpdateCallback OnAfterUpdateInterface[Cursor]
	OnAfterCursorDeleteCallback OnAfterDeleteInterface[Cursor]
	OnAfterCursorReadCallback   OnAfterReadInterface[Cursor]

	Docs           map[*Doc]any
	Docs_mapString map[string]*Doc

	// insertion point for slice of pointers maps
	OnAfterDocCreateCallback OnAfterCreateInterface[Doc]
	OnAfterDocUpdateCallback OnAfterUpdateInterface[Doc]
	OnAfterDocDeleteCallback OnAfterDeleteInterface[Doc]
	OnAfterDocReadCallback   OnAfterReadInterface[Doc]

	Forms           map[*Form]any
	Forms_mapString map[string]*Form

	// insertion point for slice of pointers maps
	OnAfterFormCreateCallback OnAfterCreateInterface[Form]
	OnAfterFormUpdateCallback OnAfterUpdateInterface[Form]
	OnAfterFormDeleteCallback OnAfterDeleteInterface[Form]
	OnAfterFormReadCallback   OnAfterReadInterface[Form]

	Loads           map[*Load]any
	Loads_mapString map[string]*Load

	// insertion point for slice of pointers maps
	OnAfterLoadCreateCallback OnAfterCreateInterface[Load]
	OnAfterLoadUpdateCallback OnAfterUpdateInterface[Load]
	OnAfterLoadDeleteCallback OnAfterDeleteInterface[Load]
	OnAfterLoadReadCallback   OnAfterReadInterface[Load]

	Sliders           map[*Slider]any
	Sliders_mapString map[string]*Slider

	// insertion point for slice of pointers maps
	OnAfterSliderCreateCallback OnAfterCreateInterface[Slider]
	OnAfterSliderUpdateCallback OnAfterUpdateInterface[Slider]
	OnAfterSliderDeleteCallback OnAfterDeleteInterface[Slider]
	OnAfterSliderReadCallback   OnAfterReadInterface[Slider]

	Splits           map[*Split]any
	Splits_mapString map[string]*Split

	// insertion point for slice of pointers maps
	OnAfterSplitCreateCallback OnAfterCreateInterface[Split]
	OnAfterSplitUpdateCallback OnAfterUpdateInterface[Split]
	OnAfterSplitDeleteCallback OnAfterDeleteInterface[Split]
	OnAfterSplitReadCallback   OnAfterReadInterface[Split]

	Svgs           map[*Svg]any
	Svgs_mapString map[string]*Svg

	// insertion point for slice of pointers maps
	OnAfterSvgCreateCallback OnAfterCreateInterface[Svg]
	OnAfterSvgUpdateCallback OnAfterUpdateInterface[Svg]
	OnAfterSvgDeleteCallback OnAfterDeleteInterface[Svg]
	OnAfterSvgReadCallback   OnAfterReadInterface[Svg]

	Tables           map[*Table]any
	Tables_mapString map[string]*Table

	// insertion point for slice of pointers maps
	OnAfterTableCreateCallback OnAfterCreateInterface[Table]
	OnAfterTableUpdateCallback OnAfterUpdateInterface[Table]
	OnAfterTableDeleteCallback OnAfterDeleteInterface[Table]
	OnAfterTableReadCallback   OnAfterReadInterface[Table]

	Tones           map[*Tone]any
	Tones_mapString map[string]*Tone

	// insertion point for slice of pointers maps
	OnAfterToneCreateCallback OnAfterCreateInterface[Tone]
	OnAfterToneUpdateCallback OnAfterUpdateInterface[Tone]
	OnAfterToneDeleteCallback OnAfterDeleteInterface[Tone]
	OnAfterToneReadCallback   OnAfterReadInterface[Tone]

	Trees           map[*Tree]any
	Trees_mapString map[string]*Tree

	// insertion point for slice of pointers maps
	OnAfterTreeCreateCallback OnAfterCreateInterface[Tree]
	OnAfterTreeUpdateCallback OnAfterUpdateInterface[Tree]
	OnAfterTreeDeleteCallback OnAfterDeleteInterface[Tree]
	OnAfterTreeReadCallback   OnAfterReadInterface[Tree]

	Views           map[*View]any
	Views_mapString map[string]*View

	// insertion point for slice of pointers maps
	View_RootAsSplitAreas_reverseMap map[*AsSplitArea]*View

	OnAfterViewCreateCallback OnAfterCreateInterface[View]
	OnAfterViewUpdateCallback OnAfterUpdateInterface[View]
	OnAfterViewDeleteCallback OnAfterDeleteInterface[View]
	OnAfterViewReadCallback   OnAfterReadInterface[View]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	Order            uint
	Map_Staged_Order map[any]uint
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gong/lib/split/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitAsSplit(assplit *AsSplit)
	CheckoutAsSplit(assplit *AsSplit)
	CommitAsSplitArea(assplitarea *AsSplitArea)
	CheckoutAsSplitArea(assplitarea *AsSplitArea)
	CommitButton(button *Button)
	CheckoutButton(button *Button)
	CommitCursor(cursor *Cursor)
	CheckoutCursor(cursor *Cursor)
	CommitDoc(doc *Doc)
	CheckoutDoc(doc *Doc)
	CommitForm(form *Form)
	CheckoutForm(form *Form)
	CommitLoad(load *Load)
	CheckoutLoad(load *Load)
	CommitSlider(slider *Slider)
	CheckoutSlider(slider *Slider)
	CommitSplit(split *Split)
	CheckoutSplit(split *Split)
	CommitSvg(svg *Svg)
	CheckoutSvg(svg *Svg)
	CommitTable(table *Table)
	CheckoutTable(table *Table)
	CommitTone(tone *Tone)
	CheckoutTone(tone *Tone)
	CommitTree(tree *Tree)
	CheckoutTree(tree *Tree)
	CommitView(view *View)
	CheckoutView(view *View)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		AsSplits:           make(map[*AsSplit]any),
		AsSplits_mapString: make(map[string]*AsSplit),

		AsSplitAreas:           make(map[*AsSplitArea]any),
		AsSplitAreas_mapString: make(map[string]*AsSplitArea),

		Buttons:           make(map[*Button]any),
		Buttons_mapString: make(map[string]*Button),

		Cursors:           make(map[*Cursor]any),
		Cursors_mapString: make(map[string]*Cursor),

		Docs:           make(map[*Doc]any),
		Docs_mapString: make(map[string]*Doc),

		Forms:           make(map[*Form]any),
		Forms_mapString: make(map[string]*Form),

		Loads:           make(map[*Load]any),
		Loads_mapString: make(map[string]*Load),

		Sliders:           make(map[*Slider]any),
		Sliders_mapString: make(map[string]*Slider),

		Splits:           make(map[*Split]any),
		Splits_mapString: make(map[string]*Split),

		Svgs:           make(map[*Svg]any),
		Svgs_mapString: make(map[string]*Svg),

		Tables:           make(map[*Table]any),
		Tables_mapString: make(map[string]*Table),

		Tones:           make(map[*Tone]any),
		Tones_mapString: make(map[string]*Tone),

		Trees:           make(map[*Tree]any),
		Trees_mapString: make(map[string]*Tree),

		Views:           make(map[*View]any),
		Views_mapString: make(map[string]*View),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		Map_Staged_Order: make(map[any]uint),
	}

	return
}

func (stage *StageStruct) GetName() string {
	return stage.name
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AsSplit"] = len(stage.AsSplits)
	stage.Map_GongStructName_InstancesNb["AsSplitArea"] = len(stage.AsSplitAreas)
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Cursor"] = len(stage.Cursors)
	stage.Map_GongStructName_InstancesNb["Doc"] = len(stage.Docs)
	stage.Map_GongStructName_InstancesNb["Form"] = len(stage.Forms)
	stage.Map_GongStructName_InstancesNb["Load"] = len(stage.Loads)
	stage.Map_GongStructName_InstancesNb["Slider"] = len(stage.Sliders)
	stage.Map_GongStructName_InstancesNb["Split"] = len(stage.Splits)
	stage.Map_GongStructName_InstancesNb["Svg"] = len(stage.Svgs)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)
	stage.Map_GongStructName_InstancesNb["Tone"] = len(stage.Tones)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)
	stage.Map_GongStructName_InstancesNb["View"] = len(stage.Views)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["AsSplit"] = len(stage.AsSplits)
	stage.Map_GongStructName_InstancesNb["AsSplitArea"] = len(stage.AsSplitAreas)
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Cursor"] = len(stage.Cursors)
	stage.Map_GongStructName_InstancesNb["Doc"] = len(stage.Docs)
	stage.Map_GongStructName_InstancesNb["Form"] = len(stage.Forms)
	stage.Map_GongStructName_InstancesNb["Load"] = len(stage.Loads)
	stage.Map_GongStructName_InstancesNb["Slider"] = len(stage.Sliders)
	stage.Map_GongStructName_InstancesNb["Split"] = len(stage.Splits)
	stage.Map_GongStructName_InstancesNb["Svg"] = len(stage.Svgs)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)
	stage.Map_GongStructName_InstancesNb["Tone"] = len(stage.Tones)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)
	stage.Map_GongStructName_InstancesNb["View"] = len(stage.Views)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts assplit to the model stage
func (assplit *AsSplit) Stage(stage *StageStruct) *AsSplit {

	if _, ok := stage.AsSplits[assplit]; !ok {
		stage.AsSplits[assplit] = __member
		stage.Map_Staged_Order[assplit] = stage.Order
		stage.Order++
	}
	stage.AsSplits_mapString[assplit.Name] = assplit

	return assplit
}

// Unstage removes assplit off the model stage
func (assplit *AsSplit) Unstage(stage *StageStruct) *AsSplit {
	delete(stage.AsSplits, assplit)
	delete(stage.AsSplits_mapString, assplit.Name)
	return assplit
}

// UnstageVoid removes assplit off the model stage
func (assplit *AsSplit) UnstageVoid(stage *StageStruct) {
	delete(stage.AsSplits, assplit)
	delete(stage.AsSplits_mapString, assplit.Name)
}

// commit assplit to the back repo (if it is already staged)
func (assplit *AsSplit) Commit(stage *StageStruct) *AsSplit {
	if _, ok := stage.AsSplits[assplit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAsSplit(assplit)
		}
	}
	return assplit
}

func (assplit *AsSplit) CommitVoid(stage *StageStruct) {
	assplit.Commit(stage)
}

// Checkout assplit to the back repo (if it is already staged)
func (assplit *AsSplit) Checkout(stage *StageStruct) *AsSplit {
	if _, ok := stage.AsSplits[assplit]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAsSplit(assplit)
		}
	}
	return assplit
}

// for satisfaction of GongStruct interface
func (assplit *AsSplit) GetName() (res string) {
	return assplit.Name
}

// Stage puts assplitarea to the model stage
func (assplitarea *AsSplitArea) Stage(stage *StageStruct) *AsSplitArea {

	if _, ok := stage.AsSplitAreas[assplitarea]; !ok {
		stage.AsSplitAreas[assplitarea] = __member
		stage.Map_Staged_Order[assplitarea] = stage.Order
		stage.Order++
	}
	stage.AsSplitAreas_mapString[assplitarea.Name] = assplitarea

	return assplitarea
}

// Unstage removes assplitarea off the model stage
func (assplitarea *AsSplitArea) Unstage(stage *StageStruct) *AsSplitArea {
	delete(stage.AsSplitAreas, assplitarea)
	delete(stage.AsSplitAreas_mapString, assplitarea.Name)
	return assplitarea
}

// UnstageVoid removes assplitarea off the model stage
func (assplitarea *AsSplitArea) UnstageVoid(stage *StageStruct) {
	delete(stage.AsSplitAreas, assplitarea)
	delete(stage.AsSplitAreas_mapString, assplitarea.Name)
}

// commit assplitarea to the back repo (if it is already staged)
func (assplitarea *AsSplitArea) Commit(stage *StageStruct) *AsSplitArea {
	if _, ok := stage.AsSplitAreas[assplitarea]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitAsSplitArea(assplitarea)
		}
	}
	return assplitarea
}

func (assplitarea *AsSplitArea) CommitVoid(stage *StageStruct) {
	assplitarea.Commit(stage)
}

// Checkout assplitarea to the back repo (if it is already staged)
func (assplitarea *AsSplitArea) Checkout(stage *StageStruct) *AsSplitArea {
	if _, ok := stage.AsSplitAreas[assplitarea]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutAsSplitArea(assplitarea)
		}
	}
	return assplitarea
}

// for satisfaction of GongStruct interface
func (assplitarea *AsSplitArea) GetName() (res string) {
	return assplitarea.Name
}

// Stage puts button to the model stage
func (button *Button) Stage(stage *StageStruct) *Button {

	if _, ok := stage.Buttons[button]; !ok {
		stage.Buttons[button] = __member
		stage.Map_Staged_Order[button] = stage.Order
		stage.Order++
	}
	stage.Buttons_mapString[button.Name] = button

	return button
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *StageStruct) *Button {
	delete(stage.Buttons, button)
	delete(stage.Buttons_mapString, button.Name)
	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *StageStruct) {
	delete(stage.Buttons, button)
	delete(stage.Buttons_mapString, button.Name)
}

// commit button to the back repo (if it is already staged)
func (button *Button) Commit(stage *StageStruct) *Button {
	if _, ok := stage.Buttons[button]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitButton(button)
		}
	}
	return button
}

func (button *Button) CommitVoid(stage *StageStruct) {
	button.Commit(stage)
}

// Checkout button to the back repo (if it is already staged)
func (button *Button) Checkout(stage *StageStruct) *Button {
	if _, ok := stage.Buttons[button]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutButton(button)
		}
	}
	return button
}

// for satisfaction of GongStruct interface
func (button *Button) GetName() (res string) {
	return button.Name
}

// Stage puts cursor to the model stage
func (cursor *Cursor) Stage(stage *StageStruct) *Cursor {

	if _, ok := stage.Cursors[cursor]; !ok {
		stage.Cursors[cursor] = __member
		stage.Map_Staged_Order[cursor] = stage.Order
		stage.Order++
	}
	stage.Cursors_mapString[cursor.Name] = cursor

	return cursor
}

// Unstage removes cursor off the model stage
func (cursor *Cursor) Unstage(stage *StageStruct) *Cursor {
	delete(stage.Cursors, cursor)
	delete(stage.Cursors_mapString, cursor.Name)
	return cursor
}

// UnstageVoid removes cursor off the model stage
func (cursor *Cursor) UnstageVoid(stage *StageStruct) {
	delete(stage.Cursors, cursor)
	delete(stage.Cursors_mapString, cursor.Name)
}

// commit cursor to the back repo (if it is already staged)
func (cursor *Cursor) Commit(stage *StageStruct) *Cursor {
	if _, ok := stage.Cursors[cursor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCursor(cursor)
		}
	}
	return cursor
}

func (cursor *Cursor) CommitVoid(stage *StageStruct) {
	cursor.Commit(stage)
}

// Checkout cursor to the back repo (if it is already staged)
func (cursor *Cursor) Checkout(stage *StageStruct) *Cursor {
	if _, ok := stage.Cursors[cursor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCursor(cursor)
		}
	}
	return cursor
}

// for satisfaction of GongStruct interface
func (cursor *Cursor) GetName() (res string) {
	return cursor.Name
}

// Stage puts doc to the model stage
func (doc *Doc) Stage(stage *StageStruct) *Doc {

	if _, ok := stage.Docs[doc]; !ok {
		stage.Docs[doc] = __member
		stage.Map_Staged_Order[doc] = stage.Order
		stage.Order++
	}
	stage.Docs_mapString[doc.Name] = doc

	return doc
}

// Unstage removes doc off the model stage
func (doc *Doc) Unstage(stage *StageStruct) *Doc {
	delete(stage.Docs, doc)
	delete(stage.Docs_mapString, doc.Name)
	return doc
}

// UnstageVoid removes doc off the model stage
func (doc *Doc) UnstageVoid(stage *StageStruct) {
	delete(stage.Docs, doc)
	delete(stage.Docs_mapString, doc.Name)
}

// commit doc to the back repo (if it is already staged)
func (doc *Doc) Commit(stage *StageStruct) *Doc {
	if _, ok := stage.Docs[doc]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDoc(doc)
		}
	}
	return doc
}

func (doc *Doc) CommitVoid(stage *StageStruct) {
	doc.Commit(stage)
}

// Checkout doc to the back repo (if it is already staged)
func (doc *Doc) Checkout(stage *StageStruct) *Doc {
	if _, ok := stage.Docs[doc]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDoc(doc)
		}
	}
	return doc
}

// for satisfaction of GongStruct interface
func (doc *Doc) GetName() (res string) {
	return doc.Name
}

// Stage puts form to the model stage
func (form *Form) Stage(stage *StageStruct) *Form {

	if _, ok := stage.Forms[form]; !ok {
		stage.Forms[form] = __member
		stage.Map_Staged_Order[form] = stage.Order
		stage.Order++
	}
	stage.Forms_mapString[form.Name] = form

	return form
}

// Unstage removes form off the model stage
func (form *Form) Unstage(stage *StageStruct) *Form {
	delete(stage.Forms, form)
	delete(stage.Forms_mapString, form.Name)
	return form
}

// UnstageVoid removes form off the model stage
func (form *Form) UnstageVoid(stage *StageStruct) {
	delete(stage.Forms, form)
	delete(stage.Forms_mapString, form.Name)
}

// commit form to the back repo (if it is already staged)
func (form *Form) Commit(stage *StageStruct) *Form {
	if _, ok := stage.Forms[form]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitForm(form)
		}
	}
	return form
}

func (form *Form) CommitVoid(stage *StageStruct) {
	form.Commit(stage)
}

// Checkout form to the back repo (if it is already staged)
func (form *Form) Checkout(stage *StageStruct) *Form {
	if _, ok := stage.Forms[form]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutForm(form)
		}
	}
	return form
}

// for satisfaction of GongStruct interface
func (form *Form) GetName() (res string) {
	return form.Name
}

// Stage puts load to the model stage
func (load *Load) Stage(stage *StageStruct) *Load {

	if _, ok := stage.Loads[load]; !ok {
		stage.Loads[load] = __member
		stage.Map_Staged_Order[load] = stage.Order
		stage.Order++
	}
	stage.Loads_mapString[load.Name] = load

	return load
}

// Unstage removes load off the model stage
func (load *Load) Unstage(stage *StageStruct) *Load {
	delete(stage.Loads, load)
	delete(stage.Loads_mapString, load.Name)
	return load
}

// UnstageVoid removes load off the model stage
func (load *Load) UnstageVoid(stage *StageStruct) {
	delete(stage.Loads, load)
	delete(stage.Loads_mapString, load.Name)
}

// commit load to the back repo (if it is already staged)
func (load *Load) Commit(stage *StageStruct) *Load {
	if _, ok := stage.Loads[load]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLoad(load)
		}
	}
	return load
}

func (load *Load) CommitVoid(stage *StageStruct) {
	load.Commit(stage)
}

// Checkout load to the back repo (if it is already staged)
func (load *Load) Checkout(stage *StageStruct) *Load {
	if _, ok := stage.Loads[load]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLoad(load)
		}
	}
	return load
}

// for satisfaction of GongStruct interface
func (load *Load) GetName() (res string) {
	return load.Name
}

// Stage puts slider to the model stage
func (slider *Slider) Stage(stage *StageStruct) *Slider {

	if _, ok := stage.Sliders[slider]; !ok {
		stage.Sliders[slider] = __member
		stage.Map_Staged_Order[slider] = stage.Order
		stage.Order++
	}
	stage.Sliders_mapString[slider.Name] = slider

	return slider
}

// Unstage removes slider off the model stage
func (slider *Slider) Unstage(stage *StageStruct) *Slider {
	delete(stage.Sliders, slider)
	delete(stage.Sliders_mapString, slider.Name)
	return slider
}

// UnstageVoid removes slider off the model stage
func (slider *Slider) UnstageVoid(stage *StageStruct) {
	delete(stage.Sliders, slider)
	delete(stage.Sliders_mapString, slider.Name)
}

// commit slider to the back repo (if it is already staged)
func (slider *Slider) Commit(stage *StageStruct) *Slider {
	if _, ok := stage.Sliders[slider]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSlider(slider)
		}
	}
	return slider
}

func (slider *Slider) CommitVoid(stage *StageStruct) {
	slider.Commit(stage)
}

// Checkout slider to the back repo (if it is already staged)
func (slider *Slider) Checkout(stage *StageStruct) *Slider {
	if _, ok := stage.Sliders[slider]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSlider(slider)
		}
	}
	return slider
}

// for satisfaction of GongStruct interface
func (slider *Slider) GetName() (res string) {
	return slider.Name
}

// Stage puts split to the model stage
func (split *Split) Stage(stage *StageStruct) *Split {

	if _, ok := stage.Splits[split]; !ok {
		stage.Splits[split] = __member
		stage.Map_Staged_Order[split] = stage.Order
		stage.Order++
	}
	stage.Splits_mapString[split.Name] = split

	return split
}

// Unstage removes split off the model stage
func (split *Split) Unstage(stage *StageStruct) *Split {
	delete(stage.Splits, split)
	delete(stage.Splits_mapString, split.Name)
	return split
}

// UnstageVoid removes split off the model stage
func (split *Split) UnstageVoid(stage *StageStruct) {
	delete(stage.Splits, split)
	delete(stage.Splits_mapString, split.Name)
}

// commit split to the back repo (if it is already staged)
func (split *Split) Commit(stage *StageStruct) *Split {
	if _, ok := stage.Splits[split]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSplit(split)
		}
	}
	return split
}

func (split *Split) CommitVoid(stage *StageStruct) {
	split.Commit(stage)
}

// Checkout split to the back repo (if it is already staged)
func (split *Split) Checkout(stage *StageStruct) *Split {
	if _, ok := stage.Splits[split]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSplit(split)
		}
	}
	return split
}

// for satisfaction of GongStruct interface
func (split *Split) GetName() (res string) {
	return split.Name
}

// Stage puts svg to the model stage
func (svg *Svg) Stage(stage *StageStruct) *Svg {

	if _, ok := stage.Svgs[svg]; !ok {
		stage.Svgs[svg] = __member
		stage.Map_Staged_Order[svg] = stage.Order
		stage.Order++
	}
	stage.Svgs_mapString[svg.Name] = svg

	return svg
}

// Unstage removes svg off the model stage
func (svg *Svg) Unstage(stage *StageStruct) *Svg {
	delete(stage.Svgs, svg)
	delete(stage.Svgs_mapString, svg.Name)
	return svg
}

// UnstageVoid removes svg off the model stage
func (svg *Svg) UnstageVoid(stage *StageStruct) {
	delete(stage.Svgs, svg)
	delete(stage.Svgs_mapString, svg.Name)
}

// commit svg to the back repo (if it is already staged)
func (svg *Svg) Commit(stage *StageStruct) *Svg {
	if _, ok := stage.Svgs[svg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSvg(svg)
		}
	}
	return svg
}

func (svg *Svg) CommitVoid(stage *StageStruct) {
	svg.Commit(stage)
}

// Checkout svg to the back repo (if it is already staged)
func (svg *Svg) Checkout(stage *StageStruct) *Svg {
	if _, ok := stage.Svgs[svg]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSvg(svg)
		}
	}
	return svg
}

// for satisfaction of GongStruct interface
func (svg *Svg) GetName() (res string) {
	return svg.Name
}

// Stage puts table to the model stage
func (table *Table) Stage(stage *StageStruct) *Table {

	if _, ok := stage.Tables[table]; !ok {
		stage.Tables[table] = __member
		stage.Map_Staged_Order[table] = stage.Order
		stage.Order++
	}
	stage.Tables_mapString[table.Name] = table

	return table
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *StageStruct) *Table {
	delete(stage.Tables, table)
	delete(stage.Tables_mapString, table.Name)
	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *StageStruct) {
	delete(stage.Tables, table)
	delete(stage.Tables_mapString, table.Name)
}

// commit table to the back repo (if it is already staged)
func (table *Table) Commit(stage *StageStruct) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTable(table)
		}
	}
	return table
}

func (table *Table) CommitVoid(stage *StageStruct) {
	table.Commit(stage)
}

// Checkout table to the back repo (if it is already staged)
func (table *Table) Checkout(stage *StageStruct) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTable(table)
		}
	}
	return table
}

// for satisfaction of GongStruct interface
func (table *Table) GetName() (res string) {
	return table.Name
}

// Stage puts tone to the model stage
func (tone *Tone) Stage(stage *StageStruct) *Tone {

	if _, ok := stage.Tones[tone]; !ok {
		stage.Tones[tone] = __member
		stage.Map_Staged_Order[tone] = stage.Order
		stage.Order++
	}
	stage.Tones_mapString[tone.Name] = tone

	return tone
}

// Unstage removes tone off the model stage
func (tone *Tone) Unstage(stage *StageStruct) *Tone {
	delete(stage.Tones, tone)
	delete(stage.Tones_mapString, tone.Name)
	return tone
}

// UnstageVoid removes tone off the model stage
func (tone *Tone) UnstageVoid(stage *StageStruct) {
	delete(stage.Tones, tone)
	delete(stage.Tones_mapString, tone.Name)
}

// commit tone to the back repo (if it is already staged)
func (tone *Tone) Commit(stage *StageStruct) *Tone {
	if _, ok := stage.Tones[tone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTone(tone)
		}
	}
	return tone
}

func (tone *Tone) CommitVoid(stage *StageStruct) {
	tone.Commit(stage)
}

// Checkout tone to the back repo (if it is already staged)
func (tone *Tone) Checkout(stage *StageStruct) *Tone {
	if _, ok := stage.Tones[tone]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTone(tone)
		}
	}
	return tone
}

// for satisfaction of GongStruct interface
func (tone *Tone) GetName() (res string) {
	return tone.Name
}

// Stage puts tree to the model stage
func (tree *Tree) Stage(stage *StageStruct) *Tree {

	if _, ok := stage.Trees[tree]; !ok {
		stage.Trees[tree] = __member
		stage.Map_Staged_Order[tree] = stage.Order
		stage.Order++
	}
	stage.Trees_mapString[tree.Name] = tree

	return tree
}

// Unstage removes tree off the model stage
func (tree *Tree) Unstage(stage *StageStruct) *Tree {
	delete(stage.Trees, tree)
	delete(stage.Trees_mapString, tree.Name)
	return tree
}

// UnstageVoid removes tree off the model stage
func (tree *Tree) UnstageVoid(stage *StageStruct) {
	delete(stage.Trees, tree)
	delete(stage.Trees_mapString, tree.Name)
}

// commit tree to the back repo (if it is already staged)
func (tree *Tree) Commit(stage *StageStruct) *Tree {
	if _, ok := stage.Trees[tree]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTree(tree)
		}
	}
	return tree
}

func (tree *Tree) CommitVoid(stage *StageStruct) {
	tree.Commit(stage)
}

// Checkout tree to the back repo (if it is already staged)
func (tree *Tree) Checkout(stage *StageStruct) *Tree {
	if _, ok := stage.Trees[tree]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTree(tree)
		}
	}
	return tree
}

// for satisfaction of GongStruct interface
func (tree *Tree) GetName() (res string) {
	return tree.Name
}

// Stage puts view to the model stage
func (view *View) Stage(stage *StageStruct) *View {

	if _, ok := stage.Views[view]; !ok {
		stage.Views[view] = __member
		stage.Map_Staged_Order[view] = stage.Order
		stage.Order++
	}
	stage.Views_mapString[view.Name] = view

	return view
}

// Unstage removes view off the model stage
func (view *View) Unstage(stage *StageStruct) *View {
	delete(stage.Views, view)
	delete(stage.Views_mapString, view.Name)
	return view
}

// UnstageVoid removes view off the model stage
func (view *View) UnstageVoid(stage *StageStruct) {
	delete(stage.Views, view)
	delete(stage.Views_mapString, view.Name)
}

// commit view to the back repo (if it is already staged)
func (view *View) Commit(stage *StageStruct) *View {
	if _, ok := stage.Views[view]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitView(view)
		}
	}
	return view
}

func (view *View) CommitVoid(stage *StageStruct) {
	view.Commit(stage)
}

// Checkout view to the back repo (if it is already staged)
func (view *View) Checkout(stage *StageStruct) *View {
	if _, ok := stage.Views[view]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutView(view)
		}
	}
	return view
}

// for satisfaction of GongStruct interface
func (view *View) GetName() (res string) {
	return view.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMAsSplit(AsSplit *AsSplit)
	CreateORMAsSplitArea(AsSplitArea *AsSplitArea)
	CreateORMButton(Button *Button)
	CreateORMCursor(Cursor *Cursor)
	CreateORMDoc(Doc *Doc)
	CreateORMForm(Form *Form)
	CreateORMLoad(Load *Load)
	CreateORMSlider(Slider *Slider)
	CreateORMSplit(Split *Split)
	CreateORMSvg(Svg *Svg)
	CreateORMTable(Table *Table)
	CreateORMTone(Tone *Tone)
	CreateORMTree(Tree *Tree)
	CreateORMView(View *View)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMAsSplit(AsSplit *AsSplit)
	DeleteORMAsSplitArea(AsSplitArea *AsSplitArea)
	DeleteORMButton(Button *Button)
	DeleteORMCursor(Cursor *Cursor)
	DeleteORMDoc(Doc *Doc)
	DeleteORMForm(Form *Form)
	DeleteORMLoad(Load *Load)
	DeleteORMSlider(Slider *Slider)
	DeleteORMSplit(Split *Split)
	DeleteORMSvg(Svg *Svg)
	DeleteORMTable(Table *Table)
	DeleteORMTone(Tone *Tone)
	DeleteORMTree(Tree *Tree)
	DeleteORMView(View *View)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.AsSplits = make(map[*AsSplit]any)
	stage.AsSplits_mapString = make(map[string]*AsSplit)

	stage.AsSplitAreas = make(map[*AsSplitArea]any)
	stage.AsSplitAreas_mapString = make(map[string]*AsSplitArea)

	stage.Buttons = make(map[*Button]any)
	stage.Buttons_mapString = make(map[string]*Button)

	stage.Cursors = make(map[*Cursor]any)
	stage.Cursors_mapString = make(map[string]*Cursor)

	stage.Docs = make(map[*Doc]any)
	stage.Docs_mapString = make(map[string]*Doc)

	stage.Forms = make(map[*Form]any)
	stage.Forms_mapString = make(map[string]*Form)

	stage.Loads = make(map[*Load]any)
	stage.Loads_mapString = make(map[string]*Load)

	stage.Sliders = make(map[*Slider]any)
	stage.Sliders_mapString = make(map[string]*Slider)

	stage.Splits = make(map[*Split]any)
	stage.Splits_mapString = make(map[string]*Split)

	stage.Svgs = make(map[*Svg]any)
	stage.Svgs_mapString = make(map[string]*Svg)

	stage.Tables = make(map[*Table]any)
	stage.Tables_mapString = make(map[string]*Table)

	stage.Tones = make(map[*Tone]any)
	stage.Tones_mapString = make(map[string]*Tone)

	stage.Trees = make(map[*Tree]any)
	stage.Trees_mapString = make(map[string]*Tree)

	stage.Views = make(map[*View]any)
	stage.Views_mapString = make(map[string]*View)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.AsSplits = nil
	stage.AsSplits_mapString = nil

	stage.AsSplitAreas = nil
	stage.AsSplitAreas_mapString = nil

	stage.Buttons = nil
	stage.Buttons_mapString = nil

	stage.Cursors = nil
	stage.Cursors_mapString = nil

	stage.Docs = nil
	stage.Docs_mapString = nil

	stage.Forms = nil
	stage.Forms_mapString = nil

	stage.Loads = nil
	stage.Loads_mapString = nil

	stage.Sliders = nil
	stage.Sliders_mapString = nil

	stage.Splits = nil
	stage.Splits_mapString = nil

	stage.Svgs = nil
	stage.Svgs_mapString = nil

	stage.Tables = nil
	stage.Tables_mapString = nil

	stage.Tones = nil
	stage.Tones_mapString = nil

	stage.Trees = nil
	stage.Trees_mapString = nil

	stage.Views = nil
	stage.Views_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for assplit := range stage.AsSplits {
		assplit.Unstage(stage)
	}

	for assplitarea := range stage.AsSplitAreas {
		assplitarea.Unstage(stage)
	}

	for button := range stage.Buttons {
		button.Unstage(stage)
	}

	for cursor := range stage.Cursors {
		cursor.Unstage(stage)
	}

	for doc := range stage.Docs {
		doc.Unstage(stage)
	}

	for form := range stage.Forms {
		form.Unstage(stage)
	}

	for load := range stage.Loads {
		load.Unstage(stage)
	}

	for slider := range stage.Sliders {
		slider.Unstage(stage)
	}

	for split := range stage.Splits {
		split.Unstage(stage)
	}

	for svg := range stage.Svgs {
		svg.Unstage(stage)
	}

	for table := range stage.Tables {
		table.Unstage(stage)
	}

	for tone := range stage.Tones {
		tone.Unstage(stage)
	}

	for tree := range stage.Trees {
		tree.Unstage(stage)
	}

	for view := range stage.Views {
		view.Unstage(stage)
	}

}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface {
}

type GongtructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type PointerToGongstruct interface {
	GetName() string
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]any) (sortedSlice []T) {

	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, CompareGongstructByName)

	return
}

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

	set := GetGongstructInstancesSetFromPointerType[T](stage)
	sortedSlice = SortGongstructSetByName(*set)

	return
}

type GongstructSet interface {
	map[any]any
}

type GongstructMapString interface {
	map[any]any
}

// GongGetSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*AsSplit]any:
		return any(&stage.AsSplits).(*Type)
	case map[*AsSplitArea]any:
		return any(&stage.AsSplitAreas).(*Type)
	case map[*Button]any:
		return any(&stage.Buttons).(*Type)
	case map[*Cursor]any:
		return any(&stage.Cursors).(*Type)
	case map[*Doc]any:
		return any(&stage.Docs).(*Type)
	case map[*Form]any:
		return any(&stage.Forms).(*Type)
	case map[*Load]any:
		return any(&stage.Loads).(*Type)
	case map[*Slider]any:
		return any(&stage.Sliders).(*Type)
	case map[*Split]any:
		return any(&stage.Splits).(*Type)
	case map[*Svg]any:
		return any(&stage.Svgs).(*Type)
	case map[*Table]any:
		return any(&stage.Tables).(*Type)
	case map[*Tone]any:
		return any(&stage.Tones).(*Type)
	case map[*Tree]any:
		return any(&stage.Trees).(*Type)
	case map[*View]any:
		return any(&stage.Views).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*AsSplit:
		return any(&stage.AsSplits_mapString).(*Type)
	case map[string]*AsSplitArea:
		return any(&stage.AsSplitAreas_mapString).(*Type)
	case map[string]*Button:
		return any(&stage.Buttons_mapString).(*Type)
	case map[string]*Cursor:
		return any(&stage.Cursors_mapString).(*Type)
	case map[string]*Doc:
		return any(&stage.Docs_mapString).(*Type)
	case map[string]*Form:
		return any(&stage.Forms_mapString).(*Type)
	case map[string]*Load:
		return any(&stage.Loads_mapString).(*Type)
	case map[string]*Slider:
		return any(&stage.Sliders_mapString).(*Type)
	case map[string]*Split:
		return any(&stage.Splits_mapString).(*Type)
	case map[string]*Svg:
		return any(&stage.Svgs_mapString).(*Type)
	case map[string]*Table:
		return any(&stage.Tables_mapString).(*Type)
	case map[string]*Tone:
		return any(&stage.Tones_mapString).(*Type)
	case map[string]*Tree:
		return any(&stage.Trees_mapString).(*Type)
	case map[string]*View:
		return any(&stage.Views_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AsSplit:
		return any(&stage.AsSplits).(*map[*Type]any)
	case AsSplitArea:
		return any(&stage.AsSplitAreas).(*map[*Type]any)
	case Button:
		return any(&stage.Buttons).(*map[*Type]any)
	case Cursor:
		return any(&stage.Cursors).(*map[*Type]any)
	case Doc:
		return any(&stage.Docs).(*map[*Type]any)
	case Form:
		return any(&stage.Forms).(*map[*Type]any)
	case Load:
		return any(&stage.Loads).(*map[*Type]any)
	case Slider:
		return any(&stage.Sliders).(*map[*Type]any)
	case Split:
		return any(&stage.Splits).(*map[*Type]any)
	case Svg:
		return any(&stage.Svgs).(*map[*Type]any)
	case Table:
		return any(&stage.Tables).(*map[*Type]any)
	case Tone:
		return any(&stage.Tones).(*map[*Type]any)
	case Tree:
		return any(&stage.Trees).(*map[*Type]any)
	case View:
		return any(&stage.Views).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *AsSplit:
		return any(&stage.AsSplits).(*map[Type]any)
	case *AsSplitArea:
		return any(&stage.AsSplitAreas).(*map[Type]any)
	case *Button:
		return any(&stage.Buttons).(*map[Type]any)
	case *Cursor:
		return any(&stage.Cursors).(*map[Type]any)
	case *Doc:
		return any(&stage.Docs).(*map[Type]any)
	case *Form:
		return any(&stage.Forms).(*map[Type]any)
	case *Load:
		return any(&stage.Loads).(*map[Type]any)
	case *Slider:
		return any(&stage.Sliders).(*map[Type]any)
	case *Split:
		return any(&stage.Splits).(*map[Type]any)
	case *Svg:
		return any(&stage.Svgs).(*map[Type]any)
	case *Table:
		return any(&stage.Tables).(*map[Type]any)
	case *Tone:
		return any(&stage.Tones).(*map[Type]any)
	case *Tree:
		return any(&stage.Trees).(*map[Type]any)
	case *View:
		return any(&stage.Views).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case AsSplit:
		return any(&stage.AsSplits_mapString).(*map[string]*Type)
	case AsSplitArea:
		return any(&stage.AsSplitAreas_mapString).(*map[string]*Type)
	case Button:
		return any(&stage.Buttons_mapString).(*map[string]*Type)
	case Cursor:
		return any(&stage.Cursors_mapString).(*map[string]*Type)
	case Doc:
		return any(&stage.Docs_mapString).(*map[string]*Type)
	case Form:
		return any(&stage.Forms_mapString).(*map[string]*Type)
	case Load:
		return any(&stage.Loads_mapString).(*map[string]*Type)
	case Slider:
		return any(&stage.Sliders_mapString).(*map[string]*Type)
	case Split:
		return any(&stage.Splits_mapString).(*map[string]*Type)
	case Svg:
		return any(&stage.Svgs_mapString).(*map[string]*Type)
	case Table:
		return any(&stage.Tables_mapString).(*map[string]*Type)
	case Tone:
		return any(&stage.Tones_mapString).(*map[string]*Type)
	case Tree:
		return any(&stage.Trees_mapString).(*map[string]*Type)
	case View:
		return any(&stage.Views_mapString).(*map[string]*Type)
	default:
		return nil
	}
}

// GetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case AsSplit:
		return any(&AsSplit{
			// Initialisation of associations
			// field is initialized with an instance of AsSplitArea with the name of the field
			AsSplitAreas: []*AsSplitArea{{Name: "AsSplitAreas"}},
		}).(*Type)
	case AsSplitArea:
		return any(&AsSplitArea{
			// Initialisation of associations
			// field is initialized with an instance of AsSplit with the name of the field
			AsSplit: &AsSplit{Name: "AsSplit"},
			// field is initialized with an instance of Button with the name of the field
			Button: &Button{Name: "Button"},
			// field is initialized with an instance of Cursor with the name of the field
			Cursor: &Cursor{Name: "Cursor"},
			// field is initialized with an instance of Doc with the name of the field
			Doc: &Doc{Name: "Doc"},
			// field is initialized with an instance of Form with the name of the field
			Form: &Form{Name: "Form"},
			// field is initialized with an instance of Load with the name of the field
			Load: &Load{Name: "Load"},
			// field is initialized with an instance of Slider with the name of the field
			Slider: &Slider{Name: "Slider"},
			// field is initialized with an instance of Split with the name of the field
			Split: &Split{Name: "Split"},
			// field is initialized with an instance of Svg with the name of the field
			Svg: &Svg{Name: "Svg"},
			// field is initialized with an instance of Table with the name of the field
			Table: &Table{Name: "Table"},
			// field is initialized with an instance of Tone with the name of the field
			Tone: &Tone{Name: "Tone"},
			// field is initialized with an instance of Tree with the name of the field
			Tree: &Tree{Name: "Tree"},
		}).(*Type)
	case Button:
		return any(&Button{
			// Initialisation of associations
		}).(*Type)
	case Cursor:
		return any(&Cursor{
			// Initialisation of associations
		}).(*Type)
	case Doc:
		return any(&Doc{
			// Initialisation of associations
		}).(*Type)
	case Form:
		return any(&Form{
			// Initialisation of associations
		}).(*Type)
	case Load:
		return any(&Load{
			// Initialisation of associations
		}).(*Type)
	case Slider:
		return any(&Slider{
			// Initialisation of associations
		}).(*Type)
	case Split:
		return any(&Split{
			// Initialisation of associations
		}).(*Type)
	case Svg:
		return any(&Svg{
			// Initialisation of associations
		}).(*Type)
	case Table:
		return any(&Table{
			// Initialisation of associations
		}).(*Type)
	case Tone:
		return any(&Tone{
			// Initialisation of associations
		}).(*Type)
	case Tree:
		return any(&Tree{
			// Initialisation of associations
		}).(*Type)
	case View:
		return any(&View{
			// Initialisation of associations
			// field is initialized with an instance of AsSplitArea with the name of the field
			RootAsSplitAreas: []*AsSplitArea{{Name: "RootAsSplitAreas"}},
		}).(*Type)
	default:
		return nil
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AsSplit
	case AsSplit:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of AsSplitArea
	case AsSplitArea:
		switch fieldname {
		// insertion point for per direct association field
		case "AsSplit":
			res := make(map[*AsSplit][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.AsSplit != nil {
					assplit_ := assplitarea.AsSplit
					var assplitareas []*AsSplitArea
					_, ok := res[assplit_]
					if ok {
						assplitareas = res[assplit_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[assplit_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Button":
			res := make(map[*Button][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Button != nil {
					button_ := assplitarea.Button
					var assplitareas []*AsSplitArea
					_, ok := res[button_]
					if ok {
						assplitareas = res[button_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[button_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Cursor":
			res := make(map[*Cursor][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Cursor != nil {
					cursor_ := assplitarea.Cursor
					var assplitareas []*AsSplitArea
					_, ok := res[cursor_]
					if ok {
						assplitareas = res[cursor_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[cursor_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Doc":
			res := make(map[*Doc][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Doc != nil {
					doc_ := assplitarea.Doc
					var assplitareas []*AsSplitArea
					_, ok := res[doc_]
					if ok {
						assplitareas = res[doc_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[doc_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Form":
			res := make(map[*Form][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Form != nil {
					form_ := assplitarea.Form
					var assplitareas []*AsSplitArea
					_, ok := res[form_]
					if ok {
						assplitareas = res[form_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[form_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Load":
			res := make(map[*Load][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Load != nil {
					load_ := assplitarea.Load
					var assplitareas []*AsSplitArea
					_, ok := res[load_]
					if ok {
						assplitareas = res[load_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[load_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Slider":
			res := make(map[*Slider][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Slider != nil {
					slider_ := assplitarea.Slider
					var assplitareas []*AsSplitArea
					_, ok := res[slider_]
					if ok {
						assplitareas = res[slider_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[slider_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Split":
			res := make(map[*Split][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Split != nil {
					split_ := assplitarea.Split
					var assplitareas []*AsSplitArea
					_, ok := res[split_]
					if ok {
						assplitareas = res[split_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[split_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Svg":
			res := make(map[*Svg][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Svg != nil {
					svg_ := assplitarea.Svg
					var assplitareas []*AsSplitArea
					_, ok := res[svg_]
					if ok {
						assplitareas = res[svg_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[svg_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Table":
			res := make(map[*Table][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Table != nil {
					table_ := assplitarea.Table
					var assplitareas []*AsSplitArea
					_, ok := res[table_]
					if ok {
						assplitareas = res[table_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[table_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tone":
			res := make(map[*Tone][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Tone != nil {
					tone_ := assplitarea.Tone
					var assplitareas []*AsSplitArea
					_, ok := res[tone_]
					if ok {
						assplitareas = res[tone_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[tone_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		case "Tree":
			res := make(map[*Tree][]*AsSplitArea)
			for assplitarea := range stage.AsSplitAreas {
				if assplitarea.Tree != nil {
					tree_ := assplitarea.Tree
					var assplitareas []*AsSplitArea
					_, ok := res[tree_]
					if ok {
						assplitareas = res[tree_]
					} else {
						assplitareas = make([]*AsSplitArea, 0)
					}
					assplitareas = append(assplitareas, assplitarea)
					res[tree_] = assplitareas
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cursor
	case Cursor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Doc
	case Doc:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Form
	case Form:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Load
	case Load:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slider
	case Slider:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Split
	case Split:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Svg
	case Svg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tone
	case Tone:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of View
	case View:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetSliceOfPointersReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..N) between one staged Gongstruct instances and many others
//
// The function provides a map with keys as instances of End and values to *Start instances
// the map is construed by iterating over all Start instances and populating keys with End instances
// and values with the Start instances
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of AsSplit
	case AsSplit:
		switch fieldname {
		// insertion point for per direct association field
		case "AsSplitAreas":
			res := make(map[*AsSplitArea]*AsSplit)
			for assplit := range stage.AsSplits {
				for _, assplitarea_ := range assplit.AsSplitAreas {
					res[assplitarea_] = assplit
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of AsSplitArea
	case AsSplitArea:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cursor
	case Cursor:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Doc
	case Doc:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Form
	case Form:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Load
	case Load:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slider
	case Slider:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Split
	case Split:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Svg
	case Svg:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tone
	case Tone:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of View
	case View:
		switch fieldname {
		// insertion point for per direct association field
		case "RootAsSplitAreas":
			res := make(map[*AsSplitArea]*View)
			for view := range stage.Views {
				for _, assplitarea_ := range view.RootAsSplitAreas {
					res[assplitarea_] = view
				}
			}
			return any(res).(map[*End]*Start)
		}
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AsSplit:
		res = "AsSplit"
	case AsSplitArea:
		res = "AsSplitArea"
	case Button:
		res = "Button"
	case Cursor:
		res = "Cursor"
	case Doc:
		res = "Doc"
	case Form:
		res = "Form"
	case Load:
		res = "Load"
	case Slider:
		res = "Slider"
	case Split:
		res = "Split"
	case Svg:
		res = "Svg"
	case Table:
		res = "Table"
	case Tone:
		res = "Tone"
	case Tree:
		res = "Tree"
	case View:
		res = "View"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AsSplit:
		res = "AsSplit"
	case *AsSplitArea:
		res = "AsSplitArea"
	case *Button:
		res = "Button"
	case *Cursor:
		res = "Cursor"
	case *Doc:
		res = "Doc"
	case *Form:
		res = "Form"
	case *Load:
		res = "Load"
	case *Slider:
		res = "Slider"
	case *Split:
		res = "Split"
	case *Svg:
		res = "Svg"
	case *Table:
		res = "Table"
	case *Tone:
		res = "Tone"
	case *Tree:
		res = "Tree"
	case *View:
		res = "View"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case AsSplit:
		res = []string{"Name", "Direction", "AsSplitAreas"}
	case AsSplitArea:
		res = []string{"Name", "ShowNameInHeader", "Size", "IsAny", "AsSplit", "Button", "Cursor", "Doc", "Form", "Load", "Slider", "Split", "Svg", "Table", "Tone", "Tree", "HasDiv", "DivStyle"}
	case Button:
		res = []string{"Name", "StackName"}
	case Cursor:
		res = []string{"Name", "StackName", "Style"}
	case Doc:
		res = []string{"Name", "StackName"}
	case Form:
		res = []string{"Name", "StackName", "FormName"}
	case Load:
		res = []string{"Name", "StackName"}
	case Slider:
		res = []string{"Name", "StackName"}
	case Split:
		res = []string{"Name", "StackName"}
	case Svg:
		res = []string{"Name", "StackName", "Style"}
	case Table:
		res = []string{"Name", "StackName", "TableName"}
	case Tone:
		res = []string{"Name", "StackName"}
	case Tree:
		res = []string{"Name", "StackName", "TreeName"}
	case View:
		res = []string{"Name", "RootAsSplitAreas"}
	}
	return
}

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type Gongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case AsSplit:
		var rf ReverseField
		_ = rf
	case AsSplitArea:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "AsSplit"
		rf.Fieldname = "AsSplitAreas"
		res = append(res, rf)
		rf.GongstructName = "View"
		rf.Fieldname = "RootAsSplitAreas"
		res = append(res, rf)
	case Button:
		var rf ReverseField
		_ = rf
	case Cursor:
		var rf ReverseField
		_ = rf
	case Doc:
		var rf ReverseField
		_ = rf
	case Form:
		var rf ReverseField
		_ = rf
	case Load:
		var rf ReverseField
		_ = rf
	case Slider:
		var rf ReverseField
		_ = rf
	case Split:
		var rf ReverseField
		_ = rf
	case Svg:
		var rf ReverseField
		_ = rf
	case Table:
		var rf ReverseField
		_ = rf
	case Tone:
		var rf ReverseField
		_ = rf
	case Tree:
		var rf ReverseField
		_ = rf
	case View:
		var rf ReverseField
		_ = rf
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *AsSplit:
		res = []string{"Name", "Direction", "AsSplitAreas"}
	case *AsSplitArea:
		res = []string{"Name", "ShowNameInHeader", "Size", "IsAny", "AsSplit", "Button", "Cursor", "Doc", "Form", "Load", "Slider", "Split", "Svg", "Table", "Tone", "Tree", "HasDiv", "DivStyle"}
	case *Button:
		res = []string{"Name", "StackName"}
	case *Cursor:
		res = []string{"Name", "StackName", "Style"}
	case *Doc:
		res = []string{"Name", "StackName"}
	case *Form:
		res = []string{"Name", "StackName", "FormName"}
	case *Load:
		res = []string{"Name", "StackName"}
	case *Slider:
		res = []string{"Name", "StackName"}
	case *Split:
		res = []string{"Name", "StackName"}
	case *Svg:
		res = []string{"Name", "StackName", "Style"}
	case *Table:
		res = []string{"Name", "StackName", "TableName"}
	case *Tone:
		res = []string{"Name", "StackName"}
	case *Tree:
		res = []string{"Name", "StackName", "TreeName"}
	case *View:
		res = []string{"Name", "RootAsSplitAreas"}
	}
	return
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt    GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat  GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool   GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeOthers GongFieldValueType = "GongFieldValueTypeOthers"
)

type GongFieldValue struct {
	valueString string
	GongFieldValueType
	valueInt   int
	valueFloat float64
	valueBool  bool
}

func (gongValueField *GongFieldValue) GetValueString() string {
	return gongValueField.valueString
}

func (gongValueField *GongFieldValue) GetValueInt() int {
	return gongValueField.valueInt
}

func (gongValueField *GongFieldValue) GetValueFloat() float64 {
	return gongValueField.valueFloat
}

func (gongValueField *GongFieldValue) GetValueBool() bool {
	return gongValueField.valueBool
}

func GetFieldStringValueFromPointer(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *AsSplit:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Direction":
			enum := inferedInstance.Direction
			res.valueString = enum.ToCodeString()
		case "AsSplitAreas":
			for idx, __instance__ := range inferedInstance.AsSplitAreas {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *AsSplitArea:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ShowNameInHeader":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowNameInHeader)
			res.valueBool = inferedInstance.ShowNameInHeader
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Size":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Size)
			res.valueFloat = inferedInstance.Size
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "IsAny":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsAny)
			res.valueBool = inferedInstance.IsAny
			res.GongFieldValueType = GongFieldValueTypeBool
		case "AsSplit":
			if inferedInstance.AsSplit != nil {
				res.valueString = inferedInstance.AsSplit.Name
			}
		case "Button":
			if inferedInstance.Button != nil {
				res.valueString = inferedInstance.Button.Name
			}
		case "Cursor":
			if inferedInstance.Cursor != nil {
				res.valueString = inferedInstance.Cursor.Name
			}
		case "Doc":
			if inferedInstance.Doc != nil {
				res.valueString = inferedInstance.Doc.Name
			}
		case "Form":
			if inferedInstance.Form != nil {
				res.valueString = inferedInstance.Form.Name
			}
		case "Load":
			if inferedInstance.Load != nil {
				res.valueString = inferedInstance.Load.Name
			}
		case "Slider":
			if inferedInstance.Slider != nil {
				res.valueString = inferedInstance.Slider.Name
			}
		case "Split":
			if inferedInstance.Split != nil {
				res.valueString = inferedInstance.Split.Name
			}
		case "Svg":
			if inferedInstance.Svg != nil {
				res.valueString = inferedInstance.Svg.Name
			}
		case "Table":
			if inferedInstance.Table != nil {
				res.valueString = inferedInstance.Table.Name
			}
		case "Tone":
			if inferedInstance.Tone != nil {
				res.valueString = inferedInstance.Tone.Name
			}
		case "Tree":
			if inferedInstance.Tree != nil {
				res.valueString = inferedInstance.Tree.Name
			}
		case "HasDiv":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasDiv)
			res.valueBool = inferedInstance.HasDiv
			res.GongFieldValueType = GongFieldValueTypeBool
		case "DivStyle":
			res.valueString = inferedInstance.DivStyle
		}
	case *Button:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case *Cursor:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "Style":
			res.valueString = inferedInstance.Style
		}
	case *Doc:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case *Form:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "FormName":
			res.valueString = inferedInstance.FormName
		}
	case *Load:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case *Slider:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case *Split:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case *Svg:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "Style":
			res.valueString = inferedInstance.Style
		}
	case *Table:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "TableName":
			res.valueString = inferedInstance.TableName
		}
	case *Tone:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case *Tree:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "TreeName":
			res.valueString = inferedInstance.TreeName
		}
	case *View:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "RootAsSplitAreas":
			for idx, __instance__ := range inferedInstance.RootAsSplitAreas {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case AsSplit:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Direction":
			enum := inferedInstance.Direction
			res.valueString = enum.ToCodeString()
		case "AsSplitAreas":
			for idx, __instance__ := range inferedInstance.AsSplitAreas {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case AsSplitArea:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ShowNameInHeader":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ShowNameInHeader)
			res.valueBool = inferedInstance.ShowNameInHeader
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Size":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Size)
			res.valueFloat = inferedInstance.Size
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "IsAny":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsAny)
			res.valueBool = inferedInstance.IsAny
			res.GongFieldValueType = GongFieldValueTypeBool
		case "AsSplit":
			if inferedInstance.AsSplit != nil {
				res.valueString = inferedInstance.AsSplit.Name
			}
		case "Button":
			if inferedInstance.Button != nil {
				res.valueString = inferedInstance.Button.Name
			}
		case "Cursor":
			if inferedInstance.Cursor != nil {
				res.valueString = inferedInstance.Cursor.Name
			}
		case "Doc":
			if inferedInstance.Doc != nil {
				res.valueString = inferedInstance.Doc.Name
			}
		case "Form":
			if inferedInstance.Form != nil {
				res.valueString = inferedInstance.Form.Name
			}
		case "Load":
			if inferedInstance.Load != nil {
				res.valueString = inferedInstance.Load.Name
			}
		case "Slider":
			if inferedInstance.Slider != nil {
				res.valueString = inferedInstance.Slider.Name
			}
		case "Split":
			if inferedInstance.Split != nil {
				res.valueString = inferedInstance.Split.Name
			}
		case "Svg":
			if inferedInstance.Svg != nil {
				res.valueString = inferedInstance.Svg.Name
			}
		case "Table":
			if inferedInstance.Table != nil {
				res.valueString = inferedInstance.Table.Name
			}
		case "Tone":
			if inferedInstance.Tone != nil {
				res.valueString = inferedInstance.Tone.Name
			}
		case "Tree":
			if inferedInstance.Tree != nil {
				res.valueString = inferedInstance.Tree.Name
			}
		case "HasDiv":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasDiv)
			res.valueBool = inferedInstance.HasDiv
			res.GongFieldValueType = GongFieldValueTypeBool
		case "DivStyle":
			res.valueString = inferedInstance.DivStyle
		}
	case Button:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case Cursor:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "Style":
			res.valueString = inferedInstance.Style
		}
	case Doc:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case Form:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "FormName":
			res.valueString = inferedInstance.FormName
		}
	case Load:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case Slider:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case Split:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case Svg:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "Style":
			res.valueString = inferedInstance.Style
		}
	case Table:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "TableName":
			res.valueString = inferedInstance.TableName
		}
	case Tone:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		}
	case Tree:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "StackName":
			res.valueString = inferedInstance.StackName
		case "TreeName":
			res.valueString = inferedInstance.TreeName
		}
	case View:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "RootAsSplitAreas":
			for idx, __instance__ := range inferedInstance.RootAsSplitAreas {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
