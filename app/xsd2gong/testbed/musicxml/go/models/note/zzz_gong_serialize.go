// generated code - do not edit
package note

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
			for instance := range stage.Accidentals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Accidental", instances, (*Accidental)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AccidentalMarks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AccidentalMark", instances, (*AccidentalMark)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Arpeggiates {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Arpeggiate", instances, (*Arpeggiate)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Arrows {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Arrow", instances, (*Arrow)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Articulationss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Articulations", instances, (*Articulations)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Assesss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Assess", instances, (*Assess)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Backups {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Backup", instances, (*Backup)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Beams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Beam", instances, (*Beam)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Bends {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bend", instances, (*Bend)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BreathMarks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BreathMark", instances, (*BreathMark)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Caesuras {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Caesura", instances, (*Caesura)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Elisions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Elision", instances, (*Elision)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmptyLines {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmptyLine", instances, (*EmptyLine)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Extends {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Extend", instances, (*Extend)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Figures {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Figure", instances, (*Figure)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FiguredBasss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FiguredBass", instances, (*FiguredBass)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Forwards {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Forward", instances, (*Forward)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Glissandos {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Glissando", instances, (*Glissando)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Graces {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Grace", instances, (*Grace)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HammerOnPullOffs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HammerOnPullOff", instances, (*HammerOnPullOff)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Handbells {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Handbell", instances, (*Handbell)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HarmonCloseds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HarmonClosed", instances, (*HarmonClosed)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HarmonMutes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HarmonMute", instances, (*HarmonMute)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harmonics {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harmonic", instances, (*Harmonic)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HeelToes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HeelToe", instances, (*HeelToe)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Holes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Hole", instances, (*Hole)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HoleCloseds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HoleClosed", instances, (*HoleClosed)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Instruments {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Instrument", instances, (*Instrument)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Listens {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Listen", instances, (*Listen)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lyrics {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Lyric", instances, (*Lyric)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Mordents {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Mordent", instances, (*Mordent)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NonArpeggiates {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NonArpeggiate", instances, (*NonArpeggiate)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Notationss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Notations", instances, (*Notations)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Notes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Note", instances, (*Note)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoteType", instances, (*NoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Noteheads {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Notehead", instances, (*Notehead)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NoteheadTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NoteheadText", instances, (*NoteheadText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Ornamentss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Ornaments", instances, (*Ornaments)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OtherNotations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OtherNotation", instances, (*OtherNotation)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OtherPlacementTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OtherPlacementText", instances, (*OtherPlacementText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Pitchs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pitch", instances, (*Pitch)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PlacementTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PlacementText", instances, (*PlacementText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Releases {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Release", instances, (*Release)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Rests {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Rest", instances, (*Rest)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Slides {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Slide", instances, (*Slide)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Slurs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Slur", instances, (*Slur)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Stems {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Stem", instances, (*Stem)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StrongAccents {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StrongAccent", instances, (*StrongAccent)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Taps {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tap", instances, (*Tap)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Technicals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Technical", instances, (*Technical)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TextElementDatas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TextElementData", instances, (*TextElementData)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Ties {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tie", instances, (*Tie)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tieds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tied", instances, (*Tied)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tremolos {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tremolo", instances, (*Tremolo)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tuplets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tuplet", instances, (*Tuplet)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TupletDots {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TupletDot", instances, (*TupletDot)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TupletNumbers {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TupletNumber", instances, (*TupletNumber)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TupletPortions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TupletPortion", instances, (*TupletPortion)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TupletTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TupletType", instances, (*TupletType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Unpitcheds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Unpitched", instances, (*Unpitched)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Waits {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Wait", instances, (*Wait)(nil).GongGetFieldHeaders(), addIDs)
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
