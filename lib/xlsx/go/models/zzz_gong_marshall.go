// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const marshallRes = `package {{PackageName}}

import (
	"time"

	"{{ModelsPackageName}}"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.Stage) {

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

const MetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}` + "{}"

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

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, modelsPackageName, packageName string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Fatalln(name + " is not a go filename")
	}

	log.Printf("%s Marshalling %s", time.Now().Format("2006-01-02 15:04:05.000000"), name)
	newBase := filepath.Base(file.Name())

	res := marshallRes
	res = strings.ReplaceAll(res, "{{databaseName}}", strings.ReplaceAll(newBase, ".go", ""))
	res = strings.ReplaceAll(res, "{{PackageName}}", packageName)
	res = strings.ReplaceAll(res, "{{ModelsPackageName}}", modelsPackageName)

	// map of identifiers
	// var StageMapDstructIds map[*Dstruct]string
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_DisplaySelection_Identifiers := make(map[*DisplaySelection]string)
	_ = map_DisplaySelection_Identifiers

	displayselectionOrdered := []*DisplaySelection{}
	for displayselection := range stage.DisplaySelections {
		displayselectionOrdered = append(displayselectionOrdered, displayselection)
	}
	sort.Slice(displayselectionOrdered[:], func(i, j int) bool {
		displayselectioni := displayselectionOrdered[i]
		displayselectionj := displayselectionOrdered[j]
		displayselectioni_order, oki := stage.DisplaySelectionMap_Staged_Order[displayselectioni]
		displayselectionj_order, okj := stage.DisplaySelectionMap_Staged_Order[displayselectionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return displayselectioni_order < displayselectionj_order
	})
	if len(displayselectionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, displayselection := range displayselectionOrdered {

		id = generatesIdentifier("DisplaySelection", idx, displayselection.Name)
		map_DisplaySelection_Identifiers[displayselection] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DisplaySelection")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", displayselection.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(displayselection.Name))
		initializerStatements += setValueField

	}

	map_XLCell_Identifiers := make(map[*XLCell]string)
	_ = map_XLCell_Identifiers

	xlcellOrdered := []*XLCell{}
	for xlcell := range stage.XLCells {
		xlcellOrdered = append(xlcellOrdered, xlcell)
	}
	sort.Slice(xlcellOrdered[:], func(i, j int) bool {
		xlcelli := xlcellOrdered[i]
		xlcellj := xlcellOrdered[j]
		xlcelli_order, oki := stage.XLCellMap_Staged_Order[xlcelli]
		xlcellj_order, okj := stage.XLCellMap_Staged_Order[xlcellj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xlcelli_order < xlcellj_order
	})
	if len(xlcellOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xlcell := range xlcellOrdered {

		id = generatesIdentifier("XLCell", idx, xlcell.Name)
		map_XLCell_Identifiers[xlcell] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLCell")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlcell.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlcell.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.Y))
		initializerStatements += setValueField

	}

	map_XLFile_Identifiers := make(map[*XLFile]string)
	_ = map_XLFile_Identifiers

	xlfileOrdered := []*XLFile{}
	for xlfile := range stage.XLFiles {
		xlfileOrdered = append(xlfileOrdered, xlfile)
	}
	sort.Slice(xlfileOrdered[:], func(i, j int) bool {
		xlfilei := xlfileOrdered[i]
		xlfilej := xlfileOrdered[j]
		xlfilei_order, oki := stage.XLFileMap_Staged_Order[xlfilei]
		xlfilej_order, okj := stage.XLFileMap_Staged_Order[xlfilej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xlfilei_order < xlfilej_order
	})
	if len(xlfileOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xlfile := range xlfileOrdered {

		id = generatesIdentifier("XLFile", idx, xlfile.Name)
		map_XLFile_Identifiers[xlfile] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLFile")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlfile.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlfile.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbSheets")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlfile.NbSheets))
		initializerStatements += setValueField

	}

	map_XLRow_Identifiers := make(map[*XLRow]string)
	_ = map_XLRow_Identifiers

	xlrowOrdered := []*XLRow{}
	for xlrow := range stage.XLRows {
		xlrowOrdered = append(xlrowOrdered, xlrow)
	}
	sort.Slice(xlrowOrdered[:], func(i, j int) bool {
		xlrowi := xlrowOrdered[i]
		xlrowj := xlrowOrdered[j]
		xlrowi_order, oki := stage.XLRowMap_Staged_Order[xlrowi]
		xlrowj_order, okj := stage.XLRowMap_Staged_Order[xlrowj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xlrowi_order < xlrowj_order
	})
	if len(xlrowOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xlrow := range xlrowOrdered {

		id = generatesIdentifier("XLRow", idx, xlrow.Name)
		map_XLRow_Identifiers[xlrow] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLRow")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlrow.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlrow.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RowIndex")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlrow.RowIndex))
		initializerStatements += setValueField

	}

	map_XLSheet_Identifiers := make(map[*XLSheet]string)
	_ = map_XLSheet_Identifiers

	xlsheetOrdered := []*XLSheet{}
	for xlsheet := range stage.XLSheets {
		xlsheetOrdered = append(xlsheetOrdered, xlsheet)
	}
	sort.Slice(xlsheetOrdered[:], func(i, j int) bool {
		xlsheeti := xlsheetOrdered[i]
		xlsheetj := xlsheetOrdered[j]
		xlsheeti_order, oki := stage.XLSheetMap_Staged_Order[xlsheeti]
		xlsheetj_order, okj := stage.XLSheetMap_Staged_Order[xlsheetj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xlsheeti_order < xlsheetj_order
	})
	if len(xlsheetOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for idx, xlsheet := range xlsheetOrdered {

		id = generatesIdentifier("XLSheet", idx, xlsheet.Name)
		map_XLSheet_Identifiers[xlsheet] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XLSheet")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xlsheet.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsheet.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxRow")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxRow))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxCol")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxCol))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbRows")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.NbRows))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(displayselectionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DisplaySelection instances pointers"
	}
	for idx, displayselection := range displayselectionOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("DisplaySelection", idx, displayselection.Name)
		map_DisplaySelection_Identifiers[displayselection] = id

		// Initialisation of values
		if displayselection.XLFile != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "XLFile")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XLFile_Identifiers[displayselection.XLFile])
			pointersInitializesStatements += setPointerField
		}

		if displayselection.XLSheet != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "XLSheet")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XLSheet_Identifiers[displayselection.XLSheet])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlcellOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLCell instances pointers"
	}
	for idx, xlcell := range xlcellOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XLCell", idx, xlcell.Name)
		map_XLCell_Identifiers[xlcell] = id

		// Initialisation of values
	}

	if len(xlfileOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLFile instances pointers"
	}
	for idx, xlfile := range xlfileOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XLFile", idx, xlfile.Name)
		map_XLFile_Identifiers[xlfile] = id

		// Initialisation of values
		for _, _xlsheet := range xlfile.Sheets {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sheets")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XLSheet_Identifiers[_xlsheet])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlrowOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLRow instances pointers"
	}
	for idx, xlrow := range xlrowOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XLRow", idx, xlrow.Name)
		map_XLRow_Identifiers[xlrow] = id

		// Initialisation of values
		for _, _xlcell := range xlrow.Cells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XLCell_Identifiers[_xlcell])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlsheetOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLSheet instances pointers"
	}
	for idx, xlsheet := range xlsheetOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("XLSheet", idx, xlsheet.Name)
		map_XLSheet_Identifiers[xlsheet] = id

		// Initialisation of values
		for _, _xlrow := range xlsheet.Rows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XLRow_Identifiers[_xlrow])
			pointersInitializesStatements += setPointerField
		}

		for _, _xlcell := range xlsheet.SheetCells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SheetCells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_XLCell_Identifiers[_xlcell])
			pointersInitializesStatements += setPointerField
		}

	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	if stage.MetaPackageImportAlias != "" {
		res = strings.ReplaceAll(res, "{{ImportPackageDeclaration}}",
			fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))

		res = strings.ReplaceAll(res, "{{ImportPackageDummyDeclaration}}",
			fmt.Sprintf("\nvar _ %s.Stage",
				stage.MetaPackageImportAlias))

		var entries string

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
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(0),", value.Ident, value.Ident)
			case GONG__ENUM_CAST_STRING:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s(\"\"),", value.Ident, value.Ident)
			case GONG__FIELD_VALUE:
				// substitute the second point with "{})."
				joker := "__substitute_for_first_point__"
				valueIdentifier := strings.Replace(value.Ident, ".", joker, 1)
				valueIdentifier = strings.Replace(valueIdentifier, ".", "{}).", 1)
				valueIdentifier = strings.Replace(valueIdentifier, joker, ".", 1)
				entries += fmt.Sprintf("\n\n\t\"%s\": (%s,", value.Ident, valueIdentifier)
			case GONG__IDENTIFIER_CONST:
				entries += fmt.Sprintf("\n\n\t\"%s\": %s,", value.Ident, value.Ident)
			case GONG__STRUCT_INSTANCE:
				entries += fmt.Sprintf("\n\n\t\"%s\": &(%s{}),", value.Ident, value.Ident)
			}
		}

		// res = strings.ReplaceAll(res, "{{EntriesDocLinkStringDocLinkIdentifier}}", entries)
	}

	fmt.Fprintln(file, res)
}

// unique identifier per struct
func generatesIdentifier(gongStructName string, idx int, instanceName string) (identifier string) {

	identifier = instanceName
	// Make a Regex to say we only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	processedString := reg.ReplaceAllString(instanceName, "_")

	identifier = fmt.Sprintf("__%s__%06d_%s", gongStructName, idx, processedString)

	return
}
