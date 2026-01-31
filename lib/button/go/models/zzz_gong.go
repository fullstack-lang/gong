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
type Stage struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between
	// succesive commit
	isInDeltaMode bool

	// insertion point for definition of arrays registering instances
	Buttons                map[*Button]struct{}
	Buttons_reference      map[*Button]*Button
	Buttons_referenceOrder map[*Button]uint // diff Unstage needs the reference order
	Buttons_mapString      map[string]*Button

	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback OnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback OnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback OnAfterDeleteInterface[Button]
	OnAfterButtonReadCallback   OnAfterReadInterface[Button]

	ButtonToggles                map[*ButtonToggle]struct{}
	ButtonToggles_reference      map[*ButtonToggle]*ButtonToggle
	ButtonToggles_referenceOrder map[*ButtonToggle]uint // diff Unstage needs the reference order
	ButtonToggles_mapString      map[string]*ButtonToggle

	// insertion point for slice of pointers maps
	OnAfterButtonToggleCreateCallback OnAfterCreateInterface[ButtonToggle]
	OnAfterButtonToggleUpdateCallback OnAfterUpdateInterface[ButtonToggle]
	OnAfterButtonToggleDeleteCallback OnAfterDeleteInterface[ButtonToggle]
	OnAfterButtonToggleReadCallback   OnAfterReadInterface[ButtonToggle]

	Groups                map[*Group]struct{}
	Groups_reference      map[*Group]*Group
	Groups_referenceOrder map[*Group]uint // diff Unstage needs the reference order
	Groups_mapString      map[string]*Group

	// insertion point for slice of pointers maps
	Group_Buttons_reverseMap map[*Button]*Group

	OnAfterGroupCreateCallback OnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback OnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback OnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   OnAfterReadInterface[Group]

	GroupToogles                map[*GroupToogle]struct{}
	GroupToogles_reference      map[*GroupToogle]*GroupToogle
	GroupToogles_referenceOrder map[*GroupToogle]uint // diff Unstage needs the reference order
	GroupToogles_mapString      map[string]*GroupToogle

	// insertion point for slice of pointers maps
	GroupToogle_ButtonToggles_reverseMap map[*ButtonToggle]*GroupToogle

	OnAfterGroupToogleCreateCallback OnAfterCreateInterface[GroupToogle]
	OnAfterGroupToogleUpdateCallback OnAfterUpdateInterface[GroupToogle]
	OnAfterGroupToogleDeleteCallback OnAfterDeleteInterface[GroupToogle]
	OnAfterGroupToogleReadCallback   OnAfterReadInterface[GroupToogle]

	Layouts                map[*Layout]struct{}
	Layouts_reference      map[*Layout]*Layout
	Layouts_referenceOrder map[*Layout]uint // diff Unstage needs the reference order
	Layouts_mapString      map[string]*Layout

	// insertion point for slice of pointers maps
	Layout_Groups_reverseMap map[*Group]*Layout

	Layout_GroupToogles_reverseMap map[*GroupToogle]*Layout

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

	ButtonToggleOrder            uint
	ButtonToggleMap_Staged_Order map[*ButtonToggle]uint

	GroupOrder            uint
	GroupMap_Staged_Order map[*Group]uint

	GroupToogleOrder            uint
	GroupToogleMap_Staged_Order map[*GroupToogle]uint

	LayoutOrder            uint
	LayoutMap_Staged_Order map[*Layout]uint

	// end of insertion point

	NamedStructs []*NamedStruct

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]ModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF ProbeIF

	forwardCommits  []string
	backwardCommits []string

	// when navigating the commit history
	// navigationMode is set to Navigating
	navigationMode gongStageNavigationMode
	commitsBehind  int // the number of commits the stage is behind the front of the history
}

type gongStageNavigationMode string

const (
	GongNavigationModeNormal gongStageNavigationMode = "Normal"
	// when the mode is navigating, each commit backward and forward
	// it is possible to go apply the nbCommitsBackward forward commits
	GongNavigationModeNavigating gongStageNavigationMode = "Navigating"
)

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {

	if len(stage.backwardCommits) == 0 {
		return errors.New("no backward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	if stage.commitsBehind >= len(stage.backwardCommits) {
		return errors.New("no more backward commit to apply")
	}

	commitToApply := stage.backwardCommits[len(stage.backwardCommits)-1-stage.commitsBehind]

	// umarshall the backward commit to the stage

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind++
	err := GongParseAstString(stage, commitToApply, true)
	if err != nil {
		log.Println("error during ApplyBackwardCommit: ", err)
		return err
	}

	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetForwardCommits() []string {
	return stage.forwardCommits
}

func (stage *Stage) GetBackwardCommits() []string {
	return stage.backwardCommits
}

func (stage *Stage) ApplyForwardCommit() error {
	if stage.navigationMode == GongNavigationModeNormal && stage.commitsBehind != 0 {
		return errors.New("in navigation mode normal, cannot have commitsBehind != 0")
	}

	if stage.commitsBehind == 0 {
		return errors.New("no more forward commit to apply")
	}

	if stage.navigationMode == GongNavigationModeNormal {
		stage.navigationMode = GongNavigationModeNavigating
	}

	commitToApply := stage.forwardCommits[len(stage.forwardCommits)-1-stage.commitsBehind+1]

	// the parsing of the commit will call the UX update
	// therefore, it is important to stage.commitsBehind before because it is used in the
	// UX
	stage.commitsBehind--
	err := GongParseAstString(stage, commitToApply, true)
	if err != nil {
		log.Println("error during ApplyForwardCommit: ", err)
		return err
	}
	stage.ComputeReferenceAndOrders()

	return nil
}

func (stage *Stage) GetCommitsBehind() int {
	return stage.commitsBehind
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {

	newCommitsLen := len(stage.forwardCommits) - stage.GetCommitsBehind()

	stage.forwardCommits = stage.forwardCommits[:newCommitsLen]
	stage.backwardCommits = stage.backwardCommits[:newCommitsLen]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation 
	var maxButtonOrder uint
	var foundButton bool
	for _, order := range stage.ButtonMap_Staged_Order {
		if !foundButton || order > maxButtonOrder {
			maxButtonOrder = order
			foundButton = true
		}
	}
	if foundButton {
		stage.ButtonOrder = maxButtonOrder + 1
	} else {
		stage.ButtonOrder = 0
	}

	var maxButtonToggleOrder uint
	var foundButtonToggle bool
	for _, order := range stage.ButtonToggleMap_Staged_Order {
		if !foundButtonToggle || order > maxButtonToggleOrder {
			maxButtonToggleOrder = order
			foundButtonToggle = true
		}
	}
	if foundButtonToggle {
		stage.ButtonToggleOrder = maxButtonToggleOrder + 1
	} else {
		stage.ButtonToggleOrder = 0
	}

	var maxGroupOrder uint
	var foundGroup bool
	for _, order := range stage.GroupMap_Staged_Order {
		if !foundGroup || order > maxGroupOrder {
			maxGroupOrder = order
			foundGroup = true
		}
	}
	if foundGroup {
		stage.GroupOrder = maxGroupOrder + 1
	} else {
		stage.GroupOrder = 0
	}

	var maxGroupToogleOrder uint
	var foundGroupToogle bool
	for _, order := range stage.GroupToogleMap_Staged_Order {
		if !foundGroupToogle || order > maxGroupToogleOrder {
			maxGroupToogleOrder = order
			foundGroupToogle = true
		}
	}
	if foundGroupToogle {
		stage.GroupToogleOrder = maxGroupToogleOrder + 1
	} else {
		stage.GroupToogleOrder = 0
	}

	var maxLayoutOrder uint
	var foundLayout bool
	for _, order := range stage.LayoutMap_Staged_Order {
		if !foundLayout || order > maxLayoutOrder {
			maxLayoutOrder = order
			foundLayout = true
		}
	}
	if foundLayout {
		stage.LayoutOrder = maxLayoutOrder + 1
	} else {
		stage.LayoutOrder = 0
	}

}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF ProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() ProbeIF {
	if stage.probeIF == nil {
		return nil
	}

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
	case *ButtonToggle:
		tmp := GetStructInstancesByOrder(stage.ButtonToggles, stage.ButtonToggleMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *ButtonToggle implements.
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
	case *GroupToogle:
		tmp := GetStructInstancesByOrder(stage.GroupToogles, stage.GroupToogleMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *GroupToogle implements.
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
	case "Button":
		res = GetNamedStructInstances(stage.Buttons, stage.ButtonMap_Staged_Order)
	case "ButtonToggle":
		res = GetNamedStructInstances(stage.ButtonToggles, stage.ButtonToggleMap_Staged_Order)
	case "Group":
		res = GetNamedStructInstances(stage.Groups, stage.GroupMap_Staged_Order)
	case "GroupToogle":
		res = GetNamedStructInstances(stage.GroupToogles, stage.GroupToogleMap_Staged_Order)
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
	CommitButtonToggle(buttontoggle *ButtonToggle)
	CheckoutButtonToggle(buttontoggle *ButtonToggle)
	CommitGroup(group *Group)
	CheckoutGroup(group *Group)
	CommitGroupToogle(grouptoogle *GroupToogle)
	CheckoutGroupToogle(grouptoogle *GroupToogle)
	CommitLayout(layout *Layout)
	CheckoutLayout(layout *Layout)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(name string) (stage *Stage) {

	stage = &Stage{ // insertion point for array initiatialisation
		Buttons:           make(map[*Button]struct{}),
		Buttons_mapString: make(map[string]*Button),

		ButtonToggles:           make(map[*ButtonToggle]struct{}),
		ButtonToggles_mapString: make(map[string]*ButtonToggle),

		Groups:           make(map[*Group]struct{}),
		Groups_mapString: make(map[string]*Group),

		GroupToogles:           make(map[*GroupToogle]struct{}),
		GroupToogles_mapString: make(map[string]*GroupToogle),

		Layouts:           make(map[*Layout]struct{}),
		Layouts_mapString: make(map[string]*Layout),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		ButtonMap_Staged_Order: make(map[*Button]uint),

		ButtonToggleMap_Staged_Order: make(map[*ButtonToggle]uint),

		GroupMap_Staged_Order: make(map[*Group]uint),

		GroupToogleMap_Staged_Order: make(map[*GroupToogle]uint),

		LayoutMap_Staged_Order: make(map[*Layout]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"Button": &ButtonUnmarshaller{},

			"ButtonToggle": &ButtonToggleUnmarshaller{},

			"Group": &GroupUnmarshaller{},

			"GroupToogle": &GroupToogleUnmarshaller{},

			"Layout": &LayoutUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Button"},
			{name: "ButtonToggle"},
			{name: "Group"},
			{name: "GroupToogle"},
			{name: "Layout"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {

	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Button:
		return stage.ButtonMap_Staged_Order[instance]
	case *ButtonToggle:
		return stage.ButtonToggleMap_Staged_Order[instance]
	case *Group:
		return stage.GroupMap_Staged_Order[instance]
	case *GroupToogle:
		return stage.GroupToogleMap_Staged_Order[instance]
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
	case *ButtonToggle:
		return stage.ButtonToggleMap_Staged_Order[instance]
	case *Group:
		return stage.GroupMap_Staged_Order[instance]
	case *GroupToogle:
		return stage.GroupToogleMap_Staged_Order[instance]
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

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.isInDeltaMode && stage.navigationMode == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
		stage.ResetHard()
	}

	if stage.IsInDeltaMode() {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().Refresh()
		}
	}
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["ButtonToggle"] = len(stage.ButtonToggles)
	stage.Map_GongStructName_InstancesNb["Group"] = len(stage.Groups)
	stage.Map_GongStructName_InstancesNb["GroupToogle"] = len(stage.GroupToogles)
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
		stage.Buttons[button] = struct{}{}
		stage.ButtonMap_Staged_Order[button] = stage.ButtonOrder
		stage.ButtonOrder++
	}
	stage.Buttons_mapString[button.Name] = button

	return button
}

// StagePreserveOrder puts button to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ButtonOrder
// - update stage.ButtonOrder accordingly
func (button *Button) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.Buttons[button]; !ok {
		stage.Buttons[button] = struct{}{}

		if order > stage.ButtonOrder {
			stage.ButtonOrder = order
		}
		stage.ButtonMap_Staged_Order[button] = order
		stage.ButtonOrder++
	}
	stage.Buttons_mapString[button.Name] = button
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	delete(stage.Buttons, button)
	delete(stage.ButtonMap_Staged_Order, button)
	delete(stage.Buttons_mapString, button.Name)

	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *Stage) {
	delete(stage.Buttons, button)
	delete(stage.ButtonMap_Staged_Order, button)
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

// for satisfaction of GongStruct interface
func (button *Button) SetName(name string) {
	button.Name = name
}

// Stage puts buttontoggle to the model stage
func (buttontoggle *ButtonToggle) Stage(stage *Stage) *ButtonToggle {

	if _, ok := stage.ButtonToggles[buttontoggle]; !ok {
		stage.ButtonToggles[buttontoggle] = struct{}{}
		stage.ButtonToggleMap_Staged_Order[buttontoggle] = stage.ButtonToggleOrder
		stage.ButtonToggleOrder++
	}
	stage.ButtonToggles_mapString[buttontoggle.Name] = buttontoggle

	return buttontoggle
}

// StagePreserveOrder puts buttontoggle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ButtonToggleOrder
// - update stage.ButtonToggleOrder accordingly
func (buttontoggle *ButtonToggle) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.ButtonToggles[buttontoggle]; !ok {
		stage.ButtonToggles[buttontoggle] = struct{}{}

		if order > stage.ButtonToggleOrder {
			stage.ButtonToggleOrder = order
		}
		stage.ButtonToggleMap_Staged_Order[buttontoggle] = order
		stage.ButtonToggleOrder++
	}
	stage.ButtonToggles_mapString[buttontoggle.Name] = buttontoggle
}

// Unstage removes buttontoggle off the model stage
func (buttontoggle *ButtonToggle) Unstage(stage *Stage) *ButtonToggle {
	delete(stage.ButtonToggles, buttontoggle)
	delete(stage.ButtonToggleMap_Staged_Order, buttontoggle)
	delete(stage.ButtonToggles_mapString, buttontoggle.Name)

	return buttontoggle
}

// UnstageVoid removes buttontoggle off the model stage
func (buttontoggle *ButtonToggle) UnstageVoid(stage *Stage) {
	delete(stage.ButtonToggles, buttontoggle)
	delete(stage.ButtonToggleMap_Staged_Order, buttontoggle)
	delete(stage.ButtonToggles_mapString, buttontoggle.Name)
}

// commit buttontoggle to the back repo (if it is already staged)
func (buttontoggle *ButtonToggle) Commit(stage *Stage) *ButtonToggle {
	if _, ok := stage.ButtonToggles[buttontoggle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitButtonToggle(buttontoggle)
		}
	}
	return buttontoggle
}

func (buttontoggle *ButtonToggle) CommitVoid(stage *Stage) {
	buttontoggle.Commit(stage)
}

func (buttontoggle *ButtonToggle) StageVoid(stage *Stage) {
	buttontoggle.Stage(stage)
}

// Checkout buttontoggle to the back repo (if it is already staged)
func (buttontoggle *ButtonToggle) Checkout(stage *Stage) *ButtonToggle {
	if _, ok := stage.ButtonToggles[buttontoggle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutButtonToggle(buttontoggle)
		}
	}
	return buttontoggle
}

// for satisfaction of GongStruct interface
func (buttontoggle *ButtonToggle) GetName() (res string) {
	return buttontoggle.Name
}

// for satisfaction of GongStruct interface
func (buttontoggle *ButtonToggle) SetName(name string) {
	buttontoggle.Name = name
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
		stage.GroupMap_Staged_Order[group] = order
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

// Stage puts grouptoogle to the model stage
func (grouptoogle *GroupToogle) Stage(stage *Stage) *GroupToogle {

	if _, ok := stage.GroupToogles[grouptoogle]; !ok {
		stage.GroupToogles[grouptoogle] = struct{}{}
		stage.GroupToogleMap_Staged_Order[grouptoogle] = stage.GroupToogleOrder
		stage.GroupToogleOrder++
	}
	stage.GroupToogles_mapString[grouptoogle.Name] = grouptoogle

	return grouptoogle
}

// StagePreserveOrder puts grouptoogle to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.GroupToogleOrder
// - update stage.GroupToogleOrder accordingly
func (grouptoogle *GroupToogle) StagePreserveOrder(stage *Stage, order uint) {

	if _, ok := stage.GroupToogles[grouptoogle]; !ok {
		stage.GroupToogles[grouptoogle] = struct{}{}

		if order > stage.GroupToogleOrder {
			stage.GroupToogleOrder = order
		}
		stage.GroupToogleMap_Staged_Order[grouptoogle] = order
		stage.GroupToogleOrder++
	}
	stage.GroupToogles_mapString[grouptoogle.Name] = grouptoogle
}

// Unstage removes grouptoogle off the model stage
func (grouptoogle *GroupToogle) Unstage(stage *Stage) *GroupToogle {
	delete(stage.GroupToogles, grouptoogle)
	delete(stage.GroupToogleMap_Staged_Order, grouptoogle)
	delete(stage.GroupToogles_mapString, grouptoogle.Name)

	return grouptoogle
}

// UnstageVoid removes grouptoogle off the model stage
func (grouptoogle *GroupToogle) UnstageVoid(stage *Stage) {
	delete(stage.GroupToogles, grouptoogle)
	delete(stage.GroupToogleMap_Staged_Order, grouptoogle)
	delete(stage.GroupToogles_mapString, grouptoogle.Name)
}

// commit grouptoogle to the back repo (if it is already staged)
func (grouptoogle *GroupToogle) Commit(stage *Stage) *GroupToogle {
	if _, ok := stage.GroupToogles[grouptoogle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitGroupToogle(grouptoogle)
		}
	}
	return grouptoogle
}

func (grouptoogle *GroupToogle) CommitVoid(stage *Stage) {
	grouptoogle.Commit(stage)
}

func (grouptoogle *GroupToogle) StageVoid(stage *Stage) {
	grouptoogle.Stage(stage)
}

// Checkout grouptoogle to the back repo (if it is already staged)
func (grouptoogle *GroupToogle) Checkout(stage *Stage) *GroupToogle {
	if _, ok := stage.GroupToogles[grouptoogle]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutGroupToogle(grouptoogle)
		}
	}
	return grouptoogle
}

// for satisfaction of GongStruct interface
func (grouptoogle *GroupToogle) GetName() (res string) {
	return grouptoogle.Name
}

// for satisfaction of GongStruct interface
func (grouptoogle *GroupToogle) SetName(name string) {
	grouptoogle.Name = name
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
		stage.LayoutMap_Staged_Order[layout] = order
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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMButton(Button *Button)
	CreateORMButtonToggle(ButtonToggle *ButtonToggle)
	CreateORMGroup(Group *Group)
	CreateORMGroupToogle(GroupToogle *GroupToogle)
	CreateORMLayout(Layout *Layout)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMButton(Button *Button)
	DeleteORMButtonToggle(ButtonToggle *ButtonToggle)
	DeleteORMGroup(Group *Group)
	DeleteORMGroupToogle(GroupToogle *GroupToogle)
	DeleteORMLayout(Layout *Layout)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Buttons = make(map[*Button]struct{})
	stage.Buttons_mapString = make(map[string]*Button)
	stage.ButtonMap_Staged_Order = make(map[*Button]uint)
	stage.ButtonOrder = 0

	stage.ButtonToggles = make(map[*ButtonToggle]struct{})
	stage.ButtonToggles_mapString = make(map[string]*ButtonToggle)
	stage.ButtonToggleMap_Staged_Order = make(map[*ButtonToggle]uint)
	stage.ButtonToggleOrder = 0

	stage.Groups = make(map[*Group]struct{})
	stage.Groups_mapString = make(map[string]*Group)
	stage.GroupMap_Staged_Order = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.GroupToogles = make(map[*GroupToogle]struct{})
	stage.GroupToogles_mapString = make(map[string]*GroupToogle)
	stage.GroupToogleMap_Staged_Order = make(map[*GroupToogle]uint)
	stage.GroupToogleOrder = 0

	stage.Layouts = make(map[*Layout]struct{})
	stage.Layouts_mapString = make(map[string]*Layout)
	stage.LayoutMap_Staged_Order = make(map[*Layout]uint)
	stage.LayoutOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.Buttons = nil
	stage.Buttons_mapString = nil

	stage.ButtonToggles = nil
	stage.ButtonToggles_mapString = nil

	stage.Groups = nil
	stage.Groups_mapString = nil

	stage.GroupToogles = nil
	stage.GroupToogles_mapString = nil

	stage.Layouts = nil
	stage.Layouts_mapString = nil

}

func (stage *Stage) Unstage() { // insertion point for array nil
	for button := range stage.Buttons {
		button.Unstage(stage)
	}

	for buttontoggle := range stage.ButtonToggles {
		buttontoggle.Unstage(stage)
	}

	for group := range stage.Groups {
		group.Unstage(stage)
	}

	for grouptoogle := range stage.GroupToogles {
		grouptoogle.Unstage(stage)
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
	SetName(string)
	CommitVoid(*Stage)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongClean(stage *Stage) (modified bool)
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
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
	case map[*Button]any:
		return any(&stage.Buttons).(*Type)
	case map[*ButtonToggle]any:
		return any(&stage.ButtonToggles).(*Type)
	case map[*Group]any:
		return any(&stage.Groups).(*Type)
	case map[*GroupToogle]any:
		return any(&stage.GroupToogles).(*Type)
	case map[*Layout]any:
		return any(&stage.Layouts).(*Type)
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
	case *Button:
		return any(stage.Buttons_mapString).(map[string]Type)
	case *ButtonToggle:
		return any(stage.ButtonToggles_mapString).(map[string]Type)
	case *Group:
		return any(stage.Groups_mapString).(map[string]Type)
	case *GroupToogle:
		return any(stage.GroupToogles_mapString).(map[string]Type)
	case *Layout:
		return any(stage.Layouts_mapString).(map[string]Type)
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
	case Button:
		return any(&stage.Buttons).(*map[*Type]struct{})
	case ButtonToggle:
		return any(&stage.ButtonToggles).(*map[*Type]struct{})
	case Group:
		return any(&stage.Groups).(*map[*Type]struct{})
	case GroupToogle:
		return any(&stage.GroupToogles).(*map[*Type]struct{})
	case Layout:
		return any(&stage.Layouts).(*map[*Type]struct{})
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
	case *Button:
		return any(&stage.Buttons).(*map[Type]struct{})
	case *ButtonToggle:
		return any(&stage.ButtonToggles).(*map[Type]struct{})
	case *Group:
		return any(&stage.Groups).(*map[Type]struct{})
	case *GroupToogle:
		return any(&stage.GroupToogles).(*map[Type]struct{})
	case *Layout:
		return any(&stage.Layouts).(*map[Type]struct{})
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
	case Button:
		return any(&stage.Buttons_mapString).(*map[string]*Type)
	case ButtonToggle:
		return any(&stage.ButtonToggles_mapString).(*map[string]*Type)
	case Group:
		return any(&stage.Groups_mapString).(*map[string]*Type)
	case GroupToogle:
		return any(&stage.GroupToogles_mapString).(*map[string]*Type)
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
	case ButtonToggle:
		return any(&ButtonToggle{
			// Initialisation of associations
		}).(*Type)
	case Group:
		return any(&Group{
			// Initialisation of associations
			// field is initialized with an instance of Button with the name of the field
			Buttons: []*Button{{Name: "Buttons"}},
		}).(*Type)
	case GroupToogle:
		return any(&GroupToogle{
			// Initialisation of associations
			// field is initialized with an instance of ButtonToggle with the name of the field
			ButtonToggles: []*ButtonToggle{{Name: "ButtonToggles"}},
		}).(*Type)
	case Layout:
		return any(&Layout{
			// Initialisation of associations
			// field is initialized with an instance of Group with the name of the field
			Groups: []*Group{{Name: "Groups"}},
			// field is initialized with an instance of GroupToogle with the name of the field
			GroupToogles: []*GroupToogle{{Name: "GroupToogles"}},
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
	// reverse maps of direct associations of ButtonToggle
	case ButtonToggle:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Group
	case Group:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of GroupToogle
	case GroupToogle:
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
	// reverse maps of direct associations of ButtonToggle
	case ButtonToggle:
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
	// reverse maps of direct associations of GroupToogle
	case GroupToogle:
		switch fieldname {
		// insertion point for per direct association field
		case "ButtonToggles":
			res := make(map[*ButtonToggle][]*GroupToogle)
			for grouptoogle := range stage.GroupToogles {
				for _, buttontoggle_ := range grouptoogle.ButtonToggles {
					res[buttontoggle_] = append(res[buttontoggle_], grouptoogle)
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
		case "GroupToogles":
			res := make(map[*GroupToogle][]*Layout)
			for layout := range stage.Layouts {
				for _, grouptoogle_ := range layout.GroupToogles {
					res[grouptoogle_] = append(res[grouptoogle_], layout)
				}
			}
			return any(res).(map[*End][]*Start)
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
	case *Button:
		res = "Button"
	case *ButtonToggle:
		res = "ButtonToggle"
	case *Group:
		res = "Group"
	case *GroupToogle:
		res = "GroupToogle"
	case *Layout:
		res = "Layout"
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
	case *Button:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Group"
		rf.Fieldname = "Buttons"
		res = append(res, rf)
	case *ButtonToggle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "GroupToogle"
		rf.Fieldname = "ButtonToggles"
		res = append(res, rf)
	case *Group:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layout"
		rf.Fieldname = "Groups"
		res = append(res, rf)
	case *GroupToogle:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Layout"
		rf.Fieldname = "GroupToogles"
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

func (buttontoggle *ButtonToggle) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:               "IsChecked",
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
			Name:                 "Buttons",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Button",
		},
		{
			Name:               "NbColumns",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (grouptoogle *GroupToogle) GongGetFieldHeaders() (res []GongFieldHeader) {
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
			Name:                 "ButtonToggles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "ButtonToggle",
		},
		{
			Name:               "IsSingleSelector",
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
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:                 "GroupToogles",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "GroupToogle",
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
func (button *Button) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
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
func (buttontoggle *ButtonToggle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = buttontoggle.Name
	case "Label":
		res.valueString = buttontoggle.Label
	case "Icon":
		res.valueString = buttontoggle.Icon
	case "IsDisabled":
		res.valueString = fmt.Sprintf("%t", buttontoggle.IsDisabled)
		res.valueBool = buttontoggle.IsDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", buttontoggle.IsChecked)
		res.valueBool = buttontoggle.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
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
func (grouptoogle *GroupToogle) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = grouptoogle.Name
	case "Percentage":
		res.valueString = fmt.Sprintf("%f", grouptoogle.Percentage)
		res.valueFloat = grouptoogle.Percentage
		res.GongFieldValueType = GongFieldValueTypeFloat
	case "ButtonToggles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range grouptoogle.ButtonToggles {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "IsSingleSelector":
		res.valueString = fmt.Sprintf("%t", grouptoogle.IsSingleSelector)
		res.valueBool = grouptoogle.IsSingleSelector
		res.GongFieldValueType = GongFieldValueTypeBool
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
	case "GroupToogles":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range layout.GroupToogles {
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

	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

// insertion point for generic set gongstruct field value
func (button *Button) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		button.Name = value.GetValueString()
	case "Label":
		button.Label = value.GetValueString()
	case "Icon":
		button.Icon = value.GetValueString()
	case "IsDisabled":
		button.IsDisabled = value.GetValueBool()
	case "Color":
		button.Color.FromCodeString(value.GetValueString())
	case "MatButtonType":
		button.MatButtonType.FromCodeString(value.GetValueString())
	case "MatButtonAppearance":
		button.MatButtonAppearance.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (buttontoggle *ButtonToggle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		buttontoggle.Name = value.GetValueString()
	case "Label":
		buttontoggle.Label = value.GetValueString()
	case "Icon":
		buttontoggle.Icon = value.GetValueString()
	case "IsDisabled":
		buttontoggle.IsDisabled = value.GetValueBool()
	case "IsChecked":
		buttontoggle.IsChecked = value.GetValueBool()
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
	case "Buttons":
		group.Buttons = make([]*Button, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Buttons {
					if stage.ButtonMap_Staged_Order[__instance__] == uint(id) {
						group.Buttons = append(group.Buttons, __instance__)
						break
					}
				}
			}
		}
	case "NbColumns":
		group.NbColumns = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (grouptoogle *GroupToogle) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		grouptoogle.Name = value.GetValueString()
	case "Percentage":
		grouptoogle.Percentage = value.GetValueFloat()
	case "ButtonToggles":
		grouptoogle.ButtonToggles = make([]*ButtonToggle, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.ButtonToggles {
					if stage.ButtonToggleMap_Staged_Order[__instance__] == uint(id) {
						grouptoogle.ButtonToggles = append(grouptoogle.ButtonToggles, __instance__)
						break
					}
				}
			}
		}
	case "IsSingleSelector":
		grouptoogle.IsSingleSelector = value.GetValueBool()
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
	case "GroupToogles":
		layout.GroupToogles = make([]*GroupToogle, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.GroupToogles {
					if stage.GroupToogleMap_Staged_Order[__instance__] == uint(id) {
						layout.GroupToogles = append(layout.GroupToogles, __instance__)
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

func SetFieldStringValueFromPointer(instance GongstructIF, fieldName string, value GongFieldValue, stage *Stage) error {
	return instance.GongSetFieldValue(fieldName, value, stage)
}

// insertion point for generic get gongstruct name
func (button *Button) GongGetGongstructName() string {
	return "Button"
}

func (buttontoggle *ButtonToggle) GongGetGongstructName() string {
	return "ButtonToggle"
}

func (group *Group) GongGetGongstructName() string {
	return "Group"
}

func (grouptoogle *GroupToogle) GongGetGongstructName() string {
	return "GroupToogle"
}

func (layout *Layout) GongGetGongstructName() string {
	return "Layout"
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {

	// insertion point for generic get gongstruct name
	stage.Buttons_mapString = make(map[string]*Button)
	for button := range stage.Buttons {
		stage.Buttons_mapString[button.Name] = button
	}

	stage.ButtonToggles_mapString = make(map[string]*ButtonToggle)
	for buttontoggle := range stage.ButtonToggles {
		stage.ButtonToggles_mapString[buttontoggle.Name] = buttontoggle
	}

	stage.Groups_mapString = make(map[string]*Group)
	for group := range stage.Groups {
		stage.Groups_mapString[group.Name] = group
	}

	stage.GroupToogles_mapString = make(map[string]*GroupToogle)
	for grouptoogle := range stage.GroupToogles {
		stage.GroupToogles_mapString[grouptoogle.Name] = grouptoogle
	}

	stage.Layouts_mapString = make(map[string]*Layout)
	for layout := range stage.Layouts {
		stage.Layouts_mapString[layout.Name] = layout
	}

}

// Last line of the template
