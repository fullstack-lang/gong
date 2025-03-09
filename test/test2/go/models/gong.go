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
	path string

	// insertion point for definition of arrays registering instances
	As           map[*A]any
	As_mapString map[string]*A

	// insertion point for slice of pointers maps
	A_Bs_reverseMap map[*B]*A

	OnAfterACreateCallback OnAfterCreateInterface[A]
	OnAfterAUpdateCallback OnAfterUpdateInterface[A]
	OnAfterADeleteCallback OnAfterDeleteInterface[A]
	OnAfterAReadCallback   OnAfterReadInterface[A]

	Bs           map[*B]any
	Bs_mapString map[string]*B

	// insertion point for slice of pointers maps
	OnAfterBCreateCallback OnAfterCreateInterface[B]
	OnAfterBUpdateCallback OnAfterUpdateInterface[B]
	OnAfterBDeleteCallback OnAfterDeleteInterface[B]
	OnAfterBReadCallback   OnAfterReadInterface[B]

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
	return "github.com/fullstack-lang/gong/test/test2/go/models"
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
	CommitA(a *A)
	CheckoutA(a *A)
	CommitB(b *B)
	CheckoutB(b *B)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		As:           make(map[*A]any),
		As_mapString: make(map[string]*A),

		Bs:           make(map[*B]any),
		Bs_mapString: make(map[string]*B),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		Map_Staged_Order: make(map[any]uint),
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
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
	stage.Map_GongStructName_InstancesNb["A"] = len(stage.As)
	stage.Map_GongStructName_InstancesNb["B"] = len(stage.Bs)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["A"] = len(stage.As)
	stage.Map_GongStructName_InstancesNb["B"] = len(stage.Bs)

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
// Stage puts a to the model stage
func (a *A) Stage(stage *StageStruct) *A {

	if _, ok := stage.As[a]; !ok {
		stage.As[a] = __member
		stage.Map_Staged_Order[a] = stage.Order
		stage.Order++
	}
	stage.As_mapString[a.Name] = a

	return a
}

// Unstage removes a off the model stage
func (a *A) Unstage(stage *StageStruct) *A {
	delete(stage.As, a)
	delete(stage.As_mapString, a.Name)
	return a
}

// UnstageVoid removes a off the model stage
func (a *A) UnstageVoid(stage *StageStruct) {
	delete(stage.As, a)
	delete(stage.As_mapString, a.Name)
}

// commit a to the back repo (if it is already staged)
func (a *A) Commit(stage *StageStruct) *A {
	if _, ok := stage.As[a]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA(a)
		}
	}
	return a
}

func (a *A) CommitVoid(stage *StageStruct) {
	a.Commit(stage)
}

// Checkout a to the back repo (if it is already staged)
func (a *A) Checkout(stage *StageStruct) *A {
	if _, ok := stage.As[a]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutA(a)
		}
	}
	return a
}

// for satisfaction of GongStruct interface
func (a *A) GetName() (res string) {
	return a.Name
}

// Stage puts b to the model stage
func (b *B) Stage(stage *StageStruct) *B {

	if _, ok := stage.Bs[b]; !ok {
		stage.Bs[b] = __member
		stage.Map_Staged_Order[b] = stage.Order
		stage.Order++
	}
	stage.Bs_mapString[b.Name] = b

	return b
}

// Unstage removes b off the model stage
func (b *B) Unstage(stage *StageStruct) *B {
	delete(stage.Bs, b)
	delete(stage.Bs_mapString, b.Name)
	return b
}

// UnstageVoid removes b off the model stage
func (b *B) UnstageVoid(stage *StageStruct) {
	delete(stage.Bs, b)
	delete(stage.Bs_mapString, b.Name)
}

// commit b to the back repo (if it is already staged)
func (b *B) Commit(stage *StageStruct) *B {
	if _, ok := stage.Bs[b]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitB(b)
		}
	}
	return b
}

func (b *B) CommitVoid(stage *StageStruct) {
	b.Commit(stage)
}

// Checkout b to the back repo (if it is already staged)
func (b *B) Checkout(stage *StageStruct) *B {
	if _, ok := stage.Bs[b]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutB(b)
		}
	}
	return b
}

// for satisfaction of GongStruct interface
func (b *B) GetName() (res string) {
	return b.Name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMA(A *A)
	CreateORMB(B *B)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMA(A *A)
	DeleteORMB(B *B)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.As = make(map[*A]any)
	stage.As_mapString = make(map[string]*A)

	stage.Bs = make(map[*B]any)
	stage.Bs_mapString = make(map[string]*B)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
	stage.As = nil
	stage.As_mapString = nil

	stage.Bs = nil
	stage.Bs_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
	for a := range stage.As {
		a.Unstage(stage)
	}

	for b := range stage.Bs {
		b.Unstage(stage)
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
	case map[*A]any:
		return any(&stage.As).(*Type)
	case map[*B]any:
		return any(&stage.Bs).(*Type)
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
	case map[string]*A:
		return any(&stage.As_mapString).(*Type)
	case map[string]*B:
		return any(&stage.Bs_mapString).(*Type)
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
	case A:
		return any(&stage.As).(*map[*Type]any)
	case B:
		return any(&stage.Bs).(*map[*Type]any)
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
	case *A:
		return any(&stage.As).(*map[Type]any)
	case *B:
		return any(&stage.Bs).(*map[Type]any)
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
	case A:
		return any(&stage.As_mapString).(*map[string]*Type)
	case B:
		return any(&stage.Bs_mapString).(*map[string]*Type)
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
	case A:
		return any(&A{
			// Initialisation of associations
			// field is initialized with an instance of B with the name of the field
			B: &B{Name: "B"},
			// field is initialized with an instance of B with the name of the field
			Bs: []*B{{Name: "Bs"}},
		}).(*Type)
	case B:
		return any(&B{
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of A
	case A:
		switch fieldname {
		// insertion point for per direct association field
		case "B":
			res := make(map[*B][]*A)
			for a := range stage.As {
				if a.B != nil {
					b_ := a.B
					var as []*A
					_, ok := res[b_]
					if ok {
						as = res[b_]
					} else {
						as = make([]*A, 0)
					}
					as = append(as, a)
					res[b_] = as
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of B
	case B:
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
	// reverse maps of direct associations of A
	case A:
		switch fieldname {
		// insertion point for per direct association field
		case "Bs":
			res := make(map[*B]*A)
			for a := range stage.As {
				for _, b_ := range a.Bs {
					res[b_] = a
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of B
	case B:
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
	case A:
		res = "A"
	case B:
		res = "B"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *A:
		res = "A"
	case *B:
		res = "B"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case A:
		res = []string{"Name", "NumberField", "B", "Bs"}
	case B:
		res = []string{"Name"}
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
	case A:
		var rf ReverseField
		_ = rf
	case B:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "A"
		rf.Fieldname = "Bs"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *A:
		res = []string{"Name", "NumberField", "B", "Bs"}
	case *B:
		res = []string{"Name"}
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
	case *A:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "NumberField":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NumberField)
			res.valueInt = inferedInstance.NumberField
			res.GongFieldValueType = GongFieldValueTypeInt
		case "B":
			if inferedInstance.B != nil {
				res.valueString = inferedInstance.B.Name
			}
		case "Bs":
			for idx, __instance__ := range inferedInstance.Bs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *B:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue(instance any, fieldName string) (res GongFieldValue) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case A:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "NumberField":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NumberField)
			res.valueInt = inferedInstance.NumberField
			res.GongFieldValueType = GongFieldValueTypeInt
		case "B":
			if inferedInstance.B != nil {
				res.valueString = inferedInstance.B.Name
			}
		case "Bs":
			for idx, __instance__ := range inferedInstance.Bs {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case B:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
