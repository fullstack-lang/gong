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

	cursor_go "github.com/fullstack-lang/gong/lib/cursor/go"
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
	Cursors           map[*Cursor]any
	Cursors_mapString map[string]*Cursor

	// insertion point for slice of pointers maps
	OnAfterCursorCreateCallback OnAfterCreateInterface[Cursor]
	OnAfterCursorUpdateCallback OnAfterUpdateInterface[Cursor]
	OnAfterCursorDeleteCallback OnAfterDeleteInterface[Cursor]
	OnAfterCursorReadCallback   OnAfterReadInterface[Cursor]

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
	CursorOrder            uint
	CursorMap_Staged_Order map[*Cursor]uint

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
	case *Cursor:
		tmp := GetStructInstancesByOrder(stage.Cursors, stage.CursorMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Cursor implements.
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
	case "Cursor":
		res = GetNamedStructInstances(stage.Cursors, stage.CursorMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/cursor/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return cursor_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return cursor_go.GoDiagramsDir
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
	CommitCursor(cursor *Cursor)
	CheckoutCursor(cursor *Cursor)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Cursors:           make(map[*Cursor]any),
		Cursors_mapString: make(map[string]*Cursor),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		CursorMap_Staged_Order: make(map[*Cursor]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Cursor"},
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
	case *Cursor:
		return stage.CursorMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Cursor:
		return stage.CursorMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["Cursor"] = len(stage.Cursors)
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
// Stage puts cursor to the model stage
func (cursor *Cursor) Stage(stage *Stage) *Cursor {

	if _, ok := stage.Cursors[cursor]; !ok {
		stage.Cursors[cursor] = __member
		stage.CursorMap_Staged_Order[cursor] = stage.CursorOrder
		stage.CursorOrder++
		stage.new[cursor] = struct{}{}
		delete(stage.deleted, cursor)
	} else {
		if _, ok := stage.new[cursor]; !ok {
			stage.modified[cursor] = struct{}{}
		}
	}
	stage.Cursors_mapString[cursor.Name] = cursor

	return cursor
}

// Unstage removes cursor off the model stage
func (cursor *Cursor) Unstage(stage *Stage) *Cursor {
	delete(stage.Cursors, cursor)
	delete(stage.Cursors_mapString, cursor.Name)

	if _, ok := stage.reference[cursor]; ok {
		stage.deleted[cursor] = struct{}{}
	} else {
		delete(stage.new, cursor)
	}
	return cursor
}

// UnstageVoid removes cursor off the model stage
func (cursor *Cursor) UnstageVoid(stage *Stage) {
	delete(stage.Cursors, cursor)
	delete(stage.Cursors_mapString, cursor.Name)
}

// commit cursor to the back repo (if it is already staged)
func (cursor *Cursor) Commit(stage *Stage) *Cursor {
	if _, ok := stage.Cursors[cursor]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCursor(cursor)
		}
	}
	return cursor
}

func (cursor *Cursor) CommitVoid(stage *Stage) {
	cursor.Commit(stage)
}

func (cursor *Cursor) StageVoid(stage *Stage) {
	cursor.Stage(stage)
}

// Checkout cursor to the back repo (if it is already staged)
func (cursor *Cursor) Checkout(stage *Stage) *Cursor {
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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCursor(Cursor *Cursor)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCursor(Cursor *Cursor)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Cursors = make(map[*Cursor]any)
	stage.Cursors_mapString = make(map[string]*Cursor)
	stage.CursorMap_Staged_Order = make(map[*Cursor]uint)
	stage.CursorOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Cursors = nil
	stage.Cursors_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for cursor := range stage.Cursors {
		cursor.Unstage(stage)
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
	case map[*Cursor]any:
		return any(&stage.Cursors).(*Type)
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
	case map[string]*Cursor:
		return any(&stage.Cursors_mapString).(*Type)
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
	case Cursor:
		return any(&stage.Cursors).(*map[*Type]any)
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
	case *Cursor:
		return any(&stage.Cursors).(*map[Type]any)
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
	case Cursor:
		return any(&stage.Cursors_mapString).(*map[string]*Type)
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
	case Cursor:
		return any(&Cursor{
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
	// reverse maps of direct associations of Cursor
	case Cursor:
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
	// reverse maps of direct associations of Cursor
	case Cursor:
		switch fieldname {
		// insertion point for per direct association field
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
	case *Cursor:
		res = "Cursor"
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
	case *Cursor:
		var rf ReverseField
		_ = rf
	}
	return
}

// insertion point for get fields header method
func (cursor *Cursor) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StartX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "EndX",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y1",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y2",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "DurationSeconds",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Color",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "FillOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Stroke",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeOpacity",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeWidth",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArray",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "StrokeDashArrayWhenSelected",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Transform",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "IsPlaying",
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
func (cursor *Cursor) GongGetFieldValueString(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = cursor.Name
		case "StartX":
			res.valueString = fmt.Sprintf("%f", cursor.StartX)
			res.valueFloat = cursor.StartX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "EndX":
			res.valueString = fmt.Sprintf("%f", cursor.EndX)
			res.valueFloat = cursor.EndX
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y1":
			res.valueString = fmt.Sprintf("%f", cursor.Y1)
			res.valueFloat = cursor.Y1
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Y2":
			res.valueString = fmt.Sprintf("%f", cursor.Y2)
			res.valueFloat = cursor.Y2
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "DurationSeconds":
			res.valueString = fmt.Sprintf("%f", cursor.DurationSeconds)
			res.valueFloat = cursor.DurationSeconds
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Color":
			res.valueString = cursor.Color
		case "FillOpacity":
			res.valueString = fmt.Sprintf("%f", cursor.FillOpacity)
			res.valueFloat = cursor.FillOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "Stroke":
			res.valueString = cursor.Stroke
		case "StrokeOpacity":
			res.valueString = fmt.Sprintf("%f", cursor.StrokeOpacity)
			res.valueFloat = cursor.StrokeOpacity
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeWidth":
			res.valueString = fmt.Sprintf("%f", cursor.StrokeWidth)
			res.valueFloat = cursor.StrokeWidth
			res.GongFieldValueType = GongFieldValueTypeFloat
		case "StrokeDashArray":
			res.valueString = cursor.StrokeDashArray
		case "StrokeDashArrayWhenSelected":
			res.valueString = cursor.StrokeDashArrayWhenSelected
		case "Transform":
			res.valueString = cursor.Transform
		case "IsPlaying":
			res.valueString = fmt.Sprintf("%t", cursor.IsPlaying)
			res.valueBool = cursor.IsPlaying
			res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}


func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValueString(fieldName, stage)
	return
}

// Last line of the template
