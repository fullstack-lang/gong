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
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: ` + "`" + `{{GeneratedFieldNameValue}}` + "`" + `}).Stage(stage)`

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
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	cellOrdered := []*Cell{}
	for cell := range stage.Cells {
		cellOrdered = append(cellOrdered, cell)
	}
	sort.Slice(cellOrdered[:], func(i, j int) bool {
		celli := cellOrdered[i]
		cellj := cellOrdered[j]
		celli_order, oki := stage.CellMap_Staged_Order[celli]
		cellj_order, okj := stage.CellMap_Staged_Order[cellj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return celli_order < cellj_order
	})
	if len(cellOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, cell := range cellOrdered {

		identifiersDecl += cell.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += cell.GongMarshallField(stage, "Name")
		pointersInitializesStatements += cell.GongMarshallField(stage, "CellString")
		pointersInitializesStatements += cell.GongMarshallField(stage, "CellFloat64")
		pointersInitializesStatements += cell.GongMarshallField(stage, "CellInt")
		pointersInitializesStatements += cell.GongMarshallField(stage, "CellBool")
		pointersInitializesStatements += cell.GongMarshallField(stage, "CellIcon")
	}

	cellbooleanOrdered := []*CellBoolean{}
	for cellboolean := range stage.CellBooleans {
		cellbooleanOrdered = append(cellbooleanOrdered, cellboolean)
	}
	sort.Slice(cellbooleanOrdered[:], func(i, j int) bool {
		cellbooleani := cellbooleanOrdered[i]
		cellbooleanj := cellbooleanOrdered[j]
		cellbooleani_order, oki := stage.CellBooleanMap_Staged_Order[cellbooleani]
		cellbooleanj_order, okj := stage.CellBooleanMap_Staged_Order[cellbooleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellbooleani_order < cellbooleanj_order
	})
	if len(cellbooleanOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, cellboolean := range cellbooleanOrdered {

		identifiersDecl += cellboolean.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += cellboolean.GongMarshallField(stage, "Name")
		initializerStatements += cellboolean.GongMarshallField(stage, "Value")
	}

	cellfloat64Ordered := []*CellFloat64{}
	for cellfloat64 := range stage.CellFloat64s {
		cellfloat64Ordered = append(cellfloat64Ordered, cellfloat64)
	}
	sort.Slice(cellfloat64Ordered[:], func(i, j int) bool {
		cellfloat64i := cellfloat64Ordered[i]
		cellfloat64j := cellfloat64Ordered[j]
		cellfloat64i_order, oki := stage.CellFloat64Map_Staged_Order[cellfloat64i]
		cellfloat64j_order, okj := stage.CellFloat64Map_Staged_Order[cellfloat64j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellfloat64i_order < cellfloat64j_order
	})
	if len(cellfloat64Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, cellfloat64 := range cellfloat64Ordered {

		identifiersDecl += cellfloat64.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += cellfloat64.GongMarshallField(stage, "Name")
		initializerStatements += cellfloat64.GongMarshallField(stage, "Value")
	}

	celliconOrdered := []*CellIcon{}
	for cellicon := range stage.CellIcons {
		celliconOrdered = append(celliconOrdered, cellicon)
	}
	sort.Slice(celliconOrdered[:], func(i, j int) bool {
		celliconi := celliconOrdered[i]
		celliconj := celliconOrdered[j]
		celliconi_order, oki := stage.CellIconMap_Staged_Order[celliconi]
		celliconj_order, okj := stage.CellIconMap_Staged_Order[celliconj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return celliconi_order < celliconj_order
	})
	if len(celliconOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, cellicon := range celliconOrdered {

		identifiersDecl += cellicon.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += cellicon.GongMarshallField(stage, "Name")
		initializerStatements += cellicon.GongMarshallField(stage, "Icon")
		initializerStatements += cellicon.GongMarshallField(stage, "NeedsConfirmation")
		initializerStatements += cellicon.GongMarshallField(stage, "ConfirmationMessage")
	}

	cellintOrdered := []*CellInt{}
	for cellint := range stage.CellInts {
		cellintOrdered = append(cellintOrdered, cellint)
	}
	sort.Slice(cellintOrdered[:], func(i, j int) bool {
		cellinti := cellintOrdered[i]
		cellintj := cellintOrdered[j]
		cellinti_order, oki := stage.CellIntMap_Staged_Order[cellinti]
		cellintj_order, okj := stage.CellIntMap_Staged_Order[cellintj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellinti_order < cellintj_order
	})
	if len(cellintOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, cellint := range cellintOrdered {

		identifiersDecl += cellint.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += cellint.GongMarshallField(stage, "Name")
		initializerStatements += cellint.GongMarshallField(stage, "Value")
	}

	cellstringOrdered := []*CellString{}
	for cellstring := range stage.CellStrings {
		cellstringOrdered = append(cellstringOrdered, cellstring)
	}
	sort.Slice(cellstringOrdered[:], func(i, j int) bool {
		cellstringi := cellstringOrdered[i]
		cellstringj := cellstringOrdered[j]
		cellstringi_order, oki := stage.CellStringMap_Staged_Order[cellstringi]
		cellstringj_order, okj := stage.CellStringMap_Staged_Order[cellstringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return cellstringi_order < cellstringj_order
	})
	if len(cellstringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, cellstring := range cellstringOrdered {

		identifiersDecl += cellstring.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += cellstring.GongMarshallField(stage, "Name")
		initializerStatements += cellstring.GongMarshallField(stage, "Value")
	}

	checkboxOrdered := []*CheckBox{}
	for checkbox := range stage.CheckBoxs {
		checkboxOrdered = append(checkboxOrdered, checkbox)
	}
	sort.Slice(checkboxOrdered[:], func(i, j int) bool {
		checkboxi := checkboxOrdered[i]
		checkboxj := checkboxOrdered[j]
		checkboxi_order, oki := stage.CheckBoxMap_Staged_Order[checkboxi]
		checkboxj_order, okj := stage.CheckBoxMap_Staged_Order[checkboxj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return checkboxi_order < checkboxj_order
	})
	if len(checkboxOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, checkbox := range checkboxOrdered {

		identifiersDecl += checkbox.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += checkbox.GongMarshallField(stage, "Name")
		initializerStatements += checkbox.GongMarshallField(stage, "Value")
	}

	displayedcolumnOrdered := []*DisplayedColumn{}
	for displayedcolumn := range stage.DisplayedColumns {
		displayedcolumnOrdered = append(displayedcolumnOrdered, displayedcolumn)
	}
	sort.Slice(displayedcolumnOrdered[:], func(i, j int) bool {
		displayedcolumni := displayedcolumnOrdered[i]
		displayedcolumnj := displayedcolumnOrdered[j]
		displayedcolumni_order, oki := stage.DisplayedColumnMap_Staged_Order[displayedcolumni]
		displayedcolumnj_order, okj := stage.DisplayedColumnMap_Staged_Order[displayedcolumnj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return displayedcolumni_order < displayedcolumnj_order
	})
	if len(displayedcolumnOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, displayedcolumn := range displayedcolumnOrdered {

		identifiersDecl += displayedcolumn.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += displayedcolumn.GongMarshallField(stage, "Name")
	}

	formdivOrdered := []*FormDiv{}
	for formdiv := range stage.FormDivs {
		formdivOrdered = append(formdivOrdered, formdiv)
	}
	sort.Slice(formdivOrdered[:], func(i, j int) bool {
		formdivi := formdivOrdered[i]
		formdivj := formdivOrdered[j]
		formdivi_order, oki := stage.FormDivMap_Staged_Order[formdivi]
		formdivj_order, okj := stage.FormDivMap_Staged_Order[formdivj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formdivi_order < formdivj_order
	})
	if len(formdivOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formdiv := range formdivOrdered {

		identifiersDecl += formdiv.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formdiv.GongMarshallField(stage, "Name")
		pointersInitializesStatements += formdiv.GongMarshallField(stage, "FormFields")
		pointersInitializesStatements += formdiv.GongMarshallField(stage, "CheckBoxs")
		pointersInitializesStatements += formdiv.GongMarshallField(stage, "FormEditAssocButton")
		pointersInitializesStatements += formdiv.GongMarshallField(stage, "FormSortAssocButton")
	}

	formeditassocbuttonOrdered := []*FormEditAssocButton{}
	for formeditassocbutton := range stage.FormEditAssocButtons {
		formeditassocbuttonOrdered = append(formeditassocbuttonOrdered, formeditassocbutton)
	}
	sort.Slice(formeditassocbuttonOrdered[:], func(i, j int) bool {
		formeditassocbuttoni := formeditassocbuttonOrdered[i]
		formeditassocbuttonj := formeditassocbuttonOrdered[j]
		formeditassocbuttoni_order, oki := stage.FormEditAssocButtonMap_Staged_Order[formeditassocbuttoni]
		formeditassocbuttonj_order, okj := stage.FormEditAssocButtonMap_Staged_Order[formeditassocbuttonj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formeditassocbuttoni_order < formeditassocbuttonj_order
	})
	if len(formeditassocbuttonOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formeditassocbutton := range formeditassocbuttonOrdered {

		identifiersDecl += formeditassocbutton.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "Name")
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "Label")
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "AssociationStorage")
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "HasChanged")
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "IsForSavePurpose")
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "HasToolTip")
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "ToolTipText")
		initializerStatements += formeditassocbutton.GongMarshallField(stage, "MatTooltipShowDelay")
	}

	formfieldOrdered := []*FormField{}
	for formfield := range stage.FormFields {
		formfieldOrdered = append(formfieldOrdered, formfield)
	}
	sort.Slice(formfieldOrdered[:], func(i, j int) bool {
		formfieldi := formfieldOrdered[i]
		formfieldj := formfieldOrdered[j]
		formfieldi_order, oki := stage.FormFieldMap_Staged_Order[formfieldi]
		formfieldj_order, okj := stage.FormFieldMap_Staged_Order[formfieldj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldi_order < formfieldj_order
	})
	if len(formfieldOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfield := range formfieldOrdered {

		identifiersDecl += formfield.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfield.GongMarshallField(stage, "Name")
		initializerStatements += formfield.GongMarshallField(stage, "InputTypeEnum")
		initializerStatements += formfield.GongMarshallField(stage, "Label")
		initializerStatements += formfield.GongMarshallField(stage, "Placeholder")
		pointersInitializesStatements += formfield.GongMarshallField(stage, "FormFieldString")
		pointersInitializesStatements += formfield.GongMarshallField(stage, "FormFieldFloat64")
		pointersInitializesStatements += formfield.GongMarshallField(stage, "FormFieldInt")
		pointersInitializesStatements += formfield.GongMarshallField(stage, "FormFieldDate")
		pointersInitializesStatements += formfield.GongMarshallField(stage, "FormFieldTime")
		pointersInitializesStatements += formfield.GongMarshallField(stage, "FormFieldDateTime")
		pointersInitializesStatements += formfield.GongMarshallField(stage, "FormFieldSelect")
		initializerStatements += formfield.GongMarshallField(stage, "HasBespokeWidth")
		initializerStatements += formfield.GongMarshallField(stage, "BespokeWidthPx")
		initializerStatements += formfield.GongMarshallField(stage, "HasBespokeHeight")
		initializerStatements += formfield.GongMarshallField(stage, "BespokeHeightPx")
	}

	formfielddateOrdered := []*FormFieldDate{}
	for formfielddate := range stage.FormFieldDates {
		formfielddateOrdered = append(formfielddateOrdered, formfielddate)
	}
	sort.Slice(formfielddateOrdered[:], func(i, j int) bool {
		formfielddatei := formfielddateOrdered[i]
		formfielddatej := formfielddateOrdered[j]
		formfielddatei_order, oki := stage.FormFieldDateMap_Staged_Order[formfielddatei]
		formfielddatej_order, okj := stage.FormFieldDateMap_Staged_Order[formfielddatej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfielddatei_order < formfielddatej_order
	})
	if len(formfielddateOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfielddate := range formfielddateOrdered {

		identifiersDecl += formfielddate.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfielddate.GongMarshallField(stage, "Name")
		initializerStatements += formfielddate.GongMarshallField(stage, "Value")
	}

	formfielddatetimeOrdered := []*FormFieldDateTime{}
	for formfielddatetime := range stage.FormFieldDateTimes {
		formfielddatetimeOrdered = append(formfielddatetimeOrdered, formfielddatetime)
	}
	sort.Slice(formfielddatetimeOrdered[:], func(i, j int) bool {
		formfielddatetimei := formfielddatetimeOrdered[i]
		formfielddatetimej := formfielddatetimeOrdered[j]
		formfielddatetimei_order, oki := stage.FormFieldDateTimeMap_Staged_Order[formfielddatetimei]
		formfielddatetimej_order, okj := stage.FormFieldDateTimeMap_Staged_Order[formfielddatetimej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfielddatetimei_order < formfielddatetimej_order
	})
	if len(formfielddatetimeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfielddatetime := range formfielddatetimeOrdered {

		identifiersDecl += formfielddatetime.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfielddatetime.GongMarshallField(stage, "Name")
		initializerStatements += formfielddatetime.GongMarshallField(stage, "Value")
	}

	formfieldfloat64Ordered := []*FormFieldFloat64{}
	for formfieldfloat64 := range stage.FormFieldFloat64s {
		formfieldfloat64Ordered = append(formfieldfloat64Ordered, formfieldfloat64)
	}
	sort.Slice(formfieldfloat64Ordered[:], func(i, j int) bool {
		formfieldfloat64i := formfieldfloat64Ordered[i]
		formfieldfloat64j := formfieldfloat64Ordered[j]
		formfieldfloat64i_order, oki := stage.FormFieldFloat64Map_Staged_Order[formfieldfloat64i]
		formfieldfloat64j_order, okj := stage.FormFieldFloat64Map_Staged_Order[formfieldfloat64j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldfloat64i_order < formfieldfloat64j_order
	})
	if len(formfieldfloat64Ordered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfieldfloat64 := range formfieldfloat64Ordered {

		identifiersDecl += formfieldfloat64.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfieldfloat64.GongMarshallField(stage, "Name")
		initializerStatements += formfieldfloat64.GongMarshallField(stage, "Value")
		initializerStatements += formfieldfloat64.GongMarshallField(stage, "HasMinValidator")
		initializerStatements += formfieldfloat64.GongMarshallField(stage, "MinValue")
		initializerStatements += formfieldfloat64.GongMarshallField(stage, "HasMaxValidator")
		initializerStatements += formfieldfloat64.GongMarshallField(stage, "MaxValue")
	}

	formfieldintOrdered := []*FormFieldInt{}
	for formfieldint := range stage.FormFieldInts {
		formfieldintOrdered = append(formfieldintOrdered, formfieldint)
	}
	sort.Slice(formfieldintOrdered[:], func(i, j int) bool {
		formfieldinti := formfieldintOrdered[i]
		formfieldintj := formfieldintOrdered[j]
		formfieldinti_order, oki := stage.FormFieldIntMap_Staged_Order[formfieldinti]
		formfieldintj_order, okj := stage.FormFieldIntMap_Staged_Order[formfieldintj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldinti_order < formfieldintj_order
	})
	if len(formfieldintOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfieldint := range formfieldintOrdered {

		identifiersDecl += formfieldint.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfieldint.GongMarshallField(stage, "Name")
		initializerStatements += formfieldint.GongMarshallField(stage, "Value")
		initializerStatements += formfieldint.GongMarshallField(stage, "HasMinValidator")
		initializerStatements += formfieldint.GongMarshallField(stage, "MinValue")
		initializerStatements += formfieldint.GongMarshallField(stage, "HasMaxValidator")
		initializerStatements += formfieldint.GongMarshallField(stage, "MaxValue")
	}

	formfieldselectOrdered := []*FormFieldSelect{}
	for formfieldselect := range stage.FormFieldSelects {
		formfieldselectOrdered = append(formfieldselectOrdered, formfieldselect)
	}
	sort.Slice(formfieldselectOrdered[:], func(i, j int) bool {
		formfieldselecti := formfieldselectOrdered[i]
		formfieldselectj := formfieldselectOrdered[j]
		formfieldselecti_order, oki := stage.FormFieldSelectMap_Staged_Order[formfieldselecti]
		formfieldselectj_order, okj := stage.FormFieldSelectMap_Staged_Order[formfieldselectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldselecti_order < formfieldselectj_order
	})
	if len(formfieldselectOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfieldselect := range formfieldselectOrdered {

		identifiersDecl += formfieldselect.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfieldselect.GongMarshallField(stage, "Name")
		pointersInitializesStatements += formfieldselect.GongMarshallField(stage, "Value")
		pointersInitializesStatements += formfieldselect.GongMarshallField(stage, "Options")
		initializerStatements += formfieldselect.GongMarshallField(stage, "CanBeEmpty")
		initializerStatements += formfieldselect.GongMarshallField(stage, "PreserveInitialOrder")
	}

	formfieldstringOrdered := []*FormFieldString{}
	for formfieldstring := range stage.FormFieldStrings {
		formfieldstringOrdered = append(formfieldstringOrdered, formfieldstring)
	}
	sort.Slice(formfieldstringOrdered[:], func(i, j int) bool {
		formfieldstringi := formfieldstringOrdered[i]
		formfieldstringj := formfieldstringOrdered[j]
		formfieldstringi_order, oki := stage.FormFieldStringMap_Staged_Order[formfieldstringi]
		formfieldstringj_order, okj := stage.FormFieldStringMap_Staged_Order[formfieldstringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldstringi_order < formfieldstringj_order
	})
	if len(formfieldstringOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfieldstring := range formfieldstringOrdered {

		identifiersDecl += formfieldstring.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfieldstring.GongMarshallField(stage, "Name")
		initializerStatements += formfieldstring.GongMarshallField(stage, "Value")
		initializerStatements += formfieldstring.GongMarshallField(stage, "IsTextArea")
	}

	formfieldtimeOrdered := []*FormFieldTime{}
	for formfieldtime := range stage.FormFieldTimes {
		formfieldtimeOrdered = append(formfieldtimeOrdered, formfieldtime)
	}
	sort.Slice(formfieldtimeOrdered[:], func(i, j int) bool {
		formfieldtimei := formfieldtimeOrdered[i]
		formfieldtimej := formfieldtimeOrdered[j]
		formfieldtimei_order, oki := stage.FormFieldTimeMap_Staged_Order[formfieldtimei]
		formfieldtimej_order, okj := stage.FormFieldTimeMap_Staged_Order[formfieldtimej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formfieldtimei_order < formfieldtimej_order
	})
	if len(formfieldtimeOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formfieldtime := range formfieldtimeOrdered {

		identifiersDecl += formfieldtime.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formfieldtime.GongMarshallField(stage, "Name")
		initializerStatements += formfieldtime.GongMarshallField(stage, "Value")
		initializerStatements += formfieldtime.GongMarshallField(stage, "Step")
	}

	formgroupOrdered := []*FormGroup{}
	for formgroup := range stage.FormGroups {
		formgroupOrdered = append(formgroupOrdered, formgroup)
	}
	sort.Slice(formgroupOrdered[:], func(i, j int) bool {
		formgroupi := formgroupOrdered[i]
		formgroupj := formgroupOrdered[j]
		formgroupi_order, oki := stage.FormGroupMap_Staged_Order[formgroupi]
		formgroupj_order, okj := stage.FormGroupMap_Staged_Order[formgroupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formgroupi_order < formgroupj_order
	})
	if len(formgroupOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formgroup := range formgroupOrdered {

		identifiersDecl += formgroup.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formgroup.GongMarshallField(stage, "Name")
		initializerStatements += formgroup.GongMarshallField(stage, "Label")
		pointersInitializesStatements += formgroup.GongMarshallField(stage, "FormDivs")
		initializerStatements += formgroup.GongMarshallField(stage, "HasSuppressButton")
		initializerStatements += formgroup.GongMarshallField(stage, "HasSuppressButtonBeenPressed")
	}

	formsortassocbuttonOrdered := []*FormSortAssocButton{}
	for formsortassocbutton := range stage.FormSortAssocButtons {
		formsortassocbuttonOrdered = append(formsortassocbuttonOrdered, formsortassocbutton)
	}
	sort.Slice(formsortassocbuttonOrdered[:], func(i, j int) bool {
		formsortassocbuttoni := formsortassocbuttonOrdered[i]
		formsortassocbuttonj := formsortassocbuttonOrdered[j]
		formsortassocbuttoni_order, oki := stage.FormSortAssocButtonMap_Staged_Order[formsortassocbuttoni]
		formsortassocbuttonj_order, okj := stage.FormSortAssocButtonMap_Staged_Order[formsortassocbuttonj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return formsortassocbuttoni_order < formsortassocbuttonj_order
	})
	if len(formsortassocbuttonOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, formsortassocbutton := range formsortassocbuttonOrdered {

		identifiersDecl += formsortassocbutton.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += formsortassocbutton.GongMarshallField(stage, "Name")
		initializerStatements += formsortassocbutton.GongMarshallField(stage, "Label")
		initializerStatements += formsortassocbutton.GongMarshallField(stage, "HasToolTip")
		initializerStatements += formsortassocbutton.GongMarshallField(stage, "ToolTipText")
		initializerStatements += formsortassocbutton.GongMarshallField(stage, "MatTooltipShowDelay")
		pointersInitializesStatements += formsortassocbutton.GongMarshallField(stage, "FormEditAssocButton")
	}

	optionOrdered := []*Option{}
	for option := range stage.Options {
		optionOrdered = append(optionOrdered, option)
	}
	sort.Slice(optionOrdered[:], func(i, j int) bool {
		optioni := optionOrdered[i]
		optionj := optionOrdered[j]
		optioni_order, oki := stage.OptionMap_Staged_Order[optioni]
		optionj_order, okj := stage.OptionMap_Staged_Order[optionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return optioni_order < optionj_order
	})
	if len(optionOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, option := range optionOrdered {

		identifiersDecl += option.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += option.GongMarshallField(stage, "Name")
	}

	rowOrdered := []*Row{}
	for row := range stage.Rows {
		rowOrdered = append(rowOrdered, row)
	}
	sort.Slice(rowOrdered[:], func(i, j int) bool {
		rowi := rowOrdered[i]
		rowj := rowOrdered[j]
		rowi_order, oki := stage.RowMap_Staged_Order[rowi]
		rowj_order, okj := stage.RowMap_Staged_Order[rowj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return rowi_order < rowj_order
	})
	if len(rowOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, row := range rowOrdered {

		identifiersDecl += row.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += row.GongMarshallField(stage, "Name")
		pointersInitializesStatements += row.GongMarshallField(stage, "Cells")
		initializerStatements += row.GongMarshallField(stage, "IsChecked")
	}

	tableOrdered := []*Table{}
	for table := range stage.Tables {
		tableOrdered = append(tableOrdered, table)
	}
	sort.Slice(tableOrdered[:], func(i, j int) bool {
		tablei := tableOrdered[i]
		tablej := tableOrdered[j]
		tablei_order, oki := stage.TableMap_Staged_Order[tablei]
		tablej_order, okj := stage.TableMap_Staged_Order[tablej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return tablei_order < tablej_order
	})
	if len(tableOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, table := range tableOrdered {

		identifiersDecl += table.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += table.GongMarshallField(stage, "Name")
		pointersInitializesStatements += table.GongMarshallField(stage, "DisplayedColumns")
		pointersInitializesStatements += table.GongMarshallField(stage, "Rows")
		initializerStatements += table.GongMarshallField(stage, "HasFiltering")
		initializerStatements += table.GongMarshallField(stage, "HasColumnSorting")
		initializerStatements += table.GongMarshallField(stage, "HasPaginator")
		initializerStatements += table.GongMarshallField(stage, "HasCheckableRows")
		initializerStatements += table.GongMarshallField(stage, "HasSaveButton")
		initializerStatements += table.GongMarshallField(stage, "SaveButtonLabel")
		initializerStatements += table.GongMarshallField(stage, "CanDragDropRows")
		initializerStatements += table.GongMarshallField(stage, "HasCloseButton")
		initializerStatements += table.GongMarshallField(stage, "SavingInProgress")
		initializerStatements += table.GongMarshallField(stage, "NbOfStickyColumns")
	}

	// insertion initialization of objects to stage
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

	for _, checkbox := range checkboxOrdered {
		_ = checkbox
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

	for _, formdiv := range formdivOrdered {
		_ = formdiv
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formeditassocbutton := range formeditassocbuttonOrdered {
		_ = formeditassocbutton
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfield := range formfieldOrdered {
		_ = formfield
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfielddate := range formfielddateOrdered {
		_ = formfielddate
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfielddatetime := range formfielddatetimeOrdered {
		_ = formfielddatetime
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldfloat64 := range formfieldfloat64Ordered {
		_ = formfieldfloat64
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldint := range formfieldintOrdered {
		_ = formfieldint
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldselect := range formfieldselectOrdered {
		_ = formfieldselect
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldstring := range formfieldstringOrdered {
		_ = formfieldstring
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formfieldtime := range formfieldtimeOrdered {
		_ = formfieldtime
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formgroup := range formgroupOrdered {
		_ = formgroup
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, formsortassocbutton := range formsortassocbuttonOrdered {
		_ = formsortassocbutton
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, option := range optionOrdered {
		_ = option
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

	for _, table := range tableOrdered {
		_ = table
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
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
func (cell *Cell) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cell.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cell.Name))
		initializerStatements += setValueField

	case "CellString":
		if cell.CellString != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", cell.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CellString")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", cell.CellString.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "CellFloat64":
		if cell.CellFloat64 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", cell.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CellFloat64")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", cell.CellFloat64.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "CellInt":
		if cell.CellInt != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", cell.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CellInt")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", cell.CellInt.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "CellBool":
		if cell.CellBool != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", cell.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CellBool")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", cell.CellBool.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "CellIcon":
		if cell.CellIcon != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", cell.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CellIcon")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", cell.CellIcon.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Cell", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (cellboolean *CellBoolean) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellboolean.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellboolean.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", cellboolean.Value))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct CellBoolean", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (cellfloat64 *CellFloat64) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellfloat64.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", cellfloat64.Value))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct CellFloat64", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (cellicon *CellIcon) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellicon.Name))
		initializerStatements += setValueField
	case "Icon":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Icon")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellicon.Icon))
		initializerStatements += setValueField
	case "NeedsConfirmation":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NeedsConfirmation")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", cellicon.NeedsConfirmation))
		initializerStatements += setValueField
	case "ConfirmationMessage":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellicon.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ConfirmationMessage")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellicon.ConfirmationMessage))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct CellIcon", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (cellint *CellInt) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellint.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", cellint.Value))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct CellInt", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (cellstring *CellString) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellstring.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", cellstring.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(cellstring.Value))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct CellString", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (checkbox *CheckBox) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(checkbox.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", checkbox.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", checkbox.Value))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct CheckBox", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (displayedcolumn *DisplayedColumn) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", displayedcolumn.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(displayedcolumn.Name))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct DisplayedColumn", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formdiv *FormDiv) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formdiv.Name))
		initializerStatements += setValueField

	case "FormFields":
		for _, _formfield := range formdiv.FormFields {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFields")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _formfield.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "CheckBoxs":
		for _, _checkbox := range formdiv.CheckBoxs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "CheckBoxs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _checkbox.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormEditAssocButton":
		if formdiv.FormEditAssocButton != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormEditAssocButton")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formdiv.FormEditAssocButton.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormSortAssocButton":
		if formdiv.FormSortAssocButton != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formdiv.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormSortAssocButton")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formdiv.FormSortAssocButton.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormDiv", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formeditassocbutton *FormEditAssocButton) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formeditassocbutton.Name))
		initializerStatements += setValueField
	case "Label":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formeditassocbutton.Label))
		initializerStatements += setValueField
	case "AssociationStorage":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "AssociationStorage")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formeditassocbutton.AssociationStorage))
		initializerStatements += setValueField
	case "HasChanged":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasChanged")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formeditassocbutton.HasChanged))
		initializerStatements += setValueField
	case "IsForSavePurpose":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsForSavePurpose")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formeditassocbutton.IsForSavePurpose))
		initializerStatements += setValueField
	case "HasToolTip":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasToolTip")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formeditassocbutton.HasToolTip))
		initializerStatements += setValueField
	case "ToolTipText":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ToolTipText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formeditassocbutton.ToolTipText))
		initializerStatements += setValueField
	case "MatTooltipShowDelay":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formeditassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MatTooltipShowDelay")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formeditassocbutton.MatTooltipShowDelay))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct FormEditAssocButton", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfield *FormField) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfield.Name))
		initializerStatements += setValueField
	case "InputTypeEnum":
		if formfield.InputTypeEnum != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "InputTypeEnum")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+formfield.InputTypeEnum.ToCodeString())
			initializerStatements += setValueField
		}
	case "Label":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfield.Label))
		initializerStatements += setValueField
	case "Placeholder":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Placeholder")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfield.Placeholder))
		initializerStatements += setValueField
	case "HasBespokeWidth":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasBespokeWidth")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfield.HasBespokeWidth))
		initializerStatements += setValueField
	case "BespokeWidthPx":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeWidthPx")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfield.BespokeWidthPx))
		initializerStatements += setValueField
	case "HasBespokeHeight":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasBespokeHeight")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfield.HasBespokeHeight))
		initializerStatements += setValueField
	case "BespokeHeightPx":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "BespokeHeightPx")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfield.BespokeHeightPx))
		initializerStatements += setValueField

	case "FormFieldString":
		if formfield.FormFieldString != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldString")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfield.FormFieldString.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormFieldFloat64":
		if formfield.FormFieldFloat64 != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldFloat64")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfield.FormFieldFloat64.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormFieldInt":
		if formfield.FormFieldInt != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldInt")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfield.FormFieldInt.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormFieldDate":
		if formfield.FormFieldDate != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldDate")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfield.FormFieldDate.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormFieldTime":
		if formfield.FormFieldTime != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldTime")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfield.FormFieldTime.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormFieldDateTime":
		if formfield.FormFieldDateTime != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldDateTime")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfield.FormFieldDateTime.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "FormFieldSelect":
		if formfield.FormFieldSelect != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfield.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormFieldSelect")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfield.FormFieldSelect.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormField", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfielddate *FormFieldDate) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfielddate.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfielddate.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", formfielddate.Value.String())
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldDate", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfielddatetime *FormFieldDateTime) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfielddatetime.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfielddatetime.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", formfielddatetime.Value.String())
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldDateTime", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfieldfloat64 *FormFieldFloat64) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldfloat64.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.Value))
		initializerStatements += setValueField
	case "HasMinValidator":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMinValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldfloat64.HasMinValidator))
		initializerStatements += setValueField
	case "MinValue":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.MinValue))
		initializerStatements += setValueField
	case "HasMaxValidator":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMaxValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldfloat64.HasMaxValidator))
		initializerStatements += setValueField
	case "MaxValue":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldfloat64.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldfloat64.MaxValue))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldFloat64", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfieldint *FormFieldInt) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldint.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.Value))
		initializerStatements += setValueField
	case "HasMinValidator":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMinValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldint.HasMinValidator))
		initializerStatements += setValueField
	case "MinValue":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MinValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.MinValue))
		initializerStatements += setValueField
	case "HasMaxValidator":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasMaxValidator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldint.HasMaxValidator))
		initializerStatements += setValueField
	case "MaxValue":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldint.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MaxValue")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", formfieldint.MaxValue))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldInt", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfieldselect *FormFieldSelect) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldselect.Name))
		initializerStatements += setValueField
	case "CanBeEmpty":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanBeEmpty")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldselect.CanBeEmpty))
		initializerStatements += setValueField
	case "PreserveInitialOrder":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "PreserveInitialOrder")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldselect.PreserveInitialOrder))
		initializerStatements += setValueField

	case "Value":
		if formfieldselect.Value != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Value")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formfieldselect.Value.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Options":
		for _, _option := range formfieldselect.Options {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formfieldselect.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Options")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _option.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldSelect", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfieldstring *FormFieldString) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldstring.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldstring.Value))
		initializerStatements += setValueField
	case "IsTextArea":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldstring.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsTextArea")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formfieldstring.IsTextArea))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldString", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formfieldtime *FormFieldTime) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formfieldtime.Name))
		initializerStatements += setValueField
	case "Value":
		setValueField = TimeInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Value")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", formfieldtime.Value.String())
		initializerStatements += setValueField
	case "Step":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formfieldtime.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Step")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", formfieldtime.Step))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct FormFieldTime", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formgroup *FormGroup) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formgroup.Name))
		initializerStatements += setValueField
	case "Label":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formgroup.Label))
		initializerStatements += setValueField
	case "HasSuppressButton":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasSuppressButton")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formgroup.HasSuppressButton))
		initializerStatements += setValueField
	case "HasSuppressButtonBeenPressed":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasSuppressButtonBeenPressed")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formgroup.HasSuppressButtonBeenPressed))
		initializerStatements += setValueField

	case "FormDivs":
		for _, _formdiv := range formgroup.FormDivs {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formgroup.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormDivs")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _formdiv.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormGroup", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (formsortassocbutton *FormSortAssocButton) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formsortassocbutton.Name))
		initializerStatements += setValueField
	case "Label":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Label")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formsortassocbutton.Label))
		initializerStatements += setValueField
	case "HasToolTip":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasToolTip")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", formsortassocbutton.HasToolTip))
		initializerStatements += setValueField
	case "ToolTipText":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ToolTipText")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formsortassocbutton.ToolTipText))
		initializerStatements += setValueField
	case "MatTooltipShowDelay":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MatTooltipShowDelay")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(formsortassocbutton.MatTooltipShowDelay))
		initializerStatements += setValueField

	case "FormEditAssocButton":
		if formsortassocbutton.FormEditAssocButton != nil {
			setPointerField = PointerFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", formsortassocbutton.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "FormEditAssocButton")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", formsortassocbutton.FormEditAssocButton.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct FormSortAssocButton", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (option *Option) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", option.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(option.Name))
		initializerStatements += setValueField

	default:
		log.Panicf("Unknown field %s for Gongstruct Option", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (row *Row) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", row.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(row.Name))
		initializerStatements += setValueField
	case "IsChecked":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", row.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "IsChecked")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", row.IsChecked))
		initializerStatements += setValueField

	case "Cells":
		for _, _cell := range row.Cells {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", row.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Cells")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _cell.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Row", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}

func (table *Table) GongMarshallField(stage *Stage, fieldName string) (res string) {
	var setValueField, setPointerField string
	_ = setValueField
	_ = setPointerField
	initializerStatements := ""
	_ = initializerStatements
	pointersInitializesStatements := ""
	_ = pointersInitializesStatements

	switch fieldName {
	case "Name":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.Name))
		initializerStatements += setValueField
	case "HasFiltering":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasFiltering")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasFiltering))
		initializerStatements += setValueField
	case "HasColumnSorting":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasColumnSorting")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasColumnSorting))
		initializerStatements += setValueField
	case "HasPaginator":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasPaginator")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasPaginator))
		initializerStatements += setValueField
	case "HasCheckableRows":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasCheckableRows")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasCheckableRows))
		initializerStatements += setValueField
	case "HasSaveButton":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasSaveButton")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasSaveButton))
		initializerStatements += setValueField
	case "SaveButtonLabel":
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SaveButtonLabel")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(table.SaveButtonLabel))
		initializerStatements += setValueField
	case "CanDragDropRows":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "CanDragDropRows")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.CanDragDropRows))
		initializerStatements += setValueField
	case "HasCloseButton":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "HasCloseButton")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.HasCloseButton))
		initializerStatements += setValueField
	case "SavingInProgress":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "SavingInProgress")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", table.SavingInProgress))
		initializerStatements += setValueField
	case "NbOfStickyColumns":
		setValueField = NumberInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", table.GongGetIdentifier(stage))
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "NbOfStickyColumns")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", table.NbOfStickyColumns))
		initializerStatements += setValueField

	case "DisplayedColumns":
		for _, _displayedcolumn := range table.DisplayedColumns {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "DisplayedColumns")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _displayedcolumn.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	case "Rows":
		for _, _row := range table.Rows {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", table.GongGetIdentifier(stage))
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Rows")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", _row.GongGetIdentifier(stage))
			pointersInitializesStatements += setPointerField
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}

	// temporary kludge to reuse existing template code
	res = initializerStatements + pointersInitializesStatements
	return
}
