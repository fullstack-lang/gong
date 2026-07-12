// generated code - do not edit
package models

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *Stage, filename string) {
	SerializeStage2(stage, filename, false)
}

func SerializeStage2(stage *Stage, filename string, addIDs bool) {
	f := buildExcelizeFile(stage, addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func buildExcelizeFile(stage *Stage, addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelizePointerToGongstruct2[*ArcNormalVectorShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*ArcNormalVectorShapeGrid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*AxesShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*BaseVectorShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*BaseVectorShapeGrid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*CircleGridShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*EndArcShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*EndArcShapeGrid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*EndArcShapeV2](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*EndArcShapeV2Grid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*ExplanationTextShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*GridPathShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*GrowthCurveBezierShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*GrowthCurveBezierShapeGrid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*GrowthCurveRhombusGridShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*GrowthCurveRhombusShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*GrowthVectorShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*InitialRhombusGridShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*InitialRhombusShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Library](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*NextCircleShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*PerpendicularVector](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*PerpendicularVectorGrid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*PerpendicularVectorGridHalfway](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*PerpendicularVectorHalfway](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Plant](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*PlantCircumferenceShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*PlantDiagram](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Rendered3DShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*RhombusShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*RotatedRhombusGridShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*RotatedRhombusShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*StackGrowthCurveBezierShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*StackOfGrowthCurve](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*StartArcShape](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*StartArcShapeGrid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*StartArcShapeV2](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*StartArcShapeV2Grid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*TopEndArcShapeV2](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*TopEndArcShapeV2Grid](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*TopStartArcShapeV2](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*TopStartArcShapeV2Grid](stage, f, addIDs)
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func SerializeStageAsBytes(stage *Stage, addIDs bool) ([]byte, error) {
	f := buildExcelizeFile(stage, addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
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

func SerializeExcelizePointerToGongstruct[Type PointerToGongstruct](stage *Stage, f *excelize.File) {
	SerializeExcelizePointerToGongstruct2[Type](stage, f, false)
}

func SerializeExcelizePointerToGongstruct2[Type PointerToGongstruct](stage *Stage, f *excelize.File, addIDs bool) {
	sheetName := GetPointerToGongstructName[Type]()

	sheetName = shortenString(sheetName)

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *GetGongstructInstancesSetFromPointerType[Type](stage)

	var sortedSlice []Type
	for key := range set {
		sortedSlice = append(sortedSlice, key)
	}
	slices.SortFunc(sortedSlice, func(a, b Type) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range GetFieldsFromPointer[Type]() {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", IntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range GetFieldsFromPointer[Type]() {
			fieldStringValue := GetFieldStringValueFromPointer(instance, fieldName.Name, stage)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}

	// // Autofit all columns according to their text content
	// cols, err := f.GetCols(sheetName)
	// if err != nil {
	// 	log.Panicln("SerializeExcelize")
	// }
	// for idx, col := range cols {
	// 	largestWidth := 0
	// 	for _, rowCell := range col {
	// 		cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
	// 		if cellWidth > largestWidth {
	// 			largestWidth = cellWidth
	// 		}
	// 	}
	// 	name, err := excelize.ColumnNumberToName(idx + 1)
	// 	if err != nil {
	// 		log.Panicln("SerializeExcelize")
	// 	}
	// 	f.SetColWidth(sheetName, name, name, float64(largestWidth))
	// }
}
