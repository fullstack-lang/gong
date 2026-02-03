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

	xlsx_go "github.com/fullstack-lang/gong/lib/xlsx/go"
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
	ProbeTreeSidebarSuffix       = ":sidebar of the probe"
	ProbeTableSuffix             = ":table of the probe"
	ProbeNotificationTableSuffix = ":notification table of the probe"
	ProbeFormSuffix              = ":form of the probe"
	ProbeSplitSuffix             = ":probe of the probe"
)

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

	// insertion point for definition of arrays registering instances
	DisplaySelections                map[*DisplaySelection]struct{}
	DisplaySelections_reference      map[*DisplaySelection]*DisplaySelection
	DisplaySelections_referenceOrder map[*DisplaySelection]uint // diff Unstage needs the reference order
	DisplaySelections_mapString      map[string]*DisplaySelection

	// insertion point for slice of pointers maps
	OnAfterDisplaySelectionCreateCallback OnAfterCreateInterface[DisplaySelection]
	OnAfterDisplaySelectionUpdateCallback OnAfterUpdateInterface[DisplaySelection]
	OnAfterDisplaySelectionDeleteCallback OnAfterDeleteInterface[DisplaySelection]
	OnAfterDisplaySelectionReadCallback   OnAfterReadInterface[DisplaySelection]

	XLCells                map[*XLCell]struct{}
	XLCells_reference      map[*XLCell]*XLCell
	XLCells_referenceOrder map[*XLCell]uint // diff Unstage needs the reference order
	XLCells_mapString      map[string]*XLCell

	// insertion point for slice of pointers maps
	OnAfterXLCellCreateCallback OnAfterCreateInterface[XLCell]
	OnAfterXLCellUpdateCallback OnAfterUpdateInterface[XLCell]
	OnAfterXLCellDeleteCallback OnAfterDeleteInterface[XLCell]
	OnAfterXLCellReadCallback   OnAfterReadInterface[XLCell]

	XLFiles                map[*XLFile]struct{}
	XLFiles_reference      map[*XLFile]*XLFile
	XLFiles_referenceOrder map[*XLFile]uint // diff Unstage needs the reference order
	XLFiles_mapString      map[string]*XLFile

	// insertion point for slice of pointers maps
	XLFile_Sheets_reverseMap map[*XLSheet]*XLFile

	OnAfterXLFileCreateCallback OnAfterCreateInterface[XLFile]
	OnAfterXLFileUpdateCallback OnAfterUpdateInterface[XLFile]
	OnAfterXLFileDeleteCallback OnAfterDeleteInterface[XLFile]
	OnAfterXLFileReadCallback   OnAfterReadInterface[XLFile]

	XLRows                map[*XLRow]struct{}
	XLRows_reference      map[*XLRow]*XLRow
	XLRows_referenceOrder map[*XLRow]uint // diff Unstage needs the reference order
	XLRows_mapString      map[string]*XLRow

	// insertion point for slice of pointers maps
	XLRow_Cells_reverseMap map[*XLCell]*XLRow

	OnAfterXLRowCreateCallback OnAfterCreateInterface[XLRow]
	OnAfterXLRowUpdateCallback OnAfterUpdateInterface[XLRow]
	OnAfterXLRowDeleteCallback OnAfterDeleteInterface[XLRow]
	OnAfterXLRowReadCallback   OnAfterReadInterface[XLRow]

	XLSheets                map[*XLSheet]struct{}
	XLSheets_reference      map[*XLSheet]*XLSheet
	XLSheets_referenceOrder map[*XLSheet]uint // diff Unstage needs the reference order
	XLSheets_mapString      map[string]*XLSheet

	// insertion point for slice of pointers maps
	XLSheet_Rows_reverseMap map[*XLRow]*XLSheet

	XLSheet_SheetCells_reverseMap map[*XLCell]*XLSheet

	OnAfterXLSheetCreateCallback OnAfterCreateInterface[XLSheet]
	OnAfterXLSheetUpdateCallback OnAfterUpdateInterface[XLSheet]
	OnAfterXLSheetDeleteCallback OnAfterDeleteInterface[XLSheet]
	OnAfterXLSheetReadCallback   OnAfterReadInterface[XLSheet]

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
	DisplaySelectionOrder            uint
	DisplaySelectionMap_Staged_Order map[*DisplaySelection]uint

	XLCellOrder            uint
	XLCellMap_Staged_Order map[*XLCell]uint

	XLFileOrder            uint
	XLFileMap_Staged_Order map[*XLFile]uint

	XLRowOrder            uint
	XLRowMap_Staged_Order map[*XLRow]uint

	XLSheetOrder            uint
	XLSheetMap_Staged_Order map[*XLSheet]uint

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

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}
}

// Orphans removes all commits
func (stage *Stage) Orphans() {
	stage.forwardCommits = stage.forwardCommits[:0]
	stage.backwardCommits = stage.backwardCommits[:0]
	stage.commitsBehind = 0
	stage.navigationMode = GongNavigationModeNormal

	stage.ComputeInstancesNb()
	if stage.OnInitCommitCallback != nil {
		stage.OnInitCommitCallback.BeforeCommit(stage)
	}
	if stage.OnInitCommitFromBackCallback != nil {
		stage.OnInitCommitFromBackCallback.BeforeCommit(stage)
	}
}

// recomputeOrders recomputes the next order for each struct
// this is necessary because the order might have been incremented
// during the commits that have been discarded
// insertion point for max order recomputation
func (stage *Stage) recomputeOrders() {
	// insertion point for max order recomputation
	var maxDisplaySelectionOrder uint
	var foundDisplaySelection bool
	for _, order := range stage.DisplaySelectionMap_Staged_Order {
		if !foundDisplaySelection || order > maxDisplaySelectionOrder {
			maxDisplaySelectionOrder = order
			foundDisplaySelection = true
		}
	}
	if foundDisplaySelection {
		stage.DisplaySelectionOrder = maxDisplaySelectionOrder + 1
	} else {
		stage.DisplaySelectionOrder = 0
	}

	var maxXLCellOrder uint
	var foundXLCell bool
	for _, order := range stage.XLCellMap_Staged_Order {
		if !foundXLCell || order > maxXLCellOrder {
			maxXLCellOrder = order
			foundXLCell = true
		}
	}
	if foundXLCell {
		stage.XLCellOrder = maxXLCellOrder + 1
	} else {
		stage.XLCellOrder = 0
	}

	var maxXLFileOrder uint
	var foundXLFile bool
	for _, order := range stage.XLFileMap_Staged_Order {
		if !foundXLFile || order > maxXLFileOrder {
			maxXLFileOrder = order
			foundXLFile = true
		}
	}
	if foundXLFile {
		stage.XLFileOrder = maxXLFileOrder + 1
	} else {
		stage.XLFileOrder = 0
	}

	var maxXLRowOrder uint
	var foundXLRow bool
	for _, order := range stage.XLRowMap_Staged_Order {
		if !foundXLRow || order > maxXLRowOrder {
			maxXLRowOrder = order
			foundXLRow = true
		}
	}
	if foundXLRow {
		stage.XLRowOrder = maxXLRowOrder + 1
	} else {
		stage.XLRowOrder = 0
	}

	var maxXLSheetOrder uint
	var foundXLSheet bool
	for _, order := range stage.XLSheetMap_Staged_Order {
		if !foundXLSheet || order > maxXLSheetOrder {
			maxXLSheetOrder = order
			foundXLSheet = true
		}
	}
	if foundXLSheet {
		stage.XLSheetOrder = maxXLSheetOrder + 1
	} else {
		stage.XLSheetOrder = 0
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

func GetStructInstancesByOrderAuto[T PointerToGongstruct](stage *Stage) (res []T) {
	var t T
	switch any(t).(type) {
	// insertion point for case
	case *DisplaySelection:
		tmp := GetStructInstancesByOrder(stage.DisplaySelections, stage.DisplaySelectionMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *DisplaySelection implements.
			res = append(res, any(v).(T))
		}
		return res
	case *XLCell:
		tmp := GetStructInstancesByOrder(stage.XLCells, stage.XLCellMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *XLCell implements.
			res = append(res, any(v).(T))
		}
		return res
	case *XLFile:
		tmp := GetStructInstancesByOrder(stage.XLFiles, stage.XLFileMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *XLFile implements.
			res = append(res, any(v).(T))
		}
		return res
	case *XLRow:
		tmp := GetStructInstancesByOrder(stage.XLRows, stage.XLRowMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *XLRow implements.
			res = append(res, any(v).(T))
		}
		return res
	case *XLSheet:
		tmp := GetStructInstancesByOrder(stage.XLSheets, stage.XLSheetMap_Staged_Order)

		// Create a new slice of the generic type T with the same capacity.
		res = make([]T, 0, len(tmp))

		// Iterate over the source slice and perform a type assertion on each element.
		for _, v := range tmp {
			// Assert that the element 'v' can be treated as type 'T'.
			// Note: This relies on the constraint that PointerToGongstruct
			// is an interface that *XLSheet implements.
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
	case "DisplaySelection":
		res = GetNamedStructInstances(stage.DisplaySelections, stage.DisplaySelectionMap_Staged_Order)
	case "XLCell":
		res = GetNamedStructInstances(stage.XLCells, stage.XLCellMap_Staged_Order)
	case "XLFile":
		res = GetNamedStructInstances(stage.XLFiles, stage.XLFileMap_Staged_Order)
	case "XLRow":
		res = GetNamedStructInstances(stage.XLRows, stage.XLRowMap_Staged_Order)
	case "XLSheet":
		res = GetNamedStructInstances(stage.XLSheets, stage.XLSheetMap_Staged_Order)
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
	return "github.com/fullstack-lang/gong/lib/xlsx/go/models"
}

func (stage *Stage) GetMap_GongStructName_InstancesNb() map[string]int {
	return stage.Map_GongStructName_InstancesNb
}

func (stage *Stage) GetModelsEmbededDir() embed.FS {
	return xlsx_go.GoModelsDir
}

func (stage *Stage) GetDigramsEmbededDir() embed.FS {
	return xlsx_go.GoDiagramsDir
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
	CommitDisplaySelection(displayselection *DisplaySelection)
	CheckoutDisplaySelection(displayselection *DisplaySelection)
	CommitXLCell(xlcell *XLCell)
	CheckoutXLCell(xlcell *XLCell)
	CommitXLFile(xlfile *XLFile)
	CheckoutXLFile(xlfile *XLFile)
	CommitXLRow(xlrow *XLRow)
	CheckoutXLRow(xlrow *XLRow)
	CommitXLSheet(xlsheet *XLSheet)
	CheckoutXLSheet(xlsheet *XLSheet)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

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

		name: name,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here

		// insertion point for order map initialisations
		DisplaySelectionMap_Staged_Order: make(map[*DisplaySelection]uint),

		XLCellMap_Staged_Order: make(map[*XLCell]uint),

		XLFileMap_Staged_Order: make(map[*XLFile]uint),

		XLRowMap_Staged_Order: make(map[*XLRow]uint),

		XLSheetMap_Staged_Order: make(map[*XLSheet]uint),

		// end of insertion point
		GongUnmarshallers: map[string]ModelUnmarshaller{ // insertion point for unmarshallers
			"DisplaySelection": &DisplaySelectionUnmarshaller{},

			"XLCell": &XLCellUnmarshaller{},

			"XLFile": &XLFileUnmarshaller{},

			"XLRow": &XLRowUnmarshaller{},

			"XLSheet": &XLSheetUnmarshaller{},

			// end of insertion point
		},

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "DisplaySelection"},
			{name: "XLCell"},
			{name: "XLFile"},
			{name: "XLRow"},
			{name: "XLSheet"},
		}, // end of insertion point

		navigationMode: GongNavigationModeNormal,
	}

	return
}

func GetOrder[Type Gongstruct](stage *Stage, instance *Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DisplaySelection:
		return stage.DisplaySelectionMap_Staged_Order[instance]
	case *XLCell:
		return stage.XLCellMap_Staged_Order[instance]
	case *XLFile:
		return stage.XLFileMap_Staged_Order[instance]
	case *XLRow:
		return stage.XLRowMap_Staged_Order[instance]
	case *XLSheet:
		return stage.XLSheetMap_Staged_Order[instance]
	default:
		return 0 // should not happen
	}
}

func GetOrderPointerGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) uint {
	switch instance := any(instance).(type) {
	// insertion point for order map initialisations
	case *DisplaySelection:
		return stage.DisplaySelectionMap_Staged_Order[instance]
	case *XLCell:
		return stage.XLCellMap_Staged_Order[instance]
	case *XLFile:
		return stage.XLFileMap_Staged_Order[instance]
	case *XLRow:
		return stage.XLRowMap_Staged_Order[instance]
	case *XLSheet:
		return stage.XLSheetMap_Staged_Order[instance]
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
	if _, ok := stage.DisplaySelections[displayselection]; !ok {
		stage.DisplaySelections[displayselection] = struct{}{}
		stage.DisplaySelectionMap_Staged_Order[displayselection] = stage.DisplaySelectionOrder
		stage.DisplaySelectionOrder++
	}
	stage.DisplaySelections_mapString[displayselection.Name] = displayselection

	return displayselection
}

// StagePreserveOrder puts displayselection to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.DisplaySelectionOrder
// - update stage.DisplaySelectionOrder accordingly
func (displayselection *DisplaySelection) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.DisplaySelections[displayselection]; !ok {
		stage.DisplaySelections[displayselection] = struct{}{}

		if order > stage.DisplaySelectionOrder {
			stage.DisplaySelectionOrder = order
		}
		stage.DisplaySelectionMap_Staged_Order[displayselection] = order
		stage.DisplaySelectionOrder++
	}
	stage.DisplaySelections_mapString[displayselection.Name] = displayselection
}

// Unstage removes displayselection off the model stage
func (displayselection *DisplaySelection) Unstage(stage *Stage) *DisplaySelection {
	delete(stage.DisplaySelections, displayselection)
	delete(stage.DisplaySelectionMap_Staged_Order, displayselection)
	delete(stage.DisplaySelections_mapString, displayselection.Name)

	return displayselection
}

// UnstageVoid removes displayselection off the model stage
func (displayselection *DisplaySelection) UnstageVoid(stage *Stage) {
	delete(stage.DisplaySelections, displayselection)
	delete(stage.DisplaySelectionMap_Staged_Order, displayselection)
	delete(stage.DisplaySelections_mapString, displayselection.Name)
}

// commit displayselection to the back repo (if it is already staged)
func (displayselection *DisplaySelection) Commit(stage *Stage) *DisplaySelection {
	if _, ok := stage.DisplaySelections[displayselection]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDisplaySelection(displayselection)
		}
	}
	return displayselection
}

func (displayselection *DisplaySelection) CommitVoid(stage *Stage) {
	displayselection.Commit(stage)
}

func (displayselection *DisplaySelection) StageVoid(stage *Stage) {
	displayselection.Stage(stage)
}

// Checkout displayselection to the back repo (if it is already staged)
func (displayselection *DisplaySelection) Checkout(stage *Stage) *DisplaySelection {
	if _, ok := stage.DisplaySelections[displayselection]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutDisplaySelection(displayselection)
		}
	}
	return displayselection
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
	if _, ok := stage.XLCells[xlcell]; !ok {
		stage.XLCells[xlcell] = struct{}{}
		stage.XLCellMap_Staged_Order[xlcell] = stage.XLCellOrder
		stage.XLCellOrder++
	}
	stage.XLCells_mapString[xlcell.Name] = xlcell

	return xlcell
}

// StagePreserveOrder puts xlcell to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLCellOrder
// - update stage.XLCellOrder accordingly
func (xlcell *XLCell) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.XLCells[xlcell]; !ok {
		stage.XLCells[xlcell] = struct{}{}

		if order > stage.XLCellOrder {
			stage.XLCellOrder = order
		}
		stage.XLCellMap_Staged_Order[xlcell] = order
		stage.XLCellOrder++
	}
	stage.XLCells_mapString[xlcell.Name] = xlcell
}

// Unstage removes xlcell off the model stage
func (xlcell *XLCell) Unstage(stage *Stage) *XLCell {
	delete(stage.XLCells, xlcell)
	delete(stage.XLCellMap_Staged_Order, xlcell)
	delete(stage.XLCells_mapString, xlcell.Name)

	return xlcell
}

// UnstageVoid removes xlcell off the model stage
func (xlcell *XLCell) UnstageVoid(stage *Stage) {
	delete(stage.XLCells, xlcell)
	delete(stage.XLCellMap_Staged_Order, xlcell)
	delete(stage.XLCells_mapString, xlcell.Name)
}

// commit xlcell to the back repo (if it is already staged)
func (xlcell *XLCell) Commit(stage *Stage) *XLCell {
	if _, ok := stage.XLCells[xlcell]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXLCell(xlcell)
		}
	}
	return xlcell
}

func (xlcell *XLCell) CommitVoid(stage *Stage) {
	xlcell.Commit(stage)
}

func (xlcell *XLCell) StageVoid(stage *Stage) {
	xlcell.Stage(stage)
}

// Checkout xlcell to the back repo (if it is already staged)
func (xlcell *XLCell) Checkout(stage *Stage) *XLCell {
	if _, ok := stage.XLCells[xlcell]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXLCell(xlcell)
		}
	}
	return xlcell
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
	if _, ok := stage.XLFiles[xlfile]; !ok {
		stage.XLFiles[xlfile] = struct{}{}
		stage.XLFileMap_Staged_Order[xlfile] = stage.XLFileOrder
		stage.XLFileOrder++
	}
	stage.XLFiles_mapString[xlfile.Name] = xlfile

	return xlfile
}

// StagePreserveOrder puts xlfile to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLFileOrder
// - update stage.XLFileOrder accordingly
func (xlfile *XLFile) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.XLFiles[xlfile]; !ok {
		stage.XLFiles[xlfile] = struct{}{}

		if order > stage.XLFileOrder {
			stage.XLFileOrder = order
		}
		stage.XLFileMap_Staged_Order[xlfile] = order
		stage.XLFileOrder++
	}
	stage.XLFiles_mapString[xlfile.Name] = xlfile
}

// Unstage removes xlfile off the model stage
func (xlfile *XLFile) Unstage(stage *Stage) *XLFile {
	delete(stage.XLFiles, xlfile)
	delete(stage.XLFileMap_Staged_Order, xlfile)
	delete(stage.XLFiles_mapString, xlfile.Name)

	return xlfile
}

// UnstageVoid removes xlfile off the model stage
func (xlfile *XLFile) UnstageVoid(stage *Stage) {
	delete(stage.XLFiles, xlfile)
	delete(stage.XLFileMap_Staged_Order, xlfile)
	delete(stage.XLFiles_mapString, xlfile.Name)
}

// commit xlfile to the back repo (if it is already staged)
func (xlfile *XLFile) Commit(stage *Stage) *XLFile {
	if _, ok := stage.XLFiles[xlfile]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXLFile(xlfile)
		}
	}
	return xlfile
}

func (xlfile *XLFile) CommitVoid(stage *Stage) {
	xlfile.Commit(stage)
}

func (xlfile *XLFile) StageVoid(stage *Stage) {
	xlfile.Stage(stage)
}

// Checkout xlfile to the back repo (if it is already staged)
func (xlfile *XLFile) Checkout(stage *Stage) *XLFile {
	if _, ok := stage.XLFiles[xlfile]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXLFile(xlfile)
		}
	}
	return xlfile
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
	if _, ok := stage.XLRows[xlrow]; !ok {
		stage.XLRows[xlrow] = struct{}{}
		stage.XLRowMap_Staged_Order[xlrow] = stage.XLRowOrder
		stage.XLRowOrder++
	}
	stage.XLRows_mapString[xlrow.Name] = xlrow

	return xlrow
}

// StagePreserveOrder puts xlrow to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLRowOrder
// - update stage.XLRowOrder accordingly
func (xlrow *XLRow) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.XLRows[xlrow]; !ok {
		stage.XLRows[xlrow] = struct{}{}

		if order > stage.XLRowOrder {
			stage.XLRowOrder = order
		}
		stage.XLRowMap_Staged_Order[xlrow] = order
		stage.XLRowOrder++
	}
	stage.XLRows_mapString[xlrow.Name] = xlrow
}

// Unstage removes xlrow off the model stage
func (xlrow *XLRow) Unstage(stage *Stage) *XLRow {
	delete(stage.XLRows, xlrow)
	delete(stage.XLRowMap_Staged_Order, xlrow)
	delete(stage.XLRows_mapString, xlrow.Name)

	return xlrow
}

// UnstageVoid removes xlrow off the model stage
func (xlrow *XLRow) UnstageVoid(stage *Stage) {
	delete(stage.XLRows, xlrow)
	delete(stage.XLRowMap_Staged_Order, xlrow)
	delete(stage.XLRows_mapString, xlrow.Name)
}

// commit xlrow to the back repo (if it is already staged)
func (xlrow *XLRow) Commit(stage *Stage) *XLRow {
	if _, ok := stage.XLRows[xlrow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXLRow(xlrow)
		}
	}
	return xlrow
}

func (xlrow *XLRow) CommitVoid(stage *Stage) {
	xlrow.Commit(stage)
}

func (xlrow *XLRow) StageVoid(stage *Stage) {
	xlrow.Stage(stage)
}

// Checkout xlrow to the back repo (if it is already staged)
func (xlrow *XLRow) Checkout(stage *Stage) *XLRow {
	if _, ok := stage.XLRows[xlrow]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXLRow(xlrow)
		}
	}
	return xlrow
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
	if _, ok := stage.XLSheets[xlsheet]; !ok {
		stage.XLSheets[xlsheet] = struct{}{}
		stage.XLSheetMap_Staged_Order[xlsheet] = stage.XLSheetOrder
		stage.XLSheetOrder++
	}
	stage.XLSheets_mapString[xlsheet.Name] = xlsheet

	return xlsheet
}

// StagePreserveOrder puts xlsheet to the model stage, and if the astrtuct
// was not staged before:
//
// - force the order if the order is equal or greater than the stage.XLSheetOrder
// - update stage.XLSheetOrder accordingly
func (xlsheet *XLSheet) StagePreserveOrder(stage *Stage, order uint) {
	if _, ok := stage.XLSheets[xlsheet]; !ok {
		stage.XLSheets[xlsheet] = struct{}{}

		if order > stage.XLSheetOrder {
			stage.XLSheetOrder = order
		}
		stage.XLSheetMap_Staged_Order[xlsheet] = order
		stage.XLSheetOrder++
	}
	stage.XLSheets_mapString[xlsheet.Name] = xlsheet
}

// Unstage removes xlsheet off the model stage
func (xlsheet *XLSheet) Unstage(stage *Stage) *XLSheet {
	delete(stage.XLSheets, xlsheet)
	delete(stage.XLSheetMap_Staged_Order, xlsheet)
	delete(stage.XLSheets_mapString, xlsheet.Name)

	return xlsheet
}

// UnstageVoid removes xlsheet off the model stage
func (xlsheet *XLSheet) UnstageVoid(stage *Stage) {
	delete(stage.XLSheets, xlsheet)
	delete(stage.XLSheetMap_Staged_Order, xlsheet)
	delete(stage.XLSheets_mapString, xlsheet.Name)
}

// commit xlsheet to the back repo (if it is already staged)
func (xlsheet *XLSheet) Commit(stage *Stage) *XLSheet {
	if _, ok := stage.XLSheets[xlsheet]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitXLSheet(xlsheet)
		}
	}
	return xlsheet
}

func (xlsheet *XLSheet) CommitVoid(stage *Stage) {
	xlsheet.Commit(stage)
}

func (xlsheet *XLSheet) StageVoid(stage *Stage) {
	xlsheet.Stage(stage)
}

// Checkout xlsheet to the back repo (if it is already staged)
func (xlsheet *XLSheet) Checkout(stage *Stage) *XLSheet {
	if _, ok := stage.XLSheets[xlsheet]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutXLSheet(xlsheet)
		}
	}
	return xlsheet
}

// for satisfaction of GongStruct interface
func (xlsheet *XLSheet) GetName() (res string) {
	return xlsheet.Name
}

// for satisfaction of GongStruct interface
func (xlsheet *XLSheet) SetName(name string) {
	xlsheet.Name = name
}

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMDisplaySelection(DisplaySelection *DisplaySelection)
	CreateORMXLCell(XLCell *XLCell)
	CreateORMXLFile(XLFile *XLFile)
	CreateORMXLRow(XLRow *XLRow)
	CreateORMXLSheet(XLSheet *XLSheet)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMDisplaySelection(DisplaySelection *DisplaySelection)
	DeleteORMXLCell(XLCell *XLCell)
	DeleteORMXLFile(XLFile *XLFile)
	DeleteORMXLRow(XLRow *XLRow)
	DeleteORMXLSheet(XLSheet *XLSheet)
}

func (stage *Stage) Reset() { // insertion point for array reset
	stage.DisplaySelections = make(map[*DisplaySelection]struct{})
	stage.DisplaySelections_mapString = make(map[string]*DisplaySelection)
	stage.DisplaySelectionMap_Staged_Order = make(map[*DisplaySelection]uint)
	stage.DisplaySelectionOrder = 0

	stage.XLCells = make(map[*XLCell]struct{})
	stage.XLCells_mapString = make(map[string]*XLCell)
	stage.XLCellMap_Staged_Order = make(map[*XLCell]uint)
	stage.XLCellOrder = 0

	stage.XLFiles = make(map[*XLFile]struct{})
	stage.XLFiles_mapString = make(map[string]*XLFile)
	stage.XLFileMap_Staged_Order = make(map[*XLFile]uint)
	stage.XLFileOrder = 0

	stage.XLRows = make(map[*XLRow]struct{})
	stage.XLRows_mapString = make(map[string]*XLRow)
	stage.XLRowMap_Staged_Order = make(map[*XLRow]uint)
	stage.XLRowOrder = 0

	stage.XLSheets = make(map[*XLSheet]struct{})
	stage.XLSheets_mapString = make(map[string]*XLSheet)
	stage.XLSheetMap_Staged_Order = make(map[*XLSheet]uint)
	stage.XLSheetOrder = 0

	if stage.GetProbeIF() != nil {
		stage.GetProbeIF().ResetNotifications()
	}
	if stage.IsInDeltaMode() {
		stage.ComputeReferenceAndOrders()
	}
}

func (stage *Stage) Nil() { // insertion point for array nil
	stage.DisplaySelections = nil
	stage.DisplaySelections_mapString = nil

	stage.XLCells = nil
	stage.XLCells_mapString = nil

	stage.XLFiles = nil
	stage.XLFiles_mapString = nil

	stage.XLRows = nil
	stage.XLRows_mapString = nil

	stage.XLSheets = nil
	stage.XLSheets_mapString = nil

	// end of insertion point for array nil
}

func (stage *Stage) Unstage() { // insertion point for array nil
	for displayselection := range stage.DisplaySelections {
		displayselection.Unstage(stage)
	}

	for xlcell := range stage.XLCells {
		xlcell.Unstage(stage)
	}

	for xlfile := range stage.XLFiles {
		xlfile.Unstage(stage)
	}

	for xlrow := range stage.XLRows {
		xlrow.Unstage(stage)
	}

	for xlsheet := range stage.XLSheets {
		xlsheet.Unstage(stage)
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
	case map[*DisplaySelection]any:
		return any(&stage.DisplaySelections).(*Type)
	case map[*XLCell]any:
		return any(&stage.XLCells).(*Type)
	case map[*XLFile]any:
		return any(&stage.XLFiles).(*Type)
	case map[*XLRow]any:
		return any(&stage.XLRows).(*Type)
	case map[*XLSheet]any:
		return any(&stage.XLSheets).(*Type)
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
	case *DisplaySelection:
		return any(stage.DisplaySelections_mapString).(map[string]Type)
	case *XLCell:
		return any(stage.XLCells_mapString).(map[string]Type)
	case *XLFile:
		return any(stage.XLFiles_mapString).(map[string]Type)
	case *XLRow:
		return any(stage.XLRows_mapString).(map[string]Type)
	case *XLSheet:
		return any(stage.XLSheets_mapString).(map[string]Type)
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
	case DisplaySelection:
		return any(&stage.DisplaySelections).(*map[*Type]struct{})
	case XLCell:
		return any(&stage.XLCells).(*map[*Type]struct{})
	case XLFile:
		return any(&stage.XLFiles).(*map[*Type]struct{})
	case XLRow:
		return any(&stage.XLRows).(*map[*Type]struct{})
	case XLSheet:
		return any(&stage.XLSheets).(*map[*Type]struct{})
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
	case *DisplaySelection:
		return any(&stage.DisplaySelections).(*map[Type]struct{})
	case *XLCell:
		return any(&stage.XLCells).(*map[Type]struct{})
	case *XLFile:
		return any(&stage.XLFiles).(*map[Type]struct{})
	case *XLRow:
		return any(&stage.XLRows).(*map[Type]struct{})
	case *XLSheet:
		return any(&stage.XLSheets).(*map[Type]struct{})
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
	case DisplaySelection:
		return any(&stage.DisplaySelections_mapString).(*map[string]*Type)
	case XLCell:
		return any(&stage.XLCells_mapString).(*map[string]*Type)
	case XLFile:
		return any(&stage.XLFiles_mapString).(*map[string]*Type)
	case XLRow:
		return any(&stage.XLRows_mapString).(*map[string]*Type)
	case XLSheet:
		return any(&stage.XLSheets_mapString).(*map[string]*Type)
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
	case DisplaySelection:
		return any(&DisplaySelection{
			// Initialisation of associations
			// field is initialized with an instance of XLFile with the name of the field
			XLFile: &XLFile{Name: "XLFile"},
			// field is initialized with an instance of XLSheet with the name of the field
			XLSheet: &XLSheet{Name: "XLSheet"},
		}).(*Type)
	case XLCell:
		return any(&XLCell{
			// Initialisation of associations
		}).(*Type)
	case XLFile:
		return any(&XLFile{
			// Initialisation of associations
			// field is initialized with an instance of XLSheet with the name of the field
			Sheets: []*XLSheet{{Name: "Sheets"}},
		}).(*Type)
	case XLRow:
		return any(&XLRow{
			// Initialisation of associations
			// field is initialized with an instance of XLCell with the name of the field
			Cells: []*XLCell{{Name: "Cells"}},
		}).(*Type)
	case XLSheet:
		return any(&XLSheet{
			// Initialisation of associations
			// field is initialized with an instance of XLRow with the name of the field
			Rows: []*XLRow{{Name: "Rows"}},
			// field is initialized with an instance of XLCell with the name of the field
			SheetCells: []*XLCell{{Name: "SheetCells"}},
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

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type GongstructIF]() (res string) {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *DisplaySelection:
		res = "DisplaySelection"
	case *XLCell:
		res = "XLCell"
	case *XLFile:
		res = "XLFile"
	case *XLRow:
		res = "XLRow"
	case *XLSheet:
		res = "XLSheet"
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
	case *DisplaySelection:
		var rf ReverseField
		_ = rf
	case *XLCell:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "XLRow"
		rf.Fieldname = "Cells"
		res = append(res, rf)
		rf.GongstructName = "XLSheet"
		rf.Fieldname = "SheetCells"
		res = append(res, rf)
	case *XLFile:
		var rf ReverseField
		_ = rf
	case *XLRow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "XLSheet"
		rf.Fieldname = "Rows"
		res = append(res, rf)
	case *XLSheet:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "XLFile"
		rf.Fieldname = "Sheets"
		res = append(res, rf)
	}
	return
}

// insertion point for get fields header method
func (displayselection *DisplaySelection) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "X",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "Y",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
	}
	return
}

func (xlfile *XLFile) GongGetFieldHeaders() (res []GongFieldHeader) {
	// insertion point for list of field headers
	res = []GongFieldHeader{
		{
			Name:               "Name",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NbSheets",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "RowIndex",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MaxRow",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "MaxCol",
			GongFieldValueType: GongFieldValueTypeBasicKind,
		},
		{
			Name:               "NbRows",
			GongFieldValueType: GongFieldValueTypeBasicKind,
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
func (displayselection *DisplaySelection) GongGetFieldValue(fieldName string, stage *Stage) (res GongFieldValue) {
	switch fieldName {
	// string value of fields
	case "Name":
		res.valueString = displayselection.Name
	case "XLFile":
		res.GongFieldValueType = GongFieldValueTypePointer
		if displayselection.XLFile != nil {
			res.valueString = displayselection.XLFile.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, displayselection.XLFile))
		}
	case "XLSheet":
		res.GongFieldValueType = GongFieldValueTypePointer
		if displayselection.XLSheet != nil {
			res.valueString = displayselection.XLSheet.Name
			res.ids = fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, displayselection.XLSheet))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
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
			res.ids += fmt.Sprintf("%d", GetOrderPointerGongstruct(stage, __instance__))
		}
	case "SheetCells":
		res.GongFieldValueType = GongFieldValueTypeSliceOfPointers
		for idx, __instance__ := range xlsheet.SheetCells {
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
func (displayselection *DisplaySelection) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		displayselection.Name = value.GetValueString()
	case "XLFile":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			displayselection.XLFile = nil
			for __instance__ := range stage.XLFiles {
				if stage.XLFileMap_Staged_Order[__instance__] == uint(id) {
					displayselection.XLFile = __instance__
					break
				}
			}
		}
	case "XLSheet":
		var id int
		if _, err := fmt.Sscanf(value.ids, "%d", &id); err == nil {
			displayselection.XLSheet = nil
			for __instance__ := range stage.XLSheets {
				if stage.XLSheetMap_Staged_Order[__instance__] == uint(id) {
					displayselection.XLSheet = __instance__
					break
				}
			}
		}
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (xlcell *XLCell) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		xlcell.Name = value.GetValueString()
	case "X":
		xlcell.X = int(value.GetValueInt())
	case "Y":
		xlcell.Y = int(value.GetValueInt())
	default:
		return fmt.Errorf("unknown field %s", fieldName)
	}
	return nil
}

func (xlfile *XLFile) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		xlfile.Name = value.GetValueString()
	case "NbSheets":
		xlfile.NbSheets = int(value.GetValueInt())
	case "Sheets":
		xlfile.Sheets = make([]*XLSheet, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.XLSheets {
					if stage.XLSheetMap_Staged_Order[__instance__] == uint(id) {
						xlfile.Sheets = append(xlfile.Sheets, __instance__)
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

func (xlrow *XLRow) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		xlrow.Name = value.GetValueString()
	case "RowIndex":
		xlrow.RowIndex = int(value.GetValueInt())
	case "Cells":
		xlrow.Cells = make([]*XLCell, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.XLCells {
					if stage.XLCellMap_Staged_Order[__instance__] == uint(id) {
						xlrow.Cells = append(xlrow.Cells, __instance__)
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

func (xlsheet *XLSheet) GongSetFieldValue(fieldName string, value GongFieldValue, stage *Stage) error {
	switch fieldName {
	// insertion point for per field code
	case "Name":
		xlsheet.Name = value.GetValueString()
	case "MaxRow":
		xlsheet.MaxRow = int(value.GetValueInt())
	case "MaxCol":
		xlsheet.MaxCol = int(value.GetValueInt())
	case "NbRows":
		xlsheet.NbRows = int(value.GetValueInt())
	case "Rows":
		xlsheet.Rows = make([]*XLRow, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.XLRows {
					if stage.XLRowMap_Staged_Order[__instance__] == uint(id) {
						xlsheet.Rows = append(xlsheet.Rows, __instance__)
						break
					}
				}
			}
		}
	case "SheetCells":
		xlsheet.SheetCells = make([]*XLCell, 0)
		ids := strings.Split(value.ids, ";")
		for _, idStr := range ids {
			var id int
			if _, err := fmt.Sscanf(idStr, "%d", &id); err == nil {
				for __instance__ := range stage.XLCells {
					if stage.XLCellMap_Staged_Order[__instance__] == uint(id) {
						xlsheet.SheetCells = append(xlsheet.SheetCells, __instance__)
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

func GetGongstructNameFromPointer(instance GongstructIF) (res string) {
	res = instance.GongGetGongstructName()
	return
}

func (stage *Stage) ResetMapStrings() {
	// insertion point for generic get gongstruct name
	stage.DisplaySelections_mapString = make(map[string]*DisplaySelection)
	for displayselection := range stage.DisplaySelections {
		stage.DisplaySelections_mapString[displayselection.Name] = displayselection
	}

	stage.XLCells_mapString = make(map[string]*XLCell)
	for xlcell := range stage.XLCells {
		stage.XLCells_mapString[xlcell.Name] = xlcell
	}

	stage.XLFiles_mapString = make(map[string]*XLFile)
	for xlfile := range stage.XLFiles {
		stage.XLFiles_mapString[xlfile.Name] = xlfile
	}

	stage.XLRows_mapString = make(map[string]*XLRow)
	for xlrow := range stage.XLRows {
		stage.XLRows_mapString[xlrow.Name] = xlrow
	}

	stage.XLSheets_mapString = make(map[string]*XLSheet)
	for xlsheet := range stage.XLSheets {
		stage.XLSheets_mapString[xlsheet.Name] = xlsheet
	}

	// end of insertion point for generic get gongstruct name
}

// Last line of the template
