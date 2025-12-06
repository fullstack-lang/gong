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

	test2_go "github.com/fullstack-lang/gong/test/test2/go"
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
	As           map[*A]struct{}
	As_mapString map[string]*A

	// insertion point for slice of pointers maps
	A_Bs_reverseMap map[*B]*A

	OnAfterACreateCallback OnAfterCreateInterface[A]
	OnAfterAUpdateCallback OnAfterUpdateInterface[A]
	OnAfterADeleteCallback OnAfterDeleteInterface[A]
	OnAfterAReadCallback   OnAfterReadInterface[A]

	Bs           map[*B]struct{}
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
	// insertion point for order fields declaration
	AOrder            uint
	AMap_Staged_Order map[*A]uint

	BOrder            uint
	BMap_Staged_Order map[*B]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
	modified  map[GongstructIF]struct{}
	new       map[GongstructIF]struct{}
	deleted   map[GongstructIF]struct{}
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
	case *A:
		tmp := GetStructInstancesByOrder(stage.As, stage.AMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *A implements.
			res = append(res, any(v).(T))
		}
		return res
	case *B:
		tmp := GetStructInstancesByOrder(stage.Bs, stage.BMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *B implements.
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
	case "A":
		res = GetNamedStructInstances(stage.As, stage.AMap_Staged_Order)
	case "B":
		res = GetNamedStructInstances(stage.Bs, stage.BMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/test/test2/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return test2_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return test2_go.GoDiagramsDir
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
	CommitA(a *A)
	CheckoutA(a *A)
	CommitB(b *B)
	CheckoutB(b *B)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		As:           make(map[*A]struct{}),
		As_mapString: make(map[string]*A),

		Bs:           make(map[*B]struct{}),
		Bs_mapString: make(map[string]*B),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		AMap_Staged_Order: make(map[*A]uint),

		BMap_Staged_Order: make(map[*B]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "A"},
			{name: "B"},
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
	case *A:
		return stage.AMap_Staged_Order[instance]
	case *B:
		return stage.BMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *A:
		return stage.AMap_Staged_Order[instance]
	case *B:
		return stage.BMap_Staged_Order[instance]
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
	stage.ComputeReference()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["A"] = len(stage.As)
	stage.Map_GongStructName_InstancesNb["B"] = len(stage.Bs)
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
// Stage puts a to the model stage
func (a *A) Stage(stage *Stage) *A {

	if _, ok := stage.As[a]; !ok {
		stage.As[a] = struct{}{}
		stage.AMap_Staged_Order[a] = stage.AOrder
		stage.AOrder++
		stage.new[a] = struct{}{}
		delete(stage.deleted, a)
	} else {
		if _, ok := stage.new[a]; !ok {
			stage.modified[a] = struct{}{}
		}
	}
	stage.As_mapString[a.Name] = a

	return a
}

// Unstage removes a off the model stage
func (a *A) Unstage(stage *Stage) *A {
	delete(stage.As, a)
	delete(stage.As_mapString, a.Name)

	if _, ok := stage.reference[a]; ok {
		stage.deleted[a] = struct{}{}
	} else {
		delete(stage.new, a)
	}
	return a
}

// UnstageVoid removes a off the model stage
func (a *A) UnstageVoid(stage *Stage) {
	delete(stage.As, a)
	delete(stage.As_mapString, a.Name)
}

// commit a to the back repo (if it is already staged)
func (a *A) Commit(stage *Stage) *A {
	if _, ok := stage.As[a]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitA(a)
		}
	}
	return a
}

func (a *A) CommitVoid(stage *Stage) {
	a.Commit(stage)
}

func (a *A) StageVoid(stage *Stage) {
	a.Stage(stage)
}

// Checkout a to the back repo (if it is already staged)
func (a *A) Checkout(stage *Stage) *A {
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
func (b *B) Stage(stage *Stage) *B {

	if _, ok := stage.Bs[b]; !ok {
		stage.Bs[b] = struct{}{}
		stage.BMap_Staged_Order[b] = stage.BOrder
		stage.BOrder++
		stage.new[b] = struct{}{}
		delete(stage.deleted, b)
	} else {
		if _, ok := stage.new[b]; !ok {
			stage.modified[b] = struct{}{}
		}
	}
	stage.Bs_mapString[b.Name] = b

	return b
}

// Unstage removes b off the model stage
func (b *B) Unstage(stage *Stage) *B {
	delete(stage.Bs, b)
	delete(stage.Bs_mapString, b.Name)

	if _, ok := stage.reference[b]; ok {
		stage.deleted[b] = struct{}{}
	} else {
		delete(stage.new, b)
	}
	return b
}

// UnstageVoid removes b off the model stage
func (b *B) UnstageVoid(stage *Stage) {
	delete(stage.Bs, b)
	delete(stage.Bs_mapString, b.Name)
}

// commit b to the back repo (if it is already staged)
func (b *B) Commit(stage *Stage) *B {
	if _, ok := stage.Bs[b]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitB(b)
		}
	}
	return b
}

func (b *B) CommitVoid(stage *Stage) {
	b.Commit(stage)
}

func (b *B) StageVoid(stage *Stage) {
	b.Stage(stage)
}

// Checkout b to the back repo (if it is already staged)
func (b *B) Checkout(stage *Stage) *B {
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

func (stage *Stage) Reset() { // insertion point for array reset
	stage.As = make(map[*A]struct{})
	stage.As_mapString = make(map[string]*A)
	stage.AMap_Staged_Order = make(map[*A]uint)
	stage.AOrder = 0

	stage.Bs = make(map[*B]struct{})
	stage.Bs_mapString = make(map[string]*B)
	stage.BMap_Staged_Order = make(map[*B]uint)
	stage.BOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.As = nil
	stage.As_mapString = nil

	stage.Bs = nil
	stage.Bs_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
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
type GongstructIF interface {
	GetName() string
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
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
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
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case A:
		return any(&stage.As).(*map[*Type]struct{})
	case B:
		return any(&stage.Bs).(*map[*Type]struct{})
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
	case *A:
		return any(&stage.As).(*map[Type]struct{})
	case *B:
		return any(&stage.Bs).(*map[Type]struct{})
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *Stage) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
	// reverse maps of direct associations of A
	case A:
		switch fieldname {
		// insertion point for per direct association field
		case "Bs":
			res := make(map[*B][]*A)
			for a := range stage.As {
				for _, b_ := range a.Bs {
					res[b_] = append(res[b_], a)
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

type ReverseField struct {
	GongstructName string
	Fieldname      string
}

func GetReverseFields[Type PointerToGongstruct]() (res []ReverseField) {

	res = make([]ReverseField, 0)

	var ret Type

	switch any(ret).(type) {

	// insertion point for generic get gongstruct name
	case *A:
		var rf ReverseField
		_ = rf
	case *B:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "A"
		rf.Fieldname = "Bs"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (a *A) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NumberField",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "B",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "B",
		},
		{
			Name:                 "Bs",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "B",
		},
	}
	return
}

func (b *B) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
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
func (a *A) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = a.Name
	case "NumberField":
		res.valueString = fmt.Sprintf("%d", a.NumberField)
		res.valueInt = a.NumberField
		res.GongFieldValueType = GongFieldValueTypeInt
	case "B":
		res.GongFieldValueType = GongFieldValueTypePointer
		if a.B != nil {
			res.valueString = a.B.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, a.B))
		}
	case "Bs":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range a.Bs {
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
func (b *B) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = b.Name
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (a *A) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		a.Name = value.GetValueString()
	case "NumberField":
		a.NumberField = int(value.GetValueInt())
	case "B":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			a.B = nil
			for __instance__ := range stage.Bs {
				if stage.BMap_Staged_Order[__instance__] == uint(id) {
					a.B = __instance__
					break
				}
			}
		}
	case "Bs":
		a.Bs = make([]*B, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Bs {
					if stage.BMap_Staged_Order[__instance__] == uint(id) {
						a.Bs = append(a.Bs, __instance__)
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

func (b *B) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		b.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (a *A) GongGetGongstructName() string {
	return "A"
}

func (b *B) GongGetGongstructName() string {
	return "B"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.As_mapString = make(map[string]*A)
	for a := range stage.As {
		stage.As_mapString[a.Name] = a
	}

	stage.Bs_mapString = make(map[string]*B)
	for b := range stage.Bs {
		stage.Bs_mapString[b.Name] = b
	}

}
// Last line of the template
