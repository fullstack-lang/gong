// generated code - do not edit
package common

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.AccidentalTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AccidentalText", instances, (*AccidentalText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Appearances {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Appearance", instances, (*Appearance)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Bookmarks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bookmark", instances, (*Bookmark)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Codas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Coda", instances, (*Coda)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Distances {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Distance", instances, (*Distance)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Dynamicss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Dynamics", instances, (*Dynamics)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Emptys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty", instances, (*Empty)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyPlacements {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyPlacement", instances, (*EmptyPlacement)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyPlacementSmufls {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyPlacementSmufl", instances, (*EmptyPlacementSmufl)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyPrintObjectStyleAligns {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyPrintObjectStyleAlign", instances, (*EmptyPrintObjectStyleAlign)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyPrintStyles {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyPrintStyle", instances, (*EmptyPrintStyle)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyPrintStyleAligns {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyPrintStyleAlign", instances, (*EmptyPrintStyleAlign)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyPrintStyleAlignIds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyPrintStyleAlignId", instances, (*EmptyPrintStyleAlignId)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyTrillSounds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyTrillSound", instances, (*EmptyTrillSound)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Encodings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Encoding", instances, (*Encoding)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Fermatas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Fermata", instances, (*Fermata)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Fingerings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Fingering", instances, (*Fingering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FormattedSymbols {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FormattedSymbol", instances, (*FormattedSymbol)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FormattedSymbolIds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FormattedSymbolId", instances, (*FormattedSymbolId)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FormattedTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FormattedText", instances, (*FormattedText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FormattedTextIds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FormattedTextId", instances, (*FormattedTextId)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Frets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Fret", instances, (*Fret)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Glyphs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Glyph", instances, (*Glyph)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HorizontalTurns {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HorizontalTurn", instances, (*HorizontalTurn)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Identifications {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Identification", instances, (*Identification)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Levels {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Level", instances, (*Level)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LineWidths {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LineWidth", instances, (*LineWidth)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Links {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Link", instances, (*Link)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MeasureLayouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MeasureLayout", instances, (*MeasureLayout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MidiDevices {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MidiDevice", instances, (*MidiDevice)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MidiInstruments {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MidiInstrument", instances, (*MidiInstrument)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Miscellaneouss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Miscellaneous", instances, (*Miscellaneous)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MiscellaneousFields {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MiscellaneousField", instances, (*MiscellaneousField)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NameDisplays {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NameDisplay", instances, (*NameDisplay)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoteSizes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoteSize", instances, (*NoteSize)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OtherAppearances {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OtherAppearance", instances, (*OtherAppearance)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OtherPlays {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OtherPlay", instances, (*OtherPlay)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OtherTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OtherText", instances, (*OtherText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PageLayouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PageLayout", instances, (*PageLayout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PageMarginss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PageMargins", instances, (*PageMargins)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Plays {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Play", instances, (*Play)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Scalings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Scaling", instances, (*Scaling)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Segnos {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Segno", instances, (*Segno)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StaffLayouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StaffLayout", instances, (*StaffLayout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StringTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StringType", instances, (*StringType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Supportss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Supports", instances, (*Supports)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SystemDividerss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SystemDividers", instances, (*SystemDividers)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SystemLayouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SystemLayout", instances, (*SystemLayout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SystemMarginss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SystemMargins", instances, (*SystemMargins)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TypedTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TypedText", instances, (*TypedText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WavyLines {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WavyLine", instances, (*WavyLine)(nil).GongGetFieldHeaders(), addIDs)
		}
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

	var tab GongExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// GongTabulator is an interface for writing to a table strings
type GongTabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type Tabulator = GongTabulator

type GongExcelizeTabulator struct {
	f *excelize.File
}

type ExcelizeTabulator = GongExcelizeTabulator

func (tab *GongExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *GongExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *GongExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *GongExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
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
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type GongstructPtr](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type GongstructPtr](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
