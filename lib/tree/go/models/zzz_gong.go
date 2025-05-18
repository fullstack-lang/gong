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

	tree_go "github.com/fullstack-lang/gong/lib/tree/go"
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
	name string

	// insertion point for definition of arrays registering instances
	Buttons           map[*Button]any
	Buttons_mapString map[string]*Button

	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback OnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback OnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback OnAfterDeleteInterface[Button]
	OnAfterButtonReadCallback   OnAfterReadInterface[Button]

	Nodes           map[*Node]any
	Nodes_mapString map[string]*Node

	// insertion point for slice of pointers maps
	Node_Children_reverseMap map[*Node]*Node

	Node_Buttons_reverseMap map[*Button]*Node

	OnAfterNodeCreateCallback OnAfterCreateInterface[Node]
	OnAfterNodeUpdateCallback OnAfterUpdateInterface[Node]
	OnAfterNodeDeleteCallback OnAfterDeleteInterface[Node]
	OnAfterNodeReadCallback   OnAfterReadInterface[Node]

	SVGIcons           map[*SVGIcon]any
	SVGIcons_mapString map[string]*SVGIcon

	// insertion point for slice of pointers maps
	OnAfterSVGIconCreateCallback OnAfterCreateInterface[SVGIcon]
	OnAfterSVGIconUpdateCallback OnAfterUpdateInterface[SVGIcon]
	OnAfterSVGIconDeleteCallback OnAfterDeleteInterface[SVGIcon]
	OnAfterSVGIconReadCallback   OnAfterReadInterface[SVGIcon]

	Trees           map[*Tree]any
	Trees_mapString map[string]*Tree

	// insertion point for slice of pointers maps
	Tree_RootNodes_reverseMap map[*Node]*Tree

	OnAfterTreeCreateCallback OnAfterCreateInterface[Tree]
	OnAfterTreeUpdateCallback OnAfterUpdateInterface[Tree]
	OnAfterTreeDeleteCallback OnAfterDeleteInterface[Tree]
	OnAfterTreeReadCallback   OnAfterReadInterface[Tree]

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

	NodeOrder            uint
	NodeMap_Staged_Order map[*Node]uint

	SVGIconOrder            uint
	SVGIconMap_Staged_Order map[*SVGIcon]uint

	TreeOrder            uint
	TreeMap_Staged_Order map[*Tree]uint

	// end of insertion point

	NamedStructs []*NamedStruct
}

// GetNamedStructs implements models.ProbebStage.
func (stage *Stage) GetNamedStructsNames() (res []string) {

	for _, namedStruct := range stage.NamedStructs {
		res = append(res, namedStruct.name)
	}

	return
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

func (stage *Stage) GetNamedStructNamesByOrder(namedStructName string) (res []string) {

	switch namedStructName {
	// insertion point for case
	case "Button":
		res = GetNamedStructInstances(stage.Buttons, stage.ButtonMap_Staged_Order)
	case "Node":
		res = GetNamedStructInstances(stage.Nodes, stage.NodeMap_Staged_Order)
	case "SVGIcon":
		res = GetNamedStructInstances(stage.SVGIcons, stage.SVGIconMap_Staged_Order)
	case "Tree":
		res = GetNamedStructInstances(stage.Trees, stage.TreeMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/tree/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return tree_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return tree_go.GoDiagramsDir
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
	CommitNode(node *Node)
	CheckoutNode(node *Node)
	CommitSVGIcon(svgicon *SVGIcon)
	CheckoutSVGIcon(svgicon *SVGIcon)
	CommitTree(tree *Tree)
	CheckoutTree(tree *Tree)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Buttons:           make(map[*Button]any),
		Buttons_mapString: make(map[string]*Button),

		Nodes:           make(map[*Node]any),
		Nodes_mapString: make(map[string]*Node),

		SVGIcons:           make(map[*SVGIcon]any),
		SVGIcons_mapString: make(map[string]*SVGIcon),

		Trees:           make(map[*Tree]any),
		Trees_mapString: make(map[string]*Tree),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ButtonMap_Staged_Order: make(map[*Button]uint),

		NodeMap_Staged_Order: make(map[*Node]uint),

		SVGIconMap_Staged_Order: make(map[*SVGIcon]uint),

		TreeMap_Staged_Order: make(map[*Tree]uint),

		// end of insertion point

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Button"},
			{name: "Node"},
			{name: "SVGIcon"},
			{name: "Tree"},
		}, // end of insertion point
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Button:
		return stage.ButtonMap_Staged_Order[instance]
	case *Node:
		return stage.NodeMap_Staged_Order[instance]
	case *SVGIcon:
		return stage.SVGIconMap_Staged_Order[instance]
	case *Tree:
		return stage.TreeMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Button:
		return stage.ButtonMap_Staged_Order[instance]
	case *Node:
		return stage.NodeMap_Staged_Order[instance]
	case *SVGIcon:
		return stage.SVGIconMap_Staged_Order[instance]
	case *Tree:
		return stage.TreeMap_Staged_Order[instance]
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
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["SVGIcon"] = len(stage.SVGIcons)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)

}

func (stage *Stage) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Node"] = len(stage.Nodes)
	stage.Map_GongStructName_InstancesNb["SVGIcon"] = len(stage.SVGIcons)
	stage.Map_GongStructName_InstancesNb["Tree"] = len(stage.Trees)

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
	}
	stage.Buttons_mapString[button.Name] = button

	return button
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	delete(stage.Buttons, button)
	delete(stage.Buttons_mapString, button.Name)
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

// Stage puts node to the model stage
func (node *Node) Stage(stage *Stage) *Node {

	if _, ok := stage.Nodes[node]; !ok {
		stage.Nodes[node] = __member
		stage.NodeMap_Staged_Order[node] = stage.NodeOrder
		stage.NodeOrder++
	}
	stage.Nodes_mapString[node.Name] = node

	return node
}

// Unstage removes node off the model stage
func (node *Node) Unstage(stage *Stage) *Node {
	delete(stage.Nodes, node)
	delete(stage.Nodes_mapString, node.Name)
	return node
}

// UnstageVoid removes node off the model stage
func (node *Node) UnstageVoid(stage *Stage) {
	delete(stage.Nodes, node)
	delete(stage.Nodes_mapString, node.Name)
}

// commit node to the back repo (if it is already staged)
func (node *Node) Commit(stage *Stage) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitNode(node)
		}
	}
	return node
}

func (node *Node) CommitVoid(stage *Stage) {
	node.Commit(stage)
}

// Checkout node to the back repo (if it is already staged)
func (node *Node) Checkout(stage *Stage) *Node {
	if _, ok := stage.Nodes[node]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutNode(node)
		}
	}
	return node
}

// for satisfaction of GongStruct interface
func (node *Node) GetName() (res string) {
	return node.Name
}

// Stage puts svgicon to the model stage
func (svgicon *SVGIcon) Stage(stage *Stage) *SVGIcon {

	if _, ok := stage.SVGIcons[svgicon]; !ok {
		stage.SVGIcons[svgicon] = __member
		stage.SVGIconMap_Staged_Order[svgicon] = stage.SVGIconOrder
		stage.SVGIconOrder++
	}
	stage.SVGIcons_mapString[svgicon.Name] = svgicon

	return svgicon
}

// Unstage removes svgicon off the model stage
func (svgicon *SVGIcon) Unstage(stage *Stage) *SVGIcon {
	delete(stage.SVGIcons, svgicon)
	delete(stage.SVGIcons_mapString, svgicon.Name)
	return svgicon
}

// UnstageVoid removes svgicon off the model stage
func (svgicon *SVGIcon) UnstageVoid(stage *Stage) {
	delete(stage.SVGIcons, svgicon)
	delete(stage.SVGIcons_mapString, svgicon.Name)
}

// commit svgicon to the back repo (if it is already staged)
func (svgicon *SVGIcon) Commit(stage *Stage) *SVGIcon {
	if _, ok := stage.SVGIcons[svgicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitSVGIcon(svgicon)
		}
	}
	return svgicon
}

func (svgicon *SVGIcon) CommitVoid(stage *Stage) {
	svgicon.Commit(stage)
}

// Checkout svgicon to the back repo (if it is already staged)
func (svgicon *SVGIcon) Checkout(stage *Stage) *SVGIcon {
	if _, ok := stage.SVGIcons[svgicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutSVGIcon(svgicon)
		}
	}
	return svgicon
}

// for satisfaction of GongStruct interface
func (svgicon *SVGIcon) GetName() (res string) {
	return svgicon.Name
}

// Stage puts tree to the model stage
func (tree *Tree) Stage(stage *Stage) *Tree {

	if _, ok := stage.Trees[tree]; !ok {
		stage.Trees[tree] = __member
		stage.TreeMap_Staged_Order[tree] = stage.TreeOrder
		stage.TreeOrder++
	}
	stage.Trees_mapString[tree.Name] = tree

	return tree
}

// Unstage removes tree off the model stage
func (tree *Tree) Unstage(stage *Stage) *Tree {
	delete(stage.Trees, tree)
	delete(stage.Trees_mapString, tree.Name)
	return tree
}

// UnstageVoid removes tree off the model stage
func (tree *Tree) UnstageVoid(stage *Stage) {
	delete(stage.Trees, tree)
	delete(stage.Trees_mapString, tree.Name)
}

// commit tree to the back repo (if it is already staged)
func (tree *Tree) Commit(stage *Stage) *Tree {
	if _, ok := stage.Trees[tree]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTree(tree)
		}
	}
	return tree
}

func (tree *Tree) CommitVoid(stage *Stage) {
	tree.Commit(stage)
}

// Checkout tree to the back repo (if it is already staged)
func (tree *Tree) Checkout(stage *Stage) *Tree {
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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMButton(Button *Button)
	CreateORMNode(Node *Node)
	CreateORMSVGIcon(SVGIcon *SVGIcon)
	CreateORMTree(Tree *Tree)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMButton(Button *Button)
	DeleteORMNode(Node *Node)
	DeleteORMSVGIcon(SVGIcon *SVGIcon)
	DeleteORMTree(Tree *Tree)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Buttons = make(map[*Button]any)
	stage.Buttons_mapString = make(map[string]*Button)
	stage.ButtonMap_Staged_Order = make(map[*Button]uint)
	stage.ButtonOrder = 0

	stage.Nodes = make(map[*Node]any)
	stage.Nodes_mapString = make(map[string]*Node)
	stage.NodeMap_Staged_Order = make(map[*Node]uint)
	stage.NodeOrder = 0

	stage.SVGIcons = make(map[*SVGIcon]any)
	stage.SVGIcons_mapString = make(map[string]*SVGIcon)
	stage.SVGIconMap_Staged_Order = make(map[*SVGIcon]uint)
	stage.SVGIconOrder = 0

	stage.Trees = make(map[*Tree]any)
	stage.Trees_mapString = make(map[string]*Tree)
	stage.TreeMap_Staged_Order = make(map[*Tree]uint)
	stage.TreeOrder = 0

}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Buttons = nil
	stage.Buttons_mapString = nil

	stage.Nodes = nil
	stage.Nodes_mapString = nil

	stage.SVGIcons = nil
	stage.SVGIcons_mapString = nil

	stage.Trees = nil
	stage.Trees_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for button := range stage.Buttons {
		button.Unstage(stage)
	}

	for node := range stage.Nodes {
		node.Unstage(stage)
	}

	for svgicon := range stage.SVGIcons {
		svgicon.Unstage(stage)
	}

	for tree := range stage.Trees {
		tree.Unstage(stage)
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
	case map[*Button]any:
		return any(&stage.Buttons).(*Type)
	case map[*Node]any:
		return any(&stage.Nodes).(*Type)
	case map[*SVGIcon]any:
		return any(&stage.SVGIcons).(*Type)
	case map[*Tree]any:
		return any(&stage.Trees).(*Type)
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
	case map[string]*Node:
		return any(&stage.Nodes_mapString).(*Type)
	case map[string]*SVGIcon:
		return any(&stage.SVGIcons_mapString).(*Type)
	case map[string]*Tree:
		return any(&stage.Trees_mapString).(*Type)
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
	case Node:
		return any(&stage.Nodes).(*map[*Type]any)
	case SVGIcon:
		return any(&stage.SVGIcons).(*map[*Type]any)
	case Tree:
		return any(&stage.Trees).(*map[*Type]any)
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
	case *Node:
		return any(&stage.Nodes).(*map[Type]any)
	case *SVGIcon:
		return any(&stage.SVGIcons).(*map[Type]any)
	case *Tree:
		return any(&stage.Trees).(*map[Type]any)
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
	case Node:
		return any(&stage.Nodes_mapString).(*map[string]*Type)
	case SVGIcon:
		return any(&stage.SVGIcons_mapString).(*map[string]*Type)
	case Tree:
		return any(&stage.Trees_mapString).(*map[string]*Type)
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
			// field is initialized with an instance of SVGIcon with the name of the field
			SVGIcon: &SVGIcon{Name: "SVGIcon"},
		}).(*Type)
	case Node:
		return any(&Node{
			// Initialisation of associations
			// field is initialized with an instance of SVGIcon with the name of the field
			PreceedingSVGIcon: &SVGIcon{Name: "PreceedingSVGIcon"},
			// field is initialized with an instance of Node with the name of the field
			Children: []*Node{{Name: "Children"}},
			// field is initialized with an instance of Button with the name of the field
			Buttons: []*Button{{Name: "Buttons"}},
		}).(*Type)
	case SVGIcon:
		return any(&SVGIcon{
			// Initialisation of associations
		}).(*Type)
	case Tree:
		return any(&Tree{
			// Initialisation of associations
			// field is initialized with an instance of Node with the name of the field
			RootNodes: []*Node{{Name: "RootNodes"}},
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
		case "SVGIcon":
			res := make(map[*SVGIcon][]*Button)
			for button := range stage.Buttons {
				if button.SVGIcon != nil {
					svgicon_ := button.SVGIcon
					var buttons []*Button
					_, ok := res[svgicon_]
					if ok {
						buttons = res[svgicon_]
					} else {
						buttons = make([]*Button, 0)
					}
					buttons = append(buttons, button)
					res[svgicon_] = buttons
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		case "PreceedingSVGIcon":
			res := make(map[*SVGIcon][]*Node)
			for node := range stage.Nodes {
				if node.PreceedingSVGIcon != nil {
					svgicon_ := node.PreceedingSVGIcon
					var nodes []*Node
					_, ok := res[svgicon_]
					if ok {
						nodes = res[svgicon_]
					} else {
						nodes = make([]*Node, 0)
					}
					nodes = append(nodes, node)
					res[svgicon_] = nodes
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SVGIcon
	case SVGIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
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
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Node
	case Node:
		switch fieldname {
		// insertion point for per direct association field
		case "Children":
			res := make(map[*Node]*Node)
			for node := range stage.Nodes {
				for _, node_ := range node.Children {
					res[node_] = node
				}
			}
			return any(res).(map[*End]*Start)
		case "Buttons":
			res := make(map[*Button]*Node)
			for node := range stage.Nodes {
				for _, button_ := range node.Buttons {
					res[button_] = node
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of SVGIcon
	case SVGIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Tree
	case Tree:
		switch fieldname {
		// insertion point for per direct association field
		case "RootNodes":
			res := make(map[*Node]*Tree)
			for tree := range stage.Trees {
				for _, node_ := range tree.RootNodes {
					res[node_] = tree
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
	case Node:
		res = "Node"
	case SVGIcon:
		res = "SVGIcon"
	case Tree:
		res = "Tree"
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
	case *Node:
		res = "Node"
	case *SVGIcon:
		res = "SVGIcon"
	case *Tree:
		res = "Tree"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Button:
		res = []string{"Name", "Icon", "SVGIcon", "IsDisabled", "HasToolTip", "ToolTipText", "ToolTipPosition"}
	case Node:
		res = []string{"Name", "FontStyle", "BackgroundColor", "IsExpanded", "HasCheckboxButton", "IsChecked", "IsCheckboxDisabled", "CheckboxHasToolTip", "CheckboxToolTipText", "CheckboxToolTipPosition", "HasSecondCheckboxButton", "IsSecondCheckboxChecked", "IsSecondCheckboxDisabled", "TextAfterSecondCheckbox", "SecondCheckboxHasToolTip", "SecondCheckboxToolTipText", "SecondCheckboxToolTipPosition", "IsInEditMode", "IsNodeClickable", "IsWithPreceedingIcon", "PreceedingIcon", "PreceedingSVGIcon", "Children", "Buttons"}
	case SVGIcon:
		res = []string{"Name", "SVG"}
	case Tree:
		res = []string{"Name", "RootNodes"}
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
		rf.GongstructName = "Node"
		rf.Fieldname = "Buttons"
		res = append(res, rf)
	case Node:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Node"
		rf.Fieldname = "Children"
		res = append(res, rf)
		rf.GongstructName = "Tree"
		rf.Fieldname = "RootNodes"
		res = append(res, rf)
	case SVGIcon:
		var rf ReverseField
		_ = rf
	case Tree:
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
		res = []string{"Name", "Icon", "SVGIcon", "IsDisabled", "HasToolTip", "ToolTipText", "ToolTipPosition"}
	case *Node:
		res = []string{"Name", "FontStyle", "BackgroundColor", "IsExpanded", "HasCheckboxButton", "IsChecked", "IsCheckboxDisabled", "CheckboxHasToolTip", "CheckboxToolTipText", "CheckboxToolTipPosition", "HasSecondCheckboxButton", "IsSecondCheckboxChecked", "IsSecondCheckboxDisabled", "TextAfterSecondCheckbox", "SecondCheckboxHasToolTip", "SecondCheckboxToolTipText", "SecondCheckboxToolTipPosition", "IsInEditMode", "IsNodeClickable", "IsWithPreceedingIcon", "PreceedingIcon", "PreceedingSVGIcon", "Children", "Buttons"}
	case *SVGIcon:
		res = []string{"Name", "SVG"}
	case *Tree:
		res = []string{"Name", "RootNodes"}
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
		case "Icon":
			res.valueString = inferedInstance.Icon
		case "SVGIcon":
			if inferedInstance.SVGIcon != nil {
				res.valueString = inferedInstance.SVGIcon.Name
			}
		case "IsDisabled":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisabled)
			res.valueBool = inferedInstance.IsDisabled
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasToolTip":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasToolTip)
			res.valueBool = inferedInstance.HasToolTip
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ToolTipText":
			res.valueString = inferedInstance.ToolTipText
		case "ToolTipPosition":
			enum := inferedInstance.ToolTipPosition
			res.valueString = enum.ToCodeString()
		}
	case *Node:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "FontStyle":
			enum := inferedInstance.FontStyle
			res.valueString = enum.ToCodeString()
		case "BackgroundColor":
			res.valueString = inferedInstance.BackgroundColor
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasCheckboxButton":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasCheckboxButton)
			res.valueBool = inferedInstance.HasCheckboxButton
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsChecked":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsChecked)
			res.valueBool = inferedInstance.IsChecked
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsCheckboxDisabled":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsCheckboxDisabled)
			res.valueBool = inferedInstance.IsCheckboxDisabled
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CheckboxHasToolTip":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CheckboxHasToolTip)
			res.valueBool = inferedInstance.CheckboxHasToolTip
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CheckboxToolTipText":
			res.valueString = inferedInstance.CheckboxToolTipText
		case "CheckboxToolTipPosition":
			enum := inferedInstance.CheckboxToolTipPosition
			res.valueString = enum.ToCodeString()
		case "HasSecondCheckboxButton":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasSecondCheckboxButton)
			res.valueBool = inferedInstance.HasSecondCheckboxButton
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsSecondCheckboxChecked":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSecondCheckboxChecked)
			res.valueBool = inferedInstance.IsSecondCheckboxChecked
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsSecondCheckboxDisabled":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSecondCheckboxDisabled)
			res.valueBool = inferedInstance.IsSecondCheckboxDisabled
			res.GongFieldValueType = GongFieldValueTypeBool
		case "TextAfterSecondCheckbox":
			res.valueString = inferedInstance.TextAfterSecondCheckbox
		case "SecondCheckboxHasToolTip":
			res.valueString = fmt.Sprintf("%t", inferedInstance.SecondCheckboxHasToolTip)
			res.valueBool = inferedInstance.SecondCheckboxHasToolTip
			res.GongFieldValueType = GongFieldValueTypeBool
		case "SecondCheckboxToolTipText":
			res.valueString = inferedInstance.SecondCheckboxToolTipText
		case "SecondCheckboxToolTipPosition":
			enum := inferedInstance.SecondCheckboxToolTipPosition
			res.valueString = enum.ToCodeString()
		case "IsInEditMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInEditMode)
			res.valueBool = inferedInstance.IsInEditMode
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsNodeClickable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsNodeClickable)
			res.valueBool = inferedInstance.IsNodeClickable
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsWithPreceedingIcon":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsWithPreceedingIcon)
			res.valueBool = inferedInstance.IsWithPreceedingIcon
			res.GongFieldValueType = GongFieldValueTypeBool
		case "PreceedingIcon":
			res.valueString = inferedInstance.PreceedingIcon
		case "PreceedingSVGIcon":
			if inferedInstance.PreceedingSVGIcon != nil {
				res.valueString = inferedInstance.PreceedingSVGIcon.Name
			}
		case "Children":
			for idx, __instance__ := range inferedInstance.Children {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Buttons":
			for idx, __instance__ := range inferedInstance.Buttons {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *SVGIcon:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "SVG":
			res.valueString = inferedInstance.SVG
		}
	case *Tree:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "RootNodes":
			for idx, __instance__ := range inferedInstance.RootNodes {
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
		case "Icon":
			res.valueString = inferedInstance.Icon
		case "SVGIcon":
			if inferedInstance.SVGIcon != nil {
				res.valueString = inferedInstance.SVGIcon.Name
			}
		case "IsDisabled":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsDisabled)
			res.valueBool = inferedInstance.IsDisabled
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasToolTip":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasToolTip)
			res.valueBool = inferedInstance.HasToolTip
			res.GongFieldValueType = GongFieldValueTypeBool
		case "ToolTipText":
			res.valueString = inferedInstance.ToolTipText
		case "ToolTipPosition":
			enum := inferedInstance.ToolTipPosition
			res.valueString = enum.ToCodeString()
		}
	case Node:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "FontStyle":
			enum := inferedInstance.FontStyle
			res.valueString = enum.ToCodeString()
		case "BackgroundColor":
			res.valueString = inferedInstance.BackgroundColor
		case "IsExpanded":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsExpanded)
			res.valueBool = inferedInstance.IsExpanded
			res.GongFieldValueType = GongFieldValueTypeBool
		case "HasCheckboxButton":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasCheckboxButton)
			res.valueBool = inferedInstance.HasCheckboxButton
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsChecked":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsChecked)
			res.valueBool = inferedInstance.IsChecked
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsCheckboxDisabled":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsCheckboxDisabled)
			res.valueBool = inferedInstance.IsCheckboxDisabled
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CheckboxHasToolTip":
			res.valueString = fmt.Sprintf("%t", inferedInstance.CheckboxHasToolTip)
			res.valueBool = inferedInstance.CheckboxHasToolTip
			res.GongFieldValueType = GongFieldValueTypeBool
		case "CheckboxToolTipText":
			res.valueString = inferedInstance.CheckboxToolTipText
		case "CheckboxToolTipPosition":
			enum := inferedInstance.CheckboxToolTipPosition
			res.valueString = enum.ToCodeString()
		case "HasSecondCheckboxButton":
			res.valueString = fmt.Sprintf("%t", inferedInstance.HasSecondCheckboxButton)
			res.valueBool = inferedInstance.HasSecondCheckboxButton
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsSecondCheckboxChecked":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSecondCheckboxChecked)
			res.valueBool = inferedInstance.IsSecondCheckboxChecked
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsSecondCheckboxDisabled":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsSecondCheckboxDisabled)
			res.valueBool = inferedInstance.IsSecondCheckboxDisabled
			res.GongFieldValueType = GongFieldValueTypeBool
		case "TextAfterSecondCheckbox":
			res.valueString = inferedInstance.TextAfterSecondCheckbox
		case "SecondCheckboxHasToolTip":
			res.valueString = fmt.Sprintf("%t", inferedInstance.SecondCheckboxHasToolTip)
			res.valueBool = inferedInstance.SecondCheckboxHasToolTip
			res.GongFieldValueType = GongFieldValueTypeBool
		case "SecondCheckboxToolTipText":
			res.valueString = inferedInstance.SecondCheckboxToolTipText
		case "SecondCheckboxToolTipPosition":
			enum := inferedInstance.SecondCheckboxToolTipPosition
			res.valueString = enum.ToCodeString()
		case "IsInEditMode":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsInEditMode)
			res.valueBool = inferedInstance.IsInEditMode
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsNodeClickable":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsNodeClickable)
			res.valueBool = inferedInstance.IsNodeClickable
			res.GongFieldValueType = GongFieldValueTypeBool
		case "IsWithPreceedingIcon":
			res.valueString = fmt.Sprintf("%t", inferedInstance.IsWithPreceedingIcon)
			res.valueBool = inferedInstance.IsWithPreceedingIcon
			res.GongFieldValueType = GongFieldValueTypeBool
		case "PreceedingIcon":
			res.valueString = inferedInstance.PreceedingIcon
		case "PreceedingSVGIcon":
			if inferedInstance.PreceedingSVGIcon != nil {
				res.valueString = inferedInstance.PreceedingSVGIcon.Name
			}
		case "Children":
			for idx, __instance__ := range inferedInstance.Children {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "Buttons":
			for idx, __instance__ := range inferedInstance.Buttons {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case SVGIcon:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "SVG":
			res.valueString = inferedInstance.SVG
		}
	case Tree:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "RootNodes":
			for idx, __instance__ := range inferedInstance.RootNodes {
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
