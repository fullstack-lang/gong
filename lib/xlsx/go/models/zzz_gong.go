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
	"time"

	gong_runtime "github.com/fullstack-lang/gong/pkg/runtime"
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

type GongMarshallingMode = gong_runtime.GongMarshallingMode

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal = gong_runtime.GongMarshallingNormal

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit = gong_runtime.GongMarshallingAppendCommit
)

type gongStageNavigationMode = gong_runtime.StageNavigationMode

const (
	GongNavigationModeNormal     = gong_runtime.GongNavigationModeNormal
	GongNavigationModeNavigating = gong_runtime.GongNavigationModeNavigating
)

const (
	GongProbeTreeSidebarSuffix           = gong_runtime.GongProbeTreeSidebarSuffix
	GongProbeNavigationTreeSidebarSuffix = gong_runtime.GongProbeNavigationTreeSidebarSuffix
	GongProbeTableSuffix                 = gong_runtime.GongProbeTableSuffix
	GongProbeNotificationTableSuffix     = gong_runtime.GongProbeNotificationTableSuffix
	GongProbeFormSuffix                  = gong_runtime.GongProbeFormSuffix
	GongProbeSplitSuffix                 = gong_runtime.GongProbeSplitSuffix
	GongProbeLoadSuffix                  = gong_runtime.GongProbeLoadSuffix
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
	gong_runtime.StageCore

	// insertion point for definition of arrays registering instances
	DisplaySelections                map[*DisplaySelection]struct{}
	DisplaySelections_instance       map[*DisplaySelection]*DisplaySelection
	DisplaySelections_mapString      map[string]*DisplaySelection
	DisplaySelectionOrder            uint
	DisplaySelection_stagedOrder     map[*DisplaySelection]uint
	DisplaySelection_orderStaged     map[uint]*DisplaySelection
	DisplaySelections_reference      map[*DisplaySelection]*DisplaySelection
	DisplaySelections_referenceOrder map[*DisplaySelection]uint

	// insertion point for slice of pointers maps
	OnAfterDisplaySelectionCreateCallback GongOnAfterCreateInterface[DisplaySelection]
	OnAfterDisplaySelectionUpdateCallback GongOnAfterUpdateInterface[DisplaySelection]
	OnAfterDisplaySelectionDeleteCallback GongOnAfterDeleteInterface[DisplaySelection]

	XLCells                map[*XLCell]struct{}
	XLCells_instance       map[*XLCell]*XLCell
	XLCells_mapString      map[string]*XLCell
	XLCellOrder            uint
	XLCell_stagedOrder     map[*XLCell]uint
	XLCell_orderStaged     map[uint]*XLCell
	XLCells_reference      map[*XLCell]*XLCell
	XLCells_referenceOrder map[*XLCell]uint

	// insertion point for slice of pointers maps
	OnAfterXLCellCreateCallback GongOnAfterCreateInterface[XLCell]
	OnAfterXLCellUpdateCallback GongOnAfterUpdateInterface[XLCell]
	OnAfterXLCellDeleteCallback GongOnAfterDeleteInterface[XLCell]

	XLFiles                map[*XLFile]struct{}
	XLFiles_instance       map[*XLFile]*XLFile
	XLFiles_mapString      map[string]*XLFile
	XLFileOrder            uint
	XLFile_stagedOrder     map[*XLFile]uint
	XLFile_orderStaged     map[uint]*XLFile
	XLFiles_reference      map[*XLFile]*XLFile
	XLFiles_referenceOrder map[*XLFile]uint

	// insertion point for slice of pointers maps
	XLFile_Sheets_reverseMap map[*XLSheet]*XLFile

	OnAfterXLFileCreateCallback GongOnAfterCreateInterface[XLFile]
	OnAfterXLFileUpdateCallback GongOnAfterUpdateInterface[XLFile]
	OnAfterXLFileDeleteCallback GongOnAfterDeleteInterface[XLFile]

	XLRows                map[*XLRow]struct{}
	XLRows_instance       map[*XLRow]*XLRow
	XLRows_mapString      map[string]*XLRow
	XLRowOrder            uint
	XLRow_stagedOrder     map[*XLRow]uint
	XLRow_orderStaged     map[uint]*XLRow
	XLRows_reference      map[*XLRow]*XLRow
	XLRows_referenceOrder map[*XLRow]uint

	// insertion point for slice of pointers maps
	XLRow_Cells_reverseMap map[*XLCell]*XLRow

	OnAfterXLRowCreateCallback GongOnAfterCreateInterface[XLRow]
	OnAfterXLRowUpdateCallback GongOnAfterUpdateInterface[XLRow]
	OnAfterXLRowDeleteCallback GongOnAfterDeleteInterface[XLRow]

	XLSheets                map[*XLSheet]struct{}
	XLSheets_instance       map[*XLSheet]*XLSheet
	XLSheets_mapString      map[string]*XLSheet
	XLSheetOrder            uint
	XLSheet_stagedOrder     map[*XLSheet]uint
	XLSheet_orderStaged     map[uint]*XLSheet
	XLSheets_reference      map[*XLSheet]*XLSheet
	XLSheets_referenceOrder map[*XLSheet]uint

	// insertion point for slice of pointers maps
	XLSheet_Rows_reverseMap map[*XLRow]*XLSheet

	XLSheet_SheetCells_reverseMap map[*XLCell]*XLSheet

	OnAfterXLSheetCreateCallback GongOnAfterCreateInterface[XLSheet]
	OnAfterXLSheetUpdateCallback GongOnAfterUpdateInterface[XLSheet]
	OnAfterXLSheetDeleteCallback GongOnAfterDeleteInterface[XLSheet]

	BackRepo GongBackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          GongOnInitCommitInterface
	OnInitCommitFromFrontCallback GongOnInitCommitInterface
	OnInitCommitFromBackCallback  GongOnInitCommitInterface

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
}

type GongStage = Stage

// RegisterBeforeCommit adds a hook that runs before the commit happens
func (s *Stage) RegisterBeforeCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterBeforeCommitHook(func() { hook(s) })
}

// RegisterAfterCommit adds a hook that runs after the commit succeeds
func (s *Stage) RegisterAfterCommit(hook func(stage *Stage)) {
	s.StageCore.RegisterAfterCommitHook(func() { hook(s) })
}

// ApplyBackwardCommit applies the commit before the current one
func (stage *Stage) ApplyBackwardCommit() error {
	return stage.StageCore.ApplyBackwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ApplyForwardCommit applies the commit after the current one
func (stage *Stage) ApplyForwardCommit() error {
	return stage.StageCore.ApplyForwardCommit(stage.ParseAstString, stage.ComputeReferenceAndOrders)
}

// ResetHard removes the more recent
// commitsBehind forward/backward Commits from the
// stage
func (stage *Stage) ResetHard() {
	stage.StageCore.ResetHardCore()

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// Squash removes all commits and marshals the stage as a single commit
func (stage *Stage) Squash() {
	stage.StageCore.SquashCore()

	// insertion point for clear references
	__gong__clearReferences(&stage.DisplaySelections_reference, &stage.DisplaySelections_instance, &stage.DisplaySelections_referenceOrder)

	__gong__clearReferences(&stage.XLCells_reference, &stage.XLCells_instance, &stage.XLCells_referenceOrder)

	__gong__clearReferences(&stage.XLFiles_reference, &stage.XLFiles_instance, &stage.XLFiles_referenceOrder)

	__gong__clearReferences(&stage.XLRows_reference, &stage.XLRows_instance, &stage.XLRows_referenceOrder)

	__gong__clearReferences(&stage.XLSheets_reference, &stage.XLSheets_instance, &stage.XLSheets_referenceOrder)

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}

	// 1. Run all Before Commit hooks
	stage.StageCore.RunBeforeCommitHooks()

	stage.StageCore.EndSquash()

	// 2. Run all After Commit hooks
	stage.StageCore.RunAfterCommitHooks()
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	stage.DisplaySelectionOrder = __gong__recomputeOrder(stage.DisplaySelection_stagedOrder)

	stage.XLCellOrder = __gong__recomputeOrder(stage.XLCell_stagedOrder)

	stage.XLFileOrder = __gong__recomputeOrder(stage.XLFile_stagedOrder)

	stage.XLRowOrder = __gong__recomputeOrder(stage.XLRow_stagedOrder)

	stage.XLSheetOrder = __gong__recomputeOrder(stage.XLSheet_stagedOrder)

	// end of insertion point for max order recomputation
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

// insertion point for stage ops
func (*DisplaySelection) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.DisplaySelections, stage.DisplaySelection_stagedOrder)
}

func (*DisplaySelection) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.DisplaySelection_orderStaged[order]
}

func (*DisplaySelection) GongGetInstancesMapByName(stage *Stage) any {
	return stage.DisplaySelections_mapString
}

func (*DisplaySelection) GongGetInstancesSet(stage *Stage) any {
	return &stage.DisplaySelections
}

func (*DisplaySelection) GongNewInstance() any {
	return new(DisplaySelection)
}

func (*XLCell) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.XLCells, stage.XLCell_stagedOrder)
}

func (*XLCell) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.XLCell_orderStaged[order]
}

func (*XLCell) GongGetInstancesMapByName(stage *Stage) any {
	return stage.XLCells_mapString
}

func (*XLCell) GongGetInstancesSet(stage *Stage) any {
	return &stage.XLCells
}

func (*XLCell) GongNewInstance() any {
	return new(XLCell)
}

func (*XLFile) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.XLFiles, stage.XLFile_stagedOrder)
}

func (*XLFile) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.XLFile_orderStaged[order]
}

func (*XLFile) GongGetInstancesMapByName(stage *Stage) any {
	return stage.XLFiles_mapString
}

func (*XLFile) GongGetInstancesSet(stage *Stage) any {
	return &stage.XLFiles
}

func (*XLFile) GongNewInstance() any {
	return new(XLFile)
}

func (*XLRow) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.XLRows, stage.XLRow_stagedOrder)
}

func (*XLRow) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.XLRow_orderStaged[order]
}

func (*XLRow) GongGetInstancesMapByName(stage *Stage) any {
	return stage.XLRows_mapString
}

func (*XLRow) GongGetInstancesSet(stage *Stage) any {
	return &stage.XLRows
}

func (*XLRow) GongNewInstance() any {
	return new(XLRow)
}

func (*XLSheet) GongGetInstancesByOrder(stage *Stage) any {
	return __gong__getStructInstancesByOrder(stage.XLSheets, stage.XLSheet_stagedOrder)
}

func (*XLSheet) GongGetInstanceFromOrder(stage *Stage, order uint) any {
	return stage.XLSheet_orderStaged[order]
}

func (*XLSheet) GongGetInstancesMapByName(stage *Stage) any {
	return stage.XLSheets_mapString
}

func (*XLSheet) GongGetInstancesSet(stage *Stage) any {
	return &stage.XLSheets
}

func (*XLSheet) GongNewInstance() any {
	return new(XLSheet)
}


// GetInstancesByOrder is the Stage method returning a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func (stage *Stage) GetInstancesByOrder[T GongstructPtr]() (res []T) {
	if stage == nil {
		return nil
	}
	var t T
	return t.GongGetInstancesByOrder(stage).([]T)
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
	return "github.com/fullstack-lang/gong/lib/xlsx/go/models"
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
		DisplaySelections:           make(map[*DisplaySelection]struct{}),
		DisplaySelections_mapString: make(map[string]*DisplaySelection),

		XLCells:           make(map[*XLCell]struct{}),
		XLCells_mapString: make(map[string]*XLCell),

		XLFiles:           make(map[*XLFile]struct{}),
		XLFiles_mapString: make(map[string]*XLFile),

		XLRows:           make(map[*XLRow]struct{}),
		XLRows_mapString: make(map[string]*XLRow),

		XLSheets:           make(map[*XLSheet]struct{}),
		XLSheets_mapString: make(map[string]*XLSheet),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		DisplaySelection_stagedOrder: make(map[*DisplaySelection]uint),
		DisplaySelection_orderStaged: make(map[uint]*DisplaySelection),
		DisplaySelections_reference:  make(map[*DisplaySelection]*DisplaySelection),

		XLCell_stagedOrder: make(map[*XLCell]uint),
		XLCell_orderStaged: make(map[uint]*XLCell),
		XLCells_reference:  make(map[*XLCell]*XLCell),

		XLFile_stagedOrder: make(map[*XLFile]uint),
		XLFile_orderStaged: make(map[uint]*XLFile),
		XLFiles_reference:  make(map[*XLFile]*XLFile),

		XLRow_stagedOrder: make(map[*XLRow]uint),
		XLRow_orderStaged: make(map[uint]*XLRow),
		XLRows_reference:  make(map[*XLRow]*XLRow),

		XLSheet_stagedOrder: make(map[*XLSheet]uint),
		XLSheet_orderStaged: make(map[uint]*XLSheet),
		XLSheets_reference:  make(map[*XLSheet]*XLSheet),

		// end of insertion point
		GongUnmarshallers: map[string]GongModelUnmarshaller{ // insertion point for unmarshallers
			"DisplaySelection": &DisplaySelectionUnmarshaller{},

			"XLCell": &XLCellUnmarshaller{},

			"XLFile": &XLFileUnmarshaller{},

			"XLRow": &XLRowUnmarshaller{},

			"XLSheet": &XLSheetUnmarshaller{},

			// end of insertion point
		},
	}
	stage.StageCore.SetName(name)
	stage.StageCore.SetNavigationMode(GongNavigationModeNormal)

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
	if stage == nil {
		return
	}
	var t Type
	val := t.GongGetInstanceFromOrder(stage, order)
	if val != nil {
		res = val.(Type)
	}
	return
}

func (stage *Stage) CommitWithSuspendedCallbacks() {
	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	tmp2 := stage.StageCore.GetBeforeCommitHooks()
	stage.StageCore.ClearBeforeCommitHooks()
	tmp3 := stage.StageCore.GetAfterCommitHooks()
	stage.StageCore.ClearAfterCommitHooks()
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
	stage.StageCore.SetBeforeCommitHooks(tmp2)
	stage.StageCore.SetAfterCommitHooks(tmp3)
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
	stage.StageCore.RunBeforeCommitHooks()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}
	stage.ComputeInstancesNb()

	// if a commit is applied when in navigation mode
	// this will reset the commits behind and swith the
	// naviagation
	if stage.IsInDeltaMode() && stage.GetNavigationMode() == GongNavigationModeNavigating && stage.GetCommitsBehind() > 0 {
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
	stage.StageCore.RunAfterCommitHooks()
}

func (stage *Stage) ComputeInstancesNb() {
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["DisplaySelection"] = len(stage.DisplaySelections)
	stage.Map_GongStructName_InstancesNb["XLCell"] = len(stage.XLCells)
	stage.Map_GongStructName_InstancesNb["XLFile"] = len(stage.XLFiles)
	stage.Map_GongStructName_InstancesNb["XLRow"] = len(stage.XLRows)
	stage.Map_GongStructName_InstancesNb["XLSheet"] = len(stage.XLSheets)
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
// Stage puts displayselection to the model stage
func (displayselection *DisplaySelection) Stage(stage *Stage) *DisplaySelection {
	__gong__stage(stage.DisplaySelections, stage.DisplaySelection_stagedOrder, stage.DisplaySelection_orderStaged, &stage.DisplaySelectionOrder, stage.DisplaySelections_mapString, displayselection, displayselection.Name)
	return displayselection
}

// StagePreserveOrder puts displayselection to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DisplaySelectionOrder
// - update stage.DisplaySelectionOrder accordingly
func (displayselection *DisplaySelection) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.DisplaySelections, stage.DisplaySelection_stagedOrder, stage.DisplaySelection_orderStaged, &stage.DisplaySelectionOrder, stage.DisplaySelections_mapString, displayselection, order, displayselection.Name)
}

// Unstage removes displayselection off the model stage
func (displayselection *DisplaySelection) Unstage(stage *Stage) *DisplaySelection {
	__gong__unstage(stage.DisplaySelections, stage.DisplaySelections_mapString, displayselection, displayselection.Name)
	return displayselection
}

// UnstageVoid removes displayselection off the model stage
func (displayselection *DisplaySelection) UnstageVoid(stage *Stage) {
	displayselection.Unstage(stage)
}

func (displayselection *DisplaySelection) StageVoid(stage *Stage) {
	displayselection.Stage(stage)
}

// for satisfaction of GongStruct interface
func (displayselection *DisplaySelection) GetName() (res string) {
	return displayselection.Name
}

// for satisfaction of GongStruct interface
func (displayselection *DisplaySelection) SetName(name string) {
	displayselection.Name = name
}

// Stage puts xlcell to the model stage
func (xlcell *XLCell) Stage(stage *Stage) *XLCell {
	__gong__stage(stage.XLCells, stage.XLCell_stagedOrder, stage.XLCell_orderStaged, &stage.XLCellOrder, stage.XLCells_mapString, xlcell, xlcell.Name)
	return xlcell
}

// StagePreserveOrder puts xlcell to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLCellOrder
// - update stage.XLCellOrder accordingly
func (xlcell *XLCell) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.XLCells, stage.XLCell_stagedOrder, stage.XLCell_orderStaged, &stage.XLCellOrder, stage.XLCells_mapString, xlcell, order, xlcell.Name)
}

// Unstage removes xlcell off the model stage
func (xlcell *XLCell) Unstage(stage *Stage) *XLCell {
	__gong__unstage(stage.XLCells, stage.XLCells_mapString, xlcell, xlcell.Name)
	return xlcell
}

// UnstageVoid removes xlcell off the model stage
func (xlcell *XLCell) UnstageVoid(stage *Stage) {
	xlcell.Unstage(stage)
}

func (xlcell *XLCell) StageVoid(stage *Stage) {
	xlcell.Stage(stage)
}

// for satisfaction of GongStruct interface
func (xlcell *XLCell) GetName() (res string) {
	return xlcell.Name
}

// for satisfaction of GongStruct interface
func (xlcell *XLCell) SetName(name string) {
	xlcell.Name = name
}

// Stage puts xlfile to the model stage
func (xlfile *XLFile) Stage(stage *Stage) *XLFile {
	__gong__stage(stage.XLFiles, stage.XLFile_stagedOrder, stage.XLFile_orderStaged, &stage.XLFileOrder, stage.XLFiles_mapString, xlfile, xlfile.Name)
	return xlfile
}

// StagePreserveOrder puts xlfile to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLFileOrder
// - update stage.XLFileOrder accordingly
func (xlfile *XLFile) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.XLFiles, stage.XLFile_stagedOrder, stage.XLFile_orderStaged, &stage.XLFileOrder, stage.XLFiles_mapString, xlfile, order, xlfile.Name)
}

// Unstage removes xlfile off the model stage
func (xlfile *XLFile) Unstage(stage *Stage) *XLFile {
	__gong__unstage(stage.XLFiles, stage.XLFiles_mapString, xlfile, xlfile.Name)
	return xlfile
}

// UnstageVoid removes xlfile off the model stage
func (xlfile *XLFile) UnstageVoid(stage *Stage) {
	xlfile.Unstage(stage)
}

func (xlfile *XLFile) StageVoid(stage *Stage) {
	xlfile.Stage(stage)
}

// for satisfaction of GongStruct interface
func (xlfile *XLFile) GetName() (res string) {
	return xlfile.Name
}

// for satisfaction of GongStruct interface
func (xlfile *XLFile) SetName(name string) {
	xlfile.Name = name
}

// Stage puts xlrow to the model stage
func (xlrow *XLRow) Stage(stage *Stage) *XLRow {
	__gong__stage(stage.XLRows, stage.XLRow_stagedOrder, stage.XLRow_orderStaged, &stage.XLRowOrder, stage.XLRows_mapString, xlrow, xlrow.Name)
	return xlrow
}

// StagePreserveOrder puts xlrow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLRowOrder
// - update stage.XLRowOrder accordingly
func (xlrow *XLRow) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.XLRows, stage.XLRow_stagedOrder, stage.XLRow_orderStaged, &stage.XLRowOrder, stage.XLRows_mapString, xlrow, order, xlrow.Name)
}

// Unstage removes xlrow off the model stage
func (xlrow *XLRow) Unstage(stage *Stage) *XLRow {
	__gong__unstage(stage.XLRows, stage.XLRows_mapString, xlrow, xlrow.Name)
	return xlrow
}

// UnstageVoid removes xlrow off the model stage
func (xlrow *XLRow) UnstageVoid(stage *Stage) {
	xlrow.Unstage(stage)
}

func (xlrow *XLRow) StageVoid(stage *Stage) {
	xlrow.Stage(stage)
}

// for satisfaction of GongStruct interface
func (xlrow *XLRow) GetName() (res string) {
	return xlrow.Name
}

// for satisfaction of GongStruct interface
func (xlrow *XLRow) SetName(name string) {
	xlrow.Name = name
}

// Stage puts xlsheet to the model stage
func (xlsheet *XLSheet) Stage(stage *Stage) *XLSheet {
	__gong__stage(stage.XLSheets, stage.XLSheet_stagedOrder, stage.XLSheet_orderStaged, &stage.XLSheetOrder, stage.XLSheets_mapString, xlsheet, xlsheet.Name)
	return xlsheet
}

// StagePreserveOrder puts xlsheet to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLSheetOrder
// - update stage.XLSheetOrder accordingly
func (xlsheet *XLSheet) StagePreserveOrder(stage *Stage, order uint) {
	__gong__stagePreserveOrder(stage.XLSheets, stage.XLSheet_stagedOrder, stage.XLSheet_orderStaged, &stage.XLSheetOrder, stage.XLSheets_mapString, xlsheet, order, xlsheet.Name)
}

// Unstage removes xlsheet off the model stage
func (xlsheet *XLSheet) Unstage(stage *Stage) *XLSheet {
	__gong__unstage(stage.XLSheets, stage.XLSheets_mapString, xlsheet, xlsheet.Name)
	return xlsheet
}

// UnstageVoid removes xlsheet off the model stage
func (xlsheet *XLSheet) UnstageVoid(stage *Stage) {
	xlsheet.Unstage(stage)
}

func (xlsheet *XLSheet) StageVoid(stage *Stage) {
	xlsheet.Stage(stage)
}

// for satisfaction of GongStruct interface
func (xlsheet *XLSheet) GetName() (res string) {
	return xlsheet.Name
}

// for satisfaction of GongStruct interface
func (xlsheet *XLSheet) SetName(name string) {
	xlsheet.Name = name
}

func (stage *Stage) Reset() { // insertion point for array reset
	__gong__resetStageType(&stage.DisplaySelections, &stage.DisplaySelections_mapString, &stage.DisplaySelection_stagedOrder, &stage.DisplaySelectionOrder)

	__gong__resetStageType(&stage.XLCells, &stage.XLCells_mapString, &stage.XLCell_stagedOrder, &stage.XLCellOrder)

	__gong__resetStageType(&stage.XLFiles, &stage.XLFiles_mapString, &stage.XLFile_stagedOrder, &stage.XLFileOrder)

	__gong__resetStageType(&stage.XLRows, &stage.XLRows_mapString, &stage.XLRow_stagedOrder, &stage.XLRowOrder)

	__gong__resetStageType(&stage.XLSheets, &stage.XLSheets_mapString, &stage.XLSheet_stagedOrder, &stage.XLSheetOrder)

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
type Gongstruct interface {
	GongGetAssociationName() any
}

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

	GongGetInstancesByOrder(stage *Stage) any
	GongGetInstanceFromOrder(stage *Stage, order uint) any
	GongGetInstancesMapByName(stage *Stage) any
	GongGetInstancesSet(stage *Stage) any
	GongNewInstance() any
	GongGetReverseFields() []GongReverseField
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
	if stage == nil {
		return nil
	}
	var t Type
	return t.GongGetInstancesMapByName(stage).(map[string]Type)
}

// GetInstancesSet is the Stage method returning the set of staged instances (pointer-type constraint).
func (stage *Stage) GetInstancesSet[Type GongstructPtr]() *map[Type]struct{} {
	if stage == nil {
		return nil
	}
	var t Type
	return t.GongGetInstancesSet(stage).(*map[Type]struct{})
}

// insertion point for instance with special fields
func (DisplaySelection) GongGetAssociationName() any {
	return &DisplaySelection{
			XLFile: &XLFile{Name: "XLFile"},
			XLSheet: &XLSheet{Name: "XLSheet"},
	}
}

func (XLCell) GongGetAssociationName() any {
	return &XLCell{
	}
}

func (XLFile) GongGetAssociationName() any {
	return &XLFile{
			Sheets: []*XLSheet{{Name: "Sheets"}},
	}
}

func (XLRow) GongGetAssociationName() any {
	return &XLRow{
			Cells: []*XLCell{{Name: "Cells"}},
	}
}

func (XLSheet) GongGetAssociationName() any {
	return &XLSheet{
			Rows: []*XLRow{{Name: "Rows"}},
			SheetCells: []*XLCell{{Name: "SheetCells"}},
	}
}


// GongGetAssociationName is a generic function that returns an instance of Type
// where each association is filled with an instance whose name is the name of the association
//
// This function can be handy for generating navigation function that are refactorable
func GongGetAssociationName[Type Gongstruct]() *Type {
	var t Type
	return t.GongGetAssociationName().(*Type)
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
	// reverse maps of direct associations of DisplaySelection
	case DisplaySelection:
		switch fieldname {
		// insertion point for per direct association field
		case "XLFile":
			res := make(map[*XLFile][]*DisplaySelection)
			for displayselection := range stage.DisplaySelections {
				if displayselection.XLFile != nil {
					xlfile_ := displayselection.XLFile
					var displayselections []*DisplaySelection
					_, ok := res[xlfile_]
					if ok {
						displayselections = res[xlfile_]
					} else {
						displayselections = make([]*DisplaySelection, 0)
					}
					displayselections = append(displayselections, displayselection)
					res[xlfile_] = displayselections
				}
			}
			return any(res).(map[*End][]*Start)
		case "XLSheet":
			res := make(map[*XLSheet][]*DisplaySelection)
			for displayselection := range stage.DisplaySelections {
				if displayselection.XLSheet != nil {
					xlsheet_ := displayselection.XLSheet
					var displayselections []*DisplaySelection
					_, ok := res[xlsheet_]
					if ok {
						displayselections = res[xlsheet_]
					} else {
						displayselections = make([]*DisplaySelection, 0)
					}
					displayselections = append(displayselections, displayselection)
					res[xlsheet_] = displayselections
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of XLCell
	case XLCell:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XLFile
	case XLFile:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XLRow
	case XLRow:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XLSheet
	case XLSheet:
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
	// reverse maps of direct associations of DisplaySelection
	case DisplaySelection:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XLCell
	case XLCell:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of XLFile
	case XLFile:
		switch fieldname {
		// insertion point for per direct association field
		case "Sheets":
			res := make(map[*XLSheet][]*XLFile)
			for xlfile := range stage.XLFiles {
				for _, xlsheet_ := range xlfile.Sheets {
					res[xlsheet_] = append(res[xlsheet_], xlfile)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of XLRow
	case XLRow:
		switch fieldname {
		// insertion point for per direct association field
		case "Cells":
			res := make(map[*XLCell][]*XLRow)
			for xlrow := range stage.XLRows {
				for _, xlcell_ := range xlrow.Cells {
					res[xlcell_] = append(res[xlcell_], xlrow)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of XLSheet
	case XLSheet:
		switch fieldname {
		// insertion point for per direct association field
		case "Rows":
			res := make(map[*XLRow][]*XLSheet)
			for xlsheet := range stage.XLSheets {
				for _, xlrow_ := range xlsheet.Rows {
					res[xlrow_] = append(res[xlrow_], xlsheet)
				}
			}
			return any(res).(map[*End][]*Start)
		case "SheetCells":
			res := make(map[*XLCell][]*XLSheet)
			for xlsheet := range stage.XLSheets {
				for _, xlcell_ := range xlsheet.SheetCells {
					res[xlcell_] = append(res[xlcell_], xlsheet)
				}
			}
			return any(res).(map[*End][]*Start)
		}
	}
	return nil
}

// GongNewInstance creates a new instance of the Gongstruct
func GongNewInstance[Type GongstructPtr]() (res Type) {
	var t Type
	return t.GongNewInstance().(Type)
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
	var t Type
	return t.GongGetGongstructName()
}

func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	return GongGetPointerToGongstructName[Type]()
}

type GongReverseField struct {
	GongstructName string
	Fieldname      string
}

type ReverseField = GongReverseField

// insertion point for generic get reverse fields
func (*DisplaySelection) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*XLCell) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "XLRow",
			Fieldname: "Cells",
		},
		{
			GongstructName: "XLSheet",
			Fieldname: "SheetCells",
		},
	}
}

func (*XLFile) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
	}
}

func (*XLRow) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "XLSheet",
			Fieldname: "Rows",
		},
	}
}

func (*XLSheet) GongGetReverseFields() []GongReverseField {
	return []GongReverseField{ 
		{
			GongstructName: "XLFile",
			Fieldname: "Sheets",
		},
	}
}


func GongGetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	var t Type
	return t.GongGetReverseFields()
}

func GetReverseFields[Type GongstructIF]() (res []GongReverseField) {
	return GongGetReverseFields[Type]()
}

// insertion point for get fields header method
func (displayselection *DisplaySelection) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:                 "XLFile",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "XLFile",
		},
		{
			Name:                 "XLSheet",
			GongFieldValueType:   GongFieldValueTypePointer,
			TargetGongstructName: "XLSheet",
		},
	}
	return
}

func (xlcell *XLCell) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeInt,
		},
	}
	return
}

func (xlfile *XLFile) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "NbSheets",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Sheets",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "XLSheet",
		},
	}
	return
}

func (xlrow *XLRow) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "RowIndex",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Cells",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "XLCell",
		},
	}
	return
}

func (xlsheet *XLSheet) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeString,
		},
		{
			Name:               "MaxRow",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "MaxCol",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:               "NbRows",
			GongFieldValueType: GongFieldValueTypeInt,
		},
		{
			Name:                 "Rows",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "XLRow",
		},
		{
			Name:                 "SheetCells",
			GongFieldValueType:   GongFieldValueTypeSliceOfPointers,
			TargetGongstructName: "XLCell",
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
func (displayselection *DisplaySelection) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = displayselection.Name
	case "XLFile":
		res.GongFieldValueType = GongFieldValueTypePointer
		if displayselection.XLFile != nil {
			res.valueString = displayselection.XLFile.Name
			res.ids = displayselection.XLFile.GongGetUUID(stage)
		}
	case "XLSheet":
		res.GongFieldValueType = GongFieldValueTypePointer
		if displayselection.XLSheet != nil {
			res.valueString = displayselection.XLSheet.Name
			res.ids = displayselection.XLSheet.GongGetUUID(stage)
		}
	}
	return
}

func (xlcell *XLCell) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = xlcell.Name
	case "X":
		res.valueString = fmt.Sprintf("%d", xlcell.X)
		res.valueInt = xlcell.X
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Y":
		res.valueString = fmt.Sprintf("%d", xlcell.Y)
		res.valueInt = xlcell.Y
		res.GongFieldValueType = GongFieldValueTypeInt
	}
	return
}

func (xlfile *XLFile) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = xlfile.Name
	case "NbSheets":
		res.valueString = fmt.Sprintf("%d", xlfile.NbSheets)
		res.valueInt = xlfile.NbSheets
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Sheets":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range xlfile.Sheets {
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

func (xlrow *XLRow) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = xlrow.Name
	case "RowIndex":
		res.valueString = fmt.Sprintf("%d", xlrow.RowIndex)
		res.valueInt = xlrow.RowIndex
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Cells":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range xlrow.Cells {
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

func (xlsheet *XLSheet) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = xlsheet.Name
	case "MaxRow":
		res.valueString = fmt.Sprintf("%d", xlsheet.MaxRow)
		res.valueInt = xlsheet.MaxRow
		res.GongFieldValueType = GongFieldValueTypeInt
	case "MaxCol":
		res.valueString = fmt.Sprintf("%d", xlsheet.MaxCol)
		res.valueInt = xlsheet.MaxCol
		res.GongFieldValueType = GongFieldValueTypeInt
	case "NbRows":
		res.valueString = fmt.Sprintf("%d", xlsheet.NbRows)
		res.valueInt = xlsheet.NbRows
		res.GongFieldValueType = GongFieldValueTypeInt
	case "Rows":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range xlsheet.Rows {
			if idx > 0 {
				res.valueString += "\n"
				res.ids += ";"
			}
			res.valueString += __instance__.Name
			res.ids += __instance__.GongGetUUID(stage)
		}
	case "SheetCells":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range xlsheet.SheetCells {
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
func (displayselection *DisplaySelection) GongGetGongstructName() string {
	return "DisplaySelection"
}

func (xlcell *XLCell) GongGetGongstructName() string {
	return "XLCell"
}

func (xlfile *XLFile) GongGetGongstructName() string {
	return "XLFile"
}

func (xlrow *XLRow) GongGetGongstructName() string {
	return "XLRow"
}

func (xlsheet *XLSheet) GongGetGongstructName() string {
	return "XLSheet"
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
	__gong__rebuildMapString(stage.DisplaySelections, &stage.DisplaySelections_mapString)

	__gong__rebuildMapString(stage.XLCells, &stage.XLCells_mapString)

	__gong__rebuildMapString(stage.XLFiles, &stage.XLFiles_mapString)

	__gong__rebuildMapString(stage.XLRows, &stage.XLRows_mapString)

	__gong__rebuildMapString(stage.XLSheets, &stage.XLSheets_mapString)

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
