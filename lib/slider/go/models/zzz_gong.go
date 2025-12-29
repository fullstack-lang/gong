// generated code - do not edit
package models

import (
	"cmp"
	"embed"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"time"

	slider_go "github.com/fullstack-lang/gong/lib/slider/go"
)

// can be used for
//
//	days := __Gong__Abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var _ = __Gong__Abs
var _ = strings.Clone("")

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
const ProbeNotificationTableSuffix = ":notification table of the probe"
const ProbeFormSuffix = ":form of the probe"
const ProbeSplitSuffix = ":probe of the probe"

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")
var _ = errUnkownEnum

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var __member __void
var _ = __member

// GongStructInterface is the interface met by GongStructs
// It allows runtime reflexion of instances (without the hassle of the "reflect" package)
type GongStructInterface interface {
	GetName() (res string)
	// GetID() (res int)
	// GetFields() (res []string)
	// GetFieldStringValue(fieldName string) (res string)
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
}

// Stage enables storage of staged instances
// swagger:ignore
type Stage struct {
	name string

	// insertion point for definition of arrays registering instances
	Checkboxs           map[*Checkbox]struct{}
	Checkboxs_reference map[*Checkbox]*Checkbox
	Checkboxs_mapString map[string]*Checkbox

	// insertion point for slice of pointers maps
	OnAfterCheckboxCreateCallback OnAfterCreateInterface[Checkbox]
	OnAfterCheckboxUpdateCallback OnAfterUpdateInterface[Checkbox]
	OnAfterCheckboxDeleteCallback OnAfterDeleteInterface[Checkbox]
	OnAfterCheckboxReadCallback   OnAfterReadInterface[Checkbox]

	Groups           map[*Group]struct{}
	Groups_reference map[*Group]*Group
	Groups_mapString map[string]*Group

	// insertion point for slice of pointers maps
	Group_Sliders_reverseMap map[*Slider]*Group

	Group_Checkboxes_reverseMap map[*Checkbox]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	Layouts           map[*Layout]struct{}
	Layouts_reference map[*Layout]*Layout
	Layouts_mapString map[string]*Layout

	// insertion point for slice of pointers maps
	Layout_Groups_reverseMap map[*Group]*Layout

	OnAfterLayoutCreateCallback OnAfterCreateInterface[Layout]
	OnAfterLayoutUpdateCallback OnAfterUpdateInterface[Layout]
	OnAfterLayoutDeleteCallback OnAfterDeleteInterface[Layout]
	OnAfterLayoutReadCallback   OnAfterReadInterface[Layout]

	Sliders           map[*Slider]struct{}
	Sliders_reference map[*Slider]*Slider
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

	NamedStructs []*NamedStruct

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
	return stage.probeIF
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []string) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	for _, instance := range orderedSet {
		res = append(res, instance.GetName())
	}

	return
}

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Checkbox:
		tmp := GetStructInstancesByOrder(stage.Checkboxs, stage.CheckboxMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Checkbox implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Group:
		tmp := GetStructInstancesByOrder(stage.Groups, stage.GroupMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Group implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Layout:
		tmp := GetStructInstancesByOrder(stage.Layouts, stage.LayoutMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Layout implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Slider:
		tmp := GetStructInstancesByOrder(stage.Sliders, stage.SliderMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Slider implements.
			res = append(res, any(v).(T))
		}
		return res

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]struct{}, order map[T]uint) (res []T) {

	orderedSet := []T{}
	for instance := range set {
		orderedSet = append(orderedSet, instance)
	}
	sort.Slice(orderedSet[:], func(i, j int) bool {
		instancei := orderedSet[i]
		instancej := orderedSet[j]
		i_order, oki := order[instancei]
		j_order, okj := order[instancej]
		if !oki || !okj {
			log.Fatalf("GetNamedStructInstances: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case
	case "Checkbox":
		res = GetNamedStructInstances(stage.Checkboxs, stage.CheckboxMap_Staged_Order)
	case "Group":
		res = GetNamedStructInstances(stage.Groups, stage.GroupMap_Staged_Order)
	case "Layout":
		res = GetNamedStructInstances(stage.Layouts, stage.LayoutMap_Staged_Order)
	case "Slider":
		res = GetNamedStructInstances(stage.Sliders, stage.SliderMap_Staged_Order)
	}

	return
}

type NamedStruct struct {
	name string
}

func (namedStruct *NamedStruct) GetName() string {
	return namedStruct.name
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
		Checkboxs:           make(map[*Checkbox]struct{}),
		Checkboxs_mapString: make(map[string]*Checkbox),

		Groups:           make(map[*Group]struct{}),
		Groups_mapString: make(map[string]*Group),

		Layouts:           make(map[*Layout]struct{}),
		Layouts_mapString: make(map[string]*Layout),

		Sliders:           make(map[*Slider]struct{}),
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Checkbox"},
			{name: "Group"},
			{name: "Layout"},
			{name: "Slider"},
		}, // end of insertion point
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

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

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

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()
	stage.ComputeDifference()
	stage.ComputeReference()

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().AddNotification(time.Now(), "Commit performed")
		stage.GetProbeIF().CommitNotificationTable()
	}
}

func (stage *Stage) ComputeInstancesNb() {
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
	stage.ComputeInstancesNb()
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
		stage.Checkboxs[checkbox] = struct{}{}
		stage.CheckboxMap_Staged_Order[checkbox] = stage.CheckboxOrder
		stage.CheckboxOrder++
	}
	stage.Checkboxs_mapString[checkbox.Name] = checkbox

	return checkbox
}

// StagePreserveOrder puts checkbox to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CheckboxOrder
// - update stage.CheckboxOrder accordingly
func (checkbox *Checkbox) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Checkboxs[checkbox]; !ok {
		stage.Checkboxs[checkbox] = struct{}{}

		if order > stage.CheckboxOrder {
			stage.CheckboxOrder = order
		}
		stage.CheckboxMap_Staged_Order[checkbox] = stage.CheckboxOrder
		stage.CheckboxOrder++
	}
	stage.Checkboxs_mapString[checkbox.Name] = checkbox
}

// Unstage removes checkbox off the model stage
func (checkbox *Checkbox) Unstage(stage *Stage) *Checkbox {
	delete(stage.Checkboxs, checkbox)
	delete(stage.CheckboxMap_Staged_Order, checkbox)
	delete(stage.Checkboxs_mapString, checkbox.Name)

	return checkbox
}

// UnstageVoid removes checkbox off the model stage
func (checkbox *Checkbox) UnstageVoid(stage *Stage) {
	delete(stage.Checkboxs, checkbox)
	delete(stage.CheckboxMap_Staged_Order, checkbox)
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

func (checkbox *Checkbox) StageVoid(stage *Stage) {
	checkbox.Stage(stage)
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

// for satisfaction of GongStruct interface
func (checkbox *Checkbox) SetName(name string) {
	checkbox.Name = name
}

// Stage puts group to the model stage
func (group *Group) Stage(stage *Stage) *Group {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = struct{}{}
		stage.GroupMap_Staged_Order[group] = stage.GroupOrder
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group

	return group
}

// StagePreserveOrder puts group to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupOrder
// - update stage.GroupOrder accordingly
func (group *Group) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = struct{}{}

		if order > stage.GroupOrder {
			stage.GroupOrder = order
		}
		stage.GroupMap_Staged_Order[group] = stage.GroupOrder
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	delete(stage.Groups, group)
	delete(stage.GroupMap_Staged_Order, group)
	delete(stage.Groups_mapString, group.Name)

	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	delete(stage.Groups, group)
	delete(stage.GroupMap_Staged_Order, group)
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

func (group *Group) StageVoid(stage *Stage) {
	group.Stage(stage)
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

// for satisfaction of GongStruct interface
func (group *Group) SetName(name string) {
	group.Name = name
}

// Stage puts layout to the model stage
func (layout *Layout) Stage(stage *Stage) *Layout {

	if _, ok := stage.Layouts[layout]; !ok {
		stage.Layouts[layout] = struct{}{}
		stage.LayoutMap_Staged_Order[layout] = stage.LayoutOrder
		stage.LayoutOrder++
	}
	stage.Layouts_mapString[layout.Name] = layout

	return layout
}

// StagePreserveOrder puts layout to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.LayoutOrder
// - update stage.LayoutOrder accordingly
func (layout *Layout) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Layouts[layout]; !ok {
		stage.Layouts[layout] = struct{}{}

		if order > stage.LayoutOrder {
			stage.LayoutOrder = order
		}
		stage.LayoutMap_Staged_Order[layout] = stage.LayoutOrder
		stage.LayoutOrder++
	}
	stage.Layouts_mapString[layout.Name] = layout
}

// Unstage removes layout off the model stage
func (layout *Layout) Unstage(stage *Stage) *Layout {
	delete(stage.Layouts, layout)
	delete(stage.LayoutMap_Staged_Order, layout)
	delete(stage.Layouts_mapString, layout.Name)

	return layout
}

// UnstageVoid removes layout off the model stage
func (layout *Layout) UnstageVoid(stage *Stage) {
	delete(stage.Layouts, layout)
	delete(stage.LayoutMap_Staged_Order, layout)
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

func (layout *Layout) StageVoid(stage *Stage) {
	layout.Stage(stage)
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

// for satisfaction of GongStruct interface
func (layout *Layout) SetName(name string) {
	layout.Name = name
}

// Stage puts slider to the model stage
func (slider *Slider) Stage(stage *Stage) *Slider {

	if _, ok := stage.Sliders[slider]; !ok {
		stage.Sliders[slider] = struct{}{}
		stage.SliderMap_Staged_Order[slider] = stage.SliderOrder
		stage.SliderOrder++
	}
	stage.Sliders_mapString[slider.Name] = slider

	return slider
}

// StagePreserveOrder puts slider to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SliderOrder
// - update stage.SliderOrder accordingly
func (slider *Slider) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Sliders[slider]; !ok {
		stage.Sliders[slider] = struct{}{}

		if order > stage.SliderOrder {
			stage.SliderOrder = order
		}
		stage.SliderMap_Staged_Order[slider] = stage.SliderOrder
		stage.SliderOrder++
	}
	stage.Sliders_mapString[slider.Name] = slider
}

// Unstage removes slider off the model stage
func (slider *Slider) Unstage(stage *Stage) *Slider {
	delete(stage.Sliders, slider)
	delete(stage.SliderMap_Staged_Order, slider)
	delete(stage.Sliders_mapString, slider.Name)

	return slider
}

// UnstageVoid removes slider off the model stage
func (slider *Slider) UnstageVoid(stage *Stage) {
	delete(stage.Sliders, slider)
	delete(stage.SliderMap_Staged_Order, slider)
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

func (slider *Slider) StageVoid(stage *Stage) {
	slider.Stage(stage)
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

// for satisfaction of GongStruct interface
func (slider *Slider) SetName(name string) {
	slider.Name = name
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
	stage.Checkboxs = make(map[*Checkbox]struct{})
	stage.Checkboxs_mapString = make(map[string]*Checkbox)
	stage.CheckboxMap_Staged_Order = make(map[*Checkbox]uint)
	stage.CheckboxOrder = 0

	stage.Groups = make(map[*Group]struct{})
	stage.Groups_mapString = make(map[string]*Group)
	stage.GroupMap_Staged_Order = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.Layouts = make(map[*Layout]struct{})
	stage.Layouts_mapString = make(map[string]*Layout)
	stage.LayoutMap_Staged_Order = make(map[*Layout]uint)
	stage.LayoutOrder = 0

	stage.Sliders = make(map[*Slider]struct{})
	stage.Sliders_mapString = make(map[string]*Slider)
	stage.SliderMap_Staged_Order = make(map[*Slider]uint)
	stage.SliderOrder = 0

	stage.ComputeReference()
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
type GongstructIF interface {
	GetName() string
	SetName(string)
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *ReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *ReverseField) GongstructIF
}
type PointerToGongstruct interface {
	GongstructIF
	comparable
}

func CompareGongstructByName[T PointerToGongstruct](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func SortGongstructSetByName[T PointerToGongstruct](set map[T]struct{}) (sortedSlice []T) {

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

// GongGetMap returns the map of staged Gonstruct instance by their name
// Can be usefull if names are unique
func GongGetMap[Type GongstructIF](stage *Stage) map[string]Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Checkbox:
		return any(stage.Checkboxs_mapString).(map[string]Type)
	case *Group:
		return any(stage.Groups_mapString).(map[string]Type)
	case *Layout:
		return any(stage.Layouts_mapString).(map[string]Type)
	case *Slider:
		return any(stage.Sliders_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Checkbox:
		return any(&stage.Checkboxs).(*map[*Type]struct{})
	case Group:
		return any(&stage.Groups).(*map[*Type]struct{})
	case Layout:
		return any(&stage.Layouts).(*map[*Type]struct{})
	case Slider:
		return any(&stage.Sliders).(*map[*Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Checkbox:
		return any(&stage.Checkboxs).(*map[Type]struct{})
	case *Group:
		return any(&stage.Groups).(*map[Type]struct{})
	case *Layout:
		return any(&stage.Layouts).(*map[Type]struct{})
	case *Slider:
		return any(&stage.Sliders).(*map[Type]struct{})
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

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
			res := make(map[*Slider][]*Group)
			for group := range stage.Groups {
				for _, slider_ := range group.Sliders {
					res[slider_] = append(res[slider_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Checkboxes":
			res := make(map[*Checkbox][]*Group)
			for group := range stage.Groups {
				for _, checkbox_ := range group.Checkboxes {
					res[checkbox_] = append(res[checkbox_], group)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Layout
	case Layout:
		switch fieldname {
		// insertion point for per direct association field
		case "Groups":
			res := make(map[*Group][]*Layout)
			for layout := range stage.Layouts {
				for _, group_ := range layout.Groups {
					res[group_] = append(res[group_], layout)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Slider
	case Slider:
		switch fieldname {
		// insertion point for per direct association field
		}
	}
	return nil
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {

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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type GongstructIF]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *Checkbox:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Checkboxes"
		res = append(res, rf)
	case *Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layout"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case *Layout:
		var rf ReverseField
		_ = rf
	case *Slider:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Sliders"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (checkbox *Checkbox) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ValueBool",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LabelForTrue",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "LabelForFalse",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (group *Group) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Percentage",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Sliders",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Slider",
		},
		{
			Name:                 "Checkboxes",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Checkbox",
		},
	}
	return
}

func (layout *Layout) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
	}
	return
}

func (slider *Slider) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsFloat64",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsInt",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MinInt",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MaxInt",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StepInt",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ValueInt",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MinFloat64",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MaxFloat64",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StepFloat64",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "ValueFloat64",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {

	var ret Type
	return ret.GongGetFieldHeaders()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeBasicKind       GongFieldValueType = "GongFieldValueTypeBasicKind"
	GongFieldValueTypePointer         GongFieldValueType = "GongFieldValueTypePointer"
	GongFieldValueTypeSliceOfPointers GongFieldValueType = "GongFieldValueTypeSliceOfPointers"
)

type GongFieldValue struct {
	GongFieldValueType
	valueString string
	valueInt    int
	valueFloat  float64
	valueBool   bool

	// in case of a pointer, the ID of the pointed element
	// in case of a slice of pointers, the IDs, separated by semi columbs
	ids string
}

type GongFieldHeader struct {
	Name string
	GongFieldValueType
	TargetGongstructName string
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

// insertion point for generic get gongstruct field value
func (checkbox *Checkbox) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = checkbox.Name
	case "ValueBool":
		res.valueString = fmt.Sprintf("%t", checkbox.ValueBool)
		res.valueBool = checkbox.ValueBool
		res.GongFieldValueType = GongFieldValueTypeBool
	case "LabelForTrue":
		res.valueString = checkbox.LabelForTrue
	case "LabelForFalse":
		res.valueString = checkbox.LabelForFalse
	}
	return
}
func (group *Group) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = group.Name
	case "Percentage":
		res.valueString = fmt.Sprintf("%f", group.Percentage)
		res.valueFloat = group.Percentage
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "Sliders":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Sliders {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "Checkboxes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Checkboxes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func (layout *Layout) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = layout.Name
	case "Groups":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layout.Groups {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	}
	return
}
func (slider *Slider) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = slider.Name
	case "IsFloat64":
		res.valueString = fmt.Sprintf("%t", slider.IsFloat64)
		res.valueBool = slider.IsFloat64
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsInt":
		res.valueString = fmt.Sprintf("%t", slider.IsInt)
		res.valueBool = slider.IsInt
		res.GongFieldValueType = GongFieldValueTypeBool
	case "MinInt":
		res.valueString = fmt.Sprintf("%d", slider.MinInt)
		res.valueInt = slider.MinInt
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MaxInt":
		res.valueString = fmt.Sprintf("%d", slider.MaxInt)
		res.valueInt = slider.MaxInt
		res.GongFieldValueType = GongFieldValueTypeInt
	case "StepInt":
		res.valueString = fmt.Sprintf("%d", slider.StepInt)
		res.valueInt = slider.StepInt
		res.GongFieldValueType = GongFieldValueTypeInt
	case "ValueInt":
		res.valueString = fmt.Sprintf("%d", slider.ValueInt)
		res.valueInt = slider.ValueInt
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MinFloat64":
		res.valueString = fmt.Sprintf("%f", slider.MinFloat64)
		res.valueFloat = slider.MinFloat64
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "MaxFloat64":
		res.valueString = fmt.Sprintf("%f", slider.MaxFloat64)
		res.valueFloat = slider.MaxFloat64
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "StepFloat64":
		res.valueString = fmt.Sprintf("%f", slider.StepFloat64)
		res.valueFloat = slider.StepFloat64
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ValueFloat64":
		res.valueString = fmt.Sprintf("%f", slider.ValueFloat64)
		res.valueFloat = slider.ValueFloat64
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (checkbox *Checkbox) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		checkbox.Name = value.GetValueString()
	case "ValueBool":
		checkbox.ValueBool = value.GetValueBool()
	case "LabelForTrue":
		checkbox.LabelForTrue = value.GetValueString()
	case "LabelForFalse":
		checkbox.LabelForFalse = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (group *Group) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		group.Name = value.GetValueString()
	case "Percentage":
		group.Percentage = value.GetValueFloat()
	case "Sliders":
		group.Sliders = make([]*Slider, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Sliders {
					if stage.SliderMap_Staged_Order[__instance__] == uint(id) {
						group.Sliders = append(group.Sliders, __instance__)
						break
					}
				}
			}
		}
	case "Checkboxes":
		group.Checkboxes = make([]*Checkbox, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Checkboxs {
					if stage.CheckboxMap_Staged_Order[__instance__] == uint(id) {
						group.Checkboxes = append(group.Checkboxes, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (layout *Layout) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		layout.Name = value.GetValueString()
	case "Groups":
		layout.Groups = make([]*Group, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Groups {
					if stage.GroupMap_Staged_Order[__instance__] == uint(id) {
						layout.Groups = append(layout.Groups, __instance__)
						break
					}
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (slider *Slider) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		slider.Name = value.GetValueString()
	case "IsFloat64":
		slider.IsFloat64 = value.GetValueBool()
	case "IsInt":
		slider.IsInt = value.GetValueBool()
	case "MinInt":
		slider.MinInt = int(value.GetValueInt())
	case "MaxInt":
		slider.MaxInt = int(value.GetValueInt())
	case "StepInt":
		slider.StepInt = int(value.GetValueInt())
	case "ValueInt":
		slider.ValueInt = int(value.GetValueInt())
	case "MinFloat64":
		slider.MinFloat64 = value.GetValueFloat()
	case "MaxFloat64":
		slider.MaxFloat64 = value.GetValueFloat()
	case "StepFloat64":
		slider.StepFloat64 = value.GetValueFloat()
	case "ValueFloat64":
		slider.ValueFloat64 = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (checkbox *Checkbox) GongGetGongstructName() string {
	return "Checkbox"
}

func (group *Group) GongGetGongstructName() string {
	return "Group"
}

func (layout *Layout) GongGetGongstructName() string {
	return "Layout"
}

func (slider *Slider) GongGetGongstructName() string {
	return "Slider"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Checkboxs_mapString = make(map[string]*Checkbox)
	for checkbox := range stage.Checkboxs {
		stage.Checkboxs_mapString[checkbox.Name] = checkbox
	}

	stage.Groups_mapString = make(map[string]*Group)
	for group := range stage.Groups {
		stage.Groups_mapString[group.Name] = group
	}

	stage.Layouts_mapString = make(map[string]*Layout)
	for layout := range stage.Layouts {
		stage.Layouts_mapString[layout.Name] = layout
	}

	stage.Sliders_mapString = make(map[string]*Slider)
	for slider := range stage.Sliders {
		stage.Sliders_mapString[slider.Name] = slider
	}

}

// Last line of the template
