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
	Buttons           map[*Button]any
	Buttons_mapString map[string]*Button

	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback OnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback OnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback OnAfterDeleteInterface[Button]
	OnAfterButtonReadCallback   OnAfterReadInterface[Button]

	Groups           map[*Group]any
	Groups_mapString map[string]*Group

	// insertion point for slice of pointers maps
	Group_Buttons_reverseMap map[*Button]*Group

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
	return "github.com/fullstack-lang/gong/lib/button/go/models"
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
	CommitButton(button *Button)
	CheckoutButton(button *Button)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitLayout(layout *Layout)
	CheckoutLayout(layout *Layout)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Buttons:           make(map[*Button]any),
		Buttons_mapString: make(map[string]*Button),

		Groups:           make(map[*Group]any),
		Groups_mapString: make(map[string]*Group),

		Layouts:           make(map[*Layout]any),
		Layouts_mapString: make(map[string]*Layout),

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
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Layout"] = len(stage.Layouts)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Layout"] = len(stage.Layouts)

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

// Stage puts group to the model stage
func (group *Group) Stage(stage *StageStruct) *Group {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = __member
		stage.Map_Staged_Order[group] = stage.Order
		stage.Order++
	}
	stage.Groups_mapString[group.Name] = group

	return group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *StageStruct) *Group {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *StageStruct) {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)
}

// commit group to the back repo (if it is already staged)
func (group *Group) Commit(stage *StageStruct) *Group {
	if _, ok := stage.Groups[group]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroup(group)
		}
	}
	return group
}

func (group *Group) CommitVoid(stage *StageStruct) {
	group.Commit(stage)
}

// Checkout group to the back repo (if it is already staged)
func (group *Group) Checkout(stage *StageStruct) *Group {
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
func (layout *Layout) Stage(stage *StageStruct) *Layout {

	if _, ok := stage.Layouts[layout]; !ok {
		stage.Layouts[layout] = __member
		stage.Map_Staged_Order[layout] = stage.Order
		stage.Order++
	}
	stage.Layouts_mapString[layout.Name] = layout

	return layout
}

// Unstage removes layout off the model stage
func (layout *Layout) Unstage(stage *StageStruct) *Layout {
	delete(stage.Layouts, layout)
	delete(stage.Layouts_mapString, layout.Name)
	return layout
}

// UnstageVoid removes layout off the model stage
func (layout *Layout) UnstageVoid(stage *StageStruct) {
	delete(stage.Layouts, layout)
	delete(stage.Layouts_mapString, layout.Name)
}

// commit layout to the back repo (if it is already staged)
func (layout *Layout) Commit(stage *StageStruct) *Layout {
	if _, ok := stage.Layouts[layout]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitLayout(layout)
		}
	}
	return layout
}

func (layout *Layout) CommitVoid(stage *StageStruct) {
	layout.Commit(stage)
}

// Checkout layout to the back repo (if it is already staged)
func (layout *Layout) Checkout(stage *StageStruct) *Layout {
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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMButton(Button *Button)
	CreateORMGroup(Group *Group)
	CreateORMLayout(Layout *Layout)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMButton(Button *Button)
	DeleteORMGroup(Group *Group)
	DeleteORMLayout(Layout *Layout)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Buttons = make(map[*Button]any)
	stage.Buttons_mapString = make(map[string]*Button)

	stage.Groups = make(map[*Group]any)
	stage.Groups_mapString = make(map[string]*Group)

	stage.Layouts = make(map[*Layout]any)
	stage.Layouts_mapString = make(map[string]*Layout)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.Buttons = nil
	stage.Buttons_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.Layouts = nil
	stage.Layouts_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for button := range stage.Buttons {
		button.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
	}

	for layout := range stage.Layouts {
		layout.Unstage(stage)
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
	case map[*Button]any:
		return any(&stage.Buttons).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
	case map[*Layout]any:
		return any(&stage.Layouts).(*Type)
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
	case map[string]*Button:
		return any(&stage.Buttons_mapString).(*Type)
	case map[string]*Group:
		return any(&stage.Groups_mapString).(*Type)
	case map[string]*Layout:
		return any(&stage.Layouts_mapString).(*Type)
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
	case Button:
		return any(&stage.Buttons).(*map[*Type]any)
	case Group:
		return any(&stage.Groups).(*map[*Type]any)
	case Layout:
		return any(&stage.Layouts).(*map[*Type]any)
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
	case *Button:
		return any(&stage.Buttons).(*map[Type]any)
	case *Group:
		return any(&stage.Groups).(*map[Type]any)
	case *Layout:
		return any(&stage.Layouts).(*map[Type]any)
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
	case Button:
		return any(&stage.Buttons_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
	case Layout:
		return any(&stage.Layouts_mapString).(*map[string]*Type)
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
	case Button:
		return any(&Button{
			// Initialisation of associations
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of Button with the name of the field
			Buttons: []*Button{{Name: "Buttons"}},
		}).(*Type)
	case Layout:
		return any(&Layout{
			// Initialisation of associations
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
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
	// reverse maps of direct associations of Button
	case Button:
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
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		case "Buttons":
			res := make(map[*Button]*Group)
			for group := range stage.Groups {
				for _, button_ := range group.Buttons {
					res[button_] = group
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
	}
	return nil
}

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Button:
		res = "Button"
	case Group:
		res = "Group"
	case Layout:
		res = "Layout"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Button:
		res = "Button"
	case *Group:
		res = "Group"
	case *Layout:
		res = "Layout"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Button:
		res = []string{"Name", "Label", "Icon"}
	case Group:
		res = []string{"Name", "Percentage", "Buttons"}
	case Layout:
		res = []string{"Name", "Groups"}
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
	case Button:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Buttons"
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
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *Button:
		res = []string{"Name", "Label", "Icon"}
	case *Group:
		res = []string{"Name", "Percentage", "Buttons"}
	case *Layout:
		res = []string{"Name", "Groups"}
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
	case *Button:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Label":
			res.valueString = inferedInstance.Label
		case "Icon":
			res.valueString = inferedInstance.Icon
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
		case "Buttons":
			for idx, __instance__ := range inferedInstance.Buttons {
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
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Button:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "Label":
			res.valueString = inferedInstance.Label
		case "Icon":
			res.valueString = inferedInstance.Icon
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
		case "Buttons":
			for idx, __instance__ := range inferedInstance.Buttons {
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
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
