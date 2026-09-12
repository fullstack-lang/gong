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
		stage.SerializeExcelizePointer2[*A_directive](f, addIDs)
		stage.SerializeExcelizePointer2[*A_measure](f, addIDs)
		stage.SerializeExcelizePointer2[*A_measure_1](f, addIDs)
		stage.SerializeExcelizePointer2[*A_part](f, addIDs)
		stage.SerializeExcelizePointer2[*A_part_1](f, addIDs)
		stage.SerializeExcelizePointer2[*Accidental](f, addIDs)
		stage.SerializeExcelizePointer2[*Accidental_mark](f, addIDs)
		stage.SerializeExcelizePointer2[*Accidental_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Accord](f, addIDs)
		stage.SerializeExcelizePointer2[*Accordion_registration](f, addIDs)
		stage.SerializeExcelizePointer2[*Appearance](f, addIDs)
		stage.SerializeExcelizePointer2[*Arpeggiate](f, addIDs)
		stage.SerializeExcelizePointer2[*Arrow](f, addIDs)
		stage.SerializeExcelizePointer2[*Articulations](f, addIDs)
		stage.SerializeExcelizePointer2[*Assess](f, addIDs)
		stage.SerializeExcelizePointer2[*Attributes](f, addIDs)
		stage.SerializeExcelizePointer2[*Backup](f, addIDs)
		stage.SerializeExcelizePointer2[*Bar_style_color](f, addIDs)
		stage.SerializeExcelizePointer2[*Barline](f, addIDs)
		stage.SerializeExcelizePointer2[*Barre](f, addIDs)
		stage.SerializeExcelizePointer2[*Bass](f, addIDs)
		stage.SerializeExcelizePointer2[*Bass_step](f, addIDs)
		stage.SerializeExcelizePointer2[*Beam](f, addIDs)
		stage.SerializeExcelizePointer2[*Beat_repeat](f, addIDs)
		stage.SerializeExcelizePointer2[*Beat_unit_tied](f, addIDs)
		stage.SerializeExcelizePointer2[*Beater](f, addIDs)
		stage.SerializeExcelizePointer2[*Bend](f, addIDs)
		stage.SerializeExcelizePointer2[*Bookmark](f, addIDs)
		stage.SerializeExcelizePointer2[*Bracket](f, addIDs)
		stage.SerializeExcelizePointer2[*Breath_mark](f, addIDs)
		stage.SerializeExcelizePointer2[*Caesura](f, addIDs)
		stage.SerializeExcelizePointer2[*Cancel](f, addIDs)
		stage.SerializeExcelizePointer2[*Clef](f, addIDs)
		stage.SerializeExcelizePointer2[*Coda](f, addIDs)
		stage.SerializeExcelizePointer2[*Credit](f, addIDs)
		stage.SerializeExcelizePointer2[*Dashes](f, addIDs)
		stage.SerializeExcelizePointer2[*Defaults](f, addIDs)
		stage.SerializeExcelizePointer2[*Degree](f, addIDs)
		stage.SerializeExcelizePointer2[*Degree_alter](f, addIDs)
		stage.SerializeExcelizePointer2[*Degree_type](f, addIDs)
		stage.SerializeExcelizePointer2[*Degree_value](f, addIDs)
		stage.SerializeExcelizePointer2[*Direction](f, addIDs)
		stage.SerializeExcelizePointer2[*Direction_type](f, addIDs)
		stage.SerializeExcelizePointer2[*Distance](f, addIDs)
		stage.SerializeExcelizePointer2[*Double](f, addIDs)
		stage.SerializeExcelizePointer2[*Dynamics](f, addIDs)
		stage.SerializeExcelizePointer2[*Effect](f, addIDs)
		stage.SerializeExcelizePointer2[*Elision](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_font](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_line](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_placement](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_placement_smufl](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_print_object_style_align](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_print_style](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_print_style_align](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_print_style_align_id](f, addIDs)
		stage.SerializeExcelizePointer2[*Empty_trill_sound](f, addIDs)
		stage.SerializeExcelizePointer2[*Encoding](f, addIDs)
		stage.SerializeExcelizePointer2[*Ending](f, addIDs)
		stage.SerializeExcelizePointer2[*Extend](f, addIDs)
		stage.SerializeExcelizePointer2[*Feature](f, addIDs)
		stage.SerializeExcelizePointer2[*Fermata](f, addIDs)
		stage.SerializeExcelizePointer2[*Figure](f, addIDs)
		stage.SerializeExcelizePointer2[*Figured_bass](f, addIDs)
		stage.SerializeExcelizePointer2[*Fingering](f, addIDs)
		stage.SerializeExcelizePointer2[*First_fret](f, addIDs)
		stage.SerializeExcelizePointer2[*For_part](f, addIDs)
		stage.SerializeExcelizePointer2[*Formatted_symbol](f, addIDs)
		stage.SerializeExcelizePointer2[*Formatted_symbol_id](f, addIDs)
		stage.SerializeExcelizePointer2[*Formatted_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Formatted_text_id](f, addIDs)
		stage.SerializeExcelizePointer2[*Forward](f, addIDs)
		stage.SerializeExcelizePointer2[*Frame](f, addIDs)
		stage.SerializeExcelizePointer2[*Frame_note](f, addIDs)
		stage.SerializeExcelizePointer2[*Fret](f, addIDs)
		stage.SerializeExcelizePointer2[*Glass](f, addIDs)
		stage.SerializeExcelizePointer2[*Glissando](f, addIDs)
		stage.SerializeExcelizePointer2[*Glyph](f, addIDs)
		stage.SerializeExcelizePointer2[*Grace](f, addIDs)
		stage.SerializeExcelizePointer2[*Group_barline](f, addIDs)
		stage.SerializeExcelizePointer2[*Group_name](f, addIDs)
		stage.SerializeExcelizePointer2[*Group_symbol](f, addIDs)
		stage.SerializeExcelizePointer2[*Grouping](f, addIDs)
		stage.SerializeExcelizePointer2[*Hammer_on_pull_off](f, addIDs)
		stage.SerializeExcelizePointer2[*Handbell](f, addIDs)
		stage.SerializeExcelizePointer2[*Harmon_closed](f, addIDs)
		stage.SerializeExcelizePointer2[*Harmon_mute](f, addIDs)
		stage.SerializeExcelizePointer2[*Harmonic](f, addIDs)
		stage.SerializeExcelizePointer2[*Harmony](f, addIDs)
		stage.SerializeExcelizePointer2[*Harmony_alter](f, addIDs)
		stage.SerializeExcelizePointer2[*Harp_pedals](f, addIDs)
		stage.SerializeExcelizePointer2[*Heel_toe](f, addIDs)
		stage.SerializeExcelizePointer2[*Hole](f, addIDs)
		stage.SerializeExcelizePointer2[*Hole_closed](f, addIDs)
		stage.SerializeExcelizePointer2[*Horizontal_turn](f, addIDs)
		stage.SerializeExcelizePointer2[*Identification](f, addIDs)
		stage.SerializeExcelizePointer2[*Image](f, addIDs)
		stage.SerializeExcelizePointer2[*Instrument](f, addIDs)
		stage.SerializeExcelizePointer2[*Instrument_change](f, addIDs)
		stage.SerializeExcelizePointer2[*Instrument_link](f, addIDs)
		stage.SerializeExcelizePointer2[*Interchangeable](f, addIDs)
		stage.SerializeExcelizePointer2[*Inversion](f, addIDs)
		stage.SerializeExcelizePointer2[*Key](f, addIDs)
		stage.SerializeExcelizePointer2[*Key_accidental](f, addIDs)
		stage.SerializeExcelizePointer2[*Key_octave](f, addIDs)
		stage.SerializeExcelizePointer2[*Kind](f, addIDs)
		stage.SerializeExcelizePointer2[*Level](f, addIDs)
		stage.SerializeExcelizePointer2[*Line_detail](f, addIDs)
		stage.SerializeExcelizePointer2[*Line_width](f, addIDs)
		stage.SerializeExcelizePointer2[*Link](f, addIDs)
		stage.SerializeExcelizePointer2[*Listen](f, addIDs)
		stage.SerializeExcelizePointer2[*Listening](f, addIDs)
		stage.SerializeExcelizePointer2[*Lyric](f, addIDs)
		stage.SerializeExcelizePointer2[*Lyric_font](f, addIDs)
		stage.SerializeExcelizePointer2[*Lyric_language](f, addIDs)
		stage.SerializeExcelizePointer2[*Measure_layout](f, addIDs)
		stage.SerializeExcelizePointer2[*Measure_numbering](f, addIDs)
		stage.SerializeExcelizePointer2[*Measure_repeat](f, addIDs)
		stage.SerializeExcelizePointer2[*Measure_style](f, addIDs)
		stage.SerializeExcelizePointer2[*Membrane](f, addIDs)
		stage.SerializeExcelizePointer2[*Metal](f, addIDs)
		stage.SerializeExcelizePointer2[*Metronome](f, addIDs)
		stage.SerializeExcelizePointer2[*Metronome_beam](f, addIDs)
		stage.SerializeExcelizePointer2[*Metronome_note](f, addIDs)
		stage.SerializeExcelizePointer2[*Metronome_tied](f, addIDs)
		stage.SerializeExcelizePointer2[*Metronome_tuplet](f, addIDs)
		stage.SerializeExcelizePointer2[*Midi_device](f, addIDs)
		stage.SerializeExcelizePointer2[*Midi_instrument](f, addIDs)
		stage.SerializeExcelizePointer2[*Miscellaneous](f, addIDs)
		stage.SerializeExcelizePointer2[*Miscellaneous_field](f, addIDs)
		stage.SerializeExcelizePointer2[*Mordent](f, addIDs)
		stage.SerializeExcelizePointer2[*Multiple_rest](f, addIDs)
		stage.SerializeExcelizePointer2[*Name_display](f, addIDs)
		stage.SerializeExcelizePointer2[*Non_arpeggiate](f, addIDs)
		stage.SerializeExcelizePointer2[*Notations](f, addIDs)
		stage.SerializeExcelizePointer2[*Note](f, addIDs)
		stage.SerializeExcelizePointer2[*Note_size](f, addIDs)
		stage.SerializeExcelizePointer2[*Note_type](f, addIDs)
		stage.SerializeExcelizePointer2[*Notehead](f, addIDs)
		stage.SerializeExcelizePointer2[*Notehead_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Numeral](f, addIDs)
		stage.SerializeExcelizePointer2[*Numeral_key](f, addIDs)
		stage.SerializeExcelizePointer2[*Numeral_root](f, addIDs)
		stage.SerializeExcelizePointer2[*Octave_shift](f, addIDs)
		stage.SerializeExcelizePointer2[*Offset](f, addIDs)
		stage.SerializeExcelizePointer2[*Opus](f, addIDs)
		stage.SerializeExcelizePointer2[*Ornaments](f, addIDs)
		stage.SerializeExcelizePointer2[*Other_appearance](f, addIDs)
		stage.SerializeExcelizePointer2[*Other_direction](f, addIDs)
		stage.SerializeExcelizePointer2[*Other_listening](f, addIDs)
		stage.SerializeExcelizePointer2[*Other_notation](f, addIDs)
		stage.SerializeExcelizePointer2[*Other_placement_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Other_play](f, addIDs)
		stage.SerializeExcelizePointer2[*Other_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Page_layout](f, addIDs)
		stage.SerializeExcelizePointer2[*Page_margins](f, addIDs)
		stage.SerializeExcelizePointer2[*Part_clef](f, addIDs)
		stage.SerializeExcelizePointer2[*Part_group](f, addIDs)
		stage.SerializeExcelizePointer2[*Part_link](f, addIDs)
		stage.SerializeExcelizePointer2[*Part_list](f, addIDs)
		stage.SerializeExcelizePointer2[*Part_name](f, addIDs)
		stage.SerializeExcelizePointer2[*Part_symbol](f, addIDs)
		stage.SerializeExcelizePointer2[*Part_transpose](f, addIDs)
		stage.SerializeExcelizePointer2[*Pedal](f, addIDs)
		stage.SerializeExcelizePointer2[*Pedal_tuning](f, addIDs)
		stage.SerializeExcelizePointer2[*Per_minute](f, addIDs)
		stage.SerializeExcelizePointer2[*Percussion](f, addIDs)
		stage.SerializeExcelizePointer2[*Pitch](f, addIDs)
		stage.SerializeExcelizePointer2[*Pitched](f, addIDs)
		stage.SerializeExcelizePointer2[*Placement_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Play](f, addIDs)
		stage.SerializeExcelizePointer2[*Player](f, addIDs)
		stage.SerializeExcelizePointer2[*Principal_voice](f, addIDs)
		stage.SerializeExcelizePointer2[*Print](f, addIDs)
		stage.SerializeExcelizePointer2[*Release](f, addIDs)
		stage.SerializeExcelizePointer2[*Repeat](f, addIDs)
		stage.SerializeExcelizePointer2[*Rest](f, addIDs)
		stage.SerializeExcelizePointer2[*Root](f, addIDs)
		stage.SerializeExcelizePointer2[*Root_step](f, addIDs)
		stage.SerializeExcelizePointer2[*Scaling](f, addIDs)
		stage.SerializeExcelizePointer2[*Scordatura](f, addIDs)
		stage.SerializeExcelizePointer2[*Score_instrument](f, addIDs)
		stage.SerializeExcelizePointer2[*Score_part](f, addIDs)
		stage.SerializeExcelizePointer2[*Score_partwise](f, addIDs)
		stage.SerializeExcelizePointer2[*Score_timewise](f, addIDs)
		stage.SerializeExcelizePointer2[*Segno](f, addIDs)
		stage.SerializeExcelizePointer2[*Slash](f, addIDs)
		stage.SerializeExcelizePointer2[*Slide](f, addIDs)
		stage.SerializeExcelizePointer2[*Slur](f, addIDs)
		stage.SerializeExcelizePointer2[*Sound](f, addIDs)
		stage.SerializeExcelizePointer2[*Staff_details](f, addIDs)
		stage.SerializeExcelizePointer2[*Staff_divide](f, addIDs)
		stage.SerializeExcelizePointer2[*Staff_layout](f, addIDs)
		stage.SerializeExcelizePointer2[*Staff_size](f, addIDs)
		stage.SerializeExcelizePointer2[*Staff_tuning](f, addIDs)
		stage.SerializeExcelizePointer2[*Stem](f, addIDs)
		stage.SerializeExcelizePointer2[*Stick](f, addIDs)
		stage.SerializeExcelizePointer2[*String_mute](f, addIDs)
		stage.SerializeExcelizePointer2[*String_type](f, addIDs)
		stage.SerializeExcelizePointer2[*Strong_accent](f, addIDs)
		stage.SerializeExcelizePointer2[*Style_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Supports](f, addIDs)
		stage.SerializeExcelizePointer2[*Swing](f, addIDs)
		stage.SerializeExcelizePointer2[*Sync](f, addIDs)
		stage.SerializeExcelizePointer2[*System_dividers](f, addIDs)
		stage.SerializeExcelizePointer2[*System_layout](f, addIDs)
		stage.SerializeExcelizePointer2[*System_margins](f, addIDs)
		stage.SerializeExcelizePointer2[*Tap](f, addIDs)
		stage.SerializeExcelizePointer2[*Technical](f, addIDs)
		stage.SerializeExcelizePointer2[*Text_element_data](f, addIDs)
		stage.SerializeExcelizePointer2[*Tie](f, addIDs)
		stage.SerializeExcelizePointer2[*Tied](f, addIDs)
		stage.SerializeExcelizePointer2[*Time](f, addIDs)
		stage.SerializeExcelizePointer2[*Time_modification](f, addIDs)
		stage.SerializeExcelizePointer2[*Timpani](f, addIDs)
		stage.SerializeExcelizePointer2[*Transpose](f, addIDs)
		stage.SerializeExcelizePointer2[*Tremolo](f, addIDs)
		stage.SerializeExcelizePointer2[*Tuplet](f, addIDs)
		stage.SerializeExcelizePointer2[*Tuplet_dot](f, addIDs)
		stage.SerializeExcelizePointer2[*Tuplet_number](f, addIDs)
		stage.SerializeExcelizePointer2[*Tuplet_portion](f, addIDs)
		stage.SerializeExcelizePointer2[*Tuplet_type](f, addIDs)
		stage.SerializeExcelizePointer2[*Typed_text](f, addIDs)
		stage.SerializeExcelizePointer2[*Unpitched](f, addIDs)
		stage.SerializeExcelizePointer2[*Virtual_instrument](f, addIDs)
		stage.SerializeExcelizePointer2[*Wait](f, addIDs)
		stage.SerializeExcelizePointer2[*Wavy_line](f, addIDs)
		stage.SerializeExcelizePointer2[*Wedge](f, addIDs)
		stage.SerializeExcelizePointer2[*Wood](f, addIDs)
		stage.SerializeExcelizePointer2[*Work](f, addIDs)
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
