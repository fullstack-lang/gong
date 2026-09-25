// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strings"
	"sync"
	"time"
)

// can be used for
//
//	days := __gong__abs(int(int(inferedInstance.ComputedDuration.Hours()) / 24))
func __gong__abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	_ = __gong__abs
	_ = strings.Clone("")
)

const (
	GongProbeTreeSidebarSuffix           = ":sidebar of the probe"
	GongProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	GongProbeTableSuffix                 = ":table of the probe"
	GongProbeNotificationTableSuffix     = ":notification table of the probe"
	GongProbeFormSuffix                  = ":form of the probe"
	GongProbeSplitSuffix                 = ":probe of the probe"
	GongProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNavigationTreeSidebarSuffix
}

func (stage *Stage) GetProbeFormStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeFormSuffix
}

func (stage *Stage) GetProbeTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeTableSuffix
}

func (stage *Stage) GetProbeNotificationTableStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeNotificationTableSuffix
}

func (stage *Stage) GetProbeSplitStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeSplitSuffix
}

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + GongProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var _ = fmt.Sprintf

// idem for math package when not need by generated code
var _ = math.E

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

// MetaPackageImport represents a package import needed by a meta/diagram file
type MetaPackageImport struct {
	Alias string
	Path  string
}

// Stage enables storage of staged instances
type Stage struct {
	name string

	// isInDeltaMode is true when the stage is used to compute difference between
	// succesive commit
	isInDeltaMode bool

	// gongMarshallingMode set the marshalling mode
	gongMarshallingMode GongMarshallingMode
	// some stages have semantic rules that forbids them to be empty
	// like for git, the commit #0 (genesis commit) cannot be rolled back
	isWithGenesisCommit bool

	// insertion point for definition of arrays registering instances
	Checkboxs                map[*Checkbox]struct{}
	Checkboxs_instance       map[*Checkbox]*Checkbox
	Checkboxs_mapString      map[string]*Checkbox
	CheckboxOrder            uint
	Checkbox_stagedOrder     map[*Checkbox]uint
	Checkbox_orderStaged     map[uint]*Checkbox
	Checkboxs_reference      map[*Checkbox]*Checkbox
	Checkboxs_referenceOrder map[*Checkbox]uint

	// insertion point for slice of pointers maps
	OnAfterCheckboxCreateCallback GongOnAfterCreateInterface[Checkbox]
	OnAfterCheckboxUpdateCallback GongOnAfterUpdateInterface[Checkbox]
	OnAfterCheckboxDeleteCallback GongOnAfterDeleteInterface[Checkbox]
	OnAfterCheckboxReadCallback   GongOnAfterReadInterface[Checkbox]

	Groups                map[*Group]struct{}
	Groups_instance       map[*Group]*Group
	Groups_mapString      map[string]*Group
	GroupOrder            uint
	Group_stagedOrder     map[*Group]uint
	Group_orderStaged     map[uint]*Group
	Groups_reference      map[*Group]*Group
	Groups_referenceOrder map[*Group]uint

	// insertion point for slice of pointers maps
	Group_Sliders_reverseMap map[*Slider]*Group

	Group_Checkboxes_reverseMap map[*Checkbox]*Group

	OnAfterGroupCreateCallback GongOnAfterCreateInterface[Group]
	OnAfterGroupUpdateCallback GongOnAfterUpdateInterface[Group]
	OnAfterGroupDeleteCallback GongOnAfterDeleteInterface[Group]
	OnAfterGroupReadCallback   GongOnAfterReadInterface[Group]

	Layouts                map[*Layout]struct{}
	Layouts_instance       map[*Layout]*Layout
	Layouts_mapString      map[string]*Layout
	LayoutOrder            uint
	Layout_stagedOrder     map[*Layout]uint
	Layout_orderStaged     map[uint]*Layout
	Layouts_reference      map[*Layout]*Layout
	Layouts_referenceOrder map[*Layout]uint

	// insertion point for slice of pointers maps
	Layout_Groups_reverseMap map[*Group]*Layout

	OnAfterLayoutCreateCallback GongOnAfterCreateInterface[Layout]
	OnAfterLayoutUpdateCallback GongOnAfterUpdateInterface[Layout]
	OnAfterLayoutDeleteCallback GongOnAfterDeleteInterface[Layout]
	OnAfterLayoutReadCallback   GongOnAfterReadInterface[Layout]

	Sliders                map[*Slider]struct{}
	Sliders_instance       map[*Slider]*Slider
	Sliders_mapString      map[string]*Slider
	SliderOrder            uint
	Slider_stagedOrder     map[*Slider]uint
	Slider_orderStaged     map[uint]*Slider
	Sliders_reference      map[*Slider]*Slider
	Sliders_referenceOrder map[*Slider]uint

	// insertion point for slice of pointers maps
	OnAfterSliderCreateCallback GongOnAfterCreateInterface[Slider]
	OnAfterSliderUpdateCallback GongOnAfterUpdateInterface[Slider]
	OnAfterSliderDeleteCallback GongOnAfterDeleteInterface[Slider]
	OnAfterSliderReadCallback   GongOnAfterReadInterface[Slider]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string
	MetaPackageImports     []*MetaPackageImport

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here

	// store the stage order of each instance in order to
	// preserve this order when serializing them
	// insertion point for order fields declaration
	// end of insertion point

	// GongUnmarshallers is the registry of all model unmarshallers
	GongUnmarshallers map[string]GongModelUnmarshaller

	// probeIF is the interface to the probe that allows log
	// commit event to the probe
	probeIF GongProbeIF

	forwardCommits  []string
	backwardCommits []string

	// when navigating the commit history
	// navigationMode is set to Navigating
	navigationMode gongStageNavigationMode
	commitsBehind  int // the number of commits the stage is behind the front of the history

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool

	modified bool

	lock sync.RWMutex
}

type GongStage = Stage

func (s *Stage) SetGongMarshallingMode(mode GongMarshallingMode) {
	s.gongMarshallingMode = mode
}

func (s *Stage) GetGongMarshallingMode() GongMarshallingMode {
	return s.gongMarshallingMode
}

func (s *Stage) SetIsWithGenesisCommit(isWithGenesisCommit bool) {
	s.isWithGenesisCommit = isWithGenesisCommit
}

func (s *Stage) GetIsWithGenesisCommit() bool {
	return s.isWithGenesisCommit
}

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.beforeCommitHooks = append(s.beforeCommitHooks, hook)
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.afterCommitHooks = append(s.afterCommitHooks, hook)
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

	if stage.isWithGenesisCommit && stage.commitsBehind >= len(stage.backwardCommits)-1 {
		return errors.New("cannot rollback genesis commit")
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
	stage.isApplyingBackwardCommit = true
	err := stage.ParseAstString(commitToApply, true)
	stage.isApplyingBackwardCommit = false
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
	stage.isApplyingForwardCommit = true
	err := stage.ParseAstString(commitToApply, true)
	stage.isApplyingForwardCommit = false
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

func (stage *Stage) Lock() {
	stage.lock.Lock()
}

func (stage *Stage) Unlock() {
	stage.lock.Unlock()
}

func (stage *Stage) RLock() {
	stage.lock.RLock()
}

func (stage *Stage) RUnlock() {
	stage.lock.RUnlock()
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

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.modified = true
	stage.isSquashing = true

	// insertion point for clear references
	stage.Checkboxs_reference = make(map[*Checkbox]*Checkbox)
	stage.Checkboxs_instance = make(map[*Checkbox]*Checkbox)
	stage.Checkboxs_referenceOrder = make(map[*Checkbox]uint)

	stage.Groups_reference = make(map[*Group]*Group)
	stage.Groups_instance = make(map[*Group]*Group)
	stage.Groups_referenceOrder = make(map[*Group]uint)

	stage.Layouts_reference = make(map[*Layout]*Layout)
	stage.Layouts_instance = make(map[*Layout]*Layout)
	stage.Layouts_referenceOrder = make(map[*Layout]uint)

	stage.Sliders_reference = make(map[*Slider]*Slider)
	stage.Sliders_instance = make(map[*Slider]*Slider)
	stage.Sliders_referenceOrder = make(map[*Slider]uint)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
	}

	stage.isSquashing = false
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	var maxCheckboxOrder uint
	var foundCheckbox bool
	for _, order := range stage.Checkbox_stagedOrder {
		if !foundCheckbox || order > maxCheckboxOrder {
			maxCheckboxOrder = order
			foundCheckbox = true
		}
	}
	if foundCheckbox {
		stage.CheckboxOrder = maxCheckboxOrder + 1
	} else {
		stage.CheckboxOrder = 0
	}

	var maxGroupOrder uint
	var foundGroup bool
	for _, order := range stage.Group_stagedOrder {
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

	var maxLayoutOrder uint
	var foundLayout bool
	for _, order := range stage.Layout_stagedOrder {
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

	var maxSliderOrder uint
	var foundSlider bool
	for _, order := range stage.Slider_stagedOrder {
		if !foundSlider || order > maxSliderOrder {
			maxSliderOrder = order
			foundSlider = true
		}
	}
	if foundSlider {
		stage.SliderOrder = maxSliderOrder + 1
	} else {
		stage.SliderOrder = 0
	}

	// end of insertion point for max order recomputation
}

func (stage *Stage) SetDeltaMode(inDeltaMode bool) {
	stage.isInDeltaMode = inDeltaMode
}

func (stage *Stage) IsInDeltaMode() bool {
	return stage.isInDeltaMode
}

func (stage *Stage) SetProbeIF(probeIF GongProbeIF) {
	stage.probeIF = probeIF
}

func (stage *Stage) GetProbeIF() GongProbeIF {
	if stage.probeIF == nil {
		return nil
	}

	return stage.probeIF
}

// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Checkbox:
		tmp := __gong__getStructInstancesByOrder(stage.Checkboxs, stage.Checkbox_stagedOrder)

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
		tmp := __gong__getStructInstancesByOrder(stage.Groups, stage.Group_stagedOrder)

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
		tmp := __gong__getStructInstancesByOrder(stage.Layouts, stage.Layout_stagedOrder)

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
		tmp := __gong__getStructInstancesByOrder(stage.Sliders, stage.Slider_stagedOrder)

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

func __gong__getStructInstancesByOrder[T GongstructPtr](set map[T]struct{}, order map[T]uint) (res []T) {
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
			log.Fatalf("getStructInstancesByOrder: pointer not found")
		}
		return i_order < j_order
	})

	res = append(res, orderedSet...)

	return
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/lib/slider/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type GongOnInitCommitInterface interface {
	BeforeCommit(stage *Stage)
}

type OnInitCommitInterface = GongOnInitCommitInterface

// GongOnAfterCreateInterface callback when an instance is updated from the front
type GongOnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *Stage,
		instance *Type)
}

type OnAfterCreateInterface[Type Gongstruct] = GongOnAfterCreateInterface[Type]

// GongOnAfterReadInterface callback when an instance is updated from the front
type GongOnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *Stage,
		instance *Type)
}

type OnAfterReadInterface[Type Gongstruct] = GongOnAfterReadInterface[Type]

// GongOnAfterUpdateInterface callback when an instance is updated from the front
type GongOnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *Stage, old, new *Type)
}

type OnAfterUpdateInterface[Type Gongstruct] = GongOnAfterUpdateInterface[Type]

// GongOnAfterDeleteInterface callback when an instance is updated from the front
type GongOnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *Stage,
		staged, front *Type)
}

type OnAfterDeleteInterface[Type Gongstruct] = GongOnAfterDeleteInterface[Type]

type GongBackRepoInterface interface {
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

type BackRepoInterface = GongBackRepoInterface

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
		Checkbox_stagedOrder: make(map[*Checkbox]uint),
		Checkbox_orderStaged: make(map[uint]*Checkbox),
		Checkboxs_reference:  make(map[*Checkbox]*Checkbox),

		Group_stagedOrder: make(map[*Group]uint),
		Group_orderStaged: make(map[uint]*Group),
		Groups_reference:  make(map[*Group]*Group),

		Layout_stagedOrder: make(map[*Layout]uint),
		Layout_orderStaged: make(map[uint]*Layout),
		Layouts_reference:  make(map[*Layout]*Layout),

		Slider_stagedOrder: make(map[*Slider]uint),
		Slider_orderStaged: make(map[uint]*Slider),
		Sliders_reference:  make(map[*Slider]*Slider),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Checkbox": &CheckboxUnmarshaller{},

			"Group": &GroupUnmarshaller{},

			"Layout": &LayoutUnmarshaller{},

			"Slider": &SliderUnmarshaller{},

			// end of insertion point
		},

		navigationMode: GongNavigationModeNormal,
	}

	return
}

// GetOrder is the Stage method returning the order of a gongstruct instance.
func (stage *Stage) GetOrder(instance GongstructIF) uint {
	if instance != nil {
		return instance.GongGetOrder(stage)
	}
	return 0
}

// GetInstanceFromOrder is the Stage method returning a gongstruct instance from its order.
func (stage *Stage) GetInstanceFromOrder[Type GongstructPtr](order uint) (res Type) {
	var t Type
	switch any(t).(type) {
	// insertion point for order map initialisations
	case *Checkbox:
		return any(stage.Checkbox_orderStaged[order]).(Type)
	case *Group:
		return any(stage.Group_orderStaged[order]).(Type)
	case *Layout:
		return any(stage.Layout_orderStaged[order]).(Type)
	case *Slider:
		return any(stage.Slider_orderStaged[order]).(Type)
	default:
		return // should not happen
	}
}

func (stage *Stage) GetName() string {
	return stage.name
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.beforeCommitHooks
	stage.beforeCommitHooks = nil
	tmp3 := stage.afterCommitHooks
	stage.afterCommitHooks = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.beforeCommitHooks = tmp2
	stage.afterCommitHooks = tmp3
}

func (stage *Stage) Commit() {
	stage.ComputeReverseMaps()

	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	for _, hook := range stage.beforeCommitHooks {
		hook(stage)
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
		if stage.probeIF != nil {
			stage.probeIF.RefreshNavigationTree()
		}
	}

	// 2. Run all After Commit hooks
	for _, hook := range stage.afterCommitHooks {
		hook(stage)
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
		stage.Checkbox_stagedOrder[checkbox] = stage.CheckboxOrder
		stage.Checkbox_orderStaged[stage.CheckboxOrder] = checkbox
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
		stage.Checkbox_stagedOrder[checkbox] = order
		stage.Checkbox_orderStaged[order] = checkbox
		stage.CheckboxOrder++
	}
	stage.Checkboxs_mapString[checkbox.Name] = checkbox
}

// Unstage removes checkbox off the model stage
func (checkbox *Checkbox) Unstage(stage *Stage) *Checkbox {
	delete(stage.Checkboxs, checkbox)
	// issue1150
	// delete(stage.Checkbox_stagedOrder, checkbox)
	delete(stage.Checkboxs_mapString, checkbox.Name)

	return checkbox
}

// UnstageVoid removes checkbox off the model stage
func (checkbox *Checkbox) UnstageVoid(stage *Stage) {
	delete(stage.Checkboxs, checkbox)
	// issue1150
	// delete(stage.Checkbox_stagedOrder, checkbox)
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
		stage.Group_stagedOrder[group] = stage.GroupOrder
		stage.Group_orderStaged[stage.GroupOrder] = group
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
		stage.Group_stagedOrder[group] = order
		stage.Group_orderStaged[order] = group
		stage.GroupOrder++
	}
	stage.Groups_mapString[group.Name] = group
}

// Unstage removes group off the model stage
func (group *Group) Unstage(stage *Stage) *Group {
	delete(stage.Groups, group)
	// issue1150
	// delete(stage.Group_stagedOrder, group)
	delete(stage.Groups_mapString, group.Name)

	return group
}

// UnstageVoid removes group off the model stage
func (group *Group) UnstageVoid(stage *Stage) {
	delete(stage.Groups, group)
	// issue1150
	// delete(stage.Group_stagedOrder, group)
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
		stage.Layout_stagedOrder[layout] = stage.LayoutOrder
		stage.Layout_orderStaged[stage.LayoutOrder] = layout
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
		stage.Layout_stagedOrder[layout] = order
		stage.Layout_orderStaged[order] = layout
		stage.LayoutOrder++
	}
	stage.Layouts_mapString[layout.Name] = layout
}

// Unstage removes layout off the model stage
func (layout *Layout) Unstage(stage *Stage) *Layout {
	delete(stage.Layouts, layout)
	// issue1150
	// delete(stage.Layout_stagedOrder, layout)
	delete(stage.Layouts_mapString, layout.Name)

	return layout
}

// UnstageVoid removes layout off the model stage
func (layout *Layout) UnstageVoid(stage *Stage) {
	delete(stage.Layouts, layout)
	// issue1150
	// delete(stage.Layout_stagedOrder, layout)
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
		stage.Slider_stagedOrder[slider] = stage.SliderOrder
		stage.Slider_orderStaged[stage.SliderOrder] = slider
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
		stage.Slider_stagedOrder[slider] = order
		stage.Slider_orderStaged[order] = slider
		stage.SliderOrder++
	}
	stage.Sliders_mapString[slider.Name] = slider
}

// Unstage removes slider off the model stage
func (slider *Slider) Unstage(stage *Stage) *Slider {
	delete(stage.Sliders, slider)
	// issue1150
	// delete(stage.Slider_stagedOrder, slider)
	delete(stage.Sliders_mapString, slider.Name)

	return slider
}

// UnstageVoid removes slider off the model stage
func (slider *Slider) UnstageVoid(stage *Stage) {
	delete(stage.Sliders, slider)
	// issue1150
	// delete(stage.Slider_stagedOrder, slider)
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

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Checkboxs = make(map[*Checkbox]struct{})
	stage.Checkboxs_mapString = make(map[string]*Checkbox)
	stage.Checkbox_stagedOrder = make(map[*Checkbox]uint)
	stage.CheckboxOrder = 0

	stage.Groups = make(map[*Group]struct{})
	stage.Groups_mapString = make(map[string]*Group)
	stage.Group_stagedOrder = make(map[*Group]uint)
	stage.GroupOrder = 0

	stage.Layouts = make(map[*Layout]struct{})
	stage.Layouts_mapString = make(map[string]*Layout)
	stage.Layout_stagedOrder = make(map[*Layout]uint)
	stage.LayoutOrder = 0

	stage.Sliders = make(map[*Slider]struct{})
	stage.Sliders_mapString = make(map[string]*Slider)
	stage.Slider_stagedOrder = make(map[*Slider]uint)
	stage.SliderOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct any

type GongstructBasicField interface {
	int | float64 | bool | string | time.Time | time.Duration
}

type GongtructBasicField = GongstructBasicField

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type GongstructIF interface {
	GetName() string
	SetName(string)
	StageVoid(*Stage)
	UnstageVoid(stage *Stage)
	GongGetFieldHeaders() []GongFieldHeader
	GongGetFieldValue(fieldName string, stage *Stage) GongFieldValue
	GongGetGongstructName() string
	GongGetOrder(stage *Stage) uint
	GongGetReferenceIdentifier(stage *Stage) string
	GongGetIdentifier(stage *Stage) string
	GongCopy() GongstructIF
	GongGetReverseFieldOwnerName(stage *Stage, reverseField *GongReverseField) string
	GongGetReverseFieldOwner(stage *Stage, reverseField *GongReverseField) GongstructIF
	GongGetUUID(stage *Stage) string
	GongAfterCreateFromFront(stage *Stage)
	GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF)
	GongAfterDeleteFromFront(stage *Stage, front GongstructIF)
	GongIsStaged(stage *Stage) bool
	GongStageBranch(stage *Stage)
	GongUnstageBranch(stage *Stage)
}
type GongstructPtr interface {
	GongstructIF
	comparable
}

type PointerToGongstruct = GongstructPtr

func GongCompareGongstructByName[T GongstructPtr](a, b T) int {
	return cmp.Compare(a.GetName(), b.GetName())
}

func GongSortGongstructSetByName[T GongstructPtr](set map[T]struct{}) (sortedSlice []T) {
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, GongCompareGongstructByName)

	return
}

// GetInstancesSorted is the Stage method returning sorted instances of a gongstruct.
func (stage *Stage) GetInstancesSorted[T GongstructPtr]() (sortedSlice []T) {
	set := stage.GetInstancesSet[T]()
	sortedSlice = GongSortGongstructSetByName(*set)

	return
}

// GetInstancesMapByName is the Stage method returning a map of staged instances by their name.
func (stage *Stage) GetInstancesMapByName[Type GongstructIF]() map[string]Type {
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

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
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

// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for instance with special fields
	case Group:
		return any(&Group{
			Sliders: []*Slider{{Name: "Sliders"}},
			Checkboxes: []*Checkbox{{Name: "Checkboxes"}},
		}).(*Type)
	case Layout:
		return any(&Layout{
			Groups: []*Group{{Name: "Groups"}},
		}).(*Type)
	default:
		return &ret
	}
}

// GetPointerReverseMap allows backtrack navigation of any Start.Fieldname
// associations (0..1) that is a pointer from one staged Gongstruct (type Start)
// instances to another (type End)
//
// The function provides a map with keys as instances of End and values to arrays of *Start
// the map is construed by iterating over all Start instances and populationg keys with End instances
// and values with slice of Start instances
// GetPointerReverseMap is the Stage method for backtrack navigation of pointer associations.
func (stage *Stage) GetPointerReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
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

// GetSliceOfPointersReverseMap is the Stage method for backtrack navigation of slice-of-pointers associations.
func (stage *Stage) GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string) map[*End][]*Start {
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

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *Checkbox:
		res = any(new(Checkbox)).(Type)
	case *Group:
		res = any(new(Group)).(Type)
	case *Layout:
		res = any(new(Layout)).(Type)
	case *Slider:
		res = any(new(Slider)).(Type)
	}
	return res
}

func NewInstance[Type GongstructPtr]() (res Type) {
	return GongNewInstance[Type]()
}

func (stage *Stage) GongNewInstance[Type GongstructPtr]() (res Type) {
	res = GongNewInstance[Type]()
	var zero Type
	if res != zero {
		res.StageVoid(stage)
	}
	return res
}

func (stage *Stage) NewInstance[Type GongstructPtr]() (res Type) {
	return stage.GongNewInstance[Type]()
}

// GongGetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GongGetPointerToGongstructName[Type GongstructIF]() (res string) {
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

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	res = make([]GongReverseField, 0)

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

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (checkbox *Checkbox) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "ValueBool",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "LabelForTrue",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "LabelForFalse",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (group *Group) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Percentage",
			GongFieldValueType: GongFieldValueTypeFloat,
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
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Groups",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Group",
		},
		{
			Name:               "IsWithCustomGutterSize",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "GutterSize",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (slider *Slider) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "IsFloat64",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "IsInt",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "MinInt",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MaxInt",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "StepInt",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "ValueInt",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MinFloat64",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "MaxFloat64",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "StepFloat64",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "ValueFloat64",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
		{
			Name:               "IsDisabled",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

// GongGetFieldsFromPointer return the array of the fields
func GongGetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
}

func GetFieldsFromPointer[Type GongstructPtr]() (res []GongFieldHeader) {
	return GongGetFieldsFromPointer[Type]()
}

type GongFieldValueType string

const (
	GongFieldValueTypeInt             GongFieldValueType = "GongFieldValueTypeInt"
	GongFieldValueTypeIntDuration     GongFieldValueType = "GongFieldValueTypeIntDuration"
	GongFieldValueTypeFloat           GongFieldValueType = "GongFieldValueTypeFloat"
	GongFieldValueTypeBool            GongFieldValueType = "GongFieldValueTypeBool"
	GongFieldValueTypeString          GongFieldValueType = "GongFieldValueTypeString"
	GongFieldValueTypeDate            GongFieldValueType = "GongFieldValueTypeDate"
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
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Checkboxes":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range group.Checkboxes {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
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
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsWithCustomGutterSize":
		res.valueString = fmt.Sprintf("%t", layout.IsWithCustomGutterSize)
		res.valueBool = layout.IsWithCustomGutterSize
		res.GongFieldValueType = GongFieldValueTypeBool
	case "GutterSize":
		res.valueString = fmt.Sprintf("%f", layout.GutterSize)
		res.valueFloat = layout.GutterSize
		res.GongFieldValueType = GongFieldValueTypeFloat
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
	case "IsDisabled":
		res.valueString = fmt.Sprintf("%t", slider.IsDisabled)
		res.valueBool = slider.IsDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
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

func GongGetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	return GongGetGongstructNameFromPointer(instance)
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

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
