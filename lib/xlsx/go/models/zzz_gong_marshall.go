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
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`/* */

const StringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + "`" + `{{GeneratedFieldNameValue}}` + "`"

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
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
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
	for _, displayselection := range displayselectionOrdered {
	
		identifiersDecl += displayselection.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(displayselection.Name))
		initializerStatements += setValueField

	}

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
	for _, xlcell := range xlcellOrdered {
	
		identifiersDecl += xlcell.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlcell.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.X))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.Y))
		initializerStatements += setValueField

	}

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
	for _, xlfile := range xlfileOrdered {
	
		identifiersDecl += xlfile.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlfile.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbSheets")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlfile.NbSheets))
		initializerStatements += setValueField

	}

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
	for _, xlrow := range xlrowOrdered {
	
		identifiersDecl += xlrow.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlrow.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RowIndex")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlrow.RowIndex))
		initializerStatements += setValueField

	}

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
	for _, xlsheet := range xlsheetOrdered {
	
		identifiersDecl += xlsheet.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsheet.Name))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxRow")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxRow))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxCol")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxCol))
		initializerStatements += setValueField

		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbRows")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.NbRows))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(displayselectionOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of DisplaySelection instances pointers"
	}
	for _, displayselection := range displayselectionOrdered {
		_ = displayselection
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		if displayselection.XLFile != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "XLFile")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", displayselection.XLFile.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		if displayselection.XLSheet != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "XLSheet")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", displayselection.XLSheet.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlcellOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLCell instances pointers"
	}
	for _, xlcell := range xlcellOrdered {
		_ = xlcell
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
	}

	if len(xlfileOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLFile instances pointers"
	}
	for _, xlfile := range xlfileOrdered {
		_ = xlfile
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _xlsheet := range xlfile.Sheets {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sheets")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlsheet.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlrowOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLRow instances pointers"
	}
	for _, xlrow := range xlrowOrdered {
		_ = xlrow
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _xlcell := range xlrow.Cells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlcell.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	}

	if len(xlsheetOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of XLSheet instances pointers"
	}
	for _, xlsheet := range xlsheetOrdered {
		_ = xlsheet
		var setPointerField string
		_ = setPointerField

		// Initialisation of values
		for _, _xlrow := range xlsheet.Rows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlrow.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

		for _, _xlcell := range xlsheet.SheetCells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SheetCells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlcell.GongGetIdentifier(stage))
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
	return
}

// insertion initialization of objects to stage
func (displayselection *DisplaySelection) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(displayselection.Name))
		initializerStatements += setValueField

	case "XLFile":
		if displayselection.XLFile != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "XLFile")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", displayselection.XLFile.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "XLSheet":
		if displayselection.XLSheet != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "XLSheet")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", displayselection.XLSheet.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct DisplaySelection", fieldName)
	}
	return
}

func (xlcell *XLCell) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlcell.Name))
		initializerStatements += setValueField
	case "X":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "X")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.X))
		initializerStatements += setValueField
	case "Y":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Y")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.Y))
		initializerStatements += setValueField


	default:
		log.Panicf("Unknown field %s for Gongstruct XLCell", fieldName)
	}
	return
}

func (xlfile *XLFile) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlfile.Name))
		initializerStatements += setValueField
	case "NbSheets":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbSheets")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlfile.NbSheets))
		initializerStatements += setValueField

	case "Sheets":
		for _, _xlsheet := range xlfile.Sheets {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Sheets")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlsheet.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct XLFile", fieldName)
	}
	return
}

func (xlrow *XLRow) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlrow.Name))
		initializerStatements += setValueField
	case "RowIndex":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "RowIndex")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlrow.RowIndex))
		initializerStatements += setValueField

	case "Cells":
		for _, _xlcell := range xlrow.Cells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlcell.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct XLRow", fieldName)
	}
	return
}

func (xlsheet *XLSheet) GongMarshallField(stage *Stage, fieldName string) (setValueField, setPointerField string) {
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(xlsheet.Name))
		initializerStatements += setValueField
	case "MaxRow":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxRow")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxRow))
		initializerStatements += setValueField
	case "MaxCol":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxCol")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxCol))
		initializerStatements += setValueField
	case "NbRows":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbRows")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.NbRows))
		initializerStatements += setValueField

	case "Rows":
		for _, _xlrow := range xlsheet.Rows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlrow.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "SheetCells":
		for _, _xlcell := range xlsheet.SheetCells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "SheetCells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _xlcell.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct XLSheet", fieldName)
	}
	return
}
