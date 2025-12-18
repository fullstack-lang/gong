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

	project_go "github.com/fullstack-lang/gong/dsme/project/go"
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
	Products           map[*Product]struct{}
	Products_mapString map[string]*Product

	// insertion point for slice of pointers maps
	OnAfterProductCreateCallback OnAfterCreateInterface[Product]
	OnAfterProductUpdateCallback OnAfterUpdateInterface[Product]
	OnAfterProductDeleteCallback OnAfterDeleteInterface[Product]
	OnAfterProductReadCallback   OnAfterReadInterface[Product]

	Projects           map[*Project]struct{}
	Projects_mapString map[string]*Project

	// insertion point for slice of pointers maps
	Project_RootTasks_reverseMap map[*Task]*Project

	Project_RootProducts_reverseMap map[*Product]*Project

	OnAfterProjectCreateCallback OnAfterCreateInterface[Project]
	OnAfterProjectUpdateCallback OnAfterUpdateInterface[Project]
	OnAfterProjectDeleteCallback OnAfterDeleteInterface[Project]
	OnAfterProjectReadCallback   OnAfterReadInterface[Project]

	Roots           map[*Root]struct{}
	Roots_mapString map[string]*Root

	// insertion point for slice of pointers maps
	Root_Projects_reverseMap map[*Project]*Root

	OnAfterRootCreateCallback OnAfterCreateInterface[Root]
	OnAfterRootUpdateCallback OnAfterUpdateInterface[Root]
	OnAfterRootDeleteCallback OnAfterDeleteInterface[Root]
	OnAfterRootReadCallback   OnAfterReadInterface[Root]

	Tasks           map[*Task]struct{}
	Tasks_mapString map[string]*Task

	// insertion point for slice of pointers maps
	OnAfterTaskCreateCallback OnAfterCreateInterface[Task]
	OnAfterTaskUpdateCallback OnAfterUpdateInterface[Task]
	OnAfterTaskDeleteCallback OnAfterDeleteInterface[Task]
	OnAfterTaskReadCallback   OnAfterReadInterface[Task]

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
	ProductOrder            uint
	ProductMap_Staged_Order map[*Product]uint

	ProjectOrder            uint
	ProjectMap_Staged_Order map[*Project]uint

	RootOrder            uint
	RootMap_Staged_Order map[*Root]uint

	TaskOrder            uint
	TaskMap_Staged_Order map[*Task]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// for the computation of the diff at each commit we need
	reference map[GongstructIF]GongstructIF
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
	case *Product:
		tmp := GetStructInstancesByOrder(stage.Products, stage.ProductMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Product implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Project:
		tmp := GetStructInstancesByOrder(stage.Projects, stage.ProjectMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Project implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Root:
		tmp := GetStructInstancesByOrder(stage.Roots, stage.RootMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Root implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Task:
		tmp := GetStructInstancesByOrder(stage.Tasks, stage.TaskMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Task implements.
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
	case "Product":
		res = GetNamedStructInstances(stage.Products, stage.ProductMap_Staged_Order)
	case "Project":
		res = GetNamedStructInstances(stage.Projects, stage.ProjectMap_Staged_Order)
	case "Root":
		res = GetNamedStructInstances(stage.Roots, stage.RootMap_Staged_Order)
	case "Task":
		res = GetNamedStructInstances(stage.Tasks, stage.TaskMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/dsme/project/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return project_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return project_go.GoDiagramsDir
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
	CommitProduct(product *Product)
	CheckoutProduct(product *Product)
	CommitProject(project *Project)
	CheckoutProject(project *Project)
	CommitRoot(root *Root)
	CheckoutRoot(root *Root)
	CommitTask(task *Task)
	CheckoutTask(task *Task)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Products:           make(map[*Product]struct{}),
		Products_mapString: make(map[string]*Product),

		Projects:           make(map[*Project]struct{}),
		Projects_mapString: make(map[string]*Project),

		Roots:           make(map[*Root]struct{}),
		Roots_mapString: make(map[string]*Root),

		Tasks:           make(map[*Task]struct{}),
		Tasks_mapString: make(map[string]*Task),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ProductMap_Staged_Order: make(map[*Product]uint),

		ProjectMap_Staged_Order: make(map[*Project]uint),

		RootMap_Staged_Order: make(map[*Root]uint),

		TaskMap_Staged_Order: make(map[*Task]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Product"},
			{name: "Project"},
			{name: "Root"},
			{name: "Task"},
		}, // end of insertion point

		reference: make(map[GongstructIF]GongstructIF),
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Product:
		return stage.ProductMap_Staged_Order[instance]
	case *Project:
		return stage.ProjectMap_Staged_Order[instance]
	case *Root:
		return stage.RootMap_Staged_Order[instance]
	case *Task:
		return stage.TaskMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Product:
		return stage.ProductMap_Staged_Order[instance]
	case *Project:
		return stage.ProjectMap_Staged_Order[instance]
	case *Root:
		return stage.RootMap_Staged_Order[instance]
	case *Task:
		return stage.TaskMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["Product"] = len(stage.Products)
	stage.Map_GongStructName_InstancesNb["Project"] = len(stage.Projects)
	stage.Map_GongStructName_InstancesNb["Root"] = len(stage.Roots)
	stage.Map_GongStructName_InstancesNb["Task"] = len(stage.Tasks)
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
// Stage puts product to the model stage
func (product *Product) Stage(stage *Stage) *Product {

	if _, ok := stage.Products[product]; !ok {
		stage.Products[product] = struct{}{}
		stage.ProductMap_Staged_Order[product] = stage.ProductOrder
		stage.ProductOrder++
	}
	stage.Products_mapString[product.Name] = product

	return product
}

// StagePreserveOrder puts product to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProductOrder
// - update stage.ProductOrder accordingly
func (product *Product) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Products[product]; !ok {
		stage.Products[product] = struct{}{}

		if order > stage.ProductOrder {
			stage.ProductOrder = order
		}
		stage.ProductMap_Staged_Order[product] = stage.ProductOrder
		stage.ProductOrder++
	}
	stage.Products_mapString[product.Name] = product
}

// Unstage removes product off the model stage
func (product *Product) Unstage(stage *Stage) *Product {
	delete(stage.Products, product)
	delete(stage.ProductMap_Staged_Order, product)
	delete(stage.Products_mapString, product.Name)

	return product
}

// UnstageVoid removes product off the model stage
func (product *Product) UnstageVoid(stage *Stage) {
	delete(stage.Products, product)
	delete(stage.ProductMap_Staged_Order, product)
	delete(stage.Products_mapString, product.Name)
}

// commit product to the back repo (if it is already staged)
func (product *Product) Commit(stage *Stage) *Product {
	if _, ok := stage.Products[product]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProduct(product)
		}
	}
	return product
}

func (product *Product) CommitVoid(stage *Stage) {
	product.Commit(stage)
}

func (product *Product) StageVoid(stage *Stage) {
	product.Stage(stage)
}

// Checkout product to the back repo (if it is already staged)
func (product *Product) Checkout(stage *Stage) *Product {
	if _, ok := stage.Products[product]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProduct(product)
		}
	}
	return product
}

// for satisfaction of GongStruct interface
func (product *Product) GetName() (res string) {
	return product.Name
}

// for satisfaction of GongStruct interface
func (product *Product) SetName(name string) (){
	product.Name = name
}

// Stage puts project to the model stage
func (project *Project) Stage(stage *Stage) *Project {

	if _, ok := stage.Projects[project]; !ok {
		stage.Projects[project] = struct{}{}
		stage.ProjectMap_Staged_Order[project] = stage.ProjectOrder
		stage.ProjectOrder++
	}
	stage.Projects_mapString[project.Name] = project

	return project
}

// StagePreserveOrder puts project to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ProjectOrder
// - update stage.ProjectOrder accordingly
func (project *Project) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Projects[project]; !ok {
		stage.Projects[project] = struct{}{}

		if order > stage.ProjectOrder {
			stage.ProjectOrder = order
		}
		stage.ProjectMap_Staged_Order[project] = stage.ProjectOrder
		stage.ProjectOrder++
	}
	stage.Projects_mapString[project.Name] = project
}

// Unstage removes project off the model stage
func (project *Project) Unstage(stage *Stage) *Project {
	delete(stage.Projects, project)
	delete(stage.ProjectMap_Staged_Order, project)
	delete(stage.Projects_mapString, project.Name)

	return project
}

// UnstageVoid removes project off the model stage
func (project *Project) UnstageVoid(stage *Stage) {
	delete(stage.Projects, project)
	delete(stage.ProjectMap_Staged_Order, project)
	delete(stage.Projects_mapString, project.Name)
}

// commit project to the back repo (if it is already staged)
func (project *Project) Commit(stage *Stage) *Project {
	if _, ok := stage.Projects[project]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitProject(project)
		}
	}
	return project
}

func (project *Project) CommitVoid(stage *Stage) {
	project.Commit(stage)
}

func (project *Project) StageVoid(stage *Stage) {
	project.Stage(stage)
}

// Checkout project to the back repo (if it is already staged)
func (project *Project) Checkout(stage *Stage) *Project {
	if _, ok := stage.Projects[project]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutProject(project)
		}
	}
	return project
}

// for satisfaction of GongStruct interface
func (project *Project) GetName() (res string) {
	return project.Name
}

// for satisfaction of GongStruct interface
func (project *Project) SetName(name string) (){
	project.Name = name
}

// Stage puts root to the model stage
func (root *Root) Stage(stage *Stage) *Root {

	if _, ok := stage.Roots[root]; !ok {
		stage.Roots[root] = struct{}{}
		stage.RootMap_Staged_Order[root] = stage.RootOrder
		stage.RootOrder++
	}
	stage.Roots_mapString[root.Name] = root

	return root
}

// StagePreserveOrder puts root to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RootOrder
// - update stage.RootOrder accordingly
func (root *Root) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Roots[root]; !ok {
		stage.Roots[root] = struct{}{}

		if order > stage.RootOrder {
			stage.RootOrder = order
		}
		stage.RootMap_Staged_Order[root] = stage.RootOrder
		stage.RootOrder++
	}
	stage.Roots_mapString[root.Name] = root
}

// Unstage removes root off the model stage
func (root *Root) Unstage(stage *Stage) *Root {
	delete(stage.Roots, root)
	delete(stage.RootMap_Staged_Order, root)
	delete(stage.Roots_mapString, root.Name)

	return root
}

// UnstageVoid removes root off the model stage
func (root *Root) UnstageVoid(stage *Stage) {
	delete(stage.Roots, root)
	delete(stage.RootMap_Staged_Order, root)
	delete(stage.Roots_mapString, root.Name)
}

// commit root to the back repo (if it is already staged)
func (root *Root) Commit(stage *Stage) *Root {
	if _, ok := stage.Roots[root]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRoot(root)
		}
	}
	return root
}

func (root *Root) CommitVoid(stage *Stage) {
	root.Commit(stage)
}

func (root *Root) StageVoid(stage *Stage) {
	root.Stage(stage)
}

// Checkout root to the back repo (if it is already staged)
func (root *Root) Checkout(stage *Stage) *Root {
	if _, ok := stage.Roots[root]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRoot(root)
		}
	}
	return root
}

// for satisfaction of GongStruct interface
func (root *Root) GetName() (res string) {
	return root.Name
}

// for satisfaction of GongStruct interface
func (root *Root) SetName(name string) (){
	root.Name = name
}

// Stage puts task to the model stage
func (task *Task) Stage(stage *Stage) *Task {

	if _, ok := stage.Tasks[task]; !ok {
		stage.Tasks[task] = struct{}{}
		stage.TaskMap_Staged_Order[task] = stage.TaskOrder
		stage.TaskOrder++
	}
	stage.Tasks_mapString[task.Name] = task

	return task
}

// StagePreserveOrder puts task to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TaskOrder
// - update stage.TaskOrder accordingly
func (task *Task) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Tasks[task]; !ok {
		stage.Tasks[task] = struct{}{}

		if order > stage.TaskOrder {
			stage.TaskOrder = order
		}
		stage.TaskMap_Staged_Order[task] = stage.TaskOrder
		stage.TaskOrder++
	}
	stage.Tasks_mapString[task.Name] = task
}

// Unstage removes task off the model stage
func (task *Task) Unstage(stage *Stage) *Task {
	delete(stage.Tasks, task)
	delete(stage.TaskMap_Staged_Order, task)
	delete(stage.Tasks_mapString, task.Name)

	return task
}

// UnstageVoid removes task off the model stage
func (task *Task) UnstageVoid(stage *Stage) {
	delete(stage.Tasks, task)
	delete(stage.TaskMap_Staged_Order, task)
	delete(stage.Tasks_mapString, task.Name)
}

// commit task to the back repo (if it is already staged)
func (task *Task) Commit(stage *Stage) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTask(task)
		}
	}
	return task
}

func (task *Task) CommitVoid(stage *Stage) {
	task.Commit(stage)
}

func (task *Task) StageVoid(stage *Stage) {
	task.Stage(stage)
}

// Checkout task to the back repo (if it is already staged)
func (task *Task) Checkout(stage *Stage) *Task {
	if _, ok := stage.Tasks[task]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTask(task)
		}
	}
	return task
}

// for satisfaction of GongStruct interface
func (task *Task) GetName() (res string) {
	return task.Name
}

// for satisfaction of GongStruct interface
func (task *Task) SetName(name string) (){
	task.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMProduct(Product *Product)
	CreateORMProject(Project *Project)
	CreateORMRoot(Root *Root)
	CreateORMTask(Task *Task)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMProduct(Product *Product)
	DeleteORMProject(Project *Project)
	DeleteORMRoot(Root *Root)
	DeleteORMTask(Task *Task)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Products = make(map[*Product]struct{})
	stage.Products_mapString = make(map[string]*Product)
	stage.ProductMap_Staged_Order = make(map[*Product]uint)
	stage.ProductOrder = 0

	stage.Projects = make(map[*Project]struct{})
	stage.Projects_mapString = make(map[string]*Project)
	stage.ProjectMap_Staged_Order = make(map[*Project]uint)
	stage.ProjectOrder = 0

	stage.Roots = make(map[*Root]struct{})
	stage.Roots_mapString = make(map[string]*Root)
	stage.RootMap_Staged_Order = make(map[*Root]uint)
	stage.RootOrder = 0

	stage.Tasks = make(map[*Task]struct{})
	stage.Tasks_mapString = make(map[string]*Task)
	stage.TaskMap_Staged_Order = make(map[*Task]uint)
	stage.TaskOrder = 0

	stage.ComputeReference()
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Products = nil
	stage.Products_mapString = nil

	stage.Projects = nil
	stage.Projects_mapString = nil

	stage.Roots = nil
	stage.Roots_mapString = nil

	stage.Tasks = nil
	stage.Tasks_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for product := range stage.Products {
		product.Unstage(stage)
	}

	for project := range stage.Projects {
		project.Unstage(stage)
	}

	for root := range stage.Roots {
		root.Unstage(stage)
	}

	for task := range stage.Tasks {
		task.Unstage(stage)
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
	case map[*Product]any:
		return any(&stage.Products).(*Type)
	case map[*Project]any:
		return any(&stage.Projects).(*Type)
	case map[*Root]any:
		return any(&stage.Roots).(*Type)
	case map[*Task]any:
		return any(&stage.Tasks).(*Type)
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
	case *Product:
		return any(stage.Products_mapString).(map[string]Type)
	case *Project:
		return any(stage.Projects_mapString).(map[string]Type)
	case *Root:
		return any(stage.Roots_mapString).(map[string]Type)
	case *Task:
		return any(stage.Tasks_mapString).(map[string]Type)
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
	case Product:
		return any(&stage.Products).(*map[*Type]struct{})
	case Project:
		return any(&stage.Projects).(*map[*Type]struct{})
	case Root:
		return any(&stage.Roots).(*map[*Type]struct{})
	case Task:
		return any(&stage.Tasks).(*map[*Type]struct{})
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
	case *Product:
		return any(&stage.Products).(*map[Type]struct{})
	case *Project:
		return any(&stage.Projects).(*map[Type]struct{})
	case *Root:
		return any(&stage.Roots).(*map[Type]struct{})
	case *Task:
		return any(&stage.Tasks).(*map[Type]struct{})
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
	case Product:
		return any(&stage.Products_mapString).(*map[string]*Type)
	case Project:
		return any(&stage.Projects_mapString).(*map[string]*Type)
	case Root:
		return any(&stage.Roots_mapString).(*map[string]*Type)
	case Task:
		return any(&stage.Tasks_mapString).(*map[string]*Type)
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
	case Product:
		return any(&Product{
			// Initialisation of associations
			// field is initialized with an instance of Product with the name of the field
			ParentProduct: &Product{Name: "ParentProduct"},
		}).(*Type)
	case Project:
		return any(&Project{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			RootTasks: []*Task{{Name: "RootTasks"}},
			// field is initialized with an instance of Product with the name of the field
			RootProducts: []*Product{{Name: "RootProducts"}},
		}).(*Type)
	case Root:
		return any(&Root{
			// Initialisation of associations
			// field is initialized with an instance of Project with the name of the field
			Projects: []*Project{{Name: "Projects"}},
		}).(*Type)
	case Task:
		return any(&Task{
			// Initialisation of associations
			// field is initialized with an instance of Task with the name of the field
			ParentTask: &Task{Name: "ParentTask"},
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
	// reverse maps of direct associations of Product
	case Product:
		switch fieldname {
		// insertion point for per direct association field
		case "ParentProduct":
			res := make(map[*Product][]*Product)
			for product := range stage.Products {
				if product.ParentProduct != nil {
					product_ := product.ParentProduct
					var products []*Product
					_, ok := res[product_]
					if ok {
						products = res[product_]
					} else {
						products = make([]*Product, 0)
					}
					products = append(products, product)
					res[product_] = products
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Project
	case Project:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Root
	case Root:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Task
	case Task:
		switch fieldname {
		// insertion point for per direct association field
		case "ParentTask":
			res := make(map[*Task][]*Task)
			for task := range stage.Tasks {
				if task.ParentTask != nil {
					task_ := task.ParentTask
					var tasks []*Task
					_, ok := res[task_]
					if ok {
						tasks = res[task_]
					} else {
						tasks = make([]*Task, 0)
					}
					tasks = append(tasks, task)
					res[task_] = tasks
				}
			}
			return any(res).(map[*End][]*Start)
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
	// reverse maps of direct associations of Product
	case Product:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Project
	case Project:
		switch fieldname {
		// insertion point for per direct association field
		case "RootTasks":
			res := make(map[*Task][]*Project)
			for project := range stage.Projects {
				for _, task_ := range project.RootTasks {
					res[task_] = append(res[task_], project)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RootProducts":
			res := make(map[*Product][]*Project)
			for project := range stage.Projects {
				for _, product_ := range project.RootProducts {
					res[product_] = append(res[product_], project)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Root
	case Root:
		switch fieldname {
		// insertion point for per direct association field
		case "Projects":
			res := make(map[*Project][]*Root)
			for root := range stage.Roots {
				for _, project_ := range root.Projects {
					res[project_] = append(res[project_], root)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Task
	case Task:
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
	case *Product:
		res = "Product"
	case *Project:
		res = "Project"
	case *Root:
		res = "Root"
	case *Task:
		res = "Task"
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
	case *Product:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Project"
		rf.Fieldname = "RootProducts"
		res = append(res, rf)
	case *Project:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Root"
		rf.Fieldname = "Projects"
		res = append(res, rf)
	case *Root:
		var rf ReverseField
		_ = rf
	case *Task:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Project"
		rf.Fieldname = "RootTasks"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (product *Product) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "ParentProduct",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Product",
		},
		{
			Name:               "IsExpanded",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (project *Project) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "RootTasks",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Task",
		},
		{
			Name:                 "RootProducts",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Product",
		},
	}
	return
}

func (root *Root) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "Projects",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Project",
		},
	}
	return
}

func (task *Task) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:                 "ParentTask",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "Task",
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
func (product *Product) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = product.Name
	case "ParentProduct":
		res.GongFieldValueType = GongFieldValueTypePointer
		if product.ParentProduct != nil {
			res.valueString = product.ParentProduct.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, product.ParentProduct))
		}
	case "IsExpanded":
		res.valueString = fmt.Sprintf("%t", product.IsExpanded)
		res.valueBool = product.IsExpanded
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}
func (project *Project) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = project.Name
	case "RootTasks":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range project.RootTasks {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "RootProducts":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range project.RootProducts {
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
func (root *Root) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = root.Name
	case "Projects":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range root.Projects {
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
func (task *Task) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = task.Name
	case "ParentTask":
		res.GongFieldValueType = GongFieldValueTypePointer
		if task.ParentTask != nil {
			res.valueString = task.ParentTask.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, task.ParentTask))
		}
	}
	return
}
func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (product *Product) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		product.Name = value.GetValueString()
	case "ParentProduct":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			product.ParentProduct = nil
			for __instance__ := range stage.Products {
				if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
					product.ParentProduct = __instance__
					break
				}
			}
		}
	case "IsExpanded":
		product.IsExpanded = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (project *Project) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		project.Name = value.GetValueString()
	case "RootTasks":
		project.RootTasks = make([]*Task, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Tasks {
					if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
						project.RootTasks = append(project.RootTasks, __instance__)
						break
					}
				}
			}
		}
	case "RootProducts":
		project.RootProducts = make([]*Product, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Products {
					if stage.ProductMap_Staged_Order[__instance__] == uint(id) {
						project.RootProducts = append(project.RootProducts, __instance__)
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

func (root *Root) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		root.Name = value.GetValueString()
	case "Projects":
		root.Projects = make([]*Project, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Projects {
					if stage.ProjectMap_Staged_Order[__instance__] == uint(id) {
						root.Projects = append(root.Projects, __instance__)
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

func (task *Task) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		task.Name = value.GetValueString()
	case "ParentTask":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			task.ParentTask = nil
			for __instance__ := range stage.Tasks {
				if stage.TaskMap_Staged_Order[__instance__] == uint(id) {
					task.ParentTask = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (product *Product) GongGetGongstructName() string {
	return "Product"
}

func (project *Project) GongGetGongstructName() string {
	return "Project"
}

func (root *Root) GongGetGongstructName() string {
	return "Root"
}

func (task *Task) GongGetGongstructName() string {
	return "Task"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Products_mapString = make(map[string]*Product)
	for product := range stage.Products {
		stage.Products_mapString[product.Name] = product
	}

	stage.Projects_mapString = make(map[string]*Project)
	for project := range stage.Projects {
		stage.Projects_mapString[project.Name] = project
	}

	stage.Roots_mapString = make(map[string]*Root)
	for root := range stage.Roots {
		stage.Roots_mapString[root.Name] = root
	}

	stage.Tasks_mapString = make(map[string]*Task)
	for task := range stage.Tasks {
		stage.Tasks_mapString[task.Name] = task
	}

}
// Last line of the template
