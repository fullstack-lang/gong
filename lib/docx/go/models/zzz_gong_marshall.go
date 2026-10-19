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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/docx/go/models"
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

	if stage != nil {
		for _, body := range __gong__sortStageSetInstances(stage.Bodys, stage.Body_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			bodyIdent := body.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Body{Name: %s}).Stage(stage)", bodyIdent, __gong__toRawStringLiteral(body.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", bodyIdent, __gong__toRawStringLiteral(body.Name)))
			for _, elem := range body.Paragraphs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", bodyIdent, bodyIdent, targetIdent))
			}
			for _, elem := range body.Tables {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Tables = append(%s.Tables, %s)", bodyIdent, bodyIdent, targetIdent))
			}
			if body.LastParagraph != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := body.LastParagraph.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LastParagraph = %s", bodyIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, document := range __gong__sortStageSetInstances(stage.Documents, stage.Document_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			documentIdent := document.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Document{Name: %s}).Stage(stage)", documentIdent, __gong__toRawStringLiteral(document.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", documentIdent, __gong__toRawStringLiteral(document.Name)))
			if document.File != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := document.File.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.File = %s", documentIdent, targetIdent))
			}
			if document.Root != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := document.Root.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Root = %s", documentIdent, targetIdent))
			}
			if document.Body != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := document.Body.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Body = %s", documentIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, docx := range __gong__sortStageSetInstances(stage.Docxs, stage.Docx_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			docxIdent := docx.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Docx{Name: %s}).Stage(stage)", docxIdent, __gong__toRawStringLiteral(docx.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", docxIdent, __gong__toRawStringLiteral(docx.Name)))
			for _, elem := range docx.Files {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Files = append(%s.Files, %s)", docxIdent, docxIdent, targetIdent))
			}
			if docx.Document != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := docx.Document.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Document = %s", docxIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, file := range __gong__sortStageSetInstances(stage.Files, stage.File_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			fileIdent := file.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.File{Name: %s}).Stage(stage)", fileIdent, __gong__toRawStringLiteral(file.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", fileIdent, __gong__toRawStringLiteral(file.Name)))
		}
	}
	if stage != nil {
		for _, node := range __gong__sortStageSetInstances(stage.Nodes, stage.Node_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			nodeIdent := node.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Node{Name: %s}).Stage(stage)", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", nodeIdent, __gong__toRawStringLiteral(node.Name)))
			for _, elem := range node.Nodes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Nodes = append(%s.Nodes, %s)", nodeIdent, nodeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, paragraph := range __gong__sortStageSetInstances(stage.Paragraphs, stage.Paragraph_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			paragraphIdent := paragraph.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Paragraph{Name: %s}).Stage(stage)", paragraphIdent, __gong__toRawStringLiteral(paragraph.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.CollatedText = %s", paragraphIdent, __gong__toRawStringLiteral(paragraph.CollatedText)))
			if paragraph.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraph.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphIdent, targetIdent))
			}
			if paragraph.ParagraphProperties != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraph.ParagraphProperties.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParagraphProperties = %s", paragraphIdent, targetIdent))
			}
			for _, elem := range paragraph.Runes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Runes = append(%s.Runes, %s)", paragraphIdent, paragraphIdent, targetIdent))
			}
			if paragraph.Next != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraph.Next.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Next = %s", paragraphIdent, targetIdent))
			}
			if paragraph.Previous != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraph.Previous.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Previous = %s", paragraphIdent, targetIdent))
			}
			if paragraph.EnclosingBody != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraph.EnclosingBody.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingBody = %s", paragraphIdent, targetIdent))
			}
			if paragraph.EnclosingTableColumn != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraph.EnclosingTableColumn.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingTableColumn = %s", paragraphIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, paragraphproperties := range __gong__sortStageSetInstances(stage.ParagraphPropertiess, stage.ParagraphProperties_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			paragraphpropertiesIdent := paragraphproperties.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParagraphProperties{Name: %s}).Stage(stage)", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphpropertiesIdent, __gong__toRawStringLiteral(paragraphproperties.Content)))
			if paragraphproperties.ParagraphStyle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraphproperties.ParagraphStyle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ParagraphStyle = %s", paragraphpropertiesIdent, targetIdent))
			}
			if paragraphproperties.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraphproperties.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphpropertiesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, paragraphstyle := range __gong__sortStageSetInstances(stage.ParagraphStyles, stage.ParagraphStyle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			paragraphstyleIdent := paragraphstyle.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ParagraphStyle{Name: %s}).Stage(stage)", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.ValAttr = %s", paragraphstyleIdent, __gong__toRawStringLiteral(paragraphstyle.ValAttr)))
			if paragraphstyle.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := paragraphstyle.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", paragraphstyleIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, rune := range __gong__sortStageSetInstances(stage.Runes, stage.Rune_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runeIdent := rune.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Rune{Name: %s}).Stage(stage)", runeIdent, __gong__toRawStringLiteral(rune.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runeIdent, __gong__toRawStringLiteral(rune.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", runeIdent, __gong__toRawStringLiteral(rune.Content)))
			if rune.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rune.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", runeIdent, targetIdent))
			}
			if rune.Text != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rune.Text.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Text = %s", runeIdent, targetIdent))
			}
			if rune.RuneProperties != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rune.RuneProperties.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RuneProperties = %s", runeIdent, targetIdent))
			}
			if rune.EnclosingParagraph != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := rune.EnclosingParagraph.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingParagraph = %s", runeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, runeproperties := range __gong__sortStageSetInstances(stage.RunePropertiess, stage.RuneProperties_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runepropertiesIdent := runeproperties.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RuneProperties{Name: %s}).Stage(stage)", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsBold = %t", runepropertiesIdent, runeproperties.IsBold))
			values.WriteString(fmt.Sprintf("\n\t%s.IsStrike = %t", runepropertiesIdent, runeproperties.IsStrike))
			values.WriteString(fmt.Sprintf("\n\t%s.IsItalic = %t", runepropertiesIdent, runeproperties.IsItalic))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", runepropertiesIdent, __gong__toRawStringLiteral(runeproperties.Content)))
			if runeproperties.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := runeproperties.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", runepropertiesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, table := range __gong__sortStageSetInstances(stage.Tables, stage.Table_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tableIdent := table.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Table{Name: %s}).Stage(stage)", tableIdent, __gong__toRawStringLiteral(table.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tableIdent, __gong__toRawStringLiteral(table.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tableIdent, __gong__toRawStringLiteral(table.Content)))
			if table.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := table.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tableIdent, targetIdent))
			}
			if table.TableProperties != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := table.TableProperties.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableProperties = %s", tableIdent, targetIdent))
			}
			for _, elem := range table.TableRows {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableRows = append(%s.TableRows, %s)", tableIdent, tableIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tablecolumn := range __gong__sortStageSetInstances(stage.TableColumns, stage.TableColumn_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablecolumnIdent := tablecolumn.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableColumn{Name: %s}).Stage(stage)", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablecolumnIdent, __gong__toRawStringLiteral(tablecolumn.Content)))
			if tablecolumn.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := tablecolumn.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablecolumnIdent, targetIdent))
			}
			for _, elem := range tablecolumn.Paragraphs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", tablecolumnIdent, tablecolumnIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tableproperties := range __gong__sortStageSetInstances(stage.TablePropertiess, stage.TableProperties_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablepropertiesIdent := tableproperties.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableProperties{Name: %s}).Stage(stage)", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablepropertiesIdent, __gong__toRawStringLiteral(tableproperties.Content)))
			if tableproperties.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := tableproperties.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablepropertiesIdent, targetIdent))
			}
			if tableproperties.TableStyle != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := tableproperties.TableStyle.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableStyle = %s", tablepropertiesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tablerow := range __gong__sortStageSetInstances(stage.TableRows, stage.TableRow_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablerowIdent := tablerow.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableRow{Name: %s}).Stage(stage)", tablerowIdent, __gong__toRawStringLiteral(tablerow.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablerowIdent, __gong__toRawStringLiteral(tablerow.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablerowIdent, __gong__toRawStringLiteral(tablerow.Content)))
			if tablerow.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := tablerow.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablerowIdent, targetIdent))
			}
			for _, elem := range tablerow.TableColumns {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TableColumns = append(%s.TableColumns, %s)", tablerowIdent, tablerowIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, tablestyle := range __gong__sortStageSetInstances(stage.TableStyles, stage.TableStyle_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			tablestyleIdent := tablestyle.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.TableStyle{Name: %s}).Stage(stage)", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.Val = %s", tablestyleIdent, __gong__toRawStringLiteral(tablestyle.Val)))
			if tablestyle.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := tablestyle.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", tablestyleIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, text := range __gong__sortStageSetInstances(stage.Texts, stage.Text_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			textIdent := text.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Text{Name: %s}).Stage(stage)", textIdent, __gong__toRawStringLiteral(text.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", textIdent, __gong__toRawStringLiteral(text.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", textIdent, __gong__toRawStringLiteral(text.Content)))
			values.WriteString(fmt.Sprintf("\n\t%s.PreserveWhiteSpace = %t", textIdent, text.PreserveWhiteSpace))
			if text.Node != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := text.Node.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Node = %s", textIdent, targetIdent))
			}
			if text.EnclosingRune != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := text.EnclosingRune.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EnclosingRune = %s", textIdent, targetIdent))
			}
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
go run github.com/fullstack-lang/gong/lib/docx/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/docx/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/docx/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/docx/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (body *Body) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := body.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", body.Name)

	case "Paragraphs":
		var sb strings.Builder
		for _, _paragraph := range body.Paragraphs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Paragraphs", _paragraph.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Tables":
		var sb strings.Builder
		for _, _table := range body.Tables {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Tables", _table.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "LastParagraph":
		if body.LastParagraph != nil {
			res = __gong__marshallPointer(ident, "LastParagraph", body.LastParagraph.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LastParagraph", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Body", fieldName)
	}
	return
}

func (document *Document) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := document.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", document.Name)

	case "File":
		if document.File != nil {
			res = __gong__marshallPointer(ident, "File", document.File.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "File", "nil")
		}
	case "Root":
		if document.Root != nil {
			res = __gong__marshallPointer(ident, "Root", document.Root.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Root", "nil")
		}
	case "Body":
		if document.Body != nil {
			res = __gong__marshallPointer(ident, "Body", document.Body.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Body", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Document", fieldName)
	}
	return
}

func (docx *Docx) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := docx.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", docx.Name)

	case "Files":
		var sb strings.Builder
		for _, _file := range docx.Files {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Files", _file.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Document":
		if docx.Document != nil {
			res = __gong__marshallPointer(ident, "Document", docx.Document.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Document", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Docx", fieldName)
	}
	return
}

func (file *File) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := file.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", file.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct File", fieldName)
	}
	return
}

func (node *Node) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := node.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", node.Name)

	case "Nodes":
		var sb strings.Builder
		for _, _node := range node.Nodes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Nodes", _node.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Node", fieldName)
	}
	return
}

func (paragraph *Paragraph) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := paragraph.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", paragraph.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", paragraph.Content)
	case "CollatedText":
		res = __gong__marshallString(ident, "CollatedText", paragraph.CollatedText)

	case "Node":
		if paragraph.Node != nil {
			res = __gong__marshallPointer(ident, "Node", paragraph.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	case "ParagraphProperties":
		if paragraph.ParagraphProperties != nil {
			res = __gong__marshallPointer(ident, "ParagraphProperties", paragraph.ParagraphProperties.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ParagraphProperties", "nil")
		}
	case "Runes":
		var sb strings.Builder
		for _, _rune := range paragraph.Runes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Runes", _rune.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Next":
		if paragraph.Next != nil {
			res = __gong__marshallPointer(ident, "Next", paragraph.Next.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Next", "nil")
		}
	case "Previous":
		if paragraph.Previous != nil {
			res = __gong__marshallPointer(ident, "Previous", paragraph.Previous.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Previous", "nil")
		}
	case "EnclosingBody":
		if paragraph.EnclosingBody != nil {
			res = __gong__marshallPointer(ident, "EnclosingBody", paragraph.EnclosingBody.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnclosingBody", "nil")
		}
	case "EnclosingTableColumn":
		if paragraph.EnclosingTableColumn != nil {
			res = __gong__marshallPointer(ident, "EnclosingTableColumn", paragraph.EnclosingTableColumn.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnclosingTableColumn", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Paragraph", fieldName)
	}
	return
}

func (paragraphproperties *ParagraphProperties) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := paragraphproperties.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", paragraphproperties.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", paragraphproperties.Content)

	case "ParagraphStyle":
		if paragraphproperties.ParagraphStyle != nil {
			res = __gong__marshallPointer(ident, "ParagraphStyle", paragraphproperties.ParagraphStyle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ParagraphStyle", "nil")
		}
	case "Node":
		if paragraphproperties.Node != nil {
			res = __gong__marshallPointer(ident, "Node", paragraphproperties.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParagraphProperties", fieldName)
	}
	return
}

func (paragraphstyle *ParagraphStyle) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := paragraphstyle.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", paragraphstyle.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", paragraphstyle.Content)
	case "ValAttr":
		res = __gong__marshallString(ident, "ValAttr", paragraphstyle.ValAttr)

	case "Node":
		if paragraphstyle.Node != nil {
			res = __gong__marshallPointer(ident, "Node", paragraphstyle.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ParagraphStyle", fieldName)
	}
	return
}

func (rune *Rune) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := rune.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", rune.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", rune.Content)

	case "Node":
		if rune.Node != nil {
			res = __gong__marshallPointer(ident, "Node", rune.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	case "Text":
		if rune.Text != nil {
			res = __gong__marshallPointer(ident, "Text", rune.Text.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Text", "nil")
		}
	case "RuneProperties":
		if rune.RuneProperties != nil {
			res = __gong__marshallPointer(ident, "RuneProperties", rune.RuneProperties.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "RuneProperties", "nil")
		}
	case "EnclosingParagraph":
		if rune.EnclosingParagraph != nil {
			res = __gong__marshallPointer(ident, "EnclosingParagraph", rune.EnclosingParagraph.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnclosingParagraph", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Rune", fieldName)
	}
	return
}

func (runeproperties *RuneProperties) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := runeproperties.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", runeproperties.Name)
	case "IsBold":
		res = __gong__marshallBool(ident, "IsBold", runeproperties.IsBold)
	case "IsStrike":
		res = __gong__marshallBool(ident, "IsStrike", runeproperties.IsStrike)
	case "IsItalic":
		res = __gong__marshallBool(ident, "IsItalic", runeproperties.IsItalic)
	case "Content":
		res = __gong__marshallString(ident, "Content", runeproperties.Content)

	case "Node":
		if runeproperties.Node != nil {
			res = __gong__marshallPointer(ident, "Node", runeproperties.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RuneProperties", fieldName)
	}
	return
}

func (table *Table) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := table.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", table.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", table.Content)

	case "Node":
		if table.Node != nil {
			res = __gong__marshallPointer(ident, "Node", table.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	case "TableProperties":
		if table.TableProperties != nil {
			res = __gong__marshallPointer(ident, "TableProperties", table.TableProperties.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TableProperties", "nil")
		}
	case "TableRows":
		var sb strings.Builder
		for _, _tablerow := range table.TableRows {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TableRows", _tablerow.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Table", fieldName)
	}
	return
}

func (tablecolumn *TableColumn) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tablecolumn.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tablecolumn.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", tablecolumn.Content)

	case "Node":
		if tablecolumn.Node != nil {
			res = __gong__marshallPointer(ident, "Node", tablecolumn.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	case "Paragraphs":
		var sb strings.Builder
		for _, _paragraph := range tablecolumn.Paragraphs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Paragraphs", _paragraph.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TableColumn", fieldName)
	}
	return
}

func (tableproperties *TableProperties) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tableproperties.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tableproperties.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", tableproperties.Content)

	case "Node":
		if tableproperties.Node != nil {
			res = __gong__marshallPointer(ident, "Node", tableproperties.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	case "TableStyle":
		if tableproperties.TableStyle != nil {
			res = __gong__marshallPointer(ident, "TableStyle", tableproperties.TableStyle.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TableStyle", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TableProperties", fieldName)
	}
	return
}

func (tablerow *TableRow) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tablerow.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tablerow.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", tablerow.Content)

	case "Node":
		if tablerow.Node != nil {
			res = __gong__marshallPointer(ident, "Node", tablerow.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	case "TableColumns":
		var sb strings.Builder
		for _, _tablecolumn := range tablerow.TableColumns {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TableColumns", _tablecolumn.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TableRow", fieldName)
	}
	return
}

func (tablestyle *TableStyle) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := tablestyle.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", tablestyle.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", tablestyle.Content)
	case "Val":
		res = __gong__marshallString(ident, "Val", tablestyle.Val)

	case "Node":
		if tablestyle.Node != nil {
			res = __gong__marshallPointer(ident, "Node", tablestyle.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TableStyle", fieldName)
	}
	return
}

func (text *Text) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := text.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", text.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", text.Content)
	case "PreserveWhiteSpace":
		res = __gong__marshallBool(ident, "PreserveWhiteSpace", text.PreserveWhiteSpace)

	case "Node":
		if text.Node != nil {
			res = __gong__marshallPointer(ident, "Node", text.Node.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Node", "nil")
		}
	case "EnclosingRune":
		if text.EnclosingRune != nil {
			res = __gong__marshallPointer(ident, "EnclosingRune", text.EnclosingRune.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EnclosingRune", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct Text", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (body *Body) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(body.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(body.GongMarshallField(stage, "Paragraphs"))
		pointersInitializesStatements.WriteString(body.GongMarshallField(stage, "Tables"))
		pointersInitializesStatements.WriteString(body.GongMarshallField(stage, "LastParagraph"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (document *Document) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(document.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(document.GongMarshallField(stage, "File"))
		pointersInitializesStatements.WriteString(document.GongMarshallField(stage, "Root"))
		pointersInitializesStatements.WriteString(document.GongMarshallField(stage, "Body"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (docx *Docx) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(docx.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(docx.GongMarshallField(stage, "Files"))
		pointersInitializesStatements.WriteString(docx.GongMarshallField(stage, "Document"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (file *File) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(file.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (node *Node) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(node.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(node.GongMarshallField(stage, "Nodes"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (paragraph *Paragraph) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(paragraph.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(paragraph.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(paragraph.GongMarshallField(stage, "Node"))
		pointersInitializesStatements.WriteString(paragraph.GongMarshallField(stage, "ParagraphProperties"))
		pointersInitializesStatements.WriteString(paragraph.GongMarshallField(stage, "Runes"))
		initializerStatements.WriteString(paragraph.GongMarshallField(stage, "CollatedText"))
		pointersInitializesStatements.WriteString(paragraph.GongMarshallField(stage, "Next"))
		pointersInitializesStatements.WriteString(paragraph.GongMarshallField(stage, "Previous"))
		pointersInitializesStatements.WriteString(paragraph.GongMarshallField(stage, "EnclosingBody"))
		pointersInitializesStatements.WriteString(paragraph.GongMarshallField(stage, "EnclosingTableColumn"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (paragraphproperties *ParagraphProperties) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(paragraphproperties.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(paragraphproperties.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(paragraphproperties.GongMarshallField(stage, "ParagraphStyle"))
		pointersInitializesStatements.WriteString(paragraphproperties.GongMarshallField(stage, "Node"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (paragraphstyle *ParagraphStyle) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(paragraphstyle.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(paragraphstyle.GongMarshallField(stage, "Node"))
		initializerStatements.WriteString(paragraphstyle.GongMarshallField(stage, "Content"))
		initializerStatements.WriteString(paragraphstyle.GongMarshallField(stage, "ValAttr"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (rune *Rune) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(rune.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(rune.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(rune.GongMarshallField(stage, "Node"))
		pointersInitializesStatements.WriteString(rune.GongMarshallField(stage, "Text"))
		pointersInitializesStatements.WriteString(rune.GongMarshallField(stage, "RuneProperties"))
		pointersInitializesStatements.WriteString(rune.GongMarshallField(stage, "EnclosingParagraph"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (runeproperties *RuneProperties) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(runeproperties.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(runeproperties.GongMarshallField(stage, "Node"))
		initializerStatements.WriteString(runeproperties.GongMarshallField(stage, "IsBold"))
		initializerStatements.WriteString(runeproperties.GongMarshallField(stage, "IsStrike"))
		initializerStatements.WriteString(runeproperties.GongMarshallField(stage, "IsItalic"))
		initializerStatements.WriteString(runeproperties.GongMarshallField(stage, "Content"))
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
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "Node"))
		initializerStatements.WriteString(table.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "TableProperties"))
		pointersInitializesStatements.WriteString(table.GongMarshallField(stage, "TableRows"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tablecolumn *TableColumn) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tablecolumn.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tablecolumn.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(tablecolumn.GongMarshallField(stage, "Node"))
		pointersInitializesStatements.WriteString(tablecolumn.GongMarshallField(stage, "Paragraphs"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tableproperties *TableProperties) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tableproperties.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(tableproperties.GongMarshallField(stage, "Node"))
		initializerStatements.WriteString(tableproperties.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(tableproperties.GongMarshallField(stage, "TableStyle"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tablerow *TableRow) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tablerow.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(tablerow.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(tablerow.GongMarshallField(stage, "Node"))
		pointersInitializesStatements.WriteString(tablerow.GongMarshallField(stage, "TableColumns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (tablestyle *TableStyle) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(tablestyle.GongMarshallField(stage, "Name"))
		pointersInitializesStatements.WriteString(tablestyle.GongMarshallField(stage, "Node"))
		initializerStatements.WriteString(tablestyle.GongMarshallField(stage, "Content"))
		initializerStatements.WriteString(tablestyle.GongMarshallField(stage, "Val"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (text *Text) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(text.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "Content"))
		pointersInitializesStatements.WriteString(text.GongMarshallField(stage, "Node"))
		initializerStatements.WriteString(text.GongMarshallField(stage, "PreserveWhiteSpace"))
		pointersInitializesStatements.WriteString(text.GongMarshallField(stage, "EnclosingRune"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
