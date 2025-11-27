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
	"time"

	button_go "github.com/fullstack-lang/gong/lib/button/go"
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

const ProbeTreeSidebarSuffix = ":sidebar of the probe"
const ProbeTableSuffix = ":table of the probe"
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

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeSplitSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

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
	name               string
	commitId           uint // commitId is updated at each commit
	commitTimeStamp    time.Time
	contentWhenParsed  string
	commitIdWhenParsed uint
	generatesDiff      bool

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
	// insertion point for order fields declaration
	ButtonOrder            uint
	ButtonMap_Staged_Order map[*Button]uint

	GroupOrder            uint
	GroupMap_Staged_Order map[*Group]uint

	LayoutOrder            uint
	LayoutMap_Staged_Order map[*Layout]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
	modified  map[GongstructIF]struct{}
	new       map[GongstructIF]struct{}
	deleted   map[GongstructIF]struct{}
}

func (stage *Stage) GetCommitId() uint {
	return stage.commitId
}

func (stage *Stage) GetCommitTS() time.Time {
	return stage.commitTimeStamp
}

func (stage *Stage) SetGeneratesDiff(generatesDiff bool) {
	stage.generatesDiff = generatesDiff
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
}

func (stage *Stage) GetReference() map[GongstructIF]GongstructIF {
	return stage.reference
}

func (stage *Stage) GetModified() map[GongstructIF]struct{} {
	return stage.modified
}

func (stage *Stage) GetNew() map[GongstructIF]struct{} {
	return stage.new
}

func (stage *Stage) GetDeleted() map[GongstructIF]struct{} {
	return stage.deleted
}

func GetNamedStructInstances[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []string) {

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
	case *Button:
		tmp := GetStructInstancesByOrder(stage.Buttons, stage.ButtonMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Button implements.
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

	}
	return
}

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []T) {

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
	case "Button":
		res = GetNamedStructInstances(stage.Buttons, stage.ButtonMap_Staged_Order)
	case "Group":
		res = GetNamedStructInstances(stage.Groups, stage.GroupMap_Staged_Order)
	case "Layout":
		res = GetNamedStructInstances(stage.Layouts, stage.LayoutMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/button/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return button_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return button_go.GoDiagramsDir
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
	CommitButton(button *Button)
	CheckoutButton(button *Button)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitLayout(layout *Layout)
	CheckoutLayout(layout *Layout)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
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

		// insertion point for order map initialisations
		ButtonMap_Staged_Order: make(map[*Button]uint),

		GroupMap_Staged_Order: make(map[*Group]uint),

		LayoutMap_Staged_Order: make(map[*Layout]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Button"},
			{name: "Group"},
			{name: "Layout"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
		new:       make(map[GongstructIF]struct{}),
		modified:  make(map[GongstructIF]struct{}),
		deleted:   make(map[GongstructIF]struct{}),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Button:
		return stage.ButtonMap_Staged_Order[instance]
	case *Group:
		return stage.GroupMap_Staged_Order[instance]
	case *Layout:
		return stage.LayoutMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Button:
		return stage.ButtonMap_Staged_Order[instance]
	case *Group:
		return stage.GroupMap_Staged_Order[instance]
	case *Layout:
		return stage.LayoutMap_Staged_Order[instance]
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
	stage.commitId++
	stage.commitTimeStamp = time.Now()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()
	stage.ComputeReference()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["Layout"] = len(stage.Layouts)
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
// Stage puts button to the model stage
func (button *Button) Stage(stage *Stage) *Button {

	if _, ok := stage.Buttons[button]; !ok {
		stage.Buttons[button] = __member
		stage.ButtonMap_Staged_Order[button] = stage.ButtonOrder
		stage.ButtonOrder++
		stage.new[button] = struct{}{}
		delete(stage.deleted, button)
	} else {
		if _, ok := stage.new[button]; !ok {
			stage.modified[button] = struct{}{}
		}
	}
	stage.Buttons_mapString[button.Name] = button

	return button
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	delete(stage.Buttons, button)
	delete(stage.Buttons_mapString, button.Name)

	if _, ok := stage.reference[button]; ok {
		stage.deleted[button] = struct{}{}
	} else {
		delete(stage.new, button)
	}
	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *Stage) {
	delete(stage.Buttons, button)
	delete(stage.Buttons_mapString, button.Name)
}

// commit button to the back repo (if it is already staged)
func (button *Button) Commit(stage *Stage) *Button {
	if _, ok := stage.Buttons[button]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitButton(button)
		}
	}
	return button
}

func (button *Button) CommitVoid(stage *Stage) {
	button.Commit(stage)
}

func (button *Button) StageVoid(stage *Stage) {
	button.Stage(stage)
}

// Checkout button to the back repo (if it is already staged)
func (button *Button) Checkout(stage *Stage) *Button {
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
func (group *Group) Stage(stage *Stage) *Group {

	if _, ok := stage.Groups[group]; !ok {
		stage.Groups[group] = __member
		stage.GroupMap_Staged_Order[group] = stage.GroupOrder
		stage.GroupOrder++
		stage.new[group] = struct{}{}
		delete(stage.deleted, group)
	} else {
		if _, ok := stage.new[group]; !ok {
			stage.modified[group] = struct{}{}
		}
	}
	stage.Groups_mapString[group.Name] = group

	return group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	delete(stage.Groups, group)
	delete(stage.Groups_mapString, group.Name)

	if _, ok := stage.reference[group]; ok {
		stage.deleted[group] = struct{}{}
	} else {
		delete(stage.new, group)
	}
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

// Stage puts layout to the model stage
func (layout *Layout) Stage(stage *Stage) *Layout {

	if _, ok := stage.Layouts[layout]; !ok {
		stage.Layouts[layout] = __member
		stage.LayoutMap_Staged_Order[layout] = stage.LayoutOrder
		stage.LayoutOrder++
		stage.new[layout] = struct{}{}
		delete(stage.deleted, layout)
	} else {
		if _, ok := stage.new[layout]; !ok {
			stage.modified[layout] = struct{}{}
		}
	}
	stage.Layouts_mapString[layout.Name] = layout

	return layout
}

// Unstage removes layout off the model stage
func (layout *Layout) Unstage(stage *Stage) *Layout {
	delete(stage.Layouts, layout)
	delete(stage.Layouts_mapString, layout.Name)

	if _, ok := stage.reference[layout]; ok {
		stage.deleted[layout] = struct{}{}
	} else {
		delete(stage.new, layout)
	}
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

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Buttons = make(map[*Button]any)
	stage.Buttons_mapString = make(map[string]*Button)
	stage.ButtonMap_Staged_Order = make(map[*Button]uint)
	stage.ButtonOrder = 0

	stage.Groups = make(map[*Group]any)
	stage.Groups_mapString = make(map[string]*Group)
	stage.GroupMap_Staged_Order = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.Layouts = make(map[*Layout]any)
	stage.Layouts_mapString = make(map[string]*Layout)
	stage.LayoutMap_Staged_Order = make(map[*Layout]uint)
	stage.LayoutOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Buttons = nil
	stage.Buttons_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.Layouts = nil
	stage.Layouts_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
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
type GongstructIF interface {
	GetName() string
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage)
	GongGetFieldValueString(fieldName string, stage *Stage) GongFieldValue
	GongCopy() GongstructIF
}
type PointerToGongstruct interface {
	GongstructIF
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
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
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
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]any {
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
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *Stage) *map[Type]any {
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
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

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
			res := make(map[*Button][]*Group)
			for group := range stage.Groups {
				for _, button_ := range group.Buttons {
					res[button_] = append(res[button_], group)
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
	}
	return nil
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type PointerToGongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *Button:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Buttons"
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
	}
	return
}

// insertion point for get fields header method
func (button *Button) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Label",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Icon",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsDisabled",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MatButtonType",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MatButtonAppearance",
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
			Name:               "Buttons",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
		},
		{
			Name:               "NbColumns",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			Name:               "Groups",
			GongFieldValueType: GongFieldValueTypeSliceOfPointers,
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
	GongFieldValueType
	Name string
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
func (button *Button) GongGetFieldValueString(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = button.Name
		case "Label":
			res.valueString = button.Label
		case "Icon":
			res.valueString = button.Icon
		case "IsDisabled":
			res.valueString = fmt.Sprintf("%t", button.IsDisabled)
			res.valueBool = button.IsDisabled
			res.GongFieldValueType = GongFieldValueTypeBool
		case "Color":
			enum := button.Color
			res.valueString = enum.ToCodeString()
		case "MatButtonType":
			enum := button.MatButtonType
			res.valueString = enum.ToCodeString()
		case "MatButtonAppearance":
			enum := button.MatButtonAppearance
			res.valueString = enum.ToCodeString()
	}
	return
}
func (group *Group) GongGetFieldValueString(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = group.Name
		case "Percentage":
			res.valueString = fmt.Sprintf("%f", group.Percentage)
			res.valueFloat = group.Percentage
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Buttons":
			res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
			for idx, __instance__ := range group.Buttons {
				if idx > 0 {
					res.valueString += "\n"
					res.ids += ";"
				}
				res.valueString += __instance__.Name
				res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
			}
		case "NbColumns":
			res.valueString = fmt.Sprintf("%d", group.NbColumns)
			res.valueInt = group.NbColumns
			res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}
func (layout *Layout) GongGetFieldValueString(fieldName string, stage *Stage) (res GongFieldValue) {
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


func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValueString(fieldName, stage)
	return
}

// Last line of the template
