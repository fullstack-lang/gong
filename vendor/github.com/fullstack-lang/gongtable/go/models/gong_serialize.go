// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *StageStruct, filename string) {

	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelize[Cell](stage, f)
		SerializeExcelize[CellBoolean](stage, f)
		SerializeExcelize[CellFloat64](stage, f)
		SerializeExcelize[CellIcon](stage, f)
		SerializeExcelize[CellInt](stage, f)
		SerializeExcelize[CellString](stage, f)
		SerializeExcelize[CheckBox](stage, f)
		SerializeExcelize[DisplayedColumn](stage, f)
		SerializeExcelize[FormDiv](stage, f)
		SerializeExcelize[FormEditAssocButton](stage, f)
		SerializeExcelize[FormField](stage, f)
		SerializeExcelize[FormFieldDate](stage, f)
		SerializeExcelize[FormFieldDateTime](stage, f)
		SerializeExcelize[FormFieldFloat64](stage, f)
		SerializeExcelize[FormFieldInt](stage, f)
		SerializeExcelize[FormFieldSelect](stage, f)
		SerializeExcelize[FormFieldString](stage, f)
		SerializeExcelize[FormFieldTime](stage, f)
		SerializeExcelize[FormGroup](stage, f)
		SerializeExcelize[FormSortAssocButton](stage, f)
		SerializeExcelize[Option](stage, f)
		SerializeExcelize[Row](stage, f)
		SerializeExcelize[Table](stage, f)
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
		if err := f.SaveAs(filename); err != nil {
			fmt.Println("cannot write xl file : ", err)
		}
	}

}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

func Serialize[Type Gongstruct](stage *StageStruct, tab Tabulator) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	tab.AddSheet(sheetName)

	headerRowIndex := tab.AddRow(sheetName)
	for colIndex, fieldName := range GetFields[Type]() {
		tab.AddCell(sheetName, headerRowIndex, colIndex, fieldName)
		// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}

	set := *GetGongstructInstancesSet[Type](stage)
	for instance := range set {
		line := tab.AddRow(sheetName)
		for index, fieldName := range GetFields[Type]() {
			tab.AddCell(sheetName, line, index, GetFieldStringValue(
				any(*instance).(Type), fieldName).valueString)
			// f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), GetFieldStringValue(
			// 	any(*instance).(Type), fieldName))
		}
	}
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

func SerializeExcelize[Type Gongstruct](stage *StageStruct, f *excelize.File) {
	sheetName := GetGongstructName[Type]()

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *GetGongstructInstancesSet[Type](stage)
	line := 1

	for index, fieldName := range GetFields[Type]() {
		f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldName)
	}
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", IntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for instance := range set {
		line = line + 1
		for index, fieldName := range GetFields[Type]() {
			fieldStringValue := GetFieldStringValue(any(*instance).(Type), fieldName)
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
		}
	}

	// Autofit all columns according to their text content
	cols, err := f.GetCols(sheetName)
	if err != nil {
		log.Panicln("SerializeExcelize")
	}
	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			log.Panicln("SerializeExcelize")
		}
		f.SetColWidth(sheetName, name, name, float64(largestWidth))
	}
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}
