// generated code - do not edit
package models

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"sort"
	"strings"
	"time"
)

var (
	_ = sort.Slice
	_ = time.Hour
)

const GongIdentifiersDecls = `
	{{Identifier}} := (&models.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

const GongUnstageStmt = `
	{{Identifier}}.Unstage(stage)`


const GongStringInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongMetaFieldStructInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = ` + `{{GeneratedFieldNameValue}}`

const GongStringEnumInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongNumberInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongPointerFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = {{GeneratedFieldNameValue}}`

const GongSliceOfPointersFieldInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}} = append({{Identifier}}.{{GeneratedFieldName}}, {{GeneratedFieldNameValue}})`

const GongTimeInitStatement = `
	{{Identifier}}.{{GeneratedFieldName}}, _ = time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", "{{GeneratedFieldNameValue}}")`

// __gong__toRawStringLiteral formats a string into safe Go source code,
// using backticks to preserve newlines and readability.
func __gong__toRawStringLiteral(s string) string {
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

func __gong__marshallString[T ~string](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongStringInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(string(val)))
}

func __gong__marshallInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%d", val))
}

func __gong__marshallBool[T ~bool](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%t", val))
}

func __gong__marshallFloat[T ~float32 | ~float64](ident, fieldName string, val T) string {
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", fmt.Sprintf("%f", val))
}

func __gong__marshallTime(ident, fieldName, valStr string) string {
	res := strings.ReplaceAll(GongTimeInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", valStr)
}

func __gong__marshallPointer(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongPointerFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallSliceOfPointers(ident, fieldName, targetIdent string) string {
	res := strings.ReplaceAll(GongSliceOfPointersFieldInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", targetIdent)
}

func __gong__marshallEnumString(ident, fieldName, codeStr string) string {
	val := "\"\""
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongStringEnumInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallEnumInt(ident, fieldName, codeStr string) string {
	val := "0"
	if codeStr != "" {
		val = "models." + codeStr
	}
	res := strings.ReplaceAll(GongNumberInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

func __gong__marshallMeta(ident, fieldName, val string) string {
	res := strings.ReplaceAll(GongMetaFieldStructInitStatement, "{{Identifier}}", ident)
	res = strings.ReplaceAll(res, "{{GeneratedFieldName}}", fieldName)
	return strings.ReplaceAll(res, "{{GeneratedFieldNameValue}}", val)
}

// __gong__sortStageSetInstances sorts instances by their staged order
func __gong__sortStageSetInstances[T comparable](instances map[T]struct{}, orderMap map[T]uint) []T {
	ordered := make([]T, 0, len(instances))
	for inst := range instances {
		ordered = append(ordered, inst)
	}
	sort.Slice(ordered, func(i, j int) bool {
		return orderMap[ordered[i]] < orderMap[ordered[j]]
	})
	return ordered
}

// MarshallFile marshall the stage content into a file as an instanciation into a stage
// according to the marshalling policy of the stage.
//
// In GongMarshallingAppendCommit mode, it will append the last commit to the file.
// In other modes, it will rewrite the entire file.
func (stage *Stage) MarshallFile(filename string, args ...string) {

	if stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		contentBytes, err := os.ReadFile(filename)

		// if the file does not exist, marshall the full stage
		if os.IsNotExist(err) {
			file, createErr := os.Create(filename)
			if createErr != nil {
				log.Fatal(createErr.Error())
			}
			defer file.Close()
			stage.Marshall(file, args...)
			return
		}
		if err != nil {
			log.Fatal(err.Error())
		}

		content := string(contentBytes)

		if stage.IsSquashing() {
			// we squash: we want to clear the current function body
			// and let the append logic write the squashed commit
			const fnSig = "func _(stage *models.Stage) {"
			if before, _, ok := strings.Cut(content, fnSig); ok {
				content = before + fnSig + "\n}\n"
			}
		}

		if stage.IsApplyingBackwardCommit() {
			// we are going backward, we need to remove the last forward commit from the file

			// because commitsBehind has been incremented before the call to this function
			// the index of the commit to remove is len(forwardCommits) - commitsBehind
			forwardCommits := stage.GetForwardCommits()
			commitIndexToRemove := len(forwardCommits) - stage.GetCommitsBehind()

			if commitIndexToRemove < 0 || commitIndexToRemove >= len(forwardCommits) {
				return // Should not happen if history is consistent
			}

			commitToRemove := forwardCommits[commitIndexToRemove]

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
					stage.Marshall(file, args...)
				}
			}
			return // we are done for the backward case
		}

		if stage.IsApplyingForwardCommit() {
			// bypass the modified check
		} else if !stage.IsModified() {
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
			stage.Marshall(file, args...)
			return
		}

		contentBeforeBrace := content[:lastBrace]
		trimmedContentBeforeBrace := strings.TrimSpace(contentBeforeBrace)
		emptyBody := stage.IsSquashing() ||
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *models.Stage) {") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// insertion point for setup of pointers") ||
			strings.HasSuffix(trimmedContentBeforeBrace, "// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)\n\t// ------------------------------------------------------------------------")

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

		stage.Marshall(file, args...)
	}
}

// Marshall marshall the stage content into the file as an instanciation into a stage
func (stage *Stage) Marshall(file *os.File, args ...string) {

	name := file.Name()

	if !strings.HasSuffix(name, ".go") {
		log.Println(name + " is not a go filename")
	}

	log.Printf("Marshalling %s", name)

	res, err := stage.MarshallToString(args...)
	if err != nil {
		log.Fatalln("Error marshalling to string:", err)
	}

	fmt.Fprintln(file, res)
}

// MarshallToString marshall the stage content into a string
func (stage *Stage) MarshallToString(args ...string) (res string, err error) {
	packageName := "main"
	modelsPackageName := "github.com/fullstack-lang/gong/lib/ssg/go/models"
	if len(args) == 1 {
		packageName = args[0]
	} else if len(args) >= 2 {
		if args[0] != "" {
			modelsPackageName = args[0]
		}
		packageName = args[1]
	}

	var declarations strings.Builder
	var values strings.Builder
	var pointers strings.Builder
	var lastStageDecl string
	var lastStageVal string
	var lastStagePtr string
	_ = lastStageDecl
	_ = lastStageVal
	_ = lastStagePtr
	localAliasPrefix := "__models"
	if stage.IsInDeltaMode() || stage.GetGongMarshallingMode() == GongMarshallingAppendCommit {
		localAliasPrefix = ""
	}
	_ = localAliasPrefix

	if stage != nil {
		for _, chapter := range __gong__sortStageSetInstances(stage.Chapters, stage.Chapter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			chapterIdent := localAliasPrefix + chapter.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Chapter{Name: %s}).Stage(stage)", chapterIdent, __gong__toRawStringLiteral(chapter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", chapterIdent, __gong__toRawStringLiteral(chapter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", chapterIdent, __gong__toRawStringLiteral(chapter.MardownContent)))
			for _, elem := range chapter.Sections {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sections = append(%s.Sections, %s)", chapterIdent, chapterIdent, targetIdent))
			}
			for _, elem := range chapter.Pages {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Pages = append(%s.Pages, %s)", chapterIdent, chapterIdent, targetIdent))
			}
			for _, elem := range chapter.SubChapters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SubChapters = append(%s.SubChapters, %s)", chapterIdent, chapterIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, content := range __gong__sortStageSetInstances(stage.Contents, stage.Content_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			contentIdent := localAliasPrefix + content.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Content{Name: %s}).Stage(stage)", contentIdent, __gong__toRawStringLiteral(content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", contentIdent, __gong__toRawStringLiteral(content.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", contentIdent, __gong__toRawStringLiteral(content.MardownContent)))
			values.WriteString(fmt.Sprintf("\n\t%s.ContentPath = %s", contentIdent, __gong__toRawStringLiteral(content.ContentPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.OutputPath = %s", contentIdent, __gong__toRawStringLiteral(content.OutputPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.StaticPath = %s", contentIdent, __gong__toRawStringLiteral(content.StaticPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.LogoSVGFile = %s", contentIdent, __gong__toRawStringLiteral(content.LogoSVGFile)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokeLogoFileName = %t", contentIdent, content.IsBespokeLogoFileName))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokeLogoFileName = %s", contentIdent, __gong__toRawStringLiteral(content.BespokeLogoFileName)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBespokePageTileLogoFileName = %t", contentIdent, content.IsBespokePageTileLogoFileName))
			values.WriteString(fmt.Sprintf("\n\t%s.BespokePageTileLogoFileName = %s", contentIdent, __gong__toRawStringLiteral(content.BespokePageTileLogoFileName)))
			if content.Target.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Target = models.%s", contentIdent, content.Target.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Target = \"\"", contentIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.VersionInfo = %s", contentIdent, __gong__toRawStringLiteral(content.VersionInfo)))
			for _, elem := range content.Chapters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Chapters = append(%s.Chapters, %s)", contentIdent, contentIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, downloadablefile := range __gong__sortStageSetInstances(stage.DownloadableFiles, stage.DownloadableFile_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			downloadablefileIdent := localAliasPrefix + downloadablefile.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DownloadableFile{Name: %s}).Stage(stage)", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", downloadablefileIdent, __gong__toRawStringLiteral(downloadablefile.Base64Content)))
		}
	}
	if stage != nil {
		for _, jpgimage := range __gong__sortStageSetInstances(stage.JpgImages, stage.JpgImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			jpgimageIdent := localAliasPrefix + jpgimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.JpgImage{Name: %s}).Stage(stage)", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", jpgimageIdent, __gong__toRawStringLiteral(jpgimage.Base64Content)))
		}
	}
	if stage != nil {
		for _, page := range __gong__sortStageSetInstances(stage.Pages, stage.Page_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pageIdent := localAliasPrefix + page.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Page{Name: %s}).Stage(stage)", pageIdent, __gong__toRawStringLiteral(page.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pageIdent, __gong__toRawStringLiteral(page.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", pageIdent, __gong__toRawStringLiteral(page.MardownContent)))
			for _, elem := range page.Sections {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Sections = append(%s.Sections, %s)", pageIdent, pageIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, pngimage := range __gong__sortStageSetInstances(stage.PngImages, stage.PngImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pngimageIdent := localAliasPrefix + pngimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.PngImage{Name: %s}).Stage(stage)", pngimageIdent, __gong__toRawStringLiteral(pngimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pngimageIdent, __gong__toRawStringLiteral(pngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", pngimageIdent, __gong__toRawStringLiteral(pngimage.Base64Content)))
		}
	}
	if stage != nil {
		for _, section := range __gong__sortStageSetInstances(stage.Sections, stage.Section_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			sectionIdent := localAliasPrefix + section.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Section{Name: %s}).Stage(stage)", sectionIdent, __gong__toRawStringLiteral(section.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", sectionIdent, __gong__toRawStringLiteral(section.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MardownContent = %s", sectionIdent, __gong__toRawStringLiteral(section.MardownContent)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsImage = %t", sectionIdent, section.IsImage))
			values.WriteString(fmt.Sprintf("\n\t%s.IsDownloadableFile = %t", sectionIdent, section.IsDownloadableFile))
			if section.SvgImage != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + section.SvgImage.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SvgImage = %s", sectionIdent, targetIdent))
			}
			if section.PngImage != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + section.PngImage.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PngImage = %s", sectionIdent, targetIdent))
			}
			if section.JpgImage != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + section.JpgImage.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.JpgImage = %s", sectionIdent, targetIdent))
			}
			if section.DownloadableFile != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + section.DownloadableFile.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DownloadableFile = %s", sectionIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, svgimage := range __gong__sortStageSetInstances(stage.SvgImages, stage.SvgImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			svgimageIdent := localAliasPrefix + svgimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SvgImage{Name: %s}).Stage(stage)", svgimageIdent, __gong__toRawStringLiteral(svgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", svgimageIdent, __gong__toRawStringLiteral(svgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", svgimageIdent, __gong__toRawStringLiteral(svgimage.Content)))
		}
	}

	var goModuleVersion string
	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		goModuleVersion = buildInfo.Main.Version
	}

	goModuleVersionWithoutDirty := strings.ReplaceAll(goModuleVersion, "+dirty", "")
	if goModuleVersionWithoutDirty == "" || goModuleVersionWithoutDirty == "(devel)" {
		goModuleVersionWithoutDirty = "latest"
	}

	var metaImports strings.Builder
	var metaDummies strings.Builder
	if len(stage.MetaPackageImports) > 0 {
		for _, imp := range stage.MetaPackageImports {
			metaImports.WriteString(fmt.Sprintf("\n\t%s %s", imp.Alias, imp.Path))
			metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", imp.Alias))
		}
	} else if stage.MetaPackageImportAlias != "" {
		metaImports.WriteString(fmt.Sprintf("\n\t%s %s", stage.MetaPackageImportAlias, stage.MetaPackageImportPath))
		metaDummies.WriteString(fmt.Sprintf("\nvar _ %s.Stage", stage.MetaPackageImportAlias))
	}

	res = fmt.Sprintf(`// file generated by gong
// go module version: %s
/* to edit this file with the default probe, use the command
go run github.com/fullstack-lang/gong/lib/ssg/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/ssg/go/models"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *models.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *models.Stage) {

	// ------------------------------------------------------------------------
	// Phase 1: Declarations (in topological order: leaves first)
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 2: Value Initializations
	// ------------------------------------------------------------------------%s

	// ------------------------------------------------------------------------
	// Phase 3: Pointer Setups (Intra-stage and Cross-stage pointers)
	// ------------------------------------------------------------------------%s
}
`, goModuleVersion, goModuleVersionWithoutDirty, packageName, metaImports.String(), metaDummies.String(), declarations.String(), values.String(), pointers.String())

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/ssg/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/ssg/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (chapter *Chapter) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := chapter.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", chapter.Name)
	case "MardownContent":
		res = __gong__marshallString(ident, "MardownContent", chapter.MardownContent)

	case "Sections":
		var sb strings.Builder
		for _, _section := range chapter.Sections {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sections", _section.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Pages":
		var sb strings.Builder
		for _, _page := range chapter.Pages {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Pages", _page.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SubChapters":
		var sb strings.Builder
		for _, _chapter := range chapter.SubChapters {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SubChapters", _chapter.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Chapter", fieldName)
	}
	return
}

func (content *Content) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := content.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", content.Name)
	case "MardownContent":
		res = __gong__marshallString(ident, "MardownContent", content.MardownContent)
	case "ContentPath":
		res = __gong__marshallString(ident, "ContentPath", content.ContentPath)
	case "OutputPath":
		res = __gong__marshallString(ident, "OutputPath", content.OutputPath)
	case "StaticPath":
		res = __gong__marshallString(ident, "StaticPath", content.StaticPath)
	case "LogoSVGFile":
		res = __gong__marshallString(ident, "LogoSVGFile", content.LogoSVGFile)
	case "IsBespokeLogoFileName":
		res = __gong__marshallBool(ident, "IsBespokeLogoFileName", content.IsBespokeLogoFileName)
	case "BespokeLogoFileName":
		res = __gong__marshallString(ident, "BespokeLogoFileName", content.BespokeLogoFileName)
	case "IsBespokePageTileLogoFileName":
		res = __gong__marshallBool(ident, "IsBespokePageTileLogoFileName", content.IsBespokePageTileLogoFileName)
	case "BespokePageTileLogoFileName":
		res = __gong__marshallString(ident, "BespokePageTileLogoFileName", content.BespokePageTileLogoFileName)
	case "Target":
		res = __gong__marshallEnumString(ident, "Target", content.Target.ToCodeString())
	case "VersionInfo":
		res = __gong__marshallString(ident, "VersionInfo", content.VersionInfo)

	case "Chapters":
		var sb strings.Builder
		for _, _chapter := range content.Chapters {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Chapters", _chapter.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Content", fieldName)
	}
	return
}

func (downloadablefile *DownloadableFile) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := downloadablefile.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", downloadablefile.Name)
	case "Base64Content":
		res = __gong__marshallString(ident, "Base64Content", downloadablefile.Base64Content)

	default:
		log.Panicf("Unknown field %s for Gongstruct DownloadableFile", fieldName)
	}
	return
}

func (jpgimage *JpgImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := jpgimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", jpgimage.Name)
	case "Base64Content":
		res = __gong__marshallString(ident, "Base64Content", jpgimage.Base64Content)

	default:
		log.Panicf("Unknown field %s for Gongstruct JpgImage", fieldName)
	}
	return
}

func (page *Page) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := page.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", page.Name)
	case "MardownContent":
		res = __gong__marshallString(ident, "MardownContent", page.MardownContent)

	case "Sections":
		var sb strings.Builder
		for _, _section := range page.Sections {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Sections", _section.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Page", fieldName)
	}
	return
}

func (pngimage *PngImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pngimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pngimage.Name)
	case "Base64Content":
		res = __gong__marshallString(ident, "Base64Content", pngimage.Base64Content)

	default:
		log.Panicf("Unknown field %s for Gongstruct PngImage", fieldName)
	}
	return
}

func (section *Section) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := section.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", section.Name)
	case "MardownContent":
		res = __gong__marshallString(ident, "MardownContent", section.MardownContent)
	case "IsImage":
		res = __gong__marshallBool(ident, "IsImage", section.IsImage)
	case "IsDownloadableFile":
		res = __gong__marshallBool(ident, "IsDownloadableFile", section.IsDownloadableFile)

	case "SvgImage":
		if section.SvgImage != nil {
			res = __gong__marshallPointer(ident, "SvgImage", section.SvgImage.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SvgImage", "nil")
		}
	case "PngImage":
		if section.PngImage != nil {
			res = __gong__marshallPointer(ident, "PngImage", section.PngImage.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PngImage", "nil")
		}
	case "JpgImage":
		if section.JpgImage != nil {
			res = __gong__marshallPointer(ident, "JpgImage", section.JpgImage.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "JpgImage", "nil")
		}
	case "DownloadableFile":
		if section.DownloadableFile != nil {
			res = __gong__marshallPointer(ident, "DownloadableFile", section.DownloadableFile.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DownloadableFile", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Section", fieldName)
	}
	return
}

func (svgimage *SvgImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := svgimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", svgimage.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", svgimage.Content)

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
