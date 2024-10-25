// generated code - do not edit
package models

import (
	"cmp"
	"errors"
	"fmt"
	"math"
	"slices"
	"time"
)

func __Gong__Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// errUnkownEnum is returns when a value cannot match enum values
var errUnkownEnum = errors.New("unkown enum")

// needed to avoid when fmt package is not needed by generated code
var __dummy__fmt_variable fmt.Scanner

// idem for math package when not need by generated code
var __dummy_math_variable = math.E

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

// StageStruct enables storage of staged instances
// swagger:ignore
type StageStruct struct {
	path string

	// insertion point for definition of arrays registering instances
	Cells           map[*Cell]any
	Cells_mapString map[string]*Cell

	// insertion point for slice of pointers maps
	OnAfterCellCreateCallback OnAfterCreateInterface[Cell]
	OnAfterCellUpdateCallback OnAfterUpdateInterface[Cell]
	OnAfterCellDeleteCallback OnAfterDeleteInterface[Cell]
	OnAfterCellReadCallback   OnAfterReadInterface[Cell]

	CellBooleans           map[*CellBoolean]any
	CellBooleans_mapString map[string]*CellBoolean

	// insertion point for slice of pointers maps
	OnAfterCellBooleanCreateCallback OnAfterCreateInterface[CellBoolean]
	OnAfterCellBooleanUpdateCallback OnAfterUpdateInterface[CellBoolean]
	OnAfterCellBooleanDeleteCallback OnAfterDeleteInterface[CellBoolean]
	OnAfterCellBooleanReadCallback   OnAfterReadInterface[CellBoolean]

	CellFloat64s           map[*CellFloat64]any
	CellFloat64s_mapString map[string]*CellFloat64

	// insertion point for slice of pointers maps
	OnAfterCellFloat64CreateCallback OnAfterCreateInterface[CellFloat64]
	OnAfterCellFloat64UpdateCallback OnAfterUpdateInterface[CellFloat64]
	OnAfterCellFloat64DeleteCallback OnAfterDeleteInterface[CellFloat64]
	OnAfterCellFloat64ReadCallback   OnAfterReadInterface[CellFloat64]

	CellIcons           map[*CellIcon]any
	CellIcons_mapString map[string]*CellIcon

	// insertion point for slice of pointers maps
	OnAfterCellIconCreateCallback OnAfterCreateInterface[CellIcon]
	OnAfterCellIconUpdateCallback OnAfterUpdateInterface[CellIcon]
	OnAfterCellIconDeleteCallback OnAfterDeleteInterface[CellIcon]
	OnAfterCellIconReadCallback   OnAfterReadInterface[CellIcon]

	CellInts           map[*CellInt]any
	CellInts_mapString map[string]*CellInt

	// insertion point for slice of pointers maps
	OnAfterCellIntCreateCallback OnAfterCreateInterface[CellInt]
	OnAfterCellIntUpdateCallback OnAfterUpdateInterface[CellInt]
	OnAfterCellIntDeleteCallback OnAfterDeleteInterface[CellInt]
	OnAfterCellIntReadCallback   OnAfterReadInterface[CellInt]

	CellStrings           map[*CellString]any
	CellStrings_mapString map[string]*CellString

	// insertion point for slice of pointers maps
	OnAfterCellStringCreateCallback OnAfterCreateInterface[CellString]
	OnAfterCellStringUpdateCallback OnAfterUpdateInterface[CellString]
	OnAfterCellStringDeleteCallback OnAfterDeleteInterface[CellString]
	OnAfterCellStringReadCallback   OnAfterReadInterface[CellString]

	CheckBoxs           map[*CheckBox]any
	CheckBoxs_mapString map[string]*CheckBox

	// insertion point for slice of pointers maps
	OnAfterCheckBoxCreateCallback OnAfterCreateInterface[CheckBox]
	OnAfterCheckBoxUpdateCallback OnAfterUpdateInterface[CheckBox]
	OnAfterCheckBoxDeleteCallback OnAfterDeleteInterface[CheckBox]
	OnAfterCheckBoxReadCallback   OnAfterReadInterface[CheckBox]

	DisplayedColumns           map[*DisplayedColumn]any
	DisplayedColumns_mapString map[string]*DisplayedColumn

	// insertion point for slice of pointers maps
	OnAfterDisplayedColumnCreateCallback OnAfterCreateInterface[DisplayedColumn]
	OnAfterDisplayedColumnUpdateCallback OnAfterUpdateInterface[DisplayedColumn]
	OnAfterDisplayedColumnDeleteCallback OnAfterDeleteInterface[DisplayedColumn]
	OnAfterDisplayedColumnReadCallback   OnAfterReadInterface[DisplayedColumn]

	FormDivs           map[*FormDiv]any
	FormDivs_mapString map[string]*FormDiv

	// insertion point for slice of pointers maps
	FormDiv_FormFields_reverseMap map[*FormField]*FormDiv

	FormDiv_CheckBoxs_reverseMap map[*CheckBox]*FormDiv

	OnAfterFormDivCreateCallback OnAfterCreateInterface[FormDiv]
	OnAfterFormDivUpdateCallback OnAfterUpdateInterface[FormDiv]
	OnAfterFormDivDeleteCallback OnAfterDeleteInterface[FormDiv]
	OnAfterFormDivReadCallback   OnAfterReadInterface[FormDiv]

	FormEditAssocButtons           map[*FormEditAssocButton]any
	FormEditAssocButtons_mapString map[string]*FormEditAssocButton

	// insertion point for slice of pointers maps
	OnAfterFormEditAssocButtonCreateCallback OnAfterCreateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonUpdateCallback OnAfterUpdateInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonDeleteCallback OnAfterDeleteInterface[FormEditAssocButton]
	OnAfterFormEditAssocButtonReadCallback   OnAfterReadInterface[FormEditAssocButton]

	FormFields           map[*FormField]any
	FormFields_mapString map[string]*FormField

	// insertion point for slice of pointers maps
	OnAfterFormFieldCreateCallback OnAfterCreateInterface[FormField]
	OnAfterFormFieldUpdateCallback OnAfterUpdateInterface[FormField]
	OnAfterFormFieldDeleteCallback OnAfterDeleteInterface[FormField]
	OnAfterFormFieldReadCallback   OnAfterReadInterface[FormField]

	FormFieldDates           map[*FormFieldDate]any
	FormFieldDates_mapString map[string]*FormFieldDate

	// insertion point for slice of pointers maps
	OnAfterFormFieldDateCreateCallback OnAfterCreateInterface[FormFieldDate]
	OnAfterFormFieldDateUpdateCallback OnAfterUpdateInterface[FormFieldDate]
	OnAfterFormFieldDateDeleteCallback OnAfterDeleteInterface[FormFieldDate]
	OnAfterFormFieldDateReadCallback   OnAfterReadInterface[FormFieldDate]

	FormFieldDateTimes           map[*FormFieldDateTime]any
	FormFieldDateTimes_mapString map[string]*FormFieldDateTime

	// insertion point for slice of pointers maps
	OnAfterFormFieldDateTimeCreateCallback OnAfterCreateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeUpdateCallback OnAfterUpdateInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeDeleteCallback OnAfterDeleteInterface[FormFieldDateTime]
	OnAfterFormFieldDateTimeReadCallback   OnAfterReadInterface[FormFieldDateTime]

	FormFieldFloat64s           map[*FormFieldFloat64]any
	FormFieldFloat64s_mapString map[string]*FormFieldFloat64

	// insertion point for slice of pointers maps
	OnAfterFormFieldFloat64CreateCallback OnAfterCreateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64UpdateCallback OnAfterUpdateInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64DeleteCallback OnAfterDeleteInterface[FormFieldFloat64]
	OnAfterFormFieldFloat64ReadCallback   OnAfterReadInterface[FormFieldFloat64]

	FormFieldInts           map[*FormFieldInt]any
	FormFieldInts_mapString map[string]*FormFieldInt

	// insertion point for slice of pointers maps
	OnAfterFormFieldIntCreateCallback OnAfterCreateInterface[FormFieldInt]
	OnAfterFormFieldIntUpdateCallback OnAfterUpdateInterface[FormFieldInt]
	OnAfterFormFieldIntDeleteCallback OnAfterDeleteInterface[FormFieldInt]
	OnAfterFormFieldIntReadCallback   OnAfterReadInterface[FormFieldInt]

	FormFieldSelects           map[*FormFieldSelect]any
	FormFieldSelects_mapString map[string]*FormFieldSelect

	// insertion point for slice of pointers maps
	FormFieldSelect_Options_reverseMap map[*Option]*FormFieldSelect

	OnAfterFormFieldSelectCreateCallback OnAfterCreateInterface[FormFieldSelect]
	OnAfterFormFieldSelectUpdateCallback OnAfterUpdateInterface[FormFieldSelect]
	OnAfterFormFieldSelectDeleteCallback OnAfterDeleteInterface[FormFieldSelect]
	OnAfterFormFieldSelectReadCallback   OnAfterReadInterface[FormFieldSelect]

	FormFieldStrings           map[*FormFieldString]any
	FormFieldStrings_mapString map[string]*FormFieldString

	// insertion point for slice of pointers maps
	OnAfterFormFieldStringCreateCallback OnAfterCreateInterface[FormFieldString]
	OnAfterFormFieldStringUpdateCallback OnAfterUpdateInterface[FormFieldString]
	OnAfterFormFieldStringDeleteCallback OnAfterDeleteInterface[FormFieldString]
	OnAfterFormFieldStringReadCallback   OnAfterReadInterface[FormFieldString]

	FormFieldTimes           map[*FormFieldTime]any
	FormFieldTimes_mapString map[string]*FormFieldTime

	// insertion point for slice of pointers maps
	OnAfterFormFieldTimeCreateCallback OnAfterCreateInterface[FormFieldTime]
	OnAfterFormFieldTimeUpdateCallback OnAfterUpdateInterface[FormFieldTime]
	OnAfterFormFieldTimeDeleteCallback OnAfterDeleteInterface[FormFieldTime]
	OnAfterFormFieldTimeReadCallback   OnAfterReadInterface[FormFieldTime]

	FormGroups           map[*FormGroup]any
	FormGroups_mapString map[string]*FormGroup

	// insertion point for slice of pointers maps
	FormGroup_FormDivs_reverseMap map[*FormDiv]*FormGroup

	OnAfterFormGroupCreateCallback OnAfterCreateInterface[FormGroup]
	OnAfterFormGroupUpdateCallback OnAfterUpdateInterface[FormGroup]
	OnAfterFormGroupDeleteCallback OnAfterDeleteInterface[FormGroup]
	OnAfterFormGroupReadCallback   OnAfterReadInterface[FormGroup]

	FormSortAssocButtons           map[*FormSortAssocButton]any
	FormSortAssocButtons_mapString map[string]*FormSortAssocButton

	// insertion point for slice of pointers maps
	OnAfterFormSortAssocButtonCreateCallback OnAfterCreateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonUpdateCallback OnAfterUpdateInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonDeleteCallback OnAfterDeleteInterface[FormSortAssocButton]
	OnAfterFormSortAssocButtonReadCallback   OnAfterReadInterface[FormSortAssocButton]

	Options           map[*Option]any
	Options_mapString map[string]*Option

	// insertion point for slice of pointers maps
	OnAfterOptionCreateCallback OnAfterCreateInterface[Option]
	OnAfterOptionUpdateCallback OnAfterUpdateInterface[Option]
	OnAfterOptionDeleteCallback OnAfterDeleteInterface[Option]
	OnAfterOptionReadCallback   OnAfterReadInterface[Option]

	Rows           map[*Row]any
	Rows_mapString map[string]*Row

	// insertion point for slice of pointers maps
	Row_Cells_reverseMap map[*Cell]*Row

	OnAfterRowCreateCallback OnAfterCreateInterface[Row]
	OnAfterRowUpdateCallback OnAfterUpdateInterface[Row]
	OnAfterRowDeleteCallback OnAfterDeleteInterface[Row]
	OnAfterRowReadCallback   OnAfterReadInterface[Row]

	Tables           map[*Table]any
	Tables_mapString map[string]*Table

	// insertion point for slice of pointers maps
	Table_DisplayedColumns_reverseMap map[*DisplayedColumn]*Table

	Table_Rows_reverseMap map[*Row]*Table

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

	// store the number of instance per gongstruct
	Map_GongStructName_InstancesNb map[string]int

	// store meta package import
	MetaPackageImportPath  string
	MetaPackageImportAlias string

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	// map to enable docLink renaming when an identifier is renamed
	Map_DocLink_Renaming map[string]GONG__Identifier
	// the to be removed stops here
}

func (stage *StageStruct) GetType() string {
	return "github.com/fullstack-lang/gongtable/go/models"
}

type GONG__Identifier struct {
	Ident string
	Type  GONG__ExpressionType
}

type OnInitCommitInterface interface {
	BeforeCommit(stage *StageStruct)
}

// OnAfterCreateInterface callback when an instance is updated from the front
type OnAfterCreateInterface[Type Gongstruct] interface {
	OnAfterCreate(stage *StageStruct,
		instance *Type)
}

// OnAfterReadInterface callback when an instance is updated from the front
type OnAfterReadInterface[Type Gongstruct] interface {
	OnAfterRead(stage *StageStruct,
		instance *Type)
}

// OnAfterUpdateInterface callback when an instance is updated from the front
type OnAfterUpdateInterface[Type Gongstruct] interface {
	OnAfterUpdate(stage *StageStruct, old, new *Type)
}

// OnAfterDeleteInterface callback when an instance is updated from the front
type OnAfterDeleteInterface[Type Gongstruct] interface {
	OnAfterDelete(stage *StageStruct,
		staged, front *Type)
}

type BackRepoInterface interface {
	Commit(stage *StageStruct)
	Checkout(stage *StageStruct)
	Backup(stage *StageStruct, dirPath string)
	Restore(stage *StageStruct, dirPath string)
	BackupXL(stage *StageStruct, dirPath string)
	RestoreXL(stage *StageStruct, dirPath string)
	// insertion point for Commit and Checkout signatures
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
	CommitCheckBox(checkbox *CheckBox)
	CheckoutCheckBox(checkbox *CheckBox)
	CommitDisplayedColumn(displayedcolumn *DisplayedColumn)
	CheckoutDisplayedColumn(displayedcolumn *DisplayedColumn)
	CommitFormDiv(formdiv *FormDiv)
	CheckoutFormDiv(formdiv *FormDiv)
	CommitFormEditAssocButton(formeditassocbutton *FormEditAssocButton)
	CheckoutFormEditAssocButton(formeditassocbutton *FormEditAssocButton)
	CommitFormField(formfield *FormField)
	CheckoutFormField(formfield *FormField)
	CommitFormFieldDate(formfielddate *FormFieldDate)
	CheckoutFormFieldDate(formfielddate *FormFieldDate)
	CommitFormFieldDateTime(formfielddatetime *FormFieldDateTime)
	CheckoutFormFieldDateTime(formfielddatetime *FormFieldDateTime)
	CommitFormFieldFloat64(formfieldfloat64 *FormFieldFloat64)
	CheckoutFormFieldFloat64(formfieldfloat64 *FormFieldFloat64)
	CommitFormFieldInt(formfieldint *FormFieldInt)
	CheckoutFormFieldInt(formfieldint *FormFieldInt)
	CommitFormFieldSelect(formfieldselect *FormFieldSelect)
	CheckoutFormFieldSelect(formfieldselect *FormFieldSelect)
	CommitFormFieldString(formfieldstring *FormFieldString)
	CheckoutFormFieldString(formfieldstring *FormFieldString)
	CommitFormFieldTime(formfieldtime *FormFieldTime)
	CheckoutFormFieldTime(formfieldtime *FormFieldTime)
	CommitFormGroup(formgroup *FormGroup)
	CheckoutFormGroup(formgroup *FormGroup)
	CommitFormSortAssocButton(formsortassocbutton *FormSortAssocButton)
	CheckoutFormSortAssocButton(formsortassocbutton *FormSortAssocButton)
	CommitOption(option *Option)
	CheckoutOption(option *Option)
	CommitRow(row *Row)
	CheckoutRow(row *Row)
	CommitTable(table *Table)
	CheckoutTable(table *Table)
	GetLastCommitFromBackNb() uint
	GetLastPushFromFrontNb() uint
}

func NewStage(path string) (stage *StageStruct) {

	stage = &StageStruct{ // insertion point for array initiatialisation
		Cells:           make(map[*Cell]any),
		Cells_mapString: make(map[string]*Cell),

		CellBooleans:           make(map[*CellBoolean]any),
		CellBooleans_mapString: make(map[string]*CellBoolean),

		CellFloat64s:           make(map[*CellFloat64]any),
		CellFloat64s_mapString: make(map[string]*CellFloat64),

		CellIcons:           make(map[*CellIcon]any),
		CellIcons_mapString: make(map[string]*CellIcon),

		CellInts:           make(map[*CellInt]any),
		CellInts_mapString: make(map[string]*CellInt),

		CellStrings:           make(map[*CellString]any),
		CellStrings_mapString: make(map[string]*CellString),

		CheckBoxs:           make(map[*CheckBox]any),
		CheckBoxs_mapString: make(map[string]*CheckBox),

		DisplayedColumns:           make(map[*DisplayedColumn]any),
		DisplayedColumns_mapString: make(map[string]*DisplayedColumn),

		FormDivs:           make(map[*FormDiv]any),
		FormDivs_mapString: make(map[string]*FormDiv),

		FormEditAssocButtons:           make(map[*FormEditAssocButton]any),
		FormEditAssocButtons_mapString: make(map[string]*FormEditAssocButton),

		FormFields:           make(map[*FormField]any),
		FormFields_mapString: make(map[string]*FormField),

		FormFieldDates:           make(map[*FormFieldDate]any),
		FormFieldDates_mapString: make(map[string]*FormFieldDate),

		FormFieldDateTimes:           make(map[*FormFieldDateTime]any),
		FormFieldDateTimes_mapString: make(map[string]*FormFieldDateTime),

		FormFieldFloat64s:           make(map[*FormFieldFloat64]any),
		FormFieldFloat64s_mapString: make(map[string]*FormFieldFloat64),

		FormFieldInts:           make(map[*FormFieldInt]any),
		FormFieldInts_mapString: make(map[string]*FormFieldInt),

		FormFieldSelects:           make(map[*FormFieldSelect]any),
		FormFieldSelects_mapString: make(map[string]*FormFieldSelect),

		FormFieldStrings:           make(map[*FormFieldString]any),
		FormFieldStrings_mapString: make(map[string]*FormFieldString),

		FormFieldTimes:           make(map[*FormFieldTime]any),
		FormFieldTimes_mapString: make(map[string]*FormFieldTime),

		FormGroups:           make(map[*FormGroup]any),
		FormGroups_mapString: make(map[string]*FormGroup),

		FormSortAssocButtons:           make(map[*FormSortAssocButton]any),
		FormSortAssocButtons_mapString: make(map[string]*FormSortAssocButton),

		Options:           make(map[*Option]any),
		Options_mapString: make(map[string]*Option),

		Rows:           make(map[*Row]any),
		Rows_mapString: make(map[string]*Row),

		Tables:           make(map[*Table]any),
		Tables_mapString: make(map[string]*Table),

		// end of insertion point
		Map_GongStructName_InstancesNb: make(map[string]int),

		path: path,

		// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
		Map_DocLink_Renaming: make(map[string]GONG__Identifier),
		// the to be removed stops here
	}

	return
}

func (stage *StageStruct) GetPath() string {
	return stage.path
}

func (stage *StageStruct) CommitWithSuspendedCallbacks() {

	tmp := stage.OnInitCommitFromBackCallback
	stage.OnInitCommitFromBackCallback = nil
	stage.Commit()
	stage.OnInitCommitFromBackCallback = tmp
}

func (stage *StageStruct) Commit() {
	stage.ComputeReverseMaps()

	if stage.BackRepo != nil {
		stage.BackRepo.Commit(stage)
	}

	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Cell"] = len(stage.Cells)
	stage.Map_GongStructName_InstancesNb["CellBoolean"] = len(stage.CellBooleans)
	stage.Map_GongStructName_InstancesNb["CellFloat64"] = len(stage.CellFloat64s)
	stage.Map_GongStructName_InstancesNb["CellIcon"] = len(stage.CellIcons)
	stage.Map_GongStructName_InstancesNb["CellInt"] = len(stage.CellInts)
	stage.Map_GongStructName_InstancesNb["CellString"] = len(stage.CellStrings)
	stage.Map_GongStructName_InstancesNb["CheckBox"] = len(stage.CheckBoxs)
	stage.Map_GongStructName_InstancesNb["DisplayedColumn"] = len(stage.DisplayedColumns)
	stage.Map_GongStructName_InstancesNb["FormDiv"] = len(stage.FormDivs)
	stage.Map_GongStructName_InstancesNb["FormEditAssocButton"] = len(stage.FormEditAssocButtons)
	stage.Map_GongStructName_InstancesNb["FormField"] = len(stage.FormFields)
	stage.Map_GongStructName_InstancesNb["FormFieldDate"] = len(stage.FormFieldDates)
	stage.Map_GongStructName_InstancesNb["FormFieldDateTime"] = len(stage.FormFieldDateTimes)
	stage.Map_GongStructName_InstancesNb["FormFieldFloat64"] = len(stage.FormFieldFloat64s)
	stage.Map_GongStructName_InstancesNb["FormFieldInt"] = len(stage.FormFieldInts)
	stage.Map_GongStructName_InstancesNb["FormFieldSelect"] = len(stage.FormFieldSelects)
	stage.Map_GongStructName_InstancesNb["FormFieldString"] = len(stage.FormFieldStrings)
	stage.Map_GongStructName_InstancesNb["FormFieldTime"] = len(stage.FormFieldTimes)
	stage.Map_GongStructName_InstancesNb["FormGroup"] = len(stage.FormGroups)
	stage.Map_GongStructName_InstancesNb["FormSortAssocButton"] = len(stage.FormSortAssocButtons)
	stage.Map_GongStructName_InstancesNb["Option"] = len(stage.Options)
	stage.Map_GongStructName_InstancesNb["Row"] = len(stage.Rows)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)

}

func (stage *StageStruct) Checkout() {
	if stage.BackRepo != nil {
		stage.BackRepo.Checkout(stage)
	}

	stage.ComputeReverseMaps()
	// insertion point for computing the map of number of instances per gongstruct
	stage.Map_GongStructName_InstancesNb["Cell"] = len(stage.Cells)
	stage.Map_GongStructName_InstancesNb["CellBoolean"] = len(stage.CellBooleans)
	stage.Map_GongStructName_InstancesNb["CellFloat64"] = len(stage.CellFloat64s)
	stage.Map_GongStructName_InstancesNb["CellIcon"] = len(stage.CellIcons)
	stage.Map_GongStructName_InstancesNb["CellInt"] = len(stage.CellInts)
	stage.Map_GongStructName_InstancesNb["CellString"] = len(stage.CellStrings)
	stage.Map_GongStructName_InstancesNb["CheckBox"] = len(stage.CheckBoxs)
	stage.Map_GongStructName_InstancesNb["DisplayedColumn"] = len(stage.DisplayedColumns)
	stage.Map_GongStructName_InstancesNb["FormDiv"] = len(stage.FormDivs)
	stage.Map_GongStructName_InstancesNb["FormEditAssocButton"] = len(stage.FormEditAssocButtons)
	stage.Map_GongStructName_InstancesNb["FormField"] = len(stage.FormFields)
	stage.Map_GongStructName_InstancesNb["FormFieldDate"] = len(stage.FormFieldDates)
	stage.Map_GongStructName_InstancesNb["FormFieldDateTime"] = len(stage.FormFieldDateTimes)
	stage.Map_GongStructName_InstancesNb["FormFieldFloat64"] = len(stage.FormFieldFloat64s)
	stage.Map_GongStructName_InstancesNb["FormFieldInt"] = len(stage.FormFieldInts)
	stage.Map_GongStructName_InstancesNb["FormFieldSelect"] = len(stage.FormFieldSelects)
	stage.Map_GongStructName_InstancesNb["FormFieldString"] = len(stage.FormFieldStrings)
	stage.Map_GongStructName_InstancesNb["FormFieldTime"] = len(stage.FormFieldTimes)
	stage.Map_GongStructName_InstancesNb["FormGroup"] = len(stage.FormGroups)
	stage.Map_GongStructName_InstancesNb["FormSortAssocButton"] = len(stage.FormSortAssocButtons)
	stage.Map_GongStructName_InstancesNb["Option"] = len(stage.Options)
	stage.Map_GongStructName_InstancesNb["Row"] = len(stage.Rows)
	stage.Map_GongStructName_InstancesNb["Table"] = len(stage.Tables)

}

// backup generates backup files in the dirPath
func (stage *StageStruct) Backup(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Backup(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) Restore(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.Restore(stage, dirPath)
	}
}

// backup generates backup files in the dirPath
func (stage *StageStruct) BackupXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.BackupXL(stage, dirPath)
	}
}

// Restore resets Stage & BackRepo and restores their content from the restore files in dirPath
func (stage *StageStruct) RestoreXL(dirPath string) {
	if stage.BackRepo != nil {
		stage.BackRepo.RestoreXL(stage, dirPath)
	}
}

// insertion point for cumulative sub template with model space calls
// Stage puts cell to the model stage
func (cell *Cell) Stage(stage *StageStruct) *Cell {
	stage.Cells[cell] = __member
	stage.Cells_mapString[cell.Name] = cell

	return cell
}

// Unstage removes cell off the model stage
func (cell *Cell) Unstage(stage *StageStruct) *Cell {
	delete(stage.Cells, cell)
	delete(stage.Cells_mapString, cell.Name)
	return cell
}

// UnstageVoid removes cell off the model stage
func (cell *Cell) UnstageVoid(stage *StageStruct) {
	delete(stage.Cells, cell)
	delete(stage.Cells_mapString, cell.Name)
}

// commit cell to the back repo (if it is already staged)
func (cell *Cell) Commit(stage *StageStruct) *Cell {
	if _, ok := stage.Cells[cell]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCell(cell)
		}
	}
	return cell
}

func (cell *Cell) CommitVoid(stage *StageStruct) {
	cell.Commit(stage)
}

// Checkout cell to the back repo (if it is already staged)
func (cell *Cell) Checkout(stage *StageStruct) *Cell {
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

// Stage puts cellboolean to the model stage
func (cellboolean *CellBoolean) Stage(stage *StageStruct) *CellBoolean {
	stage.CellBooleans[cellboolean] = __member
	stage.CellBooleans_mapString[cellboolean.Name] = cellboolean

	return cellboolean
}

// Unstage removes cellboolean off the model stage
func (cellboolean *CellBoolean) Unstage(stage *StageStruct) *CellBoolean {
	delete(stage.CellBooleans, cellboolean)
	delete(stage.CellBooleans_mapString, cellboolean.Name)
	return cellboolean
}

// UnstageVoid removes cellboolean off the model stage
func (cellboolean *CellBoolean) UnstageVoid(stage *StageStruct) {
	delete(stage.CellBooleans, cellboolean)
	delete(stage.CellBooleans_mapString, cellboolean.Name)
}

// commit cellboolean to the back repo (if it is already staged)
func (cellboolean *CellBoolean) Commit(stage *StageStruct) *CellBoolean {
	if _, ok := stage.CellBooleans[cellboolean]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellBoolean(cellboolean)
		}
	}
	return cellboolean
}

func (cellboolean *CellBoolean) CommitVoid(stage *StageStruct) {
	cellboolean.Commit(stage)
}

// Checkout cellboolean to the back repo (if it is already staged)
func (cellboolean *CellBoolean) Checkout(stage *StageStruct) *CellBoolean {
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

// Stage puts cellfloat64 to the model stage
func (cellfloat64 *CellFloat64) Stage(stage *StageStruct) *CellFloat64 {
	stage.CellFloat64s[cellfloat64] = __member
	stage.CellFloat64s_mapString[cellfloat64.Name] = cellfloat64

	return cellfloat64
}

// Unstage removes cellfloat64 off the model stage
func (cellfloat64 *CellFloat64) Unstage(stage *StageStruct) *CellFloat64 {
	delete(stage.CellFloat64s, cellfloat64)
	delete(stage.CellFloat64s_mapString, cellfloat64.Name)
	return cellfloat64
}

// UnstageVoid removes cellfloat64 off the model stage
func (cellfloat64 *CellFloat64) UnstageVoid(stage *StageStruct) {
	delete(stage.CellFloat64s, cellfloat64)
	delete(stage.CellFloat64s_mapString, cellfloat64.Name)
}

// commit cellfloat64 to the back repo (if it is already staged)
func (cellfloat64 *CellFloat64) Commit(stage *StageStruct) *CellFloat64 {
	if _, ok := stage.CellFloat64s[cellfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellFloat64(cellfloat64)
		}
	}
	return cellfloat64
}

func (cellfloat64 *CellFloat64) CommitVoid(stage *StageStruct) {
	cellfloat64.Commit(stage)
}

// Checkout cellfloat64 to the back repo (if it is already staged)
func (cellfloat64 *CellFloat64) Checkout(stage *StageStruct) *CellFloat64 {
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

// Stage puts cellicon to the model stage
func (cellicon *CellIcon) Stage(stage *StageStruct) *CellIcon {
	stage.CellIcons[cellicon] = __member
	stage.CellIcons_mapString[cellicon.Name] = cellicon

	return cellicon
}

// Unstage removes cellicon off the model stage
func (cellicon *CellIcon) Unstage(stage *StageStruct) *CellIcon {
	delete(stage.CellIcons, cellicon)
	delete(stage.CellIcons_mapString, cellicon.Name)
	return cellicon
}

// UnstageVoid removes cellicon off the model stage
func (cellicon *CellIcon) UnstageVoid(stage *StageStruct) {
	delete(stage.CellIcons, cellicon)
	delete(stage.CellIcons_mapString, cellicon.Name)
}

// commit cellicon to the back repo (if it is already staged)
func (cellicon *CellIcon) Commit(stage *StageStruct) *CellIcon {
	if _, ok := stage.CellIcons[cellicon]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellIcon(cellicon)
		}
	}
	return cellicon
}

func (cellicon *CellIcon) CommitVoid(stage *StageStruct) {
	cellicon.Commit(stage)
}

// Checkout cellicon to the back repo (if it is already staged)
func (cellicon *CellIcon) Checkout(stage *StageStruct) *CellIcon {
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

// Stage puts cellint to the model stage
func (cellint *CellInt) Stage(stage *StageStruct) *CellInt {
	stage.CellInts[cellint] = __member
	stage.CellInts_mapString[cellint.Name] = cellint

	return cellint
}

// Unstage removes cellint off the model stage
func (cellint *CellInt) Unstage(stage *StageStruct) *CellInt {
	delete(stage.CellInts, cellint)
	delete(stage.CellInts_mapString, cellint.Name)
	return cellint
}

// UnstageVoid removes cellint off the model stage
func (cellint *CellInt) UnstageVoid(stage *StageStruct) {
	delete(stage.CellInts, cellint)
	delete(stage.CellInts_mapString, cellint.Name)
}

// commit cellint to the back repo (if it is already staged)
func (cellint *CellInt) Commit(stage *StageStruct) *CellInt {
	if _, ok := stage.CellInts[cellint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellInt(cellint)
		}
	}
	return cellint
}

func (cellint *CellInt) CommitVoid(stage *StageStruct) {
	cellint.Commit(stage)
}

// Checkout cellint to the back repo (if it is already staged)
func (cellint *CellInt) Checkout(stage *StageStruct) *CellInt {
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

// Stage puts cellstring to the model stage
func (cellstring *CellString) Stage(stage *StageStruct) *CellString {
	stage.CellStrings[cellstring] = __member
	stage.CellStrings_mapString[cellstring.Name] = cellstring

	return cellstring
}

// Unstage removes cellstring off the model stage
func (cellstring *CellString) Unstage(stage *StageStruct) *CellString {
	delete(stage.CellStrings, cellstring)
	delete(stage.CellStrings_mapString, cellstring.Name)
	return cellstring
}

// UnstageVoid removes cellstring off the model stage
func (cellstring *CellString) UnstageVoid(stage *StageStruct) {
	delete(stage.CellStrings, cellstring)
	delete(stage.CellStrings_mapString, cellstring.Name)
}

// commit cellstring to the back repo (if it is already staged)
func (cellstring *CellString) Commit(stage *StageStruct) *CellString {
	if _, ok := stage.CellStrings[cellstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCellString(cellstring)
		}
	}
	return cellstring
}

func (cellstring *CellString) CommitVoid(stage *StageStruct) {
	cellstring.Commit(stage)
}

// Checkout cellstring to the back repo (if it is already staged)
func (cellstring *CellString) Checkout(stage *StageStruct) *CellString {
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

// Stage puts checkbox to the model stage
func (checkbox *CheckBox) Stage(stage *StageStruct) *CheckBox {
	stage.CheckBoxs[checkbox] = __member
	stage.CheckBoxs_mapString[checkbox.Name] = checkbox

	return checkbox
}

// Unstage removes checkbox off the model stage
func (checkbox *CheckBox) Unstage(stage *StageStruct) *CheckBox {
	delete(stage.CheckBoxs, checkbox)
	delete(stage.CheckBoxs_mapString, checkbox.Name)
	return checkbox
}

// UnstageVoid removes checkbox off the model stage
func (checkbox *CheckBox) UnstageVoid(stage *StageStruct) {
	delete(stage.CheckBoxs, checkbox)
	delete(stage.CheckBoxs_mapString, checkbox.Name)
}

// commit checkbox to the back repo (if it is already staged)
func (checkbox *CheckBox) Commit(stage *StageStruct) *CheckBox {
	if _, ok := stage.CheckBoxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitCheckBox(checkbox)
		}
	}
	return checkbox
}

func (checkbox *CheckBox) CommitVoid(stage *StageStruct) {
	checkbox.Commit(stage)
}

// Checkout checkbox to the back repo (if it is already staged)
func (checkbox *CheckBox) Checkout(stage *StageStruct) *CheckBox {
	if _, ok := stage.CheckBoxs[checkbox]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutCheckBox(checkbox)
		}
	}
	return checkbox
}

// for satisfaction of GongStruct interface
func (checkbox *CheckBox) GetName() (res string) {
	return checkbox.Name
}

// Stage puts displayedcolumn to the model stage
func (displayedcolumn *DisplayedColumn) Stage(stage *StageStruct) *DisplayedColumn {
	stage.DisplayedColumns[displayedcolumn] = __member
	stage.DisplayedColumns_mapString[displayedcolumn.Name] = displayedcolumn

	return displayedcolumn
}

// Unstage removes displayedcolumn off the model stage
func (displayedcolumn *DisplayedColumn) Unstage(stage *StageStruct) *DisplayedColumn {
	delete(stage.DisplayedColumns, displayedcolumn)
	delete(stage.DisplayedColumns_mapString, displayedcolumn.Name)
	return displayedcolumn
}

// UnstageVoid removes displayedcolumn off the model stage
func (displayedcolumn *DisplayedColumn) UnstageVoid(stage *StageStruct) {
	delete(stage.DisplayedColumns, displayedcolumn)
	delete(stage.DisplayedColumns_mapString, displayedcolumn.Name)
}

// commit displayedcolumn to the back repo (if it is already staged)
func (displayedcolumn *DisplayedColumn) Commit(stage *StageStruct) *DisplayedColumn {
	if _, ok := stage.DisplayedColumns[displayedcolumn]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitDisplayedColumn(displayedcolumn)
		}
	}
	return displayedcolumn
}

func (displayedcolumn *DisplayedColumn) CommitVoid(stage *StageStruct) {
	displayedcolumn.Commit(stage)
}

// Checkout displayedcolumn to the back repo (if it is already staged)
func (displayedcolumn *DisplayedColumn) Checkout(stage *StageStruct) *DisplayedColumn {
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

// Stage puts formdiv to the model stage
func (formdiv *FormDiv) Stage(stage *StageStruct) *FormDiv {
	stage.FormDivs[formdiv] = __member
	stage.FormDivs_mapString[formdiv.Name] = formdiv

	return formdiv
}

// Unstage removes formdiv off the model stage
func (formdiv *FormDiv) Unstage(stage *StageStruct) *FormDiv {
	delete(stage.FormDivs, formdiv)
	delete(stage.FormDivs_mapString, formdiv.Name)
	return formdiv
}

// UnstageVoid removes formdiv off the model stage
func (formdiv *FormDiv) UnstageVoid(stage *StageStruct) {
	delete(stage.FormDivs, formdiv)
	delete(stage.FormDivs_mapString, formdiv.Name)
}

// commit formdiv to the back repo (if it is already staged)
func (formdiv *FormDiv) Commit(stage *StageStruct) *FormDiv {
	if _, ok := stage.FormDivs[formdiv]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormDiv(formdiv)
		}
	}
	return formdiv
}

func (formdiv *FormDiv) CommitVoid(stage *StageStruct) {
	formdiv.Commit(stage)
}

// Checkout formdiv to the back repo (if it is already staged)
func (formdiv *FormDiv) Checkout(stage *StageStruct) *FormDiv {
	if _, ok := stage.FormDivs[formdiv]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormDiv(formdiv)
		}
	}
	return formdiv
}

// for satisfaction of GongStruct interface
func (formdiv *FormDiv) GetName() (res string) {
	return formdiv.Name
}

// Stage puts formeditassocbutton to the model stage
func (formeditassocbutton *FormEditAssocButton) Stage(stage *StageStruct) *FormEditAssocButton {
	stage.FormEditAssocButtons[formeditassocbutton] = __member
	stage.FormEditAssocButtons_mapString[formeditassocbutton.Name] = formeditassocbutton

	return formeditassocbutton
}

// Unstage removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) Unstage(stage *StageStruct) *FormEditAssocButton {
	delete(stage.FormEditAssocButtons, formeditassocbutton)
	delete(stage.FormEditAssocButtons_mapString, formeditassocbutton.Name)
	return formeditassocbutton
}

// UnstageVoid removes formeditassocbutton off the model stage
func (formeditassocbutton *FormEditAssocButton) UnstageVoid(stage *StageStruct) {
	delete(stage.FormEditAssocButtons, formeditassocbutton)
	delete(stage.FormEditAssocButtons_mapString, formeditassocbutton.Name)
}

// commit formeditassocbutton to the back repo (if it is already staged)
func (formeditassocbutton *FormEditAssocButton) Commit(stage *StageStruct) *FormEditAssocButton {
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormEditAssocButton(formeditassocbutton)
		}
	}
	return formeditassocbutton
}

func (formeditassocbutton *FormEditAssocButton) CommitVoid(stage *StageStruct) {
	formeditassocbutton.Commit(stage)
}

// Checkout formeditassocbutton to the back repo (if it is already staged)
func (formeditassocbutton *FormEditAssocButton) Checkout(stage *StageStruct) *FormEditAssocButton {
	if _, ok := stage.FormEditAssocButtons[formeditassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormEditAssocButton(formeditassocbutton)
		}
	}
	return formeditassocbutton
}

// for satisfaction of GongStruct interface
func (formeditassocbutton *FormEditAssocButton) GetName() (res string) {
	return formeditassocbutton.Name
}

// Stage puts formfield to the model stage
func (formfield *FormField) Stage(stage *StageStruct) *FormField {
	stage.FormFields[formfield] = __member
	stage.FormFields_mapString[formfield.Name] = formfield

	return formfield
}

// Unstage removes formfield off the model stage
func (formfield *FormField) Unstage(stage *StageStruct) *FormField {
	delete(stage.FormFields, formfield)
	delete(stage.FormFields_mapString, formfield.Name)
	return formfield
}

// UnstageVoid removes formfield off the model stage
func (formfield *FormField) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFields, formfield)
	delete(stage.FormFields_mapString, formfield.Name)
}

// commit formfield to the back repo (if it is already staged)
func (formfield *FormField) Commit(stage *StageStruct) *FormField {
	if _, ok := stage.FormFields[formfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormField(formfield)
		}
	}
	return formfield
}

func (formfield *FormField) CommitVoid(stage *StageStruct) {
	formfield.Commit(stage)
}

// Checkout formfield to the back repo (if it is already staged)
func (formfield *FormField) Checkout(stage *StageStruct) *FormField {
	if _, ok := stage.FormFields[formfield]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormField(formfield)
		}
	}
	return formfield
}

// for satisfaction of GongStruct interface
func (formfield *FormField) GetName() (res string) {
	return formfield.Name
}

// Stage puts formfielddate to the model stage
func (formfielddate *FormFieldDate) Stage(stage *StageStruct) *FormFieldDate {
	stage.FormFieldDates[formfielddate] = __member
	stage.FormFieldDates_mapString[formfielddate.Name] = formfielddate

	return formfielddate
}

// Unstage removes formfielddate off the model stage
func (formfielddate *FormFieldDate) Unstage(stage *StageStruct) *FormFieldDate {
	delete(stage.FormFieldDates, formfielddate)
	delete(stage.FormFieldDates_mapString, formfielddate.Name)
	return formfielddate
}

// UnstageVoid removes formfielddate off the model stage
func (formfielddate *FormFieldDate) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldDates, formfielddate)
	delete(stage.FormFieldDates_mapString, formfielddate.Name)
}

// commit formfielddate to the back repo (if it is already staged)
func (formfielddate *FormFieldDate) Commit(stage *StageStruct) *FormFieldDate {
	if _, ok := stage.FormFieldDates[formfielddate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldDate(formfielddate)
		}
	}
	return formfielddate
}

func (formfielddate *FormFieldDate) CommitVoid(stage *StageStruct) {
	formfielddate.Commit(stage)
}

// Checkout formfielddate to the back repo (if it is already staged)
func (formfielddate *FormFieldDate) Checkout(stage *StageStruct) *FormFieldDate {
	if _, ok := stage.FormFieldDates[formfielddate]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldDate(formfielddate)
		}
	}
	return formfielddate
}

// for satisfaction of GongStruct interface
func (formfielddate *FormFieldDate) GetName() (res string) {
	return formfielddate.Name
}

// Stage puts formfielddatetime to the model stage
func (formfielddatetime *FormFieldDateTime) Stage(stage *StageStruct) *FormFieldDateTime {
	stage.FormFieldDateTimes[formfielddatetime] = __member
	stage.FormFieldDateTimes_mapString[formfielddatetime.Name] = formfielddatetime

	return formfielddatetime
}

// Unstage removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) Unstage(stage *StageStruct) *FormFieldDateTime {
	delete(stage.FormFieldDateTimes, formfielddatetime)
	delete(stage.FormFieldDateTimes_mapString, formfielddatetime.Name)
	return formfielddatetime
}

// UnstageVoid removes formfielddatetime off the model stage
func (formfielddatetime *FormFieldDateTime) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldDateTimes, formfielddatetime)
	delete(stage.FormFieldDateTimes_mapString, formfielddatetime.Name)
}

// commit formfielddatetime to the back repo (if it is already staged)
func (formfielddatetime *FormFieldDateTime) Commit(stage *StageStruct) *FormFieldDateTime {
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldDateTime(formfielddatetime)
		}
	}
	return formfielddatetime
}

func (formfielddatetime *FormFieldDateTime) CommitVoid(stage *StageStruct) {
	formfielddatetime.Commit(stage)
}

// Checkout formfielddatetime to the back repo (if it is already staged)
func (formfielddatetime *FormFieldDateTime) Checkout(stage *StageStruct) *FormFieldDateTime {
	if _, ok := stage.FormFieldDateTimes[formfielddatetime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldDateTime(formfielddatetime)
		}
	}
	return formfielddatetime
}

// for satisfaction of GongStruct interface
func (formfielddatetime *FormFieldDateTime) GetName() (res string) {
	return formfielddatetime.Name
}

// Stage puts formfieldfloat64 to the model stage
func (formfieldfloat64 *FormFieldFloat64) Stage(stage *StageStruct) *FormFieldFloat64 {
	stage.FormFieldFloat64s[formfieldfloat64] = __member
	stage.FormFieldFloat64s_mapString[formfieldfloat64.Name] = formfieldfloat64

	return formfieldfloat64
}

// Unstage removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) Unstage(stage *StageStruct) *FormFieldFloat64 {
	delete(stage.FormFieldFloat64s, formfieldfloat64)
	delete(stage.FormFieldFloat64s_mapString, formfieldfloat64.Name)
	return formfieldfloat64
}

// UnstageVoid removes formfieldfloat64 off the model stage
func (formfieldfloat64 *FormFieldFloat64) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldFloat64s, formfieldfloat64)
	delete(stage.FormFieldFloat64s_mapString, formfieldfloat64.Name)
}

// commit formfieldfloat64 to the back repo (if it is already staged)
func (formfieldfloat64 *FormFieldFloat64) Commit(stage *StageStruct) *FormFieldFloat64 {
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldFloat64(formfieldfloat64)
		}
	}
	return formfieldfloat64
}

func (formfieldfloat64 *FormFieldFloat64) CommitVoid(stage *StageStruct) {
	formfieldfloat64.Commit(stage)
}

// Checkout formfieldfloat64 to the back repo (if it is already staged)
func (formfieldfloat64 *FormFieldFloat64) Checkout(stage *StageStruct) *FormFieldFloat64 {
	if _, ok := stage.FormFieldFloat64s[formfieldfloat64]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldFloat64(formfieldfloat64)
		}
	}
	return formfieldfloat64
}

// for satisfaction of GongStruct interface
func (formfieldfloat64 *FormFieldFloat64) GetName() (res string) {
	return formfieldfloat64.Name
}

// Stage puts formfieldint to the model stage
func (formfieldint *FormFieldInt) Stage(stage *StageStruct) *FormFieldInt {
	stage.FormFieldInts[formfieldint] = __member
	stage.FormFieldInts_mapString[formfieldint.Name] = formfieldint

	return formfieldint
}

// Unstage removes formfieldint off the model stage
func (formfieldint *FormFieldInt) Unstage(stage *StageStruct) *FormFieldInt {
	delete(stage.FormFieldInts, formfieldint)
	delete(stage.FormFieldInts_mapString, formfieldint.Name)
	return formfieldint
}

// UnstageVoid removes formfieldint off the model stage
func (formfieldint *FormFieldInt) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldInts, formfieldint)
	delete(stage.FormFieldInts_mapString, formfieldint.Name)
}

// commit formfieldint to the back repo (if it is already staged)
func (formfieldint *FormFieldInt) Commit(stage *StageStruct) *FormFieldInt {
	if _, ok := stage.FormFieldInts[formfieldint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldInt(formfieldint)
		}
	}
	return formfieldint
}

func (formfieldint *FormFieldInt) CommitVoid(stage *StageStruct) {
	formfieldint.Commit(stage)
}

// Checkout formfieldint to the back repo (if it is already staged)
func (formfieldint *FormFieldInt) Checkout(stage *StageStruct) *FormFieldInt {
	if _, ok := stage.FormFieldInts[formfieldint]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldInt(formfieldint)
		}
	}
	return formfieldint
}

// for satisfaction of GongStruct interface
func (formfieldint *FormFieldInt) GetName() (res string) {
	return formfieldint.Name
}

// Stage puts formfieldselect to the model stage
func (formfieldselect *FormFieldSelect) Stage(stage *StageStruct) *FormFieldSelect {
	stage.FormFieldSelects[formfieldselect] = __member
	stage.FormFieldSelects_mapString[formfieldselect.Name] = formfieldselect

	return formfieldselect
}

// Unstage removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) Unstage(stage *StageStruct) *FormFieldSelect {
	delete(stage.FormFieldSelects, formfieldselect)
	delete(stage.FormFieldSelects_mapString, formfieldselect.Name)
	return formfieldselect
}

// UnstageVoid removes formfieldselect off the model stage
func (formfieldselect *FormFieldSelect) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldSelects, formfieldselect)
	delete(stage.FormFieldSelects_mapString, formfieldselect.Name)
}

// commit formfieldselect to the back repo (if it is already staged)
func (formfieldselect *FormFieldSelect) Commit(stage *StageStruct) *FormFieldSelect {
	if _, ok := stage.FormFieldSelects[formfieldselect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldSelect(formfieldselect)
		}
	}
	return formfieldselect
}

func (formfieldselect *FormFieldSelect) CommitVoid(stage *StageStruct) {
	formfieldselect.Commit(stage)
}

// Checkout formfieldselect to the back repo (if it is already staged)
func (formfieldselect *FormFieldSelect) Checkout(stage *StageStruct) *FormFieldSelect {
	if _, ok := stage.FormFieldSelects[formfieldselect]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldSelect(formfieldselect)
		}
	}
	return formfieldselect
}

// for satisfaction of GongStruct interface
func (formfieldselect *FormFieldSelect) GetName() (res string) {
	return formfieldselect.Name
}

// Stage puts formfieldstring to the model stage
func (formfieldstring *FormFieldString) Stage(stage *StageStruct) *FormFieldString {
	stage.FormFieldStrings[formfieldstring] = __member
	stage.FormFieldStrings_mapString[formfieldstring.Name] = formfieldstring

	return formfieldstring
}

// Unstage removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) Unstage(stage *StageStruct) *FormFieldString {
	delete(stage.FormFieldStrings, formfieldstring)
	delete(stage.FormFieldStrings_mapString, formfieldstring.Name)
	return formfieldstring
}

// UnstageVoid removes formfieldstring off the model stage
func (formfieldstring *FormFieldString) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldStrings, formfieldstring)
	delete(stage.FormFieldStrings_mapString, formfieldstring.Name)
}

// commit formfieldstring to the back repo (if it is already staged)
func (formfieldstring *FormFieldString) Commit(stage *StageStruct) *FormFieldString {
	if _, ok := stage.FormFieldStrings[formfieldstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldString(formfieldstring)
		}
	}
	return formfieldstring
}

func (formfieldstring *FormFieldString) CommitVoid(stage *StageStruct) {
	formfieldstring.Commit(stage)
}

// Checkout formfieldstring to the back repo (if it is already staged)
func (formfieldstring *FormFieldString) Checkout(stage *StageStruct) *FormFieldString {
	if _, ok := stage.FormFieldStrings[formfieldstring]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldString(formfieldstring)
		}
	}
	return formfieldstring
}

// for satisfaction of GongStruct interface
func (formfieldstring *FormFieldString) GetName() (res string) {
	return formfieldstring.Name
}

// Stage puts formfieldtime to the model stage
func (formfieldtime *FormFieldTime) Stage(stage *StageStruct) *FormFieldTime {
	stage.FormFieldTimes[formfieldtime] = __member
	stage.FormFieldTimes_mapString[formfieldtime.Name] = formfieldtime

	return formfieldtime
}

// Unstage removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) Unstage(stage *StageStruct) *FormFieldTime {
	delete(stage.FormFieldTimes, formfieldtime)
	delete(stage.FormFieldTimes_mapString, formfieldtime.Name)
	return formfieldtime
}

// UnstageVoid removes formfieldtime off the model stage
func (formfieldtime *FormFieldTime) UnstageVoid(stage *StageStruct) {
	delete(stage.FormFieldTimes, formfieldtime)
	delete(stage.FormFieldTimes_mapString, formfieldtime.Name)
}

// commit formfieldtime to the back repo (if it is already staged)
func (formfieldtime *FormFieldTime) Commit(stage *StageStruct) *FormFieldTime {
	if _, ok := stage.FormFieldTimes[formfieldtime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormFieldTime(formfieldtime)
		}
	}
	return formfieldtime
}

func (formfieldtime *FormFieldTime) CommitVoid(stage *StageStruct) {
	formfieldtime.Commit(stage)
}

// Checkout formfieldtime to the back repo (if it is already staged)
func (formfieldtime *FormFieldTime) Checkout(stage *StageStruct) *FormFieldTime {
	if _, ok := stage.FormFieldTimes[formfieldtime]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormFieldTime(formfieldtime)
		}
	}
	return formfieldtime
}

// for satisfaction of GongStruct interface
func (formfieldtime *FormFieldTime) GetName() (res string) {
	return formfieldtime.Name
}

// Stage puts formgroup to the model stage
func (formgroup *FormGroup) Stage(stage *StageStruct) *FormGroup {
	stage.FormGroups[formgroup] = __member
	stage.FormGroups_mapString[formgroup.Name] = formgroup

	return formgroup
}

// Unstage removes formgroup off the model stage
func (formgroup *FormGroup) Unstage(stage *StageStruct) *FormGroup {
	delete(stage.FormGroups, formgroup)
	delete(stage.FormGroups_mapString, formgroup.Name)
	return formgroup
}

// UnstageVoid removes formgroup off the model stage
func (formgroup *FormGroup) UnstageVoid(stage *StageStruct) {
	delete(stage.FormGroups, formgroup)
	delete(stage.FormGroups_mapString, formgroup.Name)
}

// commit formgroup to the back repo (if it is already staged)
func (formgroup *FormGroup) Commit(stage *StageStruct) *FormGroup {
	if _, ok := stage.FormGroups[formgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormGroup(formgroup)
		}
	}
	return formgroup
}

func (formgroup *FormGroup) CommitVoid(stage *StageStruct) {
	formgroup.Commit(stage)
}

// Checkout formgroup to the back repo (if it is already staged)
func (formgroup *FormGroup) Checkout(stage *StageStruct) *FormGroup {
	if _, ok := stage.FormGroups[formgroup]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormGroup(formgroup)
		}
	}
	return formgroup
}

// for satisfaction of GongStruct interface
func (formgroup *FormGroup) GetName() (res string) {
	return formgroup.Name
}

// Stage puts formsortassocbutton to the model stage
func (formsortassocbutton *FormSortAssocButton) Stage(stage *StageStruct) *FormSortAssocButton {
	stage.FormSortAssocButtons[formsortassocbutton] = __member
	stage.FormSortAssocButtons_mapString[formsortassocbutton.Name] = formsortassocbutton

	return formsortassocbutton
}

// Unstage removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) Unstage(stage *StageStruct) *FormSortAssocButton {
	delete(stage.FormSortAssocButtons, formsortassocbutton)
	delete(stage.FormSortAssocButtons_mapString, formsortassocbutton.Name)
	return formsortassocbutton
}

// UnstageVoid removes formsortassocbutton off the model stage
func (formsortassocbutton *FormSortAssocButton) UnstageVoid(stage *StageStruct) {
	delete(stage.FormSortAssocButtons, formsortassocbutton)
	delete(stage.FormSortAssocButtons_mapString, formsortassocbutton.Name)
}

// commit formsortassocbutton to the back repo (if it is already staged)
func (formsortassocbutton *FormSortAssocButton) Commit(stage *StageStruct) *FormSortAssocButton {
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitFormSortAssocButton(formsortassocbutton)
		}
	}
	return formsortassocbutton
}

func (formsortassocbutton *FormSortAssocButton) CommitVoid(stage *StageStruct) {
	formsortassocbutton.Commit(stage)
}

// Checkout formsortassocbutton to the back repo (if it is already staged)
func (formsortassocbutton *FormSortAssocButton) Checkout(stage *StageStruct) *FormSortAssocButton {
	if _, ok := stage.FormSortAssocButtons[formsortassocbutton]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutFormSortAssocButton(formsortassocbutton)
		}
	}
	return formsortassocbutton
}

// for satisfaction of GongStruct interface
func (formsortassocbutton *FormSortAssocButton) GetName() (res string) {
	return formsortassocbutton.Name
}

// Stage puts option to the model stage
func (option *Option) Stage(stage *StageStruct) *Option {
	stage.Options[option] = __member
	stage.Options_mapString[option.Name] = option

	return option
}

// Unstage removes option off the model stage
func (option *Option) Unstage(stage *StageStruct) *Option {
	delete(stage.Options, option)
	delete(stage.Options_mapString, option.Name)
	return option
}

// UnstageVoid removes option off the model stage
func (option *Option) UnstageVoid(stage *StageStruct) {
	delete(stage.Options, option)
	delete(stage.Options_mapString, option.Name)
}

// commit option to the back repo (if it is already staged)
func (option *Option) Commit(stage *StageStruct) *Option {
	if _, ok := stage.Options[option]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitOption(option)
		}
	}
	return option
}

func (option *Option) CommitVoid(stage *StageStruct) {
	option.Commit(stage)
}

// Checkout option to the back repo (if it is already staged)
func (option *Option) Checkout(stage *StageStruct) *Option {
	if _, ok := stage.Options[option]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CheckoutOption(option)
		}
	}
	return option
}

// for satisfaction of GongStruct interface
func (option *Option) GetName() (res string) {
	return option.Name
}

// Stage puts row to the model stage
func (row *Row) Stage(stage *StageStruct) *Row {
	stage.Rows[row] = __member
	stage.Rows_mapString[row.Name] = row

	return row
}

// Unstage removes row off the model stage
func (row *Row) Unstage(stage *StageStruct) *Row {
	delete(stage.Rows, row)
	delete(stage.Rows_mapString, row.Name)
	return row
}

// UnstageVoid removes row off the model stage
func (row *Row) UnstageVoid(stage *StageStruct) {
	delete(stage.Rows, row)
	delete(stage.Rows_mapString, row.Name)
}

// commit row to the back repo (if it is already staged)
func (row *Row) Commit(stage *StageStruct) *Row {
	if _, ok := stage.Rows[row]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitRow(row)
		}
	}
	return row
}

func (row *Row) CommitVoid(stage *StageStruct) {
	row.Commit(stage)
}

// Checkout row to the back repo (if it is already staged)
func (row *Row) Checkout(stage *StageStruct) *Row {
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

// Stage puts table to the model stage
func (table *Table) Stage(stage *StageStruct) *Table {
	stage.Tables[table] = __member
	stage.Tables_mapString[table.Name] = table

	return table
}

// Unstage removes table off the model stage
func (table *Table) Unstage(stage *StageStruct) *Table {
	delete(stage.Tables, table)
	delete(stage.Tables_mapString, table.Name)
	return table
}

// UnstageVoid removes table off the model stage
func (table *Table) UnstageVoid(stage *StageStruct) {
	delete(stage.Tables, table)
	delete(stage.Tables_mapString, table.Name)
}

// commit table to the back repo (if it is already staged)
func (table *Table) Commit(stage *StageStruct) *Table {
	if _, ok := stage.Tables[table]; ok {
		if stage.BackRepo != nil {
			stage.BackRepo.CommitTable(table)
		}
	}
	return table
}

func (table *Table) CommitVoid(stage *StageStruct) {
	table.Commit(stage)
}

// Checkout table to the back repo (if it is already staged)
func (table *Table) Checkout(stage *StageStruct) *Table {
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

// swagger:ignore
type AllModelsStructCreateInterface interface { // insertion point for Callbacks on creation
	CreateORMCell(Cell *Cell)
	CreateORMCellBoolean(CellBoolean *CellBoolean)
	CreateORMCellFloat64(CellFloat64 *CellFloat64)
	CreateORMCellIcon(CellIcon *CellIcon)
	CreateORMCellInt(CellInt *CellInt)
	CreateORMCellString(CellString *CellString)
	CreateORMCheckBox(CheckBox *CheckBox)
	CreateORMDisplayedColumn(DisplayedColumn *DisplayedColumn)
	CreateORMFormDiv(FormDiv *FormDiv)
	CreateORMFormEditAssocButton(FormEditAssocButton *FormEditAssocButton)
	CreateORMFormField(FormField *FormField)
	CreateORMFormFieldDate(FormFieldDate *FormFieldDate)
	CreateORMFormFieldDateTime(FormFieldDateTime *FormFieldDateTime)
	CreateORMFormFieldFloat64(FormFieldFloat64 *FormFieldFloat64)
	CreateORMFormFieldInt(FormFieldInt *FormFieldInt)
	CreateORMFormFieldSelect(FormFieldSelect *FormFieldSelect)
	CreateORMFormFieldString(FormFieldString *FormFieldString)
	CreateORMFormFieldTime(FormFieldTime *FormFieldTime)
	CreateORMFormGroup(FormGroup *FormGroup)
	CreateORMFormSortAssocButton(FormSortAssocButton *FormSortAssocButton)
	CreateORMOption(Option *Option)
	CreateORMRow(Row *Row)
	CreateORMTable(Table *Table)
}

type AllModelsStructDeleteInterface interface { // insertion point for Callbacks on deletion
	DeleteORMCell(Cell *Cell)
	DeleteORMCellBoolean(CellBoolean *CellBoolean)
	DeleteORMCellFloat64(CellFloat64 *CellFloat64)
	DeleteORMCellIcon(CellIcon *CellIcon)
	DeleteORMCellInt(CellInt *CellInt)
	DeleteORMCellString(CellString *CellString)
	DeleteORMCheckBox(CheckBox *CheckBox)
	DeleteORMDisplayedColumn(DisplayedColumn *DisplayedColumn)
	DeleteORMFormDiv(FormDiv *FormDiv)
	DeleteORMFormEditAssocButton(FormEditAssocButton *FormEditAssocButton)
	DeleteORMFormField(FormField *FormField)
	DeleteORMFormFieldDate(FormFieldDate *FormFieldDate)
	DeleteORMFormFieldDateTime(FormFieldDateTime *FormFieldDateTime)
	DeleteORMFormFieldFloat64(FormFieldFloat64 *FormFieldFloat64)
	DeleteORMFormFieldInt(FormFieldInt *FormFieldInt)
	DeleteORMFormFieldSelect(FormFieldSelect *FormFieldSelect)
	DeleteORMFormFieldString(FormFieldString *FormFieldString)
	DeleteORMFormFieldTime(FormFieldTime *FormFieldTime)
	DeleteORMFormGroup(FormGroup *FormGroup)
	DeleteORMFormSortAssocButton(FormSortAssocButton *FormSortAssocButton)
	DeleteORMOption(Option *Option)
	DeleteORMRow(Row *Row)
	DeleteORMTable(Table *Table)
}

func (stage *StageStruct) Reset() { // insertion point for array reset
	stage.Cells = make(map[*Cell]any)
	stage.Cells_mapString = make(map[string]*Cell)

	stage.CellBooleans = make(map[*CellBoolean]any)
	stage.CellBooleans_mapString = make(map[string]*CellBoolean)

	stage.CellFloat64s = make(map[*CellFloat64]any)
	stage.CellFloat64s_mapString = make(map[string]*CellFloat64)

	stage.CellIcons = make(map[*CellIcon]any)
	stage.CellIcons_mapString = make(map[string]*CellIcon)

	stage.CellInts = make(map[*CellInt]any)
	stage.CellInts_mapString = make(map[string]*CellInt)

	stage.CellStrings = make(map[*CellString]any)
	stage.CellStrings_mapString = make(map[string]*CellString)

	stage.CheckBoxs = make(map[*CheckBox]any)
	stage.CheckBoxs_mapString = make(map[string]*CheckBox)

	stage.DisplayedColumns = make(map[*DisplayedColumn]any)
	stage.DisplayedColumns_mapString = make(map[string]*DisplayedColumn)

	stage.FormDivs = make(map[*FormDiv]any)
	stage.FormDivs_mapString = make(map[string]*FormDiv)

	stage.FormEditAssocButtons = make(map[*FormEditAssocButton]any)
	stage.FormEditAssocButtons_mapString = make(map[string]*FormEditAssocButton)

	stage.FormFields = make(map[*FormField]any)
	stage.FormFields_mapString = make(map[string]*FormField)

	stage.FormFieldDates = make(map[*FormFieldDate]any)
	stage.FormFieldDates_mapString = make(map[string]*FormFieldDate)

	stage.FormFieldDateTimes = make(map[*FormFieldDateTime]any)
	stage.FormFieldDateTimes_mapString = make(map[string]*FormFieldDateTime)

	stage.FormFieldFloat64s = make(map[*FormFieldFloat64]any)
	stage.FormFieldFloat64s_mapString = make(map[string]*FormFieldFloat64)

	stage.FormFieldInts = make(map[*FormFieldInt]any)
	stage.FormFieldInts_mapString = make(map[string]*FormFieldInt)

	stage.FormFieldSelects = make(map[*FormFieldSelect]any)
	stage.FormFieldSelects_mapString = make(map[string]*FormFieldSelect)

	stage.FormFieldStrings = make(map[*FormFieldString]any)
	stage.FormFieldStrings_mapString = make(map[string]*FormFieldString)

	stage.FormFieldTimes = make(map[*FormFieldTime]any)
	stage.FormFieldTimes_mapString = make(map[string]*FormFieldTime)

	stage.FormGroups = make(map[*FormGroup]any)
	stage.FormGroups_mapString = make(map[string]*FormGroup)

	stage.FormSortAssocButtons = make(map[*FormSortAssocButton]any)
	stage.FormSortAssocButtons_mapString = make(map[string]*FormSortAssocButton)

	stage.Options = make(map[*Option]any)
	stage.Options_mapString = make(map[string]*Option)

	stage.Rows = make(map[*Row]any)
	stage.Rows_mapString = make(map[string]*Row)

	stage.Tables = make(map[*Table]any)
	stage.Tables_mapString = make(map[string]*Table)

}

func (stage *StageStruct) Nil() { // insertion point for array nil
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

	stage.CheckBoxs = nil
	stage.CheckBoxs_mapString = nil

	stage.DisplayedColumns = nil
	stage.DisplayedColumns_mapString = nil

	stage.FormDivs = nil
	stage.FormDivs_mapString = nil

	stage.FormEditAssocButtons = nil
	stage.FormEditAssocButtons_mapString = nil

	stage.FormFields = nil
	stage.FormFields_mapString = nil

	stage.FormFieldDates = nil
	stage.FormFieldDates_mapString = nil

	stage.FormFieldDateTimes = nil
	stage.FormFieldDateTimes_mapString = nil

	stage.FormFieldFloat64s = nil
	stage.FormFieldFloat64s_mapString = nil

	stage.FormFieldInts = nil
	stage.FormFieldInts_mapString = nil

	stage.FormFieldSelects = nil
	stage.FormFieldSelects_mapString = nil

	stage.FormFieldStrings = nil
	stage.FormFieldStrings_mapString = nil

	stage.FormFieldTimes = nil
	stage.FormFieldTimes_mapString = nil

	stage.FormGroups = nil
	stage.FormGroups_mapString = nil

	stage.FormSortAssocButtons = nil
	stage.FormSortAssocButtons_mapString = nil

	stage.Options = nil
	stage.Options_mapString = nil

	stage.Rows = nil
	stage.Rows_mapString = nil

	stage.Tables = nil
	stage.Tables_mapString = nil

}

func (stage *StageStruct) Unstage() { // insertion point for array nil
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

	for checkbox := range stage.CheckBoxs {
		checkbox.Unstage(stage)
	}

	for displayedcolumn := range stage.DisplayedColumns {
		displayedcolumn.Unstage(stage)
	}

	for formdiv := range stage.FormDivs {
		formdiv.Unstage(stage)
	}

	for formeditassocbutton := range stage.FormEditAssocButtons {
		formeditassocbutton.Unstage(stage)
	}

	for formfield := range stage.FormFields {
		formfield.Unstage(stage)
	}

	for formfielddate := range stage.FormFieldDates {
		formfielddate.Unstage(stage)
	}

	for formfielddatetime := range stage.FormFieldDateTimes {
		formfielddatetime.Unstage(stage)
	}

	for formfieldfloat64 := range stage.FormFieldFloat64s {
		formfieldfloat64.Unstage(stage)
	}

	for formfieldint := range stage.FormFieldInts {
		formfieldint.Unstage(stage)
	}

	for formfieldselect := range stage.FormFieldSelects {
		formfieldselect.Unstage(stage)
	}

	for formfieldstring := range stage.FormFieldStrings {
		formfieldstring.Unstage(stage)
	}

	for formfieldtime := range stage.FormFieldTimes {
		formfieldtime.Unstage(stage)
	}

	for formgroup := range stage.FormGroups {
		formgroup.Unstage(stage)
	}

	for formsortassocbutton := range stage.FormSortAssocButtons {
		formsortassocbutton.Unstage(stage)
	}

	for option := range stage.Options {
		option.Unstage(stage)
	}

	for row := range stage.Rows {
		row.Unstage(stage)
	}

	for table := range stage.Tables {
		table.Unstage(stage)
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
	CommitVoid(*StageStruct)
	UnstageVoid(stage *StageStruct)
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

func GetGongstrucsSorted[T PointerToGongstruct](stage *StageStruct) (sortedSlice []T) {

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
func GongGetSet[Type GongstructSet](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
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
	case map[*CheckBox]any:
		return any(&stage.CheckBoxs).(*Type)
	case map[*DisplayedColumn]any:
		return any(&stage.DisplayedColumns).(*Type)
	case map[*FormDiv]any:
		return any(&stage.FormDivs).(*Type)
	case map[*FormEditAssocButton]any:
		return any(&stage.FormEditAssocButtons).(*Type)
	case map[*FormField]any:
		return any(&stage.FormFields).(*Type)
	case map[*FormFieldDate]any:
		return any(&stage.FormFieldDates).(*Type)
	case map[*FormFieldDateTime]any:
		return any(&stage.FormFieldDateTimes).(*Type)
	case map[*FormFieldFloat64]any:
		return any(&stage.FormFieldFloat64s).(*Type)
	case map[*FormFieldInt]any:
		return any(&stage.FormFieldInts).(*Type)
	case map[*FormFieldSelect]any:
		return any(&stage.FormFieldSelects).(*Type)
	case map[*FormFieldString]any:
		return any(&stage.FormFieldStrings).(*Type)
	case map[*FormFieldTime]any:
		return any(&stage.FormFieldTimes).(*Type)
	case map[*FormGroup]any:
		return any(&stage.FormGroups).(*Type)
	case map[*FormSortAssocButton]any:
		return any(&stage.FormSortAssocButtons).(*Type)
	case map[*Option]any:
		return any(&stage.Options).(*Type)
	case map[*Row]any:
		return any(&stage.Rows).(*Type)
	case map[*Table]any:
		return any(&stage.Tables).(*Type)
	default:
		return nil
	}
}

// GongGetMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GongGetMap[Type GongstructMapString](stage *StageStruct) *Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case map[string]*Cell:
		return any(&stage.Cells_mapString).(*Type)
	case map[string]*CellBoolean:
		return any(&stage.CellBooleans_mapString).(*Type)
	case map[string]*CellFloat64:
		return any(&stage.CellFloat64s_mapString).(*Type)
	case map[string]*CellIcon:
		return any(&stage.CellIcons_mapString).(*Type)
	case map[string]*CellInt:
		return any(&stage.CellInts_mapString).(*Type)
	case map[string]*CellString:
		return any(&stage.CellStrings_mapString).(*Type)
	case map[string]*CheckBox:
		return any(&stage.CheckBoxs_mapString).(*Type)
	case map[string]*DisplayedColumn:
		return any(&stage.DisplayedColumns_mapString).(*Type)
	case map[string]*FormDiv:
		return any(&stage.FormDivs_mapString).(*Type)
	case map[string]*FormEditAssocButton:
		return any(&stage.FormEditAssocButtons_mapString).(*Type)
	case map[string]*FormField:
		return any(&stage.FormFields_mapString).(*Type)
	case map[string]*FormFieldDate:
		return any(&stage.FormFieldDates_mapString).(*Type)
	case map[string]*FormFieldDateTime:
		return any(&stage.FormFieldDateTimes_mapString).(*Type)
	case map[string]*FormFieldFloat64:
		return any(&stage.FormFieldFloat64s_mapString).(*Type)
	case map[string]*FormFieldInt:
		return any(&stage.FormFieldInts_mapString).(*Type)
	case map[string]*FormFieldSelect:
		return any(&stage.FormFieldSelects_mapString).(*Type)
	case map[string]*FormFieldString:
		return any(&stage.FormFieldStrings_mapString).(*Type)
	case map[string]*FormFieldTime:
		return any(&stage.FormFieldTimes_mapString).(*Type)
	case map[string]*FormGroup:
		return any(&stage.FormGroups_mapString).(*Type)
	case map[string]*FormSortAssocButton:
		return any(&stage.FormSortAssocButtons_mapString).(*Type)
	case map[string]*Option:
		return any(&stage.Options_mapString).(*Type)
	case map[string]*Row:
		return any(&stage.Rows_mapString).(*Type)
	case map[string]*Table:
		return any(&stage.Tables_mapString).(*Type)
	default:
		return nil
	}
}

// GetGongstructInstancesSet returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSet[Type Gongstruct](stage *StageStruct) *map[*Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case Cell:
		return any(&stage.Cells).(*map[*Type]any)
	case CellBoolean:
		return any(&stage.CellBooleans).(*map[*Type]any)
	case CellFloat64:
		return any(&stage.CellFloat64s).(*map[*Type]any)
	case CellIcon:
		return any(&stage.CellIcons).(*map[*Type]any)
	case CellInt:
		return any(&stage.CellInts).(*map[*Type]any)
	case CellString:
		return any(&stage.CellStrings).(*map[*Type]any)
	case CheckBox:
		return any(&stage.CheckBoxs).(*map[*Type]any)
	case DisplayedColumn:
		return any(&stage.DisplayedColumns).(*map[*Type]any)
	case FormDiv:
		return any(&stage.FormDivs).(*map[*Type]any)
	case FormEditAssocButton:
		return any(&stage.FormEditAssocButtons).(*map[*Type]any)
	case FormField:
		return any(&stage.FormFields).(*map[*Type]any)
	case FormFieldDate:
		return any(&stage.FormFieldDates).(*map[*Type]any)
	case FormFieldDateTime:
		return any(&stage.FormFieldDateTimes).(*map[*Type]any)
	case FormFieldFloat64:
		return any(&stage.FormFieldFloat64s).(*map[*Type]any)
	case FormFieldInt:
		return any(&stage.FormFieldInts).(*map[*Type]any)
	case FormFieldSelect:
		return any(&stage.FormFieldSelects).(*map[*Type]any)
	case FormFieldString:
		return any(&stage.FormFieldStrings).(*map[*Type]any)
	case FormFieldTime:
		return any(&stage.FormFieldTimes).(*map[*Type]any)
	case FormGroup:
		return any(&stage.FormGroups).(*map[*Type]any)
	case FormSortAssocButton:
		return any(&stage.FormSortAssocButtons).(*map[*Type]any)
	case Option:
		return any(&stage.Options).(*map[*Type]any)
	case Row:
		return any(&stage.Rows).(*map[*Type]any)
	case Table:
		return any(&stage.Tables).(*map[*Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesSetFromPointerType returns the set staged GongstructType instances
// it is usefull because it allows refactoring of gongstruct identifier
func GetGongstructInstancesSetFromPointerType[Type PointerToGongstruct](stage *StageStruct) *map[Type]any {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
	case *Cell:
		return any(&stage.Cells).(*map[Type]any)
	case *CellBoolean:
		return any(&stage.CellBooleans).(*map[Type]any)
	case *CellFloat64:
		return any(&stage.CellFloat64s).(*map[Type]any)
	case *CellIcon:
		return any(&stage.CellIcons).(*map[Type]any)
	case *CellInt:
		return any(&stage.CellInts).(*map[Type]any)
	case *CellString:
		return any(&stage.CellStrings).(*map[Type]any)
	case *CheckBox:
		return any(&stage.CheckBoxs).(*map[Type]any)
	case *DisplayedColumn:
		return any(&stage.DisplayedColumns).(*map[Type]any)
	case *FormDiv:
		return any(&stage.FormDivs).(*map[Type]any)
	case *FormEditAssocButton:
		return any(&stage.FormEditAssocButtons).(*map[Type]any)
	case *FormField:
		return any(&stage.FormFields).(*map[Type]any)
	case *FormFieldDate:
		return any(&stage.FormFieldDates).(*map[Type]any)
	case *FormFieldDateTime:
		return any(&stage.FormFieldDateTimes).(*map[Type]any)
	case *FormFieldFloat64:
		return any(&stage.FormFieldFloat64s).(*map[Type]any)
	case *FormFieldInt:
		return any(&stage.FormFieldInts).(*map[Type]any)
	case *FormFieldSelect:
		return any(&stage.FormFieldSelects).(*map[Type]any)
	case *FormFieldString:
		return any(&stage.FormFieldStrings).(*map[Type]any)
	case *FormFieldTime:
		return any(&stage.FormFieldTimes).(*map[Type]any)
	case *FormGroup:
		return any(&stage.FormGroups).(*map[Type]any)
	case *FormSortAssocButton:
		return any(&stage.FormSortAssocButtons).(*map[Type]any)
	case *Option:
		return any(&stage.Options).(*map[Type]any)
	case *Row:
		return any(&stage.Rows).(*map[Type]any)
	case *Table:
		return any(&stage.Tables).(*map[Type]any)
	default:
		return nil
	}
}

// GetGongstructInstancesMap returns the map of staged GongstructType instances
// it is usefull because it allows refactoring of gong struct identifier
func GetGongstructInstancesMap[Type Gongstruct](stage *StageStruct) *map[string]*Type {
	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get functions
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
	case CheckBox:
		return any(&stage.CheckBoxs_mapString).(*map[string]*Type)
	case DisplayedColumn:
		return any(&stage.DisplayedColumns_mapString).(*map[string]*Type)
	case FormDiv:
		return any(&stage.FormDivs_mapString).(*map[string]*Type)
	case FormEditAssocButton:
		return any(&stage.FormEditAssocButtons_mapString).(*map[string]*Type)
	case FormField:
		return any(&stage.FormFields_mapString).(*map[string]*Type)
	case FormFieldDate:
		return any(&stage.FormFieldDates_mapString).(*map[string]*Type)
	case FormFieldDateTime:
		return any(&stage.FormFieldDateTimes_mapString).(*map[string]*Type)
	case FormFieldFloat64:
		return any(&stage.FormFieldFloat64s_mapString).(*map[string]*Type)
	case FormFieldInt:
		return any(&stage.FormFieldInts_mapString).(*map[string]*Type)
	case FormFieldSelect:
		return any(&stage.FormFieldSelects_mapString).(*map[string]*Type)
	case FormFieldString:
		return any(&stage.FormFieldStrings_mapString).(*map[string]*Type)
	case FormFieldTime:
		return any(&stage.FormFieldTimes_mapString).(*map[string]*Type)
	case FormGroup:
		return any(&stage.FormGroups_mapString).(*map[string]*Type)
	case FormSortAssocButton:
		return any(&stage.FormSortAssocButtons_mapString).(*map[string]*Type)
	case Option:
		return any(&stage.Options_mapString).(*map[string]*Type)
	case Row:
		return any(&stage.Rows_mapString).(*map[string]*Type)
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
	case CheckBox:
		return any(&CheckBox{
			// Initialisation of associations
		}).(*Type)
	case DisplayedColumn:
		return any(&DisplayedColumn{
			// Initialisation of associations
		}).(*Type)
	case FormDiv:
		return any(&FormDiv{
			// Initialisation of associations
			// field is initialized with an instance of FormField with the name of the field
			FormFields: []*FormField{{Name: "FormFields"}},
			// field is initialized with an instance of CheckBox with the name of the field
			CheckBoxs: []*CheckBox{{Name: "CheckBoxs"}},
			// field is initialized with an instance of FormEditAssocButton with the name of the field
			FormEditAssocButton: &FormEditAssocButton{Name: "FormEditAssocButton"},
			// field is initialized with an instance of FormSortAssocButton with the name of the field
			FormSortAssocButton: &FormSortAssocButton{Name: "FormSortAssocButton"},
		}).(*Type)
	case FormEditAssocButton:
		return any(&FormEditAssocButton{
			// Initialisation of associations
		}).(*Type)
	case FormField:
		return any(&FormField{
			// Initialisation of associations
			// field is initialized with an instance of FormFieldString with the name of the field
			FormFieldString: &FormFieldString{Name: "FormFieldString"},
			// field is initialized with an instance of FormFieldFloat64 with the name of the field
			FormFieldFloat64: &FormFieldFloat64{Name: "FormFieldFloat64"},
			// field is initialized with an instance of FormFieldInt with the name of the field
			FormFieldInt: &FormFieldInt{Name: "FormFieldInt"},
			// field is initialized with an instance of FormFieldDate with the name of the field
			FormFieldDate: &FormFieldDate{Name: "FormFieldDate"},
			// field is initialized with an instance of FormFieldTime with the name of the field
			FormFieldTime: &FormFieldTime{Name: "FormFieldTime"},
			// field is initialized with an instance of FormFieldDateTime with the name of the field
			FormFieldDateTime: &FormFieldDateTime{Name: "FormFieldDateTime"},
			// field is initialized with an instance of FormFieldSelect with the name of the field
			FormFieldSelect: &FormFieldSelect{Name: "FormFieldSelect"},
		}).(*Type)
	case FormFieldDate:
		return any(&FormFieldDate{
			// Initialisation of associations
		}).(*Type)
	case FormFieldDateTime:
		return any(&FormFieldDateTime{
			// Initialisation of associations
		}).(*Type)
	case FormFieldFloat64:
		return any(&FormFieldFloat64{
			// Initialisation of associations
		}).(*Type)
	case FormFieldInt:
		return any(&FormFieldInt{
			// Initialisation of associations
		}).(*Type)
	case FormFieldSelect:
		return any(&FormFieldSelect{
			// Initialisation of associations
			// field is initialized with an instance of Option with the name of the field
			Value: &Option{Name: "Value"},
			// field is initialized with an instance of Option with the name of the field
			Options: []*Option{{Name: "Options"}},
		}).(*Type)
	case FormFieldString:
		return any(&FormFieldString{
			// Initialisation of associations
		}).(*Type)
	case FormFieldTime:
		return any(&FormFieldTime{
			// Initialisation of associations
		}).(*Type)
	case FormGroup:
		return any(&FormGroup{
			// Initialisation of associations
			// field is initialized with an instance of FormDiv with the name of the field
			FormDivs: []*FormDiv{{Name: "FormDivs"}},
		}).(*Type)
	case FormSortAssocButton:
		return any(&FormSortAssocButton{
			// Initialisation of associations
		}).(*Type)
	case Option:
		return any(&Option{
			// Initialisation of associations
		}).(*Type)
	case Row:
		return any(&Row{
			// Initialisation of associations
			// field is initialized with an instance of Cell with the name of the field
			Cells: []*Cell{{Name: "Cells"}},
		}).(*Type)
	case Table:
		return any(&Table{
			// Initialisation of associations
			// field is initialized with an instance of DisplayedColumn with the name of the field
			DisplayedColumns: []*DisplayedColumn{{Name: "DisplayedColumns"}},
			// field is initialized with an instance of Row with the name of the field
			Rows: []*Row{{Name: "Rows"}},
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
func GetPointerReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End][]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
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
	// reverse maps of direct associations of CheckBox
	case CheckBox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DisplayedColumn
	case DisplayedColumn:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormDiv
	case FormDiv:
		switch fieldname {
		// insertion point for per direct association field
		case "FormEditAssocButton":
			res := make(map[*FormEditAssocButton][]*FormDiv)
			for formdiv := range stage.FormDivs {
				if formdiv.FormEditAssocButton != nil {
					formeditassocbutton_ := formdiv.FormEditAssocButton
					var formdivs []*FormDiv
					_, ok := res[formeditassocbutton_]
					if ok {
						formdivs = res[formeditassocbutton_]
					} else {
						formdivs = make([]*FormDiv, 0)
					}
					formdivs = append(formdivs, formdiv)
					res[formeditassocbutton_] = formdivs
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormSortAssocButton":
			res := make(map[*FormSortAssocButton][]*FormDiv)
			for formdiv := range stage.FormDivs {
				if formdiv.FormSortAssocButton != nil {
					formsortassocbutton_ := formdiv.FormSortAssocButton
					var formdivs []*FormDiv
					_, ok := res[formsortassocbutton_]
					if ok {
						formdivs = res[formsortassocbutton_]
					} else {
						formdivs = make([]*FormDiv, 0)
					}
					formdivs = append(formdivs, formdiv)
					res[formsortassocbutton_] = formdivs
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormEditAssocButton
	case FormEditAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormField
	case FormField:
		switch fieldname {
		// insertion point for per direct association field
		case "FormFieldString":
			res := make(map[*FormFieldString][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldString != nil {
					formfieldstring_ := formfield.FormFieldString
					var formfields []*FormField
					_, ok := res[formfieldstring_]
					if ok {
						formfields = res[formfieldstring_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldstring_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldFloat64":
			res := make(map[*FormFieldFloat64][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldFloat64 != nil {
					formfieldfloat64_ := formfield.FormFieldFloat64
					var formfields []*FormField
					_, ok := res[formfieldfloat64_]
					if ok {
						formfields = res[formfieldfloat64_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldfloat64_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldInt":
			res := make(map[*FormFieldInt][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldInt != nil {
					formfieldint_ := formfield.FormFieldInt
					var formfields []*FormField
					_, ok := res[formfieldint_]
					if ok {
						formfields = res[formfieldint_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldint_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldDate":
			res := make(map[*FormFieldDate][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldDate != nil {
					formfielddate_ := formfield.FormFieldDate
					var formfields []*FormField
					_, ok := res[formfielddate_]
					if ok {
						formfields = res[formfielddate_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfielddate_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldTime":
			res := make(map[*FormFieldTime][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldTime != nil {
					formfieldtime_ := formfield.FormFieldTime
					var formfields []*FormField
					_, ok := res[formfieldtime_]
					if ok {
						formfields = res[formfieldtime_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldtime_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldDateTime":
			res := make(map[*FormFieldDateTime][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldDateTime != nil {
					formfielddatetime_ := formfield.FormFieldDateTime
					var formfields []*FormField
					_, ok := res[formfielddatetime_]
					if ok {
						formfields = res[formfielddatetime_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfielddatetime_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		case "FormFieldSelect":
			res := make(map[*FormFieldSelect][]*FormField)
			for formfield := range stage.FormFields {
				if formfield.FormFieldSelect != nil {
					formfieldselect_ := formfield.FormFieldSelect
					var formfields []*FormField
					_, ok := res[formfieldselect_]
					if ok {
						formfields = res[formfieldselect_]
					} else {
						formfields = make([]*FormField, 0)
					}
					formfields = append(formfields, formfield)
					res[formfieldselect_] = formfields
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormFieldDate
	case FormFieldDate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDateTime
	case FormFieldDateTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldFloat64
	case FormFieldFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldInt
	case FormFieldInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldSelect
	case FormFieldSelect:
		switch fieldname {
		// insertion point for per direct association field
		case "Value":
			res := make(map[*Option][]*FormFieldSelect)
			for formfieldselect := range stage.FormFieldSelects {
				if formfieldselect.Value != nil {
					option_ := formfieldselect.Value
					var formfieldselects []*FormFieldSelect
					_, ok := res[option_]
					if ok {
						formfieldselects = res[option_]
					} else {
						formfieldselects = make([]*FormFieldSelect, 0)
					}
					formfieldselects = append(formfieldselects, formfieldselect)
					res[option_] = formfieldselects
				}
			}
			return any(res).(map[*End][]*Start)
		}
	// reverse maps of direct associations of FormFieldString
	case FormFieldString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldTime
	case FormFieldTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormGroup
	case FormGroup:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormSortAssocButton
	case FormSortAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Option
	case Option:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Row
	case Row:
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
func GetSliceOfPointersReverseMap[Start, End Gongstruct](fieldname string, stage *StageStruct) map[*End]*Start {

	var ret Start

	switch any(ret).(type) {
	// insertion point of functions that provide maps for reverse associations
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
	// reverse maps of direct associations of CheckBox
	case CheckBox:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of DisplayedColumn
	case DisplayedColumn:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormDiv
	case FormDiv:
		switch fieldname {
		// insertion point for per direct association field
		case "FormFields":
			res := make(map[*FormField]*FormDiv)
			for formdiv := range stage.FormDivs {
				for _, formfield_ := range formdiv.FormFields {
					res[formfield_] = formdiv
				}
			}
			return any(res).(map[*End]*Start)
		case "CheckBoxs":
			res := make(map[*CheckBox]*FormDiv)
			for formdiv := range stage.FormDivs {
				for _, checkbox_ := range formdiv.CheckBoxs {
					res[checkbox_] = formdiv
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of FormEditAssocButton
	case FormEditAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormField
	case FormField:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDate
	case FormFieldDate:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldDateTime
	case FormFieldDateTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldFloat64
	case FormFieldFloat64:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldInt
	case FormFieldInt:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldSelect
	case FormFieldSelect:
		switch fieldname {
		// insertion point for per direct association field
		case "Options":
			res := make(map[*Option]*FormFieldSelect)
			for formfieldselect := range stage.FormFieldSelects {
				for _, option_ := range formfieldselect.Options {
					res[option_] = formfieldselect
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of FormFieldString
	case FormFieldString:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormFieldTime
	case FormFieldTime:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of FormGroup
	case FormGroup:
		switch fieldname {
		// insertion point for per direct association field
		case "FormDivs":
			res := make(map[*FormDiv]*FormGroup)
			for formgroup := range stage.FormGroups {
				for _, formdiv_ := range formgroup.FormDivs {
					res[formdiv_] = formgroup
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of FormSortAssocButton
	case FormSortAssocButton:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Option
	case Option:
		switch fieldname {
		// insertion point for per direct association field
		}
	// reverse maps of direct associations of Row
	case Row:
		switch fieldname {
		// insertion point for per direct association field
		case "Cells":
			res := make(map[*Cell]*Row)
			for row := range stage.Rows {
				for _, cell_ := range row.Cells {
					res[cell_] = row
				}
			}
			return any(res).(map[*End]*Start)
		}
	// reverse maps of direct associations of Table
	case Table:
		switch fieldname {
		// insertion point for per direct association field
		case "DisplayedColumns":
			res := make(map[*DisplayedColumn]*Table)
			for table := range stage.Tables {
				for _, displayedcolumn_ := range table.DisplayedColumns {
					res[displayedcolumn_] = table
				}
			}
			return any(res).(map[*End]*Start)
		case "Rows":
			res := make(map[*Row]*Table)
			for table := range stage.Tables {
				for _, row_ := range table.Rows {
					res[row_] = table
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
	case Cell:
		res = "Cell"
	case CellBoolean:
		res = "CellBoolean"
	case CellFloat64:
		res = "CellFloat64"
	case CellIcon:
		res = "CellIcon"
	case CellInt:
		res = "CellInt"
	case CellString:
		res = "CellString"
	case CheckBox:
		res = "CheckBox"
	case DisplayedColumn:
		res = "DisplayedColumn"
	case FormDiv:
		res = "FormDiv"
	case FormEditAssocButton:
		res = "FormEditAssocButton"
	case FormField:
		res = "FormField"
	case FormFieldDate:
		res = "FormFieldDate"
	case FormFieldDateTime:
		res = "FormFieldDateTime"
	case FormFieldFloat64:
		res = "FormFieldFloat64"
	case FormFieldInt:
		res = "FormFieldInt"
	case FormFieldSelect:
		res = "FormFieldSelect"
	case FormFieldString:
		res = "FormFieldString"
	case FormFieldTime:
		res = "FormFieldTime"
	case FormGroup:
		res = "FormGroup"
	case FormSortAssocButton:
		res = "FormSortAssocButton"
	case Option:
		res = "Option"
	case Row:
		res = "Row"
	case Table:
		res = "Table"
	}
	return res
}

// GetPointerToGongstructName returns the name of the Gongstruct
// this can be usefull if one want program robust to refactoring
func GetPointerToGongstructName[Type PointerToGongstruct]() (res string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
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
	case *CheckBox:
		res = "CheckBox"
	case *DisplayedColumn:
		res = "DisplayedColumn"
	case *FormDiv:
		res = "FormDiv"
	case *FormEditAssocButton:
		res = "FormEditAssocButton"
	case *FormField:
		res = "FormField"
	case *FormFieldDate:
		res = "FormFieldDate"
	case *FormFieldDateTime:
		res = "FormFieldDateTime"
	case *FormFieldFloat64:
		res = "FormFieldFloat64"
	case *FormFieldInt:
		res = "FormFieldInt"
	case *FormFieldSelect:
		res = "FormFieldSelect"
	case *FormFieldString:
		res = "FormFieldString"
	case *FormFieldTime:
		res = "FormFieldTime"
	case *FormGroup:
		res = "FormGroup"
	case *FormSortAssocButton:
		res = "FormSortAssocButton"
	case *Option:
		res = "Option"
	case *Row:
		res = "Row"
	case *Table:
		res = "Table"
	}
	return res
}

// GetFields return the array of the fields
func GetFields[Type Gongstruct]() (res []string) {

	var ret Type

	switch any(ret).(type) {
	// insertion point for generic get gongstruct name
	case Cell:
		res = []string{"Name", "CellString", "CellFloat64", "CellInt", "CellBool", "CellIcon"}
	case CellBoolean:
		res = []string{"Name", "Value"}
	case CellFloat64:
		res = []string{"Name", "Value"}
	case CellIcon:
		res = []string{"Name", "Icon"}
	case CellInt:
		res = []string{"Name", "Value"}
	case CellString:
		res = []string{"Name", "Value"}
	case CheckBox:
		res = []string{"Name", "Value"}
	case DisplayedColumn:
		res = []string{"Name"}
	case FormDiv:
		res = []string{"Name", "FormFields", "CheckBoxs", "FormEditAssocButton", "FormSortAssocButton"}
	case FormEditAssocButton:
		res = []string{"Name", "Label"}
	case FormField:
		res = []string{"Name", "InputTypeEnum", "Label", "Placeholder", "FormFieldString", "FormFieldFloat64", "FormFieldInt", "FormFieldDate", "FormFieldTime", "FormFieldDateTime", "FormFieldSelect", "HasBespokeWidth", "BespokeWidthPx", "HasBespokeHeight", "BespokeHeightPx"}
	case FormFieldDate:
		res = []string{"Name", "Value"}
	case FormFieldDateTime:
		res = []string{"Name", "Value"}
	case FormFieldFloat64:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case FormFieldInt:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case FormFieldSelect:
		res = []string{"Name", "Value", "Options", "CanBeEmpty"}
	case FormFieldString:
		res = []string{"Name", "Value", "IsTextArea"}
	case FormFieldTime:
		res = []string{"Name", "Value", "Step"}
	case FormGroup:
		res = []string{"Name", "Label", "FormDivs", "HasSuppressButton", "HasSuppressButtonBeenPressed"}
	case FormSortAssocButton:
		res = []string{"Name", "Label"}
	case Option:
		res = []string{"Name"}
	case Row:
		res = []string{"Name", "Cells", "IsChecked"}
	case Table:
		res = []string{"Name", "DisplayedColumns", "Rows", "HasFiltering", "HasColumnSorting", "HasPaginator", "HasCheckableRows", "HasSaveButton", "CanDragDropRows", "HasCloseButton", "SavingInProgress", "NbOfStickyColumns"}
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
	case Cell:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Row"
		rf.Fieldname = "Cells"
		res = append(res, rf)
	case CellBoolean:
		var rf ReverseField
		_ = rf
	case CellFloat64:
		var rf ReverseField
		_ = rf
	case CellIcon:
		var rf ReverseField
		_ = rf
	case CellInt:
		var rf ReverseField
		_ = rf
	case CellString:
		var rf ReverseField
		_ = rf
	case CheckBox:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormDiv"
		rf.Fieldname = "CheckBoxs"
		res = append(res, rf)
	case DisplayedColumn:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Table"
		rf.Fieldname = "DisplayedColumns"
		res = append(res, rf)
	case FormDiv:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormGroup"
		rf.Fieldname = "FormDivs"
		res = append(res, rf)
	case FormEditAssocButton:
		var rf ReverseField
		_ = rf
	case FormField:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormDiv"
		rf.Fieldname = "FormFields"
		res = append(res, rf)
	case FormFieldDate:
		var rf ReverseField
		_ = rf
	case FormFieldDateTime:
		var rf ReverseField
		_ = rf
	case FormFieldFloat64:
		var rf ReverseField
		_ = rf
	case FormFieldInt:
		var rf ReverseField
		_ = rf
	case FormFieldSelect:
		var rf ReverseField
		_ = rf
	case FormFieldString:
		var rf ReverseField
		_ = rf
	case FormFieldTime:
		var rf ReverseField
		_ = rf
	case FormGroup:
		var rf ReverseField
		_ = rf
	case FormSortAssocButton:
		var rf ReverseField
		_ = rf
	case Option:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "FormFieldSelect"
		rf.Fieldname = "Options"
		res = append(res, rf)
	case Row:
		var rf ReverseField
		_ = rf
		rf.GongstructName = "Table"
		rf.Fieldname = "Rows"
		res = append(res, rf)
	case Table:
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
	case *Cell:
		res = []string{"Name", "CellString", "CellFloat64", "CellInt", "CellBool", "CellIcon"}
	case *CellBoolean:
		res = []string{"Name", "Value"}
	case *CellFloat64:
		res = []string{"Name", "Value"}
	case *CellIcon:
		res = []string{"Name", "Icon"}
	case *CellInt:
		res = []string{"Name", "Value"}
	case *CellString:
		res = []string{"Name", "Value"}
	case *CheckBox:
		res = []string{"Name", "Value"}
	case *DisplayedColumn:
		res = []string{"Name"}
	case *FormDiv:
		res = []string{"Name", "FormFields", "CheckBoxs", "FormEditAssocButton", "FormSortAssocButton"}
	case *FormEditAssocButton:
		res = []string{"Name", "Label"}
	case *FormField:
		res = []string{"Name", "InputTypeEnum", "Label", "Placeholder", "FormFieldString", "FormFieldFloat64", "FormFieldInt", "FormFieldDate", "FormFieldTime", "FormFieldDateTime", "FormFieldSelect", "HasBespokeWidth", "BespokeWidthPx", "HasBespokeHeight", "BespokeHeightPx"}
	case *FormFieldDate:
		res = []string{"Name", "Value"}
	case *FormFieldDateTime:
		res = []string{"Name", "Value"}
	case *FormFieldFloat64:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case *FormFieldInt:
		res = []string{"Name", "Value", "HasMinValidator", "MinValue", "HasMaxValidator", "MaxValue"}
	case *FormFieldSelect:
		res = []string{"Name", "Value", "Options", "CanBeEmpty"}
	case *FormFieldString:
		res = []string{"Name", "Value", "IsTextArea"}
	case *FormFieldTime:
		res = []string{"Name", "Value", "Step"}
	case *FormGroup:
		res = []string{"Name", "Label", "FormDivs", "HasSuppressButton", "HasSuppressButtonBeenPressed"}
	case *FormSortAssocButton:
		res = []string{"Name", "Label"}
	case *Option:
		res = []string{"Name"}
	case *Row:
		res = []string{"Name", "Cells", "IsChecked"}
	case *Table:
		res = []string{"Name", "DisplayedColumns", "Rows", "HasFiltering", "HasColumnSorting", "HasPaginator", "HasCheckableRows", "HasSaveButton", "CanDragDropRows", "HasCloseButton", "SavingInProgress", "NbOfStickyColumns"}
	}
	return
}

func GetFieldStringValueFromPointer[Type PointerToGongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case *Cell:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CellString":
			if inferedInstance.CellString != nil {
				res = inferedInstance.CellString.Name
			}
		case "CellFloat64":
			if inferedInstance.CellFloat64 != nil {
				res = inferedInstance.CellFloat64.Name
			}
		case "CellInt":
			if inferedInstance.CellInt != nil {
				res = inferedInstance.CellInt.Name
			}
		case "CellBool":
			if inferedInstance.CellBool != nil {
				res = inferedInstance.CellBool.Name
			}
		case "CellIcon":
			if inferedInstance.CellIcon != nil {
				res = inferedInstance.CellIcon.Name
			}
		}
	case *CellBoolean:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%t", inferedInstance.Value)
		}
	case *CellFloat64:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%f", inferedInstance.Value)
		}
	case *CellIcon:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Icon":
			res = inferedInstance.Icon
		}
	case *CellInt:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%d", inferedInstance.Value)
		}
	case *CellString:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case *CheckBox:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%t", inferedInstance.Value)
		}
	case *DisplayedColumn:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *FormDiv:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "FormFields":
			for idx, __instance__ := range inferedInstance.FormFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CheckBoxs":
			for idx, __instance__ := range inferedInstance.CheckBoxs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "FormEditAssocButton":
			if inferedInstance.FormEditAssocButton != nil {
				res = inferedInstance.FormEditAssocButton.Name
			}
		case "FormSortAssocButton":
			if inferedInstance.FormSortAssocButton != nil {
				res = inferedInstance.FormSortAssocButton.Name
			}
		}
	case *FormEditAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case *FormField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "InputTypeEnum":
			enum := inferedInstance.InputTypeEnum
			res = enum.ToCodeString()
		case "Label":
			res = inferedInstance.Label
		case "Placeholder":
			res = inferedInstance.Placeholder
		case "FormFieldString":
			if inferedInstance.FormFieldString != nil {
				res = inferedInstance.FormFieldString.Name
			}
		case "FormFieldFloat64":
			if inferedInstance.FormFieldFloat64 != nil {
				res = inferedInstance.FormFieldFloat64.Name
			}
		case "FormFieldInt":
			if inferedInstance.FormFieldInt != nil {
				res = inferedInstance.FormFieldInt.Name
			}
		case "FormFieldDate":
			if inferedInstance.FormFieldDate != nil {
				res = inferedInstance.FormFieldDate.Name
			}
		case "FormFieldTime":
			if inferedInstance.FormFieldTime != nil {
				res = inferedInstance.FormFieldTime.Name
			}
		case "FormFieldDateTime":
			if inferedInstance.FormFieldDateTime != nil {
				res = inferedInstance.FormFieldDateTime.Name
			}
		case "FormFieldSelect":
			if inferedInstance.FormFieldSelect != nil {
				res = inferedInstance.FormFieldSelect.Name
			}
		case "HasBespokeWidth":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeWidth)
		case "BespokeWidthPx":
			res = fmt.Sprintf("%d", inferedInstance.BespokeWidthPx)
		case "HasBespokeHeight":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeHeight)
		case "BespokeHeightPx":
			res = fmt.Sprintf("%d", inferedInstance.BespokeHeightPx)
		}
	case *FormFieldDate:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case *FormFieldDateTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case *FormFieldFloat64:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%f", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%f", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%f", inferedInstance.MaxValue)
		}
	case *FormFieldInt:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%d", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%d", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%d", inferedInstance.MaxValue)
		}
	case *FormFieldSelect:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			if inferedInstance.Value != nil {
				res = inferedInstance.Value.Name
			}
		case "Options":
			for idx, __instance__ := range inferedInstance.Options {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CanBeEmpty":
			res = fmt.Sprintf("%t", inferedInstance.CanBeEmpty)
		}
	case *FormFieldString:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		case "IsTextArea":
			res = fmt.Sprintf("%t", inferedInstance.IsTextArea)
		}
	case *FormFieldTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		case "Step":
			res = fmt.Sprintf("%f", inferedInstance.Step)
		}
	case *FormGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		case "FormDivs":
			for idx, __instance__ := range inferedInstance.FormDivs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "HasSuppressButton":
			res = fmt.Sprintf("%t", inferedInstance.HasSuppressButton)
		case "HasSuppressButtonBeenPressed":
			res = fmt.Sprintf("%t", inferedInstance.HasSuppressButtonBeenPressed)
		}
	case *FormSortAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case *Option:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case *Row:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Cells":
			for idx, __instance__ := range inferedInstance.Cells {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "IsChecked":
			res = fmt.Sprintf("%t", inferedInstance.IsChecked)
		}
	case *Table:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DisplayedColumns":
			for idx, __instance__ := range inferedInstance.DisplayedColumns {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Rows":
			for idx, __instance__ := range inferedInstance.Rows {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "HasFiltering":
			res = fmt.Sprintf("%t", inferedInstance.HasFiltering)
		case "HasColumnSorting":
			res = fmt.Sprintf("%t", inferedInstance.HasColumnSorting)
		case "HasPaginator":
			res = fmt.Sprintf("%t", inferedInstance.HasPaginator)
		case "HasCheckableRows":
			res = fmt.Sprintf("%t", inferedInstance.HasCheckableRows)
		case "HasSaveButton":
			res = fmt.Sprintf("%t", inferedInstance.HasSaveButton)
		case "CanDragDropRows":
			res = fmt.Sprintf("%t", inferedInstance.CanDragDropRows)
		case "HasCloseButton":
			res = fmt.Sprintf("%t", inferedInstance.HasCloseButton)
		case "SavingInProgress":
			res = fmt.Sprintf("%t", inferedInstance.SavingInProgress)
		case "NbOfStickyColumns":
			res = fmt.Sprintf("%d", inferedInstance.NbOfStickyColumns)
		}
	default:
		_ = inferedInstance
	}
	return
}

func GetFieldStringValue[Type Gongstruct](instance Type, fieldName string) (res string) {

	switch inferedInstance := any(instance).(type) {
	// insertion point for generic get gongstruct field value
	case Cell:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "CellString":
			if inferedInstance.CellString != nil {
				res = inferedInstance.CellString.Name
			}
		case "CellFloat64":
			if inferedInstance.CellFloat64 != nil {
				res = inferedInstance.CellFloat64.Name
			}
		case "CellInt":
			if inferedInstance.CellInt != nil {
				res = inferedInstance.CellInt.Name
			}
		case "CellBool":
			if inferedInstance.CellBool != nil {
				res = inferedInstance.CellBool.Name
			}
		case "CellIcon":
			if inferedInstance.CellIcon != nil {
				res = inferedInstance.CellIcon.Name
			}
		}
	case CellBoolean:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%t", inferedInstance.Value)
		}
	case CellFloat64:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%f", inferedInstance.Value)
		}
	case CellIcon:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Icon":
			res = inferedInstance.Icon
		}
	case CellInt:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%d", inferedInstance.Value)
		}
	case CellString:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		}
	case CheckBox:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%t", inferedInstance.Value)
		}
	case DisplayedColumn:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case FormDiv:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "FormFields":
			for idx, __instance__ := range inferedInstance.FormFields {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CheckBoxs":
			for idx, __instance__ := range inferedInstance.CheckBoxs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "FormEditAssocButton":
			if inferedInstance.FormEditAssocButton != nil {
				res = inferedInstance.FormEditAssocButton.Name
			}
		case "FormSortAssocButton":
			if inferedInstance.FormSortAssocButton != nil {
				res = inferedInstance.FormSortAssocButton.Name
			}
		}
	case FormEditAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case FormField:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "InputTypeEnum":
			enum := inferedInstance.InputTypeEnum
			res = enum.ToCodeString()
		case "Label":
			res = inferedInstance.Label
		case "Placeholder":
			res = inferedInstance.Placeholder
		case "FormFieldString":
			if inferedInstance.FormFieldString != nil {
				res = inferedInstance.FormFieldString.Name
			}
		case "FormFieldFloat64":
			if inferedInstance.FormFieldFloat64 != nil {
				res = inferedInstance.FormFieldFloat64.Name
			}
		case "FormFieldInt":
			if inferedInstance.FormFieldInt != nil {
				res = inferedInstance.FormFieldInt.Name
			}
		case "FormFieldDate":
			if inferedInstance.FormFieldDate != nil {
				res = inferedInstance.FormFieldDate.Name
			}
		case "FormFieldTime":
			if inferedInstance.FormFieldTime != nil {
				res = inferedInstance.FormFieldTime.Name
			}
		case "FormFieldDateTime":
			if inferedInstance.FormFieldDateTime != nil {
				res = inferedInstance.FormFieldDateTime.Name
			}
		case "FormFieldSelect":
			if inferedInstance.FormFieldSelect != nil {
				res = inferedInstance.FormFieldSelect.Name
			}
		case "HasBespokeWidth":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeWidth)
		case "BespokeWidthPx":
			res = fmt.Sprintf("%d", inferedInstance.BespokeWidthPx)
		case "HasBespokeHeight":
			res = fmt.Sprintf("%t", inferedInstance.HasBespokeHeight)
		case "BespokeHeightPx":
			res = fmt.Sprintf("%d", inferedInstance.BespokeHeightPx)
		}
	case FormFieldDate:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case FormFieldDateTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		}
	case FormFieldFloat64:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%f", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%f", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%f", inferedInstance.MaxValue)
		}
	case FormFieldInt:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = fmt.Sprintf("%d", inferedInstance.Value)
		case "HasMinValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMinValidator)
		case "MinValue":
			res = fmt.Sprintf("%d", inferedInstance.MinValue)
		case "HasMaxValidator":
			res = fmt.Sprintf("%t", inferedInstance.HasMaxValidator)
		case "MaxValue":
			res = fmt.Sprintf("%d", inferedInstance.MaxValue)
		}
	case FormFieldSelect:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			if inferedInstance.Value != nil {
				res = inferedInstance.Value.Name
			}
		case "Options":
			for idx, __instance__ := range inferedInstance.Options {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "CanBeEmpty":
			res = fmt.Sprintf("%t", inferedInstance.CanBeEmpty)
		}
	case FormFieldString:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value
		case "IsTextArea":
			res = fmt.Sprintf("%t", inferedInstance.IsTextArea)
		}
	case FormFieldTime:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Value":
			res = inferedInstance.Value.String()
		case "Step":
			res = fmt.Sprintf("%f", inferedInstance.Step)
		}
	case FormGroup:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		case "FormDivs":
			for idx, __instance__ := range inferedInstance.FormDivs {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "HasSuppressButton":
			res = fmt.Sprintf("%t", inferedInstance.HasSuppressButton)
		case "HasSuppressButtonBeenPressed":
			res = fmt.Sprintf("%t", inferedInstance.HasSuppressButtonBeenPressed)
		}
	case FormSortAssocButton:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Label":
			res = inferedInstance.Label
		}
	case Option:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		}
	case Row:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "Cells":
			for idx, __instance__ := range inferedInstance.Cells {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "IsChecked":
			res = fmt.Sprintf("%t", inferedInstance.IsChecked)
		}
	case Table:
		switch fieldName {
		// string value of fields
		case "Name":
			res = inferedInstance.Name
		case "DisplayedColumns":
			for idx, __instance__ := range inferedInstance.DisplayedColumns {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "Rows":
			for idx, __instance__ := range inferedInstance.Rows {
				if idx > 0 {
					res += "\n"
				}
				res += __instance__.Name
			}
		case "HasFiltering":
			res = fmt.Sprintf("%t", inferedInstance.HasFiltering)
		case "HasColumnSorting":
			res = fmt.Sprintf("%t", inferedInstance.HasColumnSorting)
		case "HasPaginator":
			res = fmt.Sprintf("%t", inferedInstance.HasPaginator)
		case "HasCheckableRows":
			res = fmt.Sprintf("%t", inferedInstance.HasCheckableRows)
		case "HasSaveButton":
			res = fmt.Sprintf("%t", inferedInstance.HasSaveButton)
		case "CanDragDropRows":
			res = fmt.Sprintf("%t", inferedInstance.CanDragDropRows)
		case "HasCloseButton":
			res = fmt.Sprintf("%t", inferedInstance.HasCloseButton)
		case "SavingInProgress":
			res = fmt.Sprintf("%t", inferedInstance.SavingInProgress)
		case "NbOfStickyColumns":
			res = fmt.Sprintf("%d", inferedInstance.NbOfStickyColumns)
		}
	default:
		_ = inferedInstance
	}
	return
}

// Last line of the template
