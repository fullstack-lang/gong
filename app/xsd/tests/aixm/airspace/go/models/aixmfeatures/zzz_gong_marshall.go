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
	modelsPackageName := "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"
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
		for _, a_extension_13 := range __gong__sortStageSetInstances(stage.A_extension_13s, stage.A_extension_13_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_13Ident := a_extension_13.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_13{Name: %s}).Stage(stage)", a_extension_13Ident, __gong__toRawStringLiteral(a_extension_13.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_13Ident, __gong__toRawStringLiteral(a_extension_13.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_13Ident, a_extension_13.Owns))
		}
	}
	if stage != nil {
		for _, a_extension_14 := range __gong__sortStageSetInstances(stage.A_extension_14s, stage.A_extension_14_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			a_extension_14Ident := a_extension_14.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.A_extension_14{Name: %s}).Stage(stage)", a_extension_14Ident, __gong__toRawStringLiteral(a_extension_14.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", a_extension_14Ident, __gong__toRawStringLiteral(a_extension_14.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", a_extension_14Ident, a_extension_14.Owns))
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
		for _, aircraftcharacteristicpropertytype := range __gong__sortStageSetInstances(stage.AircraftCharacteristicPropertyTypes, stage.AircraftCharacteristicPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aircraftcharacteristicpropertytypeIdent := aircraftcharacteristicpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftCharacteristicPropertyType{Name: %s}).Stage(stage)", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", aircraftcharacteristicpropertytypeIdent, __gong__toRawStringLiteral(aircraftcharacteristicpropertytype.NilReason)))
			if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := aircraftcharacteristicpropertytype.AircraftCharacteristic.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AircraftCharacteristic = %s", aircraftcharacteristicpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, aircraftcharacteristictype := range __gong__sortStageSetInstances(stage.AircraftCharacteristicTypes, stage.AircraftCharacteristicType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			aircraftcharacteristictypeIdent := aircraftcharacteristictype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AircraftCharacteristicType{Name: %s}).Stage(stage)", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", aircraftcharacteristictypeIdent, __gong__toRawStringLiteral(aircraftcharacteristictype.Description)))
			for _, elem := range aircraftcharacteristictype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", aircraftcharacteristictypeIdent, aircraftcharacteristictypeIdent, targetIdent))
			}
			for _, elem := range aircraftcharacteristictype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", aircraftcharacteristictypeIdent, aircraftcharacteristictypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspaceactivationpropertytype := range __gong__sortStageSetInstances(stage.AirspaceActivationPropertyTypes, stage.AirspaceActivationPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspaceactivationpropertytypeIdent := airspaceactivationpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceActivationPropertyType{Name: %s}).Stage(stage)", airspaceactivationpropertytypeIdent, __gong__toRawStringLiteral(airspaceactivationpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspaceactivationpropertytypeIdent, __gong__toRawStringLiteral(airspaceactivationpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspaceactivationpropertytypeIdent, __gong__toRawStringLiteral(airspaceactivationpropertytype.NilReason)))
			if airspaceactivationpropertytype.AirspaceActivation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspaceactivationpropertytype.AirspaceActivation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceActivation = %s", airspaceactivationpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspaceactivationtype := range __gong__sortStageSetInstances(stage.AirspaceActivationTypes, stage.AirspaceActivationType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspaceactivationtypeIdent := airspaceactivationtype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceActivationType{Name: %s}).Stage(stage)", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspaceactivationtypeIdent, __gong__toRawStringLiteral(airspaceactivationtype.Description)))
			for _, elem := range airspaceactivationtype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Levels {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Levels = append(%s.Levels, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.User {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.User = append(%s.User, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Aircraft {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Aircraft = append(%s.Aircraft, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
			for _, elem := range airspaceactivationtype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspaceactivationtypeIdent, airspaceactivationtypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacegeometrycomponentpropertytype := range __gong__sortStageSetInstances(stage.AirspaceGeometryComponentPropertyTypes, stage.AirspaceGeometryComponentPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacegeometrycomponentpropertytypeIdent := airspacegeometrycomponentpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceGeometryComponentPropertyType{Name: %s}).Stage(stage)", airspacegeometrycomponentpropertytypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponentpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacegeometrycomponentpropertytypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponentpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacegeometrycomponentpropertytypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponentpropertytype.NilReason)))
			if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacegeometrycomponentpropertytype.AirspaceGeometryComponent.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceGeometryComponent = %s", airspacegeometrycomponentpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacegeometrycomponenttype := range __gong__sortStageSetInstances(stage.AirspaceGeometryComponentTypes, stage.AirspaceGeometryComponentType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacegeometrycomponenttypeIdent := airspacegeometrycomponenttype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceGeometryComponentType{Name: %s}).Stage(stage)", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacegeometrycomponenttypeIdent, __gong__toRawStringLiteral(airspacegeometrycomponenttype.Description)))
			for _, elem := range airspacegeometrycomponenttype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacegeometrycomponenttypeIdent, airspacegeometrycomponenttypeIdent, targetIdent))
			}
			if airspacegeometrycomponenttype.TheAirspaceVolume != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacegeometrycomponenttype.TheAirspaceVolume.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheAirspaceVolume = %s", airspacegeometrycomponenttypeIdent, targetIdent))
			}
			for _, elem := range airspacegeometrycomponenttype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacegeometrycomponenttypeIdent, airspacegeometrycomponenttypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacelayerclasspropertytype := range __gong__sortStageSetInstances(stage.AirspaceLayerClassPropertyTypes, stage.AirspaceLayerClassPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacelayerclasspropertytypeIdent := airspacelayerclasspropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerClassPropertyType{Name: %s}).Stage(stage)", airspacelayerclasspropertytypeIdent, __gong__toRawStringLiteral(airspacelayerclasspropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayerclasspropertytypeIdent, __gong__toRawStringLiteral(airspacelayerclasspropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacelayerclasspropertytypeIdent, __gong__toRawStringLiteral(airspacelayerclasspropertytype.NilReason)))
			if airspacelayerclasspropertytype.AirspaceLayerClass != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacelayerclasspropertytype.AirspaceLayerClass.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceLayerClass = %s", airspacelayerclasspropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacelayerclasstype := range __gong__sortStageSetInstances(stage.AirspaceLayerClassTypes, stage.AirspaceLayerClassType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacelayerclasstypeIdent := airspacelayerclasstype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerClassType{Name: %s}).Stage(stage)", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacelayerclasstypeIdent, __gong__toRawStringLiteral(airspacelayerclasstype.Description)))
			for _, elem := range airspacelayerclasstype.TimeInterval {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeInterval = append(%s.TimeInterval, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.SpecialDateAuthority {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.SpecialDateAuthority = append(%s.SpecialDateAuthority, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.AssociatedLevels {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AssociatedLevels = append(%s.AssociatedLevels, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
			for _, elem := range airspacelayerclasstype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacelayerclasstypeIdent, airspacelayerclasstypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacelayerpropertytype := range __gong__sortStageSetInstances(stage.AirspaceLayerPropertyTypes, stage.AirspaceLayerPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacelayerpropertytypeIdent := airspacelayerpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerPropertyType{Name: %s}).Stage(stage)", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacelayerpropertytypeIdent, __gong__toRawStringLiteral(airspacelayerpropertytype.NilReason)))
			if airspacelayerpropertytype.AirspaceLayer != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacelayerpropertytype.AirspaceLayer.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceLayer = %s", airspacelayerpropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacelayertype := range __gong__sortStageSetInstances(stage.AirspaceLayerTypes, stage.AirspaceLayerType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacelayertypeIdent := airspacelayertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceLayerType{Name: %s}).Stage(stage)", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacelayertypeIdent, __gong__toRawStringLiteral(airspacelayertype.Description)))
			if airspacelayertype.DiscreteLevelSeries != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacelayertype.DiscreteLevelSeries.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.DiscreteLevelSeries = %s", airspacelayertypeIdent, targetIdent))
			}
			for _, elem := range airspacelayertype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacelayertypeIdent, airspacelayertypeIdent, targetIdent))
			}
			for _, elem := range airspacelayertype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacelayertypeIdent, airspacelayertypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacepropertytype := range __gong__sortStageSetInstances(stage.AirspacePropertyTypes, stage.AirspacePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacepropertytypeIdent := airspacepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspacePropertyType{Name: %s}).Stage(stage)", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airspacepropertytypeIdent, airspacepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", airspacepropertytypeIdent, __gong__toRawStringLiteral(airspacepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, airspacetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AirspaceTimeSlicePropertyTypes, stage.AirspaceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacetimeslicepropertytypeIdent := airspacetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceTimeSlicePropertyType{Name: %s}).Stage(stage)", airspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(airspacetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", airspacetimeslicepropertytypeIdent, airspacetimeslicepropertytype.Owns))
			if airspacetimeslicepropertytype.AirspaceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacetimeslicepropertytype.AirspaceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceTimeSlice = %s", airspacetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacetimeslicetype := range __gong__sortStageSetInstances(stage.AirspaceTimeSliceTypes, stage.AirspaceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacetimeslicetypeIdent := airspacetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceTimeSliceType{Name: %s}).Stage(stage)", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", airspacetimeslicetypeIdent, __gong__toRawStringLiteral(airspacetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", airspacetimeslicetypeIdent, airspacetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", airspacetimeslicetypeIdent, airspacetimeslicetype.CorrectionNumber))
			for _, elem := range airspacetimeslicetype.Class {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Class = append(%s.Class, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			if airspacetimeslicetype.ProtectedRoute != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacetimeslicetype.ProtectedRoute.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ProtectedRoute = %s", airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.GeometryComponent {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeometryComponent = append(%s.GeometryComponent, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.Activation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Activation = append(%s.Activation, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range airspacetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacetimeslicetypeIdent, airspacetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacetype := range __gong__sortStageSetInstances(stage.AirspaceTypes, stage.AirspaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacetypeIdent := airspacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceType{Name: %s}).Stage(stage)", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacetypeIdent, __gong__toRawStringLiteral(airspacetype.Description)))
			for _, elem := range airspacetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", airspacetypeIdent, airspacetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacevolumedependencypropertytype := range __gong__sortStageSetInstances(stage.AirspaceVolumeDependencyPropertyTypes, stage.AirspaceVolumeDependencyPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacevolumedependencypropertytypeIdent := airspacevolumedependencypropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumeDependencyPropertyType{Name: %s}).Stage(stage)", airspacevolumedependencypropertytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencypropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumedependencypropertytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencypropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacevolumedependencypropertytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencypropertytype.NilReason)))
			if airspacevolumedependencypropertytype.AirspaceVolumeDependency != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacevolumedependencypropertytype.AirspaceVolumeDependency.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceVolumeDependency = %s", airspacevolumedependencypropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacevolumedependencytype := range __gong__sortStageSetInstances(stage.AirspaceVolumeDependencyTypes, stage.AirspaceVolumeDependencyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacevolumedependencytypeIdent := airspacevolumedependencytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumeDependencyType{Name: %s}).Stage(stage)", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacevolumedependencytypeIdent, __gong__toRawStringLiteral(airspacevolumedependencytype.Description)))
			for _, elem := range airspacevolumedependencytype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacevolumedependencytypeIdent, airspacevolumedependencytypeIdent, targetIdent))
			}
			if airspacevolumedependencytype.TheAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacevolumedependencytype.TheAirspace.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TheAirspace = %s", airspacevolumedependencytypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumedependencytype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacevolumedependencytypeIdent, airspacevolumedependencytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacevolumepropertytype := range __gong__sortStageSetInstances(stage.AirspaceVolumePropertyTypes, stage.AirspaceVolumePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacevolumepropertytypeIdent := airspacevolumepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumePropertyType{Name: %s}).Stage(stage)", airspacevolumepropertytypeIdent, __gong__toRawStringLiteral(airspacevolumepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumepropertytypeIdent, __gong__toRawStringLiteral(airspacevolumepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", airspacevolumepropertytypeIdent, __gong__toRawStringLiteral(airspacevolumepropertytype.NilReason)))
			if airspacevolumepropertytype.AirspaceVolume != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacevolumepropertytype.AirspaceVolume.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AirspaceVolume = %s", airspacevolumepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, airspacevolumetype := range __gong__sortStageSetInstances(stage.AirspaceVolumeTypes, stage.AirspaceVolumeType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			airspacevolumetypeIdent := airspacevolumetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AirspaceVolumeType{Name: %s}).Stage(stage)", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", airspacevolumetypeIdent, __gong__toRawStringLiteral(airspacevolumetype.Description)))
			if airspacevolumetype.HorizontalProjection != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacevolumetype.HorizontalProjection.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.HorizontalProjection = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.Centreline != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacevolumetype.Centreline.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Centreline = %s", airspacevolumetypeIdent, targetIdent))
			}
			if airspacevolumetype.ContributorAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := airspacevolumetype.ContributorAirspace.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ContributorAirspace = %s", airspacevolumetypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", airspacevolumetypeIdent, airspacevolumetypeIdent, targetIdent))
			}
			for _, elem := range airspacevolumetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", airspacevolumetypeIdent, airspacevolumetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforairspacetimeslicepropertytype := range __gong__sortStageSetInstances(stage.AuthorityForAirspaceTimeSlicePropertyTypes, stage.AuthorityForAirspaceTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforairspacetimeslicepropertytypeIdent := authorityforairspacetimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAirspaceTimeSlicePropertyType{Name: %s}).Stage(stage)", authorityforairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforairspacetimeslicepropertytypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", authorityforairspacetimeslicepropertytypeIdent, authorityforairspacetimeslicepropertytype.Owns))
			if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AuthorityForAirspaceTimeSlice = %s", authorityforairspacetimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforairspacetimeslicetype := range __gong__sortStageSetInstances(stage.AuthorityForAirspaceTimeSliceTypes, stage.AuthorityForAirspaceTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforairspacetimeslicetypeIdent := authorityforairspacetimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAirspaceTimeSliceType{Name: %s}).Stage(stage)", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", authorityforairspacetimeslicetypeIdent, __gong__toRawStringLiteral(authorityforairspacetimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetype.CorrectionNumber))
			if authorityforairspacetimeslicetype.ResponsibleOrganisation != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforairspacetimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ResponsibleOrganisation = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			if authorityforairspacetimeslicetype.AssignedAirspace != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := authorityforairspacetimeslicetype.AssignedAirspace.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.AssignedAirspace = %s", authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range authorityforairspacetimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetypeIdent, targetIdent))
			}
			for _, elem := range authorityforairspacetimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", authorityforairspacetimeslicetypeIdent, authorityforairspacetimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, authorityforairspacetype := range __gong__sortStageSetInstances(stage.AuthorityForAirspaceTypes, stage.AuthorityForAirspaceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			authorityforairspacetypeIdent := authorityforairspacetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.AuthorityForAirspaceType{Name: %s}).Stage(stage)", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", authorityforairspacetypeIdent, __gong__toRawStringLiteral(authorityforairspacetype.Description)))
			for _, elem := range authorityforairspacetype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", authorityforairspacetypeIdent, authorityforairspacetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, curvepropertytype := range __gong__sortStageSetInstances(stage.CurvePropertyTypes, stage.CurvePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			curvepropertytypeIdent := curvepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.CurvePropertyType{Name: %s}).Stage(stage)", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", curvepropertytypeIdent, __gong__toRawStringLiteral(curvepropertytype.NilReason)))
			if curvepropertytype.Curve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvepropertytype.Curve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Curve = %s", curvepropertytypeIdent, targetIdent))
			}
			if curvepropertytype.ElevatedCurve != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := curvepropertytype.ElevatedCurve.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.ElevatedCurve = %s", curvepropertytypeIdent, targetIdent))
			}
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
		for _, geobordertimeslicepropertytype := range __gong__sortStageSetInstances(stage.GeoBorderTimeSlicePropertyTypes, stage.GeoBorderTimeSlicePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geobordertimeslicepropertytypeIdent := geobordertimeslicepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GeoBorderTimeSlicePropertyType{Name: %s}).Stage(stage)", geobordertimeslicepropertytypeIdent, __gong__toRawStringLiteral(geobordertimeslicepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geobordertimeslicepropertytypeIdent, __gong__toRawStringLiteral(geobordertimeslicepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", geobordertimeslicepropertytypeIdent, geobordertimeslicepropertytype.Owns))
			if geobordertimeslicepropertytype.GeoBorderTimeSlice != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geobordertimeslicepropertytype.GeoBorderTimeSlice.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.GeoBorderTimeSlice = %s", geobordertimeslicepropertytypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, geobordertimeslicetype := range __gong__sortStageSetInstances(stage.GeoBorderTimeSliceTypes, stage.GeoBorderTimeSliceType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geobordertimeslicetypeIdent := geobordertimeslicetype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GeoBorderTimeSliceType{Name: %s}).Stage(stage)", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Description)))
			values.WriteString(fmt.Sprintf("\n\t%s.Interpretation = %s", geobordertimeslicetypeIdent, __gong__toRawStringLiteral(geobordertimeslicetype.Interpretation)))
			values.WriteString(fmt.Sprintf("\n\t%s.SequenceNumber = %d", geobordertimeslicetypeIdent, geobordertimeslicetype.SequenceNumber))
			values.WriteString(fmt.Sprintf("\n\t%s.CorrectionNumber = %d", geobordertimeslicetypeIdent, geobordertimeslicetype.CorrectionNumber))
			if geobordertimeslicetype.Border != nil {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := geobordertimeslicetype.Border.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Border = %s", geobordertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range geobordertimeslicetype.Annotation {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Annotation = append(%s.Annotation, %s)", geobordertimeslicetypeIdent, geobordertimeslicetypeIdent, targetIdent))
			}
			for _, elem := range geobordertimeslicetype.Extension {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.Extension = append(%s.Extension, %s)", geobordertimeslicetypeIdent, geobordertimeslicetypeIdent, targetIdent))
			}
		}
	}
	if stage != nil {
		for _, geobordertype := range __gong__sortStageSetInstances(stage.GeoBorderTypes, stage.GeoBorderType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			geobordertypeIdent := geobordertype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.GeoBorderType{Name: %s}).Stage(stage)", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Id = %s", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Id)))
			values.WriteString(fmt.Sprintf("\n\t%s.Description = %s", geobordertypeIdent, __gong__toRawStringLiteral(geobordertype.Description)))
			for _, elem := range geobordertype.TimeSlice {
				if lastStagePtr != "Stage" {
					if pointers.Len() > 0 {
						pointers.WriteString("\n")
					}
					lastStagePtr = "Stage"
				}
				targetIdent := elem.GongGetIdentifier(stage)
				pointers.WriteString(fmt.Sprintf("\n\t%s.TimeSlice = append(%s.TimeSlice, %s)", geobordertypeIdent, geobordertypeIdent, targetIdent))
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
		for _, routepropertytype := range __gong__sortStageSetInstances(stage.RoutePropertyTypes, stage.RoutePropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			routepropertytypeIdent := routepropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.RoutePropertyType{Name: %s}).Stage(stage)", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", routepropertytypeIdent, routepropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", routepropertytypeIdent, __gong__toRawStringLiteral(routepropertytype.Actuate)))
		}
	}
	if stage != nil {
		for _, standardlevelcolumnpropertytype := range __gong__sortStageSetInstances(stage.StandardLevelColumnPropertyTypes, stage.StandardLevelColumnPropertyType_stagedOrder) {
			if lastStageDecl != "Stage" {
				if declarations.Len() > 0 {
					declarations.WriteString("\n")
				}
				lastStageDecl = "Stage"
			}
			standardlevelcolumnpropertytypeIdent := standardlevelcolumnpropertytype.GongGetIdentifier(stage)
			declarations.WriteString(fmt.Sprintf("\n\t%s := (&aixmfeatures.StandardLevelColumnPropertyType{Name: %s}).Stage(stage)", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Name)))
			if lastStageVal != "Stage" {
				if values.Len() > 0 {
					values.WriteString("\n")
				}
				lastStageVal = "Stage"
			}
			values.WriteString(fmt.Sprintf("\n\t%s.Name = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Name)))
			values.WriteString(fmt.Sprintf("\n\t%s.Owns = %t", standardlevelcolumnpropertytypeIdent, standardlevelcolumnpropertytype.Owns))
			values.WriteString(fmt.Sprintf("\n\t%s.NilReason = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.NilReason)))
			values.WriteString(fmt.Sprintf("\n\t%s.Type_XSD = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Type_XSD)))
			values.WriteString(fmt.Sprintf("\n\t%s.Href = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Href)))
			values.WriteString(fmt.Sprintf("\n\t%s.Role = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Role)))
			values.WriteString(fmt.Sprintf("\n\t%s.Arcrole = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Arcrole)))
			values.WriteString(fmt.Sprintf("\n\t%s.Title = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Title)))
			values.WriteString(fmt.Sprintf("\n\t%s.Show = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Show)))
			values.WriteString(fmt.Sprintf("\n\t%s.Actuate = %s", standardlevelcolumnpropertytypeIdent, __gong__toRawStringLiteral(standardlevelcolumnpropertytype.Actuate)))
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
go run github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures/probe/cmd/probe@%s stage.go
*/
package %s

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures"%s
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

	if modelsPackageName != "" && modelsPackageName != "github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures" {
		res = strings.ReplaceAll(res, "\"github.com/fullstack-lang/gong/app/xsd/tests/aixm/airspace/go/models/aixmfeatures\"", fmt.Sprintf("%q", modelsPackageName))
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

func (a_extension_13 *A_extension_13) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_13.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_13.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_13.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_13", fieldName)
	}
	return
}

func (a_extension_14 *A_extension_14) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := a_extension_14.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", a_extension_14.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", a_extension_14.Owns)

	default:
		log.Panicf("Unknown field %s for Gongstruct A_extension_14", fieldName)
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

func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aircraftcharacteristicpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aircraftcharacteristicpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", aircraftcharacteristicpropertytype.NilReason)

	case "AircraftCharacteristic":
		if aircraftcharacteristicpropertytype.AircraftCharacteristic != nil {
			res = __gong__marshallPointer(ident, "AircraftCharacteristic", aircraftcharacteristicpropertytype.AircraftCharacteristic.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AircraftCharacteristic", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AircraftCharacteristicPropertyType", fieldName)
	}
	return
}

func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := aircraftcharacteristictype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", aircraftcharacteristictype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", aircraftcharacteristictype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", aircraftcharacteristictype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range aircraftcharacteristictype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_12 := range aircraftcharacteristictype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_12.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AircraftCharacteristicType", fieldName)
	}
	return
}

func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspaceactivationpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspaceactivationpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspaceactivationpropertytype.NilReason)

	case "AirspaceActivation":
		if airspaceactivationpropertytype.AirspaceActivation != nil {
			res = __gong__marshallPointer(ident, "AirspaceActivation", airspaceactivationpropertytype.AirspaceActivation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceActivation", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceActivationPropertyType", fieldName)
	}
	return
}

func (airspaceactivationtype *AirspaceActivationType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspaceactivationtype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspaceactivationtype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspaceactivationtype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspaceactivationtype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range airspaceactivationtype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspaceactivationtype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range airspaceactivationtype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Levels":
		var sb strings.Builder
		for _, _airspacelayerpropertytype := range airspaceactivationtype.Levels {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Levels", _airspacelayerpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "User":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range airspaceactivationtype.User {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "User", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Aircraft":
		var sb strings.Builder
		for _, _aircraftcharacteristicpropertytype := range airspaceactivationtype.Aircraft {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Aircraft", _aircraftcharacteristicpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_1 := range airspaceactivationtype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_1.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceActivationType", fieldName)
	}
	return
}

func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacegeometrycomponentpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacegeometrycomponentpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacegeometrycomponentpropertytype.NilReason)

	case "AirspaceGeometryComponent":
		if airspacegeometrycomponentpropertytype.AirspaceGeometryComponent != nil {
			res = __gong__marshallPointer(ident, "AirspaceGeometryComponent", airspacegeometrycomponentpropertytype.AirspaceGeometryComponent.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceGeometryComponent", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceGeometryComponentPropertyType", fieldName)
	}
	return
}

func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacegeometrycomponenttype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacegeometrycomponenttype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacegeometrycomponenttype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacegeometrycomponenttype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacegeometrycomponenttype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheAirspaceVolume":
		if airspacegeometrycomponenttype.TheAirspaceVolume != nil {
			res = __gong__marshallPointer(ident, "TheAirspaceVolume", airspacegeometrycomponenttype.TheAirspaceVolume.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheAirspaceVolume", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_2 := range airspacegeometrycomponenttype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_2.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceGeometryComponentType", fieldName)
	}
	return
}

func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacelayerclasspropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacelayerclasspropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacelayerclasspropertytype.NilReason)

	case "AirspaceLayerClass":
		if airspacelayerclasspropertytype.AirspaceLayerClass != nil {
			res = __gong__marshallPointer(ident, "AirspaceLayerClass", airspacelayerclasspropertytype.AirspaceLayerClass.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceLayerClass", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceLayerClassPropertyType", fieldName)
	}
	return
}

func (airspacelayerclasstype *AirspaceLayerClassType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacelayerclasstype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacelayerclasstype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacelayerclasstype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacelayerclasstype.Description)

	case "TimeInterval":
		var sb strings.Builder
		for _, _timesheetpropertytype := range airspacelayerclasstype.TimeInterval {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeInterval", _timesheetpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacelayerclasstype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "SpecialDateAuthority":
		var sb strings.Builder
		for _, _organisationauthoritypropertytype := range airspacelayerclasstype.SpecialDateAuthority {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "SpecialDateAuthority", _organisationauthoritypropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "AssociatedLevels":
		var sb strings.Builder
		for _, _airspacelayerpropertytype := range airspacelayerclasstype.AssociatedLevels {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "AssociatedLevels", _airspacelayerpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_3 := range airspacelayerclasstype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_3.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceLayerClassType", fieldName)
	}
	return
}

func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacelayerpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacelayerpropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacelayerpropertytype.NilReason)

	case "AirspaceLayer":
		if airspacelayerpropertytype.AirspaceLayer != nil {
			res = __gong__marshallPointer(ident, "AirspaceLayer", airspacelayerpropertytype.AirspaceLayer.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceLayer", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceLayerPropertyType", fieldName)
	}
	return
}

func (airspacelayertype *AirspaceLayerType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacelayertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacelayertype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacelayertype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacelayertype.Description)

	case "DiscreteLevelSeries":
		if airspacelayertype.DiscreteLevelSeries != nil {
			res = __gong__marshallPointer(ident, "DiscreteLevelSeries", airspacelayertype.DiscreteLevelSeries.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "DiscreteLevelSeries", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacelayertype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_13 := range airspacelayertype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_13.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceLayerType", fieldName)
	}
	return
}

func (airspacepropertytype *AirspacePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airspacepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", airspacepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", airspacepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", airspacepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", airspacepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", airspacepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", airspacepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", airspacepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct AirspacePropertyType", fieldName)
	}
	return
}

func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", airspacetimeslicepropertytype.Owns)

	case "AirspaceTimeSlice":
		if airspacetimeslicepropertytype.AirspaceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AirspaceTimeSlice", airspacetimeslicepropertytype.AirspaceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceTimeSlicePropertyType", fieldName)
	}
	return
}

func (airspacetimeslicetype *AirspaceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", airspacetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", airspacetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", airspacetimeslicetype.CorrectionNumber)

	case "Class":
		var sb strings.Builder
		for _, _airspacelayerclasspropertytype := range airspacetimeslicetype.Class {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Class", _airspacelayerclasspropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "ProtectedRoute":
		if airspacetimeslicetype.ProtectedRoute != nil {
			res = __gong__marshallPointer(ident, "ProtectedRoute", airspacetimeslicetype.ProtectedRoute.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ProtectedRoute", "nil")
		}
	case "GeometryComponent":
		var sb strings.Builder
		for _, _airspacegeometrycomponentpropertytype := range airspacetimeslicetype.GeometryComponent {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "GeometryComponent", _airspacegeometrycomponentpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Activation":
		var sb strings.Builder
		for _, _airspaceactivationpropertytype := range airspacetimeslicetype.Activation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Activation", _airspaceactivationpropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension := range airspacetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceTimeSliceType", fieldName)
	}
	return
}

func (airspacetype *AirspaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _airspacetimeslicepropertytype := range airspacetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _airspacetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceType", fieldName)
	}
	return
}

func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacevolumedependencypropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacevolumedependencypropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacevolumedependencypropertytype.NilReason)

	case "AirspaceVolumeDependency":
		if airspacevolumedependencypropertytype.AirspaceVolumeDependency != nil {
			res = __gong__marshallPointer(ident, "AirspaceVolumeDependency", airspacevolumedependencypropertytype.AirspaceVolumeDependency.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceVolumeDependency", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceVolumeDependencyPropertyType", fieldName)
	}
	return
}

func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacevolumedependencytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacevolumedependencytype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacevolumedependencytype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacevolumedependencytype.Description)

	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacevolumedependencytype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "TheAirspace":
		if airspacevolumedependencytype.TheAirspace != nil {
			res = __gong__marshallPointer(ident, "TheAirspace", airspacevolumedependencytype.TheAirspace.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "TheAirspace", "nil")
		}
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_5 := range airspacevolumedependencytype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_5.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceVolumeDependencyType", fieldName)
	}
	return
}

func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacevolumepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacevolumepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", airspacevolumepropertytype.NilReason)

	case "AirspaceVolume":
		if airspacevolumepropertytype.AirspaceVolume != nil {
			res = __gong__marshallPointer(ident, "AirspaceVolume", airspacevolumepropertytype.AirspaceVolume.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AirspaceVolume", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceVolumePropertyType", fieldName)
	}
	return
}

func (airspacevolumetype *AirspaceVolumeType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := airspacevolumetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", airspacevolumetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", airspacevolumetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", airspacevolumetype.Description)

	case "HorizontalProjection":
		if airspacevolumetype.HorizontalProjection != nil {
			res = __gong__marshallPointer(ident, "HorizontalProjection", airspacevolumetype.HorizontalProjection.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "HorizontalProjection", "nil")
		}
	case "Centreline":
		if airspacevolumetype.Centreline != nil {
			res = __gong__marshallPointer(ident, "Centreline", airspacevolumetype.Centreline.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Centreline", "nil")
		}
	case "ContributorAirspace":
		if airspacevolumetype.ContributorAirspace != nil {
			res = __gong__marshallPointer(ident, "ContributorAirspace", airspacevolumetype.ContributorAirspace.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ContributorAirspace", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range airspacevolumetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_4 := range airspacevolumetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_4.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AirspaceVolumeType", fieldName)
	}
	return
}

func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforairspacetimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforairspacetimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", authorityforairspacetimeslicepropertytype.Owns)

	case "AuthorityForAirspaceTimeSlice":
		if authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice != nil {
			res = __gong__marshallPointer(ident, "AuthorityForAirspaceTimeSlice", authorityforairspacetimeslicepropertytype.AuthorityForAirspaceTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AuthorityForAirspaceTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForAirspaceTimeSlicePropertyType", fieldName)
	}
	return
}

func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforairspacetimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforairspacetimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", authorityforairspacetimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", authorityforairspacetimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", authorityforairspacetimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", authorityforairspacetimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", authorityforairspacetimeslicetype.CorrectionNumber)

	case "ResponsibleOrganisation":
		if authorityforairspacetimeslicetype.ResponsibleOrganisation != nil {
			res = __gong__marshallPointer(ident, "ResponsibleOrganisation", authorityforairspacetimeslicetype.ResponsibleOrganisation.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ResponsibleOrganisation", "nil")
		}
	case "AssignedAirspace":
		if authorityforairspacetimeslicetype.AssignedAirspace != nil {
			res = __gong__marshallPointer(ident, "AssignedAirspace", authorityforairspacetimeslicetype.AssignedAirspace.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "AssignedAirspace", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range authorityforairspacetimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_6 := range authorityforairspacetimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_6.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForAirspaceTimeSliceType", fieldName)
	}
	return
}

func (authorityforairspacetype *AuthorityForAirspaceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := authorityforairspacetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", authorityforairspacetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", authorityforairspacetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", authorityforairspacetype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _authorityforairspacetimeslicepropertytype := range authorityforairspacetype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _authorityforairspacetimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct AuthorityForAirspaceType", fieldName)
	}
	return
}

func (curvepropertytype *CurvePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := curvepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", curvepropertytype.Name)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", curvepropertytype.NilReason)

	case "Curve":
		if curvepropertytype.Curve != nil {
			res = __gong__marshallPointer(ident, "Curve", curvepropertytype.Curve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Curve", "nil")
		}
	case "ElevatedCurve":
		if curvepropertytype.ElevatedCurve != nil {
			res = __gong__marshallPointer(ident, "ElevatedCurve", curvepropertytype.ElevatedCurve.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "ElevatedCurve", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct CurvePropertyType", fieldName)
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
		for _, _a_extension_8 := range elevatedcurvetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_8.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedCurveType", fieldName)
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
		for _, _a_extension_9 := range elevatedsurfacetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_9.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct ElevatedSurfaceType", fieldName)
	}
	return
}

func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := geobordertimeslicepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", geobordertimeslicepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", geobordertimeslicepropertytype.Owns)

	case "GeoBorderTimeSlice":
		if geobordertimeslicepropertytype.GeoBorderTimeSlice != nil {
			res = __gong__marshallPointer(ident, "GeoBorderTimeSlice", geobordertimeslicepropertytype.GeoBorderTimeSlice.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "GeoBorderTimeSlice", "nil")
		}
	default:
		log.Panicf("Unknown field %s for Gongstruct GeoBorderTimeSlicePropertyType", fieldName)
	}
	return
}

func (geobordertimeslicetype *GeoBorderTimeSliceType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := geobordertimeslicetype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", geobordertimeslicetype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", geobordertimeslicetype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", geobordertimeslicetype.Description)
	case "Interpretation":
		res = __gong__marshallString(ident, "Interpretation", geobordertimeslicetype.Interpretation)
	case "SequenceNumber":
		res = __gong__marshallInt(ident, "SequenceNumber", geobordertimeslicetype.SequenceNumber)
	case "CorrectionNumber":
		res = __gong__marshallInt(ident, "CorrectionNumber", geobordertimeslicetype.CorrectionNumber)

	case "Border":
		if geobordertimeslicetype.Border != nil {
			res = __gong__marshallPointer(ident, "Border", geobordertimeslicetype.Border.GongGetIdentifier(stage))
		} else {
			res = __gong__marshallPointer(ident, "Border", "nil")
		}
	case "Annotation":
		var sb strings.Builder
		for _, _notepropertytype := range geobordertimeslicetype.Annotation {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Annotation", _notepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	case "Extension":
		var sb strings.Builder
		for _, _a_extension_7 := range geobordertimeslicetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_7.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GeoBorderTimeSliceType", fieldName)
	}
	return
}

func (geobordertype *GeoBorderType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := geobordertype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", geobordertype.Name)
	case "Id":
		res = __gong__marshallString(ident, "Id", geobordertype.Id)
	case "Description":
		res = __gong__marshallString(ident, "Description", geobordertype.Description)

	case "TimeSlice":
		var sb strings.Builder
		for _, _geobordertimeslicepropertytype := range geobordertype.TimeSlice {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "TimeSlice", _geobordertimeslicepropertytype.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct GeoBorderType", fieldName)
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
		for _, _a_extension_10 := range linguisticnotetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_10.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct LinguisticNoteType", fieldName)
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
		for _, _a_extension_11 := range notetype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_11.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct NoteType", fieldName)
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

func (routepropertytype *RoutePropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := routepropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", routepropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", routepropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", routepropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", routepropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", routepropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", routepropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", routepropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", routepropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", routepropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", routepropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct RoutePropertyType", fieldName)
	}
	return
}

func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallField(stage *Stage, fieldName string) (res string) {
	ident := standardlevelcolumnpropertytype.GongGetIdentifier(stage)
	_ = ident
	switch fieldName {
	case "Name":
		res = __gong__marshallString(ident, "Name", standardlevelcolumnpropertytype.Name)
	case "Owns":
		res = __gong__marshallBool(ident, "Owns", standardlevelcolumnpropertytype.Owns)
	case "NilReason":
		res = __gong__marshallString(ident, "NilReason", standardlevelcolumnpropertytype.NilReason)
	case "Type_XSD":
		res = __gong__marshallString(ident, "Type_XSD", standardlevelcolumnpropertytype.Type_XSD)
	case "Href":
		res = __gong__marshallString(ident, "Href", standardlevelcolumnpropertytype.Href)
	case "Role":
		res = __gong__marshallString(ident, "Role", standardlevelcolumnpropertytype.Role)
	case "Arcrole":
		res = __gong__marshallString(ident, "Arcrole", standardlevelcolumnpropertytype.Arcrole)
	case "Title":
		res = __gong__marshallString(ident, "Title", standardlevelcolumnpropertytype.Title)
	case "Show":
		res = __gong__marshallString(ident, "Show", standardlevelcolumnpropertytype.Show)
	case "Actuate":
		res = __gong__marshallString(ident, "Actuate", standardlevelcolumnpropertytype.Actuate)

	default:
		log.Panicf("Unknown field %s for Gongstruct StandardLevelColumnPropertyType", fieldName)
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
		for _, _a_extension_14 := range timesheettype.Extension {
			sb.WriteString(__gong__marshallSliceOfPointers(ident, "Extension", _a_extension_14.GongGetIdentifier(stage)))
		}
		res = sb.String()
	default:
		log.Panicf("Unknown field %s for Gongstruct TimesheetType", fieldName)
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
func (a_extension_13 *A_extension_13) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_13.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_13.GongMarshallField(stage, "Owns"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (a_extension_14 *A_extension_14) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(a_extension_14.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(a_extension_14.GongMarshallField(stage, "Owns"))
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
func (aircraftcharacteristicpropertytype *AircraftCharacteristicPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aircraftcharacteristicpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aircraftcharacteristicpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(aircraftcharacteristicpropertytype.GongMarshallField(stage, "AircraftCharacteristic"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (aircraftcharacteristictype *AircraftCharacteristicType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(aircraftcharacteristictype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspaceactivationpropertytype *AirspaceActivationPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspaceactivationpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspaceactivationpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(airspaceactivationpropertytype.GongMarshallField(stage, "AirspaceActivation"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspaceactivationtype *AirspaceActivationType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "Levels"))
		pointersInitializesStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "User"))
		pointersInitializesStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "Aircraft"))
		pointersInitializesStatements.WriteString(airspaceactivationtype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacegeometrycomponentpropertytype *AirspaceGeometryComponentPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacegeometrycomponentpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacegeometrycomponentpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(airspacegeometrycomponentpropertytype.GongMarshallField(stage, "AirspaceGeometryComponent"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacegeometrycomponenttype *AirspaceGeometryComponentType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacegeometrycomponenttype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacegeometrycomponenttype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacegeometrycomponenttype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacegeometrycomponenttype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacegeometrycomponenttype.GongMarshallField(stage, "TheAirspaceVolume"))
		pointersInitializesStatements.WriteString(airspacegeometrycomponenttype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacelayerclasspropertytype *AirspaceLayerClassPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacelayerclasspropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacelayerclasspropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(airspacelayerclasspropertytype.GongMarshallField(stage, "AirspaceLayerClass"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacelayerclasstype *AirspaceLayerClassType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "TimeInterval"))
		pointersInitializesStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "SpecialDateAuthority"))
		pointersInitializesStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "AssociatedLevels"))
		pointersInitializesStatements.WriteString(airspacelayerclasstype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacelayerpropertytype *AirspaceLayerPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacelayerpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacelayerpropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(airspacelayerpropertytype.GongMarshallField(stage, "AirspaceLayer"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacelayertype *AirspaceLayerType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacelayertype.GongMarshallField(stage, "DiscreteLevelSeries"))
		pointersInitializesStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacelayertype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacepropertytype *AirspacePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(airspacepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacetimeslicepropertytype *AirspaceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(airspacetimeslicepropertytype.GongMarshallField(stage, "AirspaceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacetimeslicetype *AirspaceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Class"))
		pointersInitializesStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "ProtectedRoute"))
		pointersInitializesStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "GeometryComponent"))
		pointersInitializesStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Activation"))
		pointersInitializesStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacetype *AirspaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacevolumedependencypropertytype *AirspaceVolumeDependencyPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacevolumedependencypropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacevolumedependencypropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(airspacevolumedependencypropertytype.GongMarshallField(stage, "AirspaceVolumeDependency"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacevolumedependencytype *AirspaceVolumeDependencyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacevolumedependencytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacevolumedependencytype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacevolumedependencytype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacevolumedependencytype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacevolumedependencytype.GongMarshallField(stage, "TheAirspace"))
		pointersInitializesStatements.WriteString(airspacevolumedependencytype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacevolumepropertytype *AirspaceVolumePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacevolumepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacevolumepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(airspacevolumepropertytype.GongMarshallField(stage, "AirspaceVolume"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (airspacevolumetype *AirspaceVolumeType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "HorizontalProjection"))
		pointersInitializesStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "Centreline"))
		pointersInitializesStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "ContributorAirspace"))
		pointersInitializesStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(airspacevolumetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforairspacetimeslicepropertytype *AuthorityForAirspaceTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforairspacetimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforairspacetimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(authorityforairspacetimeslicepropertytype.GongMarshallField(stage, "AuthorityForAirspaceTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforairspacetimeslicetype *AuthorityForAirspaceTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "ResponsibleOrganisation"))
		pointersInitializesStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "AssignedAirspace"))
		pointersInitializesStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(authorityforairspacetimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (authorityforairspacetype *AuthorityForAirspaceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(authorityforairspacetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(authorityforairspacetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(authorityforairspacetype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(authorityforairspacetype.GongMarshallField(stage, "TimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (curvepropertytype *CurvePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(curvepropertytype.GongMarshallField(stage, "NilReason"))
		pointersInitializesStatements.WriteString(curvepropertytype.GongMarshallField(stage, "Curve"))
		pointersInitializesStatements.WriteString(curvepropertytype.GongMarshallField(stage, "ElevatedCurve"))
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
func (geobordertimeslicepropertytype *GeoBorderTimeSlicePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(geobordertimeslicepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(geobordertimeslicepropertytype.GongMarshallField(stage, "Owns"))
		pointersInitializesStatements.WriteString(geobordertimeslicepropertytype.GongMarshallField(stage, "GeoBorderTimeSlice"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (geobordertimeslicetype *GeoBorderTimeSliceType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "Description"))
		initializerStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "Interpretation"))
		initializerStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "SequenceNumber"))
		initializerStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
		pointersInitializesStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "Border"))
		pointersInitializesStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "Annotation"))
		pointersInitializesStatements.WriteString(geobordertimeslicetype.GongMarshallField(stage, "Extension"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (geobordertype *GeoBorderType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(geobordertype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(geobordertype.GongMarshallField(stage, "Id"))
		initializerStatements.WriteString(geobordertype.GongMarshallField(stage, "Description"))
		pointersInitializesStatements.WriteString(geobordertype.GongMarshallField(stage, "TimeSlice"))
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
func (routepropertytype *RoutePropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(routepropertytype.GongMarshallField(stage, "Actuate"))
	}
	initRes = initializerStatements.String()
	ptrRes = pointersInitializesStatements.String()
	return
}
func (standardlevelcolumnpropertytype *StandardLevelColumnPropertyType) GongMarshallAllFields(stage *Stage) (initRes string, ptrRes string) {

	var initializerStatements strings.Builder
	var pointersInitializesStatements strings.Builder
	{ // Insertion point for basic fields value assignment
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Name"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Owns"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "NilReason"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Type_XSD"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Href"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Role"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Arcrole"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Title"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Show"))
		initializerStatements.WriteString(standardlevelcolumnpropertytype.GongMarshallField(stage, "Actuate"))
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
