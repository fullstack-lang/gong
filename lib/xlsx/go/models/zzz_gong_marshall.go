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

	// insertion point for declaration of instances to stage{{Identifiers}}

	// insertion point for initialization of values{{ValueInitializers}}

	// insertion point for setup of pointers{{PointersInitializers}}
}`

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`

// previous version does not hanldle embedded structs (https://github.com/golang/go/issues/9859)
// simpler version but the name of the instance cannot be human read before the fields initialization
const IdentifiersDeclsWithoutNameInit = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)` /* */

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
	var identifiersDecl strings.Builder
	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder

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
		identifiersDecl.WriteString("\n")
	}
	for _, displayselection := range displayselectionOrdered {

		identifiersDecl.WriteString(displayselection.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(displayselection.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(displayselection.GongMarshallField(stage, "XLFile"))
		pointersInitializesStatements.WriteString(displayselection.GongMarshallField(stage, "XLSheet"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, xlcell := range xlcellOrdered {

		identifiersDecl.WriteString(xlcell.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlcell.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlcell.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(xlcell.GongMarshallField(stage, "Y"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, xlfile := range xlfileOrdered {

		identifiersDecl.WriteString(xlfile.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlfile.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlfile.GongMarshallField(stage, "NbSheets"))
		pointersInitializesStatements.WriteString(xlfile.GongMarshallField(stage, "Sheets"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, xlrow := range xlrowOrdered {

		identifiersDecl.WriteString(xlrow.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlrow.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlrow.GongMarshallField(stage, "RowIndex"))
		pointersInitializesStatements.WriteString(xlrow.GongMarshallField(stage, "Cells"))
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
		identifiersDecl.WriteString("\n")
	}
	for _, xlsheet := range xlsheetOrdered {

		identifiersDecl.WriteString(xlsheet.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "MaxRow"))
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "MaxCol"))
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "NbRows"))
		pointersInitializesStatements.WriteString(xlsheet.GongMarshallField(stage, "Rows"))
		pointersInitializesStatements.WriteString(xlsheet.GongMarshallField(stage, "SheetCells"))
	}

	// insertion initialization of objects to stage
	for _, displayselection := range displayselectionOrdered {
		_ = displayselection
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, xlcell := range xlcellOrdered {
		_ = xlcell
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, xlfile := range xlfileOrdered {
		_ = xlfile
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, xlrow := range xlrowOrdered {
		_ = xlrow
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, xlsheet := range xlsheetOrdered {
		_ = xlsheet
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
func (displayselection *DisplaySelection) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(displayselection.Name))

	case "XLFile":
		if displayselection.XLFile != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XLFile")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", displayselection.XLFile.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XLFile")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "XLSheet":
		if displayselection.XLSheet != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XLSheet")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", displayselection.XLSheet.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", displayselection.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "XLSheet")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DisplaySelection", fieldName)
	}
	return
}

func (xlcell *XLCell) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xlcell.Name))
	case "X":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "X")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.X))
	case "Y":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlcell.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Y")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlcell.Y))

	default:
		log.Panicf("Unknown field %s for Gongstruct XLCell", fieldName)
	}
	return
}

func (xlfile *XLFile) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xlfile.Name))
	case "NbSheets":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbSheets")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlfile.NbSheets))

	case "Sheets":
		var sb strings.Builder
		for _, _xlsheet := range xlfile.Sheets {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", xlfile.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Sheets")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _xlsheet.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct XLFile", fieldName)
	}
	return
}

func (xlrow *XLRow) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xlrow.Name))
	case "RowIndex":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RowIndex")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlrow.RowIndex))

	case "Cells":
		var sb strings.Builder
		for _, _xlcell := range xlrow.Cells {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", xlrow.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Cells")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _xlcell.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct XLRow", fieldName)
	}
	return
}

func (xlsheet *XLSheet) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xlsheet.Name))
	case "MaxRow":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MaxRow")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxRow))
	case "MaxCol":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MaxCol")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.MaxCol))
	case "NbRows":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "NbRows")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", xlsheet.NbRows))

	case "Rows":
		var sb strings.Builder
		for _, _xlrow := range xlsheet.Rows {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Rows")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _xlrow.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SheetCells":
		var sb strings.Builder
		for _, _xlcell := range xlsheet.SheetCells {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", xlsheet.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SheetCells")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _xlcell.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct XLSheet", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (displayselection *DisplaySelection) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(displayselection.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(displayselection.GongMarshallField(stage, "XLFile"))
		pointersInitializesStatements.WriteString(displayselection.GongMarshallField(stage, "XLSheet"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (xlcell *XLCell) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlcell.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlcell.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(xlcell.GongMarshallField(stage, "Y"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (xlfile *XLFile) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlfile.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlfile.GongMarshallField(stage, "NbSheets"))
		pointersInitializesStatements.WriteString(xlfile.GongMarshallField(stage, "Sheets"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (xlrow *XLRow) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlrow.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlrow.GongMarshallField(stage, "RowIndex"))
		pointersInitializesStatements.WriteString(xlrow.GongMarshallField(stage, "Cells"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (xlsheet *XLSheet) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "MaxRow"))
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "MaxCol"))
		initializerStatements.WriteString(xlsheet.GongMarshallField(stage, "NbRows"))
		pointersInitializesStatements.WriteString(xlsheet.GongMarshallField(stage, "Rows"))
		pointersInitializesStatements.WriteString(xlsheet.GongMarshallField(stage, "SheetCells"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
