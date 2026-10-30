// generated code - do not edit
package aixmfeatures

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
			for instance := range stage.A_extensions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension", instances, (*A_extension)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_1", instances, (*A_extension_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_10s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_10", instances, (*A_extension_10)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_11s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_11", instances, (*A_extension_11)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_12s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_12", instances, (*A_extension_12)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_13s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_13", instances, (*A_extension_13)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_14s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_14", instances, (*A_extension_14)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_2s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_2", instances, (*A_extension_2)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_3s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_3", instances, (*A_extension_3)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_4s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_4", instances, (*A_extension_4)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_5s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_5", instances, (*A_extension_5)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_6s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_6", instances, (*A_extension_6)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_7s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_7", instances, (*A_extension_7)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_8s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_8", instances, (*A_extension_8)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_extension_9s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_extension_9", instances, (*A_extension_9)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AbstractPropertiesWithScheduleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AbstractPropertiesWithScheduleType", instances, (*AbstractPropertiesWithScheduleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftCharacteristicPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftCharacteristicPropertyType", instances, (*AircraftCharacteristicPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AircraftCharacteristicTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AircraftCharacteristicType", instances, (*AircraftCharacteristicType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceActivationPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceActivationPropertyType", instances, (*AirspaceActivationPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceActivationTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceActivationType", instances, (*AirspaceActivationType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceGeometryComponentPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceGeometryComponentPropertyType", instances, (*AirspaceGeometryComponentPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceGeometryComponentTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceGeometryComponentType", instances, (*AirspaceGeometryComponentType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerClassPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerClassPropertyType", instances, (*AirspaceLayerClassPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerClassTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerClassType", instances, (*AirspaceLayerClassType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerPropertyType", instances, (*AirspaceLayerPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceLayerTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceLayerType", instances, (*AirspaceLayerType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspacePropertyType", instances, (*AirspacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceTimeSlicePropertyType", instances, (*AirspaceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceTimeSliceType", instances, (*AirspaceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceType", instances, (*AirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumeDependencyPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumeDependencyPropertyType", instances, (*AirspaceVolumeDependencyPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumeDependencyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumeDependencyType", instances, (*AirspaceVolumeDependencyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumePropertyType", instances, (*AirspaceVolumePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AirspaceVolumeTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AirspaceVolumeType", instances, (*AirspaceVolumeType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAirspaceTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAirspaceTimeSlicePropertyType", instances, (*AuthorityForAirspaceTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAirspaceTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAirspaceTimeSliceType", instances, (*AuthorityForAirspaceTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AuthorityForAirspaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AuthorityForAirspaceType", instances, (*AuthorityForAirspaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurvePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurvePropertyType", instances, (*CurvePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.CurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "CurveType", instances, (*CurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedCurveTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedCurveType", instances, (*ElevatedCurveType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ElevatedSurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ElevatedSurfaceType", instances, (*ElevatedSurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeoBorderTimeSlicePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeoBorderTimeSlicePropertyType", instances, (*GeoBorderTimeSlicePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeoBorderTimeSliceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeoBorderTimeSliceType", instances, (*GeoBorderTimeSliceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.GeoBorderTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "GeoBorderType", instances, (*GeoBorderType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinguisticNotePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinguisticNotePropertyType", instances, (*LinguisticNotePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.LinguisticNoteTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "LinguisticNoteType", instances, (*LinguisticNoteType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.NotePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "NotePropertyType", instances, (*NotePropertyType)(nil).GongGetFieldHeaders(), addIDs)
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
			for instance := range stage.OrganisationAuthorityPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "OrganisationAuthorityPropertyType", instances, (*OrganisationAuthorityPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RoutePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RoutePropertyType", instances, (*RoutePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StandardLevelColumnPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StandardLevelColumnPropertyType", instances, (*StandardLevelColumnPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfacePropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfacePropertyType", instances, (*SurfacePropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SurfaceTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SurfaceType", instances, (*SurfaceType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimesheetPropertyTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimesheetPropertyType", instances, (*TimesheetPropertyType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TimesheetTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TimesheetType", instances, (*TimesheetType)(nil).GongGetFieldHeaders(), addIDs)
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
