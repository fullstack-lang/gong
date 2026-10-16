// generated code - do not edit
package direction

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
			for instance := range stage.Accords {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Accord", instances, (*Accord)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AccordionRegistrations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AccordionRegistration", instances, (*AccordionRegistration)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Barres {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Barre", instances, (*Barre)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Basss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bass", instances, (*Bass)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BassSteps {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BassStep", instances, (*BassStep)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.BeatUnitTieds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "BeatUnitTied", instances, (*BeatUnitTied)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Beaters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Beater", instances, (*Beater)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Brackets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bracket", instances, (*Bracket)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Dashess {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Dashes", instances, (*Dashes)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Degrees {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Degree", instances, (*Degree)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DegreeAlters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DegreeAlter", instances, (*DegreeAlter)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DegreeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DegreeType", instances, (*DegreeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DegreeValues {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DegreeValue", instances, (*DegreeValue)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Directions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Direction", instances, (*Direction)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DirectionTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DirectionType", instances, (*DirectionType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Effects {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Effect", instances, (*Effect)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Features {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Feature", instances, (*Feature)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FirstFrets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FirstFret", instances, (*FirstFret)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Frames {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Frame", instances, (*Frame)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FrameNotes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FrameNote", instances, (*FrameNote)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Glasss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Glass", instances, (*Glass)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Groupings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Grouping", instances, (*Grouping)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harmonys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harmony", instances, (*Harmony)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HarmonyAlters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HarmonyAlter", instances, (*HarmonyAlter)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.HarpPedalss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "HarpPedals", instances, (*HarpPedals)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Images {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Image", instances, (*Image)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.InstrumentChanges {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "InstrumentChange", instances, (*InstrumentChange)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Inversions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Inversion", instances, (*Inversion)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Kinds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Kind", instances, (*Kind)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Listenings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Listening", instances, (*Listening)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MeasureNumberings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MeasureNumbering", instances, (*MeasureNumbering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Membranes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Membrane", instances, (*Membrane)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metal", instances, (*Metal)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metronomes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metronome", instances, (*Metronome)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MetronomeBeams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MetronomeBeam", instances, (*MetronomeBeam)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MetronomeNotes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MetronomeNote", instances, (*MetronomeNote)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MetronomeTieds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MetronomeTied", instances, (*MetronomeTied)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MetronomeTuplets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MetronomeTuplet", instances, (*MetronomeTuplet)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Numerals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Numeral", instances, (*Numeral)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NumeralKeys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NumeralKey", instances, (*NumeralKey)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NumeralRoots {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NumeralRoot", instances, (*NumeralRoot)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OctaveShifts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OctaveShift", instances, (*OctaveShift)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Offsets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Offset", instances, (*Offset)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OtherDirections {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OtherDirection", instances, (*OtherDirection)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.OtherListenings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OtherListening", instances, (*OtherListening)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Pedals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pedal", instances, (*Pedal)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PedalTunings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PedalTuning", instances, (*PedalTuning)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PerMinutes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PerMinute", instances, (*PerMinute)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Percussions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Percussion", instances, (*Percussion)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Pitcheds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pitched", instances, (*Pitched)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.PrincipalVoices {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "PrincipalVoice", instances, (*PrincipalVoice)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Prints {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Print", instances, (*Print)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Roots {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Root", instances, (*Root)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RootSteps {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RootStep", instances, (*RootStep)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Scordaturas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Scordatura", instances, (*Scordatura)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Sounds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Sound", instances, (*Sound)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StaffDivides {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StaffDivide", instances, (*StaffDivide)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Sticks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Stick", instances, (*Stick)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StringMutes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StringMute", instances, (*StringMute)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StyleTexts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StyleText", instances, (*StyleText)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Swings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Swing", instances, (*Swing)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Syncs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Sync", instances, (*Sync)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimeModifications {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimeModification", instances, (*TimeModification)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Timpanis {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Timpani", instances, (*Timpani)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.VirtualInstruments {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "VirtualInstrument", instances, (*VirtualInstrument)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Wedges {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Wedge", instances, (*Wedge)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Woods {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Wood", instances, (*Wood)(nil).GongGetFieldHeaders(), addIDs)
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
