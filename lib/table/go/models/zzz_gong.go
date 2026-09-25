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
	Buttons                map[*Button]struct{}
	Buttons_instance       map[*Button]*Button
	Buttons_mapString      map[string]*Button
	ButtonOrder            uint
	Button_stagedOrder     map[*Button]uint
	Button_orderStaged     map[uint]*Button
	Buttons_reference      map[*Button]*Button
	Buttons_referenceOrder map[*Button]uint

	// insertion point for slice of pointers maps
	OnAfterButtonCreateCallback GongOnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback GongOnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback GongOnAfterDeleteInterface[Button]

	Cells                map[*Cell]struct{}
	Cells_instance       map[*Cell]*Cell
	Cells_mapString      map[string]*Cell
	CellOrder            uint
	Cell_stagedOrder     map[*Cell]uint
	Cell_orderStaged     map[uint]*Cell
	Cells_reference      map[*Cell]*Cell
	Cells_referenceOrder map[*Cell]uint

	// insertion point for slice of pointers maps
	OnAfterCellCreateCallback GongOnAfterCreateInterface[Cell]
	OnAfterCellUpdateCallback GongOnAfterUpdateInterface[Cell]
	OnAfterCellDeleteCallback GongOnAfterDeleteInterface[Cell]

	CellBooleans                map[*CellBoolean]struct{}
	CellBooleans_instance       map[*CellBoolean]*CellBoolean
	CellBooleans_mapString      map[string]*CellBoolean
	CellBooleanOrder            uint
	CellBoolean_stagedOrder     map[*CellBoolean]uint
	CellBoolean_orderStaged     map[uint]*CellBoolean
	CellBooleans_reference      map[*CellBoolean]*CellBoolean
	CellBooleans_referenceOrder map[*CellBoolean]uint

	// insertion point for slice of pointers maps
	OnAfterCellBooleanCreateCallback GongOnAfterCreateInterface[CellBoolean]
	OnAfterCellBooleanUpdateCallback GongOnAfterUpdateInterface[CellBoolean]
	OnAfterCellBooleanDeleteCallback GongOnAfterDeleteInterface[CellBoolean]

	CellFloat64s                map[*CellFloat64]struct{}
	CellFloat64s_instance       map[*CellFloat64]*CellFloat64
	CellFloat64s_mapString      map[string]*CellFloat64
	CellFloat64Order            uint
	CellFloat64_stagedOrder     map[*CellFloat64]uint
	CellFloat64_orderStaged     map[uint]*CellFloat64
	CellFloat64s_reference      map[*CellFloat64]*CellFloat64
	CellFloat64s_referenceOrder map[*CellFloat64]uint

	// insertion point for slice of pointers maps
	OnAfterCellFloat64CreateCallback GongOnAfterCreateInterface[CellFloat64]
	OnAfterCellFloat64UpdateCallback GongOnAfterUpdateInterface[CellFloat64]
	OnAfterCellFloat64DeleteCallback GongOnAfterDeleteInterface[CellFloat64]

	CellIcons                map[*CellIcon]struct{}
	CellIcons_instance       map[*CellIcon]*CellIcon
	CellIcons_mapString      map[string]*CellIcon
	CellIconOrder            uint
	CellIcon_stagedOrder     map[*CellIcon]uint
	CellIcon_orderStaged     map[uint]*CellIcon
	CellIcons_reference      map[*CellIcon]*CellIcon
	CellIcons_referenceOrder map[*CellIcon]uint

	// insertion point for slice of pointers maps
	OnAfterCellIconCreateCallback GongOnAfterCreateInterface[CellIcon]
	OnAfterCellIconUpdateCallback GongOnAfterUpdateInterface[CellIcon]
	OnAfterCellIconDeleteCallback GongOnAfterDeleteInterface[CellIcon]

	CellInts                map[*CellInt]struct{}
	CellInts_instance       map[*CellInt]*CellInt
	CellInts_mapString      map[string]*CellInt
	CellIntOrder            uint
	CellInt_stagedOrder     map[*CellInt]uint
	CellInt_orderStaged     map[uint]*CellInt
	CellInts_reference      map[*CellInt]*CellInt
	CellInts_referenceOrder map[*CellInt]uint

	// insertion point for slice of pointers maps
	OnAfterCellIntCreateCallback GongOnAfterCreateInterface[CellInt]
	OnAfterCellIntUpdateCallback GongOnAfterUpdateInterface[CellInt]
	OnAfterCellIntDeleteCallback GongOnAfterDeleteInterface[CellInt]

	CellStrings                map[*CellString]struct{}
	CellStrings_instance       map[*CellString]*CellString
	CellStrings_mapString      map[string]*CellString
	CellStringOrder            uint
	CellString_stagedOrder     map[*CellString]uint
	CellString_orderStaged     map[uint]*CellString
	CellStrings_reference      map[*CellString]*CellString
	CellStrings_referenceOrder map[*CellString]uint

	// insertion point for slice of pointers maps
	OnAfterCellStringCreateCallback GongOnAfterCreateInterface[CellString]
	OnAfterCellStringUpdateCallback GongOnAfterUpdateInterface[CellString]
	OnAfterCellStringDeleteCallback GongOnAfterDeleteInterface[CellString]

	DisplayedColumns                map[*DisplayedColumn]struct{}
	DisplayedColumns_instance       map[*DisplayedColumn]*DisplayedColumn
	DisplayedColumns_mapString      map[string]*DisplayedColumn
	DisplayedColumnOrder            uint
	DisplayedColumn_stagedOrder     map[*DisplayedColumn]uint
	DisplayedColumn_orderStaged     map[uint]*DisplayedColumn
	DisplayedColumns_reference      map[*DisplayedColumn]*DisplayedColumn
	DisplayedColumns_referenceOrder map[*DisplayedColumn]uint

	// insertion point for slice of pointers maps
	OnAfterDisplayedColumnCreateCallback GongOnAfterCreateInterface[DisplayedColumn]
	OnAfterDisplayedColumnUpdateCallback GongOnAfterUpdateInterface[DisplayedColumn]
	OnAfterDisplayedColumnDeleteCallback GongOnAfterDeleteInterface[DisplayedColumn]

	Rows                map[*Row]struct{}
	Rows_instance       map[*Row]*Row
	Rows_mapString      map[string]*Row
	RowOrder            uint
	Row_stagedOrder     map[*Row]uint
	Row_orderStaged     map[uint]*Row
	Rows_reference      map[*Row]*Row
	Rows_referenceOrder map[*Row]uint

	// insertion point for slice of pointers maps
	Row_Cells_reverseMap map[*Cell]*Row

	OnAfterRowCreateCallback GongOnAfterCreateInterface[Row]
	OnAfterRowUpdateCallback GongOnAfterUpdateInterface[Row]
	OnAfterRowDeleteCallback GongOnAfterDeleteInterface[Row]

	SVGIcons                map[*SVGIcon]struct{}
	SVGIcons_instance       map[*SVGIcon]*SVGIcon
	SVGIcons_mapString      map[string]*SVGIcon
	SVGIconOrder            uint
	SVGIcon_stagedOrder     map[*SVGIcon]uint
	SVGIcon_orderStaged     map[uint]*SVGIcon
	SVGIcons_reference      map[*SVGIcon]*SVGIcon
	SVGIcons_referenceOrder map[*SVGIcon]uint

	// insertion point for slice of pointers maps
	OnAfterSVGIconCreateCallback GongOnAfterCreateInterface[SVGIcon]
	OnAfterSVGIconUpdateCallback GongOnAfterUpdateInterface[SVGIcon]
	OnAfterSVGIconDeleteCallback GongOnAfterDeleteInterface[SVGIcon]

	Tables                map[*Table]struct{}
	Tables_instance       map[*Table]*Table
	Tables_mapString      map[string]*Table
	TableOrder            uint
	Table_stagedOrder     map[*Table]uint
	Table_orderStaged     map[uint]*Table
	Tables_reference      map[*Table]*Table
	Tables_referenceOrder map[*Table]uint

	// insertion point for slice of pointers maps
	Table_DisplayedColumns_reverseMap map[*DisplayedColumn]*Table

	Table_Rows_reverseMap map[*Row]*Table

	Table_RowsSelectedForBulkDelete_reverseMap map[*Row]*Table

	Table_Buttons_reverseMap map[*Button]*Table

	OnAfterTableCreateCallback GongOnAfterCreateInterface[Table]
	OnAfterTableUpdateCallback GongOnAfterUpdateInterface[Table]
	OnAfterTableDeleteCallback GongOnAfterDeleteInterface[Table]

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
	__gong__clearReferences(&stage.Buttons_reference, &stage.Buttons_instance, &stage.Buttons_referenceOrder)

	__gong__clearReferences(&stage.Cells_reference, &stage.Cells_instance, &stage.Cells_referenceOrder)

	__gong__clearReferences(&stage.CellBooleans_reference, &stage.CellBooleans_instance, &stage.CellBooleans_referenceOrder)

	__gong__clearReferences(&stage.CellFloat64s_reference, &stage.CellFloat64s_instance, &stage.CellFloat64s_referenceOrder)

	__gong__clearReferences(&stage.CellIcons_reference, &stage.CellIcons_instance, &stage.CellIcons_referenceOrder)

	__gong__clearReferences(&stage.CellInts_reference, &stage.CellInts_instance, &stage.CellInts_referenceOrder)

	__gong__clearReferences(&stage.CellStrings_reference, &stage.CellStrings_instance, &stage.CellStrings_referenceOrder)

	__gong__clearReferences(&stage.DisplayedColumns_reference, &stage.DisplayedColumns_instance, &stage.DisplayedColumns_referenceOrder)

	__gong__clearReferences(&stage.Rows_reference, &stage.Rows_instance, &stage.Rows_referenceOrder)

	__gong__clearReferences(&stage.SVGIcons_reference, &stage.SVGIcons_instance, &stage.SVGIcons_referenceOrder)

	__gong__clearReferences(&stage.Tables_reference, &stage.Tables_instance, &stage.Tables_referenceOrder)

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
	stage.ButtonOrder = __gong__recomputeOrder(stage.Button_stagedOrder)

	stage.CellOrder = __gong__recomputeOrder(stage.Cell_stagedOrder)

	stage.CellBooleanOrder = __gong__recomputeOrder(stage.CellBoolean_stagedOrder)

	stage.CellFloat64Order = __gong__recomputeOrder(stage.CellFloat64_stagedOrder)

	stage.CellIconOrder = __gong__recomputeOrder(stage.CellIcon_stagedOrder)

	stage.CellIntOrder = __gong__recomputeOrder(stage.CellInt_stagedOrder)

	stage.CellStringOrder = __gong__recomputeOrder(stage.CellString_stagedOrder)

	stage.DisplayedColumnOrder = __gong__recomputeOrder(stage.DisplayedColumn_stagedOrder)

	stage.RowOrder = __gong__recomputeOrder(stage.Row_stagedOrder)

	stage.SVGIconOrder = __gong__recomputeOrder(stage.SVGIcon_stagedOrder)

	stage.TableOrder = __gong__recomputeOrder(stage.Table_stagedOrder)

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
	case *Button:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Buttons, stage.Button_stagedOrder))
	case *Cell:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Cells, stage.Cell_stagedOrder))
	case *CellBoolean:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.CellBooleans, stage.CellBoolean_stagedOrder))
	case *CellFloat64:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.CellFloat64s, stage.CellFloat64_stagedOrder))
	case *CellIcon:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.CellIcons, stage.CellIcon_stagedOrder))
	case *CellInt:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.CellInts, stage.CellInt_stagedOrder))
	case *CellString:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.CellStrings, stage.CellString_stagedOrder))
	case *DisplayedColumn:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.DisplayedColumns, stage.DisplayedColumn_stagedOrder))
	case *Row:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Rows, stage.Row_stagedOrder))
	case *SVGIcon:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.SVGIcons, stage.SVGIcon_stagedOrder))
	case *Table:
		return __gong__castSlice[T](__gong__getStructInstancesByOrder(stage.Tables, stage.Table_stagedOrder))

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

func __gong__castSlice[T any, S any](s []S) []T {
	res := make([]T, len(s))
	for i, v := range s {
		res[i] = any(v).(T)
	}
	return res
}

func __gong__stage[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	order *uint,
	mapString map[string]T,
	instance T,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		stagedOrder[instance] = *order
		orderStaged[*order] = instance
		*order++
	}
	mapString[name] = instance
}

func __gong__stagePreserveOrder[T comparable](
	instances map[T]struct{},
	stagedOrder map[T]uint,
	orderStaged map[uint]T,
	currentOrder *uint,
	mapString map[string]T,
	instance T,
	order uint,
	name string,
) {
	if _, ok := instances[instance]; !ok {
		instances[instance] = struct{}{}
		if order > *currentOrder {
			*currentOrder = order
		}
		stagedOrder[instance] = order
		orderStaged[order] = instance
		*currentOrder++
	}
	mapString[name] = instance
}

func __gong__unstage[T comparable](
	instances map[T]struct{},
	mapString map[string]T,
	instance T,
	name string,
) {
	delete(instances, instance)
	delete(mapString, name)
}

func __gong__recomputeOrder[T comparable](stagedOrder map[T]uint) uint {
	var maxOrder uint
	var found bool
	for _, order := range stagedOrder {
		if !found || order > maxOrder {
			maxOrder = order
			found = true
		}
	}
	if found {
		return maxOrder + 1
	}
	return 0
}

func __gong__rebuildMapString[T interface {
	comparable
	GetName() string
}](staged map[T]struct{}, mapString *map[string]T) {
	*mapString = make(map[string]T, len(staged))
	for instance := range staged {
		(*mapString)[instance.GetName()] = instance
	}
}

func __gong__clearReferences[T comparable](ref *map[T]T, inst *map[T]T, refOrder *map[T]uint) {
	*ref = make(map[T]T)
	*inst = make(map[T]T)
	*refOrder = make(map[T]uint)
}

func __gong__resetStageType[T comparable](staged *map[T]struct{}, mapString *map[string]T, stagedOrder *map[T]uint, order *uint) {
	*staged = make(map[T]struct{})
	*mapString = make(map[string]T)
	*stagedOrder = make(map[T]uint)
	*order = 0
}

func (stage *Stage) GetType() string {
	return "github.com/fullstack-lang/gong/lib/table/go/models"
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
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

type BackRepoInterface = GongBackRepoInterface

func NewStage(name string) (stage *Stage) {
	stage = &Stage{ // insertion point for array initiatialisation
		Buttons:           make(map[*Button]struct{}),
		Buttons_mapString: make(map[string]*Button),

		Cells:           make(map[*Cell]struct{}),
		Cells_mapString: make(map[string]*Cell),

		CellBooleans:           make(map[*CellBoolean]struct{}),
		CellBooleans_mapString: make(map[string]*CellBoolean),

		CellFloat64s:           make(map[*CellFloat64]struct{}),
		CellFloat64s_mapString: make(map[string]*CellFloat64),

		CellIcons:           make(map[*CellIcon]struct{}),
		CellIcons_mapString: make(map[string]*CellIcon),

		CellInts:           make(map[*CellInt]struct{}),
		CellInts_mapString: make(map[string]*CellInt),

		CellStrings:           make(map[*CellString]struct{}),
		CellStrings_mapString: make(map[string]*CellString),

		DisplayedColumns:           make(map[*DisplayedColumn]struct{}),
		DisplayedColumns_mapString: make(map[string]*DisplayedColumn),

		Rows:           make(map[*Row]struct{}),
		Rows_mapString: make(map[string]*Row),

		SVGIcons:           make(map[*SVGIcon]struct{}),
		SVGIcons_mapString: make(map[string]*SVGIcon),

		Tables:           make(map[*Table]struct{}),
		Tables_mapString: make(map[string]*Table),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		Button_stagedOrder: make(map[*Button]uint),
		Button_orderStaged: make(map[uint]*Button),
		Buttons_reference:  make(map[*Button]*Button),

		Cell_stagedOrder: make(map[*Cell]uint),
		Cell_orderStaged: make(map[uint]*Cell),
		Cells_reference:  make(map[*Cell]*Cell),

		CellBoolean_stagedOrder: make(map[*CellBoolean]uint),
		CellBoolean_orderStaged: make(map[uint]*CellBoolean),
		CellBooleans_reference:  make(map[*CellBoolean]*CellBoolean),

		CellFloat64_stagedOrder: make(map[*CellFloat64]uint),
		CellFloat64_orderStaged: make(map[uint]*CellFloat64),
		CellFloat64s_reference:  make(map[*CellFloat64]*CellFloat64),

		CellIcon_stagedOrder: make(map[*CellIcon]uint),
		CellIcon_orderStaged: make(map[uint]*CellIcon),
		CellIcons_reference:  make(map[*CellIcon]*CellIcon),

		CellInt_stagedOrder: make(map[*CellInt]uint),
		CellInt_orderStaged: make(map[uint]*CellInt),
		CellInts_reference:  make(map[*CellInt]*CellInt),

		CellString_stagedOrder: make(map[*CellString]uint),
		CellString_orderStaged: make(map[uint]*CellString),
		CellStrings_reference:  make(map[*CellString]*CellString),

		DisplayedColumn_stagedOrder: make(map[*DisplayedColumn]uint),
		DisplayedColumn_orderStaged: make(map[uint]*DisplayedColumn),
		DisplayedColumns_reference:  make(map[*DisplayedColumn]*DisplayedColumn),

		Row_stagedOrder: make(map[*Row]uint),
		Row_orderStaged: make(map[uint]*Row),
		Rows_reference:  make(map[*Row]*Row),

		SVGIcon_stagedOrder: make(map[*SVGIcon]uint),
		SVGIcon_orderStaged: make(map[uint]*SVGIcon),
		SVGIcons_reference:  make(map[*SVGIcon]*SVGIcon),

		Table_stagedOrder: make(map[*Table]uint),
		Table_orderStaged: make(map[uint]*Table),
		Tables_reference:  make(map[*Table]*Table),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"Button": &ButtonUnmarshaller{},

			"Cell": &CellUnmarshaller{},

			"CellBoolean": &CellBooleanUnmarshaller{},

			"CellFloat64": &CellFloat64Unmarshaller{},

			"CellIcon": &CellIconUnmarshaller{},

			"CellInt": &CellIntUnmarshaller{},

			"CellString": &CellStringUnmarshaller{},

			"DisplayedColumn": &DisplayedColumnUnmarshaller{},

			"Row": &RowUnmarshaller{},

			"SVGIcon": &SVGIconUnmarshaller{},

			"Table": &TableUnmarshaller{},

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
	case *Button:
		return any(stage.Button_orderStaged[order]).(Type)
	case *Cell:
		return any(stage.Cell_orderStaged[order]).(Type)
	case *CellBoolean:
		return any(stage.CellBoolean_orderStaged[order]).(Type)
	case *CellFloat64:
		return any(stage.CellFloat64_orderStaged[order]).(Type)
	case *CellIcon:
		return any(stage.CellIcon_orderStaged[order]).(Type)
	case *CellInt:
		return any(stage.CellInt_orderStaged[order]).(Type)
	case *CellString:
		return any(stage.CellString_orderStaged[order]).(Type)
	case *DisplayedColumn:
		return any(stage.DisplayedColumn_orderStaged[order]).(Type)
	case *Row:
		return any(stage.Row_orderStaged[order]).(Type)
	case *SVGIcon:
		return any(stage.SVGIcon_orderStaged[order]).(Type)
	case *Table:
		return any(stage.Table_orderStaged[order]).(Type)
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
	stage.Map_GongStructName_InstancesNb["Button"] = len(stage.Buttons)
	stage.Map_GongStructName_InstancesNb["Cell"] = len(stage.Cells)
	stage.Map_GongStructName_InstancesNb["CellBoolean"] = len(stage.CellBooleans)
	stage.Map_GongStructName_InstancesNb["CellFloat64"] = len(stage.CellFloat64s)
	stage.Map_GongStructName_InstancesNb["CellIcon"] = len(stage.CellIcons)
	stage.Map_GongStructName_InstancesNb["CellInt"] = len(stage.CellInts)
	stage.Map_GongStructName_InstancesNb["CellString"] = len(stage.CellStrings)
	stage.Map_GongStructName_InstancesNb["DisplayedColumn"] = len(stage.DisplayedColumns)
	stage.Map_GongStructName_InstancesNb["Row"] = len(stage.Rows)
	stage.Map_GongStructName_InstancesNb["SVGIcon"] = len(stage.SVGIcons)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)
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
	__gong__stage(stage.Buttons, stage.Button_stagedOrder, stage.Button_orderStaged, &stage.ButtonOrder, stage.Buttons_mapString, button, button.Name)
	return button
}

// StagePreserveOrder puts button to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.ButtonOrder
// - update stage.ButtonOrder accordingly
func (button *Button) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Buttons, stage.Button_stagedOrder, stage.Button_orderStaged, &stage.ButtonOrder, stage.Buttons_mapString, button, order, button.Name)
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	__gong__unstage(stage.Buttons, stage.Buttons_mapString, button, button.Name)
	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *Stage) {
	button.Unstage(stage)
}

func (button *Button) StageVoid(stage *Stage) {
	button.Stage(stage)
}

// for satisfaction of GongStruct interface
func (button *Button) GetName() (res string) {
	return button.Name
}

// for satisfaction of GongStruct interface
func (button *Button) SetName(name string) {
	button.Name = name
}

// Stage puts cell to the model stage
func (cell *Cell) Stage(stage *Stage) *Cell {
	__gong__stage(stage.Cells, stage.Cell_stagedOrder, stage.Cell_orderStaged, &stage.CellOrder, stage.Cells_mapString, cell, cell.Name)
	return cell
}

// StagePreserveOrder puts cell to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellOrder
// - update stage.CellOrder accordingly
func (cell *Cell) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Cells, stage.Cell_stagedOrder, stage.Cell_orderStaged, &stage.CellOrder, stage.Cells_mapString, cell, order, cell.Name)
}

// Unstage removes cell off the model stage
func (cell *Cell) Unstage(stage *Stage) *Cell {
	__gong__unstage(stage.Cells, stage.Cells_mapString, cell, cell.Name)
	return cell
}

// UnstageVoid removes cell off the model stage
func (cell *Cell) UnstageVoid(stage *Stage) {
	cell.Unstage(stage)
}

func (cell *Cell) StageVoid(stage *Stage) {
	cell.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cell *Cell) GetName() (res string) {
	return cell.Name
}

// for satisfaction of GongStruct interface
func (cell *Cell) SetName(name string) {
	cell.Name = name
}

// Stage puts cellboolean to the model stage
func (cellboolean *CellBoolean) Stage(stage *Stage) *CellBoolean {
	__gong__stage(stage.CellBooleans, stage.CellBoolean_stagedOrder, stage.CellBoolean_orderStaged, &stage.CellBooleanOrder, stage.CellBooleans_mapString, cellboolean, cellboolean.Name)
	return cellboolean
}

// StagePreserveOrder puts cellboolean to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellBooleanOrder
// - update stage.CellBooleanOrder accordingly
func (cellboolean *CellBoolean) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CellBooleans, stage.CellBoolean_stagedOrder, stage.CellBoolean_orderStaged, &stage.CellBooleanOrder, stage.CellBooleans_mapString, cellboolean, order, cellboolean.Name)
}

// Unstage removes cellboolean off the model stage
func (cellboolean *CellBoolean) Unstage(stage *Stage) *CellBoolean {
	__gong__unstage(stage.CellBooleans, stage.CellBooleans_mapString, cellboolean, cellboolean.Name)
	return cellboolean
}

// UnstageVoid removes cellboolean off the model stage
func (cellboolean *CellBoolean) UnstageVoid(stage *Stage) {
	cellboolean.Unstage(stage)
}

func (cellboolean *CellBoolean) StageVoid(stage *Stage) {
	cellboolean.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cellboolean *CellBoolean) GetName() (res string) {
	return cellboolean.Name
}

// for satisfaction of GongStruct interface
func (cellboolean *CellBoolean) SetName(name string) {
	cellboolean.Name = name
}

// Stage puts cellfloat64 to the model stage
func (cellfloat64 *CellFloat64) Stage(stage *Stage) *CellFloat64 {
	__gong__stage(stage.CellFloat64s, stage.CellFloat64_stagedOrder, stage.CellFloat64_orderStaged, &stage.CellFloat64Order, stage.CellFloat64s_mapString, cellfloat64, cellfloat64.Name)
	return cellfloat64
}

// StagePreserveOrder puts cellfloat64 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellFloat64Order
// - update stage.CellFloat64Order accordingly
func (cellfloat64 *CellFloat64) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CellFloat64s, stage.CellFloat64_stagedOrder, stage.CellFloat64_orderStaged, &stage.CellFloat64Order, stage.CellFloat64s_mapString, cellfloat64, order, cellfloat64.Name)
}

// Unstage removes cellfloat64 off the model stage
func (cellfloat64 *CellFloat64) Unstage(stage *Stage) *CellFloat64 {
	__gong__unstage(stage.CellFloat64s, stage.CellFloat64s_mapString, cellfloat64, cellfloat64.Name)
	return cellfloat64
}

// UnstageVoid removes cellfloat64 off the model stage
func (cellfloat64 *CellFloat64) UnstageVoid(stage *Stage) {
	cellfloat64.Unstage(stage)
}

func (cellfloat64 *CellFloat64) StageVoid(stage *Stage) {
	cellfloat64.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cellfloat64 *CellFloat64) GetName() (res string) {
	return cellfloat64.Name
}

// for satisfaction of GongStruct interface
func (cellfloat64 *CellFloat64) SetName(name string) {
	cellfloat64.Name = name
}

// Stage puts cellicon to the model stage
func (cellicon *CellIcon) Stage(stage *Stage) *CellIcon {
	__gong__stage(stage.CellIcons, stage.CellIcon_stagedOrder, stage.CellIcon_orderStaged, &stage.CellIconOrder, stage.CellIcons_mapString, cellicon, cellicon.Name)
	return cellicon
}

// StagePreserveOrder puts cellicon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellIconOrder
// - update stage.CellIconOrder accordingly
func (cellicon *CellIcon) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CellIcons, stage.CellIcon_stagedOrder, stage.CellIcon_orderStaged, &stage.CellIconOrder, stage.CellIcons_mapString, cellicon, order, cellicon.Name)
}

// Unstage removes cellicon off the model stage
func (cellicon *CellIcon) Unstage(stage *Stage) *CellIcon {
	__gong__unstage(stage.CellIcons, stage.CellIcons_mapString, cellicon, cellicon.Name)
	return cellicon
}

// UnstageVoid removes cellicon off the model stage
func (cellicon *CellIcon) UnstageVoid(stage *Stage) {
	cellicon.Unstage(stage)
}

func (cellicon *CellIcon) StageVoid(stage *Stage) {
	cellicon.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cellicon *CellIcon) GetName() (res string) {
	return cellicon.Name
}

// for satisfaction of GongStruct interface
func (cellicon *CellIcon) SetName(name string) {
	cellicon.Name = name
}

// Stage puts cellint to the model stage
func (cellint *CellInt) Stage(stage *Stage) *CellInt {
	__gong__stage(stage.CellInts, stage.CellInt_stagedOrder, stage.CellInt_orderStaged, &stage.CellIntOrder, stage.CellInts_mapString, cellint, cellint.Name)
	return cellint
}

// StagePreserveOrder puts cellint to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellIntOrder
// - update stage.CellIntOrder accordingly
func (cellint *CellInt) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CellInts, stage.CellInt_stagedOrder, stage.CellInt_orderStaged, &stage.CellIntOrder, stage.CellInts_mapString, cellint, order, cellint.Name)
}

// Unstage removes cellint off the model stage
func (cellint *CellInt) Unstage(stage *Stage) *CellInt {
	__gong__unstage(stage.CellInts, stage.CellInts_mapString, cellint, cellint.Name)
	return cellint
}

// UnstageVoid removes cellint off the model stage
func (cellint *CellInt) UnstageVoid(stage *Stage) {
	cellint.Unstage(stage)
}

func (cellint *CellInt) StageVoid(stage *Stage) {
	cellint.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cellint *CellInt) GetName() (res string) {
	return cellint.Name
}

// for satisfaction of GongStruct interface
func (cellint *CellInt) SetName(name string) {
	cellint.Name = name
}

// Stage puts cellstring to the model stage
func (cellstring *CellString) Stage(stage *Stage) *CellString {
	__gong__stage(stage.CellStrings, stage.CellString_stagedOrder, stage.CellString_orderStaged, &stage.CellStringOrder, stage.CellStrings_mapString, cellstring, cellstring.Name)
	return cellstring
}

// StagePreserveOrder puts cellstring to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellStringOrder
// - update stage.CellStringOrder accordingly
func (cellstring *CellString) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.CellStrings, stage.CellString_stagedOrder, stage.CellString_orderStaged, &stage.CellStringOrder, stage.CellStrings_mapString, cellstring, order, cellstring.Name)
}

// Unstage removes cellstring off the model stage
func (cellstring *CellString) Unstage(stage *Stage) *CellString {
	__gong__unstage(stage.CellStrings, stage.CellStrings_mapString, cellstring, cellstring.Name)
	return cellstring
}

// UnstageVoid removes cellstring off the model stage
func (cellstring *CellString) UnstageVoid(stage *Stage) {
	cellstring.Unstage(stage)
}

func (cellstring *CellString) StageVoid(stage *Stage) {
	cellstring.Stage(stage)
}

// for satisfaction of GongStruct interface
func (cellstring *CellString) GetName() (res string) {
	return cellstring.Name
}

// for satisfaction of GongStruct interface
func (cellstring *CellString) SetName(name string) {
	cellstring.Name = name
}

// Stage puts displayedcolumn to the model stage
func (displayedcolumn *DisplayedColumn) Stage(stage *Stage) *DisplayedColumn {
	__gong__stage(stage.DisplayedColumns, stage.DisplayedColumn_stagedOrder, stage.DisplayedColumn_orderStaged, &stage.DisplayedColumnOrder, stage.DisplayedColumns_mapString, displayedcolumn, displayedcolumn.Name)
	return displayedcolumn
}

// StagePreserveOrder puts displayedcolumn to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DisplayedColumnOrder
// - update stage.DisplayedColumnOrder accordingly
func (displayedcolumn *DisplayedColumn) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DisplayedColumns, stage.DisplayedColumn_stagedOrder, stage.DisplayedColumn_orderStaged, &stage.DisplayedColumnOrder, stage.DisplayedColumns_mapString, displayedcolumn, order, displayedcolumn.Name)
}

// Unstage removes displayedcolumn off the model stage
func (displayedcolumn *DisplayedColumn) Unstage(stage *Stage) *DisplayedColumn {
	__gong__unstage(stage.DisplayedColumns, stage.DisplayedColumns_mapString, displayedcolumn, displayedcolumn.Name)
	return displayedcolumn
}

// UnstageVoid removes displayedcolumn off the model stage
func (displayedcolumn *DisplayedColumn) UnstageVoid(stage *Stage) {
	displayedcolumn.Unstage(stage)
}

func (displayedcolumn *DisplayedColumn) StageVoid(stage *Stage) {
	displayedcolumn.Stage(stage)
}

// for satisfaction of GongStruct interface
func (displayedcolumn *DisplayedColumn) GetName() (res string) {
	return displayedcolumn.Name
}

// for satisfaction of GongStruct interface
func (displayedcolumn *DisplayedColumn) SetName(name string) {
	displayedcolumn.Name = name
}

// Stage puts row to the model stage
func (row *Row) Stage(stage *Stage) *Row {
	__gong__stage(stage.Rows, stage.Row_stagedOrder, stage.Row_orderStaged, &stage.RowOrder, stage.Rows_mapString, row, row.Name)
	return row
}

// StagePreserveOrder puts row to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RowOrder
// - update stage.RowOrder accordingly
func (row *Row) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Rows, stage.Row_stagedOrder, stage.Row_orderStaged, &stage.RowOrder, stage.Rows_mapString, row, order, row.Name)
}

// Unstage removes row off the model stage
func (row *Row) Unstage(stage *Stage) *Row {
	__gong__unstage(stage.Rows, stage.Rows_mapString, row, row.Name)
	return row
}

// UnstageVoid removes row off the model stage
func (row *Row) UnstageVoid(stage *Stage) {
	row.Unstage(stage)
}

func (row *Row) StageVoid(stage *Stage) {
	row.Stage(stage)
}

// for satisfaction of GongStruct interface
func (row *Row) GetName() (res string) {
	return row.Name
}

// for satisfaction of GongStruct interface
func (row *Row) SetName(name string) {
	row.Name = name
}

// Stage puts svgicon to the model stage
func (svgicon *SVGIcon) Stage(stage *Stage) *SVGIcon {
	__gong__stage(stage.SVGIcons, stage.SVGIcon_stagedOrder, stage.SVGIcon_orderStaged, &stage.SVGIconOrder, stage.SVGIcons_mapString, svgicon, svgicon.Name)
	return svgicon
}

// StagePreserveOrder puts svgicon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SVGIconOrder
// - update stage.SVGIconOrder accordingly
func (svgicon *SVGIcon) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.SVGIcons, stage.SVGIcon_stagedOrder, stage.SVGIcon_orderStaged, &stage.SVGIconOrder, stage.SVGIcons_mapString, svgicon, order, svgicon.Name)
}

// Unstage removes svgicon off the model stage
func (svgicon *SVGIcon) Unstage(stage *Stage) *SVGIcon {
	__gong__unstage(stage.SVGIcons, stage.SVGIcons_mapString, svgicon, svgicon.Name)
	return svgicon
}

// UnstageVoid removes svgicon off the model stage
func (svgicon *SVGIcon) UnstageVoid(stage *Stage) {
	svgicon.Unstage(stage)
}

func (svgicon *SVGIcon) StageVoid(stage *Stage) {
	svgicon.Stage(stage)
}

// for satisfaction of GongStruct interface
func (svgicon *SVGIcon) GetName() (res string) {
	return svgicon.Name
}

// for satisfaction of GongStruct interface
func (svgicon *SVGIcon) SetName(name string) {
	svgicon.Name = name
}

// Stage puts table to the model stage
func (table *Table) Stage(stage *Stage) *Table {
	__gong__stage(stage.Tables, stage.Table_stagedOrder, stage.Table_orderStaged, &stage.TableOrder, stage.Tables_mapString, table, table.Name)
	return table
}

// StagePreserveOrder puts table to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableOrder
// - update stage.TableOrder accordingly
func (table *Table) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.Tables, stage.Table_stagedOrder, stage.Table_orderStaged, &stage.TableOrder, stage.Tables_mapString, table, order, table.Name)
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *Stage) *Table {
	__gong__unstage(stage.Tables, stage.Tables_mapString, table, table.Name)
	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *Stage) {
	table.Unstage(stage)
}

func (table *Table) StageVoid(stage *Stage) {
	table.Stage(stage)
}

// for satisfaction of GongStruct interface
func (table *Table) GetName() (res string) {
	return table.Name
}

// for satisfaction of GongStruct interface
func (table *Table) SetName(name string) {
	table.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.Buttons, &stage.Buttons_mapString, &stage.Button_stagedOrder, &stage.ButtonOrder)

	__gong__resetStageType(&stage.Cells, &stage.Cells_mapString, &stage.Cell_stagedOrder, &stage.CellOrder)

	__gong__resetStageType(&stage.CellBooleans, &stage.CellBooleans_mapString, &stage.CellBoolean_stagedOrder, &stage.CellBooleanOrder)

	__gong__resetStageType(&stage.CellFloat64s, &stage.CellFloat64s_mapString, &stage.CellFloat64_stagedOrder, &stage.CellFloat64Order)

	__gong__resetStageType(&stage.CellIcons, &stage.CellIcons_mapString, &stage.CellIcon_stagedOrder, &stage.CellIconOrder)

	__gong__resetStageType(&stage.CellInts, &stage.CellInts_mapString, &stage.CellInt_stagedOrder, &stage.CellIntOrder)

	__gong__resetStageType(&stage.CellStrings, &stage.CellStrings_mapString, &stage.CellString_stagedOrder, &stage.CellStringOrder)

	__gong__resetStageType(&stage.DisplayedColumns, &stage.DisplayedColumns_mapString, &stage.DisplayedColumn_stagedOrder, &stage.DisplayedColumnOrder)

	__gong__resetStageType(&stage.Rows, &stage.Rows_mapString, &stage.Row_stagedOrder, &stage.RowOrder)

	__gong__resetStageType(&stage.SVGIcons, &stage.SVGIcons_mapString, &stage.SVGIcon_stagedOrder, &stage.SVGIconOrder)

	__gong__resetStageType(&stage.Tables, &stage.Tables_mapString, &stage.Table_stagedOrder, &stage.TableOrder)

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
	case *Button:
		return any(stage.Buttons_mapString).(map[string]Type)
	case *Cell:
		return any(stage.Cells_mapString).(map[string]Type)
	case *CellBoolean:
		return any(stage.CellBooleans_mapString).(map[string]Type)
	case *CellFloat64:
		return any(stage.CellFloat64s_mapString).(map[string]Type)
	case *CellIcon:
		return any(stage.CellIcons_mapString).(map[string]Type)
	case *CellInt:
		return any(stage.CellInts_mapString).(map[string]Type)
	case *CellString:
		return any(stage.CellStrings_mapString).(map[string]Type)
	case *DisplayedColumn:
		return any(stage.DisplayedColumns_mapString).(map[string]Type)
	case *Row:
		return any(stage.Rows_mapString).(map[string]Type)
	case *SVGIcon:
		return any(stage.SVGIcons_mapString).(map[string]Type)
	case *Table:
		return any(stage.Tables_mapString).(map[string]Type)
	default:
		return nil
	}
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Button:
		return any(&stage.Buttons).(*map[Type]struct{})
	case *Cell:
		return any(&stage.Cells).(*map[Type]struct{})
	case *CellBoolean:
		return any(&stage.CellBooleans).(*map[Type]struct{})
	case *CellFloat64:
		return any(&stage.CellFloat64s).(*map[Type]struct{})
	case *CellIcon:
		return any(&stage.CellIcons).(*map[Type]struct{})
	case *CellInt:
		return any(&stage.CellInts).(*map[Type]struct{})
	case *CellString:
		return any(&stage.CellStrings).(*map[Type]struct{})
	case *DisplayedColumn:
		return any(&stage.DisplayedColumns).(*map[Type]struct{})
	case *Row:
		return any(&stage.Rows).(*map[Type]struct{})
	case *SVGIcon:
		return any(&stage.SVGIcons).(*map[Type]struct{})
	case *Table:
		return any(&stage.Tables).(*map[Type]struct{})
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
	case Button:
		return any(&Button{
			SVGIcon: &SVGIcon{Name: "SVGIcon"},
		}).(*Type)
	case Cell:
		return any(&Cell{
			CellString: &CellString{Name: "CellString"},
			CellFloat64: &CellFloat64{Name: "CellFloat64"},
			CellInt: &CellInt{Name: "CellInt"},
			CellBool: &CellBoolean{Name: "CellBool"},
			CellIcon: &CellIcon{Name: "CellIcon"},
		}).(*Type)
	case Row:
		return any(&Row{
			Cells: []*Cell{{Name: "Cells"}},
		}).(*Type)
	case Table:
		return any(&Table{
			DisplayedColumns: []*DisplayedColumn{{Name: "DisplayedColumns"}},
			Rows: []*Row{{Name: "Rows"}},
			RowsSelectedForBulkDelete: []*Row{{Name: "RowsSelectedForBulkDelete"}},
			Buttons: []*Button{{Name: "Buttons"}},
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
	// reverse maps of direct associations of Cell
	case Cell:
		switch fieldname {
		// insertion point for per direct association field
		case "CellString":
			res := make(map[*CellString][]*Cell)
			for cell := range stage.Cells {
				if cell.CellString != nil {
					cellstring_ := cell.CellString
					var cells []*Cell
					_, ok := res[cellstring_]
					if ok {
						cells = res[cellstring_]
					} else {
						cells = make([]*Cell, 0)
					}
					cells = append(cells, cell)
					res[cellstring_] = cells
				}
			}
			return any(res).(map[*End][]*Start)
		case "CellFloat64":
			res := make(map[*CellFloat64][]*Cell)
			for cell := range stage.Cells {
				if cell.CellFloat64 != nil {
					cellfloat64_ := cell.CellFloat64
					var cells []*Cell
					_, ok := res[cellfloat64_]
					if ok {
						cells = res[cellfloat64_]
					} else {
						cells = make([]*Cell, 0)
					}
					cells = append(cells, cell)
					res[cellfloat64_] = cells
				}
			}
			return any(res).(map[*End][]*Start)
		case "CellInt":
			res := make(map[*CellInt][]*Cell)
			for cell := range stage.Cells {
				if cell.CellInt != nil {
					cellint_ := cell.CellInt
					var cells []*Cell
					_, ok := res[cellint_]
					if ok {
						cells = res[cellint_]
					} else {
						cells = make([]*Cell, 0)
					}
					cells = append(cells, cell)
					res[cellint_] = cells
				}
			}
			return any(res).(map[*End][]*Start)
		case "CellBool":
			res := make(map[*CellBoolean][]*Cell)
			for cell := range stage.Cells {
				if cell.CellBool != nil {
					cellboolean_ := cell.CellBool
					var cells []*Cell
					_, ok := res[cellboolean_]
					if ok {
						cells = res[cellboolean_]
					} else {
						cells = make([]*Cell, 0)
					}
					cells = append(cells, cell)
					res[cellboolean_] = cells
				}
			}
			return any(res).(map[*End][]*Start)
		case "CellIcon":
			res := make(map[*CellIcon][]*Cell)
			for cell := range stage.Cells {
				if cell.CellIcon != nil {
					cellicon_ := cell.CellIcon
					var cells []*Cell
					_, ok := res[cellicon_]
					if ok {
						cells = res[cellicon_]
					} else {
						cells = make([]*Cell, 0)
					}
					cells = append(cells, cell)
					res[cellicon_] = cells
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of CellBoolean
	case CellBoolean:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellFloat64
	case CellFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellIcon
	case CellIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellInt
	case CellInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellString
	case CellString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DisplayedColumn
	case DisplayedColumn:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Row
	case Row:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of SVGIcon
	case SVGIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
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
	// reverse maps of direct associations of Button
	case Button:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Cell
	case Cell:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellBoolean
	case CellBoolean:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellFloat64
	case CellFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellIcon
	case CellIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellInt
	case CellInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of CellString
	case CellString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DisplayedColumn
	case DisplayedColumn:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Row
	case Row:
		switch fieldname {
		// insertion point for per direct association field
		case "Cells":
			res := make(map[*Cell][]*Row)
			for row := range stage.Rows {
				for _, cell_ := range row.Cells {
					res[cell_] = append(res[cell_], row)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of SVGIcon
	case SVGIcon:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		case "DisplayedColumns":
			res := make(map[*DisplayedColumn][]*Table)
			for table := range stage.Tables {
				for _, displayedcolumn_ := range table.DisplayedColumns {
					res[displayedcolumn_] = append(res[displayedcolumn_], table)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Rows":
			res := make(map[*Row][]*Table)
			for table := range stage.Tables {
				for _, row_ := range table.Rows {
					res[row_] = append(res[row_], table)
				}
			}
			return any(res).(map[*End][]*Start)
		case "RowsSelectedForBulkDelete":
			res := make(map[*Row][]*Table)
			for table := range stage.Tables {
				for _, row_ := range table.RowsSelectedForBulkDelete {
					res[row_] = append(res[row_], table)
				}
			}
			return any(res).(map[*End][]*Start)
		case "Buttons":
			res := make(map[*Button][]*Table)
			for table := range stage.Tables {
				for _, button_ := range table.Buttons {
					res[button_] = append(res[button_], table)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic new instance
	case *Button:
		res = any(new(Button)).(Type)
	case *Cell:
		res = any(new(Cell)).(Type)
	case *CellBoolean:
		res = any(new(CellBoolean)).(Type)
	case *CellFloat64:
		res = any(new(CellFloat64)).(Type)
	case *CellIcon:
		res = any(new(CellIcon)).(Type)
	case *CellInt:
		res = any(new(CellInt)).(Type)
	case *CellString:
		res = any(new(CellString)).(Type)
	case *DisplayedColumn:
		res = any(new(DisplayedColumn)).(Type)
	case *Row:
		res = any(new(Row)).(Type)
	case *SVGIcon:
		res = any(new(SVGIcon)).(Type)
	case *Table:
		res = any(new(Table)).(Type)
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
	case *Button:
		res = "Button"
	case *Cell:
		res = "Cell"
	case *CellBoolean:
		res = "CellBoolean"
	case *CellFloat64:
		res = "CellFloat64"
	case *CellIcon:
		res = "CellIcon"
	case *CellInt:
		res = "CellInt"
	case *CellString:
		res = "CellString"
	case *DisplayedColumn:
		res = "DisplayedColumn"
	case *Row:
		res = "Row"
	case *SVGIcon:
		res = "SVGIcon"
	case *Table:
		res = "Table"
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
	case *Button:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Table"
		rf.Fieldname = "Buttons"
		res = append(res, rf)
	case *Cell:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Row"
		rf.Fieldname = "Cells"
		res = append(res, rf)
	case *CellBoolean:
		var rf ReverseField
		_ = rf
	case *CellFloat64:
		var rf ReverseField
		_ = rf
	case *CellIcon:
		var rf ReverseField
		_ = rf
	case *CellInt:
		var rf ReverseField
		_ = rf
	case *CellString:
		var rf ReverseField
		_ = rf
	case *DisplayedColumn:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Table"
		rf.Fieldname = "DisplayedColumns"
		res = append(res, rf)
	case *Row:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Table"
		rf.Fieldname = "Rows"
		res = append(res, rf)
		rf.GongstructName = "Table"
		rf.Fieldname = "RowsSelectedForBulkDelete"
		res = append(res, rf)
	case *SVGIcon:
		var rf ReverseField
		_ = rf
	case *Table:
		var rf ReverseField
		_ = rf
	}
	return
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (button *Button) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Icon",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "SVGIcon",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "SVGIcon",
		},
		{
			Name:               "IsDisabled",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasToolTip",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ToolTipText",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "ToolTipPosition",
			GongFieldValueType:   GongFieldValueTypeString,
			TargetGongstructName: "ToolTipPositionEnum",
		},
	}
	return
}

func (cell *Cell) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "CellString",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CellString",
		},
		{
			Name:                 "CellFloat64",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CellFloat64",
		},
		{
			Name:                 "CellInt",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CellInt",
		},
		{
			Name:                 "CellBool",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CellBoolean",
		},
		{
			Name:                 "CellIcon",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "CellIcon",
		},
	}
	return
}

func (cellboolean *CellBoolean) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (cellfloat64 *CellFloat64) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeFloat,
		},
	}
	return
}

func (cellicon *CellIcon) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Icon",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NeedsConfirmation",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "ConfirmationMessage",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (cellint *CellInt) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (cellstring *CellString) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "Value",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (displayedcolumn *DisplayedColumn) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (row *Row) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "Cells",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Cell",
		},
		{
			Name:               "IsChecked",
			GongFieldValueType: GongFieldValueTypeBool,
		},
	}
	return
}

func (svgicon *SVGIcon) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "SVG",
			GongFieldValueType: GongFieldValueTypeString,
		},
	}
	return
}

func (table *Table) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "DisplayedColumns",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "DisplayedColumn",
		},
		{
			Name:                 "Rows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Row",
		},
		{
			Name:               "HasFiltering",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasColumnSorting",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasPaginator",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasCheckableRows",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasSaveButton",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "SaveButtonLabel",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "HasBulkDeleteButton",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "BulkDeleteButtonTooltip",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "RowsSelectedForBulkDelete",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Row",
		},
		{
			Name:               "CanDragDropRows",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "HasCloseButton",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "SavingInProgress",
			GongFieldValueType: GongFieldValueTypeBool,
		},
		{
			Name:               "NbOfStickyColumns",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Buttons",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "Button",
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
func (button *Button) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = button.Name
	case "Icon":
		res.valueString = button.Icon
	case "SVGIcon":
		res.GongFieldValueType = GongFieldValueTypePointer
		if button.SVGIcon != nil {
			res.valueString = button.SVGIcon.Name
			res.ids = button.SVGIcon.GongGetUUID(stage)
		}
	case "IsDisabled":
		res.valueString = fmt.Sprintf("%t", button.IsDisabled)
		res.valueBool = button.IsDisabled
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasToolTip":
		res.valueString = fmt.Sprintf("%t", button.HasToolTip)
		res.valueBool = button.HasToolTip
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ToolTipText":
		res.valueString = button.ToolTipText
	case "ToolTipPosition":
		enum := button.ToolTipPosition
		res.valueString = enum.ToCodeString()
	}
	return
}

func (cell *Cell) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cell.Name
	case "CellString":
		res.GongFieldValueType = GongFieldValueTypePointer
		if cell.CellString != nil {
			res.valueString = cell.CellString.Name
			res.ids = cell.CellString.GongGetUUID(stage)
		}
	case "CellFloat64":
		res.GongFieldValueType = GongFieldValueTypePointer
		if cell.CellFloat64 != nil {
			res.valueString = cell.CellFloat64.Name
			res.ids = cell.CellFloat64.GongGetUUID(stage)
		}
	case "CellInt":
		res.GongFieldValueType = GongFieldValueTypePointer
		if cell.CellInt != nil {
			res.valueString = cell.CellInt.Name
			res.ids = cell.CellInt.GongGetUUID(stage)
		}
	case "CellBool":
		res.GongFieldValueType = GongFieldValueTypePointer
		if cell.CellBool != nil {
			res.valueString = cell.CellBool.Name
			res.ids = cell.CellBool.GongGetUUID(stage)
		}
	case "CellIcon":
		res.GongFieldValueType = GongFieldValueTypePointer
		if cell.CellIcon != nil {
			res.valueString = cell.CellIcon.Name
			res.ids = cell.CellIcon.GongGetUUID(stage)
		}
	}
	return
}

func (cellboolean *CellBoolean) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cellboolean.Name
	case "Value":
		res.valueString = fmt.Sprintf("%t", cellboolean.Value)
		res.valueBool = cellboolean.Value
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (cellfloat64 *CellFloat64) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cellfloat64.Name
	case "Value":
		res.valueString = fmt.Sprintf("%f", cellfloat64.Value)
		res.valueFloat = cellfloat64.Value
		res.GongFieldValueType = GongFieldValueTypeFloat
	}
	return
}

func (cellicon *CellIcon) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cellicon.Name
	case "Icon":
		res.valueString = cellicon.Icon
	case "NeedsConfirmation":
		res.valueString = fmt.Sprintf("%t", cellicon.NeedsConfirmation)
		res.valueBool = cellicon.NeedsConfirmation
		res.GongFieldValueType = GongFieldValueTypeBool
	case "ConfirmationMessage":
		res.valueString = cellicon.ConfirmationMessage
	}
	return
}

func (cellint *CellInt) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cellint.Name
	case "Value":
		res.valueString = fmt.Sprintf("%d", cellint.Value)
		res.valueInt = cellint.Value
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (cellstring *CellString) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = cellstring.Name
	case "Value":
		res.valueString = cellstring.Value
	}
	return
}

func (displayedcolumn *DisplayedColumn) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = displayedcolumn.Name
	}
	return
}

func (row *Row) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = row.Name
	case "Cells":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range row.Cells {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "IsChecked":
		res.valueString = fmt.Sprintf("%t", row.IsChecked)
		res.valueBool = row.IsChecked
		res.GongFieldValueType = GongFieldValueTypeBool
	}
	return
}

func (svgicon *SVGIcon) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = svgicon.Name
	case "SVG":
		res.valueString = svgicon.SVG
	}
	return
}

func (table *Table) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = table.Name
	case "DisplayedColumns":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range table.DisplayedColumns {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "Rows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range table.Rows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "HasFiltering":
		res.valueString = fmt.Sprintf("%t", table.HasFiltering)
		res.valueBool = table.HasFiltering
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasColumnSorting":
		res.valueString = fmt.Sprintf("%t", table.HasColumnSorting)
		res.valueBool = table.HasColumnSorting
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasPaginator":
		res.valueString = fmt.Sprintf("%t", table.HasPaginator)
		res.valueBool = table.HasPaginator
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasCheckableRows":
		res.valueString = fmt.Sprintf("%t", table.HasCheckableRows)
		res.valueBool = table.HasCheckableRows
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasSaveButton":
		res.valueString = fmt.Sprintf("%t", table.HasSaveButton)
		res.valueBool = table.HasSaveButton
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SaveButtonLabel":
		res.valueString = table.SaveButtonLabel
	case "HasBulkDeleteButton":
		res.valueString = fmt.Sprintf("%t", table.HasBulkDeleteButton)
		res.valueBool = table.HasBulkDeleteButton
		res.GongFieldValueType = GongFieldValueTypeBool
	case "BulkDeleteButtonTooltip":
		res.valueString = table.BulkDeleteButtonTooltip
	case "RowsSelectedForBulkDelete":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range table.RowsSelectedForBulkDelete {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "CanDragDropRows":
		res.valueString = fmt.Sprintf("%t", table.CanDragDropRows)
		res.valueBool = table.CanDragDropRows
		res.GongFieldValueType = GongFieldValueTypeBool
	case "HasCloseButton":
		res.valueString = fmt.Sprintf("%t", table.HasCloseButton)
		res.valueBool = table.HasCloseButton
		res.GongFieldValueType = GongFieldValueTypeBool
	case "SavingInProgress":
		res.valueString = fmt.Sprintf("%t", table.SavingInProgress)
		res.valueBool = table.SavingInProgress
		res.GongFieldValueType = GongFieldValueTypeBool
	case "NbOfStickyColumns":
		res.valueString = fmt.Sprintf("%d", table.NbOfStickyColumns)
		res.valueInt = table.NbOfStickyColumns
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Buttons":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range table.Buttons {
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

func (stage *Stage) GetFieldStringValueFromPointer(instance GongstructIF, fieldName string) (res GongFieldValue) {
	res = instance.GongGetFieldValue(fieldName, stage)
	return
}

func GetFieldStringValueFromPointer(instance GongstructIF, fieldName string, stage *Stage) (res GongFieldValue) {
	return stage.GetFieldStringValueFromPointer(instance, fieldName)
}

// insertion point for generic get gongstruct name
func (button *Button) GongGetGongstructName() string {
	return "Button"
}

func (cell *Cell) GongGetGongstructName() string {
	return "Cell"
}

func (cellboolean *CellBoolean) GongGetGongstructName() string {
	return "CellBoolean"
}

func (cellfloat64 *CellFloat64) GongGetGongstructName() string {
	return "CellFloat64"
}

func (cellicon *CellIcon) GongGetGongstructName() string {
	return "CellIcon"
}

func (cellint *CellInt) GongGetGongstructName() string {
	return "CellInt"
}

func (cellstring *CellString) GongGetGongstructName() string {
	return "CellString"
}

func (displayedcolumn *DisplayedColumn) GongGetGongstructName() string {
	return "DisplayedColumn"
}

func (row *Row) GongGetGongstructName() string {
	return "Row"
}

func (svgicon *SVGIcon) GongGetGongstructName() string {
	return "SVGIcon"
}

func (table *Table) GongGetGongstructName() string {
	return "Table"
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
	__gong__rebuildMapString(stage.Buttons, &stage.Buttons_mapString)

	__gong__rebuildMapString(stage.Cells, &stage.Cells_mapString)

	__gong__rebuildMapString(stage.CellBooleans, &stage.CellBooleans_mapString)

	__gong__rebuildMapString(stage.CellFloat64s, &stage.CellFloat64s_mapString)

	__gong__rebuildMapString(stage.CellIcons, &stage.CellIcons_mapString)

	__gong__rebuildMapString(stage.CellInts, &stage.CellInts_mapString)

	__gong__rebuildMapString(stage.CellStrings, &stage.CellStrings_mapString)

	__gong__rebuildMapString(stage.DisplayedColumns, &stage.DisplayedColumns_mapString)

	__gong__rebuildMapString(stage.Rows, &stage.Rows_mapString)

	__gong__rebuildMapString(stage.SVGIcons, &stage.SVGIcons_mapString)

	__gong__rebuildMapString(stage.Tables, &stage.Tables_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
