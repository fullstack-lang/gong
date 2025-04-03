// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"

	slider_go "github.com/fullstack-lang/gong/lib/slider/go"
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

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
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

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Checkboxs           map[*Checkbox]any
	Checkboxs_mapString map[string]*Checkbox

	// insertion point for slice of pointers maps
	OnAfterCheckboxCreateCallback OnAfterCreateInterface[Checkbox]
	OnAfterCheckboxUpdateCallback OnAfterUpdateInterface[Checkbox]
	OnAfterCheckboxDeleteCallback OnAfterDeleteInterface[Checkbox]
	OnAfterCheckboxReadCallback   OnAfterReadInterface[Checkbox]

	Groups           map[*Group]any
	Groups_mapString map[string]*Group

	// insertion point for slice of pointers maps
	Group_Sliders_reverseMap map[*Slider]*Group

	Group_Checkboxes_reverseMap map[*Checkbox]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	Layouts           map[*Layout]any
	Layouts_mapString map[string]*Layout

	// insertion point for slice of pointers maps
	Layout_Groups_reverseMap map[*Group]*Layout

	OnAfterLayoutCreateCallback OnAfterCreateInterface[Layout]
	OnAfterLayoutUpdateCallback OnAfterUpdateInterface[Layout]
	OnAfterLayoutDeleteCallback OnAfterDeleteInterface[Layout]
	OnAfterLayoutReadCallback   OnAfterReadInterface[Layout]

	Sliders           map[*Slider]any
	Sliders_mapString map[string]*Slider

	// insertion point for slice of pointers maps
	OnAfterSliderCreateCallback OnAfterCreateInterface[Slider]
	OnAfterSliderUpdateCallback OnAfterUpdateInterface[Slider]
	OnAfterSliderDeleteCallback OnAfterDeleteInterface[Slider]
	OnAfterSliderReadCallback   OnAfterReadInterface[Slider]

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
	// insertion point for order fields declaration
	CheckboxOrder            uint
	CheckboxMap_Staged_Order map[*Checkbox]uint

	GroupOrder            uint
	GroupMap_Staged_Order map[*Group]uint

	LayoutOrder            uint
	LayoutMap_Staged_Order map[*Layout]uint

	SliderOrder            uint
	SliderMap_Staged_Order map[*Slider]uint

	// end of insertion point
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/lib/slider/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return slider_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return slider_go.GoDiagramsDir
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *Stage)
	Checkout(stage *Stage)
	Backup(stage *Stage, dirPath string)
	Restore(stage *Stage, dirPath string)
	BackupXL(stage *Stage, dirPath string)
	RestoreXL(stage *Stage, dirPath string)
	// insertion point for Commit and Checkout signatures
	CommitCheckbox(checkbox *Checkbox)
	CheckoutCheckbox(checkbox *Checkbox)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitLayout(layout *Layout)
	CheckoutLayout(layout *Layout)
	CommitSlider(slider *Slider)
	CheckoutSlider(slider *Slider)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Checkboxs:           make(map[*Checkbox]any),
		Checkboxs_mapString: make(map[string]*Checkbox),

		Groups:           make(map[*Group]any),
		Groups_mapString: make(map[string]*Group),

		Layouts:           make(map[*Layout]any),
		Layouts_mapString: make(map[string]*Layout),

		Sliders:           make(map[*Slider]any),
		Sliders_mapString: make(map[string]*Slider),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CheckboxMap_Staged_Order: make(map[*Checkbox]uint),

		GroupMap_Staged_Order: make(map[*Group]uint),

		LayoutMap_Staged_Order: make(map[*Layout]uint),

		SliderMap_Staged_Order: make(map[*Slider]uint),

		// end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Checkbox:
		return stage.CheckboxMap_Staged_Order[instance]
	case *Group:
		return stage.GroupMap_Staged_Order[instance]
	case *Layout:
		return stage.LayoutMap_Staged_Order[instance]
	case *Slider:
		return stage.SliderMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Checkbox"] = len(stage.Checkboxs)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Layout"] = len(stage.Layouts)
	stage.Map_GongStructName_InstancesNb["Slider"] = len(stage.Sliders)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Checkbox"] = len(stage.Checkboxs)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Layout"] = len(stage.Layouts)
	stage.Map_GongStructName_InstancesNb["Slider"] = len(stage.Sliders)

}

// backup generates backup files in the dirPath
func (stage *Stage) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *Stage) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *Stage) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts checkbox to the model stage
func (checkbox *Checkbox) Stage(stage *Stage) *Checkbox {

	if _, ok := stage.Checkboxs[checkbox]; !ok {
		stage.Checkboxs[checkbox] = __member
		stage.CheckboxMap_Staged_Order[checkbox] = stage.CheckboxOrder
		stage.CheckboxOrder++
	}
	stage.Checkboxs_mapString[checkbox.Name] = checkbox

	return checkbox
}

// Unstage removes checkbox off the model stage
func (checkbox *Checkbox) Unstage(stage *Stage) *Checkbox {
	delete(stage.Checkboxs, checkbox)
	delete(stage.Checkboxs_mapString, checkbox.Name)
	return checkbox
}

// UnstageVoid removes checkbox off the model stage
func (checkbox *Checkbox) UnstageVoid(stage *Stage) {
	delete(stage.Checkboxs, checkbox)
	delete(stage.Checkboxs_mapString, checkbox.Name)
}

// commit checkbox to the back repo (if it is already staged)
func (checkbox *Checkbox) Commit(stage *Stage) *Checkbox {
	if _, ok := stage.Checkboxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCheckbox(checkbox)
		}
	}
	return checkbox
}

func (checkbox *Checkbox) CommitVoid(stage *Stage) {
	checkbox.Commit(stage)
}

// Checkout checkbox to the back repo (if it is already staged)
func (checkbox *Checkbox) Checkout(stage *Stage) *Checkbox {
	if _, ok := stage.Checkboxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCheckbox(checkbox)
		}
	}
	return checkbox
}

// for satisfaction of GongStruct interface
func (checkbox *Checkbox) GetName() (res string) {
	return checkbox.Name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *Stage) *Group {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = __member
		stage.GroupMap_Staged_Order[group] = stage.GroupOrder
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group

	return group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
}

// commit group to the back repo (if it is already staged)
func (group *Group) Commit(stage *Stage) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroup(group)
		}
	}
	return group
}

func (group *Group) CommitVoid(stage *Stage) {
	group.Commit(stage)
}

// Checkout group to the back repo (if it is already staged)
func (group *Group) Checkout(stage *Stage) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroup(group)
		}
	}
	return group
}

// for satisfaction of GongStruct interface
func (group *Group) GetName() (res string) {
	return group.Name
}

// Stage puts layout to the model stage
func (layout *Layout) Stage(stage *Stage) *Layout {

	if _, ok := stage.Layouts[layout]; !ok {
		stage.Layouts[layout] = __member
		stage.LayoutMap_Staged_Order[layout] = stage.LayoutOrder
		stage.LayoutOrder++
	}
	stage.Layouts_mapString[layout.Name] = layout

	return layout
}

// Unstage removes layout off the model stage
func (layout *Layout) Unstage(stage *Stage) *Layout {
	delete(stage.Layouts, layout)
	delete(stage.Layouts_mapString, layout.Name)
	return layout
}

// UnstageVoid removes layout off the model stage
func (layout *Layout) UnstageVoid(stage *Stage) {
	delete(stage.Layouts, layout)
	delete(stage.Layouts_mapString, layout.Name)
}

// commit layout to the back repo (if it is already staged)
func (layout *Layout) Commit(stage *Stage) *Layout {
	if _, ok := stage.Layouts[layout]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLayout(layout)
		}
	}
	return layout
}

func (layout *Layout) CommitVoid(stage *Stage) {
	layout.Commit(stage)
}

// Checkout layout to the back repo (if it is already staged)
func (layout *Layout) Checkout(stage *Stage) *Layout {
	if _, ok := stage.Layouts[layout]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutLayout(layout)
		}
	}
	return layout
}

// for satisfaction of GongStruct interface
func (layout *Layout) GetName() (res string) {
	return layout.Name
}

// Stage puts slider to the model stage
func (slider *Slider) Stage(stage *Stage) *Slider {

	if _, ok := stage.Sliders[slider]; !ok {
		stage.Sliders[slider] = __member
		stage.SliderMap_Staged_Order[slider] = stage.SliderOrder
		stage.SliderOrder++
	}
	stage.Sliders_mapString[slider.Name] = slider

	return slider
}

// Unstage removes slider off the model stage
func (slider *Slider) Unstage(stage *Stage) *Slider {
	delete(stage.Sliders, slider)
	delete(stage.Sliders_mapString, slider.Name)
	return slider
}

// UnstageVoid removes slider off the model stage
func (slider *Slider) UnstageVoid(stage *Stage) {
	delete(stage.Sliders, slider)
	delete(stage.Sliders_mapString, slider.Name)
}

// commit slider to the back repo (if it is already staged)
func (slider *Slider) Commit(stage *Stage) *Slider {
	if _, ok := stage.Sliders[slider]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSlider(slider)
		}
	}
	return slider
}

func (slider *Slider) CommitVoid(stage *Stage) {
	slider.Commit(stage)
}

// Checkout slider to the back repo (if it is already staged)
func (slider *Slider) Checkout(stage *Stage) *Slider {
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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCheckbox(Checkbox *Checkbox)
	CreateORMGroup(Group *Group)
	CreateORMLayout(Layout *Layout)
	CreateORMSlider(Slider *Slider)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCheckbox(Checkbox *Checkbox)
	DeleteORMGroup(Group *Group)
	DeleteORMLayout(Layout *Layout)
	DeleteORMSlider(Slider *Slider)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Checkboxs = make(map[*Checkbox]any)
	stage.Checkboxs_mapString = make(map[string]*Checkbox)
	stage.CheckboxMap_Staged_Order = make(map[*Checkbox]uint)
	stage.CheckboxOrder = 0

	stage.Groups = make(map[*Group]any)
	stage.Groups_mapString = make(map[string]*Group)
	stage.GroupMap_Staged_Order = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.Layouts = make(map[*Layout]any)
	stage.Layouts_mapString = make(map[string]*Layout)
	stage.LayoutMap_Staged_Order = make(map[*Layout]uint)
	stage.LayoutOrder = 0

	stage.Sliders = make(map[*Slider]any)
	stage.Sliders_mapString = make(map[string]*Slider)
	stage.SliderMap_Staged_Order = make(map[*Slider]uint)
	stage.SliderOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Checkboxs = nil
	stage.Checkboxs_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.Layouts = nil
	stage.Layouts_mapString = nil

	stage.Sliders = nil
	stage.Sliders_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for checkbox := range stage.Checkboxs {
		checkbox.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
	}

	for layout := range stage.Layouts {
		layout.Unstage(stage)
	}

	for slider := range stage.Sliders {
		slider.Unstage(stage)
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
	CommitVoid(*Stage)
	UnstageVoid(stage *Stage)
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

func GetGongstrucsSorted[T PointerToGongstruct](stage *Stage) (sortedSlice []T) {

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
func GongGetSet[Type GongstructSet](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[*Checkbox]any:
		return any(&stage.Checkboxs).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
	case map[*Layout]any:
		return any(&stage.Layouts).(*Type)
	case map[*Slider]any:
		return any(&stage.Sliders).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Checkbox:
		return any(&stage.Checkboxs_mapString).(*Type)
	case map[string]*Group:
		return any(&stage.Groups_mapString).(*Type)
	case map[string]*Layout:
		return any(&stage.Layouts_mapString).(*Type)
	case map[string]*Slider:
		return any(&stage.Sliders_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Checkbox:
		return any(&stage.Checkboxs).(*map[*Type]any)
	case Group:
		return any(&stage.Groups).(*map[*Type]any)
	case Layout:
		return any(&stage.Layouts).(*map[*Type]any)
	case Slider:
		return any(&stage.Sliders).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Checkbox:
		return any(&stage.Checkboxs).(*map[Type]any)
	case *Group:
		return any(&stage.Groups).(*map[Type]any)
	case *Layout:
		return any(&stage.Layouts).(*map[Type]any)
	case *Slider:
		return any(&stage.Sliders).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Checkbox:
		return any(&stage.Checkboxs_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
	case Layout:
		return any(&stage.Layouts_mapString).(*map[string]*Type)
	case Slider:
		return any(&stage.Sliders_mapString).(*map[string]*Type)
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
	case Checkbox:
		return any(&Checkbox{
			// Initialisation of associations
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of Slider with the name of the field
			Sliders: []*Slider{{Name: "Sliders"}},
			// field is initialized with an instance of Checkbox with the name of the field
			Checkboxes: []*Checkbox{{Name: "Checkboxes"}},
		}).(*Type)
	case Layout:
		return any(&Layout{
			// Initialisation of associations
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
		}).(*Type)
	case Slider:
		return any(&Slider{
			// Initialisation of associations
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Checkbox
	case Checkbox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Layout
	case Layout:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Slider
	case Slider:
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of Checkbox
	case Checkbox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "Sliders":
			res := make(map[*Slider]*Group)
			for group := range stage.Groups {
				for _, slider_ := range group.Sliders {
					res[slider_] = group
				}
			}
			return any(res).(map[*End]*Start)
		case "Checkboxes":
			res := make(map[*Checkbox]*Group)
			for group := range stage.Groups {
				for _, checkbox_ := range group.Checkboxes {
					res[checkbox_] = group
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Layout
	case Layout:
		switch fieldname {
		// insertion point for per direct association field
		case "Groups":
			res := make(map[*Group]*Layout)
			for layout := range stage.Layouts {
				for _, group_ := range layout.Groups {
					res[group_] = layout
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Slider
	case Slider:
		switch fieldname {
		// insertion point for per direct association field
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
	case Checkbox:
		res = "Checkbox"
	case Group:
		res = "Group"
	case Layout:
		res = "Layout"
	case Slider:
		res = "Slider"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Checkbox:
		res = "Checkbox"
	case *Group:
		res = "Group"
	case *Layout:
		res = "Layout"
	case *Slider:
		res = "Slider"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Checkbox:
		res = []string{"Name", "ValueBool", "LabelForTrue", "LabelForFalse"}
	case Group:
		res = []string{"Name", "Percentage", "Sliders", "Checkboxes"}
	case Layout:
		res = []string{"Name", "Groups"}
	case Slider:
		res = []string{"Name", "IsFloat64", "IsInt", "MinInt", "MaxInt", "StepInt", "ValueInt", "MinFloat64", "MaxFloat64", "StepFloat64", "ValueFloat64"}
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
	case Checkbox:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Checkboxes"
		res = append(res, rf)
	case Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layout"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case Layout:
		var rf ReverseField
		_ = rf
	case Slider:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Sliders"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Checkbox:
		res = []string{"Name", "ValueBool", "LabelForTrue", "LabelForFalse"}
	case *Group:
		res = []string{"Name", "Percentage", "Sliders", "Checkboxes"}
	case *Layout:
		res = []string{"Name", "Groups"}
	case *Slider:
		res = []string{"Name", "IsFloat64", "IsInt", "MinInt", "MaxInt", "StepInt", "ValueInt", "MinFloat64", "MaxFloat64", "StepFloat64", "ValueFloat64"}
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
	case *Checkbox:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ValueBool":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ValueBool)
			res.valueBool = inferedInstance.ValueBool
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LabelForTrue":
			res.valueString = inferedInstance.LabelForTrue
		case "LabelForFalse":
			res.valueString = inferedInstance.LabelForFalse
		}
	case *Group:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Percentage":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Percentage)
			res.valueFloat = inferedInstance.Percentage
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Sliders":
			for idx, __instance__ := range inferedInstance.Sliders {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Checkboxes":
			for idx, __instance__ := range inferedInstance.Checkboxes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Layout:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *Slider:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsFloat64":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsFloat64)
			res.valueBool = inferedInstance.IsFloat64
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsInt":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInt)
			res.valueBool = inferedInstance.IsInt
			res.GongFieldValueType = GongFieldValueTypeBool
		case "MinInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MinInt)
			res.valueInt = inferedInstance.MinInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "MaxInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MaxInt)
			res.valueInt = inferedInstance.MaxInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "StepInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.StepInt)
			res.valueInt = inferedInstance.StepInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ValueInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ValueInt)
			res.valueInt = inferedInstance.ValueInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "MinFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MinFloat64)
			res.valueFloat = inferedInstance.MinFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "MaxFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MaxFloat64)
			res.valueFloat = inferedInstance.MaxFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StepFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StepFloat64)
			res.valueFloat = inferedInstance.StepFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ValueFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ValueFloat64)
			res.valueFloat = inferedInstance.ValueFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Checkbox:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "ValueBool":
			res.valueString = fmt.Sprintf("%t", inferedInstance.ValueBool)
			res.valueBool = inferedInstance.ValueBool
			res.GongFieldValueType = GongFieldValueTypeBool
		case "LabelForTrue":
			res.valueString = inferedInstance.LabelForTrue
		case "LabelForFalse":
			res.valueString = inferedInstance.LabelForFalse
		}
	case Group:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Percentage":
			res.valueString = fmt.Sprintf("%f", inferedInstance.Percentage)
			res.valueFloat = inferedInstance.Percentage
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Sliders":
			for idx, __instance__ := range inferedInstance.Sliders {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Checkboxes":
			for idx, __instance__ := range inferedInstance.Checkboxes {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Layout:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Groups":
			for idx, __instance__ := range inferedInstance.Groups {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case Slider:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "IsFloat64":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsFloat64)
			res.valueBool = inferedInstance.IsFloat64
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsInt":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInt)
			res.valueBool = inferedInstance.IsInt
			res.GongFieldValueType = GongFieldValueTypeBool
		case "MinInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MinInt)
			res.valueInt = inferedInstance.MinInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "MaxInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MaxInt)
			res.valueInt = inferedInstance.MaxInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "StepInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.StepInt)
			res.valueInt = inferedInstance.StepInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "ValueInt":
			res.valueString = fmt.Sprintf("%d", inferedInstance.ValueInt)
			res.valueInt = inferedInstance.ValueInt
			res.GongFieldValueType = GongFieldValueTypeInt
		case "MinFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MinFloat64)
			res.valueFloat = inferedInstance.MinFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "MaxFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.MaxFloat64)
			res.valueFloat = inferedInstance.MaxFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StepFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.StepFloat64)
			res.valueFloat = inferedInstance.StepFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "ValueFloat64":
			res.valueString = fmt.Sprintf("%f", inferedInstance.ValueFloat64)
			res.valueFloat = inferedInstance.ValueFloat64
			res.GongFieldValueType = GongFieldValueTypeFloat
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
