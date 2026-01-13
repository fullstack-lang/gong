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
	identifiersDecl := ""
	initializerStatements := ""
	pointersInitializesStatements := ""

	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	chapterOrdered := []*Chapter{}
	for chapter := range stage.Chapters {
		chapterOrdered = append(chapterOrdered, chapter)
	}
	sort.Slice(chapterOrdered[:], func(i, j int) bool {
		chapteri := chapterOrdered[i]
		chapterj := chapterOrdered[j]
		chapteri_order, oki := stage.ChapterMap_Staged_Order[chapteri]
		chapterj_order, okj := stage.ChapterMap_Staged_Order[chapterj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return chapteri_order < chapterj_order
	})
	if len(chapterOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, chapter := range chapterOrdered {

		identifiersDecl += chapter.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += chapter.GongMarshallField(stage, "Name")
		initializerStatements += chapter.GongMarshallField(stage, "MardownContent")
		pointersInitializesStatements += chapter.GongMarshallField(stage, "Pages")
	}

	contentOrdered := []*Content{}
	for content := range stage.Contents {
		contentOrdered = append(contentOrdered, content)
	}
	sort.Slice(contentOrdered[:], func(i, j int) bool {
		contenti := contentOrdered[i]
		contentj := contentOrdered[j]
		contenti_order, oki := stage.ContentMap_Staged_Order[contenti]
		contentj_order, okj := stage.ContentMap_Staged_Order[contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return contenti_order < contentj_order
	})
	if len(contentOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, content := range contentOrdered {

		identifiersDecl += content.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += content.GongMarshallField(stage, "Name")
		initializerStatements += content.GongMarshallField(stage, "MardownContent")
		initializerStatements += content.GongMarshallField(stage, "ContentPath")
		initializerStatements += content.GongMarshallField(stage, "OutputPath")
		initializerStatements += content.GongMarshallField(stage, "LayoutPath")
		initializerStatements += content.GongMarshallField(stage, "StaticPath")
		initializerStatements += content.GongMarshallField(stage, "IsBespokeLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "BespokeLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "IsBespokePageTileLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "BespokePageTileLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "Target")
		pointersInitializesStatements += content.GongMarshallField(stage, "Chapters")
		initializerStatements += content.GongMarshallField(stage, "VersionInfo")
	}

	pageOrdered := []*Page{}
	for page := range stage.Pages {
		pageOrdered = append(pageOrdered, page)
	}
	sort.Slice(pageOrdered[:], func(i, j int) bool {
		pagei := pageOrdered[i]
		pagej := pageOrdered[j]
		pagei_order, oki := stage.PageMap_Staged_Order[pagei]
		pagej_order, okj := stage.PageMap_Staged_Order[pagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return pagei_order < pagej_order
	})
	if len(pageOrdered) > 0 {
		identifiersDecl += "\n"
	}
	for _, page := range pageOrdered {

		identifiersDecl += page.GongMarshallIdentifier(stage)

		initializerStatements += "\n"
		// Insertion point for basic fields value assignment
		initializerStatements += page.GongMarshallField(stage, "Name")
		initializerStatements += page.GongMarshallField(stage, "MardownContent")
	}

	// insertion initialization of objects to stage
	for _, chapter := range chapterOrdered {
		_ = chapter
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, content := range contentOrdered {
		_ = content
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, page := range pageOrdered {
		_ = page
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

// insertion point for marshall field methods
func (chapter *Chapter) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", chapter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(chapter.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", chapter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(chapter.MardownContent))

	case "Pages":
		for _, _page := range chapter.Pages {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", chapter.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Pages")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _page.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Chapter", fieldName)
	}
	return
}

func (content *Content) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.MardownContent))
	case "ContentPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ContentPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.ContentPath))
	case "OutputPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OutputPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.OutputPath))
	case "LayoutPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LayoutPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.LayoutPath))
	case "StaticPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StaticPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.StaticPath))
	case "IsBespokeLogoFileName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBespokeLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", content.IsBespokeLogoFileName))
	case "BespokeLogoFileName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokeLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.BespokeLogoFileName))
	case "IsBespokePageTileLogoFileName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBespokePageTileLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", content.IsBespokePageTileLogoFileName))
	case "BespokePageTileLogoFileName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokePageTileLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.BespokePageTileLogoFileName))
	case "Target":
		if content.Target.ToCodeString() != "" {
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Target")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "models."+content.Target.ToCodeString())
		} else {
			// in case of empty enum, we need to unstage the previous value
			res = StringEnumInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Target")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "\"\"")
		}
	case "VersionInfo":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "VersionInfo")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(content.VersionInfo))

	case "Chapters":
		for _, _chapter := range content.Chapters {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", content.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Chapters")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _chapter.GongGetIdentifier(stage))
			res += tmp
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Content", fieldName)
	}
	return
}

func (page *Page) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", page.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(page.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", page.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", string(page.MardownContent))

	default:
		log.Panicf("Unknown field %s for Gongstruct Page", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (chapter *Chapter) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += chapter.GongMarshallField(stage, "Name")
		initializerStatements += chapter.GongMarshallField(stage, "MardownContent")
		pointersInitializesStatements += chapter.GongMarshallField(stage, "Pages")
	}
	return
}
func (content *Content) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += content.GongMarshallField(stage, "Name")
		initializerStatements += content.GongMarshallField(stage, "MardownContent")
		initializerStatements += content.GongMarshallField(stage, "ContentPath")
		initializerStatements += content.GongMarshallField(stage, "OutputPath")
		initializerStatements += content.GongMarshallField(stage, "LayoutPath")
		initializerStatements += content.GongMarshallField(stage, "StaticPath")
		initializerStatements += content.GongMarshallField(stage, "IsBespokeLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "BespokeLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "IsBespokePageTileLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "BespokePageTileLogoFileName")
		initializerStatements += content.GongMarshallField(stage, "Target")
		pointersInitializesStatements += content.GongMarshallField(stage, "Chapters")
		initializerStatements += content.GongMarshallField(stage, "VersionInfo")
	}
	return
}
func (page *Page) GongMarshallAllFields(stage *Stage) (initializerStatements string, pointersInitializesStatements string) {

	{ // Insertion point for basic fields value assignment
		initializerStatements += page.GongMarshallField(stage, "Name")
		initializerStatements += page.GongMarshallField(stage, "MardownContent")
	}
	return
}
