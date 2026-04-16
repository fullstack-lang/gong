// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const marshallRes = `package {{PackageName}}

import (
	"slices"
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const StringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const NumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const PointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const SliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const TimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// ToRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func ToRawStringLiteral(s string) string {
	// Step 1: Replace every backtick with a closing backtick,
	// a double-quoted backtick, and an opening backtick.
	escaped := strings.ReplaceAll(s, "`", "` + \"`\" + `")

	// Step 2: Wrap the entire resulting string in backticks.
	result := "`" + escaped + "`"

	// Step 3: Clean up any empty raw strings (``) at the boundaries
	// just in case your original string started or ended with a backtick.
	result = strings.ReplaceAll(result, "`` + ", "")
	result = strings.ReplaceAll(result, " + ``", "")

	return result
}

// MarshallFile marshall the stage content into a file as an instanciation into a stage
// according to the marshalling policy of the stage.
//
// In GongMarshallingAppendCommit mode, it will append the last commit to the file.
// In other modes, it will rewrite the entire file.
func (stage *Stage) MarshallFile(filename, modelsPackageName, packageName string) {

	if stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		contentBytes, err := os.ReadFile(filename)

		// if the file does not exist, marshall the full stage
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, modelsPackageName, packageName)
			return
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		content := string(contentBytes)

		if stage.isSquashing {
			// we squash: we want to clear the current function body
			// and let the append logic write the squashed commit
			firstBrace := strings.Index(content, "func _(stage *models.Stage) {")
			if firstBrace != -1 {
				firstBrace += len("func _(stage *models.Stage) {")
				content = content[:firstBrace] + "\n}"
			}
		}

		if stage.isApplyingBackwardCommit {
			// we are going backward, we need to remove the last forward commit from the file

			// because commitsBehind has been incremented before the call to this function
			// the index of the commit to remove is len(forwardCommits) - commitsBehind
			commitIndexToRemove := len(stage.forwardCommits) - stage.GetCommitsBehind()

			if commitIndexToRemove < 0 || commitIndexToRemove >= len(stage.forwardCommits) {
				return // Should not happen if history is consistent
			}

			commitToRemove := stage.forwardCommits[commitIndexToRemove]

			lastIndex := strings.LastIndex(content, commitToRemove+"\n")
			if lastIndex != -1 {
				newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove)+1:]
				err = os.WriteFile(filename, []byte(newContent), 0644)
				if err != nil {
					log.Fatal(err.Error())
				}
			} else {
				lastIndex = strings.LastIndex(content, commitToRemove)
				if lastIndex != -1 {
					newContent := content[:lastIndex] + content[lastIndex+len(commitToRemove):]
					err = os.WriteFile(filename, []byte(newContent), 0644)
					if err != nil {
						log.Fatal(err.Error())
					}
				} else {
					// The commit block was not found. This typically happens for the initial
					// commit which is formatted differently (the lines after func _(stage *models.Stage) {).
					// We rewrite the entire file with the current (rewound) stage state to safely remove it.
					file, createErr := os.Create(filename)
					if createErr != nil {
						log.Fatal(createErr.Error())
					}
					defer file.Close()
					stage.Marshall(file, modelsPackageName, packageName)
				}
			}
			return // we are done for the backward case
		}

		if stage.isApplyingForwardCommit {
			// bypass the modified check
		} else if !stage.modified {
			return
		}

		forwardCommits := stage.GetForwardCommits()
		if len(forwardCommits) == 0 {
			return // nothing to do
		}

		activeCommits := len(forwardCommits) - stage.GetCommitsBehind()
		if activeCommits <= 0 {
			return
		}
		forwardCommit := forwardCommits[activeCommits-1]

		// append before the ending brace of the func
		lastBrace := strings.LastIndex(content, "}")
		if lastBrace == -1 {
			// if no brace, something is wrong with the file, so we rewrite it
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, modelsPackageName, packageName)
			return
		}

		contentBeforeBrace := content[:lastBrace]
		trimmedContentBeforeBrace := strings.TrimSpace(contentBeforeBrace)
		emptyBody := stage.isSquashing ||
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// insertion point for setup of pointers")

		// check if the file ends with stage.Commit() before the brace
		if !emptyBody && !strings.HasSuffix(trimmedContentBeforeBrace, "stage.Commit()") {
			contentBeforeBrace = contentBeforeBrace + "\n\tstage.Commit()\n"
		}

		// insert the commit statements before the last brace
		newContent := contentBeforeBrace + forwardCommit + "\n" + content[lastBrace:]

		err = os.WriteFile(filename, []byte(newContent), 0644)
		if err != nil {
			log.Fatal(err.Error())
		}

	} else {
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Marshall(file, modelsPackageName, packageName)
	}
}

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(modelsPackageName, packageName)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(modelsPackageName, packageName string) (res string, err error) {

	res = marshallRes
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	buttonOrdered := []*Button{}
	for button := range stage.Buttons {
		buttonOrdered = append(buttonOrdered, button)
	}
	sort.Slice(buttonOrdered[:], func(i, j int) bool {
		buttoni := buttonOrdered[i]
		buttonj := buttonOrdered[j]
		buttoni_order, oki := stage.Button_stagedOrder[buttoni]
		buttonj_order, okj := stage.Button_stagedOrder[buttonj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return buttoni_order < buttonj_order
	})
	if len(buttonOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, button := range buttonOrdered {

		identifiersDecl.WriteString(button.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(button.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "Icon"))
		pointersInitializesStatements.WriteString(button.GongMarshallField(stage, "SVGIcon"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "IsDisabled"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "ToolTipPosition"))
	}

	cellOrdered := []*Cell{}
	for cell := range stage.Cells {
		cellOrdered = append(cellOrdered, cell)
	}
	sort.Slice(cellOrdered[:], func(i, j int) bool {
		celli := cellOrdered[i]
		cellj := cellOrdered[j]
		celli_order, oki := stage.Cell_stagedOrder[celli]
		cellj_order, okj := stage.Cell_stagedOrder[cellj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return celli_order < cellj_order
	})
	if len(cellOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, cell := range cellOrdered {

		identifiersDecl.WriteString(cell.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(cell.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellString"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellFloat64"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellInt"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellBool"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellIcon"))
	}

	cellbooleanOrdered := []*CellBoolean{}
	for cellboolean := range stage.CellBooleans {
		cellbooleanOrdered = append(cellbooleanOrdered, cellboolean)
	}
	sort.Slice(cellbooleanOrdered[:], func(i, j int) bool {
		cellbooleani := cellbooleanOrdered[i]
		cellbooleanj := cellbooleanOrdered[j]
		cellbooleani_order, oki := stage.CellBoolean_stagedOrder[cellbooleani]
		cellbooleanj_order, okj := stage.CellBoolean_stagedOrder[cellbooleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellbooleani_order < cellbooleanj_order
	})
	if len(cellbooleanOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, cellboolean := range cellbooleanOrdered {

		identifiersDecl.WriteString(cellboolean.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellboolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellboolean.GongMarshallField(stage, "Value"))
	}

	cellfloat64Ordered := []*CellFloat64{}
	for cellfloat64 := range stage.CellFloat64s {
		cellfloat64Ordered = append(cellfloat64Ordered, cellfloat64)
	}
	sort.Slice(cellfloat64Ordered[:], func(i, j int) bool {
		cellfloat64i := cellfloat64Ordered[i]
		cellfloat64j := cellfloat64Ordered[j]
		cellfloat64i_order, oki := stage.CellFloat64_stagedOrder[cellfloat64i]
		cellfloat64j_order, okj := stage.CellFloat64_stagedOrder[cellfloat64j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellfloat64i_order < cellfloat64j_order
	})
	if len(cellfloat64Ordered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, cellfloat64 := range cellfloat64Ordered {

		identifiersDecl.WriteString(cellfloat64.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellfloat64.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellfloat64.GongMarshallField(stage, "Value"))
	}

	celliconOrdered := []*CellIcon{}
	for cellicon := range stage.CellIcons {
		celliconOrdered = append(celliconOrdered, cellicon)
	}
	sort.Slice(celliconOrdered[:], func(i, j int) bool {
		celliconi := celliconOrdered[i]
		celliconj := celliconOrdered[j]
		celliconi_order, oki := stage.CellIcon_stagedOrder[celliconi]
		celliconj_order, okj := stage.CellIcon_stagedOrder[celliconj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return celliconi_order < celliconj_order
	})
	if len(celliconOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, cellicon := range celliconOrdered {

		identifiersDecl.WriteString(cellicon.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "Icon"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "NeedsConfirmation"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "ConfirmationMessage"))
	}

	cellintOrdered := []*CellInt{}
	for cellint := range stage.CellInts {
		cellintOrdered = append(cellintOrdered, cellint)
	}
	sort.Slice(cellintOrdered[:], func(i, j int) bool {
		cellinti := cellintOrdered[i]
		cellintj := cellintOrdered[j]
		cellinti_order, oki := stage.CellInt_stagedOrder[cellinti]
		cellintj_order, okj := stage.CellInt_stagedOrder[cellintj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellinti_order < cellintj_order
	})
	if len(cellintOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, cellint := range cellintOrdered {

		identifiersDecl.WriteString(cellint.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellint.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellint.GongMarshallField(stage, "Value"))
	}

	cellstringOrdered := []*CellString{}
	for cellstring := range stage.CellStrings {
		cellstringOrdered = append(cellstringOrdered, cellstring)
	}
	sort.Slice(cellstringOrdered[:], func(i, j int) bool {
		cellstringi := cellstringOrdered[i]
		cellstringj := cellstringOrdered[j]
		cellstringi_order, oki := stage.CellString_stagedOrder[cellstringi]
		cellstringj_order, okj := stage.CellString_stagedOrder[cellstringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellstringi_order < cellstringj_order
	})
	if len(cellstringOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, cellstring := range cellstringOrdered {

		identifiersDecl.WriteString(cellstring.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellstring.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellstring.GongMarshallField(stage, "Value"))
	}

	displayedcolumnOrdered := []*DisplayedColumn{}
	for displayedcolumn := range stage.DisplayedColumns {
		displayedcolumnOrdered = append(displayedcolumnOrdered, displayedcolumn)
	}
	sort.Slice(displayedcolumnOrdered[:], func(i, j int) bool {
		displayedcolumni := displayedcolumnOrdered[i]
		displayedcolumnj := displayedcolumnOrdered[j]
		displayedcolumni_order, oki := stage.DisplayedColumn_stagedOrder[displayedcolumni]
		displayedcolumnj_order, okj := stage.DisplayedColumn_stagedOrder[displayedcolumnj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return displayedcolumni_order < displayedcolumnj_order
	})
	if len(displayedcolumnOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, displayedcolumn := range displayedcolumnOrdered {

		identifiersDecl.WriteString(displayedcolumn.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(displayedcolumn.GongMarshallField(stage, "Name"))
	}

	rowOrdered := []*Row{}
	for row := range stage.Rows {
		rowOrdered = append(rowOrdered, row)
	}
	sort.Slice(rowOrdered[:], func(i, j int) bool {
		rowi := rowOrdered[i]
		rowj := rowOrdered[j]
		rowi_order, oki := stage.Row_stagedOrder[rowi]
		rowj_order, okj := stage.Row_stagedOrder[rowj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rowi_order < rowj_order
	})
	if len(rowOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, row := range rowOrdered {

		identifiersDecl.WriteString(row.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(row.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(row.GongMarshallField(stage, "Cells"))
		initializerStatements.WriteString(row.GongMarshallField(stage, "IsChecked"))
	}

	svgiconOrdered := []*SVGIcon{}
	for svgicon := range stage.SVGIcons {
		svgiconOrdered = append(svgiconOrdered, svgicon)
	}
	sort.Slice(svgiconOrdered[:], func(i, j int) bool {
		svgiconi := svgiconOrdered[i]
		svgiconj := svgiconOrdered[j]
		svgiconi_order, oki := stage.SVGIcon_stagedOrder[svgiconi]
		svgiconj_order, okj := stage.SVGIcon_stagedOrder[svgiconj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return svgiconi_order < svgiconj_order
	})
	if len(svgiconOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, svgicon := range svgiconOrdered {

		identifiersDecl.WriteString(svgicon.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(svgicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svgicon.GongMarshallField(stage, "SVG"))
	}

	tableOrdered := []*Table{}
	for table := range stage.Tables {
		tableOrdered = append(tableOrdered, table)
	}
	sort.Slice(tableOrdered[:], func(i, j int) bool {
		tablei := tableOrdered[i]
		tablej := tableOrdered[j]
		tablei_order, oki := stage.Table_stagedOrder[tablei]
		tablej_order, okj := stage.Table_stagedOrder[tablej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tablei_order < tablej_order
	})
	if len(tableOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, table := range tableOrdered {

		identifiersDecl.WriteString(table.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(table.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "DisplayedColumns"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "Rows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasFiltering"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasColumnSorting"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasPaginator"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasCheckableRows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasSaveButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "SaveButtonLabel"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasBulkDeleteButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "BulkDeleteButtonTooltip"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "BulkDeleteSelectedRowsIDsJson"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "CanDragDropRows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasCloseButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "SavingInProgress"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "NbOfStickyColumns"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "Buttons"))
	}

	// insertion initialization of objects to stage
	for _, button := range buttonOrdered {
		_ = button
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, cell := range cellOrdered {
		_ = cell
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, cellboolean := range cellbooleanOrdered {
		_ = cellboolean
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, cellfloat64 := range cellfloat64Ordered {
		_ = cellfloat64
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, cellicon := range celliconOrdered {
		_ = cellicon
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, cellint := range cellintOrdered {
		_ = cellint
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, cellstring := range cellstringOrdered {
		_ = cellstring
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, displayedcolumn := range displayedcolumnOrdered {
		_ = displayedcolumn
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, row := range rowOrdered {
		_ = row
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, svgicon := range svgiconOrdered {
		_ = svgicon
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, table := range tableOrdered {
		_ = table
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl.String())
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements.String())
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements.String())

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries strings.Builder

		// regenerate the map of doc link renaming
		// the key and value are set to the value because
		// if it has been renamed, this is the new value that matters
		valuesOrdered := make([]GONG__Identifier, 0)
		for _, value := range stage.Map_DocLink_Renaming {
			valuesOrdered = append(valuesOrdered, value)
		}
		sort.Slice(valuesOrdered[:], func(i, j int) bool {
			return valuesOrdered[i].Ident < valuesOrdered[j].Ident
		})
		for _, value := range valuesOrdered {

			// get the number of points in the value to find if it is a field
			// or a struct

			switch value.Type {
			case GONG__ENUM_CAST_INT:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident))
			case GONG__ENUM_CAST_STRING:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident))
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier))
			case GONG__IDENTIFIER_CONST:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident))
			case GONG__STRUCT_INSTANCE:
				entries.WriteString(fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident))
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries.String())
	}
	return
}

// insertion point for marshall field methods
func (button *Button) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(button.Name))
	case "Icon":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Icon")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(button.Icon))
	case "IsDisabled":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDisabled")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", button.IsDisabled))
	case "HasToolTip":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasToolTip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", button.HasToolTip))
	case "ToolTipText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(button.ToolTipText))
	case "ToolTipPosition":
		if button.ToolTipPosition.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+button.ToolTipPosition.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ToolTipPosition")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	case "SVGIcon":
		if button.SVGIcon != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVGIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", button.SVGIcon.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", button.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVGIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Button", fieldName)
	}
	return
}

func (cell *Cell) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cell.Name))

	case "CellString":
		if cell.CellString != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellString")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", cell.CellString.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellString")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "CellFloat64":
		if cell.CellFloat64 != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellFloat64")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", cell.CellFloat64.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellFloat64")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "CellInt":
		if cell.CellInt != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellInt")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", cell.CellInt.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellInt")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "CellBool":
		if cell.CellBool != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellBool")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", cell.CellBool.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellBool")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "CellIcon":
		if cell.CellIcon != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", cell.CellIcon.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", cell.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CellIcon")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Cell", fieldName)
	}
	return
}

func (cellboolean *CellBoolean) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellboolean.Name))
	case "Value":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", cellboolean.Value))

	default:
		log.Panicf("Unknown field %s for Gongstruct CellBoolean", fieldName)
	}
	return
}

func (cellfloat64 *CellFloat64) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellfloat64.Name))
	case "Value":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", cellfloat64.Value))

	default:
		log.Panicf("Unknown field %s for Gongstruct CellFloat64", fieldName)
	}
	return
}

func (cellicon *CellIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellicon.Name))
	case "Icon":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Icon")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellicon.Icon))
	case "NeedsConfirmation":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NeedsConfirmation")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", cellicon.NeedsConfirmation))
	case "ConfirmationMessage":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ConfirmationMessage")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellicon.ConfirmationMessage))

	default:
		log.Panicf("Unknown field %s for Gongstruct CellIcon", fieldName)
	}
	return
}

func (cellint *CellInt) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellint.Name))
	case "Value":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellint.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", cellint.Value))

	default:
		log.Panicf("Unknown field %s for Gongstruct CellInt", fieldName)
	}
	return
}

func (cellstring *CellString) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellstring.Name))
	case "Value":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Value")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(cellstring.Value))

	default:
		log.Panicf("Unknown field %s for Gongstruct CellString", fieldName)
	}
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", displayedcolumn.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(displayedcolumn.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct DisplayedColumn", fieldName)
	}
	return
}

func (row *Row) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", row.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(row.Name))
	case "IsChecked":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", row.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsChecked")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", row.IsChecked))

	case "Cells":
		var sb strings.Builder
		for _, _cell := range row.Cells {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", row.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Cells")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _cell.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Row", fieldName)
	}
	return
}

func (svgicon *SVGIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(svgicon.Name))
	case "SVG":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgicon.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SVG")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(svgicon.SVG))

	default:
		log.Panicf("Unknown field %s for Gongstruct SVGIcon", fieldName)
	}
	return
}

func (table *Table) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(table.Name))
	case "HasFiltering":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasFiltering")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasFiltering))
	case "HasColumnSorting":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasColumnSorting")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasColumnSorting))
	case "HasPaginator":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasPaginator")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasPaginator))
	case "HasCheckableRows":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasCheckableRows")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasCheckableRows))
	case "HasSaveButton":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasSaveButton")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasSaveButton))
	case "SaveButtonLabel":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SaveButtonLabel")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(table.SaveButtonLabel))
	case "HasBulkDeleteButton":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasBulkDeleteButton")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasBulkDeleteButton))
	case "BulkDeleteButtonTooltip":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BulkDeleteButtonTooltip")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(table.BulkDeleteButtonTooltip))
	case "BulkDeleteSelectedRowsIDsJson":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BulkDeleteSelectedRowsIDsJson")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(table.BulkDeleteSelectedRowsIDsJson))
	case "CanDragDropRows":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CanDragDropRows")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.CanDragDropRows))
	case "HasCloseButton":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "HasCloseButton")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasCloseButton))
	case "SavingInProgress":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SavingInProgress")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.SavingInProgress))
	case "NbOfStickyColumns":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", table.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbOfStickyColumns")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", table.NbOfStickyColumns))

	case "DisplayedColumns":
		var sb strings.Builder
		for _, _displayedcolumn := range table.DisplayedColumns {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", table.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DisplayedColumns")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _displayedcolumn.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Rows":
		var sb strings.Builder
		for _, _row := range table.Rows {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", table.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Rows")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _row.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Buttons":
		var sb strings.Builder
		for _, _button := range table.Buttons {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", table.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Buttons")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _button.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (button *Button) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(button.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "Icon"))
		pointersInitializesStatements.WriteString(button.GongMarshallField(stage, "SVGIcon"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "IsDisabled"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "HasToolTip"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "ToolTipText"))
		initializerStatements.WriteString(button.GongMarshallField(stage, "ToolTipPosition"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cell *Cell) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cell.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellString"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellFloat64"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellInt"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellBool"))
		pointersInitializesStatements.WriteString(cell.GongMarshallField(stage, "CellIcon"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellboolean *CellBoolean) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellboolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellboolean.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellfloat64 *CellFloat64) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellfloat64.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellfloat64.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellicon *CellIcon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "Icon"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "NeedsConfirmation"))
		initializerStatements.WriteString(cellicon.GongMarshallField(stage, "ConfirmationMessage"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellint *CellInt) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellint.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellint.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (cellstring *CellString) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(cellstring.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(cellstring.GongMarshallField(stage, "Value"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (displayedcolumn *DisplayedColumn) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(displayedcolumn.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (row *Row) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(row.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(row.GongMarshallField(stage, "Cells"))
		initializerStatements.WriteString(row.GongMarshallField(stage, "IsChecked"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (svgicon *SVGIcon) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(svgicon.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svgicon.GongMarshallField(stage, "SVG"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (table *Table) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(table.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "DisplayedColumns"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "Rows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasFiltering"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasColumnSorting"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasPaginator"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasCheckableRows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasSaveButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "SaveButtonLabel"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasBulkDeleteButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "BulkDeleteButtonTooltip"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "BulkDeleteSelectedRowsIDsJson"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "CanDragDropRows"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "HasCloseButton"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "SavingInProgress"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "NbOfStickyColumns"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "Buttons"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
