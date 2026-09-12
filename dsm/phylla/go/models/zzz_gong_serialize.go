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
		stage.SerializeExcelizePointer2[*Angle0Shape](f, addIDs)
		stage.SerializeExcelizePointer2[*ArcNormalVectorShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ArcNormalVectorShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*AxesShape](f, addIDs)
		stage.SerializeExcelizePointer2[*BaseVectorShape](f, addIDs)
		stage.SerializeExcelizePointer2[*BaseVectorShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*ChosenP1P2PairShape](f, addIDs)
		stage.SerializeExcelizePointer2[*CircleGridShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Circumference3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Clock2DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*Clock3DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*ClockAbstract](f, addIDs)
		stage.SerializeExcelizePointer2[*ClockTopCurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*CutLine3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EndArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EndArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*EndHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EndHalfwayArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*ExplanationTextShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Eye3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EyeCornersSampledPoints3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EyeSampledPoints3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EyeSeatBottomCurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EyeStoolBottomCurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*EyeVolume3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*GridPathShape](f, addIDs)
		stage.SerializeExcelizePointer2[*GrowthCurve2D](f, addIDs)
		stage.SerializeExcelizePointer2[*GrowthCurve2DRibbon](f, addIDs)
		stage.SerializeExcelizePointer2[*GrowthCurve2DRibbonEndShape](f, addIDs)
		stage.SerializeExcelizePointer2[*GrowthCurve2DRibbonStartShape](f, addIDs)
		stage.SerializeExcelizePointer2[*GrowthCurveRhombusGridShape](f, addIDs)
		stage.SerializeExcelizePointer2[*GrowthCurveRhombusShape](f, addIDs)
		stage.SerializeExcelizePointer2[*GrowthVectorShape](f, addIDs)
		stage.SerializeExcelizePointer2[*InitialRhombusGridShape](f, addIDs)
		stage.SerializeExcelizePointer2[*InitialRhombusShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Key3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*KeyHole3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*KeyHoleShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Leaves3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Library](f, addIDs)
		stage.SerializeExcelizePointer2[*MidArcVectorShape](f, addIDs)
		stage.SerializeExcelizePointer2[*MidArcVectorShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*MusicAbstract](f, addIDs)
		stage.SerializeExcelizePointer2[*OriginalPoints3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ParastichyMCurves3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ParastichyNCurves3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DRibbon](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DRibbonEndShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DRibbonStartShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectory](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectoryP1CurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectoryP1P2](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectoryP1P2PairLineShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectoryP1PointShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectoryP2CurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectoryP2PointShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyGrowthCurve2DTrajectoryShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyRotatedSeatBottomCurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyRotatedSeatTopCurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PartiallyRotatedTorusShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PerpendicularVector](f, addIDs)
		stage.SerializeExcelizePointer2[*PerpendicularVectorGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*PerpendicularVectorGridHalfway](f, addIDs)
		stage.SerializeExcelizePointer2[*PerpendicularVectorHalfway](f, addIDs)
		stage.SerializeExcelizePointer2[*Plant2DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*Plant3DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*PlantAbstract](f, addIDs)
		stage.SerializeExcelizePointer2[*PlantCircumferenceShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PointsAndLines3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*PxShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Rendered3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*RhombusShape](f, addIDs)
		stage.SerializeExcelizePointer2[*RhombusStuff](f, addIDs)
		stage.SerializeExcelizePointer2[*RotatedRhombusGridShape](f, addIDs)
		stage.SerializeExcelizePointer2[*RotatedRhombusShape](f, addIDs)
		stage.SerializeExcelizePointer2[*RotatedSampledPoints3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*RotatedSeatAndLegs3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*SampledPoints3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Seat3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*SeatAndLegs3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*SeatBottomCurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*SeatTopCurveShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedBottomTopStartArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedBottomTopStartArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftGrowthCurve2DRibbon](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftGrowthCurve2DRibbonEndShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftGrowthCurve2DRibbonStartShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftPartiallyGrowthCurve2DRibbon](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftPartiallyGrowthCurve2DRibbonEndShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftPartiallyGrowthCurve2DRibbonStartShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftStackGrowthCurveEndArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftStackGrowthCurveStartArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftStackNormalVector](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftStackOfGrowthCurve](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedLeftStackOfNormalVector](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedRightGrowthCurve2DRibbon](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedRightGrowthCurve2DRibbonEndShape](f, addIDs)
		stage.SerializeExcelizePointer2[*ShiftedRightGrowthCurve2DRibbonStartShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackGrowthCurve2DEndHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackGrowthCurve2DRibbonEndShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackGrowthCurve2DRibbonStartShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackGrowthCurve2DStartHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackOfGrowthCurve2D](f, addIDs)
		stage.SerializeExcelizePointer2[*StackOfGrowthCurve2DByGrowthVector](f, addIDs)
		stage.SerializeExcelizePointer2[*StackOfGrowthCurve2DRibbon](f, addIDs)
		stage.SerializeExcelizePointer2[*StackOfPartiallyRotatedTorusShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackOfRotatedGrowthCurve2D](f, addIDs)
		stage.SerializeExcelizePointer2[*StackOfRotatedGrowthCurve2DRibbon](f, addIDs)
		stage.SerializeExcelizePointer2[*StackRotatedGrowthCurve2DEndArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackRotatedGrowthCurve2DRibbonEndShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackRotatedGrowthCurve2DRibbonStartShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StackRotatedGrowthCurve2DStartArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StartArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StartArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*StartHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*StartHalfwayArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*StemCylinder3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Stool2DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*Stool3DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*StoolAbstract](f, addIDs)
		stage.SerializeExcelizePointer2[*TiledFloor3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopEndArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopEndArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*TopEndHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopEndHalfwayArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*TopGrowthCurve2D](f, addIDs)
		stage.SerializeExcelizePointer2[*TopMidArcVectorShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopMidArcVectorShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStackGrowthCurve2DEndHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStackGrowthCurve2DStartHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStackOfGrowthCurve2D](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStackOfRotatedGrowthCurve2D](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStackOfRotatedGrowthCurve2DEndArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStackOfRotatedGrowthCurve2DStartArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStartArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStartArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStartHalfwayArcShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TopStartHalfwayArcShapeGrid](f, addIDs)
		stage.SerializeExcelizePointer2[*Torus3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TorusEdge3DShape](f, addIDs)
		stage.SerializeExcelizePointer2[*TorusStackShape](f, addIDs)
		stage.SerializeExcelizePointer2[*Vase2DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*Vase3DDiagram](f, addIDs)
		stage.SerializeExcelizePointer2[*VaseAbstract](f, addIDs)
		stage.SerializeExcelizePointer2[*VerticalTorusStackShape](f, addIDs)
		stage.SerializeExcelizePointer2[*VolumeKey3DShape](f, addIDs)
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

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type PointerToGongstruct](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type PointerToGongstruct](f *excelize.File, addIDs bool) {
	sheetName := GetPointerToGongstructName[Type]()

	sheetName = shortenString(sheetName)

	// Create a new sheet.
	f.NewSheet(sheetName)

	set := *stage.GetInstancesSet[Type]()

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

// SerializeExcelizePointerToGongstruct is a backward-compatible forwarder.
func SerializeExcelizePointerToGongstruct[Type PointerToGongstruct](stage *Stage, f *excelize.File) {
	stage.SerializeExcelizePointer[Type](f)
}

// SerializeExcelizePointerToGongstruct2 is a backward-compatible forwarder.
func SerializeExcelizePointerToGongstruct2[Type PointerToGongstruct](stage *Stage, f *excelize.File, addIDs bool) {
	stage.SerializeExcelizePointer2[Type](f, addIDs)
}
