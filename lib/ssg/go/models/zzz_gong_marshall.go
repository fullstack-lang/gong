// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
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

	const __write__local_time = "{{LocalTimeStamp}}"
	const __write__utc_time__ = "{{UTCTimeStamp}}"

	const __commitId__ = "{{CommitId}}"

	// Declaration of instances to stage{{Identifiers}}

	// Setup of values{{ValueInitializers}}

	// Setup of pointers{{PointersInitializers}}
}
`

const IdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{}).Stage(stage)`

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

	if stage.generatesDiff {
		diff := computeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.delta", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
		diff = ComputeDiff(stage.contentWhenParsed, res)
		os.WriteFile(fmt.Sprintf("%s-%.10d-%.10d.diff", name, stage.commitIdWhenParsed, stage.commitId), []byte(diff), os.FileMode(0666))
	}
	stage.contentWhenParsed = res
	stage.commitIdWhenParsed = stage.commitId

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

	id := ""
	_ = id
	decl := ""
	_ = decl
	setValueField := ""
	_ = setValueField

	// insertion initialization of objects to stage
	map_Chapter_Identifiers := make(map[*Chapter]string)
	_ = map_Chapter_Identifiers

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
	for idx, chapter := range chapterOrdered {

		id = generatesIdentifier("Chapter", idx, chapter.Name)
		map_Chapter_Identifiers[chapter] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Chapter")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", chapter.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(chapter.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MardownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(chapter.MardownContent))
		initializerStatements += setValueField

	}

	map_Content_Identifiers := make(map[*Content]string)
	_ = map_Content_Identifiers

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
	for idx, content := range contentOrdered {

		id = generatesIdentifier("Content", idx, content.Name)
		map_Content_Identifiers[content] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Content")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", content.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MardownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.MardownContent))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "ContentPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.ContentPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "OutputPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.OutputPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "LayoutPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.LayoutPath))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "StaticPath")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.StaticPath))
		initializerStatements += setValueField

		if content.Target != "" {
			setValueField = StringEnumInitStatement
			setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Target")
			setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", "models."+content.Target.ToCodeString())
			initializerStatements += setValueField
		}

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "VersionInfo")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(content.VersionInfo))
		initializerStatements += setValueField

	}

	map_Page_Identifiers := make(map[*Page]string)
	_ = map_Page_Identifiers

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
	for idx, page := range pageOrdered {

		id = generatesIdentifier("Page", idx, page.Name)
		map_Page_Identifiers[page] = id

		decl = IdentifiersDecls
		decl = strings.ReplaceAll(decl, "{{Identifier}}", id)
		decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page")
		decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", page.Name)
		identifiersDecl += decl

		initializerStatements += "\n"
		// Initialisation of values
		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "Name")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(page.Name))
		initializerStatements += setValueField

		setValueField = StringInitStatement
		setValueField = strings.ReplaceAll(setValueField, "{{Identifier}}", id)
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldName}}", "MardownContent")
		setValueField = strings.ReplaceAll(setValueField, "{{GeneratedFieldNameValue}}", string(page.MardownContent))
		initializerStatements += setValueField

	}

	// insertion initialization of objects to stage
	if len(chapterOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Chapter instances pointers"
	}
	for idx, chapter := range chapterOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Chapter", idx, chapter.Name)
		map_Chapter_Identifiers[chapter] = id

		// Initialisation of values
		for _, _page := range chapter.Pages {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Pages")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Page_Identifiers[_page])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(contentOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Content instances pointers"
	}
	for idx, content := range contentOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Content", idx, content.Name)
		map_Content_Identifiers[content] = id

		// Initialisation of values
		for _, _chapter := range content.Chapters {
			setPointerField = SliceOfPointersFieldInitStatement
			setPointerField = strings.ReplaceAll(setPointerField, "{{Identifier}}", id)
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldName}}", "Chapters")
			setPointerField = strings.ReplaceAll(setPointerField, "{{GeneratedFieldNameValue}}", map_Chapter_Identifiers[_chapter])
			pointersInitializesStatements += setPointerField
		}

	}

	if len(pageOrdered) > 0 {
		pointersInitializesStatements += "\n\t// setup of Page instances pointers"
	}
	for idx, page := range pageOrdered {
		var setPointerField string
		_ = setPointerField

		id = generatesIdentifier("Page", idx, page.Name)
		map_Page_Identifiers[page] = id

		// Initialisation of values
	}

	res = strings.ReplaceAll(res, "{{Identifiers}}", identifiersDecl)
	res = strings.ReplaceAll(res, "{{ValueInitializers}}", initializerStatements)
	res = strings.ReplaceAll(res, "{{PointersInitializers}}", pointersInitializesStatements)

	// Local time with timezone
	localTimestamp := stage.commitTimeStamp.Format("2006-01-02 15:04:05.000000 MST")

	// UTC time
	utcTimestamp := stage.commitTimeStamp.UTC().Format("2006-01-02 15:04:05.000000 UTC")
	res = strings.ReplaceAll(res, "{{LocalTimeStamp}}", localTimestamp)
	res = strings.ReplaceAll(res, "{{UTCTimeStamp}}", utcTimestamp)
	res = strings.ReplaceAll(res, "{{CommitId}}", fmt.Sprintf("%.10d", stage.commitId))

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

// computeDiff calculates the git-style unified diff between two strings.
func computeDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffToDelta(diffs)
}

// computePrettyDiff calculates the git-style unified diff between two strings.
func computePrettyDiff(a, b string) string {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(a, b, false)
	return dmp.DiffPrettyHtml(diffs)
}

// applyDiff reconstructs the original string 'a' from the new string 'b' and the diff string 'c'.
func applyDiff(b, c string) (string, error) {
	dmp := diffmatchpatch.New()
	diffs, err := dmp.DiffFromDelta(b, c)
	if err != nil {
		return "", err
	}
	patches := dmp.PatchMake(b, diffs)
	// We are applying the patch in reverse to get from 'b' to 'a'.
	// The library's PatchApply function returns the new string and a slice of booleans indicating the success of each patch application.
	result, _ := dmp.PatchApply(patches, b)
	return result, nil
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
