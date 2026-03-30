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
	alternative_idOrdered := []*ALTERNATIVE_ID{}
	for alternative_id := range stage.ALTERNATIVE_IDs {
		alternative_idOrdered = append(alternative_idOrdered, alternative_id)
	}
	sort.Slice(alternative_idOrdered[:], func(i, j int) bool {
		alternative_idi := alternative_idOrdered[i]
		alternative_idj := alternative_idOrdered[j]
		alternative_idi_order, oki := stage.ALTERNATIVE_IDMap_Staged_Order[alternative_idi]
		alternative_idj_order, okj := stage.ALTERNATIVE_IDMap_Staged_Order[alternative_idj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return alternative_idi_order < alternative_idj_order
	})
	if len(alternative_idOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, alternative_id := range alternative_idOrdered {

		identifiersDecl.WriteString(alternative_id.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(alternative_id.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(alternative_id.GongMarshallField(stage, "IDENTIFIER"))
	}

	attribute_definition_booleanOrdered := []*ATTRIBUTE_DEFINITION_BOOLEAN{}
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		attribute_definition_booleanOrdered = append(attribute_definition_booleanOrdered, attribute_definition_boolean)
	}
	sort.Slice(attribute_definition_booleanOrdered[:], func(i, j int) bool {
		attribute_definition_booleani := attribute_definition_booleanOrdered[i]
		attribute_definition_booleanj := attribute_definition_booleanOrdered[j]
		attribute_definition_booleani_order, oki := stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_booleani]
		attribute_definition_booleanj_order, okj := stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_booleani_order < attribute_definition_booleanj_order
	})
	if len(attribute_definition_booleanOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_definition_boolean := range attribute_definition_booleanOrdered {

		identifiersDecl.WriteString(attribute_definition_boolean.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "TYPE"))
	}

	attribute_definition_dateOrdered := []*ATTRIBUTE_DEFINITION_DATE{}
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		attribute_definition_dateOrdered = append(attribute_definition_dateOrdered, attribute_definition_date)
	}
	sort.Slice(attribute_definition_dateOrdered[:], func(i, j int) bool {
		attribute_definition_datei := attribute_definition_dateOrdered[i]
		attribute_definition_datej := attribute_definition_dateOrdered[j]
		attribute_definition_datei_order, oki := stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_datei]
		attribute_definition_datej_order, okj := stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_datei_order < attribute_definition_datej_order
	})
	if len(attribute_definition_dateOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_definition_date := range attribute_definition_dateOrdered {

		identifiersDecl.WriteString(attribute_definition_date.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "TYPE"))
	}

	attribute_definition_enumerationOrdered := []*ATTRIBUTE_DEFINITION_ENUMERATION{}
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		attribute_definition_enumerationOrdered = append(attribute_definition_enumerationOrdered, attribute_definition_enumeration)
	}
	sort.Slice(attribute_definition_enumerationOrdered[:], func(i, j int) bool {
		attribute_definition_enumerationi := attribute_definition_enumerationOrdered[i]
		attribute_definition_enumerationj := attribute_definition_enumerationOrdered[j]
		attribute_definition_enumerationi_order, oki := stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumerationi]
		attribute_definition_enumerationj_order, okj := stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_enumerationi_order < attribute_definition_enumerationj_order
	})
	if len(attribute_definition_enumerationOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {

		identifiersDecl.WriteString(attribute_definition_enumeration.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "MULTI_VALUED"))
		pointersInitializesStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "TYPE"))
	}

	attribute_definition_integerOrdered := []*ATTRIBUTE_DEFINITION_INTEGER{}
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		attribute_definition_integerOrdered = append(attribute_definition_integerOrdered, attribute_definition_integer)
	}
	sort.Slice(attribute_definition_integerOrdered[:], func(i, j int) bool {
		attribute_definition_integeri := attribute_definition_integerOrdered[i]
		attribute_definition_integerj := attribute_definition_integerOrdered[j]
		attribute_definition_integeri_order, oki := stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integeri]
		attribute_definition_integerj_order, okj := stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_integeri_order < attribute_definition_integerj_order
	})
	if len(attribute_definition_integerOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_definition_integer := range attribute_definition_integerOrdered {

		identifiersDecl.WriteString(attribute_definition_integer.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "TYPE"))
	}

	attribute_definition_realOrdered := []*ATTRIBUTE_DEFINITION_REAL{}
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		attribute_definition_realOrdered = append(attribute_definition_realOrdered, attribute_definition_real)
	}
	sort.Slice(attribute_definition_realOrdered[:], func(i, j int) bool {
		attribute_definition_reali := attribute_definition_realOrdered[i]
		attribute_definition_realj := attribute_definition_realOrdered[j]
		attribute_definition_reali_order, oki := stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_reali]
		attribute_definition_realj_order, okj := stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_reali_order < attribute_definition_realj_order
	})
	if len(attribute_definition_realOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_definition_real := range attribute_definition_realOrdered {

		identifiersDecl.WriteString(attribute_definition_real.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "TYPE"))
	}

	attribute_definition_stringOrdered := []*ATTRIBUTE_DEFINITION_STRING{}
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		attribute_definition_stringOrdered = append(attribute_definition_stringOrdered, attribute_definition_string)
	}
	sort.Slice(attribute_definition_stringOrdered[:], func(i, j int) bool {
		attribute_definition_stringi := attribute_definition_stringOrdered[i]
		attribute_definition_stringj := attribute_definition_stringOrdered[j]
		attribute_definition_stringi_order, oki := stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_stringi]
		attribute_definition_stringj_order, okj := stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_stringi_order < attribute_definition_stringj_order
	})
	if len(attribute_definition_stringOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_definition_string := range attribute_definition_stringOrdered {

		identifiersDecl.WriteString(attribute_definition_string.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "TYPE"))
	}

	attribute_definition_xhtmlOrdered := []*ATTRIBUTE_DEFINITION_XHTML{}
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		attribute_definition_xhtmlOrdered = append(attribute_definition_xhtmlOrdered, attribute_definition_xhtml)
	}
	sort.Slice(attribute_definition_xhtmlOrdered[:], func(i, j int) bool {
		attribute_definition_xhtmli := attribute_definition_xhtmlOrdered[i]
		attribute_definition_xhtmlj := attribute_definition_xhtmlOrdered[j]
		attribute_definition_xhtmli_order, oki := stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtmli]
		attribute_definition_xhtmlj_order, okj := stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_definition_xhtmli_order < attribute_definition_xhtmlj_order
	})
	if len(attribute_definition_xhtmlOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {

		identifiersDecl.WriteString(attribute_definition_xhtml.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "TYPE"))
	}

	attribute_value_booleanOrdered := []*ATTRIBUTE_VALUE_BOOLEAN{}
	for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		attribute_value_booleanOrdered = append(attribute_value_booleanOrdered, attribute_value_boolean)
	}
	sort.Slice(attribute_value_booleanOrdered[:], func(i, j int) bool {
		attribute_value_booleani := attribute_value_booleanOrdered[i]
		attribute_value_booleanj := attribute_value_booleanOrdered[j]
		attribute_value_booleani_order, oki := stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_booleani]
		attribute_value_booleanj_order, okj := stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_booleani_order < attribute_value_booleanj_order
	})
	if len(attribute_value_booleanOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_value_boolean := range attribute_value_booleanOrdered {

		identifiersDecl.WriteString(attribute_value_boolean.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "DEFINITION"))
	}

	attribute_value_dateOrdered := []*ATTRIBUTE_VALUE_DATE{}
	for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
		attribute_value_dateOrdered = append(attribute_value_dateOrdered, attribute_value_date)
	}
	sort.Slice(attribute_value_dateOrdered[:], func(i, j int) bool {
		attribute_value_datei := attribute_value_dateOrdered[i]
		attribute_value_datej := attribute_value_dateOrdered[j]
		attribute_value_datei_order, oki := stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_datei]
		attribute_value_datej_order, okj := stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_datei_order < attribute_value_datej_order
	})
	if len(attribute_value_dateOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_value_date := range attribute_value_dateOrdered {

		identifiersDecl.WriteString(attribute_value_date.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_date.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_date.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_date.GongMarshallField(stage, "DEFINITION"))
	}

	attribute_value_enumerationOrdered := []*ATTRIBUTE_VALUE_ENUMERATION{}
	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		attribute_value_enumerationOrdered = append(attribute_value_enumerationOrdered, attribute_value_enumeration)
	}
	sort.Slice(attribute_value_enumerationOrdered[:], func(i, j int) bool {
		attribute_value_enumerationi := attribute_value_enumerationOrdered[i]
		attribute_value_enumerationj := attribute_value_enumerationOrdered[j]
		attribute_value_enumerationi_order, oki := stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumerationi]
		attribute_value_enumerationj_order, okj := stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_enumerationi_order < attribute_value_enumerationj_order
	})
	if len(attribute_value_enumerationOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_value_enumeration := range attribute_value_enumerationOrdered {

		identifiersDecl.WriteString(attribute_value_enumeration.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_enumeration.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(attribute_value_enumeration.GongMarshallField(stage, "DEFINITION"))
		pointersInitializesStatements.WriteString(attribute_value_enumeration.GongMarshallField(stage, "VALUES"))
	}

	attribute_value_integerOrdered := []*ATTRIBUTE_VALUE_INTEGER{}
	for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
		attribute_value_integerOrdered = append(attribute_value_integerOrdered, attribute_value_integer)
	}
	sort.Slice(attribute_value_integerOrdered[:], func(i, j int) bool {
		attribute_value_integeri := attribute_value_integerOrdered[i]
		attribute_value_integerj := attribute_value_integerOrdered[j]
		attribute_value_integeri_order, oki := stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integeri]
		attribute_value_integerj_order, okj := stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_integeri_order < attribute_value_integerj_order
	})
	if len(attribute_value_integerOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_value_integer := range attribute_value_integerOrdered {

		identifiersDecl.WriteString(attribute_value_integer.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "DEFINITION"))
	}

	attribute_value_realOrdered := []*ATTRIBUTE_VALUE_REAL{}
	for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
		attribute_value_realOrdered = append(attribute_value_realOrdered, attribute_value_real)
	}
	sort.Slice(attribute_value_realOrdered[:], func(i, j int) bool {
		attribute_value_reali := attribute_value_realOrdered[i]
		attribute_value_realj := attribute_value_realOrdered[j]
		attribute_value_reali_order, oki := stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_reali]
		attribute_value_realj_order, okj := stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_reali_order < attribute_value_realj_order
	})
	if len(attribute_value_realOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_value_real := range attribute_value_realOrdered {

		identifiersDecl.WriteString(attribute_value_real.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_real.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_real.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_real.GongMarshallField(stage, "DEFINITION"))
	}

	attribute_value_stringOrdered := []*ATTRIBUTE_VALUE_STRING{}
	for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
		attribute_value_stringOrdered = append(attribute_value_stringOrdered, attribute_value_string)
	}
	sort.Slice(attribute_value_stringOrdered[:], func(i, j int) bool {
		attribute_value_stringi := attribute_value_stringOrdered[i]
		attribute_value_stringj := attribute_value_stringOrdered[j]
		attribute_value_stringi_order, oki := stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_stringi]
		attribute_value_stringj_order, okj := stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_stringi_order < attribute_value_stringj_order
	})
	if len(attribute_value_stringOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_value_string := range attribute_value_stringOrdered {

		identifiersDecl.WriteString(attribute_value_string.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_string.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_string.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_string.GongMarshallField(stage, "DEFINITION"))
	}

	attribute_value_xhtmlOrdered := []*ATTRIBUTE_VALUE_XHTML{}
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		attribute_value_xhtmlOrdered = append(attribute_value_xhtmlOrdered, attribute_value_xhtml)
	}
	sort.Slice(attribute_value_xhtmlOrdered[:], func(i, j int) bool {
		attribute_value_xhtmli := attribute_value_xhtmlOrdered[i]
		attribute_value_xhtmlj := attribute_value_xhtmlOrdered[j]
		attribute_value_xhtmli_order, oki := stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtmli]
		attribute_value_xhtmlj_order, okj := stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return attribute_value_xhtmli_order < attribute_value_xhtmlj_order
	})
	if len(attribute_value_xhtmlOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, attribute_value_xhtml := range attribute_value_xhtmlOrdered {

		identifiersDecl.WriteString(attribute_value_xhtml.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "IS_SIMPLIFIED"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "THE_ORIGINAL_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "DEFINITION"))
	}

	a_alternative_idOrdered := []*A_ALTERNATIVE_ID{}
	for a_alternative_id := range stage.A_ALTERNATIVE_IDs {
		a_alternative_idOrdered = append(a_alternative_idOrdered, a_alternative_id)
	}
	sort.Slice(a_alternative_idOrdered[:], func(i, j int) bool {
		a_alternative_idi := a_alternative_idOrdered[i]
		a_alternative_idj := a_alternative_idOrdered[j]
		a_alternative_idi_order, oki := stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_idi]
		a_alternative_idj_order, okj := stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_idj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_alternative_idi_order < a_alternative_idj_order
	})
	if len(a_alternative_idOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_alternative_id := range a_alternative_idOrdered {

		identifiersDecl.WriteString(a_alternative_id.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_alternative_id.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_alternative_id.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}

	a_attribute_definition_boolean_refOrdered := []*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{}
	for a_attribute_definition_boolean_ref := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		a_attribute_definition_boolean_refOrdered = append(a_attribute_definition_boolean_refOrdered, a_attribute_definition_boolean_ref)
	}
	sort.Slice(a_attribute_definition_boolean_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_boolean_refi := a_attribute_definition_boolean_refOrdered[i]
		a_attribute_definition_boolean_refj := a_attribute_definition_boolean_refOrdered[j]
		a_attribute_definition_boolean_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_refi]
		a_attribute_definition_boolean_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_boolean_refi_order < a_attribute_definition_boolean_refj_order
	})
	if len(a_attribute_definition_boolean_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_definition_boolean_ref := range a_attribute_definition_boolean_refOrdered {

		identifiersDecl.WriteString(a_attribute_definition_boolean_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_boolean_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_boolean_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN_REF"))
	}

	a_attribute_definition_date_refOrdered := []*A_ATTRIBUTE_DEFINITION_DATE_REF{}
	for a_attribute_definition_date_ref := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		a_attribute_definition_date_refOrdered = append(a_attribute_definition_date_refOrdered, a_attribute_definition_date_ref)
	}
	sort.Slice(a_attribute_definition_date_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_date_refi := a_attribute_definition_date_refOrdered[i]
		a_attribute_definition_date_refj := a_attribute_definition_date_refOrdered[j]
		a_attribute_definition_date_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_refi]
		a_attribute_definition_date_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_date_refi_order < a_attribute_definition_date_refj_order
	})
	if len(a_attribute_definition_date_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_definition_date_ref := range a_attribute_definition_date_refOrdered {

		identifiersDecl.WriteString(a_attribute_definition_date_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_date_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_date_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE_REF"))
	}

	a_attribute_definition_enumeration_refOrdered := []*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{}
	for a_attribute_definition_enumeration_ref := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		a_attribute_definition_enumeration_refOrdered = append(a_attribute_definition_enumeration_refOrdered, a_attribute_definition_enumeration_ref)
	}
	sort.Slice(a_attribute_definition_enumeration_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_enumeration_refi := a_attribute_definition_enumeration_refOrdered[i]
		a_attribute_definition_enumeration_refj := a_attribute_definition_enumeration_refOrdered[j]
		a_attribute_definition_enumeration_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_refi]
		a_attribute_definition_enumeration_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_enumeration_refi_order < a_attribute_definition_enumeration_refj_order
	})
	if len(a_attribute_definition_enumeration_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_definition_enumeration_ref := range a_attribute_definition_enumeration_refOrdered {

		identifiersDecl.WriteString(a_attribute_definition_enumeration_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_enumeration_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_enumeration_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION_REF"))
	}

	a_attribute_definition_integer_refOrdered := []*A_ATTRIBUTE_DEFINITION_INTEGER_REF{}
	for a_attribute_definition_integer_ref := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		a_attribute_definition_integer_refOrdered = append(a_attribute_definition_integer_refOrdered, a_attribute_definition_integer_ref)
	}
	sort.Slice(a_attribute_definition_integer_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_integer_refi := a_attribute_definition_integer_refOrdered[i]
		a_attribute_definition_integer_refj := a_attribute_definition_integer_refOrdered[j]
		a_attribute_definition_integer_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_refi]
		a_attribute_definition_integer_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_integer_refi_order < a_attribute_definition_integer_refj_order
	})
	if len(a_attribute_definition_integer_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_definition_integer_ref := range a_attribute_definition_integer_refOrdered {

		identifiersDecl.WriteString(a_attribute_definition_integer_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_integer_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_integer_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER_REF"))
	}

	a_attribute_definition_real_refOrdered := []*A_ATTRIBUTE_DEFINITION_REAL_REF{}
	for a_attribute_definition_real_ref := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		a_attribute_definition_real_refOrdered = append(a_attribute_definition_real_refOrdered, a_attribute_definition_real_ref)
	}
	sort.Slice(a_attribute_definition_real_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_real_refi := a_attribute_definition_real_refOrdered[i]
		a_attribute_definition_real_refj := a_attribute_definition_real_refOrdered[j]
		a_attribute_definition_real_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_refi]
		a_attribute_definition_real_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_real_refi_order < a_attribute_definition_real_refj_order
	})
	if len(a_attribute_definition_real_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_definition_real_ref := range a_attribute_definition_real_refOrdered {

		identifiersDecl.WriteString(a_attribute_definition_real_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_real_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_real_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL_REF"))
	}

	a_attribute_definition_string_refOrdered := []*A_ATTRIBUTE_DEFINITION_STRING_REF{}
	for a_attribute_definition_string_ref := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		a_attribute_definition_string_refOrdered = append(a_attribute_definition_string_refOrdered, a_attribute_definition_string_ref)
	}
	sort.Slice(a_attribute_definition_string_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_string_refi := a_attribute_definition_string_refOrdered[i]
		a_attribute_definition_string_refj := a_attribute_definition_string_refOrdered[j]
		a_attribute_definition_string_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_refi]
		a_attribute_definition_string_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_string_refi_order < a_attribute_definition_string_refj_order
	})
	if len(a_attribute_definition_string_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_definition_string_ref := range a_attribute_definition_string_refOrdered {

		identifiersDecl.WriteString(a_attribute_definition_string_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_string_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_string_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING_REF"))
	}

	a_attribute_definition_xhtml_refOrdered := []*A_ATTRIBUTE_DEFINITION_XHTML_REF{}
	for a_attribute_definition_xhtml_ref := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		a_attribute_definition_xhtml_refOrdered = append(a_attribute_definition_xhtml_refOrdered, a_attribute_definition_xhtml_ref)
	}
	sort.Slice(a_attribute_definition_xhtml_refOrdered[:], func(i, j int) bool {
		a_attribute_definition_xhtml_refi := a_attribute_definition_xhtml_refOrdered[i]
		a_attribute_definition_xhtml_refj := a_attribute_definition_xhtml_refOrdered[j]
		a_attribute_definition_xhtml_refi_order, oki := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_refi]
		a_attribute_definition_xhtml_refj_order, okj := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_definition_xhtml_refi_order < a_attribute_definition_xhtml_refj_order
	})
	if len(a_attribute_definition_xhtml_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_definition_xhtml_ref := range a_attribute_definition_xhtml_refOrdered {

		identifiersDecl.WriteString(a_attribute_definition_xhtml_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_xhtml_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_xhtml_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML_REF"))
	}

	a_attribute_value_booleanOrdered := []*A_ATTRIBUTE_VALUE_BOOLEAN{}
	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		a_attribute_value_booleanOrdered = append(a_attribute_value_booleanOrdered, a_attribute_value_boolean)
	}
	sort.Slice(a_attribute_value_booleanOrdered[:], func(i, j int) bool {
		a_attribute_value_booleani := a_attribute_value_booleanOrdered[i]
		a_attribute_value_booleanj := a_attribute_value_booleanOrdered[j]
		a_attribute_value_booleani_order, oki := stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_booleani]
		a_attribute_value_booleanj_order, okj := stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_booleani_order < a_attribute_value_booleanj_order
	})
	if len(a_attribute_value_booleanOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_boolean := range a_attribute_value_booleanOrdered {

		identifiersDecl.WriteString(a_attribute_value_boolean.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_boolean.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_boolean.GongMarshallField(stage, "ATTRIBUTE_VALUE_BOOLEAN"))
	}

	a_attribute_value_dateOrdered := []*A_ATTRIBUTE_VALUE_DATE{}
	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		a_attribute_value_dateOrdered = append(a_attribute_value_dateOrdered, a_attribute_value_date)
	}
	sort.Slice(a_attribute_value_dateOrdered[:], func(i, j int) bool {
		a_attribute_value_datei := a_attribute_value_dateOrdered[i]
		a_attribute_value_datej := a_attribute_value_dateOrdered[j]
		a_attribute_value_datei_order, oki := stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_datei]
		a_attribute_value_datej_order, okj := stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_datei_order < a_attribute_value_datej_order
	})
	if len(a_attribute_value_dateOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_date := range a_attribute_value_dateOrdered {

		identifiersDecl.WriteString(a_attribute_value_date.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_date.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_date.GongMarshallField(stage, "ATTRIBUTE_VALUE_DATE"))
	}

	a_attribute_value_enumerationOrdered := []*A_ATTRIBUTE_VALUE_ENUMERATION{}
	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		a_attribute_value_enumerationOrdered = append(a_attribute_value_enumerationOrdered, a_attribute_value_enumeration)
	}
	sort.Slice(a_attribute_value_enumerationOrdered[:], func(i, j int) bool {
		a_attribute_value_enumerationi := a_attribute_value_enumerationOrdered[i]
		a_attribute_value_enumerationj := a_attribute_value_enumerationOrdered[j]
		a_attribute_value_enumerationi_order, oki := stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumerationi]
		a_attribute_value_enumerationj_order, okj := stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_enumerationi_order < a_attribute_value_enumerationj_order
	})
	if len(a_attribute_value_enumerationOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_enumeration := range a_attribute_value_enumerationOrdered {

		identifiersDecl.WriteString(a_attribute_value_enumeration.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_enumeration.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_enumeration.GongMarshallField(stage, "ATTRIBUTE_VALUE_ENUMERATION"))
	}

	a_attribute_value_integerOrdered := []*A_ATTRIBUTE_VALUE_INTEGER{}
	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		a_attribute_value_integerOrdered = append(a_attribute_value_integerOrdered, a_attribute_value_integer)
	}
	sort.Slice(a_attribute_value_integerOrdered[:], func(i, j int) bool {
		a_attribute_value_integeri := a_attribute_value_integerOrdered[i]
		a_attribute_value_integerj := a_attribute_value_integerOrdered[j]
		a_attribute_value_integeri_order, oki := stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integeri]
		a_attribute_value_integerj_order, okj := stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_integeri_order < a_attribute_value_integerj_order
	})
	if len(a_attribute_value_integerOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_integer := range a_attribute_value_integerOrdered {

		identifiersDecl.WriteString(a_attribute_value_integer.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_integer.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_integer.GongMarshallField(stage, "ATTRIBUTE_VALUE_INTEGER"))
	}

	a_attribute_value_realOrdered := []*A_ATTRIBUTE_VALUE_REAL{}
	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		a_attribute_value_realOrdered = append(a_attribute_value_realOrdered, a_attribute_value_real)
	}
	sort.Slice(a_attribute_value_realOrdered[:], func(i, j int) bool {
		a_attribute_value_reali := a_attribute_value_realOrdered[i]
		a_attribute_value_realj := a_attribute_value_realOrdered[j]
		a_attribute_value_reali_order, oki := stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_reali]
		a_attribute_value_realj_order, okj := stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_reali_order < a_attribute_value_realj_order
	})
	if len(a_attribute_value_realOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_real := range a_attribute_value_realOrdered {

		identifiersDecl.WriteString(a_attribute_value_real.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_real.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_real.GongMarshallField(stage, "ATTRIBUTE_VALUE_REAL"))
	}

	a_attribute_value_stringOrdered := []*A_ATTRIBUTE_VALUE_STRING{}
	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		a_attribute_value_stringOrdered = append(a_attribute_value_stringOrdered, a_attribute_value_string)
	}
	sort.Slice(a_attribute_value_stringOrdered[:], func(i, j int) bool {
		a_attribute_value_stringi := a_attribute_value_stringOrdered[i]
		a_attribute_value_stringj := a_attribute_value_stringOrdered[j]
		a_attribute_value_stringi_order, oki := stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_stringi]
		a_attribute_value_stringj_order, okj := stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_stringi_order < a_attribute_value_stringj_order
	})
	if len(a_attribute_value_stringOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_string := range a_attribute_value_stringOrdered {

		identifiersDecl.WriteString(a_attribute_value_string.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_string.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_string.GongMarshallField(stage, "ATTRIBUTE_VALUE_STRING"))
	}

	a_attribute_value_xhtmlOrdered := []*A_ATTRIBUTE_VALUE_XHTML{}
	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		a_attribute_value_xhtmlOrdered = append(a_attribute_value_xhtmlOrdered, a_attribute_value_xhtml)
	}
	sort.Slice(a_attribute_value_xhtmlOrdered[:], func(i, j int) bool {
		a_attribute_value_xhtmli := a_attribute_value_xhtmlOrdered[i]
		a_attribute_value_xhtmlj := a_attribute_value_xhtmlOrdered[j]
		a_attribute_value_xhtmli_order, oki := stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtmli]
		a_attribute_value_xhtmlj_order, okj := stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_xhtmli_order < a_attribute_value_xhtmlj_order
	})
	if len(a_attribute_value_xhtmlOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_xhtml := range a_attribute_value_xhtmlOrdered {

		identifiersDecl.WriteString(a_attribute_value_xhtml.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_xhtml.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml.GongMarshallField(stage, "ATTRIBUTE_VALUE_XHTML"))
	}

	a_attribute_value_xhtml_1Ordered := []*A_ATTRIBUTE_VALUE_XHTML_1{}
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		a_attribute_value_xhtml_1Ordered = append(a_attribute_value_xhtml_1Ordered, a_attribute_value_xhtml_1)
	}
	sort.Slice(a_attribute_value_xhtml_1Ordered[:], func(i, j int) bool {
		a_attribute_value_xhtml_1i := a_attribute_value_xhtml_1Ordered[i]
		a_attribute_value_xhtml_1j := a_attribute_value_xhtml_1Ordered[j]
		a_attribute_value_xhtml_1i_order, oki := stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1i]
		a_attribute_value_xhtml_1j_order, okj := stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_attribute_value_xhtml_1i_order < a_attribute_value_xhtml_1j_order
	})
	if len(a_attribute_value_xhtml_1Ordered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_attribute_value_xhtml_1 := range a_attribute_value_xhtml_1Ordered {

		identifiersDecl.WriteString(a_attribute_value_xhtml_1.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_BOOLEAN"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_DATE"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_ENUMERATION"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_INTEGER"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_REAL"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_STRING"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_XHTML"))
	}

	a_childrenOrdered := []*A_CHILDREN{}
	for a_children := range stage.A_CHILDRENs {
		a_childrenOrdered = append(a_childrenOrdered, a_children)
	}
	sort.Slice(a_childrenOrdered[:], func(i, j int) bool {
		a_childreni := a_childrenOrdered[i]
		a_childrenj := a_childrenOrdered[j]
		a_childreni_order, oki := stage.A_CHILDRENMap_Staged_Order[a_childreni]
		a_childrenj_order, okj := stage.A_CHILDRENMap_Staged_Order[a_childrenj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_childreni_order < a_childrenj_order
	})
	if len(a_childrenOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_children := range a_childrenOrdered {

		identifiersDecl.WriteString(a_children.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_children.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_children.GongMarshallField(stage, "SPEC_HIERARCHY"))
	}

	a_core_contentOrdered := []*A_CORE_CONTENT{}
	for a_core_content := range stage.A_CORE_CONTENTs {
		a_core_contentOrdered = append(a_core_contentOrdered, a_core_content)
	}
	sort.Slice(a_core_contentOrdered[:], func(i, j int) bool {
		a_core_contenti := a_core_contentOrdered[i]
		a_core_contentj := a_core_contentOrdered[j]
		a_core_contenti_order, oki := stage.A_CORE_CONTENTMap_Staged_Order[a_core_contenti]
		a_core_contentj_order, okj := stage.A_CORE_CONTENTMap_Staged_Order[a_core_contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_core_contenti_order < a_core_contentj_order
	})
	if len(a_core_contentOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_core_content := range a_core_contentOrdered {

		identifiersDecl.WriteString(a_core_content.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_core_content.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_core_content.GongMarshallField(stage, "REQ_IF_CONTENT"))
	}

	a_datatypesOrdered := []*A_DATATYPES{}
	for a_datatypes := range stage.A_DATATYPESs {
		a_datatypesOrdered = append(a_datatypesOrdered, a_datatypes)
	}
	sort.Slice(a_datatypesOrdered[:], func(i, j int) bool {
		a_datatypesi := a_datatypesOrdered[i]
		a_datatypesj := a_datatypesOrdered[j]
		a_datatypesi_order, oki := stage.A_DATATYPESMap_Staged_Order[a_datatypesi]
		a_datatypesj_order, okj := stage.A_DATATYPESMap_Staged_Order[a_datatypesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatypesi_order < a_datatypesj_order
	})
	if len(a_datatypesOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatypes := range a_datatypesOrdered {

		identifiersDecl.WriteString(a_datatypes.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatypes.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_BOOLEAN"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_DATE"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_ENUMERATION"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_INTEGER"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_REAL"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_STRING"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_XHTML"))
	}

	a_datatype_definition_boolean_refOrdered := []*A_DATATYPE_DEFINITION_BOOLEAN_REF{}
	for a_datatype_definition_boolean_ref := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		a_datatype_definition_boolean_refOrdered = append(a_datatype_definition_boolean_refOrdered, a_datatype_definition_boolean_ref)
	}
	sort.Slice(a_datatype_definition_boolean_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_boolean_refi := a_datatype_definition_boolean_refOrdered[i]
		a_datatype_definition_boolean_refj := a_datatype_definition_boolean_refOrdered[j]
		a_datatype_definition_boolean_refi_order, oki := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_refi]
		a_datatype_definition_boolean_refj_order, okj := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_boolean_refi_order < a_datatype_definition_boolean_refj_order
	})
	if len(a_datatype_definition_boolean_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatype_definition_boolean_ref := range a_datatype_definition_boolean_refOrdered {

		identifiersDecl.WriteString(a_datatype_definition_boolean_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_boolean_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_boolean_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_BOOLEAN_REF"))
	}

	a_datatype_definition_date_refOrdered := []*A_DATATYPE_DEFINITION_DATE_REF{}
	for a_datatype_definition_date_ref := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		a_datatype_definition_date_refOrdered = append(a_datatype_definition_date_refOrdered, a_datatype_definition_date_ref)
	}
	sort.Slice(a_datatype_definition_date_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_date_refi := a_datatype_definition_date_refOrdered[i]
		a_datatype_definition_date_refj := a_datatype_definition_date_refOrdered[j]
		a_datatype_definition_date_refi_order, oki := stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_refi]
		a_datatype_definition_date_refj_order, okj := stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_date_refi_order < a_datatype_definition_date_refj_order
	})
	if len(a_datatype_definition_date_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatype_definition_date_ref := range a_datatype_definition_date_refOrdered {

		identifiersDecl.WriteString(a_datatype_definition_date_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_date_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_date_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_DATE_REF"))
	}

	a_datatype_definition_enumeration_refOrdered := []*A_DATATYPE_DEFINITION_ENUMERATION_REF{}
	for a_datatype_definition_enumeration_ref := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		a_datatype_definition_enumeration_refOrdered = append(a_datatype_definition_enumeration_refOrdered, a_datatype_definition_enumeration_ref)
	}
	sort.Slice(a_datatype_definition_enumeration_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_enumeration_refi := a_datatype_definition_enumeration_refOrdered[i]
		a_datatype_definition_enumeration_refj := a_datatype_definition_enumeration_refOrdered[j]
		a_datatype_definition_enumeration_refi_order, oki := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_refi]
		a_datatype_definition_enumeration_refj_order, okj := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_enumeration_refi_order < a_datatype_definition_enumeration_refj_order
	})
	if len(a_datatype_definition_enumeration_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatype_definition_enumeration_ref := range a_datatype_definition_enumeration_refOrdered {

		identifiersDecl.WriteString(a_datatype_definition_enumeration_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_enumeration_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_enumeration_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_ENUMERATION_REF"))
	}

	a_datatype_definition_integer_refOrdered := []*A_DATATYPE_DEFINITION_INTEGER_REF{}
	for a_datatype_definition_integer_ref := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		a_datatype_definition_integer_refOrdered = append(a_datatype_definition_integer_refOrdered, a_datatype_definition_integer_ref)
	}
	sort.Slice(a_datatype_definition_integer_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_integer_refi := a_datatype_definition_integer_refOrdered[i]
		a_datatype_definition_integer_refj := a_datatype_definition_integer_refOrdered[j]
		a_datatype_definition_integer_refi_order, oki := stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_refi]
		a_datatype_definition_integer_refj_order, okj := stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_integer_refi_order < a_datatype_definition_integer_refj_order
	})
	if len(a_datatype_definition_integer_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatype_definition_integer_ref := range a_datatype_definition_integer_refOrdered {

		identifiersDecl.WriteString(a_datatype_definition_integer_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_integer_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_integer_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_INTEGER_REF"))
	}

	a_datatype_definition_real_refOrdered := []*A_DATATYPE_DEFINITION_REAL_REF{}
	for a_datatype_definition_real_ref := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		a_datatype_definition_real_refOrdered = append(a_datatype_definition_real_refOrdered, a_datatype_definition_real_ref)
	}
	sort.Slice(a_datatype_definition_real_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_real_refi := a_datatype_definition_real_refOrdered[i]
		a_datatype_definition_real_refj := a_datatype_definition_real_refOrdered[j]
		a_datatype_definition_real_refi_order, oki := stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_refi]
		a_datatype_definition_real_refj_order, okj := stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_real_refi_order < a_datatype_definition_real_refj_order
	})
	if len(a_datatype_definition_real_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatype_definition_real_ref := range a_datatype_definition_real_refOrdered {

		identifiersDecl.WriteString(a_datatype_definition_real_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_real_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_real_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_REAL_REF"))
	}

	a_datatype_definition_string_refOrdered := []*A_DATATYPE_DEFINITION_STRING_REF{}
	for a_datatype_definition_string_ref := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		a_datatype_definition_string_refOrdered = append(a_datatype_definition_string_refOrdered, a_datatype_definition_string_ref)
	}
	sort.Slice(a_datatype_definition_string_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_string_refi := a_datatype_definition_string_refOrdered[i]
		a_datatype_definition_string_refj := a_datatype_definition_string_refOrdered[j]
		a_datatype_definition_string_refi_order, oki := stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_refi]
		a_datatype_definition_string_refj_order, okj := stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_string_refi_order < a_datatype_definition_string_refj_order
	})
	if len(a_datatype_definition_string_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatype_definition_string_ref := range a_datatype_definition_string_refOrdered {

		identifiersDecl.WriteString(a_datatype_definition_string_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_string_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_string_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_STRING_REF"))
	}

	a_datatype_definition_xhtml_refOrdered := []*A_DATATYPE_DEFINITION_XHTML_REF{}
	for a_datatype_definition_xhtml_ref := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		a_datatype_definition_xhtml_refOrdered = append(a_datatype_definition_xhtml_refOrdered, a_datatype_definition_xhtml_ref)
	}
	sort.Slice(a_datatype_definition_xhtml_refOrdered[:], func(i, j int) bool {
		a_datatype_definition_xhtml_refi := a_datatype_definition_xhtml_refOrdered[i]
		a_datatype_definition_xhtml_refj := a_datatype_definition_xhtml_refOrdered[j]
		a_datatype_definition_xhtml_refi_order, oki := stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_refi]
		a_datatype_definition_xhtml_refj_order, okj := stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_datatype_definition_xhtml_refi_order < a_datatype_definition_xhtml_refj_order
	})
	if len(a_datatype_definition_xhtml_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_datatype_definition_xhtml_ref := range a_datatype_definition_xhtml_refOrdered {

		identifiersDecl.WriteString(a_datatype_definition_xhtml_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_xhtml_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_xhtml_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_XHTML_REF"))
	}

	a_editable_attsOrdered := []*A_EDITABLE_ATTS{}
	for a_editable_atts := range stage.A_EDITABLE_ATTSs {
		a_editable_attsOrdered = append(a_editable_attsOrdered, a_editable_atts)
	}
	sort.Slice(a_editable_attsOrdered[:], func(i, j int) bool {
		a_editable_attsi := a_editable_attsOrdered[i]
		a_editable_attsj := a_editable_attsOrdered[j]
		a_editable_attsi_order, oki := stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_attsi]
		a_editable_attsj_order, okj := stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_attsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_editable_attsi_order < a_editable_attsj_order
	})
	if len(a_editable_attsOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_editable_atts := range a_editable_attsOrdered {

		identifiersDecl.WriteString(a_editable_atts.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML_REF"))
	}

	a_enum_value_refOrdered := []*A_ENUM_VALUE_REF{}
	for a_enum_value_ref := range stage.A_ENUM_VALUE_REFs {
		a_enum_value_refOrdered = append(a_enum_value_refOrdered, a_enum_value_ref)
	}
	sort.Slice(a_enum_value_refOrdered[:], func(i, j int) bool {
		a_enum_value_refi := a_enum_value_refOrdered[i]
		a_enum_value_refj := a_enum_value_refOrdered[j]
		a_enum_value_refi_order, oki := stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_refi]
		a_enum_value_refj_order, okj := stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_enum_value_refi_order < a_enum_value_refj_order
	})
	if len(a_enum_value_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_enum_value_ref := range a_enum_value_refOrdered {

		identifiersDecl.WriteString(a_enum_value_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_enum_value_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_enum_value_ref.GongMarshallField(stage, "ENUM_VALUE_REF"))
	}

	a_objectOrdered := []*A_OBJECT{}
	for a_object := range stage.A_OBJECTs {
		a_objectOrdered = append(a_objectOrdered, a_object)
	}
	sort.Slice(a_objectOrdered[:], func(i, j int) bool {
		a_objecti := a_objectOrdered[i]
		a_objectj := a_objectOrdered[j]
		a_objecti_order, oki := stage.A_OBJECTMap_Staged_Order[a_objecti]
		a_objectj_order, okj := stage.A_OBJECTMap_Staged_Order[a_objectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_objecti_order < a_objectj_order
	})
	if len(a_objectOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_object := range a_objectOrdered {

		identifiersDecl.WriteString(a_object.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_object.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_object.GongMarshallField(stage, "SPEC_OBJECT_REF"))
	}

	a_propertiesOrdered := []*A_PROPERTIES{}
	for a_properties := range stage.A_PROPERTIESs {
		a_propertiesOrdered = append(a_propertiesOrdered, a_properties)
	}
	sort.Slice(a_propertiesOrdered[:], func(i, j int) bool {
		a_propertiesi := a_propertiesOrdered[i]
		a_propertiesj := a_propertiesOrdered[j]
		a_propertiesi_order, oki := stage.A_PROPERTIESMap_Staged_Order[a_propertiesi]
		a_propertiesj_order, okj := stage.A_PROPERTIESMap_Staged_Order[a_propertiesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_propertiesi_order < a_propertiesj_order
	})
	if len(a_propertiesOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_properties := range a_propertiesOrdered {

		identifiersDecl.WriteString(a_properties.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_properties.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_properties.GongMarshallField(stage, "EMBEDDED_VALUE"))
	}

	a_relation_group_type_refOrdered := []*A_RELATION_GROUP_TYPE_REF{}
	for a_relation_group_type_ref := range stage.A_RELATION_GROUP_TYPE_REFs {
		a_relation_group_type_refOrdered = append(a_relation_group_type_refOrdered, a_relation_group_type_ref)
	}
	sort.Slice(a_relation_group_type_refOrdered[:], func(i, j int) bool {
		a_relation_group_type_refi := a_relation_group_type_refOrdered[i]
		a_relation_group_type_refj := a_relation_group_type_refOrdered[j]
		a_relation_group_type_refi_order, oki := stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_refi]
		a_relation_group_type_refj_order, okj := stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_relation_group_type_refi_order < a_relation_group_type_refj_order
	})
	if len(a_relation_group_type_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_relation_group_type_ref := range a_relation_group_type_refOrdered {

		identifiersDecl.WriteString(a_relation_group_type_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_relation_group_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_relation_group_type_ref.GongMarshallField(stage, "RELATION_GROUP_TYPE_REF"))
	}

	a_source_1Ordered := []*A_SOURCE_1{}
	for a_source_1 := range stage.A_SOURCE_1s {
		a_source_1Ordered = append(a_source_1Ordered, a_source_1)
	}
	sort.Slice(a_source_1Ordered[:], func(i, j int) bool {
		a_source_1i := a_source_1Ordered[i]
		a_source_1j := a_source_1Ordered[j]
		a_source_1i_order, oki := stage.A_SOURCE_1Map_Staged_Order[a_source_1i]
		a_source_1j_order, okj := stage.A_SOURCE_1Map_Staged_Order[a_source_1j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_source_1i_order < a_source_1j_order
	})
	if len(a_source_1Ordered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_source_1 := range a_source_1Ordered {

		identifiersDecl.WriteString(a_source_1.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_source_1.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_source_1.GongMarshallField(stage, "SPEC_OBJECT_REF"))
	}

	a_source_specification_1Ordered := []*A_SOURCE_SPECIFICATION_1{}
	for a_source_specification_1 := range stage.A_SOURCE_SPECIFICATION_1s {
		a_source_specification_1Ordered = append(a_source_specification_1Ordered, a_source_specification_1)
	}
	sort.Slice(a_source_specification_1Ordered[:], func(i, j int) bool {
		a_source_specification_1i := a_source_specification_1Ordered[i]
		a_source_specification_1j := a_source_specification_1Ordered[j]
		a_source_specification_1i_order, oki := stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1i]
		a_source_specification_1j_order, okj := stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1j]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_source_specification_1i_order < a_source_specification_1j_order
	})
	if len(a_source_specification_1Ordered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_source_specification_1 := range a_source_specification_1Ordered {

		identifiersDecl.WriteString(a_source_specification_1.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_source_specification_1.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_source_specification_1.GongMarshallField(stage, "SPECIFICATION_REF"))
	}

	a_specificationsOrdered := []*A_SPECIFICATIONS{}
	for a_specifications := range stage.A_SPECIFICATIONSs {
		a_specificationsOrdered = append(a_specificationsOrdered, a_specifications)
	}
	sort.Slice(a_specificationsOrdered[:], func(i, j int) bool {
		a_specificationsi := a_specificationsOrdered[i]
		a_specificationsj := a_specificationsOrdered[j]
		a_specificationsi_order, oki := stage.A_SPECIFICATIONSMap_Staged_Order[a_specificationsi]
		a_specificationsj_order, okj := stage.A_SPECIFICATIONSMap_Staged_Order[a_specificationsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_specificationsi_order < a_specificationsj_order
	})
	if len(a_specificationsOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_specifications := range a_specificationsOrdered {

		identifiersDecl.WriteString(a_specifications.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_specifications.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_specifications.GongMarshallField(stage, "SPECIFICATION"))
	}

	a_specification_type_refOrdered := []*A_SPECIFICATION_TYPE_REF{}
	for a_specification_type_ref := range stage.A_SPECIFICATION_TYPE_REFs {
		a_specification_type_refOrdered = append(a_specification_type_refOrdered, a_specification_type_ref)
	}
	sort.Slice(a_specification_type_refOrdered[:], func(i, j int) bool {
		a_specification_type_refi := a_specification_type_refOrdered[i]
		a_specification_type_refj := a_specification_type_refOrdered[j]
		a_specification_type_refi_order, oki := stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_refi]
		a_specification_type_refj_order, okj := stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_specification_type_refi_order < a_specification_type_refj_order
	})
	if len(a_specification_type_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_specification_type_ref := range a_specification_type_refOrdered {

		identifiersDecl.WriteString(a_specification_type_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_specification_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_specification_type_ref.GongMarshallField(stage, "SPECIFICATION_TYPE_REF"))
	}

	a_specified_valuesOrdered := []*A_SPECIFIED_VALUES{}
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		a_specified_valuesOrdered = append(a_specified_valuesOrdered, a_specified_values)
	}
	sort.Slice(a_specified_valuesOrdered[:], func(i, j int) bool {
		a_specified_valuesi := a_specified_valuesOrdered[i]
		a_specified_valuesj := a_specified_valuesOrdered[j]
		a_specified_valuesi_order, oki := stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_valuesi]
		a_specified_valuesj_order, okj := stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_valuesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_specified_valuesi_order < a_specified_valuesj_order
	})
	if len(a_specified_valuesOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_specified_values := range a_specified_valuesOrdered {

		identifiersDecl.WriteString(a_specified_values.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_specified_values.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_specified_values.GongMarshallField(stage, "ENUM_VALUE"))
	}

	a_spec_attributesOrdered := []*A_SPEC_ATTRIBUTES{}
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		a_spec_attributesOrdered = append(a_spec_attributesOrdered, a_spec_attributes)
	}
	sort.Slice(a_spec_attributesOrdered[:], func(i, j int) bool {
		a_spec_attributesi := a_spec_attributesOrdered[i]
		a_spec_attributesj := a_spec_attributesOrdered[j]
		a_spec_attributesi_order, oki := stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributesi]
		a_spec_attributesj_order, okj := stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_attributesi_order < a_spec_attributesj_order
	})
	if len(a_spec_attributesOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_attributes := range a_spec_attributesOrdered {

		identifiersDecl.WriteString(a_spec_attributes.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML"))
	}

	a_spec_objectsOrdered := []*A_SPEC_OBJECTS{}
	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		a_spec_objectsOrdered = append(a_spec_objectsOrdered, a_spec_objects)
	}
	sort.Slice(a_spec_objectsOrdered[:], func(i, j int) bool {
		a_spec_objectsi := a_spec_objectsOrdered[i]
		a_spec_objectsj := a_spec_objectsOrdered[j]
		a_spec_objectsi_order, oki := stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objectsi]
		a_spec_objectsj_order, okj := stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objectsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_objectsi_order < a_spec_objectsj_order
	})
	if len(a_spec_objectsOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_objects := range a_spec_objectsOrdered {

		identifiersDecl.WriteString(a_spec_objects.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_objects.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_objects.GongMarshallField(stage, "SPEC_OBJECT"))
	}

	a_spec_object_type_refOrdered := []*A_SPEC_OBJECT_TYPE_REF{}
	for a_spec_object_type_ref := range stage.A_SPEC_OBJECT_TYPE_REFs {
		a_spec_object_type_refOrdered = append(a_spec_object_type_refOrdered, a_spec_object_type_ref)
	}
	sort.Slice(a_spec_object_type_refOrdered[:], func(i, j int) bool {
		a_spec_object_type_refi := a_spec_object_type_refOrdered[i]
		a_spec_object_type_refj := a_spec_object_type_refOrdered[j]
		a_spec_object_type_refi_order, oki := stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_refi]
		a_spec_object_type_refj_order, okj := stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_object_type_refi_order < a_spec_object_type_refj_order
	})
	if len(a_spec_object_type_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_object_type_ref := range a_spec_object_type_refOrdered {

		identifiersDecl.WriteString(a_spec_object_type_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_object_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_spec_object_type_ref.GongMarshallField(stage, "SPEC_OBJECT_TYPE_REF"))
	}

	a_spec_relationsOrdered := []*A_SPEC_RELATIONS{}
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		a_spec_relationsOrdered = append(a_spec_relationsOrdered, a_spec_relations)
	}
	sort.Slice(a_spec_relationsOrdered[:], func(i, j int) bool {
		a_spec_relationsi := a_spec_relationsOrdered[i]
		a_spec_relationsj := a_spec_relationsOrdered[j]
		a_spec_relationsi_order, oki := stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relationsi]
		a_spec_relationsj_order, okj := stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relationsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relationsi_order < a_spec_relationsj_order
	})
	if len(a_spec_relationsOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_relations := range a_spec_relationsOrdered {

		identifiersDecl.WriteString(a_spec_relations.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relations.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_relations.GongMarshallField(stage, "SPEC_RELATION"))
	}

	a_spec_relation_groupsOrdered := []*A_SPEC_RELATION_GROUPS{}
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		a_spec_relation_groupsOrdered = append(a_spec_relation_groupsOrdered, a_spec_relation_groups)
	}
	sort.Slice(a_spec_relation_groupsOrdered[:], func(i, j int) bool {
		a_spec_relation_groupsi := a_spec_relation_groupsOrdered[i]
		a_spec_relation_groupsj := a_spec_relation_groupsOrdered[j]
		a_spec_relation_groupsi_order, oki := stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groupsi]
		a_spec_relation_groupsj_order, okj := stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groupsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relation_groupsi_order < a_spec_relation_groupsj_order
	})
	if len(a_spec_relation_groupsOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_relation_groups := range a_spec_relation_groupsOrdered {

		identifiersDecl.WriteString(a_spec_relation_groups.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relation_groups.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_relation_groups.GongMarshallField(stage, "RELATION_GROUP"))
	}

	a_spec_relation_refOrdered := []*A_SPEC_RELATION_REF{}
	for a_spec_relation_ref := range stage.A_SPEC_RELATION_REFs {
		a_spec_relation_refOrdered = append(a_spec_relation_refOrdered, a_spec_relation_ref)
	}
	sort.Slice(a_spec_relation_refOrdered[:], func(i, j int) bool {
		a_spec_relation_refi := a_spec_relation_refOrdered[i]
		a_spec_relation_refj := a_spec_relation_refOrdered[j]
		a_spec_relation_refi_order, oki := stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_refi]
		a_spec_relation_refj_order, okj := stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relation_refi_order < a_spec_relation_refj_order
	})
	if len(a_spec_relation_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_relation_ref := range a_spec_relation_refOrdered {

		identifiersDecl.WriteString(a_spec_relation_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relation_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_spec_relation_ref.GongMarshallField(stage, "SPEC_RELATION_REF"))
	}

	a_spec_relation_type_refOrdered := []*A_SPEC_RELATION_TYPE_REF{}
	for a_spec_relation_type_ref := range stage.A_SPEC_RELATION_TYPE_REFs {
		a_spec_relation_type_refOrdered = append(a_spec_relation_type_refOrdered, a_spec_relation_type_ref)
	}
	sort.Slice(a_spec_relation_type_refOrdered[:], func(i, j int) bool {
		a_spec_relation_type_refi := a_spec_relation_type_refOrdered[i]
		a_spec_relation_type_refj := a_spec_relation_type_refOrdered[j]
		a_spec_relation_type_refi_order, oki := stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_refi]
		a_spec_relation_type_refj_order, okj := stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_refj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_relation_type_refi_order < a_spec_relation_type_refj_order
	})
	if len(a_spec_relation_type_refOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_relation_type_ref := range a_spec_relation_type_refOrdered {

		identifiersDecl.WriteString(a_spec_relation_type_ref.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relation_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_spec_relation_type_ref.GongMarshallField(stage, "SPEC_RELATION_TYPE_REF"))
	}

	a_spec_typesOrdered := []*A_SPEC_TYPES{}
	for a_spec_types := range stage.A_SPEC_TYPESs {
		a_spec_typesOrdered = append(a_spec_typesOrdered, a_spec_types)
	}
	sort.Slice(a_spec_typesOrdered[:], func(i, j int) bool {
		a_spec_typesi := a_spec_typesOrdered[i]
		a_spec_typesj := a_spec_typesOrdered[j]
		a_spec_typesi_order, oki := stage.A_SPEC_TYPESMap_Staged_Order[a_spec_typesi]
		a_spec_typesj_order, okj := stage.A_SPEC_TYPESMap_Staged_Order[a_spec_typesj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_spec_typesi_order < a_spec_typesj_order
	})
	if len(a_spec_typesOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_spec_types := range a_spec_typesOrdered {

		identifiersDecl.WriteString(a_spec_types.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_types.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "RELATION_GROUP_TYPE"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "SPEC_OBJECT_TYPE"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "SPEC_RELATION_TYPE"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "SPECIFICATION_TYPE"))
	}

	a_the_headerOrdered := []*A_THE_HEADER{}
	for a_the_header := range stage.A_THE_HEADERs {
		a_the_headerOrdered = append(a_the_headerOrdered, a_the_header)
	}
	sort.Slice(a_the_headerOrdered[:], func(i, j int) bool {
		a_the_headeri := a_the_headerOrdered[i]
		a_the_headerj := a_the_headerOrdered[j]
		a_the_headeri_order, oki := stage.A_THE_HEADERMap_Staged_Order[a_the_headeri]
		a_the_headerj_order, okj := stage.A_THE_HEADERMap_Staged_Order[a_the_headerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_the_headeri_order < a_the_headerj_order
	})
	if len(a_the_headerOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_the_header := range a_the_headerOrdered {

		identifiersDecl.WriteString(a_the_header.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_the_header.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_the_header.GongMarshallField(stage, "REQ_IF_HEADER"))
	}

	a_tool_extensionsOrdered := []*A_TOOL_EXTENSIONS{}
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		a_tool_extensionsOrdered = append(a_tool_extensionsOrdered, a_tool_extensions)
	}
	sort.Slice(a_tool_extensionsOrdered[:], func(i, j int) bool {
		a_tool_extensionsi := a_tool_extensionsOrdered[i]
		a_tool_extensionsj := a_tool_extensionsOrdered[j]
		a_tool_extensionsi_order, oki := stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensionsi]
		a_tool_extensionsj_order, okj := stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensionsj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return a_tool_extensionsi_order < a_tool_extensionsj_order
	})
	if len(a_tool_extensionsOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, a_tool_extensions := range a_tool_extensionsOrdered {

		identifiersDecl.WriteString(a_tool_extensions.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_tool_extensions.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_tool_extensions.GongMarshallField(stage, "REQ_IF_TOOL_EXTENSION"))
	}

	datatype_definition_booleanOrdered := []*DATATYPE_DEFINITION_BOOLEAN{}
	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		datatype_definition_booleanOrdered = append(datatype_definition_booleanOrdered, datatype_definition_boolean)
	}
	sort.Slice(datatype_definition_booleanOrdered[:], func(i, j int) bool {
		datatype_definition_booleani := datatype_definition_booleanOrdered[i]
		datatype_definition_booleanj := datatype_definition_booleanOrdered[j]
		datatype_definition_booleani_order, oki := stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_booleani]
		datatype_definition_booleanj_order, okj := stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_booleanj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_booleani_order < datatype_definition_booleanj_order
	})
	if len(datatype_definition_booleanOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datatype_definition_boolean := range datatype_definition_booleanOrdered {

		identifiersDecl.WriteString(datatype_definition_boolean.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}

	datatype_definition_dateOrdered := []*DATATYPE_DEFINITION_DATE{}
	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		datatype_definition_dateOrdered = append(datatype_definition_dateOrdered, datatype_definition_date)
	}
	sort.Slice(datatype_definition_dateOrdered[:], func(i, j int) bool {
		datatype_definition_datei := datatype_definition_dateOrdered[i]
		datatype_definition_datej := datatype_definition_dateOrdered[j]
		datatype_definition_datei_order, oki := stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_datei]
		datatype_definition_datej_order, okj := stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_datej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_datei_order < datatype_definition_datej_order
	})
	if len(datatype_definition_dateOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datatype_definition_date := range datatype_definition_dateOrdered {

		identifiersDecl.WriteString(datatype_definition_date.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}

	datatype_definition_enumerationOrdered := []*DATATYPE_DEFINITION_ENUMERATION{}
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		datatype_definition_enumerationOrdered = append(datatype_definition_enumerationOrdered, datatype_definition_enumeration)
	}
	sort.Slice(datatype_definition_enumerationOrdered[:], func(i, j int) bool {
		datatype_definition_enumerationi := datatype_definition_enumerationOrdered[i]
		datatype_definition_enumerationj := datatype_definition_enumerationOrdered[j]
		datatype_definition_enumerationi_order, oki := stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumerationi]
		datatype_definition_enumerationj_order, okj := stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumerationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_enumerationi_order < datatype_definition_enumerationj_order
	})
	if len(datatype_definition_enumerationOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {

		identifiersDecl.WriteString(datatype_definition_enumeration.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "SPECIFIED_VALUES"))
	}

	datatype_definition_integerOrdered := []*DATATYPE_DEFINITION_INTEGER{}
	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		datatype_definition_integerOrdered = append(datatype_definition_integerOrdered, datatype_definition_integer)
	}
	sort.Slice(datatype_definition_integerOrdered[:], func(i, j int) bool {
		datatype_definition_integeri := datatype_definition_integerOrdered[i]
		datatype_definition_integerj := datatype_definition_integerOrdered[j]
		datatype_definition_integeri_order, oki := stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integeri]
		datatype_definition_integerj_order, okj := stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_integeri_order < datatype_definition_integerj_order
	})
	if len(datatype_definition_integerOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datatype_definition_integer := range datatype_definition_integerOrdered {

		identifiersDecl.WriteString(datatype_definition_integer.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "MAX"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "MIN"))
		pointersInitializesStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}

	datatype_definition_realOrdered := []*DATATYPE_DEFINITION_REAL{}
	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		datatype_definition_realOrdered = append(datatype_definition_realOrdered, datatype_definition_real)
	}
	sort.Slice(datatype_definition_realOrdered[:], func(i, j int) bool {
		datatype_definition_reali := datatype_definition_realOrdered[i]
		datatype_definition_realj := datatype_definition_realOrdered[j]
		datatype_definition_reali_order, oki := stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_reali]
		datatype_definition_realj_order, okj := stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_realj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_reali_order < datatype_definition_realj_order
	})
	if len(datatype_definition_realOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datatype_definition_real := range datatype_definition_realOrdered {

		identifiersDecl.WriteString(datatype_definition_real.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "ACCURACY"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "MAX"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "MIN"))
		pointersInitializesStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}

	datatype_definition_stringOrdered := []*DATATYPE_DEFINITION_STRING{}
	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		datatype_definition_stringOrdered = append(datatype_definition_stringOrdered, datatype_definition_string)
	}
	sort.Slice(datatype_definition_stringOrdered[:], func(i, j int) bool {
		datatype_definition_stringi := datatype_definition_stringOrdered[i]
		datatype_definition_stringj := datatype_definition_stringOrdered[j]
		datatype_definition_stringi_order, oki := stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_stringi]
		datatype_definition_stringj_order, okj := stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_stringj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_stringi_order < datatype_definition_stringj_order
	})
	if len(datatype_definition_stringOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datatype_definition_string := range datatype_definition_stringOrdered {

		identifiersDecl.WriteString(datatype_definition_string.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "MAX_LENGTH"))
		pointersInitializesStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}

	datatype_definition_xhtmlOrdered := []*DATATYPE_DEFINITION_XHTML{}
	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		datatype_definition_xhtmlOrdered = append(datatype_definition_xhtmlOrdered, datatype_definition_xhtml)
	}
	sort.Slice(datatype_definition_xhtmlOrdered[:], func(i, j int) bool {
		datatype_definition_xhtmli := datatype_definition_xhtmlOrdered[i]
		datatype_definition_xhtmlj := datatype_definition_xhtmlOrdered[j]
		datatype_definition_xhtmli_order, oki := stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtmli]
		datatype_definition_xhtmlj_order, okj := stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtmlj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return datatype_definition_xhtmli_order < datatype_definition_xhtmlj_order
	})
	if len(datatype_definition_xhtmlOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {

		identifiersDecl.WriteString(datatype_definition_xhtml.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}

	embedded_valueOrdered := []*EMBEDDED_VALUE{}
	for embedded_value := range stage.EMBEDDED_VALUEs {
		embedded_valueOrdered = append(embedded_valueOrdered, embedded_value)
	}
	sort.Slice(embedded_valueOrdered[:], func(i, j int) bool {
		embedded_valuei := embedded_valueOrdered[i]
		embedded_valuej := embedded_valueOrdered[j]
		embedded_valuei_order, oki := stage.EMBEDDED_VALUEMap_Staged_Order[embedded_valuei]
		embedded_valuej_order, okj := stage.EMBEDDED_VALUEMap_Staged_Order[embedded_valuej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return embedded_valuei_order < embedded_valuej_order
	})
	if len(embedded_valueOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, embedded_value := range embedded_valueOrdered {

		identifiersDecl.WriteString(embedded_value.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(embedded_value.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(embedded_value.GongMarshallField(stage, "KEY"))
		initializerStatements.WriteString(embedded_value.GongMarshallField(stage, "OTHER_CONTENT"))
	}

	enum_valueOrdered := []*ENUM_VALUE{}
	for enum_value := range stage.ENUM_VALUEs {
		enum_valueOrdered = append(enum_valueOrdered, enum_value)
	}
	sort.Slice(enum_valueOrdered[:], func(i, j int) bool {
		enum_valuei := enum_valueOrdered[i]
		enum_valuej := enum_valueOrdered[j]
		enum_valuei_order, oki := stage.ENUM_VALUEMap_Staged_Order[enum_valuei]
		enum_valuej_order, okj := stage.ENUM_VALUEMap_Staged_Order[enum_valuej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return enum_valuei_order < enum_valuej_order
	})
	if len(enum_valueOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, enum_value := range enum_valueOrdered {

		identifiersDecl.WriteString(enum_value.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(enum_value.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(enum_value.GongMarshallField(stage, "PROPERTIES"))
	}

	relation_groupOrdered := []*RELATION_GROUP{}
	for relation_group := range stage.RELATION_GROUPs {
		relation_groupOrdered = append(relation_groupOrdered, relation_group)
	}
	sort.Slice(relation_groupOrdered[:], func(i, j int) bool {
		relation_groupi := relation_groupOrdered[i]
		relation_groupj := relation_groupOrdered[j]
		relation_groupi_order, oki := stage.RELATION_GROUPMap_Staged_Order[relation_groupi]
		relation_groupj_order, okj := stage.RELATION_GROUPMap_Staged_Order[relation_groupj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return relation_groupi_order < relation_groupj_order
	})
	if len(relation_groupOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, relation_group := range relation_groupOrdered {

		identifiersDecl.WriteString(relation_group.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "SOURCE_SPECIFICATION"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "SPEC_RELATIONS"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "TARGET_SPECIFICATION"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "TYPE"))
	}

	relation_group_typeOrdered := []*RELATION_GROUP_TYPE{}
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		relation_group_typeOrdered = append(relation_group_typeOrdered, relation_group_type)
	}
	sort.Slice(relation_group_typeOrdered[:], func(i, j int) bool {
		relation_group_typei := relation_group_typeOrdered[i]
		relation_group_typej := relation_group_typeOrdered[j]
		relation_group_typei_order, oki := stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_typei]
		relation_group_typej_order, okj := stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return relation_group_typei_order < relation_group_typej_order
	})
	if len(relation_group_typeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, relation_group_type := range relation_group_typeOrdered {

		identifiersDecl.WriteString(relation_group_type.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(relation_group_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(relation_group_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}

	req_ifOrdered := []*REQ_IF{}
	for req_if := range stage.REQ_IFs {
		req_ifOrdered = append(req_ifOrdered, req_if)
	}
	sort.Slice(req_ifOrdered[:], func(i, j int) bool {
		req_ifi := req_ifOrdered[i]
		req_ifj := req_ifOrdered[j]
		req_ifi_order, oki := stage.REQ_IFMap_Staged_Order[req_ifi]
		req_ifj_order, okj := stage.REQ_IFMap_Staged_Order[req_ifj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_ifi_order < req_ifj_order
	})
	if len(req_ifOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, req_if := range req_ifOrdered {

		identifiersDecl.WriteString(req_if.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(req_if.GongMarshallField(stage, "Lang"))
		pointersInitializesStatements.WriteString(req_if.GongMarshallField(stage, "THE_HEADER"))
		pointersInitializesStatements.WriteString(req_if.GongMarshallField(stage, "CORE_CONTENT"))
		pointersInitializesStatements.WriteString(req_if.GongMarshallField(stage, "TOOL_EXTENSIONS"))
	}

	req_if_contentOrdered := []*REQ_IF_CONTENT{}
	for req_if_content := range stage.REQ_IF_CONTENTs {
		req_if_contentOrdered = append(req_if_contentOrdered, req_if_content)
	}
	sort.Slice(req_if_contentOrdered[:], func(i, j int) bool {
		req_if_contenti := req_if_contentOrdered[i]
		req_if_contentj := req_if_contentOrdered[j]
		req_if_contenti_order, oki := stage.REQ_IF_CONTENTMap_Staged_Order[req_if_contenti]
		req_if_contentj_order, okj := stage.REQ_IF_CONTENTMap_Staged_Order[req_if_contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_if_contenti_order < req_if_contentj_order
	})
	if len(req_if_contentOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, req_if_content := range req_if_contentOrdered {

		identifiersDecl.WriteString(req_if_content.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if_content.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "DATATYPES"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_TYPES"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_OBJECTS"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_RELATIONS"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPECIFICATIONS"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_RELATION_GROUPS"))
	}

	req_if_headerOrdered := []*REQ_IF_HEADER{}
	for req_if_header := range stage.REQ_IF_HEADERs {
		req_if_headerOrdered = append(req_if_headerOrdered, req_if_header)
	}
	sort.Slice(req_if_headerOrdered[:], func(i, j int) bool {
		req_if_headeri := req_if_headerOrdered[i]
		req_if_headerj := req_if_headerOrdered[j]
		req_if_headeri_order, oki := stage.REQ_IF_HEADERMap_Staged_Order[req_if_headeri]
		req_if_headerj_order, okj := stage.REQ_IF_HEADERMap_Staged_Order[req_if_headerj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_if_headeri_order < req_if_headerj_order
	})
	if len(req_if_headerOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, req_if_header := range req_if_headerOrdered {

		identifiersDecl.WriteString(req_if_header.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "COMMENT"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "CREATION_TIME"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "REPOSITORY_ID"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "REQ_IF_TOOL_ID"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "REQ_IF_VERSION"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "SOURCE_TOOL_ID"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "TITLE"))
	}

	req_if_tool_extensionOrdered := []*REQ_IF_TOOL_EXTENSION{}
	for req_if_tool_extension := range stage.REQ_IF_TOOL_EXTENSIONs {
		req_if_tool_extensionOrdered = append(req_if_tool_extensionOrdered, req_if_tool_extension)
	}
	sort.Slice(req_if_tool_extensionOrdered[:], func(i, j int) bool {
		req_if_tool_extensioni := req_if_tool_extensionOrdered[i]
		req_if_tool_extensionj := req_if_tool_extensionOrdered[j]
		req_if_tool_extensioni_order, oki := stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extensioni]
		req_if_tool_extensionj_order, okj := stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extensionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return req_if_tool_extensioni_order < req_if_tool_extensionj_order
	})
	if len(req_if_tool_extensionOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, req_if_tool_extension := range req_if_tool_extensionOrdered {

		identifiersDecl.WriteString(req_if_tool_extension.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if_tool_extension.GongMarshallField(stage, "Name"))
	}

	specificationOrdered := []*SPECIFICATION{}
	for specification := range stage.SPECIFICATIONs {
		specificationOrdered = append(specificationOrdered, specification)
	}
	sort.Slice(specificationOrdered[:], func(i, j int) bool {
		specificationi := specificationOrdered[i]
		specificationj := specificationOrdered[j]
		specificationi_order, oki := stage.SPECIFICATIONMap_Staged_Order[specificationi]
		specificationj_order, okj := stage.SPECIFICATIONMap_Staged_Order[specificationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return specificationi_order < specificationj_order
	})
	if len(specificationOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, specification := range specificationOrdered {

		identifiersDecl.WriteString(specification.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(specification.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "CHILDREN"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "VALUES"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "TYPE"))
	}

	specification_typeOrdered := []*SPECIFICATION_TYPE{}
	for specification_type := range stage.SPECIFICATION_TYPEs {
		specification_typeOrdered = append(specification_typeOrdered, specification_type)
	}
	sort.Slice(specification_typeOrdered[:], func(i, j int) bool {
		specification_typei := specification_typeOrdered[i]
		specification_typej := specification_typeOrdered[j]
		specification_typei_order, oki := stage.SPECIFICATION_TYPEMap_Staged_Order[specification_typei]
		specification_typej_order, okj := stage.SPECIFICATION_TYPEMap_Staged_Order[specification_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return specification_typei_order < specification_typej_order
	})
	if len(specification_typeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, specification_type := range specification_typeOrdered {

		identifiersDecl.WriteString(specification_type.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(specification_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(specification_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}

	spec_hierarchyOrdered := []*SPEC_HIERARCHY{}
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		spec_hierarchyOrdered = append(spec_hierarchyOrdered, spec_hierarchy)
	}
	sort.Slice(spec_hierarchyOrdered[:], func(i, j int) bool {
		spec_hierarchyi := spec_hierarchyOrdered[i]
		spec_hierarchyj := spec_hierarchyOrdered[j]
		spec_hierarchyi_order, oki := stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchyi]
		spec_hierarchyj_order, okj := stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchyj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_hierarchyi_order < spec_hierarchyj_order
	})
	if len(spec_hierarchyOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spec_hierarchy := range spec_hierarchyOrdered {

		identifiersDecl.WriteString(spec_hierarchy.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "IS_TABLE_INTERNAL"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "CHILDREN"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "EDITABLE_ATTS"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "OBJECT"))
	}

	spec_objectOrdered := []*SPEC_OBJECT{}
	for spec_object := range stage.SPEC_OBJECTs {
		spec_objectOrdered = append(spec_objectOrdered, spec_object)
	}
	sort.Slice(spec_objectOrdered[:], func(i, j int) bool {
		spec_objecti := spec_objectOrdered[i]
		spec_objectj := spec_objectOrdered[j]
		spec_objecti_order, oki := stage.SPEC_OBJECTMap_Staged_Order[spec_objecti]
		spec_objectj_order, okj := stage.SPEC_OBJECTMap_Staged_Order[spec_objectj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_objecti_order < spec_objectj_order
	})
	if len(spec_objectOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spec_object := range spec_objectOrdered {

		identifiersDecl.WriteString(spec_object.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_object.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_object.GongMarshallField(stage, "VALUES"))
		pointersInitializesStatements.WriteString(spec_object.GongMarshallField(stage, "TYPE"))
	}

	spec_object_typeOrdered := []*SPEC_OBJECT_TYPE{}
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		spec_object_typeOrdered = append(spec_object_typeOrdered, spec_object_type)
	}
	sort.Slice(spec_object_typeOrdered[:], func(i, j int) bool {
		spec_object_typei := spec_object_typeOrdered[i]
		spec_object_typej := spec_object_typeOrdered[j]
		spec_object_typei_order, oki := stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_typei]
		spec_object_typej_order, okj := stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_object_typei_order < spec_object_typej_order
	})
	if len(spec_object_typeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spec_object_type := range spec_object_typeOrdered {

		identifiersDecl.WriteString(spec_object_type.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_object_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_object_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}

	spec_relationOrdered := []*SPEC_RELATION{}
	for spec_relation := range stage.SPEC_RELATIONs {
		spec_relationOrdered = append(spec_relationOrdered, spec_relation)
	}
	sort.Slice(spec_relationOrdered[:], func(i, j int) bool {
		spec_relationi := spec_relationOrdered[i]
		spec_relationj := spec_relationOrdered[j]
		spec_relationi_order, oki := stage.SPEC_RELATIONMap_Staged_Order[spec_relationi]
		spec_relationj_order, okj := stage.SPEC_RELATIONMap_Staged_Order[spec_relationj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_relationi_order < spec_relationj_order
	})
	if len(spec_relationOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spec_relation := range spec_relationOrdered {

		identifiersDecl.WriteString(spec_relation.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "VALUES"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "SOURCE"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "TARGET"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "TYPE"))
	}

	spec_relation_typeOrdered := []*SPEC_RELATION_TYPE{}
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		spec_relation_typeOrdered = append(spec_relation_typeOrdered, spec_relation_type)
	}
	sort.Slice(spec_relation_typeOrdered[:], func(i, j int) bool {
		spec_relation_typei := spec_relation_typeOrdered[i]
		spec_relation_typej := spec_relation_typeOrdered[j]
		spec_relation_typei_order, oki := stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_typei]
		spec_relation_typej_order, okj := stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_typej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return spec_relation_typei_order < spec_relation_typej_order
	})
	if len(spec_relation_typeOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, spec_relation_type := range spec_relation_typeOrdered {

		identifiersDecl.WriteString(spec_relation_type.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_relation_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_relation_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}

	xhtml_contentOrdered := []*XHTML_CONTENT{}
	for xhtml_content := range stage.XHTML_CONTENTs {
		xhtml_contentOrdered = append(xhtml_contentOrdered, xhtml_content)
	}
	sort.Slice(xhtml_contentOrdered[:], func(i, j int) bool {
		xhtml_contenti := xhtml_contentOrdered[i]
		xhtml_contentj := xhtml_contentOrdered[j]
		xhtml_contenti_order, oki := stage.XHTML_CONTENTMap_Staged_Order[xhtml_contenti]
		xhtml_contentj_order, okj := stage.XHTML_CONTENTMap_Staged_Order[xhtml_contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return xhtml_contenti_order < xhtml_contentj_order
	})
	if len(xhtml_contentOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, xhtml_content := range xhtml_contentOrdered {

		identifiersDecl.WriteString(xhtml_content.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(xhtml_content.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xhtml_content.GongMarshallField(stage, "EnclosedText"))
		initializerStatements.WriteString(xhtml_content.GongMarshallField(stage, "PureText"))
	}

	// insertion initialization of objects to stage
	for _, alternative_id := range alternative_idOrdered {
		_ = alternative_id
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_definition_boolean := range attribute_definition_booleanOrdered {
		_ = attribute_definition_boolean
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_definition_date := range attribute_definition_dateOrdered {
		_ = attribute_definition_date
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_definition_enumeration := range attribute_definition_enumerationOrdered {
		_ = attribute_definition_enumeration
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_definition_integer := range attribute_definition_integerOrdered {
		_ = attribute_definition_integer
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_definition_real := range attribute_definition_realOrdered {
		_ = attribute_definition_real
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_definition_string := range attribute_definition_stringOrdered {
		_ = attribute_definition_string
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_definition_xhtml := range attribute_definition_xhtmlOrdered {
		_ = attribute_definition_xhtml
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_value_boolean := range attribute_value_booleanOrdered {
		_ = attribute_value_boolean
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_value_date := range attribute_value_dateOrdered {
		_ = attribute_value_date
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_value_enumeration := range attribute_value_enumerationOrdered {
		_ = attribute_value_enumeration
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_value_integer := range attribute_value_integerOrdered {
		_ = attribute_value_integer
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_value_real := range attribute_value_realOrdered {
		_ = attribute_value_real
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_value_string := range attribute_value_stringOrdered {
		_ = attribute_value_string
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, attribute_value_xhtml := range attribute_value_xhtmlOrdered {
		_ = attribute_value_xhtml
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_alternative_id := range a_alternative_idOrdered {
		_ = a_alternative_id
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_definition_boolean_ref := range a_attribute_definition_boolean_refOrdered {
		_ = a_attribute_definition_boolean_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_definition_date_ref := range a_attribute_definition_date_refOrdered {
		_ = a_attribute_definition_date_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_definition_enumeration_ref := range a_attribute_definition_enumeration_refOrdered {
		_ = a_attribute_definition_enumeration_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_definition_integer_ref := range a_attribute_definition_integer_refOrdered {
		_ = a_attribute_definition_integer_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_definition_real_ref := range a_attribute_definition_real_refOrdered {
		_ = a_attribute_definition_real_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_definition_string_ref := range a_attribute_definition_string_refOrdered {
		_ = a_attribute_definition_string_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_definition_xhtml_ref := range a_attribute_definition_xhtml_refOrdered {
		_ = a_attribute_definition_xhtml_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_boolean := range a_attribute_value_booleanOrdered {
		_ = a_attribute_value_boolean
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_date := range a_attribute_value_dateOrdered {
		_ = a_attribute_value_date
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_enumeration := range a_attribute_value_enumerationOrdered {
		_ = a_attribute_value_enumeration
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_integer := range a_attribute_value_integerOrdered {
		_ = a_attribute_value_integer
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_real := range a_attribute_value_realOrdered {
		_ = a_attribute_value_real
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_string := range a_attribute_value_stringOrdered {
		_ = a_attribute_value_string
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_xhtml := range a_attribute_value_xhtmlOrdered {
		_ = a_attribute_value_xhtml
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_attribute_value_xhtml_1 := range a_attribute_value_xhtml_1Ordered {
		_ = a_attribute_value_xhtml_1
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_children := range a_childrenOrdered {
		_ = a_children
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_core_content := range a_core_contentOrdered {
		_ = a_core_content
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatypes := range a_datatypesOrdered {
		_ = a_datatypes
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatype_definition_boolean_ref := range a_datatype_definition_boolean_refOrdered {
		_ = a_datatype_definition_boolean_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatype_definition_date_ref := range a_datatype_definition_date_refOrdered {
		_ = a_datatype_definition_date_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatype_definition_enumeration_ref := range a_datatype_definition_enumeration_refOrdered {
		_ = a_datatype_definition_enumeration_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatype_definition_integer_ref := range a_datatype_definition_integer_refOrdered {
		_ = a_datatype_definition_integer_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatype_definition_real_ref := range a_datatype_definition_real_refOrdered {
		_ = a_datatype_definition_real_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatype_definition_string_ref := range a_datatype_definition_string_refOrdered {
		_ = a_datatype_definition_string_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_datatype_definition_xhtml_ref := range a_datatype_definition_xhtml_refOrdered {
		_ = a_datatype_definition_xhtml_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_editable_atts := range a_editable_attsOrdered {
		_ = a_editable_atts
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_enum_value_ref := range a_enum_value_refOrdered {
		_ = a_enum_value_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_object := range a_objectOrdered {
		_ = a_object
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_properties := range a_propertiesOrdered {
		_ = a_properties
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_relation_group_type_ref := range a_relation_group_type_refOrdered {
		_ = a_relation_group_type_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_source_1 := range a_source_1Ordered {
		_ = a_source_1
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_source_specification_1 := range a_source_specification_1Ordered {
		_ = a_source_specification_1
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_specifications := range a_specificationsOrdered {
		_ = a_specifications
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_specification_type_ref := range a_specification_type_refOrdered {
		_ = a_specification_type_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_specified_values := range a_specified_valuesOrdered {
		_ = a_specified_values
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_attributes := range a_spec_attributesOrdered {
		_ = a_spec_attributes
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_objects := range a_spec_objectsOrdered {
		_ = a_spec_objects
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_object_type_ref := range a_spec_object_type_refOrdered {
		_ = a_spec_object_type_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_relations := range a_spec_relationsOrdered {
		_ = a_spec_relations
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_relation_groups := range a_spec_relation_groupsOrdered {
		_ = a_spec_relation_groups
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_relation_ref := range a_spec_relation_refOrdered {
		_ = a_spec_relation_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_relation_type_ref := range a_spec_relation_type_refOrdered {
		_ = a_spec_relation_type_ref
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_spec_types := range a_spec_typesOrdered {
		_ = a_spec_types
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_the_header := range a_the_headerOrdered {
		_ = a_the_header
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, a_tool_extensions := range a_tool_extensionsOrdered {
		_ = a_tool_extensions
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datatype_definition_boolean := range datatype_definition_booleanOrdered {
		_ = datatype_definition_boolean
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datatype_definition_date := range datatype_definition_dateOrdered {
		_ = datatype_definition_date
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datatype_definition_enumeration := range datatype_definition_enumerationOrdered {
		_ = datatype_definition_enumeration
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datatype_definition_integer := range datatype_definition_integerOrdered {
		_ = datatype_definition_integer
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datatype_definition_real := range datatype_definition_realOrdered {
		_ = datatype_definition_real
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datatype_definition_string := range datatype_definition_stringOrdered {
		_ = datatype_definition_string
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, datatype_definition_xhtml := range datatype_definition_xhtmlOrdered {
		_ = datatype_definition_xhtml
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, embedded_value := range embedded_valueOrdered {
		_ = embedded_value
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, enum_value := range enum_valueOrdered {
		_ = enum_value
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, relation_group := range relation_groupOrdered {
		_ = relation_group
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, relation_group_type := range relation_group_typeOrdered {
		_ = relation_group_type
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, req_if := range req_ifOrdered {
		_ = req_if
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, req_if_content := range req_if_contentOrdered {
		_ = req_if_content
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, req_if_header := range req_if_headerOrdered {
		_ = req_if_header
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, req_if_tool_extension := range req_if_tool_extensionOrdered {
		_ = req_if_tool_extension
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, specification := range specificationOrdered {
		_ = specification
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, specification_type := range specification_typeOrdered {
		_ = specification_type
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spec_hierarchy := range spec_hierarchyOrdered {
		_ = spec_hierarchy
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spec_object := range spec_objectOrdered {
		_ = spec_object
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spec_object_type := range spec_object_typeOrdered {
		_ = spec_object_type
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spec_relation := range spec_relationOrdered {
		_ = spec_relation
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, spec_relation_type := range spec_relation_typeOrdered {
		_ = spec_relation_type
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, xhtml_content := range xhtml_contentOrdered {
		_ = xhtml_content
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
func (alternative_id *ALTERNATIVE_ID) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", alternative_id.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(alternative_id.Name))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", alternative_id.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(alternative_id.IDENTIFIER))

	default:
		log.Panicf("Unknown field %s for Gongstruct ALTERNATIVE_ID", fieldName)
	}
	return
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_boolean.IS_EDITABLE))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_boolean.LONG_NAME))

	case "ALTERNATIVE_ID":
		if attribute_definition_boolean.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_boolean.DEFAULT_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_boolean.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if attribute_definition_boolean.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_boolean.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_BOOLEAN", fieldName)
	}
	return
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_date.IS_EDITABLE))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_date.LONG_NAME))

	case "ALTERNATIVE_ID":
		if attribute_definition_date.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_date.DEFAULT_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_date.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if attribute_definition_date.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_date.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_DATE", fieldName)
	}
	return
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration.IS_EDITABLE))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_enumeration.LONG_NAME))
	case "MULTI_VALUED":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MULTI_VALUED")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_enumeration.MULTI_VALUED))

	case "ALTERNATIVE_ID":
		if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_enumeration.DEFAULT_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if attribute_definition_enumeration.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_ENUMERATION", fieldName)
	}
	return
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_integer.IS_EDITABLE))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_integer.LONG_NAME))

	case "ALTERNATIVE_ID":
		if attribute_definition_integer.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_integer.DEFAULT_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_integer.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if attribute_definition_integer.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_integer.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_INTEGER", fieldName)
	}
	return
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_real.IS_EDITABLE))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_real.LONG_NAME))

	case "ALTERNATIVE_ID":
		if attribute_definition_real.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_real.DEFAULT_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_real.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if attribute_definition_real.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_real.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_REAL", fieldName)
	}
	return
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_string.IS_EDITABLE))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_string.LONG_NAME))

	case "ALTERNATIVE_ID":
		if attribute_definition_string.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_string.DEFAULT_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_string.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if attribute_definition_string.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_string.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_STRING", fieldName)
	}
	return
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_definition_xhtml.IS_EDITABLE))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_definition_xhtml.LONG_NAME))

	case "ALTERNATIVE_ID":
		if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_xhtml.DEFAULT_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFAULT_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if attribute_definition_xhtml.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_XHTML", fieldName)
	}
	return
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_boolean.Name))
	case "THE_VALUE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_VALUE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_value_boolean.THE_VALUE))

	case "DEFINITION":
		if attribute_value_boolean.DEFINITION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_boolean.DEFINITION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_BOOLEAN", fieldName)
	}
	return
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_date.Name))
	case "THE_VALUE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_VALUE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_date.THE_VALUE))

	case "DEFINITION":
		if attribute_value_date.DEFINITION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_date.DEFINITION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_DATE", fieldName)
	}
	return
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_enumeration.Name))

	case "DEFINITION":
		if attribute_value_enumeration.DEFINITION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_enumeration.DEFINITION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "VALUES":
		if attribute_value_enumeration.VALUES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_enumeration.VALUES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_ENUMERATION", fieldName)
	}
	return
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_integer.Name))
	case "THE_VALUE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_VALUE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", attribute_value_integer.THE_VALUE))

	case "DEFINITION":
		if attribute_value_integer.DEFINITION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_integer.DEFINITION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_INTEGER", fieldName)
	}
	return
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_real.Name))
	case "THE_VALUE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_VALUE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", attribute_value_real.THE_VALUE))

	case "DEFINITION":
		if attribute_value_real.DEFINITION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_real.DEFINITION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_REAL", fieldName)
	}
	return
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_string.Name))
	case "THE_VALUE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_VALUE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_string.THE_VALUE))

	case "DEFINITION":
		if attribute_value_string.DEFINITION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_string.DEFINITION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_STRING", fieldName)
	}
	return
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(attribute_value_xhtml.Name))
	case "IS_SIMPLIFIED":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_SIMPLIFIED")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", attribute_value_xhtml.IS_SIMPLIFIED))

	case "THE_VALUE":
		if attribute_value_xhtml.THE_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_xhtml.THE_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "THE_ORIGINAL_VALUE":
		if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_ORIGINAL_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_xhtml.THE_ORIGINAL_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_ORIGINAL_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DEFINITION":
		if attribute_value_xhtml.DEFINITION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", attribute_value_xhtml.DEFINITION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DEFINITION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_XHTML", fieldName)
	}
	return
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_alternative_id.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_alternative_id.Name))

	case "ALTERNATIVE_ID":
		if a_alternative_id.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_alternative_id.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", a_alternative_id.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_alternative_id.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ALTERNATIVE_ID", fieldName)
	}
	return
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_boolean_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_boolean_ref.Name))
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_boolean_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_boolean_ref.ATTRIBUTE_DEFINITION_BOOLEAN_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", fieldName)
	}
	return
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_date_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_date_ref.Name))
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_date_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_date_ref.ATTRIBUTE_DEFINITION_DATE_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_DATE_REF", fieldName)
	}
	return
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_enumeration_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_enumeration_ref.Name))
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_enumeration_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_enumeration_ref.ATTRIBUTE_DEFINITION_ENUMERATION_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", fieldName)
	}
	return
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_integer_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_integer_ref.Name))
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_integer_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_integer_ref.ATTRIBUTE_DEFINITION_INTEGER_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_INTEGER_REF", fieldName)
	}
	return
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_real_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_real_ref.Name))
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_real_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_real_ref.ATTRIBUTE_DEFINITION_REAL_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_REAL_REF", fieldName)
	}
	return
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_string_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_string_ref.Name))
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_string_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_string_ref.ATTRIBUTE_DEFINITION_STRING_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_STRING_REF", fieldName)
	}
	return
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_xhtml_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_xhtml_ref.Name))
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_definition_xhtml_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_definition_xhtml_ref.ATTRIBUTE_DEFINITION_XHTML_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_XHTML_REF", fieldName)
	}
	return
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_boolean.Name))

	case "ATTRIBUTE_VALUE_BOOLEAN":
		var sb strings.Builder
		for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_boolean.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_BOOLEAN")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_boolean.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_BOOLEAN", fieldName)
	}
	return
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_date.Name))

	case "ATTRIBUTE_VALUE_DATE":
		var sb strings.Builder
		for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_date.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_DATE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_date.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_DATE", fieldName)
	}
	return
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_enumeration.Name))

	case "ATTRIBUTE_VALUE_ENUMERATION":
		var sb strings.Builder
		for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_enumeration.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_ENUMERATION")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_enumeration.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_ENUMERATION", fieldName)
	}
	return
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_integer.Name))

	case "ATTRIBUTE_VALUE_INTEGER":
		var sb strings.Builder
		for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_integer.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_INTEGER")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_integer.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_INTEGER", fieldName)
	}
	return
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_real.Name))

	case "ATTRIBUTE_VALUE_REAL":
		var sb strings.Builder
		for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_real.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_REAL")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_real.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_REAL", fieldName)
	}
	return
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_string.Name))

	case "ATTRIBUTE_VALUE_STRING":
		var sb strings.Builder
		for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_string.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_STRING")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_string.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_STRING", fieldName)
	}
	return
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_xhtml.Name))

	case "ATTRIBUTE_VALUE_XHTML":
		var sb strings.Builder
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_XHTML")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_xhtml.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_XHTML", fieldName)
	}
	return
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_attribute_value_xhtml_1.Name))

	case "ATTRIBUTE_VALUE_BOOLEAN":
		var sb strings.Builder
		for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_BOOLEAN")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_boolean.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_DATE":
		var sb strings.Builder
		for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_DATE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_date.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_ENUMERATION":
		var sb strings.Builder
		for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_ENUMERATION")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_enumeration.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_INTEGER":
		var sb strings.Builder
		for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_INTEGER")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_integer.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_REAL":
		var sb strings.Builder
		for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_REAL")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_real.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_STRING":
		var sb strings.Builder
		for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_STRING")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_string.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_XHTML":
		var sb strings.Builder
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_VALUE_XHTML")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_value_xhtml.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_XHTML_1", fieldName)
	}
	return
}

func (a_children *A_CHILDREN) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_children.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_children.Name))

	case "SPEC_HIERARCHY":
		var sb strings.Builder
		for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_children.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SPEC_HIERARCHY")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spec_hierarchy.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_CHILDREN", fieldName)
	}
	return
}

func (a_core_content *A_CORE_CONTENT) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_core_content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_core_content.Name))

	case "REQ_IF_CONTENT":
		if a_core_content.REQ_IF_CONTENT != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_core_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "REQ_IF_CONTENT")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", a_core_content.REQ_IF_CONTENT.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_core_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "REQ_IF_CONTENT")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_CORE_CONTENT", fieldName)
	}
	return
}

func (a_datatypes *A_DATATYPES) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatypes.Name))

	case "DATATYPE_DEFINITION_BOOLEAN":
		var sb strings.Builder
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_BOOLEAN")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datatype_definition_boolean.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_DATE":
		var sb strings.Builder
		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_DATE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datatype_definition_date.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_ENUMERATION":
		var sb strings.Builder
		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_ENUMERATION")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datatype_definition_enumeration.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_INTEGER":
		var sb strings.Builder
		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_INTEGER")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datatype_definition_integer.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_REAL":
		var sb strings.Builder
		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_REAL")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datatype_definition_real.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_STRING":
		var sb strings.Builder
		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_STRING")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datatype_definition_string.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_XHTML":
		var sb strings.Builder
		for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_XHTML")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _datatype_definition_xhtml.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPES", fieldName)
	}
	return
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_boolean_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_boolean_ref.Name))
	case "DATATYPE_DEFINITION_BOOLEAN_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_boolean_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_BOOLEAN_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_boolean_ref.DATATYPE_DEFINITION_BOOLEAN_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_BOOLEAN_REF", fieldName)
	}
	return
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_date_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_date_ref.Name))
	case "DATATYPE_DEFINITION_DATE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_date_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_DATE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_date_ref.DATATYPE_DEFINITION_DATE_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_DATE_REF", fieldName)
	}
	return
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_enumeration_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_enumeration_ref.Name))
	case "DATATYPE_DEFINITION_ENUMERATION_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_enumeration_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_ENUMERATION_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_enumeration_ref.DATATYPE_DEFINITION_ENUMERATION_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_ENUMERATION_REF", fieldName)
	}
	return
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_integer_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_integer_ref.Name))
	case "DATATYPE_DEFINITION_INTEGER_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_integer_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_INTEGER_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_integer_ref.DATATYPE_DEFINITION_INTEGER_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_INTEGER_REF", fieldName)
	}
	return
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_real_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_real_ref.Name))
	case "DATATYPE_DEFINITION_REAL_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_real_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_REAL_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_real_ref.DATATYPE_DEFINITION_REAL_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_REAL_REF", fieldName)
	}
	return
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_string_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_string_ref.Name))
	case "DATATYPE_DEFINITION_STRING_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_string_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_STRING_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_string_ref.DATATYPE_DEFINITION_STRING_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_STRING_REF", fieldName)
	}
	return
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_xhtml_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_xhtml_ref.Name))
	case "DATATYPE_DEFINITION_XHTML_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_datatype_definition_xhtml_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPE_DEFINITION_XHTML_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_datatype_definition_xhtml_ref.DATATYPE_DEFINITION_XHTML_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_XHTML_REF", fieldName)
	}
	return
}

func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.Name))
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_BOOLEAN_REF))
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_DATE_REF))
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_ENUMERATION_REF))
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_INTEGER_REF))
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_REAL_REF))
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_STRING_REF))
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_editable_atts.ATTRIBUTE_DEFINITION_XHTML_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_EDITABLE_ATTS", fieldName)
	}
	return
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_enum_value_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_enum_value_ref.Name))
	case "ENUM_VALUE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_enum_value_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ENUM_VALUE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_enum_value_ref.ENUM_VALUE_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ENUM_VALUE_REF", fieldName)
	}
	return
}

func (a_object *A_OBJECT) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_object.Name))
	case "SPEC_OBJECT_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_OBJECT_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_object.SPEC_OBJECT_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_OBJECT", fieldName)
	}
	return
}

func (a_properties *A_PROPERTIES) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_properties.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_properties.Name))

	case "EMBEDDED_VALUE":
		if a_properties.EMBEDDED_VALUE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_properties.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EMBEDDED_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", a_properties.EMBEDDED_VALUE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_properties.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EMBEDDED_VALUE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_PROPERTIES", fieldName)
	}
	return
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_relation_group_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_relation_group_type_ref.Name))
	case "RELATION_GROUP_TYPE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_relation_group_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "RELATION_GROUP_TYPE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_relation_group_type_ref.RELATION_GROUP_TYPE_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_RELATION_GROUP_TYPE_REF", fieldName)
	}
	return
}

func (a_source_1 *A_SOURCE_1) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_source_1.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_source_1.Name))
	case "SPEC_OBJECT_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_source_1.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_OBJECT_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_source_1.SPEC_OBJECT_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SOURCE_1", fieldName)
	}
	return
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_source_specification_1.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_source_specification_1.Name))
	case "SPECIFICATION_REF":
		if a_source_specification_1.SPECIFICATION_REF.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_source_specification_1.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPECIFICATION_REF")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+a_source_specification_1.SPECIFICATION_REF.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_source_specification_1.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPECIFICATION_REF")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SOURCE_SPECIFICATION_1", fieldName)
	}
	return
}

func (a_specifications *A_SPECIFICATIONS) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_specifications.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_specifications.Name))

	case "SPECIFICATION":
		var sb strings.Builder
		for _, _specification := range a_specifications.SPECIFICATION {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_specifications.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SPECIFICATION")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _specification.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPECIFICATIONS", fieldName)
	}
	return
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_specification_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_specification_type_ref.Name))
	case "SPECIFICATION_TYPE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_specification_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPECIFICATION_TYPE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_specification_type_ref.SPECIFICATION_TYPE_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPECIFICATION_TYPE_REF", fieldName)
	}
	return
}

func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_specified_values.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_specified_values.Name))

	case "ENUM_VALUE":
		var sb strings.Builder
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_specified_values.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ENUM_VALUE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _enum_value.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPECIFIED_VALUES", fieldName)
	}
	return
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_attributes.Name))

	case "ATTRIBUTE_DEFINITION_BOOLEAN":
		var sb strings.Builder
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_BOOLEAN")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_definition_boolean.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_DATE":
		var sb strings.Builder
		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_DATE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_definition_date.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_ENUMERATION":
		var sb strings.Builder
		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_ENUMERATION")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_definition_enumeration.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_INTEGER":
		var sb strings.Builder
		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_INTEGER")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_definition_integer.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_REAL":
		var sb strings.Builder
		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_REAL")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_definition_real.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_STRING":
		var sb strings.Builder
		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_STRING")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_definition_string.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_XHTML":
		var sb strings.Builder
		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "ATTRIBUTE_DEFINITION_XHTML")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _attribute_definition_xhtml.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_ATTRIBUTES", fieldName)
	}
	return
}

func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_objects.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_objects.Name))

	case "SPEC_OBJECT":
		var sb strings.Builder
		for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_objects.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SPEC_OBJECT")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spec_object.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_OBJECTS", fieldName)
	}
	return
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_object_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_object_type_ref.Name))
	case "SPEC_OBJECT_TYPE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_object_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_OBJECT_TYPE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_object_type_ref.SPEC_OBJECT_TYPE_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_OBJECT_TYPE_REF", fieldName)
	}
	return
}

func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_relations.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_relations.Name))

	case "SPEC_RELATION":
		var sb strings.Builder
		for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_relations.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SPEC_RELATION")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spec_relation.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATIONS", fieldName)
	}
	return
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_relation_groups.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_relation_groups.Name))

	case "RELATION_GROUP":
		var sb strings.Builder
		for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_relation_groups.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RELATION_GROUP")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _relation_group.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATION_GROUPS", fieldName)
	}
	return
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_relation_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_relation_ref.Name))
	case "SPEC_RELATION_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_relation_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATION_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_relation_ref.SPEC_RELATION_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATION_REF", fieldName)
	}
	return
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_relation_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_relation_type_ref.Name))
	case "SPEC_RELATION_TYPE_REF":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_relation_type_ref.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATION_TYPE_REF")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_relation_type_ref.SPEC_RELATION_TYPE_REF))

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATION_TYPE_REF", fieldName)
	}
	return
}

func (a_spec_types *A_SPEC_TYPES) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_spec_types.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_spec_types.Name))

	case "RELATION_GROUP_TYPE":
		var sb strings.Builder
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_types.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "RELATION_GROUP_TYPE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _relation_group_type.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SPEC_OBJECT_TYPE":
		var sb strings.Builder
		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_types.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SPEC_OBJECT_TYPE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spec_object_type.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SPEC_RELATION_TYPE":
		var sb strings.Builder
		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_types.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SPEC_RELATION_TYPE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _spec_relation_type.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SPECIFICATION_TYPE":
		var sb strings.Builder
		for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_spec_types.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SPECIFICATION_TYPE")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _specification_type.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_TYPES", fieldName)
	}
	return
}

func (a_the_header *A_THE_HEADER) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_the_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_the_header.Name))

	case "REQ_IF_HEADER":
		if a_the_header.REQ_IF_HEADER != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_the_header.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "REQ_IF_HEADER")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", a_the_header.REQ_IF_HEADER.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", a_the_header.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "REQ_IF_HEADER")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_THE_HEADER", fieldName)
	}
	return
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", a_tool_extensions.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(a_tool_extensions.Name))

	case "REQ_IF_TOOL_EXTENSION":
		var sb strings.Builder
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", a_tool_extensions.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "REQ_IF_TOOL_EXTENSION")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _req_if_tool_extension.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_TOOL_EXTENSIONS", fieldName)
	}
	return
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_boolean.LONG_NAME))

	case "ALTERNATIVE_ID":
		if datatype_definition_boolean.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_BOOLEAN", fieldName)
	}
	return
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_date.LONG_NAME))

	case "ALTERNATIVE_ID":
		if datatype_definition_date.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_DATE", fieldName)
	}
	return
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_enumeration.LONG_NAME))

	case "ALTERNATIVE_ID":
		if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPECIFIED_VALUES":
		if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPECIFIED_VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_enumeration.SPECIFIED_VALUES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPECIFIED_VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_ENUMERATION", fieldName)
	}
	return
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_integer.LONG_NAME))
	case "MAX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MAX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_integer.MAX))
	case "MIN":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MIN")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_integer.MIN))

	case "ALTERNATIVE_ID":
		if datatype_definition_integer.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_INTEGER", fieldName)
	}
	return
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.Name))
	case "ACCURACY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ACCURACY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_real.ACCURACY))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_real.LONG_NAME))
	case "MAX":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MAX")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatype_definition_real.MAX))
	case "MIN":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MIN")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", datatype_definition_real.MIN))

	case "ALTERNATIVE_ID":
		if datatype_definition_real.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_REAL", fieldName)
	}
	return
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_string.LONG_NAME))
	case "MAX_LENGTH":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MAX_LENGTH")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", datatype_definition_string.MAX_LENGTH))

	case "ALTERNATIVE_ID":
		if datatype_definition_string.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_STRING", fieldName)
	}
	return
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(datatype_definition_xhtml.LONG_NAME))

	case "ALTERNATIVE_ID":
		if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", datatype_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_XHTML", fieldName)
	}
	return
}

func (embedded_value *EMBEDDED_VALUE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", embedded_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(embedded_value.Name))
	case "KEY":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", embedded_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "KEY")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", embedded_value.KEY))
	case "OTHER_CONTENT":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", embedded_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OTHER_CONTENT")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(embedded_value.OTHER_CONTENT))

	default:
		log.Panicf("Unknown field %s for Gongstruct EMBEDDED_VALUE", fieldName)
	}
	return
}

func (enum_value *ENUM_VALUE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(enum_value.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(enum_value.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(enum_value.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(enum_value.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(enum_value.LONG_NAME))

	case "ALTERNATIVE_ID":
		if enum_value.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", enum_value.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "PROPERTIES":
		if enum_value.PROPERTIES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PROPERTIES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", enum_value.PROPERTIES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PROPERTIES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ENUM_VALUE", fieldName)
	}
	return
}

func (relation_group *RELATION_GROUP) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group.LONG_NAME))

	case "ALTERNATIVE_ID":
		if relation_group.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", relation_group.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SOURCE_SPECIFICATION":
		if relation_group.SOURCE_SPECIFICATION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SOURCE_SPECIFICATION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", relation_group.SOURCE_SPECIFICATION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SOURCE_SPECIFICATION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_RELATIONS":
		if relation_group.SPEC_RELATIONS != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", relation_group.SPEC_RELATIONS.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TARGET_SPECIFICATION":
		if relation_group.TARGET_SPECIFICATION != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TARGET_SPECIFICATION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", relation_group.TARGET_SPECIFICATION.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TARGET_SPECIFICATION")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if relation_group.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", relation_group.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RELATION_GROUP", fieldName)
	}
	return
}

func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group_type.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group_type.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group_type.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group_type.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(relation_group_type.LONG_NAME))

	case "ALTERNATIVE_ID":
		if relation_group_type.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", relation_group_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if relation_group_type.SPEC_ATTRIBUTES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", relation_group_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RELATION_GROUP_TYPE", fieldName)
	}
	return
}

func (req_if *REQ_IF) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if.Name))
	case "Lang":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Lang")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if.Lang))

	case "THE_HEADER":
		if req_if.THE_HEADER != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_HEADER")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if.THE_HEADER.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "THE_HEADER")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "CORE_CONTENT":
		if req_if.CORE_CONTENT != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CORE_CONTENT")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if.CORE_CONTENT.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CORE_CONTENT")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TOOL_EXTENSIONS":
		if req_if.TOOL_EXTENSIONS != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TOOL_EXTENSIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if.TOOL_EXTENSIONS.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TOOL_EXTENSIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF", fieldName)
	}
	return
}

func (req_if_content *REQ_IF_CONTENT) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_content.Name))

	case "DATATYPES":
		if req_if_content.DATATYPES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if_content.DATATYPES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DATATYPES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_TYPES":
		if req_if_content.SPEC_TYPES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_TYPES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if_content.SPEC_TYPES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_TYPES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_OBJECTS":
		if req_if_content.SPEC_OBJECTS != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_OBJECTS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if_content.SPEC_OBJECTS.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_OBJECTS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_RELATIONS":
		if req_if_content.SPEC_RELATIONS != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if_content.SPEC_RELATIONS.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPECIFICATIONS":
		if req_if_content.SPECIFICATIONS != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPECIFICATIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if_content.SPECIFICATIONS.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPECIFICATIONS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_RELATION_GROUPS":
		if req_if_content.SPEC_RELATION_GROUPS != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATION_GROUPS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", req_if_content.SPEC_RELATION_GROUPS.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_RELATION_GROUPS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF_CONTENT", fieldName)
	}
	return
}

func (req_if_header *REQ_IF_HEADER) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.Name))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.IDENTIFIER))
	case "COMMENT":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "COMMENT")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.COMMENT))
	case "CREATION_TIME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CREATION_TIME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.CREATION_TIME))
	case "REPOSITORY_ID":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "REPOSITORY_ID")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.REPOSITORY_ID))
	case "REQ_IF_TOOL_ID":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "REQ_IF_TOOL_ID")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.REQ_IF_TOOL_ID))
	case "REQ_IF_VERSION":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "REQ_IF_VERSION")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.REQ_IF_VERSION))
	case "SOURCE_TOOL_ID":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SOURCE_TOOL_ID")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.SOURCE_TOOL_ID))
	case "TITLE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TITLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_header.TITLE))

	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF_HEADER", fieldName)
	}
	return
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", req_if_tool_extension.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(req_if_tool_extension.Name))

	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF_TOOL_EXTENSION", fieldName)
	}
	return
}

func (specification *SPECIFICATION) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification.LONG_NAME))

	case "ALTERNATIVE_ID":
		if specification.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", specification.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "CHILDREN":
		if specification.CHILDREN != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CHILDREN")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", specification.CHILDREN.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CHILDREN")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "VALUES":
		if specification.VALUES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", specification.VALUES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if specification.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", specification.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPECIFICATION", fieldName)
	}
	return
}

func (specification_type *SPECIFICATION_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification_type.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification_type.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification_type.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification_type.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(specification_type.LONG_NAME))

	case "ALTERNATIVE_ID":
		if specification_type.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", specification_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if specification_type.SPEC_ATTRIBUTES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", specification_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPECIFICATION_TYPE", fieldName)
	}
	return
}

func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.IDENTIFIER))
	case "IS_EDITABLE":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_EDITABLE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_hierarchy.IS_EDITABLE))
	case "IS_TABLE_INTERNAL":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IS_TABLE_INTERNAL")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", spec_hierarchy.IS_TABLE_INTERNAL))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_hierarchy.LONG_NAME))

	case "ALTERNATIVE_ID":
		if spec_hierarchy.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_hierarchy.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "CHILDREN":
		if spec_hierarchy.CHILDREN != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CHILDREN")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_hierarchy.CHILDREN.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "CHILDREN")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "EDITABLE_ATTS":
		if spec_hierarchy.EDITABLE_ATTS != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EDITABLE_ATTS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_hierarchy.EDITABLE_ATTS.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EDITABLE_ATTS")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "OBJECT":
		if spec_hierarchy.OBJECT != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OBJECT")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_hierarchy.OBJECT.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OBJECT")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_HIERARCHY", fieldName)
	}
	return
}

func (spec_object *SPEC_OBJECT) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object.LONG_NAME))

	case "ALTERNATIVE_ID":
		if spec_object.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_object.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "VALUES":
		if spec_object.VALUES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_object.VALUES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if spec_object.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_object.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_OBJECT", fieldName)
	}
	return
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object_type.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object_type.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object_type.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object_type.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_object_type.LONG_NAME))

	case "ALTERNATIVE_ID":
		if spec_object_type.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_object_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if spec_object_type.SPEC_ATTRIBUTES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_object_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_OBJECT_TYPE", fieldName)
	}
	return
}

func (spec_relation *SPEC_RELATION) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation.LONG_NAME))

	case "ALTERNATIVE_ID":
		if spec_relation.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_relation.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "VALUES":
		if spec_relation.VALUES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_relation.VALUES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VALUES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SOURCE":
		if spec_relation.SOURCE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SOURCE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_relation.SOURCE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SOURCE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TARGET":
		if spec_relation.TARGET != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TARGET")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_relation.TARGET.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TARGET")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "TYPE":
		if spec_relation.TYPE != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_relation.TYPE.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "TYPE")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_RELATION", fieldName)
	}
	return
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation_type.Name))
	case "DESC":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DESC")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation_type.DESC))
	case "IDENTIFIER":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IDENTIFIER")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation_type.IDENTIFIER))
	case "LAST_CHANGE":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LAST_CHANGE")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation_type.LAST_CHANGE))
	case "LONG_NAME":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LONG_NAME")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(spec_relation_type.LONG_NAME))

	case "ALTERNATIVE_ID":
		if spec_relation_type.ALTERNATIVE_ID != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_relation_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ALTERNATIVE_ID")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if spec_relation_type.SPEC_ATTRIBUTES != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", spec_relation_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SPEC_ATTRIBUTES")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_RELATION_TYPE", fieldName)
	}
	return
}

func (xhtml_content *XHTML_CONTENT) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xhtml_content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xhtml_content.Name))
	case "EnclosedText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xhtml_content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "EnclosedText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xhtml_content.EnclosedText))
	case "PureText":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", xhtml_content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PureText")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(xhtml_content.PureText))

	default:
		log.Panicf("Unknown field %s for Gongstruct XHTML_CONTENT", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (alternative_id *ALTERNATIVE_ID) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(alternative_id.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(alternative_id.GongMarshallField(stage, "IDENTIFIER"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_boolean.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_date.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "MULTI_VALUED"))
		pointersInitializesStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_enumeration.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_integer.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_real.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_string.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "DEFAULT_VALUE"))
		pointersInitializesStatements.WriteString(attribute_definition_xhtml.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "DEFINITION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_date.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_date.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_date.GongMarshallField(stage, "DEFINITION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_enumeration.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(attribute_value_enumeration.GongMarshallField(stage, "DEFINITION"))
		pointersInitializesStatements.WriteString(attribute_value_enumeration.GongMarshallField(stage, "VALUES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "DEFINITION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_real.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_real.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_real.GongMarshallField(stage, "DEFINITION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_string.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_string.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_string.GongMarshallField(stage, "DEFINITION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "IS_SIMPLIFIED"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "THE_ORIGINAL_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "DEFINITION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_alternative_id.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_alternative_id.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_boolean_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_boolean_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_date_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_date_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_enumeration_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_enumeration_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_integer_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_integer_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_real_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_real_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_string_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_string_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_definition_xhtml_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_attribute_definition_xhtml_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_boolean.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_boolean.GongMarshallField(stage, "ATTRIBUTE_VALUE_BOOLEAN"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_date.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_date.GongMarshallField(stage, "ATTRIBUTE_VALUE_DATE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_enumeration.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_enumeration.GongMarshallField(stage, "ATTRIBUTE_VALUE_ENUMERATION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_integer.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_integer.GongMarshallField(stage, "ATTRIBUTE_VALUE_INTEGER"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_real.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_real.GongMarshallField(stage, "ATTRIBUTE_VALUE_REAL"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_string.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_string.GongMarshallField(stage, "ATTRIBUTE_VALUE_STRING"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_xhtml.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml.GongMarshallField(stage, "ATTRIBUTE_VALUE_XHTML"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_BOOLEAN"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_DATE"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_ENUMERATION"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_INTEGER"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_REAL"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_STRING"))
		pointersInitializesStatements.WriteString(a_attribute_value_xhtml_1.GongMarshallField(stage, "ATTRIBUTE_VALUE_XHTML"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_children *A_CHILDREN) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_children.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_children.GongMarshallField(stage, "SPEC_HIERARCHY"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_core_content *A_CORE_CONTENT) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_core_content.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_core_content.GongMarshallField(stage, "REQ_IF_CONTENT"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatypes *A_DATATYPES) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatypes.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_BOOLEAN"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_DATE"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_ENUMERATION"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_INTEGER"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_REAL"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_STRING"))
		pointersInitializesStatements.WriteString(a_datatypes.GongMarshallField(stage, "DATATYPE_DEFINITION_XHTML"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_boolean_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_boolean_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_BOOLEAN_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_date_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_date_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_DATE_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_enumeration_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_enumeration_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_ENUMERATION_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_integer_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_integer_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_INTEGER_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_real_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_real_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_REAL_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_string_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_string_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_STRING_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_datatype_definition_xhtml_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_datatype_definition_xhtml_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_XHTML_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING_REF"))
		initializerStatements.WriteString(a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_enum_value_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_enum_value_ref.GongMarshallField(stage, "ENUM_VALUE_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_object *A_OBJECT) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_object.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_object.GongMarshallField(stage, "SPEC_OBJECT_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_properties *A_PROPERTIES) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_properties.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_properties.GongMarshallField(stage, "EMBEDDED_VALUE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_relation_group_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_relation_group_type_ref.GongMarshallField(stage, "RELATION_GROUP_TYPE_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_source_1 *A_SOURCE_1) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_source_1.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_source_1.GongMarshallField(stage, "SPEC_OBJECT_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_source_specification_1.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_source_specification_1.GongMarshallField(stage, "SPECIFICATION_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_specifications *A_SPECIFICATIONS) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_specifications.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_specifications.GongMarshallField(stage, "SPECIFICATION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_specification_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_specification_type_ref.GongMarshallField(stage, "SPECIFICATION_TYPE_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_specified_values.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_specified_values.GongMarshallField(stage, "ENUM_VALUE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING"))
		pointersInitializesStatements.WriteString(a_spec_attributes.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_objects.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_objects.GongMarshallField(stage, "SPEC_OBJECT"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_object_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_spec_object_type_ref.GongMarshallField(stage, "SPEC_OBJECT_TYPE_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relations.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_relations.GongMarshallField(stage, "SPEC_RELATION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relation_groups.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_relation_groups.GongMarshallField(stage, "RELATION_GROUP"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relation_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_spec_relation_ref.GongMarshallField(stage, "SPEC_RELATION_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_relation_type_ref.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_spec_relation_type_ref.GongMarshallField(stage, "SPEC_RELATION_TYPE_REF"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_spec_types *A_SPEC_TYPES) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_spec_types.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "RELATION_GROUP_TYPE"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "SPEC_OBJECT_TYPE"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "SPEC_RELATION_TYPE"))
		pointersInitializesStatements.WriteString(a_spec_types.GongMarshallField(stage, "SPECIFICATION_TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_the_header *A_THE_HEADER) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_the_header.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_the_header.GongMarshallField(stage, "REQ_IF_HEADER"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_tool_extensions.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(a_tool_extensions.GongMarshallField(stage, "REQ_IF_TOOL_EXTENSION"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(datatype_definition_enumeration.GongMarshallField(stage, "SPECIFIED_VALUES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "MAX"))
		initializerStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "MIN"))
		pointersInitializesStatements.WriteString(datatype_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "ACCURACY"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "MAX"))
		initializerStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "MIN"))
		pointersInitializesStatements.WriteString(datatype_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "LONG_NAME"))
		initializerStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "MAX_LENGTH"))
		pointersInitializesStatements.WriteString(datatype_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(datatype_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (embedded_value *EMBEDDED_VALUE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(embedded_value.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(embedded_value.GongMarshallField(stage, "KEY"))
		initializerStatements.WriteString(embedded_value.GongMarshallField(stage, "OTHER_CONTENT"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (enum_value *ENUM_VALUE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(enum_value.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(enum_value.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(enum_value.GongMarshallField(stage, "PROPERTIES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (relation_group *RELATION_GROUP) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(relation_group.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "SOURCE_SPECIFICATION"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "SPEC_RELATIONS"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "TARGET_SPECIFICATION"))
		pointersInitializesStatements.WriteString(relation_group.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(relation_group_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(relation_group_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(relation_group_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (req_if *REQ_IF) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(req_if.GongMarshallField(stage, "Lang"))
		pointersInitializesStatements.WriteString(req_if.GongMarshallField(stage, "THE_HEADER"))
		pointersInitializesStatements.WriteString(req_if.GongMarshallField(stage, "CORE_CONTENT"))
		pointersInitializesStatements.WriteString(req_if.GongMarshallField(stage, "TOOL_EXTENSIONS"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (req_if_content *REQ_IF_CONTENT) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if_content.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "DATATYPES"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_TYPES"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_OBJECTS"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_RELATIONS"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPECIFICATIONS"))
		pointersInitializesStatements.WriteString(req_if_content.GongMarshallField(stage, "SPEC_RELATION_GROUPS"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (req_if_header *REQ_IF_HEADER) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "COMMENT"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "CREATION_TIME"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "REPOSITORY_ID"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "REQ_IF_TOOL_ID"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "REQ_IF_VERSION"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "SOURCE_TOOL_ID"))
		initializerStatements.WriteString(req_if_header.GongMarshallField(stage, "TITLE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(req_if_tool_extension.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specification *SPECIFICATION) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specification.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(specification.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "CHILDREN"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "VALUES"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specification_type *SPECIFICATION_TYPE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(specification_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(specification_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(specification_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "IS_EDITABLE"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "IS_TABLE_INTERNAL"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "CHILDREN"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "EDITABLE_ATTS"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "OBJECT"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spec_object *SPEC_OBJECT) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_object.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_object.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_object.GongMarshallField(stage, "VALUES"))
		pointersInitializesStatements.WriteString(spec_object.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_object_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_object_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_object_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spec_relation *SPEC_RELATION) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_relation.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "VALUES"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "SOURCE"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "TARGET"))
		pointersInitializesStatements.WriteString(spec_relation.GongMarshallField(stage, "TYPE"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "DESC"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "IDENTIFIER"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "LAST_CHANGE"))
		initializerStatements.WriteString(spec_relation_type.GongMarshallField(stage, "LONG_NAME"))
		pointersInitializesStatements.WriteString(spec_relation_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		pointersInitializesStatements.WriteString(spec_relation_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (xhtml_content *XHTML_CONTENT) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(xhtml_content.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(xhtml_content.GongMarshallField(stage, "EnclosedText"))
		initializerStatements.WriteString(xhtml_content.GongMarshallField(stage, "PureText"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
