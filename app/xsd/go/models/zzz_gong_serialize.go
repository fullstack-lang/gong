// generated code - do not edit
package models

import (
	"archive/zip"
	"bytes"
	"cmp"
	"encoding/csv"
	"fmt"
	"io"
	"os"
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

func (stage *Stage) SerializeStageCSV(filename string) {
	stage.SerializeStageCSV2(filename, false)
}

func (stage *Stage) SerializeStageCSV2(filename string, addIDs bool) {
	csvZipBytes, err := stage.SerializeStageCSVAsBytes(addIDs)
	if err != nil {
		fmt.Println("cannot build csv zip : ", err)
		return
	}
	if err := os.WriteFile(filename, csvZipBytes, 0644); err != nil {
		fmt.Println("cannot write csv zip file : ", err)
	}
}

// SerializeStageCSVAsBytes serializes the stage to an in-memory ZIP archive of CSV files and returns the bytes.
func (stage *Stage) SerializeStageCSVAsBytes(addIDs bool) ([]byte, error) {
	var buf bytes.Buffer
	zw := zip.NewWriter(&buf)
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.Alls {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "All", instances, (*All)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Annotations {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Annotation", instances, (*Annotation)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Attributes {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Attribute", instances, (*Attribute)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.AttributeGroups {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "AttributeGroup", instances, (*AttributeGroup)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Choices {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Choice", instances, (*Choice)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.ComplexContents {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "ComplexContent", instances, (*ComplexContent)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.ComplexTypes {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "ComplexType", instances, (*ComplexType)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Documentations {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Documentation", instances, (*Documentation)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Elements {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Element", instances, (*Element)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Enumerations {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Enumeration", instances, (*Enumeration)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Extensions {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Extension", instances, (*Extension)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.FieldXPaths {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "FieldXPath", instances, (*FieldXPath)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.FractionDigits {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "FractionDigit", instances, (*FractionDigit)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Groups {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Group", instances, (*Group)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Keys {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Key", instances, (*Key)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.KeyRefs {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "KeyRef", instances, (*KeyRef)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lengths {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Length", instances, (*Length)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lists {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "List", instances, (*List)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.MaxExclusives {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "MaxExclusive", instances, (*MaxExclusive)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.MaxInclusives {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "MaxInclusive", instances, (*MaxInclusive)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.MaxLengths {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "MaxLength", instances, (*MaxLength)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinExclusives {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "MinExclusive", instances, (*MinExclusive)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinInclusives {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "MinInclusive", instances, (*MinInclusive)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinLengths {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "MinLength", instances, (*MinLength)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Patterns {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Pattern", instances, (*Pattern)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Restrictions {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Restriction", instances, (*Restriction)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Schemas {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Schema", instances, (*Schema)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Selectors {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Selector", instances, (*Selector)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Sequences {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Sequence", instances, (*Sequence)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.SimpleContents {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "SimpleContent", instances, (*SimpleContent)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.SimpleTypes {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "SimpleType", instances, (*SimpleType)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.TotalDigits {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "TotalDigit", instances, (*TotalDigit)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Unions {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Union", instances, (*Union)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.Uniques {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "Unique", instances, (*Unique)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.WhiteSpaces {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "WhiteSpace", instances, (*WhiteSpace)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.XSDImports {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "XSDImport", instances, (*XSDImport)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
		{
			var instances []GongstructIF
			for instance := range stage.XSDIncludes {
				instances = append(instances, instance)
			}
			if err := stage.SerializeCSVZip(zw, "XSDInclude", instances, (*XSDInclude)(nil).GongGetFieldHeaders(), addIDs); err != nil {
				_ = zw.Close()
				return nil, err
			}
		}
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.Alls {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "All", instances, (*All)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Annotations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Annotation", instances, (*Annotation)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Attributes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Attribute", instances, (*Attribute)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.AttributeGroups {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "AttributeGroup", instances, (*AttributeGroup)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Choices {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Choice", instances, (*Choice)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ComplexContents {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ComplexContent", instances, (*ComplexContent)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ComplexTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ComplexType", instances, (*ComplexType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Documentations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Documentation", instances, (*Documentation)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Elements {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Element", instances, (*Element)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Enumerations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Enumeration", instances, (*Enumeration)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Extensions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Extension", instances, (*Extension)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FieldXPaths {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FieldXPath", instances, (*FieldXPath)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.FractionDigits {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "FractionDigit", instances, (*FractionDigit)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Groups {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Group", instances, (*Group)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Keys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Key", instances, (*Key)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.KeyRefs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "KeyRef", instances, (*KeyRef)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lengths {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Length", instances, (*Length)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lists {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "List", instances, (*List)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MaxExclusives {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MaxExclusive", instances, (*MaxExclusive)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MaxInclusives {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MaxInclusive", instances, (*MaxInclusive)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MaxLengths {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MaxLength", instances, (*MaxLength)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinExclusives {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MinExclusive", instances, (*MinExclusive)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinInclusives {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MinInclusive", instances, (*MinInclusive)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.MinLengths {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "MinLength", instances, (*MinLength)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Patterns {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pattern", instances, (*Pattern)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Restrictions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Restriction", instances, (*Restriction)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Schemas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Schema", instances, (*Schema)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Selectors {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Selector", instances, (*Selector)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Sequences {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Sequence", instances, (*Sequence)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SimpleContents {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SimpleContent", instances, (*SimpleContent)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SimpleTypes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SimpleType", instances, (*SimpleType)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.TotalDigits {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "TotalDigit", instances, (*TotalDigit)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Unions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Union", instances, (*Union)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Uniques {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Unique", instances, (*Unique)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.WhiteSpaces {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "WhiteSpace", instances, (*WhiteSpace)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.XSDImports {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "XSDImport", instances, (*XSDImport)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.XSDIncludes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "XSDInclude", instances, (*XSDInclude)(nil).GongGetFieldHeaders(), addIDs)
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

// SerializeCSVZip writes a CSV entry named "<name>.csv" into zw with optional IDs.
func (stage *Stage) SerializeCSVZip(zw *zip.Writer, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) error {
	w, err := zw.Create(name + ".csv")
	if err != nil {
		return err
	}
	return stage.SerializeCSV(w, instances, fields, addIDs)
}

// SerializeCSV is the Stage method for CSV serialization with optional IDs.
func (stage *Stage) SerializeCSV(w io.Writer, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) error {
	cw := csv.NewWriter(w)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	if !addIDs {
		headers := make([]string, len(fields))
		for index, fieldHeader := range fields {
			headers[index] = fieldHeader.Name
		}
		if err := cw.Write(headers); err != nil {
			return err
		}
	} else {
		headers := make([]string, 2*len(fields))
		for index, fieldHeader := range fields {
			headers[2*index] = fieldHeader.Name
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				headers[2*index+1] = fieldHeader.Name + ":" + fieldHeader.TargetGongstructName + ":ID"
			case GongFieldValueTypeSliceOfPointers:
				headers[2*index+1] = fieldHeader.Name + ":" + fieldHeader.TargetGongstructName + ":IDs"
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					headers[2*index+1] = "ID"
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
					headers[2*index+1] = header
				}
			}
		}
		if err := cw.Write(headers); err != nil {
			return err
		}
	}

	for _, instance := range sortedSlice {
		if !addIDs {
			row := make([]string, len(fields))
			for index, fieldName := range fields {
				fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
				row[index] = fieldStringValue.GetValueString()
			}
			if err := cw.Write(row); err != nil {
				return err
			}
		} else {
			row := make([]string, 2*len(fields))
			for index, fieldName := range fields {
				fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
				row[2*index] = fieldStringValue.GetValueString()
				if index == 0 {
					row[2*index+1] = instance.GongGetUUID(stage)
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						row[2*index+1] = fieldStringValue.ids
					}
				}
			}
			if err := cw.Write(row); err != nil {
				return err
			}
		}
	}

	cw.Flush()
	return cw.Error()
}

// SerializeCSVPointer is the Stage method for CSV serialization.
func (stage *Stage) SerializeCSVPointer[Type GongstructPtr](w io.Writer) error {
	return stage.SerializeCSVPointer2[Type](w, false)
}

// SerializeCSVPointer2 is the Stage method for CSV serialization with optional IDs.
func (stage *Stage) SerializeCSVPointer2[Type GongstructPtr](w io.Writer, addIDs bool) error {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	return stage.SerializeCSV(w, instances, ret.GongGetFieldHeaders(), addIDs)
}

// SerializeCSVZipPointer is the Stage method for ZIP CSV serialization of a Gongstruct type.
func (stage *Stage) SerializeCSVZipPointer[Type GongstructPtr](zw *zip.Writer) error {
	return stage.SerializeCSVZipPointer2[Type](zw, false)
}

// SerializeCSVZipPointer2 is the Stage method for ZIP CSV serialization of a Gongstruct type with optional IDs.
func (stage *Stage) SerializeCSVZipPointer2[Type GongstructPtr](zw *zip.Writer, addIDs bool) error {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	return stage.SerializeCSVZip(zw, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
