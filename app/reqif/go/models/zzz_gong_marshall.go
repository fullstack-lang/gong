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
	modelsPackageName := "github.com/fullstack-lang/gong/app/reqif/go/models"
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
		for _, alternative_id := range __gong__sortStageSetInstances(stage.ALTERNATIVE_IDs, stage.ALTERNATIVE_ID_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			alternative_idIdent := localAliasPrefix + alternative_id.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ALTERNATIVE_ID{Name: %s}).Stage(stage)", alternative_idIdent, __gong__toRawStringLiteral(alternative_id.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", alternative_idIdent, __gong__toRawStringLiteral(alternative_id.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", alternative_idIdent, __gong__toRawStringLiteral(alternative_id.IDENTIFIER)))
		}
	}
	if stage != nil {
		for _, attribute_definition_boolean := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_BOOLEANs, stage.ATTRIBUTE_DEFINITION_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_booleanIdent := localAliasPrefix + attribute_definition_boolean.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_BOOLEAN{Name: %s}).Stage(stage)", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_booleanIdent, attribute_definition_boolean.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_booleanIdent, __gong__toRawStringLiteral(attribute_definition_boolean.LONG_NAME)))
			if attribute_definition_boolean.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_booleanIdent, targetIdent))
			}
			if attribute_definition_boolean.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_boolean.DEFAULT_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_booleanIdent, targetIdent))
			}
			if attribute_definition_boolean.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_boolean.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_booleanIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_definition_boolean_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings, stage.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_boolean_renderingIdent := localAliasPrefix + attribute_definition_boolean_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering{Name: %s}).Stage(stage)", attribute_definition_boolean_renderingIdent, __gong__toRawStringLiteral(attribute_definition_boolean_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_boolean_renderingIdent, __gong__toRawStringLiteral(attribute_definition_boolean_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_boolean_renderingIdent, attribute_definition_boolean_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_definition_date := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_DATEs, stage.ATTRIBUTE_DEFINITION_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_dateIdent := localAliasPrefix + attribute_definition_date.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_DATE{Name: %s}).Stage(stage)", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_dateIdent, attribute_definition_date.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_dateIdent, __gong__toRawStringLiteral(attribute_definition_date.LONG_NAME)))
			if attribute_definition_date.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_dateIdent, targetIdent))
			}
			if attribute_definition_date.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_date.DEFAULT_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_dateIdent, targetIdent))
			}
			if attribute_definition_date.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_date.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_dateIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_definition_date_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_DATE_Renderings, stage.ATTRIBUTE_DEFINITION_DATE_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_date_renderingIdent := localAliasPrefix + attribute_definition_date_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_DATE_Rendering{Name: %s}).Stage(stage)", attribute_definition_date_renderingIdent, __gong__toRawStringLiteral(attribute_definition_date_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_date_renderingIdent, __gong__toRawStringLiteral(attribute_definition_date_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_date_renderingIdent, attribute_definition_date_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_definition_enumeration := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, stage.ATTRIBUTE_DEFINITION_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_enumerationIdent := localAliasPrefix + attribute_definition_enumeration.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_ENUMERATION{Name: %s}).Stage(stage)", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_enumerationIdent, attribute_definition_enumeration.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_enumerationIdent, __gong__toRawStringLiteral(attribute_definition_enumeration.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MULTI_VALUED = %t", attribute_definition_enumerationIdent, attribute_definition_enumeration.MULTI_VALUED))
			if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_enumerationIdent, targetIdent))
			}
			if attribute_definition_enumeration.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_enumeration.DEFAULT_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_enumerationIdent, targetIdent))
			}
			if attribute_definition_enumeration.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_enumeration.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_enumerationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_definition_enumeration_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings, stage.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_enumeration_renderingIdent := localAliasPrefix + attribute_definition_enumeration_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering{Name: %s}).Stage(stage)", attribute_definition_enumeration_renderingIdent, __gong__toRawStringLiteral(attribute_definition_enumeration_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_enumeration_renderingIdent, __gong__toRawStringLiteral(attribute_definition_enumeration_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_enumeration_renderingIdent, attribute_definition_enumeration_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_definition_integer := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_INTEGERs, stage.ATTRIBUTE_DEFINITION_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_integerIdent := localAliasPrefix + attribute_definition_integer.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_INTEGER{Name: %s}).Stage(stage)", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_integerIdent, attribute_definition_integer.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_integerIdent, __gong__toRawStringLiteral(attribute_definition_integer.LONG_NAME)))
			if attribute_definition_integer.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_integerIdent, targetIdent))
			}
			if attribute_definition_integer.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_integer.DEFAULT_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_integerIdent, targetIdent))
			}
			if attribute_definition_integer.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_integer.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_integerIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_definition_integer_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings, stage.ATTRIBUTE_DEFINITION_INTEGER_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_integer_renderingIdent := localAliasPrefix + attribute_definition_integer_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_INTEGER_Rendering{Name: %s}).Stage(stage)", attribute_definition_integer_renderingIdent, __gong__toRawStringLiteral(attribute_definition_integer_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_integer_renderingIdent, __gong__toRawStringLiteral(attribute_definition_integer_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_integer_renderingIdent, attribute_definition_integer_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_definition_real := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_REALs, stage.ATTRIBUTE_DEFINITION_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_realIdent := localAliasPrefix + attribute_definition_real.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_REAL{Name: %s}).Stage(stage)", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_realIdent, attribute_definition_real.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_realIdent, __gong__toRawStringLiteral(attribute_definition_real.LONG_NAME)))
			if attribute_definition_real.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_realIdent, targetIdent))
			}
			if attribute_definition_real.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_real.DEFAULT_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_realIdent, targetIdent))
			}
			if attribute_definition_real.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_real.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_realIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_definition_real_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_REAL_Renderings, stage.ATTRIBUTE_DEFINITION_REAL_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_real_renderingIdent := localAliasPrefix + attribute_definition_real_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_REAL_Rendering{Name: %s}).Stage(stage)", attribute_definition_real_renderingIdent, __gong__toRawStringLiteral(attribute_definition_real_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_real_renderingIdent, __gong__toRawStringLiteral(attribute_definition_real_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_real_renderingIdent, attribute_definition_real_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_definition_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_Renderings, stage.ATTRIBUTE_DEFINITION_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_renderingIdent := localAliasPrefix + attribute_definition_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_Rendering{Name: %s}).Stage(stage)", attribute_definition_renderingIdent, __gong__toRawStringLiteral(attribute_definition_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_renderingIdent, __gong__toRawStringLiteral(attribute_definition_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_renderingIdent, attribute_definition_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_renderingIdent, attribute_definition_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_renderingIdent, attribute_definition_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_renderingIdent, attribute_definition_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_definition_string := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_STRINGs, stage.ATTRIBUTE_DEFINITION_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_stringIdent := localAliasPrefix + attribute_definition_string.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_STRING{Name: %s}).Stage(stage)", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_stringIdent, attribute_definition_string.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_stringIdent, __gong__toRawStringLiteral(attribute_definition_string.LONG_NAME)))
			if attribute_definition_string.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_stringIdent, targetIdent))
			}
			if attribute_definition_string.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_string.DEFAULT_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_stringIdent, targetIdent))
			}
			if attribute_definition_string.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_string.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_stringIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_definition_string_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_STRING_Renderings, stage.ATTRIBUTE_DEFINITION_STRING_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_string_renderingIdent := localAliasPrefix + attribute_definition_string_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_STRING_Rendering{Name: %s}).Stage(stage)", attribute_definition_string_renderingIdent, __gong__toRawStringLiteral(attribute_definition_string_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_string_renderingIdent, __gong__toRawStringLiteral(attribute_definition_string_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_string_renderingIdent, attribute_definition_string_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_definition_xhtml := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_XHTMLs, stage.ATTRIBUTE_DEFINITION_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_xhtmlIdent := localAliasPrefix + attribute_definition_xhtml.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_XHTML{Name: %s}).Stage(stage)", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", attribute_definition_xhtmlIdent, attribute_definition_xhtml.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", attribute_definition_xhtmlIdent, __gong__toRawStringLiteral(attribute_definition_xhtml.LONG_NAME)))
			if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", attribute_definition_xhtmlIdent, targetIdent))
			}
			if attribute_definition_xhtml.DEFAULT_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_xhtml.DEFAULT_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFAULT_VALUE = %s", attribute_definition_xhtmlIdent, targetIdent))
			}
			if attribute_definition_xhtml.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_definition_xhtml.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", attribute_definition_xhtmlIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_definition_xhtml_rendering := range __gong__sortStageSetInstances(stage.ATTRIBUTE_DEFINITION_XHTML_Renderings, stage.ATTRIBUTE_DEFINITION_XHTML_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_definition_xhtml_renderingIdent := localAliasPrefix + attribute_definition_xhtml_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_DEFINITION_XHTML_Rendering{Name: %s}).Stage(stage)", attribute_definition_xhtml_renderingIdent, __gong__toRawStringLiteral(attribute_definition_xhtml_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_definition_xhtml_renderingIdent, __gong__toRawStringLiteral(attribute_definition_xhtml_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTable = %t", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.ShowInTable))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInTitle = %t", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.ShowInTitle))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowInSubject = %t", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.ShowInSubject))
			values.WriteString(fmt.Sprintf("\n\t%s.Rank = %d", attribute_definition_xhtml_renderingIdent, attribute_definition_xhtml_rendering.Rank))
		}
	}
	if stage != nil {
		for _, attribute_value_boolean := range __gong__sortStageSetInstances(stage.ATTRIBUTE_VALUE_BOOLEANs, stage.ATTRIBUTE_VALUE_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_booleanIdent := localAliasPrefix + attribute_value_boolean.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_BOOLEAN{Name: %s}).Stage(stage)", attribute_value_booleanIdent, __gong__toRawStringLiteral(attribute_value_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_booleanIdent, __gong__toRawStringLiteral(attribute_value_boolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %t", attribute_value_booleanIdent, attribute_value_boolean.THE_VALUE))
			if attribute_value_boolean.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_boolean.DEFINITION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_booleanIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_value_date := range __gong__sortStageSetInstances(stage.ATTRIBUTE_VALUE_DATEs, stage.ATTRIBUTE_VALUE_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_dateIdent := localAliasPrefix + attribute_value_date.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_DATE{Name: %s}).Stage(stage)", attribute_value_dateIdent, __gong__toRawStringLiteral(attribute_value_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_dateIdent, __gong__toRawStringLiteral(attribute_value_date.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %s", attribute_value_dateIdent, __gong__toRawStringLiteral(attribute_value_date.THE_VALUE)))
			if attribute_value_date.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_date.DEFINITION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_dateIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_value_enumeration := range __gong__sortStageSetInstances(stage.ATTRIBUTE_VALUE_ENUMERATIONs, stage.ATTRIBUTE_VALUE_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_enumerationIdent := localAliasPrefix + attribute_value_enumeration.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_ENUMERATION{Name: %s}).Stage(stage)", attribute_value_enumerationIdent, __gong__toRawStringLiteral(attribute_value_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_enumerationIdent, __gong__toRawStringLiteral(attribute_value_enumeration.Name)))
			if attribute_value_enumeration.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_enumeration.DEFINITION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_enumerationIdent, targetIdent))
			}
			if attribute_value_enumeration.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_enumeration.VALUES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", attribute_value_enumerationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_value_integer := range __gong__sortStageSetInstances(stage.ATTRIBUTE_VALUE_INTEGERs, stage.ATTRIBUTE_VALUE_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_integerIdent := localAliasPrefix + attribute_value_integer.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_INTEGER{Name: %s}).Stage(stage)", attribute_value_integerIdent, __gong__toRawStringLiteral(attribute_value_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_integerIdent, __gong__toRawStringLiteral(attribute_value_integer.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %d", attribute_value_integerIdent, attribute_value_integer.THE_VALUE))
			if attribute_value_integer.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_integer.DEFINITION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_integerIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_value_real := range __gong__sortStageSetInstances(stage.ATTRIBUTE_VALUE_REALs, stage.ATTRIBUTE_VALUE_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_realIdent := localAliasPrefix + attribute_value_real.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_REAL{Name: %s}).Stage(stage)", attribute_value_realIdent, __gong__toRawStringLiteral(attribute_value_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_realIdent, __gong__toRawStringLiteral(attribute_value_real.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %f", attribute_value_realIdent, attribute_value_real.THE_VALUE))
			if attribute_value_real.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_real.DEFINITION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_realIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_value_string := range __gong__sortStageSetInstances(stage.ATTRIBUTE_VALUE_STRINGs, stage.ATTRIBUTE_VALUE_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_stringIdent := localAliasPrefix + attribute_value_string.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_STRING{Name: %s}).Stage(stage)", attribute_value_stringIdent, __gong__toRawStringLiteral(attribute_value_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_stringIdent, __gong__toRawStringLiteral(attribute_value_string.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %s", attribute_value_stringIdent, __gong__toRawStringLiteral(attribute_value_string.THE_VALUE)))
			if attribute_value_string.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_string.DEFINITION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_stringIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, attribute_value_xhtml := range __gong__sortStageSetInstances(stage.ATTRIBUTE_VALUE_XHTMLs, stage.ATTRIBUTE_VALUE_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			attribute_value_xhtmlIdent := localAliasPrefix + attribute_value_xhtml.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ATTRIBUTE_VALUE_XHTML{Name: %s}).Stage(stage)", attribute_value_xhtmlIdent, __gong__toRawStringLiteral(attribute_value_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", attribute_value_xhtmlIdent, __gong__toRawStringLiteral(attribute_value_xhtml.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_SIMPLIFIED = %t", attribute_value_xhtmlIdent, attribute_value_xhtml.IS_SIMPLIFIED))
			if attribute_value_xhtml.DEFINITION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_xhtml.DEFINITION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DEFINITION = %s", attribute_value_xhtmlIdent, targetIdent))
			}
			if attribute_value_xhtml.THE_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_xhtml.THE_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.THE_VALUE = %s", attribute_value_xhtmlIdent, targetIdent))
			}
			if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + attribute_value_xhtml.THE_ORIGINAL_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.THE_ORIGINAL_VALUE = %s", attribute_value_xhtmlIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_alternative_id := range __gong__sortStageSetInstances(stage.A_ALTERNATIVE_IDs, stage.A_ALTERNATIVE_ID_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_alternative_idIdent := localAliasPrefix + a_alternative_id.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ALTERNATIVE_ID{Name: %s}).Stage(stage)", a_alternative_idIdent, __gong__toRawStringLiteral(a_alternative_id.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_alternative_idIdent, __gong__toRawStringLiteral(a_alternative_id.Name)))
			if a_alternative_id.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + a_alternative_id.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", a_alternative_idIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_definition_boolean_ref := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs, stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_boolean_refIdent := localAliasPrefix + a_attribute_definition_boolean_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF{Name: %s}).Stage(stage)", a_attribute_definition_boolean_refIdent, __gong__toRawStringLiteral(a_attribute_definition_boolean_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_boolean_refIdent, __gong__toRawStringLiteral(a_attribute_definition_boolean_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_BOOLEAN_REF = %s", a_attribute_definition_boolean_refIdent, __gong__toRawStringLiteral(a_attribute_definition_boolean_ref.ATTRIBUTE_DEFINITION_BOOLEAN_REF)))
		}
	}
	if stage != nil {
		for _, a_attribute_definition_date_ref := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs, stage.A_ATTRIBUTE_DEFINITION_DATE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_date_refIdent := localAliasPrefix + a_attribute_definition_date_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_DATE_REF{Name: %s}).Stage(stage)", a_attribute_definition_date_refIdent, __gong__toRawStringLiteral(a_attribute_definition_date_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_date_refIdent, __gong__toRawStringLiteral(a_attribute_definition_date_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_DATE_REF = %s", a_attribute_definition_date_refIdent, __gong__toRawStringLiteral(a_attribute_definition_date_ref.ATTRIBUTE_DEFINITION_DATE_REF)))
		}
	}
	if stage != nil {
		for _, a_attribute_definition_enumeration_ref := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs, stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_enumeration_refIdent := localAliasPrefix + a_attribute_definition_enumeration_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF{Name: %s}).Stage(stage)", a_attribute_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_attribute_definition_enumeration_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_attribute_definition_enumeration_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_ENUMERATION_REF = %s", a_attribute_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_attribute_definition_enumeration_ref.ATTRIBUTE_DEFINITION_ENUMERATION_REF)))
		}
	}
	if stage != nil {
		for _, a_attribute_definition_integer_ref := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs, stage.A_ATTRIBUTE_DEFINITION_INTEGER_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_integer_refIdent := localAliasPrefix + a_attribute_definition_integer_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_INTEGER_REF{Name: %s}).Stage(stage)", a_attribute_definition_integer_refIdent, __gong__toRawStringLiteral(a_attribute_definition_integer_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_integer_refIdent, __gong__toRawStringLiteral(a_attribute_definition_integer_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_INTEGER_REF = %s", a_attribute_definition_integer_refIdent, __gong__toRawStringLiteral(a_attribute_definition_integer_ref.ATTRIBUTE_DEFINITION_INTEGER_REF)))
		}
	}
	if stage != nil {
		for _, a_attribute_definition_real_ref := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs, stage.A_ATTRIBUTE_DEFINITION_REAL_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_real_refIdent := localAliasPrefix + a_attribute_definition_real_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_REAL_REF{Name: %s}).Stage(stage)", a_attribute_definition_real_refIdent, __gong__toRawStringLiteral(a_attribute_definition_real_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_real_refIdent, __gong__toRawStringLiteral(a_attribute_definition_real_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_REAL_REF = %s", a_attribute_definition_real_refIdent, __gong__toRawStringLiteral(a_attribute_definition_real_ref.ATTRIBUTE_DEFINITION_REAL_REF)))
		}
	}
	if stage != nil {
		for _, a_attribute_definition_string_ref := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs, stage.A_ATTRIBUTE_DEFINITION_STRING_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_string_refIdent := localAliasPrefix + a_attribute_definition_string_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_STRING_REF{Name: %s}).Stage(stage)", a_attribute_definition_string_refIdent, __gong__toRawStringLiteral(a_attribute_definition_string_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_string_refIdent, __gong__toRawStringLiteral(a_attribute_definition_string_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_STRING_REF = %s", a_attribute_definition_string_refIdent, __gong__toRawStringLiteral(a_attribute_definition_string_ref.ATTRIBUTE_DEFINITION_STRING_REF)))
		}
	}
	if stage != nil {
		for _, a_attribute_definition_xhtml_ref := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs, stage.A_ATTRIBUTE_DEFINITION_XHTML_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_definition_xhtml_refIdent := localAliasPrefix + a_attribute_definition_xhtml_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_DEFINITION_XHTML_REF{Name: %s}).Stage(stage)", a_attribute_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_attribute_definition_xhtml_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_attribute_definition_xhtml_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_XHTML_REF = %s", a_attribute_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_attribute_definition_xhtml_ref.ATTRIBUTE_DEFINITION_XHTML_REF)))
		}
	}
	if stage != nil {
		for _, a_attribute_value_boolean := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_BOOLEANs, stage.A_ATTRIBUTE_VALUE_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_booleanIdent := localAliasPrefix + a_attribute_value_boolean.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_BOOLEAN{Name: %s}).Stage(stage)", a_attribute_value_booleanIdent, __gong__toRawStringLiteral(a_attribute_value_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_booleanIdent, __gong__toRawStringLiteral(a_attribute_value_boolean.Name)))
			for _, elem := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_BOOLEAN = append(%s.ATTRIBUTE_VALUE_BOOLEAN, %s)", a_attribute_value_booleanIdent, a_attribute_value_booleanIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_value_date := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_DATEs, stage.A_ATTRIBUTE_VALUE_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_dateIdent := localAliasPrefix + a_attribute_value_date.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_DATE{Name: %s}).Stage(stage)", a_attribute_value_dateIdent, __gong__toRawStringLiteral(a_attribute_value_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_dateIdent, __gong__toRawStringLiteral(a_attribute_value_date.Name)))
			for _, elem := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_DATE = append(%s.ATTRIBUTE_VALUE_DATE, %s)", a_attribute_value_dateIdent, a_attribute_value_dateIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_value_enumeration := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs, stage.A_ATTRIBUTE_VALUE_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_enumerationIdent := localAliasPrefix + a_attribute_value_enumeration.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_ENUMERATION{Name: %s}).Stage(stage)", a_attribute_value_enumerationIdent, __gong__toRawStringLiteral(a_attribute_value_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_enumerationIdent, __gong__toRawStringLiteral(a_attribute_value_enumeration.Name)))
			for _, elem := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_ENUMERATION = append(%s.ATTRIBUTE_VALUE_ENUMERATION, %s)", a_attribute_value_enumerationIdent, a_attribute_value_enumerationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_value_integer := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_INTEGERs, stage.A_ATTRIBUTE_VALUE_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_integerIdent := localAliasPrefix + a_attribute_value_integer.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_INTEGER{Name: %s}).Stage(stage)", a_attribute_value_integerIdent, __gong__toRawStringLiteral(a_attribute_value_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_integerIdent, __gong__toRawStringLiteral(a_attribute_value_integer.Name)))
			for _, elem := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_INTEGER = append(%s.ATTRIBUTE_VALUE_INTEGER, %s)", a_attribute_value_integerIdent, a_attribute_value_integerIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_value_real := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_REALs, stage.A_ATTRIBUTE_VALUE_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_realIdent := localAliasPrefix + a_attribute_value_real.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_REAL{Name: %s}).Stage(stage)", a_attribute_value_realIdent, __gong__toRawStringLiteral(a_attribute_value_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_realIdent, __gong__toRawStringLiteral(a_attribute_value_real.Name)))
			for _, elem := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_REAL = append(%s.ATTRIBUTE_VALUE_REAL, %s)", a_attribute_value_realIdent, a_attribute_value_realIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_value_string := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_STRINGs, stage.A_ATTRIBUTE_VALUE_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_stringIdent := localAliasPrefix + a_attribute_value_string.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_STRING{Name: %s}).Stage(stage)", a_attribute_value_stringIdent, __gong__toRawStringLiteral(a_attribute_value_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_stringIdent, __gong__toRawStringLiteral(a_attribute_value_string.Name)))
			for _, elem := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_STRING = append(%s.ATTRIBUTE_VALUE_STRING, %s)", a_attribute_value_stringIdent, a_attribute_value_stringIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_value_xhtml := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_XHTMLs, stage.A_ATTRIBUTE_VALUE_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_xhtmlIdent := localAliasPrefix + a_attribute_value_xhtml.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_XHTML{Name: %s}).Stage(stage)", a_attribute_value_xhtmlIdent, __gong__toRawStringLiteral(a_attribute_value_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_xhtmlIdent, __gong__toRawStringLiteral(a_attribute_value_xhtml.Name)))
			for _, elem := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_XHTML = append(%s.ATTRIBUTE_VALUE_XHTML, %s)", a_attribute_value_xhtmlIdent, a_attribute_value_xhtmlIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_attribute_value_xhtml_1 := range __gong__sortStageSetInstances(stage.A_ATTRIBUTE_VALUE_XHTML_1s, stage.A_ATTRIBUTE_VALUE_XHTML_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_attribute_value_xhtml_1Ident := localAliasPrefix + a_attribute_value_xhtml_1.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ATTRIBUTE_VALUE_XHTML_1{Name: %s}).Stage(stage)", a_attribute_value_xhtml_1Ident, __gong__toRawStringLiteral(a_attribute_value_xhtml_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_attribute_value_xhtml_1Ident, __gong__toRawStringLiteral(a_attribute_value_xhtml_1.Name)))
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_BOOLEAN = append(%s.ATTRIBUTE_VALUE_BOOLEAN, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_DATE = append(%s.ATTRIBUTE_VALUE_DATE, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_ENUMERATION = append(%s.ATTRIBUTE_VALUE_ENUMERATION, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_INTEGER = append(%s.ATTRIBUTE_VALUE_INTEGER, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_REAL = append(%s.ATTRIBUTE_VALUE_REAL, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_STRING = append(%s.ATTRIBUTE_VALUE_STRING, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
			for _, elem := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_VALUE_XHTML = append(%s.ATTRIBUTE_VALUE_XHTML, %s)", a_attribute_value_xhtml_1Ident, a_attribute_value_xhtml_1Ident, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_children := range __gong__sortStageSetInstances(stage.A_CHILDRENs, stage.A_CHILDREN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_childrenIdent := localAliasPrefix + a_children.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_CHILDREN{Name: %s}).Stage(stage)", a_childrenIdent, __gong__toRawStringLiteral(a_children.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_childrenIdent, __gong__toRawStringLiteral(a_children.Name)))
			for _, elem := range a_children.SPEC_HIERARCHY {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_HIERARCHY = append(%s.SPEC_HIERARCHY, %s)", a_childrenIdent, a_childrenIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_core_content := range __gong__sortStageSetInstances(stage.A_CORE_CONTENTs, stage.A_CORE_CONTENT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_core_contentIdent := localAliasPrefix + a_core_content.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_CORE_CONTENT{Name: %s}).Stage(stage)", a_core_contentIdent, __gong__toRawStringLiteral(a_core_content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_core_contentIdent, __gong__toRawStringLiteral(a_core_content.Name)))
			if a_core_content.REQ_IF_CONTENT != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + a_core_content.REQ_IF_CONTENT.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_CONTENT = %s", a_core_contentIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_datatypes := range __gong__sortStageSetInstances(stage.A_DATATYPESs, stage.A_DATATYPES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatypesIdent := localAliasPrefix + a_datatypes.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPES{Name: %s}).Stage(stage)", a_datatypesIdent, __gong__toRawStringLiteral(a_datatypes.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatypesIdent, __gong__toRawStringLiteral(a_datatypes.Name)))
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_BOOLEAN = append(%s.DATATYPE_DEFINITION_BOOLEAN, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_DATE = append(%s.DATATYPE_DEFINITION_DATE, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_ENUMERATION = append(%s.DATATYPE_DEFINITION_ENUMERATION, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_INTEGER = append(%s.DATATYPE_DEFINITION_INTEGER, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_REAL = append(%s.DATATYPE_DEFINITION_REAL, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_STRING = append(%s.DATATYPE_DEFINITION_STRING, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
			for _, elem := range a_datatypes.DATATYPE_DEFINITION_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_XHTML = append(%s.DATATYPE_DEFINITION_XHTML, %s)", a_datatypesIdent, a_datatypesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_datatype_definition_boolean_ref := range __gong__sortStageSetInstances(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs, stage.A_DATATYPE_DEFINITION_BOOLEAN_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_boolean_refIdent := localAliasPrefix + a_datatype_definition_boolean_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_BOOLEAN_REF{Name: %s}).Stage(stage)", a_datatype_definition_boolean_refIdent, __gong__toRawStringLiteral(a_datatype_definition_boolean_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_boolean_refIdent, __gong__toRawStringLiteral(a_datatype_definition_boolean_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_BOOLEAN_REF = %s", a_datatype_definition_boolean_refIdent, __gong__toRawStringLiteral(a_datatype_definition_boolean_ref.DATATYPE_DEFINITION_BOOLEAN_REF)))
		}
	}
	if stage != nil {
		for _, a_datatype_definition_date_ref := range __gong__sortStageSetInstances(stage.A_DATATYPE_DEFINITION_DATE_REFs, stage.A_DATATYPE_DEFINITION_DATE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_date_refIdent := localAliasPrefix + a_datatype_definition_date_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_DATE_REF{Name: %s}).Stage(stage)", a_datatype_definition_date_refIdent, __gong__toRawStringLiteral(a_datatype_definition_date_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_date_refIdent, __gong__toRawStringLiteral(a_datatype_definition_date_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_DATE_REF = %s", a_datatype_definition_date_refIdent, __gong__toRawStringLiteral(a_datatype_definition_date_ref.DATATYPE_DEFINITION_DATE_REF)))
		}
	}
	if stage != nil {
		for _, a_datatype_definition_enumeration_ref := range __gong__sortStageSetInstances(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs, stage.A_DATATYPE_DEFINITION_ENUMERATION_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_enumeration_refIdent := localAliasPrefix + a_datatype_definition_enumeration_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_ENUMERATION_REF{Name: %s}).Stage(stage)", a_datatype_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_datatype_definition_enumeration_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_datatype_definition_enumeration_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_ENUMERATION_REF = %s", a_datatype_definition_enumeration_refIdent, __gong__toRawStringLiteral(a_datatype_definition_enumeration_ref.DATATYPE_DEFINITION_ENUMERATION_REF)))
		}
	}
	if stage != nil {
		for _, a_datatype_definition_integer_ref := range __gong__sortStageSetInstances(stage.A_DATATYPE_DEFINITION_INTEGER_REFs, stage.A_DATATYPE_DEFINITION_INTEGER_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_integer_refIdent := localAliasPrefix + a_datatype_definition_integer_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_INTEGER_REF{Name: %s}).Stage(stage)", a_datatype_definition_integer_refIdent, __gong__toRawStringLiteral(a_datatype_definition_integer_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_integer_refIdent, __gong__toRawStringLiteral(a_datatype_definition_integer_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_INTEGER_REF = %s", a_datatype_definition_integer_refIdent, __gong__toRawStringLiteral(a_datatype_definition_integer_ref.DATATYPE_DEFINITION_INTEGER_REF)))
		}
	}
	if stage != nil {
		for _, a_datatype_definition_real_ref := range __gong__sortStageSetInstances(stage.A_DATATYPE_DEFINITION_REAL_REFs, stage.A_DATATYPE_DEFINITION_REAL_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_real_refIdent := localAliasPrefix + a_datatype_definition_real_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_REAL_REF{Name: %s}).Stage(stage)", a_datatype_definition_real_refIdent, __gong__toRawStringLiteral(a_datatype_definition_real_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_real_refIdent, __gong__toRawStringLiteral(a_datatype_definition_real_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_REAL_REF = %s", a_datatype_definition_real_refIdent, __gong__toRawStringLiteral(a_datatype_definition_real_ref.DATATYPE_DEFINITION_REAL_REF)))
		}
	}
	if stage != nil {
		for _, a_datatype_definition_string_ref := range __gong__sortStageSetInstances(stage.A_DATATYPE_DEFINITION_STRING_REFs, stage.A_DATATYPE_DEFINITION_STRING_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_string_refIdent := localAliasPrefix + a_datatype_definition_string_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_STRING_REF{Name: %s}).Stage(stage)", a_datatype_definition_string_refIdent, __gong__toRawStringLiteral(a_datatype_definition_string_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_string_refIdent, __gong__toRawStringLiteral(a_datatype_definition_string_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_STRING_REF = %s", a_datatype_definition_string_refIdent, __gong__toRawStringLiteral(a_datatype_definition_string_ref.DATATYPE_DEFINITION_STRING_REF)))
		}
	}
	if stage != nil {
		for _, a_datatype_definition_xhtml_ref := range __gong__sortStageSetInstances(stage.A_DATATYPE_DEFINITION_XHTML_REFs, stage.A_DATATYPE_DEFINITION_XHTML_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_datatype_definition_xhtml_refIdent := localAliasPrefix + a_datatype_definition_xhtml_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_DATATYPE_DEFINITION_XHTML_REF{Name: %s}).Stage(stage)", a_datatype_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_datatype_definition_xhtml_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_datatype_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_datatype_definition_xhtml_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DATATYPE_DEFINITION_XHTML_REF = %s", a_datatype_definition_xhtml_refIdent, __gong__toRawStringLiteral(a_datatype_definition_xhtml_ref.DATATYPE_DEFINITION_XHTML_REF)))
		}
	}
	if stage != nil {
		for _, a_editable_atts := range __gong__sortStageSetInstances(stage.A_EDITABLE_ATTSs, stage.A_EDITABLE_ATTS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_editable_attsIdent := localAliasPrefix + a_editable_atts.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_EDITABLE_ATTS{Name: %s}).Stage(stage)", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_BOOLEAN_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_BOOLEAN_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_DATE_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_DATE_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_ENUMERATION_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_ENUMERATION_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_INTEGER_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_INTEGER_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_REAL_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_REAL_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_STRING_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_STRING_REF)))
			values.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_XHTML_REF = %s", a_editable_attsIdent, __gong__toRawStringLiteral(a_editable_atts.ATTRIBUTE_DEFINITION_XHTML_REF)))
		}
	}
	if stage != nil {
		for _, a_enum_value_ref := range __gong__sortStageSetInstances(stage.A_ENUM_VALUE_REFs, stage.A_ENUM_VALUE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_enum_value_refIdent := localAliasPrefix + a_enum_value_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_ENUM_VALUE_REF{Name: %s}).Stage(stage)", a_enum_value_refIdent, __gong__toRawStringLiteral(a_enum_value_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_enum_value_refIdent, __gong__toRawStringLiteral(a_enum_value_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ENUM_VALUE_REF = %s", a_enum_value_refIdent, __gong__toRawStringLiteral(a_enum_value_ref.ENUM_VALUE_REF)))
		}
	}
	if stage != nil {
		for _, a_object := range __gong__sortStageSetInstances(stage.A_OBJECTs, stage.A_OBJECT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_objectIdent := localAliasPrefix + a_object.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_OBJECT{Name: %s}).Stage(stage)", a_objectIdent, __gong__toRawStringLiteral(a_object.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_objectIdent, __gong__toRawStringLiteral(a_object.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_REF = %s", a_objectIdent, __gong__toRawStringLiteral(a_object.SPEC_OBJECT_REF)))
		}
	}
	if stage != nil {
		for _, a_properties := range __gong__sortStageSetInstances(stage.A_PROPERTIESs, stage.A_PROPERTIES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_propertiesIdent := localAliasPrefix + a_properties.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_PROPERTIES{Name: %s}).Stage(stage)", a_propertiesIdent, __gong__toRawStringLiteral(a_properties.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_propertiesIdent, __gong__toRawStringLiteral(a_properties.Name)))
			if a_properties.EMBEDDED_VALUE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + a_properties.EMBEDDED_VALUE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EMBEDDED_VALUE = %s", a_propertiesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_relation_group_type_ref := range __gong__sortStageSetInstances(stage.A_RELATION_GROUP_TYPE_REFs, stage.A_RELATION_GROUP_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_relation_group_type_refIdent := localAliasPrefix + a_relation_group_type_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_RELATION_GROUP_TYPE_REF{Name: %s}).Stage(stage)", a_relation_group_type_refIdent, __gong__toRawStringLiteral(a_relation_group_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_relation_group_type_refIdent, __gong__toRawStringLiteral(a_relation_group_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.RELATION_GROUP_TYPE_REF = %s", a_relation_group_type_refIdent, __gong__toRawStringLiteral(a_relation_group_type_ref.RELATION_GROUP_TYPE_REF)))
		}
	}
	if stage != nil {
		for _, a_source_1 := range __gong__sortStageSetInstances(stage.A_SOURCE_1s, stage.A_SOURCE_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_source_1Ident := localAliasPrefix + a_source_1.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SOURCE_1{Name: %s}).Stage(stage)", a_source_1Ident, __gong__toRawStringLiteral(a_source_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_source_1Ident, __gong__toRawStringLiteral(a_source_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_REF = %s", a_source_1Ident, __gong__toRawStringLiteral(a_source_1.SPEC_OBJECT_REF)))
		}
	}
	if stage != nil {
		for _, a_source_specification_1 := range __gong__sortStageSetInstances(stage.A_SOURCE_SPECIFICATION_1s, stage.A_SOURCE_SPECIFICATION_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_source_specification_1Ident := localAliasPrefix + a_source_specification_1.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SOURCE_SPECIFICATION_1{Name: %s}).Stage(stage)", a_source_specification_1Ident, __gong__toRawStringLiteral(a_source_specification_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_source_specification_1Ident, __gong__toRawStringLiteral(a_source_specification_1.Name)))
			if a_source_specification_1.SPECIFICATION_REF.ToCodeString() != "" {
				values.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION_REF = models.%s", a_source_specification_1Ident, a_source_specification_1.SPECIFICATION_REF.ToCodeString()))
			} else {
				values.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION_REF = \"\"", a_source_specification_1Ident))
			}
		}
	}
	if stage != nil {
		for _, a_specifications := range __gong__sortStageSetInstances(stage.A_SPECIFICATIONSs, stage.A_SPECIFICATIONS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_specificationsIdent := localAliasPrefix + a_specifications.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPECIFICATIONS{Name: %s}).Stage(stage)", a_specificationsIdent, __gong__toRawStringLiteral(a_specifications.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_specificationsIdent, __gong__toRawStringLiteral(a_specifications.Name)))
			for _, elem := range a_specifications.SPECIFICATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION = append(%s.SPECIFICATION, %s)", a_specificationsIdent, a_specificationsIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_specification_type_ref := range __gong__sortStageSetInstances(stage.A_SPECIFICATION_TYPE_REFs, stage.A_SPECIFICATION_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_specification_type_refIdent := localAliasPrefix + a_specification_type_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPECIFICATION_TYPE_REF{Name: %s}).Stage(stage)", a_specification_type_refIdent, __gong__toRawStringLiteral(a_specification_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_specification_type_refIdent, __gong__toRawStringLiteral(a_specification_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION_TYPE_REF = %s", a_specification_type_refIdent, __gong__toRawStringLiteral(a_specification_type_ref.SPECIFICATION_TYPE_REF)))
		}
	}
	if stage != nil {
		for _, a_specified_values := range __gong__sortStageSetInstances(stage.A_SPECIFIED_VALUESs, stage.A_SPECIFIED_VALUES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_specified_valuesIdent := localAliasPrefix + a_specified_values.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPECIFIED_VALUES{Name: %s}).Stage(stage)", a_specified_valuesIdent, __gong__toRawStringLiteral(a_specified_values.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_specified_valuesIdent, __gong__toRawStringLiteral(a_specified_values.Name)))
			for _, elem := range a_specified_values.ENUM_VALUE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ENUM_VALUE = append(%s.ENUM_VALUE, %s)", a_specified_valuesIdent, a_specified_valuesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_spec_attributes := range __gong__sortStageSetInstances(stage.A_SPEC_ATTRIBUTESs, stage.A_SPEC_ATTRIBUTES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_attributesIdent := localAliasPrefix + a_spec_attributes.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_ATTRIBUTES{Name: %s}).Stage(stage)", a_spec_attributesIdent, __gong__toRawStringLiteral(a_spec_attributes.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_attributesIdent, __gong__toRawStringLiteral(a_spec_attributes.Name)))
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_BOOLEAN = append(%s.ATTRIBUTE_DEFINITION_BOOLEAN, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_DATE = append(%s.ATTRIBUTE_DEFINITION_DATE, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_ENUMERATION = append(%s.ATTRIBUTE_DEFINITION_ENUMERATION, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_INTEGER = append(%s.ATTRIBUTE_DEFINITION_INTEGER, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_REAL = append(%s.ATTRIBUTE_DEFINITION_REAL, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_STRING = append(%s.ATTRIBUTE_DEFINITION_STRING, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
			for _, elem := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ATTRIBUTE_DEFINITION_XHTML = append(%s.ATTRIBUTE_DEFINITION_XHTML, %s)", a_spec_attributesIdent, a_spec_attributesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_spec_objects := range __gong__sortStageSetInstances(stage.A_SPEC_OBJECTSs, stage.A_SPEC_OBJECTS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_objectsIdent := localAliasPrefix + a_spec_objects.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_OBJECTS{Name: %s}).Stage(stage)", a_spec_objectsIdent, __gong__toRawStringLiteral(a_spec_objects.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_objectsIdent, __gong__toRawStringLiteral(a_spec_objects.Name)))
			for _, elem := range a_spec_objects.SPEC_OBJECT {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT = append(%s.SPEC_OBJECT, %s)", a_spec_objectsIdent, a_spec_objectsIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_spec_object_type_ref := range __gong__sortStageSetInstances(stage.A_SPEC_OBJECT_TYPE_REFs, stage.A_SPEC_OBJECT_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_object_type_refIdent := localAliasPrefix + a_spec_object_type_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_OBJECT_TYPE_REF{Name: %s}).Stage(stage)", a_spec_object_type_refIdent, __gong__toRawStringLiteral(a_spec_object_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_object_type_refIdent, __gong__toRawStringLiteral(a_spec_object_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_TYPE_REF = %s", a_spec_object_type_refIdent, __gong__toRawStringLiteral(a_spec_object_type_ref.SPEC_OBJECT_TYPE_REF)))
		}
	}
	if stage != nil {
		for _, a_spec_relations := range __gong__sortStageSetInstances(stage.A_SPEC_RELATIONSs, stage.A_SPEC_RELATIONS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relationsIdent := localAliasPrefix + a_spec_relations.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATIONS{Name: %s}).Stage(stage)", a_spec_relationsIdent, __gong__toRawStringLiteral(a_spec_relations.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relationsIdent, __gong__toRawStringLiteral(a_spec_relations.Name)))
			for _, elem := range a_spec_relations.SPEC_RELATION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION = append(%s.SPEC_RELATION, %s)", a_spec_relationsIdent, a_spec_relationsIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_spec_relation_groups := range __gong__sortStageSetInstances(stage.A_SPEC_RELATION_GROUPSs, stage.A_SPEC_RELATION_GROUPS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relation_groupsIdent := localAliasPrefix + a_spec_relation_groups.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATION_GROUPS{Name: %s}).Stage(stage)", a_spec_relation_groupsIdent, __gong__toRawStringLiteral(a_spec_relation_groups.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relation_groupsIdent, __gong__toRawStringLiteral(a_spec_relation_groups.Name)))
			for _, elem := range a_spec_relation_groups.RELATION_GROUP {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RELATION_GROUP = append(%s.RELATION_GROUP, %s)", a_spec_relation_groupsIdent, a_spec_relation_groupsIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_spec_relation_ref := range __gong__sortStageSetInstances(stage.A_SPEC_RELATION_REFs, stage.A_SPEC_RELATION_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relation_refIdent := localAliasPrefix + a_spec_relation_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATION_REF{Name: %s}).Stage(stage)", a_spec_relation_refIdent, __gong__toRawStringLiteral(a_spec_relation_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relation_refIdent, __gong__toRawStringLiteral(a_spec_relation_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_REF = %s", a_spec_relation_refIdent, __gong__toRawStringLiteral(a_spec_relation_ref.SPEC_RELATION_REF)))
		}
	}
	if stage != nil {
		for _, a_spec_relation_type_ref := range __gong__sortStageSetInstances(stage.A_SPEC_RELATION_TYPE_REFs, stage.A_SPEC_RELATION_TYPE_REF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_relation_type_refIdent := localAliasPrefix + a_spec_relation_type_ref.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_RELATION_TYPE_REF{Name: %s}).Stage(stage)", a_spec_relation_type_refIdent, __gong__toRawStringLiteral(a_spec_relation_type_ref.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_relation_type_refIdent, __gong__toRawStringLiteral(a_spec_relation_type_ref.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_TYPE_REF = %s", a_spec_relation_type_refIdent, __gong__toRawStringLiteral(a_spec_relation_type_ref.SPEC_RELATION_TYPE_REF)))
		}
	}
	if stage != nil {
		for _, a_spec_types := range __gong__sortStageSetInstances(stage.A_SPEC_TYPESs, stage.A_SPEC_TYPES_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_spec_typesIdent := localAliasPrefix + a_spec_types.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_SPEC_TYPES{Name: %s}).Stage(stage)", a_spec_typesIdent, __gong__toRawStringLiteral(a_spec_types.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_spec_typesIdent, __gong__toRawStringLiteral(a_spec_types.Name)))
			for _, elem := range a_spec_types.RELATION_GROUP_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.RELATION_GROUP_TYPE = append(%s.RELATION_GROUP_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
			for _, elem := range a_spec_types.SPEC_OBJECT_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECT_TYPE = append(%s.SPEC_OBJECT_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
			for _, elem := range a_spec_types.SPEC_RELATION_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_TYPE = append(%s.SPEC_RELATION_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
			for _, elem := range a_spec_types.SPECIFICATION_TYPE {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATION_TYPE = append(%s.SPECIFICATION_TYPE, %s)", a_spec_typesIdent, a_spec_typesIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_the_header := range __gong__sortStageSetInstances(stage.A_THE_HEADERs, stage.A_THE_HEADER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_the_headerIdent := localAliasPrefix + a_the_header.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_THE_HEADER{Name: %s}).Stage(stage)", a_the_headerIdent, __gong__toRawStringLiteral(a_the_header.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_the_headerIdent, __gong__toRawStringLiteral(a_the_header.Name)))
			if a_the_header.REQ_IF_HEADER != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + a_the_header.REQ_IF_HEADER.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_HEADER = %s", a_the_headerIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, a_tool_extensions := range __gong__sortStageSetInstances(stage.A_TOOL_EXTENSIONSs, stage.A_TOOL_EXTENSIONS_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_tool_extensionsIdent := localAliasPrefix + a_tool_extensions.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.A_TOOL_EXTENSIONS{Name: %s}).Stage(stage)", a_tool_extensionsIdent, __gong__toRawStringLiteral(a_tool_extensions.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_tool_extensionsIdent, __gong__toRawStringLiteral(a_tool_extensions.Name)))
			for _, elem := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_TOOL_EXTENSION = append(%s.REQ_IF_TOOL_EXTENSION, %s)", a_tool_extensionsIdent, a_tool_extensionsIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datatype_definition_boolean := range __gong__sortStageSetInstances(stage.DATATYPE_DEFINITION_BOOLEANs, stage.DATATYPE_DEFINITION_BOOLEAN_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_booleanIdent := localAliasPrefix + datatype_definition_boolean.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_BOOLEAN{Name: %s}).Stage(stage)", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_booleanIdent, __gong__toRawStringLiteral(datatype_definition_boolean.LONG_NAME)))
			if datatype_definition_boolean.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_booleanIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datatype_definition_date := range __gong__sortStageSetInstances(stage.DATATYPE_DEFINITION_DATEs, stage.DATATYPE_DEFINITION_DATE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_dateIdent := localAliasPrefix + datatype_definition_date.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_DATE{Name: %s}).Stage(stage)", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_dateIdent, __gong__toRawStringLiteral(datatype_definition_date.LONG_NAME)))
			if datatype_definition_date.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_dateIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datatype_definition_enumeration := range __gong__sortStageSetInstances(stage.DATATYPE_DEFINITION_ENUMERATIONs, stage.DATATYPE_DEFINITION_ENUMERATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_enumerationIdent := localAliasPrefix + datatype_definition_enumeration.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_ENUMERATION{Name: %s}).Stage(stage)", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_enumerationIdent, __gong__toRawStringLiteral(datatype_definition_enumeration.LONG_NAME)))
			if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_enumerationIdent, targetIdent))
			}
			if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_enumeration.SPECIFIED_VALUES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFIED_VALUES = %s", datatype_definition_enumerationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datatype_definition_integer := range __gong__sortStageSetInstances(stage.DATATYPE_DEFINITION_INTEGERs, stage.DATATYPE_DEFINITION_INTEGER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_integerIdent := localAliasPrefix + datatype_definition_integer.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_INTEGER{Name: %s}).Stage(stage)", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_integerIdent, __gong__toRawStringLiteral(datatype_definition_integer.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MAX = %d", datatype_definition_integerIdent, datatype_definition_integer.MAX))
			values.WriteString(fmt.Sprintf("\n\t%s.MIN = %d", datatype_definition_integerIdent, datatype_definition_integer.MIN))
			if datatype_definition_integer.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_integerIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datatype_definition_real := range __gong__sortStageSetInstances(stage.DATATYPE_DEFINITION_REALs, stage.DATATYPE_DEFINITION_REAL_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_realIdent := localAliasPrefix + datatype_definition_real.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_REAL{Name: %s}).Stage(stage)", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.ACCURACY = %d", datatype_definition_realIdent, datatype_definition_real.ACCURACY))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_realIdent, __gong__toRawStringLiteral(datatype_definition_real.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MAX = %f", datatype_definition_realIdent, datatype_definition_real.MAX))
			values.WriteString(fmt.Sprintf("\n\t%s.MIN = %f", datatype_definition_realIdent, datatype_definition_real.MIN))
			if datatype_definition_real.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_realIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datatype_definition_string := range __gong__sortStageSetInstances(stage.DATATYPE_DEFINITION_STRINGs, stage.DATATYPE_DEFINITION_STRING_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_stringIdent := localAliasPrefix + datatype_definition_string.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_STRING{Name: %s}).Stage(stage)", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_stringIdent, __gong__toRawStringLiteral(datatype_definition_string.LONG_NAME)))
			values.WriteString(fmt.Sprintf("\n\t%s.MAX_LENGTH = %d", datatype_definition_stringIdent, datatype_definition_string.MAX_LENGTH))
			if datatype_definition_string.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_stringIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, datatype_definition_xhtml := range __gong__sortStageSetInstances(stage.DATATYPE_DEFINITION_XHTMLs, stage.DATATYPE_DEFINITION_XHTML_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			datatype_definition_xhtmlIdent := localAliasPrefix + datatype_definition_xhtml.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.DATATYPE_DEFINITION_XHTML{Name: %s}).Stage(stage)", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", datatype_definition_xhtmlIdent, __gong__toRawStringLiteral(datatype_definition_xhtml.LONG_NAME)))
			if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + datatype_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", datatype_definition_xhtmlIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, embedded_value := range __gong__sortStageSetInstances(stage.EMBEDDED_VALUEs, stage.EMBEDDED_VALUE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embedded_valueIdent := localAliasPrefix + embedded_value.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EMBEDDED_VALUE{Name: %s}).Stage(stage)", embedded_valueIdent, __gong__toRawStringLiteral(embedded_value.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embedded_valueIdent, __gong__toRawStringLiteral(embedded_value.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.KEY = %d", embedded_valueIdent, embedded_value.KEY))
			values.WriteString(fmt.Sprintf("\n\t%s.OTHER_CONTENT = %s", embedded_valueIdent, __gong__toRawStringLiteral(embedded_value.OTHER_CONTENT)))
		}
	}
	if stage != nil {
		for _, enum_value := range __gong__sortStageSetInstances(stage.ENUM_VALUEs, stage.ENUM_VALUE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			enum_valueIdent := localAliasPrefix + enum_value.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.ENUM_VALUE{Name: %s}).Stage(stage)", enum_valueIdent, __gong__toRawStringLiteral(enum_value.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", enum_valueIdent, __gong__toRawStringLiteral(enum_value.LONG_NAME)))
			if enum_value.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + enum_value.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", enum_valueIdent, targetIdent))
			}
			if enum_value.PROPERTIES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + enum_value.PROPERTIES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.PROPERTIES = %s", enum_valueIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, embeddedjpgimage := range __gong__sortStageSetInstances(stage.EmbeddedJpgImages, stage.EmbeddedJpgImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embeddedjpgimageIdent := localAliasPrefix + embeddedjpgimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EmbeddedJpgImage{Name: %s}).Stage(stage)", embeddedjpgimageIdent, __gong__toRawStringLiteral(embeddedjpgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embeddedjpgimageIdent, __gong__toRawStringLiteral(embeddedjpgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", embeddedjpgimageIdent, __gong__toRawStringLiteral(embeddedjpgimage.Base64Content)))
		}
	}
	if stage != nil {
		for _, embeddedpngimage := range __gong__sortStageSetInstances(stage.EmbeddedPngImages, stage.EmbeddedPngImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embeddedpngimageIdent := localAliasPrefix + embeddedpngimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EmbeddedPngImage{Name: %s}).Stage(stage)", embeddedpngimageIdent, __gong__toRawStringLiteral(embeddedpngimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embeddedpngimageIdent, __gong__toRawStringLiteral(embeddedpngimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Base64Content = %s", embeddedpngimageIdent, __gong__toRawStringLiteral(embeddedpngimage.Base64Content)))
		}
	}
	if stage != nil {
		for _, embeddedsvgimage := range __gong__sortStageSetInstances(stage.EmbeddedSvgImages, stage.EmbeddedSvgImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			embeddedsvgimageIdent := localAliasPrefix + embeddedsvgimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.EmbeddedSvgImage{Name: %s}).Stage(stage)", embeddedsvgimageIdent, __gong__toRawStringLiteral(embeddedsvgimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", embeddedsvgimageIdent, __gong__toRawStringLiteral(embeddedsvgimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Content = %s", embeddedsvgimageIdent, __gong__toRawStringLiteral(embeddedsvgimage.Content)))
		}
	}
	if stage != nil {
		for _, kill := range __gong__sortStageSetInstances(stage.Kills, stage.Kill_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			killIdent := localAliasPrefix + kill.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Kill{Name: %s}).Stage(stage)", killIdent, __gong__toRawStringLiteral(kill.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", killIdent, __gong__toRawStringLiteral(kill.Name)))
		}
	}
	if stage != nil {
		for _, map_identifier_bool := range __gong__sortStageSetInstances(stage.Map_identifier_bools, stage.Map_identifier_bool_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			map_identifier_boolIdent := localAliasPrefix + map_identifier_bool.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.Map_identifier_bool{Name: %s}).Stage(stage)", map_identifier_boolIdent, __gong__toRawStringLiteral(map_identifier_bool.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", map_identifier_boolIdent, __gong__toRawStringLiteral(map_identifier_bool.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Value = %t", map_identifier_boolIdent, map_identifier_bool.Value))
		}
	}
	if stage != nil {
		for _, relation_group := range __gong__sortStageSetInstances(stage.RELATION_GROUPs, stage.RELATION_GROUP_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			relation_groupIdent := localAliasPrefix + relation_group.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RELATION_GROUP{Name: %s}).Stage(stage)", relation_groupIdent, __gong__toRawStringLiteral(relation_group.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", relation_groupIdent, __gong__toRawStringLiteral(relation_group.LONG_NAME)))
			if relation_group.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + relation_group.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.SOURCE_SPECIFICATION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + relation_group.SOURCE_SPECIFICATION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SOURCE_SPECIFICATION = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.SPEC_RELATIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + relation_group.SPEC_RELATIONS.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATIONS = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.TARGET_SPECIFICATION != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + relation_group.TARGET_SPECIFICATION.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TARGET_SPECIFICATION = %s", relation_groupIdent, targetIdent))
			}
			if relation_group.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + relation_group.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", relation_groupIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, relation_group_type := range __gong__sortStageSetInstances(stage.RELATION_GROUP_TYPEs, stage.RELATION_GROUP_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			relation_group_typeIdent := localAliasPrefix + relation_group_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.RELATION_GROUP_TYPE{Name: %s}).Stage(stage)", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", relation_group_typeIdent, __gong__toRawStringLiteral(relation_group_type.LONG_NAME)))
			if relation_group_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + relation_group_type.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", relation_group_typeIdent, targetIdent))
			}
			if relation_group_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + relation_group_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", relation_group_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, req_if := range __gong__sortStageSetInstances(stage.REQ_IFs, stage.REQ_IF_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_ifIdent := localAliasPrefix + req_if.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF{Name: %s}).Stage(stage)", req_ifIdent, __gong__toRawStringLiteral(req_if.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_ifIdent, __gong__toRawStringLiteral(req_if.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Lang = %s", req_ifIdent, __gong__toRawStringLiteral(req_if.Lang)))
			if req_if.THE_HEADER != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if.THE_HEADER.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.THE_HEADER = %s", req_ifIdent, targetIdent))
			}
			if req_if.CORE_CONTENT != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if.CORE_CONTENT.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CORE_CONTENT = %s", req_ifIdent, targetIdent))
			}
			if req_if.TOOL_EXTENSIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if.TOOL_EXTENSIONS.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TOOL_EXTENSIONS = %s", req_ifIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, req_if_content := range __gong__sortStageSetInstances(stage.REQ_IF_CONTENTs, stage.REQ_IF_CONTENT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_if_contentIdent := localAliasPrefix + req_if_content.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF_CONTENT{Name: %s}).Stage(stage)", req_if_contentIdent, __gong__toRawStringLiteral(req_if_content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_if_contentIdent, __gong__toRawStringLiteral(req_if_content.Name)))
			if req_if_content.DATATYPES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if_content.DATATYPES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DATATYPES = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_TYPES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if_content.SPEC_TYPES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_TYPES = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_OBJECTS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if_content.SPEC_OBJECTS.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_OBJECTS = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_RELATIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if_content.SPEC_RELATIONS.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATIONS = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPECIFICATIONS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if_content.SPECIFICATIONS.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPECIFICATIONS = %s", req_if_contentIdent, targetIdent))
			}
			if req_if_content.SPEC_RELATION_GROUPS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + req_if_content.SPEC_RELATION_GROUPS.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_RELATION_GROUPS = %s", req_if_contentIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, req_if_header := range __gong__sortStageSetInstances(stage.REQ_IF_HEADERs, stage.REQ_IF_HEADER_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_if_headerIdent := localAliasPrefix + req_if_header.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF_HEADER{Name: %s}).Stage(stage)", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.COMMENT = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.COMMENT)))
			values.WriteString(fmt.Sprintf("\n\t%s.CREATION_TIME = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.CREATION_TIME)))
			values.WriteString(fmt.Sprintf("\n\t%s.REPOSITORY_ID = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.REPOSITORY_ID)))
			values.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_TOOL_ID = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.REQ_IF_TOOL_ID)))
			values.WriteString(fmt.Sprintf("\n\t%s.REQ_IF_VERSION = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.REQ_IF_VERSION)))
			values.WriteString(fmt.Sprintf("\n\t%s.SOURCE_TOOL_ID = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.SOURCE_TOOL_ID)))
			values.WriteString(fmt.Sprintf("\n\t%s.TITLE = %s", req_if_headerIdent, __gong__toRawStringLiteral(req_if_header.TITLE)))
		}
	}
	if stage != nil {
		for _, req_if_tool_extension := range __gong__sortStageSetInstances(stage.REQ_IF_TOOL_EXTENSIONs, stage.REQ_IF_TOOL_EXTENSION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			req_if_tool_extensionIdent := localAliasPrefix + req_if_tool_extension.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.REQ_IF_TOOL_EXTENSION{Name: %s}).Stage(stage)", req_if_tool_extensionIdent, __gong__toRawStringLiteral(req_if_tool_extension.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", req_if_tool_extensionIdent, __gong__toRawStringLiteral(req_if_tool_extension.Name)))
		}
	}
	if stage != nil {
		for _, specification := range __gong__sortStageSetInstances(stage.SPECIFICATIONs, stage.SPECIFICATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specificationIdent := localAliasPrefix + specification.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPECIFICATION{Name: %s}).Stage(stage)", specificationIdent, __gong__toRawStringLiteral(specification.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specificationIdent, __gong__toRawStringLiteral(specification.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", specificationIdent, __gong__toRawStringLiteral(specification.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", specificationIdent, __gong__toRawStringLiteral(specification.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", specificationIdent, __gong__toRawStringLiteral(specification.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", specificationIdent, __gong__toRawStringLiteral(specification.LONG_NAME)))
			if specification.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + specification.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", specificationIdent, targetIdent))
			}
			if specification.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + specification.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", specificationIdent, targetIdent))
			}
			if specification.CHILDREN != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + specification.CHILDREN.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CHILDREN = %s", specificationIdent, targetIdent))
			}
			if specification.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + specification.VALUES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", specificationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, specification_rendering := range __gong__sortStageSetInstances(stage.SPECIFICATION_Renderings, stage.SPECIFICATION_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specification_renderingIdent := localAliasPrefix + specification_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPECIFICATION_Rendering{Name: %s}).Stage(stage)", specification_renderingIdent, __gong__toRawStringLiteral(specification_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specification_renderingIdent, __gong__toRawStringLiteral(specification_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNodeExpanded = %t", specification_renderingIdent, specification_rendering.IsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.IsSelected = %t", specification_renderingIdent, specification_rendering.IsSelected))
			values.WriteString(fmt.Sprintf("\n\t%s.IsWithHeadingNumbering = %t", specification_renderingIdent, specification_rendering.IsWithHeadingNumbering))
		}
	}
	if stage != nil {
		for _, specification_type := range __gong__sortStageSetInstances(stage.SPECIFICATION_TYPEs, stage.SPECIFICATION_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specification_typeIdent := localAliasPrefix + specification_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPECIFICATION_TYPE{Name: %s}).Stage(stage)", specification_typeIdent, __gong__toRawStringLiteral(specification_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", specification_typeIdent, __gong__toRawStringLiteral(specification_type.LONG_NAME)))
			if specification_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + specification_type.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", specification_typeIdent, targetIdent))
			}
			if specification_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + specification_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", specification_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, spec_hierarchy := range __gong__sortStageSetInstances(stage.SPEC_HIERARCHYs, stage.SPEC_HIERARCHY_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_hierarchyIdent := localAliasPrefix + spec_hierarchy.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_HIERARCHY{Name: %s}).Stage(stage)", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_EDITABLE = %t", spec_hierarchyIdent, spec_hierarchy.IS_EDITABLE))
			values.WriteString(fmt.Sprintf("\n\t%s.IS_TABLE_INTERNAL = %t", spec_hierarchyIdent, spec_hierarchy.IS_TABLE_INTERNAL))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_hierarchyIdent, __gong__toRawStringLiteral(spec_hierarchy.LONG_NAME)))
			if spec_hierarchy.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_hierarchy.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_hierarchyIdent, targetIdent))
			}
			if spec_hierarchy.OBJECT != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_hierarchy.OBJECT.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.OBJECT = %s", spec_hierarchyIdent, targetIdent))
			}
			if spec_hierarchy.CHILDREN != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_hierarchy.CHILDREN.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.CHILDREN = %s", spec_hierarchyIdent, targetIdent))
			}
			if spec_hierarchy.EDITABLE_ATTS != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_hierarchy.EDITABLE_ATTS.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.EDITABLE_ATTS = %s", spec_hierarchyIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, spec_object := range __gong__sortStageSetInstances(stage.SPEC_OBJECTs, stage.SPEC_OBJECT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_objectIdent := localAliasPrefix + spec_object.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_OBJECT{Name: %s}).Stage(stage)", spec_objectIdent, __gong__toRawStringLiteral(spec_object.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_objectIdent, __gong__toRawStringLiteral(spec_object.LONG_NAME)))
			if spec_object.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_object.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_objectIdent, targetIdent))
			}
			if spec_object.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_object.VALUES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", spec_objectIdent, targetIdent))
			}
			if spec_object.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_object.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", spec_objectIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, spec_object_type := range __gong__sortStageSetInstances(stage.SPEC_OBJECT_TYPEs, stage.SPEC_OBJECT_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_object_typeIdent := localAliasPrefix + spec_object_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_OBJECT_TYPE{Name: %s}).Stage(stage)", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_object_typeIdent, __gong__toRawStringLiteral(spec_object_type.LONG_NAME)))
			if spec_object_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_object_type.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_object_typeIdent, targetIdent))
			}
			if spec_object_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_object_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", spec_object_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, spec_object_type_rendering := range __gong__sortStageSetInstances(stage.SPEC_OBJECT_TYPE_Renderings, stage.SPEC_OBJECT_TYPE_Rendering_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_object_type_renderingIdent := localAliasPrefix + spec_object_type_rendering.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_OBJECT_TYPE_Rendering{Name: %s}).Stage(stage)", spec_object_type_renderingIdent, __gong__toRawStringLiteral(spec_object_type_rendering.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_object_type_renderingIdent, __gong__toRawStringLiteral(spec_object_type_rendering.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.IsNodeExpanded = %t", spec_object_type_renderingIdent, spec_object_type_rendering.IsNodeExpanded))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowIdentifier = %t", spec_object_type_renderingIdent, spec_object_type_rendering.ShowIdentifier))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowName = %t", spec_object_type_renderingIdent, spec_object_type_rendering.ShowName))
			values.WriteString(fmt.Sprintf("\n\t%s.ShowRelations = %t", spec_object_type_renderingIdent, spec_object_type_rendering.ShowRelations))
			values.WriteString(fmt.Sprintf("\n\t%s.IsHeading = %t", spec_object_type_renderingIdent, spec_object_type_rendering.IsHeading))
		}
	}
	if stage != nil {
		for _, spec_relation := range __gong__sortStageSetInstances(stage.SPEC_RELATIONs, stage.SPEC_RELATION_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_relationIdent := localAliasPrefix + spec_relation.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_RELATION{Name: %s}).Stage(stage)", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_relationIdent, __gong__toRawStringLiteral(spec_relation.LONG_NAME)))
			if spec_relation.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_relation.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.VALUES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_relation.VALUES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VALUES = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.SOURCE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_relation.SOURCE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SOURCE = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.TARGET != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_relation.TARGET.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TARGET = %s", spec_relationIdent, targetIdent))
			}
			if spec_relation.TYPE != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_relation.TYPE.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TYPE = %s", spec_relationIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, spec_relation_type := range __gong__sortStageSetInstances(stage.SPEC_RELATION_TYPEs, stage.SPEC_RELATION_TYPE_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			spec_relation_typeIdent := localAliasPrefix + spec_relation_type.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.SPEC_RELATION_TYPE{Name: %s}).Stage(stage)", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.DESC = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.DESC)))
			values.WriteString(fmt.Sprintf("\n\t%s.IDENTIFIER = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.IDENTIFIER)))
			values.WriteString(fmt.Sprintf("\n\t%s.LAST_CHANGE = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.LAST_CHANGE)))
			values.WriteString(fmt.Sprintf("\n\t%s.LONG_NAME = %s", spec_relation_typeIdent, __gong__toRawStringLiteral(spec_relation_type.LONG_NAME)))
			if spec_relation_type.ALTERNATIVE_ID != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_relation_type.ALTERNATIVE_ID.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ALTERNATIVE_ID = %s", spec_relation_typeIdent, targetIdent))
			}
			if spec_relation_type.SPEC_ATTRIBUTES != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + spec_relation_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SPEC_ATTRIBUTES = %s", spec_relation_typeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, staticwebsite := range __gong__sortStageSetInstances(stage.StaticWebSites, stage.StaticWebSite_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsiteIdent := localAliasPrefix + staticwebsite.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSite{Name: %s}).Stage(stage)", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MarkdownContent = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.MarkdownContent)))
			values.WriteString(fmt.Sprintf("\n\t%s.InputImagesDir = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.InputImagesDir)))
			values.WriteString(fmt.Sprintf("\n\t%s.OutputStaticWebDir = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.OutputStaticWebDir)))
			values.WriteString(fmt.Sprintf("\n\t%s.VersionInfo = %s", staticwebsiteIdent, __gong__toRawStringLiteral(staticwebsite.VersionInfo)))
			for _, elem := range staticwebsite.Chapters {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Chapters = append(%s.Chapters, %s)", staticwebsiteIdent, staticwebsiteIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, staticwebsitechapter := range __gong__sortStageSetInstances(stage.StaticWebSiteChapters, stage.StaticWebSiteChapter_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsitechapterIdent := localAliasPrefix + staticwebsitechapter.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteChapter{Name: %s}).Stage(stage)", staticwebsitechapterIdent, __gong__toRawStringLiteral(staticwebsitechapter.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsitechapterIdent, __gong__toRawStringLiteral(staticwebsitechapter.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.MarkdownContent = %s", staticwebsitechapterIdent, __gong__toRawStringLiteral(staticwebsitechapter.MarkdownContent)))
			for _, elem := range staticwebsitechapter.Paragraphs {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Paragraphs = append(%s.Paragraphs, %s)", staticwebsitechapterIdent, staticwebsitechapterIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, staticwebsitegeneratedimage := range __gong__sortStageSetInstances(stage.StaticWebSiteGeneratedImages, stage.StaticWebSiteGeneratedImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsitegeneratedimageIdent := localAliasPrefix + staticwebsitegeneratedimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteGeneratedImage{Name: %s}).Stage(stage)", staticwebsitegeneratedimageIdent, __gong__toRawStringLiteral(staticwebsitegeneratedimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsitegeneratedimageIdent, __gong__toRawStringLiteral(staticwebsitegeneratedimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceDirectoryPath = %s", staticwebsitegeneratedimageIdent, __gong__toRawStringLiteral(staticwebsitegeneratedimage.SourceDirectoryPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", staticwebsitegeneratedimageIdent, staticwebsitegeneratedimage.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", staticwebsitegeneratedimageIdent, staticwebsitegeneratedimage.Height))
		}
	}
	if stage != nil {
		for _, staticwebsiteimage := range __gong__sortStageSetInstances(stage.StaticWebSiteImages, stage.StaticWebSiteImage_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsiteimageIdent := localAliasPrefix + staticwebsiteimage.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteImage{Name: %s}).Stage(stage)", staticwebsiteimageIdent, __gong__toRawStringLiteral(staticwebsiteimage.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsiteimageIdent, __gong__toRawStringLiteral(staticwebsiteimage.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.SourceDirectoryPath = %s", staticwebsiteimageIdent, __gong__toRawStringLiteral(staticwebsiteimage.SourceDirectoryPath)))
			values.WriteString(fmt.Sprintf("\n\t%s.Width = %d", staticwebsiteimageIdent, staticwebsiteimage.Width))
			values.WriteString(fmt.Sprintf("\n\t%s.Height = %d", staticwebsiteimageIdent, staticwebsiteimage.Height))
		}
	}
	if stage != nil {
		for _, staticwebsiteparagraph := range __gong__sortStageSetInstances(stage.StaticWebSiteParagraphs, stage.StaticWebSiteParagraph_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			staticwebsiteparagraphIdent := localAliasPrefix + staticwebsiteparagraph.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.StaticWebSiteParagraph{Name: %s}).Stage(stage)", staticwebsiteparagraphIdent, __gong__toRawStringLiteral(staticwebsiteparagraph.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", staticwebsiteparagraphIdent, __gong__toRawStringLiteral(staticwebsiteparagraph.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.LegendMarkdownContent = %s", staticwebsiteparagraphIdent, __gong__toRawStringLiteral(staticwebsiteparagraph.LegendMarkdownContent)))
			if staticwebsiteparagraph.Image != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := localAliasPrefix + staticwebsiteparagraph.Image.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Image = %s", staticwebsiteparagraphIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, xhtml_content := range __gong__sortStageSetInstances(stage.XHTML_CONTENTs, stage.XHTML_CONTENT_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			xhtml_contentIdent := localAliasPrefix + xhtml_content.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&models.XHTML_CONTENT{Name: %s}).Stage(stage)", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.EnclosedText = %s", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.EnclosedText)))
			values.WriteString(fmt.Sprintf("\n\t%s.PureText = %s", xhtml_contentIdent, __gong__toRawStringLiteral(xhtml_content.PureText)))
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
go run github.com/fullstack-lang/gong/app/reqif/go/models/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/reqif/go/models"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/reqif/go/models" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/reqif/go/models\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (alternative_id *ALTERNATIVE_ID) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := alternative_id.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", alternative_id.Name)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", alternative_id.IDENTIFIER)

	default:
		log.Panicf("Unknown field %s for Gongstruct ALTERNATIVE_ID", fieldName)
	}
	return
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_boolean.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_boolean.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", attribute_definition_boolean.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", attribute_definition_boolean.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", attribute_definition_boolean.IS_EDITABLE)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", attribute_definition_boolean.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", attribute_definition_boolean.LONG_NAME)

	case "ALTERNATIVE_ID":
		if attribute_definition_boolean.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", attribute_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_boolean.DEFAULT_VALUE != nil {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", attribute_definition_boolean.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", "nil")
		}
	case "TYPE":
		if attribute_definition_boolean.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", attribute_definition_boolean.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_BOOLEAN", fieldName)
	}
	return
}

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_boolean_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_boolean_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_boolean_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_boolean_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_boolean_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_boolean_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_BOOLEAN_Rendering", fieldName)
	}
	return
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_date.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_date.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", attribute_definition_date.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", attribute_definition_date.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", attribute_definition_date.IS_EDITABLE)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", attribute_definition_date.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", attribute_definition_date.LONG_NAME)

	case "ALTERNATIVE_ID":
		if attribute_definition_date.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", attribute_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_date.DEFAULT_VALUE != nil {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", attribute_definition_date.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", "nil")
		}
	case "TYPE":
		if attribute_definition_date.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", attribute_definition_date.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_DATE", fieldName)
	}
	return
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_date_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_date_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_date_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_date_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_date_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_date_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_DATE_Rendering", fieldName)
	}
	return
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_enumeration.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_enumeration.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", attribute_definition_enumeration.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", attribute_definition_enumeration.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", attribute_definition_enumeration.IS_EDITABLE)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", attribute_definition_enumeration.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", attribute_definition_enumeration.LONG_NAME)
	case "MULTI_VALUED":
		res = __gong__marshallBool(ident, "MULTI_VALUED", attribute_definition_enumeration.MULTI_VALUED)

	case "ALTERNATIVE_ID":
		if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", attribute_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_enumeration.DEFAULT_VALUE != nil {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", attribute_definition_enumeration.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", "nil")
		}
	case "TYPE":
		if attribute_definition_enumeration.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", attribute_definition_enumeration.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_ENUMERATION", fieldName)
	}
	return
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_enumeration_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_enumeration_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_enumeration_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_enumeration_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_enumeration_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_enumeration_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_ENUMERATION_Rendering", fieldName)
	}
	return
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_integer.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_integer.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", attribute_definition_integer.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", attribute_definition_integer.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", attribute_definition_integer.IS_EDITABLE)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", attribute_definition_integer.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", attribute_definition_integer.LONG_NAME)

	case "ALTERNATIVE_ID":
		if attribute_definition_integer.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", attribute_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_integer.DEFAULT_VALUE != nil {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", attribute_definition_integer.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", "nil")
		}
	case "TYPE":
		if attribute_definition_integer.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", attribute_definition_integer.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_INTEGER", fieldName)
	}
	return
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_integer_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_integer_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_integer_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_integer_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_integer_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_integer_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_INTEGER_Rendering", fieldName)
	}
	return
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_real.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_real.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", attribute_definition_real.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", attribute_definition_real.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", attribute_definition_real.IS_EDITABLE)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", attribute_definition_real.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", attribute_definition_real.LONG_NAME)

	case "ALTERNATIVE_ID":
		if attribute_definition_real.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", attribute_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_real.DEFAULT_VALUE != nil {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", attribute_definition_real.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", "nil")
		}
	case "TYPE":
		if attribute_definition_real.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", attribute_definition_real.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_REAL", fieldName)
	}
	return
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_real_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_real_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_real_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_real_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_real_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_real_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_REAL_Rendering", fieldName)
	}
	return
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_Rendering", fieldName)
	}
	return
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_string.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_string.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", attribute_definition_string.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", attribute_definition_string.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", attribute_definition_string.IS_EDITABLE)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", attribute_definition_string.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", attribute_definition_string.LONG_NAME)

	case "ALTERNATIVE_ID":
		if attribute_definition_string.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", attribute_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_string.DEFAULT_VALUE != nil {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", attribute_definition_string.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", "nil")
		}
	case "TYPE":
		if attribute_definition_string.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", attribute_definition_string.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_STRING", fieldName)
	}
	return
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_string_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_string_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_string_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_string_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_string_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_string_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_STRING_Rendering", fieldName)
	}
	return
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_xhtml.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_xhtml.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", attribute_definition_xhtml.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", attribute_definition_xhtml.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", attribute_definition_xhtml.IS_EDITABLE)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", attribute_definition_xhtml.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", attribute_definition_xhtml.LONG_NAME)

	case "ALTERNATIVE_ID":
		if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", attribute_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "DEFAULT_VALUE":
		if attribute_definition_xhtml.DEFAULT_VALUE != nil {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", attribute_definition_xhtml.DEFAULT_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFAULT_VALUE", "nil")
		}
	case "TYPE":
		if attribute_definition_xhtml.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", attribute_definition_xhtml.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_XHTML", fieldName)
	}
	return
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_definition_xhtml_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_definition_xhtml_rendering.Name)
	case "ShowInTable":
		res = __gong__marshallBool(ident, "ShowInTable", attribute_definition_xhtml_rendering.ShowInTable)
	case "ShowInTitle":
		res = __gong__marshallBool(ident, "ShowInTitle", attribute_definition_xhtml_rendering.ShowInTitle)
	case "ShowInSubject":
		res = __gong__marshallBool(ident, "ShowInSubject", attribute_definition_xhtml_rendering.ShowInSubject)
	case "Rank":
		res = __gong__marshallInt(ident, "Rank", attribute_definition_xhtml_rendering.Rank)

	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_DEFINITION_XHTML_Rendering", fieldName)
	}
	return
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_value_boolean.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_value_boolean.Name)
	case "THE_VALUE":
		res = __gong__marshallBool(ident, "THE_VALUE", attribute_value_boolean.THE_VALUE)

	case "DEFINITION":
		if attribute_value_boolean.DEFINITION != nil {
			res = __gong__marshallPointer(ident, "DEFINITION", attribute_value_boolean.DEFINITION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFINITION", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_BOOLEAN", fieldName)
	}
	return
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_value_date.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_value_date.Name)
	case "THE_VALUE":
		res = __gong__marshallString(ident, "THE_VALUE", attribute_value_date.THE_VALUE)

	case "DEFINITION":
		if attribute_value_date.DEFINITION != nil {
			res = __gong__marshallPointer(ident, "DEFINITION", attribute_value_date.DEFINITION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFINITION", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_DATE", fieldName)
	}
	return
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_value_enumeration.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_value_enumeration.Name)

	case "DEFINITION":
		if attribute_value_enumeration.DEFINITION != nil {
			res = __gong__marshallPointer(ident, "DEFINITION", attribute_value_enumeration.DEFINITION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFINITION", "nil")
		}
	case "VALUES":
		if attribute_value_enumeration.VALUES != nil {
			res = __gong__marshallPointer(ident, "VALUES", attribute_value_enumeration.VALUES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VALUES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_ENUMERATION", fieldName)
	}
	return
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_value_integer.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_value_integer.Name)
	case "THE_VALUE":
		res = __gong__marshallInt(ident, "THE_VALUE", attribute_value_integer.THE_VALUE)

	case "DEFINITION":
		if attribute_value_integer.DEFINITION != nil {
			res = __gong__marshallPointer(ident, "DEFINITION", attribute_value_integer.DEFINITION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFINITION", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_INTEGER", fieldName)
	}
	return
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_value_real.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_value_real.Name)
	case "THE_VALUE":
		res = __gong__marshallFloat(ident, "THE_VALUE", attribute_value_real.THE_VALUE)

	case "DEFINITION":
		if attribute_value_real.DEFINITION != nil {
			res = __gong__marshallPointer(ident, "DEFINITION", attribute_value_real.DEFINITION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFINITION", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_REAL", fieldName)
	}
	return
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_value_string.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_value_string.Name)
	case "THE_VALUE":
		res = __gong__marshallString(ident, "THE_VALUE", attribute_value_string.THE_VALUE)

	case "DEFINITION":
		if attribute_value_string.DEFINITION != nil {
			res = __gong__marshallPointer(ident, "DEFINITION", attribute_value_string.DEFINITION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFINITION", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_STRING", fieldName)
	}
	return
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := attribute_value_xhtml.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", attribute_value_xhtml.Name)
	case "IS_SIMPLIFIED":
		res = __gong__marshallBool(ident, "IS_SIMPLIFIED", attribute_value_xhtml.IS_SIMPLIFIED)

	case "DEFINITION":
		if attribute_value_xhtml.DEFINITION != nil {
			res = __gong__marshallPointer(ident, "DEFINITION", attribute_value_xhtml.DEFINITION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DEFINITION", "nil")
		}
	case "THE_VALUE":
		if attribute_value_xhtml.THE_VALUE != nil {
			res = __gong__marshallPointer(ident, "THE_VALUE", attribute_value_xhtml.THE_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "THE_VALUE", "nil")
		}
	case "THE_ORIGINAL_VALUE":
		if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
			res = __gong__marshallPointer(ident, "THE_ORIGINAL_VALUE", attribute_value_xhtml.THE_ORIGINAL_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "THE_ORIGINAL_VALUE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ATTRIBUTE_VALUE_XHTML", fieldName)
	}
	return
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_alternative_id.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_alternative_id.Name)

	case "ALTERNATIVE_ID":
		if a_alternative_id.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", a_alternative_id.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ALTERNATIVE_ID", fieldName)
	}
	return
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_definition_boolean_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_definition_boolean_ref.Name)
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_BOOLEAN_REF", a_attribute_definition_boolean_ref.ATTRIBUTE_DEFINITION_BOOLEAN_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", fieldName)
	}
	return
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_definition_date_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_definition_date_ref.Name)
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_DATE_REF", a_attribute_definition_date_ref.ATTRIBUTE_DEFINITION_DATE_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_DATE_REF", fieldName)
	}
	return
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_definition_enumeration_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_definition_enumeration_ref.Name)
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_ENUMERATION_REF", a_attribute_definition_enumeration_ref.ATTRIBUTE_DEFINITION_ENUMERATION_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", fieldName)
	}
	return
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_definition_integer_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_definition_integer_ref.Name)
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_INTEGER_REF", a_attribute_definition_integer_ref.ATTRIBUTE_DEFINITION_INTEGER_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_INTEGER_REF", fieldName)
	}
	return
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_definition_real_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_definition_real_ref.Name)
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_REAL_REF", a_attribute_definition_real_ref.ATTRIBUTE_DEFINITION_REAL_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_REAL_REF", fieldName)
	}
	return
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_definition_string_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_definition_string_ref.Name)
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_STRING_REF", a_attribute_definition_string_ref.ATTRIBUTE_DEFINITION_STRING_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_STRING_REF", fieldName)
	}
	return
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_definition_xhtml_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_definition_xhtml_ref.Name)
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_XHTML_REF", a_attribute_definition_xhtml_ref.ATTRIBUTE_DEFINITION_XHTML_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_DEFINITION_XHTML_REF", fieldName)
	}
	return
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_boolean.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_boolean.Name)

	case "ATTRIBUTE_VALUE_BOOLEAN":
		var sb strings.Builder
		for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_BOOLEAN", _attribute_value_boolean.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_BOOLEAN", fieldName)
	}
	return
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_date.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_date.Name)

	case "ATTRIBUTE_VALUE_DATE":
		var sb strings.Builder
		for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_DATE", _attribute_value_date.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_DATE", fieldName)
	}
	return
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_enumeration.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_enumeration.Name)

	case "ATTRIBUTE_VALUE_ENUMERATION":
		var sb strings.Builder
		for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_ENUMERATION", _attribute_value_enumeration.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_ENUMERATION", fieldName)
	}
	return
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_integer.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_integer.Name)

	case "ATTRIBUTE_VALUE_INTEGER":
		var sb strings.Builder
		for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_INTEGER", _attribute_value_integer.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_INTEGER", fieldName)
	}
	return
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_real.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_real.Name)

	case "ATTRIBUTE_VALUE_REAL":
		var sb strings.Builder
		for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_REAL", _attribute_value_real.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_REAL", fieldName)
	}
	return
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_string.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_string.Name)

	case "ATTRIBUTE_VALUE_STRING":
		var sb strings.Builder
		for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_STRING", _attribute_value_string.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_STRING", fieldName)
	}
	return
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_xhtml.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_xhtml.Name)

	case "ATTRIBUTE_VALUE_XHTML":
		var sb strings.Builder
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_XHTML", _attribute_value_xhtml.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_XHTML", fieldName)
	}
	return
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_attribute_value_xhtml_1.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_attribute_value_xhtml_1.Name)

	case "ATTRIBUTE_VALUE_BOOLEAN":
		var sb strings.Builder
		for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_BOOLEAN", _attribute_value_boolean.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_DATE":
		var sb strings.Builder
		for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_DATE", _attribute_value_date.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_ENUMERATION":
		var sb strings.Builder
		for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_ENUMERATION", _attribute_value_enumeration.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_INTEGER":
		var sb strings.Builder
		for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_INTEGER", _attribute_value_integer.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_REAL":
		var sb strings.Builder
		for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_REAL", _attribute_value_real.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_STRING":
		var sb strings.Builder
		for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_STRING", _attribute_value_string.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_VALUE_XHTML":
		var sb strings.Builder
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_VALUE_XHTML", _attribute_value_xhtml.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_ATTRIBUTE_VALUE_XHTML_1", fieldName)
	}
	return
}

func (a_children *A_CHILDREN) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_children.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_children.Name)

	case "SPEC_HIERARCHY":
		var sb strings.Builder
		for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SPEC_HIERARCHY", _spec_hierarchy.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_CHILDREN", fieldName)
	}
	return
}

func (a_core_content *A_CORE_CONTENT) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_core_content.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_core_content.Name)

	case "REQ_IF_CONTENT":
		if a_core_content.REQ_IF_CONTENT != nil {
			res = __gong__marshallPointer(ident, "REQ_IF_CONTENT", a_core_content.REQ_IF_CONTENT.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "REQ_IF_CONTENT", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_CORE_CONTENT", fieldName)
	}
	return
}

func (a_datatypes *A_DATATYPES) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatypes.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatypes.Name)

	case "DATATYPE_DEFINITION_BOOLEAN":
		var sb strings.Builder
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DATATYPE_DEFINITION_BOOLEAN", _datatype_definition_boolean.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_DATE":
		var sb strings.Builder
		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DATATYPE_DEFINITION_DATE", _datatype_definition_date.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_ENUMERATION":
		var sb strings.Builder
		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DATATYPE_DEFINITION_ENUMERATION", _datatype_definition_enumeration.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_INTEGER":
		var sb strings.Builder
		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DATATYPE_DEFINITION_INTEGER", _datatype_definition_integer.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_REAL":
		var sb strings.Builder
		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DATATYPE_DEFINITION_REAL", _datatype_definition_real.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_STRING":
		var sb strings.Builder
		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DATATYPE_DEFINITION_STRING", _datatype_definition_string.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "DATATYPE_DEFINITION_XHTML":
		var sb strings.Builder
		for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "DATATYPE_DEFINITION_XHTML", _datatype_definition_xhtml.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPES", fieldName)
	}
	return
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatype_definition_boolean_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatype_definition_boolean_ref.Name)
	case "DATATYPE_DEFINITION_BOOLEAN_REF":
		res = __gong__marshallString(ident, "DATATYPE_DEFINITION_BOOLEAN_REF", a_datatype_definition_boolean_ref.DATATYPE_DEFINITION_BOOLEAN_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_BOOLEAN_REF", fieldName)
	}
	return
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatype_definition_date_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatype_definition_date_ref.Name)
	case "DATATYPE_DEFINITION_DATE_REF":
		res = __gong__marshallString(ident, "DATATYPE_DEFINITION_DATE_REF", a_datatype_definition_date_ref.DATATYPE_DEFINITION_DATE_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_DATE_REF", fieldName)
	}
	return
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatype_definition_enumeration_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatype_definition_enumeration_ref.Name)
	case "DATATYPE_DEFINITION_ENUMERATION_REF":
		res = __gong__marshallString(ident, "DATATYPE_DEFINITION_ENUMERATION_REF", a_datatype_definition_enumeration_ref.DATATYPE_DEFINITION_ENUMERATION_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_ENUMERATION_REF", fieldName)
	}
	return
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatype_definition_integer_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatype_definition_integer_ref.Name)
	case "DATATYPE_DEFINITION_INTEGER_REF":
		res = __gong__marshallString(ident, "DATATYPE_DEFINITION_INTEGER_REF", a_datatype_definition_integer_ref.DATATYPE_DEFINITION_INTEGER_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_INTEGER_REF", fieldName)
	}
	return
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatype_definition_real_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatype_definition_real_ref.Name)
	case "DATATYPE_DEFINITION_REAL_REF":
		res = __gong__marshallString(ident, "DATATYPE_DEFINITION_REAL_REF", a_datatype_definition_real_ref.DATATYPE_DEFINITION_REAL_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_REAL_REF", fieldName)
	}
	return
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatype_definition_string_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatype_definition_string_ref.Name)
	case "DATATYPE_DEFINITION_STRING_REF":
		res = __gong__marshallString(ident, "DATATYPE_DEFINITION_STRING_REF", a_datatype_definition_string_ref.DATATYPE_DEFINITION_STRING_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_STRING_REF", fieldName)
	}
	return
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_datatype_definition_xhtml_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_datatype_definition_xhtml_ref.Name)
	case "DATATYPE_DEFINITION_XHTML_REF":
		res = __gong__marshallString(ident, "DATATYPE_DEFINITION_XHTML_REF", a_datatype_definition_xhtml_ref.DATATYPE_DEFINITION_XHTML_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_DATATYPE_DEFINITION_XHTML_REF", fieldName)
	}
	return
}

func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_editable_atts.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_editable_atts.Name)
	case "ATTRIBUTE_DEFINITION_BOOLEAN_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_BOOLEAN_REF", a_editable_atts.ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	case "ATTRIBUTE_DEFINITION_DATE_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_DATE_REF", a_editable_atts.ATTRIBUTE_DEFINITION_DATE_REF)
	case "ATTRIBUTE_DEFINITION_ENUMERATION_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_ENUMERATION_REF", a_editable_atts.ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	case "ATTRIBUTE_DEFINITION_INTEGER_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_INTEGER_REF", a_editable_atts.ATTRIBUTE_DEFINITION_INTEGER_REF)
	case "ATTRIBUTE_DEFINITION_REAL_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_REAL_REF", a_editable_atts.ATTRIBUTE_DEFINITION_REAL_REF)
	case "ATTRIBUTE_DEFINITION_STRING_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_STRING_REF", a_editable_atts.ATTRIBUTE_DEFINITION_STRING_REF)
	case "ATTRIBUTE_DEFINITION_XHTML_REF":
		res = __gong__marshallString(ident, "ATTRIBUTE_DEFINITION_XHTML_REF", a_editable_atts.ATTRIBUTE_DEFINITION_XHTML_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_EDITABLE_ATTS", fieldName)
	}
	return
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_enum_value_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_enum_value_ref.Name)
	case "ENUM_VALUE_REF":
		res = __gong__marshallString(ident, "ENUM_VALUE_REF", a_enum_value_ref.ENUM_VALUE_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_ENUM_VALUE_REF", fieldName)
	}
	return
}

func (a_object *A_OBJECT) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_object.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_object.Name)
	case "SPEC_OBJECT_REF":
		res = __gong__marshallString(ident, "SPEC_OBJECT_REF", a_object.SPEC_OBJECT_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_OBJECT", fieldName)
	}
	return
}

func (a_properties *A_PROPERTIES) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_properties.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_properties.Name)

	case "EMBEDDED_VALUE":
		if a_properties.EMBEDDED_VALUE != nil {
			res = __gong__marshallPointer(ident, "EMBEDDED_VALUE", a_properties.EMBEDDED_VALUE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EMBEDDED_VALUE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_PROPERTIES", fieldName)
	}
	return
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_relation_group_type_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_relation_group_type_ref.Name)
	case "RELATION_GROUP_TYPE_REF":
		res = __gong__marshallString(ident, "RELATION_GROUP_TYPE_REF", a_relation_group_type_ref.RELATION_GROUP_TYPE_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_RELATION_GROUP_TYPE_REF", fieldName)
	}
	return
}

func (a_source_1 *A_SOURCE_1) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_source_1.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_source_1.Name)
	case "SPEC_OBJECT_REF":
		res = __gong__marshallString(ident, "SPEC_OBJECT_REF", a_source_1.SPEC_OBJECT_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SOURCE_1", fieldName)
	}
	return
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_source_specification_1.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_source_specification_1.Name)
	case "SPECIFICATION_REF":
		res = __gong__marshallEnumString(ident, "SPECIFICATION_REF", a_source_specification_1.SPECIFICATION_REF.ToCodeString())

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SOURCE_SPECIFICATION_1", fieldName)
	}
	return
}

func (a_specifications *A_SPECIFICATIONS) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_specifications.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_specifications.Name)

	case "SPECIFICATION":
		var sb strings.Builder
		for _, _specification := range a_specifications.SPECIFICATION {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SPECIFICATION", _specification.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPECIFICATIONS", fieldName)
	}
	return
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_specification_type_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_specification_type_ref.Name)
	case "SPECIFICATION_TYPE_REF":
		res = __gong__marshallString(ident, "SPECIFICATION_TYPE_REF", a_specification_type_ref.SPECIFICATION_TYPE_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPECIFICATION_TYPE_REF", fieldName)
	}
	return
}

func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_specified_values.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_specified_values.Name)

	case "ENUM_VALUE":
		var sb strings.Builder
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ENUM_VALUE", _enum_value.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPECIFIED_VALUES", fieldName)
	}
	return
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_attributes.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_attributes.Name)

	case "ATTRIBUTE_DEFINITION_BOOLEAN":
		var sb strings.Builder
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_DEFINITION_BOOLEAN", _attribute_definition_boolean.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_DATE":
		var sb strings.Builder
		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_DEFINITION_DATE", _attribute_definition_date.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_ENUMERATION":
		var sb strings.Builder
		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_DEFINITION_ENUMERATION", _attribute_definition_enumeration.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_INTEGER":
		var sb strings.Builder
		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_DEFINITION_INTEGER", _attribute_definition_integer.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_REAL":
		var sb strings.Builder
		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_DEFINITION_REAL", _attribute_definition_real.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_STRING":
		var sb strings.Builder
		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_DEFINITION_STRING", _attribute_definition_string.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ATTRIBUTE_DEFINITION_XHTML":
		var sb strings.Builder
		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "ATTRIBUTE_DEFINITION_XHTML", _attribute_definition_xhtml.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_ATTRIBUTES", fieldName)
	}
	return
}

func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_objects.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_objects.Name)

	case "SPEC_OBJECT":
		var sb strings.Builder
		for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SPEC_OBJECT", _spec_object.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_OBJECTS", fieldName)
	}
	return
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_object_type_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_object_type_ref.Name)
	case "SPEC_OBJECT_TYPE_REF":
		res = __gong__marshallString(ident, "SPEC_OBJECT_TYPE_REF", a_spec_object_type_ref.SPEC_OBJECT_TYPE_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_OBJECT_TYPE_REF", fieldName)
	}
	return
}

func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_relations.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_relations.Name)

	case "SPEC_RELATION":
		var sb strings.Builder
		for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SPEC_RELATION", _spec_relation.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATIONS", fieldName)
	}
	return
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_relation_groups.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_relation_groups.Name)

	case "RELATION_GROUP":
		var sb strings.Builder
		for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RELATION_GROUP", _relation_group.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATION_GROUPS", fieldName)
	}
	return
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_relation_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_relation_ref.Name)
	case "SPEC_RELATION_REF":
		res = __gong__marshallString(ident, "SPEC_RELATION_REF", a_spec_relation_ref.SPEC_RELATION_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATION_REF", fieldName)
	}
	return
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_relation_type_ref.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_relation_type_ref.Name)
	case "SPEC_RELATION_TYPE_REF":
		res = __gong__marshallString(ident, "SPEC_RELATION_TYPE_REF", a_spec_relation_type_ref.SPEC_RELATION_TYPE_REF)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_RELATION_TYPE_REF", fieldName)
	}
	return
}

func (a_spec_types *A_SPEC_TYPES) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_spec_types.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_spec_types.Name)

	case "RELATION_GROUP_TYPE":
		var sb strings.Builder
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "RELATION_GROUP_TYPE", _relation_group_type.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SPEC_OBJECT_TYPE":
		var sb strings.Builder
		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SPEC_OBJECT_TYPE", _spec_object_type.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SPEC_RELATION_TYPE":
		var sb strings.Builder
		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SPEC_RELATION_TYPE", _spec_relation_type.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SPECIFICATION_TYPE":
		var sb strings.Builder
		for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SPECIFICATION_TYPE", _specification_type.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_SPEC_TYPES", fieldName)
	}
	return
}

func (a_the_header *A_THE_HEADER) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_the_header.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_the_header.Name)

	case "REQ_IF_HEADER":
		if a_the_header.REQ_IF_HEADER != nil {
			res = __gong__marshallPointer(ident, "REQ_IF_HEADER", a_the_header.REQ_IF_HEADER.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "REQ_IF_HEADER", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct A_THE_HEADER", fieldName)
	}
	return
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_tool_extensions.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_tool_extensions.Name)

	case "REQ_IF_TOOL_EXTENSION":
		var sb strings.Builder
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "REQ_IF_TOOL_EXTENSION", _req_if_tool_extension.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct A_TOOL_EXTENSIONS", fieldName)
	}
	return
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datatype_definition_boolean.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datatype_definition_boolean.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", datatype_definition_boolean.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", datatype_definition_boolean.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", datatype_definition_boolean.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", datatype_definition_boolean.LONG_NAME)

	case "ALTERNATIVE_ID":
		if datatype_definition_boolean.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", datatype_definition_boolean.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_BOOLEAN", fieldName)
	}
	return
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datatype_definition_date.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datatype_definition_date.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", datatype_definition_date.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", datatype_definition_date.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", datatype_definition_date.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", datatype_definition_date.LONG_NAME)

	case "ALTERNATIVE_ID":
		if datatype_definition_date.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", datatype_definition_date.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_DATE", fieldName)
	}
	return
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datatype_definition_enumeration.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datatype_definition_enumeration.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", datatype_definition_enumeration.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", datatype_definition_enumeration.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", datatype_definition_enumeration.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", datatype_definition_enumeration.LONG_NAME)

	case "ALTERNATIVE_ID":
		if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", datatype_definition_enumeration.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "SPECIFIED_VALUES":
		if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
			res = __gong__marshallPointer(ident, "SPECIFIED_VALUES", datatype_definition_enumeration.SPECIFIED_VALUES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPECIFIED_VALUES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_ENUMERATION", fieldName)
	}
	return
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datatype_definition_integer.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datatype_definition_integer.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", datatype_definition_integer.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", datatype_definition_integer.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", datatype_definition_integer.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", datatype_definition_integer.LONG_NAME)
	case "MAX":
		res = __gong__marshallInt(ident, "MAX", datatype_definition_integer.MAX)
	case "MIN":
		res = __gong__marshallInt(ident, "MIN", datatype_definition_integer.MIN)

	case "ALTERNATIVE_ID":
		if datatype_definition_integer.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", datatype_definition_integer.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_INTEGER", fieldName)
	}
	return
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datatype_definition_real.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datatype_definition_real.Name)
	case "ACCURACY":
		res = __gong__marshallInt(ident, "ACCURACY", datatype_definition_real.ACCURACY)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", datatype_definition_real.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", datatype_definition_real.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", datatype_definition_real.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", datatype_definition_real.LONG_NAME)
	case "MAX":
		res = __gong__marshallFloat(ident, "MAX", datatype_definition_real.MAX)
	case "MIN":
		res = __gong__marshallFloat(ident, "MIN", datatype_definition_real.MIN)

	case "ALTERNATIVE_ID":
		if datatype_definition_real.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", datatype_definition_real.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_REAL", fieldName)
	}
	return
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datatype_definition_string.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datatype_definition_string.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", datatype_definition_string.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", datatype_definition_string.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", datatype_definition_string.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", datatype_definition_string.LONG_NAME)
	case "MAX_LENGTH":
		res = __gong__marshallInt(ident, "MAX_LENGTH", datatype_definition_string.MAX_LENGTH)

	case "ALTERNATIVE_ID":
		if datatype_definition_string.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", datatype_definition_string.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_STRING", fieldName)
	}
	return
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := datatype_definition_xhtml.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", datatype_definition_xhtml.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", datatype_definition_xhtml.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", datatype_definition_xhtml.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", datatype_definition_xhtml.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", datatype_definition_xhtml.LONG_NAME)

	case "ALTERNATIVE_ID":
		if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", datatype_definition_xhtml.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct DATATYPE_DEFINITION_XHTML", fieldName)
	}
	return
}

func (embedded_value *EMBEDDED_VALUE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := embedded_value.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", embedded_value.Name)
	case "KEY":
		res = __gong__marshallInt(ident, "KEY", embedded_value.KEY)
	case "OTHER_CONTENT":
		res = __gong__marshallString(ident, "OTHER_CONTENT", embedded_value.OTHER_CONTENT)

	default:
		log.Panicf("Unknown field %s for Gongstruct EMBEDDED_VALUE", fieldName)
	}
	return
}

func (enum_value *ENUM_VALUE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := enum_value.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", enum_value.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", enum_value.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", enum_value.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", enum_value.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", enum_value.LONG_NAME)

	case "ALTERNATIVE_ID":
		if enum_value.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", enum_value.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "PROPERTIES":
		if enum_value.PROPERTIES != nil {
			res = __gong__marshallPointer(ident, "PROPERTIES", enum_value.PROPERTIES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "PROPERTIES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ENUM_VALUE", fieldName)
	}
	return
}

func (embeddedjpgimage *EmbeddedJpgImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := embeddedjpgimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", embeddedjpgimage.Name)
	case "Base64Content":
		res = __gong__marshallString(ident, "Base64Content", embeddedjpgimage.Base64Content)

	default:
		log.Panicf("Unknown field %s for Gongstruct EmbeddedJpgImage", fieldName)
	}
	return
}

func (embeddedpngimage *EmbeddedPngImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := embeddedpngimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", embeddedpngimage.Name)
	case "Base64Content":
		res = __gong__marshallString(ident, "Base64Content", embeddedpngimage.Base64Content)

	default:
		log.Panicf("Unknown field %s for Gongstruct EmbeddedPngImage", fieldName)
	}
	return
}

func (embeddedsvgimage *EmbeddedSvgImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := embeddedsvgimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", embeddedsvgimage.Name)
	case "Content":
		res = __gong__marshallString(ident, "Content", embeddedsvgimage.Content)

	default:
		log.Panicf("Unknown field %s for Gongstruct EmbeddedSvgImage", fieldName)
	}
	return
}

func (kill *Kill) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := kill.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", kill.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct Kill", fieldName)
	}
	return
}

func (map_identifier_bool *Map_identifier_bool) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := map_identifier_bool.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", map_identifier_bool.Name)
	case "Value":
		res = __gong__marshallBool(ident, "Value", map_identifier_bool.Value)

	default:
		log.Panicf("Unknown field %s for Gongstruct Map_identifier_bool", fieldName)
	}
	return
}

func (relation_group *RELATION_GROUP) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := relation_group.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", relation_group.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", relation_group.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", relation_group.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", relation_group.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", relation_group.LONG_NAME)

	case "ALTERNATIVE_ID":
		if relation_group.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", relation_group.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "SOURCE_SPECIFICATION":
		if relation_group.SOURCE_SPECIFICATION != nil {
			res = __gong__marshallPointer(ident, "SOURCE_SPECIFICATION", relation_group.SOURCE_SPECIFICATION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SOURCE_SPECIFICATION", "nil")
		}
	case "SPEC_RELATIONS":
		if relation_group.SPEC_RELATIONS != nil {
			res = __gong__marshallPointer(ident, "SPEC_RELATIONS", relation_group.SPEC_RELATIONS.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_RELATIONS", "nil")
		}
	case "TARGET_SPECIFICATION":
		if relation_group.TARGET_SPECIFICATION != nil {
			res = __gong__marshallPointer(ident, "TARGET_SPECIFICATION", relation_group.TARGET_SPECIFICATION.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TARGET_SPECIFICATION", "nil")
		}
	case "TYPE":
		if relation_group.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", relation_group.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RELATION_GROUP", fieldName)
	}
	return
}

func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := relation_group_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", relation_group_type.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", relation_group_type.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", relation_group_type.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", relation_group_type.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", relation_group_type.LONG_NAME)

	case "ALTERNATIVE_ID":
		if relation_group_type.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", relation_group_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if relation_group_type.SPEC_ATTRIBUTES != nil {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", relation_group_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct RELATION_GROUP_TYPE", fieldName)
	}
	return
}

func (req_if *REQ_IF) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := req_if.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", req_if.Name)
	case "Lang":
		res = __gong__marshallString(ident, "Lang", req_if.Lang)

	case "THE_HEADER":
		if req_if.THE_HEADER != nil {
			res = __gong__marshallPointer(ident, "THE_HEADER", req_if.THE_HEADER.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "THE_HEADER", "nil")
		}
	case "CORE_CONTENT":
		if req_if.CORE_CONTENT != nil {
			res = __gong__marshallPointer(ident, "CORE_CONTENT", req_if.CORE_CONTENT.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CORE_CONTENT", "nil")
		}
	case "TOOL_EXTENSIONS":
		if req_if.TOOL_EXTENSIONS != nil {
			res = __gong__marshallPointer(ident, "TOOL_EXTENSIONS", req_if.TOOL_EXTENSIONS.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TOOL_EXTENSIONS", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF", fieldName)
	}
	return
}

func (req_if_content *REQ_IF_CONTENT) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := req_if_content.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", req_if_content.Name)

	case "DATATYPES":
		if req_if_content.DATATYPES != nil {
			res = __gong__marshallPointer(ident, "DATATYPES", req_if_content.DATATYPES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DATATYPES", "nil")
		}
	case "SPEC_TYPES":
		if req_if_content.SPEC_TYPES != nil {
			res = __gong__marshallPointer(ident, "SPEC_TYPES", req_if_content.SPEC_TYPES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_TYPES", "nil")
		}
	case "SPEC_OBJECTS":
		if req_if_content.SPEC_OBJECTS != nil {
			res = __gong__marshallPointer(ident, "SPEC_OBJECTS", req_if_content.SPEC_OBJECTS.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_OBJECTS", "nil")
		}
	case "SPEC_RELATIONS":
		if req_if_content.SPEC_RELATIONS != nil {
			res = __gong__marshallPointer(ident, "SPEC_RELATIONS", req_if_content.SPEC_RELATIONS.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_RELATIONS", "nil")
		}
	case "SPECIFICATIONS":
		if req_if_content.SPECIFICATIONS != nil {
			res = __gong__marshallPointer(ident, "SPECIFICATIONS", req_if_content.SPECIFICATIONS.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPECIFICATIONS", "nil")
		}
	case "SPEC_RELATION_GROUPS":
		if req_if_content.SPEC_RELATION_GROUPS != nil {
			res = __gong__marshallPointer(ident, "SPEC_RELATION_GROUPS", req_if_content.SPEC_RELATION_GROUPS.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_RELATION_GROUPS", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF_CONTENT", fieldName)
	}
	return
}

func (req_if_header *REQ_IF_HEADER) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := req_if_header.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", req_if_header.Name)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", req_if_header.IDENTIFIER)
	case "COMMENT":
		res = __gong__marshallString(ident, "COMMENT", req_if_header.COMMENT)
	case "CREATION_TIME":
		res = __gong__marshallString(ident, "CREATION_TIME", req_if_header.CREATION_TIME)
	case "REPOSITORY_ID":
		res = __gong__marshallString(ident, "REPOSITORY_ID", req_if_header.REPOSITORY_ID)
	case "REQ_IF_TOOL_ID":
		res = __gong__marshallString(ident, "REQ_IF_TOOL_ID", req_if_header.REQ_IF_TOOL_ID)
	case "REQ_IF_VERSION":
		res = __gong__marshallString(ident, "REQ_IF_VERSION", req_if_header.REQ_IF_VERSION)
	case "SOURCE_TOOL_ID":
		res = __gong__marshallString(ident, "SOURCE_TOOL_ID", req_if_header.SOURCE_TOOL_ID)
	case "TITLE":
		res = __gong__marshallString(ident, "TITLE", req_if_header.TITLE)

	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF_HEADER", fieldName)
	}
	return
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := req_if_tool_extension.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", req_if_tool_extension.Name)

	default:
		log.Panicf("Unknown field %s for Gongstruct REQ_IF_TOOL_EXTENSION", fieldName)
	}
	return
}

func (specification *SPECIFICATION) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specification.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specification.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", specification.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", specification.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", specification.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", specification.LONG_NAME)

	case "ALTERNATIVE_ID":
		if specification.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", specification.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "TYPE":
		if specification.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", specification.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	case "CHILDREN":
		if specification.CHILDREN != nil {
			res = __gong__marshallPointer(ident, "CHILDREN", specification.CHILDREN.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CHILDREN", "nil")
		}
	case "VALUES":
		if specification.VALUES != nil {
			res = __gong__marshallPointer(ident, "VALUES", specification.VALUES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VALUES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPECIFICATION", fieldName)
	}
	return
}

func (specification_rendering *SPECIFICATION_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specification_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specification_rendering.Name)
	case "IsNodeExpanded":
		res = __gong__marshallBool(ident, "IsNodeExpanded", specification_rendering.IsNodeExpanded)
	case "IsSelected":
		res = __gong__marshallBool(ident, "IsSelected", specification_rendering.IsSelected)
	case "IsWithHeadingNumbering":
		res = __gong__marshallBool(ident, "IsWithHeadingNumbering", specification_rendering.IsWithHeadingNumbering)

	default:
		log.Panicf("Unknown field %s for Gongstruct SPECIFICATION_Rendering", fieldName)
	}
	return
}

func (specification_type *SPECIFICATION_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specification_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specification_type.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", specification_type.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", specification_type.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", specification_type.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", specification_type.LONG_NAME)

	case "ALTERNATIVE_ID":
		if specification_type.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", specification_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if specification_type.SPEC_ATTRIBUTES != nil {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", specification_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPECIFICATION_TYPE", fieldName)
	}
	return
}

func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := spec_hierarchy.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", spec_hierarchy.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", spec_hierarchy.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", spec_hierarchy.IDENTIFIER)
	case "IS_EDITABLE":
		res = __gong__marshallBool(ident, "IS_EDITABLE", spec_hierarchy.IS_EDITABLE)
	case "IS_TABLE_INTERNAL":
		res = __gong__marshallBool(ident, "IS_TABLE_INTERNAL", spec_hierarchy.IS_TABLE_INTERNAL)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", spec_hierarchy.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", spec_hierarchy.LONG_NAME)

	case "ALTERNATIVE_ID":
		if spec_hierarchy.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", spec_hierarchy.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "OBJECT":
		if spec_hierarchy.OBJECT != nil {
			res = __gong__marshallPointer(ident, "OBJECT", spec_hierarchy.OBJECT.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "OBJECT", "nil")
		}
	case "CHILDREN":
		if spec_hierarchy.CHILDREN != nil {
			res = __gong__marshallPointer(ident, "CHILDREN", spec_hierarchy.CHILDREN.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "CHILDREN", "nil")
		}
	case "EDITABLE_ATTS":
		if spec_hierarchy.EDITABLE_ATTS != nil {
			res = __gong__marshallPointer(ident, "EDITABLE_ATTS", spec_hierarchy.EDITABLE_ATTS.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "EDITABLE_ATTS", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_HIERARCHY", fieldName)
	}
	return
}

func (spec_object *SPEC_OBJECT) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := spec_object.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", spec_object.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", spec_object.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", spec_object.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", spec_object.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", spec_object.LONG_NAME)

	case "ALTERNATIVE_ID":
		if spec_object.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", spec_object.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "VALUES":
		if spec_object.VALUES != nil {
			res = __gong__marshallPointer(ident, "VALUES", spec_object.VALUES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VALUES", "nil")
		}
	case "TYPE":
		if spec_object.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", spec_object.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_OBJECT", fieldName)
	}
	return
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := spec_object_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", spec_object_type.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", spec_object_type.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", spec_object_type.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", spec_object_type.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", spec_object_type.LONG_NAME)

	case "ALTERNATIVE_ID":
		if spec_object_type.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", spec_object_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if spec_object_type.SPEC_ATTRIBUTES != nil {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", spec_object_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_OBJECT_TYPE", fieldName)
	}
	return
}

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := spec_object_type_rendering.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", spec_object_type_rendering.Name)
	case "IsNodeExpanded":
		res = __gong__marshallBool(ident, "IsNodeExpanded", spec_object_type_rendering.IsNodeExpanded)
	case "ShowIdentifier":
		res = __gong__marshallBool(ident, "ShowIdentifier", spec_object_type_rendering.ShowIdentifier)
	case "ShowName":
		res = __gong__marshallBool(ident, "ShowName", spec_object_type_rendering.ShowName)
	case "ShowRelations":
		res = __gong__marshallBool(ident, "ShowRelations", spec_object_type_rendering.ShowRelations)
	case "IsHeading":
		res = __gong__marshallBool(ident, "IsHeading", spec_object_type_rendering.IsHeading)

	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_OBJECT_TYPE_Rendering", fieldName)
	}
	return
}

func (spec_relation *SPEC_RELATION) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := spec_relation.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", spec_relation.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", spec_relation.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", spec_relation.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", spec_relation.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", spec_relation.LONG_NAME)

	case "ALTERNATIVE_ID":
		if spec_relation.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", spec_relation.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "VALUES":
		if spec_relation.VALUES != nil {
			res = __gong__marshallPointer(ident, "VALUES", spec_relation.VALUES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VALUES", "nil")
		}
	case "SOURCE":
		if spec_relation.SOURCE != nil {
			res = __gong__marshallPointer(ident, "SOURCE", spec_relation.SOURCE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SOURCE", "nil")
		}
	case "TARGET":
		if spec_relation.TARGET != nil {
			res = __gong__marshallPointer(ident, "TARGET", spec_relation.TARGET.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TARGET", "nil")
		}
	case "TYPE":
		if spec_relation.TYPE != nil {
			res = __gong__marshallPointer(ident, "TYPE", spec_relation.TYPE.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TYPE", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_RELATION", fieldName)
	}
	return
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := spec_relation_type.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", spec_relation_type.Name)
	case "DESC":
		res = __gong__marshallString(ident, "DESC", spec_relation_type.DESC)
	case "IDENTIFIER":
		res = __gong__marshallString(ident, "IDENTIFIER", spec_relation_type.IDENTIFIER)
	case "LAST_CHANGE":
		res = __gong__marshallString(ident, "LAST_CHANGE", spec_relation_type.LAST_CHANGE)
	case "LONG_NAME":
		res = __gong__marshallString(ident, "LONG_NAME", spec_relation_type.LONG_NAME)

	case "ALTERNATIVE_ID":
		if spec_relation_type.ALTERNATIVE_ID != nil {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", spec_relation_type.ALTERNATIVE_ID.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ALTERNATIVE_ID", "nil")
		}
	case "SPEC_ATTRIBUTES":
		if spec_relation_type.SPEC_ATTRIBUTES != nil {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", spec_relation_type.SPEC_ATTRIBUTES.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SPEC_ATTRIBUTES", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SPEC_RELATION_TYPE", fieldName)
	}
	return
}

func (staticwebsite *StaticWebSite) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := staticwebsite.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", staticwebsite.Name)
	case "MarkdownContent":
		res = __gong__marshallString(ident, "MarkdownContent", staticwebsite.MarkdownContent)
	case "InputImagesDir":
		res = __gong__marshallString(ident, "InputImagesDir", staticwebsite.InputImagesDir)
	case "OutputStaticWebDir":
		res = __gong__marshallString(ident, "OutputStaticWebDir", staticwebsite.OutputStaticWebDir)
	case "VersionInfo":
		res = __gong__marshallString(ident, "VersionInfo", staticwebsite.VersionInfo)

	case "Chapters":
		var sb strings.Builder
		for _, _staticwebsitechapter := range staticwebsite.Chapters {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Chapters", _staticwebsitechapter.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StaticWebSite", fieldName)
	}
	return
}

func (staticwebsitechapter *StaticWebSiteChapter) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := staticwebsitechapter.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", staticwebsitechapter.Name)
	case "MarkdownContent":
		res = __gong__marshallString(ident, "MarkdownContent", staticwebsitechapter.MarkdownContent)

	case "Paragraphs":
		var sb strings.Builder
		for _, _staticwebsiteparagraph := range staticwebsitechapter.Paragraphs {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Paragraphs", _staticwebsiteparagraph.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct StaticWebSiteChapter", fieldName)
	}
	return
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := staticwebsitegeneratedimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", staticwebsitegeneratedimage.Name)
	case "SourceDirectoryPath":
		res = __gong__marshallString(ident, "SourceDirectoryPath", staticwebsitegeneratedimage.SourceDirectoryPath)
	case "Width":
		res = __gong__marshallInt(ident, "Width", staticwebsitegeneratedimage.Width)
	case "Height":
		res = __gong__marshallInt(ident, "Height", staticwebsitegeneratedimage.Height)

	default:
		log.Panicf("Unknown field %s for Gongstruct StaticWebSiteGeneratedImage", fieldName)
	}
	return
}

func (staticwebsiteimage *StaticWebSiteImage) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := staticwebsiteimage.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", staticwebsiteimage.Name)
	case "SourceDirectoryPath":
		res = __gong__marshallString(ident, "SourceDirectoryPath", staticwebsiteimage.SourceDirectoryPath)
	case "Width":
		res = __gong__marshallInt(ident, "Width", staticwebsiteimage.Width)
	case "Height":
		res = __gong__marshallInt(ident, "Height", staticwebsiteimage.Height)

	default:
		log.Panicf("Unknown field %s for Gongstruct StaticWebSiteImage", fieldName)
	}
	return
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := staticwebsiteparagraph.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", staticwebsiteparagraph.Name)
	case "LegendMarkdownContent":
		res = __gong__marshallString(ident, "LegendMarkdownContent", staticwebsiteparagraph.LegendMarkdownContent)

	case "Image":
		if staticwebsiteparagraph.Image != nil {
			res = __gong__marshallPointer(ident, "Image", staticwebsiteparagraph.Image.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Image", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct StaticWebSiteParagraph", fieldName)
	}
	return
}

func (xhtml_content *XHTML_CONTENT) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := xhtml_content.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", xhtml_content.Name)
	case "EnclosedText":
		res = __gong__marshallString(ident, "EnclosedText", xhtml_content.EnclosedText)
	case "PureText":
		res = __gong__marshallString(ident, "PureText", xhtml_content.PureText)

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
func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_boolean_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_boolean_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_boolean_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_boolean_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_boolean_rendering.GongMarshallField(stage, "Rank"))
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
func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_date_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_date_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_date_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_date_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_date_rendering.GongMarshallField(stage, "Rank"))
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
func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_enumeration_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_enumeration_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_enumeration_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_enumeration_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_enumeration_rendering.GongMarshallField(stage, "Rank"))
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
func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_integer_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_integer_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_integer_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_integer_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_integer_rendering.GongMarshallField(stage, "Rank"))
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
func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_real_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_real_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_real_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_real_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_real_rendering.GongMarshallField(stage, "Rank"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_rendering.GongMarshallField(stage, "Rank"))
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
func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_string_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_string_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_string_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_string_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_string_rendering.GongMarshallField(stage, "Rank"))
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
func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(attribute_definition_xhtml_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(attribute_definition_xhtml_rendering.GongMarshallField(stage, "ShowInTable"))
		initializerStatements.WriteString(attribute_definition_xhtml_rendering.GongMarshallField(stage, "ShowInTitle"))
		initializerStatements.WriteString(attribute_definition_xhtml_rendering.GongMarshallField(stage, "ShowInSubject"))
		initializerStatements.WriteString(attribute_definition_xhtml_rendering.GongMarshallField(stage, "Rank"))
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
		pointersInitializesStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "DEFINITION"))
		initializerStatements.WriteString(attribute_value_boolean.GongMarshallField(stage, "THE_VALUE"))
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
		pointersInitializesStatements.WriteString(attribute_value_date.GongMarshallField(stage, "DEFINITION"))
		initializerStatements.WriteString(attribute_value_date.GongMarshallField(stage, "THE_VALUE"))
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
		pointersInitializesStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "DEFINITION"))
		initializerStatements.WriteString(attribute_value_integer.GongMarshallField(stage, "THE_VALUE"))
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
		pointersInitializesStatements.WriteString(attribute_value_real.GongMarshallField(stage, "DEFINITION"))
		initializerStatements.WriteString(attribute_value_real.GongMarshallField(stage, "THE_VALUE"))
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
		pointersInitializesStatements.WriteString(attribute_value_string.GongMarshallField(stage, "DEFINITION"))
		initializerStatements.WriteString(attribute_value_string.GongMarshallField(stage, "THE_VALUE"))
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
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "DEFINITION"))
		initializerStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "IS_SIMPLIFIED"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "THE_VALUE"))
		pointersInitializesStatements.WriteString(attribute_value_xhtml.GongMarshallField(stage, "THE_ORIGINAL_VALUE"))
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
func (embeddedjpgimage *EmbeddedJpgImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(embeddedjpgimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(embeddedjpgimage.GongMarshallField(stage, "Base64Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (embeddedpngimage *EmbeddedPngImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(embeddedpngimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(embeddedpngimage.GongMarshallField(stage, "Base64Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (embeddedsvgimage *EmbeddedSvgImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(embeddedsvgimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(embeddedsvgimage.GongMarshallField(stage, "Content"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (kill *Kill) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(kill.GongMarshallField(stage, "Name"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (map_identifier_bool *Map_identifier_bool) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(map_identifier_bool.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(map_identifier_bool.GongMarshallField(stage, "Value"))
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
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "TYPE"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "CHILDREN"))
		pointersInitializesStatements.WriteString(specification.GongMarshallField(stage, "VALUES"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specification_rendering *SPECIFICATION_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specification_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specification_rendering.GongMarshallField(stage, "IsNodeExpanded"))
		initializerStatements.WriteString(specification_rendering.GongMarshallField(stage, "IsSelected"))
		initializerStatements.WriteString(specification_rendering.GongMarshallField(stage, "IsWithHeadingNumbering"))
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
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "OBJECT"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "CHILDREN"))
		pointersInitializesStatements.WriteString(spec_hierarchy.GongMarshallField(stage, "EDITABLE_ATTS"))
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
func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(spec_object_type_rendering.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(spec_object_type_rendering.GongMarshallField(stage, "IsNodeExpanded"))
		initializerStatements.WriteString(spec_object_type_rendering.GongMarshallField(stage, "ShowIdentifier"))
		initializerStatements.WriteString(spec_object_type_rendering.GongMarshallField(stage, "ShowName"))
		initializerStatements.WriteString(spec_object_type_rendering.GongMarshallField(stage, "ShowRelations"))
		initializerStatements.WriteString(spec_object_type_rendering.GongMarshallField(stage, "IsHeading"))
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
func (staticwebsite *StaticWebSite) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(staticwebsite.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(staticwebsite.GongMarshallField(stage, "MarkdownContent"))
		pointersInitializesStatements.WriteString(staticwebsite.GongMarshallField(stage, "Chapters"))
		initializerStatements.WriteString(staticwebsite.GongMarshallField(stage, "InputImagesDir"))
		initializerStatements.WriteString(staticwebsite.GongMarshallField(stage, "OutputStaticWebDir"))
		initializerStatements.WriteString(staticwebsite.GongMarshallField(stage, "VersionInfo"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (staticwebsitechapter *StaticWebSiteChapter) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(staticwebsitechapter.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(staticwebsitechapter.GongMarshallField(stage, "MarkdownContent"))
		pointersInitializesStatements.WriteString(staticwebsitechapter.GongMarshallField(stage, "Paragraphs"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(staticwebsitegeneratedimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(staticwebsitegeneratedimage.GongMarshallField(stage, "SourceDirectoryPath"))
		initializerStatements.WriteString(staticwebsitegeneratedimage.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(staticwebsitegeneratedimage.GongMarshallField(stage, "Height"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (staticwebsiteimage *StaticWebSiteImage) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(staticwebsiteimage.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(staticwebsiteimage.GongMarshallField(stage, "SourceDirectoryPath"))
		initializerStatements.WriteString(staticwebsiteimage.GongMarshallField(stage, "Width"))
		initializerStatements.WriteString(staticwebsiteimage.GongMarshallField(stage, "Height"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (staticwebsiteparagraph *StaticWebSiteParagraph) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(staticwebsiteparagraph.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(staticwebsiteparagraph.GongMarshallField(stage, "LegendMarkdownContent"))
		pointersInitializesStatements.WriteString(staticwebsiteparagraph.GongMarshallField(stage, "Image"))
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
