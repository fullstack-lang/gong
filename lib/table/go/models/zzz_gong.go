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
	"sync"
	"time"

	table_go "github.com/fullstack-lang/gong/lib/table/go"
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

var (
	_ = __Gong__Abs
	_ = strings.Clone("")
)

const (
	ProbeTreeSidebarSuffix           = ":sidebar of the probe"
	ProbeNavigationTreeSidebarSuffix = ":sidebar of the probe, navigation"
	ProbeTableSuffix                 = ":table of the probe"
	ProbeNotificationTableSuffix     = ":notification table of the probe"
	ProbeFormSuffix                  = ":form of the probe"
	ProbeSplitSuffix                 = ":probe of the probe"
	ProbeLoadSuffix                  = ":load of the probe"
)

type GongMarshallingMode string

const (
	// the whole stage is generated at each marshall. This is the default
	GongMarshallingNormal GongMarshallingMode = "GongMarshallingNormal"

	// only the last commit is append to the marshall file
	GongMarshallingAppendCommit GongMarshallingMode = "GongMarshallingAppendCommit"
)

func (stage *Stage) GetProbeTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeTreeSidebarSuffix
}

func (stage *Stage) GetProbeNavigationTreeSidebarStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeNavigationTreeSidebarSuffix
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

func (stage *Stage) GetProbeLoadStageName() string {
	return stage.GetType() + ":" + stage.GetName() + ProbeLoadSuffix
}

// errUnkownEnum is returns when a value cannot match enum values
var (
	errUnkownEnum = errors.New("unkown enum")
	_             = errUnkownEnum
)

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

var _ = __dummy__fmt_variable

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

var _ = __dummy_math_variable

// swagger:ignore
type __void any

// needed for creating set of instances in the stage
var (
	__member __void
	_        = __member
)

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
	OnAfterButtonCreateCallback OnAfterCreateInterface[Button]
	OnAfterButtonUpdateCallback OnAfterUpdateInterface[Button]
	OnAfterButtonDeleteCallback OnAfterDeleteInterface[Button]
	OnAfterButtonReadCallback   OnAfterReadInterface[Button]

	Cells                map[*Cell]struct{}
	Cells_instance       map[*Cell]*Cell
	Cells_mapString      map[string]*Cell
	CellOrder            uint
	Cell_stagedOrder     map[*Cell]uint
	Cell_orderStaged     map[uint]*Cell
	Cells_reference      map[*Cell]*Cell
	Cells_referenceOrder map[*Cell]uint

	// insertion point for slice of pointers maps
	OnAfterCellCreateCallback OnAfterCreateInterface[Cell]
	OnAfterCellUpdateCallback OnAfterUpdateInterface[Cell]
	OnAfterCellDeleteCallback OnAfterDeleteInterface[Cell]
	OnAfterCellReadCallback   OnAfterReadInterface[Cell]

	CellBooleans                map[*CellBoolean]struct{}
	CellBooleans_instance       map[*CellBoolean]*CellBoolean
	CellBooleans_mapString      map[string]*CellBoolean
	CellBooleanOrder            uint
	CellBoolean_stagedOrder     map[*CellBoolean]uint
	CellBoolean_orderStaged     map[uint]*CellBoolean
	CellBooleans_reference      map[*CellBoolean]*CellBoolean
	CellBooleans_referenceOrder map[*CellBoolean]uint

	// insertion point for slice of pointers maps
	OnAfterCellBooleanCreateCallback OnAfterCreateInterface[CellBoolean]
	OnAfterCellBooleanUpdateCallback OnAfterUpdateInterface[CellBoolean]
	OnAfterCellBooleanDeleteCallback OnAfterDeleteInterface[CellBoolean]
	OnAfterCellBooleanReadCallback   OnAfterReadInterface[CellBoolean]

	CellFloat64s                map[*CellFloat64]struct{}
	CellFloat64s_instance       map[*CellFloat64]*CellFloat64
	CellFloat64s_mapString      map[string]*CellFloat64
	CellFloat64Order            uint
	CellFloat64_stagedOrder     map[*CellFloat64]uint
	CellFloat64_orderStaged     map[uint]*CellFloat64
	CellFloat64s_reference      map[*CellFloat64]*CellFloat64
	CellFloat64s_referenceOrder map[*CellFloat64]uint

	// insertion point for slice of pointers maps
	OnAfterCellFloat64CreateCallback OnAfterCreateInterface[CellFloat64]
	OnAfterCellFloat64UpdateCallback OnAfterUpdateInterface[CellFloat64]
	OnAfterCellFloat64DeleteCallback OnAfterDeleteInterface[CellFloat64]
	OnAfterCellFloat64ReadCallback   OnAfterReadInterface[CellFloat64]

	CellIcons                map[*CellIcon]struct{}
	CellIcons_instance       map[*CellIcon]*CellIcon
	CellIcons_mapString      map[string]*CellIcon
	CellIconOrder            uint
	CellIcon_stagedOrder     map[*CellIcon]uint
	CellIcon_orderStaged     map[uint]*CellIcon
	CellIcons_reference      map[*CellIcon]*CellIcon
	CellIcons_referenceOrder map[*CellIcon]uint

	// insertion point for slice of pointers maps
	OnAfterCellIconCreateCallback OnAfterCreateInterface[CellIcon]
	OnAfterCellIconUpdateCallback OnAfterUpdateInterface[CellIcon]
	OnAfterCellIconDeleteCallback OnAfterDeleteInterface[CellIcon]
	OnAfterCellIconReadCallback   OnAfterReadInterface[CellIcon]

	CellInts                map[*CellInt]struct{}
	CellInts_instance       map[*CellInt]*CellInt
	CellInts_mapString      map[string]*CellInt
	CellIntOrder            uint
	CellInt_stagedOrder     map[*CellInt]uint
	CellInt_orderStaged     map[uint]*CellInt
	CellInts_reference      map[*CellInt]*CellInt
	CellInts_referenceOrder map[*CellInt]uint

	// insertion point for slice of pointers maps
	OnAfterCellIntCreateCallback OnAfterCreateInterface[CellInt]
	OnAfterCellIntUpdateCallback OnAfterUpdateInterface[CellInt]
	OnAfterCellIntDeleteCallback OnAfterDeleteInterface[CellInt]
	OnAfterCellIntReadCallback   OnAfterReadInterface[CellInt]

	CellStrings                map[*CellString]struct{}
	CellStrings_instance       map[*CellString]*CellString
	CellStrings_mapString      map[string]*CellString
	CellStringOrder            uint
	CellString_stagedOrder     map[*CellString]uint
	CellString_orderStaged     map[uint]*CellString
	CellStrings_reference      map[*CellString]*CellString
	CellStrings_referenceOrder map[*CellString]uint

	// insertion point for slice of pointers maps
	OnAfterCellStringCreateCallback OnAfterCreateInterface[CellString]
	OnAfterCellStringUpdateCallback OnAfterUpdateInterface[CellString]
	OnAfterCellStringDeleteCallback OnAfterDeleteInterface[CellString]
	OnAfterCellStringReadCallback   OnAfterReadInterface[CellString]

	DisplayedColumns                map[*DisplayedColumn]struct{}
	DisplayedColumns_instance       map[*DisplayedColumn]*DisplayedColumn
	DisplayedColumns_mapString      map[string]*DisplayedColumn
	DisplayedColumnOrder            uint
	DisplayedColumn_stagedOrder     map[*DisplayedColumn]uint
	DisplayedColumn_orderStaged     map[uint]*DisplayedColumn
	DisplayedColumns_reference      map[*DisplayedColumn]*DisplayedColumn
	DisplayedColumns_referenceOrder map[*DisplayedColumn]uint

	// insertion point for slice of pointers maps
	OnAfterDisplayedColumnCreateCallback OnAfterCreateInterface[DisplayedColumn]
	OnAfterDisplayedColumnUpdateCallback OnAfterUpdateInterface[DisplayedColumn]
	OnAfterDisplayedColumnDeleteCallback OnAfterDeleteInterface[DisplayedColumn]
	OnAfterDisplayedColumnReadCallback   OnAfterReadInterface[DisplayedColumn]

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

	OnAfterRowCreateCallback OnAfterCreateInterface[Row]
	OnAfterRowUpdateCallback OnAfterUpdateInterface[Row]
	OnAfterRowDeleteCallback OnAfterDeleteInterface[Row]
	OnAfterRowReadCallback   OnAfterReadInterface[Row]

	SVGIcons                map[*SVGIcon]struct{}
	SVGIcons_instance       map[*SVGIcon]*SVGIcon
	SVGIcons_mapString      map[string]*SVGIcon
	SVGIconOrder            uint
	SVGIcon_stagedOrder     map[*SVGIcon]uint
	SVGIcon_orderStaged     map[uint]*SVGIcon
	SVGIcons_reference      map[*SVGIcon]*SVGIcon
	SVGIcons_referenceOrder map[*SVGIcon]uint

	// insertion point for slice of pointers maps
	OnAfterSVGIconCreateCallback OnAfterCreateInterface[SVGIcon]
	OnAfterSVGIconUpdateCallback OnAfterUpdateInterface[SVGIcon]
	OnAfterSVGIconDeleteCallback OnAfterDeleteInterface[SVGIcon]
	OnAfterSVGIconReadCallback   OnAfterReadInterface[SVGIcon]

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

	OnAfterTableCreateCallback OnAfterCreateInterface[Table]
	OnAfterTableUpdateCallback OnAfterUpdateInterface[Table]
	OnAfterTableDeleteCallback OnAfterDeleteInterface[Table]
	OnAfterTableReadCallback   OnAfterReadInterface[Table]

	AllModelsStructCreateCallback AllModelsStructCreateInterface

	AllModelsStructDeleteCallback AllModelsStructDeleteInterface

	BackRepo BackRepoInterface

	// if set will be called before each commit to the back repo
	OnInitCommitCallback          OnInitCommitInterface
	OnInitCommitFromFrontCallback OnInitCommitInterface
	OnInitCommitFromBackCallback  OnInitCommitInterface

	// Private slices to hold the registered hooks
	beforeCommitHooks []func(stage *Stage)
	afterCommitHooks  []func(stage *Stage)

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

	isApplyingBackwardCommit bool
	isApplyingForwardCommit  bool
	isSquashing              bool

	modified bool

	lock sync.RWMutex
}

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
	err := GongParseAstString(stage, commitToApply, true)
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
	err := GongParseAstString(stage, commitToApply, true)
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
	stage.Buttons_reference = make(map[*Button]*Button)
	stage.Buttons_instance = make(map[*Button]*Button)
	stage.Buttons_referenceOrder = make(map[*Button]uint)

	stage.Cells_reference = make(map[*Cell]*Cell)
	stage.Cells_instance = make(map[*Cell]*Cell)
	stage.Cells_referenceOrder = make(map[*Cell]uint)

	stage.CellBooleans_reference = make(map[*CellBoolean]*CellBoolean)
	stage.CellBooleans_instance = make(map[*CellBoolean]*CellBoolean)
	stage.CellBooleans_referenceOrder = make(map[*CellBoolean]uint)

	stage.CellFloat64s_reference = make(map[*CellFloat64]*CellFloat64)
	stage.CellFloat64s_instance = make(map[*CellFloat64]*CellFloat64)
	stage.CellFloat64s_referenceOrder = make(map[*CellFloat64]uint)

	stage.CellIcons_reference = make(map[*CellIcon]*CellIcon)
	stage.CellIcons_instance = make(map[*CellIcon]*CellIcon)
	stage.CellIcons_referenceOrder = make(map[*CellIcon]uint)

	stage.CellInts_reference = make(map[*CellInt]*CellInt)
	stage.CellInts_instance = make(map[*CellInt]*CellInt)
	stage.CellInts_referenceOrder = make(map[*CellInt]uint)

	stage.CellStrings_reference = make(map[*CellString]*CellString)
	stage.CellStrings_instance = make(map[*CellString]*CellString)
	stage.CellStrings_referenceOrder = make(map[*CellString]uint)

	stage.DisplayedColumns_reference = make(map[*DisplayedColumn]*DisplayedColumn)
	stage.DisplayedColumns_instance = make(map[*DisplayedColumn]*DisplayedColumn)
	stage.DisplayedColumns_referenceOrder = make(map[*DisplayedColumn]uint)

	stage.Rows_reference = make(map[*Row]*Row)
	stage.Rows_instance = make(map[*Row]*Row)
	stage.Rows_referenceOrder = make(map[*Row]uint)

	stage.SVGIcons_reference = make(map[*SVGIcon]*SVGIcon)
	stage.SVGIcons_instance = make(map[*SVGIcon]*SVGIcon)
	stage.SVGIcons_referenceOrder = make(map[*SVGIcon]uint)

	stage.Tables_reference = make(map[*Table]*Table)
	stage.Tables_instance = make(map[*Table]*Table)
	stage.Tables_referenceOrder = make(map[*Table]uint)

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
	var maxButtonOrder uint
	var foundButton bool
	for _, order := range stage.Button_stagedOrder {
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

	var maxCellOrder uint
	var foundCell bool
	for _, order := range stage.Cell_stagedOrder {
		if !foundCell || order > maxCellOrder {
			maxCellOrder = order
			foundCell = true
		}
	}
	if foundCell {
		stage.CellOrder = maxCellOrder + 1
	} else {
		stage.CellOrder = 0
	}

	var maxCellBooleanOrder uint
	var foundCellBoolean bool
	for _, order := range stage.CellBoolean_stagedOrder {
		if !foundCellBoolean || order > maxCellBooleanOrder {
			maxCellBooleanOrder = order
			foundCellBoolean = true
		}
	}
	if foundCellBoolean {
		stage.CellBooleanOrder = maxCellBooleanOrder + 1
	} else {
		stage.CellBooleanOrder = 0
	}

	var maxCellFloat64Order uint
	var foundCellFloat64 bool
	for _, order := range stage.CellFloat64_stagedOrder {
		if !foundCellFloat64 || order > maxCellFloat64Order {
			maxCellFloat64Order = order
			foundCellFloat64 = true
		}
	}
	if foundCellFloat64 {
		stage.CellFloat64Order = maxCellFloat64Order + 1
	} else {
		stage.CellFloat64Order = 0
	}

	var maxCellIconOrder uint
	var foundCellIcon bool
	for _, order := range stage.CellIcon_stagedOrder {
		if !foundCellIcon || order > maxCellIconOrder {
			maxCellIconOrder = order
			foundCellIcon = true
		}
	}
	if foundCellIcon {
		stage.CellIconOrder = maxCellIconOrder + 1
	} else {
		stage.CellIconOrder = 0
	}

	var maxCellIntOrder uint
	var foundCellInt bool
	for _, order := range stage.CellInt_stagedOrder {
		if !foundCellInt || order > maxCellIntOrder {
			maxCellIntOrder = order
			foundCellInt = true
		}
	}
	if foundCellInt {
		stage.CellIntOrder = maxCellIntOrder + 1
	} else {
		stage.CellIntOrder = 0
	}

	var maxCellStringOrder uint
	var foundCellString bool
	for _, order := range stage.CellString_stagedOrder {
		if !foundCellString || order > maxCellStringOrder {
			maxCellStringOrder = order
			foundCellString = true
		}
	}
	if foundCellString {
		stage.CellStringOrder = maxCellStringOrder + 1
	} else {
		stage.CellStringOrder = 0
	}

	var maxDisplayedColumnOrder uint
	var foundDisplayedColumn bool
	for _, order := range stage.DisplayedColumn_stagedOrder {
		if !foundDisplayedColumn || order > maxDisplayedColumnOrder {
			maxDisplayedColumnOrder = order
			foundDisplayedColumn = true
		}
	}
	if foundDisplayedColumn {
		stage.DisplayedColumnOrder = maxDisplayedColumnOrder + 1
	} else {
		stage.DisplayedColumnOrder = 0
	}

	var maxRowOrder uint
	var foundRow bool
	for _, order := range stage.Row_stagedOrder {
		if !foundRow || order > maxRowOrder {
			maxRowOrder = order
			foundRow = true
		}
	}
	if foundRow {
		stage.RowOrder = maxRowOrder + 1
	} else {
		stage.RowOrder = 0
	}

	var maxSVGIconOrder uint
	var foundSVGIcon bool
	for _, order := range stage.SVGIcon_stagedOrder {
		if !foundSVGIcon || order > maxSVGIconOrder {
			maxSVGIconOrder = order
			foundSVGIcon = true
		}
	}
	if foundSVGIcon {
		stage.SVGIconOrder = maxSVGIconOrder + 1
	} else {
		stage.SVGIconOrder = 0
	}

	var maxTableOrder uint
	var foundTable bool
	for _, order := range stage.Table_stagedOrder {
		if !foundTable || order > maxTableOrder {
			maxTableOrder = order
			foundTable = true
		}
	}
	if foundTable {
		stage.TableOrder = maxTableOrder + 1
	} else {
		stage.TableOrder = 0
	}

	// end of insertion point for max order recomputation
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

// GetStructInstancesByOrderAuto returns a slice of generic pointers to gongstructs
// ordered by their order in the stage.
func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *Button:
		tmp := GetStructInstancesByOrder(stage.Buttons, stage.Button_stagedOrder)

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
	case *Cell:
		tmp := GetStructInstancesByOrder(stage.Cells, stage.Cell_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Cell implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CellBoolean:
		tmp := GetStructInstancesByOrder(stage.CellBooleans, stage.CellBoolean_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CellBoolean implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CellFloat64:
		tmp := GetStructInstancesByOrder(stage.CellFloat64s, stage.CellFloat64_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CellFloat64 implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CellIcon:
		tmp := GetStructInstancesByOrder(stage.CellIcons, stage.CellIcon_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CellIcon implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CellInt:
		tmp := GetStructInstancesByOrder(stage.CellInts, stage.CellInt_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CellInt implements.
			res = append(res, any(v).(T))
		}
		return res
	case *CellString:
		tmp := GetStructInstancesByOrder(stage.CellStrings, stage.CellString_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *CellString implements.
			res = append(res, any(v).(T))
		}
		return res
	case *DisplayedColumn:
		tmp := GetStructInstancesByOrder(stage.DisplayedColumns, stage.DisplayedColumn_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DisplayedColumn implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Row:
		tmp := GetStructInstancesByOrder(stage.Rows, stage.Row_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Row implements.
			res = append(res, any(v).(T))
		}
		return res
	case *SVGIcon:
		tmp := GetStructInstancesByOrder(stage.SVGIcons, stage.SVGIcon_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *SVGIcon implements.
			res = append(res, any(v).(T))
		}
		return res
	case *Table:
		tmp := GetStructInstancesByOrder(stage.Tables, stage.Table_stagedOrder)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *Table implements.
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
		res = GetNamedStructInstances(stage.Buttons, stage.Button_stagedOrder)
	case "Cell":
		res = GetNamedStructInstances(stage.Cells, stage.Cell_stagedOrder)
	case "CellBoolean":
		res = GetNamedStructInstances(stage.CellBooleans, stage.CellBoolean_stagedOrder)
	case "CellFloat64":
		res = GetNamedStructInstances(stage.CellFloat64s, stage.CellFloat64_stagedOrder)
	case "CellIcon":
		res = GetNamedStructInstances(stage.CellIcons, stage.CellIcon_stagedOrder)
	case "CellInt":
		res = GetNamedStructInstances(stage.CellInts, stage.CellInt_stagedOrder)
	case "CellString":
		res = GetNamedStructInstances(stage.CellStrings, stage.CellString_stagedOrder)
	case "DisplayedColumn":
		res = GetNamedStructInstances(stage.DisplayedColumns, stage.DisplayedColumn_stagedOrder)
	case "Row":
		res = GetNamedStructInstances(stage.Rows, stage.Row_stagedOrder)
	case "SVGIcon":
		res = GetNamedStructInstances(stage.SVGIcons, stage.SVGIcon_stagedOrder)
	case "Table":
		res = GetNamedStructInstances(stage.Tables, stage.Table_stagedOrder)
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
	return "github.com/fullstack-lang/gong/lib/table/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return table_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return table_go.GoDiagramsDir
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
	CommitCell(cell *Cell)
	CheckoutCell(cell *Cell)
	CommitCellBoolean(cellboolean *CellBoolean)
	CheckoutCellBoolean(cellboolean *CellBoolean)
	CommitCellFloat64(cellfloat64 *CellFloat64)
	CheckoutCellFloat64(cellfloat64 *CellFloat64)
	CommitCellIcon(cellicon *CellIcon)
	CheckoutCellIcon(cellicon *CellIcon)
	CommitCellInt(cellint *CellInt)
	CheckoutCellInt(cellint *CellInt)
	CommitCellString(cellstring *CellString)
	CheckoutCellString(cellstring *CellString)
	CommitDisplayedColumn(displayedcolumn *DisplayedColumn)
	CheckoutDisplayedColumn(displayedcolumn *DisplayedColumn)
	CommitRow(row *Row)
	CheckoutRow(row *Row)
	CommitSVGIcon(svgicon *SVGIcon)
	CheckoutSVGIcon(svgicon *SVGIcon)
	CommitTable(table *Table)
	CheckoutTable(table *Table)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "Button"},
			{name: "Cell"},
			{name: "CellBoolean"},
			{name: "CellFloat64"},
			{name: "CellIcon"},
			{name: "CellInt"},
			{name: "CellString"},
			{name: "DisplayedColumn"},
			{name: "Row"},
			{name: "SVGIcon"},
			{name: "Table"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Button:
		return stage.Button_stagedOrder[instance]
	case *Cell:
		return stage.Cell_stagedOrder[instance]
	case *CellBoolean:
		return stage.CellBoolean_stagedOrder[instance]
	case *CellFloat64:
		return stage.CellFloat64_stagedOrder[instance]
	case *CellIcon:
		return stage.CellIcon_stagedOrder[instance]
	case *CellInt:
		return stage.CellInt_stagedOrder[instance]
	case *CellString:
		return stage.CellString_stagedOrder[instance]
	case *DisplayedColumn:
		return stage.DisplayedColumn_stagedOrder[instance]
	case *Row:
		return stage.Row_stagedOrder[instance]
	case *SVGIcon:
		return stage.SVGIcon_stagedOrder[instance]
	case *Table:
		return stage.Table_stagedOrder[instance]
	default:
		return 0 // should not happen
	}
}

func GongGetInstanceFromOrder[Type PointerToGongstruct](stage *Stage, order uint) (res Type) {
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

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *Button:
		return stage.Button_stagedOrder[instance]
	case *Cell:
		return stage.Cell_stagedOrder[instance]
	case *CellBoolean:
		return stage.CellBoolean_stagedOrder[instance]
	case *CellFloat64:
		return stage.CellFloat64_stagedOrder[instance]
	case *CellIcon:
		return stage.CellIcon_stagedOrder[instance]
	case *CellInt:
		return stage.CellInt_stagedOrder[instance]
	case *CellString:
		return stage.CellString_stagedOrder[instance]
	case *DisplayedColumn:
		return stage.DisplayedColumn_stagedOrder[instance]
	case *Row:
		return stage.Row_stagedOrder[instance]
	case *SVGIcon:
		return stage.SVGIcon_stagedOrder[instance]
	case *Table:
		return stage.Table_stagedOrder[instance]
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
	if _, ok := stage.Buttons[button]; !ok {
		stage.Buttons[button] = struct{}{}
		stage.Button_stagedOrder[button] = stage.ButtonOrder
		stage.Button_orderStaged[stage.ButtonOrder] = button
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
		stage.Button_stagedOrder[button] = order
		stage.Button_orderStaged[order] = button
		stage.ButtonOrder++
	}
	stage.Buttons_mapString[button.Name] = button
}

// Unstage removes button off the model stage
func (button *Button) Unstage(stage *Stage) *Button {
	delete(stage.Buttons, button)
	// issue1150
	// delete(stage.Button_stagedOrder, button)
	delete(stage.Buttons_mapString, button.Name)

	return button
}

// UnstageVoid removes button off the model stage
func (button *Button) UnstageVoid(stage *Stage) {
	delete(stage.Buttons, button)
	// issue1150
	// delete(stage.Button_stagedOrder, button)
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

// Stage puts cell to the model stage
func (cell *Cell) Stage(stage *Stage) *Cell {
	if _, ok := stage.Cells[cell]; !ok {
		stage.Cells[cell] = struct{}{}
		stage.Cell_stagedOrder[cell] = stage.CellOrder
		stage.Cell_orderStaged[stage.CellOrder] = cell
		stage.CellOrder++
	}
	stage.Cells_mapString[cell.Name] = cell

	return cell
}

// StagePreserveOrder puts cell to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellOrder
// - update stage.CellOrder accordingly
func (cell *Cell) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Cells[cell]; !ok {
		stage.Cells[cell] = struct{}{}

		if order > stage.CellOrder {
			stage.CellOrder = order
		}
		stage.Cell_stagedOrder[cell] = order
		stage.Cell_orderStaged[order] = cell
		stage.CellOrder++
	}
	stage.Cells_mapString[cell.Name] = cell
}

// Unstage removes cell off the model stage
func (cell *Cell) Unstage(stage *Stage) *Cell {
	delete(stage.Cells, cell)
	// issue1150
	// delete(stage.Cell_stagedOrder, cell)
	delete(stage.Cells_mapString, cell.Name)

	return cell
}

// UnstageVoid removes cell off the model stage
func (cell *Cell) UnstageVoid(stage *Stage) {
	delete(stage.Cells, cell)
	// issue1150
	// delete(stage.Cell_stagedOrder, cell)
	delete(stage.Cells_mapString, cell.Name)
}

// commit cell to the back repo (if it is already staged)
func (cell *Cell) Commit(stage *Stage) *Cell {
	if _, ok := stage.Cells[cell]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCell(cell)
		}
	}
	return cell
}

func (cell *Cell) CommitVoid(stage *Stage) {
	cell.Commit(stage)
}

func (cell *Cell) StageVoid(stage *Stage) {
	cell.Stage(stage)
}

// Checkout cell to the back repo (if it is already staged)
func (cell *Cell) Checkout(stage *Stage) *Cell {
	if _, ok := stage.Cells[cell]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCell(cell)
		}
	}
	return cell
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
	if _, ok := stage.CellBooleans[cellboolean]; !ok {
		stage.CellBooleans[cellboolean] = struct{}{}
		stage.CellBoolean_stagedOrder[cellboolean] = stage.CellBooleanOrder
		stage.CellBoolean_orderStaged[stage.CellBooleanOrder] = cellboolean
		stage.CellBooleanOrder++
	}
	stage.CellBooleans_mapString[cellboolean.Name] = cellboolean

	return cellboolean
}

// StagePreserveOrder puts cellboolean to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellBooleanOrder
// - update stage.CellBooleanOrder accordingly
func (cellboolean *CellBoolean) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CellBooleans[cellboolean]; !ok {
		stage.CellBooleans[cellboolean] = struct{}{}

		if order > stage.CellBooleanOrder {
			stage.CellBooleanOrder = order
		}
		stage.CellBoolean_stagedOrder[cellboolean] = order
		stage.CellBoolean_orderStaged[order] = cellboolean
		stage.CellBooleanOrder++
	}
	stage.CellBooleans_mapString[cellboolean.Name] = cellboolean
}

// Unstage removes cellboolean off the model stage
func (cellboolean *CellBoolean) Unstage(stage *Stage) *CellBoolean {
	delete(stage.CellBooleans, cellboolean)
	// issue1150
	// delete(stage.CellBoolean_stagedOrder, cellboolean)
	delete(stage.CellBooleans_mapString, cellboolean.Name)

	return cellboolean
}

// UnstageVoid removes cellboolean off the model stage
func (cellboolean *CellBoolean) UnstageVoid(stage *Stage) {
	delete(stage.CellBooleans, cellboolean)
	// issue1150
	// delete(stage.CellBoolean_stagedOrder, cellboolean)
	delete(stage.CellBooleans_mapString, cellboolean.Name)
}

// commit cellboolean to the back repo (if it is already staged)
func (cellboolean *CellBoolean) Commit(stage *Stage) *CellBoolean {
	if _, ok := stage.CellBooleans[cellboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellBoolean(cellboolean)
		}
	}
	return cellboolean
}

func (cellboolean *CellBoolean) CommitVoid(stage *Stage) {
	cellboolean.Commit(stage)
}

func (cellboolean *CellBoolean) StageVoid(stage *Stage) {
	cellboolean.Stage(stage)
}

// Checkout cellboolean to the back repo (if it is already staged)
func (cellboolean *CellBoolean) Checkout(stage *Stage) *CellBoolean {
	if _, ok := stage.CellBooleans[cellboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCellBoolean(cellboolean)
		}
	}
	return cellboolean
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
	if _, ok := stage.CellFloat64s[cellfloat64]; !ok {
		stage.CellFloat64s[cellfloat64] = struct{}{}
		stage.CellFloat64_stagedOrder[cellfloat64] = stage.CellFloat64Order
		stage.CellFloat64_orderStaged[stage.CellFloat64Order] = cellfloat64
		stage.CellFloat64Order++
	}
	stage.CellFloat64s_mapString[cellfloat64.Name] = cellfloat64

	return cellfloat64
}

// StagePreserveOrder puts cellfloat64 to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellFloat64Order
// - update stage.CellFloat64Order accordingly
func (cellfloat64 *CellFloat64) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CellFloat64s[cellfloat64]; !ok {
		stage.CellFloat64s[cellfloat64] = struct{}{}

		if order > stage.CellFloat64Order {
			stage.CellFloat64Order = order
		}
		stage.CellFloat64_stagedOrder[cellfloat64] = order
		stage.CellFloat64_orderStaged[order] = cellfloat64
		stage.CellFloat64Order++
	}
	stage.CellFloat64s_mapString[cellfloat64.Name] = cellfloat64
}

// Unstage removes cellfloat64 off the model stage
func (cellfloat64 *CellFloat64) Unstage(stage *Stage) *CellFloat64 {
	delete(stage.CellFloat64s, cellfloat64)
	// issue1150
	// delete(stage.CellFloat64_stagedOrder, cellfloat64)
	delete(stage.CellFloat64s_mapString, cellfloat64.Name)

	return cellfloat64
}

// UnstageVoid removes cellfloat64 off the model stage
func (cellfloat64 *CellFloat64) UnstageVoid(stage *Stage) {
	delete(stage.CellFloat64s, cellfloat64)
	// issue1150
	// delete(stage.CellFloat64_stagedOrder, cellfloat64)
	delete(stage.CellFloat64s_mapString, cellfloat64.Name)
}

// commit cellfloat64 to the back repo (if it is already staged)
func (cellfloat64 *CellFloat64) Commit(stage *Stage) *CellFloat64 {
	if _, ok := stage.CellFloat64s[cellfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellFloat64(cellfloat64)
		}
	}
	return cellfloat64
}

func (cellfloat64 *CellFloat64) CommitVoid(stage *Stage) {
	cellfloat64.Commit(stage)
}

func (cellfloat64 *CellFloat64) StageVoid(stage *Stage) {
	cellfloat64.Stage(stage)
}

// Checkout cellfloat64 to the back repo (if it is already staged)
func (cellfloat64 *CellFloat64) Checkout(stage *Stage) *CellFloat64 {
	if _, ok := stage.CellFloat64s[cellfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCellFloat64(cellfloat64)
		}
	}
	return cellfloat64
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
	if _, ok := stage.CellIcons[cellicon]; !ok {
		stage.CellIcons[cellicon] = struct{}{}
		stage.CellIcon_stagedOrder[cellicon] = stage.CellIconOrder
		stage.CellIcon_orderStaged[stage.CellIconOrder] = cellicon
		stage.CellIconOrder++
	}
	stage.CellIcons_mapString[cellicon.Name] = cellicon

	return cellicon
}

// StagePreserveOrder puts cellicon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellIconOrder
// - update stage.CellIconOrder accordingly
func (cellicon *CellIcon) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CellIcons[cellicon]; !ok {
		stage.CellIcons[cellicon] = struct{}{}

		if order > stage.CellIconOrder {
			stage.CellIconOrder = order
		}
		stage.CellIcon_stagedOrder[cellicon] = order
		stage.CellIcon_orderStaged[order] = cellicon
		stage.CellIconOrder++
	}
	stage.CellIcons_mapString[cellicon.Name] = cellicon
}

// Unstage removes cellicon off the model stage
func (cellicon *CellIcon) Unstage(stage *Stage) *CellIcon {
	delete(stage.CellIcons, cellicon)
	// issue1150
	// delete(stage.CellIcon_stagedOrder, cellicon)
	delete(stage.CellIcons_mapString, cellicon.Name)

	return cellicon
}

// UnstageVoid removes cellicon off the model stage
func (cellicon *CellIcon) UnstageVoid(stage *Stage) {
	delete(stage.CellIcons, cellicon)
	// issue1150
	// delete(stage.CellIcon_stagedOrder, cellicon)
	delete(stage.CellIcons_mapString, cellicon.Name)
}

// commit cellicon to the back repo (if it is already staged)
func (cellicon *CellIcon) Commit(stage *Stage) *CellIcon {
	if _, ok := stage.CellIcons[cellicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellIcon(cellicon)
		}
	}
	return cellicon
}

func (cellicon *CellIcon) CommitVoid(stage *Stage) {
	cellicon.Commit(stage)
}

func (cellicon *CellIcon) StageVoid(stage *Stage) {
	cellicon.Stage(stage)
}

// Checkout cellicon to the back repo (if it is already staged)
func (cellicon *CellIcon) Checkout(stage *Stage) *CellIcon {
	if _, ok := stage.CellIcons[cellicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCellIcon(cellicon)
		}
	}
	return cellicon
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
	if _, ok := stage.CellInts[cellint]; !ok {
		stage.CellInts[cellint] = struct{}{}
		stage.CellInt_stagedOrder[cellint] = stage.CellIntOrder
		stage.CellInt_orderStaged[stage.CellIntOrder] = cellint
		stage.CellIntOrder++
	}
	stage.CellInts_mapString[cellint.Name] = cellint

	return cellint
}

// StagePreserveOrder puts cellint to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellIntOrder
// - update stage.CellIntOrder accordingly
func (cellint *CellInt) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CellInts[cellint]; !ok {
		stage.CellInts[cellint] = struct{}{}

		if order > stage.CellIntOrder {
			stage.CellIntOrder = order
		}
		stage.CellInt_stagedOrder[cellint] = order
		stage.CellInt_orderStaged[order] = cellint
		stage.CellIntOrder++
	}
	stage.CellInts_mapString[cellint.Name] = cellint
}

// Unstage removes cellint off the model stage
func (cellint *CellInt) Unstage(stage *Stage) *CellInt {
	delete(stage.CellInts, cellint)
	// issue1150
	// delete(stage.CellInt_stagedOrder, cellint)
	delete(stage.CellInts_mapString, cellint.Name)

	return cellint
}

// UnstageVoid removes cellint off the model stage
func (cellint *CellInt) UnstageVoid(stage *Stage) {
	delete(stage.CellInts, cellint)
	// issue1150
	// delete(stage.CellInt_stagedOrder, cellint)
	delete(stage.CellInts_mapString, cellint.Name)
}

// commit cellint to the back repo (if it is already staged)
func (cellint *CellInt) Commit(stage *Stage) *CellInt {
	if _, ok := stage.CellInts[cellint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellInt(cellint)
		}
	}
	return cellint
}

func (cellint *CellInt) CommitVoid(stage *Stage) {
	cellint.Commit(stage)
}

func (cellint *CellInt) StageVoid(stage *Stage) {
	cellint.Stage(stage)
}

// Checkout cellint to the back repo (if it is already staged)
func (cellint *CellInt) Checkout(stage *Stage) *CellInt {
	if _, ok := stage.CellInts[cellint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCellInt(cellint)
		}
	}
	return cellint
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
	if _, ok := stage.CellStrings[cellstring]; !ok {
		stage.CellStrings[cellstring] = struct{}{}
		stage.CellString_stagedOrder[cellstring] = stage.CellStringOrder
		stage.CellString_orderStaged[stage.CellStringOrder] = cellstring
		stage.CellStringOrder++
	}
	stage.CellStrings_mapString[cellstring.Name] = cellstring

	return cellstring
}

// StagePreserveOrder puts cellstring to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.CellStringOrder
// - update stage.CellStringOrder accordingly
func (cellstring *CellString) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.CellStrings[cellstring]; !ok {
		stage.CellStrings[cellstring] = struct{}{}

		if order > stage.CellStringOrder {
			stage.CellStringOrder = order
		}
		stage.CellString_stagedOrder[cellstring] = order
		stage.CellString_orderStaged[order] = cellstring
		stage.CellStringOrder++
	}
	stage.CellStrings_mapString[cellstring.Name] = cellstring
}

// Unstage removes cellstring off the model stage
func (cellstring *CellString) Unstage(stage *Stage) *CellString {
	delete(stage.CellStrings, cellstring)
	// issue1150
	// delete(stage.CellString_stagedOrder, cellstring)
	delete(stage.CellStrings_mapString, cellstring.Name)

	return cellstring
}

// UnstageVoid removes cellstring off the model stage
func (cellstring *CellString) UnstageVoid(stage *Stage) {
	delete(stage.CellStrings, cellstring)
	// issue1150
	// delete(stage.CellString_stagedOrder, cellstring)
	delete(stage.CellStrings_mapString, cellstring.Name)
}

// commit cellstring to the back repo (if it is already staged)
func (cellstring *CellString) Commit(stage *Stage) *CellString {
	if _, ok := stage.CellStrings[cellstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellString(cellstring)
		}
	}
	return cellstring
}

func (cellstring *CellString) CommitVoid(stage *Stage) {
	cellstring.Commit(stage)
}

func (cellstring *CellString) StageVoid(stage *Stage) {
	cellstring.Stage(stage)
}

// Checkout cellstring to the back repo (if it is already staged)
func (cellstring *CellString) Checkout(stage *Stage) *CellString {
	if _, ok := stage.CellStrings[cellstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCellString(cellstring)
		}
	}
	return cellstring
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
	if _, ok := stage.DisplayedColumns[displayedcolumn]; !ok {
		stage.DisplayedColumns[displayedcolumn] = struct{}{}
		stage.DisplayedColumn_stagedOrder[displayedcolumn] = stage.DisplayedColumnOrder
		stage.DisplayedColumn_orderStaged[stage.DisplayedColumnOrder] = displayedcolumn
		stage.DisplayedColumnOrder++
	}
	stage.DisplayedColumns_mapString[displayedcolumn.Name] = displayedcolumn

	return displayedcolumn
}

// StagePreserveOrder puts displayedcolumn to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DisplayedColumnOrder
// - update stage.DisplayedColumnOrder accordingly
func (displayedcolumn *DisplayedColumn) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DisplayedColumns[displayedcolumn]; !ok {
		stage.DisplayedColumns[displayedcolumn] = struct{}{}

		if order > stage.DisplayedColumnOrder {
			stage.DisplayedColumnOrder = order
		}
		stage.DisplayedColumn_stagedOrder[displayedcolumn] = order
		stage.DisplayedColumn_orderStaged[order] = displayedcolumn
		stage.DisplayedColumnOrder++
	}
	stage.DisplayedColumns_mapString[displayedcolumn.Name] = displayedcolumn
}

// Unstage removes displayedcolumn off the model stage
func (displayedcolumn *DisplayedColumn) Unstage(stage *Stage) *DisplayedColumn {
	delete(stage.DisplayedColumns, displayedcolumn)
	// issue1150
	// delete(stage.DisplayedColumn_stagedOrder, displayedcolumn)
	delete(stage.DisplayedColumns_mapString, displayedcolumn.Name)

	return displayedcolumn
}

// UnstageVoid removes displayedcolumn off the model stage
func (displayedcolumn *DisplayedColumn) UnstageVoid(stage *Stage) {
	delete(stage.DisplayedColumns, displayedcolumn)
	// issue1150
	// delete(stage.DisplayedColumn_stagedOrder, displayedcolumn)
	delete(stage.DisplayedColumns_mapString, displayedcolumn.Name)
}

// commit displayedcolumn to the back repo (if it is already staged)
func (displayedcolumn *DisplayedColumn) Commit(stage *Stage) *DisplayedColumn {
	if _, ok := stage.DisplayedColumns[displayedcolumn]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDisplayedColumn(displayedcolumn)
		}
	}
	return displayedcolumn
}

func (displayedcolumn *DisplayedColumn) CommitVoid(stage *Stage) {
	displayedcolumn.Commit(stage)
}

func (displayedcolumn *DisplayedColumn) StageVoid(stage *Stage) {
	displayedcolumn.Stage(stage)
}

// Checkout displayedcolumn to the back repo (if it is already staged)
func (displayedcolumn *DisplayedColumn) Checkout(stage *Stage) *DisplayedColumn {
	if _, ok := stage.DisplayedColumns[displayedcolumn]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDisplayedColumn(displayedcolumn)
		}
	}
	return displayedcolumn
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
	if _, ok := stage.Rows[row]; !ok {
		stage.Rows[row] = struct{}{}
		stage.Row_stagedOrder[row] = stage.RowOrder
		stage.Row_orderStaged[stage.RowOrder] = row
		stage.RowOrder++
	}
	stage.Rows_mapString[row.Name] = row

	return row
}

// StagePreserveOrder puts row to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.RowOrder
// - update stage.RowOrder accordingly
func (row *Row) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Rows[row]; !ok {
		stage.Rows[row] = struct{}{}

		if order > stage.RowOrder {
			stage.RowOrder = order
		}
		stage.Row_stagedOrder[row] = order
		stage.Row_orderStaged[order] = row
		stage.RowOrder++
	}
	stage.Rows_mapString[row.Name] = row
}

// Unstage removes row off the model stage
func (row *Row) Unstage(stage *Stage) *Row {
	delete(stage.Rows, row)
	// issue1150
	// delete(stage.Row_stagedOrder, row)
	delete(stage.Rows_mapString, row.Name)

	return row
}

// UnstageVoid removes row off the model stage
func (row *Row) UnstageVoid(stage *Stage) {
	delete(stage.Rows, row)
	// issue1150
	// delete(stage.Row_stagedOrder, row)
	delete(stage.Rows_mapString, row.Name)
}

// commit row to the back repo (if it is already staged)
func (row *Row) Commit(stage *Stage) *Row {
	if _, ok := stage.Rows[row]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRow(row)
		}
	}
	return row
}

func (row *Row) CommitVoid(stage *Stage) {
	row.Commit(stage)
}

func (row *Row) StageVoid(stage *Stage) {
	row.Stage(stage)
}

// Checkout row to the back repo (if it is already staged)
func (row *Row) Checkout(stage *Stage) *Row {
	if _, ok := stage.Rows[row]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutRow(row)
		}
	}
	return row
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
	if _, ok := stage.SVGIcons[svgicon]; !ok {
		stage.SVGIcons[svgicon] = struct{}{}
		stage.SVGIcon_stagedOrder[svgicon] = stage.SVGIconOrder
		stage.SVGIcon_orderStaged[stage.SVGIconOrder] = svgicon
		stage.SVGIconOrder++
	}
	stage.SVGIcons_mapString[svgicon.Name] = svgicon

	return svgicon
}

// StagePreserveOrder puts svgicon to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.SVGIconOrder
// - update stage.SVGIconOrder accordingly
func (svgicon *SVGIcon) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.SVGIcons[svgicon]; !ok {
		stage.SVGIcons[svgicon] = struct{}{}

		if order > stage.SVGIconOrder {
			stage.SVGIconOrder = order
		}
		stage.SVGIcon_stagedOrder[svgicon] = order
		stage.SVGIcon_orderStaged[order] = svgicon
		stage.SVGIconOrder++
	}
	stage.SVGIcons_mapString[svgicon.Name] = svgicon
}

// Unstage removes svgicon off the model stage
func (svgicon *SVGIcon) Unstage(stage *Stage) *SVGIcon {
	delete(stage.SVGIcons, svgicon)
	// issue1150
	// delete(stage.SVGIcon_stagedOrder, svgicon)
	delete(stage.SVGIcons_mapString, svgicon.Name)

	return svgicon
}

// UnstageVoid removes svgicon off the model stage
func (svgicon *SVGIcon) UnstageVoid(stage *Stage) {
	delete(stage.SVGIcons, svgicon)
	// issue1150
	// delete(stage.SVGIcon_stagedOrder, svgicon)
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

func (svgicon *SVGIcon) StageVoid(stage *Stage) {
	svgicon.Stage(stage)
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

// for satisfaction of GongStruct interface
func (svgicon *SVGIcon) SetName(name string) {
	svgicon.Name = name
}

// Stage puts table to the model stage
func (table *Table) Stage(stage *Stage) *Table {
	if _, ok := stage.Tables[table]; !ok {
		stage.Tables[table] = struct{}{}
		stage.Table_stagedOrder[table] = stage.TableOrder
		stage.Table_orderStaged[stage.TableOrder] = table
		stage.TableOrder++
	}
	stage.Tables_mapString[table.Name] = table

	return table
}

// StagePreserveOrder puts table to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.TableOrder
// - update stage.TableOrder accordingly
func (table *Table) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.Tables[table]; !ok {
		stage.Tables[table] = struct{}{}

		if order > stage.TableOrder {
			stage.TableOrder = order
		}
		stage.Table_stagedOrder[table] = order
		stage.Table_orderStaged[order] = table
		stage.TableOrder++
	}
	stage.Tables_mapString[table.Name] = table
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *Stage) *Table {
	delete(stage.Tables, table)
	// issue1150
	// delete(stage.Table_stagedOrder, table)
	delete(stage.Tables_mapString, table.Name)

	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *Stage) {
	delete(stage.Tables, table)
	// issue1150
	// delete(stage.Table_stagedOrder, table)
	delete(stage.Tables_mapString, table.Name)
}

// commit table to the back repo (if it is already staged)
func (table *Table) Commit(stage *Stage) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTable(table)
		}
	}
	return table
}

func (table *Table) CommitVoid(stage *Stage) {
	table.Commit(stage)
}

func (table *Table) StageVoid(stage *Stage) {
	table.Stage(stage)
}

// Checkout table to the back repo (if it is already staged)
func (table *Table) Checkout(stage *Stage) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutTable(table)
		}
	}
	return table
}

// for satisfaction of GongStruct interface
func (table *Table) GetName() (res string) {
	return table.Name
}

// for satisfaction of GongStruct interface
func (table *Table) SetName(name string) {
	table.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMButton(Button *Button)
	CreateORMCell(Cell *Cell)
	CreateORMCellBoolean(CellBoolean *CellBoolean)
	CreateORMCellFloat64(CellFloat64 *CellFloat64)
	CreateORMCellIcon(CellIcon *CellIcon)
	CreateORMCellInt(CellInt *CellInt)
	CreateORMCellString(CellString *CellString)
	CreateORMDisplayedColumn(DisplayedColumn *DisplayedColumn)
	CreateORMRow(Row *Row)
	CreateORMSVGIcon(SVGIcon *SVGIcon)
	CreateORMTable(Table *Table)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMButton(Button *Button)
	DeleteORMCell(Cell *Cell)
	DeleteORMCellBoolean(CellBoolean *CellBoolean)
	DeleteORMCellFloat64(CellFloat64 *CellFloat64)
	DeleteORMCellIcon(CellIcon *CellIcon)
	DeleteORMCellInt(CellInt *CellInt)
	DeleteORMCellString(CellString *CellString)
	DeleteORMDisplayedColumn(DisplayedColumn *DisplayedColumn)
	DeleteORMRow(Row *Row)
	DeleteORMSVGIcon(SVGIcon *SVGIcon)
	DeleteORMTable(Table *Table)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.Buttons = make(map[*Button]struct{})
	stage.Buttons_mapString = make(map[string]*Button)
	stage.Button_stagedOrder = make(map[*Button]uint)
	stage.ButtonOrder = 0

	stage.Cells = make(map[*Cell]struct{})
	stage.Cells_mapString = make(map[string]*Cell)
	stage.Cell_stagedOrder = make(map[*Cell]uint)
	stage.CellOrder = 0

	stage.CellBooleans = make(map[*CellBoolean]struct{})
	stage.CellBooleans_mapString = make(map[string]*CellBoolean)
	stage.CellBoolean_stagedOrder = make(map[*CellBoolean]uint)
	stage.CellBooleanOrder = 0

	stage.CellFloat64s = make(map[*CellFloat64]struct{})
	stage.CellFloat64s_mapString = make(map[string]*CellFloat64)
	stage.CellFloat64_stagedOrder = make(map[*CellFloat64]uint)
	stage.CellFloat64Order = 0

	stage.CellIcons = make(map[*CellIcon]struct{})
	stage.CellIcons_mapString = make(map[string]*CellIcon)
	stage.CellIcon_stagedOrder = make(map[*CellIcon]uint)
	stage.CellIconOrder = 0

	stage.CellInts = make(map[*CellInt]struct{})
	stage.CellInts_mapString = make(map[string]*CellInt)
	stage.CellInt_stagedOrder = make(map[*CellInt]uint)
	stage.CellIntOrder = 0

	stage.CellStrings = make(map[*CellString]struct{})
	stage.CellStrings_mapString = make(map[string]*CellString)
	stage.CellString_stagedOrder = make(map[*CellString]uint)
	stage.CellStringOrder = 0

	stage.DisplayedColumns = make(map[*DisplayedColumn]struct{})
	stage.DisplayedColumns_mapString = make(map[string]*DisplayedColumn)
	stage.DisplayedColumn_stagedOrder = make(map[*DisplayedColumn]uint)
	stage.DisplayedColumnOrder = 0

	stage.Rows = make(map[*Row]struct{})
	stage.Rows_mapString = make(map[string]*Row)
	stage.Row_stagedOrder = make(map[*Row]uint)
	stage.RowOrder = 0

	stage.SVGIcons = make(map[*SVGIcon]struct{})
	stage.SVGIcons_mapString = make(map[string]*SVGIcon)
	stage.SVGIcon_stagedOrder = make(map[*SVGIcon]uint)
	stage.SVGIconOrder = 0

	stage.Tables = make(map[*Table]struct{})
	stage.Tables_mapString = make(map[string]*Table)
	stage.Table_stagedOrder = make(map[*Table]uint)
	stage.TableOrder = 0

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

	stage.Cells = nil
	stage.Cells_mapString = nil

	stage.CellBooleans = nil
	stage.CellBooleans_mapString = nil

	stage.CellFloat64s = nil
	stage.CellFloat64s_mapString = nil

	stage.CellIcons = nil
	stage.CellIcons_mapString = nil

	stage.CellInts = nil
	stage.CellInts_mapString = nil

	stage.CellStrings = nil
	stage.CellStrings_mapString = nil

	stage.DisplayedColumns = nil
	stage.DisplayedColumns_mapString = nil

	stage.Rows = nil
	stage.Rows_mapString = nil

	stage.SVGIcons = nil
	stage.SVGIcons_mapString = nil

	stage.Tables = nil
	stage.Tables_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for button := range stage.Buttons {
		button.Unstage(stage)
	}

	for cell := range stage.Cells {
		cell.Unstage(stage)
	}

	for cellboolean := range stage.CellBooleans {
		cellboolean.Unstage(stage)
	}

	for cellfloat64 := range stage.CellFloat64s {
		cellfloat64.Unstage(stage)
	}

	for cellicon := range stage.CellIcons {
		cellicon.Unstage(stage)
	}

	for cellint := range stage.CellInts {
		cellint.Unstage(stage)
	}

	for cellstring := range stage.CellStrings {
		cellstring.Unstage(stage)
	}

	for displayedcolumn := range stage.DisplayedColumns {
		displayedcolumn.Unstage(stage)
	}

	for row := range stage.Rows {
		row.Unstage(stage)
	}

	for svgicon := range stage.SVGIcons {
		svgicon.Unstage(stage)
	}

	for table := range stage.Tables {
		table.Unstage(stage)
	}

	// end of insertion point for array nil
}

// Gongstruct is the type parameter for generated generic function that allows
// - access to staged instances
// - navigation between staged instances by going backward association links between gongstruct
// - full refactoring of Gongstruct identifiers / fields
type Gongstruct interface{}

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
	GongGetUUID(stage *Stage) string
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
	case map[*Cell]any:
		return any(&stage.Cells).(*Type)
	case map[*CellBoolean]any:
		return any(&stage.CellBooleans).(*Type)
	case map[*CellFloat64]any:
		return any(&stage.CellFloat64s).(*Type)
	case map[*CellIcon]any:
		return any(&stage.CellIcons).(*Type)
	case map[*CellInt]any:
		return any(&stage.CellInts).(*Type)
	case map[*CellString]any:
		return any(&stage.CellStrings).(*Type)
	case map[*DisplayedColumn]any:
		return any(&stage.DisplayedColumns).(*Type)
	case map[*Row]any:
		return any(&stage.Rows).(*Type)
	case map[*SVGIcon]any:
		return any(&stage.SVGIcons).(*Type)
	case map[*Table]any:
		return any(&stage.Tables).(*Type)
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

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *Stage) *map[*Type]struct{} {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Button:
		return any(&stage.Buttons).(*map[*Type]struct{})
	case Cell:
		return any(&stage.Cells).(*map[*Type]struct{})
	case CellBoolean:
		return any(&stage.CellBooleans).(*map[*Type]struct{})
	case CellFloat64:
		return any(&stage.CellFloat64s).(*map[*Type]struct{})
	case CellIcon:
		return any(&stage.CellIcons).(*map[*Type]struct{})
	case CellInt:
		return any(&stage.CellInts).(*map[*Type]struct{})
	case CellString:
		return any(&stage.CellStrings).(*map[*Type]struct{})
	case DisplayedColumn:
		return any(&stage.DisplayedColumns).(*map[*Type]struct{})
	case Row:
		return any(&stage.Rows).(*map[*Type]struct{})
	case SVGIcon:
		return any(&stage.SVGIcons).(*map[*Type]struct{})
	case Table:
		return any(&stage.Tables).(*map[*Type]struct{})
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

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *Stage) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Button:
		return any(&stage.Buttons_mapString).(*map[string]*Type)
	case Cell:
		return any(&stage.Cells_mapString).(*map[string]*Type)
	case CellBoolean:
		return any(&stage.CellBooleans_mapString).(*map[string]*Type)
	case CellFloat64:
		return any(&stage.CellFloat64s_mapString).(*map[string]*Type)
	case CellIcon:
		return any(&stage.CellIcons_mapString).(*map[string]*Type)
	case CellInt:
		return any(&stage.CellInts_mapString).(*map[string]*Type)
	case CellString:
		return any(&stage.CellStrings_mapString).(*map[string]*Type)
	case DisplayedColumn:
		return any(&stage.DisplayedColumns_mapString).(*map[string]*Type)
	case Row:
		return any(&stage.Rows_mapString).(*map[string]*Type)
	case SVGIcon:
		return any(&stage.SVGIcons_mapString).(*map[string]*Type)
	case Table:
		return any(&stage.Tables_mapString).(*map[string]*Type)
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
	case Cell:
		return any(&Cell{
			// Initialisation of associations
			// field is initialized with an instance of CellString with the name of the field
			CellString: &CellString{Name: "CellString"},
			// field is initialized with an instance of CellFloat64 with the name of the field
			CellFloat64: &CellFloat64{Name: "CellFloat64"},
			// field is initialized with an instance of CellInt with the name of the field
			CellInt: &CellInt{Name: "CellInt"},
			// field is initialized with an instance of CellBoolean with the name of the field
			CellBool: &CellBoolean{Name: "CellBool"},
			// field is initialized with an instance of CellIcon with the name of the field
			CellIcon: &CellIcon{Name: "CellIcon"},
		}).(*Type)
	case CellBoolean:
		return any(&CellBoolean{
			// Initialisation of associations
		}).(*Type)
	case CellFloat64:
		return any(&CellFloat64{
			// Initialisation of associations
		}).(*Type)
	case CellIcon:
		return any(&CellIcon{
			// Initialisation of associations
		}).(*Type)
	case CellInt:
		return any(&CellInt{
			// Initialisation of associations
		}).(*Type)
	case CellString:
		return any(&CellString{
			// Initialisation of associations
		}).(*Type)
	case DisplayedColumn:
		return any(&DisplayedColumn{
			// Initialisation of associations
		}).(*Type)
	case Row:
		return any(&Row{
			// Initialisation of associations
			// field is initialized with an instance of Cell with the name of the field
			Cells: []*Cell{{Name: "Cells"}},
		}).(*Type)
	case SVGIcon:
		return any(&SVGIcon{
			// Initialisation of associations
		}).(*Type)
	case Table:
		return any(&Table{
			// Initialisation of associations
			// field is initialized with an instance of DisplayedColumn with the name of the field
			DisplayedColumns: []*DisplayedColumn{{Name: "DisplayedColumns"}},
			// field is initialized with an instance of Row with the name of the field
			Rows: []*Row{{Name: "Rows"}},
			// field is initialized with an instance of Row with the name of the field
			RowsSelectedForBulkDelete: []*Row{{Name: "RowsSelectedForBulkDelete"}},
			// field is initialized with an instance of Button with the name of the field
			Buttons: []*Button{{Name: "Buttons"}},
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
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

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []GongFieldHeader) {
	var ret Type
	return ret.GongGetFieldHeaders()
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
	case "Icon":
		button.Icon = value.GetValueString()
	case "SVGIcon":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			button.SVGIcon = nil
			for __instance__ := range stage.SVGIcons {
				if stage.SVGIcon_stagedOrder[__instance__] == uint(id) {
					button.SVGIcon = __instance__
					break
				}
			}
		}
	case "IsDisabled":
		button.IsDisabled = value.GetValueBool()
	case "HasToolTip":
		button.HasToolTip = value.GetValueBool()
	case "ToolTipText":
		button.ToolTipText = value.GetValueString()
	case "ToolTipPosition":
		button.ToolTipPosition.FromCodeString(value.GetValueString())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (cell *Cell) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cell.Name = value.GetValueString()
	case "CellString":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			cell.CellString = nil
			for __instance__ := range stage.CellStrings {
				if stage.CellString_stagedOrder[__instance__] == uint(id) {
					cell.CellString = __instance__
					break
				}
			}
		}
	case "CellFloat64":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			cell.CellFloat64 = nil
			for __instance__ := range stage.CellFloat64s {
				if stage.CellFloat64_stagedOrder[__instance__] == uint(id) {
					cell.CellFloat64 = __instance__
					break
				}
			}
		}
	case "CellInt":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			cell.CellInt = nil
			for __instance__ := range stage.CellInts {
				if stage.CellInt_stagedOrder[__instance__] == uint(id) {
					cell.CellInt = __instance__
					break
				}
			}
		}
	case "CellBool":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			cell.CellBool = nil
			for __instance__ := range stage.CellBooleans {
				if stage.CellBoolean_stagedOrder[__instance__] == uint(id) {
					cell.CellBool = __instance__
					break
				}
			}
		}
	case "CellIcon":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			cell.CellIcon = nil
			for __instance__ := range stage.CellIcons {
				if stage.CellIcon_stagedOrder[__instance__] == uint(id) {
					cell.CellIcon = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (cellboolean *CellBoolean) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cellboolean.Name = value.GetValueString()
	case "Value":
		cellboolean.Value = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (cellfloat64 *CellFloat64) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cellfloat64.Name = value.GetValueString()
	case "Value":
		cellfloat64.Value = value.GetValueFloat()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (cellicon *CellIcon) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cellicon.Name = value.GetValueString()
	case "Icon":
		cellicon.Icon = value.GetValueString()
	case "NeedsConfirmation":
		cellicon.NeedsConfirmation = value.GetValueBool()
	case "ConfirmationMessage":
		cellicon.ConfirmationMessage = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (cellint *CellInt) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cellint.Name = value.GetValueString()
	case "Value":
		cellint.Value = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (cellstring *CellString) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		cellstring.Name = value.GetValueString()
	case "Value":
		cellstring.Value = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (displayedcolumn *DisplayedColumn) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		displayedcolumn.Name = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (row *Row) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		row.Name = value.GetValueString()
	case "Cells":
		row.Cells = make([]*Cell, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Cells {
					if stage.Cell_stagedOrder[__instance__] == uint(id) {
						row.Cells = append(row.Cells, __instance__)
						break
					}
				}
			}
		}
	case "IsChecked":
		row.IsChecked = value.GetValueBool()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (svgicon *SVGIcon) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		svgicon.Name = value.GetValueString()
	case "SVG":
		svgicon.SVG = value.GetValueString()
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (table *Table) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		table.Name = value.GetValueString()
	case "DisplayedColumns":
		table.DisplayedColumns = make([]*DisplayedColumn, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.DisplayedColumns {
					if stage.DisplayedColumn_stagedOrder[__instance__] == uint(id) {
						table.DisplayedColumns = append(table.DisplayedColumns, __instance__)
						break
					}
				}
			}
		}
	case "Rows":
		table.Rows = make([]*Row, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Rows {
					if stage.Row_stagedOrder[__instance__] == uint(id) {
						table.Rows = append(table.Rows, __instance__)
						break
					}
				}
			}
		}
	case "HasFiltering":
		table.HasFiltering = value.GetValueBool()
	case "HasColumnSorting":
		table.HasColumnSorting = value.GetValueBool()
	case "HasPaginator":
		table.HasPaginator = value.GetValueBool()
	case "HasCheckableRows":
		table.HasCheckableRows = value.GetValueBool()
	case "HasSaveButton":
		table.HasSaveButton = value.GetValueBool()
	case "SaveButtonLabel":
		table.SaveButtonLabel = value.GetValueString()
	case "HasBulkDeleteButton":
		table.HasBulkDeleteButton = value.GetValueBool()
	case "BulkDeleteButtonTooltip":
		table.BulkDeleteButtonTooltip = value.GetValueString()
	case "RowsSelectedForBulkDelete":
		table.RowsSelectedForBulkDelete = make([]*Row, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Rows {
					if stage.Row_stagedOrder[__instance__] == uint(id) {
						table.RowsSelectedForBulkDelete = append(table.RowsSelectedForBulkDelete, __instance__)
						break
					}
				}
			}
		}
	case "CanDragDropRows":
		table.CanDragDropRows = value.GetValueBool()
	case "HasCloseButton":
		table.HasCloseButton = value.GetValueBool()
	case "SavingInProgress":
		table.SavingInProgress = value.GetValueBool()
	case "NbOfStickyColumns":
		table.NbOfStickyColumns = int(value.GetValueInt())
	case "Buttons":
		table.Buttons = make([]*Button, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.Buttons {
					if stage.Button_stagedOrder[__instance__] == uint(id) {
						table.Buttons = append(table.Buttons, __instance__)
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

	stage.Cells_mapString = make(map[string]*Cell)
	for cell := range stage.Cells {
		stage.Cells_mapString[cell.Name] = cell
	}

	stage.CellBooleans_mapString = make(map[string]*CellBoolean)
	for cellboolean := range stage.CellBooleans {
		stage.CellBooleans_mapString[cellboolean.Name] = cellboolean
	}

	stage.CellFloat64s_mapString = make(map[string]*CellFloat64)
	for cellfloat64 := range stage.CellFloat64s {
		stage.CellFloat64s_mapString[cellfloat64.Name] = cellfloat64
	}

	stage.CellIcons_mapString = make(map[string]*CellIcon)
	for cellicon := range stage.CellIcons {
		stage.CellIcons_mapString[cellicon.Name] = cellicon
	}

	stage.CellInts_mapString = make(map[string]*CellInt)
	for cellint := range stage.CellInts {
		stage.CellInts_mapString[cellint.Name] = cellint
	}

	stage.CellStrings_mapString = make(map[string]*CellString)
	for cellstring := range stage.CellStrings {
		stage.CellStrings_mapString[cellstring.Name] = cellstring
	}

	stage.DisplayedColumns_mapString = make(map[string]*DisplayedColumn)
	for displayedcolumn := range stage.DisplayedColumns {
		stage.DisplayedColumns_mapString[displayedcolumn.Name] = displayedcolumn
	}

	stage.Rows_mapString = make(map[string]*Row)
	for row := range stage.Rows {
		stage.Rows_mapString[row.Name] = row
	}

	stage.SVGIcons_mapString = make(map[string]*SVGIcon)
	for svgicon := range stage.SVGIcons {
		stage.SVGIcons_mapString[svgicon.Name] = svgicon
	}

	stage.Tables_mapString = make(map[string]*Table)
	for table := range stage.Tables {
		stage.Tables_mapString[table.Name] = table
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
