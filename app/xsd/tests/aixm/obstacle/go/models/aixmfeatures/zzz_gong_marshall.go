// generated code - do not edit
package aixmfeatures

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
	{{Identifier}} := (&aixmfeatures.{{GeneratedStructName}}{Name: {{GeneratedFieldNameValue}}}).Stage(stage)`

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
			const fnSig = "func _(stage *aixmfeatures.Stage) {"
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
			strings.HasSuffix(trimmedContentBeforeBrace, "func _(stage *aixmfeatures.Stage) {") ||
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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"
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
		for _, a_extension := range __gong__sortStageSetInstances(stage.A_extensions, stage.A_extension_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extensionIdent := a_extension.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension{Name: %s}).Stage(stage)", a_extensionIdent, __gong__toRawStringLiteral(a_extension.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extensionIdent, __gong__toRawStringLiteral(a_extension.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extensionIdent, a_extension.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_1 := range __gong__sortStageSetInstances(stage.A_extension_1s, stage.A_extension_1_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_1Ident := a_extension_1.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_1{Name: %s}).Stage(stage)", a_extension_1Ident, __gong__toRawStringLiteral(a_extension_1.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_1Ident, __gong__toRawStringLiteral(a_extension_1.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_1Ident, a_extension_1.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_10 := range __gong__sortStageSetInstances(stage.A_extension_10s, stage.A_extension_10_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_10Ident := a_extension_10.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_10{Name: %s}).Stage(stage)", a_extension_10Ident, __gong__toRawStringLiteral(a_extension_10.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_10Ident, __gong__toRawStringLiteral(a_extension_10.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_10Ident, a_extension_10.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_11 := range __gong__sortStageSetInstances(stage.A_extension_11s, stage.A_extension_11_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_11Ident := a_extension_11.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_11{Name: %s}).Stage(stage)", a_extension_11Ident, __gong__toRawStringLiteral(a_extension_11.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_11Ident, __gong__toRawStringLiteral(a_extension_11.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_11Ident, a_extension_11.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_12 := range __gong__sortStageSetInstances(stage.A_extension_12s, stage.A_extension_12_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_12Ident := a_extension_12.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_12{Name: %s}).Stage(stage)", a_extension_12Ident, __gong__toRawStringLiteral(a_extension_12.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_12Ident, __gong__toRawStringLiteral(a_extension_12.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_12Ident, a_extension_12.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_2 := range __gong__sortStageSetInstances(stage.A_extension_2s, stage.A_extension_2_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_2Ident := a_extension_2.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_2{Name: %s}).Stage(stage)", a_extension_2Ident, __gong__toRawStringLiteral(a_extension_2.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_2Ident, __gong__toRawStringLiteral(a_extension_2.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_2Ident, a_extension_2.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_3 := range __gong__sortStageSetInstances(stage.A_extension_3s, stage.A_extension_3_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_3Ident := a_extension_3.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_3{Name: %s}).Stage(stage)", a_extension_3Ident, __gong__toRawStringLiteral(a_extension_3.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_3Ident, __gong__toRawStringLiteral(a_extension_3.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_3Ident, a_extension_3.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_4 := range __gong__sortStageSetInstances(stage.A_extension_4s, stage.A_extension_4_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_4Ident := a_extension_4.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_4{Name: %s}).Stage(stage)", a_extension_4Ident, __gong__toRawStringLiteral(a_extension_4.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_4Ident, __gong__toRawStringLiteral(a_extension_4.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_4Ident, a_extension_4.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_5 := range __gong__sortStageSetInstances(stage.A_extension_5s, stage.A_extension_5_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_5Ident := a_extension_5.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_5{Name: %s}).Stage(stage)", a_extension_5Ident, __gong__toRawStringLiteral(a_extension_5.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_5Ident, __gong__toRawStringLiteral(a_extension_5.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_5Ident, a_extension_5.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_6 := range __gong__sortStageSetInstances(stage.A_extension_6s, stage.A_extension_6_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_6Ident := a_extension_6.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_6{Name: %s}).Stage(stage)", a_extension_6Ident, __gong__toRawStringLiteral(a_extension_6.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_6Ident, __gong__toRawStringLiteral(a_extension_6.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_6Ident, a_extension_6.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_7 := range __gong__sortStageSetInstances(stage.A_extension_7s, stage.A_extension_7_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_7Ident := a_extension_7.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_7{Name: %s}).Stage(stage)", a_extension_7Ident, __gong__toRawStringLiteral(a_extension_7.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_7Ident, __gong__toRawStringLiteral(a_extension_7.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_7Ident, a_extension_7.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_8 := range __gong__sortStageSetInstances(stage.A_extension_8s, stage.A_extension_8_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_8Ident := a_extension_8.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_8{Name: %s}).Stage(stage)", a_extension_8Ident, __gong__toRawStringLiteral(a_extension_8.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_8Ident, __gong__toRawStringLiteral(a_extension_8.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_8Ident, a_extension_8.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_9 := range __gong__sortStageSetInstances(stage.A_extension_9s, stage.A_extension_9_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_9Ident := a_extension_9.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_9{Name: %s}).Stage(stage)", a_extension_9Ident, __gong__toRawStringLiteral(a_extension_9.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_9Ident, __gong__toRawStringLiteral(a_extension_9.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_9Ident, a_extension_9.Owns))
		}
	}
	if stage != nil {
		for _, abstractpropertieswithscheduletype := range __gong__sortStageSetInstances(stage.AbstractPropertiesWithScheduleTypes, stage.AbstractPropertiesWithScheduleType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			abstractpropertieswithscheduletypeIdent := abstractpropertieswithscheduletype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AbstractPropertiesWithScheduleType{Name: %s}).Stage(stage)", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", abstractpropertieswithscheduletypeIdent, __gong__toRawStringLiteral(abstractpropertieswithscheduletype.Description)))
		}
	}
	if stage != nil {
		for _, aeronauticalgroundlighttimeslicepropertytype := range __gong__sortStageSetInstances(stage.AeronauticalGroundLightTimeSlicePropertyTypes, stage.AeronauticalGroundLightTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aeronauticalgroundlighttimeslicepropertytypeIdent := aeronauticalgroundlighttimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AeronauticalGroundLightTimeSlicePropertyType{Name: %s}).Stage(stage)", aeronauticalgroundlighttimeslicepropertytypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aeronauticalgroundlighttimeslicepropertytypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", aeronauticalgroundlighttimeslicepropertytypeIdent, aeronauticalgroundlighttimeslicepropertytype.Owns))
			if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AeronauticalGroundLightTimeSlice = %s", aeronauticalgroundlighttimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aeronauticalgroundlighttimeslicetype := range __gong__sortStageSetInstances(stage.AeronauticalGroundLightTimeSliceTypes, stage.AeronauticalGroundLightTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aeronauticalgroundlighttimeslicetypeIdent := aeronauticalgroundlighttimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AeronauticalGroundLightTimeSliceType{Name: %s}).Stage(stage)", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", aeronauticalgroundlighttimeslicetypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetype.CorrectionNumber))
			if aeronauticalgroundlighttimeslicetype.StructureBeacon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.StructureBeacon.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.StructureBeacon = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.AerodromeBeacon.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AerodromeBeacon = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			if aeronauticalgroundlighttimeslicetype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aeronauticalgroundlighttimeslicetype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aeronauticalgroundlighttimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
			for _, elem := range aeronauticalgroundlighttimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aeronauticalgroundlighttimeslicetypeIdent, aeronauticalgroundlighttimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aeronauticalgroundlighttype := range __gong__sortStageSetInstances(stage.AeronauticalGroundLightTypes, stage.AeronauticalGroundLightType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aeronauticalgroundlighttypeIdent := aeronauticalgroundlighttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AeronauticalGroundLightType{Name: %s}).Stage(stage)", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aeronauticalgroundlighttypeIdent, __gong__toRawStringLiteral(aeronauticalgroundlighttype.Description)))
			for _, elem := range aeronauticalgroundlighttype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", aeronauticalgroundlighttypeIdent, aeronauticalgroundlighttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airportheliportpropertytype := range __gong__sortStageSetInstances(stage.AirportHeliportPropertyTypes, stage.AirportHeliportPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airportheliportpropertytypeIdent := airportheliportpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirportHeliportPropertyType{Name: %s}).Stage(stage)", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airportheliportpropertytypeIdent, airportheliportpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", airportheliportpropertytypeIdent, __gong__toRawStringLiteral(airportheliportpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, curvetype := range __gong__sortStageSetInstances(stage.CurveTypes, stage.CurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curvetypeIdent := curvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CurveType{Name: %s}).Stage(stage)", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", curvetypeIdent, __gong__toRawStringLiteral(curvetype.SrsName)))
			for _, elem := range curvetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", curvetypeIdent, curvetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedcurvepropertytype := range __gong__sortStageSetInstances(stage.ElevatedCurvePropertyTypes, stage.ElevatedCurvePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedcurvepropertytypeIdent := elevatedcurvepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedCurvePropertyType{Name: %s}).Stage(stage)", elevatedcurvepropertytypeIdent, __gong__toRawStringLiteral(elevatedcurvepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedcurvepropertytypeIdent, __gong__toRawStringLiteral(elevatedcurvepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", elevatedcurvepropertytypeIdent, __gong__toRawStringLiteral(elevatedcurvepropertytype.NilReason)))
			if elevatedcurvepropertytype.ElevatedCurve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elevatedcurvepropertytype.ElevatedCurve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedCurve = %s", elevatedcurvepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedcurvetype := range __gong__sortStageSetInstances(stage.ElevatedCurveTypes, stage.ElevatedCurveType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedcurvetypeIdent := elevatedcurvetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedCurveType{Name: %s}).Stage(stage)", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedcurvetypeIdent, __gong__toRawStringLiteral(elevatedcurvetype.SrsName)))
			for _, elem := range elevatedcurvetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedcurvetypeIdent, elevatedcurvetypeIdent, targetIdent))
			}
			for _, elem := range elevatedcurvetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedcurvetypeIdent, elevatedcurvetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedpointpropertytype := range __gong__sortStageSetInstances(stage.ElevatedPointPropertyTypes, stage.ElevatedPointPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedpointpropertytypeIdent := elevatedpointpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedPointPropertyType{Name: %s}).Stage(stage)", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", elevatedpointpropertytypeIdent, __gong__toRawStringLiteral(elevatedpointpropertytype.NilReason)))
			if elevatedpointpropertytype.ElevatedPoint != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elevatedpointpropertytype.ElevatedPoint.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedPoint = %s", elevatedpointpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedpointtype := range __gong__sortStageSetInstances(stage.ElevatedPointTypes, stage.ElevatedPointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedpointtypeIdent := elevatedpointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedPointType{Name: %s}).Stage(stage)", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedpointtypeIdent, __gong__toRawStringLiteral(elevatedpointtype.SrsName)))
			for _, elem := range elevatedpointtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedpointtypeIdent, elevatedpointtypeIdent, targetIdent))
			}
			for _, elem := range elevatedpointtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedpointtypeIdent, elevatedpointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedsurfacepropertytype := range __gong__sortStageSetInstances(stage.ElevatedSurfacePropertyTypes, stage.ElevatedSurfacePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedsurfacepropertytypeIdent := elevatedsurfacepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedSurfacePropertyType{Name: %s}).Stage(stage)", elevatedsurfacepropertytypeIdent, __gong__toRawStringLiteral(elevatedsurfacepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedsurfacepropertytypeIdent, __gong__toRawStringLiteral(elevatedsurfacepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", elevatedsurfacepropertytypeIdent, __gong__toRawStringLiteral(elevatedsurfacepropertytype.NilReason)))
			if elevatedsurfacepropertytype.ElevatedSurface != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elevatedsurfacepropertytype.ElevatedSurface.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedSurface = %s", elevatedsurfacepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, elevatedsurfacetype := range __gong__sortStageSetInstances(stage.ElevatedSurfaceTypes, stage.ElevatedSurfaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			elevatedsurfacetypeIdent := elevatedsurfacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ElevatedSurfaceType{Name: %s}).Stage(stage)", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", elevatedsurfacetypeIdent, __gong__toRawStringLiteral(elevatedsurfacetype.SrsName)))
			for _, elem := range elevatedsurfacetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", elevatedsurfacetypeIdent, elevatedsurfacetypeIdent, targetIdent))
			}
			for _, elem := range elevatedsurfacetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", elevatedsurfacetypeIdent, elevatedsurfacetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, groundlightsystempropertytype := range __gong__sortStageSetInstances(stage.GroundLightSystemPropertyTypes, stage.GroundLightSystemPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			groundlightsystempropertytypeIdent := groundlightsystempropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GroundLightSystemPropertyType{Name: %s}).Stage(stage)", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", groundlightsystempropertytypeIdent, groundlightsystempropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", groundlightsystempropertytypeIdent, __gong__toRawStringLiteral(groundlightsystempropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, lightelementpropertytype := range __gong__sortStageSetInstances(stage.LightElementPropertyTypes, stage.LightElementPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lightelementpropertytypeIdent := lightelementpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementPropertyType{Name: %s}).Stage(stage)", lightelementpropertytypeIdent, __gong__toRawStringLiteral(lightelementpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementpropertytypeIdent, __gong__toRawStringLiteral(lightelementpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", lightelementpropertytypeIdent, __gong__toRawStringLiteral(lightelementpropertytype.NilReason)))
			if lightelementpropertytype.LightElement != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := lightelementpropertytype.LightElement.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightElement = %s", lightelementpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, lightelementstatuspropertytype := range __gong__sortStageSetInstances(stage.LightElementStatusPropertyTypes, stage.LightElementStatusPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lightelementstatuspropertytypeIdent := lightelementstatuspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementStatusPropertyType{Name: %s}).Stage(stage)", lightelementstatuspropertytypeIdent, __gong__toRawStringLiteral(lightelementstatuspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementstatuspropertytypeIdent, __gong__toRawStringLiteral(lightelementstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", lightelementstatuspropertytypeIdent, __gong__toRawStringLiteral(lightelementstatuspropertytype.NilReason)))
			if lightelementstatuspropertytype.LightElementStatus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := lightelementstatuspropertytype.LightElementStatus.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightElementStatus = %s", lightelementstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, lightelementstatustype := range __gong__sortStageSetInstances(stage.LightElementStatusTypes, stage.LightElementStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lightelementstatustypeIdent := lightelementstatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementStatusType{Name: %s}).Stage(stage)", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", lightelementstatustypeIdent, __gong__toRawStringLiteral(lightelementstatustype.Description)))
			for _, elem := range lightelementstatustype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
			for _, elem := range lightelementstatustype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", lightelementstatustypeIdent, lightelementstatustypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, lightelementtype := range __gong__sortStageSetInstances(stage.LightElementTypes, stage.LightElementType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			lightelementtypeIdent := lightelementtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LightElementType{Name: %s}).Stage(stage)", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", lightelementtypeIdent, __gong__toRawStringLiteral(lightelementtype.Description)))
			if lightelementtype.Location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := lightelementtype.Location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Location = %s", lightelementtypeIdent, targetIdent))
			}
			for _, elem := range lightelementtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", lightelementtypeIdent, lightelementtypeIdent, targetIdent))
			}
			for _, elem := range lightelementtype.Availability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Availability = append(%s.Availability, %s)", lightelementtypeIdent, lightelementtypeIdent, targetIdent))
			}
			for _, elem := range lightelementtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", lightelementtypeIdent, lightelementtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, linguisticnotepropertytype := range __gong__sortStageSetInstances(stage.LinguisticNotePropertyTypes, stage.LinguisticNotePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linguisticnotepropertytypeIdent := linguisticnotepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LinguisticNotePropertyType{Name: %s}).Stage(stage)", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", linguisticnotepropertytypeIdent, __gong__toRawStringLiteral(linguisticnotepropertytype.NilReason)))
			if linguisticnotepropertytype.LinguisticNote != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := linguisticnotepropertytype.LinguisticNote.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LinguisticNote = %s", linguisticnotepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, linguisticnotetype := range __gong__sortStageSetInstances(stage.LinguisticNoteTypes, stage.LinguisticNoteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			linguisticnotetypeIdent := linguisticnotetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.LinguisticNoteType{Name: %s}).Stage(stage)", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", linguisticnotetypeIdent, __gong__toRawStringLiteral(linguisticnotetype.Description)))
			for _, elem := range linguisticnotetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", linguisticnotetypeIdent, linguisticnotetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, markerbeaconpropertytype := range __gong__sortStageSetInstances(stage.MarkerBeaconPropertyTypes, stage.MarkerBeaconPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			markerbeaconpropertytypeIdent := markerbeaconpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.MarkerBeaconPropertyType{Name: %s}).Stage(stage)", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", markerbeaconpropertytypeIdent, markerbeaconpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", markerbeaconpropertytypeIdent, __gong__toRawStringLiteral(markerbeaconpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, navaidequipmentpropertytype := range __gong__sortStageSetInstances(stage.NavaidEquipmentPropertyTypes, stage.NavaidEquipmentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			navaidequipmentpropertytypeIdent := navaidequipmentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NavaidEquipmentPropertyType{Name: %s}).Stage(stage)", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", navaidequipmentpropertytypeIdent, navaidequipmentpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", navaidequipmentpropertytypeIdent, __gong__toRawStringLiteral(navaidequipmentpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, notepropertytype := range __gong__sortStageSetInstances(stage.NotePropertyTypes, stage.NotePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notepropertytypeIdent := notepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NotePropertyType{Name: %s}).Stage(stage)", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", notepropertytypeIdent, __gong__toRawStringLiteral(notepropertytype.NilReason)))
			if notepropertytype.Note != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := notepropertytype.Note.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Note = %s", notepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, notetype := range __gong__sortStageSetInstances(stage.NoteTypes, stage.NoteType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			notetypeIdent := notetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.NoteType{Name: %s}).Stage(stage)", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", notetypeIdent, __gong__toRawStringLiteral(notetype.Description)))
			for _, elem := range notetype.TranslatedNote {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TranslatedNote = append(%s.TranslatedNote, %s)", notetypeIdent, notetypeIdent, targetIdent))
			}
			for _, elem := range notetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", notetypeIdent, notetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstacleareatimeslicepropertytype := range __gong__sortStageSetInstances(stage.ObstacleAreaTimeSlicePropertyTypes, stage.ObstacleAreaTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstacleareatimeslicepropertytypeIdent := obstacleareatimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAreaTimeSlicePropertyType{Name: %s}).Stage(stage)", obstacleareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleareatimeslicepropertytypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", obstacleareatimeslicepropertytypeIdent, obstacleareatimeslicepropertytype.Owns))
			if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ObstacleAreaTimeSlice = %s", obstacleareatimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstacleareatimeslicetype := range __gong__sortStageSetInstances(stage.ObstacleAreaTimeSliceTypes, stage.ObstacleAreaTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstacleareatimeslicetypeIdent := obstacleareatimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAreaTimeSliceType{Name: %s}).Stage(stage)", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", obstacleareatimeslicetypeIdent, __gong__toRawStringLiteral(obstacleareatimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", obstacleareatimeslicetypeIdent, obstacleareatimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", obstacleareatimeslicetypeIdent, obstacleareatimeslicetype.CorrectionNumber))
			if obstacleareatimeslicetype.Reference_ownerOrganisation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleareatimeslicetype.Reference_ownerOrganisation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reference_ownerOrganisation = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.Reference_ownerRunway != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleareatimeslicetype.Reference_ownerRunway.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reference_ownerRunway = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.Reference_ownerAirport != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleareatimeslicetype.Reference_ownerAirport.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Reference_ownerAirport = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			if obstacleareatimeslicetype.SurfaceExtent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := obstacleareatimeslicetype.SurfaceExtent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SurfaceExtent = %s", obstacleareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatimeslicetype.Obstacle {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Obstacle = append(%s.Obstacle, %s)", obstacleareatimeslicetypeIdent, obstacleareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", obstacleareatimeslicetypeIdent, obstacleareatimeslicetypeIdent, targetIdent))
			}
			for _, elem := range obstacleareatimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", obstacleareatimeslicetypeIdent, obstacleareatimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, obstacleareatype := range __gong__sortStageSetInstances(stage.ObstacleAreaTypes, stage.ObstacleAreaType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			obstacleareatypeIdent := obstacleareatype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ObstacleAreaType{Name: %s}).Stage(stage)", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", obstacleareatypeIdent, __gong__toRawStringLiteral(obstacleareatype.Description)))
			for _, elem := range obstacleareatype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", obstacleareatypeIdent, obstacleareatypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, organisationauthoritypropertytype := range __gong__sortStageSetInstances(stage.OrganisationAuthorityPropertyTypes, stage.OrganisationAuthorityPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			organisationauthoritypropertytypeIdent := organisationauthoritypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.OrganisationAuthorityPropertyType{Name: %s}).Stage(stage)", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", organisationauthoritypropertytypeIdent, organisationauthoritypropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", organisationauthoritypropertytypeIdent, __gong__toRawStringLiteral(organisationauthoritypropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, passengerservicepropertytype := range __gong__sortStageSetInstances(stage.PassengerServicePropertyTypes, stage.PassengerServicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			passengerservicepropertytypeIdent := passengerservicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PassengerServicePropertyType{Name: %s}).Stage(stage)", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", passengerservicepropertytypeIdent, passengerservicepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", passengerservicepropertytypeIdent, __gong__toRawStringLiteral(passengerservicepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, pointtype := range __gong__sortStageSetInstances(stage.PointTypes, stage.PointType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			pointtypeIdent := pointtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.PointType{Name: %s}).Stage(stage)", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", pointtypeIdent, __gong__toRawStringLiteral(pointtype.SrsName)))
			for _, elem := range pointtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", pointtypeIdent, pointtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, runwaydirectionpropertytype := range __gong__sortStageSetInstances(stage.RunwayDirectionPropertyTypes, stage.RunwayDirectionPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			runwaydirectionpropertytypeIdent := runwaydirectionpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RunwayDirectionPropertyType{Name: %s}).Stage(stage)", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", runwaydirectionpropertytypeIdent, runwaydirectionpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", runwaydirectionpropertytypeIdent, __gong__toRawStringLiteral(runwaydirectionpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, servicepropertytype := range __gong__sortStageSetInstances(stage.ServicePropertyTypes, stage.ServicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			servicepropertytypeIdent := servicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.ServicePropertyType{Name: %s}).Stage(stage)", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", servicepropertytypeIdent, servicepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", servicepropertytypeIdent, __gong__toRawStringLiteral(servicepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, specialnavigationstationpropertytype := range __gong__sortStageSetInstances(stage.SpecialNavigationStationPropertyTypes, stage.SpecialNavigationStationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			specialnavigationstationpropertytypeIdent := specialnavigationstationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SpecialNavigationStationPropertyType{Name: %s}).Stage(stage)", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", specialnavigationstationpropertytypeIdent, specialnavigationstationpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", specialnavigationstationpropertytypeIdent, __gong__toRawStringLiteral(specialnavigationstationpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, surfacepropertytype := range __gong__sortStageSetInstances(stage.SurfacePropertyTypes, stage.SurfacePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surfacepropertytypeIdent := surfacepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurfacePropertyType{Name: %s}).Stage(stage)", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", surfacepropertytypeIdent, __gong__toRawStringLiteral(surfacepropertytype.NilReason)))
			if surfacepropertytype.Surface != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surfacepropertytype.Surface.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Surface = %s", surfacepropertytypeIdent, targetIdent))
			}
			if surfacepropertytype.ElevatedSurface != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := surfacepropertytype.ElevatedSurface.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedSurface = %s", surfacepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, surfacetype := range __gong__sortStageSetInstances(stage.SurfaceTypes, stage.SurfaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			surfacetypeIdent := surfacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.SurfaceType{Name: %s}).Stage(stage)", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.SrsName = %s", surfacetypeIdent, __gong__toRawStringLiteral(surfacetype.SrsName)))
			for _, elem := range surfacetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", surfacetypeIdent, surfacetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, timesheetpropertytype := range __gong__sortStageSetInstances(stage.TimesheetPropertyTypes, stage.TimesheetPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timesheetpropertytypeIdent := timesheetpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TimesheetPropertyType{Name: %s}).Stage(stage)", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", timesheetpropertytypeIdent, __gong__toRawStringLiteral(timesheetpropertytype.NilReason)))
			if timesheetpropertytype.Timesheet != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := timesheetpropertytype.Timesheet.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Timesheet = %s", timesheetpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, timesheettype := range __gong__sortStageSetInstances(stage.TimesheetTypes, stage.TimesheetType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			timesheettypeIdent := timesheettype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.TimesheetType{Name: %s}).Stage(stage)", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", timesheettypeIdent, __gong__toRawStringLiteral(timesheettype.Description)))
			for _, elem := range timesheettype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", timesheettypeIdent, timesheettypeIdent, targetIdent))
			}
			for _, elem := range timesheettype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", timesheettypeIdent, timesheettypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, unitpropertytype := range __gong__sortStageSetInstances(stage.UnitPropertyTypes, stage.UnitPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			unitpropertytypeIdent := unitpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.UnitPropertyType{Name: %s}).Stage(stage)", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", unitpropertytypeIdent, unitpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", unitpropertytypeIdent, __gong__toRawStringLiteral(unitpropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, verticalstructurelightingstatuspropertytype := range __gong__sortStageSetInstances(stage.VerticalStructureLightingStatusPropertyTypes, stage.VerticalStructureLightingStatusPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructurelightingstatuspropertytypeIdent := verticalstructurelightingstatuspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureLightingStatusPropertyType{Name: %s}).Stage(stage)", verticalstructurelightingstatuspropertytypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatuspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurelightingstatuspropertytypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatuspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", verticalstructurelightingstatuspropertytypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatuspropertytype.NilReason)))
			if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalStructureLightingStatus = %s", verticalstructurelightingstatuspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, verticalstructurelightingstatustype := range __gong__sortStageSetInstances(stage.VerticalStructureLightingStatusTypes, stage.VerticalStructureLightingStatusType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructurelightingstatustypeIdent := verticalstructurelightingstatustype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureLightingStatusType{Name: %s}).Stage(stage)", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructurelightingstatustypeIdent, __gong__toRawStringLiteral(verticalstructurelightingstatustype.Description)))
			for _, elem := range verticalstructurelightingstatustype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
			for _, elem := range verticalstructurelightingstatustype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", verticalstructurelightingstatustypeIdent, verticalstructurelightingstatustypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, verticalstructurepartpropertytype := range __gong__sortStageSetInstances(stage.VerticalStructurePartPropertyTypes, stage.VerticalStructurePartPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructurepartpropertytypeIdent := verticalstructurepartpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructurePartPropertyType{Name: %s}).Stage(stage)", verticalstructurepartpropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepartpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurepartpropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepartpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", verticalstructurepartpropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepartpropertytype.NilReason)))
			if verticalstructurepartpropertytype.VerticalStructurePart != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := verticalstructurepartpropertytype.VerticalStructurePart.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalStructurePart = %s", verticalstructurepartpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, verticalstructureparttype := range __gong__sortStageSetInstances(stage.VerticalStructurePartTypes, stage.VerticalStructurePartType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructureparttypeIdent := verticalstructureparttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructurePartType{Name: %s}).Stage(stage)", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructureparttypeIdent, __gong__toRawStringLiteral(verticalstructureparttype.Description)))
			for _, elem := range verticalstructureparttype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.HorizontalProjection_surfaceExtent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := verticalstructureparttype.HorizontalProjection_surfaceExtent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection_surfaceExtent = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.HorizontalProjection_linearExtent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := verticalstructureparttype.HorizontalProjection_linearExtent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection_linearExtent = %s", verticalstructureparttypeIdent, targetIdent))
			}
			if verticalstructureparttype.HorizontalProjection_location != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := verticalstructureparttype.HorizontalProjection_location.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection_location = %s", verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.Lighting {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Lighting = append(%s.Lighting, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
			for _, elem := range verticalstructureparttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", verticalstructureparttypeIdent, verticalstructureparttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, verticalstructurepropertytype := range __gong__sortStageSetInstances(stage.VerticalStructurePropertyTypes, stage.VerticalStructurePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructurepropertytypeIdent := verticalstructurepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructurePropertyType{Name: %s}).Stage(stage)", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", verticalstructurepropertytypeIdent, verticalstructurepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", verticalstructurepropertytypeIdent, __gong__toRawStringLiteral(verticalstructurepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, verticalstructuretimeslicepropertytype := range __gong__sortStageSetInstances(stage.VerticalStructureTimeSlicePropertyTypes, stage.VerticalStructureTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructuretimeslicepropertytypeIdent := verticalstructuretimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureTimeSlicePropertyType{Name: %s}).Stage(stage)", verticalstructuretimeslicepropertytypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructuretimeslicepropertytypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", verticalstructuretimeslicepropertytypeIdent, verticalstructuretimeslicepropertytype.Owns))
			if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.VerticalStructureTimeSlice = %s", verticalstructuretimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, verticalstructuretimeslicetype := range __gong__sortStageSetInstances(stage.VerticalStructureTimeSliceTypes, stage.VerticalStructureTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructuretimeslicetypeIdent := verticalstructuretimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureTimeSliceType{Name: %s}).Stage(stage)", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", verticalstructuretimeslicetypeIdent, __gong__toRawStringLiteral(verticalstructuretimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetype.CorrectionNumber))
			if verticalstructuretimeslicetype.Marker != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := verticalstructuretimeslicetype.Marker.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Marker = %s", verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.Part {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Part = append(%s.Part, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedPassengerService {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedPassengerService = append(%s.HostedPassengerService, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.SupportedGroundLight {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SupportedGroundLight = append(%s.SupportedGroundLight, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedNavaidEquipment {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedNavaidEquipment = append(%s.HostedNavaidEquipment, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedSpecialNavStation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedSpecialNavStation = append(%s.HostedSpecialNavStation, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedUnit {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedUnit = append(%s.HostedUnit, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.HostedOrganisation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HostedOrganisation = append(%s.HostedOrganisation, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.SupportedService {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SupportedService = append(%s.SupportedService, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.LightingAvailability {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.LightingAvailability = append(%s.LightingAvailability, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
			for _, elem := range verticalstructuretimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", verticalstructuretimeslicetypeIdent, verticalstructuretimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, verticalstructuretype := range __gong__sortStageSetInstances(stage.VerticalStructureTypes, stage.VerticalStructureType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			verticalstructuretypeIdent := verticalstructuretype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.VerticalStructureType{Name: %s}).Stage(stage)", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", verticalstructuretypeIdent, __gong__toRawStringLiteral(verticalstructuretype.Description)))
			for _, elem := range verticalstructuretype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", verticalstructuretypeIdent, verticalstructuretypeIdent, targetIdent))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures"%s
)

var (
	_ time.Time
	_ = slices.Index[[]int, int]

	_ *aixmfeatures.Stage
)%s

// function will stage objects across all coordinated stages
func _(stage *aixmfeatures.Stage) {

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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/obstacle/go/models/aixmfeatures\"", fmt.Sprintf("%q", modelsPackageName))
	}

	return res, nil
}

// insertion point for marshall field methods
func (a_extension *A_extension) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension", fieldName)
	}
	return
}

func (a_extension_1 *A_extension_1) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_1.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_1.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_1.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_1", fieldName)
	}
	return
}

func (a_extension_10 *A_extension_10) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_10.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_10.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_10.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_10", fieldName)
	}
	return
}

func (a_extension_11 *A_extension_11) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_11.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_11.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_11.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_11", fieldName)
	}
	return
}

func (a_extension_12 *A_extension_12) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_12.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_12.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_12.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_12", fieldName)
	}
	return
}

func (a_extension_2 *A_extension_2) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_2.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_2.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_2.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_2", fieldName)
	}
	return
}

func (a_extension_3 *A_extension_3) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_3.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_3.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_3.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_3", fieldName)
	}
	return
}

func (a_extension_4 *A_extension_4) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_4.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_4.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_4.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_4", fieldName)
	}
	return
}

func (a_extension_5 *A_extension_5) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_5.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_5.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_5.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_5", fieldName)
	}
	return
}

func (a_extension_6 *A_extension_6) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_6.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_6.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_6.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_6", fieldName)
	}
	return
}

func (a_extension_7 *A_extension_7) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_7.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_7.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_7.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_7", fieldName)
	}
	return
}

func (a_extension_8 *A_extension_8) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_8.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_8.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_8.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_8", fieldName)
	}
	return
}

func (a_extension_9 *A_extension_9) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_9.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_9.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_9.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_9", fieldName)
	}
	return
}

func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := abstractpropertieswithscheduletype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", abstractpropertieswithscheduletype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", abstractpropertieswithscheduletype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", abstractpropertieswithscheduletype.Description)

	default:
		log.Panicf("Unknown field %s for Gongstruct AbstractPropertiesWithScheduleType", fieldName)
	}
	return
}

func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aeronauticalgroundlighttimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aeronauticalgroundlighttimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", aeronauticalgroundlighttimeslicepropertytype.Owns)

	case "AeronauticalGroundLightTimeSlice":
		if aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AeronauticalGroundLightTimeSlice", aeronauticalgroundlighttimeslicepropertytype.AeronauticalGroundLightTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AeronauticalGroundLightTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AeronauticalGroundLightTimeSlicePropertyType", fieldName)
	}
	return
}

func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aeronauticalgroundlighttimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aeronauticalgroundlighttimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aeronauticalgroundlighttimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aeronauticalgroundlighttimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", aeronauticalgroundlighttimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", aeronauticalgroundlighttimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", aeronauticalgroundlighttimeslicetype.CorrectionNumber)

	case "StructureBeacon":
		if aeronauticalgroundlighttimeslicetype.StructureBeacon != nil {
			res = __gong__marshallPointer(ident, "StructureBeacon", aeronauticalgroundlighttimeslicetype.StructureBeacon.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "StructureBeacon", "nil")
		}
	case "AerodromeBeacon":
		if aeronauticalgroundlighttimeslicetype.AerodromeBeacon != nil {
			res = __gong__marshallPointer(ident, "AerodromeBeacon", aeronauticalgroundlighttimeslicetype.AerodromeBeacon.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AerodromeBeacon", "nil")
		}
	case "Location":
		if aeronauticalgroundlighttimeslicetype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", aeronauticalgroundlighttimeslicetype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aeronauticalgroundlighttimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_3 := range aeronauticalgroundlighttimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AeronauticalGroundLightTimeSliceType", fieldName)
	}
	return
}

func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aeronauticalgroundlighttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aeronauticalgroundlighttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aeronauticalgroundlighttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aeronauticalgroundlighttype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _aeronauticalgroundlighttimeslicepropertytype := range aeronauticalgroundlighttype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _aeronauticalgroundlighttimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AeronauticalGroundLightType", fieldName)
	}
	return
}

func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airportheliportpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airportheliportpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airportheliportpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airportheliportpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", airportheliportpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", airportheliportpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", airportheliportpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", airportheliportpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", airportheliportpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", airportheliportpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", airportheliportpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AirportHeliportPropertyType", fieldName)
	}
	return
}

func (curvetype *CurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", curvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", curvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", curvetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range curvetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct CurveType", fieldName)
	}
	return
}

func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedcurvepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedcurvepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", elevatedcurvepropertytype.NilReason)

	case "ElevatedCurve":
		if elevatedcurvepropertytype.ElevatedCurve != nil {
			res = __gong__marshallPointer(ident, "ElevatedCurve", elevatedcurvepropertytype.ElevatedCurve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedCurve", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedCurvePropertyType", fieldName)
	}
	return
}

func (elevatedcurvetype *ElevatedCurveType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedcurvetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedcurvetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevatedcurvetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevatedcurvetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", elevatedcurvetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range elevatedcurvetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension := range elevatedcurvetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedCurveType", fieldName)
	}
	return
}

func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedpointpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedpointpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", elevatedpointpropertytype.NilReason)

	case "ElevatedPoint":
		if elevatedpointpropertytype.ElevatedPoint != nil {
			res = __gong__marshallPointer(ident, "ElevatedPoint", elevatedpointpropertytype.ElevatedPoint.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedPoint", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedPointPropertyType", fieldName)
	}
	return
}

func (elevatedpointtype *ElevatedPointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedpointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedpointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevatedpointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevatedpointtype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", elevatedpointtype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range elevatedpointtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_1 := range elevatedpointtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_1.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedPointType", fieldName)
	}
	return
}

func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedsurfacepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedsurfacepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", elevatedsurfacepropertytype.NilReason)

	case "ElevatedSurface":
		if elevatedsurfacepropertytype.ElevatedSurface != nil {
			res = __gong__marshallPointer(ident, "ElevatedSurface", elevatedsurfacepropertytype.ElevatedSurface.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedSurface", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedSurfacePropertyType", fieldName)
	}
	return
}

func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := elevatedsurfacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", elevatedsurfacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", elevatedsurfacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", elevatedsurfacetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", elevatedsurfacetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range elevatedsurfacetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_2 := range elevatedsurfacetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_2.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedSurfaceType", fieldName)
	}
	return
}

func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := groundlightsystempropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", groundlightsystempropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", groundlightsystempropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", groundlightsystempropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", groundlightsystempropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", groundlightsystempropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", groundlightsystempropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", groundlightsystempropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", groundlightsystempropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", groundlightsystempropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", groundlightsystempropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct GroundLightSystemPropertyType", fieldName)
	}
	return
}

func (lightelementpropertytype *LightElementPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lightelementpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lightelementpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", lightelementpropertytype.NilReason)

	case "LightElement":
		if lightelementpropertytype.LightElement != nil {
			res = __gong__marshallPointer(ident, "LightElement", lightelementpropertytype.LightElement.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LightElement", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LightElementPropertyType", fieldName)
	}
	return
}

func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lightelementstatuspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lightelementstatuspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", lightelementstatuspropertytype.NilReason)

	case "LightElementStatus":
		if lightelementstatuspropertytype.LightElementStatus != nil {
			res = __gong__marshallPointer(ident, "LightElementStatus", lightelementstatuspropertytype.LightElementStatus.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LightElementStatus", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LightElementStatusPropertyType", fieldName)
	}
	return
}

func (lightelementstatustype *LightElementStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lightelementstatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lightelementstatustype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", lightelementstatustype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", lightelementstatustype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range lightelementstatustype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range lightelementstatustype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range lightelementstatustype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_11 := range lightelementstatustype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_11.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LightElementStatusType", fieldName)
	}
	return
}

func (lightelementtype *LightElementType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := lightelementtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", lightelementtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", lightelementtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", lightelementtype.Description)

	case "Location":
		if lightelementtype.Location != nil {
			res = __gong__marshallPointer(ident, "Location", lightelementtype.Location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Location", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range lightelementtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Availability":
		var sb strings.Builder
		for _, _lightelementstatuspropertytype := range lightelementtype.Availability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Availability", _lightelementstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_10 := range lightelementtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_10.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LightElementType", fieldName)
	}
	return
}

func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linguisticnotepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linguisticnotepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", linguisticnotepropertytype.NilReason)

	case "LinguisticNote":
		if linguisticnotepropertytype.LinguisticNote != nil {
			res = __gong__marshallPointer(ident, "LinguisticNote", linguisticnotepropertytype.LinguisticNote.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "LinguisticNote", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNotePropertyType", fieldName)
	}
	return
}

func (linguisticnotetype *LinguisticNoteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := linguisticnotetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", linguisticnotetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", linguisticnotetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", linguisticnotetype.Description)

	case "Extension":
		var sb strings.Builder
		for _, _a_extension_4 := range linguisticnotetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_4.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNoteType", fieldName)
	}
	return
}

func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := markerbeaconpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", markerbeaconpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", markerbeaconpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", markerbeaconpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", markerbeaconpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", markerbeaconpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", markerbeaconpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", markerbeaconpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", markerbeaconpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", markerbeaconpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", markerbeaconpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct MarkerBeaconPropertyType", fieldName)
	}
	return
}

func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := navaidequipmentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", navaidequipmentpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", navaidequipmentpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", navaidequipmentpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", navaidequipmentpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", navaidequipmentpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", navaidequipmentpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", navaidequipmentpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", navaidequipmentpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", navaidequipmentpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", navaidequipmentpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct NavaidEquipmentPropertyType", fieldName)
	}
	return
}

func (notepropertytype *NotePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", notepropertytype.NilReason)

	case "Note":
		if notepropertytype.Note != nil {
			res = __gong__marshallPointer(ident, "Note", notepropertytype.Note.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Note", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct NotePropertyType", fieldName)
	}
	return
}

func (notetype *NoteType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := notetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", notetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", notetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", notetype.Description)

	case "TranslatedNote":
		var sb strings.Builder
		for _, _linguisticnotepropertytype := range notetype.TranslatedNote {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TranslatedNote", _linguisticnotepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_5 := range notetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_5.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteType", fieldName)
	}
	return
}

func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstacleareatimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstacleareatimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", obstacleareatimeslicepropertytype.Owns)

	case "ObstacleAreaTimeSlice":
		if obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice != nil {
			res = __gong__marshallPointer(ident, "ObstacleAreaTimeSlice", obstacleareatimeslicepropertytype.ObstacleAreaTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ObstacleAreaTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstacleAreaTimeSlicePropertyType", fieldName)
	}
	return
}

func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstacleareatimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstacleareatimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", obstacleareatimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", obstacleareatimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", obstacleareatimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", obstacleareatimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", obstacleareatimeslicetype.CorrectionNumber)

	case "Reference_ownerOrganisation":
		if obstacleareatimeslicetype.Reference_ownerOrganisation != nil {
			res = __gong__marshallPointer(ident, "Reference_ownerOrganisation", obstacleareatimeslicetype.Reference_ownerOrganisation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Reference_ownerOrganisation", "nil")
		}
	case "Reference_ownerRunway":
		if obstacleareatimeslicetype.Reference_ownerRunway != nil {
			res = __gong__marshallPointer(ident, "Reference_ownerRunway", obstacleareatimeslicetype.Reference_ownerRunway.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Reference_ownerRunway", "nil")
		}
	case "Reference_ownerAirport":
		if obstacleareatimeslicetype.Reference_ownerAirport != nil {
			res = __gong__marshallPointer(ident, "Reference_ownerAirport", obstacleareatimeslicetype.Reference_ownerAirport.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Reference_ownerAirport", "nil")
		}
	case "SurfaceExtent":
		if obstacleareatimeslicetype.SurfaceExtent != nil {
			res = __gong__marshallPointer(ident, "SurfaceExtent", obstacleareatimeslicetype.SurfaceExtent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "SurfaceExtent", "nil")
		}
	case "Obstacle":
		var sb strings.Builder
		for _, _verticalstructurepropertytype := range obstacleareatimeslicetype.Obstacle {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Obstacle", _verticalstructurepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range obstacleareatimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_6 := range obstacleareatimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_6.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstacleAreaTimeSliceType", fieldName)
	}
	return
}

func (obstacleareatype *ObstacleAreaType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := obstacleareatype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", obstacleareatype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", obstacleareatype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", obstacleareatype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _obstacleareatimeslicepropertytype := range obstacleareatype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _obstacleareatimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ObstacleAreaType", fieldName)
	}
	return
}

func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := organisationauthoritypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", organisationauthoritypropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", organisationauthoritypropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", organisationauthoritypropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", organisationauthoritypropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", organisationauthoritypropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", organisationauthoritypropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", organisationauthoritypropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", organisationauthoritypropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", organisationauthoritypropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", organisationauthoritypropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct OrganisationAuthorityPropertyType", fieldName)
	}
	return
}

func (passengerservicepropertytype *PassengerServicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := passengerservicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", passengerservicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", passengerservicepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", passengerservicepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", passengerservicepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", passengerservicepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", passengerservicepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", passengerservicepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", passengerservicepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", passengerservicepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", passengerservicepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct PassengerServicePropertyType", fieldName)
	}
	return
}

func (pointtype *PointType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := pointtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", pointtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", pointtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", pointtype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", pointtype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range pointtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct PointType", fieldName)
	}
	return
}

func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := runwaydirectionpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", runwaydirectionpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", runwaydirectionpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", runwaydirectionpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", runwaydirectionpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", runwaydirectionpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", runwaydirectionpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", runwaydirectionpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", runwaydirectionpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", runwaydirectionpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", runwaydirectionpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RunwayDirectionPropertyType", fieldName)
	}
	return
}

func (servicepropertytype *ServicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := servicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", servicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", servicepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", servicepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", servicepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", servicepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", servicepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", servicepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", servicepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", servicepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", servicepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct ServicePropertyType", fieldName)
	}
	return
}

func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := specialnavigationstationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", specialnavigationstationpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", specialnavigationstationpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", specialnavigationstationpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", specialnavigationstationpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", specialnavigationstationpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", specialnavigationstationpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", specialnavigationstationpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", specialnavigationstationpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", specialnavigationstationpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", specialnavigationstationpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct SpecialNavigationStationPropertyType", fieldName)
	}
	return
}

func (surfacepropertytype *SurfacePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surfacepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surfacepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", surfacepropertytype.NilReason)

	case "Surface":
		if surfacepropertytype.Surface != nil {
			res = __gong__marshallPointer(ident, "Surface", surfacepropertytype.Surface.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Surface", "nil")
		}
	case "ElevatedSurface":
		if surfacepropertytype.ElevatedSurface != nil {
			res = __gong__marshallPointer(ident, "ElevatedSurface", surfacepropertytype.ElevatedSurface.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedSurface", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct SurfacePropertyType", fieldName)
	}
	return
}

func (surfacetype *SurfaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := surfacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", surfacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", surfacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", surfacetype.Description)
	case "SrsName":
		res = __gong__marshallString(ident, "SrsName", surfacetype.SrsName)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range surfacetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct SurfaceType", fieldName)
	}
	return
}

func (timesheetpropertytype *TimesheetPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timesheetpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timesheetpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", timesheetpropertytype.NilReason)

	case "Timesheet":
		if timesheetpropertytype.Timesheet != nil {
			res = __gong__marshallPointer(ident, "Timesheet", timesheetpropertytype.Timesheet.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Timesheet", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct TimesheetPropertyType", fieldName)
	}
	return
}

func (timesheettype *TimesheetType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := timesheettype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", timesheettype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", timesheettype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", timesheettype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range timesheettype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_12 := range timesheettype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_12.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TimesheetType", fieldName)
	}
	return
}

func (unitpropertytype *UnitPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := unitpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", unitpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", unitpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", unitpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", unitpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", unitpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", unitpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", unitpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", unitpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", unitpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", unitpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct UnitPropertyType", fieldName)
	}
	return
}

func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructurelightingstatuspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructurelightingstatuspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", verticalstructurelightingstatuspropertytype.NilReason)

	case "VerticalStructureLightingStatus":
		if verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus != nil {
			res = __gong__marshallPointer(ident, "VerticalStructureLightingStatus", verticalstructurelightingstatuspropertytype.VerticalStructureLightingStatus.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VerticalStructureLightingStatus", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructureLightingStatusPropertyType", fieldName)
	}
	return
}

func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructurelightingstatustype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructurelightingstatustype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", verticalstructurelightingstatustype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", verticalstructurelightingstatustype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range verticalstructurelightingstatustype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range verticalstructurelightingstatustype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range verticalstructurelightingstatustype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_8 := range verticalstructurelightingstatustype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_8.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructureLightingStatusType", fieldName)
	}
	return
}

func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructurepartpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructurepartpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", verticalstructurepartpropertytype.NilReason)

	case "VerticalStructurePart":
		if verticalstructurepartpropertytype.VerticalStructurePart != nil {
			res = __gong__marshallPointer(ident, "VerticalStructurePart", verticalstructurepartpropertytype.VerticalStructurePart.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VerticalStructurePart", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructurePartPropertyType", fieldName)
	}
	return
}

func (verticalstructureparttype *VerticalStructurePartType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructureparttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructureparttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", verticalstructureparttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", verticalstructureparttype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range verticalstructureparttype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range verticalstructureparttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range verticalstructureparttype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "HorizontalProjection_surfaceExtent":
		if verticalstructureparttype.HorizontalProjection_surfaceExtent != nil {
			res = __gong__marshallPointer(ident, "HorizontalProjection_surfaceExtent", verticalstructureparttype.HorizontalProjection_surfaceExtent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HorizontalProjection_surfaceExtent", "nil")
		}
	case "HorizontalProjection_linearExtent":
		if verticalstructureparttype.HorizontalProjection_linearExtent != nil {
			res = __gong__marshallPointer(ident, "HorizontalProjection_linearExtent", verticalstructureparttype.HorizontalProjection_linearExtent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HorizontalProjection_linearExtent", "nil")
		}
	case "HorizontalProjection_location":
		if verticalstructureparttype.HorizontalProjection_location != nil {
			res = __gong__marshallPointer(ident, "HorizontalProjection_location", verticalstructureparttype.HorizontalProjection_location.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HorizontalProjection_location", "nil")
		}
	case "Lighting":
		var sb strings.Builder
		for _, _lightelementpropertytype := range verticalstructureparttype.Lighting {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Lighting", _lightelementpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_9 := range verticalstructureparttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_9.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructurePartType", fieldName)
	}
	return
}

func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructurepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructurepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", verticalstructurepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", verticalstructurepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", verticalstructurepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", verticalstructurepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", verticalstructurepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", verticalstructurepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", verticalstructurepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", verticalstructurepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", verticalstructurepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructurePropertyType", fieldName)
	}
	return
}

func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructuretimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructuretimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", verticalstructuretimeslicepropertytype.Owns)

	case "VerticalStructureTimeSlice":
		if verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice != nil {
			res = __gong__marshallPointer(ident, "VerticalStructureTimeSlice", verticalstructuretimeslicepropertytype.VerticalStructureTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "VerticalStructureTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructureTimeSlicePropertyType", fieldName)
	}
	return
}

func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructuretimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructuretimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", verticalstructuretimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", verticalstructuretimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", verticalstructuretimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", verticalstructuretimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", verticalstructuretimeslicetype.CorrectionNumber)

	case "Marker":
		if verticalstructuretimeslicetype.Marker != nil {
			res = __gong__marshallPointer(ident, "Marker", verticalstructuretimeslicetype.Marker.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Marker", "nil")
		}
	case "Part":
		var sb strings.Builder
		for _, _verticalstructurepartpropertytype := range verticalstructuretimeslicetype.Part {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Part", _verticalstructurepartpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "HostedPassengerService":
		var sb strings.Builder
		for _, _passengerservicepropertytype := range verticalstructuretimeslicetype.HostedPassengerService {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HostedPassengerService", _passengerservicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SupportedGroundLight":
		var sb strings.Builder
		for _, _groundlightsystempropertytype := range verticalstructuretimeslicetype.SupportedGroundLight {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SupportedGroundLight", _groundlightsystempropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "HostedNavaidEquipment":
		var sb strings.Builder
		for _, _navaidequipmentpropertytype := range verticalstructuretimeslicetype.HostedNavaidEquipment {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HostedNavaidEquipment", _navaidequipmentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "HostedSpecialNavStation":
		var sb strings.Builder
		for _, _specialnavigationstationpropertytype := range verticalstructuretimeslicetype.HostedSpecialNavStation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HostedSpecialNavStation", _specialnavigationstationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "HostedUnit":
		var sb strings.Builder
		for _, _unitpropertytype := range verticalstructuretimeslicetype.HostedUnit {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HostedUnit", _unitpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "HostedOrganisation":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range verticalstructuretimeslicetype.HostedOrganisation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "HostedOrganisation", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SupportedService":
		var sb strings.Builder
		for _, _servicepropertytype := range verticalstructuretimeslicetype.SupportedService {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SupportedService", _servicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range verticalstructuretimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "LightingAvailability":
		var sb strings.Builder
		for _, _verticalstructurelightingstatuspropertytype := range verticalstructuretimeslicetype.LightingAvailability {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "LightingAvailability", _verticalstructurelightingstatuspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_7 := range verticalstructuretimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_7.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructureTimeSliceType", fieldName)
	}
	return
}

func (verticalstructuretype *VerticalStructureType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := verticalstructuretype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", verticalstructuretype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", verticalstructuretype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", verticalstructuretype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _verticalstructuretimeslicepropertytype := range verticalstructuretype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _verticalstructuretimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct VerticalStructureType", fieldName)
	}
	return
}

// insertion point for marshall all fields methods
func (a_extension *A_extension) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_1 *A_extension_1) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_1.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_1.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_10 *A_extension_10) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_10.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_10.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_11 *A_extension_11) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_11.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_11.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_12 *A_extension_12) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_12.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_12.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_2 *A_extension_2) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_2.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_2.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_3 *A_extension_3) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_3.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_3.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_4 *A_extension_4) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_4.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_4.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_5 *A_extension_5) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_5.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_5.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_6 *A_extension_6) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_6.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_6.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_7 *A_extension_7) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_7.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_7.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_8 *A_extension_8) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_8.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_8.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_9 *A_extension_9) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_9.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_9.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (abstractpropertieswithscheduletype *AbstractPropertiesWithScheduleType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(abstractpropertieswithscheduletype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(abstractpropertieswithscheduletype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(abstractpropertieswithscheduletype.GongMarshallField(stage, "Description"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aeronauticalgroundlighttimeslicepropertytype *AeronauticalGroundLightTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(aeronauticalgroundlighttimeslicepropertytype.GongMarshallField(stage, "AeronauticalGroundLightTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aeronauticalgroundlighttimeslicetype *AeronauticalGroundLightTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "StructureBeacon"))
		pointersInitializesStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "AerodromeBeacon"))
		pointersInitializesStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aeronauticalgroundlighttimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aeronauticalgroundlighttype *AeronauticalGroundLightType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aeronauticalgroundlighttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aeronauticalgroundlighttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aeronauticalgroundlighttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aeronauticalgroundlighttype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airportheliportpropertytype *AirportHeliportPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(airportheliportpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curvetype *CurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(curvetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(curvetype.GongMarshallField(stage, "Annotation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedcurvepropertytype *ElevatedCurvePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedcurvepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedcurvepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(elevatedcurvepropertytype.GongMarshallField(stage, "ElevatedCurve"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedcurvetype *ElevatedCurveType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(elevatedcurvetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedpointpropertytype *ElevatedPointPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedpointpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedpointpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(elevatedpointpropertytype.GongMarshallField(stage, "ElevatedPoint"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedpointtype *ElevatedPointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(elevatedpointtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedsurfacepropertytype *ElevatedSurfacePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedsurfacepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedsurfacepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(elevatedsurfacepropertytype.GongMarshallField(stage, "ElevatedSurface"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (elevatedsurfacetype *ElevatedSurfaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(elevatedsurfacetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (groundlightsystempropertytype *GroundLightSystemPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(groundlightsystempropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lightelementpropertytype *LightElementPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lightelementpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lightelementpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(lightelementpropertytype.GongMarshallField(stage, "LightElement"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lightelementstatuspropertytype *LightElementStatusPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lightelementstatuspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lightelementstatuspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(lightelementstatuspropertytype.GongMarshallField(stage, "LightElementStatus"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lightelementstatustype *LightElementStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lightelementstatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lightelementstatustype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(lightelementstatustype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(lightelementstatustype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(lightelementstatustype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(lightelementstatustype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(lightelementstatustype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (lightelementtype *LightElementType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(lightelementtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(lightelementtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(lightelementtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(lightelementtype.GongMarshallField(stage, "Location"))
		pointersInitializesStatements.WriteString(lightelementtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(lightelementtype.GongMarshallField(stage, "Availability"))
		pointersInitializesStatements.WriteString(lightelementtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linguisticnotepropertytype *LinguisticNotePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linguisticnotepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linguisticnotepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(linguisticnotepropertytype.GongMarshallField(stage, "LinguisticNote"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (linguisticnotetype *LinguisticNoteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(linguisticnotetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (markerbeaconpropertytype *MarkerBeaconPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(markerbeaconpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (navaidequipmentpropertytype *NavaidEquipmentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(navaidequipmentpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notepropertytype *NotePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(notepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(notepropertytype.GongMarshallField(stage, "Note"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (notetype *NoteType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(notetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(notetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(notetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(notetype.GongMarshallField(stage, "TranslatedNote"))
		pointersInitializesStatements.WriteString(notetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstacleareatimeslicepropertytype *ObstacleAreaTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstacleareatimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstacleareatimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicepropertytype.GongMarshallField(stage, "ObstacleAreaTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstacleareatimeslicetype *ObstacleAreaTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Reference_ownerOrganisation"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Reference_ownerRunway"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Reference_ownerAirport"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "SurfaceExtent"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Obstacle"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(obstacleareatimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (obstacleareatype *ObstacleAreaType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(obstacleareatype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(obstacleareatype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(obstacleareatype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(obstacleareatype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (organisationauthoritypropertytype *OrganisationAuthorityPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(organisationauthoritypropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (passengerservicepropertytype *PassengerServicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(passengerservicepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (pointtype *PointType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(pointtype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(pointtype.GongMarshallField(stage, "Annotation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (runwaydirectionpropertytype *RunwayDirectionPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(runwaydirectionpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (servicepropertytype *ServicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(servicepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (specialnavigationstationpropertytype *SpecialNavigationStationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(specialnavigationstationpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surfacepropertytype *SurfacePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "Surface"))
		pointersInitializesStatements.WriteString(surfacepropertytype.GongMarshallField(stage, "ElevatedSurface"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (surfacetype *SurfaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(surfacetype.GongMarshallField(stage, "SrsName"))
		pointersInitializesStatements.WriteString(surfacetype.GongMarshallField(stage, "Annotation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timesheetpropertytype *TimesheetPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timesheetpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timesheetpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(timesheetpropertytype.GongMarshallField(stage, "Timesheet"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (timesheettype *TimesheetType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(timesheettype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(timesheettype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(timesheettype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(timesheettype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(timesheettype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (unitpropertytype *UnitPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(unitpropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructurelightingstatuspropertytype *VerticalStructureLightingStatusPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructurelightingstatuspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructurelightingstatuspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(verticalstructurelightingstatuspropertytype.GongMarshallField(stage, "VerticalStructureLightingStatus"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructurelightingstatustype *VerticalStructureLightingStatusType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructurelightingstatustype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructurelightingstatustype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(verticalstructurelightingstatustype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(verticalstructurelightingstatustype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(verticalstructurelightingstatustype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(verticalstructurelightingstatustype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(verticalstructurelightingstatustype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructurepartpropertytype *VerticalStructurePartPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructurepartpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructurepartpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(verticalstructurepartpropertytype.GongMarshallField(stage, "VerticalStructurePart"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructureparttype *VerticalStructurePartType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "HorizontalProjection_surfaceExtent"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "HorizontalProjection_linearExtent"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "HorizontalProjection_location"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "Lighting"))
		pointersInitializesStatements.WriteString(verticalstructureparttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructurepropertytype *VerticalStructurePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(verticalstructurepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructuretimeslicepropertytype *VerticalStructureTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructuretimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructuretimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicepropertytype.GongMarshallField(stage, "VerticalStructureTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructuretimeslicetype *VerticalStructureTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Marker"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Part"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "HostedPassengerService"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "SupportedGroundLight"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "HostedNavaidEquipment"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "HostedSpecialNavStation"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "HostedUnit"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "HostedOrganisation"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "SupportedService"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "LightingAvailability"))
		pointersInitializesStatements.WriteString(verticalstructuretimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (verticalstructuretype *VerticalStructureType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(verticalstructuretype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(verticalstructuretype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(verticalstructuretype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(verticalstructuretype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
