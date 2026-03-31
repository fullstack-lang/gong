// generated code - do not edit
package models

import (
	"cmp"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func SerializeStage(stage *Stage, filename string) {
	SerializeStage2(stage, filename, false)
}
func SerializeStage2(stage *Stage, filename string, addIDs bool) {

	f := excelize.NewFile()
	{
		// insertion point
		SerializeExcelizePointerToGongstruct2[*A_directive](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*A_measure](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*A_measure_1](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*A_part](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*A_part_1](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Accidental](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Accidental_mark](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Accidental_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Accord](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Accordion_registration](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Appearance](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Arpeggiate](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Arrow](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Articulations](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Assess](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Attributes](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Backup](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Bar_style_color](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Barline](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Barre](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Bass](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Bass_step](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Beam](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Beat_repeat](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Beat_unit_tied](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Beater](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Bend](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Bookmark](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Bracket](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Breath_mark](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Caesura](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Cancel](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Clef](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Coda](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Credit](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Dashes](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Defaults](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Degree](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Degree_alter](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Degree_type](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Degree_value](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Direction](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Direction_type](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Distance](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Double](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Dynamics](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Effect](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Elision](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_font](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_line](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_placement](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_placement_smufl](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_print_object_style_align](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_print_style](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_print_style_align](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_print_style_align_id](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Empty_trill_sound](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Encoding](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Ending](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Extend](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Feature](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Fermata](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Figure](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Figured_bass](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Fingering](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*First_fret](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*For_part](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Formatted_symbol](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Formatted_symbol_id](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Formatted_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Formatted_text_id](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Forward](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Frame](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Frame_note](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Fret](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Glass](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Glissando](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Glyph](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Grace](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Group_barline](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Group_name](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Group_symbol](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Grouping](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Hammer_on_pull_off](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Handbell](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Harmon_closed](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Harmon_mute](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Harmonic](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Harmony](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Harmony_alter](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Harp_pedals](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Heel_toe](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Hole](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Hole_closed](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Horizontal_turn](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Identification](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Image](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Instrument](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Instrument_change](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Instrument_link](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Interchangeable](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Inversion](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Key](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Key_accidental](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Key_octave](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Kind](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Level](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Line_detail](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Line_width](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Link](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Listen](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Listening](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Lyric](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Lyric_font](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Lyric_language](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Measure_layout](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Measure_numbering](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Measure_repeat](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Measure_style](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Membrane](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Metal](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Metronome](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Metronome_beam](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Metronome_note](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Metronome_tied](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Metronome_tuplet](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Midi_device](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Midi_instrument](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Miscellaneous](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Miscellaneous_field](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Mordent](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Multiple_rest](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Name_display](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Non_arpeggiate](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Notations](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Note](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Note_size](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Note_type](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Notehead](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Notehead_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Numeral](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Numeral_key](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Numeral_root](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Octave_shift](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Offset](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Opus](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Ornaments](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Other_appearance](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Other_direction](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Other_listening](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Other_notation](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Other_placement_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Other_play](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Other_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Page_layout](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Page_margins](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Part_clef](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Part_group](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Part_link](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Part_list](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Part_name](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Part_symbol](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Part_transpose](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Pedal](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Pedal_tuning](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Per_minute](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Percussion](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Pitch](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Pitched](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Placement_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Play](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Player](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Principal_voice](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Print](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Release](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Repeat](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Rest](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Root](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Root_step](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Scaling](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Scordatura](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Score_instrument](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Score_part](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Score_partwise](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Score_timewise](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Segno](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Slash](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Slide](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Slur](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Sound](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Staff_details](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Staff_divide](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Staff_layout](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Staff_size](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Staff_tuning](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Stem](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Stick](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*String_mute](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*String_type](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Strong_accent](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Style_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Supports](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Swing](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Sync](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*System_dividers](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*System_layout](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*System_margins](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tap](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Technical](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Text_element_data](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tie](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tied](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Time](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Time_modification](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Timpani](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Transpose](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tremolo](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tuplet](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tuplet_dot](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tuplet_number](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tuplet_portion](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Tuplet_type](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Typed_text](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Unpitched](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Virtual_instrument](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Wait](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Wavy_line](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Wedge](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Wood](stage, f, addIDs)
		SerializeExcelizePointerToGongstruct2[*Work](stage, f, addIDs)
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
		return
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
		return
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
				return
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return
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
		if err := f.SaveAs(filename); err != nil {
			fmt.Println("cannot write xl file : ", err)
		}
	}
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

		// 3. Add the ID value in column A
		// We use type assertion to check if the instance implements GetID()
		id := GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(instance), uint64(GetOrderPointerGongstruct(stage, instance)))

		for index, fieldName := range GetFieldsFromPointer[Type]() {
			fieldStringValue := GetFieldStringValueFromPointer(instance, fieldName.Name, stage)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", IntToLetters(int32(2*index+2)), line), id)
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

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}
