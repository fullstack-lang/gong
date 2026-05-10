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
var (
	_ time.Time
	_ = slices.Index[[]int, int]
)

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
				content = content[:firstBrace] + "\n}\n"
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
	chapterOrdered := []*Chapter{}
	for chapter := range stage.Chapters {
		chapterOrdered = append(chapterOrdered, chapter)
	}
	sort.Slice(chapterOrdered[:], func(i, j int) bool {
		chapteri := chapterOrdered[i]
		chapterj := chapterOrdered[j]
		chapteri_order, oki := stage.Chapter_stagedOrder[chapteri]
		chapterj_order, okj := stage.Chapter_stagedOrder[chapterj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return chapteri_order < chapterj_order
	})
	if len(chapterOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, chapter := range chapterOrdered {

		identifiersDecl.WriteString(chapter.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "MardownContent"))
		pointersInitializesStatements.WriteString(chapter.GongMarshallField(stage, "Sections"))
		pointersInitializesStatements.WriteString(chapter.GongMarshallField(stage, "Pages"))
		pointersInitializesStatements.WriteString(chapter.GongMarshallField(stage, "SubChapters"))
	}

	contentOrdered := []*Content{}
	for content := range stage.Contents {
		contentOrdered = append(contentOrdered, content)
	}
	sort.Slice(contentOrdered[:], func(i, j int) bool {
		contenti := contentOrdered[i]
		contentj := contentOrdered[j]
		contenti_order, oki := stage.Content_stagedOrder[contenti]
		contentj_order, okj := stage.Content_stagedOrder[contentj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return contenti_order < contentj_order
	})
	if len(contentOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, content := range contentOrdered {

		identifiersDecl.WriteString(content.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(content.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "MardownContent"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "ContentPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "OutputPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "StaticPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "IsBespokeLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "BespokeLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "IsBespokePageTileLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "BespokePageTileLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "Target"))
		pointersInitializesStatements.WriteString(content.GongMarshallField(stage, "Chapters"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "VersionInfo"))
	}

	downloadablefileOrdered := []*DownloadableFile{}
	for downloadablefile := range stage.DownloadableFiles {
		downloadablefileOrdered = append(downloadablefileOrdered, downloadablefile)
	}
	sort.Slice(downloadablefileOrdered[:], func(i, j int) bool {
		downloadablefilei := downloadablefileOrdered[i]
		downloadablefilej := downloadablefileOrdered[j]
		downloadablefilei_order, oki := stage.DownloadableFile_stagedOrder[downloadablefilei]
		downloadablefilej_order, okj := stage.DownloadableFile_stagedOrder[downloadablefilej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return downloadablefilei_order < downloadablefilej_order
	})
	if len(downloadablefileOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, downloadablefile := range downloadablefileOrdered {

		identifiersDecl.WriteString(downloadablefile.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(downloadablefile.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(downloadablefile.GongMarshallField(stage, "Base64Content"))
	}

	jpgimageOrdered := []*JpgImage{}
	for jpgimage := range stage.JpgImages {
		jpgimageOrdered = append(jpgimageOrdered, jpgimage)
	}
	sort.Slice(jpgimageOrdered[:], func(i, j int) bool {
		jpgimagei := jpgimageOrdered[i]
		jpgimagej := jpgimageOrdered[j]
		jpgimagei_order, oki := stage.JpgImage_stagedOrder[jpgimagei]
		jpgimagej_order, okj := stage.JpgImage_stagedOrder[jpgimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return jpgimagei_order < jpgimagej_order
	})
	if len(jpgimageOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, jpgimage := range jpgimageOrdered {

		identifiersDecl.WriteString(jpgimage.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(jpgimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(jpgimage.GongMarshallField(stage, "Base64Content"))
	}

	pageOrdered := []*Page{}
	for page := range stage.Pages {
		pageOrdered = append(pageOrdered, page)
	}
	sort.Slice(pageOrdered[:], func(i, j int) bool {
		pagei := pageOrdered[i]
		pagej := pageOrdered[j]
		pagei_order, oki := stage.Page_stagedOrder[pagei]
		pagej_order, okj := stage.Page_stagedOrder[pagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return pagei_order < pagej_order
	})
	if len(pageOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, page := range pageOrdered {

		identifiersDecl.WriteString(page.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(page.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(page.GongMarshallField(stage, "MardownContent"))
		pointersInitializesStatements.WriteString(page.GongMarshallField(stage, "Sections"))
	}

	pngimageOrdered := []*PngImage{}
	for pngimage := range stage.PngImages {
		pngimageOrdered = append(pngimageOrdered, pngimage)
	}
	sort.Slice(pngimageOrdered[:], func(i, j int) bool {
		pngimagei := pngimageOrdered[i]
		pngimagej := pngimageOrdered[j]
		pngimagei_order, oki := stage.PngImage_stagedOrder[pngimagei]
		pngimagej_order, okj := stage.PngImage_stagedOrder[pngimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return pngimagei_order < pngimagej_order
	})
	if len(pngimageOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, pngimage := range pngimageOrdered {

		identifiersDecl.WriteString(pngimage.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(pngimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pngimage.GongMarshallField(stage, "Base64Content"))
	}

	sectionOrdered := []*Section{}
	for section := range stage.Sections {
		sectionOrdered = append(sectionOrdered, section)
	}
	sort.Slice(sectionOrdered[:], func(i, j int) bool {
		sectioni := sectionOrdered[i]
		sectionj := sectionOrdered[j]
		sectioni_order, oki := stage.Section_stagedOrder[sectioni]
		sectionj_order, okj := stage.Section_stagedOrder[sectionj]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return sectioni_order < sectionj_order
	})
	if len(sectionOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, section := range sectionOrdered {

		identifiersDecl.WriteString(section.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(section.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(section.GongMarshallField(stage, "MardownContent"))
		initializerStatements.WriteString(section.GongMarshallField(stage, "IsImage"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "SvgImage"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "PngImage"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "JpgImage"))
		initializerStatements.WriteString(section.GongMarshallField(stage, "IsDownloadableFile"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "DownloadableFile"))
	}

	svgimageOrdered := []*SvgImage{}
	for svgimage := range stage.SvgImages {
		svgimageOrdered = append(svgimageOrdered, svgimage)
	}
	sort.Slice(svgimageOrdered[:], func(i, j int) bool {
		svgimagei := svgimageOrdered[i]
		svgimagej := svgimageOrdered[j]
		svgimagei_order, oki := stage.SvgImage_stagedOrder[svgimagei]
		svgimagej_order, okj := stage.SvgImage_stagedOrder[svgimagej]
		if !oki || !okj {
			log.Fatalln("unknown pointers")
		}
		return svgimagei_order < svgimagej_order
	})
	if len(svgimageOrdered) > 0 {
		identifiersDecl.WriteString("\n")
	}
	for _, svgimage := range svgimageOrdered {

		identifiersDecl.WriteString(svgimage.GongMarshallIdentifier(stage))

		initializerStatements.WriteString("\n")
		// Insertion point for basic fields value assignment
		initializerStatements.WriteString(svgimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svgimage.GongMarshallField(stage, "Content"))
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

	for _, downloadablefile := range downloadablefileOrdered {
		_ = downloadablefile
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, jpgimage := range jpgimageOrdered {
		_ = jpgimage
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

	for _, pngimage := range pngimageOrdered {
		_ = pngimage
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, section := range sectionOrdered {
		_ = section
		var setPointerField string
		_ = setPointerField

		// Insertion point for pointers initialization
	}

	for _, svgimage := range svgimageOrdered {
		_ = svgimage
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
func (chapter *Chapter) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", chapter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(chapter.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", chapter.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(chapter.MardownContent))

	case "Sections":
		var sb strings.Builder
		for _, _section := range chapter.Sections {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", chapter.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Sections")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _section.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "Pages":
		var sb strings.Builder
		for _, _page := range chapter.Pages {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", chapter.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Pages")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _page.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	case "SubChapters":
		var sb strings.Builder
		for _, _chapter := range chapter.SubChapters {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", chapter.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "SubChapters")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _chapter.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.MardownContent))
	case "ContentPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "ContentPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.ContentPath))
	case "OutputPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "OutputPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.OutputPath))
	case "StaticPath":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "StaticPath")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.StaticPath))
	case "LogoSVGFile":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "LogoSVGFile")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.LogoSVGFile))
	case "IsBespokeLogoFileName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBespokeLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", content.IsBespokeLogoFileName))
	case "BespokeLogoFileName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokeLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.BespokeLogoFileName))
	case "IsBespokePageTileLogoFileName":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsBespokePageTileLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", content.IsBespokePageTileLogoFileName))
	case "BespokePageTileLogoFileName":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", content.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "BespokePageTileLogoFileName")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.BespokePageTileLogoFileName))
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
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(content.VersionInfo))

	case "Chapters":
		var sb strings.Builder
		for _, _chapter := range content.Chapters {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", content.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Chapters")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _chapter.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Content", fieldName)
	}
	return
}

func (downloadablefile *DownloadableFile) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", downloadablefile.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(downloadablefile.Name))
	case "Base64Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", downloadablefile.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Base64Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(downloadablefile.Base64Content))

	default:
		log.Panicf("Unknown field %s for Gongstruct DownloadableFile", fieldName)
	}
	return
}

func (jpgimage *JpgImage) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(jpgimage.Name))
	case "Base64Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", jpgimage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Base64Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(jpgimage.Base64Content))

	default:
		log.Panicf("Unknown field %s for Gongstruct JpgImage", fieldName)
	}
	return
}

func (page *Page) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", page.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(page.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", page.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(page.MardownContent))

	case "Sections":
		var sb strings.Builder
		for _, _section := range page.Sections {
			tmp := SliceOfPointersFieldInitStatement
			tmp = strings.ReplaceAll(tmp, "{{Identifier}}", page.GongGetIdentifier(stage))
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldName}}", "Sections")
			tmp = strings.ReplaceAll(tmp, "{{GeneratedFieldNameValue}}", _section.GongGetIdentifier(stage))
			sb.WriteString(tmp)
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Page", fieldName)
	}
	return
}

func (pngimage *PngImage) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(pngimage.Name))
	case "Base64Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", pngimage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Base64Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(pngimage.Base64Content))

	default:
		log.Panicf("Unknown field %s for Gongstruct PngImage", fieldName)
	}
	return
}

func (section *Section) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(section.Name))
	case "MardownContent":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "MardownContent")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(section.MardownContent))
	case "IsImage":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsImage")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", section.IsImage))
	case "IsDownloadableFile":
		res = NumberInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "IsDownloadableFile")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", section.IsDownloadableFile))

	case "SvgImage":
		if section.SvgImage != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SvgImage")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", section.SvgImage.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "SvgImage")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "PngImage":
		if section.PngImage != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PngImage")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", section.PngImage.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "PngImage")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "JpgImage":
		if section.JpgImage != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "JpgImage")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", section.JpgImage.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "JpgImage")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	case "DownloadableFile":
		if section.DownloadableFile != nil {
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DownloadableFile")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", section.DownloadableFile.GongGetIdentifier(stage))
		} else {
			// in case of nil pointer, we need to unstage the previous value
			res = PointerFieldInitStatement
			res = strings.ReplaceAll(res, "{{Identifier}}", section.GongGetIdentifier(stage))
			res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "DownloadableFile")
			res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Section", fieldName)
	}
	return
}

func (svgimage *SvgImage) GongMarshallField(stage *Stage, fieldName string) (res string) {

	switch fieldName {
	case "Name":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Name")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(svgimage.Name))
	case "Content":
		res = StringInitStatement
		res = strings.ReplaceAll(res, "{{Identifier}}", svgimage.GongGetIdentifier(stage))
		res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", "Content")
		res = strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(svgimage.Content))

	default:
		log.Panicf("Unknown field %s for Gongstruct SvgImage", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (chapter *Chapter) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(chapter.GongMarshallField(stage, "MardownContent"))
		pointersInitializesStatements.WriteString(chapter.GongMarshallField(stage, "Sections"))
		pointersInitializesStatements.WriteString(chapter.GongMarshallField(stage, "Pages"))
		pointersInitializesStatements.WriteString(chapter.GongMarshallField(stage, "SubChapters"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (content *Content) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(content.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "MardownContent"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "ContentPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "OutputPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "StaticPath"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "LogoSVGFile"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "IsBespokeLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "BespokeLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "IsBespokePageTileLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "BespokePageTileLogoFileName"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "Target"))
		pointersInitializesStatements.WriteString(content.GongMarshallField(stage, "Chapters"))
		initializerStatements.WriteString(content.GongMarshallField(stage, "VersionInfo"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (downloadablefile *DownloadableFile) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(downloadablefile.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(downloadablefile.GongMarshallField(stage, "Base64Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (jpgimage *JpgImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(jpgimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(jpgimage.GongMarshallField(stage, "Base64Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (page *Page) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(page.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(page.GongMarshallField(stage, "MardownContent"))
		pointersInitializesStatements.WriteString(page.GongMarshallField(stage, "Sections"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pngimage *PngImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pngimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pngimage.GongMarshallField(stage, "Base64Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (section *Section) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(section.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(section.GongMarshallField(stage, "MardownContent"))
		initializerStatements.WriteString(section.GongMarshallField(stage, "IsImage"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "SvgImage"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "PngImage"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "JpgImage"))
		initializerStatements.WriteString(section.GongMarshallField(stage, "IsDownloadableFile"))
		pointersInitializesStatements.WriteString(section.GongMarshallField(stage, "DownloadableFile"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (svgimage *SvgImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(svgimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(svgimage.GongMarshallField(stage, "Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
