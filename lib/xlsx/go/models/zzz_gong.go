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
	name               string
	commitId           uint // commitId is updated at each commit
	commitTimeStamp    time.Time
	contentWhenParsed  string
	commitIdWhenParsed uint
	generatesDiff      bool

	// insertion point for definition of arrays registering instances
	DisplaySelections           map[*DisplaySelection]any
	DisplaySelections_mapString map[string]*DisplaySelection

	// insertion point for slice of pointers maps
	OnAfterDisplaySelectionCreateCallback OnAfterCreateInterface[DisplaySelection]
	OnAfterDisplaySelectionUpdateCallback OnAfterUpdateInterface[DisplaySelection]
	OnAfterDisplaySelectionDeleteCallback OnAfterDeleteInterface[DisplaySelection]
	OnAfterDisplaySelectionReadCallback   OnAfterReadInterface[DisplaySelection]

	XLCells           map[*XLCell]any
	XLCells_mapString map[string]*XLCell

	// insertion point for slice of pointers maps
	OnAfterXLCellCreateCallback OnAfterCreateInterface[XLCell]
	OnAfterXLCellUpdateCallback OnAfterUpdateInterface[XLCell]
	OnAfterXLCellDeleteCallback OnAfterDeleteInterface[XLCell]
	OnAfterXLCellReadCallback   OnAfterReadInterface[XLCell]

	XLFiles           map[*XLFile]any
	XLFiles_mapString map[string]*XLFile

	// insertion point for slice of pointers maps
	XLFile_Sheets_reverseMap map[*XLSheet]*XLFile

	OnAfterXLFileCreateCallback OnAfterCreateInterface[XLFile]
	OnAfterXLFileUpdateCallback OnAfterUpdateInterface[XLFile]
	OnAfterXLFileDeleteCallback OnAfterDeleteInterface[XLFile]
	OnAfterXLFileReadCallback   OnAfterReadInterface[XLFile]

	XLRows           map[*XLRow]any
	XLRows_mapString map[string]*XLRow

	// insertion point for slice of pointers maps
	XLRow_Cells_reverseMap map[*XLCell]*XLRow

	OnAfterXLRowCreateCallback OnAfterCreateInterface[XLRow]
	OnAfterXLRowUpdateCallback OnAfterUpdateInterface[XLRow]
	OnAfterXLRowDeleteCallback OnAfterDeleteInterface[XLRow]
	OnAfterXLRowReadCallback   OnAfterReadInterface[XLRow]

	XLSheets           map[*XLSheet]any
	XLSheets_mapString map[string]*XLSheet

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
}

func (stage *Stage) SetGeneratesDiff(generatesDiff bool) {
	stage.generatesDiff = generatesDiff
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

func GetStructInstancesByOrder[T PointerToGongstruct](set map[T]any, order map[T]uint) (res []T) {

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
		res = append(res, instance)
	}

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
		DisplaySelections:           make(map[*DisplaySelection]any),
		DisplaySelections_mapString: make(map[string]*DisplaySelection),

		XLCells:           make(map[*XLCell]any),
		XLCells_mapString: make(map[string]*XLCell),

		XLFiles:           make(map[*XLFile]any),
		XLFiles_mapString: make(map[string]*XLFile),

		XLRows:           make(map[*XLRow]any),
		XLRows_mapString: make(map[string]*XLRow),

		XLSheets:           make(map[*XLSheet]any),
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

		NamedStructs: []*NamedStruct{ // insertion point for order map initialisations
			{name: "DisplaySelection"},
			{name: "XLCell"},
			{name: "XLFile"},
			{name: "XLRow"},
			{name: "XLSheet"},
		}, // end of insertion point
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
	stage.commitId++
	stage.commitTimeStamp = time.Now()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

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
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["DisplaySelection"] = len(stage.DisplaySelections)
	stage.Map_GongStructName_InstancesNb["XLCell"] = len(stage.XLCells)
	stage.Map_GongStructName_InstancesNb["XLFile"] = len(stage.XLFiles)
	stage.Map_GongStructName_InstancesNb["XLRow"] = len(stage.XLRows)
	stage.Map_GongStructName_InstancesNb["XLSheet"] = len(stage.XLSheets)

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
		stage.DisplaySelections[displayselection] = __member
		stage.DisplaySelectionMap_Staged_Order[displayselection] = stage.DisplaySelectionOrder
		stage.DisplaySelectionOrder++
	}
	stage.DisplaySelections_mapString[displayselection.Name] = displayselection

	return displayselection
}

// Unstage removes displayselection off the model stage
func (displayselection *DisplaySelection) Unstage(stage *Stage) *DisplaySelection {
	delete(stage.DisplaySelections, displayselection)
	delete(stage.DisplaySelections_mapString, displayselection.Name)
	return displayselection
}

// UnstageVoid removes displayselection off the model stage
func (displayselection *DisplaySelection) UnstageVoid(stage *Stage) {
	delete(stage.DisplaySelections, displayselection)
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

// Stage puts xlcell to the model stage
func (xlcell *XLCell) Stage(stage *Stage) *XLCell {

	if _, ok := stage.XLCells[xlcell]; !ok {
		stage.XLCells[xlcell] = __member
		stage.XLCellMap_Staged_Order[xlcell] = stage.XLCellOrder
		stage.XLCellOrder++
	}
	stage.XLCells_mapString[xlcell.Name] = xlcell

	return xlcell
}

// Unstage removes xlcell off the model stage
func (xlcell *XLCell) Unstage(stage *Stage) *XLCell {
	delete(stage.XLCells, xlcell)
	delete(stage.XLCells_mapString, xlcell.Name)
	return xlcell
}

// UnstageVoid removes xlcell off the model stage
func (xlcell *XLCell) UnstageVoid(stage *Stage) {
	delete(stage.XLCells, xlcell)
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

// Stage puts xlfile to the model stage
func (xlfile *XLFile) Stage(stage *Stage) *XLFile {

	if _, ok := stage.XLFiles[xlfile]; !ok {
		stage.XLFiles[xlfile] = __member
		stage.XLFileMap_Staged_Order[xlfile] = stage.XLFileOrder
		stage.XLFileOrder++
	}
	stage.XLFiles_mapString[xlfile.Name] = xlfile

	return xlfile
}

// Unstage removes xlfile off the model stage
func (xlfile *XLFile) Unstage(stage *Stage) *XLFile {
	delete(stage.XLFiles, xlfile)
	delete(stage.XLFiles_mapString, xlfile.Name)
	return xlfile
}

// UnstageVoid removes xlfile off the model stage
func (xlfile *XLFile) UnstageVoid(stage *Stage) {
	delete(stage.XLFiles, xlfile)
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

// Stage puts xlrow to the model stage
func (xlrow *XLRow) Stage(stage *Stage) *XLRow {

	if _, ok := stage.XLRows[xlrow]; !ok {
		stage.XLRows[xlrow] = __member
		stage.XLRowMap_Staged_Order[xlrow] = stage.XLRowOrder
		stage.XLRowOrder++
	}
	stage.XLRows_mapString[xlrow.Name] = xlrow

	return xlrow
}

// Unstage removes xlrow off the model stage
func (xlrow *XLRow) Unstage(stage *Stage) *XLRow {
	delete(stage.XLRows, xlrow)
	delete(stage.XLRows_mapString, xlrow.Name)
	return xlrow
}

// UnstageVoid removes xlrow off the model stage
func (xlrow *XLRow) UnstageVoid(stage *Stage) {
	delete(stage.XLRows, xlrow)
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

// Stage puts xlsheet to the model stage
func (xlsheet *XLSheet) Stage(stage *Stage) *XLSheet {

	if _, ok := stage.XLSheets[xlsheet]; !ok {
		stage.XLSheets[xlsheet] = __member
		stage.XLSheetMap_Staged_Order[xlsheet] = stage.XLSheetOrder
		stage.XLSheetOrder++
	}
	stage.XLSheets_mapString[xlsheet.Name] = xlsheet

	return xlsheet
}

// Unstage removes xlsheet off the model stage
func (xlsheet *XLSheet) Unstage(stage *Stage) *XLSheet {
	delete(stage.XLSheets, xlsheet)
	delete(stage.XLSheets_mapString, xlsheet.Name)
	return xlsheet
}

// UnstageVoid removes xlsheet off the model stage
func (xlsheet *XLSheet) UnstageVoid(stage *Stage) {
	delete(stage.XLSheets, xlsheet)
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
	stage.DisplaySelections = make(map[*DisplaySelection]any)
	stage.DisplaySelections_mapString = make(map[string]*DisplaySelection)
	stage.DisplaySelectionMap_Staged_Order = make(map[*DisplaySelection]uint)
	stage.DisplaySelectionOrder = 0

	stage.XLCells = make(map[*XLCell]any)
	stage.XLCells_mapString = make(map[string]*XLCell)
	stage.XLCellMap_Staged_Order = make(map[*XLCell]uint)
	stage.XLCellOrder = 0

	stage.XLFiles = make(map[*XLFile]any)
	stage.XLFiles_mapString = make(map[string]*XLFile)
	stage.XLFileMap_Staged_Order = make(map[*XLFile]uint)
	stage.XLFileOrder = 0

	stage.XLRows = make(map[*XLRow]any)
	stage.XLRows_mapString = make(map[string]*XLRow)
	stage.XLRowMap_Staged_Order = make(map[*XLRow]uint)
	stage.XLRowOrder = 0

	stage.XLSheets = make(map[*XLSheet]any)
	stage.XLSheets_mapString = make(map[string]*XLSheet)
	stage.XLSheetMap_Staged_Order = make(map[*XLSheet]uint)
	stage.XLSheetOrder = 0

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

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *Stage) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*DisplaySelection:
		return any(&stage.DisplaySelections_mapString).(*Type)
	case map[string]*XLCell:
		return any(&stage.XLCells_mapString).(*Type)
	case map[string]*XLFile:
		return any(&stage.XLFiles_mapString).(*Type)
	case map[string]*XLRow:
		return any(&stage.XLRows_mapString).(*Type)
	case map[string]*XLSheet:
		return any(&stage.XLSheets_mapString).(*Type)
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
	case DisplaySelection:
		return any(&stage.DisplaySelections).(*map[*Type]any)
	case XLCell:
		return any(&stage.XLCells).(*map[*Type]any)
	case XLFile:
		return any(&stage.XLFiles).(*map[*Type]any)
	case XLRow:
		return any(&stage.XLRows).(*map[*Type]any)
	case XLSheet:
		return any(&stage.XLSheets).(*map[*Type]any)
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
	case *DisplaySelection:
		return any(&stage.DisplaySelections).(*map[Type]any)
	case *XLCell:
		return any(&stage.XLCells).(*map[Type]any)
	case *XLFile:
		return any(&stage.XLFiles).(*map[Type]any)
	case *XLRow:
		return any(&stage.XLRows).(*map[Type]any)
	case *XLSheet:
		return any(&stage.XLSheets).(*map[Type]any)
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

// GetGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetGongstructName[Type Gongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case DisplaySelection:
		res = "DisplaySelection"
	case XLCell:
		res = "XLCell"
	case XLFile:
		res = "XLFile"
	case XLRow:
		res = "XLRow"
	case XLSheet:
		res = "XLSheet"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

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

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case DisplaySelection:
		res = []string{"Name", "XLFile", "XLSheet"}
	case XLCell:
		res = []string{"Name", "X", "Y"}
	case XLFile:
		res = []string{"Name", "NbSheets", "Sheets"}
	case XLRow:
		res = []string{"Name", "RowIndex", "Cells"}
	case XLSheet:
		res = []string{"Name", "MaxRow", "MaxCol", "NbRows", "Rows", "SheetCells"}
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
	case DisplaySelection:
		var rf ReverseField
		_ = rf
	case XLCell:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "XLRow"
		rf.Fieldname = "Cells"
		res = append(res, rf)
		rf.GongstructName = "XLSheet"
		rf.Fieldname = "SheetCells"
		res = append(res, rf)
	case XLFile:
		var rf ReverseField
		_ = rf
	case XLRow:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "XLSheet"
		rf.Fieldname = "Rows"
		res = append(res, rf)
	case XLSheet:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "XLFile"
		rf.Fieldname = "Sheets"
		res = append(res, rf)
	}
	return
}

// GetFieldsFromPointer return the array of the fields
func GetFieldsFromPointer[Type PointerToGongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case *DisplaySelection:
		res = []string{"Name", "XLFile", "XLSheet"}
	case *XLCell:
		res = []string{"Name", "X", "Y"}
	case *XLFile:
		res = []string{"Name", "NbSheets", "Sheets"}
	case *XLRow:
		res = []string{"Name", "RowIndex", "Cells"}
	case *XLSheet:
		res = []string{"Name", "MaxRow", "MaxCol", "NbRows", "Rows", "SheetCells"}
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
	case *DisplaySelection:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "XLFile":
			if inferedInstance.XLFile != nil {
				res.valueString = inferedInstance.XLFile.Name
			}
		case "XLSheet":
			if inferedInstance.XLSheet != nil {
				res.valueString = inferedInstance.XLSheet.Name
			}
		}
	case *XLCell:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%d", inferedInstance.X)
			res.valueInt = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Y":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Y)
			res.valueInt = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case *XLFile:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "NbSheets":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbSheets)
			res.valueInt = inferedInstance.NbSheets
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Sheets":
			for idx, __instance__ := range inferedInstance.Sheets {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *XLRow:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "RowIndex":
			res.valueString = fmt.Sprintf("%d", inferedInstance.RowIndex)
			res.valueInt = inferedInstance.RowIndex
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Cells":
			for idx, __instance__ := range inferedInstance.Cells {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case *XLSheet:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MaxRow":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MaxRow)
			res.valueInt = inferedInstance.MaxRow
			res.GongFieldValueType = GongFieldValueTypeInt
		case "MaxCol":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MaxCol)
			res.valueInt = inferedInstance.MaxCol
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NbRows":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbRows)
			res.valueInt = inferedInstance.NbRows
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Rows":
			for idx, __instance__ := range inferedInstance.Rows {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SheetCells":
			for idx, __instance__ := range inferedInstance.SheetCells {
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
	case DisplaySelection:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "XLFile":
			if inferedInstance.XLFile != nil {
				res.valueString = inferedInstance.XLFile.Name
			}
		case "XLSheet":
			if inferedInstance.XLSheet != nil {
				res.valueString = inferedInstance.XLSheet.Name
			}
		}
	case XLCell:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "X":
			res.valueString = fmt.Sprintf("%d", inferedInstance.X)
			res.valueInt = inferedInstance.X
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Y":
			res.valueString = fmt.Sprintf("%d", inferedInstance.Y)
			res.valueInt = inferedInstance.Y
			res.GongFieldValueType = GongFieldValueTypeInt
		}
	case XLFile:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "NbSheets":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbSheets)
			res.valueInt = inferedInstance.NbSheets
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Sheets":
			for idx, __instance__ := range inferedInstance.Sheets {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case XLRow:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "RowIndex":
			res.valueString = fmt.Sprintf("%d", inferedInstance.RowIndex)
			res.valueInt = inferedInstance.RowIndex
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Cells":
			for idx, __instance__ := range inferedInstance.Cells {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		}
	case XLSheet:
		switch fieldName {
		// string value of fields
		case "Name":
			res.valueString = inferedInstance.Name
		case "MaxRow":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MaxRow)
			res.valueInt = inferedInstance.MaxRow
			res.GongFieldValueType = GongFieldValueTypeInt
		case "MaxCol":
			res.valueString = fmt.Sprintf("%d", inferedInstance.MaxCol)
			res.valueInt = inferedInstance.MaxCol
			res.GongFieldValueType = GongFieldValueTypeInt
		case "NbRows":
			res.valueString = fmt.Sprintf("%d", inferedInstance.NbRows)
			res.valueInt = inferedInstance.NbRows
			res.GongFieldValueType = GongFieldValueTypeInt
		case "Rows":
			for idx, __instance__ := range inferedInstance.Rows {
				if idx > 0 {
					res.valueString += "\n"
				}
				res.valueString += __instance__.Name
			}
		case "SheetCells":
			for idx, __instance__ := range inferedInstance.SheetCells {
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
