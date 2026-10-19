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
	modelsPackageName := "github.com/fullstack-lang/gong/lib/doc/go/models"
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
		for _, attributeshape := range __gong__sortStageSetInstances(stage.AttributeShapes, stage.AttributeShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attributeshapeIdent := localAliasPrefix + attributeshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.AttributeShape{Name: %s}).Stage(stage)", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Name)))
			if str, ok := attributeshape.IdentifierMeta.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.IdentifierMeta = %s", attributeshapeIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.FieldTypeAsString = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.FieldTypeAsString)))
			values.WriteString(fmt.Sprintf("\n\t%s.Structname = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Structname)))
			values.WriteString(fmt.Sprintf("\n\t%s.Fieldtypename = %s", attributeshapeIdent, __gong__toRawStringLiteral(attributeshape.Fieldtypename)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", attributeshapeIdent, attributeshape.IsHidden))
		}
	}
	if stage != nil {
		for _, classdiagram := range __gong__sortStageSetInstances(stage.Classdiagrams, stage.Classdiagram_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			classdiagramIdent := localAliasPrefix + classdiagram.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Classdiagram{Name: %s}).Stage(stage)", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsIncludedInStaticWebSite = %t", classdiagramIdent, classdiagram.IsIncludedInStaticWebSite))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowNbInstances = %t", classdiagramIdent, classdiagram.ShowNbInstances))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowMultiplicity = %t", classdiagramIdent, classdiagram.ShowMultiplicity))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowLinkNames = %t", classdiagramIdent, classdiagram.ShowLinkNames))
			values.WriteString(fmt.Sprintf("\n\t%s.IsInRenameMode = %t", classdiagramIdent, classdiagram.IsInRenameMode))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", classdiagramIdent, classdiagram.IsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongStructsIsExpanded = %t", classdiagramIdent, classdiagram.NodeGongStructsIsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongStructNodeExpansion = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.NodeGongStructNodeExpansion)))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongEnumsIsExpanded = %t", classdiagramIdent, classdiagram.NodeGongEnumsIsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongEnumNodeExpansion = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.NodeGongEnumNodeExpansion)))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongNotesIsExpanded = %t", classdiagramIdent, classdiagram.NodeGongNotesIsExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.NodeGongNoteNodeExpansion = %s", classdiagramIdent, __gong__toRawStringLiteral(classdiagram.NodeGongNoteNodeExpansion)))
			for _, elem := range classdiagram.GongStructShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongStructShapes = append(%s.GongStructShapes, %s)", classdiagramIdent, classdiagramIdent, targetIdent))
			}
			for _, elem := range classdiagram.GongEnumShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnumShapes = append(%s.GongEnumShapes, %s)", classdiagramIdent, classdiagramIdent, targetIdent))
			}
			for _, elem := range classdiagram.GongNoteShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongNoteShapes = append(%s.GongNoteShapes, %s)", classdiagramIdent, classdiagramIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, diagrampackage := range __gong__sortStageSetInstances(stage.DiagramPackages, stage.DiagramPackage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			diagrampackageIdent := localAliasPrefix + diagrampackage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DiagramPackage{Name: %s}).Stage(stage)", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Path = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.Path)))
			values.WriteString(fmt.Sprintf("\n\t%s.GongModelPath = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.GongModelPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.AbsolutePathToDiagramPackage = %s", diagrampackageIdent, __gong__toRawStringLiteral(diagrampackage.AbsolutePathToDiagramPackage)))
			for _, elem := range diagrampackage.Classdiagrams {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Classdiagrams = append(%s.Classdiagrams, %s)", diagrampackageIdent, diagrampackageIdent, targetIdent))
			}
			if diagrampackage.SelectedClassdiagram != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + diagrampackage.SelectedClassdiagram.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SelectedClassdiagram = %s", diagrampackageIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, gongenumshape := range __gong__sortStageSetInstances(stage.GongEnumShapes, stage.GongEnumShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumshapeIdent := localAliasPrefix + gongenumshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnumShape{Name: %s}).Stage(stage)", gongenumshapeIdent, __gong__toRawStringLiteral(gongenumshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumshapeIdent, __gong__toRawStringLiteral(gongenumshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", gongenumshapeIdent, gongenumshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", gongenumshapeIdent, gongenumshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", gongenumshapeIdent, gongenumshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", gongenumshapeIdent, gongenumshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongenumshapeIdent, gongenumshape.IsHidden))
			if str, ok := gongenumshape.IdentifierMeta.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.IdentifierMeta = %s", gongenumshapeIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", gongenumshapeIdent, gongenumshape.IsExpanded))
			for _, elem := range gongenumshape.GongEnumValueShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongEnumValueShapes = append(%s.GongEnumValueShapes, %s)", gongenumshapeIdent, gongenumshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, gongenumvalueshape := range __gong__sortStageSetInstances(stage.GongEnumValueShapes, stage.GongEnumValueShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongenumvalueshapeIdent := localAliasPrefix + gongenumvalueshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongEnumValueShape{Name: %s}).Stage(stage)", gongenumvalueshapeIdent, __gong__toRawStringLiteral(gongenumvalueshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongenumvalueshapeIdent, __gong__toRawStringLiteral(gongenumvalueshape.Name)))
			if str, ok := gongenumvalueshape.IdentifierMeta.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.IdentifierMeta = %s", gongenumvalueshapeIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongenumvalueshapeIdent, gongenumvalueshape.IsHidden))
		}
	}
	if stage != nil {
		for _, gongnotelinkshape := range __gong__sortStageSetInstances(stage.GongNoteLinkShapes, stage.GongNoteLinkShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongnotelinkshapeIdent := localAliasPrefix + gongnotelinkshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongNoteLinkShape{Name: %s}).Stage(stage)", gongnotelinkshapeIdent, __gong__toRawStringLiteral(gongnotelinkshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongnotelinkshapeIdent, __gong__toRawStringLiteral(gongnotelinkshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", gongnotelinkshapeIdent, __gong__toRawStringLiteral(gongnotelinkshape.Identifier)))
			if gongnotelinkshape.Type.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = models.%s", gongnotelinkshapeIdent, gongnotelinkshape.Type.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.Type = \"\"", gongnotelinkshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongnotelinkshapeIdent, gongnotelinkshape.IsHidden))
		}
	}
	if stage != nil {
		for _, gongnoteshape := range __gong__sortStageSetInstances(stage.GongNoteShapes, stage.GongNoteShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongnoteshapeIdent := localAliasPrefix + gongnoteshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongNoteShape{Name: %s}).Stage(stage)", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Identifier = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Identifier)))
			values.WriteString(fmt.Sprintf("\n\t%s.Body = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.Body)))
			values.WriteString(fmt.Sprintf("\n\t%s.BodyHTML = %s", gongnoteshapeIdent, __gong__toRawStringLiteral(gongnoteshape.BodyHTML)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", gongnoteshapeIdent, gongnoteshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", gongnoteshapeIdent, gongnoteshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", gongnoteshapeIdent, gongnoteshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", gongnoteshapeIdent, gongnoteshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongnoteshapeIdent, gongnoteshape.IsHidden))
			values.WriteString(fmt.Sprintf("\n\t%s.Matched = %t", gongnoteshapeIdent, gongnoteshape.Matched))
			values.WriteString(fmt.Sprintf("\n\t%s.IsExpanded = %t", gongnoteshapeIdent, gongnoteshape.IsExpanded))
			for _, elem := range gongnoteshape.GongNoteLinkShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GongNoteLinkShapes = append(%s.GongNoteLinkShapes, %s)", gongnoteshapeIdent, gongnoteshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, gongstructshape := range __gong__sortStageSetInstances(stage.GongStructShapes, stage.GongStructShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			gongstructshapeIdent := localAliasPrefix + gongstructshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.GongStructShape{Name: %s}).Stage(stage)", gongstructshapeIdent, __gong__toRawStringLiteral(gongstructshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", gongstructshapeIdent, __gong__toRawStringLiteral(gongstructshape.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", gongstructshapeIdent, gongstructshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", gongstructshapeIdent, gongstructshape.Y))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %f", gongstructshapeIdent, gongstructshape.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %f", gongstructshapeIdent, gongstructshape.Height))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", gongstructshapeIdent, gongstructshape.IsHidden))
			if str, ok := gongstructshape.IdentifierMeta.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.IdentifierMeta = %s", gongstructshapeIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", gongstructshapeIdent, gongstructshape.IsSelected))
			for _, elem := range gongstructshape.AttributeShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AttributeShapes = append(%s.AttributeShapes, %s)", gongstructshapeIdent, gongstructshapeIdent, targetIdent))
			}
			for _, elem := range gongstructshape.LinkShapes {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LinkShapes = append(%s.LinkShapes, %s)", gongstructshapeIdent, gongstructshapeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, linkshape := range __gong__sortStageSetInstances(stage.LinkShapes, stage.LinkShape_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linkshapeIdent := localAliasPrefix + linkshape.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.LinkShape{Name: %s}).Stage(stage)", linkshapeIdent, __gong__toRawStringLiteral(linkshape.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linkshapeIdent, __gong__toRawStringLiteral(linkshape.Name)))
			if str, ok := linkshape.IdentifierMeta.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.IdentifierMeta = %s", linkshapeIdent, str))
			}
			if str, ok := linkshape.FieldTypeIdentifierMeta.(string); ok {
				values.WriteString(fmt.Sprintf("\n\t%s.FieldTypeIdentifierMeta = %s", linkshapeIdent, str))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.FieldOffsetX = %f", linkshapeIdent, linkshape.FieldOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.FieldOffsetY = %f", linkshapeIdent, linkshape.FieldOffsetY))
			if linkshape.TargetMultiplicity.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.TargetMultiplicity = models.%s", linkshapeIdent, linkshape.TargetMultiplicity.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.TargetMultiplicity = \"\"", linkshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.TargetMultiplicityOffsetX = %f", linkshapeIdent, linkshape.TargetMultiplicityOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.TargetMultiplicityOffsetY = %f", linkshapeIdent, linkshape.TargetMultiplicityOffsetY))
			if linkshape.SourceMultiplicity.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.SourceMultiplicity = models.%s", linkshapeIdent, linkshape.SourceMultiplicity.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.SourceMultiplicity = \"\"", linkshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.SourceMultiplicityOffsetX = %f", linkshapeIdent, linkshape.SourceMultiplicityOffsetX))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceMultiplicityOffsetY = %f", linkshapeIdent, linkshape.SourceMultiplicityOffsetY))
			values.WriteString(fmt.Sprintf("\n\t%s.X = %f", linkshapeIdent, linkshape.X))
			values.WriteString(fmt.Sprintf("\n\t%s.Y = %f", linkshapeIdent, linkshape.Y))
			if linkshape.StartOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = models.%s", linkshapeIdent, linkshape.StartOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.StartOrientation = \"\"", linkshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.StartRatio = %f", linkshapeIdent, linkshape.StartRatio))
			if linkshape.EndOrientation.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = models.%s", linkshapeIdent, linkshape.EndOrientation.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.EndOrientation = \"\"", linkshapeIdent))
			}
			values.WriteString(fmt.Sprintf("\n\t%s.EndRatio = %f", linkshapeIdent, linkshape.EndRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.CornerOffsetRatio = %f", linkshapeIdent, linkshape.CornerOffsetRatio))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHidden = %t", linkshapeIdent, linkshape.IsHidden))
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
go run github.com/fullstack-lang/gong/lib/doc/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/doc/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/lib/doc/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/lib/doc/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (attributeshape *AttributeShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attributeshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attributeshape.Name)
	case "IdentifierMeta":
		if str, ok := attributeshape.IdentifierMeta.(string); ok {
			res = __gong__marshallMeta(ident, "IdentifierMeta", str)
		}
	case "FieldTypeAsString":
		res = __gong__marshallString(ident, "FieldTypeAsString", attributeshape.FieldTypeAsString)
	case "Structname":
		res = __gong__marshallString(ident, "Structname", attributeshape.Structname)
	case "Fieldtypename":
		res = __gong__marshallString(ident, "Fieldtypename", attributeshape.Fieldtypename)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", attributeshape.IsHidden)

	default:
		log.Panicf("Unknown field %s for Gongstruct AttributeShape", fieldName)
	}
	return
}

func (classdiagram *Classdiagram) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := classdiagram.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", classdiagram.Name)
	case "Description":
		res = __gong__marshallString(ident, "Description", classdiagram.Description)
	case "IsIncludedInStaticWebSite":
		res = __gong__marshallBool(ident, "IsIncludedInStaticWebSite", classdiagram.IsIncludedInStaticWebSite)
	case "ShowNbInstances":
		res = __gong__marshallBool(ident, "ShowNbInstances", classdiagram.ShowNbInstances)
	case "ShowMultiplicity":
		res = __gong__marshallBool(ident, "ShowMultiplicity", classdiagram.ShowMultiplicity)
	case "ShowLinkNames":
		res = __gong__marshallBool(ident, "ShowLinkNames", classdiagram.ShowLinkNames)
	case "IsInRenameMode":
		res = __gong__marshallBool(ident, "IsInRenameMode", classdiagram.IsInRenameMode)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", classdiagram.IsExpanded)
	case "NodeGongStructsIsExpanded":
		res = __gong__marshallBool(ident, "NodeGongStructsIsExpanded", classdiagram.NodeGongStructsIsExpanded)
	case "NodeGongStructNodeExpansion":
		res = __gong__marshallString(ident, "NodeGongStructNodeExpansion", classdiagram.NodeGongStructNodeExpansion)
	case "NodeGongEnumsIsExpanded":
		res = __gong__marshallBool(ident, "NodeGongEnumsIsExpanded", classdiagram.NodeGongEnumsIsExpanded)
	case "NodeGongEnumNodeExpansion":
		res = __gong__marshallString(ident, "NodeGongEnumNodeExpansion", classdiagram.NodeGongEnumNodeExpansion)
	case "NodeGongNotesIsExpanded":
		res = __gong__marshallBool(ident, "NodeGongNotesIsExpanded", classdiagram.NodeGongNotesIsExpanded)
	case "NodeGongNoteNodeExpansion":
		res = __gong__marshallString(ident, "NodeGongNoteNodeExpansion", classdiagram.NodeGongNoteNodeExpansion)

	case "GongStructShapes":
		var sb strings.Builder
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongStructShapes", _gongstructshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GongEnumShapes":
		var sb strings.Builder
		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongEnumShapes", _gongenumshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "GongNoteShapes":
		var sb strings.Builder
		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongNoteShapes", _gongnoteshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct Classdiagram", fieldName)
	}
	return
}

func (diagrampackage *DiagramPackage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := diagrampackage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", diagrampackage.Name)
	case "Path":
		res = __gong__marshallString(ident, "Path", diagrampackage.Path)
	case "GongModelPath":
		res = __gong__marshallString(ident, "GongModelPath", diagrampackage.GongModelPath)
	case "AbsolutePathToDiagramPackage":
		res = __gong__marshallString(ident, "AbsolutePathToDiagramPackage", diagrampackage.AbsolutePathToDiagramPackage)

	case "Classdiagrams":
		var sb strings.Builder
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Classdiagrams", _classdiagram.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SelectedClassdiagram":
		if diagrampackage.SelectedClassdiagram != nil {
			res = __gong__marshallPointer(ident, "SelectedClassdiagram", diagrampackage.SelectedClassdiagram.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SelectedClassdiagram", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DiagramPackage", fieldName)
	}
	return
}

func (gongenumshape *GongEnumShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongenumshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongenumshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", gongenumshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", gongenumshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", gongenumshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", gongenumshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", gongenumshape.IsHidden)
	case "IdentifierMeta":
		if str, ok := gongenumshape.IdentifierMeta.(string); ok {
			res = __gong__marshallMeta(ident, "IdentifierMeta", str)
		}
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", gongenumshape.IsExpanded)

	case "GongEnumValueShapes":
		var sb strings.Builder
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongEnumValueShapes", _gongenumvalueshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumShape", fieldName)
	}
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongenumvalueshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongenumvalueshape.Name)
	case "IdentifierMeta":
		if str, ok := gongenumvalueshape.IdentifierMeta.(string); ok {
			res = __gong__marshallMeta(ident, "IdentifierMeta", str)
		}
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", gongenumvalueshape.IsHidden)

	default:
		log.Panicf("Unknown field %s for Gongstruct GongEnumValueShape", fieldName)
	}
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongnotelinkshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongnotelinkshape.Name)
	case "Identifier":
		res = __gong__marshallString(ident, "Identifier", gongnotelinkshape.Identifier)
	case "Type":
		res = __gong__marshallEnumString(ident, "Type", gongnotelinkshape.Type.ToCodeString())
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", gongnotelinkshape.IsHidden)

	default:
		log.Panicf("Unknown field %s for Gongstruct GongNoteLinkShape", fieldName)
	}
	return
}

func (gongnoteshape *GongNoteShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongnoteshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongnoteshape.Name)
	case "Identifier":
		res = __gong__marshallString(ident, "Identifier", gongnoteshape.Identifier)
	case "Body":
		res = __gong__marshallString(ident, "Body", gongnoteshape.Body)
	case "BodyHTML":
		res = __gong__marshallString(ident, "BodyHTML", gongnoteshape.BodyHTML)
	case "X":
		res = __gong__marshallFloat(ident, "X", gongnoteshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", gongnoteshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", gongnoteshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", gongnoteshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", gongnoteshape.IsHidden)
	case "Matched":
		res = __gong__marshallBool(ident, "Matched", gongnoteshape.Matched)
	case "IsExpanded":
		res = __gong__marshallBool(ident, "IsExpanded", gongnoteshape.IsExpanded)

	case "GongNoteLinkShapes":
		var sb strings.Builder
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GongNoteLinkShapes", _gongnotelinkshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongNoteShape", fieldName)
	}
	return
}

func (gongstructshape *GongStructShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := gongstructshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", gongstructshape.Name)
	case "X":
		res = __gong__marshallFloat(ident, "X", gongstructshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", gongstructshape.Y)
	case "Width":
		res = __gong__marshallFloat(ident, "Width", gongstructshape.Width)
	case "Height":
		res = __gong__marshallFloat(ident, "Height", gongstructshape.Height)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", gongstructshape.IsHidden)
	case "IdentifierMeta":
		if str, ok := gongstructshape.IdentifierMeta.(string); ok {
			res = __gong__marshallMeta(ident, "IdentifierMeta", str)
		}
	case "IsSelected":
		res = __gong__marshallBool(ident, "IsSelected", gongstructshape.IsSelected)

	case "AttributeShapes":
		var sb strings.Builder
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AttributeShapes", _attributeshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "LinkShapes":
		var sb strings.Builder
		for _, _linkshape := range gongstructshape.LinkShapes {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "LinkShapes", _linkshape.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GongStructShape", fieldName)
	}
	return
}

func (linkshape *LinkShape) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linkshape.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linkshape.Name)
	case "IdentifierMeta":
		if str, ok := linkshape.IdentifierMeta.(string); ok {
			res = __gong__marshallMeta(ident, "IdentifierMeta", str)
		}
	case "FieldTypeIdentifierMeta":
		if str, ok := linkshape.FieldTypeIdentifierMeta.(string); ok {
			res = __gong__marshallMeta(ident, "FieldTypeIdentifierMeta", str)
		}
	case "FieldOffsetX":
		res = __gong__marshallFloat(ident, "FieldOffsetX", linkshape.FieldOffsetX)
	case "FieldOffsetY":
		res = __gong__marshallFloat(ident, "FieldOffsetY", linkshape.FieldOffsetY)
	case "TargetMultiplicity":
		res = __gong__marshallEnumString(ident, "TargetMultiplicity", linkshape.TargetMultiplicity.ToCodeString())
	case "TargetMultiplicityOffsetX":
		res = __gong__marshallFloat(ident, "TargetMultiplicityOffsetX", linkshape.TargetMultiplicityOffsetX)
	case "TargetMultiplicityOffsetY":
		res = __gong__marshallFloat(ident, "TargetMultiplicityOffsetY", linkshape.TargetMultiplicityOffsetY)
	case "SourceMultiplicity":
		res = __gong__marshallEnumString(ident, "SourceMultiplicity", linkshape.SourceMultiplicity.ToCodeString())
	case "SourceMultiplicityOffsetX":
		res = __gong__marshallFloat(ident, "SourceMultiplicityOffsetX", linkshape.SourceMultiplicityOffsetX)
	case "SourceMultiplicityOffsetY":
		res = __gong__marshallFloat(ident, "SourceMultiplicityOffsetY", linkshape.SourceMultiplicityOffsetY)
	case "X":
		res = __gong__marshallFloat(ident, "X", linkshape.X)
	case "Y":
		res = __gong__marshallFloat(ident, "Y", linkshape.Y)
	case "StartOrientation":
		res = __gong__marshallEnumString(ident, "StartOrientation", linkshape.StartOrientation.ToCodeString())
	case "StartRatio":
		res = __gong__marshallFloat(ident, "StartRatio", linkshape.StartRatio)
	case "EndOrientation":
		res = __gong__marshallEnumString(ident, "EndOrientation", linkshape.EndOrientation.ToCodeString())
	case "EndRatio":
		res = __gong__marshallFloat(ident, "EndRatio", linkshape.EndRatio)
	case "CornerOffsetRatio":
		res = __gong__marshallFloat(ident, "CornerOffsetRatio", linkshape.CornerOffsetRatio)
	case "IsHidden":
		res = __gong__marshallBool(ident, "IsHidden", linkshape.IsHidden)

	default:
		log.Panicf("Unknown field %s for Gongstruct LinkShape", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (attributeshape *AttributeShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "IdentifierMeta"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "FieldTypeAsString"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Structname"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "Fieldtypename"))
		initializerStatements.WriteString(attributeshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (classdiagram *Classdiagram) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsIncludedInStaticWebSite"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongStructShapes"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongEnumShapes"))
		pointersInitializesStatements.WriteString(classdiagram.GongMarshallField(stage, "GongNoteShapes"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowNbInstances"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowMultiplicity"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "ShowLinkNames"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsInRenameMode"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "IsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongStructsIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongStructNodeExpansion"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongEnumsIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongEnumNodeExpansion"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongNotesIsExpanded"))
		initializerStatements.WriteString(classdiagram.GongMarshallField(stage, "NodeGongNoteNodeExpansion"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (diagrampackage *DiagramPackage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "Path"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "GongModelPath"))
		pointersInitializesStatements.WriteString(diagrampackage.GongMarshallField(stage, "Classdiagrams"))
		pointersInitializesStatements.WriteString(diagrampackage.GongMarshallField(stage, "SelectedClassdiagram"))
		initializerStatements.WriteString(diagrampackage.GongMarshallField(stage, "AbsolutePathToDiagramPackage"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongenumshape *GongEnumShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "IdentifierMeta"))
		pointersInitializesStatements.WriteString(gongenumshape.GongMarshallField(stage, "GongEnumValueShapes"))
		initializerStatements.WriteString(gongenumshape.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongenumvalueshape *GongEnumValueShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongenumvalueshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongenumvalueshape.GongMarshallField(stage, "IdentifierMeta"))
		initializerStatements.WriteString(gongenumvalueshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongnotelinkshape *GongNoteLinkShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Identifier"))
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "Type"))
		initializerStatements.WriteString(gongnotelinkshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongnoteshape *GongNoteShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Identifier"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Body"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "BodyHTML"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "Matched"))
		pointersInitializesStatements.WriteString(gongnoteshape.GongMarshallField(stage, "GongNoteLinkShapes"))
		initializerStatements.WriteString(gongnoteshape.GongMarshallField(stage, "IsExpanded"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (gongstructshape *GongStructShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "Height"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "IsHidden"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "IdentifierMeta"))
		pointersInitializesStatements.WriteString(gongstructshape.GongMarshallField(stage, "AttributeShapes"))
		pointersInitializesStatements.WriteString(gongstructshape.GongMarshallField(stage, "LinkShapes"))
		initializerStatements.WriteString(gongstructshape.GongMarshallField(stage, "IsSelected"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linkshape *LinkShape) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "IdentifierMeta"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldTypeIdentifierMeta"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "FieldOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicity"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "TargetMultiplicityOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicity"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetX"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "SourceMultiplicityOffsetY"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "X"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "Y"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "StartOrientation"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "StartRatio"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "EndOrientation"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "EndRatio"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "CornerOffsetRatio"))
		initializerStatements.WriteString(linkshape.GongMarshallField(stage, "IsHidden"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
